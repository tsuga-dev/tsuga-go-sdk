# InputGraphVisualizationDistributionPromql

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays a PromQL histogram query grouped by &#x60;le&#x60; as a distribution chart, e.g. &#x60;sum by (le) (rate(http_duration_bucket[5m]))&#x60; | 
**Queries** | **[]string** | PromQL expressions configured for this widget. Only the first expression is rendered; later entries are stored and ignored. | 
**Precision** | Pointer to [**GraphVisualizationQueryValueConnectionPrecision**](GraphVisualizationQueryValueConnectionPrecision.md) |  | [optional] 
**Normalizer** | Pointer to [**InputNormalizer**](InputNormalizer.md) |  | [optional] 
**PercentileMarkers** | Pointer to **[]int32** | Percentile markers displayed on top of the distribution chart | [optional] 

## Methods

### NewInputGraphVisualizationDistributionPromql

`func NewInputGraphVisualizationDistributionPromql(type_ string, queries []string, ) *InputGraphVisualizationDistributionPromql`

NewInputGraphVisualizationDistributionPromql instantiates a new InputGraphVisualizationDistributionPromql object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputGraphVisualizationDistributionPromqlWithDefaults

`func NewInputGraphVisualizationDistributionPromqlWithDefaults() *InputGraphVisualizationDistributionPromql`

NewInputGraphVisualizationDistributionPromqlWithDefaults instantiates a new InputGraphVisualizationDistributionPromql object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputGraphVisualizationDistributionPromql) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputGraphVisualizationDistributionPromql) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputGraphVisualizationDistributionPromql) SetType(v string)`

SetType sets Type field to given value.


### GetQueries

`func (o *InputGraphVisualizationDistributionPromql) GetQueries() []string`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *InputGraphVisualizationDistributionPromql) GetQueriesOk() (*[]string, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *InputGraphVisualizationDistributionPromql) SetQueries(v []string)`

SetQueries sets Queries field to given value.


### GetPrecision

`func (o *InputGraphVisualizationDistributionPromql) GetPrecision() GraphVisualizationQueryValueConnectionPrecision`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *InputGraphVisualizationDistributionPromql) GetPrecisionOk() (*GraphVisualizationQueryValueConnectionPrecision, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *InputGraphVisualizationDistributionPromql) SetPrecision(v GraphVisualizationQueryValueConnectionPrecision)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *InputGraphVisualizationDistributionPromql) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetNormalizer

`func (o *InputGraphVisualizationDistributionPromql) GetNormalizer() InputNormalizer`

GetNormalizer returns the Normalizer field if non-nil, zero value otherwise.

### GetNormalizerOk

`func (o *InputGraphVisualizationDistributionPromql) GetNormalizerOk() (*InputNormalizer, bool)`

GetNormalizerOk returns a tuple with the Normalizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizer

`func (o *InputGraphVisualizationDistributionPromql) SetNormalizer(v InputNormalizer)`

SetNormalizer sets Normalizer field to given value.

### HasNormalizer

`func (o *InputGraphVisualizationDistributionPromql) HasNormalizer() bool`

HasNormalizer returns a boolean if a field has been set.

### GetPercentileMarkers

`func (o *InputGraphVisualizationDistributionPromql) GetPercentileMarkers() []int32`

GetPercentileMarkers returns the PercentileMarkers field if non-nil, zero value otherwise.

### GetPercentileMarkersOk

`func (o *InputGraphVisualizationDistributionPromql) GetPercentileMarkersOk() (*[]int32, bool)`

GetPercentileMarkersOk returns a tuple with the PercentileMarkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentileMarkers

`func (o *InputGraphVisualizationDistributionPromql) SetPercentileMarkers(v []int32)`

SetPercentileMarkers sets PercentileMarkers field to given value.

### HasPercentileMarkers

`func (o *InputGraphVisualizationDistributionPromql) HasPercentileMarkers() bool`

HasPercentileMarkers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


