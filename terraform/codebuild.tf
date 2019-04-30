module "codebuild-terraform-provider" {
  source                        = "git@github.com:intelematics/bespin-ci-cd.git?ref=feat/first-commit//terraform/modules/codebuild"
  github_auth_token             = "${var.github_auth_token}"
  codebuild_project_name        = "terraform-provider-cloudconformity"
  codebuild_project_description = "Terraform Cloud Conformity Provider"
  github_repository             = "https://github.com/intelematics/terraform-provider-cloudconformity.git"
  codebuild_role_name           = "terraform-provider-cloudconformity-role"
}

data "aws_iam_policy_document" "ecr-policy-document" {
  statement {
    sid    = "AccessECR"
    effect = "Allow"

    actions = [
      "ecr:GetDownloadUrlForLayer",
      "ecr:BatchGetImage",
      "ecr:BatchCheckLayerAvailability",
      "ecr:PutImage",
      "ecr:InitiateLayerUpload",
      "ecr:UploadLayerPart",
      "ecr:CompleteLayerUpload",
    ]

    resources = [
      "*",
    ]
  }

  statement {
    sid    = "ecrAuthorization"
    effect = "Allow"

    actions = [
      "ecr:GetAuthorizationToken",
    ]

    resources = [
      "*",
    ]
  }

  statement {
    sid    = "ecsAccess"
    effect = "Allow"

    actions = [
      "ecs:RegisterTaskDefinition",
      "ecs:DescribeTaskDefinition",
      "ecs:DescribeServices",
      "ecs:CreateService",
      "ecs:ListServices",
      "ecs:UpdateService",
    ]

    resources = [
      "*",
    ]
  }
}

resource "aws_iam_policy" "ecr-policy" {
  name   = "codebuild-ecr-policy"
  policy = "${data.aws_iam_policy_document.ecr-policy-document.json}"
}

resource "aws_iam_policy_attachment" "attach" {
  name = "codebuild-policy-attachment"

  roles = [
    "terraform-provider-cloudconformity-role",
  ]

  policy_arn = "${aws_iam_policy.ecr-policy.arn}"
}
