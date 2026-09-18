# ClientErrorResponseDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator identifying tag-policy validation details. Set by Tsuga to &#x60;tag_policy_violation&#x60;. Returned when the error code is &#x60;TAG_POLICY_VALIDATION_ERROR&#x60;. | 
**Violations** | [**[]TagPolicyViolationDetailsViolationsInner**](TagPolicyViolationDetailsViolationsInner.md) | Tag-policy violations that caused the request to fail. Built from active policies applicable to the submitted asset type and owner team. Returned on tag-policy validation errors; can contain multiple policies and multiple assets for bulk requests. | 
**RuleErrors** | [**[]GrokRuleErrorDetailsRuleErrorsInner**](GrokRuleErrorDetailsRuleErrorsInner.md) | Grok rules that failed to compile, with full structured detail. One entry per invalid rule. | 

## Methods

### NewClientErrorResponseDetails

`func NewClientErrorResponseDetails(type_ string, violations []TagPolicyViolationDetailsViolationsInner, ruleErrors []GrokRuleErrorDetailsRuleErrorsInner, ) *ClientErrorResponseDetails`

NewClientErrorResponseDetails instantiates a new ClientErrorResponseDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientErrorResponseDetailsWithDefaults

`func NewClientErrorResponseDetailsWithDefaults() *ClientErrorResponseDetails`

NewClientErrorResponseDetailsWithDefaults instantiates a new ClientErrorResponseDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ClientErrorResponseDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClientErrorResponseDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClientErrorResponseDetails) SetType(v string)`

SetType sets Type field to given value.


### GetViolations

`func (o *ClientErrorResponseDetails) GetViolations() []TagPolicyViolationDetailsViolationsInner`

GetViolations returns the Violations field if non-nil, zero value otherwise.

### GetViolationsOk

`func (o *ClientErrorResponseDetails) GetViolationsOk() (*[]TagPolicyViolationDetailsViolationsInner, bool)`

GetViolationsOk returns a tuple with the Violations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolations

`func (o *ClientErrorResponseDetails) SetViolations(v []TagPolicyViolationDetailsViolationsInner)`

SetViolations sets Violations field to given value.


### GetRuleErrors

`func (o *ClientErrorResponseDetails) GetRuleErrors() []GrokRuleErrorDetailsRuleErrorsInner`

GetRuleErrors returns the RuleErrors field if non-nil, zero value otherwise.

### GetRuleErrorsOk

`func (o *ClientErrorResponseDetails) GetRuleErrorsOk() (*[]GrokRuleErrorDetailsRuleErrorsInner, bool)`

GetRuleErrorsOk returns a tuple with the RuleErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleErrors

`func (o *ClientErrorResponseDetails) SetRuleErrors(v []GrokRuleErrorDetailsRuleErrorsInner)`

SetRuleErrors sets RuleErrors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


