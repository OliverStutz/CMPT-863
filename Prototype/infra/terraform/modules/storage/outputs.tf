output "content_bucket" {
  value = google_storage_bucket.content.name
}

output "tts_cache_bucket" {
  value = google_storage_bucket.tts_cache.name
}
