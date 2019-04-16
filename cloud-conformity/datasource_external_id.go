package cloud_conformity

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/intelematics/terraform-provider-cloudconformity/sdk"
)

func dataSourceExternalId() *schema.Resource {
	return &schema.Resource{
		Read: dataExternalIdRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataExternalIdRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)

	externalId, err := client.GetExternalId()

	if err != nil {
		return fmt.Errorf("error finding external id: %s", err)
	}

	d.SetId(externalId)

	return nil
}
