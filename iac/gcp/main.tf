# GKE Cluster
resource "google_container_cluster" "primary" {
  name     = "devops-test-gke"
  location = var.region
}

# Google SQL for Postgres
resource "google_sql_database_instance" "postgres" {
  name             = "devops-test-postgres"
  database_version = "POSTGRES_17"
  region           = var.region
}
