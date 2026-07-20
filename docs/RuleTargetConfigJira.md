# RuleTargetConfigJira

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Jira target backed by a configured Jira integration. | 
**IntegrationId** | **string** | Identifier of the Jira integration to use | 
**IntegrationName** | **string** | Human readable name of the Jira integration | 
**ProjectKey** | **string** | Key of the Jira project that alert issues are filed into, like \&quot;OPS\&quot;. | 
**IssueType** | **string** | Name of the Jira issue type created for alerts, like \&quot;Bug\&quot; or \&quot;Task\&quot;. | 

## Methods

### NewRuleTargetConfigJira

`func NewRuleTargetConfigJira(type_ string, integrationId string, integrationName string, projectKey string, issueType string, ) *RuleTargetConfigJira`

NewRuleTargetConfigJira instantiates a new RuleTargetConfigJira object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleTargetConfigJiraWithDefaults

`func NewRuleTargetConfigJiraWithDefaults() *RuleTargetConfigJira`

NewRuleTargetConfigJiraWithDefaults instantiates a new RuleTargetConfigJira object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RuleTargetConfigJira) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RuleTargetConfigJira) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RuleTargetConfigJira) SetType(v string)`

SetType sets Type field to given value.


### GetIntegrationId

`func (o *RuleTargetConfigJira) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *RuleTargetConfigJira) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *RuleTargetConfigJira) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetIntegrationName

`func (o *RuleTargetConfigJira) GetIntegrationName() string`

GetIntegrationName returns the IntegrationName field if non-nil, zero value otherwise.

### GetIntegrationNameOk

`func (o *RuleTargetConfigJira) GetIntegrationNameOk() (*string, bool)`

GetIntegrationNameOk returns a tuple with the IntegrationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationName

`func (o *RuleTargetConfigJira) SetIntegrationName(v string)`

SetIntegrationName sets IntegrationName field to given value.


### GetProjectKey

`func (o *RuleTargetConfigJira) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *RuleTargetConfigJira) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *RuleTargetConfigJira) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.


### GetIssueType

`func (o *RuleTargetConfigJira) GetIssueType() string`

GetIssueType returns the IssueType field if non-nil, zero value otherwise.

### GetIssueTypeOk

`func (o *RuleTargetConfigJira) GetIssueTypeOk() (*string, bool)`

GetIssueTypeOk returns a tuple with the IssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueType

`func (o *RuleTargetConfigJira) SetIssueType(v string)`

SetIssueType sets IssueType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


