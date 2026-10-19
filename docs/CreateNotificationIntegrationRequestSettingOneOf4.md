# CreateNotificationIntegrationRequestSettingOneOf4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Url** | **string** |  | 
**Method** | **string** |  | 
**Authentication** | [**CreateNotificationIntegrationRequestSettingOneOf4Authentication**](CreateNotificationIntegrationRequestSettingOneOf4Authentication.md) |  | 
**CustomHeaders** | **map[string]string** |  | 
**PayloadTemplate** | [**CreateNotificationIntegrationRequestSettingOneOf4PayloadTemplate**](CreateNotificationIntegrationRequestSettingOneOf4PayloadTemplate.md) |  | 
**EventNameMapping** | Pointer to [**CreateNotificationIntegrationRequestSettingOneOf4EventNameMapping**](CreateNotificationIntegrationRequestSettingOneOf4EventNameMapping.md) |  | [optional] 

## Methods

### NewCreateNotificationIntegrationRequestSettingOneOf4

`func NewCreateNotificationIntegrationRequestSettingOneOf4(type_ string, url string, method string, authentication CreateNotificationIntegrationRequestSettingOneOf4Authentication, customHeaders map[string]*string, payloadTemplate CreateNotificationIntegrationRequestSettingOneOf4PayloadTemplate, ) *CreateNotificationIntegrationRequestSettingOneOf4`

NewCreateNotificationIntegrationRequestSettingOneOf4 instantiates a new CreateNotificationIntegrationRequestSettingOneOf4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNotificationIntegrationRequestSettingOneOf4WithDefaults

`func NewCreateNotificationIntegrationRequestSettingOneOf4WithDefaults() *CreateNotificationIntegrationRequestSettingOneOf4`

NewCreateNotificationIntegrationRequestSettingOneOf4WithDefaults instantiates a new CreateNotificationIntegrationRequestSettingOneOf4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateNotificationIntegrationRequestSettingOneOf4) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateNotificationIntegrationRequestSettingOneOf4) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateNotificationIntegrationRequestSettingOneOf4) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *CreateNotificationIntegrationRequestSettingOneOf4) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateNotificationIntegrationRequestSettingOneOf4) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateNotificationIntegrationRequestSettingOneOf4) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMethod

`func (o *CreateNotificationIntegrationRequestSettingOneOf4) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *CreateNotificationIntegrationRequestSettingOneOf4) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *CreateNotificationIntegrationRequestSettingOneOf4) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetAuthentication

`func (o *CreateNotificationIntegrationRequestSettingOneOf4) GetAuthentication() CreateNotificationIntegrationRequestSettingOneOf4Authentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *CreateNotificationIntegrationRequestSettingOneOf4) GetAuthenticationOk() (*CreateNotificationIntegrationRequestSettingOneOf4Authentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *CreateNotificationIntegrationRequestSettingOneOf4) SetAuthentication(v CreateNotificationIntegrationRequestSettingOneOf4Authentication)`

SetAuthentication sets Authentication field to given value.


### GetCustomHeaders

`func (o *CreateNotificationIntegrationRequestSettingOneOf4) GetCustomHeaders() map[string]*string`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *CreateNotificationIntegrationRequestSettingOneOf4) GetCustomHeadersOk() (*map[string]*string, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *CreateNotificationIntegrationRequestSettingOneOf4) SetCustomHeaders(v map[string]*string)`

SetCustomHeaders sets CustomHeaders field to given value.


### GetPayloadTemplate

`func (o *CreateNotificationIntegrationRequestSettingOneOf4) GetPayloadTemplate() CreateNotificationIntegrationRequestSettingOneOf4PayloadTemplate`

GetPayloadTemplate returns the PayloadTemplate field if non-nil, zero value otherwise.

### GetPayloadTemplateOk

`func (o *CreateNotificationIntegrationRequestSettingOneOf4) GetPayloadTemplateOk() (*CreateNotificationIntegrationRequestSettingOneOf4PayloadTemplate, bool)`

GetPayloadTemplateOk returns a tuple with the PayloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplate

`func (o *CreateNotificationIntegrationRequestSettingOneOf4) SetPayloadTemplate(v CreateNotificationIntegrationRequestSettingOneOf4PayloadTemplate)`

SetPayloadTemplate sets PayloadTemplate field to given value.


### GetEventNameMapping

`func (o *CreateNotificationIntegrationRequestSettingOneOf4) GetEventNameMapping() CreateNotificationIntegrationRequestSettingOneOf4EventNameMapping`

GetEventNameMapping returns the EventNameMapping field if non-nil, zero value otherwise.

### GetEventNameMappingOk

`func (o *CreateNotificationIntegrationRequestSettingOneOf4) GetEventNameMappingOk() (*CreateNotificationIntegrationRequestSettingOneOf4EventNameMapping, bool)`

GetEventNameMappingOk returns a tuple with the EventNameMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNameMapping

`func (o *CreateNotificationIntegrationRequestSettingOneOf4) SetEventNameMapping(v CreateNotificationIntegrationRequestSettingOneOf4EventNameMapping)`

SetEventNameMapping sets EventNameMapping field to given value.

### HasEventNameMapping

`func (o *CreateNotificationIntegrationRequestSettingOneOf4) HasEventNameMapping() bool`

HasEventNameMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


