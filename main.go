package main

import (
	"github.com/badstreff/terraform-provider-tower/tower"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: tower.Provider,
	})
}
