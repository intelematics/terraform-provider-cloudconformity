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
					resource.TestCheckResourceAttr(name, "environment", "test"),
					resource.TestCheckResourceAttr(name, "real_time_monitoring", "true"),
					resource.TestCheckResourceAttr(name, "cost_package", "true"),
					resource.TestCheckResourceAttr(name, "security_package", "true"),
					resource.TestCheckResourceAttrPair("cloudconformity_external_id.it", "id", name, "external_id"),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
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

resource "aws_iam_role_policy_attachment" "readonly" {
  role       = "${aws_iam_role.role.name}"
  policy_arn = "arn:aws:iam::aws:policy/ReadOnlyAccess"
}

resource "cloudconformity_account" "test" {
	name = "%s"
	environment = "%s"
	role_arn = "${aws_iam_role.role.arn}"
	external_id = "${data.cloudconformity_external_id.it.id}"
}`, name, environment)
}
