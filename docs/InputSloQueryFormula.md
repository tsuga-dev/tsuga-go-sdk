# InputSloQueryFormula

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Queries** | [**[]InputSloAggregationQuery**](InputSloAggregationQuery.md) | Aggregation queries used by SLO evaluation. Each query is referenced from formulas as q1, q2, and so on. | 
**Formula** | **string** | Formula referencing query outputs, such as &#x60;q1 + q2&#x60;, to compute derived results. Defaults to &#x60;q1&#x60;. Formulas may reference only submitted queries (&#x60;q1&#x60; through &#x60;qN&#x60;); undefined query references return 400. | 

## Methods

### NewInputSloQueryFormula

`func NewInputSloQueryFormula(queries []InputSloAggregationQuery, formula string, ) *InputSloQueryFormula`

NewInputSloQueryFormula instantiates a new InputSloQueryFormula object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputSloQueryFormulaWithDefaults

`func NewInputSloQueryFormulaWithDefaults() *InputSloQueryFormula`

NewInputSloQueryFormulaWithDefaults instantiates a new InputSloQueryFormula object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueries

`func (o *InputSloQueryFormula) GetQueries() []InputSloAggregationQuery`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *InputSloQueryFormula) GetQueriesOk() (*[]InputSloAggregationQuery, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *InputSloQueryFormula) SetQueries(v []InputSloAggregationQuery)`

SetQueries sets Queries field to given value.


### GetFormula

`func (o *InputSloQueryFormula) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *InputSloQueryFormula) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *InputSloQueryFormula) SetFormula(v string)`

SetFormula sets Formula field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


