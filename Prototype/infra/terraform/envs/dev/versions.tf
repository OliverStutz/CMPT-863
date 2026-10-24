terraform {
  required_version = ">= 1.15.0"
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = ">= 7.0, < 8.0"
    }
  }
}

provider "google" {
  project = var.project_id
  region  = var.region
  # The billing budgets API needs a quota project when called with user or
  # service account credentials.
  user_project_override = true
  billing_project       = var.project_id
}
