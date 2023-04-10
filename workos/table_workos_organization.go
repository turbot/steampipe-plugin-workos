package workos

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"github.com/workos/workos-go/v2/pkg/organizations"
)

func tableWorkOSOrganization(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "workos_organization",
		Description: "Retrieve information about your organizations.",
		List: &plugin.ListConfig{
			Hydrate: listOrganizations,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getOrganization,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The Organization's unique identifier.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The Organization's name.",
			},
			{
				Name:        "allow_profiles_outside_organization",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether Connections within the Organization allow profiles that are outside of the Organization's configured User Email Domains.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The timestamp of when the Organization was created.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The timestamp of when the Organization was updated.",
			},
			{
				Name:        "domains",
				Type:        proto.ColumnType_JSON,
				Description: "The Organization's Domains.",
			},
		},
	}
}

func listOrganizations(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	apiKey, err := getAPIKey(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("workos_organization.listOrganizations", "connection_error", err)
		return nil, err
	}
	organizations.SetAPIKey(*apiKey)

	// Limiting the results
	maxLimit := 100
	if d.QueryContext.Limit != nil {
		limit := int(*d.QueryContext.Limit)
		if limit < maxLimit {
			maxLimit = limit
		}
	}

	input := organizations.ListOrganizationsOpts{
		Limit: maxLimit,
	}

	for {
		orgList, err := organizations.ListOrganizations(ctx, input)
		if err != nil {
			plugin.Logger(ctx).Error("workos_organization.listOrganizations", "api_error", err)
			return nil, err
		}
		for _, org := range orgList.Data {
			d.StreamListItem(ctx, org)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if orgList.ListMetadata.Before != "" {
			input.Before = orgList.ListMetadata.Before
		} else {
			break
		}
	}

	return nil, nil
}

func getOrganization(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQuals["id"].GetStringValue()

	// Check if id is empty.
	if id == "" {
		return nil, nil
	}

	apiKey, err := getAPIKey(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("workos_organization.getOrganization", "connection_error", err)
		return nil, err
	}
	organizations.SetAPIKey(*apiKey)

	input := organizations.GetOrganizationOpts{
		Organization: id,
	}

	org, err := organizations.GetOrganization(ctx, input)
	if err != nil {
		plugin.Logger(ctx).Error("workos_organization.getOrganization", "api_error", err)
		return nil, err
	}

	return org, nil
}
