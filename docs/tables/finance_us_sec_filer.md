# Table: finance_us_sec_filer

Filer details from the US Securities and Exchange Commission (SEC) Edgar database.

Note: A `symbol` must be provided in all queries to this table.

## Examples

### Get filer details for Apple

```sql
select
  *
from
  finance_us_sec_filer
where
  symbol = 'AAPL'
```
