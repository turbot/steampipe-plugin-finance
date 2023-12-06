---
title: "Steampipe Table: finance_quote_daily - Query Finance Quote Daily using SQL"
description: "Allows users to query daily finance quotes, providing detailed financial data for specific tickers on a daily basis."
---

# Table: finance_quote_daily - Query Finance Quote Daily using SQL

The Finance Quote Daily is a resource in the Finance service that provides detailed financial data for specific tickers on a daily basis. It includes information such as the opening price, closing price, highest price, lowest price, volume, and changes for each day. This data can be used for a variety of financial analysis and trading strategies.

## Table Usage Guide

The `finance_quote_daily` table provides valuable insights into the daily financial performance of specific tickers. As a financial analyst or trader, you can explore detailed financial data through this table, including opening and closing prices, daily highs and lows, volume, and changes. Use it to analyze market trends, evaluate trading strategies, and make informed investment decisions.

**Important Notes**
- You must specify the `symbol` in the `where` clause to query this table.
- History is limited to the last 121 months (~10 years).
- Symbol types are defined in [finance_quote](./finance_quote).

## Examples

### Johnson & Johnson daily price history (most recent first)
Analyze the daily closing prices of Johnson & Johnson stocks, listed in reverse chronological order. This can be useful for tracking stock performance over time and making informed investment decisions.

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
Discover the days when Apple's trading volume was at its highest. This can be beneficial for investors to understand the market activity and make informed decisions.

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
Discover the dates when Bitcoin had its highest closing prices. This query is useful for understanding market trends and making informed investment decisions.

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