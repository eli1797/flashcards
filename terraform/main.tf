terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.48.0"
    }
    archive = {
      source  = "hashicorp/archive"
      version = "~> 2.2.0"
    }
  }

  required_version = "~> 1.0"
}

provider "aws" {
  region = "us-east-2"
}

# resource "aws_lambda_function" "api_cards" {
#   function_name = "api.cards"

#   s3_bucket = "go-code-bucket"
#   s3_key    = aws_s3_bucket_object.lambda_go_bucket.key

#   runtime = "go1.x"
#   handler = "bin/main"

#     publish = true

#     environment {
#         variables = {
#             ENVIRONMENT = var.env
#         }
#     }

#     ## source_code_hash = filebase64sha256("lambda_function_payload.zip")

#   source_code_hash = data.archive_file.lambda_go_code.output_base64sha256

#   role = aws_iam_role.lambda_exec.arn
# }

# resource "aws_cloudwatch_log_group" "api_cards" {
#   name = "/aws/lambda/${aws_lambda_function.api_cards.function_name}"

#   retention_in_days = 30
# }

# resource "aws_iam_role" "lambda_exec" {
#   name = "serverless_lambda"

#   assume_role_policy = jsonencode({
#     Version = "2012-10-17"
#     Statement = [{
#       Action = "sts:AssumeRole"
#       Effect = "Allow"
#       Sid    = ""
#       Principal = {
#         Service = "lambda.amazonaws.com"
#       }
#       }
#     ]
#   })
# }

# resource "aws_iam_role_policy_attachment" "lambda_policy" {
#   role       = aws_iam_role.lambda_exec.name
#   policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
# }
