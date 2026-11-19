# NotificationIntegrationSettingWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**UrlLastChars** | **string** |  | 
**Method** | **string** |  | 
**Authentication** | [**NotificationIntegrationSettingWebhookAuthentication**](NotificationIntegrationSettingWebhookAuthentication.md) |  | 
**CustomHeaders** | Pointer to **map[string]string** |  | [optional] 
**PayloadTemplate** | [**NotificationIntegrationSettingWebhookPayloadTemplate**](NotificationIntegrationSettingWebhookPayloadTemplate.md) |  | 
**EventNameMapping** | Pointer to [**NotificationIntegrationSettingWebhookEventNameMapping**](NotificationIntegrationSettingWebhookEventNameMapping.md) |  | [optional] 

## Methods

### NewNotificationIntegrationSettingWebhook

`func NewNotificationIntegrationSettingWebhook(type_ string, urlLastChars string, method string, authentication NotificationIntegrationSettingWebhookAuthentication, payloadTemplate NotificationIntegrationSettingWebhookPayloadTemplate, ) *NotificationIntegrationSettingWebhook`

NewNotificationIntegrationSettingWebhook instantiates a new NotificationIntegrationSettingWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationIntegrationSettingWebhookWithDefaults

`func NewNotificationIntegrationSettingWebhookWithDefaults() *NotificationIntegrationSettingWebhook`

NewNotificationIntegrationSettingWebhookWithDefaults instantiates a new NotificationIntegrationSettingWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NotificationIntegrationSettingWebhook) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationIntegrationSettingWebhook) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationIntegrationSettingWebhook) SetType(v string)`

SetType sets Type field to given value.


### GetUrlLastChars

`func (o *NotificationIntegrationSettingWebhook) GetUrlLastChars() string`

GetUrlLastChars returns the UrlLastChars field if non-nil, zero value otherwise.

### GetUrlLastCharsOk

`func (o *NotificationIntegrationSettingWebhook) GetUrlLastCharsOk() (*string, bool)`

GetUrlLastCharsOk returns a tuple with the UrlLastChars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlLastChars

`func (o *NotificationIntegrationSettingWebhook) SetUrlLastChars(v string)`

SetUrlLastChars sets UrlLastChars field to given value.


### GetMethod

`func (o *NotificationIntegrationSettingWebhook) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *NotificationIntegrationSettingWebhook) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *NotificationIntegrationSettingWebhook) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetAuthentication

`func (o *NotificationIntegrationSettingWebhook) GetAuthentication() NotificationIntegrationSettingWebhookAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *NotificationIntegrationSettingWebhook) GetAuthenticationOk() (*NotificationIntegrationSettingWebhookAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *NotificationIntegrationSettingWebhook) SetAuthentication(v NotificationIntegrationSettingWebhookAuthentication)`

SetAuthentication sets Authentication field to given value.


### GetCustomHeaders

`func (o *NotificationIntegrationSettingWebhook) GetCustomHeaders() map[string]string`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *NotificationIntegrationSettingWebhook) GetCustomHeadersOk() (*map[string]string, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *NotificationIntegrationSettingWebhook) SetCustomHeaders(v map[string]string)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *NotificationIntegrationSettingWebhook) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.

### GetPayloadTemplate

`func (o *NotificationIntegrationSettingWebhook) GetPayloadTemplate() NotificationIntegrationSettingWebhookPayloadTemplate`

GetPayloadTemplate returns the PayloadTemplate field if non-nil, zero value otherwise.

### GetPayloadTemplateOk

`func (o *NotificationIntegrationSettingWebhook) GetPayloadTemplateOk() (*NotificationIntegrationSettingWebhookPayloadTemplate, bool)`

GetPayloadTemplateOk returns a tuple with the PayloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplate

`func (o *NotificationIntegrationSettingWebhook) SetPayloadTemplate(v NotificationIntegrationSettingWebhookPayloadTemplate)`

SetPayloadTemplate sets PayloadTemplate field to given value.


### GetEventNameMapping

`func (o *NotificationIntegrationSettingWebhook) GetEventNameMapping() NotificationIntegrationSettingWebhookEventNameMapping`

GetEventNameMapping returns the EventNameMapping field if non-nil, zero value otherwise.

### GetEventNameMappingOk

`func (o *NotificationIntegrationSettingWebhook) GetEventNameMappingOk() (*NotificationIntegrationSettingWebhookEventNameMapping, bool)`

GetEventNameMappingOk returns a tuple with the EventNameMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNameMapping

`func (o *NotificationIntegrationSettingWebhook) SetEventNameMapping(v NotificationIntegrationSettingWebhookEventNameMapping)`

SetEventNameMapping sets EventNameMapping field to given value.

### HasEventNameMapping

`func (o *NotificationIntegrationSettingWebhook) HasEventNameMapping() bool`

HasEventNameMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


