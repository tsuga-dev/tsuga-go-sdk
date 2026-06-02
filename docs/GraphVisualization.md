# GraphVisualization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays the database rows-based aggregation as a time series chart | 
**ConnectionId** | **string** | The ID of the connection to use to query the datastore. | 
**Queries** | [**[]AggregationQuery1**](AggregationQuery1.md) | Aggregations that may be combined together in the same query | 
**LegendMode** | Pointer to **string** | Controls whether and how the widget displays legend or series details (e.g. table, legend-only, or no legend) | [optional] 
**Thresholds** | Pointer to [**[]ThresholdMarker**](ThresholdMarker.md) | Threshold markers displayed on the chart | [optional] 
**YAxisSettings** | Pointer to [**InputGraphVisualizationTimeseriesConnectionYAxisSettings**](InputGraphVisualizationTimeseriesConnectionYAxisSettings.md) |  | [optional] 
**Query** | **string** | Tsuga query that selects logs to cluster into patterns | 
**ListColumns** | Pointer to [**[]WidgetListColumn1**](WidgetListColumn1.md) | Custom columns to display for each log | [optional] 
**ListColumnsSize** | Pointer to **map[string]float32** |  | [optional] 
**BackgroundMode** | Pointer to **string** | Controls whether the widget uses a solid or transparent background | [optional] 
**Conditions** | Pointer to [**[]ConditionalFormatting**](ConditionalFormatting.md) | Conditional formatting rules applied to the displayed value | [optional] 
**Precision** | Pointer to **float32** | Number of decimal places to display in the value | [optional] 
**Normalizer** | Pointer to [**Normalizer1**](Normalizer1.md) |  | [optional] 
**Source** | **string** | Data source being queried for this aggregation | 
**Formula** | Pointer to **string** | Formula referencing query outputs (e.g. q1+q2) to compute derived series | [optional] 
**Aliases** | Pointer to [**InputGraphVisualizationTimeseriesAliases**](InputGraphVisualizationTimeseriesAliases.md) |  | [optional] 
**VisibleSeries** | Pointer to **[]bool** | Flags indicating whether each query or formula series is visible | [optional] 
**GroupBy** | Pointer to [**[]AggregationGroupBy1**](AggregationGroupBy1.md) | Fields used to group the results | [optional] 
**TimeBucket** | Pointer to [**GraphVisualizationTimeseriesTimeBucket**](GraphVisualizationTimeseriesTimeBucket.md) |  | [optional] 
**Smoothing** | Pointer to **bool** | Whether to apply automatic smoothing to the rendered timeseries | [optional] 
**PercentileMarkers** | Pointer to **[]int32** | Percentile markers displayed on top of the distribution chart | [optional] 
**Palette** | Pointer to **string** | Color palette used to render the heatmap intensity gradient | [optional] 
**Columns** | [**[]TableColumn1**](TableColumn1.md) | Each column defines an independent aggregation displayed as a table column | 
**DefaultSorting** | Pointer to [**[]TableDefaultSorting1**](TableDefaultSorting1.md) |  | [optional] 
**Note** | Pointer to **string** | Markdown-compatible text shown in the note | [optional] 
**NoteColor** | Pointer to **string** | Background color used to render the note | [optional] 
**NoteAlign** | Pointer to **string** | Flex alignment keyword used for widget layout | [optional] 
**NoteJustifyContent** | Pointer to **string** | Flex alignment keyword used for widget layout | [optional] 
**Layout** | Pointer to **string** | Layout used to render log patterns | [optional] 

## Methods

### NewGraphVisualization

`func NewGraphVisualization(type_ string, connectionId string, queries []AggregationQuery1, query string, source string, columns []TableColumn1, ) *GraphVisualization`

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

`func (o *GraphVisualization) GetPrecision() float32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *GraphVisualization) GetPrecisionOk() (*float32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *GraphVisualization) SetPrecision(v float32)`

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

### GetAliases

`func (o *GraphVisualization) GetAliases() InputGraphVisualizationTimeseriesAliases`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *GraphVisualization) GetAliasesOk() (*InputGraphVisualizationTimeseriesAliases, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *GraphVisualization) SetAliases(v InputGraphVisualizationTimeseriesAliases)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *GraphVisualization) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

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

`func (o *GraphVisualization) GetGroupBy() []AggregationGroupBy1`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *GraphVisualization) GetGroupByOk() (*[]AggregationGroupBy1, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *GraphVisualization) SetGroupBy(v []AggregationGroupBy1)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *GraphVisualization) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetTimeBucket

`func (o *GraphVisualization) GetTimeBucket() GraphVisualizationTimeseriesTimeBucket`

GetTimeBucket returns the TimeBucket field if non-nil, zero value otherwise.

### GetTimeBucketOk

`func (o *GraphVisualization) GetTimeBucketOk() (*GraphVisualizationTimeseriesTimeBucket, bool)`

GetTimeBucketOk returns a tuple with the TimeBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBucket

`func (o *GraphVisualization) SetTimeBucket(v GraphVisualizationTimeseriesTimeBucket)`

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


### GetDefaultSorting

`func (o *GraphVisualization) GetDefaultSorting() []TableDefaultSorting1`

GetDefaultSorting returns the DefaultSorting field if non-nil, zero value otherwise.

### GetDefaultSortingOk

`func (o *GraphVisualization) GetDefaultSortingOk() (*[]TableDefaultSorting1, bool)`

GetDefaultSortingOk returns a tuple with the DefaultSorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSorting

`func (o *GraphVisualization) SetDefaultSorting(v []TableDefaultSorting1)`

SetDefaultSorting sets DefaultSorting field to given value.

### HasDefaultSorting

`func (o *GraphVisualization) HasDefaultSorting() bool`

HasDefaultSorting returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


