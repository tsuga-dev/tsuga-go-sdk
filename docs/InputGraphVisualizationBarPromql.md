# InputGraphVisualizationBarPromql

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays PromQL metrics query-based aggregation as a bar chart | 
**Queries** | **[]string** | PromQL expressions configured for this widget. PromQL-backed widgets require at least one query. | 
**Aliases** | Pointer to [**InputGraphVisualizationTimeseriesPromqlAliases**](InputGraphVisualizationTimeseriesPromqlAliases.md) |  | [optional] 
**TimeBucket** | Pointer to [**GraphVisualizationTimeseriesPromqlTimeBucket**](GraphVisualizationTimeseriesPromqlTimeBucket.md) |  | [optional] 
**Precision** | Pointer to **float32** | Number of decimal places to display in the value | [optional] 
**Normalizer** | Pointer to [**Normalizer1**](Normalizer1.md) |  | [optional] 
**Thresholds** | Pointer to [**[]ThresholdMarker**](ThresholdMarker.md) | Threshold markers displayed on the chart | [optional] 
**LegendMode** | Pointer to **string** | Controls whether and how the widget displays legend or series details (e.g. table, legend-only, or no legend) | [optional] 
**YAxisSettings** | Pointer to [**InputGraphVisualizationTimeseriesConnectionYAxisSettings**](InputGraphVisualizationTimeseriesConnectionYAxisSettings.md) |  | [optional] 

## Methods

### NewInputGraphVisualizationBarPromql

`func NewInputGraphVisualizationBarPromql(type_ string, queries []string, ) *InputGraphVisualizationBarPromql`

NewInputGraphVisualizationBarPromql instantiates a new InputGraphVisualizationBarPromql object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputGraphVisualizationBarPromqlWithDefaults

`func NewInputGraphVisualizationBarPromqlWithDefaults() *InputGraphVisualizationBarPromql`

NewInputGraphVisualizationBarPromqlWithDefaults instantiates a new InputGraphVisualizationBarPromql object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputGraphVisualizationBarPromql) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputGraphVisualizationBarPromql) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputGraphVisualizationBarPromql) SetType(v string)`

SetType sets Type field to given value.


### GetQueries

`func (o *InputGraphVisualizationBarPromql) GetQueries() []string`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *InputGraphVisualizationBarPromql) GetQueriesOk() (*[]string, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *InputGraphVisualizationBarPromql) SetQueries(v []string)`

SetQueries sets Queries field to given value.


### GetAliases

`func (o *InputGraphVisualizationBarPromql) GetAliases() InputGraphVisualizationTimeseriesPromqlAliases`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *InputGraphVisualizationBarPromql) GetAliasesOk() (*InputGraphVisualizationTimeseriesPromqlAliases, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *InputGraphVisualizationBarPromql) SetAliases(v InputGraphVisualizationTimeseriesPromqlAliases)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *InputGraphVisualizationBarPromql) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetTimeBucket

`func (o *InputGraphVisualizationBarPromql) GetTimeBucket() GraphVisualizationTimeseriesPromqlTimeBucket`

GetTimeBucket returns the TimeBucket field if non-nil, zero value otherwise.

### GetTimeBucketOk

`func (o *InputGraphVisualizationBarPromql) GetTimeBucketOk() (*GraphVisualizationTimeseriesPromqlTimeBucket, bool)`

GetTimeBucketOk returns a tuple with the TimeBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBucket

`func (o *InputGraphVisualizationBarPromql) SetTimeBucket(v GraphVisualizationTimeseriesPromqlTimeBucket)`

SetTimeBucket sets TimeBucket field to given value.

### HasTimeBucket

`func (o *InputGraphVisualizationBarPromql) HasTimeBucket() bool`

HasTimeBucket returns a boolean if a field has been set.

### GetPrecision

`func (o *InputGraphVisualizationBarPromql) GetPrecision() float32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *InputGraphVisualizationBarPromql) GetPrecisionOk() (*float32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *InputGraphVisualizationBarPromql) SetPrecision(v float32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *InputGraphVisualizationBarPromql) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetNormalizer

`func (o *InputGraphVisualizationBarPromql) GetNormalizer() Normalizer1`

GetNormalizer returns the Normalizer field if non-nil, zero value otherwise.

### GetNormalizerOk

`func (o *InputGraphVisualizationBarPromql) GetNormalizerOk() (*Normalizer1, bool)`

GetNormalizerOk returns a tuple with the Normalizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizer

`func (o *InputGraphVisualizationBarPromql) SetNormalizer(v Normalizer1)`

SetNormalizer sets Normalizer field to given value.

### HasNormalizer

`func (o *InputGraphVisualizationBarPromql) HasNormalizer() bool`

HasNormalizer returns a boolean if a field has been set.

### GetThresholds

`func (o *InputGraphVisualizationBarPromql) GetThresholds() []ThresholdMarker`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *InputGraphVisualizationBarPromql) GetThresholdsOk() (*[]ThresholdMarker, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *InputGraphVisualizationBarPromql) SetThresholds(v []ThresholdMarker)`

SetThresholds sets Thresholds field to given value.

### HasThresholds

`func (o *InputGraphVisualizationBarPromql) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.

### GetLegendMode

`func (o *InputGraphVisualizationBarPromql) GetLegendMode() string`

GetLegendMode returns the LegendMode field if non-nil, zero value otherwise.

### GetLegendModeOk

`func (o *InputGraphVisualizationBarPromql) GetLegendModeOk() (*string, bool)`

GetLegendModeOk returns a tuple with the LegendMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegendMode

`func (o *InputGraphVisualizationBarPromql) SetLegendMode(v string)`

SetLegendMode sets LegendMode field to given value.

### HasLegendMode

`func (o *InputGraphVisualizationBarPromql) HasLegendMode() bool`

HasLegendMode returns a boolean if a field has been set.

### GetYAxisSettings

`func (o *InputGraphVisualizationBarPromql) GetYAxisSettings() InputGraphVisualizationTimeseriesConnectionYAxisSettings`

GetYAxisSettings returns the YAxisSettings field if non-nil, zero value otherwise.

### GetYAxisSettingsOk

`func (o *InputGraphVisualizationBarPromql) GetYAxisSettingsOk() (*InputGraphVisualizationTimeseriesConnectionYAxisSettings, bool)`

GetYAxisSettingsOk returns a tuple with the YAxisSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYAxisSettings

`func (o *InputGraphVisualizationBarPromql) SetYAxisSettings(v InputGraphVisualizationTimeseriesConnectionYAxisSettings)`

SetYAxisSettings sets YAxisSettings field to given value.

### HasYAxisSettings

`func (o *InputGraphVisualizationBarPromql) HasYAxisSettings() bool`

HasYAxisSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


