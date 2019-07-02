module "codebuild-terraform-provider" {
  source                         = "git@github.com:intelematics/infrastructure-ci-cd.git//terraform/modules/codebuild?ref=feat/BE-322/slack-notif"
  github_auth_token              = var.github_auth_token
  codebuild_project_name         = "terraform-provider-cloudconformity"
  codebuild_project_description  = "Terraform Cloud Conformity Provider"
  github_repository              = "https://github.com/intelematics/terraform-provider-cloudconformity.git"
  codebuild_role_name            = "terraform-provider-cloudconformity-role"
  slack_notification_lambda_name = "bespinbuildstatus_codebuild_alert_slack_bot"
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
  policy = data.aws_iam_policy_document.ecr-policy-document.json
}

resource "aws_iam_policy_attachment" "attach" {
  name = "codebuild-policy-attachment"

  roles = [
    "terraform-provider-cloudconformity-role",
  ]

  policy_arn = aws_iam_policy.ecr-policy.arn
}

resource "aws_ecr_repository" "terraform-provider-cloudconformity" {
  name = "terraform-provider-cloudconformity"
}

data "aws_iam_policy_document" "ecr_permission_policy_doc" {
  statement {
    sid    = "AllowPull"
    effect = "Allow"

    principals {
      identifiers = ["*"]
      type        = "*"
    }

    actions = ["ecr:GetDownloadUrlForLayer",
      "ecr:BatchGetImage",
      "ecr:BatchCheckLayerAvailability",
    ]
  }
}

resource "aws_ecr_repository_policy" "ecr_permission_policy" {
  policy     = data.aws_iam_policy_document.ecr_permission_policy_doc.json
  repository = aws_ecr_repository.terraform-provider-cloudconformity.name
}

resource "aws_ecr_lifecycle_policy" "lifecycle-policy" {
  repository = aws_ecr_repository.terraform-provider-cloudconformity.name

  policy = <<EOF
{
    "rules": [
        {
            "rulePriority": 1,
            "description": "Keep last 30 images",
            "selection": {
                "tagStatus": "any",
                "countType": "imageCountMoreThan",
                "countNumber": 30
            },
            "action": {
                "type": "expire"
            }
        }
    ]
}
EOF
}
