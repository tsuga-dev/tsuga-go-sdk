# InputGraphVisualizationTimeseries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays the aggregation as a time series chart | 
**Source** | **string** | Data source being queried for this aggregation | 
**Queries** | [**[]AggregationQuery1**](AggregationQuery1.md) | Aggregations that may be combined together in the same query | 
**Formula** | Pointer to **string** | Formula referencing query outputs (e.g. q1+q2) to compute derived series | [optional] 
**Aliases** | Pointer to [**InputGraphVisualizationTimeseriesAliases**](InputGraphVisualizationTimeseriesAliases.md) |  | [optional] 
**VisibleSeries** | Pointer to **[]bool** | Flags indicating whether each query or formula series is visible | [optional] 
**GroupBy** | Pointer to [**[]AggregationGroupBy1**](AggregationGroupBy1.md) | Fields used to group the results | [optional] 
**TimeBucket** | Pointer to [**GraphVisualizationTimeseriesTimeBucket**](GraphVisualizationTimeseriesTimeBucket.md) |  | [optional] 
**Normalizer** | Pointer to [**Normalizer1**](Normalizer1.md) |  | [optional] 
**Precision** | Pointer to **float32** | Number of decimal places to display in the value | [optional] 
**LegendMode** | Pointer to **string** | Controls whether and how the widget displays legend or series details (e.g. table, legend-only, or no legend) | [optional] 
**Thresholds** | Pointer to [**[]ThresholdMarker**](ThresholdMarker.md) | Threshold markers displayed on the chart | [optional] 
**YAxisSettings** | Pointer to [**InputGraphVisualizationTimeseriesConnectionYAxisSettings**](InputGraphVisualizationTimeseriesConnectionYAxisSettings.md) |  | [optional] 
**Smoothing** | Pointer to **bool** | Whether to apply automatic smoothing to the rendered timeseries | [optional] 

## Methods

### NewInputGraphVisualizationTimeseries

`func NewInputGraphVisualizationTimeseries(type_ string, source string, queries []AggregationQuery1, ) *InputGraphVisualizationTimeseries`

NewInputGraphVisualizationTimeseries instantiates a new InputGraphVisualizationTimeseries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputGraphVisualizationTimeseriesWithDefaults

`func NewInputGraphVisualizationTimeseriesWithDefaults() *InputGraphVisualizationTimeseries`

NewInputGraphVisualizationTimeseriesWithDefaults instantiates a new InputGraphVisualizationTimeseries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputGraphVisualizationTimeseries) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputGraphVisualizationTimeseries) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputGraphVisualizationTimeseries) SetType(v string)`

SetType sets Type field to given value.


### GetSource

`func (o *InputGraphVisualizationTimeseries) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InputGraphVisualizationTimeseries) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InputGraphVisualizationTimeseries) SetSource(v string)`

SetSource sets Source field to given value.


### GetQueries

`func (o *InputGraphVisualizationTimeseries) GetQueries() []AggregationQuery1`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *InputGraphVisualizationTimeseries) GetQueriesOk() (*[]AggregationQuery1, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *InputGraphVisualizationTimeseries) SetQueries(v []AggregationQuery1)`

SetQueries sets Queries field to given value.


### GetFormula

`func (o *InputGraphVisualizationTimeseries) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *InputGraphVisualizationTimeseries) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *InputGraphVisualizationTimeseries) SetFormula(v string)`

SetFormula sets Formula field to given value.

### HasFormula

`func (o *InputGraphVisualizationTimeseries) HasFormula() bool`

HasFormula returns a boolean if a field has been set.

### GetAliases

`func (o *InputGraphVisualizationTimeseries) GetAliases() InputGraphVisualizationTimeseriesAliases`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *InputGraphVisualizationTimeseries) GetAliasesOk() (*InputGraphVisualizationTimeseriesAliases, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *InputGraphVisualizationTimeseries) SetAliases(v InputGraphVisualizationTimeseriesAliases)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *InputGraphVisualizationTimeseries) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetVisibleSeries

`func (o *InputGraphVisualizationTimeseries) GetVisibleSeries() []bool`

GetVisibleSeries returns the VisibleSeries field if non-nil, zero value otherwise.

### GetVisibleSeriesOk

`func (o *InputGraphVisualizationTimeseries) GetVisibleSeriesOk() (*[]bool, bool)`

GetVisibleSeriesOk returns a tuple with the VisibleSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleSeries

`func (o *InputGraphVisualizationTimeseries) SetVisibleSeries(v []bool)`

SetVisibleSeries sets VisibleSeries field to given value.

### HasVisibleSeries

`func (o *InputGraphVisualizationTimeseries) HasVisibleSeries() bool`

HasVisibleSeries returns a boolean if a field has been set.

### GetGroupBy

`func (o *InputGraphVisualizationTimeseries) GetGroupBy() []AggregationGroupBy1`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *InputGraphVisualizationTimeseries) GetGroupByOk() (*[]AggregationGroupBy1, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *InputGraphVisualizationTimeseries) SetGroupBy(v []AggregationGroupBy1)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *InputGraphVisualizationTimeseries) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetTimeBucket

`func (o *InputGraphVisualizationTimeseries) GetTimeBucket() GraphVisualizationTimeseriesTimeBucket`

GetTimeBucket returns the TimeBucket field if non-nil, zero value otherwise.

### GetTimeBucketOk

`func (o *InputGraphVisualizationTimeseries) GetTimeBucketOk() (*GraphVisualizationTimeseriesTimeBucket, bool)`

GetTimeBucketOk returns a tuple with the TimeBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBucket

`func (o *InputGraphVisualizationTimeseries) SetTimeBucket(v GraphVisualizationTimeseriesTimeBucket)`

SetTimeBucket sets TimeBucket field to given value.

### HasTimeBucket

`func (o *InputGraphVisualizationTimeseries) HasTimeBucket() bool`

HasTimeBucket returns a boolean if a field has been set.

### GetNormalizer

`func (o *InputGraphVisualizationTimeseries) GetNormalizer() Normalizer1`

GetNormalizer returns the Normalizer field if non-nil, zero value otherwise.

### GetNormalizerOk

`func (o *InputGraphVisualizationTimeseries) GetNormalizerOk() (*Normalizer1, bool)`

GetNormalizerOk returns a tuple with the Normalizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizer

`func (o *InputGraphVisualizationTimeseries) SetNormalizer(v Normalizer1)`

SetNormalizer sets Normalizer field to given value.

### HasNormalizer

`func (o *InputGraphVisualizationTimeseries) HasNormalizer() bool`

HasNormalizer returns a boolean if a field has been set.

### GetPrecision

`func (o *InputGraphVisualizationTimeseries) GetPrecision() float32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *InputGraphVisualizationTimeseries) GetPrecisionOk() (*float32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *InputGraphVisualizationTimeseries) SetPrecision(v float32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *InputGraphVisualizationTimeseries) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetLegendMode

`func (o *InputGraphVisualizationTimeseries) GetLegendMode() string`

GetLegendMode returns the LegendMode field if non-nil, zero value otherwise.

### GetLegendModeOk

`func (o *InputGraphVisualizationTimeseries) GetLegendModeOk() (*string, bool)`

GetLegendModeOk returns a tuple with the LegendMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegendMode

`func (o *InputGraphVisualizationTimeseries) SetLegendMode(v string)`

SetLegendMode sets LegendMode field to given value.

### HasLegendMode

`func (o *InputGraphVisualizationTimeseries) HasLegendMode() bool`

HasLegendMode returns a boolean if a field has been set.

### GetThresholds

`func (o *InputGraphVisualizationTimeseries) GetThresholds() []ThresholdMarker`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *InputGraphVisualizationTimeseries) GetThresholdsOk() (*[]ThresholdMarker, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *InputGraphVisualizationTimeseries) SetThresholds(v []ThresholdMarker)`

SetThresholds sets Thresholds field to given value.

### HasThresholds

`func (o *InputGraphVisualizationTimeseries) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.

### GetYAxisSettings

`func (o *InputGraphVisualizationTimeseries) GetYAxisSettings() InputGraphVisualizationTimeseriesConnectionYAxisSettings`

GetYAxisSettings returns the YAxisSettings field if non-nil, zero value otherwise.

### GetYAxisSettingsOk

`func (o *InputGraphVisualizationTimeseries) GetYAxisSettingsOk() (*InputGraphVisualizationTimeseriesConnectionYAxisSettings, bool)`

GetYAxisSettingsOk returns a tuple with the YAxisSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYAxisSettings

`func (o *InputGraphVisualizationTimeseries) SetYAxisSettings(v InputGraphVisualizationTimeseriesConnectionYAxisSettings)`

SetYAxisSettings sets YAxisSettings field to given value.

### HasYAxisSettings

`func (o *InputGraphVisualizationTimeseries) HasYAxisSettings() bool`

HasYAxisSettings returns a boolean if a field has been set.

### GetSmoothing

`func (o *InputGraphVisualizationTimeseries) GetSmoothing() bool`

GetSmoothing returns the Smoothing field if non-nil, zero value otherwise.

### GetSmoothingOk

`func (o *InputGraphVisualizationTimeseries) GetSmoothingOk() (*bool, bool)`

GetSmoothingOk returns a tuple with the Smoothing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmoothing

`func (o *InputGraphVisualizationTimeseries) SetSmoothing(v bool)`

SetSmoothing sets Smoothing field to given value.

### HasSmoothing

`func (o *InputGraphVisualizationTimeseries) HasSmoothing() bool`

HasSmoothing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


