variable "env" {
  description = "Environment"
  type        = string

  validation {
    condition     = contains(["tst"], var.env)
    error_message = "Not a valid environment."
  }
}