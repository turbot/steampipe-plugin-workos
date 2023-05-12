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
		Description: "Retrieve information about your groups.",
		List: &plugin.ListConfig{
			ParentHydrate: listDirectories,
			Hydrate:       listGroups,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "directory_id",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getGroup,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The Group's unique identifier.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The Group's name.",
			},
			{
				Name:        "directory_id",
				Type:        proto.ColumnType_STRING,
				Description: "The identifier of the directory the group belongs to.",
				Transform:   transform.FromField("DirectoryID"),
			},
			{
				Name:        "organization_id",
				Type:        proto.ColumnType_STRING,
				Description: "The identifier for the organization in which the directory resides.",
				Transform:   transform.FromField("OrganizationID"),
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The Group's created at date.",
			},
			{
				Name:        "idp_id",
				Type:        proto.ColumnType_STRING,
				Description: "The Group's unique identifier assigned by the directory provider.",
				Transform:   transform.FromField("IdpID"),
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The Group's updated at date.",
			},
			{
				Name:        "raw_attributes",
				Type:        proto.ColumnType_JSON,
				Description: "The Group's raw attributes in raw encoded JSON.",
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

func listGroups(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	dir := h.Item.(directorysync.Directory)
	directory_id := d.EqualsQualString("directory_id")

	// check if the provided directory_id is not matching with the parentHydrate
	if directory_id != "" && directory_id != dir.ID {
		return nil, nil
	}

	apiKey, err := getAPIKey(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("workos_group.listGroups", "connection_error", err)
		return nil, err
	}
	directorysync.SetAPIKey(*apiKey)

	// Limiting the results
	maxLimit := 100
	if d.QueryContext.Limit != nil {
		limit := int(*d.QueryContext.Limit)
		if limit < maxLimit {
			maxLimit = limit
		}
	}

	input := directorysync.ListGroupsOpts{
		Limit:     maxLimit,
		Directory: dir.ID,
	}

	for {
		groupList, err := directorysync.ListGroups(ctx, input)
		if err != nil {
			plugin.Logger(ctx).Error("workos_group.listGroups", "api_error", err)
			return nil, err
		}
		for _, group := range groupList.Data {
			d.StreamListItem(ctx, group)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if groupList.ListMetadata.Before != "" {
			input.Before = groupList.ListMetadata.Before
		} else {
			break
		}
	}

	return nil, nil
}

func getGroup(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQualString("id")

	// Check if id is empty.
	if id == "" {
		return nil, nil
	}

	apiKey, err := getAPIKey(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("workos_group.getGroup", "connection_error", err)
		return nil, err
	}
	directorysync.SetAPIKey(*apiKey)

	input := directorysync.GetGroupOpts{
		Group: id,
	}

	group, err := directorysync.GetGroup(ctx, input)
	if err != nil {
		plugin.Logger(ctx).Error("workos_group.getGroup", "api_error", err)
		return nil, err
	}

	return group, nil
}
