# http-backtest

HTTP Backtest in golang allows to run a series of comparison between two environments.

It can be used to verify that a migration, a new feature or a bug fix did not introduced a regression.

#### What is a regression

Regressions are a type of software bug. An example of regression is a feature that works but no longer works after new code releases. This issue is caused by several factors such as system upgrades, feature enhancements, and previous bug fixes through patching. The most common cause of regression is usually bug fixes.

#### Use Cases

You need to compare API responses between local and staging.

It can also be a safety net to catch divergence between two environments.

- Version Upgrade
- Migration
- New Feature
- Bug fixes

#### Supported body formats : 

- json
- xml
- csv
- html
- text

#### Requirements for best results : 

- Same data in both environment
- Idempotent HTTP requests (An HTTP request is idempotent if the intended effect on the server of making a single request is the same as the effect of making several identical requests)