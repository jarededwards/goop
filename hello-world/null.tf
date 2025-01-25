##! uncomment this block to run terraform locally
terraform {
  required_providers {
    civo = {
      source = "civo/civo"
    }
  }
}
provider "civo" {
  region = "nyc1"
}

variable "cluster_name" {
  type    = string
}

variable "node_count" {
  type    = string
}

variable "node_type" {
  type    = string
  default = "g4s.kube.medium"
}

variable "token" {
  type    = string
}


provider "civo" {
  token = var.token
}

resource "civo_network" "cluster" {
  label = var.cluster_name
}

resource "civo_firewall" "cluster" {
  name                 = var.cluster_name
  network_id           = civo_network.cluster.id
  create_default_rules = true
}

resource "civo_kubernetes_cluster" "cluster" {
  name        = var.cluster_name
  network_id  = civo_network.cluster.id
  firewall_id = civo_firewall.cluster.id
  pools {
    label      = var.cluster_name
    size       = var.node_type
    node_count = tonumber(var.node_count)
  }
}
