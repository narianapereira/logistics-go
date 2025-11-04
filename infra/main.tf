# --- PROVIDERS AND CONFIG ---


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

