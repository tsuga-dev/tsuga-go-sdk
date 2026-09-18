# RuleTargetConfigServiceNow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | ServiceNow target backed by a configured ServiceNow integration. | 
**IntegrationId** | **string** | Identifier of the ServiceNow integration to use | 
**IntegrationName** | **string** | Human readable name of the ServiceNow integration | 
**RenotifyConfig** | Pointer to [**RuleTargetConfigSlackRenotifyConfig**](RuleTargetConfigSlackRenotifyConfig.md) |  | [optional] 

## Methods

### NewRuleTargetConfigServiceNow

`func NewRuleTargetConfigServiceNow(type_ string, integrationId string, integrationName string, ) *RuleTargetConfigServiceNow`

NewRuleTargetConfigServiceNow instantiates a new RuleTargetConfigServiceNow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleTargetConfigServiceNowWithDefaults

`func NewRuleTargetConfigServiceNowWithDefaults() *RuleTargetConfigServiceNow`

NewRuleTargetConfigServiceNowWithDefaults instantiates a new RuleTargetConfigServiceNow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RuleTargetConfigServiceNow) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RuleTargetConfigServiceNow) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RuleTargetConfigServiceNow) SetType(v string)`

SetType sets Type field to given value.


### GetIntegrationId

`func (o *RuleTargetConfigServiceNow) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *RuleTargetConfigServiceNow) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *RuleTargetConfigServiceNow) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetIntegrationName

`func (o *RuleTargetConfigServiceNow) GetIntegrationName() string`

GetIntegrationName returns the IntegrationName field if non-nil, zero value otherwise.

### GetIntegrationNameOk

`func (o *RuleTargetConfigServiceNow) GetIntegrationNameOk() (*string, bool)`

GetIntegrationNameOk returns a tuple with the IntegrationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationName

`func (o *RuleTargetConfigServiceNow) SetIntegrationName(v string)`

SetIntegrationName sets IntegrationName field to given value.


### GetRenotifyConfig

`func (o *RuleTargetConfigServiceNow) GetRenotifyConfig() RuleTargetConfigSlackRenotifyConfig`

GetRenotifyConfig returns the RenotifyConfig field if non-nil, zero value otherwise.

### GetRenotifyConfigOk

`func (o *RuleTargetConfigServiceNow) GetRenotifyConfigOk() (*RuleTargetConfigSlackRenotifyConfig, bool)`

GetRenotifyConfigOk returns a tuple with the RenotifyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenotifyConfig

`func (o *RuleTargetConfigServiceNow) SetRenotifyConfig(v RuleTargetConfigSlackRenotifyConfig)`

SetRenotifyConfig sets RenotifyConfig field to given value.

### HasRenotifyConfig

`func (o *RuleTargetConfigServiceNow) HasRenotifyConfig() bool`

HasRenotifyConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


