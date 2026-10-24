variable "project_id" {
  type = string
}

variable "service_name" {
  description = "Cloud Run service name, for the 5xx and log filters."
  type        = string
}

variable "host" {
  description = "Host for the uptime check: the custom domain if set, else the run.app host."
  type        = string
}

variable "alert_email" {
  description = "Email for alerts and budget notices. Empty creates no channel."
  type        = string
  default     = ""
}

variable "billing_account" {
  description = "Billing account id (XXXXXX-XXXXXX-XXXXXX). Empty skips the budget, to be created by hand."
  type        = string
  default     = ""
}

variable "budget_amount" {
  description = "Monthly budget in whole currency units."
  type        = number
  default     = 25
}

variable "budget_currency" {
  description = "Must match the billing account currency."
  type        = string
  default     = "USD"
}

variable "error_5xx_threshold" {
  type    = number
  default = 5
}

variable "rejected_event_threshold" {
  type    = number
  default = 50
}

# DECISION: the document names a log-based metric for rejected events but not the log
# field. Default assumes the API writes a structured log line with
# jsonPayload.event_valid=false for each stored rejected event; the backend must match
# this, or this variable is changed.
variable "rejected_event_log_filter" {
  type    = string
  default = "jsonPayload.event_valid=false"
}
