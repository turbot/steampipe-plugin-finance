![image](https://hub.steampipe.io/images/plugins/turbot/finance-social-graphic.png)

# Finance Plugin for Steampipe

Use SQL to query financial data including quotes and public company information.

* **[Get started →](https://hub.steampipe.io/plugins/turbot/finance)**
* Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/finance/tables)
* Community: [Discussion forums](https://github.com/turbot/steampipe/discussions)
* Get involved: [Issues](https://github.com/turbot/steampipe-plugin-finance/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):
```shell
steampipe plugin install finance
```

Run a query:
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

## Developing

Prerequisites:
- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone git@github.com:turbot/steampipe-plugin-finance
cd steampipe-plugin-finance
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:
```
make
```

Configure the plugin:
```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/finance.spc
```

Try it!
```
steampipe query
> .inspect finance
```

Further reading:
* [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
* [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [MPL-2.0 open source license](https://github.com/turbot/steampipe-plugin-finance/blob/main/LICENSE).

`help wanted` issues:
- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Finance Plugin](https://github.com/turbot/steampipe-plugin-finance/labels/help%20wanted)
