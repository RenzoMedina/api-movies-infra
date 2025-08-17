variable "aws_instance_type" {
  description = "The type of AWS instance to create"
  type        = string
  default     = "t2.micro"
}

variable "key_name" {
  description = "The name of the SSH key pair to use for the instance"
  type        = string
  default     = "myvm-demo"
}

variable "aws_region" {
  description = "The AWS region to deploy resources in"
  type        = string
  default     = "us-east-1"
}

variable "aws_instance_name" {
  description = "The name of the AWS instance"
  type        = string
  default     = "VM-DevOps"
}