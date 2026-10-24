# Enables the APIs the study needs (architecture 12.1).
locals {
  # DECISION: iam.googleapis.com (service accounts) and billingbudgets.googleapis.com
  # (budget alert) are added to the list in 12.1, which does not name them but whose
  # modules need them.
  services = toset(concat([
    "run.googleapis.com",
    "firestore.googleapis.com",
    "storage.googleapis.com",
    "secretmanager.googleapis.com",
    "artifactregistry.googleapis.com",
    "texttospeech.googleapis.com",
    "cloudscheduler.googleapis.com",
    "monitoring.googleapis.com",
    "logging.googleapis.com",
    "iam.googleapis.com",
    "billingbudgets.googleapis.com",
  ], var.extra_services))
}

resource "google_project_service" "this" {
  for_each = local.services

  project            = var.project_id
  service            = each.value
  disable_on_destroy = false
}

# 12.3: ci-deploy calls Text-to-Speech, which needs serviceUsageConsumer.
resource "google_project_iam_member" "service_usage_consumer" {
  for_each = var.service_usage_consumers

  project = var.project_id
  role    = "roles/serviceusage.serviceUsageConsumer"
  member  = each.value
}
