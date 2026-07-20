# GraphVisualizationPiePromql

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays PromQL metrics query-based aggregation as a pie chart | 
**Queries** | **[]string** | PromQL expressions configured for this widget. PromQL-backed widgets require at least one query. | 
**Aliases** | Pointer to [**GraphVisualizationTimeseriesPromqlAliases**](GraphVisualizationTimeseriesPromqlAliases.md) |  | [optional] 
**Normalizer** | Pointer to [**Normalizer**](Normalizer.md) |  | [optional] 
**Precision** | Pointer to **float32** | Number of decimal places to display in the value | [optional] 
**LegendMode** | Pointer to **string** | Controls whether and how the widget displays legend or series details (e.g. table, legend-only, or no legend) | [optional] 

## Methods

### NewGraphVisualizationPiePromql

`func NewGraphVisualizationPiePromql(type_ string, queries []string, ) *GraphVisualizationPiePromql`

NewGraphVisualizationPiePromql instantiates a new GraphVisualizationPiePromql object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphVisualizationPiePromqlWithDefaults

`func NewGraphVisualizationPiePromqlWithDefaults() *GraphVisualizationPiePromql`

NewGraphVisualizationPiePromqlWithDefaults instantiates a new GraphVisualizationPiePromql object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GraphVisualizationPiePromql) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GraphVisualizationPiePromql) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GraphVisualizationPiePromql) SetType(v string)`

SetType sets Type field to given value.


### GetQueries

`func (o *GraphVisualizationPiePromql) GetQueries() []string`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *GraphVisualizationPiePromql) GetQueriesOk() (*[]string, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *GraphVisualizationPiePromql) SetQueries(v []string)`

SetQueries sets Queries field to given value.


### GetAliases

`func (o *GraphVisualizationPiePromql) GetAliases() GraphVisualizationTimeseriesPromqlAliases`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *GraphVisualizationPiePromql) GetAliasesOk() (*GraphVisualizationTimeseriesPromqlAliases, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *GraphVisualizationPiePromql) SetAliases(v GraphVisualizationTimeseriesPromqlAliases)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *GraphVisualizationPiePromql) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetNormalizer

`func (o *GraphVisualizationPiePromql) GetNormalizer() Normalizer`

GetNormalizer returns the Normalizer field if non-nil, zero value otherwise.

### GetNormalizerOk

`func (o *GraphVisualizationPiePromql) GetNormalizerOk() (*Normalizer, bool)`

GetNormalizerOk returns a tuple with the Normalizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizer

`func (o *GraphVisualizationPiePromql) SetNormalizer(v Normalizer)`

SetNormalizer sets Normalizer field to given value.

### HasNormalizer

`func (o *GraphVisualizationPiePromql) HasNormalizer() bool`

HasNormalizer returns a boolean if a field has been set.

### GetPrecision

`func (o *GraphVisualizationPiePromql) GetPrecision() float32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *GraphVisualizationPiePromql) GetPrecisionOk() (*float32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *GraphVisualizationPiePromql) SetPrecision(v float32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *GraphVisualizationPiePromql) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetLegendMode

`func (o *GraphVisualizationPiePromql) GetLegendMode() string`

GetLegendMode returns the LegendMode field if non-nil, zero value otherwise.

### GetLegendModeOk

`func (o *GraphVisualizationPiePromql) GetLegendModeOk() (*string, bool)`

GetLegendModeOk returns a tuple with the LegendMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegendMode

`func (o *GraphVisualizationPiePromql) SetLegendMode(v string)`

SetLegendMode sets LegendMode field to given value.

### HasLegendMode

`func (o *GraphVisualizationPiePromql) HasLegendMode() bool`

HasLegendMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


