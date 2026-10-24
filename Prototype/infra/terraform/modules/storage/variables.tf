variable "project_id" {
  type = string
}

variable "region" {
  type = string
}

variable "content_bucket_name" {
  description = "Content bucket name; empty means <project_id>-content."
  type        = string
  default     = ""
}

variable "tts_cache_bucket_name" {
  description = "TTS cache bucket name; empty means <project_id>-tts-cache."
  type        = string
  default     = ""
}

variable "content_viewers" {
  description = "Map of static label to IAM member (static keys keep for_each plannable), with objectViewer on the content bucket (the API runtime account)."
  type        = map(string)
  default     = {}
}

variable "object_admins" {
  description = "Map of static label to IAM member (static keys keep for_each plannable), with objectAdmin on the content and TTS cache buckets (the CI deployer)."
  type        = map(string)
  default     = {}
}

variable "force_destroy" {
  description = "Allow terraform destroy to delete non-empty buckets."
  type        = bool
  default     = false
}

variable "log_retention_days" {
  type    = number
  default = 7
}
