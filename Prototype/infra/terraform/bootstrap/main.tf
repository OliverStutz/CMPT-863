# One-time bootstrap per GCP project, applied by Oliver with his own (owner)
# credentials: Terraform state bucket, Workload Identity Federation for GitHub
# Actions (no JSON keys), and the CI service accounts of 12.3.

resource "google_project_service" "bootstrap" {
  for_each = toset([
    "iam.googleapis.com",
    "iamcredentials.googleapis.com",
    "sts.googleapis.com",
    "cloudresourcemanager.googleapis.com",
    "serviceusage.googleapis.com",
    "storage.googleapis.com",
  ])

  project            = var.project_id
  service            = each.value
  disable_on_destroy = false
}

########## State bucket ##########

resource "google_storage_bucket" "tfstate" {
  project                     = var.project_id
  name                        = var.state_bucket_name != "" ? var.state_bucket_name : "${var.project_id}-tfstate"
  location                    = var.region
  uniform_bucket_level_access = true
  public_access_prevention    = "enforced"
  force_destroy               = false

  versioning {
    enabled = true
  }

  lifecycle_rule {
    condition {
      num_newer_versions = 20
    }
    action {
      type = "Delete"
    }
  }

  depends_on = [google_project_service.bootstrap]
}

########## Workload Identity Federation for GitHub Actions ##########

resource "google_iam_workload_identity_pool" "github" {
  project                   = var.project_id
  workload_identity_pool_id = "github-actions"
  display_name              = "GitHub Actions"

  depends_on = [google_project_service.bootstrap]
}

resource "google_iam_workload_identity_pool_provider" "github" {
  project                            = var.project_id
  workload_identity_pool_id          = google_iam_workload_identity_pool.github.workload_identity_pool_id
  workload_identity_pool_provider_id = "github"
  display_name                       = "GitHub OIDC"

  attribute_mapping = {
    "google.subject"       = "assertion.sub"
    "attribute.repository" = "assertion.repository"
    "attribute.ref"        = "assertion.ref"
  }
  # Only tokens from this one repository are accepted.
  attribute_condition = "assertion.repository == \"${var.github_repository}\""

  oidc {
    issuer_uri = "https://token.actions.githubusercontent.com"
  }
}

locals {
  pool        = google_iam_workload_identity_pool.github.name
  repo_member = "principalSet://iam.googleapis.com/${local.pool}/attribute.repository/${var.github_repository}"
  main_member = "principalSet://iam.googleapis.com/${local.pool}/attribute.ref/refs/heads/main"
}

########## CI service accounts (12.3) ##########

resource "google_service_account" "ci_deploy" {
  project      = var.project_id
  account_id   = "ci-deploy"
  display_name = "CI: image push, content upload"
}

resource "google_service_account" "ci_terraform_plan" {
  project      = var.project_id
  account_id   = "ci-terraform-plan"
  display_name = "CI: terraform plan (read-only plus state lock)"
}

resource "google_service_account" "ci_terraform_apply" {
  project      = var.project_id
  account_id   = "ci-terraform-apply"
  display_name = "CI: terraform apply from main only"
}

# Any workflow in the repository may use ci-deploy and ci-terraform-plan.
resource "google_service_account_iam_member" "wif_deploy" {
  service_account_id = google_service_account.ci_deploy.name
  role               = "roles/iam.workloadIdentityUser"
  member             = local.repo_member
}

resource "google_service_account_iam_member" "wif_plan" {
  service_account_id = google_service_account.ci_terraform_plan.name
  role               = "roles/iam.workloadIdentityUser"
  member             = local.repo_member
}

# ci-terraform-apply only from refs/heads/main (the pool already pins the repository).
# Prod approval is added with a protected GitHub environment "prod".
resource "google_service_account_iam_member" "wif_apply" {
  service_account_id = google_service_account.ci_terraform_apply.name
  role               = "roles/iam.workloadIdentityUser"
  member             = local.main_member
}

# ci-deploy's own roles (Artifact Registry writer on the repo, objectAdmin on the
# buckets, run.developer, act-as on the runtime accounts, serviceUsageConsumer) are
# granted by the env config, scoped to the resources it creates.

# ci-terraform-plan: viewer, securityReviewer, objectUser on the state bucket (12.3).
resource "google_project_iam_member" "plan" {
  for_each = toset(["roles/viewer", "roles/iam.securityReviewer"])

  project = var.project_id
  role    = each.value
  member  = "serviceAccount:${google_service_account.ci_terraform_plan.email}"
}

resource "google_storage_bucket_iam_member" "plan_state" {
  bucket = google_storage_bucket.tfstate.name
  role   = "roles/storage.objectUser"
  member = "serviceAccount:${google_service_account.ci_terraform_plan.email}"
}

# ci-terraform-apply: "the admin roles Terraform needs" (12.3), listed per service
# rather than roles/owner.
resource "google_project_iam_member" "apply" {
  for_each = toset([
    "roles/serviceusage.serviceUsageAdmin",
    "roles/run.admin",
    "roles/cloudscheduler.admin",
    "roles/datastore.owner",
    "roles/storage.admin",
    "roles/secretmanager.admin",
    "roles/artifactregistry.admin",
    "roles/iam.serviceAccountAdmin",
    "roles/iam.serviceAccountUser",
    "roles/resourcemanager.projectIamAdmin",
    "roles/monitoring.admin",
    "roles/logging.configWriter",
  ])

  project = var.project_id
  role    = each.value
  member  = "serviceAccount:${google_service_account.ci_terraform_apply.email}"
}
# The budget needs roles/billing.costsManager on the billing account for
# ci-terraform-apply. That is granted by hand (a billing account admin), or the
# budget is created by hand and billing_account left empty in the env tfvars.
