# Table: finance_quote_hourly

Historical price data by day for a given symbol.

Note:
* A `symbol` must be provided in all queries to this table.
* History is limited to the last 121 months (~10 years).
* Symbol types are defined in [finance_quote](./finance_quote).

## Examples

### Johnson & Johnson daily price history (most recent first)

```sql
select
  timestamp,
  close
from
  finance_quote_daily
where
  symbol = 'JNJ'
order by
  timestamp desc
```

### Top 10 days by volume for Apple

```sql
select
  timestamp,
  volume
from
  finance_quote_daily
where
  symbol = 'AAPL'
order by
  volume desc
limit
  10
```

### Top 10 days by highest closing price for Bitcoin

```sql
select
  timestamp,
  close
from
  finance_quote_daily
where
  symbol = 'BTC-USD'
order by
  close desc
limit
  10
```
