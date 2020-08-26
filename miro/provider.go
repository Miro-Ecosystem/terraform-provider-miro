package miro

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_key": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Access key for Miro API",
				Required:    true,
			},
		},
		ResourcesMap:   map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{},
		ConfigureFunc:  providerConfigureFunc,
	}
}

func providerConfigureFunc(data *schema.ResourceData) (interface{}, error) {
	key := data.Get("access_key").(string)
	_ = key
	return nil, nil
}
