# --- PROVIDERS AND CONFIG ---

provider "google" {
  region  = var.gcp_region
}

resource "google_container_cluster" "logistics_go_cluster" {
  name                     = "logistics-go-cluster"
  location                 = var.gcp_region
  
  enable_autopilot         = true 
  
  deletion_protection      = false 
  
  initial_node_count       = 1
}

data "google_client_config" "current" {}

provider "kubernetes" {
  host                   = "https://${google_container_cluster.logistics_go_cluster.endpoint}"
  token                  = data.google_client_config.current.access_token
  cluster_ca_certificate = base64decode(google_container_cluster.logistics_go_cluster.master_auth[0].cluster_ca_certificate)
}

# --- YAML AND DEPLOY ---

data "local_file" "manifest_yaml" {
  filename = "manifests/deployment.yaml"
}

locals {
  # Substitutes the IMAGE_PLACEHOLDER for the actual tag for the docker image
  rendered_manifest = replace(
    data.local_file.manifest_yaml.content,
    "IMAGE_PLACEHOLDER",
    var.app_docker_image
  )
}

resource "kubernetes_manifest" "logistics_app" {
  manifest = yamldecode(local.rendered_manifest)
}

# --- OUTPUT ---

data "kubernetes_service" "logistics_service_data" {
  metadata {
    name = "logistics-service"
  }
  depends_on = [kubernetes_manifest.logistics_app]

}

# Prints external api in actions
output "app_external_ip" {
    value = data.kubernetes_service.logistics_service_data.status[0].load_balancer[0].ingress[0].ip
}