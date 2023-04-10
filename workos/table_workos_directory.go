package workos

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"github.com/workos/workos-go/v2/pkg/directorysync"
)

func tableWorkOSDirectory(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "workos_directory",
		Description: "Retrieve information about your directories.",
		List: &plugin.ListConfig{
			Hydrate: listDirectories,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "domain",
					Require: plugin.Optional,
				},
				{
					Name:    "name",
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
			Hydrate:    getDirectory,
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

func listDirectories(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	apiKey, err := getAPIKey(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("workos_directory.listDirectories", "connection_error", err)
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

	input := directorysync.ListDirectoriesOpts{
		Limit: maxLimit,
	}

	if d.EqualsQuals["domain"] != nil {
		input.Domain = d.EqualsQuals["domain"].GetStringValue()
	}
	if d.EqualsQuals["name"] != nil {
		input.Search = d.EqualsQuals["name"].GetStringValue()
	}
	if d.EqualsQuals["organization_id"] != nil {
		input.OrganizationID = d.EqualsQuals["organization_id"].GetStringValue()
	}

	for {
		dirList, err := directorysync.ListDirectories(ctx, input)
		if err != nil {
			plugin.Logger(ctx).Error("workos_directory.listDirectories", "api_error", err)
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

func getDirectory(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQuals["id"].GetStringValue()

	// Check if id is empty.
	if id == "" {
		return nil, nil
	}

	apiKey, err := getAPIKey(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("workos_directory.getDirectory", "connection_error", err)
		return nil, err
	}
	directorysync.SetAPIKey(*apiKey)

	input := directorysync.GetDirectoryOpts{
		Directory: id,
	}

	dir, err := directorysync.GetDirectory(ctx, input)
	if err != nil {
		plugin.Logger(ctx).Error("workos_directory.getDirectory", "api_error", err)
		return nil, err
	}

	return dir, nil
}
