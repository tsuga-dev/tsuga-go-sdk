# InputGraphVisualizationQueryValuePromql

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Configures a PromQL metrics query-based aggregation as a single value | 
**Queries** | **[]string** | PromQL expressions configured for this widget. PromQL-backed widgets require at least one query. | 
**Aliases** | Pointer to [**InputGraphVisualizationTimeseriesPromqlAliases**](InputGraphVisualizationTimeseriesPromqlAliases.md) |  | [optional] 
**BackgroundMode** | Pointer to **string** | Controls whether the widget uses a solid or transparent background | [optional] 
**Conditions** | Pointer to [**[]ConditionalFormatting**](ConditionalFormatting.md) | Conditional formatting rules applied to the displayed value | [optional] 
**Precision** | Pointer to **float32** | Number of decimal places to display in the value | [optional] 
**Normalizer** | Pointer to [**Normalizer1**](Normalizer1.md) |  | [optional] 
**LegendMode** | Pointer to **string** | Controls whether and how the widget displays legend or series details (e.g. table, legend-only, or no legend) | [optional] 

## Methods

### NewInputGraphVisualizationQueryValuePromql

`func NewInputGraphVisualizationQueryValuePromql(type_ string, queries []string, ) *InputGraphVisualizationQueryValuePromql`

NewInputGraphVisualizationQueryValuePromql instantiates a new InputGraphVisualizationQueryValuePromql object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputGraphVisualizationQueryValuePromqlWithDefaults

`func NewInputGraphVisualizationQueryValuePromqlWithDefaults() *InputGraphVisualizationQueryValuePromql`

NewInputGraphVisualizationQueryValuePromqlWithDefaults instantiates a new InputGraphVisualizationQueryValuePromql object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputGraphVisualizationQueryValuePromql) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputGraphVisualizationQueryValuePromql) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputGraphVisualizationQueryValuePromql) SetType(v string)`

SetType sets Type field to given value.


### GetQueries

`func (o *InputGraphVisualizationQueryValuePromql) GetQueries() []string`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *InputGraphVisualizationQueryValuePromql) GetQueriesOk() (*[]string, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *InputGraphVisualizationQueryValuePromql) SetQueries(v []string)`

SetQueries sets Queries field to given value.


### GetAliases

`func (o *InputGraphVisualizationQueryValuePromql) GetAliases() InputGraphVisualizationTimeseriesPromqlAliases`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *InputGraphVisualizationQueryValuePromql) GetAliasesOk() (*InputGraphVisualizationTimeseriesPromqlAliases, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *InputGraphVisualizationQueryValuePromql) SetAliases(v InputGraphVisualizationTimeseriesPromqlAliases)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *InputGraphVisualizationQueryValuePromql) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetBackgroundMode

`func (o *InputGraphVisualizationQueryValuePromql) GetBackgroundMode() string`

GetBackgroundMode returns the BackgroundMode field if non-nil, zero value otherwise.

### GetBackgroundModeOk

`func (o *InputGraphVisualizationQueryValuePromql) GetBackgroundModeOk() (*string, bool)`

GetBackgroundModeOk returns a tuple with the BackgroundMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundMode

`func (o *InputGraphVisualizationQueryValuePromql) SetBackgroundMode(v string)`

SetBackgroundMode sets BackgroundMode field to given value.

### HasBackgroundMode

`func (o *InputGraphVisualizationQueryValuePromql) HasBackgroundMode() bool`

HasBackgroundMode returns a boolean if a field has been set.

### GetConditions

`func (o *InputGraphVisualizationQueryValuePromql) GetConditions() []ConditionalFormatting`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *InputGraphVisualizationQueryValuePromql) GetConditionsOk() (*[]ConditionalFormatting, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *InputGraphVisualizationQueryValuePromql) SetConditions(v []ConditionalFormatting)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *InputGraphVisualizationQueryValuePromql) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetPrecision

`func (o *InputGraphVisualizationQueryValuePromql) GetPrecision() float32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *InputGraphVisualizationQueryValuePromql) GetPrecisionOk() (*float32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *InputGraphVisualizationQueryValuePromql) SetPrecision(v float32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *InputGraphVisualizationQueryValuePromql) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetNormalizer

`func (o *InputGraphVisualizationQueryValuePromql) GetNormalizer() Normalizer1`

GetNormalizer returns the Normalizer field if non-nil, zero value otherwise.

### GetNormalizerOk

`func (o *InputGraphVisualizationQueryValuePromql) GetNormalizerOk() (*Normalizer1, bool)`

GetNormalizerOk returns a tuple with the Normalizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizer

`func (o *InputGraphVisualizationQueryValuePromql) SetNormalizer(v Normalizer1)`

SetNormalizer sets Normalizer field to given value.

### HasNormalizer

`func (o *InputGraphVisualizationQueryValuePromql) HasNormalizer() bool`

HasNormalizer returns a boolean if a field has been set.

### GetLegendMode

`func (o *InputGraphVisualizationQueryValuePromql) GetLegendMode() string`

GetLegendMode returns the LegendMode field if non-nil, zero value otherwise.

### GetLegendModeOk

`func (o *InputGraphVisualizationQueryValuePromql) GetLegendModeOk() (*string, bool)`

GetLegendModeOk returns a tuple with the LegendMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegendMode

`func (o *InputGraphVisualizationQueryValuePromql) SetLegendMode(v string)`

SetLegendMode sets LegendMode field to given value.

### HasLegendMode

`func (o *InputGraphVisualizationQueryValuePromql) HasLegendMode() bool`

HasLegendMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


