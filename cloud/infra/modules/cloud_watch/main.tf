resource "aws_cloudwatch_event_rule" "every_ten_minute" {
  name                = "every-ten-minute"
  description         = "Fires every 10 minutes"
  schedule_expression = "rate(10 minutes)"
}

resource "aws_cloudwatch_event_target" "check_ten_minute" {
  rule      = aws_cloudwatch_event_rule.every_ten_minute.name
  target_id = "lambda"
  arn       = var.lambda_function_arn
}

resource "aws_lambda_permission" "allow_cloudwatch_to_call_check_foo" {
  statement_id  = "AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = var.lambda_function_name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.every_ten_minute.arn
}
