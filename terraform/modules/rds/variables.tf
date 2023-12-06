variable "db_name" {
  description = "The name of the database to create when the DB instance is created."
  type        = string
  default     = "tulsadoom"
}

variable "db_username" {
  description = "Username for the master DB user."
  type        = string
}

variable "db_password" {
  description = "Password for the master DB user."
  type        = string
}

variable "vpc_id" {
  description = "The VPC ID where the cluster and workers will be deployed."
  type        = string
}

variable "subnet_ids" {
  description = "A list of private subnet IDs to launch resources in."
  type        = list(string)
}
