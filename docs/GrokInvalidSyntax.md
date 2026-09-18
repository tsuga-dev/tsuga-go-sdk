# GrokInvalidSyntax

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator identifying a rule the grok parser could not parse at all. | 
**Error** | [**GrokInvalidSyntaxError**](GrokInvalidSyntaxError.md) |  | 

## Methods

### NewGrokInvalidSyntax

`func NewGrokInvalidSyntax(type_ string, error_ GrokInvalidSyntaxError, ) *GrokInvalidSyntax`

NewGrokInvalidSyntax instantiates a new GrokInvalidSyntax object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrokInvalidSyntaxWithDefaults

`func NewGrokInvalidSyntaxWithDefaults() *GrokInvalidSyntax`

NewGrokInvalidSyntaxWithDefaults instantiates a new GrokInvalidSyntax object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GrokInvalidSyntax) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GrokInvalidSyntax) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GrokInvalidSyntax) SetType(v string)`

SetType sets Type field to given value.


### GetError

`func (o *GrokInvalidSyntax) GetError() GrokInvalidSyntaxError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GrokInvalidSyntax) GetErrorOk() (*GrokInvalidSyntaxError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GrokInvalidSyntax) SetError(v GrokInvalidSyntaxError)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


