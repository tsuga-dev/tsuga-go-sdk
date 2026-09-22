# SloQueryFormula1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Queries** | [**[]MonitorAggregationQuery**](MonitorAggregationQuery.md) | Aggregation queries used by alerting and SLO evaluation. Each query is referenced from formulas as q1, q2, and so on. | 
**Formula** | **string** | Formula referencing submitted query outputs, such as &#x60;q1 + q2&#x60;. References must be within &#x60;q1&#x60; through &#x60;qN&#x60; for the submitted queries. | 

## Methods

### NewSloQueryFormula1

`func NewSloQueryFormula1(queries []MonitorAggregationQuery, formula string, ) *SloQueryFormula1`

NewSloQueryFormula1 instantiates a new SloQueryFormula1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSloQueryFormula1WithDefaults

`func NewSloQueryFormula1WithDefaults() *SloQueryFormula1`

NewSloQueryFormula1WithDefaults instantiates a new SloQueryFormula1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueries

`func (o *SloQueryFormula1) GetQueries() []MonitorAggregationQuery`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *SloQueryFormula1) GetQueriesOk() (*[]MonitorAggregationQuery, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *SloQueryFormula1) SetQueries(v []MonitorAggregationQuery)`

SetQueries sets Queries field to given value.


### GetFormula

`func (o *SloQueryFormula1) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *SloQueryFormula1) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *SloQueryFormula1) SetFormula(v string)`

SetFormula sets Formula field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


