package finance

import (
	"context"
	"time"

	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func tableFinanceQuoteHourly(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "finance_quote_hourly",
		Description: "Hourly historical quotes for a given symbol.",
		List: &plugin.ListConfig{
			Hydrate:    listQuoteHourly,
			KeyColumns: plugin.SingleColumn("symbol"),
		},
		Columns: usSecHistoryColumns(),
	}
}

func listQuoteHourly(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	symbol := quals["symbol"].GetStringValue()

	// Hourly for 13 months
	t := time.Now()
	thirteenMonthsAgo := t.AddDate(0, -13, 0)
	params := &chart.Params{
		Symbol:   symbol,
		Start:    datetime.New(&thirteenMonthsAgo),
		End:      datetime.New(&t),
		Interval: datetime.OneHour,
	}

	iter := chart.Get(params)

	for iter.Next() {
		b := iter.Bar()
		d.StreamListItem(ctx, b)
	}
	if err := iter.Err(); err != nil {
		plugin.Logger(ctx).Error("finance_quote_hourly.listQuoteHourly", "query_error", err)
		return nil, err
	}
	return nil, nil
}
