package workos

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// Plugin creates this (workos) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-workos",
		DefaultTransform: transform.FromCamel(),
		DefaultIgnoreConfig: &plugin.IgnoreConfig{
			ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"404"}),
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"workos_connection":   tableWorkOSConnection(ctx),
			"workos_directory":    tableWorkOSDirectory(ctx),
			"workos_group":        tableWorkOSGroup(ctx),
			"workos_organization": tableWorkOSOrganization(ctx),
			"workos_user":         tableWorkOSUser(ctx),
		},
	}
	return p
}
