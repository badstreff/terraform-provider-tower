package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider - main entrypoint for terraform providers
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{},
	}
}
