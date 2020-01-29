package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Config is used to store all provider settings
type Config struct {
	token string
	url   string
}

// type ProviderNetboxClient struct {
// 	configuration Config
// }

// func (c *Config) Client() (interface{}, error) {
// 	cfg := Config{
// 		token: c.token,
// 		url:   c.url,
// 	}

// 	terraformNetboxClient := ProviderNetboxClient{
// 		configuration: cfg,
// 	}

// 	return &terraformNetboxClient, nil
// }

// Provider hello
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"netbox_server": resourceServer(),
		},
		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETBOX_TOKEN", nil),
				Description: "API key used to access Netbox",
			},
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETBOX_URL", nil),
				Description: "Full url including protocol (e.g http://) to a netbox api",
			},
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		token: d.Get("token").(string),
		url:   d.Get("url").(string),
	}
	return &config, nil
}
