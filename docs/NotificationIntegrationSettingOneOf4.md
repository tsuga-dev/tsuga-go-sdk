# NotificationIntegrationSettingOneOf4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**UrlLastChars** | **string** |  | 
**Method** | **string** |  | 
**Authentication** | [**NotificationIntegrationSettingOneOf4Authentication**](NotificationIntegrationSettingOneOf4Authentication.md) |  | 
**CustomHeaders** | Pointer to **map[string]string** |  | [optional] 
**PayloadTemplate** | [**CreateNotificationIntegrationRequestSettingOneOf4PayloadTemplate**](CreateNotificationIntegrationRequestSettingOneOf4PayloadTemplate.md) |  | 
**EventNameMapping** | Pointer to [**NotificationIntegrationSettingOneOf4EventNameMapping**](NotificationIntegrationSettingOneOf4EventNameMapping.md) |  | [optional] 

## Methods

### NewNotificationIntegrationSettingOneOf4

`func NewNotificationIntegrationSettingOneOf4(type_ string, urlLastChars string, method string, authentication NotificationIntegrationSettingOneOf4Authentication, payloadTemplate CreateNotificationIntegrationRequestSettingOneOf4PayloadTemplate, ) *NotificationIntegrationSettingOneOf4`

NewNotificationIntegrationSettingOneOf4 instantiates a new NotificationIntegrationSettingOneOf4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationIntegrationSettingOneOf4WithDefaults

`func NewNotificationIntegrationSettingOneOf4WithDefaults() *NotificationIntegrationSettingOneOf4`

NewNotificationIntegrationSettingOneOf4WithDefaults instantiates a new NotificationIntegrationSettingOneOf4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NotificationIntegrationSettingOneOf4) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationIntegrationSettingOneOf4) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationIntegrationSettingOneOf4) SetType(v string)`

SetType sets Type field to given value.


### GetUrlLastChars

`func (o *NotificationIntegrationSettingOneOf4) GetUrlLastChars() string`

GetUrlLastChars returns the UrlLastChars field if non-nil, zero value otherwise.

### GetUrlLastCharsOk

`func (o *NotificationIntegrationSettingOneOf4) GetUrlLastCharsOk() (*string, bool)`

GetUrlLastCharsOk returns a tuple with the UrlLastChars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlLastChars

`func (o *NotificationIntegrationSettingOneOf4) SetUrlLastChars(v string)`

SetUrlLastChars sets UrlLastChars field to given value.


### GetMethod

`func (o *NotificationIntegrationSettingOneOf4) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *NotificationIntegrationSettingOneOf4) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *NotificationIntegrationSettingOneOf4) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetAuthentication

`func (o *NotificationIntegrationSettingOneOf4) GetAuthentication() NotificationIntegrationSettingOneOf4Authentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *NotificationIntegrationSettingOneOf4) GetAuthenticationOk() (*NotificationIntegrationSettingOneOf4Authentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *NotificationIntegrationSettingOneOf4) SetAuthentication(v NotificationIntegrationSettingOneOf4Authentication)`

SetAuthentication sets Authentication field to given value.


### GetCustomHeaders

`func (o *NotificationIntegrationSettingOneOf4) GetCustomHeaders() map[string]string`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *NotificationIntegrationSettingOneOf4) GetCustomHeadersOk() (*map[string]string, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *NotificationIntegrationSettingOneOf4) SetCustomHeaders(v map[string]string)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *NotificationIntegrationSettingOneOf4) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.

### GetPayloadTemplate

`func (o *NotificationIntegrationSettingOneOf4) GetPayloadTemplate() CreateNotificationIntegrationRequestSettingOneOf4PayloadTemplate`

GetPayloadTemplate returns the PayloadTemplate field if non-nil, zero value otherwise.

### GetPayloadTemplateOk

`func (o *NotificationIntegrationSettingOneOf4) GetPayloadTemplateOk() (*CreateNotificationIntegrationRequestSettingOneOf4PayloadTemplate, bool)`

GetPayloadTemplateOk returns a tuple with the PayloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplate

`func (o *NotificationIntegrationSettingOneOf4) SetPayloadTemplate(v CreateNotificationIntegrationRequestSettingOneOf4PayloadTemplate)`

SetPayloadTemplate sets PayloadTemplate field to given value.


### GetEventNameMapping

`func (o *NotificationIntegrationSettingOneOf4) GetEventNameMapping() NotificationIntegrationSettingOneOf4EventNameMapping`

GetEventNameMapping returns the EventNameMapping field if non-nil, zero value otherwise.

### GetEventNameMappingOk

`func (o *NotificationIntegrationSettingOneOf4) GetEventNameMappingOk() (*NotificationIntegrationSettingOneOf4EventNameMapping, bool)`

GetEventNameMappingOk returns a tuple with the EventNameMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNameMapping

`func (o *NotificationIntegrationSettingOneOf4) SetEventNameMapping(v NotificationIntegrationSettingOneOf4EventNameMapping)`

SetEventNameMapping sets EventNameMapping field to given value.

### HasEventNameMapping

`func (o *NotificationIntegrationSettingOneOf4) HasEventNameMapping() bool`

HasEventNameMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


