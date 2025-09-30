# This place defines where the state of terraform will be stored
# Bucket on gcp
terraform {
  backend "gcs" {
    bucket  = "terraform-state-logistics-go" 
    
    prefix  = "gke/logistics-go-app/production"
  }
}