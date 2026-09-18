# RuleTargetInputIncidentIo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | incident.io target backed by a configured incident.io integration. | 
**IntegrationId** | **string** | Identifier of the incident.io integration to use | 
**RenotifyConfig** | Pointer to [**RuleTargetInputSlackRenotifyConfig**](RuleTargetInputSlackRenotifyConfig.md) |  | [optional] 

## Methods

### NewRuleTargetInputIncidentIo

`func NewRuleTargetInputIncidentIo(type_ string, integrationId string, ) *RuleTargetInputIncidentIo`

NewRuleTargetInputIncidentIo instantiates a new RuleTargetInputIncidentIo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleTargetInputIncidentIoWithDefaults

`func NewRuleTargetInputIncidentIoWithDefaults() *RuleTargetInputIncidentIo`

NewRuleTargetInputIncidentIoWithDefaults instantiates a new RuleTargetInputIncidentIo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RuleTargetInputIncidentIo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RuleTargetInputIncidentIo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RuleTargetInputIncidentIo) SetType(v string)`

SetType sets Type field to given value.


### GetIntegrationId

`func (o *RuleTargetInputIncidentIo) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *RuleTargetInputIncidentIo) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *RuleTargetInputIncidentIo) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetRenotifyConfig

`func (o *RuleTargetInputIncidentIo) GetRenotifyConfig() RuleTargetInputSlackRenotifyConfig`

GetRenotifyConfig returns the RenotifyConfig field if non-nil, zero value otherwise.

### GetRenotifyConfigOk

`func (o *RuleTargetInputIncidentIo) GetRenotifyConfigOk() (*RuleTargetInputSlackRenotifyConfig, bool)`

GetRenotifyConfigOk returns a tuple with the RenotifyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenotifyConfig

`func (o *RuleTargetInputIncidentIo) SetRenotifyConfig(v RuleTargetInputSlackRenotifyConfig)`

SetRenotifyConfig sets RenotifyConfig field to given value.

### HasRenotifyConfig

`func (o *RuleTargetInputIncidentIo) HasRenotifyConfig() bool`

HasRenotifyConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


