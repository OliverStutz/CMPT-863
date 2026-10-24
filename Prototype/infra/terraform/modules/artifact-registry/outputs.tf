output "repository_id" {
  value = google_artifact_registry_repository.this.repository_id
}

output "image_base" {
  description = "Prefix for images in this repository, e.g. <image_base>/api:<tag>."
  value       = "${var.region}-docker.pkg.dev/${var.project_id}/${google_artifact_registry_repository.this.repository_id}"
}
