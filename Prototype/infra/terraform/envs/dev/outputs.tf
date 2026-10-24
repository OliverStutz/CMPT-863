output "service_url" {
  value = module.cloud_run.service_uri
}

output "content_bucket" {
  value = module.storage.content_bucket
}

output "tts_cache_bucket" {
  value = module.storage.tts_cache_bucket
}

output "image_base" {
  description = "Push images to <image_base>/api."
  value       = module.artifact_registry.image_base
}

output "admin_password_hash_secret" {
  value = module.secrets.admin_password_hash_secret
}

output "domain_dns_records" {
  value = module.cloud_run.domain_dns_records
}
