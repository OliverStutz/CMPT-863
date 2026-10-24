# One Cloud Run service serves both SPAs, the content assets and both APIs (2, 3.4).
# Also: the runtime accounts (12.3), the sweeper Cloud Run job with its Cloud
# Scheduler trigger (5.4), and an optional domain mapping (12.1).

locals {
  # Before the first image is pushed, a public placeholder lets the first apply succeed.
  image = var.image_digest != "" ? var.image_digest : "us-docker.pkg.dev/cloudrun/container/hello"
}

########## Service accounts and project-level roles (12.3) ##########

resource "google_service_account" "api" {
  project      = var.project_id
  account_id   = "api-runtime"
  display_name = "Study API runtime"
}

resource "google_service_account" "sweeper" {
  project      = var.project_id
  account_id   = "sweeper-runtime"
  display_name = "Study sweeper job runtime"
}

resource "google_service_account" "scheduler" {
  project      = var.project_id
  account_id   = "scheduler-invoker"
  display_name = "Cloud Scheduler invoker for the sweeper job"
}

locals {
  project_roles = {
    "api-log"     = { member = "serviceAccount:${google_service_account.api.email}", role = "roles/logging.logWriter" }
    "api-metric"  = { member = "serviceAccount:${google_service_account.api.email}", role = "roles/monitoring.metricWriter" }
    "sweeper-log" = { member = "serviceAccount:${google_service_account.sweeper.email}", role = "roles/logging.logWriter" }
  }
}

resource "google_project_iam_member" "runtime" {
  for_each = local.project_roles

  project = var.project_id
  role    = each.value.role
  member  = each.value.member
}

########## API service ##########

resource "google_cloud_run_v2_service" "api" {
  project             = var.project_id
  name                = var.service_name
  location            = var.region
  ingress             = "INGRESS_TRAFFIC_ALL"
  deletion_protection = var.deletion_protection

  template {
    service_account                  = google_service_account.api.email
    timeout                          = "60s"
    max_instance_request_concurrency = var.concurrency

    scaling {
      min_instance_count = var.min_instances
      max_instance_count = var.max_instances
    }

    containers {
      image = local.image

      ports {
        container_port = 8080
      }

      resources {
        limits = {
          cpu    = var.cpu
          memory = var.memory
        }
        cpu_idle          = true # CPU allocated only during requests (12.4)
        startup_cpu_boost = true
      }

      env {
        name  = "CONTENT_BUCKET"
        value = var.content_bucket
      }
      env {
        name  = "TRUSTED_PROXY_HOPS"
        value = "1"
      }
      env {
        name  = "FIRESTORE_PROJECT"
        value = var.project_id
      }

      # Secret contract agreed with the backend: plain env vars holding Secret Manager
      # resource names (projects/<p>/secrets/<id>); the service reads
      # ".../versions/latest" itself and caches it (8.2, 12.4), so a rotated value takes
      # effect without a redeploy. No secret value ever passes through Terraform.
      env {
        name  = "ADMIN_PASSWORD_HASH_SECRET"
        value = var.admin_password_hash_secret
      }
      env {
        name  = "ADMIN_SESSION_KEY_SECRET"
        value = var.admin_session_key_secret
      }
      env {
        name  = "ADMIN_ACCESS_KEY_SECRET"
        value = var.admin_access_key_secret
      }

      dynamic "env" {
        for_each = var.extra_env
        content {
          name  = env.key
          value = env.value
        }
      }
    }
  }

  traffic {
    type    = "TRAFFIC_TARGET_ALLOCATION_TYPE_LATEST"
    percent = 100
  }
}

# Public: participants are anonymous and the service does its own auth (12.1).
# An organisation policy restricting allUsers (domain restricted sharing) blocks this.
resource "google_cloud_run_v2_service_iam_member" "public" {
  project  = var.project_id
  location = google_cloud_run_v2_service.api.location
  name     = google_cloud_run_v2_service.api.name
  role     = "roles/run.invoker"
  member   = "allUsers"
}

########## Sweeper job and its trigger (5.4) ##########

resource "google_cloud_run_v2_job" "sweeper" {
  project             = var.project_id
  name                = "sweeper"
  location            = var.region
  deletion_protection = var.deletion_protection

  template {
    task_count = 1

    template {
      service_account = google_service_account.sweeper.email
      timeout         = "300s"
      max_retries     = 1

      containers {
        image = local.image
        # Same image as the service; the sweeper binary sits next to /api.
        command = var.image_digest != "" ? ["/sweeper"] : null

        resources {
          limits = {
            cpu    = "1"
            memory = "512Mi"
          }
        }

        env {
          name  = "FIRESTORE_PROJECT"
          value = var.project_id
        }
      }
    }
  }
}

resource "google_cloud_run_v2_job_iam_member" "scheduler_invoker" {
  project  = var.project_id
  location = google_cloud_run_v2_job.sweeper.location
  name     = google_cloud_run_v2_job.sweeper.name
  role     = "roles/run.invoker"
  member   = "serviceAccount:${google_service_account.scheduler.email}"
}

resource "google_cloud_scheduler_job" "sweeper" {
  project          = var.project_id
  region           = var.region
  name             = "sweeper-every-5-min"
  description      = "Marks sessions without events for 30 minutes as Abandoned."
  schedule         = "*/5 * * * *"
  time_zone        = "Etc/UTC"
  attempt_deadline = "60s"

  retry_config {
    retry_count = 0
  }

  http_target {
    http_method = "POST"
    uri         = "https://run.googleapis.com/v2/projects/${var.project_id}/locations/${var.region}/jobs/${google_cloud_run_v2_job.sweeper.name}:run"

    oauth_token {
      service_account_email = google_service_account.scheduler.email
      scope                 = "https://www.googleapis.com/auth/cloud-platform"
    }
  }

  depends_on = [google_cloud_run_v2_job_iam_member.scheduler_invoker]
}

########## CI deployer (12.3) ##########
# The ci-deploy account itself is created in infra/terraform/bootstrap together with
# Workload Identity Federation; here it gets run.developer and act-as on the two
# runtime accounts.
# DECISION: 12.3 gives ci-deploy run.developer; the brief said "run admin". The
# document's narrower role is used, because terraform apply (which needs run.admin to
# set the allUsers invoker policy) runs as ci-terraform-apply, not as ci-deploy.

resource "google_project_iam_member" "deployer_run" {
  count = var.deployer_member != "" ? 1 : 0

  project = var.project_id
  role    = "roles/run.developer"
  member  = var.deployer_member
}

resource "google_service_account_iam_member" "deployer_act_as_api" {
  count = var.deployer_member != "" ? 1 : 0

  service_account_id = google_service_account.api.name
  role               = "roles/iam.serviceAccountUser"
  member             = var.deployer_member
}

resource "google_service_account_iam_member" "deployer_act_as_sweeper" {
  count = var.deployer_member != "" ? 1 : 0

  service_account_id = google_service_account.sweeper.name
  role               = "roles/iam.serviceAccountUser"
  member             = var.deployer_member
}

########## Optional custom domain (OPEN Q) ##########
# Cloud Run domain mappings are not available in every region and have been a preview
# feature. Check availability for var.region before setting var.domain; if they are
# not available the fallback is a load balancer just for the domain (12.1), which
# brings back its fixed monthly cost. The domain must be verified for the project
# (Search Console) by the account that applies this.
resource "google_cloud_run_domain_mapping" "this" {
  count = var.domain != "" ? 1 : 0

  project  = var.project_id
  location = var.region
  name     = var.domain

  metadata {
    namespace = var.project_id
  }

  spec {
    route_name = google_cloud_run_v2_service.api.name
  }
}
