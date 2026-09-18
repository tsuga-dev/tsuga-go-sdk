# GrokInvalidFilterError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | **int32** | Byte offset of the invalid filter. | 
**InvalidFilter** | **string** | The filter that is invalid. | 
**Message** | **string** | Why the filter is invalid. | 

## Methods

### NewGrokInvalidFilterError

`func NewGrokInvalidFilterError(offset int32, invalidFilter string, message string, ) *GrokInvalidFilterError`

NewGrokInvalidFilterError instantiates a new GrokInvalidFilterError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrokInvalidFilterErrorWithDefaults

`func NewGrokInvalidFilterErrorWithDefaults() *GrokInvalidFilterError`

NewGrokInvalidFilterErrorWithDefaults instantiates a new GrokInvalidFilterError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *GrokInvalidFilterError) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GrokInvalidFilterError) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GrokInvalidFilterError) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetInvalidFilter

`func (o *GrokInvalidFilterError) GetInvalidFilter() string`

GetInvalidFilter returns the InvalidFilter field if non-nil, zero value otherwise.

### GetInvalidFilterOk

`func (o *GrokInvalidFilterError) GetInvalidFilterOk() (*string, bool)`

GetInvalidFilterOk returns a tuple with the InvalidFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidFilter

`func (o *GrokInvalidFilterError) SetInvalidFilter(v string)`

SetInvalidFilter sets InvalidFilter field to given value.


### GetMessage

`func (o *GrokInvalidFilterError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GrokInvalidFilterError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GrokInvalidFilterError) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


