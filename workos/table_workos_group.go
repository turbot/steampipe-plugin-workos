package workos

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"github.com/workos/workos-go/v2/pkg/directorysync"
)

func tableWorkOSGroup(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "workos_group",
		Description: "Retrieve information about your Groups.",
		List: &plugin.ListConfig{
			ParentHydrate: listDirectories,
			Hydrate:       listGroups,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getGroup,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "Directory unique identifier.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "Directory name.",
			},
			{
				Name:        "state",
				Type:        proto.ColumnType_STRING,
				Description: "Linked status for the Directory.",
			},
			{
				Name:        "organization_id",
				Type:        proto.ColumnType_STRING,
				Description: "Identifier for the Directory's Organization.",
				Transform:   transform.FromField("OrganizationID"),
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The timestamp of when the Directory was created.",
			},
			{
				Name:        "domain",
				Type:        proto.ColumnType_STRING,
				Description: "Directory domain.",
			},
			{
				Name:        "external_key",
				Type:        proto.ColumnType_STRING,
				Description: "Externally used identifier for the Directory.",
			},
			{
				Name:        "idp_id",
				Type:        proto.ColumnType_STRING,
				Description: "The user's directory provider's Identifier.",
				Transform:   transform.FromField("IdpID"),
			},
			{
				Name:        "type",
				Type:        proto.ColumnType_STRING,
				Description: "Type of the directory.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The timestamp of when the Directory was updated.",
			},
		},
	}
}

func listGroups(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	dir := h.Item.(directorysync.Directory)

	apiKey, err := getAPIKey(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listGroups", "connection_error", err)
		return nil, err
	}
	directorysync.SetAPIKey(*apiKey)

	// Limiting the results
	maxLimit := 100
	if d.QueryContext.Limit != nil {
		limit := int(*d.QueryContext.Limit)
		if limit < maxLimit {
			if limit < 1 {
				maxLimit = 1
			} else {
				maxLimit = limit
			}
		}
	}

	input := directorysync.ListGroupsOpts{
		Limit:     maxLimit,
		Directory: dir.ID,
	}

	for {
		dirList, err := directorysync.ListGroups(ctx, input)
		if err != nil {
			plugin.Logger(ctx).Error("listGroups", "api_error", err)
			return nil, err
		}
		for _, dir := range dirList.Data {
			d.StreamListItem(ctx, dir)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if dirList.ListMetadata.Before != "" {
			input.Before = dirList.ListMetadata.Before
		} else {
			break
		}
	}

	return nil, nil
}

func getGroup(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQuals["id"].GetStringValue()

	// Check if id is empty.
	if id == "" {
		return nil, nil
	}

	apiKey, err := getAPIKey(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getGroup", "connection_error", err)
		return nil, err
	}

	directorysync.SetAPIKey(*apiKey)
	input := directorysync.GetGroupOpts{
		Group: id,
	}
	dir, err := directorysync.GetGroup(ctx, input)
	if err != nil {
		plugin.Logger(ctx).Error("getGroup", "api_error", err)
		return nil, err
	}

	return dir, nil
}
