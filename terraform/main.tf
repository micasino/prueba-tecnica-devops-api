provider "google" {}


terraform {
  required_version = ">=0.51.1"

  required_providers {

    tfe = {
      version = "~> 0.50.0"
    }

    google = {
      source  = "hashicorp/google"
      version = "~> 3.74.0"
    }
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "~> 2.0.3"
      version = "2.24.0"
    }
    helm = {
      source  = "hashicorp/helm"
      version = "~> 2.1.0"
    }

  }
}


resource "google_compute_network" "iac_vpc" {
  name    = "iac-vpc"
  project = "iacchalenge"
}

resource "google_compute_subnetwork" "gke_subnet" {
  name          = "gke-subnet"
  project       = "iacchalenge"
  ip_cidr_range = "10.0.1.0/24"
  network       = google_compute_network.iac_vpc.name
  region        = "us-central1"
}

resource "google_compute_subnetwork" "sql_subnet" {
  name          = "sql-subnet"
  project       = "iacchalenge"
  ip_cidr_range = "10.0.2.0/24"
  network       = google_compute_network.iac_vpc.name
  region        = "us-central1"
}

resource "google_container_cluster" "instance_gke_cluster" {
  name       = "iac-gke-cluster"
  project    = "iacchalenge"
  location   = "us-central1"
  network    = google_compute_network.iac_vpc.name
  subnetwork = google_compute_subnetwork.gke_subnet.self_link

  initial_node_count = 1

  node_config {
    preemptible  = false
    machine_type = "e2-medium"
  }
}

resource "null_resource" "configure_kubeconfig" {
  provisioner "local-exec" {
    command = <<-EOT
      gcloud container clusters get-credentials iac-gke-cluster --region us-central1 --project iacchalenge
    EOT
  }

  depends_on = [
    google_container_cluster.instance_gke_cluster,
  ]
}

resource "google_redis_instance" "iac_redis_instance" {
  name           = "iac-redis-instance"
  project        = "iacchalenge"
  memory_size_gb = 1
  region         = "us-central1"
  tier           = "BASIC"

  authorized_network = "iac-vpc"
}

resource "google_sql_database_instance" "iac_sql_instance" {
  name             = "iac-sql-instance"
  database_version = "POSTGRES_13"
  region           = "us-central1"
  project          = "iacchalenge"
  settings {
    tier = "db-f1-micro"
  }

  depends_on = [google_compute_subnetwork.sql_subnet]
}

###### hasta aca todo OK ########



provider "kubernetes" {
  config_path = "~/.kube/config"
  host = "https://35.184.112.207:443"
}

provider "helm" {
  kubernetes {
    config_path = "~/.kube/config"
  }
}
#output "kubeconfig_path" {
#  value = data.kubernetes_config_map.example.data["kubeconfig"]
#}
