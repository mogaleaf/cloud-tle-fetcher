output "lambda_bucket_name" {
  description = "Name of the S3 bucket used to store function code."

  value = aws_s3_bucket.lambda_bucket.id
}
output "lambda_function_name" {
  description = "Name of the S3 bucket used to store function code."

  value = aws_lambda_function.function_notification.function_name
}
output "invoke_arn" {
  description = "Name of the S3 bucket used to store function code."

  value = aws_lambda_function.function_notification.invoke_arn
}
output "arn" {
  description = "Name of the S3 bucket used to store function code."

  value = aws_lambda_function.function_notification.arn
}
