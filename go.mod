module github.com/turbot/steampipe-plugin-finance

go 1.16

require (
	github.com/piquette/edgr v0.0.1
	github.com/piquette/finance-go v1.0.0
	github.com/shopspring/decimal v1.2.0
	github.com/turbot/steampipe-plugin-sdk v0.2.7
)

replace github.com/piquette/edgr => github.com/e-gineer/edgr v0.0.2-0.20210901021602-7664639af765
