# Table: finance_eps_history

Query EPS for any symbol including:

| Type       | Symbol    | Name     |
| :------------- | :----------: | -----------: |
| Equities | `AMZN` | Amazon.com, Inc. |
| International Equities | `'RELIANCE.NS'` | Reliance Industries (Indian Stock Exchange in INR) |

Note: A `symbol` must be provided in all queries to this table.

## Examples

### EPS history for Amazon stock

```sql
select 
    symbol,
    quarter, 
    estimate, 
    actual, 
    difference, 
    suprise_percent 
from finance_eps_history 
where 
    symbol='AMZN'
```

### Global exchange EPS for Reliance & State Bank Of India

```sql
select 
    * 
from 
    finance_eps_history 
where 
    symbol 
in('RELIANCE.NS','SBIN.NS')
```
