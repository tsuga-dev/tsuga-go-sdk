# Changelog

All notable changes to the SDK are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.13.0] - 2026-10-28

### Added

- `MonitorConfigurationLog`, `MonitorConfigurationMetric`, `MonitorConfigurationTrace` and their `Input` variants: `ConsecutiveWindows`, the number of consecutive `Timeframe` windows that must all meet the conditions before the monitor alerts. Leave it unset or set it to 1 for single-window evaluation.

## [1.12.0] - 2026-10-15

### Added

- `AdaptiveNotificationRulesAPI`: new API for adaptive notification rules. On update, leave `MonitorIds` or `SloIds` unset to keep the current list, or send an empty list to clear it.

## [1.11.0] - 2026-10-15

### Added

- `MonitorFoldersAPI`: new API for monitor folders.
- `CreateMonitorRequest`, `UpdateMonitorRequest`: `FolderId`, filing the monitor into a monitor folder. On update, leave it unset to keep the current folder or set it to nil to remove it.
- `Monitor`: `FolderId`, the folder holding the monitor.
- `QueryMonitorsRequestFilters`: `FolderId`, restricting results to a folder, or to unfiled monitors when set to nil.

## [1.10.1] - 2026-10-12

### Changed

- Breaking: model names no longer depend on spec ordering. Request-side models are prefixed with `Input` and numbered response models lose their suffix.
- Request-side models used when creating or updating resources (old name → new name):
  - `AggregationGroupBy` → `InputAggregationGroupBy`
  - `AggregationQuery1` → `InputAggregationQuery`
  - `AllPublicTeams` → `InputAllPublicTeams`
  - `AllTeams` → `InputAllTeams`
  - `BurnRateAlert` → `InputBurnRateAlert`
  - `ConditionalFormatting` → `InputConditionalFormatting`
  - `CustomLink1` → `InputCustomLink`
  - `EventSloConfiguration` → `InputEventSloConfiguration`
  - `GaugeColorThreshold` → `InputGaugeColorThreshold`
  - `Graph` → `InputGraph`
  - `Graph1` → `InputGraph`
  - `GraphLayout` (as `UpdateDashboardGraphRequest.Layout`) → `UpdateDashboardGraphRequestLayout` (same fields)
  - `GraphVisualization` (the request oneOf, used by `Graph`, `Graph1` and `UpdateDashboardGraphRequest`) → `UpdateDashboardGraphRequestVisualization` (used by `InputGraph` and `UpdateDashboardGraphRequest`)
  - `ListDefaultSorting1` → `InputListDefaultSorting`
  - `MonitorAggregationQuery1` → `InputMonitorAggregationQuery`
  - `Normalizer` → `InputNormalizer`
  - `SloAggregationQuery` → `InputSloAggregationQuery`
  - `SloQueryFormula` → `InputSloQueryFormula`
  - `SpecificTeams` → `InputSpecificTeams`
  - `SpecificTeams1` → `InputSpecificTeams`
  - `TableColumn1` → `InputTableColumn`
  - `TableDefaultSorting1` → `InputTableDefaultSorting`
  - `Tag` → `InputTag`
  - `TelemetryTagPolicy` → `InputTelemetryTagPolicy`
  - `ThresholdAlert` → `InputThresholdAlert`
  - `ThresholdMarker` → `InputThresholdMarker`
  - `TimeSliceSloConfiguration` → `InputTimeSliceSloConfiguration`
  - `TsugaAssetTagPolicy` → `InputTsugaAssetTagPolicy`
  - `UpdateNotificationRuleRequestTeamsFilter` → `CreateNotificationRuleRequestTeamsFilter`
  - `WidgetListColumn1` → `InputWidgetListColumn`
- Response-side models that lost their numeric suffix (old name → new name):
  - `AggregationGroupBy1` → `AggregationGroupBy`
  - `BurnRateAlert1` → `BurnRateAlert`
  - `EventSloConfiguration1` → `EventSloConfiguration`
  - `Graph2` → `Graph`
  - `Graph1Layout` → `GraphLayout`
  - `Graph2Visualization` → `GraphVisualization`
  - `Normalizer1` → `Normalizer`
  - `SloQueryFormula1` → `SloQueryFormula`
  - `SpecificTeams2` → `SpecificTeams`
  - `Tag1` → `Tag`
  - `TelemetryTagPolicy1` → `TelemetryTagPolicy`
  - `ThresholdAlert1` → `ThresholdAlert`
  - `TimeSliceSloConfiguration1` → `TimeSliceSloConfiguration`
- Renamed oneOf members also rename the wrapper fields that hold them, for example `RuleTeamsFilter.SpecificTeams2` → `.SpecificTeams` and `UpdateSloRequestAlertsInnerConfiguration.BurnRateAlert` → `.InputBurnRateAlert`, and their convenience constructors, for example `SpecificTeams2AsRuleTeamsFilter` → `SpecificTeamsAsRuleTeamsFilter` and `TelemetryTagPolicy1AsTagPolicyConfiguration` → `TelemetryTagPolicyAsTagPolicyConfiguration`.

- Breaking: model types that no longer exist (see the renames above for replacements):
  - `AggregationGroupBy1`
  - `AggregationQuery1`
  - `BurnRateAlert1`
  - `CustomLink1`
  - `EventSloConfiguration1`
  - `Graph1`
  - `Graph1Layout`
  - `Graph2`
  - `Graph2Visualization`
  - `ListDefaultSorting1`
  - `MonitorAggregationQuery1`
  - `Normalizer1`
  - `SloAggregationQuery`
  - `SloQueryFormula1`
  - `SpecificTeams1`
  - `SpecificTeams2`
  - `TableColumn1`
  - `TableDefaultSorting1`
  - `Tag1`
  - `TelemetryTagPolicy1`
  - `ThresholdAlert1`
  - `TimeSliceSloConfiguration1`
  - `UpdateNotificationRuleRequestTeamsFilter`
  - `WidgetListColumn1`

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
