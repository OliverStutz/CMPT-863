# Copy these into the GitHub environment variables (see infra/README.md).
output "TF_STATE_BUCKET" {
  value = google_storage_bucket.tfstate.name
}

output "GCP_WIF_PROVIDER" {
  value = google_iam_workload_identity_pool_provider.github.name
}

output "GCP_DEPLOY_SA" {
  value = google_service_account.ci_deploy.email
}

output "GCP_TERRAFORM_PLAN_SA" {
  value = google_service_account.ci_terraform_plan.email
}

output "GCP_TERRAFORM_SA" {
  value = google_service_account.ci_terraform_apply.email
}
