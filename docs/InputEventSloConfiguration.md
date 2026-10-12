# InputEventSloConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Configuration kind. Set to &#x60;event&#x60; for SLOs that compute SLI from good and total query formulas over the rolling timeframe. | 
**DataSource** | **string** | Telemetry source queried by this aggregation: &#x60;logs&#x60;, &#x60;metrics&#x60;, &#x60;traces&#x60;, or &#x60;rum&#x60;. | 
**GoodQuery** | [**InputSloQueryFormula**](InputSloQueryFormula.md) |  | 
**TotalQuery** | [**InputSloQueryFormula**](InputSloQueryFormula.md) |  | 
**GroupByFields** | Pointer to [**[]InputAggregationGroupBy**](InputAggregationGroupBy.md) | Nested grouping levels applied to aggregation results, outermost first (e.g. group by service, then by level within each service). Each level splits results further, so the response contains one result per unique combination of group values instead of one aggregated total. Defaults to an empty array (one ungrouped result) when omitted. | [optional] 
**NoDataBehavior** | **string** | How to treat SLO with no data: good (count as meeting the SLO), bad (count as breaching it) | 

## Methods

### NewInputEventSloConfiguration

`func NewInputEventSloConfiguration(type_ string, dataSource string, goodQuery InputSloQueryFormula, totalQuery InputSloQueryFormula, noDataBehavior string, ) *InputEventSloConfiguration`

NewInputEventSloConfiguration instantiates a new InputEventSloConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputEventSloConfigurationWithDefaults

`func NewInputEventSloConfigurationWithDefaults() *InputEventSloConfiguration`

NewInputEventSloConfigurationWithDefaults instantiates a new InputEventSloConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputEventSloConfiguration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputEventSloConfiguration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputEventSloConfiguration) SetType(v string)`

SetType sets Type field to given value.


### GetDataSource

`func (o *InputEventSloConfiguration) GetDataSource() string`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *InputEventSloConfiguration) GetDataSourceOk() (*string, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *InputEventSloConfiguration) SetDataSource(v string)`

SetDataSource sets DataSource field to given value.


### GetGoodQuery

`func (o *InputEventSloConfiguration) GetGoodQuery() InputSloQueryFormula`

GetGoodQuery returns the GoodQuery field if non-nil, zero value otherwise.

### GetGoodQueryOk

`func (o *InputEventSloConfiguration) GetGoodQueryOk() (*InputSloQueryFormula, bool)`

GetGoodQueryOk returns a tuple with the GoodQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodQuery

`func (o *InputEventSloConfiguration) SetGoodQuery(v InputSloQueryFormula)`

SetGoodQuery sets GoodQuery field to given value.


### GetTotalQuery

`func (o *InputEventSloConfiguration) GetTotalQuery() InputSloQueryFormula`

GetTotalQuery returns the TotalQuery field if non-nil, zero value otherwise.

### GetTotalQueryOk

`func (o *InputEventSloConfiguration) GetTotalQueryOk() (*InputSloQueryFormula, bool)`

GetTotalQueryOk returns a tuple with the TotalQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalQuery

`func (o *InputEventSloConfiguration) SetTotalQuery(v InputSloQueryFormula)`

SetTotalQuery sets TotalQuery field to given value.


### GetGroupByFields

`func (o *InputEventSloConfiguration) GetGroupByFields() []InputAggregationGroupBy`

GetGroupByFields returns the GroupByFields field if non-nil, zero value otherwise.

### GetGroupByFieldsOk

`func (o *InputEventSloConfiguration) GetGroupByFieldsOk() (*[]InputAggregationGroupBy, bool)`

GetGroupByFieldsOk returns a tuple with the GroupByFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByFields

`func (o *InputEventSloConfiguration) SetGroupByFields(v []InputAggregationGroupBy)`

SetGroupByFields sets GroupByFields field to given value.

### HasGroupByFields

`func (o *InputEventSloConfiguration) HasGroupByFields() bool`

HasGroupByFields returns a boolean if a field has been set.

### GetNoDataBehavior

`func (o *InputEventSloConfiguration) GetNoDataBehavior() string`

GetNoDataBehavior returns the NoDataBehavior field if non-nil, zero value otherwise.

### GetNoDataBehaviorOk

`func (o *InputEventSloConfiguration) GetNoDataBehaviorOk() (*string, bool)`

GetNoDataBehaviorOk returns a tuple with the NoDataBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDataBehavior

`func (o *InputEventSloConfiguration) SetNoDataBehavior(v string)`

SetNoDataBehavior sets NoDataBehavior field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


