package finance

import (
	"context"
	"fmt"

	"github.com/turbot/edgr/core"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
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
			{Name: "cik", Type: proto.ColumnType_STRING, Transform: transform.FromField("CIK"), Description: "Central Index Key (CIK), if available for the company. The CIK is used to identify entities that are regulated by the Securities and Exchange Commission (SEC)."},
			{Name: "currency", Type: proto.ColumnType_STRING, Description: "Currency the symbol is traded in using."},
			{Name: "is_enabled", Type: proto.ColumnType_BOOL, Description: "True if the symbol is enabled for trading on IEX."},
			{Name: "date", Type: proto.ColumnType_TIMESTAMP, Description: "Date the symbol reference data was generated."},
			{Name: "exchange", Type: proto.ColumnType_STRING, Description: "Exchange symbol."},
			{Name: "exchange_name", Type: proto.ColumnType_STRING, Description: "Exchange name."},
			{Name: "exchange_segment", Type: proto.ColumnType_STRING, Description: "Exchange segment."},
			{Name: "exchange_segment_name", Type: proto.ColumnType_STRING, Description: "Exchange segment name."},
			{Name: "exchange_suffix", Type: proto.ColumnType_STRING, Description: "Exchange segment suffix."},
			{Name: "figi", Type: proto.ColumnType_STRING, Transform: transform.FromField("FIGI"), Description: "OpenFIGI id for the security, if available."},
			{Name: "iex_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("IEXID"), Description: "Unique ID applied by IEX to track securities through symbol changes."},
			{Name: "lei", Type: proto.ColumnType_STRING, Transform: transform.FromField("LEI"), Description: "Legal Entity Identifier (LEI) for the security, if available."},
			{Name: "region", Type: proto.ColumnType_STRING, Description: "Country code for the symbol."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "common issue typepossible values are: ad - ADR, cs - Common Stock, cef - Closed End Fund, et - ETF, oef - Open Ended Fund, ps - Preferred Stock, rt - Right, struct - Structured Product, ut - Unit, wi - When Issued, wt - Warrant, empty - Other."},
		},
	}
}

func listUsSecPublicCompany(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	// client := resty.New()

	// We were using the SDK from github.com/piquette/edgr/core to retrieve Security Symbols Reference Data.
	// However, the SDK is currently encountering a Forbidden error when accessing the endpoint https://api.iextrading.com/1.0/ref-data/symbols,
	// despite using a valid API token.
	// We raised a support request (https://github.com/piquette/edgr/issues/3) to address the error; however, it has not yet been resolved.
	// Fortunately, the new API endpoint, as described in the documentation (https://iexcloud.io/docs/core/REF_DATA#security-symbols-reference-data),
	// is now providing the expected data.
	// For further discussions regarding the legacy and new API endpoints, please refer to this discussion thread: https://infinitekind.tenderapp.com/discussions/problems/57171-iex-trading.

	config := GetConfig(d.Connection)
	if config.IEXAPIToken == nil {
		return nil, fmt.Errorf("partial credentials found in connection config, missing: iex_api_token")
	}

	companies, err := core.GetPublicCompaniesWithHeaders(map[string]string{"token": *config.IEXAPIToken}, map[string]string{"User-Agent": "Steampipe/v0.x"})
	if err != nil {
		plugin.Logger(ctx).Error("finance_us_sec_public_company.listUsSecPublicCompany", "query_error", err)
		return nil, err
	}
	for _, c := range companies {
		d.StreamListItem(ctx, c)
	}
	return nil, nil
}
