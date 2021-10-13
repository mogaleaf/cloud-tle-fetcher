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

#module "deploy_tle_fetcher" {
#  source = "./modules/deploy_aws_lambda"
#
#  bucket_name           = "satellite-leo-planning-tle-fetcher-bucket"
#  binary_name           = "tle-fetcher.zip"
#  binary_source_dir     = "../tle_fetcher/lambda/build"
#  lambda_function_name  = "satellite-leo-planning-tle-fetcher"
#  lambda_handler        = "lambda"
#  iam_arn = aws_iam_role.lambda_exec.arn
#  environment_variables = {
#    "RDS_HOSTNAME" = aws_db_instance.datastore.address,
#    "RDS_PORT"     = aws_db_instance.datastore.port,
#    "RDS_USER"     = aws_db_instance.datastore.username,
#    "RDS_PASSWORD" = aws_db_instance.datastore.password,
#    "RDS_SCHEMA"   = var.datastore_schema,
#    "AWS_LAMBDA"   = "true",
#  }
#  depends_on            = [
#    aws_db_instance.datastore,
#  ]
#}

resource "aws_iam_role_policy" "lambda_policy" {
  name = "lambda_policy"
  role = aws_iam_role.lambda_exec.id

  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Action" : ["logs:*"],
        "Effect" : "Allow",
        "Resource" : ["arn:aws:logs:*:*:*"]
      },
      {
        "Action" : ["execute-api:ManageConnections"],
        "Effect" : "Allow",
        "Resource" : ["arn:aws:execute-api:*:*:*/@connections/*"]
      },
      {
        "Effect" : "Allow",
        "Action" : [
          "dynamodb:BatchGetItem",
          "dynamodb:GetItem",
          "dynamodb:Query",
          "dynamodb:Scan",
          "dynamodb:BatchWriteItem",
          "dynamodb:PutItem",
          "dynamodb:UpdateItem",
          "dynamodb:GetShardIterator",
          "dynamodb:DescribeStream",
          "dynamodb:GetRecords",
          "dynamodb:ListStreams"

        ],
        "Resource" : [
          module.deploy_dynamo_db.arn_tle,
          module.deploy_dynamo_db.arn_com,
          module.deploy_dynamo_db.arn_satellite,
          module.deploy_dynamo_db.arn_stream_tle,
        ]
      },
    ]
  })
}
resource "aws_iam_role" "lambda_exec" {
  name               = "serverless_lambda"
  assume_role_policy = jsonencode({
    Version   = "2012-10-17"
    Statement = [
      {
        Action    = "sts:AssumeRole"
        Effect    = "Allow"
        Sid       = ""
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })

}

module "deploy_dynamo_db" {
  source = "./modules/dynamodb"
}
module "deploy_lambdas" {
  source  = "./modules/lambdas"
  iam_arn = aws_iam_role.lambda_exec.arn
}
module "deploy_gateway" {
  source                  = "./modules/gateway"
  lambda_join_lambda_arn  = module.deploy_lambdas.lambda_join_arn
  lambda_join_lambda_name = module.deploy_lambdas.lambda_join_name
}
module "cloud_watch" {
  source  = "./modules/cloud_watch"
  lambda_function_arn = module.deploy_lambdas.lambda_fetch_arn
  lambda_function_name = module.deploy_lambdas.lambda_fetch_name
}
resource "aws_lambda_event_source_mapping" "add_trigger_receive" {
  event_source_arn  = module.deploy_dynamo_db.arn_stream_tle
  function_name     = module.deploy_lambdas.lambda_receive_name
  starting_position = "LATEST"
}
