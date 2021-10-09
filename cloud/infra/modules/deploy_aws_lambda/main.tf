resource "aws_s3_bucket" "lambda_bucket" {
  bucket =  var.bucket_name

  acl           = "private"
  force_destroy = true
}

data "archive_file" "zip_binary" {
  type = "zip"

  source_dir  = var.binary_source_dir
  output_path = "${path.module}/bin/${var.binary_name}"
}

resource "aws_s3_bucket_object" "bucket_upload" {
  bucket = aws_s3_bucket.lambda_bucket.id

  key    = var.binary_name
  source = data.archive_file.zip_binary.output_path

  etag = filemd5(data.archive_file.zip_binary.output_path)
}

resource "aws_lambda_function" "function_notification" {
  function_name = var.lambda_function_name

  s3_bucket = aws_s3_bucket.lambda_bucket.id
  s3_key    = aws_s3_bucket_object.bucket_upload.key

  runtime = "go1.x"
  handler = var.lambda_handler

  source_code_hash = data.archive_file.zip_binary.output_base64sha256

  role = aws_iam_role.lambda_exec.arn

  environment {
    variables = var.environment_variables
  }
}

resource "aws_iam_role" "lambda_exec" {
  name = "serverless_lambda"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Sid    = ""
      Principal = {
        Service = "lambda.amazonaws.com"
      }
    }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "lambda_policy" {
  role       = aws_iam_role.lambda_exec.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}
