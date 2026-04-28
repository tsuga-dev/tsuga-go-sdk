# FunctionTimeOffset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Shifts a query&#39;s data backward by a specified duration for easy comparison over different periods of time. | 
**Seconds** | **int32** | Number of seconds to offset | 

## Methods

### NewFunctionTimeOffset

`func NewFunctionTimeOffset(type_ string, seconds int32, ) *FunctionTimeOffset`

NewFunctionTimeOffset instantiates a new FunctionTimeOffset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionTimeOffsetWithDefaults

`func NewFunctionTimeOffsetWithDefaults() *FunctionTimeOffset`

NewFunctionTimeOffsetWithDefaults instantiates a new FunctionTimeOffset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FunctionTimeOffset) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FunctionTimeOffset) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FunctionTimeOffset) SetType(v string)`

SetType sets Type field to given value.


### GetSeconds

`func (o *FunctionTimeOffset) GetSeconds() int32`

GetSeconds returns the Seconds field if non-nil, zero value otherwise.

### GetSecondsOk

`func (o *FunctionTimeOffset) GetSecondsOk() (*int32, bool)`

GetSecondsOk returns a tuple with the Seconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeconds

`func (o *FunctionTimeOffset) SetSeconds(v int32)`

SetSeconds sets Seconds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


