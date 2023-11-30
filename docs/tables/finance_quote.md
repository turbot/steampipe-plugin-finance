---
title: "Steampipe Table: finance_quote - Query Finance Quotes using SQL"
description: "Allows users to query Finance Quotes, specifically the real-time stock prices and related financial details, providing insights into market trends and investment decisions."
---

# Table: finance_quote - Query Finance Quotes using SQL

Finance Quotes is a resource that provides real-time stock prices and related financial details. It allows users to monitor and respond to market trends, making informed investment decisions. It offers a centralized way to set up and manage alerts for various financial resources, including stocks, mutual funds, ETFs, and more.

## Table Usage Guide

The `finance_quote` table provides insights into real-time stock prices and related financial details. As a financial analyst or investor, explore specific details through this table, including stock prices, trading volume, market capitalization, and associated metadata. Utilize it to uncover information about market trends, make informed investment decisions, and monitor your investment portfolio.

## Examples

### Current price of Amazon stock
Explore the current market value of a specific stock, in this case Amazon, to aid in financial decision making. This can be particularly useful for investors seeking up-to-date information for their portfolio management.

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
Analyze the fluctuations in Bitcoin's trading range over the past year. This can provide insights into the cryptocurrency's performance and volatility, helping to inform investment decisions.

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
Analyze the settings to understand the market status of Westpac across different exchanges. This is beneficial for tracking the performance of Westpac's shares in real-time, across multiple markets.

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