variable "aws_region" {
  description = "AWS region where infrastructure will be deployed"
  type        = string
  default     = "ap-south-1"
}

variable "cluster_name" {
  description = "Name of the EKS cluster"
  type        = string
  default     = "student-career-eks"
}

variable "environment" {
  description = "Deployment environment"
  type        = string
  default     = "dev"
}