# InputFunctionTimeOffset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Shifts a query&#39;s data backward by a specified duration for easy comparison over different periods of time. | 
**Seconds** | **int32** | Number of seconds to offset | 

## Methods

### NewInputFunctionTimeOffset

`func NewInputFunctionTimeOffset(type_ string, seconds int32, ) *InputFunctionTimeOffset`

NewInputFunctionTimeOffset instantiates a new InputFunctionTimeOffset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputFunctionTimeOffsetWithDefaults

`func NewInputFunctionTimeOffsetWithDefaults() *InputFunctionTimeOffset`

NewInputFunctionTimeOffsetWithDefaults instantiates a new InputFunctionTimeOffset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputFunctionTimeOffset) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputFunctionTimeOffset) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputFunctionTimeOffset) SetType(v string)`

SetType sets Type field to given value.


### GetSeconds

`func (o *InputFunctionTimeOffset) GetSeconds() int32`

GetSeconds returns the Seconds field if non-nil, zero value otherwise.

### GetSecondsOk

`func (o *InputFunctionTimeOffset) GetSecondsOk() (*int32, bool)`

GetSecondsOk returns a tuple with the Seconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeconds

`func (o *InputFunctionTimeOffset) SetSeconds(v int32)`

SetSeconds sets Seconds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


