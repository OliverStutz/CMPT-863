# Study environment: one Cloud Run service, Firestore, private content bucket,
# Secret Manager, Cloud Scheduler-driven sweeper, monitoring (architecture 12).
# envs/dev and envs/prod are identical apart from backend prefix and tfvars.

locals {
  firestore_location = var.firestore_location != "" ? var.firestore_location : var.region
  deployer_member    = var.ci_deploy_sa_email != "" ? "serviceAccount:${var.ci_deploy_sa_email}" : ""
  deployer_map       = var.ci_deploy_sa_email != "" ? { ci_deploy = local.deployer_member } : {}
}

module "project_services" {
  source = "../../modules/project-services"

  project_id              = var.project_id
  service_usage_consumers = local.deployer_map
}

module "artifact_registry" {
  source = "../../modules/artifact-registry"

  project_id     = var.project_id
  region         = var.region
  writer_members = local.deployer_map

  depends_on = [module.project_services]
}

module "storage" {
  source = "../../modules/storage"

  project_id          = var.project_id
  region              = var.region
  content_bucket_name = local.content_bucket_name
  content_viewers     = { api_runtime = module.cloud_run.api_runtime_member }
  object_admins       = local.deployer_map
  force_destroy       = var.force_destroy_buckets

  depends_on = [module.project_services]
}

module "firestore" {
  source = "../../modules/firestore"

  project_id        = var.project_id
  location          = local.firestore_location
  delete_protection = var.firestore_delete_protection
  user_members = {
    api_runtime     = module.cloud_run.api_runtime_member
    sweeper_runtime = module.cloud_run.sweeper_runtime_member
  }

  depends_on = [module.project_services]
}

module "secrets" {
  source = "../../modules/secrets"

  project_id       = var.project_id
  accessor_members = { api_runtime = module.cloud_run.api_runtime_member }

  depends_on = [module.project_services]
}

module "cloud_run" {
  source = "../../modules/cloud-run"

  project_id                 = var.project_id
  region                     = var.region
  image_digest               = var.image_digest
  min_instances              = var.min_instances
  max_instances              = var.max_instances
  content_bucket             = local.content_bucket_name
  admin_password_hash_secret = local.admin_password_hash_secret
  admin_access_key_secret    = local.admin_access_key_secret
  admin_session_key_secret   = local.admin_session_key_secret
  domain                     = var.domain
  deployer_member            = local.deployer_member
  deletion_protection        = var.run_deletion_protection

  depends_on = [module.project_services]
}

# Names are computed here rather than read from module outputs, so that the
# storage/secrets modules (which need the runtime account) and the cloud-run module
# (which needs the names) do not form a module-level cycle.
locals {
  content_bucket_name        = "${var.project_id}-content"
  admin_password_hash_secret = "projects/${var.project_id}/secrets/admin-password-hash"
  admin_access_key_secret    = "projects/${var.project_id}/secrets/admin-access-key"
  admin_session_key_secret   = "projects/${var.project_id}/secrets/admin-session-key"
}

module "monitoring" {
  source = "../../modules/monitoring"

  project_id      = var.project_id
  service_name    = module.cloud_run.service_name
  host            = var.domain != "" ? var.domain : trimprefix(module.cloud_run.service_uri, "https://")
  alert_email     = var.alert_email
  billing_account = var.billing_account
  budget_amount   = var.budget_amount
  budget_currency = var.budget_currency

  depends_on = [module.project_services]
}
