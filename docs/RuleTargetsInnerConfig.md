# RuleTargetsInnerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Slack target backed by a configured Slack integration. | 
**Channel** | **string** | Slack channel ID that receives the notification | 
**IntegrationId** | **string** | Identifier of the Jira integration to use | 
**IntegrationName** | **string** | Human readable name of the Jira integration | 
**Addresses** | **[]string** | Email addresses that will receive the alert | 
**ProjectKey** | **string** | Key of the Jira project that alert issues are filed into, like \&quot;OPS\&quot;. | 
**IssueType** | **string** | Name of the Jira issue type created for alerts, like \&quot;Bug\&quot; or \&quot;Task\&quot;. | 

## Methods

### NewRuleTargetsInnerConfig

`func NewRuleTargetsInnerConfig(type_ string, channel string, integrationId string, integrationName string, addresses []string, projectKey string, issueType string, ) *RuleTargetsInnerConfig`

NewRuleTargetsInnerConfig instantiates a new RuleTargetsInnerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleTargetsInnerConfigWithDefaults

`func NewRuleTargetsInnerConfigWithDefaults() *RuleTargetsInnerConfig`

NewRuleTargetsInnerConfigWithDefaults instantiates a new RuleTargetsInnerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RuleTargetsInnerConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RuleTargetsInnerConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RuleTargetsInnerConfig) SetType(v string)`

SetType sets Type field to given value.


### GetChannel

`func (o *RuleTargetsInnerConfig) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *RuleTargetsInnerConfig) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *RuleTargetsInnerConfig) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetIntegrationId

`func (o *RuleTargetsInnerConfig) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *RuleTargetsInnerConfig) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *RuleTargetsInnerConfig) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetIntegrationName

`func (o *RuleTargetsInnerConfig) GetIntegrationName() string`

GetIntegrationName returns the IntegrationName field if non-nil, zero value otherwise.

### GetIntegrationNameOk

`func (o *RuleTargetsInnerConfig) GetIntegrationNameOk() (*string, bool)`

GetIntegrationNameOk returns a tuple with the IntegrationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationName

`func (o *RuleTargetsInnerConfig) SetIntegrationName(v string)`

SetIntegrationName sets IntegrationName field to given value.


### GetAddresses

`func (o *RuleTargetsInnerConfig) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *RuleTargetsInnerConfig) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *RuleTargetsInnerConfig) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetProjectKey

`func (o *RuleTargetsInnerConfig) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *RuleTargetsInnerConfig) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *RuleTargetsInnerConfig) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.


### GetIssueType

`func (o *RuleTargetsInnerConfig) GetIssueType() string`

GetIssueType returns the IssueType field if non-nil, zero value otherwise.

### GetIssueTypeOk

`func (o *RuleTargetsInnerConfig) GetIssueTypeOk() (*string, bool)`

GetIssueTypeOk returns a tuple with the IssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueType

`func (o *RuleTargetsInnerConfig) SetIssueType(v string)`

SetIssueType sets IssueType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


