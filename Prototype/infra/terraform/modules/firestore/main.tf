# Firestore native mode, database "(default)" so the free tier applies (12.1, 12.4).
# The location is fixed once the database exists (3.5).
resource "google_firestore_database" "this" {
  project                           = var.project_id
  name                              = "(default)"
  location_id                       = var.location
  type                              = "FIRESTORE_NATIVE"
  concurrency_mode                  = "OPTIMISTIC"
  point_in_time_recovery_enablement = "POINT_IN_TIME_RECOVERY_ENABLED"
  delete_protection_state           = var.delete_protection ? "DELETE_PROTECTION_ENABLED" : "DELETE_PROTECTION_DISABLED"
  deletion_policy                   = "ABANDON"
}

# Composite indexes from section 7.1.
# DECISION: the 7.1 rows "sessions updatedAt asc" and "runs code asc" are single-field
# indexes, which Firestore creates automatically; a composite index needs two or more
# fields, so they are not declared here.
locals {
  indexes = {
    sessions_run_opened = {
      collection = "sessions"
      fields     = [["runId", "ASCENDING"], ["openedAt", "DESCENDING"]]
    }
    sessions_status_lastevent = {
      collection = "sessions"
      fields     = [["status", "ASCENDING"], ["lastEventAt", "ASCENDING"]]
    }
    sessions_ip_opened = {
      collection = "sessions"
      fields     = [["ip", "ASCENDING"], ["openedAt", "DESCENDING"]]
    }
  }
}

resource "google_firestore_index" "this" {
  for_each = local.indexes

  project    = var.project_id
  database   = google_firestore_database.this.name
  collection = each.value.collection

  dynamic "fields" {
    for_each = each.value.fields
    content {
      field_path = fields.value[0]
      order      = fields.value[1]
    }
  }
}

# Backups (12.4): daily kept 7 days, weekly kept 14 weeks. No TTL policy.
resource "google_firestore_backup_schedule" "daily" {
  project   = var.project_id
  database  = google_firestore_database.this.name
  retention = "604800s" # 7 days, the maximum for a daily schedule

  daily_recurrence {}
}

resource "google_firestore_backup_schedule" "weekly" {
  project   = var.project_id
  database  = google_firestore_database.this.name
  retention = "8467200s" # 14 weeks

  weekly_recurrence {
    day = "SUNDAY"
  }
}

# 12.3: datastore.user for the API and sweeper runtime accounts. This role can only be
# granted at project level.
resource "google_project_iam_member" "user" {
  for_each = var.user_members

  project = var.project_id
  role    = "roles/datastore.user"
  member  = each.value
}
