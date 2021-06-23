package finance

import (
	"context"

	"github.com/macleanpinto/yahoo-finance/earning"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableFinanceEarningHistory(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "finance_earning_history",
		Description: "Recent four quarter earing history.",
		List: &plugin.ListConfig{
			Hydrate:    GetEarningHistory,
			KeyColumns: plugin.SingleColumn("symbol"),
		},
		Columns: []*plugin.Column{
			{Name: "symbol", Type: proto.ColumnType_STRING, Hydrate: symbolString, Transform: transform.FromValue(), Description: "Symbol to quote."},
			{Name: "quarter", Type: proto.ColumnType_STRING, Transform: transform.FromField("Quarter.Fmt"), Description: "Quaterly earnings announcement date."},
			{Name: "eps_estimate", Type: proto.ColumnType_STRING, Transform: transform.FromField("EpsEstimate.Raw"), Description: "Estimated earnings per share."},
			{Name: "eps_actual", Type: proto.ColumnType_STRING, Transform: transform.FromField("EpsActual.Fmt"), Description: "Actual earnings per share"},
			{Name: "eps_difference", Type: proto.ColumnType_STRING, Transform: transform.FromField("EpsDifference.Fmt"), Description: "Difference between actual and estimated earning per share. "},
			{Name: "eps_suprise_percent", Type: proto.ColumnType_STRING, Transform: transform.FromField("SurprisePercent.Fmt"), Description: "Earning per share suprise in percent. "},
		},
	}
}

func GetEarningHistory(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	symbol := quals["symbol"].GetStringValue()

	q, err := earning.GetEarningHistory(symbol)
	if err != nil {
		plugin.Logger(ctx).Error("finance_earning_history.GetEarningHistory", "query_error", err)
		return nil, err
	}
	for _, v := range q {
		d.StreamListItem(ctx, v)
	}
	return nil, nil
}
