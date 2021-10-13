module "deploy_fetch" {
  source                = "../deploy_aws_lambda"
  bucket_name           = "tle-fetcher-solution-fetch-bucket"
  binary_name           = "fetch.zip"
  binary_source_dir     = "${path.module}/../../../tle_fetcher_solution/lambda/fetch/build"
  lambda_function_name  = "fetch"
  lambda_handler        = "fetch"
  iam_arn = var.iam_arn
  environment_variables = {
    "REGION"     = "us-east-1",
    "AWS_LAMBDA" = "true",
  }
}

module "deploy_join" {
  source                = "../deploy_aws_lambda"
  bucket_name           = "tle-fetcher-solution-join-bucket"
  binary_name           = "join.zip"
  binary_source_dir     = "${path.module}/../../../tle_fetcher_solution/lambda/join/build"
  lambda_function_name  = "join"
  lambda_handler        = "join"
  iam_arn = var.iam_arn
  environment_variables = {
    "REGION"     = "us-east-1",
    "AWS_LAMBDA" = "true",
  }
}


module "deploy_receive" {
  source                = "../deploy_aws_lambda"
  bucket_name           = "tle-fetcher-solution-receive-bucket"
  binary_name           = "receive.zip"
  binary_source_dir     = "${path.module}/../../../tle_fetcher_solution/lambda/receive/build"
  lambda_function_name  = "receive"
  lambda_handler        = "receive"
  iam_arn = var.iam_arn
  environment_variables = {
    "REGION"     = "us-east-1",
    "AWS_LAMBDA" = "true",
  }
}
