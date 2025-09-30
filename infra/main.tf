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
  alias       = "gke"
  config_path = "~/.kube/config"
}

# --- YAML AND DEPLOY ---

data "local_file" "deployment_yaml" {
  filename = "manifests/deployment.yaml"
}

data "local_file" "service_yaml" {
  filename = "manifests/service.yaml"
}

locals {
  # Substitutes the IMAGE_PLACEHOLDER for the actual tag for the docker image
  rendered_deployment = replace(
    data.local_file.deployment_yaml.content,
    "IMAGE_PLACEHOLDER",
    var.app_docker_image
  )

  rendered_service = data.local_file.service_yaml.content
}

resource "kubernetes_manifest" "logistics_app_deployment" {
  provider = kubernetes.gke

  manifest = yamldecode(local.rendered_deployment)
}

resource "kubernetes_manifest" "logistics_app_service" {
  provider = kubernetes.gke

  manifest = yamldecode(local.rendered_service)

  depends_on = [kubernetes_manifest.logistics_app_deployment]
}

# --- OUTPUT ---

data "kubernetes_service" "logistics_service_data" {
  provider = kubernetes.gke
    
  metadata {
    name = "logistics-service"
  }
  depends_on = [kubernetes_manifest.logistics_app_service]
}

# Prints external api in actions
output "app_external_ip" {
    value = data.kubernetes_service.logistics_service_data.status[0].load_balancer[0].ingress[0].ip
}