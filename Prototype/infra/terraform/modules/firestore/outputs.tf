output "database" {
  value = google_firestore_database.this.name
}

output "location" {
  value = google_firestore_database.this.location_id
}
