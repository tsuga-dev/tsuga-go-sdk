# InputNotificationIntegrationJiraSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**SiteUrl** | **string** | Base URL of the Jira Cloud site, like \&quot;https://your-company.atlassian.net\&quot;. Tsuga files issues against this site. | 
**Email** | **string** | Atlassian account email used together with the API token for Basic authentication. | 
**ApiToken** | **string** | Atlassian API token minted for the account, used as the Basic auth password. | 

## Methods

### NewInputNotificationIntegrationJiraSetting

`func NewInputNotificationIntegrationJiraSetting(type_ string, siteUrl string, email string, apiToken string, ) *InputNotificationIntegrationJiraSetting`

NewInputNotificationIntegrationJiraSetting instantiates a new InputNotificationIntegrationJiraSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputNotificationIntegrationJiraSettingWithDefaults

`func NewInputNotificationIntegrationJiraSettingWithDefaults() *InputNotificationIntegrationJiraSetting`

NewInputNotificationIntegrationJiraSettingWithDefaults instantiates a new InputNotificationIntegrationJiraSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputNotificationIntegrationJiraSetting) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputNotificationIntegrationJiraSetting) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputNotificationIntegrationJiraSetting) SetType(v string)`

SetType sets Type field to given value.


### GetSiteUrl

`func (o *InputNotificationIntegrationJiraSetting) GetSiteUrl() string`

GetSiteUrl returns the SiteUrl field if non-nil, zero value otherwise.

### GetSiteUrlOk

`func (o *InputNotificationIntegrationJiraSetting) GetSiteUrlOk() (*string, bool)`

GetSiteUrlOk returns a tuple with the SiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteUrl

`func (o *InputNotificationIntegrationJiraSetting) SetSiteUrl(v string)`

SetSiteUrl sets SiteUrl field to given value.


### GetEmail

`func (o *InputNotificationIntegrationJiraSetting) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InputNotificationIntegrationJiraSetting) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InputNotificationIntegrationJiraSetting) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetApiToken

`func (o *InputNotificationIntegrationJiraSetting) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *InputNotificationIntegrationJiraSetting) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *InputNotificationIntegrationJiraSetting) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


