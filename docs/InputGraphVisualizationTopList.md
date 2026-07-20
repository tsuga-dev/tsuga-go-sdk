# InputGraphVisualizationTopList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays the aggregation as a ranked list of top results | 
**Source** | **string** | Telemetry source queried by this aggregation: &#x60;logs&#x60;, &#x60;metrics&#x60;, &#x60;traces&#x60;, or &#x60;rum&#x60;. | 
**Queries** | [**[]AggregationQuery1**](AggregationQuery1.md) | Aggregations that may be combined together in the same query. Each item is referenced from &#x60;formula&#x60; as q1, q2, and so on, in submission order. Limited to 15 items. For dataSource \&quot;metrics\&quot;, each aggregate&#39;s &#x60;field&#x60; is the metric name, not an attribute; to count distinct values of an attribute use unique-count with field \&quot;&lt;metricName&gt;.context.&lt;attribute&gt;\&quot; (e.g. \&quot;system.cpu.utilization.context.host.name\&quot;). | 
**Formula** | Pointer to **string** | Formula referencing query outputs, such as &#x60;q1 + q2&#x60;, to compute derived results. Defaults to &#x60;q1&#x60;. Formulas may reference only submitted queries (&#x60;q1&#x60; through &#x60;qN&#x60;); undefined query references return 400. | [optional] 
**Aliases** | Pointer to [**InputGraphVisualizationTimeseriesPromqlAliases**](InputGraphVisualizationTimeseriesPromqlAliases.md) |  | [optional] 
**VisibleSeries** | Pointer to **[]bool** | Flags indicating whether each query or formula series is visible | [optional] 
**GroupBy** | Pointer to [**[]AggregationGroupBy1**](AggregationGroupBy1.md) | Nested grouping levels applied to aggregation results, outermost first (e.g. group by service, then by level within each service). Each level splits results further, so the response contains one result per unique combination of group values instead of one aggregated total. Defaults to an empty array (one ungrouped result) when omitted. Limited to 7 levels. | [optional] 
**Normalizer** | Pointer to [**Normalizer1**](Normalizer1.md) |  | [optional] 
**Precision** | Pointer to **float32** | Number of decimal places to display in the value | [optional] 
**Conditions** | Pointer to [**[]ConditionalFormatting**](ConditionalFormatting.md) | Conditional formatting rules applied to the displayed value | [optional] 

## Methods

### NewInputGraphVisualizationTopList

`func NewInputGraphVisualizationTopList(type_ string, source string, queries []AggregationQuery1, ) *InputGraphVisualizationTopList`

NewInputGraphVisualizationTopList instantiates a new InputGraphVisualizationTopList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputGraphVisualizationTopListWithDefaults

`func NewInputGraphVisualizationTopListWithDefaults() *InputGraphVisualizationTopList`

NewInputGraphVisualizationTopListWithDefaults instantiates a new InputGraphVisualizationTopList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputGraphVisualizationTopList) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputGraphVisualizationTopList) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputGraphVisualizationTopList) SetType(v string)`

SetType sets Type field to given value.


### GetSource

`func (o *InputGraphVisualizationTopList) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InputGraphVisualizationTopList) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InputGraphVisualizationTopList) SetSource(v string)`

SetSource sets Source field to given value.


### GetQueries

`func (o *InputGraphVisualizationTopList) GetQueries() []AggregationQuery1`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *InputGraphVisualizationTopList) GetQueriesOk() (*[]AggregationQuery1, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *InputGraphVisualizationTopList) SetQueries(v []AggregationQuery1)`

SetQueries sets Queries field to given value.


### GetFormula

`func (o *InputGraphVisualizationTopList) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *InputGraphVisualizationTopList) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *InputGraphVisualizationTopList) SetFormula(v string)`

SetFormula sets Formula field to given value.

### HasFormula

`func (o *InputGraphVisualizationTopList) HasFormula() bool`

HasFormula returns a boolean if a field has been set.

### GetAliases

`func (o *InputGraphVisualizationTopList) GetAliases() InputGraphVisualizationTimeseriesPromqlAliases`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *InputGraphVisualizationTopList) GetAliasesOk() (*InputGraphVisualizationTimeseriesPromqlAliases, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *InputGraphVisualizationTopList) SetAliases(v InputGraphVisualizationTimeseriesPromqlAliases)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *InputGraphVisualizationTopList) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetVisibleSeries

`func (o *InputGraphVisualizationTopList) GetVisibleSeries() []bool`

GetVisibleSeries returns the VisibleSeries field if non-nil, zero value otherwise.

### GetVisibleSeriesOk

`func (o *InputGraphVisualizationTopList) GetVisibleSeriesOk() (*[]bool, bool)`

GetVisibleSeriesOk returns a tuple with the VisibleSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleSeries

`func (o *InputGraphVisualizationTopList) SetVisibleSeries(v []bool)`

SetVisibleSeries sets VisibleSeries field to given value.

### HasVisibleSeries

`func (o *InputGraphVisualizationTopList) HasVisibleSeries() bool`

HasVisibleSeries returns a boolean if a field has been set.

### GetGroupBy

`func (o *InputGraphVisualizationTopList) GetGroupBy() []AggregationGroupBy1`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *InputGraphVisualizationTopList) GetGroupByOk() (*[]AggregationGroupBy1, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *InputGraphVisualizationTopList) SetGroupBy(v []AggregationGroupBy1)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *InputGraphVisualizationTopList) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetNormalizer

`func (o *InputGraphVisualizationTopList) GetNormalizer() Normalizer1`

GetNormalizer returns the Normalizer field if non-nil, zero value otherwise.

### GetNormalizerOk

`func (o *InputGraphVisualizationTopList) GetNormalizerOk() (*Normalizer1, bool)`

GetNormalizerOk returns a tuple with the Normalizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizer

`func (o *InputGraphVisualizationTopList) SetNormalizer(v Normalizer1)`

SetNormalizer sets Normalizer field to given value.

### HasNormalizer

`func (o *InputGraphVisualizationTopList) HasNormalizer() bool`

HasNormalizer returns a boolean if a field has been set.

### GetPrecision

`func (o *InputGraphVisualizationTopList) GetPrecision() float32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *InputGraphVisualizationTopList) GetPrecisionOk() (*float32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *InputGraphVisualizationTopList) SetPrecision(v float32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *InputGraphVisualizationTopList) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetConditions

`func (o *InputGraphVisualizationTopList) GetConditions() []ConditionalFormatting`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *InputGraphVisualizationTopList) GetConditionsOk() (*[]ConditionalFormatting, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *InputGraphVisualizationTopList) SetConditions(v []ConditionalFormatting)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *InputGraphVisualizationTopList) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


