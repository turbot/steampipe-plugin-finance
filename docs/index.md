---
organization: Turbot
category: ["media"]
icon_url: "/images/plugins/turbot/finance.svg"
brand_color: "#55CC44"
display_name: Finance
name: finance
description: Steampipe plugin to query financial data including quotes and public company information.
og_description: Query financial data with SQL! Open source CLI. No DB required.
og_image: "/images/plugins/turbot/finance-social-graphic.png"
---

# Finance + Steampipe

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

Financial data is queried from multiple sources including [Yahoo Finance](https://finance.yahoo.com/), [US SEC Edgar](https://www.sec.gov/edgar.shtml) and the [IEX Cloud](https://iexcloud.io/documentation/)service.

For example:

```sql
select
  symbol,
  short_name,
  regular_market_price,
  regular_market_change_percent,
  regular_market_time
from
  finance_quote
where
  symbol in ('GME', 'BTC-USD', 'DOGE-USD', 'ETH-USD');
```

```
+----------+----------------------+----------------------+-------------------------------+---------------------+
| symbol   | short_name           | regular_market_price | regular_market_change_percent | regular_market_time |
+----------+----------------------+----------------------+-------------------------------+---------------------+
| GME      | GameStop Corporation | 168.74               | -0.05330589                   | 2021-05-20 13:51:33 |
| BTC-USD  | Bitcoin USD          | 41959.867            | 18.85597                      | 2021-05-20 13:51:02 |
| DOGE-USD | Dogecoin USD         | 0.41606605           | 16.540815                     | 2021-05-20 13:51:03 |
| ETH-USD  | Ethereum USD         | 2917.209             | 20.443544                     | 2021-05-20 13:51:02 |
+----------+----------------------+----------------------+-------------------------------+---------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/finance/tables)**

## Get started

### Install

Download and install the latest Finance plugin:

```bash
steampipe plugin install finance
```

### Configuration

Installing the latest finance plugin will create a config file (`~/.steampipe/config/finance.spc`) with a single connection named `finance`:

```hcl
connection "finance" {
  plugin = "finance"

  iex_api_token = "pk_3746625324373hrf7fa44d2983d34e"
}
```

- `iex_api_token`(Optional) - An additional header name and its corresponding value to be sent to the server. If the header name is configured, the value must not be empty.

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-finance
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
