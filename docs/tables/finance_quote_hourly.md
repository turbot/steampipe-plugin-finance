# Table: finance_quote_hourly

Historical price data by hour for a given symbol.

Note:
* A `symbol` must be provided in all queries to this table.
* History is limited to the last 13 months.
* Symbol types are defined in [finance_quote](./finance_quote).

## Examples

### Apple hourly price history (most recent first)

```sql
select
  timestamp,
  close
from
  finance_quote_hourly
where
  symbol = 'AAPL'
order by
  timestamp desc
```

### Hourly prices for Amazon on April 29th, 2020

```sql
select
  timestamp,
  close
from
  finance_quote_hourly
where
  symbol = 'AMZN'
  and date(timestamp) = '2020-04-29'
order by
  timestamp
```
