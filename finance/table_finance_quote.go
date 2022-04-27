package finance

import (
	"context"

	"github.com/piquette/finance-go/quote"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableFinanceQuote(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "finance_quote",
		Description: "Most recent available quote for the given symbol.",
		List: &plugin.ListConfig{
			Hydrate:    listQuote,
			KeyColumns: plugin.SingleColumn("symbol"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "symbol", Type: proto.ColumnType_STRING, Description: "Symbol to quote."},
			{Name: "short_name", Type: proto.ColumnType_STRING, Description: "Short descriptive name for the entity."},
			{Name: "regular_market_price", Type: proto.ColumnType_DOUBLE, Description: "Price in the regular market."},
			{Name: "regular_market_time", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("RegularMarketTime").Transform(transform.UnixToTimestamp), Description: "Time when the regular market data was updated."},
			// Other columns
			{Name: "ask", Type: proto.ColumnType_DOUBLE, Description: "Ask price. "},
			{Name: "ask_size", Type: proto.ColumnType_DOUBLE, Description: "Ask size."},
			{Name: "average_daily_volume_10_day", Type: proto.ColumnType_INT, Description: "Average daily volume - last 10 days."},
			{Name: "average_daily_volume_3_month", Type: proto.ColumnType_INT, Description: "Average daily volume - last 3 months."},
			{Name: "bid", Type: proto.ColumnType_DOUBLE, Description: "Bid price."},
			{Name: "bid_size", Type: proto.ColumnType_DOUBLE, Description: "Bid size."},
			{Name: "currency_id", Type: proto.ColumnType_STRING, Description: "Currency ID, e.g. AUD, USD."},
			{Name: "exchange_id", Type: proto.ColumnType_STRING, Description: "Exchange ID, e.g. NYQ, CCC."},
			{Name: "exchange_timezone_name", Type: proto.ColumnType_STRING, Description: "Timezone at the exchange."},
			{Name: "exchange_timezone_short_name", Type: proto.ColumnType_STRING, Description: "Timezone short name at the exchange."},
			{Name: "fifty_day_average", Type: proto.ColumnType_DOUBLE, Description: "50 day average price."},
			{Name: "fifty_day_average_change", Type: proto.ColumnType_DOUBLE, Description: "50 day average change."},
			{Name: "fifty_day_average_change_percent", Type: proto.ColumnType_DOUBLE, Description: "50 day average change percentage."},
			{Name: "fifty_two_week_high", Type: proto.ColumnType_DOUBLE, Description: "52 week high."},
			{Name: "fifty_two_week_high_change", Type: proto.ColumnType_DOUBLE, Description: "52 week high change."},
			{Name: "fifty_two_week_high_change_percent", Type: proto.ColumnType_DOUBLE, Description: "52 week high change percentage."},
			{Name: "fifty_two_week_low", Type: proto.ColumnType_DOUBLE, Description: "52 week low."},
			{Name: "fifty_two_week_low_change", Type: proto.ColumnType_DOUBLE, Description: "52 week low change."},
			{Name: "fifty_two_week_low_change_percent", Type: proto.ColumnType_DOUBLE, Description: "52 week low change percent."},
			{Name: "full_exchange_name", Type: proto.ColumnType_STRING, Description: "Full exchange name."},
			{Name: "gmt_offset_milliseconds", Type: proto.ColumnType_INT, Transform: transform.FromField("GMTOffSetMilliseconds"), Description: "GMT offset in milliseconds."},
			{Name: "is_tradeable", Type: proto.ColumnType_BOOL, Description: "True if the symbol is tradeable."},
			{Name: "market_id", Type: proto.ColumnType_STRING, Description: "Market identifier, e.g. us_market."},
			{Name: "market_state", Type: proto.ColumnType_STRING, Description: "Current state of the market, e.g. REGULAR, CLOSED."},
			{Name: "post_market_change", Type: proto.ColumnType_DOUBLE, Description: "Post market price change."},
			{Name: "post_market_change_percent", Type: proto.ColumnType_DOUBLE, Description: "Post market price change percentage."},
			{Name: "post_market_price", Type: proto.ColumnType_DOUBLE, Description: "Post market price."},
			{Name: "post_market_time", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("PostMarketTime").Transform(transform.UnixToTimestamp), Description: "Timestamp for post market data."},
			{Name: "pre_market_change", Type: proto.ColumnType_DOUBLE, Description: "Pre market price change."},
			{Name: "pre_market_change_percent", Type: proto.ColumnType_DOUBLE, Description: "Pre market price change percentage."},
			{Name: "pre_market_price", Type: proto.ColumnType_DOUBLE, Description: "Pre market price."},
			{Name: "pre_market_time", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("PreMarketTime").Transform(transform.UnixToTimestamp), Description: "Timestamp for pre market data."},
			{Name: "quote_delay", Type: proto.ColumnType_INT, Description: "Quote delay in minutes."},
			{Name: "quote_source", Type: proto.ColumnType_STRING, Description: "Quote source."},
			{Name: "quote_type", Type: proto.ColumnType_STRING, Description: "Quote type, e.g. EQUITY, CRYPTOCURRENCY."},
			{Name: "regular_market_change", Type: proto.ColumnType_DOUBLE, Description: "Change in price since the regular market open."},
			{Name: "regular_market_change_percent", Type: proto.ColumnType_DOUBLE, Description: "Change percentage during the regular market session."},
			{Name: "regular_market_day_high", Type: proto.ColumnType_DOUBLE, Description: "High price for the regular market day."},
			{Name: "regular_market_day_low", Type: proto.ColumnType_DOUBLE, Description: "Low price for the regular market day."},
			{Name: "regular_market_open", Type: proto.ColumnType_DOUBLE, Description: "Opening price for the regular market."},
			{Name: "regular_market_previous_close", Type: proto.ColumnType_DOUBLE, Description: "Close price of the previous regular market session."},
			{Name: "regular_market_volume", Type: proto.ColumnType_INT, Description: "Trading volume for the regular market session."},
			{Name: "source_interval", Type: proto.ColumnType_INT, Description: "Source interval in minutes."},
			{Name: "two_hundred_day_average", Type: proto.ColumnType_DOUBLE, Description: "200 day average price."},
			{Name: "two_hundred_day_average_change", Type: proto.ColumnType_DOUBLE, Description: "200 day average price change."},
			{Name: "two_hundred_day_average_change_percent", Type: proto.ColumnType_DOUBLE, Description: "200 day average price change percentage."},
		},
	}
}

func listQuote(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	symbol := quals["symbol"].GetStringValue()
	q, err := quote.Get(symbol)
	if err != nil {
		plugin.Logger(ctx).Error("finance_quote.listQuote", "query_error", err)
		return nil, err
	}
	d.StreamListItem(ctx, q)
	return nil, nil
}
