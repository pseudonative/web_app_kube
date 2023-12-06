resource "aws_db_instance" "postgres_db" {
  allocated_storage      = 20
  storage_type           = "gp2"
  engine                 = "postgres"
  engine_version         = "16.1"
  instance_class         = "db.t3.micro"
  db_name                = var.db_name
  username               = var.db_username
  password               = var.db_password
  db_subnet_group_name   = aws_db_subnet_group.subnet_group.name
  vpc_security_group_ids = [aws_security_group.db_sg.id]

  skip_final_snapshot = true
}

# DB subnet group
resource "aws_db_subnet_group" "subnet_group" {
  name       = "my-db-subnet-group"
  subnet_ids = var.subnet_ids
}

# Security group for RDS
resource "aws_security_group" "db_sg" {
  name        = "db-security-group"
  description = "Database security group"
  vpc_id      = var.vpc_id

  # Allow inbound PostgreSQL connection
  ingress {
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = ["73.78.80.250/32"]
  }
}