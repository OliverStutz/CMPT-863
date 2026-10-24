# Private content bucket, TTS cache bucket, 7-day log retention (12.1, 12.4).
locals {
  content_bucket = coalesce(var.content_bucket_name, "${var.project_id}-content")
  cache_bucket   = coalesce(var.tts_cache_bucket_name, "${var.project_id}-tts-cache")
}

# Content assets: hashed files, scene packs, resolved configs, manifests.
# Read by the API service (objectViewer), written by CI (objectAdmin).
# Uniform bucket-level access and enforced public access prevention are set
# explicitly; there is no allUsers or allAuthenticatedUsers binding, and the
# enforced prevention would reject one. Participants get assets only through the
# service's /a/* handler.
resource "google_storage_bucket" "content" {
  project                     = var.project_id
  name                        = local.content_bucket
  location                    = var.region
  storage_class               = "STANDARD"
  uniform_bucket_level_access = true
  public_access_prevention    = "enforced"
  force_destroy               = var.force_destroy
}

# Cache of generated TTS clips, used only by CI.
resource "google_storage_bucket" "tts_cache" {
  project                     = var.project_id
  name                        = local.cache_bucket
  location                    = var.region
  storage_class               = "STANDARD"
  uniform_bucket_level_access = true
  public_access_prevention    = "enforced"
  force_destroy               = var.force_destroy
}

# Guard against a public member slipping in through the variables.
resource "terraform_data" "no_public_members" {
  lifecycle {
    precondition {
      condition = length([
        for m in concat(values(var.content_viewers), values(var.object_admins)) : m
        if contains(["allUsers", "allAuthenticatedUsers"], m)
      ]) == 0
      error_message = "The content and TTS cache buckets must not be granted to allUsers or allAuthenticatedUsers."
    }
  }
}

resource "google_storage_bucket_iam_member" "content_viewer" {
  for_each = var.content_viewers

  bucket = google_storage_bucket.content.name
  role   = "roles/storage.objectViewer"
  member = each.value
}

resource "google_storage_bucket_iam_member" "content_admin" {
  for_each = var.object_admins

  bucket = google_storage_bucket.content.name
  role   = "roles/storage.objectAdmin"
  member = each.value
}

resource "google_storage_bucket_iam_member" "cache_admin" {
  for_each = var.object_admins

  bucket = google_storage_bucket.tts_cache.name
  role   = "roles/storage.objectAdmin"
  member = each.value
}

# DECISION: 12.1 lists a "log bucket with 7-day retention" under storage and 12.4 says
# request logs go "in a bucket with 7-day retention". The simplest reading is the
# project's Cloud Logging _Default bucket with retention cut to 7 days (no extra
# GCS bucket or log sink).
resource "google_logging_project_bucket_config" "default" {
  project        = var.project_id
  location       = "global"
  bucket_id      = "_Default"
  retention_days = var.log_retention_days
}
