terraform {
  backend "gcs" {
    bucket      = "iac-storage-bucket"
    prefix      = "iac-terraform-environment"
    config_path = "../states/cert_manager.tfstate"
  }
}

