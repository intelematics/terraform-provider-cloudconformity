package cloud_conformity

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAccGetExternalId(t *testing.T) {
	name := "data.cloudconformity_external_id.id"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: `data "cloudconformity_external_id" "id" {}`,
				Check: resource.TestCheckResourceAttrSet(name, "id"),
			},
		},
	})
}