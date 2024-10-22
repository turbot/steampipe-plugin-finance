## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#65](https://github.com/turbot/steampipe-plugin-finance/pull/65))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#65](https://github.com/turbot/steampipe-plugin-finance/pull/65))

## v0.9.0 [2024-07-26]

_What's new?_

- Added support for [iex_api_token](https://iexcloud.io/documentation/reference/glossary.html#token-api-token) connnection config parameter to access `finance_us_sec_public_company` table data. ([#43](https://github.com/turbot/steampipe-plugin-finance/pull/43))

_Enhancements_

- The Plugin and the Steampipe Anywhere binaries are now built with the `netgo` package. ([#61](https://github.com/turbot/steampipe-plugin-finance/pull/61))
- Added the `version` flag to the plugin's Export tool. ([#65](https://github.com/turbot/steampipe-export/pull/65))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.10.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v5100-2024-04-10) that adds support for connection key columns. ([#54](https://github.com/turbot/steampipe-plugin-finance/pull/54))

## v0.8.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#49](https://github.com/turbot/steampipe-plugin-finance/pull/49))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#49](https://github.com/turbot/steampipe-plugin-finance/pull/49))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-finance/blob/main/docs/LICENSE). ([#49](https://github.com/turbot/steampipe-plugin-finance/pull/49))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#48](https://github.com/turbot/steampipe-plugin-finance/pull/48))

## v0.7.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#38](https://github.com/turbot/steampipe-plugin-finance/pull/38))

## v0.7.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#36](https://github.com/turbot/steampipe-plugin-finance/pull/36))
- Recompiled plugin with Go version `1.21`. ([#36](https://github.com/turbot/steampipe-plugin-finance/pull/36))

## v0.6.0 [2023-08-17]

_Dependencies_

- Recompiled plugin with `finance-go v1.1.1-0.20230807033903-430a57233430` that fixes the plugin to return results consistently instead of returning errors. ([#28](https://github.com/turbot/steampipe-plugin-finance/pull/28))

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
