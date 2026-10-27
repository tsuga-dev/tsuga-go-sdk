# Function

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Normalizes the metric to events per second | 
**Base** | **int32** | The base of the logarithm | 
**Exponent** | **int32** | The exponent to raise values to | 
**Seconds** | **int32** | Number of seconds to offset | 

## Methods

### NewFunction

`func NewFunction(type_ string, base int32, exponent int32, seconds int32, ) *Function`

NewFunction instantiates a new Function object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionWithDefaults

`func NewFunctionWithDefaults() *Function`

NewFunctionWithDefaults instantiates a new Function object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Function) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Function) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Function) SetType(v string)`

SetType sets Type field to given value.


### GetBase

`func (o *Function) GetBase() int32`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *Function) GetBaseOk() (*int32, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *Function) SetBase(v int32)`

SetBase sets Base field to given value.


### GetExponent

`func (o *Function) GetExponent() int32`

GetExponent returns the Exponent field if non-nil, zero value otherwise.

### GetExponentOk

`func (o *Function) GetExponentOk() (*int32, bool)`

GetExponentOk returns a tuple with the Exponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExponent

`func (o *Function) SetExponent(v int32)`

SetExponent sets Exponent field to given value.


### GetSeconds

`func (o *Function) GetSeconds() int32`

GetSeconds returns the Seconds field if non-nil, zero value otherwise.

### GetSecondsOk

`func (o *Function) GetSecondsOk() (*int32, bool)`

GetSecondsOk returns a tuple with the Seconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeconds

`func (o *Function) SetSeconds(v int32)`

SetSeconds sets Seconds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


