## v0.5.0 [2023-05-12]

_Dependencies_

- Recompiled plugin with [finance-go v1.1.0](https://github.com/piquette/finance-go/releases/tag/v1.1.0) which fixes the upstream API errors. ([#23](https://github.com/turbot/steampipe-plugin-finance/pull/23))

## v0.4.0 [2023-03-23]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#20](https://github.com/turbot/steampipe-plugin-finance/pull/20))

## v0.3.0 [2022-09-28]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v417-2022-09-08) which includes several caching and memory management improvements. ([#17](https://github.com/turbot/steampipe-plugin-finance/pull/17))
- Recompiled plugin with Go version `1.19`. ([#17](https://github.com/turbot/steampipe-plugin-finance/pull/17))

## v0.2.1 [2022-05-23]

_Bug fixes_

- Fixed the Slack community links in README and docs/index.md files. ([#13](https://github.com/turbot/steampipe-plugin-finance/pull/13))

## v0.2.0 [2022-04-27]

_Enhancements_

- Added support for native Linux ARM and Mac M1 builds. ([#10](https://github.com/turbot/steampipe-plugin-finance/pull/10))
- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#11](https://github.com/turbot/steampipe-plugin-finance/pull/11))

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
