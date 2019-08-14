Terraform Provider for Cloud Conformity
==================

- Website: https://www.terraform.io

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

To Do
-----------
This repo still contains code used by Intelematics, to run TF with this
provider:
- Dockerfile
- TF code to deploy CodeBuild project (CI/CD) to build docker images & push to ECR (incl. buildspec.yml)

This code should be split out ASAP

License
-----------
This software is Copyright (c) 2019 Intelematics (enquiry@intelematics.com) and released under the MPL 2.0 License.
For details about this license, please see the text of LICENSE.

Maintainers
-----------

This provider plugin is maintained by the Cloud team at [Intelematics](https://www.intelematics.com/), and the Cloud Conformity team at [Cloud Conformity](https://cloudconformity.com)

Requirements
------------

-   [Cloud Conformity Account](https://cloudconformity.com)
-   [Terraform](https://www.terraform.io/downloads.html) 0.10.x
-   [Go](https://golang.org/doc/install) 1.11 (to build the provider plugin)

Usage
---------------------

```
# Requires a Cloud Conformity API Key 
# Generate this by logging in to CC, click your name in the top right then
# `User Settings` -> `API Keys` -> `New API Key`
provider "cloudconformity" {
  api_key = var.cloudconformity_key
}

# Assumes you have Aliases set up for your AWS Accounts
# This is recommended for readability, but otherwise can just use AWS Account name
data "aws_iam_account_alias" "current" {
  depends_on = [aws_iam_account_alias.alias]
}

# Retrieve CloudConformity External ID, for use in IAM trust policy
data "cloudconformity_external_id" "it" {}

# IAM Policy doc allowing CloudConformity External ID to assume a role
data "aws_iam_policy_document" "assume" {
  statement {
    effect = "Allow"

    principals {
      type = "AWS"

      identifiers = [
        "arn:aws:iam::${var.cloudconformity_account_Id}:root",
      ]
    }

    actions = [
      "sts:AssumeRole",
    ]

    condition {
      test     = "StringEquals"
      variable = "sts:ExternalId"

      values = [
        data.cloudconformity_external_id.it.id,
      ]
    }
  }
}

# Create the IAM Role for CloudConformity to assume, using the IAM Policy above
resource "aws_iam_role" "cloud_conformity_role" {
  name                 = "cloud-conformity-role"
  assume_role_policy   = data.aws_iam_policy_document.assume.json
}

resource "aws_iam_role_policy_attachment" "cloud_conformity_role_attach" {
  role       = aws_iam_role.cloud_conformity_role.name
  policy_arn = "arn:aws:iam::aws:policy/ReadOnlyAccess"
}

# Create the actual CloudConformity Account
resource "cloudconformity_account" "cloudconformity_account" {
  name        = data.aws_iam_account_alias.current.account_alias
  environment = "production"
  role_arn    = aws_iam_role.cloud_conformity_role.arn
  external_id = data.cloudconformity_external_id.it.id
}

output "account_id" {
  value = data.aws_iam_account_alias.current.account_alias
}

```

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/terraform-providers/terraform-provider-cloudconformity`

```sh
$ mkdir -p $GOPATH/src/github.com/terraform-providers; cd $GOPATH/src/github.com/terraform-providers
$ git clone git@github.com:terraform-providers/terraform-provider-cloudconformity
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-cloudconformity
$ make build
```

Using the provider
----------------------
## Fill in for each provider

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.12+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-cloudconformity
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

## Go Module Upgrade
The `go.mod` file lists all required module dependencies.  Along with `go.sum`,
it can be removed and recreated with:
```
go mod init
```

This replaces the older `dep` program and places the `vendor` directory in a
centralised location outside the project.

Go will fetch the latest versions of the dependencies to create the files.
**However**, at the time of writing, the latest version of
`terraform-provider-aws` generated compile errors.  It has been pinned to
version 1.23.0 until these are resolved.

Further details about using Go modules can be found
[here](https://blog.golang.org/using-go-modules).  The `go.mod` file will always
contain a reference to the last version retrieved with `go get`.
