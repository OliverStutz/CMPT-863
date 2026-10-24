variable "project_id" {
  description = "GCP project id (OPEN Q). Run the bootstrap once per project (dev, prod)."
  type        = string
}

variable "region" {
  type    = string
  default = "northamerica-northeast1"
}

variable "state_bucket_name" {
  description = "Terraform state bucket; empty means <project_id>-tfstate."
  type        = string
  default     = ""
}

variable "github_repository" {
  description = "GitHub repository allowed to use the pool, as owner/name."
  type        = string
}
