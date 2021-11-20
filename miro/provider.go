package miro

import (
	"context"

	"github.com/Miro-Ecosystem/go-miro/miro"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_token": {
				Type:        schema.TypeString,
				Description: "Access key for Miro API",
				Required:    true,
				Sensitive:   true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"miro_board": resourceBoard(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigureFunc,
	}
}

func providerConfigureFunc(_ context.Context, data *schema.ResourceData) (interface{}, diag.Diagnostics) {
	key := data.Get("access_token").(string)
	var diags diag.Diagnostics
	return miro.NewClient(key), diags
}
