# GraphVisualizationTimeseriesPromql

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays PromQL metrics query-based aggregation as a time series chart | 
**Queries** | **[]string** | PromQL expressions configured for this widget. PromQL-backed widgets require at least one query. | 
**Aliases** | Pointer to [**GraphVisualizationTimeseriesPromqlAliases**](GraphVisualizationTimeseriesPromqlAliases.md) |  | [optional] 
**TimeBucket** | Pointer to [**GraphVisualizationTimeseriesPromqlTimeBucket**](GraphVisualizationTimeseriesPromqlTimeBucket.md) |  | [optional] 
**Normalizer** | Pointer to [**Normalizer**](Normalizer.md) |  | [optional] 
**Precision** | Pointer to **float32** | Number of decimal places to display in the value | [optional] 
**LegendMode** | Pointer to **string** | Controls whether and how the widget displays legend or series details (e.g. table, legend-only, or no legend) | [optional] 
**Thresholds** | Pointer to [**[]ThresholdMarker**](ThresholdMarker.md) | Threshold markers displayed on the chart | [optional] 
**YAxisSettings** | Pointer to [**GraphVisualizationTimeseriesConnectionYAxisSettings**](GraphVisualizationTimeseriesConnectionYAxisSettings.md) |  | [optional] 
**Smoothing** | Pointer to **bool** | Whether to apply automatic smoothing to the rendered timeseries | [optional] 

## Methods

### NewGraphVisualizationTimeseriesPromql

`func NewGraphVisualizationTimeseriesPromql(type_ string, queries []string, ) *GraphVisualizationTimeseriesPromql`

NewGraphVisualizationTimeseriesPromql instantiates a new GraphVisualizationTimeseriesPromql object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphVisualizationTimeseriesPromqlWithDefaults

`func NewGraphVisualizationTimeseriesPromqlWithDefaults() *GraphVisualizationTimeseriesPromql`

NewGraphVisualizationTimeseriesPromqlWithDefaults instantiates a new GraphVisualizationTimeseriesPromql object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GraphVisualizationTimeseriesPromql) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GraphVisualizationTimeseriesPromql) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GraphVisualizationTimeseriesPromql) SetType(v string)`

SetType sets Type field to given value.


### GetQueries

`func (o *GraphVisualizationTimeseriesPromql) GetQueries() []string`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *GraphVisualizationTimeseriesPromql) GetQueriesOk() (*[]string, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *GraphVisualizationTimeseriesPromql) SetQueries(v []string)`

SetQueries sets Queries field to given value.


### GetAliases

`func (o *GraphVisualizationTimeseriesPromql) GetAliases() GraphVisualizationTimeseriesPromqlAliases`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *GraphVisualizationTimeseriesPromql) GetAliasesOk() (*GraphVisualizationTimeseriesPromqlAliases, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *GraphVisualizationTimeseriesPromql) SetAliases(v GraphVisualizationTimeseriesPromqlAliases)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *GraphVisualizationTimeseriesPromql) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetTimeBucket

`func (o *GraphVisualizationTimeseriesPromql) GetTimeBucket() GraphVisualizationTimeseriesPromqlTimeBucket`

GetTimeBucket returns the TimeBucket field if non-nil, zero value otherwise.

### GetTimeBucketOk

`func (o *GraphVisualizationTimeseriesPromql) GetTimeBucketOk() (*GraphVisualizationTimeseriesPromqlTimeBucket, bool)`

GetTimeBucketOk returns a tuple with the TimeBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBucket

`func (o *GraphVisualizationTimeseriesPromql) SetTimeBucket(v GraphVisualizationTimeseriesPromqlTimeBucket)`

SetTimeBucket sets TimeBucket field to given value.

### HasTimeBucket

`func (o *GraphVisualizationTimeseriesPromql) HasTimeBucket() bool`

HasTimeBucket returns a boolean if a field has been set.

### GetNormalizer

`func (o *GraphVisualizationTimeseriesPromql) GetNormalizer() Normalizer`

GetNormalizer returns the Normalizer field if non-nil, zero value otherwise.

### GetNormalizerOk

`func (o *GraphVisualizationTimeseriesPromql) GetNormalizerOk() (*Normalizer, bool)`

GetNormalizerOk returns a tuple with the Normalizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizer

`func (o *GraphVisualizationTimeseriesPromql) SetNormalizer(v Normalizer)`

SetNormalizer sets Normalizer field to given value.

### HasNormalizer

`func (o *GraphVisualizationTimeseriesPromql) HasNormalizer() bool`

HasNormalizer returns a boolean if a field has been set.

### GetPrecision

`func (o *GraphVisualizationTimeseriesPromql) GetPrecision() float32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *GraphVisualizationTimeseriesPromql) GetPrecisionOk() (*float32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *GraphVisualizationTimeseriesPromql) SetPrecision(v float32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *GraphVisualizationTimeseriesPromql) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetLegendMode

`func (o *GraphVisualizationTimeseriesPromql) GetLegendMode() string`

GetLegendMode returns the LegendMode field if non-nil, zero value otherwise.

### GetLegendModeOk

`func (o *GraphVisualizationTimeseriesPromql) GetLegendModeOk() (*string, bool)`

GetLegendModeOk returns a tuple with the LegendMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegendMode

`func (o *GraphVisualizationTimeseriesPromql) SetLegendMode(v string)`

SetLegendMode sets LegendMode field to given value.

### HasLegendMode

`func (o *GraphVisualizationTimeseriesPromql) HasLegendMode() bool`

HasLegendMode returns a boolean if a field has been set.

### GetThresholds

`func (o *GraphVisualizationTimeseriesPromql) GetThresholds() []ThresholdMarker`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *GraphVisualizationTimeseriesPromql) GetThresholdsOk() (*[]ThresholdMarker, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *GraphVisualizationTimeseriesPromql) SetThresholds(v []ThresholdMarker)`

SetThresholds sets Thresholds field to given value.

### HasThresholds

`func (o *GraphVisualizationTimeseriesPromql) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.

### GetYAxisSettings

`func (o *GraphVisualizationTimeseriesPromql) GetYAxisSettings() GraphVisualizationTimeseriesConnectionYAxisSettings`

GetYAxisSettings returns the YAxisSettings field if non-nil, zero value otherwise.

### GetYAxisSettingsOk

`func (o *GraphVisualizationTimeseriesPromql) GetYAxisSettingsOk() (*GraphVisualizationTimeseriesConnectionYAxisSettings, bool)`

GetYAxisSettingsOk returns a tuple with the YAxisSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYAxisSettings

`func (o *GraphVisualizationTimeseriesPromql) SetYAxisSettings(v GraphVisualizationTimeseriesConnectionYAxisSettings)`

SetYAxisSettings sets YAxisSettings field to given value.

### HasYAxisSettings

`func (o *GraphVisualizationTimeseriesPromql) HasYAxisSettings() bool`

HasYAxisSettings returns a boolean if a field has been set.

### GetSmoothing

`func (o *GraphVisualizationTimeseriesPromql) GetSmoothing() bool`

GetSmoothing returns the Smoothing field if non-nil, zero value otherwise.

### GetSmoothingOk

`func (o *GraphVisualizationTimeseriesPromql) GetSmoothingOk() (*bool, bool)`

GetSmoothingOk returns a tuple with the Smoothing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmoothing

`func (o *GraphVisualizationTimeseriesPromql) SetSmoothing(v bool)`

SetSmoothing sets Smoothing field to given value.

### HasSmoothing

`func (o *GraphVisualizationTimeseriesPromql) HasSmoothing() bool`

HasSmoothing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


