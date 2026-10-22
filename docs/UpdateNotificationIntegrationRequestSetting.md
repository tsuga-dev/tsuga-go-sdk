# UpdateNotificationIntegrationRequestSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**AlertSourceConfigId** | **string** |  | 
**Token** | **string** |  | 
**Url** | **string** |  | 
**IntegrationKey** | **string** |  | 
**WebhookUrl** | **string** |  | 
**Method** | Pointer to **string** |  | [optional] 
**Authentication** | Pointer to [**UpdateNotificationIntegrationRequestSettingOneOf2Authentication**](UpdateNotificationIntegrationRequestSettingOneOf2Authentication.md) |  | [optional] 
**CustomHeaders** | Pointer to **map[string]string** |  | [optional] 
**PayloadTemplate** | Pointer to [**CreateNotificationIntegrationRequestSettingOneOf4PayloadTemplate**](CreateNotificationIntegrationRequestSettingOneOf4PayloadTemplate.md) |  | [optional] 
**EventNameMapping** | Pointer to [**NullableUpdateNotificationIntegrationRequestSettingOneOf2EventNameMapping**](UpdateNotificationIntegrationRequestSettingOneOf2EventNameMapping.md) |  | [optional] 
**InstanceName** | Pointer to **NullableString** | Folder filter for dashboard query results. Optional; pass a folder ID to return dashboards in that folder, or null to return only dashboards that are not in any folder. | [optional] 
**Username** | Pointer to **NullableString** | Folder filter for dashboard query results. Optional; pass a folder ID to return dashboards in that folder, or null to return only dashboards that are not in any folder. | [optional] 
**Password** | Pointer to **NullableString** | Folder filter for dashboard query results. Optional; pass a folder ID to return dashboards in that folder, or null to return only dashboards that are not in any folder. | [optional] 
**SiteUrl** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **NullableString** | Folder filter for dashboard query results. Optional; pass a folder ID to return dashboards in that folder, or null to return only dashboards that are not in any folder. | [optional] 
**ApiToken** | Pointer to **NullableString** | Folder filter for dashboard query results. Optional; pass a folder ID to return dashboards in that folder, or null to return only dashboards that are not in any folder. | [optional] 

## Methods

### NewUpdateNotificationIntegrationRequestSetting

`func NewUpdateNotificationIntegrationRequestSetting(type_ string, alertSourceConfigId string, token string, url string, integrationKey string, webhookUrl string, ) *UpdateNotificationIntegrationRequestSetting`

NewUpdateNotificationIntegrationRequestSetting instantiates a new UpdateNotificationIntegrationRequestSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNotificationIntegrationRequestSettingWithDefaults

`func NewUpdateNotificationIntegrationRequestSettingWithDefaults() *UpdateNotificationIntegrationRequestSetting`

NewUpdateNotificationIntegrationRequestSettingWithDefaults instantiates a new UpdateNotificationIntegrationRequestSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateNotificationIntegrationRequestSetting) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateNotificationIntegrationRequestSetting) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateNotificationIntegrationRequestSetting) SetType(v string)`

SetType sets Type field to given value.


### GetAlertSourceConfigId

`func (o *UpdateNotificationIntegrationRequestSetting) GetAlertSourceConfigId() string`

GetAlertSourceConfigId returns the AlertSourceConfigId field if non-nil, zero value otherwise.

### GetAlertSourceConfigIdOk

`func (o *UpdateNotificationIntegrationRequestSetting) GetAlertSourceConfigIdOk() (*string, bool)`

GetAlertSourceConfigIdOk returns a tuple with the AlertSourceConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertSourceConfigId

`func (o *UpdateNotificationIntegrationRequestSetting) SetAlertSourceConfigId(v string)`

SetAlertSourceConfigId sets AlertSourceConfigId field to given value.


### GetToken

`func (o *UpdateNotificationIntegrationRequestSetting) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateNotificationIntegrationRequestSetting) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateNotificationIntegrationRequestSetting) SetToken(v string)`

SetToken sets Token field to given value.


### GetUrl

`func (o *UpdateNotificationIntegrationRequestSetting) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateNotificationIntegrationRequestSetting) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateNotificationIntegrationRequestSetting) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetIntegrationKey

`func (o *UpdateNotificationIntegrationRequestSetting) GetIntegrationKey() string`

GetIntegrationKey returns the IntegrationKey field if non-nil, zero value otherwise.

### GetIntegrationKeyOk

`func (o *UpdateNotificationIntegrationRequestSetting) GetIntegrationKeyOk() (*string, bool)`

GetIntegrationKeyOk returns a tuple with the IntegrationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationKey

`func (o *UpdateNotificationIntegrationRequestSetting) SetIntegrationKey(v string)`

SetIntegrationKey sets IntegrationKey field to given value.


### GetWebhookUrl

`func (o *UpdateNotificationIntegrationRequestSetting) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *UpdateNotificationIntegrationRequestSetting) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *UpdateNotificationIntegrationRequestSetting) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.


### GetMethod

`func (o *UpdateNotificationIntegrationRequestSetting) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *UpdateNotificationIntegrationRequestSetting) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *UpdateNotificationIntegrationRequestSetting) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *UpdateNotificationIntegrationRequestSetting) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetAuthentication

`func (o *UpdateNotificationIntegrationRequestSetting) GetAuthentication() UpdateNotificationIntegrationRequestSettingOneOf2Authentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *UpdateNotificationIntegrationRequestSetting) GetAuthenticationOk() (*UpdateNotificationIntegrationRequestSettingOneOf2Authentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *UpdateNotificationIntegrationRequestSetting) SetAuthentication(v UpdateNotificationIntegrationRequestSettingOneOf2Authentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *UpdateNotificationIntegrationRequestSetting) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetCustomHeaders

`func (o *UpdateNotificationIntegrationRequestSetting) GetCustomHeaders() map[string]string`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *UpdateNotificationIntegrationRequestSetting) GetCustomHeadersOk() (*map[string]string, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *UpdateNotificationIntegrationRequestSetting) SetCustomHeaders(v map[string]string)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *UpdateNotificationIntegrationRequestSetting) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.

### GetPayloadTemplate

`func (o *UpdateNotificationIntegrationRequestSetting) GetPayloadTemplate() CreateNotificationIntegrationRequestSettingOneOf4PayloadTemplate`

GetPayloadTemplate returns the PayloadTemplate field if non-nil, zero value otherwise.

### GetPayloadTemplateOk

`func (o *UpdateNotificationIntegrationRequestSetting) GetPayloadTemplateOk() (*CreateNotificationIntegrationRequestSettingOneOf4PayloadTemplate, bool)`

GetPayloadTemplateOk returns a tuple with the PayloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplate

`func (o *UpdateNotificationIntegrationRequestSetting) SetPayloadTemplate(v CreateNotificationIntegrationRequestSettingOneOf4PayloadTemplate)`

SetPayloadTemplate sets PayloadTemplate field to given value.

### HasPayloadTemplate

`func (o *UpdateNotificationIntegrationRequestSetting) HasPayloadTemplate() bool`

HasPayloadTemplate returns a boolean if a field has been set.

### GetEventNameMapping

`func (o *UpdateNotificationIntegrationRequestSetting) GetEventNameMapping() UpdateNotificationIntegrationRequestSettingOneOf2EventNameMapping`

GetEventNameMapping returns the EventNameMapping field if non-nil, zero value otherwise.

### GetEventNameMappingOk

`func (o *UpdateNotificationIntegrationRequestSetting) GetEventNameMappingOk() (*UpdateNotificationIntegrationRequestSettingOneOf2EventNameMapping, bool)`

GetEventNameMappingOk returns a tuple with the EventNameMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNameMapping

`func (o *UpdateNotificationIntegrationRequestSetting) SetEventNameMapping(v UpdateNotificationIntegrationRequestSettingOneOf2EventNameMapping)`

SetEventNameMapping sets EventNameMapping field to given value.

### HasEventNameMapping

`func (o *UpdateNotificationIntegrationRequestSetting) HasEventNameMapping() bool`

HasEventNameMapping returns a boolean if a field has been set.

### SetEventNameMappingNil

`func (o *UpdateNotificationIntegrationRequestSetting) SetEventNameMappingNil(b bool)`

 SetEventNameMappingNil sets the value for EventNameMapping to be an explicit nil

### UnsetEventNameMapping
`func (o *UpdateNotificationIntegrationRequestSetting) UnsetEventNameMapping()`

UnsetEventNameMapping ensures that no value is present for EventNameMapping, not even an explicit nil
### GetInstanceName

`func (o *UpdateNotificationIntegrationRequestSetting) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *UpdateNotificationIntegrationRequestSetting) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *UpdateNotificationIntegrationRequestSetting) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *UpdateNotificationIntegrationRequestSetting) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### SetInstanceNameNil

`func (o *UpdateNotificationIntegrationRequestSetting) SetInstanceNameNil(b bool)`

 SetInstanceNameNil sets the value for InstanceName to be an explicit nil

### UnsetInstanceName
`func (o *UpdateNotificationIntegrationRequestSetting) UnsetInstanceName()`

UnsetInstanceName ensures that no value is present for InstanceName, not even an explicit nil
### GetUsername

`func (o *UpdateNotificationIntegrationRequestSetting) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateNotificationIntegrationRequestSetting) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateNotificationIntegrationRequestSetting) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateNotificationIntegrationRequestSetting) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *UpdateNotificationIntegrationRequestSetting) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *UpdateNotificationIntegrationRequestSetting) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *UpdateNotificationIntegrationRequestSetting) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateNotificationIntegrationRequestSetting) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateNotificationIntegrationRequestSetting) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateNotificationIntegrationRequestSetting) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *UpdateNotificationIntegrationRequestSetting) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *UpdateNotificationIntegrationRequestSetting) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetSiteUrl

`func (o *UpdateNotificationIntegrationRequestSetting) GetSiteUrl() string`

GetSiteUrl returns the SiteUrl field if non-nil, zero value otherwise.

### GetSiteUrlOk

`func (o *UpdateNotificationIntegrationRequestSetting) GetSiteUrlOk() (*string, bool)`

GetSiteUrlOk returns a tuple with the SiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteUrl

`func (o *UpdateNotificationIntegrationRequestSetting) SetSiteUrl(v string)`

SetSiteUrl sets SiteUrl field to given value.

### HasSiteUrl

`func (o *UpdateNotificationIntegrationRequestSetting) HasSiteUrl() bool`

HasSiteUrl returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateNotificationIntegrationRequestSetting) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateNotificationIntegrationRequestSetting) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateNotificationIntegrationRequestSetting) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateNotificationIntegrationRequestSetting) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UpdateNotificationIntegrationRequestSetting) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UpdateNotificationIntegrationRequestSetting) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetApiToken

`func (o *UpdateNotificationIntegrationRequestSetting) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *UpdateNotificationIntegrationRequestSetting) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *UpdateNotificationIntegrationRequestSetting) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.

### HasApiToken

`func (o *UpdateNotificationIntegrationRequestSetting) HasApiToken() bool`

HasApiToken returns a boolean if a field has been set.

### SetApiTokenNil

`func (o *UpdateNotificationIntegrationRequestSetting) SetApiTokenNil(b bool)`

 SetApiTokenNil sets the value for ApiToken to be an explicit nil

### UnsetApiToken
`func (o *UpdateNotificationIntegrationRequestSetting) UnsetApiToken()`

UnsetApiToken ensures that no value is present for ApiToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


