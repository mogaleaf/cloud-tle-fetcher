resource "aws_dynamodb_table" "satellites" {
  name         = "Satellite"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "SatelliteName"

  attribute {
    name = "SatelliteName"
    type = "S"
  }

  ttl {
    attribute_name = "Expire"
    enabled        = true
  }
}
resource "aws_dynamodb_table" "tles" {
  name             = "Tle"
  billing_mode     = "PAY_PER_REQUEST"
  hash_key         = "SatelliteLastID"
  stream_enabled   = true
  stream_view_type = "NEW_IMAGE"
  attribute {
    name = "SatelliteLastID"
    type = "S"
  }
  ttl {
    attribute_name = "Expire"
    enabled        = true
  }
}
resource "aws_dynamodb_table" "websocket_client_subscription" {
  name             = "Connection"
  billing_mode     = "PAY_PER_REQUEST"
  hash_key         = "ConnectionID"

  attribute {
    name = "ConnectionID"
    type = "S"
  }
  ttl {
    attribute_name = "Expire"
    enabled        = true
  }
}
