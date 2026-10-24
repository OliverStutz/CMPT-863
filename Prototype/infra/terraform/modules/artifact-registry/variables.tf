variable "project_id" {
  type = string
}

variable "region" {
  description = "Region of the repository."
  type        = string
}

variable "repository_id" {
  type    = string
  default = "study"
}

variable "writer_members" {
  description = "Map of static label to IAM member (static keys keep for_each plannable), allowed to push images (the CI deployer)."
  type        = map(string)
  default     = {}
}
