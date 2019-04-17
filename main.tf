provider "aws" {
  region = "ap-southeast-2"
}

data "cloudconformity_external_id" "it" {}

data "aws_iam_policy_document" "assume" {
  statement {
    effect = "Allow"

    principals {
      type        = "AWS"
      identifiers = ["arn:aws:iam::717210094962:root"]
    }

    actions = ["sts:AssumeRole"]

    condition {
      test     = "StringEquals"
      variable = "sts:ExternalId"
      values   = ["${data.cloudconformity_external_id.it.id}"]
    }
  }
}

resource "aws_iam_role" "role" {
  name                 = "nc-test-role"
  assume_role_policy   = "${data.aws_iam_policy_document.assume.json}"
  permissions_boundary = "arn:aws:iam::566134440840:policy/managed-permission-boundary"
}

resource "aws_iam_role_policy_attachment" "readonly" {
  role       = "${aws_iam_role.role.name}"
  policy_arn = "arn:aws:iam::aws:policy/ReadOnlyAccess"
}

resource "null_resource" "delay" {
  provisioner "local-exec" {
    command = "sleep 10"
  }
  triggers = {
    "before" = "${aws_iam_role.role.id}"
  }
}

resource "cloudconformity_account" "test" {
  name        = "test-account"
  environment = "test"
  role_arn    = "${aws_iam_role.role.arn}"
  external_id = "${data.cloudconformity_external_id.it.id}"
  depends_on = ["null_resource.delay"]
}
