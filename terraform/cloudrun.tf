resource "google_cloud_run_v2_service" "default" {
  location = var.region
  name     = "api-server"

  template {
    scaling {
      max_instance_count = 1
    }
    containers {
      image = "us-west1-docker.pkg.dev/api-server/my-repo/server-image"
      resources {
        limits = { "memory" : "0.5Gi", "cpu" : "1" }
      }

      ports {
        container_port = "8080"
      }

      env {
        name = "DB_PASSWORD"
        value = var.db_password
      }
    }
  }

  traffic {
    type = "TRAFFIC_TARGET_ALLOCATION_TYPE_LATEST"
    percent = 100
  }
}
