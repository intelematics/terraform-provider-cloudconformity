package cloud_conformity

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"os"
	"testing"
)

var testAccProvider *schema.Provider

var testAccProviders map[string]terraform.ResourceProvider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"cloudconformity": testAccProvider,
	}
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("CLOUD_CONFORMITY_API_KEY"); v == "" {
		t.Fatal("CLOUD_CONFORMITY_API_KEY must be set for acceptance tests")
	}
}
