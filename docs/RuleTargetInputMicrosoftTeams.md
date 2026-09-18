# RuleTargetInputMicrosoftTeams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Microsoft Teams target backed by a configured Microsoft Teams integration. | 
**IntegrationId** | **string** | Identifier of the Microsoft Teams integration to use | 
**RenotifyConfig** | Pointer to [**RuleTargetInputSlackRenotifyConfig**](RuleTargetInputSlackRenotifyConfig.md) |  | [optional] 

## Methods

### NewRuleTargetInputMicrosoftTeams

`func NewRuleTargetInputMicrosoftTeams(type_ string, integrationId string, ) *RuleTargetInputMicrosoftTeams`

NewRuleTargetInputMicrosoftTeams instantiates a new RuleTargetInputMicrosoftTeams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleTargetInputMicrosoftTeamsWithDefaults

`func NewRuleTargetInputMicrosoftTeamsWithDefaults() *RuleTargetInputMicrosoftTeams`

NewRuleTargetInputMicrosoftTeamsWithDefaults instantiates a new RuleTargetInputMicrosoftTeams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RuleTargetInputMicrosoftTeams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RuleTargetInputMicrosoftTeams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RuleTargetInputMicrosoftTeams) SetType(v string)`

SetType sets Type field to given value.


### GetIntegrationId

`func (o *RuleTargetInputMicrosoftTeams) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *RuleTargetInputMicrosoftTeams) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *RuleTargetInputMicrosoftTeams) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetRenotifyConfig

`func (o *RuleTargetInputMicrosoftTeams) GetRenotifyConfig() RuleTargetInputSlackRenotifyConfig`

GetRenotifyConfig returns the RenotifyConfig field if non-nil, zero value otherwise.

### GetRenotifyConfigOk

`func (o *RuleTargetInputMicrosoftTeams) GetRenotifyConfigOk() (*RuleTargetInputSlackRenotifyConfig, bool)`

GetRenotifyConfigOk returns a tuple with the RenotifyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenotifyConfig

`func (o *RuleTargetInputMicrosoftTeams) SetRenotifyConfig(v RuleTargetInputSlackRenotifyConfig)`

SetRenotifyConfig sets RenotifyConfig field to given value.

### HasRenotifyConfig

`func (o *RuleTargetInputMicrosoftTeams) HasRenotifyConfig() bool`

HasRenotifyConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


