package workos

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"github.com/workos/workos-go/v2/pkg/sso"
)

func tableWorkOSConnection(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "workos_connection",
		Description: "Retrieve information about your connections.",
		List: &plugin.ListConfig{
			Hydrate: listConnections,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "connection_type",
					Require: plugin.Optional,
				},
				{
					Name:    "organization_id",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getConnection,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "Connection unique identifier.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "Connection name.",
			},
			{
				Name:        "state",
				Type:        proto.ColumnType_STRING,
				Description: "Connection linked state.",
			},
			{
				Name:        "organization_id",
				Type:        proto.ColumnType_STRING,
				Description: "Organization ID.",
				Transform:   transform.FromField("OrganizationID"),
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The timestamp of when the Connection was created.",
			},
			{
				Name:        "connection_type",
				Type:        proto.ColumnType_STRING,
				Description: "Connection provider type.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The timestamp of when the Connection was updated.",
			},
			{
				Name:        "domains",
				Type:        proto.ColumnType_JSON,
				Description: "Domain records for the Connection.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

func listConnections(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	apiKey, err := getAPIKey(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("workos_connection.listConnections", "connection_error", err)
		return nil, err
	}
	sso.Configure(*apiKey, "")

	// Limiting the results
	maxLimit := 100
	if d.QueryContext.Limit != nil {
		limit := int(*d.QueryContext.Limit)
		if limit < maxLimit {
			maxLimit = limit
		}
	}

	input := sso.ListConnectionsOpts{
		Limit: maxLimit,
	}

	if d.EqualsQualString("connection_type") != "" {
		input.ConnectionType = sso.ConnectionType(d.EqualsQualString("connection_type"))
	}
	if d.EqualsQualString("organization_id") != "" {
		input.OrganizationID = d.EqualsQualString("organization_id")
	}

	for {
		connections, err := sso.ListConnections(ctx, input)
		if err != nil {
			plugin.Logger(ctx).Error("workos_connection.listConnections", "api_error", err)
			return nil, err
		}
		for _, connection := range connections.Data {
			d.StreamListItem(ctx, connection)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if connections.ListMetadata.Before != "" {
			input.Before = connections.ListMetadata.Before
		} else {
			break
		}
	}

	return nil, nil
}

func getConnection(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQualString("id")

	// Check if id is empty.
	if id == "" {
		return nil, nil
	}

	apiKey, err := getAPIKey(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("workos_connection.getConnection", "connection_error", err)
		return nil, err
	}
	sso.Configure(*apiKey, "")

	input := sso.GetConnectionOpts{
		Connection: id,
	}

	connection, err := sso.GetConnection(ctx, input)
	if err != nil {
		plugin.Logger(ctx).Error("workos_connection.getConnection", "api_error", err)
		return nil, err
	}

	return connection, nil
}
