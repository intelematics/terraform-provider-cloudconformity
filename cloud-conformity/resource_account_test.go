package cloud_conformity

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/intelematics/terraform-provider-cloudconformity/sdk"
	"testing"
)

func TestAccAccountCreate__Basic(t *testing.T) {
	name := "cloudconformity_account.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersWithAws,
		CheckDestroy: testAccountDeleted,
		Steps: []resource.TestStep{
			{
				Config: testAccountSet("test-account", "test"),
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

func TestAccAccountUpdate(t *testing.T) {

	name := "cloudconformity_account.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersWithAws,
		Steps: []resource.TestStep{
			{
				Config: testAccountUpdate("test-account", "test", true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", "test-account"),
					resource.TestCheckResourceAttr(name, "environment", "test"),
					resource.TestCheckResourceAttr(name, "real_time_monitoring", "true"),
					resource.TestCheckResourceAttr(name, "cost_package", "true"),
					resource.TestCheckResourceAttr(name, "security_package", "true"),
					resource.TestCheckResourceAttr(name, "external_id", "2b1dc920-3afd-11e9-a137-bbd8fdf89dea"),
				),
			},
			{
				Config: testAccountUpdate("test-account1", "prod", false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", "test-account1"),
					resource.TestCheckResourceAttr(name, "environment", "prod"),
					resource.TestCheckResourceAttr(name, "real_time_monitoring", "true"),
					resource.TestCheckResourceAttr(name, "cost_package", "false"),
					resource.TestCheckResourceAttr(name, "security_package", "true"),
					resource.TestCheckResourceAttr(name, "external_id", "2b1dc920-3afd-11e9-a137-bbd8fdf89dea"),
				),
			},
		},
	})
}

func testAccountUpdate(name, environment string, cost_package bool) string {
	return fmt.Sprintf(`
resource "cloudconformity_account" "test" {
	name = "%s"
	environment = "%s"
	role_arn = "arn:aws:iam::566134440840:role/cloud-conformity-role"
	cost_package = %t
	external_id = "2b1dc920-3afd-11e9-a137-bbd8fdf89dea"
}`, name, environment, cost_package)
}

func testAccountSet(name, environment string) string {
	return fmt.Sprintf(`
provider "aws" {
  region = "ap-southeast-2"
}

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
      values = ["${data.cloudconformity_external_id.it.id}"]
    }
  }
}

resource "aws_iam_role" "role" {
  name = "cc-role"
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
}`, name, environment)
}

func testAccountDeleted(s *terraform.State) error {

	client := testAccProvider.Meta().(*sdk.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudconformity_account" {
			continue
		}
		_, err := client.GetAccount(rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("account not destroyed, %s ", rs.Primary.ID)
		}
	}

	return nil
}
