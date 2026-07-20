# GraphVisualizationBarPromql

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays PromQL metrics query-based aggregation as a bar chart | 
**Queries** | **[]string** | PromQL expressions configured for this widget. PromQL-backed widgets require at least one query. | 
**Aliases** | Pointer to [**GraphVisualizationTimeseriesPromqlAliases**](GraphVisualizationTimeseriesPromqlAliases.md) |  | [optional] 
**TimeBucket** | Pointer to [**GraphVisualizationTimeseriesPromqlTimeBucket**](GraphVisualizationTimeseriesPromqlTimeBucket.md) |  | [optional] 
**Precision** | Pointer to **float32** | Number of decimal places to display in the value | [optional] 
**Normalizer** | Pointer to [**Normalizer**](Normalizer.md) |  | [optional] 
**Thresholds** | Pointer to [**[]ThresholdMarker**](ThresholdMarker.md) | Threshold markers displayed on the chart | [optional] 
**LegendMode** | Pointer to **string** | Controls whether and how the widget displays legend or series details (e.g. table, legend-only, or no legend) | [optional] 
**YAxisSettings** | Pointer to [**GraphVisualizationTimeseriesConnectionYAxisSettings**](GraphVisualizationTimeseriesConnectionYAxisSettings.md) |  | [optional] 

## Methods

### NewGraphVisualizationBarPromql

`func NewGraphVisualizationBarPromql(type_ string, queries []string, ) *GraphVisualizationBarPromql`

NewGraphVisualizationBarPromql instantiates a new GraphVisualizationBarPromql object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphVisualizationBarPromqlWithDefaults

`func NewGraphVisualizationBarPromqlWithDefaults() *GraphVisualizationBarPromql`

NewGraphVisualizationBarPromqlWithDefaults instantiates a new GraphVisualizationBarPromql object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GraphVisualizationBarPromql) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GraphVisualizationBarPromql) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GraphVisualizationBarPromql) SetType(v string)`

SetType sets Type field to given value.


### GetQueries

`func (o *GraphVisualizationBarPromql) GetQueries() []string`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *GraphVisualizationBarPromql) GetQueriesOk() (*[]string, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *GraphVisualizationBarPromql) SetQueries(v []string)`

SetQueries sets Queries field to given value.


### GetAliases

`func (o *GraphVisualizationBarPromql) GetAliases() GraphVisualizationTimeseriesPromqlAliases`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *GraphVisualizationBarPromql) GetAliasesOk() (*GraphVisualizationTimeseriesPromqlAliases, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *GraphVisualizationBarPromql) SetAliases(v GraphVisualizationTimeseriesPromqlAliases)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *GraphVisualizationBarPromql) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetTimeBucket

`func (o *GraphVisualizationBarPromql) GetTimeBucket() GraphVisualizationTimeseriesPromqlTimeBucket`

GetTimeBucket returns the TimeBucket field if non-nil, zero value otherwise.

### GetTimeBucketOk

`func (o *GraphVisualizationBarPromql) GetTimeBucketOk() (*GraphVisualizationTimeseriesPromqlTimeBucket, bool)`

GetTimeBucketOk returns a tuple with the TimeBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBucket

`func (o *GraphVisualizationBarPromql) SetTimeBucket(v GraphVisualizationTimeseriesPromqlTimeBucket)`

SetTimeBucket sets TimeBucket field to given value.

### HasTimeBucket

`func (o *GraphVisualizationBarPromql) HasTimeBucket() bool`

HasTimeBucket returns a boolean if a field has been set.

### GetPrecision

`func (o *GraphVisualizationBarPromql) GetPrecision() float32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *GraphVisualizationBarPromql) GetPrecisionOk() (*float32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *GraphVisualizationBarPromql) SetPrecision(v float32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *GraphVisualizationBarPromql) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetNormalizer

`func (o *GraphVisualizationBarPromql) GetNormalizer() Normalizer`

GetNormalizer returns the Normalizer field if non-nil, zero value otherwise.

### GetNormalizerOk

`func (o *GraphVisualizationBarPromql) GetNormalizerOk() (*Normalizer, bool)`

GetNormalizerOk returns a tuple with the Normalizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizer

`func (o *GraphVisualizationBarPromql) SetNormalizer(v Normalizer)`

SetNormalizer sets Normalizer field to given value.

### HasNormalizer

`func (o *GraphVisualizationBarPromql) HasNormalizer() bool`

HasNormalizer returns a boolean if a field has been set.

### GetThresholds

`func (o *GraphVisualizationBarPromql) GetThresholds() []ThresholdMarker`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *GraphVisualizationBarPromql) GetThresholdsOk() (*[]ThresholdMarker, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *GraphVisualizationBarPromql) SetThresholds(v []ThresholdMarker)`

SetThresholds sets Thresholds field to given value.

### HasThresholds

`func (o *GraphVisualizationBarPromql) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.

### GetLegendMode

`func (o *GraphVisualizationBarPromql) GetLegendMode() string`

GetLegendMode returns the LegendMode field if non-nil, zero value otherwise.

### GetLegendModeOk

`func (o *GraphVisualizationBarPromql) GetLegendModeOk() (*string, bool)`

GetLegendModeOk returns a tuple with the LegendMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegendMode

`func (o *GraphVisualizationBarPromql) SetLegendMode(v string)`

SetLegendMode sets LegendMode field to given value.

### HasLegendMode

`func (o *GraphVisualizationBarPromql) HasLegendMode() bool`

HasLegendMode returns a boolean if a field has been set.

### GetYAxisSettings

`func (o *GraphVisualizationBarPromql) GetYAxisSettings() GraphVisualizationTimeseriesConnectionYAxisSettings`

GetYAxisSettings returns the YAxisSettings field if non-nil, zero value otherwise.

### GetYAxisSettingsOk

`func (o *GraphVisualizationBarPromql) GetYAxisSettingsOk() (*GraphVisualizationTimeseriesConnectionYAxisSettings, bool)`

GetYAxisSettingsOk returns a tuple with the YAxisSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYAxisSettings

`func (o *GraphVisualizationBarPromql) SetYAxisSettings(v GraphVisualizationTimeseriesConnectionYAxisSettings)`

SetYAxisSettings sets YAxisSettings field to given value.

### HasYAxisSettings

`func (o *GraphVisualizationBarPromql) HasYAxisSettings() bool`

HasYAxisSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


