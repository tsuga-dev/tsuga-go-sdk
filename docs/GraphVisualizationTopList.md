# GraphVisualizationTopList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays the aggregation as a ranked list of top results | 
**Source** | **string** | Data source being queried for this aggregation | 
**Queries** | [**[]GraphVisualizationTimeseriesQueriesInner**](GraphVisualizationTimeseriesQueriesInner.md) | Aggregations that may be combined together in the same query | 
**Formula** | Pointer to **string** | Formula referencing query outputs (e.g. q1+q2) to compute derived series | [optional] 
**VisibleSeries** | Pointer to **[]bool** | Flags indicating whether each query or formula series is visible | [optional] 
**GroupBy** | Pointer to [**[]GraphVisualizationTimeseriesGroupByInner**](GraphVisualizationTimeseriesGroupByInner.md) | Fields used to group the results | [optional] 
**Normalizer** | Pointer to [**GraphVisualizationTimeseriesNormalizer**](GraphVisualizationTimeseriesNormalizer.md) |  | [optional] 

## Methods

### NewGraphVisualizationTopList

`func NewGraphVisualizationTopList(type_ string, source string, queries []GraphVisualizationTimeseriesQueriesInner, ) *GraphVisualizationTopList`

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

`func (o *GraphVisualizationTopList) GetQueries() []GraphVisualizationTimeseriesQueriesInner`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *GraphVisualizationTopList) GetQueriesOk() (*[]GraphVisualizationTimeseriesQueriesInner, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *GraphVisualizationTopList) SetQueries(v []GraphVisualizationTimeseriesQueriesInner)`

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

`func (o *GraphVisualizationTopList) GetGroupBy() []GraphVisualizationTimeseriesGroupByInner`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *GraphVisualizationTopList) GetGroupByOk() (*[]GraphVisualizationTimeseriesGroupByInner, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *GraphVisualizationTopList) SetGroupBy(v []GraphVisualizationTimeseriesGroupByInner)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *GraphVisualizationTopList) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetNormalizer

`func (o *GraphVisualizationTopList) GetNormalizer() GraphVisualizationTimeseriesNormalizer`

GetNormalizer returns the Normalizer field if non-nil, zero value otherwise.

### GetNormalizerOk

`func (o *GraphVisualizationTopList) GetNormalizerOk() (*GraphVisualizationTimeseriesNormalizer, bool)`

GetNormalizerOk returns a tuple with the Normalizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizer

`func (o *GraphVisualizationTopList) SetNormalizer(v GraphVisualizationTimeseriesNormalizer)`

SetNormalizer sets Normalizer field to given value.

### HasNormalizer

`func (o *GraphVisualizationTopList) HasNormalizer() bool`

HasNormalizer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


