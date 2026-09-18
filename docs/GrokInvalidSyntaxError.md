# GrokInvalidSyntaxError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | **int32** | Byte offset in the rule where parsing failed. | 
**Detail** | **string** | Parser detail message. | 

## Methods

### NewGrokInvalidSyntaxError

`func NewGrokInvalidSyntaxError(offset int32, detail string, ) *GrokInvalidSyntaxError`

NewGrokInvalidSyntaxError instantiates a new GrokInvalidSyntaxError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrokInvalidSyntaxErrorWithDefaults

`func NewGrokInvalidSyntaxErrorWithDefaults() *GrokInvalidSyntaxError`

NewGrokInvalidSyntaxErrorWithDefaults instantiates a new GrokInvalidSyntaxError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *GrokInvalidSyntaxError) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GrokInvalidSyntaxError) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GrokInvalidSyntaxError) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetDetail

`func (o *GrokInvalidSyntaxError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *GrokInvalidSyntaxError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *GrokInvalidSyntaxError) SetDetail(v string)`

SetDetail sets Detail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


