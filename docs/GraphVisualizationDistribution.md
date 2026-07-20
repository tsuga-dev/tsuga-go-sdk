# GraphVisualizationDistribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays the aggregation as a distribution chart | 
**Source** | **string** | Telemetry source queried by this aggregation: &#x60;logs&#x60;, &#x60;metrics&#x60;, &#x60;traces&#x60;, or &#x60;rum&#x60;. | 
**Queries** | [**[]AggregationQuery**](AggregationQuery.md) | Aggregations that may be combined together in the same query. Each item is referenced from &#x60;formula&#x60; as q1, q2, and so on, in submission order. | 
**Formula** | Pointer to **string** | Formula referencing submitted query outputs, such as &#x60;q1 + q2&#x60;. References must be within &#x60;q1&#x60; through &#x60;qN&#x60; for the submitted queries. | [optional] 
**Aliases** | Pointer to [**GraphVisualizationTimeseriesPromqlAliases**](GraphVisualizationTimeseriesPromqlAliases.md) |  | [optional] 
**VisibleSeries** | Pointer to **[]bool** | Flags indicating whether each query or formula series is visible | [optional] 
**GroupBy** | Pointer to [**[]AggregationGroupBy**](AggregationGroupBy.md) | Nested grouping levels applied to the results, outermost first (e.g. group by service, then by level within each service). Each level splits results further, so the response contains one result per unique combination of group values. | [optional] 
**Precision** | Pointer to **float32** | Number of decimal places to display in the value | [optional] 
**Normalizer** | Pointer to [**Normalizer**](Normalizer.md) |  | [optional] 
**PercentileMarkers** | Pointer to **[]int32** | Percentile markers displayed on top of the distribution chart | [optional] 
**BoundsScale** | Pointer to **string** |  | [optional] 

## Methods

### NewGraphVisualizationDistribution

`func NewGraphVisualizationDistribution(type_ string, source string, queries []AggregationQuery, ) *GraphVisualizationDistribution`

NewGraphVisualizationDistribution instantiates a new GraphVisualizationDistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphVisualizationDistributionWithDefaults

`func NewGraphVisualizationDistributionWithDefaults() *GraphVisualizationDistribution`

NewGraphVisualizationDistributionWithDefaults instantiates a new GraphVisualizationDistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GraphVisualizationDistribution) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GraphVisualizationDistribution) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GraphVisualizationDistribution) SetType(v string)`

SetType sets Type field to given value.


### GetSource

`func (o *GraphVisualizationDistribution) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GraphVisualizationDistribution) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GraphVisualizationDistribution) SetSource(v string)`

SetSource sets Source field to given value.


### GetQueries

`func (o *GraphVisualizationDistribution) GetQueries() []AggregationQuery`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *GraphVisualizationDistribution) GetQueriesOk() (*[]AggregationQuery, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *GraphVisualizationDistribution) SetQueries(v []AggregationQuery)`

SetQueries sets Queries field to given value.


### GetFormula

`func (o *GraphVisualizationDistribution) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *GraphVisualizationDistribution) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *GraphVisualizationDistribution) SetFormula(v string)`

SetFormula sets Formula field to given value.

### HasFormula

`func (o *GraphVisualizationDistribution) HasFormula() bool`

HasFormula returns a boolean if a field has been set.

### GetAliases

`func (o *GraphVisualizationDistribution) GetAliases() GraphVisualizationTimeseriesPromqlAliases`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *GraphVisualizationDistribution) GetAliasesOk() (*GraphVisualizationTimeseriesPromqlAliases, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *GraphVisualizationDistribution) SetAliases(v GraphVisualizationTimeseriesPromqlAliases)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *GraphVisualizationDistribution) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetVisibleSeries

`func (o *GraphVisualizationDistribution) GetVisibleSeries() []bool`

GetVisibleSeries returns the VisibleSeries field if non-nil, zero value otherwise.

### GetVisibleSeriesOk

`func (o *GraphVisualizationDistribution) GetVisibleSeriesOk() (*[]bool, bool)`

GetVisibleSeriesOk returns a tuple with the VisibleSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleSeries

`func (o *GraphVisualizationDistribution) SetVisibleSeries(v []bool)`

SetVisibleSeries sets VisibleSeries field to given value.

### HasVisibleSeries

`func (o *GraphVisualizationDistribution) HasVisibleSeries() bool`

HasVisibleSeries returns a boolean if a field has been set.

### GetGroupBy

`func (o *GraphVisualizationDistribution) GetGroupBy() []AggregationGroupBy`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *GraphVisualizationDistribution) GetGroupByOk() (*[]AggregationGroupBy, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *GraphVisualizationDistribution) SetGroupBy(v []AggregationGroupBy)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *GraphVisualizationDistribution) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetPrecision

`func (o *GraphVisualizationDistribution) GetPrecision() float32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *GraphVisualizationDistribution) GetPrecisionOk() (*float32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *GraphVisualizationDistribution) SetPrecision(v float32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *GraphVisualizationDistribution) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetNormalizer

`func (o *GraphVisualizationDistribution) GetNormalizer() Normalizer`

GetNormalizer returns the Normalizer field if non-nil, zero value otherwise.

### GetNormalizerOk

`func (o *GraphVisualizationDistribution) GetNormalizerOk() (*Normalizer, bool)`

GetNormalizerOk returns a tuple with the Normalizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizer

`func (o *GraphVisualizationDistribution) SetNormalizer(v Normalizer)`

SetNormalizer sets Normalizer field to given value.

### HasNormalizer

`func (o *GraphVisualizationDistribution) HasNormalizer() bool`

HasNormalizer returns a boolean if a field has been set.

### GetPercentileMarkers

`func (o *GraphVisualizationDistribution) GetPercentileMarkers() []int32`

GetPercentileMarkers returns the PercentileMarkers field if non-nil, zero value otherwise.

### GetPercentileMarkersOk

`func (o *GraphVisualizationDistribution) GetPercentileMarkersOk() (*[]int32, bool)`

GetPercentileMarkersOk returns a tuple with the PercentileMarkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentileMarkers

`func (o *GraphVisualizationDistribution) SetPercentileMarkers(v []int32)`

SetPercentileMarkers sets PercentileMarkers field to given value.

### HasPercentileMarkers

`func (o *GraphVisualizationDistribution) HasPercentileMarkers() bool`

HasPercentileMarkers returns a boolean if a field has been set.

### GetBoundsScale

`func (o *GraphVisualizationDistribution) GetBoundsScale() string`

GetBoundsScale returns the BoundsScale field if non-nil, zero value otherwise.

### GetBoundsScaleOk

`func (o *GraphVisualizationDistribution) GetBoundsScaleOk() (*string, bool)`

GetBoundsScaleOk returns a tuple with the BoundsScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundsScale

`func (o *GraphVisualizationDistribution) SetBoundsScale(v string)`

SetBoundsScale sets BoundsScale field to given value.

### HasBoundsScale

`func (o *GraphVisualizationDistribution) HasBoundsScale() bool`

HasBoundsScale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


