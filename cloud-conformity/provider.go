package cloud_conformity

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/hashicorp/terraform/terraform"
	"github.com/intelematics/terraform-provider-cloudconformity/sdk"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUD_CONFORMITY_API_KEY", ""),
			},
			"region": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "ap-southeast-2",
				ValidateFunc: validation.StringInSlice([]string{"eu-west-1", "ap-southeast-2", "us-west-2"}, true),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"cloudconformity_account": dataSourceAccount(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	apiKey := d.Get("api_key").(string)
	region := d.Get("region").(string)
	return sdk.NewClient(apiKey, region), nil
}
