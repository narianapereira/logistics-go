# Variables that I will use in my main tf, that are loaded by TF and come from the VARs defined on my gb actions

variable "app_docker_image" {
  description = "The full Docker image path and tag"
  type        = string
}