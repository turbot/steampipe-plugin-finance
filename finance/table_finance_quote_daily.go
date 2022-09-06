package finance

import (
	"context"
	"time"

	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func tableFinanceQuoteDaily(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "finance_quote_daily",
		Description: "Daily historical quotes for a given symbol.",
		List: &plugin.ListConfig{
			Hydrate:    listQuoteDaily,
			KeyColumns: plugin.SingleColumn("symbol"),
		},
		Columns: usSecHistoryColumns(),
	}
}

func listQuoteDaily(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	symbol := quals["symbol"].GetStringValue()

	// Daily for 121 months (10 years)
	t := time.Now()
	sixtyOneMonthsAgo := t.AddDate(0, -121, 0)
	params := &chart.Params{
		Symbol:   symbol,
		Start:    datetime.New(&sixtyOneMonthsAgo),
		End:      datetime.New(&t),
		Interval: datetime.OneDay,
	}

	iter := chart.Get(params)

	for iter.Next() {
		b := iter.Bar()
		d.StreamListItem(ctx, b)
	}
	if err := iter.Err(); err != nil {
		plugin.Logger(ctx).Error("finance_quote_daily.listQuoteDaily", "query_error", err)
		return nil, err
	}
	return nil, nil
}
