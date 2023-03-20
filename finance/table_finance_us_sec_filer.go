package finance

import (
	"context"

	"github.com/piquette/edgr/core"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableFinanceUsSecFiler(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "finance_us_sec_filer",
		Description: "Lookup company filer details from the US SEC Edgar database.",
		List: &plugin.ListConfig{
			Hydrate:    listUsSecFiler,
			KeyColumns: plugin.SingleColumn("symbol"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "symbol", Type: proto.ColumnType_STRING, Description: "Symbol for the filer."},
			{Name: "cik", Type: proto.ColumnType_STRING, Transform: transform.FromField("CIK"), Description: "CIK (Central Index Key) of the filer."},
			// Other columns
			// Always null? - {Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the filer."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the filer."},
			{Name: "sic", Type: proto.ColumnType_STRING, Transform: transform.FromField("SIC"), Description: "SIC (Standard Industrial Classification) of the filer."},
			{Name: "sic_description", Type: proto.ColumnType_STRING, Transform: transform.FromField("SICDescription"), Description: "Description of the SIC (Standard Industrial Classification) of the filer."},
		},
	}
}

func listUsSecFiler(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.EqualsQuals
	symbol := quals["symbol"].GetStringValue()
	filer, err := core.GetFilerWithHeaders(symbol, map[string]string{"User-Agent": "Steampipe/v0.x"})
	if err != nil {
		plugin.Logger(ctx).Error("finance_us_sec_filer.listUsSecFiler", "query_error", err)
		return nil, err
	}
	d.StreamListItem(ctx, filer)
	return nil, nil
}
