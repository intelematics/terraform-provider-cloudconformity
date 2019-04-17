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
		Update: resourceAccountUpdate,

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
			"real_time_monitoring": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"cost_package": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"security_package": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
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
		Name:                  d.Get("name").(string),
		Environment:           d.Get("environment").(string),
		Role:                  d.Get("role_arn").(string),
		ExternalId:            d.Get("external_id").(string),
		HasRealTimeMonitoring: d.Get("real_time_monitoring").(bool),
		CostPackage:           d.Get("cost_package").(bool),
		SecurityPackage:       d.Get("security_package").(bool),
	}
	account, err := client.CreateAccount(accountRequest)
	if err != nil {
		return err
	}
	d.SetId(account)
	return resourceAccountRead(d, meta)
}

func resourceAccountRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	account, err := client.GetAccount(d.Id())
	if err != nil {
		return err
	}

	settings, err := client.GetAccountAccessSettings(d.Id())
	if err != nil {
		return err
	}

	_ = d.Set("name", account.Name)
	_ = d.Set("environment", account.Environment)
	_ = d.Set("real_time_monitoring", account.HasRealTimeMonitoring)
	_ = d.Set("cost_package", account.CostPackage)
	_ = d.Set("security_package", account.SecurityPackage)

	_ = d.Set("role_arn", settings.RoleArn)
	_ = d.Set("external_id", settings.ExternalId)
	_ = d.Set("account_id", account.Id)

	return nil
}

func resourceAccountDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	err := client.DeleteAccount(d.Get("account_id").(string))
	return err
}

func resourceAccountExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(*sdk.Client)
	return client.DoesAccountExist(d.Id())
}

func resourceAccountUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)

	if d.HasChange("cost_package") || d.HasChange("real_time_monitoring") || d.HasChange("security_package") {
		cost := d.Get("cost_package").(bool)
		realTime := d.Get("real_time_monitoring").(bool)
		security := d.Get("security_package").(bool)
		err := client.UpdateAccountSubscription(d.Id(), cost, realTime, security)
		if err != nil {
			return err
		}
	}

	return nil
}
