# InputGraphVisualizationPie

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays the aggregation as a pie chart | 
**Source** | **string** | Data source being queried for this aggregation | 
**Queries** | [**[]AggregationQuery1**](AggregationQuery1.md) | Aggregations that may be combined together in the same query | 
**Formula** | Pointer to **string** | Formula referencing query outputs (e.g. q1+q2) to compute derived series | [optional] 
**Aliases** | Pointer to [**InputGraphVisualizationTimeseriesAliases**](InputGraphVisualizationTimeseriesAliases.md) |  | [optional] 
**VisibleSeries** | Pointer to **[]bool** | Flags indicating whether each query or formula series is visible | [optional] 
**GroupBy** | Pointer to [**[]AggregationGroupBy1**](AggregationGroupBy1.md) | Fields used to group the results | [optional] 
**Normalizer** | Pointer to [**Normalizer1**](Normalizer1.md) |  | [optional] 
**Precision** | Pointer to **float32** | Number of decimal places to display in the value | [optional] 
**LegendMode** | Pointer to **string** | Controls whether and how the widget displays legend or series details (e.g. table, legend-only, or no legend) | [optional] 

## Methods

### NewInputGraphVisualizationPie

`func NewInputGraphVisualizationPie(type_ string, source string, queries []AggregationQuery1, ) *InputGraphVisualizationPie`

NewInputGraphVisualizationPie instantiates a new InputGraphVisualizationPie object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputGraphVisualizationPieWithDefaults

`func NewInputGraphVisualizationPieWithDefaults() *InputGraphVisualizationPie`

NewInputGraphVisualizationPieWithDefaults instantiates a new InputGraphVisualizationPie object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputGraphVisualizationPie) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputGraphVisualizationPie) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputGraphVisualizationPie) SetType(v string)`

SetType sets Type field to given value.


### GetSource

`func (o *InputGraphVisualizationPie) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InputGraphVisualizationPie) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InputGraphVisualizationPie) SetSource(v string)`

SetSource sets Source field to given value.


### GetQueries

`func (o *InputGraphVisualizationPie) GetQueries() []AggregationQuery1`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *InputGraphVisualizationPie) GetQueriesOk() (*[]AggregationQuery1, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *InputGraphVisualizationPie) SetQueries(v []AggregationQuery1)`

SetQueries sets Queries field to given value.


### GetFormula

`func (o *InputGraphVisualizationPie) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *InputGraphVisualizationPie) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *InputGraphVisualizationPie) SetFormula(v string)`

SetFormula sets Formula field to given value.

### HasFormula

`func (o *InputGraphVisualizationPie) HasFormula() bool`

HasFormula returns a boolean if a field has been set.

### GetAliases

`func (o *InputGraphVisualizationPie) GetAliases() InputGraphVisualizationTimeseriesAliases`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *InputGraphVisualizationPie) GetAliasesOk() (*InputGraphVisualizationTimeseriesAliases, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *InputGraphVisualizationPie) SetAliases(v InputGraphVisualizationTimeseriesAliases)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *InputGraphVisualizationPie) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetVisibleSeries

`func (o *InputGraphVisualizationPie) GetVisibleSeries() []bool`

GetVisibleSeries returns the VisibleSeries field if non-nil, zero value otherwise.

### GetVisibleSeriesOk

`func (o *InputGraphVisualizationPie) GetVisibleSeriesOk() (*[]bool, bool)`

GetVisibleSeriesOk returns a tuple with the VisibleSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleSeries

`func (o *InputGraphVisualizationPie) SetVisibleSeries(v []bool)`

SetVisibleSeries sets VisibleSeries field to given value.

### HasVisibleSeries

`func (o *InputGraphVisualizationPie) HasVisibleSeries() bool`

HasVisibleSeries returns a boolean if a field has been set.

### GetGroupBy

`func (o *InputGraphVisualizationPie) GetGroupBy() []AggregationGroupBy1`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *InputGraphVisualizationPie) GetGroupByOk() (*[]AggregationGroupBy1, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *InputGraphVisualizationPie) SetGroupBy(v []AggregationGroupBy1)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *InputGraphVisualizationPie) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetNormalizer

`func (o *InputGraphVisualizationPie) GetNormalizer() Normalizer1`

GetNormalizer returns the Normalizer field if non-nil, zero value otherwise.

### GetNormalizerOk

`func (o *InputGraphVisualizationPie) GetNormalizerOk() (*Normalizer1, bool)`

GetNormalizerOk returns a tuple with the Normalizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizer

`func (o *InputGraphVisualizationPie) SetNormalizer(v Normalizer1)`

SetNormalizer sets Normalizer field to given value.

### HasNormalizer

`func (o *InputGraphVisualizationPie) HasNormalizer() bool`

HasNormalizer returns a boolean if a field has been set.

### GetPrecision

`func (o *InputGraphVisualizationPie) GetPrecision() float32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *InputGraphVisualizationPie) GetPrecisionOk() (*float32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *InputGraphVisualizationPie) SetPrecision(v float32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *InputGraphVisualizationPie) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetLegendMode

`func (o *InputGraphVisualizationPie) GetLegendMode() string`

GetLegendMode returns the LegendMode field if non-nil, zero value otherwise.

### GetLegendModeOk

`func (o *InputGraphVisualizationPie) GetLegendModeOk() (*string, bool)`

GetLegendModeOk returns a tuple with the LegendMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegendMode

`func (o *InputGraphVisualizationPie) SetLegendMode(v string)`

SetLegendMode sets LegendMode field to given value.

### HasLegendMode

`func (o *InputGraphVisualizationPie) HasLegendMode() bool`

HasLegendMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


