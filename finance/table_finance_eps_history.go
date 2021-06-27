package finance

import (
	"context"

	"github.com/macleanpinto/yahoo-finance/eps"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableFinanceEpsHistory(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "finance_eps_history",
		Description: "Earning history for the most recent four quarters.",
		List: &plugin.ListConfig{
			Hydrate:    getEpsHistory,
			KeyColumns: plugin.SingleColumn("symbol"),
		},
		Columns: []*plugin.Column{
			{Name: "symbol", Type: proto.ColumnType_STRING, Hydrate: symbolString, Transform: transform.FromValue(), Description: "Symbol to fetch earnings."},
			{Name: "quarter", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Quarter.Fmt").NullIfZero(), Description: "Quaterly eps announcement date."},
			{Name: "estimate", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("EpsEstimate.Raw"), Description: "Estimated earnings per share."},
			{Name: "actual", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("EpsActual.Raw"), Description: "Actual earnings per share."},
			{Name: "difference", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("EpsDifference.Raw"), Description: "Difference between actual and estimated earning per share."},
			{Name: "suprise_percent", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("SurprisePercent.Raw").Transform(convertFloat64ToPercent), Description: "Earning per share suprise as a percentage."},
		},
	}
}

func convertFloat64ToPercent(_ context.Context, input *transform.TransformData) (interface{}, error) {
	f := input.Value.(float64)
	return 100 * f, nil
}

func getEpsHistory(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	symbol := quals["symbol"].GetStringValue()
	q, err := eps.GetEpsHistory(symbol)
	if err != nil {
		plugin.Logger(ctx).Error("finance_eps_history.getEpsHistory", "query_error", q)
		return nil, err
	}
	for _, v := range q {
		d.StreamListItem(ctx, v)
	}
	// if symbol does not exist
	return nil, nil
}
