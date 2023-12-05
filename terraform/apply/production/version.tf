terraform {
  required_version = ">= 1.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.0"
    }
  }
}

provider "aws" {
  alias  = "productionuseast1"
  region = "us-east-1"

  assume_role {
    role_arn     = "arn:aws:iam::520291287938:role/ProdFullAccess"
    session_name = "Terraform-ProdUpdate"
  }
}


terraform {
  required_version = ">= 1.5.0"
  backend "s3" {
    bucket         = "shoshone-tfstate"
    key            = "shoshonekey/crom"
    region         = "us-east-1"
    kms_key_id     = "arn:aws:kms:us-east-1:520291287938:key/4fc9e509-04c4-4881-89e7-46fb49790093"
    dynamodb_table = "shoshone-state-lock"
  }
}