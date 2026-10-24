variable "project_id" {
  type = string
}

variable "location" {
  description = "Firestore location; defaults to the region in the env. Fixed once created."
  type        = string
}

variable "delete_protection" {
  type    = bool
  default = true
}

variable "user_members" {
  description = "Map of static label to IAM member (static keys keep for_each plannable), granted roles/datastore.user."
  type        = map(string)
  default     = {}
}
