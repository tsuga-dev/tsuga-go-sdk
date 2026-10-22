# CreateNotificationIntegrationRequestSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**AlertSourceConfigId** | **NullableString** | Display name of the graph widget. Pass null to remove it. | 
**Token** | **NullableString** | Display name of the graph widget. Pass null to remove it. | 
**Url** | **string** |  | 
**IntegrationKey** | **NullableString** | Display name of the graph widget. Pass null to remove it. | 
**WebhookUrl** | **string** |  | 
**Method** | **string** |  | 
**Authentication** | [**CreateNotificationIntegrationRequestSettingOneOf4Authentication**](CreateNotificationIntegrationRequestSettingOneOf4Authentication.md) |  | 
**CustomHeaders** | **map[string]string** |  | 
**PayloadTemplate** | [**CreateNotificationIntegrationRequestSettingOneOf4PayloadTemplate**](CreateNotificationIntegrationRequestSettingOneOf4PayloadTemplate.md) |  | 
**EventNameMapping** | Pointer to [**CreateNotificationIntegrationRequestSettingOneOf4EventNameMapping**](CreateNotificationIntegrationRequestSettingOneOf4EventNameMapping.md) |  | [optional] 
**InstanceName** | **string** | Instance name of the ServiceNow instance. This corresponds to the subdomain part of the instance URL, like \&quot;my-instance\&quot; in \&quot;my-instance.service-now.com\&quot; | 
**Username** | **NullableString** | Folder filter for dashboard query results. Optional; pass a folder ID to return dashboards in that folder, or null to return only dashboards that are not in any folder. | 
**Password** | **NullableString** | Folder filter for dashboard query results. Optional; pass a folder ID to return dashboards in that folder, or null to return only dashboards that are not in any folder. | 
**SiteUrl** | **string** | Base URL of the Jira Cloud site, like \&quot;https://your-company.atlassian.net\&quot;. Tsuga files issues against this site. | 
**Email** | **string** | Atlassian account email used together with the API token for Basic authentication. | 
**ApiToken** | **string** | Atlassian API token minted for the account, used as the Basic auth password. | 

## Methods

### NewCreateNotificationIntegrationRequestSetting

`func NewCreateNotificationIntegrationRequestSetting(type_ string, alertSourceConfigId NullableString, token NullableString, url string, integrationKey NullableString, webhookUrl string, method string, authentication CreateNotificationIntegrationRequestSettingOneOf4Authentication, customHeaders map[string]*string, payloadTemplate CreateNotificationIntegrationRequestSettingOneOf4PayloadTemplate, instanceName string, username NullableString, password NullableString, siteUrl string, email string, apiToken string, ) *CreateNotificationIntegrationRequestSetting`

NewCreateNotificationIntegrationRequestSetting instantiates a new CreateNotificationIntegrationRequestSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNotificationIntegrationRequestSettingWithDefaults

`func NewCreateNotificationIntegrationRequestSettingWithDefaults() *CreateNotificationIntegrationRequestSetting`

NewCreateNotificationIntegrationRequestSettingWithDefaults instantiates a new CreateNotificationIntegrationRequestSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateNotificationIntegrationRequestSetting) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateNotificationIntegrationRequestSetting) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateNotificationIntegrationRequestSetting) SetType(v string)`

SetType sets Type field to given value.


### GetAlertSourceConfigId

`func (o *CreateNotificationIntegrationRequestSetting) GetAlertSourceConfigId() string`

GetAlertSourceConfigId returns the AlertSourceConfigId field if non-nil, zero value otherwise.

### GetAlertSourceConfigIdOk

`func (o *CreateNotificationIntegrationRequestSetting) GetAlertSourceConfigIdOk() (*string, bool)`

GetAlertSourceConfigIdOk returns a tuple with the AlertSourceConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertSourceConfigId

`func (o *CreateNotificationIntegrationRequestSetting) SetAlertSourceConfigId(v string)`

SetAlertSourceConfigId sets AlertSourceConfigId field to given value.


### SetAlertSourceConfigIdNil

`func (o *CreateNotificationIntegrationRequestSetting) SetAlertSourceConfigIdNil(b bool)`

 SetAlertSourceConfigIdNil sets the value for AlertSourceConfigId to be an explicit nil

### UnsetAlertSourceConfigId
`func (o *CreateNotificationIntegrationRequestSetting) UnsetAlertSourceConfigId()`

UnsetAlertSourceConfigId ensures that no value is present for AlertSourceConfigId, not even an explicit nil
### GetToken

`func (o *CreateNotificationIntegrationRequestSetting) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateNotificationIntegrationRequestSetting) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateNotificationIntegrationRequestSetting) SetToken(v string)`

SetToken sets Token field to given value.


### SetTokenNil

`func (o *CreateNotificationIntegrationRequestSetting) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *CreateNotificationIntegrationRequestSetting) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetUrl

`func (o *CreateNotificationIntegrationRequestSetting) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateNotificationIntegrationRequestSetting) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateNotificationIntegrationRequestSetting) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetIntegrationKey

`func (o *CreateNotificationIntegrationRequestSetting) GetIntegrationKey() string`

GetIntegrationKey returns the IntegrationKey field if non-nil, zero value otherwise.

### GetIntegrationKeyOk

`func (o *CreateNotificationIntegrationRequestSetting) GetIntegrationKeyOk() (*string, bool)`

GetIntegrationKeyOk returns a tuple with the IntegrationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationKey

`func (o *CreateNotificationIntegrationRequestSetting) SetIntegrationKey(v string)`

SetIntegrationKey sets IntegrationKey field to given value.


### SetIntegrationKeyNil

`func (o *CreateNotificationIntegrationRequestSetting) SetIntegrationKeyNil(b bool)`

 SetIntegrationKeyNil sets the value for IntegrationKey to be an explicit nil

### UnsetIntegrationKey
`func (o *CreateNotificationIntegrationRequestSetting) UnsetIntegrationKey()`

UnsetIntegrationKey ensures that no value is present for IntegrationKey, not even an explicit nil
### GetWebhookUrl

`func (o *CreateNotificationIntegrationRequestSetting) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *CreateNotificationIntegrationRequestSetting) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *CreateNotificationIntegrationRequestSetting) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.


### GetMethod

`func (o *CreateNotificationIntegrationRequestSetting) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *CreateNotificationIntegrationRequestSetting) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *CreateNotificationIntegrationRequestSetting) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetAuthentication

`func (o *CreateNotificationIntegrationRequestSetting) GetAuthentication() CreateNotificationIntegrationRequestSettingOneOf4Authentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *CreateNotificationIntegrationRequestSetting) GetAuthenticationOk() (*CreateNotificationIntegrationRequestSettingOneOf4Authentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *CreateNotificationIntegrationRequestSetting) SetAuthentication(v CreateNotificationIntegrationRequestSettingOneOf4Authentication)`

SetAuthentication sets Authentication field to given value.


### GetCustomHeaders

`func (o *CreateNotificationIntegrationRequestSetting) GetCustomHeaders() map[string]*string`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *CreateNotificationIntegrationRequestSetting) GetCustomHeadersOk() (*map[string]*string, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *CreateNotificationIntegrationRequestSetting) SetCustomHeaders(v map[string]*string)`

SetCustomHeaders sets CustomHeaders field to given value.


### GetPayloadTemplate

`func (o *CreateNotificationIntegrationRequestSetting) GetPayloadTemplate() CreateNotificationIntegrationRequestSettingOneOf4PayloadTemplate`

GetPayloadTemplate returns the PayloadTemplate field if non-nil, zero value otherwise.

### GetPayloadTemplateOk

`func (o *CreateNotificationIntegrationRequestSetting) GetPayloadTemplateOk() (*CreateNotificationIntegrationRequestSettingOneOf4PayloadTemplate, bool)`

GetPayloadTemplateOk returns a tuple with the PayloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplate

`func (o *CreateNotificationIntegrationRequestSetting) SetPayloadTemplate(v CreateNotificationIntegrationRequestSettingOneOf4PayloadTemplate)`

SetPayloadTemplate sets PayloadTemplate field to given value.


### GetEventNameMapping

`func (o *CreateNotificationIntegrationRequestSetting) GetEventNameMapping() CreateNotificationIntegrationRequestSettingOneOf4EventNameMapping`

GetEventNameMapping returns the EventNameMapping field if non-nil, zero value otherwise.

### GetEventNameMappingOk

`func (o *CreateNotificationIntegrationRequestSetting) GetEventNameMappingOk() (*CreateNotificationIntegrationRequestSettingOneOf4EventNameMapping, bool)`

GetEventNameMappingOk returns a tuple with the EventNameMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNameMapping

`func (o *CreateNotificationIntegrationRequestSetting) SetEventNameMapping(v CreateNotificationIntegrationRequestSettingOneOf4EventNameMapping)`

SetEventNameMapping sets EventNameMapping field to given value.

### HasEventNameMapping

`func (o *CreateNotificationIntegrationRequestSetting) HasEventNameMapping() bool`

HasEventNameMapping returns a boolean if a field has been set.

### GetInstanceName

`func (o *CreateNotificationIntegrationRequestSetting) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *CreateNotificationIntegrationRequestSetting) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *CreateNotificationIntegrationRequestSetting) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.


### GetUsername

`func (o *CreateNotificationIntegrationRequestSetting) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateNotificationIntegrationRequestSetting) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateNotificationIntegrationRequestSetting) SetUsername(v string)`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *CreateNotificationIntegrationRequestSetting) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *CreateNotificationIntegrationRequestSetting) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *CreateNotificationIntegrationRequestSetting) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateNotificationIntegrationRequestSetting) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateNotificationIntegrationRequestSetting) SetPassword(v string)`

SetPassword sets Password field to given value.


### SetPasswordNil

`func (o *CreateNotificationIntegrationRequestSetting) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *CreateNotificationIntegrationRequestSetting) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetSiteUrl

`func (o *CreateNotificationIntegrationRequestSetting) GetSiteUrl() string`

GetSiteUrl returns the SiteUrl field if non-nil, zero value otherwise.

### GetSiteUrlOk

`func (o *CreateNotificationIntegrationRequestSetting) GetSiteUrlOk() (*string, bool)`

GetSiteUrlOk returns a tuple with the SiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteUrl

`func (o *CreateNotificationIntegrationRequestSetting) SetSiteUrl(v string)`

SetSiteUrl sets SiteUrl field to given value.


### GetEmail

`func (o *CreateNotificationIntegrationRequestSetting) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateNotificationIntegrationRequestSetting) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateNotificationIntegrationRequestSetting) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetApiToken

`func (o *CreateNotificationIntegrationRequestSetting) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *CreateNotificationIntegrationRequestSetting) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *CreateNotificationIntegrationRequestSetting) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


