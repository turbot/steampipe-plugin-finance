package finance

import (
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/schema"
)

type financeConfig struct {
}

var ConfigSchema = map[string]*schema.Attribute{}

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
