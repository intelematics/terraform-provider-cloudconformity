package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/intelematics/terraform-provider-cloudconformity/cloud-conformity"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: cloud_conformity.Provider})
}
