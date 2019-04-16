package cloud_conformity

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-aws/aws"
	"os"
	"testing"
)

var testAccProvider *schema.Provider

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvidersWithAws map[string]terraform.ResourceProvider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProvidersWithAws = map[string]terraform.ResourceProvider{
		"cloudconformity": testAccProvider,
		"aws":             aws.Provider(),
	}

	testAccProviders = map[string]terraform.ResourceProvider{
		"cloudconformity": testAccProvider,
	}
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("CLOUD_CONFORMITY_API_KEY"); v == "" {
		t.Fatal("CLOUD_CONFORMITY_API_KEY must be set for acceptance tests")
	}
}
