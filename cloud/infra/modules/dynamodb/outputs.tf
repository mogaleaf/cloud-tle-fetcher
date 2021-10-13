output "arn_tle" {
  description = "arn instance hostname"
  value       = aws_dynamodb_table.tles.arn
}
output "arn_satellite" {
  description = "arn instance hostname"
  value       = aws_dynamodb_table.satellites.arn
}
output "arn_com" {
  description = "arn instance hostname"
  value       = aws_dynamodb_table.websocket_client_subscription.arn
}

output "arn_stream_tle" {
  description = "arn instance hostname"
  value       = aws_dynamodb_table.tles.stream_arn
}
