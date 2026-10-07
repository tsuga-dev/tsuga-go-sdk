# Changelog

All notable changes to the SDK are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.10.0] - 2026-10-07

### Added

- `CreateMonitorRequest`, `UpdateMonitorRequest`: `PauseUntil`, snoozing the monitor until the given time. On update, leave it unset to keep the current snooze or set it to nil to clear it.

## [1.9.0] - 2026-09-22

### Added

- `SlosAPI`: new API for SLOs.

## [1.8.0] - 2026-09-21

### Added

- Dashboards: `DataNormalizerPercentFraction`, a percentage normalizer for values between 0 and 1.
- Dashboards: SLO and span list widget types.
- Dashboards: relative mode on group by.
- Dashboards: line style options on timeseries widgets.

## [1.7.0] - 2026-07-20

Releases up to and including v1.7.0 predate this changelog. See the
[GitHub releases](https://github.com/tsuga-dev/tsuga-go-sdk/releases) for their contents.
