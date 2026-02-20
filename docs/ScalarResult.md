# ScalarResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Query identifier (q1, q2, ...) or \&quot;formula\&quot; for formula | 
**Group** | [**map[string]ScalarResultGroupValue**](ScalarResultGroupValue.md) | Group-by field values | 
**Value** | **float32** | Aggregated value | 

## Methods

### NewScalarResult

`func NewScalarResult(id string, group map[string]ScalarResultGroupValue, value float32, ) *ScalarResult`

NewScalarResult instantiates a new ScalarResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScalarResultWithDefaults

`func NewScalarResultWithDefaults() *ScalarResult`

NewScalarResultWithDefaults instantiates a new ScalarResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScalarResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScalarResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScalarResult) SetId(v string)`

SetId sets Id field to given value.


### GetGroup

`func (o *ScalarResult) GetGroup() map[string]ScalarResultGroupValue`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ScalarResult) GetGroupOk() (*map[string]ScalarResultGroupValue, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ScalarResult) SetGroup(v map[string]ScalarResultGroupValue)`

SetGroup sets Group field to given value.


### GetValue

`func (o *ScalarResult) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ScalarResult) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ScalarResult) SetValue(v float32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


