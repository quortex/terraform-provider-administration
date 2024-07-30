terraform {
  required_providers {
    administration = {
      source = "registry.terraform.io/quortex/administration"
    }
  }
  required_version = ">= 1.9.0"
}

provider "administration" {
  client_id     = "my_client_id"
  client_secret = "my_secret_id"
}

resource "administration_billing_plan" "premium" {
  name     = "premium"
  features = ["a", "b", "c"]
  limits   = []
}
