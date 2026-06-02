# InputGraphVisualizationDistribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays the aggregation as a distribution chart | 
**Source** | **string** | Data source being queried for this aggregation | 
**Queries** | [**[]AggregationQuery1**](AggregationQuery1.md) | Aggregations that may be combined together in the same query | 
**Formula** | Pointer to **string** | Formula referencing query outputs (e.g. q1+q2) to compute derived series | [optional] 
**Aliases** | Pointer to [**InputGraphVisualizationTimeseriesAliases**](InputGraphVisualizationTimeseriesAliases.md) |  | [optional] 
**VisibleSeries** | Pointer to **[]bool** | Flags indicating whether each query or formula series is visible | [optional] 
**GroupBy** | Pointer to [**[]AggregationGroupBy1**](AggregationGroupBy1.md) | Fields used to group the results | [optional] 
**Precision** | Pointer to **float32** | Number of decimal places to display in the value | [optional] 
**Normalizer** | Pointer to [**Normalizer1**](Normalizer1.md) |  | [optional] 
**PercentileMarkers** | Pointer to **[]int32** | Percentile markers displayed on top of the distribution chart | [optional] 

## Methods

### NewInputGraphVisualizationDistribution

`func NewInputGraphVisualizationDistribution(type_ string, source string, queries []AggregationQuery1, ) *InputGraphVisualizationDistribution`

NewInputGraphVisualizationDistribution instantiates a new InputGraphVisualizationDistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputGraphVisualizationDistributionWithDefaults

`func NewInputGraphVisualizationDistributionWithDefaults() *InputGraphVisualizationDistribution`

NewInputGraphVisualizationDistributionWithDefaults instantiates a new InputGraphVisualizationDistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputGraphVisualizationDistribution) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputGraphVisualizationDistribution) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputGraphVisualizationDistribution) SetType(v string)`

SetType sets Type field to given value.


### GetSource

`func (o *InputGraphVisualizationDistribution) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InputGraphVisualizationDistribution) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InputGraphVisualizationDistribution) SetSource(v string)`

SetSource sets Source field to given value.


### GetQueries

`func (o *InputGraphVisualizationDistribution) GetQueries() []AggregationQuery1`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *InputGraphVisualizationDistribution) GetQueriesOk() (*[]AggregationQuery1, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *InputGraphVisualizationDistribution) SetQueries(v []AggregationQuery1)`

SetQueries sets Queries field to given value.


### GetFormula

`func (o *InputGraphVisualizationDistribution) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *InputGraphVisualizationDistribution) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *InputGraphVisualizationDistribution) SetFormula(v string)`

SetFormula sets Formula field to given value.

### HasFormula

`func (o *InputGraphVisualizationDistribution) HasFormula() bool`

HasFormula returns a boolean if a field has been set.

### GetAliases

`func (o *InputGraphVisualizationDistribution) GetAliases() InputGraphVisualizationTimeseriesAliases`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *InputGraphVisualizationDistribution) GetAliasesOk() (*InputGraphVisualizationTimeseriesAliases, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *InputGraphVisualizationDistribution) SetAliases(v InputGraphVisualizationTimeseriesAliases)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *InputGraphVisualizationDistribution) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetVisibleSeries

`func (o *InputGraphVisualizationDistribution) GetVisibleSeries() []bool`

GetVisibleSeries returns the VisibleSeries field if non-nil, zero value otherwise.

### GetVisibleSeriesOk

`func (o *InputGraphVisualizationDistribution) GetVisibleSeriesOk() (*[]bool, bool)`

GetVisibleSeriesOk returns a tuple with the VisibleSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleSeries

`func (o *InputGraphVisualizationDistribution) SetVisibleSeries(v []bool)`

SetVisibleSeries sets VisibleSeries field to given value.

### HasVisibleSeries

`func (o *InputGraphVisualizationDistribution) HasVisibleSeries() bool`

HasVisibleSeries returns a boolean if a field has been set.

### GetGroupBy

`func (o *InputGraphVisualizationDistribution) GetGroupBy() []AggregationGroupBy1`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *InputGraphVisualizationDistribution) GetGroupByOk() (*[]AggregationGroupBy1, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *InputGraphVisualizationDistribution) SetGroupBy(v []AggregationGroupBy1)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *InputGraphVisualizationDistribution) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetPrecision

`func (o *InputGraphVisualizationDistribution) GetPrecision() float32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *InputGraphVisualizationDistribution) GetPrecisionOk() (*float32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *InputGraphVisualizationDistribution) SetPrecision(v float32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *InputGraphVisualizationDistribution) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetNormalizer

`func (o *InputGraphVisualizationDistribution) GetNormalizer() Normalizer1`

GetNormalizer returns the Normalizer field if non-nil, zero value otherwise.

### GetNormalizerOk

`func (o *InputGraphVisualizationDistribution) GetNormalizerOk() (*Normalizer1, bool)`

GetNormalizerOk returns a tuple with the Normalizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizer

`func (o *InputGraphVisualizationDistribution) SetNormalizer(v Normalizer1)`

SetNormalizer sets Normalizer field to given value.

### HasNormalizer

`func (o *InputGraphVisualizationDistribution) HasNormalizer() bool`

HasNormalizer returns a boolean if a field has been set.

### GetPercentileMarkers

`func (o *InputGraphVisualizationDistribution) GetPercentileMarkers() []int32`

GetPercentileMarkers returns the PercentileMarkers field if non-nil, zero value otherwise.

### GetPercentileMarkersOk

`func (o *InputGraphVisualizationDistribution) GetPercentileMarkersOk() (*[]int32, bool)`

GetPercentileMarkersOk returns a tuple with the PercentileMarkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentileMarkers

`func (o *InputGraphVisualizationDistribution) SetPercentileMarkers(v []int32)`

SetPercentileMarkers sets PercentileMarkers field to given value.

### HasPercentileMarkers

`func (o *InputGraphVisualizationDistribution) HasPercentileMarkers() bool`

HasPercentileMarkers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


