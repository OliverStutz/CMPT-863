output "uptime_check_id" {
  value = google_monitoring_uptime_check_config.healthz.uptime_check_id
}

output "notification_channels" {
  value = local.channels
}
