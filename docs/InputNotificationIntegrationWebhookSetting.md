# InputNotificationIntegrationWebhookSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Url** | **string** |  | 
**Method** | **string** |  | 
**Authentication** | [**InputNotificationIntegrationWebhookSettingAuthentication**](InputNotificationIntegrationWebhookSettingAuthentication.md) |  | 
**CustomHeaders** | **map[string]string** |  | 
**PayloadTemplate** | [**InputNotificationIntegrationWebhookSettingPayloadTemplate**](InputNotificationIntegrationWebhookSettingPayloadTemplate.md) |  | 
**EventNameMapping** | Pointer to [**InputNotificationIntegrationWebhookSettingEventNameMapping**](InputNotificationIntegrationWebhookSettingEventNameMapping.md) |  | [optional] 

## Methods

### NewInputNotificationIntegrationWebhookSetting

`func NewInputNotificationIntegrationWebhookSetting(type_ string, url string, method string, authentication InputNotificationIntegrationWebhookSettingAuthentication, customHeaders map[string]string, payloadTemplate InputNotificationIntegrationWebhookSettingPayloadTemplate, ) *InputNotificationIntegrationWebhookSetting`

NewInputNotificationIntegrationWebhookSetting instantiates a new InputNotificationIntegrationWebhookSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputNotificationIntegrationWebhookSettingWithDefaults

`func NewInputNotificationIntegrationWebhookSettingWithDefaults() *InputNotificationIntegrationWebhookSetting`

NewInputNotificationIntegrationWebhookSettingWithDefaults instantiates a new InputNotificationIntegrationWebhookSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputNotificationIntegrationWebhookSetting) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputNotificationIntegrationWebhookSetting) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputNotificationIntegrationWebhookSetting) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *InputNotificationIntegrationWebhookSetting) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InputNotificationIntegrationWebhookSetting) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InputNotificationIntegrationWebhookSetting) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMethod

`func (o *InputNotificationIntegrationWebhookSetting) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *InputNotificationIntegrationWebhookSetting) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *InputNotificationIntegrationWebhookSetting) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetAuthentication

`func (o *InputNotificationIntegrationWebhookSetting) GetAuthentication() InputNotificationIntegrationWebhookSettingAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *InputNotificationIntegrationWebhookSetting) GetAuthenticationOk() (*InputNotificationIntegrationWebhookSettingAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *InputNotificationIntegrationWebhookSetting) SetAuthentication(v InputNotificationIntegrationWebhookSettingAuthentication)`

SetAuthentication sets Authentication field to given value.


### GetCustomHeaders

`func (o *InputNotificationIntegrationWebhookSetting) GetCustomHeaders() map[string]string`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *InputNotificationIntegrationWebhookSetting) GetCustomHeadersOk() (*map[string]string, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *InputNotificationIntegrationWebhookSetting) SetCustomHeaders(v map[string]string)`

SetCustomHeaders sets CustomHeaders field to given value.


### GetPayloadTemplate

`func (o *InputNotificationIntegrationWebhookSetting) GetPayloadTemplate() InputNotificationIntegrationWebhookSettingPayloadTemplate`

GetPayloadTemplate returns the PayloadTemplate field if non-nil, zero value otherwise.

### GetPayloadTemplateOk

`func (o *InputNotificationIntegrationWebhookSetting) GetPayloadTemplateOk() (*InputNotificationIntegrationWebhookSettingPayloadTemplate, bool)`

GetPayloadTemplateOk returns a tuple with the PayloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplate

`func (o *InputNotificationIntegrationWebhookSetting) SetPayloadTemplate(v InputNotificationIntegrationWebhookSettingPayloadTemplate)`

SetPayloadTemplate sets PayloadTemplate field to given value.


### GetEventNameMapping

`func (o *InputNotificationIntegrationWebhookSetting) GetEventNameMapping() InputNotificationIntegrationWebhookSettingEventNameMapping`

GetEventNameMapping returns the EventNameMapping field if non-nil, zero value otherwise.

### GetEventNameMappingOk

`func (o *InputNotificationIntegrationWebhookSetting) GetEventNameMappingOk() (*InputNotificationIntegrationWebhookSettingEventNameMapping, bool)`

GetEventNameMappingOk returns a tuple with the EventNameMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNameMapping

`func (o *InputNotificationIntegrationWebhookSetting) SetEventNameMapping(v InputNotificationIntegrationWebhookSettingEventNameMapping)`

SetEventNameMapping sets EventNameMapping field to given value.

### HasEventNameMapping

`func (o *InputNotificationIntegrationWebhookSetting) HasEventNameMapping() bool`

HasEventNameMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


