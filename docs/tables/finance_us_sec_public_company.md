---
title: "Steampipe Table: finance_us_sec_public_company - Query Finance US SEC Public Companies using SQL"
description: "Allows users to query US SEC Public Companies, specifically the financial data and filings, providing insights into financial health and regulatory compliance."
---

# Table: finance_us_sec_public_company - Query Finance US SEC Public Companies using SQL

The US Securities and Exchange Commission (SEC) Public Company is a regulated entity that is required to disclose certain business and financial information to the public and the SEC. This includes annual reports, quarterly earnings updates, and any significant developments that could affect the company's financial health. These disclosures help investors make informed decisions and ensure that companies are adhering to regulatory standards.

## Table Usage Guide

The `finance_us_sec_public_company` table provides insights into US SEC Public Companies within the finance sector. As a financial analyst or investor, explore company-specific details through this table, including financial data, filings, and associated metadata. Utilize it to uncover information about companies, such as their financial health, regulatory compliance, and significant developments that could impact their financial status.

**Note:** To query this table, you must configure an API token in the `~/.steampipe/config/finance.spc` file, as instructed in the documentation (https://hub.steampipe.io/plugins/turbot/finance#documentation).

## Examples

### List public companies
Explore which public companies are listed in the U.S. Securities and Exchange Commission's database. This is useful for gaining insights into the landscape of publicly traded companies.

```sql+postgres
select
  *
from
  finance_us_sec_public_company
order by
  name;
```

```sql+sqlite
select
  *
from
  finance_us_sec_public_company
order by
  name;
```

### Find a company by symbol
Discover the details of a specific company by using its unique symbol. This can be useful to quickly gain insights into a company's financial data without having to search through extensive databases.

```sql+postgres
select
  *
from
  finance_us_sec_public_company
where
  symbol = 'AAPL';
```

```sql+sqlite
select
  *
from
  finance_us_sec_public_company
where
  symbol = 'AAPL';
```

### Find a company by name
Discover the specifics of a particular company by searching using its name. This can be useful for gaining insights into a company's financial information, which can be beneficial for investment decisions or competitive analysis.

```sql+postgres
select
  *
from
  finance_us_sec_public_company
where
  name ilike '%apple%'
order by
  name;
```

```sql+sqlite
select
  *
from
  finance_us_sec_public_company
where
  name like '%apple%'
order by
  name;
```

### "Apple"-ish companies with their last closing price
Discover companies with names similar to "Apple" and gain insights into their latest closing prices. This query can be used to track and compare the financial performance of similarly named companies in the market.

```sql+postgres
select
  c.symbol,
  c.name,
  f.*
from
  finance_us_sec_public_company as c,
  finance_us_sec_filer as f
where
  c.symbol = f.symbol
  and c.name ilike '%apple%'
order by
  c.symbol;
```

```sql+sqlite
select
  c.symbol,
  c.name,
  f.*
from
  finance_us_sec_public_company as c,
  finance_us_sec_filer as f
where
  c.symbol = f.symbol
  and lower(c.name) like '%apple%'
order by
  c.symbol;
```