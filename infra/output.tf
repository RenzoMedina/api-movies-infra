output "public_ip" {
  description = "IP pública de lainstancia"
  value       = aws_instance.dev_ec2.public_ip
}
