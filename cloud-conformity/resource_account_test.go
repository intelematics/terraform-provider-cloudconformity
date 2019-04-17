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
					resource.TestCheckResourceAttr(name, "external_id", "2b1dc920-3afd-11e9-a137-bbd8fdf89dea"),
				),
			},
		},
	})
}

func testAccount(name, environment string) string {
	return fmt.Sprintf(`
resource "cloudconformity_account" "test" {
	name = "%s"
	environment = "%s"
	role_arn = "arn:aws:iam::566134440840:role/cloud-conformity-role"
	external_id = "2b1dc920-3afd-11e9-a137-bbd8fdf89dea"
}`, name, environment)
}

func testAccount_tofix(name, environment string) string {
	return fmt.Sprintf(`
data "cloudconformity_external_id" "it" {}

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
      values = ["2b1dc920-3afd-11e9-a137-bbd8fdf89dea"]
    }
  }
}

resource "aws_iam_role" "role" {
  name = "nc-test-role"
  assume_role_policy = "${data.aws_iam_policy_document.assume.json}"
  permissions_boundary = "arn:aws:iam::566134440840:policy/managed-permission-boundary"
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
	depends_on = ["aws_iam_role.role","aws_iam_role_policy_attachment.readonly"]
}`, name, environment)
}
