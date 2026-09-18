# GrokRuleErrorDetailsRuleErrorsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rule** | **string** | The rule that failed to compile. | 
**RuleIndex** | **int32** | 0-based index of the failing rule in the submitted &#x60;rules&#x60; array. | 
**Error** | [**GrokRuleErrorDetailsRuleErrorsInnerError**](GrokRuleErrorDetailsRuleErrorsInnerError.md) |  | 

## Methods

### NewGrokRuleErrorDetailsRuleErrorsInner

`func NewGrokRuleErrorDetailsRuleErrorsInner(rule string, ruleIndex int32, error_ GrokRuleErrorDetailsRuleErrorsInnerError, ) *GrokRuleErrorDetailsRuleErrorsInner`

NewGrokRuleErrorDetailsRuleErrorsInner instantiates a new GrokRuleErrorDetailsRuleErrorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrokRuleErrorDetailsRuleErrorsInnerWithDefaults

`func NewGrokRuleErrorDetailsRuleErrorsInnerWithDefaults() *GrokRuleErrorDetailsRuleErrorsInner`

NewGrokRuleErrorDetailsRuleErrorsInnerWithDefaults instantiates a new GrokRuleErrorDetailsRuleErrorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRule

`func (o *GrokRuleErrorDetailsRuleErrorsInner) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *GrokRuleErrorDetailsRuleErrorsInner) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *GrokRuleErrorDetailsRuleErrorsInner) SetRule(v string)`

SetRule sets Rule field to given value.


### GetRuleIndex

`func (o *GrokRuleErrorDetailsRuleErrorsInner) GetRuleIndex() int32`

GetRuleIndex returns the RuleIndex field if non-nil, zero value otherwise.

### GetRuleIndexOk

`func (o *GrokRuleErrorDetailsRuleErrorsInner) GetRuleIndexOk() (*int32, bool)`

GetRuleIndexOk returns a tuple with the RuleIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleIndex

`func (o *GrokRuleErrorDetailsRuleErrorsInner) SetRuleIndex(v int32)`

SetRuleIndex sets RuleIndex field to given value.


### GetError

`func (o *GrokRuleErrorDetailsRuleErrorsInner) GetError() GrokRuleErrorDetailsRuleErrorsInnerError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GrokRuleErrorDetailsRuleErrorsInner) GetErrorOk() (*GrokRuleErrorDetailsRuleErrorsInnerError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GrokRuleErrorDetailsRuleErrorsInner) SetError(v GrokRuleErrorDetailsRuleErrorsInnerError)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


