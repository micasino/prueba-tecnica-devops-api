provider "google" {
    project = "iacchalenge"
    credentials = file("iacchallenge-credentials.json")
    region = "us-central1"
    zone = "us-central1-c"

}


provider "helm" {
  kubernetes {
    config_path = "~/.kube/config"
  }
}

terraform {
  required_providers {
 
    helm = {
      source  = "hashicorp/helm"
      version = "~> 2.0"
    }
    kubernetes = {
      source = "hashicorp/kubernetes"
      version = "~> 2.0"
    }
  }
  required_version = ">= 1"
}




resource "google_compute_network" "iac_vpc" {
  name = "iac-vpc"
}

resource "google_compute_subnetwork" "gke_subnet" {
  name          = "gke-subnet"
  ip_cidr_range = "10.0.1.0/24"
  network       = google_compute_network.iac_vpc.name
  region        = "us-central1"
}

resource "google_compute_subnetwork" "sql_subnet" {
  name          = "sql-subnet"
  ip_cidr_range = "10.0.2.0/24"
  network       = google_compute_network.iac_vpc.name
  region        = "us-central1"
}

resource "google_container_cluster" "instance_gke_cluster" {
  name     = "iac-gke-cluster"
  location = "us-central1"
  network  = google_compute_network.iac_vpc.name
  subnetwork = google_compute_subnetwork.gke_subnet.self_link
  
  initial_node_count = 1
  
  node_config {
    preemptible       = false
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
  name     = "iac-redis-instance"
  memory_size_gb = 1
  region   = "us-central1"
  tier     = "BASIC"

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


# Retrieve an access token as the Terraform runner
data "google_client_config" "provider" {}



provider "kubernetes" {
  host  = "https://${data.google_container_cluster.instance_gke_cluster.endpoint}"
  token = data.google_client_config.provider.access_token
  cluster_ca_certificate = base64decode(
    data.google_container_cluster.instance_gke_cluster.master_auth[0].cluster_ca_certificate,
  )
}
