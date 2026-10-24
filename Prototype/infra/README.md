# Infrastructure: study deployment on Google Cloud

Built to `02 - Architecture.md` sections 2, 3, 12, 13, 14 and the decisions of 2026-10-24: one Cloud Run service serves both SPAs, the content assets and both APIs. There is no load balancer, no CDN and no IAP. Data lives in Firestore (native mode, database `(default)`), content assets in a private GCS bucket read by the service, the shared admin password as a bcrypt hash in Secret Manager. A Cloud Scheduler job starts the `sweeper` Cloud Run job every 5 minutes.

Project `cmpt863-study` exists with billing linked; the bootstrap is applied and its state is in the bucket. The dev environment apply is the next step.

## Layout

```text
Dockerfile, .dockerignore          image: Node 26 builds both apps, Go 1.27 builds /api and /sweeper, distroless runtime
.github/workflows/ci.yml           PR: go vet/test, npm build/test, terraform fmt -check and validate
.github/workflows/deploy.yml       main: image to Artifact Registry, terraform apply, content upload, smoke test
infra/terraform/bootstrap/         once per project: state bucket, Workload Identity Federation, CI accounts
infra/terraform/modules/           project-services, artifact-registry, cloud-run, storage, firestore, secrets, monitoring
infra/terraform/envs/{dev,prod}/   call the modules; own state prefix (study/dev, study/prod) and tfvars
```

## Service accounts (12.3)

| Account | Created by | Roles |
|---|---|---|
| `api-runtime` | env (cloud-run) | `datastore.user`; `storage.objectViewer` on the content bucket; `secretmanager.secretAccessor` on `admin-password-hash`, `admin-session-key` and `admin-access-key`; `logging.logWriter`; `monitoring.metricWriter` |
| `sweeper-runtime` | env (cloud-run) | `datastore.user`; `logging.logWriter` |
| `scheduler-invoker` | env (cloud-run) | `run.invoker` on the sweeper job |
| `ci-deploy` | bootstrap | from the env: `artifactregistry.writer` on the repo; `storage.objectAdmin` on the content and TTS cache buckets; `run.developer`; `iam.serviceAccountUser` on the two runtime accounts; `serviceusage.serviceUsageConsumer` |
| `ci-terraform-plan` | bootstrap | `viewer`, `iam.securityReviewer`, `storage.objectUser` on the state bucket |
| `ci-terraform-apply` | bootstrap | per-service admin roles (run, scheduler, datastore, storage, secret manager, artifact registry, service accounts, project IAM, monitoring, logging config, service usage); usable only from `refs/heads/main` |

GitHub reaches these accounts through Workload Identity Federation. The pool accepts tokens only from the one repository named in the bootstrap. No JSON keys exist.

## Bootstrap, in order

Do this once per project (dev, then prod). Commands that create or bill anything are Oliver's to run.

1. **Project.** Create the project (`gcloud projects create <id>`) and note its id. Separate projects for dev and prod are recommended (12.2).
2. **Billing.** Link a billing account: `gcloud billing projects link <id> --billing-account=<XXXXXX-XXXXXX-XXXXXX>`.
3. **State bucket, WIF and CI accounts.** In `infra/terraform/bootstrap`, copy `terraform.tfvars.example` to `terraform.tfvars`, set `project_id` and `github_repository`, then with your own credentials (`gcloud auth application-default login`):
   `terraform init && terraform apply`.
   The bootstrap's own state is then migrated into the bucket it created (`backend.tf`, prefix `study/bootstrap`): restore the backend block and run `terraform init -migrate-state`. Done for `cmpt863-study` on 2026-10-24; both the bootstrap and the environments keep their state in `cmpt863-study-tfstate`, versioned. If you prefer to make the state bucket by hand instead: `gcloud storage buckets create gs://<id>-tfstate --location=northamerica-northeast1 --uniform-bucket-level-access --public-access-prevention`, then `gcloud storage buckets update gs://<id>-tfstate --versioning`; WIF and the CI accounts then also have to be made by hand, so the bootstrap config is the easier route.
4. **GitHub variables.** Create GitHub environments `dev` and `prod` (protect `prod` with required reviewers) and set on each, from the bootstrap outputs:
   - `GCP_PROJECT_ID`
   - `GCP_REGION` (`northamerica-northeast1`)
   - `GCP_WIF_PROVIDER` (output `GCP_WIF_PROVIDER`, `projects/<number>/locations/global/workloadIdentityPools/github-actions/providers/github`)
   - `GCP_DEPLOY_SA` (output `GCP_DEPLOY_SA`)
   - `GCP_TERRAFORM_SA` (output `GCP_TERRAFORM_SA`)
   - `TF_STATE_BUCKET` (output `TF_STATE_BUCKET`)
   - `ALERT_EMAIL` (optional)
5. **First apply, by hand.** In `infra/terraform/envs/dev`: copy `terraform.tfvars.example` to `terraform.tfvars` and fill it in (the copy is git-ignored), then
   `terraform init -backend-config="bucket=<id>-tfstate"` and `terraform apply`.
   With `image_digest` empty the service and job run Google's public hello image, so this first apply works before any image exists and creates the Artifact Registry repository the first CI run pushes to.
6. **Secrets.** The apply creates the secret containers only; add the values out of band so they never enter Terraform state:
   - admin password: `make admin-password` prints the bcrypt hash (12.4); then `printf '%s' '<hash>' | gcloud secrets versions add admin-password-hash --data-file=- --project=<id>`
   - session key: `head -c 32 /dev/urandom | gcloud secrets versions add admin-session-key --data-file=- --project=<id>`
   - admin access key: `make admin-key` prints a random key; then `printf '%s' '<key>' | gcloud secrets versions add admin-access-key --data-file=- --project=<id>`
   The access key is a URL gate: `/admin/?k=<key>` is the only way the admin page renders, and the shared password is asked for after that. Share the key only out of band (in person, or a password manager), never in the repository, an issue, a chat channel or an email thread that others can read.
   Rotating any of the three means adding a new version; the service picks it up within 5 minutes.
7. **First deploy.** Push to `main` (or run `deploy` by hand). The workflow builds and pushes the image, applies Terraform with the new `image_digest`, uploads `content/` and checks `/api/v1/healthz`.
8. **Content upload.** Done by `deploy.yml`: `content/` is mirrored to the content bucket without deleting anything. MP3s are git-ignored and uploaded only when present (generated by `tools/tts` before the upload). By hand: `gcloud storage rsync content gs://<id>-content --recursive`.
9. **Prod.** Repeat 1 to 6 for the prod project in `envs/prod` (example tfvars: `min_instances = 1`, deletion protection on), then run `deploy` by hand with `environment: prod`.

## Variables to set (`envs/*/terraform.tfvars`)

| Variable | Default | Note |
|---|---|---|
| `project_id` | none | OPEN Q |
| `region` | `northamerica-northeast1` | 3.5 |
| `firestore_location` | the region | fixed once the database exists |
| `min_instances` | 0 | 1 during active recruitment (prod example 1) |
| `max_instances` | 10 | |
| `domain` | empty | OPEN Q; see below |
| `alert_email` | empty | alert and budget notices |
| `billing_account` | empty | set to create the budget from Terraform; needs `roles/billing.costsManager` on the billing account for the applying identity, else create the budget by hand (50/90/100 %) |
| `budget_amount`, `budget_currency` | 25, USD | currency must match the billing account |
| `ci_deploy_sa_email` | empty | bootstrap output |
| `image_digest` | empty | set by CI |

Service settings fixed in the module: 1 vCPU, 512 MiB, concurrency 80, request timeout 60 s, CPU only during requests, ingress all, public (`allUsers` invoker, participants are anonymous). Environment: `CONTENT_BUCKET`, `TRUSTED_PROXY_HOPS=1`, `FIRESTORE_PROJECT`, `ADMIN_PASSWORD_HASH_SECRET`, `ADMIN_SESSION_KEY_SECRET` and `ADMIN_ACCESS_KEY_SECRET`. The last three are plain env vars holding Secret Manager resource names (`projects/<id>/secrets/admin-password-hash`, `.../admin-session-key`, `.../admin-access-key`); the service reads `versions/latest` itself and caches it for 5 minutes, so no secret value passes through Terraform or the Cloud Run configuration. An organisation policy that restricts `allUsers` would block the public invoker binding.

**Domain.** Until a domain is named the service is reached at its `run.app` URL. Setting `domain` adds a Cloud Run domain mapping; first check that domain mappings are available in the chosen region (they are not in every region and have been a preview feature) and verify the domain for the project. The output `domain_dns_records` lists the records to create at the domain host. If mappings are not available, the fallback is a load balancer for the domain, with its fixed monthly cost.

## Credential hygiene

- Nothing in `infra/` holds a secret value. Terraform creates only the secret containers (`admin-password-hash`, `admin-session-key`, `admin-access-key`); their values are added by hand with `gcloud secrets versions add`, so they never appear in code, tfvars or Terraform state. The tfvars hold project ids, an email and a billing account id, which are not secrets.
- CI reaches GCP only through Workload Identity Federation; there are no service account JSON keys anywhere.
- Git-ignored (root `.gitignore`): `envs/*/terraform.tfvars`, `envs/*/backend.hcl`, `bootstrap/terraform.tfvars` and `bootstrap/terraform.tfstate*`, plus `.terraform/` and all `*.tfstate`. The `.example` files stay committed. If CI should read a committed `terraform.tfvars`, set those values as GitHub environment variables instead (`TF_VAR_*` in `deploy.yml`).
- The content bucket (and the TTS cache bucket) has uniform bucket-level access and public access prevention `enforced`, both set explicitly in `modules/storage`, and no `allUsers` or `allAuthenticatedUsers` binding; a precondition in the module fails the plan if one is passed in. Participants receive assets only through the service's `/a/*` handler.
- `ci.yml` fails a pull request if any file matches common credential patterns (`-----BEGIN`, `"refresh_token"`, `ya29.`, Google API keys `AIza...`, `"private_key"`), excluding `.example` files and docs.

## Rough monthly cost, 40 participants (12.5)

| Item | USD per month |
|---|---|
| Cloud Run, min 0 | under 1 |
| Cloud Run, min 1 kept all month | about 10 |
| Egress | under 1 |
| Firestore | 0 (free tier) |
| Cloud Storage | under 1 |
| Text-to-Speech | 0 to 1 |
| Secret Manager, Scheduler, Artifact Registry | about 1 |
| Monitoring and logging | 0 (free tier) |
| **Total** | **about 1 to 13** |

Firestore point-in-time recovery and the daily and weekly backups add a small storage charge. Check with the GCP pricing calculator for the chosen region.

## Still needed from Oliver

- GCP project id(s) for dev and prod, and the billing account.
- The GitHub repository (`owner/name`) for the WIF pool.
- The domain, if any, and a check that Cloud Run domain mappings work in the region.
- The admin password (its bcrypt hash goes into `admin-password-hash`) and the admin access key (`admin-access-key`), both set by hand.
- Alert email and budget amount.
- Running the bootstrap and first apply (steps 1 to 6).

## Local checks

```sh
terraform fmt -recursive infra/terraform
terraform -chdir=infra/terraform/envs/dev init -backend=false && terraform -chdir=infra/terraform/envs/dev validate
docker build -t study-api:dev .
```
