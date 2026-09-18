# RuleTargetInputJira

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Jira target backed by a configured Jira integration. Jira targets file one issue per notification and do not support renotification. | 
**IntegrationId** | **string** | Identifier of the Jira integration to use | 
**ProjectKey** | **string** | Key of the Jira project that alert issues are filed into, like \&quot;OPS\&quot;. | 
**IssueType** | **string** | Name of the Jira issue type created for alerts, like \&quot;Bug\&quot; or \&quot;Task\&quot;. | 
**OpenStatus** | Pointer to **string** | Jira status the alert ticket is moved to right after it is filed for a firing alert. A later transition between two firing states leaves the ticket status alone. | [optional] 
**ClosedStatus** | Pointer to **string** | Jira status the alert ticket is moved to when the alert resolves. Only takes effect when the rule delivers recovery transitions. | [optional] 

## Methods

### NewRuleTargetInputJira

`func NewRuleTargetInputJira(type_ string, integrationId string, projectKey string, issueType string, ) *RuleTargetInputJira`

NewRuleTargetInputJira instantiates a new RuleTargetInputJira object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleTargetInputJiraWithDefaults

`func NewRuleTargetInputJiraWithDefaults() *RuleTargetInputJira`

NewRuleTargetInputJiraWithDefaults instantiates a new RuleTargetInputJira object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RuleTargetInputJira) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RuleTargetInputJira) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RuleTargetInputJira) SetType(v string)`

SetType sets Type field to given value.


### GetIntegrationId

`func (o *RuleTargetInputJira) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *RuleTargetInputJira) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *RuleTargetInputJira) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetProjectKey

`func (o *RuleTargetInputJira) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *RuleTargetInputJira) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *RuleTargetInputJira) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.


### GetIssueType

`func (o *RuleTargetInputJira) GetIssueType() string`

GetIssueType returns the IssueType field if non-nil, zero value otherwise.

### GetIssueTypeOk

`func (o *RuleTargetInputJira) GetIssueTypeOk() (*string, bool)`

GetIssueTypeOk returns a tuple with the IssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueType

`func (o *RuleTargetInputJira) SetIssueType(v string)`

SetIssueType sets IssueType field to given value.


### GetOpenStatus

`func (o *RuleTargetInputJira) GetOpenStatus() string`

GetOpenStatus returns the OpenStatus field if non-nil, zero value otherwise.

### GetOpenStatusOk

`func (o *RuleTargetInputJira) GetOpenStatusOk() (*string, bool)`

GetOpenStatusOk returns a tuple with the OpenStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStatus

`func (o *RuleTargetInputJira) SetOpenStatus(v string)`

SetOpenStatus sets OpenStatus field to given value.

### HasOpenStatus

`func (o *RuleTargetInputJira) HasOpenStatus() bool`

HasOpenStatus returns a boolean if a field has been set.

### GetClosedStatus

`func (o *RuleTargetInputJira) GetClosedStatus() string`

GetClosedStatus returns the ClosedStatus field if non-nil, zero value otherwise.

### GetClosedStatusOk

`func (o *RuleTargetInputJira) GetClosedStatusOk() (*string, bool)`

GetClosedStatusOk returns a tuple with the ClosedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedStatus

`func (o *RuleTargetInputJira) SetClosedStatus(v string)`

SetClosedStatus sets ClosedStatus field to given value.

### HasClosedStatus

`func (o *RuleTargetInputJira) HasClosedStatus() bool`

HasClosedStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


