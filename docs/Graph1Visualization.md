# Graph1Visualization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays the database rows-based aggregation as a time series chart | 
**ConnectionId** | **string** | The ID of the connection to use to query the datastore. | 
**Queries** | [**[]AggregationQuery**](AggregationQuery.md) | Aggregations that may be combined together in the same query | 
**LegendMode** | Pointer to **string** | Controls whether and how the widget displays legend or series details (e.g. table, legend-only, or no legend) | [optional] 
**Thresholds** | Pointer to [**[]ThresholdMarker**](ThresholdMarker.md) | Threshold markers displayed on the chart | [optional] 
**YAxisSettings** | Pointer to [**GraphVisualizationTimeseriesConnectionYAxisSettings**](GraphVisualizationTimeseriesConnectionYAxisSettings.md) |  | [optional] 
**Query** | **string** | Tsuga query that selects logs to cluster into patterns | 
**ListColumns** | Pointer to [**[]WidgetListColumn**](WidgetListColumn.md) | Custom columns to display for each log | [optional] 
**ListColumnsSize** | Pointer to **map[string]float32** |  | [optional] 
**BackgroundMode** | Pointer to **string** | Controls whether the widget uses a solid or transparent background | [optional] 
**Conditions** | Pointer to [**[]ConditionalFormatting**](ConditionalFormatting.md) | Conditional formatting rules applied to the displayed value | [optional] 
**Precision** | Pointer to **float32** | Number of decimal places to display in the value | [optional] 
**Normalizer** | Pointer to [**Normalizer**](Normalizer.md) |  | [optional] 
**Source** | **string** | Data source being queried for this aggregation | 
**Formula** | Pointer to **string** | Formula referencing query outputs (e.g. q1+q2) to compute derived series | [optional] 
**Aliases** | Pointer to [**GraphVisualizationTimeseriesAliases**](GraphVisualizationTimeseriesAliases.md) |  | [optional] 
**VisibleSeries** | Pointer to **[]bool** | Flags indicating whether each query or formula series is visible | [optional] 
**GroupBy** | Pointer to [**[]AggregationGroupBy**](AggregationGroupBy.md) | Fields used to group the results | [optional] 
**TimeBucket** | Pointer to [**GraphVisualizationTimeseriesTimeBucket**](GraphVisualizationTimeseriesTimeBucket.md) |  | [optional] 
**Smoothing** | Pointer to **bool** | Whether to apply automatic smoothing to the rendered timeseries | [optional] 
**PercentileMarkers** | Pointer to **[]int32** | Percentile markers displayed on top of the distribution chart | [optional] 
**Palette** | Pointer to **string** | Color palette used to render the heatmap intensity gradient | [optional] 
**Columns** | [**[]TableColumn**](TableColumn.md) | Each column defines an independent aggregation displayed as a table column | 
**DefaultSorting** | Pointer to [**[]TableDefaultSorting**](TableDefaultSorting.md) |  | [optional] 
**Note** | Pointer to **string** | Markdown-compatible text shown in the note | [optional] 
**NoteColor** | Pointer to **string** | Background color used to render the note | [optional] 
**NoteAlign** | Pointer to **string** | Flex alignment keyword used for widget layout | [optional] 
**NoteJustifyContent** | Pointer to **string** | Flex alignment keyword used for widget layout | [optional] 
**Layout** | Pointer to **string** | Layout used to render log patterns | [optional] 

## Methods

### NewGraph1Visualization

`func NewGraph1Visualization(type_ string, connectionId string, queries []AggregationQuery, query string, source string, columns []TableColumn, ) *Graph1Visualization`

NewGraph1Visualization instantiates a new Graph1Visualization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraph1VisualizationWithDefaults

`func NewGraph1VisualizationWithDefaults() *Graph1Visualization`

NewGraph1VisualizationWithDefaults instantiates a new Graph1Visualization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Graph1Visualization) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Graph1Visualization) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Graph1Visualization) SetType(v string)`

SetType sets Type field to given value.


### GetConnectionId

`func (o *Graph1Visualization) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *Graph1Visualization) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *Graph1Visualization) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetQueries

`func (o *Graph1Visualization) GetQueries() []AggregationQuery`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *Graph1Visualization) GetQueriesOk() (*[]AggregationQuery, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *Graph1Visualization) SetQueries(v []AggregationQuery)`

SetQueries sets Queries field to given value.


### GetLegendMode

`func (o *Graph1Visualization) GetLegendMode() string`

GetLegendMode returns the LegendMode field if non-nil, zero value otherwise.

### GetLegendModeOk

`func (o *Graph1Visualization) GetLegendModeOk() (*string, bool)`

GetLegendModeOk returns a tuple with the LegendMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegendMode

`func (o *Graph1Visualization) SetLegendMode(v string)`

SetLegendMode sets LegendMode field to given value.

### HasLegendMode

`func (o *Graph1Visualization) HasLegendMode() bool`

HasLegendMode returns a boolean if a field has been set.

### GetThresholds

`func (o *Graph1Visualization) GetThresholds() []ThresholdMarker`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *Graph1Visualization) GetThresholdsOk() (*[]ThresholdMarker, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *Graph1Visualization) SetThresholds(v []ThresholdMarker)`

SetThresholds sets Thresholds field to given value.

### HasThresholds

`func (o *Graph1Visualization) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.

### GetYAxisSettings

`func (o *Graph1Visualization) GetYAxisSettings() GraphVisualizationTimeseriesConnectionYAxisSettings`

GetYAxisSettings returns the YAxisSettings field if non-nil, zero value otherwise.

### GetYAxisSettingsOk

`func (o *Graph1Visualization) GetYAxisSettingsOk() (*GraphVisualizationTimeseriesConnectionYAxisSettings, bool)`

GetYAxisSettingsOk returns a tuple with the YAxisSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYAxisSettings

`func (o *Graph1Visualization) SetYAxisSettings(v GraphVisualizationTimeseriesConnectionYAxisSettings)`

SetYAxisSettings sets YAxisSettings field to given value.

### HasYAxisSettings

`func (o *Graph1Visualization) HasYAxisSettings() bool`

HasYAxisSettings returns a boolean if a field has been set.

### GetQuery

`func (o *Graph1Visualization) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Graph1Visualization) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Graph1Visualization) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetListColumns

`func (o *Graph1Visualization) GetListColumns() []WidgetListColumn`

GetListColumns returns the ListColumns field if non-nil, zero value otherwise.

### GetListColumnsOk

`func (o *Graph1Visualization) GetListColumnsOk() (*[]WidgetListColumn, bool)`

GetListColumnsOk returns a tuple with the ListColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListColumns

`func (o *Graph1Visualization) SetListColumns(v []WidgetListColumn)`

SetListColumns sets ListColumns field to given value.

### HasListColumns

`func (o *Graph1Visualization) HasListColumns() bool`

HasListColumns returns a boolean if a field has been set.

### GetListColumnsSize

`func (o *Graph1Visualization) GetListColumnsSize() map[string]float32`

GetListColumnsSize returns the ListColumnsSize field if non-nil, zero value otherwise.

### GetListColumnsSizeOk

`func (o *Graph1Visualization) GetListColumnsSizeOk() (*map[string]float32, bool)`

GetListColumnsSizeOk returns a tuple with the ListColumnsSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListColumnsSize

`func (o *Graph1Visualization) SetListColumnsSize(v map[string]float32)`

SetListColumnsSize sets ListColumnsSize field to given value.

### HasListColumnsSize

`func (o *Graph1Visualization) HasListColumnsSize() bool`

HasListColumnsSize returns a boolean if a field has been set.

### GetBackgroundMode

`func (o *Graph1Visualization) GetBackgroundMode() string`

GetBackgroundMode returns the BackgroundMode field if non-nil, zero value otherwise.

### GetBackgroundModeOk

`func (o *Graph1Visualization) GetBackgroundModeOk() (*string, bool)`

GetBackgroundModeOk returns a tuple with the BackgroundMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundMode

`func (o *Graph1Visualization) SetBackgroundMode(v string)`

SetBackgroundMode sets BackgroundMode field to given value.

### HasBackgroundMode

`func (o *Graph1Visualization) HasBackgroundMode() bool`

HasBackgroundMode returns a boolean if a field has been set.

### GetConditions

`func (o *Graph1Visualization) GetConditions() []ConditionalFormatting`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *Graph1Visualization) GetConditionsOk() (*[]ConditionalFormatting, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *Graph1Visualization) SetConditions(v []ConditionalFormatting)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *Graph1Visualization) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetPrecision

`func (o *Graph1Visualization) GetPrecision() float32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *Graph1Visualization) GetPrecisionOk() (*float32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *Graph1Visualization) SetPrecision(v float32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *Graph1Visualization) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetNormalizer

`func (o *Graph1Visualization) GetNormalizer() Normalizer`

GetNormalizer returns the Normalizer field if non-nil, zero value otherwise.

### GetNormalizerOk

`func (o *Graph1Visualization) GetNormalizerOk() (*Normalizer, bool)`

GetNormalizerOk returns a tuple with the Normalizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizer

`func (o *Graph1Visualization) SetNormalizer(v Normalizer)`

SetNormalizer sets Normalizer field to given value.

### HasNormalizer

`func (o *Graph1Visualization) HasNormalizer() bool`

HasNormalizer returns a boolean if a field has been set.

### GetSource

`func (o *Graph1Visualization) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Graph1Visualization) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Graph1Visualization) SetSource(v string)`

SetSource sets Source field to given value.


### GetFormula

`func (o *Graph1Visualization) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *Graph1Visualization) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *Graph1Visualization) SetFormula(v string)`

SetFormula sets Formula field to given value.

### HasFormula

`func (o *Graph1Visualization) HasFormula() bool`

HasFormula returns a boolean if a field has been set.

### GetAliases

`func (o *Graph1Visualization) GetAliases() GraphVisualizationTimeseriesAliases`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *Graph1Visualization) GetAliasesOk() (*GraphVisualizationTimeseriesAliases, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *Graph1Visualization) SetAliases(v GraphVisualizationTimeseriesAliases)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *Graph1Visualization) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetVisibleSeries

`func (o *Graph1Visualization) GetVisibleSeries() []bool`

GetVisibleSeries returns the VisibleSeries field if non-nil, zero value otherwise.

### GetVisibleSeriesOk

`func (o *Graph1Visualization) GetVisibleSeriesOk() (*[]bool, bool)`

GetVisibleSeriesOk returns a tuple with the VisibleSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleSeries

`func (o *Graph1Visualization) SetVisibleSeries(v []bool)`

SetVisibleSeries sets VisibleSeries field to given value.

### HasVisibleSeries

`func (o *Graph1Visualization) HasVisibleSeries() bool`

HasVisibleSeries returns a boolean if a field has been set.

### GetGroupBy

`func (o *Graph1Visualization) GetGroupBy() []AggregationGroupBy`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *Graph1Visualization) GetGroupByOk() (*[]AggregationGroupBy, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *Graph1Visualization) SetGroupBy(v []AggregationGroupBy)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *Graph1Visualization) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetTimeBucket

`func (o *Graph1Visualization) GetTimeBucket() GraphVisualizationTimeseriesTimeBucket`

GetTimeBucket returns the TimeBucket field if non-nil, zero value otherwise.

### GetTimeBucketOk

`func (o *Graph1Visualization) GetTimeBucketOk() (*GraphVisualizationTimeseriesTimeBucket, bool)`

GetTimeBucketOk returns a tuple with the TimeBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBucket

`func (o *Graph1Visualization) SetTimeBucket(v GraphVisualizationTimeseriesTimeBucket)`

SetTimeBucket sets TimeBucket field to given value.

### HasTimeBucket

`func (o *Graph1Visualization) HasTimeBucket() bool`

HasTimeBucket returns a boolean if a field has been set.

### GetSmoothing

`func (o *Graph1Visualization) GetSmoothing() bool`

GetSmoothing returns the Smoothing field if non-nil, zero value otherwise.

### GetSmoothingOk

`func (o *Graph1Visualization) GetSmoothingOk() (*bool, bool)`

GetSmoothingOk returns a tuple with the Smoothing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmoothing

`func (o *Graph1Visualization) SetSmoothing(v bool)`

SetSmoothing sets Smoothing field to given value.

### HasSmoothing

`func (o *Graph1Visualization) HasSmoothing() bool`

HasSmoothing returns a boolean if a field has been set.

### GetPercentileMarkers

`func (o *Graph1Visualization) GetPercentileMarkers() []int32`

GetPercentileMarkers returns the PercentileMarkers field if non-nil, zero value otherwise.

### GetPercentileMarkersOk

`func (o *Graph1Visualization) GetPercentileMarkersOk() (*[]int32, bool)`

GetPercentileMarkersOk returns a tuple with the PercentileMarkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentileMarkers

`func (o *Graph1Visualization) SetPercentileMarkers(v []int32)`

SetPercentileMarkers sets PercentileMarkers field to given value.

### HasPercentileMarkers

`func (o *Graph1Visualization) HasPercentileMarkers() bool`

HasPercentileMarkers returns a boolean if a field has been set.

### GetPalette

`func (o *Graph1Visualization) GetPalette() string`

GetPalette returns the Palette field if non-nil, zero value otherwise.

### GetPaletteOk

`func (o *Graph1Visualization) GetPaletteOk() (*string, bool)`

GetPaletteOk returns a tuple with the Palette field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPalette

`func (o *Graph1Visualization) SetPalette(v string)`

SetPalette sets Palette field to given value.

### HasPalette

`func (o *Graph1Visualization) HasPalette() bool`

HasPalette returns a boolean if a field has been set.

### GetColumns

`func (o *Graph1Visualization) GetColumns() []TableColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *Graph1Visualization) GetColumnsOk() (*[]TableColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *Graph1Visualization) SetColumns(v []TableColumn)`

SetColumns sets Columns field to given value.


### GetDefaultSorting

`func (o *Graph1Visualization) GetDefaultSorting() []TableDefaultSorting`

GetDefaultSorting returns the DefaultSorting field if non-nil, zero value otherwise.

### GetDefaultSortingOk

`func (o *Graph1Visualization) GetDefaultSortingOk() (*[]TableDefaultSorting, bool)`

GetDefaultSortingOk returns a tuple with the DefaultSorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSorting

`func (o *Graph1Visualization) SetDefaultSorting(v []TableDefaultSorting)`

SetDefaultSorting sets DefaultSorting field to given value.

### HasDefaultSorting

`func (o *Graph1Visualization) HasDefaultSorting() bool`

HasDefaultSorting returns a boolean if a field has been set.

### GetNote

`func (o *Graph1Visualization) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *Graph1Visualization) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *Graph1Visualization) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *Graph1Visualization) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetNoteColor

`func (o *Graph1Visualization) GetNoteColor() string`

GetNoteColor returns the NoteColor field if non-nil, zero value otherwise.

### GetNoteColorOk

`func (o *Graph1Visualization) GetNoteColorOk() (*string, bool)`

GetNoteColorOk returns a tuple with the NoteColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteColor

`func (o *Graph1Visualization) SetNoteColor(v string)`

SetNoteColor sets NoteColor field to given value.

### HasNoteColor

`func (o *Graph1Visualization) HasNoteColor() bool`

HasNoteColor returns a boolean if a field has been set.

### GetNoteAlign

`func (o *Graph1Visualization) GetNoteAlign() string`

GetNoteAlign returns the NoteAlign field if non-nil, zero value otherwise.

### GetNoteAlignOk

`func (o *Graph1Visualization) GetNoteAlignOk() (*string, bool)`

GetNoteAlignOk returns a tuple with the NoteAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteAlign

`func (o *Graph1Visualization) SetNoteAlign(v string)`

SetNoteAlign sets NoteAlign field to given value.

### HasNoteAlign

`func (o *Graph1Visualization) HasNoteAlign() bool`

HasNoteAlign returns a boolean if a field has been set.

### GetNoteJustifyContent

`func (o *Graph1Visualization) GetNoteJustifyContent() string`

GetNoteJustifyContent returns the NoteJustifyContent field if non-nil, zero value otherwise.

### GetNoteJustifyContentOk

`func (o *Graph1Visualization) GetNoteJustifyContentOk() (*string, bool)`

GetNoteJustifyContentOk returns a tuple with the NoteJustifyContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteJustifyContent

`func (o *Graph1Visualization) SetNoteJustifyContent(v string)`

SetNoteJustifyContent sets NoteJustifyContent field to given value.

### HasNoteJustifyContent

`func (o *Graph1Visualization) HasNoteJustifyContent() bool`

HasNoteJustifyContent returns a boolean if a field has been set.

### GetLayout

`func (o *Graph1Visualization) GetLayout() string`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *Graph1Visualization) GetLayoutOk() (*string, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *Graph1Visualization) SetLayout(v string)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *Graph1Visualization) HasLayout() bool`

HasLayout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


