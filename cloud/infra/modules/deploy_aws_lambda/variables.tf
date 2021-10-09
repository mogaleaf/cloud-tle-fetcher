variable "bucket_name" {
  description = "Name of the s3 bucket. Must be unique."
  type        = string
}

variable "binary_name" {
  description = "Name of the binaries to deploy in the bucket. Must be unique."
  type        = string
}

variable "binary_source_dir" {
  description = "path to find the binary"
  type        = string
}

variable "lambda_function_name" {
  description = "name of the lambda function"
  type        = string
}
