resource "google_project_iam_member" "cloudbuild_iam" {
  for_each = toset([
    "roles/run.developer",
    "roles/iam.serviceAccountUser"
  ])
  role    = each.key
  member  = "serviceAccount:${var.project_id}@cloudbuild.gserviceaccount.com"
  project = var.project_id
}


resource "google_cloudbuild_trigger" "build-trigger" {
  location = "us-central1"
  name     = "container-builder"
  filename = "./terraform/cloudbuild.yaml"

  github {
    owner = "Bayathy"
    name  = "api-server"
    push {
      branch = "^main$"
    }
  }

  include_build_logs = "INCLUDE_BUILD_LOGS_WITH_STATUS"
}