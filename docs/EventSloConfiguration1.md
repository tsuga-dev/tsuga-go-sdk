# EventSloConfiguration1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Configuration kind. Set to &#x60;event&#x60; for SLOs that compute SLI from good and total query formulas over the rolling timeframe. | 
**DataSource** | **string** | Telemetry source queried by this aggregation: &#x60;logs&#x60;, &#x60;metrics&#x60;, &#x60;traces&#x60;, or &#x60;rum&#x60;. | 
**GoodQuery** | [**SloQueryFormula1**](SloQueryFormula1.md) |  | 
**TotalQuery** | [**SloQueryFormula1**](SloQueryFormula1.md) |  | 
**GroupByFields** | Pointer to [**[]AggregationGroupBy1**](AggregationGroupBy1.md) | Nested grouping levels applied to the results, outermost first (e.g. group by service, then by level within each service). Each level splits results further, so the response contains one result per unique combination of group values. | [optional] 
**NoDataBehavior** | **string** | How to treat SLO with no data: good (count as meeting the SLO), bad (count as breaching it) | 

## Methods

### NewEventSloConfiguration1

`func NewEventSloConfiguration1(type_ string, dataSource string, goodQuery SloQueryFormula1, totalQuery SloQueryFormula1, noDataBehavior string, ) *EventSloConfiguration1`

NewEventSloConfiguration1 instantiates a new EventSloConfiguration1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventSloConfiguration1WithDefaults

`func NewEventSloConfiguration1WithDefaults() *EventSloConfiguration1`

NewEventSloConfiguration1WithDefaults instantiates a new EventSloConfiguration1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EventSloConfiguration1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventSloConfiguration1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventSloConfiguration1) SetType(v string)`

SetType sets Type field to given value.


### GetDataSource

`func (o *EventSloConfiguration1) GetDataSource() string`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *EventSloConfiguration1) GetDataSourceOk() (*string, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *EventSloConfiguration1) SetDataSource(v string)`

SetDataSource sets DataSource field to given value.


### GetGoodQuery

`func (o *EventSloConfiguration1) GetGoodQuery() SloQueryFormula1`

GetGoodQuery returns the GoodQuery field if non-nil, zero value otherwise.

### GetGoodQueryOk

`func (o *EventSloConfiguration1) GetGoodQueryOk() (*SloQueryFormula1, bool)`

GetGoodQueryOk returns a tuple with the GoodQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodQuery

`func (o *EventSloConfiguration1) SetGoodQuery(v SloQueryFormula1)`

SetGoodQuery sets GoodQuery field to given value.


### GetTotalQuery

`func (o *EventSloConfiguration1) GetTotalQuery() SloQueryFormula1`

GetTotalQuery returns the TotalQuery field if non-nil, zero value otherwise.

### GetTotalQueryOk

`func (o *EventSloConfiguration1) GetTotalQueryOk() (*SloQueryFormula1, bool)`

GetTotalQueryOk returns a tuple with the TotalQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalQuery

`func (o *EventSloConfiguration1) SetTotalQuery(v SloQueryFormula1)`

SetTotalQuery sets TotalQuery field to given value.


### GetGroupByFields

`func (o *EventSloConfiguration1) GetGroupByFields() []AggregationGroupBy1`

GetGroupByFields returns the GroupByFields field if non-nil, zero value otherwise.

### GetGroupByFieldsOk

`func (o *EventSloConfiguration1) GetGroupByFieldsOk() (*[]AggregationGroupBy1, bool)`

GetGroupByFieldsOk returns a tuple with the GroupByFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByFields

`func (o *EventSloConfiguration1) SetGroupByFields(v []AggregationGroupBy1)`

SetGroupByFields sets GroupByFields field to given value.

### HasGroupByFields

`func (o *EventSloConfiguration1) HasGroupByFields() bool`

HasGroupByFields returns a boolean if a field has been set.

### GetNoDataBehavior

`func (o *EventSloConfiguration1) GetNoDataBehavior() string`

GetNoDataBehavior returns the NoDataBehavior field if non-nil, zero value otherwise.

### GetNoDataBehaviorOk

`func (o *EventSloConfiguration1) GetNoDataBehaviorOk() (*string, bool)`

GetNoDataBehaviorOk returns a tuple with the NoDataBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDataBehavior

`func (o *EventSloConfiguration1) SetNoDataBehavior(v string)`

SetNoDataBehavior sets NoDataBehavior field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


