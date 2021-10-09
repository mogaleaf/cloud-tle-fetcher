output "lambda_bucket_name" {
  description = "Name of the S3 bucket used to store function code."

  value = module.deploy_tle_fetcher.lambda_bucket_name
}
output "rds_hostname" {
  description = "RDS instance hostname"
  value       = aws_db_instance.datastore.address
  sensitive   = true
}

output "rds_port" {
  description = "RDS instance port"
  value       = aws_db_instance.datastore.port
  sensitive   = true
}

output "rds_username" {
  description = "RDS instance root username"
  value       = aws_db_instance.datastore.username
  sensitive   = true
}

