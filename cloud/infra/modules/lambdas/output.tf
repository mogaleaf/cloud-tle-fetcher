output "lambda_join_arn" {
  value = module.deploy_join.invoke_arn
}
output "lambda_receive_name" {
  value = module.deploy_receive.lambda_function_name
}
output "lambda_join_name" {
  value = module.deploy_join.lambda_function_name
}
output "lambda_fetch_arn" {
  value = module.deploy_fetch.arn
}
output "lambda_fetch_name" {
  value = module.deploy_fetch.lambda_function_name
}
