package finance

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type financeConfig struct {
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
