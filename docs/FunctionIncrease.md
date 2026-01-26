# FunctionIncrease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | &#x60;Increase&#x60; is designed for cumulative metrics. It computes the increase of cumulative metrics automatically handling counter resets. When used with an aggregation like &#x60;sum&#x60; it combines all series without requiring an explicit group-by. | 

## Methods

### NewFunctionIncrease

`func NewFunctionIncrease(type_ string, ) *FunctionIncrease`

NewFunctionIncrease instantiates a new FunctionIncrease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionIncreaseWithDefaults

`func NewFunctionIncreaseWithDefaults() *FunctionIncrease`

NewFunctionIncreaseWithDefaults instantiates a new FunctionIncrease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FunctionIncrease) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FunctionIncrease) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FunctionIncrease) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


