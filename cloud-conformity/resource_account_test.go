package cloud_conformity

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAccAccountCreate__Basic(t *testing.T) {
	name := "cloudconformity_account.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersWithAws,
		Steps: []resource.TestStep{
			{
				Config: testAccount("test-account", "test"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", "test-account"),
				),
			},
		},
	})
}

func testAccount(name, environment string) string {
	return fmt.Sprintf(`
data "aws_iam_policy_document" "assume" {
  statement {
    effect = "Allow",
    principals {
      type = "AWS"
      identifiers = ["arn:aws:iam::717210094962:root"]
    }
    actions = ["sts:AssumeRole"]
    condition {
      test = "StringEquals"
      variable = "sts:ExternalId"
      values = ["${data.cloudconformity_external_id.it.id}"]
    }
  }
}

data "cloudconformity_external_id" "it" {}

resource "aws_iam_role" "role" {
  name = "nc-test-role"
  assume_role_policy = "${data.aws_iam_policy_document.assume.json}"
  permissions_boundary = "arn:aws:iam::531491312713:policy/managed-permission-boundary"
}

resource "cloudconformity_account" "test" {
	name = "%s"
	environment = "%s"
	role_arn = "${aws_iam_role.role.arn}"
	external_id = "${data.cloudconformity_external_id.it.id}"
}`, name, environment)
}
