# GraphVisualizationGaugePromql

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Configures a PromQL metrics query-based aggregation as a gauge | 
**Queries** | **[]string** | PromQL expressions configured for this widget. PromQL-backed widgets require at least one query. | 
**Aliases** | Pointer to [**GraphVisualizationTimeseriesPromqlAliases**](GraphVisualizationTimeseriesPromqlAliases.md) |  | [optional] 
**Max** | Pointer to **float32** | Gauge maximum value, on the raw query value before the unit is applied | [optional] 
**ColorThresholds** | Pointer to [**[]GaugeColorThreshold**](GaugeColorThreshold.md) | Color thresholds inside the gauge range, on the raw query value before the unit is applied | [optional] 
**Precision** | Pointer to [**GraphVisualizationQueryValueConnectionPrecision**](GraphVisualizationQueryValueConnectionPrecision.md) |  | [optional] 
**Normalizer** | Pointer to [**Normalizer**](Normalizer.md) |  | [optional] 

## Methods

### NewGraphVisualizationGaugePromql

`func NewGraphVisualizationGaugePromql(type_ string, queries []string, ) *GraphVisualizationGaugePromql`

NewGraphVisualizationGaugePromql instantiates a new GraphVisualizationGaugePromql object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphVisualizationGaugePromqlWithDefaults

`func NewGraphVisualizationGaugePromqlWithDefaults() *GraphVisualizationGaugePromql`

NewGraphVisualizationGaugePromqlWithDefaults instantiates a new GraphVisualizationGaugePromql object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GraphVisualizationGaugePromql) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GraphVisualizationGaugePromql) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GraphVisualizationGaugePromql) SetType(v string)`

SetType sets Type field to given value.


### GetQueries

`func (o *GraphVisualizationGaugePromql) GetQueries() []string`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *GraphVisualizationGaugePromql) GetQueriesOk() (*[]string, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *GraphVisualizationGaugePromql) SetQueries(v []string)`

SetQueries sets Queries field to given value.


### GetAliases

`func (o *GraphVisualizationGaugePromql) GetAliases() GraphVisualizationTimeseriesPromqlAliases`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *GraphVisualizationGaugePromql) GetAliasesOk() (*GraphVisualizationTimeseriesPromqlAliases, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *GraphVisualizationGaugePromql) SetAliases(v GraphVisualizationTimeseriesPromqlAliases)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *GraphVisualizationGaugePromql) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetMax

`func (o *GraphVisualizationGaugePromql) GetMax() float32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *GraphVisualizationGaugePromql) GetMaxOk() (*float32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *GraphVisualizationGaugePromql) SetMax(v float32)`

SetMax sets Max field to given value.

### HasMax

`func (o *GraphVisualizationGaugePromql) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetColorThresholds

`func (o *GraphVisualizationGaugePromql) GetColorThresholds() []GaugeColorThreshold`

GetColorThresholds returns the ColorThresholds field if non-nil, zero value otherwise.

### GetColorThresholdsOk

`func (o *GraphVisualizationGaugePromql) GetColorThresholdsOk() (*[]GaugeColorThreshold, bool)`

GetColorThresholdsOk returns a tuple with the ColorThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorThresholds

`func (o *GraphVisualizationGaugePromql) SetColorThresholds(v []GaugeColorThreshold)`

SetColorThresholds sets ColorThresholds field to given value.

### HasColorThresholds

`func (o *GraphVisualizationGaugePromql) HasColorThresholds() bool`

HasColorThresholds returns a boolean if a field has been set.

### GetPrecision

`func (o *GraphVisualizationGaugePromql) GetPrecision() GraphVisualizationQueryValueConnectionPrecision`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *GraphVisualizationGaugePromql) GetPrecisionOk() (*GraphVisualizationQueryValueConnectionPrecision, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *GraphVisualizationGaugePromql) SetPrecision(v GraphVisualizationQueryValueConnectionPrecision)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *GraphVisualizationGaugePromql) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetNormalizer

`func (o *GraphVisualizationGaugePromql) GetNormalizer() Normalizer`

GetNormalizer returns the Normalizer field if non-nil, zero value otherwise.

### GetNormalizerOk

`func (o *GraphVisualizationGaugePromql) GetNormalizerOk() (*Normalizer, bool)`

GetNormalizerOk returns a tuple with the Normalizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizer

`func (o *GraphVisualizationGaugePromql) SetNormalizer(v Normalizer)`

SetNormalizer sets Normalizer field to given value.

### HasNormalizer

`func (o *GraphVisualizationGaugePromql) HasNormalizer() bool`

HasNormalizer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


