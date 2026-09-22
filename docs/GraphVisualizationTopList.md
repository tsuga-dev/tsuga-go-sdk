# GraphVisualizationTopList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays the aggregation as a ranked list of top results | 
**Source** | **string** | Telemetry source queried by this aggregation: &#x60;logs&#x60;, &#x60;metrics&#x60;, &#x60;traces&#x60;, or &#x60;rum&#x60;. | 
**Queries** | [**[]AggregationQuery**](AggregationQuery.md) | Aggregations that may be combined together in the same query. Each item is referenced from &#x60;formula&#x60; as q1, q2, and so on, in submission order. | 
**Formula** | Pointer to **string** | Formula referencing submitted query outputs, such as &#x60;q1 + q2&#x60;. References must be within &#x60;q1&#x60; through &#x60;qN&#x60; for the submitted queries. | [optional] 
**Aliases** | Pointer to [**GraphVisualizationTimeseriesPromqlAliases**](GraphVisualizationTimeseriesPromqlAliases.md) |  | [optional] 
**VisibleSeries** | Pointer to **[]bool** | Flags indicating whether each query or formula series is visible | [optional] 
**GroupBy** | Pointer to [**[]AggregationGroupBy1**](AggregationGroupBy1.md) | Nested grouping levels applied to the results, outermost first (e.g. group by service, then by level within each service). Each level splits results further, so the response contains one result per unique combination of group values. | [optional] 
**GroupByMode** | Pointer to **string** | &#x60;absolute&#x60; keeps each group at its own value; &#x60;relative&#x60; shows it as a percentage of the ungrouped total (defaults to absolute) | [optional] 
**Normalizer** | Pointer to [**Normalizer**](Normalizer.md) |  | [optional] 
**Precision** | Pointer to [**GraphVisualizationQueryValueConnectionPrecision**](GraphVisualizationQueryValueConnectionPrecision.md) |  | [optional] 
**Conditions** | Pointer to [**[]ConditionalFormatting**](ConditionalFormatting.md) | Conditional formatting rules applied to the displayed value | [optional] 
**IsStacked** | Pointer to **bool** | Requests stacked rendering for a top-list widget. Tsuga renders stacked rows only for one count or sum query with exactly two grouped fields, no formula, non-negative values, and a single-cluster context; otherwise the widget renders as a normal top list. | [optional] 

## Methods

### NewGraphVisualizationTopList

`func NewGraphVisualizationTopList(type_ string, source string, queries []AggregationQuery, ) *GraphVisualizationTopList`

NewGraphVisualizationTopList instantiates a new GraphVisualizationTopList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphVisualizationTopListWithDefaults

`func NewGraphVisualizationTopListWithDefaults() *GraphVisualizationTopList`

NewGraphVisualizationTopListWithDefaults instantiates a new GraphVisualizationTopList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GraphVisualizationTopList) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GraphVisualizationTopList) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GraphVisualizationTopList) SetType(v string)`

SetType sets Type field to given value.


### GetSource

`func (o *GraphVisualizationTopList) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GraphVisualizationTopList) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GraphVisualizationTopList) SetSource(v string)`

SetSource sets Source field to given value.


### GetQueries

`func (o *GraphVisualizationTopList) GetQueries() []AggregationQuery`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *GraphVisualizationTopList) GetQueriesOk() (*[]AggregationQuery, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *GraphVisualizationTopList) SetQueries(v []AggregationQuery)`

SetQueries sets Queries field to given value.


### GetFormula

`func (o *GraphVisualizationTopList) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *GraphVisualizationTopList) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *GraphVisualizationTopList) SetFormula(v string)`

SetFormula sets Formula field to given value.

### HasFormula

`func (o *GraphVisualizationTopList) HasFormula() bool`

HasFormula returns a boolean if a field has been set.

### GetAliases

`func (o *GraphVisualizationTopList) GetAliases() GraphVisualizationTimeseriesPromqlAliases`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *GraphVisualizationTopList) GetAliasesOk() (*GraphVisualizationTimeseriesPromqlAliases, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *GraphVisualizationTopList) SetAliases(v GraphVisualizationTimeseriesPromqlAliases)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *GraphVisualizationTopList) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetVisibleSeries

`func (o *GraphVisualizationTopList) GetVisibleSeries() []bool`

GetVisibleSeries returns the VisibleSeries field if non-nil, zero value otherwise.

### GetVisibleSeriesOk

`func (o *GraphVisualizationTopList) GetVisibleSeriesOk() (*[]bool, bool)`

GetVisibleSeriesOk returns a tuple with the VisibleSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleSeries

`func (o *GraphVisualizationTopList) SetVisibleSeries(v []bool)`

SetVisibleSeries sets VisibleSeries field to given value.

### HasVisibleSeries

`func (o *GraphVisualizationTopList) HasVisibleSeries() bool`

HasVisibleSeries returns a boolean if a field has been set.

### GetGroupBy

`func (o *GraphVisualizationTopList) GetGroupBy() []AggregationGroupBy1`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *GraphVisualizationTopList) GetGroupByOk() (*[]AggregationGroupBy1, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *GraphVisualizationTopList) SetGroupBy(v []AggregationGroupBy1)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *GraphVisualizationTopList) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetGroupByMode

`func (o *GraphVisualizationTopList) GetGroupByMode() string`

GetGroupByMode returns the GroupByMode field if non-nil, zero value otherwise.

### GetGroupByModeOk

`func (o *GraphVisualizationTopList) GetGroupByModeOk() (*string, bool)`

GetGroupByModeOk returns a tuple with the GroupByMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByMode

`func (o *GraphVisualizationTopList) SetGroupByMode(v string)`

SetGroupByMode sets GroupByMode field to given value.

### HasGroupByMode

`func (o *GraphVisualizationTopList) HasGroupByMode() bool`

HasGroupByMode returns a boolean if a field has been set.

### GetNormalizer

`func (o *GraphVisualizationTopList) GetNormalizer() Normalizer`

GetNormalizer returns the Normalizer field if non-nil, zero value otherwise.

### GetNormalizerOk

`func (o *GraphVisualizationTopList) GetNormalizerOk() (*Normalizer, bool)`

GetNormalizerOk returns a tuple with the Normalizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizer

`func (o *GraphVisualizationTopList) SetNormalizer(v Normalizer)`

SetNormalizer sets Normalizer field to given value.

### HasNormalizer

`func (o *GraphVisualizationTopList) HasNormalizer() bool`

HasNormalizer returns a boolean if a field has been set.

### GetPrecision

`func (o *GraphVisualizationTopList) GetPrecision() GraphVisualizationQueryValueConnectionPrecision`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *GraphVisualizationTopList) GetPrecisionOk() (*GraphVisualizationQueryValueConnectionPrecision, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *GraphVisualizationTopList) SetPrecision(v GraphVisualizationQueryValueConnectionPrecision)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *GraphVisualizationTopList) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetConditions

`func (o *GraphVisualizationTopList) GetConditions() []ConditionalFormatting`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *GraphVisualizationTopList) GetConditionsOk() (*[]ConditionalFormatting, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *GraphVisualizationTopList) SetConditions(v []ConditionalFormatting)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *GraphVisualizationTopList) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetIsStacked

`func (o *GraphVisualizationTopList) GetIsStacked() bool`

GetIsStacked returns the IsStacked field if non-nil, zero value otherwise.

### GetIsStackedOk

`func (o *GraphVisualizationTopList) GetIsStackedOk() (*bool, bool)`

GetIsStackedOk returns a tuple with the IsStacked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStacked

`func (o *GraphVisualizationTopList) SetIsStacked(v bool)`

SetIsStacked sets IsStacked field to given value.

### HasIsStacked

`func (o *GraphVisualizationTopList) HasIsStacked() bool`

HasIsStacked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


