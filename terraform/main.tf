terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.48.0"
    }
  }

  required_version = "= 1.0.2"
}

provider "aws" {
  region = "us-east-2"
}

resource "aws_lambda_function" "api_cards" {
  function_name = "apiCards-${var.env}"

  s3_bucket = "go-code-bucket"
  s3_key    = "${var.env}-code"

  runtime = "go1.x"
  handler = "bin/main"

  publish = true
  timeout = 10

  environment {
    variables = {
      ENVIRONMENT = var.env
    }
  }

  source_code_hash = data.aws_s3_bucket_object.api_cards_source_code_hash.body

  role = aws_iam_role.lambda_exec.arn
}

data "aws_s3_bucket_object" "api_cards_source_code_hash" {
  bucket = "go-code-bucket"
  key    = "${var.env}-code-sha256"
}

resource "aws_cloudwatch_log_group" "api_cards_cw" {
  name = "/aws/lambda/${aws_lambda_function.api_cards.function_name}"

  retention_in_days = 7
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

# api.gateway

resource "aws_api_gateway_rest_api" "api_gw" {
  name = "cards_api_gw"
}

resource "aws_api_gateway_resource" "proxy" {
  rest_api_id = aws_api_gateway_rest_api.api_gw.id
  parent_id   = aws_api_gateway_rest_api.api_gw.root_resource_id
  path_part   = "{proxy+}"
}

resource "aws_api_gateway_method" "proxyMethod" {
  rest_api_id   = aws_api_gateway_rest_api.api_gw.id
  resource_id   = aws_api_gateway_resource.proxy.id
  http_method   = "ANY"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "gw_integrate" {
  rest_api_id = aws_api_gateway_rest_api.api_gw.id
  resource_id = aws_api_gateway_method.proxyMethod.resource_id
  http_method = aws_api_gateway_method.proxyMethod.http_method

  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.api_cards.invoke_arn
}

resource "aws_api_gateway_deployment" "apideploy" {
  depends_on = [
    aws_api_gateway_integration.gw_integrate,
  ]

  rest_api_id = aws_api_gateway_rest_api.api_gw.id
  stage_name  = var.env
}


resource "aws_lambda_permission" "apigw" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.api_cards.function_name
  principal     = "apigateway.amazonaws.com"

  # The "/*/*" portion grants access from any method on any resource
  # within the API Gateway REST API.
  source_arn = "${aws_api_gateway_rest_api.api_gw.execution_arn}/*/*"
}
