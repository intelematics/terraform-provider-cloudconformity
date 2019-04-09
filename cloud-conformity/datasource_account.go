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
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataAccountRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	accountName := d.Get("name").(string)
	accounts, err := client.ListAccounts()

	if err != nil {
		return fmt.Errorf("error finding account %s: %s", accountName, err)
	}

	var accountOverview *sdk.AccountOverview

	for _, account := range accounts {
		if accountName == account.Name {
			accountOverview = &account
			break
		}
	}

	if accountOverview == nil {
		return fmt.Errorf("unable to find account %s", accountName)
	}

	d.SetId(accountOverview.Id)
	_ = d.Set("name", accountOverview.Name)
	_ = d.Set("id", accountOverview.Id)

	return nil
}
