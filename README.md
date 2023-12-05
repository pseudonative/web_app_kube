# Terraform AWS Infrastructure Management

This repository contains the Terraform code and CI/CD workflows for managing AWS infrastructure for a web application. It leverages GitHub Actions for automation, ensuring consistent and reliable infrastructure deployment and management.

## Overview

The project uses Terraform for infrastructure as code to provision and manage AWS resources. GitHub Actions are used for continuous integration and deployment, handling Terraform plans and applies, and managing AWS credentials securely.

### Key Components

- **Terraform**: Infrastructure as Code to manage AWS resources.
- **GitHub Actions**: Automation of CI/CD pipeline.
- **AWS**: Cloud infrastructure provider.

## Prerequisites

- AWS Account with administrative access.
- GitHub Account.
- Basic understanding of Terraform and AWS.

## Setup

1. **AWS Credentials**: Ensure AWS credentials are securely stored in GitHub Secrets.
2. **GitHub Repository**: Fork or clone this repository to your GitHub account.

## Workflows

### 1. WebApp CI/CD (`web_app_cicd.yaml`)

Triggered on pushes and pull requests to `staging-**` branches. It includes steps for setting AWS credentials, checking out the code, and preparing the build.

### 2. Terraform Plan (`terraform_plan.yaml`)

Triggered by comments on pull requests. It performs a Terraform plan to show potential changes without applying them.

### 3. Terraform Apply (`terraform_apply.yaml`)

Triggered by specific issue comments. It applies the Terraform plan to the AWS environment.

### Reusable Workflows and Actions

- **Authentication and Initial Checkout Workflow**: Common steps for AWS authentication and initial repo checkout.
- **Composite Actions**: For downloading artifacts and setting up the environment.

## Usage

- To trigger a Terraform plan, comment `/rezboy_plan` on a pull request.
- To apply the Terraform plan, comment `/lgtm_apply` on the issue linked to the pull request.

## Contributing

Contributions to this project are welcome. Please adhere to the following guidelines:
- Fork the repository.
- Create a new branch for each feature or improvement.
- Submit a pull request for review.

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgments

Thanks to all contributors and maintainers of this project.

---

For more details on each workflow and action, refer to the respective YAML files in the `.github/workflows/` directory.



# Terraform AWS VPC Infrastructure

This repository contains Terraform code for creating and managing an AWS VPC infrastructure. The setup is designed to be modular and reusable, facilitating ease of deployment and management of cloud resources.

## Project Structure

- `terraform/`: Root directory for Terraform configurations.
  - `apply/production/`: Contains Terraform configurations for the production environment.
    - `main.tf`: Main Terraform file with resource definitions.
    - `variables.tf`: Variable definitions for Terraform.
    - `output.tf`: Output definitions for Terraform.
    - `version.tf`: Terraform version and provider settings.
  - `modules/vpc/`: Module for AWS VPC creation and configuration.
    - `main.tf`: VPC resource definitions.
    - `variables.tf`: Variable definitions for the VPC module.
    - `output.tf`: Output definitions for the VPC module.
    - `locals.tf`: Local values for the VPC module.
    - `version.tf`: Terraform version and provider settings for the module.

## AWS Resources Managed

- VPC
- Subnets (Public and Private)
- Route Tables and Associations
- Internet Gateway
- NAT Gateway
- Elastic IPs
- Default Network ACLs
- Default Security Groups

## Pre-requisites

- AWS Account with necessary permissions.
- Terraform installed on your local machine.
- Basic understanding of Terraform and AWS.

## Usage

1. Clone the repository to your local machine.
2. Navigate to the `terraform/apply/production` directory.
3. Initialize Terraform:
   ```sh
   terraform init
