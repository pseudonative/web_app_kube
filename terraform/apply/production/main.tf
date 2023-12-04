module "crom_vpc" {
  source = "../../modules/vpc"
  providers = {
    aws = aws.productionuseast1
  }
}