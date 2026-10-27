# Dev environment on project cmpt863-study. Nothing here is secret.
project_id = "cmpt863-study"
region     = "northamerica-northeast1"

min_instances = 0
max_instances = 10

domain = "" # none yet; the service is reached at its run.app URL

alert_email     = "oliver.stutz@gmail.com"
billing_account = "018B43-B4FE82-8F35D6"
budget_amount   = 25

ci_deploy_sa_email = "ci-deploy@cmpt863-study.iam.gserviceaccount.com" # created by the bootstrap

firestore_delete_protection = true
run_deletion_protection     = false
force_destroy_buckets       = false

# The billing account is in CAD; a budget must use the account currency.
budget_currency = "CAD"

# Current image (commit f48c1ba). CI overwrites this on deploy.
image_digest = "northamerica-northeast1-docker.pkg.dev/cmpt863-study/study/api@sha256:52cf677d1d5915c2246355c95a2234d0549ee77f28385a90e06dc647a8146052"

# Dev accepts configs whose version starts with "test" (fast end-to-end runs). Never set on prod.
extra_env = { ALLOW_TEST_CONFIGS = "true" }
