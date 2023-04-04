package workos

import (
	"context"
	"errors"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type workosConfig struct {
	APIKey *string `cty:"api_key"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"api_key": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &workosConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) workosConfig {
	if connection == nil || connection.Config == nil {
		return workosConfig{}
	}
	config, _ := connection.Config.(workosConfig)
	return config
}

func getAPIKey(ctx context.Context, d *plugin.QueryData) (*string, error) {
	workosConfig := GetConfig(d.Connection)

	if workosConfig.APIKey != nil {
		return workosConfig.APIKey, nil
	}

	return nil, errors.New("'api_key' or ('email' and 'password') must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
}
