# RuleTargetConfigEmail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Addresses** | **[]string** | Email addresses that will receive the alert | 

## Methods

### NewRuleTargetConfigEmail

`func NewRuleTargetConfigEmail(type_ string, addresses []string, ) *RuleTargetConfigEmail`

NewRuleTargetConfigEmail instantiates a new RuleTargetConfigEmail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleTargetConfigEmailWithDefaults

`func NewRuleTargetConfigEmailWithDefaults() *RuleTargetConfigEmail`

NewRuleTargetConfigEmailWithDefaults instantiates a new RuleTargetConfigEmail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RuleTargetConfigEmail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RuleTargetConfigEmail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RuleTargetConfigEmail) SetType(v string)`

SetType sets Type field to given value.


### GetAddresses

`func (o *RuleTargetConfigEmail) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *RuleTargetConfigEmail) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *RuleTargetConfigEmail) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


