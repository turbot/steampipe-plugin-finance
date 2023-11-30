---
title: "Steampipe Table: finance_us_sec_filer - Query Finance US SEC Filers using SQL"
description: "Allows users to query US SEC Filers in the Finance service. This table provides data on the filers' details, specifically their CIK, name, address, and other related information."
---

# Table: finance_us_sec_filer - Query Finance US SEC Filers using SQL

US SEC Filers are entities that are required to file statements, reports, and other forms with the U.S. Securities and Exchange Commission (SEC). These filings provide a comprehensive overview of a company's performance, financial health, and operations. The data from these filings is public and can be used for various analytical purposes.

## Table Usage Guide

The `finance_us_sec_filer` table provides insights into US SEC Filers within the Finance service. As a financial analyst or data scientist, explore filer-specific details through this table, including their CIK, name, address, and associated metadata. Utilize it to uncover information about filers, such as their location, to aid in your financial analysis and decision-making processes.

## Examples

### Get filer details for Apple
Explore comprehensive information about a specific company listed on the U.S. Securities and Exchange Commission (SEC) by its stock symbol. This is particularly useful for financial analysis or investment decision-making.

```sql
select
  *
from
  finance_us_sec_filer
where
  symbol = 'AAPL'
```