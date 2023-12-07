---
title: "Steampipe Table: finance_quote_hourly - Query Finance Quote Hourly data using SQL"
description: "Allows users to query hourly financial quotes, providing insights into hourly changes in financial data."
---

# Table: finance_quote_hourly - Query Finance Quote Hourly data using SQL

Finance Quote Hourly is a data set within the Finance service that provides hourly financial data. It offers a detailed view of financial data changes on an hourly basis, allowing for more granular analysis of financial trends and patterns. This data can be crucial for financial analysts, traders, and anyone interested in financial data analysis.

## Table Usage Guide

The `finance_quote_hourly` table provides insights into hourly financial data. As a financial analyst or trader, explore hourly financial trends and patterns through this table, including price changes, volume changes, and other relevant financial data. Utilize it to uncover detailed information about financial changes, such as sudden spikes or drops, and gain a better understanding of financial market trends.

**Important Notes**
- You must specify the `symbol` in the `where` clause to query this table.
- History is limited to the last 13 months.
- Symbol types are defined in [finance_quote](./finance_quote).

## Examples

### Apple hourly price history (most recent first)
Analyze the hourly closing prices of Apple's stock to understand its recent performance trends. This can help in making informed investment decisions by identifying patterns or changes in the stock's value.

```sql+postgres
select
  timestamp,
  close
from
  finance_quote_hourly
where
  symbol = 'AAPL'
order by
  timestamp desc;
```

```sql+sqlite
select
  timestamp,
  close
from
  finance_quote_hourly
where
  symbol = 'AAPL'
order by
  timestamp desc;
```

### Hourly prices for Amazon on April 29th, 2020
Explore the fluctuation of Amazon's stock prices on a specific date, April 29th, 2020, by analyzing the closing prices for each hour. This can assist in understanding the stock's performance and volatility throughout that day.

```sql+postgres
select
  timestamp,
  close
from
  finance_quote_hourly
where
  symbol = 'AMZN'
  and date(timestamp) = '2020-04-29'
order by
  timestamp;
```

```sql+sqlite
select
  timestamp,
  close
from
  finance_quote_hourly
where
  symbol = 'AMZN'
  and date(timestamp) = '2020-04-29'
order by
  timestamp;
```