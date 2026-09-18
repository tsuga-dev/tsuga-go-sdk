# GrokUnknownPattern

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator identifying a rule that references a pattern name the grok engine does not provide. | 
**Error** | [**GrokUnknownPatternError**](GrokUnknownPatternError.md) |  | 

## Methods

### NewGrokUnknownPattern

`func NewGrokUnknownPattern(type_ string, error_ GrokUnknownPatternError, ) *GrokUnknownPattern`

NewGrokUnknownPattern instantiates a new GrokUnknownPattern object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrokUnknownPatternWithDefaults

`func NewGrokUnknownPatternWithDefaults() *GrokUnknownPattern`

NewGrokUnknownPatternWithDefaults instantiates a new GrokUnknownPattern object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GrokUnknownPattern) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GrokUnknownPattern) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GrokUnknownPattern) SetType(v string)`

SetType sets Type field to given value.


### GetError

`func (o *GrokUnknownPattern) GetError() GrokUnknownPatternError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GrokUnknownPattern) GetErrorOk() (*GrokUnknownPatternError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GrokUnknownPattern) SetError(v GrokUnknownPatternError)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


