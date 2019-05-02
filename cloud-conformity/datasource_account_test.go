package cloud_conformity

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAccAccountGet(t *testing.T) {
	name := "data.cloudconformity_account.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceAccountConfig("FordSync US"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "id", "HiGVOhutb"),
					resource.TestCheckResourceAttr(name, "name", "FordSync US"),
				),
			},
		},
	})
}

func testDataSourceAccountConfig(accountId string) string {
	return fmt.Sprintf(`
data "cloudconformity_account" "test" {
	name = "%s"
}
`, accountId)
}
