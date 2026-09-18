# GrokUnknownPatternError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | **int32** | Byte offset of the unknown pattern. | 
**UnknownPattern** | **string** | The pattern name that is not supported. | 

## Methods

### NewGrokUnknownPatternError

`func NewGrokUnknownPatternError(offset int32, unknownPattern string, ) *GrokUnknownPatternError`

NewGrokUnknownPatternError instantiates a new GrokUnknownPatternError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrokUnknownPatternErrorWithDefaults

`func NewGrokUnknownPatternErrorWithDefaults() *GrokUnknownPatternError`

NewGrokUnknownPatternErrorWithDefaults instantiates a new GrokUnknownPatternError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *GrokUnknownPatternError) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GrokUnknownPatternError) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GrokUnknownPatternError) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetUnknownPattern

`func (o *GrokUnknownPatternError) GetUnknownPattern() string`

GetUnknownPattern returns the UnknownPattern field if non-nil, zero value otherwise.

### GetUnknownPatternOk

`func (o *GrokUnknownPatternError) GetUnknownPatternOk() (*string, bool)`

GetUnknownPatternOk returns a tuple with the UnknownPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownPattern

`func (o *GrokUnknownPatternError) SetUnknownPattern(v string)`

SetUnknownPattern sets UnknownPattern field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


