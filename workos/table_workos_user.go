package workos

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"github.com/workos/workos-go/v2/pkg/directorysync"
)

func tableWorkOSUser(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "workos_user",
		Description: "Retrieve information about your users.",
		List: &plugin.ListConfig{
			ParentHydrate: listDirectories,
			Hydrate:       listUsers,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "directory_id",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getUser,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The User's unique identifier.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "user_name",
				Type:        proto.ColumnType_STRING,
				Description: "The User's username.",
				Transform:   transform.FromField("Username"),
			},
			{
				Name:        "state",
				Type:        proto.ColumnType_STRING,
				Description: "The User's state.",
			},
			{
				Name:        "directory_id",
				Type:        proto.ColumnType_STRING,
				Description: "The identifier of the Directory the Directory User belongs to.",
				Transform:   transform.FromField("DirectoryID"),
			},
			{
				Name:        "organization_id",
				Type:        proto.ColumnType_STRING,
				Description: "The identifier for the Organization in which the Directory resides.",
				Transform:   transform.FromField("OrganizationID"),
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The User's created at date.",
			},
			{
				Name:        "first_name",
				Type:        proto.ColumnType_STRING,
				Description: "The User's first name.",
			},
			{
				Name:        "idp_id",
				Type:        proto.ColumnType_STRING,
				Description: "The User's unique identifier assigned by the Directory Provider.",
				Transform:   transform.FromField("IdpID"),
			},
			{
				Name:        "job_title",
				Type:        proto.ColumnType_STRING,
				Description: "The User's job title.",
			},
			{
				Name:        "last_name",
				Type:        proto.ColumnType_STRING,
				Description: "The User's last name.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The User's updated at date.",
			},
			{
				Name:        "custom_attributes",
				Type:        proto.ColumnType_JSON,
				Description: "The User's custom attributes in raw encoded JSON.",
			},
			{
				Name:        "emails",
				Type:        proto.ColumnType_JSON,
				Description: "The User's e-mails.",
			},
			{
				Name:        "groups",
				Type:        proto.ColumnType_JSON,
				Description: "The User's groups.",
			},
			{
				Name:        "raw_attributes",
				Type:        proto.ColumnType_JSON,
				Description: "The User's raw attributes in raw encoded JSON.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Username"),
			},
		},
	}
}

func listUsers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	dir := h.Item.(directorysync.Directory)
	directory_id := d.EqualsQualString("directory_id")

	// check if the provided directory_id is not matching with the parentHydrate
	if directory_id != "" && directory_id != dir.ID {
		return nil, nil
	}

	apiKey, err := getAPIKey(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("workos_user.listUsers", "connection_error", err)
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

	input := directorysync.ListUsersOpts{
		Limit:     maxLimit,
		Directory: dir.ID,
	}

	for {
		userList, err := directorysync.ListUsers(ctx, input)
		if err != nil {
			plugin.Logger(ctx).Error("workos_user.listUsers", "api_error", err)
			return nil, err
		}
		for _, user := range userList.Data {
			d.StreamListItem(ctx, user)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if userList.ListMetadata.Before != "" {
			input.Before = userList.ListMetadata.Before
		} else {
			break
		}
	}

	return nil, nil
}

func getUser(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQualString("id")

	// Check if id is empty.
	if id == "" {
		return nil, nil
	}

	apiKey, err := getAPIKey(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("workos_user.getUser", "connection_error", err)
		return nil, err
	}

	directorysync.SetAPIKey(*apiKey)
	input := directorysync.GetUserOpts{
		User: id,
	}

	user, err := directorysync.GetUser(ctx, input)
	if err != nil {
		plugin.Logger(ctx).Error("workos_user.getUser", "api_error", err)
		return nil, err
	}

	return user, nil
}
