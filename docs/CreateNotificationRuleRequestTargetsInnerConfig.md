# CreateNotificationRuleRequestTargetsInnerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Slack target backed by a configured Slack integration. | 
**Channel** | **string** | Slack channel ID that receives the notification | 
**IntegrationId** | **string** | Identifier of the Jira integration to use | 
**RenotifyConfig** | Pointer to [**RuleTargetInputSlackRenotifyConfig**](RuleTargetInputSlackRenotifyConfig.md) |  | [optional] 
**Addresses** | **[]string** | Email addresses that will receive the alert | 
**ProjectKey** | **string** | Key of the Jira project that alert issues are filed into, like \&quot;OPS\&quot;. | 
**IssueType** | **string** | Name of the Jira issue type created for alerts, like \&quot;Bug\&quot; or \&quot;Task\&quot;. | 
**OpenStatus** | Pointer to **string** | Jira status the alert ticket is moved to right after it is filed for a firing alert. A later transition between two firing states leaves the ticket status alone. | [optional] 
**ClosedStatus** | Pointer to **string** | Jira status the alert ticket is moved to when the alert resolves. Only takes effect when the rule delivers recovery transitions. | [optional] 

## Methods

### NewCreateNotificationRuleRequestTargetsInnerConfig

`func NewCreateNotificationRuleRequestTargetsInnerConfig(type_ string, channel string, integrationId string, addresses []string, projectKey string, issueType string, ) *CreateNotificationRuleRequestTargetsInnerConfig`

NewCreateNotificationRuleRequestTargetsInnerConfig instantiates a new CreateNotificationRuleRequestTargetsInnerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNotificationRuleRequestTargetsInnerConfigWithDefaults

`func NewCreateNotificationRuleRequestTargetsInnerConfigWithDefaults() *CreateNotificationRuleRequestTargetsInnerConfig`

NewCreateNotificationRuleRequestTargetsInnerConfigWithDefaults instantiates a new CreateNotificationRuleRequestTargetsInnerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) SetType(v string)`

SetType sets Type field to given value.


### GetChannel

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetIntegrationId

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetRenotifyConfig

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) GetRenotifyConfig() RuleTargetInputSlackRenotifyConfig`

GetRenotifyConfig returns the RenotifyConfig field if non-nil, zero value otherwise.

### GetRenotifyConfigOk

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) GetRenotifyConfigOk() (*RuleTargetInputSlackRenotifyConfig, bool)`

GetRenotifyConfigOk returns a tuple with the RenotifyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenotifyConfig

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) SetRenotifyConfig(v RuleTargetInputSlackRenotifyConfig)`

SetRenotifyConfig sets RenotifyConfig field to given value.

### HasRenotifyConfig

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) HasRenotifyConfig() bool`

HasRenotifyConfig returns a boolean if a field has been set.

### GetAddresses

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetProjectKey

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.


### GetIssueType

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) GetIssueType() string`

GetIssueType returns the IssueType field if non-nil, zero value otherwise.

### GetIssueTypeOk

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) GetIssueTypeOk() (*string, bool)`

GetIssueTypeOk returns a tuple with the IssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueType

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) SetIssueType(v string)`

SetIssueType sets IssueType field to given value.


### GetOpenStatus

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) GetOpenStatus() string`

GetOpenStatus returns the OpenStatus field if non-nil, zero value otherwise.

### GetOpenStatusOk

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) GetOpenStatusOk() (*string, bool)`

GetOpenStatusOk returns a tuple with the OpenStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStatus

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) SetOpenStatus(v string)`

SetOpenStatus sets OpenStatus field to given value.

### HasOpenStatus

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) HasOpenStatus() bool`

HasOpenStatus returns a boolean if a field has been set.

### GetClosedStatus

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) GetClosedStatus() string`

GetClosedStatus returns the ClosedStatus field if non-nil, zero value otherwise.

### GetClosedStatusOk

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) GetClosedStatusOk() (*string, bool)`

GetClosedStatusOk returns a tuple with the ClosedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedStatus

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) SetClosedStatus(v string)`

SetClosedStatus sets ClosedStatus field to given value.

### HasClosedStatus

`func (o *CreateNotificationRuleRequestTargetsInnerConfig) HasClosedStatus() bool`

HasClosedStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


