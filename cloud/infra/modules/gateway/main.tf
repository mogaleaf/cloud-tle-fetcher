resource "aws_apigatewayv2_api" "tle_fetcher_solution_gateway" {
  name                       = "tle_fetcher_solution_gateway"
  protocol_type              = "WEBSOCKET"
  route_selection_expression = "$request.body.action"
}

resource "aws_apigatewayv2_stage" "tle_fetcher_solution_gateway" {
  api_id = aws_apigatewayv2_api.tle_fetcher_solution_gateway.id
  name        = "run"
  auto_deploy = true
}

resource "aws_apigatewayv2_integration" "join" {
  api_id           = aws_apigatewayv2_api.tle_fetcher_solution_gateway.id
  integration_type = "AWS_PROXY"

  connection_type           = "INTERNET"
  content_handling_strategy = "CONVERT_TO_TEXT"
  description               = "join lambda"
  integration_method        = "POST"
  integration_uri           = var.lambda_join_lambda_arn

}

resource "aws_apigatewayv2_route" "join" {
  api_id    = aws_apigatewayv2_api.tle_fetcher_solution_gateway.id
  route_key = "$default"
  authorization_type = "NONE"
  target = "integrations/${aws_apigatewayv2_integration.join.id}"
}

resource "aws_apigatewayv2_route_response" "join" {
  api_id             = aws_apigatewayv2_api.tle_fetcher_solution_gateway.id
  route_id           = aws_apigatewayv2_route.join.id
  route_response_key = "$default"
}

resource "aws_apigatewayv2_integration_response" "join" {
  api_id                   = aws_apigatewayv2_api.tle_fetcher_solution_gateway.id
  integration_id           = aws_apigatewayv2_integration.join.id
  integration_response_key = "/200/"
}

resource "aws_lambda_permission" "api_gw" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = var.lambda_join_lambda_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_apigatewayv2_api.tle_fetcher_solution_gateway.execution_arn}/*/*"
}
