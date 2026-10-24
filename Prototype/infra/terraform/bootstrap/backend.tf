# State for the bootstrap lives in the bucket the bootstrap itself created.
# First run (bucket absent): comment this block out, apply with local state,
# then restore it and run: terraform init -migrate-state
terraform {
  backend "gcs" {
    bucket = "cmpt863-study-tfstate"
    prefix = "study/bootstrap"
  }
}
