provider "aws" {
  region = "ap-southeast-2"
}

terraform {
  backend "s3" {
    bucket = "bespin-ci-cd"
    key    = "terraform-provider-cloudconformity-codebuild"
    region = "ap-southeast-2"
  }
}
