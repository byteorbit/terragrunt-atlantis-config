resource "random_pet" "test" {
  prefix = var.prefix
  length = 2
}
terraform {
  required_providers {
    random = {
      source  = "hashicorp/random"
      version = "~> 3.6"
    }
  }
  required_version = ">= 1.5.7"
}
