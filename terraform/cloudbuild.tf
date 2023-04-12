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