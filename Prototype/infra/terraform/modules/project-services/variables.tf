variable "project_id" {
  description = "GCP project id (OPEN Q, still to be named)."
  type        = string
}

variable "extra_services" {
  description = "Additional APIs to enable."
  type        = list(string)
  default     = []
}

variable "service_usage_consumers" {
  description = "Map of static label to IAM member (static keys keep for_each plannable), (e.g. serviceAccount:ci-deploy@...) that get roles/serviceusage.serviceUsageConsumer, needed to call Text-to-Speech from CI."
  type        = map(string)
  default     = {}
}
