module "crom_vpc" {
  source = "../../modules/vpc"
  providers = {
    aws = aws.productionuseast1
  }
}

module "crom_eks" {
  source = "../../modules/eks"
  providers = {
    aws = aws.productionuseast1
  }
  vpc_id          = module.crom_vpc.vpc_id
  private_subnets = module.crom_vpc.private_subnets
}
