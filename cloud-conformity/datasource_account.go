package cloud_conformity

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/intelematics/terraform-provider-cloudconformity/sdk"
)

func dataSourceAccount() *schema.Resource {
	return &schema.Resource{
		Read: dataAccountRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataAccountRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	accountId := d.Get("id").(string)
	account, err := client.GetAccount(accountId)

	if err != nil {
		return fmt.Errorf("error finding application %q: %s", accountId, err)
	}

	d.SetId(account.Id)
	_ = d.Set("name", account.Name)
	_ = d.Set("id", account.Id)

	return nil
}
