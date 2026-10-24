output "service_name" {
  value = google_cloud_run_v2_service.api.name
}

output "service_uri" {
  description = "https://...run.app URL of the service."
  value       = google_cloud_run_v2_service.api.uri
}

output "api_runtime_member" {
  value = "serviceAccount:${google_service_account.api.email}"
}

output "sweeper_runtime_member" {
  value = "serviceAccount:${google_service_account.sweeper.email}"
}

output "scheduler_invoker_email" {
  value = google_service_account.scheduler.email
}

output "sweeper_job_name" {
  value = google_cloud_run_v2_job.sweeper.name
}

output "domain_dns_records" {
  description = "DNS records to create at the domain host, when a domain is set."
  value       = var.domain != "" ? google_cloud_run_domain_mapping.this[0].status[0].resource_records : []
}
