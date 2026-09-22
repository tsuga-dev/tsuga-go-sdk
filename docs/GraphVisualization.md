# GraphVisualization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays the database rows-based aggregation as a time series chart | 
**ConnectionId** | **string** | The ID of the connection to use to query the datastore. | 
**Queries** | [**[]AggregationQuery1**](AggregationQuery1.md) | Aggregations that may be combined together in the same query. Each item is referenced from &#x60;formula&#x60; as q1, q2, and so on, in submission order. For dataSource \&quot;metrics\&quot;, each aggregate&#39;s &#x60;field&#x60; is the metric name, not an attribute; to count distinct values of an attribute use unique-count with field \&quot;&lt;metricName&gt;.context.&lt;attribute&gt;\&quot; (e.g. \&quot;system.cpu.utilization.context.host.name\&quot;). | 
**LegendMode** | Pointer to **string** | Controls whether and how the widget displays legend or series details (e.g. table, legend-only, or no legend) | [optional] 
**Thresholds** | Pointer to [**[]ThresholdMarker**](ThresholdMarker.md) | Threshold markers displayed on the chart | [optional] 
**YAxisSettings** | Pointer to [**InputGraphVisualizationTimeseriesConnectionYAxisSettings**](InputGraphVisualizationTimeseriesConnectionYAxisSettings.md) |  | [optional] 
**LineStyleOptions** | Pointer to [**map[string]GraphVisualizationTimeseriesConnectionLineStyleOptionsValue**](GraphVisualizationTimeseriesConnectionLineStyleOptionsValue.md) | Line style of each series, keyed by 1-based query index. The last index is the formula when there is one. For widgets with a single query, only the &#x60;1&#x60; entry is read and it applies to every series. Defaults to regular. | [optional] 
**Query** | **string** | Query that selects trace spans for the list | 
**ListColumns** | Pointer to [**[]WidgetListColumn1**](WidgetListColumn1.md) | Custom columns to display for each span | [optional] 
**ListColumnsSize** | Pointer to **map[string]float32** | List column widths in pixels, keyed by the &#x60;attribute&#x60; of the matching &#x60;listColumns&#x60; entry. Columns without an entry keep their default width. | [optional] 
**IsCellWrapped** | Pointer to **bool** | Whether list widget cell text wraps instead of truncating. Applies to span list widgets. Optional; omit or set false to use truncated cells. | [optional] 
**DefaultSorting** | Pointer to [**[]ListDefaultSorting1**](ListDefaultSorting1.md) | Default sorting applied to a list widget. Optional on create or update for log, span, or connection list widgets. Users can still change sorting by selecting columns in the rendered list. | [optional] 
**BackgroundMode** | Pointer to **string** | Controls whether the widget uses a solid or transparent background | [optional] 
**Conditions** | Pointer to [**[]ConditionalFormatting**](ConditionalFormatting.md) | Conditional formatting rules applied to the displayed value | [optional] 
**Precision** | Pointer to [**GraphVisualizationQueryValueConnectionPrecision**](GraphVisualizationQueryValueConnectionPrecision.md) |  | [optional] 
**Normalizer** | Pointer to [**Normalizer1**](Normalizer1.md) |  | [optional] 
**Aliases** | Pointer to [**InputGraphVisualizationTimeseriesPromqlAliases**](InputGraphVisualizationTimeseriesPromqlAliases.md) |  | [optional] 
**TimeBucket** | Pointer to [**GraphVisualizationTimeseriesPromqlTimeBucket**](GraphVisualizationTimeseriesPromqlTimeBucket.md) |  | [optional] 
**Smoothing** | Pointer to **bool** | Whether to apply automatic smoothing to the rendered timeseries | [optional] 
**Source** | **string** | Telemetry source queried by this aggregation: &#x60;logs&#x60;, &#x60;metrics&#x60;, &#x60;traces&#x60;, or &#x60;rum&#x60;. | 
**Formula** | Pointer to **string** | Formula referencing query outputs, such as &#x60;q1 + q2&#x60;, to compute derived results. Defaults to &#x60;q1&#x60;. Formulas may reference only submitted queries (&#x60;q1&#x60; through &#x60;qN&#x60;); undefined query references return 400. | [optional] 
**VisibleSeries** | Pointer to **[]bool** | Flags indicating whether each query or formula series is visible | [optional] 
**GroupBy** | Pointer to [**[]AggregationGroupBy**](AggregationGroupBy.md) | Nested grouping levels applied to aggregation results, outermost first (e.g. group by service, then by level within each service). Each level splits results further, so the response contains one result per unique combination of group values instead of one aggregated total. Defaults to an empty array (one ungrouped result) when omitted. | [optional] 
**GroupByMode** | Pointer to **string** | &#x60;absolute&#x60; keeps each group at its own value; &#x60;relative&#x60; shows it as a percentage of the ungrouped total (defaults to absolute) | [optional] 
**IsStacked** | Pointer to **bool** | Requests stacked rendering for a top-list widget. Tsuga renders stacked rows only for one count or sum query with exactly two grouped fields, no formula, non-negative values, and a single-cluster context; otherwise the widget renders as a normal top list. | [optional] 
**Max** | Pointer to **float32** | Gauge maximum value | [optional] 
**ColorThresholds** | Pointer to [**[]GaugeColorThreshold**](GaugeColorThreshold.md) | Color thresholds inside the gauge range | [optional] 
**Group** | Pointer to **string** | Attribute that switches the count to \&quot;Groups\&quot; mode: records are grouped by this attribute, the aggregation produces one value per group, and the chart buckets those per-group values. When omitted, individual records are bucketed. | [optional] 
**PercentileMarkers** | Pointer to **[]int32** | Percentile markers displayed on top of the distribution chart | [optional] 
**BoundsScale** | Pointer to **string** | Spacing of the bucket boundaries across the distribution range. &#x60;linear&#x60; splits the range into equal-width buckets; &#x60;log&#x60; widens each bucket logarithmically, giving finer resolution near the lower bound. | [optional] 
**Palette** | Pointer to **string** | Color palette used to render the heatmap intensity gradient | [optional] 
**Columns** | [**[]TableColumn1**](TableColumn1.md) | Each column defines an independent aggregation displayed as a table column | 
**ColumnSizes** | Pointer to **map[string]float32** | Table column widths in pixels, keyed by column id: &#x60;label&#x60; for the grouping column and &#x60;col-&lt;index&gt;&#x60; for each entry in &#x60;columns&#x60;. Columns without an entry keep their default width. | [optional] 
**Note** | Pointer to **string** | Markdown-compatible text shown in the note | [optional] 
**NoteColor** | Pointer to **string** | Background color used to render the note | [optional] 
**NoteAlign** | Pointer to **string** | Flex alignment keyword used for widget layout | [optional] 
**NoteJustifyContent** | Pointer to **string** | Flex alignment keyword used for widget layout | [optional] 
**Layout** | Pointer to **string** | Layout used to render log patterns | [optional] 
**SloId** | **string** | Id of the SLO to display | 

## Methods

### NewGraphVisualization

`func NewGraphVisualization(type_ string, connectionId string, queries []AggregationQuery1, query string, source string, columns []TableColumn1, sloId string, ) *GraphVisualization`

NewGraphVisualization instantiates a new GraphVisualization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphVisualizationWithDefaults

`func NewGraphVisualizationWithDefaults() *GraphVisualization`

NewGraphVisualizationWithDefaults instantiates a new GraphVisualization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GraphVisualization) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GraphVisualization) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GraphVisualization) SetType(v string)`

SetType sets Type field to given value.


### GetConnectionId

`func (o *GraphVisualization) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *GraphVisualization) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *GraphVisualization) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetQueries

`func (o *GraphVisualization) GetQueries() []AggregationQuery1`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *GraphVisualization) GetQueriesOk() (*[]AggregationQuery1, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *GraphVisualization) SetQueries(v []AggregationQuery1)`

SetQueries sets Queries field to given value.


### GetLegendMode

`func (o *GraphVisualization) GetLegendMode() string`

GetLegendMode returns the LegendMode field if non-nil, zero value otherwise.

### GetLegendModeOk

`func (o *GraphVisualization) GetLegendModeOk() (*string, bool)`

GetLegendModeOk returns a tuple with the LegendMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegendMode

`func (o *GraphVisualization) SetLegendMode(v string)`

SetLegendMode sets LegendMode field to given value.

### HasLegendMode

`func (o *GraphVisualization) HasLegendMode() bool`

HasLegendMode returns a boolean if a field has been set.

### GetThresholds

`func (o *GraphVisualization) GetThresholds() []ThresholdMarker`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *GraphVisualization) GetThresholdsOk() (*[]ThresholdMarker, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *GraphVisualization) SetThresholds(v []ThresholdMarker)`

SetThresholds sets Thresholds field to given value.

### HasThresholds

`func (o *GraphVisualization) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.

### GetYAxisSettings

`func (o *GraphVisualization) GetYAxisSettings() InputGraphVisualizationTimeseriesConnectionYAxisSettings`

GetYAxisSettings returns the YAxisSettings field if non-nil, zero value otherwise.

### GetYAxisSettingsOk

`func (o *GraphVisualization) GetYAxisSettingsOk() (*InputGraphVisualizationTimeseriesConnectionYAxisSettings, bool)`

GetYAxisSettingsOk returns a tuple with the YAxisSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYAxisSettings

`func (o *GraphVisualization) SetYAxisSettings(v InputGraphVisualizationTimeseriesConnectionYAxisSettings)`

SetYAxisSettings sets YAxisSettings field to given value.

### HasYAxisSettings

`func (o *GraphVisualization) HasYAxisSettings() bool`

HasYAxisSettings returns a boolean if a field has been set.

### GetLineStyleOptions

`func (o *GraphVisualization) GetLineStyleOptions() map[string]GraphVisualizationTimeseriesConnectionLineStyleOptionsValue`

GetLineStyleOptions returns the LineStyleOptions field if non-nil, zero value otherwise.

### GetLineStyleOptionsOk

`func (o *GraphVisualization) GetLineStyleOptionsOk() (*map[string]GraphVisualizationTimeseriesConnectionLineStyleOptionsValue, bool)`

GetLineStyleOptionsOk returns a tuple with the LineStyleOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineStyleOptions

`func (o *GraphVisualization) SetLineStyleOptions(v map[string]GraphVisualizationTimeseriesConnectionLineStyleOptionsValue)`

SetLineStyleOptions sets LineStyleOptions field to given value.

### HasLineStyleOptions

`func (o *GraphVisualization) HasLineStyleOptions() bool`

HasLineStyleOptions returns a boolean if a field has been set.

### GetQuery

`func (o *GraphVisualization) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *GraphVisualization) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *GraphVisualization) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetListColumns

`func (o *GraphVisualization) GetListColumns() []WidgetListColumn1`

GetListColumns returns the ListColumns field if non-nil, zero value otherwise.

### GetListColumnsOk

`func (o *GraphVisualization) GetListColumnsOk() (*[]WidgetListColumn1, bool)`

GetListColumnsOk returns a tuple with the ListColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListColumns

`func (o *GraphVisualization) SetListColumns(v []WidgetListColumn1)`

SetListColumns sets ListColumns field to given value.

### HasListColumns

`func (o *GraphVisualization) HasListColumns() bool`

HasListColumns returns a boolean if a field has been set.

### GetListColumnsSize

`func (o *GraphVisualization) GetListColumnsSize() map[string]float32`

GetListColumnsSize returns the ListColumnsSize field if non-nil, zero value otherwise.

### GetListColumnsSizeOk

`func (o *GraphVisualization) GetListColumnsSizeOk() (*map[string]float32, bool)`

GetListColumnsSizeOk returns a tuple with the ListColumnsSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListColumnsSize

`func (o *GraphVisualization) SetListColumnsSize(v map[string]float32)`

SetListColumnsSize sets ListColumnsSize field to given value.

### HasListColumnsSize

`func (o *GraphVisualization) HasListColumnsSize() bool`

HasListColumnsSize returns a boolean if a field has been set.

### GetIsCellWrapped

`func (o *GraphVisualization) GetIsCellWrapped() bool`

GetIsCellWrapped returns the IsCellWrapped field if non-nil, zero value otherwise.

### GetIsCellWrappedOk

`func (o *GraphVisualization) GetIsCellWrappedOk() (*bool, bool)`

GetIsCellWrappedOk returns a tuple with the IsCellWrapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCellWrapped

`func (o *GraphVisualization) SetIsCellWrapped(v bool)`

SetIsCellWrapped sets IsCellWrapped field to given value.

### HasIsCellWrapped

`func (o *GraphVisualization) HasIsCellWrapped() bool`

HasIsCellWrapped returns a boolean if a field has been set.

### GetDefaultSorting

`func (o *GraphVisualization) GetDefaultSorting() []ListDefaultSorting1`

GetDefaultSorting returns the DefaultSorting field if non-nil, zero value otherwise.

### GetDefaultSortingOk

`func (o *GraphVisualization) GetDefaultSortingOk() (*[]ListDefaultSorting1, bool)`

GetDefaultSortingOk returns a tuple with the DefaultSorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSorting

`func (o *GraphVisualization) SetDefaultSorting(v []ListDefaultSorting1)`

SetDefaultSorting sets DefaultSorting field to given value.

### HasDefaultSorting

`func (o *GraphVisualization) HasDefaultSorting() bool`

HasDefaultSorting returns a boolean if a field has been set.

### GetBackgroundMode

`func (o *GraphVisualization) GetBackgroundMode() string`

GetBackgroundMode returns the BackgroundMode field if non-nil, zero value otherwise.

### GetBackgroundModeOk

`func (o *GraphVisualization) GetBackgroundModeOk() (*string, bool)`

GetBackgroundModeOk returns a tuple with the BackgroundMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundMode

`func (o *GraphVisualization) SetBackgroundMode(v string)`

SetBackgroundMode sets BackgroundMode field to given value.

### HasBackgroundMode

`func (o *GraphVisualization) HasBackgroundMode() bool`

HasBackgroundMode returns a boolean if a field has been set.

### GetConditions

`func (o *GraphVisualization) GetConditions() []ConditionalFormatting`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *GraphVisualization) GetConditionsOk() (*[]ConditionalFormatting, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *GraphVisualization) SetConditions(v []ConditionalFormatting)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *GraphVisualization) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetPrecision

`func (o *GraphVisualization) GetPrecision() GraphVisualizationQueryValueConnectionPrecision`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *GraphVisualization) GetPrecisionOk() (*GraphVisualizationQueryValueConnectionPrecision, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *GraphVisualization) SetPrecision(v GraphVisualizationQueryValueConnectionPrecision)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *GraphVisualization) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetNormalizer

`func (o *GraphVisualization) GetNormalizer() Normalizer1`

GetNormalizer returns the Normalizer field if non-nil, zero value otherwise.

### GetNormalizerOk

`func (o *GraphVisualization) GetNormalizerOk() (*Normalizer1, bool)`

GetNormalizerOk returns a tuple with the Normalizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizer

`func (o *GraphVisualization) SetNormalizer(v Normalizer1)`

SetNormalizer sets Normalizer field to given value.

### HasNormalizer

`func (o *GraphVisualization) HasNormalizer() bool`

HasNormalizer returns a boolean if a field has been set.

### GetAliases

`func (o *GraphVisualization) GetAliases() InputGraphVisualizationTimeseriesPromqlAliases`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *GraphVisualization) GetAliasesOk() (*InputGraphVisualizationTimeseriesPromqlAliases, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *GraphVisualization) SetAliases(v InputGraphVisualizationTimeseriesPromqlAliases)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *GraphVisualization) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetTimeBucket

`func (o *GraphVisualization) GetTimeBucket() GraphVisualizationTimeseriesPromqlTimeBucket`

GetTimeBucket returns the TimeBucket field if non-nil, zero value otherwise.

### GetTimeBucketOk

`func (o *GraphVisualization) GetTimeBucketOk() (*GraphVisualizationTimeseriesPromqlTimeBucket, bool)`

GetTimeBucketOk returns a tuple with the TimeBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBucket

`func (o *GraphVisualization) SetTimeBucket(v GraphVisualizationTimeseriesPromqlTimeBucket)`

SetTimeBucket sets TimeBucket field to given value.

### HasTimeBucket

`func (o *GraphVisualization) HasTimeBucket() bool`

HasTimeBucket returns a boolean if a field has been set.

### GetSmoothing

`func (o *GraphVisualization) GetSmoothing() bool`

GetSmoothing returns the Smoothing field if non-nil, zero value otherwise.

### GetSmoothingOk

`func (o *GraphVisualization) GetSmoothingOk() (*bool, bool)`

GetSmoothingOk returns a tuple with the Smoothing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmoothing

`func (o *GraphVisualization) SetSmoothing(v bool)`

SetSmoothing sets Smoothing field to given value.

### HasSmoothing

`func (o *GraphVisualization) HasSmoothing() bool`

HasSmoothing returns a boolean if a field has been set.

### GetSource

`func (o *GraphVisualization) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GraphVisualization) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GraphVisualization) SetSource(v string)`

SetSource sets Source field to given value.


### GetFormula

`func (o *GraphVisualization) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *GraphVisualization) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *GraphVisualization) SetFormula(v string)`

SetFormula sets Formula field to given value.

### HasFormula

`func (o *GraphVisualization) HasFormula() bool`

HasFormula returns a boolean if a field has been set.

### GetVisibleSeries

`func (o *GraphVisualization) GetVisibleSeries() []bool`

GetVisibleSeries returns the VisibleSeries field if non-nil, zero value otherwise.

### GetVisibleSeriesOk

`func (o *GraphVisualization) GetVisibleSeriesOk() (*[]bool, bool)`

GetVisibleSeriesOk returns a tuple with the VisibleSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleSeries

`func (o *GraphVisualization) SetVisibleSeries(v []bool)`

SetVisibleSeries sets VisibleSeries field to given value.

### HasVisibleSeries

`func (o *GraphVisualization) HasVisibleSeries() bool`

HasVisibleSeries returns a boolean if a field has been set.

### GetGroupBy

`func (o *GraphVisualization) GetGroupBy() []AggregationGroupBy`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *GraphVisualization) GetGroupByOk() (*[]AggregationGroupBy, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *GraphVisualization) SetGroupBy(v []AggregationGroupBy)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *GraphVisualization) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetGroupByMode

`func (o *GraphVisualization) GetGroupByMode() string`

GetGroupByMode returns the GroupByMode field if non-nil, zero value otherwise.

### GetGroupByModeOk

`func (o *GraphVisualization) GetGroupByModeOk() (*string, bool)`

GetGroupByModeOk returns a tuple with the GroupByMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByMode

`func (o *GraphVisualization) SetGroupByMode(v string)`

SetGroupByMode sets GroupByMode field to given value.

### HasGroupByMode

`func (o *GraphVisualization) HasGroupByMode() bool`

HasGroupByMode returns a boolean if a field has been set.

### GetIsStacked

`func (o *GraphVisualization) GetIsStacked() bool`

GetIsStacked returns the IsStacked field if non-nil, zero value otherwise.

### GetIsStackedOk

`func (o *GraphVisualization) GetIsStackedOk() (*bool, bool)`

GetIsStackedOk returns a tuple with the IsStacked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStacked

`func (o *GraphVisualization) SetIsStacked(v bool)`

SetIsStacked sets IsStacked field to given value.

### HasIsStacked

`func (o *GraphVisualization) HasIsStacked() bool`

HasIsStacked returns a boolean if a field has been set.

### GetMax

`func (o *GraphVisualization) GetMax() float32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *GraphVisualization) GetMaxOk() (*float32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *GraphVisualization) SetMax(v float32)`

SetMax sets Max field to given value.

### HasMax

`func (o *GraphVisualization) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetColorThresholds

`func (o *GraphVisualization) GetColorThresholds() []GaugeColorThreshold`

GetColorThresholds returns the ColorThresholds field if non-nil, zero value otherwise.

### GetColorThresholdsOk

`func (o *GraphVisualization) GetColorThresholdsOk() (*[]GaugeColorThreshold, bool)`

GetColorThresholdsOk returns a tuple with the ColorThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorThresholds

`func (o *GraphVisualization) SetColorThresholds(v []GaugeColorThreshold)`

SetColorThresholds sets ColorThresholds field to given value.

### HasColorThresholds

`func (o *GraphVisualization) HasColorThresholds() bool`

HasColorThresholds returns a boolean if a field has been set.

### GetGroup

`func (o *GraphVisualization) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GraphVisualization) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GraphVisualization) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *GraphVisualization) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetPercentileMarkers

`func (o *GraphVisualization) GetPercentileMarkers() []int32`

GetPercentileMarkers returns the PercentileMarkers field if non-nil, zero value otherwise.

### GetPercentileMarkersOk

`func (o *GraphVisualization) GetPercentileMarkersOk() (*[]int32, bool)`

GetPercentileMarkersOk returns a tuple with the PercentileMarkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentileMarkers

`func (o *GraphVisualization) SetPercentileMarkers(v []int32)`

SetPercentileMarkers sets PercentileMarkers field to given value.

### HasPercentileMarkers

`func (o *GraphVisualization) HasPercentileMarkers() bool`

HasPercentileMarkers returns a boolean if a field has been set.

### GetBoundsScale

`func (o *GraphVisualization) GetBoundsScale() string`

GetBoundsScale returns the BoundsScale field if non-nil, zero value otherwise.

### GetBoundsScaleOk

`func (o *GraphVisualization) GetBoundsScaleOk() (*string, bool)`

GetBoundsScaleOk returns a tuple with the BoundsScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundsScale

`func (o *GraphVisualization) SetBoundsScale(v string)`

SetBoundsScale sets BoundsScale field to given value.

### HasBoundsScale

`func (o *GraphVisualization) HasBoundsScale() bool`

HasBoundsScale returns a boolean if a field has been set.

### GetPalette

`func (o *GraphVisualization) GetPalette() string`

GetPalette returns the Palette field if non-nil, zero value otherwise.

### GetPaletteOk

`func (o *GraphVisualization) GetPaletteOk() (*string, bool)`

GetPaletteOk returns a tuple with the Palette field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPalette

`func (o *GraphVisualization) SetPalette(v string)`

SetPalette sets Palette field to given value.

### HasPalette

`func (o *GraphVisualization) HasPalette() bool`

HasPalette returns a boolean if a field has been set.

### GetColumns

`func (o *GraphVisualization) GetColumns() []TableColumn1`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *GraphVisualization) GetColumnsOk() (*[]TableColumn1, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *GraphVisualization) SetColumns(v []TableColumn1)`

SetColumns sets Columns field to given value.


### GetColumnSizes

`func (o *GraphVisualization) GetColumnSizes() map[string]float32`

GetColumnSizes returns the ColumnSizes field if non-nil, zero value otherwise.

### GetColumnSizesOk

`func (o *GraphVisualization) GetColumnSizesOk() (*map[string]float32, bool)`

GetColumnSizesOk returns a tuple with the ColumnSizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnSizes

`func (o *GraphVisualization) SetColumnSizes(v map[string]float32)`

SetColumnSizes sets ColumnSizes field to given value.

### HasColumnSizes

`func (o *GraphVisualization) HasColumnSizes() bool`

HasColumnSizes returns a boolean if a field has been set.

### GetNote

`func (o *GraphVisualization) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *GraphVisualization) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *GraphVisualization) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *GraphVisualization) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetNoteColor

`func (o *GraphVisualization) GetNoteColor() string`

GetNoteColor returns the NoteColor field if non-nil, zero value otherwise.

### GetNoteColorOk

`func (o *GraphVisualization) GetNoteColorOk() (*string, bool)`

GetNoteColorOk returns a tuple with the NoteColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteColor

`func (o *GraphVisualization) SetNoteColor(v string)`

SetNoteColor sets NoteColor field to given value.

### HasNoteColor

`func (o *GraphVisualization) HasNoteColor() bool`

HasNoteColor returns a boolean if a field has been set.

### GetNoteAlign

`func (o *GraphVisualization) GetNoteAlign() string`

GetNoteAlign returns the NoteAlign field if non-nil, zero value otherwise.

### GetNoteAlignOk

`func (o *GraphVisualization) GetNoteAlignOk() (*string, bool)`

GetNoteAlignOk returns a tuple with the NoteAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteAlign

`func (o *GraphVisualization) SetNoteAlign(v string)`

SetNoteAlign sets NoteAlign field to given value.

### HasNoteAlign

`func (o *GraphVisualization) HasNoteAlign() bool`

HasNoteAlign returns a boolean if a field has been set.

### GetNoteJustifyContent

`func (o *GraphVisualization) GetNoteJustifyContent() string`

GetNoteJustifyContent returns the NoteJustifyContent field if non-nil, zero value otherwise.

### GetNoteJustifyContentOk

`func (o *GraphVisualization) GetNoteJustifyContentOk() (*string, bool)`

GetNoteJustifyContentOk returns a tuple with the NoteJustifyContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteJustifyContent

`func (o *GraphVisualization) SetNoteJustifyContent(v string)`

SetNoteJustifyContent sets NoteJustifyContent field to given value.

### HasNoteJustifyContent

`func (o *GraphVisualization) HasNoteJustifyContent() bool`

HasNoteJustifyContent returns a boolean if a field has been set.

### GetLayout

`func (o *GraphVisualization) GetLayout() string`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *GraphVisualization) GetLayoutOk() (*string, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *GraphVisualization) SetLayout(v string)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *GraphVisualization) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetSloId

`func (o *GraphVisualization) GetSloId() string`

GetSloId returns the SloId field if non-nil, zero value otherwise.

### GetSloIdOk

`func (o *GraphVisualization) GetSloIdOk() (*string, bool)`

GetSloIdOk returns a tuple with the SloId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloId

`func (o *GraphVisualization) SetSloId(v string)`

SetSloId sets SloId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


