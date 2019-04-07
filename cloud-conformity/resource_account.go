package cloud_conformity

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/intelematics/terraform-provider-cloudconformity/sdk"
)

func resourceAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccountCreate,
		Read:   resourceAccountRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"environment": {
				Type:     schema.TypeString,
				Required: true,
			},
			"role_arn": {
				Type:     schema.TypeString,
				Required: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAccountCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	accountRequest := sdk.CreateAccountRequest{
		Name:        d.Get("id").(string),
		Environment: d.Get("environment").(string),
		Role:        d.Get("role_arn").(string),
		ExternalId:  d.Get("external_id").(string),
	}
	account, err := client.CreateAccount(accountRequest)
	if err != nil {
		return err
	}
	d.SetId(account.Id)
	return resourceAccountRead(d, meta)
}

func resourceAccountRead(d *schema.ResourceData, meta interface{}) error {
	// TODO: Add other properties.
	client := meta.(*sdk.Client)
	request := sdk.CreateAccountRequest{

	}
	account, err := client.CreateAccount(request)
	if err != nil {
		return err
	}

	_ := d.Set("id", account.Id)

	return nil
}
