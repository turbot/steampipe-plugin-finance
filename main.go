package main

import (
	"github.com/turbot/steampipe-plugin-finance/finance"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: finance.Plugin})
}
