---
organization: Turbot
category: ["internet"]
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

Financial data is queried from multiple sources including [Yahoo Finance](https://finance.yahoo.com/) and the [US SEC Edgar](https://www.sec.gov/edgar.shtml) service.

For example:
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

```
+--------+------------------+----------------------+
| symbol | short_name       | regular_market_price |
+--------+------------------+----------------------+
| AMZN   | Amazon.com, Inc. | 3467.42              |
+--------+------------------+----------------------+
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
  plugin     = "finance"
}
```

## Get involved

* Open source: https://github.com/turbot/steampipe-plugin-finance
* Community: [Discussion forums](https://github.com/turbot/steampipe/discussions)
