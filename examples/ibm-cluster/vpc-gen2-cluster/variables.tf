variable "flavor" {
  default = "bx2.2x8"

}

variable "worker_count" {
  default = "1"
}

variable "zone" {
  default = "us-south-1"
}

variable "resource_group" {
  default = "Default"
}

variable "name" {
  default = "talnecluster"
}

variable "region" {
  default = "us-south"
}

variable "service_instance_name" {
  default = "talne-service-instance"
}

variable "cluster_name" {
  default = "talnetestcluster"
}

variable "worker_pool_name" {
  default = "talnevpc2pool"
}

variable "kube_version" {
  type        = string
  description = "Kubernetes version that you want to set up in your cluster."
}
