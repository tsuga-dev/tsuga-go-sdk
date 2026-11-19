# InputNotificationIntegrationWebhookSettingUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Url** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Authentication** | Pointer to [**InputNotificationIntegrationWebhookSettingUpdateAuthentication**](InputNotificationIntegrationWebhookSettingUpdateAuthentication.md) |  | [optional] 
**CustomHeaders** | Pointer to **map[string]string** |  | [optional] 
**PayloadTemplate** | Pointer to [**InputNotificationIntegrationWebhookSettingUpdatePayloadTemplate**](InputNotificationIntegrationWebhookSettingUpdatePayloadTemplate.md) |  | [optional] 
**EventNameMapping** | Pointer to [**NullableInputNotificationIntegrationWebhookSettingUpdateEventNameMapping**](InputNotificationIntegrationWebhookSettingUpdateEventNameMapping.md) |  | [optional] 

## Methods

### NewInputNotificationIntegrationWebhookSettingUpdate

`func NewInputNotificationIntegrationWebhookSettingUpdate(type_ string, ) *InputNotificationIntegrationWebhookSettingUpdate`

NewInputNotificationIntegrationWebhookSettingUpdate instantiates a new InputNotificationIntegrationWebhookSettingUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputNotificationIntegrationWebhookSettingUpdateWithDefaults

`func NewInputNotificationIntegrationWebhookSettingUpdateWithDefaults() *InputNotificationIntegrationWebhookSettingUpdate`

NewInputNotificationIntegrationWebhookSettingUpdateWithDefaults instantiates a new InputNotificationIntegrationWebhookSettingUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputNotificationIntegrationWebhookSettingUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputNotificationIntegrationWebhookSettingUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputNotificationIntegrationWebhookSettingUpdate) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *InputNotificationIntegrationWebhookSettingUpdate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InputNotificationIntegrationWebhookSettingUpdate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InputNotificationIntegrationWebhookSettingUpdate) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InputNotificationIntegrationWebhookSettingUpdate) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetMethod

`func (o *InputNotificationIntegrationWebhookSettingUpdate) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *InputNotificationIntegrationWebhookSettingUpdate) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *InputNotificationIntegrationWebhookSettingUpdate) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *InputNotificationIntegrationWebhookSettingUpdate) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetAuthentication

`func (o *InputNotificationIntegrationWebhookSettingUpdate) GetAuthentication() InputNotificationIntegrationWebhookSettingUpdateAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *InputNotificationIntegrationWebhookSettingUpdate) GetAuthenticationOk() (*InputNotificationIntegrationWebhookSettingUpdateAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *InputNotificationIntegrationWebhookSettingUpdate) SetAuthentication(v InputNotificationIntegrationWebhookSettingUpdateAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *InputNotificationIntegrationWebhookSettingUpdate) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetCustomHeaders

`func (o *InputNotificationIntegrationWebhookSettingUpdate) GetCustomHeaders() map[string]string`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *InputNotificationIntegrationWebhookSettingUpdate) GetCustomHeadersOk() (*map[string]string, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *InputNotificationIntegrationWebhookSettingUpdate) SetCustomHeaders(v map[string]string)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *InputNotificationIntegrationWebhookSettingUpdate) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.

### GetPayloadTemplate

`func (o *InputNotificationIntegrationWebhookSettingUpdate) GetPayloadTemplate() InputNotificationIntegrationWebhookSettingUpdatePayloadTemplate`

GetPayloadTemplate returns the PayloadTemplate field if non-nil, zero value otherwise.

### GetPayloadTemplateOk

`func (o *InputNotificationIntegrationWebhookSettingUpdate) GetPayloadTemplateOk() (*InputNotificationIntegrationWebhookSettingUpdatePayloadTemplate, bool)`

GetPayloadTemplateOk returns a tuple with the PayloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplate

`func (o *InputNotificationIntegrationWebhookSettingUpdate) SetPayloadTemplate(v InputNotificationIntegrationWebhookSettingUpdatePayloadTemplate)`

SetPayloadTemplate sets PayloadTemplate field to given value.

### HasPayloadTemplate

`func (o *InputNotificationIntegrationWebhookSettingUpdate) HasPayloadTemplate() bool`

HasPayloadTemplate returns a boolean if a field has been set.

### GetEventNameMapping

`func (o *InputNotificationIntegrationWebhookSettingUpdate) GetEventNameMapping() InputNotificationIntegrationWebhookSettingUpdateEventNameMapping`

GetEventNameMapping returns the EventNameMapping field if non-nil, zero value otherwise.

### GetEventNameMappingOk

`func (o *InputNotificationIntegrationWebhookSettingUpdate) GetEventNameMappingOk() (*InputNotificationIntegrationWebhookSettingUpdateEventNameMapping, bool)`

GetEventNameMappingOk returns a tuple with the EventNameMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNameMapping

`func (o *InputNotificationIntegrationWebhookSettingUpdate) SetEventNameMapping(v InputNotificationIntegrationWebhookSettingUpdateEventNameMapping)`

SetEventNameMapping sets EventNameMapping field to given value.

### HasEventNameMapping

`func (o *InputNotificationIntegrationWebhookSettingUpdate) HasEventNameMapping() bool`

HasEventNameMapping returns a boolean if a field has been set.

### SetEventNameMappingNil

`func (o *InputNotificationIntegrationWebhookSettingUpdate) SetEventNameMappingNil(b bool)`

 SetEventNameMappingNil sets the value for EventNameMapping to be an explicit nil

### UnsetEventNameMapping
`func (o *InputNotificationIntegrationWebhookSettingUpdate) UnsetEventNameMapping()`

UnsetEventNameMapping ensures that no value is present for EventNameMapping, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


