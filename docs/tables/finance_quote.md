# Table: finance_quote

Query prices for any symbol including:

| Type | Example Symbol | Example Name |
|-|-|
| Equities | `AMZN` | Amazon.com, Inc. |
| International Equities | `WBC.AX` | Westpac FPO (Australian Stock Exchange in AUD) |
| Indexes | `^FTSE` | FTSE 100 |
| Options | `AMZN210507C02240000` | AMZN May 2021 2240.000 call |
| Foreign exchange pairs | `AUDUSD=X` | AUD/USD |
| Cryptocurrency pairs | `BTC-USD` | Bitcoin USD |
| Futures | `NQ=F` | Nasdaq 100 Jun 21 |
| ETFs | `VTI` | Vanguard Total Stock Market ETF |
| Mutual Funds | `VFIAX` | Vanguard 500 Index Fd Admiral S |

Note: A `symbol` must be provided in all queries to this table.

## Examples

### Current price of Amazon stock

```sql
select
  symbol,
  short_name,
  regular_market_price
from
  finance_quote
where
  symbol = 'AMZN'
```

### 52 week trading range for Bitcoin

```sql
select
  symbol,
  short_name,
  fifty_two_week_low,
  fifty_two_week_high
from
  finance_quote
where
  symbol = 'BTC-USD'
```

### Global exchange quotes for Westpac

```sql
select
  symbol,
  short_name,
  regular_market_price,
  currency_id,
  full_exchange_name,
  exchange_timezone_name,
  regular_market_time
from
  finance_quote
where
  symbol in ('WBK', 'WBC.AX', 'WBC.NZ')
```
