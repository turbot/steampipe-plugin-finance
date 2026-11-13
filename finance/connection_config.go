package finance

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type financeConfig struct {
	IEXAPIToken *string `hcl:"iex_api_token,optional"`
}

func ConfigInstance() interface{} {
	return &financeConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) financeConfig {
	if connection == nil || connection.GetConfig() == nil {
		return financeConfig{}
	}
	config, _ := connection.GetConfig().(financeConfig)
	return config
}
