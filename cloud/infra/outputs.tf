output "lambda_bucket_name" {
  description = "Name of the S3 bucket used to store function code."

  value = module.deploy_tle_fetcher.lambda_bucket_name
}
