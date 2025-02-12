terraform {
  required_version = "~> 1.10.5"
  required_providers {
    administration = {
      version = "0.0.4"
      source  = "localhost/quortex/administration"
    }
    # quortex = {
    #   version = "0.0.1"
    #   source  = "localhost/quortex/quortex"
    # }
  }
}

# defined in local/variables.tf
provider "administration" {
  auth_server   = var.auth_server
  host          = var.host
  client_id     = var.client_id
  client_secret = var.client_secret
}

# example
# resource "administration_billing_plan" "premium" {
#   name     = "premium"
#   features = ["a", "b", "c"]
#   limits   = []
# }

