terraform {
  required_version = ">= 1.15.0"
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = ">= 7.0, < 8.0"
    }
  }
  # Local state on purpose: this config creates the state bucket the envs use.
  # Keep terraform.tfstate from this directory somewhere safe (it holds no secrets),
  # or migrate it into the new bucket afterwards with a gcs backend block and
  # `terraform init -migrate-state`.
}

provider "google" {
  project = var.project_id
  region  = var.region
}
