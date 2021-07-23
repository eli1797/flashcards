output "function_name" {
  description = "Name of the Lambda function."
  value       = aws_lambda_function.api_cards.function_name
}

output "base_url" {
  description = "Base url for invoking the api.gateway trigger."
  value       = aws_api_gateway_deployment.apideploy.invoke_url
}
