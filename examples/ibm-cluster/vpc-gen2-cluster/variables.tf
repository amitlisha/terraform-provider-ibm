variable "flavor" {
  default = "cx2.2x4"
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
  default = "mycluster"
}

variable "region" {
  default = "us-south"
}

variable "service_instance_name" {
  default = "my-service-instance"
}

variable "cluster_name" {
  default = "mytestcluster"
}

variable "worker_pool_name" {
  default = "myvpc2pool"
}

variable "kube_version" {
  default = "1.21.7"
}
