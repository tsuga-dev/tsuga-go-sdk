# GrokRuleErrorDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator identifying grok rule validation details. Set by Tsuga to &#x60;grok_rule_error&#x60;. Returned when the error code is &#x60;GROK_RULE_VALIDATION_ERROR&#x60;. | 
**RuleErrors** | [**[]GrokRuleErrorDetailsRuleErrorsInner**](GrokRuleErrorDetailsRuleErrorsInner.md) | Grok rules that failed to compile, with full structured detail. One entry per invalid rule. | 

## Methods

### NewGrokRuleErrorDetails

`func NewGrokRuleErrorDetails(type_ string, ruleErrors []GrokRuleErrorDetailsRuleErrorsInner, ) *GrokRuleErrorDetails`

NewGrokRuleErrorDetails instantiates a new GrokRuleErrorDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrokRuleErrorDetailsWithDefaults

`func NewGrokRuleErrorDetailsWithDefaults() *GrokRuleErrorDetails`

NewGrokRuleErrorDetailsWithDefaults instantiates a new GrokRuleErrorDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GrokRuleErrorDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GrokRuleErrorDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GrokRuleErrorDetails) SetType(v string)`

SetType sets Type field to given value.


### GetRuleErrors

`func (o *GrokRuleErrorDetails) GetRuleErrors() []GrokRuleErrorDetailsRuleErrorsInner`

GetRuleErrors returns the RuleErrors field if non-nil, zero value otherwise.

### GetRuleErrorsOk

`func (o *GrokRuleErrorDetails) GetRuleErrorsOk() (*[]GrokRuleErrorDetailsRuleErrorsInner, bool)`

GetRuleErrorsOk returns a tuple with the RuleErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleErrors

`func (o *GrokRuleErrorDetails) SetRuleErrors(v []GrokRuleErrorDetailsRuleErrorsInner)`

SetRuleErrors sets RuleErrors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


