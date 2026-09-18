# GrokUnknownFilterError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | **int32** | Byte offset of the unknown filter. | 
**UnknownFilter** | **string** | The filter name that is not implemented. | 
**Message** | **string** | Details about the unsupported filter. | 

## Methods

### NewGrokUnknownFilterError

`func NewGrokUnknownFilterError(offset int32, unknownFilter string, message string, ) *GrokUnknownFilterError`

NewGrokUnknownFilterError instantiates a new GrokUnknownFilterError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrokUnknownFilterErrorWithDefaults

`func NewGrokUnknownFilterErrorWithDefaults() *GrokUnknownFilterError`

NewGrokUnknownFilterErrorWithDefaults instantiates a new GrokUnknownFilterError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *GrokUnknownFilterError) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GrokUnknownFilterError) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GrokUnknownFilterError) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetUnknownFilter

`func (o *GrokUnknownFilterError) GetUnknownFilter() string`

GetUnknownFilter returns the UnknownFilter field if non-nil, zero value otherwise.

### GetUnknownFilterOk

`func (o *GrokUnknownFilterError) GetUnknownFilterOk() (*string, bool)`

GetUnknownFilterOk returns a tuple with the UnknownFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFilter

`func (o *GrokUnknownFilterError) SetUnknownFilter(v string)`

SetUnknownFilter sets UnknownFilter field to given value.


### GetMessage

`func (o *GrokUnknownFilterError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GrokUnknownFilterError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GrokUnknownFilterError) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


