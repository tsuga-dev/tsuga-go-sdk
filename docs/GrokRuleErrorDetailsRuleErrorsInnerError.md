# GrokRuleErrorDetailsRuleErrorsInnerError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator identifying a rule the grok parser could not parse at all. | 
**Error** | [**GrokUnknownFilterError**](GrokUnknownFilterError.md) |  | 

## Methods

### NewGrokRuleErrorDetailsRuleErrorsInnerError

`func NewGrokRuleErrorDetailsRuleErrorsInnerError(type_ string, error_ GrokUnknownFilterError, ) *GrokRuleErrorDetailsRuleErrorsInnerError`

NewGrokRuleErrorDetailsRuleErrorsInnerError instantiates a new GrokRuleErrorDetailsRuleErrorsInnerError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrokRuleErrorDetailsRuleErrorsInnerErrorWithDefaults

`func NewGrokRuleErrorDetailsRuleErrorsInnerErrorWithDefaults() *GrokRuleErrorDetailsRuleErrorsInnerError`

NewGrokRuleErrorDetailsRuleErrorsInnerErrorWithDefaults instantiates a new GrokRuleErrorDetailsRuleErrorsInnerError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GrokRuleErrorDetailsRuleErrorsInnerError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GrokRuleErrorDetailsRuleErrorsInnerError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GrokRuleErrorDetailsRuleErrorsInnerError) SetType(v string)`

SetType sets Type field to given value.


### GetError

`func (o *GrokRuleErrorDetailsRuleErrorsInnerError) GetError() GrokUnknownFilterError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GrokRuleErrorDetailsRuleErrorsInnerError) GetErrorOk() (*GrokUnknownFilterError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GrokRuleErrorDetailsRuleErrorsInnerError) SetError(v GrokUnknownFilterError)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


