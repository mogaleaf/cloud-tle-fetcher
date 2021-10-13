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

  role = var.iam_arn
  environment {
    variables = var.environment_variables
  }

}
