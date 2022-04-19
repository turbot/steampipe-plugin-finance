package finance

import (
	"context"

	"github.com/piquette/edgr/core"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func tableFinanceUsSecPublicCompany(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "finance_us_sec_public_company",
		Description: "US public companies from the SEC Edgar database.",
		List: &plugin.ListConfig{
			Hydrate: listUsSecPublicCompany,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the company."},
			{Name: "symbol", Type: proto.ColumnType_STRING, Description: "Symbol of the company."},
		},
	}
}

func listUsSecPublicCompany(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	companies, err := core.GetPublicCompaniesWithHeaders(map[string]string{"User-Agent": "Steampipe/v0.x"})
	if err != nil {
		plugin.Logger(ctx).Error("finance_us_sec_public_company.listUsSecPublicCompany", "query_error", err)
		return nil, err
	}
	for _, c := range companies {
		d.StreamListItem(ctx, c)
	}
	return nil, nil
}
