# UpdateDashboardGraphRequestVisualization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays the database rows-based aggregation as a time series chart | 
**ConnectionId** | **string** | The ID of the connection to use to query the datastore. | 
**Queries** | [**[]InputAggregationQuery**](InputAggregationQuery.md) | Aggregations that may be combined together in the same query. Each item is referenced from &#x60;formula&#x60; as q1, q2, and so on, in submission order. For dataSource \&quot;metrics\&quot;, each aggregate&#39;s &#x60;field&#x60; is the metric name, not an attribute; to count distinct values of an attribute use unique-count with field \&quot;&lt;metricName&gt;.context.&lt;attribute&gt;\&quot; (e.g. \&quot;system.cpu.utilization.context.host.name\&quot;). | 
**LegendMode** | Pointer to **string** | Controls whether and how the widget displays legend or series details (e.g. table, legend-only, or no legend) | [optional] 
**Thresholds** | Pointer to [**[]InputThresholdMarker**](InputThresholdMarker.md) | Threshold markers displayed on the chart | [optional] 
**YAxisSettings** | Pointer to [**InputGraphVisualizationTimeseriesConnectionYAxisSettings**](InputGraphVisualizationTimeseriesConnectionYAxisSettings.md) |  | [optional] 
**LineStyleOptions** | Pointer to [**map[string]GraphVisualizationTimeseriesConnectionLineStyleOptionsValue**](GraphVisualizationTimeseriesConnectionLineStyleOptionsValue.md) | Line style of each series, keyed by 1-based query index. The last index is the formula when there is one. For widgets with a single query, only the &#x60;1&#x60; entry is read and it applies to every series. Defaults to regular. | [optional] 
**Query** | **string** | Query that selects RUM events for the list. Prefix every field with the stream it belongs to (&#x60;measurements.&#x60;, &#x60;events.&#x60;, or &#x60;exceptions.&#x60;) and filter on a single stream per query. A query without stream-prefixed fields searches the &#x60;events&#x60; stream. | 
**ListColumns** | Pointer to [**[]InputWidgetListColumn**](InputWidgetListColumn.md) | Custom columns to display for each RUM event. Set attributes without their stream prefix, e.g. &#x60;view.name&#x60; rather than &#x60;events.view.name&#x60;. | [optional] 
**ListColumnsSize** | Pointer to **map[string]float32** | List column widths in pixels, keyed by the &#x60;attribute&#x60; of the matching &#x60;listColumns&#x60; entry. Columns without an entry keep their default width. | [optional] 
**IsCellWrapped** | Pointer to **bool** | Whether list widget cell text wraps instead of truncating. Applies to RUM event list widgets. Optional; omit or set false to use truncated cells. | [optional] 
**DefaultSorting** | Pointer to [**[]InputListDefaultSorting**](InputListDefaultSorting.md) | Default sorting applied to a list widget. Optional on create or update for log, span, RUM event, or connection list widgets. Users can still change sorting by selecting columns in the rendered list. | [optional] 
**BackgroundMode** | Pointer to **string** | Controls whether the widget uses a solid or transparent background | [optional] 
**Conditions** | Pointer to [**[]InputConditionalFormatting**](InputConditionalFormatting.md) | Conditional formatting rules applied to the raw query value, before the unit is applied | [optional] 
**Precision** | Pointer to [**GraphVisualizationQueryValueConnectionPrecision**](GraphVisualizationQueryValueConnectionPrecision.md) |  | [optional] 
**Normalizer** | Pointer to [**InputNormalizer**](InputNormalizer.md) |  | [optional] 
**Aliases** | Pointer to [**InputGraphVisualizationTimeseriesPromqlAliases**](InputGraphVisualizationTimeseriesPromqlAliases.md) |  | [optional] 
**TimeBucket** | Pointer to [**GraphVisualizationTimeseriesPromqlTimeBucket**](GraphVisualizationTimeseriesPromqlTimeBucket.md) |  | [optional] 
**Smoothing** | Pointer to **bool** | Whether to apply automatic smoothing to the rendered timeseries | [optional] 
**Source** | **string** | Telemetry source queried by this aggregation: &#x60;logs&#x60;, &#x60;metrics&#x60;, &#x60;traces&#x60;, or &#x60;rum&#x60;. | 
**Formula** | Pointer to **string** | Formula referencing query outputs, such as &#x60;q1 + q2&#x60;, to compute derived results. Defaults to &#x60;q1&#x60;. Formulas may reference only submitted queries (&#x60;q1&#x60; through &#x60;qN&#x60;); undefined query references return 400. | [optional] 
**VisibleSeries** | Pointer to **[]bool** | Flags indicating whether each query or formula series is visible | [optional] 
**GroupBy** | Pointer to [**[]InputAggregationGroupBy**](InputAggregationGroupBy.md) | Nested grouping levels applied to aggregation results, outermost first (e.g. group by service, then by level within each service). Each level splits results further, so the response contains one result per unique combination of group values instead of one aggregated total. Defaults to an empty array (one ungrouped result) when omitted. | [optional] 
**GroupByMode** | Pointer to **string** | &#x60;absolute&#x60; keeps each group at its own value; &#x60;relative&#x60; shows it as a percentage of the ungrouped total (defaults to absolute) | [optional] 
**CustomLink** | Pointer to [**InputCustomLink**](InputCustomLink.md) |  | [optional] 
**IsStacked** | Pointer to **bool** | Requests stacked rendering for a top-list widget. Tsuga renders stacked rows only for one count or sum query with exactly two grouped fields, no formula, non-negative values, and a single-cluster context; otherwise the widget renders as a normal top list. | [optional] 
**Max** | Pointer to **float32** | Gauge maximum value, on the raw query value before the unit is applied | [optional] 
**ColorThresholds** | Pointer to [**[]InputGaugeColorThreshold**](InputGaugeColorThreshold.md) | Color thresholds inside the gauge range, on the raw query value before the unit is applied | [optional] 
**Group** | Pointer to **string** | Attribute that switches the count to \&quot;Groups\&quot; mode: records are grouped by this attribute, the aggregation produces one value per group, and the chart buckets those per-group values. When omitted, individual records are bucketed. | [optional] 
**PercentileMarkers** | Pointer to **[]int32** | Percentile markers displayed on top of the distribution chart | [optional] 
**BoundsScale** | Pointer to **string** | Spacing of the bucket boundaries across the distribution range. &#x60;linear&#x60; splits the range into equal-width buckets; &#x60;log&#x60; widens each bucket logarithmically, giving finer resolution near the lower bound. | [optional] 
**Palette** | Pointer to **string** | Color palette used to render the heatmap intensity gradient | [optional] 
**Columns** | [**[]InputTableColumn**](InputTableColumn.md) | Each column defines an independent aggregation displayed as a table column | 
**ColumnSizes** | Pointer to **map[string]float32** | Table column widths in pixels, keyed by column id: &#x60;label&#x60; for the grouping column and &#x60;col-&lt;index&gt;&#x60; for each entry in &#x60;columns&#x60;. Columns without an entry keep their default width. | [optional] 
**Note** | Pointer to **string** | Markdown-compatible text shown in the note | [optional] 
**NoteColor** | Pointer to **string** | Background color used to render the note | [optional] 
**NoteAlign** | Pointer to **string** | Flex alignment keyword used for widget layout | [optional] 
**NoteJustifyContent** | Pointer to **string** | Flex alignment keyword used for widget layout | [optional] 
**Layout** | Pointer to **string** | Layout used to render log patterns | [optional] 
**SloId** | **string** | Id of the SLO to display. Omit it to cover every SLO that matches &#x60;filters&#x60;. | 
**Filters** | Pointer to [**InputGraphVisualizationSloFilters**](InputGraphVisualizationSloFilters.md) |  | [optional] 
**OrganizeBy** | Pointer to **string** | Groups the results by owning team or by cluster. Omit it to list results ungrouped. | [optional] 
**Display** | **[]string** | Sections to show: &#x60;summary&#x60; for the result count per status, &#x60;table&#x60; for the list of results. | 

## Methods

### NewUpdateDashboardGraphRequestVisualization

`func NewUpdateDashboardGraphRequestVisualization(type_ string, connectionId string, queries []InputAggregationQuery, query string, source string, columns []InputTableColumn, sloId string, display []string, ) *UpdateDashboardGraphRequestVisualization`

NewUpdateDashboardGraphRequestVisualization instantiates a new UpdateDashboardGraphRequestVisualization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDashboardGraphRequestVisualizationWithDefaults

`func NewUpdateDashboardGraphRequestVisualizationWithDefaults() *UpdateDashboardGraphRequestVisualization`

NewUpdateDashboardGraphRequestVisualizationWithDefaults instantiates a new UpdateDashboardGraphRequestVisualization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateDashboardGraphRequestVisualization) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateDashboardGraphRequestVisualization) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateDashboardGraphRequestVisualization) SetType(v string)`

SetType sets Type field to given value.


### GetConnectionId

`func (o *UpdateDashboardGraphRequestVisualization) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *UpdateDashboardGraphRequestVisualization) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *UpdateDashboardGraphRequestVisualization) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetQueries

`func (o *UpdateDashboardGraphRequestVisualization) GetQueries() []InputAggregationQuery`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *UpdateDashboardGraphRequestVisualization) GetQueriesOk() (*[]InputAggregationQuery, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *UpdateDashboardGraphRequestVisualization) SetQueries(v []InputAggregationQuery)`

SetQueries sets Queries field to given value.


### GetLegendMode

`func (o *UpdateDashboardGraphRequestVisualization) GetLegendMode() string`

GetLegendMode returns the LegendMode field if non-nil, zero value otherwise.

### GetLegendModeOk

`func (o *UpdateDashboardGraphRequestVisualization) GetLegendModeOk() (*string, bool)`

GetLegendModeOk returns a tuple with the LegendMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegendMode

`func (o *UpdateDashboardGraphRequestVisualization) SetLegendMode(v string)`

SetLegendMode sets LegendMode field to given value.

### HasLegendMode

`func (o *UpdateDashboardGraphRequestVisualization) HasLegendMode() bool`

HasLegendMode returns a boolean if a field has been set.

### GetThresholds

`func (o *UpdateDashboardGraphRequestVisualization) GetThresholds() []InputThresholdMarker`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *UpdateDashboardGraphRequestVisualization) GetThresholdsOk() (*[]InputThresholdMarker, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *UpdateDashboardGraphRequestVisualization) SetThresholds(v []InputThresholdMarker)`

SetThresholds sets Thresholds field to given value.

### HasThresholds

`func (o *UpdateDashboardGraphRequestVisualization) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.

### GetYAxisSettings

`func (o *UpdateDashboardGraphRequestVisualization) GetYAxisSettings() InputGraphVisualizationTimeseriesConnectionYAxisSettings`

GetYAxisSettings returns the YAxisSettings field if non-nil, zero value otherwise.

### GetYAxisSettingsOk

`func (o *UpdateDashboardGraphRequestVisualization) GetYAxisSettingsOk() (*InputGraphVisualizationTimeseriesConnectionYAxisSettings, bool)`

GetYAxisSettingsOk returns a tuple with the YAxisSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYAxisSettings

`func (o *UpdateDashboardGraphRequestVisualization) SetYAxisSettings(v InputGraphVisualizationTimeseriesConnectionYAxisSettings)`

SetYAxisSettings sets YAxisSettings field to given value.

### HasYAxisSettings

`func (o *UpdateDashboardGraphRequestVisualization) HasYAxisSettings() bool`

HasYAxisSettings returns a boolean if a field has been set.

### GetLineStyleOptions

`func (o *UpdateDashboardGraphRequestVisualization) GetLineStyleOptions() map[string]GraphVisualizationTimeseriesConnectionLineStyleOptionsValue`

GetLineStyleOptions returns the LineStyleOptions field if non-nil, zero value otherwise.

### GetLineStyleOptionsOk

`func (o *UpdateDashboardGraphRequestVisualization) GetLineStyleOptionsOk() (*map[string]GraphVisualizationTimeseriesConnectionLineStyleOptionsValue, bool)`

GetLineStyleOptionsOk returns a tuple with the LineStyleOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineStyleOptions

`func (o *UpdateDashboardGraphRequestVisualization) SetLineStyleOptions(v map[string]GraphVisualizationTimeseriesConnectionLineStyleOptionsValue)`

SetLineStyleOptions sets LineStyleOptions field to given value.

### HasLineStyleOptions

`func (o *UpdateDashboardGraphRequestVisualization) HasLineStyleOptions() bool`

HasLineStyleOptions returns a boolean if a field has been set.

### GetQuery

`func (o *UpdateDashboardGraphRequestVisualization) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *UpdateDashboardGraphRequestVisualization) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *UpdateDashboardGraphRequestVisualization) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetListColumns

`func (o *UpdateDashboardGraphRequestVisualization) GetListColumns() []InputWidgetListColumn`

GetListColumns returns the ListColumns field if non-nil, zero value otherwise.

### GetListColumnsOk

`func (o *UpdateDashboardGraphRequestVisualization) GetListColumnsOk() (*[]InputWidgetListColumn, bool)`

GetListColumnsOk returns a tuple with the ListColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListColumns

`func (o *UpdateDashboardGraphRequestVisualization) SetListColumns(v []InputWidgetListColumn)`

SetListColumns sets ListColumns field to given value.

### HasListColumns

`func (o *UpdateDashboardGraphRequestVisualization) HasListColumns() bool`

HasListColumns returns a boolean if a field has been set.

### GetListColumnsSize

`func (o *UpdateDashboardGraphRequestVisualization) GetListColumnsSize() map[string]float32`

GetListColumnsSize returns the ListColumnsSize field if non-nil, zero value otherwise.

### GetListColumnsSizeOk

`func (o *UpdateDashboardGraphRequestVisualization) GetListColumnsSizeOk() (*map[string]float32, bool)`

GetListColumnsSizeOk returns a tuple with the ListColumnsSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListColumnsSize

`func (o *UpdateDashboardGraphRequestVisualization) SetListColumnsSize(v map[string]float32)`

SetListColumnsSize sets ListColumnsSize field to given value.

### HasListColumnsSize

`func (o *UpdateDashboardGraphRequestVisualization) HasListColumnsSize() bool`

HasListColumnsSize returns a boolean if a field has been set.

### GetIsCellWrapped

`func (o *UpdateDashboardGraphRequestVisualization) GetIsCellWrapped() bool`

GetIsCellWrapped returns the IsCellWrapped field if non-nil, zero value otherwise.

### GetIsCellWrappedOk

`func (o *UpdateDashboardGraphRequestVisualization) GetIsCellWrappedOk() (*bool, bool)`

GetIsCellWrappedOk returns a tuple with the IsCellWrapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCellWrapped

`func (o *UpdateDashboardGraphRequestVisualization) SetIsCellWrapped(v bool)`

SetIsCellWrapped sets IsCellWrapped field to given value.

### HasIsCellWrapped

`func (o *UpdateDashboardGraphRequestVisualization) HasIsCellWrapped() bool`

HasIsCellWrapped returns a boolean if a field has been set.

### GetDefaultSorting

`func (o *UpdateDashboardGraphRequestVisualization) GetDefaultSorting() []InputListDefaultSorting`

GetDefaultSorting returns the DefaultSorting field if non-nil, zero value otherwise.

### GetDefaultSortingOk

`func (o *UpdateDashboardGraphRequestVisualization) GetDefaultSortingOk() (*[]InputListDefaultSorting, bool)`

GetDefaultSortingOk returns a tuple with the DefaultSorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSorting

`func (o *UpdateDashboardGraphRequestVisualization) SetDefaultSorting(v []InputListDefaultSorting)`

SetDefaultSorting sets DefaultSorting field to given value.

### HasDefaultSorting

`func (o *UpdateDashboardGraphRequestVisualization) HasDefaultSorting() bool`

HasDefaultSorting returns a boolean if a field has been set.

### GetBackgroundMode

`func (o *UpdateDashboardGraphRequestVisualization) GetBackgroundMode() string`

GetBackgroundMode returns the BackgroundMode field if non-nil, zero value otherwise.

### GetBackgroundModeOk

`func (o *UpdateDashboardGraphRequestVisualization) GetBackgroundModeOk() (*string, bool)`

GetBackgroundModeOk returns a tuple with the BackgroundMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundMode

`func (o *UpdateDashboardGraphRequestVisualization) SetBackgroundMode(v string)`

SetBackgroundMode sets BackgroundMode field to given value.

### HasBackgroundMode

`func (o *UpdateDashboardGraphRequestVisualization) HasBackgroundMode() bool`

HasBackgroundMode returns a boolean if a field has been set.

### GetConditions

`func (o *UpdateDashboardGraphRequestVisualization) GetConditions() []InputConditionalFormatting`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *UpdateDashboardGraphRequestVisualization) GetConditionsOk() (*[]InputConditionalFormatting, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *UpdateDashboardGraphRequestVisualization) SetConditions(v []InputConditionalFormatting)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *UpdateDashboardGraphRequestVisualization) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetPrecision

`func (o *UpdateDashboardGraphRequestVisualization) GetPrecision() GraphVisualizationQueryValueConnectionPrecision`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *UpdateDashboardGraphRequestVisualization) GetPrecisionOk() (*GraphVisualizationQueryValueConnectionPrecision, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *UpdateDashboardGraphRequestVisualization) SetPrecision(v GraphVisualizationQueryValueConnectionPrecision)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *UpdateDashboardGraphRequestVisualization) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetNormalizer

`func (o *UpdateDashboardGraphRequestVisualization) GetNormalizer() InputNormalizer`

GetNormalizer returns the Normalizer field if non-nil, zero value otherwise.

### GetNormalizerOk

`func (o *UpdateDashboardGraphRequestVisualization) GetNormalizerOk() (*InputNormalizer, bool)`

GetNormalizerOk returns a tuple with the Normalizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizer

`func (o *UpdateDashboardGraphRequestVisualization) SetNormalizer(v InputNormalizer)`

SetNormalizer sets Normalizer field to given value.

### HasNormalizer

`func (o *UpdateDashboardGraphRequestVisualization) HasNormalizer() bool`

HasNormalizer returns a boolean if a field has been set.

### GetAliases

`func (o *UpdateDashboardGraphRequestVisualization) GetAliases() InputGraphVisualizationTimeseriesPromqlAliases`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *UpdateDashboardGraphRequestVisualization) GetAliasesOk() (*InputGraphVisualizationTimeseriesPromqlAliases, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *UpdateDashboardGraphRequestVisualization) SetAliases(v InputGraphVisualizationTimeseriesPromqlAliases)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *UpdateDashboardGraphRequestVisualization) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetTimeBucket

`func (o *UpdateDashboardGraphRequestVisualization) GetTimeBucket() GraphVisualizationTimeseriesPromqlTimeBucket`

GetTimeBucket returns the TimeBucket field if non-nil, zero value otherwise.

### GetTimeBucketOk

`func (o *UpdateDashboardGraphRequestVisualization) GetTimeBucketOk() (*GraphVisualizationTimeseriesPromqlTimeBucket, bool)`

GetTimeBucketOk returns a tuple with the TimeBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBucket

`func (o *UpdateDashboardGraphRequestVisualization) SetTimeBucket(v GraphVisualizationTimeseriesPromqlTimeBucket)`

SetTimeBucket sets TimeBucket field to given value.

### HasTimeBucket

`func (o *UpdateDashboardGraphRequestVisualization) HasTimeBucket() bool`

HasTimeBucket returns a boolean if a field has been set.

### GetSmoothing

`func (o *UpdateDashboardGraphRequestVisualization) GetSmoothing() bool`

GetSmoothing returns the Smoothing field if non-nil, zero value otherwise.

### GetSmoothingOk

`func (o *UpdateDashboardGraphRequestVisualization) GetSmoothingOk() (*bool, bool)`

GetSmoothingOk returns a tuple with the Smoothing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmoothing

`func (o *UpdateDashboardGraphRequestVisualization) SetSmoothing(v bool)`

SetSmoothing sets Smoothing field to given value.

### HasSmoothing

`func (o *UpdateDashboardGraphRequestVisualization) HasSmoothing() bool`

HasSmoothing returns a boolean if a field has been set.

### GetSource

`func (o *UpdateDashboardGraphRequestVisualization) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *UpdateDashboardGraphRequestVisualization) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *UpdateDashboardGraphRequestVisualization) SetSource(v string)`

SetSource sets Source field to given value.


### GetFormula

`func (o *UpdateDashboardGraphRequestVisualization) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *UpdateDashboardGraphRequestVisualization) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *UpdateDashboardGraphRequestVisualization) SetFormula(v string)`

SetFormula sets Formula field to given value.

### HasFormula

`func (o *UpdateDashboardGraphRequestVisualization) HasFormula() bool`

HasFormula returns a boolean if a field has been set.

### GetVisibleSeries

`func (o *UpdateDashboardGraphRequestVisualization) GetVisibleSeries() []bool`

GetVisibleSeries returns the VisibleSeries field if non-nil, zero value otherwise.

### GetVisibleSeriesOk

`func (o *UpdateDashboardGraphRequestVisualization) GetVisibleSeriesOk() (*[]bool, bool)`

GetVisibleSeriesOk returns a tuple with the VisibleSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleSeries

`func (o *UpdateDashboardGraphRequestVisualization) SetVisibleSeries(v []bool)`

SetVisibleSeries sets VisibleSeries field to given value.

### HasVisibleSeries

`func (o *UpdateDashboardGraphRequestVisualization) HasVisibleSeries() bool`

HasVisibleSeries returns a boolean if a field has been set.

### GetGroupBy

`func (o *UpdateDashboardGraphRequestVisualization) GetGroupBy() []InputAggregationGroupBy`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *UpdateDashboardGraphRequestVisualization) GetGroupByOk() (*[]InputAggregationGroupBy, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *UpdateDashboardGraphRequestVisualization) SetGroupBy(v []InputAggregationGroupBy)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *UpdateDashboardGraphRequestVisualization) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetGroupByMode

`func (o *UpdateDashboardGraphRequestVisualization) GetGroupByMode() string`

GetGroupByMode returns the GroupByMode field if non-nil, zero value otherwise.

### GetGroupByModeOk

`func (o *UpdateDashboardGraphRequestVisualization) GetGroupByModeOk() (*string, bool)`

GetGroupByModeOk returns a tuple with the GroupByMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByMode

`func (o *UpdateDashboardGraphRequestVisualization) SetGroupByMode(v string)`

SetGroupByMode sets GroupByMode field to given value.

### HasGroupByMode

`func (o *UpdateDashboardGraphRequestVisualization) HasGroupByMode() bool`

HasGroupByMode returns a boolean if a field has been set.

### GetCustomLink

`func (o *UpdateDashboardGraphRequestVisualization) GetCustomLink() InputCustomLink`

GetCustomLink returns the CustomLink field if non-nil, zero value otherwise.

### GetCustomLinkOk

`func (o *UpdateDashboardGraphRequestVisualization) GetCustomLinkOk() (*InputCustomLink, bool)`

GetCustomLinkOk returns a tuple with the CustomLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLink

`func (o *UpdateDashboardGraphRequestVisualization) SetCustomLink(v InputCustomLink)`

SetCustomLink sets CustomLink field to given value.

### HasCustomLink

`func (o *UpdateDashboardGraphRequestVisualization) HasCustomLink() bool`

HasCustomLink returns a boolean if a field has been set.

### GetIsStacked

`func (o *UpdateDashboardGraphRequestVisualization) GetIsStacked() bool`

GetIsStacked returns the IsStacked field if non-nil, zero value otherwise.

### GetIsStackedOk

`func (o *UpdateDashboardGraphRequestVisualization) GetIsStackedOk() (*bool, bool)`

GetIsStackedOk returns a tuple with the IsStacked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStacked

`func (o *UpdateDashboardGraphRequestVisualization) SetIsStacked(v bool)`

SetIsStacked sets IsStacked field to given value.

### HasIsStacked

`func (o *UpdateDashboardGraphRequestVisualization) HasIsStacked() bool`

HasIsStacked returns a boolean if a field has been set.

### GetMax

`func (o *UpdateDashboardGraphRequestVisualization) GetMax() float32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *UpdateDashboardGraphRequestVisualization) GetMaxOk() (*float32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *UpdateDashboardGraphRequestVisualization) SetMax(v float32)`

SetMax sets Max field to given value.

### HasMax

`func (o *UpdateDashboardGraphRequestVisualization) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetColorThresholds

`func (o *UpdateDashboardGraphRequestVisualization) GetColorThresholds() []InputGaugeColorThreshold`

GetColorThresholds returns the ColorThresholds field if non-nil, zero value otherwise.

### GetColorThresholdsOk

`func (o *UpdateDashboardGraphRequestVisualization) GetColorThresholdsOk() (*[]InputGaugeColorThreshold, bool)`

GetColorThresholdsOk returns a tuple with the ColorThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorThresholds

`func (o *UpdateDashboardGraphRequestVisualization) SetColorThresholds(v []InputGaugeColorThreshold)`

SetColorThresholds sets ColorThresholds field to given value.

### HasColorThresholds

`func (o *UpdateDashboardGraphRequestVisualization) HasColorThresholds() bool`

HasColorThresholds returns a boolean if a field has been set.

### GetGroup

`func (o *UpdateDashboardGraphRequestVisualization) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *UpdateDashboardGraphRequestVisualization) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *UpdateDashboardGraphRequestVisualization) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *UpdateDashboardGraphRequestVisualization) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetPercentileMarkers

`func (o *UpdateDashboardGraphRequestVisualization) GetPercentileMarkers() []int32`

GetPercentileMarkers returns the PercentileMarkers field if non-nil, zero value otherwise.

### GetPercentileMarkersOk

`func (o *UpdateDashboardGraphRequestVisualization) GetPercentileMarkersOk() (*[]int32, bool)`

GetPercentileMarkersOk returns a tuple with the PercentileMarkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentileMarkers

`func (o *UpdateDashboardGraphRequestVisualization) SetPercentileMarkers(v []int32)`

SetPercentileMarkers sets PercentileMarkers field to given value.

### HasPercentileMarkers

`func (o *UpdateDashboardGraphRequestVisualization) HasPercentileMarkers() bool`

HasPercentileMarkers returns a boolean if a field has been set.

### GetBoundsScale

`func (o *UpdateDashboardGraphRequestVisualization) GetBoundsScale() string`

GetBoundsScale returns the BoundsScale field if non-nil, zero value otherwise.

### GetBoundsScaleOk

`func (o *UpdateDashboardGraphRequestVisualization) GetBoundsScaleOk() (*string, bool)`

GetBoundsScaleOk returns a tuple with the BoundsScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundsScale

`func (o *UpdateDashboardGraphRequestVisualization) SetBoundsScale(v string)`

SetBoundsScale sets BoundsScale field to given value.

### HasBoundsScale

`func (o *UpdateDashboardGraphRequestVisualization) HasBoundsScale() bool`

HasBoundsScale returns a boolean if a field has been set.

### GetPalette

`func (o *UpdateDashboardGraphRequestVisualization) GetPalette() string`

GetPalette returns the Palette field if non-nil, zero value otherwise.

### GetPaletteOk

`func (o *UpdateDashboardGraphRequestVisualization) GetPaletteOk() (*string, bool)`

GetPaletteOk returns a tuple with the Palette field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPalette

`func (o *UpdateDashboardGraphRequestVisualization) SetPalette(v string)`

SetPalette sets Palette field to given value.

### HasPalette

`func (o *UpdateDashboardGraphRequestVisualization) HasPalette() bool`

HasPalette returns a boolean if a field has been set.

### GetColumns

`func (o *UpdateDashboardGraphRequestVisualization) GetColumns() []InputTableColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *UpdateDashboardGraphRequestVisualization) GetColumnsOk() (*[]InputTableColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *UpdateDashboardGraphRequestVisualization) SetColumns(v []InputTableColumn)`

SetColumns sets Columns field to given value.


### GetColumnSizes

`func (o *UpdateDashboardGraphRequestVisualization) GetColumnSizes() map[string]float32`

GetColumnSizes returns the ColumnSizes field if non-nil, zero value otherwise.

### GetColumnSizesOk

`func (o *UpdateDashboardGraphRequestVisualization) GetColumnSizesOk() (*map[string]float32, bool)`

GetColumnSizesOk returns a tuple with the ColumnSizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnSizes

`func (o *UpdateDashboardGraphRequestVisualization) SetColumnSizes(v map[string]float32)`

SetColumnSizes sets ColumnSizes field to given value.

### HasColumnSizes

`func (o *UpdateDashboardGraphRequestVisualization) HasColumnSizes() bool`

HasColumnSizes returns a boolean if a field has been set.

### GetNote

`func (o *UpdateDashboardGraphRequestVisualization) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *UpdateDashboardGraphRequestVisualization) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *UpdateDashboardGraphRequestVisualization) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *UpdateDashboardGraphRequestVisualization) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetNoteColor

`func (o *UpdateDashboardGraphRequestVisualization) GetNoteColor() string`

GetNoteColor returns the NoteColor field if non-nil, zero value otherwise.

### GetNoteColorOk

`func (o *UpdateDashboardGraphRequestVisualization) GetNoteColorOk() (*string, bool)`

GetNoteColorOk returns a tuple with the NoteColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteColor

`func (o *UpdateDashboardGraphRequestVisualization) SetNoteColor(v string)`

SetNoteColor sets NoteColor field to given value.

### HasNoteColor

`func (o *UpdateDashboardGraphRequestVisualization) HasNoteColor() bool`

HasNoteColor returns a boolean if a field has been set.

### GetNoteAlign

`func (o *UpdateDashboardGraphRequestVisualization) GetNoteAlign() string`

GetNoteAlign returns the NoteAlign field if non-nil, zero value otherwise.

### GetNoteAlignOk

`func (o *UpdateDashboardGraphRequestVisualization) GetNoteAlignOk() (*string, bool)`

GetNoteAlignOk returns a tuple with the NoteAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteAlign

`func (o *UpdateDashboardGraphRequestVisualization) SetNoteAlign(v string)`

SetNoteAlign sets NoteAlign field to given value.

### HasNoteAlign

`func (o *UpdateDashboardGraphRequestVisualization) HasNoteAlign() bool`

HasNoteAlign returns a boolean if a field has been set.

### GetNoteJustifyContent

`func (o *UpdateDashboardGraphRequestVisualization) GetNoteJustifyContent() string`

GetNoteJustifyContent returns the NoteJustifyContent field if non-nil, zero value otherwise.

### GetNoteJustifyContentOk

`func (o *UpdateDashboardGraphRequestVisualization) GetNoteJustifyContentOk() (*string, bool)`

GetNoteJustifyContentOk returns a tuple with the NoteJustifyContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteJustifyContent

`func (o *UpdateDashboardGraphRequestVisualization) SetNoteJustifyContent(v string)`

SetNoteJustifyContent sets NoteJustifyContent field to given value.

### HasNoteJustifyContent

`func (o *UpdateDashboardGraphRequestVisualization) HasNoteJustifyContent() bool`

HasNoteJustifyContent returns a boolean if a field has been set.

### GetLayout

`func (o *UpdateDashboardGraphRequestVisualization) GetLayout() string`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *UpdateDashboardGraphRequestVisualization) GetLayoutOk() (*string, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *UpdateDashboardGraphRequestVisualization) SetLayout(v string)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *UpdateDashboardGraphRequestVisualization) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetSloId

`func (o *UpdateDashboardGraphRequestVisualization) GetSloId() string`

GetSloId returns the SloId field if non-nil, zero value otherwise.

### GetSloIdOk

`func (o *UpdateDashboardGraphRequestVisualization) GetSloIdOk() (*string, bool)`

GetSloIdOk returns a tuple with the SloId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloId

`func (o *UpdateDashboardGraphRequestVisualization) SetSloId(v string)`

SetSloId sets SloId field to given value.


### GetFilters

`func (o *UpdateDashboardGraphRequestVisualization) GetFilters() InputGraphVisualizationSloFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *UpdateDashboardGraphRequestVisualization) GetFiltersOk() (*InputGraphVisualizationSloFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *UpdateDashboardGraphRequestVisualization) SetFilters(v InputGraphVisualizationSloFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *UpdateDashboardGraphRequestVisualization) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetOrganizeBy

`func (o *UpdateDashboardGraphRequestVisualization) GetOrganizeBy() string`

GetOrganizeBy returns the OrganizeBy field if non-nil, zero value otherwise.

### GetOrganizeByOk

`func (o *UpdateDashboardGraphRequestVisualization) GetOrganizeByOk() (*string, bool)`

GetOrganizeByOk returns a tuple with the OrganizeBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizeBy

`func (o *UpdateDashboardGraphRequestVisualization) SetOrganizeBy(v string)`

SetOrganizeBy sets OrganizeBy field to given value.

### HasOrganizeBy

`func (o *UpdateDashboardGraphRequestVisualization) HasOrganizeBy() bool`

HasOrganizeBy returns a boolean if a field has been set.

### GetDisplay

`func (o *UpdateDashboardGraphRequestVisualization) GetDisplay() []string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *UpdateDashboardGraphRequestVisualization) GetDisplayOk() (*[]string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *UpdateDashboardGraphRequestVisualization) SetDisplay(v []string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


