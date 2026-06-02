# InputGraphVisualizationQueryValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays the aggregation as a single value | 
**Source** | **string** | Data source being queried for this aggregation | 
**Queries** | [**[]AggregationQuery1**](AggregationQuery1.md) | Aggregations that may be combined together in the same query | 
**Formula** | Pointer to **string** | Formula referencing query outputs (e.g. q1+q2) to compute derived series | [optional] 
**Aliases** | Pointer to [**InputGraphVisualizationTimeseriesAliases**](InputGraphVisualizationTimeseriesAliases.md) |  | [optional] 
**VisibleSeries** | Pointer to **[]bool** | Flags indicating whether each query or formula series is visible | [optional] 
**BackgroundMode** | Pointer to **string** | Controls whether the widget uses a solid or transparent background | [optional] 
**Conditions** | Pointer to [**[]ConditionalFormatting**](ConditionalFormatting.md) | Conditional formatting rules applied to the displayed value | [optional] 
**Precision** | Pointer to **float32** | Number of decimal places to display in the value | [optional] 
**Normalizer** | Pointer to [**Normalizer1**](Normalizer1.md) |  | [optional] 
**LegendMode** | Pointer to **string** | Controls whether and how the widget displays legend or series details (e.g. table, legend-only, or no legend) | [optional] 

## Methods

### NewInputGraphVisualizationQueryValue

`func NewInputGraphVisualizationQueryValue(type_ string, source string, queries []AggregationQuery1, ) *InputGraphVisualizationQueryValue`

NewInputGraphVisualizationQueryValue instantiates a new InputGraphVisualizationQueryValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputGraphVisualizationQueryValueWithDefaults

`func NewInputGraphVisualizationQueryValueWithDefaults() *InputGraphVisualizationQueryValue`

NewInputGraphVisualizationQueryValueWithDefaults instantiates a new InputGraphVisualizationQueryValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputGraphVisualizationQueryValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputGraphVisualizationQueryValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputGraphVisualizationQueryValue) SetType(v string)`

SetType sets Type field to given value.


### GetSource

`func (o *InputGraphVisualizationQueryValue) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InputGraphVisualizationQueryValue) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InputGraphVisualizationQueryValue) SetSource(v string)`

SetSource sets Source field to given value.


### GetQueries

`func (o *InputGraphVisualizationQueryValue) GetQueries() []AggregationQuery1`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *InputGraphVisualizationQueryValue) GetQueriesOk() (*[]AggregationQuery1, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *InputGraphVisualizationQueryValue) SetQueries(v []AggregationQuery1)`

SetQueries sets Queries field to given value.


### GetFormula

`func (o *InputGraphVisualizationQueryValue) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *InputGraphVisualizationQueryValue) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *InputGraphVisualizationQueryValue) SetFormula(v string)`

SetFormula sets Formula field to given value.

### HasFormula

`func (o *InputGraphVisualizationQueryValue) HasFormula() bool`

HasFormula returns a boolean if a field has been set.

### GetAliases

`func (o *InputGraphVisualizationQueryValue) GetAliases() InputGraphVisualizationTimeseriesAliases`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *InputGraphVisualizationQueryValue) GetAliasesOk() (*InputGraphVisualizationTimeseriesAliases, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *InputGraphVisualizationQueryValue) SetAliases(v InputGraphVisualizationTimeseriesAliases)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *InputGraphVisualizationQueryValue) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetVisibleSeries

`func (o *InputGraphVisualizationQueryValue) GetVisibleSeries() []bool`

GetVisibleSeries returns the VisibleSeries field if non-nil, zero value otherwise.

### GetVisibleSeriesOk

`func (o *InputGraphVisualizationQueryValue) GetVisibleSeriesOk() (*[]bool, bool)`

GetVisibleSeriesOk returns a tuple with the VisibleSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleSeries

`func (o *InputGraphVisualizationQueryValue) SetVisibleSeries(v []bool)`

SetVisibleSeries sets VisibleSeries field to given value.

### HasVisibleSeries

`func (o *InputGraphVisualizationQueryValue) HasVisibleSeries() bool`

HasVisibleSeries returns a boolean if a field has been set.

### GetBackgroundMode

`func (o *InputGraphVisualizationQueryValue) GetBackgroundMode() string`

GetBackgroundMode returns the BackgroundMode field if non-nil, zero value otherwise.

### GetBackgroundModeOk

`func (o *InputGraphVisualizationQueryValue) GetBackgroundModeOk() (*string, bool)`

GetBackgroundModeOk returns a tuple with the BackgroundMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundMode

`func (o *InputGraphVisualizationQueryValue) SetBackgroundMode(v string)`

SetBackgroundMode sets BackgroundMode field to given value.

### HasBackgroundMode

`func (o *InputGraphVisualizationQueryValue) HasBackgroundMode() bool`

HasBackgroundMode returns a boolean if a field has been set.

### GetConditions

`func (o *InputGraphVisualizationQueryValue) GetConditions() []ConditionalFormatting`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *InputGraphVisualizationQueryValue) GetConditionsOk() (*[]ConditionalFormatting, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *InputGraphVisualizationQueryValue) SetConditions(v []ConditionalFormatting)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *InputGraphVisualizationQueryValue) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetPrecision

`func (o *InputGraphVisualizationQueryValue) GetPrecision() float32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *InputGraphVisualizationQueryValue) GetPrecisionOk() (*float32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *InputGraphVisualizationQueryValue) SetPrecision(v float32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *InputGraphVisualizationQueryValue) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetNormalizer

`func (o *InputGraphVisualizationQueryValue) GetNormalizer() Normalizer1`

GetNormalizer returns the Normalizer field if non-nil, zero value otherwise.

### GetNormalizerOk

`func (o *InputGraphVisualizationQueryValue) GetNormalizerOk() (*Normalizer1, bool)`

GetNormalizerOk returns a tuple with the Normalizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizer

`func (o *InputGraphVisualizationQueryValue) SetNormalizer(v Normalizer1)`

SetNormalizer sets Normalizer field to given value.

### HasNormalizer

`func (o *InputGraphVisualizationQueryValue) HasNormalizer() bool`

HasNormalizer returns a boolean if a field has been set.

### GetLegendMode

`func (o *InputGraphVisualizationQueryValue) GetLegendMode() string`

GetLegendMode returns the LegendMode field if non-nil, zero value otherwise.

### GetLegendModeOk

`func (o *InputGraphVisualizationQueryValue) GetLegendModeOk() (*string, bool)`

GetLegendModeOk returns a tuple with the LegendMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegendMode

`func (o *InputGraphVisualizationQueryValue) SetLegendMode(v string)`

SetLegendMode sets LegendMode field to given value.

### HasLegendMode

`func (o *InputGraphVisualizationQueryValue) HasLegendMode() bool`

HasLegendMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


