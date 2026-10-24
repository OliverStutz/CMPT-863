# DECISION: a backend block cannot read Terraform variables, so "bucket as a variable"
# is done with partial configuration: the bucket is passed at init time, e.g.
#   terraform init -backend-config="bucket=<state bucket>"
# or -backend-config=backend.hcl (copy backend.hcl.example). CI passes
# TF_STATE_BUCKET. The prefix keeps dev and prod state apart (12.2).
terraform {
  backend "gcs" {
    prefix = "study/prod"
  }
}
