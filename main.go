package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-workos/workos"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: workos.Plugin})
}
