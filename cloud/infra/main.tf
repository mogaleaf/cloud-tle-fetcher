terraform {
  required_providers {
    aws     = {
      source  = "hashicorp/aws"
      version = "~> 3.48.0"
    }
    random  = {
      source  = "hashicorp/random"
      version = "~> 3.1.0"
    }
    archive = {
      source  = "hashicorp/archive"
      version = "~> 2.2.0"
    }
  }

  required_version = "~> 1.0"
}

provider "aws" {
  region = var.aws_region
}

module "deploy_tle_fetcher" {
  source = "./modules/deploy_aws_lambda"

  bucket_name           = "satellite-leo-planning-tle-fetcher-bucket"
  binary_name           = "tle-fetcher.zip"
  binary_source_dir     = "../tle_fetcher/lambda/build"
  lambda_function_name  = "satellite-leo-planning-tle-fetcher"
  lambda_handler        = "tle-fetcher"
  environment_variables = {
    "RDS_HOSTNAME" = aws_db_instance.datastore.address,
    "RDS_PORT"     = aws_db_instance.datastore.port,
    "RDS_USER"     = aws_db_instance.datastore.username,
    "RDS_PASSWORD" = aws_db_instance.datastore.password,
    "RDS_SCHEMA"   = var.datastore_schema,
    "AWS_LAMBDA"   = "true",
  }
  depends_on            = [
    aws_db_instance.datastore,
  ]
}
