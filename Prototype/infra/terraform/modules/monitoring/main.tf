# Uptime check, alerts, notification channel and budget (12.1).

data "google_project" "this" {
  project_id = var.project_id
}

resource "google_monitoring_notification_channel" "email" {
  count = var.alert_email != "" ? 1 : 0

  project      = var.project_id
  display_name = "Study alerts (${var.alert_email})"
  type         = "email"
  labels = {
    email_address = var.alert_email
  }
}

locals {
  channels = [for c in google_monitoring_notification_channel.email : c.id]
}

########## Uptime check on /api/v1/healthz ##########

resource "google_monitoring_uptime_check_config" "healthz" {
  project      = var.project_id
  display_name = "study-api healthz"
  timeout      = "10s"
  period       = "300s"

  http_check {
    path         = "/api/v1/healthz"
    port         = 443
    use_ssl      = true
    validate_ssl = true
  }

  monitored_resource {
    type = "uptime_url"
    labels = {
      project_id = var.project_id
      host       = var.host
    }
  }
}

resource "google_monitoring_alert_policy" "uptime" {
  project      = var.project_id
  display_name = "study-api healthz failing"
  combiner     = "OR"

  conditions {
    display_name = "Uptime check failing"
    condition_threshold {
      filter          = "metric.type=\"monitoring.googleapis.com/uptime_check/check_passed\" AND resource.type=\"uptime_url\" AND metric.label.check_id=\"${google_monitoring_uptime_check_config.healthz.uptime_check_id}\""
      comparison      = "COMPARISON_GT"
      threshold_value = 1
      duration        = "600s"
      aggregations {
        alignment_period     = "300s"
        per_series_aligner   = "ALIGN_NEXT_OLDER"
        cross_series_reducer = "REDUCE_COUNT_FALSE"
        group_by_fields      = ["resource.label.host"]
      }
    }
  }

  notification_channels = local.channels
}

########## 5xx rate ##########

resource "google_monitoring_alert_policy" "errors_5xx" {
  project      = var.project_id
  display_name = "study-api 5xx responses"
  combiner     = "OR"

  conditions {
    display_name = "5xx responses above ${var.error_5xx_threshold} in 5 minutes"
    condition_threshold {
      filter          = "metric.type=\"run.googleapis.com/request_count\" AND resource.type=\"cloud_run_revision\" AND resource.label.service_name=\"${var.service_name}\" AND metric.label.response_code_class=\"5xx\""
      comparison      = "COMPARISON_GT"
      threshold_value = var.error_5xx_threshold
      duration        = "0s"
      aggregations {
        alignment_period     = "300s"
        per_series_aligner   = "ALIGN_SUM"
        cross_series_reducer = "REDUCE_SUM"
      }
    }
  }

  notification_channels = local.channels
}

########## Rejected-event spikes (log-based metric) ##########

resource "google_logging_metric" "rejected_events" {
  project     = var.project_id
  name        = "study_rejected_events"
  description = "Participant events stored with valid=false (section 6)."
  filter      = "resource.type=\"cloud_run_revision\" AND resource.labels.service_name=\"${var.service_name}\" AND ${var.rejected_event_log_filter}"

  metric_descriptor {
    metric_kind = "DELTA"
    value_type  = "INT64"
    unit        = "1"
  }
}

resource "google_monitoring_alert_policy" "rejected_events" {
  project      = var.project_id
  display_name = "study-api rejected-event spike"
  combiner     = "OR"

  conditions {
    display_name = "Rejected events above ${var.rejected_event_threshold} in 5 minutes"
    condition_threshold {
      filter          = "metric.type=\"logging.googleapis.com/user/${google_logging_metric.rejected_events.name}\" AND resource.type=\"cloud_run_revision\""
      comparison      = "COMPARISON_GT"
      threshold_value = var.rejected_event_threshold
      duration        = "0s"
      aggregations {
        alignment_period     = "300s"
        per_series_aligner   = "ALIGN_SUM"
        cross_series_reducer = "REDUCE_SUM"
      }
    }
  }

  notification_channels = local.channels
}

########## Budget (50/90/100 %) ##########
# Needs billing rights on the billing account for the applying identity
# (roles/billing.costsManager). Leave billing_account empty to create it by hand.
resource "google_billing_budget" "this" {
  count = var.billing_account != "" ? 1 : 0

  billing_account = var.billing_account
  display_name    = "study ${var.project_id}"

  budget_filter {
    projects = ["projects/${data.google_project.this.number}"]
  }

  amount {
    specified_amount {
      currency_code = var.budget_currency
      units         = tostring(var.budget_amount)
    }
  }

  dynamic "threshold_rules" {
    for_each = [0.5, 0.9, 1.0]
    content {
      threshold_percent = threshold_rules.value
    }
  }

  all_updates_rule {
    monitoring_notification_channels = local.channels
    disable_default_iam_recipients   = false
  }
}
