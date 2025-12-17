# GraphVisualizationBar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays the aggregation as a bar chart | 
**Source** | **string** | Data source being queried for this aggregation | 
**Queries** | [**[]AggregationQuery**](AggregationQuery.md) | Aggregations that may be combined together in the same query | 
**Formula** | Pointer to **string** | Formula referencing query outputs (e.g. q1+q2) to compute derived series | [optional] 
**VisibleSeries** | Pointer to **[]bool** | Flags indicating whether each query or formula series is visible | [optional] 
**GroupBy** | Pointer to [**[]AggregationGroupBy**](AggregationGroupBy.md) | Fields used to group the results | [optional] 
**TimeBucket** | Pointer to [**GraphVisualizationTimeseriesTimeBucket**](GraphVisualizationTimeseriesTimeBucket.md) |  | [optional] 
**Normalizer** | Pointer to [**Normalizer**](Normalizer.md) |  | [optional] 

## Methods

### NewGraphVisualizationBar

`func NewGraphVisualizationBar(type_ string, source string, queries []AggregationQuery, ) *GraphVisualizationBar`

NewGraphVisualizationBar instantiates a new GraphVisualizationBar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphVisualizationBarWithDefaults

`func NewGraphVisualizationBarWithDefaults() *GraphVisualizationBar`

NewGraphVisualizationBarWithDefaults instantiates a new GraphVisualizationBar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GraphVisualizationBar) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GraphVisualizationBar) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GraphVisualizationBar) SetType(v string)`

SetType sets Type field to given value.


### GetSource

`func (o *GraphVisualizationBar) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GraphVisualizationBar) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GraphVisualizationBar) SetSource(v string)`

SetSource sets Source field to given value.


### GetQueries

`func (o *GraphVisualizationBar) GetQueries() []AggregationQuery`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *GraphVisualizationBar) GetQueriesOk() (*[]AggregationQuery, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *GraphVisualizationBar) SetQueries(v []AggregationQuery)`

SetQueries sets Queries field to given value.


### GetFormula

`func (o *GraphVisualizationBar) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *GraphVisualizationBar) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *GraphVisualizationBar) SetFormula(v string)`

SetFormula sets Formula field to given value.

### HasFormula

`func (o *GraphVisualizationBar) HasFormula() bool`

HasFormula returns a boolean if a field has been set.

### GetVisibleSeries

`func (o *GraphVisualizationBar) GetVisibleSeries() []bool`

GetVisibleSeries returns the VisibleSeries field if non-nil, zero value otherwise.

### GetVisibleSeriesOk

`func (o *GraphVisualizationBar) GetVisibleSeriesOk() (*[]bool, bool)`

GetVisibleSeriesOk returns a tuple with the VisibleSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleSeries

`func (o *GraphVisualizationBar) SetVisibleSeries(v []bool)`

SetVisibleSeries sets VisibleSeries field to given value.

### HasVisibleSeries

`func (o *GraphVisualizationBar) HasVisibleSeries() bool`

HasVisibleSeries returns a boolean if a field has been set.

### GetGroupBy

`func (o *GraphVisualizationBar) GetGroupBy() []AggregationGroupBy`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *GraphVisualizationBar) GetGroupByOk() (*[]AggregationGroupBy, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *GraphVisualizationBar) SetGroupBy(v []AggregationGroupBy)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *GraphVisualizationBar) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetTimeBucket

`func (o *GraphVisualizationBar) GetTimeBucket() GraphVisualizationTimeseriesTimeBucket`

GetTimeBucket returns the TimeBucket field if non-nil, zero value otherwise.

### GetTimeBucketOk

`func (o *GraphVisualizationBar) GetTimeBucketOk() (*GraphVisualizationTimeseriesTimeBucket, bool)`

GetTimeBucketOk returns a tuple with the TimeBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBucket

`func (o *GraphVisualizationBar) SetTimeBucket(v GraphVisualizationTimeseriesTimeBucket)`

SetTimeBucket sets TimeBucket field to given value.

### HasTimeBucket

`func (o *GraphVisualizationBar) HasTimeBucket() bool`

HasTimeBucket returns a boolean if a field has been set.

### GetNormalizer

`func (o *GraphVisualizationBar) GetNormalizer() Normalizer`

GetNormalizer returns the Normalizer field if non-nil, zero value otherwise.

### GetNormalizerOk

`func (o *GraphVisualizationBar) GetNormalizerOk() (*Normalizer, bool)`

GetNormalizerOk returns a tuple with the Normalizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizer

`func (o *GraphVisualizationBar) SetNormalizer(v Normalizer)`

SetNormalizer sets Normalizer field to given value.

### HasNormalizer

`func (o *GraphVisualizationBar) HasNormalizer() bool`

HasNormalizer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


