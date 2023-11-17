package finance

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type financeConfig struct {
	IEXAPIToken *string `cty:"iex_api_token,optional"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"iex_api_token": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &financeConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) financeConfig {
	if connection == nil || connection.Config == nil {
		return financeConfig{}
	}
	config, _ := connection.Config.(financeConfig)
	return config
}
