package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/pauloconnor/terraform-provider-consulacl/consulacl"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: consulacl.Provider})
}
