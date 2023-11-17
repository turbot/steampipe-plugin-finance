# Table: finance_us_sec_public_company

US public companies as registered by the US Securities and Exchange Commission.

**Note:** To query this table, you must configure an API token in the `~/.steampipe/config/finance.spc` file, as instructed in the documentation (https://hub.steampipe.io/plugins/turbot/finance#documentation).

## Examples

### List public companies

```sql
select
  *
from
  finance_us_sec_public_company
order by
  name
```

### Find a company by symbol

```sql
select
  *
from
  finance_us_sec_public_company
where
  symbol = 'AAPL'
```

### Find a company by name

```sql
select
  *
from
  finance_us_sec_public_company
where
  name ilike '%apple%'
order by
  name
```

### "Apple"-ish companies with their last closing price

```sql
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
  c.symbol
```
