variable "project_id" {
  type = string
}

variable "region" {
  type = string
}

variable "service_name" {
  type    = string
  default = "study-api"
}

variable "image_digest" {
  description = "Full image reference, ideally pinned by digest (<region>-docker.pkg.dev/<p>/study/api@sha256:...). Empty uses Google's public hello image so the first apply works before any image is pushed."
  type        = string
  default     = ""
}

variable "min_instances" {
  description = "1 during active recruitment, 0 otherwise (12.4)."
  type        = number
  default     = 0
}

variable "max_instances" {
  type    = number
  default = 3
}

variable "concurrency" {
  type    = number
  default = 80
}

variable "cpu" {
  type    = string
  default = "1"
}

variable "memory" {
  type    = string
  default = "512Mi"
}

variable "content_bucket" {
  description = "Name of the private content bucket (CONTENT_BUCKET)."
  type        = string
}

variable "admin_password_hash_secret" {
  description = "Resource name of the admin-password-hash secret (ADMIN_PASSWORD_HASH_SECRET)."
  type        = string
}

variable "admin_access_key_secret" {
  description = "Resource name of the admin-access-key secret (ADMIN_ACCESS_KEY_SECRET)."
  type        = string
}

variable "admin_session_key_secret" {
  description = "Resource name of the admin-session-key secret (ADMIN_SESSION_KEY_SECRET)."
  type        = string
}

variable "extra_env" {
  description = "Further plain environment variables for the API service."
  type        = map(string)
  default     = {}
}

variable "domain" {
  description = "Custom domain (OPEN Q). Empty means only the run.app URL."
  type        = string
  default     = ""
}

variable "deployer_member" {
  description = "IAM member of the CI deployer (serviceAccount:ci-deploy@...). Empty skips its bindings."
  type        = string
  default     = ""
}

variable "deletion_protection" {
  description = "Terraform-side deletion protection for the service and job."
  type        = bool
  default     = false
}
