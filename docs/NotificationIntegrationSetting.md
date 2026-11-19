# NotificationIntegrationSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**AlertSourceConfigId** | **string** |  | 
**TokenLastChars** | **string** |  | 
**UrlLastChars** | **string** |  | 
**IntegrationKeyLastChars** | **string** |  | 
**WebhookUrlLastChars** | **string** |  | 
**Method** | **string** |  | 
**Authentication** | [**NotificationIntegrationSettingWebhookAuthentication**](NotificationIntegrationSettingWebhookAuthentication.md) |  | 
**CustomHeaders** | Pointer to **map[string]string** |  | [optional] 
**PayloadTemplate** | [**NotificationIntegrationSettingWebhookPayloadTemplate**](NotificationIntegrationSettingWebhookPayloadTemplate.md) |  | 
**EventNameMapping** | Pointer to [**NotificationIntegrationSettingWebhookEventNameMapping**](NotificationIntegrationSettingWebhookEventNameMapping.md) |  | [optional] 
**InstanceName** | **string** |  | 
**Username** | **string** |  | 
**PasswordLastChars** | **string** |  | 
**SiteUrl** | **string** |  | 
**Email** | **string** |  | 
**ApiTokenLastChars** | **string** |  | 

## Methods

### NewNotificationIntegrationSetting

`func NewNotificationIntegrationSetting(type_ string, alertSourceConfigId string, tokenLastChars string, urlLastChars string, integrationKeyLastChars string, webhookUrlLastChars string, method string, authentication NotificationIntegrationSettingWebhookAuthentication, payloadTemplate NotificationIntegrationSettingWebhookPayloadTemplate, instanceName string, username string, passwordLastChars string, siteUrl string, email string, apiTokenLastChars string, ) *NotificationIntegrationSetting`

NewNotificationIntegrationSetting instantiates a new NotificationIntegrationSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationIntegrationSettingWithDefaults

`func NewNotificationIntegrationSettingWithDefaults() *NotificationIntegrationSetting`

NewNotificationIntegrationSettingWithDefaults instantiates a new NotificationIntegrationSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NotificationIntegrationSetting) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationIntegrationSetting) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationIntegrationSetting) SetType(v string)`

SetType sets Type field to given value.


### GetAlertSourceConfigId

`func (o *NotificationIntegrationSetting) GetAlertSourceConfigId() string`

GetAlertSourceConfigId returns the AlertSourceConfigId field if non-nil, zero value otherwise.

### GetAlertSourceConfigIdOk

`func (o *NotificationIntegrationSetting) GetAlertSourceConfigIdOk() (*string, bool)`

GetAlertSourceConfigIdOk returns a tuple with the AlertSourceConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertSourceConfigId

`func (o *NotificationIntegrationSetting) SetAlertSourceConfigId(v string)`

SetAlertSourceConfigId sets AlertSourceConfigId field to given value.


### GetTokenLastChars

`func (o *NotificationIntegrationSetting) GetTokenLastChars() string`

GetTokenLastChars returns the TokenLastChars field if non-nil, zero value otherwise.

### GetTokenLastCharsOk

`func (o *NotificationIntegrationSetting) GetTokenLastCharsOk() (*string, bool)`

GetTokenLastCharsOk returns a tuple with the TokenLastChars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenLastChars

`func (o *NotificationIntegrationSetting) SetTokenLastChars(v string)`

SetTokenLastChars sets TokenLastChars field to given value.


### GetUrlLastChars

`func (o *NotificationIntegrationSetting) GetUrlLastChars() string`

GetUrlLastChars returns the UrlLastChars field if non-nil, zero value otherwise.

### GetUrlLastCharsOk

`func (o *NotificationIntegrationSetting) GetUrlLastCharsOk() (*string, bool)`

GetUrlLastCharsOk returns a tuple with the UrlLastChars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlLastChars

`func (o *NotificationIntegrationSetting) SetUrlLastChars(v string)`

SetUrlLastChars sets UrlLastChars field to given value.


### GetIntegrationKeyLastChars

`func (o *NotificationIntegrationSetting) GetIntegrationKeyLastChars() string`

GetIntegrationKeyLastChars returns the IntegrationKeyLastChars field if non-nil, zero value otherwise.

### GetIntegrationKeyLastCharsOk

`func (o *NotificationIntegrationSetting) GetIntegrationKeyLastCharsOk() (*string, bool)`

GetIntegrationKeyLastCharsOk returns a tuple with the IntegrationKeyLastChars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationKeyLastChars

`func (o *NotificationIntegrationSetting) SetIntegrationKeyLastChars(v string)`

SetIntegrationKeyLastChars sets IntegrationKeyLastChars field to given value.


### GetWebhookUrlLastChars

`func (o *NotificationIntegrationSetting) GetWebhookUrlLastChars() string`

GetWebhookUrlLastChars returns the WebhookUrlLastChars field if non-nil, zero value otherwise.

### GetWebhookUrlLastCharsOk

`func (o *NotificationIntegrationSetting) GetWebhookUrlLastCharsOk() (*string, bool)`

GetWebhookUrlLastCharsOk returns a tuple with the WebhookUrlLastChars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrlLastChars

`func (o *NotificationIntegrationSetting) SetWebhookUrlLastChars(v string)`

SetWebhookUrlLastChars sets WebhookUrlLastChars field to given value.


### GetMethod

`func (o *NotificationIntegrationSetting) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *NotificationIntegrationSetting) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *NotificationIntegrationSetting) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetAuthentication

`func (o *NotificationIntegrationSetting) GetAuthentication() NotificationIntegrationSettingWebhookAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *NotificationIntegrationSetting) GetAuthenticationOk() (*NotificationIntegrationSettingWebhookAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *NotificationIntegrationSetting) SetAuthentication(v NotificationIntegrationSettingWebhookAuthentication)`

SetAuthentication sets Authentication field to given value.


### GetCustomHeaders

`func (o *NotificationIntegrationSetting) GetCustomHeaders() map[string]string`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *NotificationIntegrationSetting) GetCustomHeadersOk() (*map[string]string, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *NotificationIntegrationSetting) SetCustomHeaders(v map[string]string)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *NotificationIntegrationSetting) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.

### GetPayloadTemplate

`func (o *NotificationIntegrationSetting) GetPayloadTemplate() NotificationIntegrationSettingWebhookPayloadTemplate`

GetPayloadTemplate returns the PayloadTemplate field if non-nil, zero value otherwise.

### GetPayloadTemplateOk

`func (o *NotificationIntegrationSetting) GetPayloadTemplateOk() (*NotificationIntegrationSettingWebhookPayloadTemplate, bool)`

GetPayloadTemplateOk returns a tuple with the PayloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplate

`func (o *NotificationIntegrationSetting) SetPayloadTemplate(v NotificationIntegrationSettingWebhookPayloadTemplate)`

SetPayloadTemplate sets PayloadTemplate field to given value.


### GetEventNameMapping

`func (o *NotificationIntegrationSetting) GetEventNameMapping() NotificationIntegrationSettingWebhookEventNameMapping`

GetEventNameMapping returns the EventNameMapping field if non-nil, zero value otherwise.

### GetEventNameMappingOk

`func (o *NotificationIntegrationSetting) GetEventNameMappingOk() (*NotificationIntegrationSettingWebhookEventNameMapping, bool)`

GetEventNameMappingOk returns a tuple with the EventNameMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNameMapping

`func (o *NotificationIntegrationSetting) SetEventNameMapping(v NotificationIntegrationSettingWebhookEventNameMapping)`

SetEventNameMapping sets EventNameMapping field to given value.

### HasEventNameMapping

`func (o *NotificationIntegrationSetting) HasEventNameMapping() bool`

HasEventNameMapping returns a boolean if a field has been set.

### GetInstanceName

`func (o *NotificationIntegrationSetting) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *NotificationIntegrationSetting) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *NotificationIntegrationSetting) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.


### GetUsername

`func (o *NotificationIntegrationSetting) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *NotificationIntegrationSetting) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *NotificationIntegrationSetting) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPasswordLastChars

`func (o *NotificationIntegrationSetting) GetPasswordLastChars() string`

GetPasswordLastChars returns the PasswordLastChars field if non-nil, zero value otherwise.

### GetPasswordLastCharsOk

`func (o *NotificationIntegrationSetting) GetPasswordLastCharsOk() (*string, bool)`

GetPasswordLastCharsOk returns a tuple with the PasswordLastChars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLastChars

`func (o *NotificationIntegrationSetting) SetPasswordLastChars(v string)`

SetPasswordLastChars sets PasswordLastChars field to given value.


### GetSiteUrl

`func (o *NotificationIntegrationSetting) GetSiteUrl() string`

GetSiteUrl returns the SiteUrl field if non-nil, zero value otherwise.

### GetSiteUrlOk

`func (o *NotificationIntegrationSetting) GetSiteUrlOk() (*string, bool)`

GetSiteUrlOk returns a tuple with the SiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteUrl

`func (o *NotificationIntegrationSetting) SetSiteUrl(v string)`

SetSiteUrl sets SiteUrl field to given value.


### GetEmail

`func (o *NotificationIntegrationSetting) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *NotificationIntegrationSetting) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *NotificationIntegrationSetting) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetApiTokenLastChars

`func (o *NotificationIntegrationSetting) GetApiTokenLastChars() string`

GetApiTokenLastChars returns the ApiTokenLastChars field if non-nil, zero value otherwise.

### GetApiTokenLastCharsOk

`func (o *NotificationIntegrationSetting) GetApiTokenLastCharsOk() (*string, bool)`

GetApiTokenLastCharsOk returns a tuple with the ApiTokenLastChars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiTokenLastChars

`func (o *NotificationIntegrationSetting) SetApiTokenLastChars(v string)`

SetApiTokenLastChars sets ApiTokenLastChars field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


