# UpdateNotificationIntegrationRequestSettingOneOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Url** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Authentication** | Pointer to [**UpdateNotificationIntegrationRequestSettingOneOf2Authentication**](UpdateNotificationIntegrationRequestSettingOneOf2Authentication.md) |  | [optional] 
**CustomHeaders** | Pointer to **map[string]string** |  | [optional] 
**PayloadTemplate** | Pointer to [**CreateNotificationIntegrationRequestSettingOneOf4PayloadTemplate**](CreateNotificationIntegrationRequestSettingOneOf4PayloadTemplate.md) |  | [optional] 
**EventNameMapping** | Pointer to [**NullableUpdateNotificationIntegrationRequestSettingOneOf2EventNameMapping**](UpdateNotificationIntegrationRequestSettingOneOf2EventNameMapping.md) |  | [optional] 

## Methods

### NewUpdateNotificationIntegrationRequestSettingOneOf2

`func NewUpdateNotificationIntegrationRequestSettingOneOf2(type_ string, ) *UpdateNotificationIntegrationRequestSettingOneOf2`

NewUpdateNotificationIntegrationRequestSettingOneOf2 instantiates a new UpdateNotificationIntegrationRequestSettingOneOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNotificationIntegrationRequestSettingOneOf2WithDefaults

`func NewUpdateNotificationIntegrationRequestSettingOneOf2WithDefaults() *UpdateNotificationIntegrationRequestSettingOneOf2`

NewUpdateNotificationIntegrationRequestSettingOneOf2WithDefaults instantiates a new UpdateNotificationIntegrationRequestSettingOneOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateNotificationIntegrationRequestSettingOneOf2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateNotificationIntegrationRequestSettingOneOf2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateNotificationIntegrationRequestSettingOneOf2) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *UpdateNotificationIntegrationRequestSettingOneOf2) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateNotificationIntegrationRequestSettingOneOf2) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateNotificationIntegrationRequestSettingOneOf2) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateNotificationIntegrationRequestSettingOneOf2) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetMethod

`func (o *UpdateNotificationIntegrationRequestSettingOneOf2) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *UpdateNotificationIntegrationRequestSettingOneOf2) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *UpdateNotificationIntegrationRequestSettingOneOf2) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *UpdateNotificationIntegrationRequestSettingOneOf2) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetAuthentication

`func (o *UpdateNotificationIntegrationRequestSettingOneOf2) GetAuthentication() UpdateNotificationIntegrationRequestSettingOneOf2Authentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *UpdateNotificationIntegrationRequestSettingOneOf2) GetAuthenticationOk() (*UpdateNotificationIntegrationRequestSettingOneOf2Authentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *UpdateNotificationIntegrationRequestSettingOneOf2) SetAuthentication(v UpdateNotificationIntegrationRequestSettingOneOf2Authentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *UpdateNotificationIntegrationRequestSettingOneOf2) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetCustomHeaders

`func (o *UpdateNotificationIntegrationRequestSettingOneOf2) GetCustomHeaders() map[string]string`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *UpdateNotificationIntegrationRequestSettingOneOf2) GetCustomHeadersOk() (*map[string]string, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *UpdateNotificationIntegrationRequestSettingOneOf2) SetCustomHeaders(v map[string]string)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *UpdateNotificationIntegrationRequestSettingOneOf2) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.

### GetPayloadTemplate

`func (o *UpdateNotificationIntegrationRequestSettingOneOf2) GetPayloadTemplate() CreateNotificationIntegrationRequestSettingOneOf4PayloadTemplate`

GetPayloadTemplate returns the PayloadTemplate field if non-nil, zero value otherwise.

### GetPayloadTemplateOk

`func (o *UpdateNotificationIntegrationRequestSettingOneOf2) GetPayloadTemplateOk() (*CreateNotificationIntegrationRequestSettingOneOf4PayloadTemplate, bool)`

GetPayloadTemplateOk returns a tuple with the PayloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplate

`func (o *UpdateNotificationIntegrationRequestSettingOneOf2) SetPayloadTemplate(v CreateNotificationIntegrationRequestSettingOneOf4PayloadTemplate)`

SetPayloadTemplate sets PayloadTemplate field to given value.

### HasPayloadTemplate

`func (o *UpdateNotificationIntegrationRequestSettingOneOf2) HasPayloadTemplate() bool`

HasPayloadTemplate returns a boolean if a field has been set.

### GetEventNameMapping

`func (o *UpdateNotificationIntegrationRequestSettingOneOf2) GetEventNameMapping() UpdateNotificationIntegrationRequestSettingOneOf2EventNameMapping`

GetEventNameMapping returns the EventNameMapping field if non-nil, zero value otherwise.

### GetEventNameMappingOk

`func (o *UpdateNotificationIntegrationRequestSettingOneOf2) GetEventNameMappingOk() (*UpdateNotificationIntegrationRequestSettingOneOf2EventNameMapping, bool)`

GetEventNameMappingOk returns a tuple with the EventNameMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNameMapping

`func (o *UpdateNotificationIntegrationRequestSettingOneOf2) SetEventNameMapping(v UpdateNotificationIntegrationRequestSettingOneOf2EventNameMapping)`

SetEventNameMapping sets EventNameMapping field to given value.

### HasEventNameMapping

`func (o *UpdateNotificationIntegrationRequestSettingOneOf2) HasEventNameMapping() bool`

HasEventNameMapping returns a boolean if a field has been set.

### SetEventNameMappingNil

`func (o *UpdateNotificationIntegrationRequestSettingOneOf2) SetEventNameMappingNil(b bool)`

 SetEventNameMappingNil sets the value for EventNameMapping to be an explicit nil

### UnsetEventNameMapping
`func (o *UpdateNotificationIntegrationRequestSettingOneOf2) UnsetEventNameMapping()`

UnsetEventNameMapping ensures that no value is present for EventNameMapping, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


