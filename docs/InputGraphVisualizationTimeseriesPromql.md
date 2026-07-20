# InputGraphVisualizationTimeseriesPromql

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays PromQL metrics query-based aggregation as a time series chart | 
**Queries** | **[]string** | PromQL expressions configured for this widget. PromQL-backed widgets require at least one query. | 
**Aliases** | Pointer to [**InputGraphVisualizationTimeseriesPromqlAliases**](InputGraphVisualizationTimeseriesPromqlAliases.md) |  | [optional] 
**TimeBucket** | Pointer to [**GraphVisualizationTimeseriesPromqlTimeBucket**](GraphVisualizationTimeseriesPromqlTimeBucket.md) |  | [optional] 
**Normalizer** | Pointer to [**Normalizer1**](Normalizer1.md) |  | [optional] 
**Precision** | Pointer to **float32** | Number of decimal places to display in the value | [optional] 
**LegendMode** | Pointer to **string** | Controls whether and how the widget displays legend or series details (e.g. table, legend-only, or no legend) | [optional] 
**Thresholds** | Pointer to [**[]ThresholdMarker**](ThresholdMarker.md) | Threshold markers displayed on the chart | [optional] 
**YAxisSettings** | Pointer to [**InputGraphVisualizationTimeseriesConnectionYAxisSettings**](InputGraphVisualizationTimeseriesConnectionYAxisSettings.md) |  | [optional] 
**Smoothing** | Pointer to **bool** | Whether to apply automatic smoothing to the rendered timeseries | [optional] 

## Methods

### NewInputGraphVisualizationTimeseriesPromql

`func NewInputGraphVisualizationTimeseriesPromql(type_ string, queries []string, ) *InputGraphVisualizationTimeseriesPromql`

NewInputGraphVisualizationTimeseriesPromql instantiates a new InputGraphVisualizationTimeseriesPromql object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputGraphVisualizationTimeseriesPromqlWithDefaults

`func NewInputGraphVisualizationTimeseriesPromqlWithDefaults() *InputGraphVisualizationTimeseriesPromql`

NewInputGraphVisualizationTimeseriesPromqlWithDefaults instantiates a new InputGraphVisualizationTimeseriesPromql object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputGraphVisualizationTimeseriesPromql) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputGraphVisualizationTimeseriesPromql) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputGraphVisualizationTimeseriesPromql) SetType(v string)`

SetType sets Type field to given value.


### GetQueries

`func (o *InputGraphVisualizationTimeseriesPromql) GetQueries() []string`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *InputGraphVisualizationTimeseriesPromql) GetQueriesOk() (*[]string, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *InputGraphVisualizationTimeseriesPromql) SetQueries(v []string)`

SetQueries sets Queries field to given value.


### GetAliases

`func (o *InputGraphVisualizationTimeseriesPromql) GetAliases() InputGraphVisualizationTimeseriesPromqlAliases`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *InputGraphVisualizationTimeseriesPromql) GetAliasesOk() (*InputGraphVisualizationTimeseriesPromqlAliases, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *InputGraphVisualizationTimeseriesPromql) SetAliases(v InputGraphVisualizationTimeseriesPromqlAliases)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *InputGraphVisualizationTimeseriesPromql) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetTimeBucket

`func (o *InputGraphVisualizationTimeseriesPromql) GetTimeBucket() GraphVisualizationTimeseriesPromqlTimeBucket`

GetTimeBucket returns the TimeBucket field if non-nil, zero value otherwise.

### GetTimeBucketOk

`func (o *InputGraphVisualizationTimeseriesPromql) GetTimeBucketOk() (*GraphVisualizationTimeseriesPromqlTimeBucket, bool)`

GetTimeBucketOk returns a tuple with the TimeBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBucket

`func (o *InputGraphVisualizationTimeseriesPromql) SetTimeBucket(v GraphVisualizationTimeseriesPromqlTimeBucket)`

SetTimeBucket sets TimeBucket field to given value.

### HasTimeBucket

`func (o *InputGraphVisualizationTimeseriesPromql) HasTimeBucket() bool`

HasTimeBucket returns a boolean if a field has been set.

### GetNormalizer

`func (o *InputGraphVisualizationTimeseriesPromql) GetNormalizer() Normalizer1`

GetNormalizer returns the Normalizer field if non-nil, zero value otherwise.

### GetNormalizerOk

`func (o *InputGraphVisualizationTimeseriesPromql) GetNormalizerOk() (*Normalizer1, bool)`

GetNormalizerOk returns a tuple with the Normalizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizer

`func (o *InputGraphVisualizationTimeseriesPromql) SetNormalizer(v Normalizer1)`

SetNormalizer sets Normalizer field to given value.

### HasNormalizer

`func (o *InputGraphVisualizationTimeseriesPromql) HasNormalizer() bool`

HasNormalizer returns a boolean if a field has been set.

### GetPrecision

`func (o *InputGraphVisualizationTimeseriesPromql) GetPrecision() float32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *InputGraphVisualizationTimeseriesPromql) GetPrecisionOk() (*float32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *InputGraphVisualizationTimeseriesPromql) SetPrecision(v float32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *InputGraphVisualizationTimeseriesPromql) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetLegendMode

`func (o *InputGraphVisualizationTimeseriesPromql) GetLegendMode() string`

GetLegendMode returns the LegendMode field if non-nil, zero value otherwise.

### GetLegendModeOk

`func (o *InputGraphVisualizationTimeseriesPromql) GetLegendModeOk() (*string, bool)`

GetLegendModeOk returns a tuple with the LegendMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegendMode

`func (o *InputGraphVisualizationTimeseriesPromql) SetLegendMode(v string)`

SetLegendMode sets LegendMode field to given value.

### HasLegendMode

`func (o *InputGraphVisualizationTimeseriesPromql) HasLegendMode() bool`

HasLegendMode returns a boolean if a field has been set.

### GetThresholds

`func (o *InputGraphVisualizationTimeseriesPromql) GetThresholds() []ThresholdMarker`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *InputGraphVisualizationTimeseriesPromql) GetThresholdsOk() (*[]ThresholdMarker, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *InputGraphVisualizationTimeseriesPromql) SetThresholds(v []ThresholdMarker)`

SetThresholds sets Thresholds field to given value.

### HasThresholds

`func (o *InputGraphVisualizationTimeseriesPromql) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.

### GetYAxisSettings

`func (o *InputGraphVisualizationTimeseriesPromql) GetYAxisSettings() InputGraphVisualizationTimeseriesConnectionYAxisSettings`

GetYAxisSettings returns the YAxisSettings field if non-nil, zero value otherwise.

### GetYAxisSettingsOk

`func (o *InputGraphVisualizationTimeseriesPromql) GetYAxisSettingsOk() (*InputGraphVisualizationTimeseriesConnectionYAxisSettings, bool)`

GetYAxisSettingsOk returns a tuple with the YAxisSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYAxisSettings

`func (o *InputGraphVisualizationTimeseriesPromql) SetYAxisSettings(v InputGraphVisualizationTimeseriesConnectionYAxisSettings)`

SetYAxisSettings sets YAxisSettings field to given value.

### HasYAxisSettings

`func (o *InputGraphVisualizationTimeseriesPromql) HasYAxisSettings() bool`

HasYAxisSettings returns a boolean if a field has been set.

### GetSmoothing

`func (o *InputGraphVisualizationTimeseriesPromql) GetSmoothing() bool`

GetSmoothing returns the Smoothing field if non-nil, zero value otherwise.

### GetSmoothingOk

`func (o *InputGraphVisualizationTimeseriesPromql) GetSmoothingOk() (*bool, bool)`

GetSmoothingOk returns a tuple with the Smoothing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmoothing

`func (o *InputGraphVisualizationTimeseriesPromql) SetSmoothing(v bool)`

SetSmoothing sets Smoothing field to given value.

### HasSmoothing

`func (o *InputGraphVisualizationTimeseriesPromql) HasSmoothing() bool`

HasSmoothing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


