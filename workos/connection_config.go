package workos

import (
	"context"
	"errors"
	"os"

	"github.com/turbot/go-kit/types"
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

	apiKey := os.Getenv("WORKOS_API_KEY")

	if workosConfig.APIKey != nil {
		apiKey = *workosConfig.APIKey
	}

	if apiKey != "" {
		return types.String(apiKey), nil
	}

	return nil, errors.New("'api_key' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe.")
}
