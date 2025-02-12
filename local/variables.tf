variable "client_id" {
  type        = string
  description = "Specify the provider client_id."
}

variable "client_secret" {
  type        = string
  description = "Specify the provider client_secret."
}

variable "auth_server" {
  type        = string
  description = "Specify the provider auth_server."
  default     = "https://auth.dev.saas-dev.quortex.io"
}

variable "host" {
  type        = string
  description = "Specify the provider host."
  default     = "http://localhost:8000"
}
