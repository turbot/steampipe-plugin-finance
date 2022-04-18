package finance

import (
	"context"

	"github.com/shopspring/decimal"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func usSecHistoryColumns() []*plugin.Column {
	return []*plugin.Column{
		// Top columns
		{Name: "symbol", Type: proto.ColumnType_STRING, Hydrate: symbolString, Transform: transform.FromValue(), Description: "Symbol to quote."},
		// Other columns
		{Name: "adjusted_close", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("AdjClose").Transform(decimalToDouble), Description: "Adjusted close price after accounting for any corporate actions."},
		{Name: "close", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("Close").Transform(decimalToDouble), Description: "Last price during the regular trading session."},
		{Name: "high", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("High").Transform(decimalToDouble), Description: "Highest price during the trading session."},
		{Name: "low", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("Low").Transform(decimalToDouble), Description: "Lowest price during the trading session."},
		{Name: "open", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("Open").Transform(decimalToDouble), Description: "Opening price during the trading session."},
		{Name: "timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Timestamp").Transform(transform.UnixToTimestamp), Description: "Timestamp of the record."},
		{Name: "volume", Type: proto.ColumnType_INT, Description: "Total trading volume (units bought and sold) during the period."},
	}
}

func symbolString(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	s := quals["symbol"].GetStringValue()
	return s, nil
}


func decimalToDouble(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	dec := d.Value.(decimal.Decimal)
	f, _ := dec.Float64()
	return f, nil
}
