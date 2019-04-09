package cloud_conformity

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/intelematics/terraform-provider-cloudconformity/sdk"
)

func resourceAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccountCreate,
		Read:   resourceAccountRead,
		Delete: resourceAccountDelete,
		Exists: resourceAccountExists,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"environment": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"role_arn": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"real_time_monitoring": {
				Type: schema.TypeBool,
				Optional: true,
				Default: true,
				ForceNew: true,
			},
			"cost_package": {
				Type: schema.TypeBool,
				Optional: true,
				Default: true,
				ForceNew: true,
			},
			"account_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAccountCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	accountRequest := sdk.CreateAccountRequest{
		Name:        d.Get("name").(string),
		Environment: d.Get("environment").(string),
		Role:        d.Get("role_arn").(string),
		ExternalId:  d.Get("external_id").(string),
		HasRealTimeMonitoring: d.Get("real_time_monitoring").(bool),
		CostPackage: d.Get("cost_package").(bool),
	}
	account, err := client.CreateAccount(accountRequest)
	if err != nil {
		return err
	}
	d.SetId(account.Id)
	return resourceAccountRead(d, meta)
}

func resourceAccountRead(d *schema.ResourceData, meta interface{}) error {
	// TODO: Add properties.

	return nil
}

func resourceAccountDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	err := client.DeleteAccount(d.Get("account_id").(string))
	return err
}

func resourceAccountExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	return false, nil
}
