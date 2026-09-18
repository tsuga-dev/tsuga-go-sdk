# RuleTargetConfigGrafanaIrm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Grafana IRM target backed by a configured Grafana IRM integration. | 
**IntegrationId** | **string** | Identifier of the Grafana IRM integration to use | 
**IntegrationName** | **string** | Human readable name of the Grafana IRM integration | 
**RenotifyConfig** | Pointer to [**RuleTargetConfigSlackRenotifyConfig**](RuleTargetConfigSlackRenotifyConfig.md) |  | [optional] 

## Methods

### NewRuleTargetConfigGrafanaIrm

`func NewRuleTargetConfigGrafanaIrm(type_ string, integrationId string, integrationName string, ) *RuleTargetConfigGrafanaIrm`

NewRuleTargetConfigGrafanaIrm instantiates a new RuleTargetConfigGrafanaIrm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleTargetConfigGrafanaIrmWithDefaults

`func NewRuleTargetConfigGrafanaIrmWithDefaults() *RuleTargetConfigGrafanaIrm`

NewRuleTargetConfigGrafanaIrmWithDefaults instantiates a new RuleTargetConfigGrafanaIrm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RuleTargetConfigGrafanaIrm) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RuleTargetConfigGrafanaIrm) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RuleTargetConfigGrafanaIrm) SetType(v string)`

SetType sets Type field to given value.


### GetIntegrationId

`func (o *RuleTargetConfigGrafanaIrm) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *RuleTargetConfigGrafanaIrm) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *RuleTargetConfigGrafanaIrm) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetIntegrationName

`func (o *RuleTargetConfigGrafanaIrm) GetIntegrationName() string`

GetIntegrationName returns the IntegrationName field if non-nil, zero value otherwise.

### GetIntegrationNameOk

`func (o *RuleTargetConfigGrafanaIrm) GetIntegrationNameOk() (*string, bool)`

GetIntegrationNameOk returns a tuple with the IntegrationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationName

`func (o *RuleTargetConfigGrafanaIrm) SetIntegrationName(v string)`

SetIntegrationName sets IntegrationName field to given value.


### GetRenotifyConfig

`func (o *RuleTargetConfigGrafanaIrm) GetRenotifyConfig() RuleTargetConfigSlackRenotifyConfig`

GetRenotifyConfig returns the RenotifyConfig field if non-nil, zero value otherwise.

### GetRenotifyConfigOk

`func (o *RuleTargetConfigGrafanaIrm) GetRenotifyConfigOk() (*RuleTargetConfigSlackRenotifyConfig, bool)`

GetRenotifyConfigOk returns a tuple with the RenotifyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenotifyConfig

`func (o *RuleTargetConfigGrafanaIrm) SetRenotifyConfig(v RuleTargetConfigSlackRenotifyConfig)`

SetRenotifyConfig sets RenotifyConfig field to given value.

### HasRenotifyConfig

`func (o *RuleTargetConfigGrafanaIrm) HasRenotifyConfig() bool`

HasRenotifyConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


