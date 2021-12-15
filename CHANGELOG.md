## v0.1.0 [2021-12-15]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk-v1.8.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v182--2021-11-22) ([#7](https://github.com/turbot/steampipe-plugin-finance/pull/7))
- Recompiled plugin with Go version 1.17 ([#7](https://github.com/turbot/steampipe-plugin-finance/pull/7))

## v0.0.3 [2021-08-31]

_Bug fixes_

- Pass User-Agent to finance_us_sec_* tables to avoid throttling errors.

## v0.0.2 [2021-07-02]

_Enhancements_

- Updated plugin category from `internet` to `media`
- Updated plugin license to Apache 2.0 per [turbot/steampipe#488](https://github.com/turbot/steampipe/issues/488)

## v0.0.1 [2021-05-20]

_What's new?_

- New tables added
  - [finance_quote](https://hub.steampipe.io/plugins/turbot/finance/tables/finance_quote)
  - [finance_quote_daily](https://hub.steampipe.io/plugins/turbot/finance/tables/finance_quote_daily)
  - [finance_quote_hourly](https://hub.steampipe.io/plugins/turbot/finance/tables/finance_quote_hourly)
  - [finance_us_sec_filer](https://hub.steampipe.io/plugins/turbot/finance/tables/finance_us_sec_filer)
  - [finance_us_sec_public_company](https://hub.steampipe.io/plugins/turbot/finance/tables/finance_us_sec_public_company)
