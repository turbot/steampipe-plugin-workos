package workos

import (
	"context"
	"errors"
	"os"

	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type workosConfig struct {
	APIKey *string `hcl:"api_key"`
}

func ConfigInstance() interface{} {
	return &workosConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) workosConfig {
	if connection == nil || connection.GetConfig() == nil {
		return workosConfig{}
	}
	config, _ := connection.GetConfig().(workosConfig)
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

	return nil, errors.New("'api_key' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
}
