variable "aws_region" {
  description = "AWS region for all resources."

  type    = string
  default = "us-east-1"
}
variable "db_user" {
  description = "RDS root user name"
  sensitive   = true
}
variable "db_password" {
  description = "RDS root user password"
  sensitive   = true
}

variable "datastore_schema" {
  description = "handler of the lambda"
  type        = string
  default     = "datastore"
}
