# InputGraphVisualizationPiePromql

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays PromQL metrics query-based aggregation as a pie chart | 
**Queries** | **[]string** | PromQL expressions configured for this widget. PromQL-backed widgets require at least one query. | 
**Aliases** | Pointer to [**InputGraphVisualizationTimeseriesPromqlAliases**](InputGraphVisualizationTimeseriesPromqlAliases.md) |  | [optional] 
**Normalizer** | Pointer to [**Normalizer1**](Normalizer1.md) |  | [optional] 
**Precision** | Pointer to **float32** | Number of decimal places to display in the value | [optional] 
**LegendMode** | Pointer to **string** | Controls whether and how the widget displays legend or series details (e.g. table, legend-only, or no legend) | [optional] 

## Methods

### NewInputGraphVisualizationPiePromql

`func NewInputGraphVisualizationPiePromql(type_ string, queries []string, ) *InputGraphVisualizationPiePromql`

NewInputGraphVisualizationPiePromql instantiates a new InputGraphVisualizationPiePromql object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputGraphVisualizationPiePromqlWithDefaults

`func NewInputGraphVisualizationPiePromqlWithDefaults() *InputGraphVisualizationPiePromql`

NewInputGraphVisualizationPiePromqlWithDefaults instantiates a new InputGraphVisualizationPiePromql object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputGraphVisualizationPiePromql) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputGraphVisualizationPiePromql) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputGraphVisualizationPiePromql) SetType(v string)`

SetType sets Type field to given value.


### GetQueries

`func (o *InputGraphVisualizationPiePromql) GetQueries() []string`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *InputGraphVisualizationPiePromql) GetQueriesOk() (*[]string, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *InputGraphVisualizationPiePromql) SetQueries(v []string)`

SetQueries sets Queries field to given value.


### GetAliases

`func (o *InputGraphVisualizationPiePromql) GetAliases() InputGraphVisualizationTimeseriesPromqlAliases`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *InputGraphVisualizationPiePromql) GetAliasesOk() (*InputGraphVisualizationTimeseriesPromqlAliases, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *InputGraphVisualizationPiePromql) SetAliases(v InputGraphVisualizationTimeseriesPromqlAliases)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *InputGraphVisualizationPiePromql) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetNormalizer

`func (o *InputGraphVisualizationPiePromql) GetNormalizer() Normalizer1`

GetNormalizer returns the Normalizer field if non-nil, zero value otherwise.

### GetNormalizerOk

`func (o *InputGraphVisualizationPiePromql) GetNormalizerOk() (*Normalizer1, bool)`

GetNormalizerOk returns a tuple with the Normalizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizer

`func (o *InputGraphVisualizationPiePromql) SetNormalizer(v Normalizer1)`

SetNormalizer sets Normalizer field to given value.

### HasNormalizer

`func (o *InputGraphVisualizationPiePromql) HasNormalizer() bool`

HasNormalizer returns a boolean if a field has been set.

### GetPrecision

`func (o *InputGraphVisualizationPiePromql) GetPrecision() float32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *InputGraphVisualizationPiePromql) GetPrecisionOk() (*float32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *InputGraphVisualizationPiePromql) SetPrecision(v float32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *InputGraphVisualizationPiePromql) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetLegendMode

`func (o *InputGraphVisualizationPiePromql) GetLegendMode() string`

GetLegendMode returns the LegendMode field if non-nil, zero value otherwise.

### GetLegendModeOk

`func (o *InputGraphVisualizationPiePromql) GetLegendModeOk() (*string, bool)`

GetLegendModeOk returns a tuple with the LegendMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegendMode

`func (o *InputGraphVisualizationPiePromql) SetLegendMode(v string)`

SetLegendMode sets LegendMode field to given value.

### HasLegendMode

`func (o *InputGraphVisualizationPiePromql) HasLegendMode() bool`

HasLegendMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


