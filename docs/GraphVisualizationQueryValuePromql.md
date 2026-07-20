# GraphVisualizationQueryValuePromql

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Configures a PromQL metrics query-based aggregation as a single value | 
**Queries** | **[]string** | PromQL expressions configured for this widget. PromQL-backed widgets require at least one query. | 
**Aliases** | Pointer to [**GraphVisualizationTimeseriesPromqlAliases**](GraphVisualizationTimeseriesPromqlAliases.md) |  | [optional] 
**BackgroundMode** | Pointer to **string** | Controls whether the widget uses a solid or transparent background | [optional] 
**Conditions** | Pointer to [**[]ConditionalFormatting**](ConditionalFormatting.md) | Conditional formatting rules applied to the displayed value | [optional] 
**Precision** | Pointer to **float32** | Number of decimal places to display in the value | [optional] 
**Normalizer** | Pointer to [**Normalizer**](Normalizer.md) |  | [optional] 
**LegendMode** | Pointer to **string** | Controls whether and how the widget displays legend or series details (e.g. table, legend-only, or no legend) | [optional] 

## Methods

### NewGraphVisualizationQueryValuePromql

`func NewGraphVisualizationQueryValuePromql(type_ string, queries []string, ) *GraphVisualizationQueryValuePromql`

NewGraphVisualizationQueryValuePromql instantiates a new GraphVisualizationQueryValuePromql object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphVisualizationQueryValuePromqlWithDefaults

`func NewGraphVisualizationQueryValuePromqlWithDefaults() *GraphVisualizationQueryValuePromql`

NewGraphVisualizationQueryValuePromqlWithDefaults instantiates a new GraphVisualizationQueryValuePromql object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GraphVisualizationQueryValuePromql) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GraphVisualizationQueryValuePromql) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GraphVisualizationQueryValuePromql) SetType(v string)`

SetType sets Type field to given value.


### GetQueries

`func (o *GraphVisualizationQueryValuePromql) GetQueries() []string`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *GraphVisualizationQueryValuePromql) GetQueriesOk() (*[]string, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *GraphVisualizationQueryValuePromql) SetQueries(v []string)`

SetQueries sets Queries field to given value.


### GetAliases

`func (o *GraphVisualizationQueryValuePromql) GetAliases() GraphVisualizationTimeseriesPromqlAliases`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *GraphVisualizationQueryValuePromql) GetAliasesOk() (*GraphVisualizationTimeseriesPromqlAliases, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *GraphVisualizationQueryValuePromql) SetAliases(v GraphVisualizationTimeseriesPromqlAliases)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *GraphVisualizationQueryValuePromql) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetBackgroundMode

`func (o *GraphVisualizationQueryValuePromql) GetBackgroundMode() string`

GetBackgroundMode returns the BackgroundMode field if non-nil, zero value otherwise.

### GetBackgroundModeOk

`func (o *GraphVisualizationQueryValuePromql) GetBackgroundModeOk() (*string, bool)`

GetBackgroundModeOk returns a tuple with the BackgroundMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundMode

`func (o *GraphVisualizationQueryValuePromql) SetBackgroundMode(v string)`

SetBackgroundMode sets BackgroundMode field to given value.

### HasBackgroundMode

`func (o *GraphVisualizationQueryValuePromql) HasBackgroundMode() bool`

HasBackgroundMode returns a boolean if a field has been set.

### GetConditions

`func (o *GraphVisualizationQueryValuePromql) GetConditions() []ConditionalFormatting`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *GraphVisualizationQueryValuePromql) GetConditionsOk() (*[]ConditionalFormatting, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *GraphVisualizationQueryValuePromql) SetConditions(v []ConditionalFormatting)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *GraphVisualizationQueryValuePromql) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetPrecision

`func (o *GraphVisualizationQueryValuePromql) GetPrecision() float32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *GraphVisualizationQueryValuePromql) GetPrecisionOk() (*float32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *GraphVisualizationQueryValuePromql) SetPrecision(v float32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *GraphVisualizationQueryValuePromql) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetNormalizer

`func (o *GraphVisualizationQueryValuePromql) GetNormalizer() Normalizer`

GetNormalizer returns the Normalizer field if non-nil, zero value otherwise.

### GetNormalizerOk

`func (o *GraphVisualizationQueryValuePromql) GetNormalizerOk() (*Normalizer, bool)`

GetNormalizerOk returns a tuple with the Normalizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizer

`func (o *GraphVisualizationQueryValuePromql) SetNormalizer(v Normalizer)`

SetNormalizer sets Normalizer field to given value.

### HasNormalizer

`func (o *GraphVisualizationQueryValuePromql) HasNormalizer() bool`

HasNormalizer returns a boolean if a field has been set.

### GetLegendMode

`func (o *GraphVisualizationQueryValuePromql) GetLegendMode() string`

GetLegendMode returns the LegendMode field if non-nil, zero value otherwise.

### GetLegendModeOk

`func (o *GraphVisualizationQueryValuePromql) GetLegendModeOk() (*string, bool)`

GetLegendModeOk returns a tuple with the LegendMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegendMode

`func (o *GraphVisualizationQueryValuePromql) SetLegendMode(v string)`

SetLegendMode sets LegendMode field to given value.

### HasLegendMode

`func (o *GraphVisualizationQueryValuePromql) HasLegendMode() bool`

HasLegendMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


