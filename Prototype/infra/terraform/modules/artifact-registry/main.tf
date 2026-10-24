# Docker repository for the API image, keeps the last 10 images (12.1).
resource "google_artifact_registry_repository" "this" {
  project       = var.project_id
  location      = var.region
  repository_id = var.repository_id
  format        = "DOCKER"
  description   = "Study API images"

  cleanup_policy_dry_run = false

  # DECISION: "keep last 10 images" is implemented as a KEEP policy for the 10 most
  # recent versions plus a DELETE policy for anything older than one day; the KEEP
  # policy wins, so only versions outside the newest 10 are removed.
  cleanup_policies {
    id     = "keep-last-10"
    action = "KEEP"
    most_recent_versions {
      keep_count = 10
    }
  }

  cleanup_policies {
    id     = "delete-older"
    action = "DELETE"
    condition {
      tag_state  = "ANY"
      older_than = "86400s"
    }
  }
}

# 12.3: ci-deploy pushes images (artifactregistry.writer on the repo only).
resource "google_artifact_registry_repository_iam_member" "writer" {
  for_each = var.writer_members

  project    = var.project_id
  location   = google_artifact_registry_repository.this.location
  repository = google_artifact_registry_repository.this.name
  role       = "roles/artifactregistry.writer"
  member     = each.value
}
