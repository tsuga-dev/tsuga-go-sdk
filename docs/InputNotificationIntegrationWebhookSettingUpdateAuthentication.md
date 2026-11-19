# InputNotificationIntegrationWebhookSettingUpdateAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Authenticate with a bearer token in the &#x60;Authorization&#x60; header. | 
**Token** | Pointer to **string** | Bearer token sent in the &#x60;Authorization&#x60; header. Omit to keep the stored value; required when switching from another authentication type. | [optional] 
**Username** | Pointer to **string** | Basic authentication username. Omit to keep the stored value; required when switching from another authentication type. | [optional] 
**Password** | Pointer to **string** | Basic authentication password. Omit to keep the stored value; required when switching from another authentication type. | [optional] 
**Key** | Pointer to **string** | Name of the HTTP header sent with each request. Omit to keep the stored value; required when switching from another authentication type. | [optional] 
**Value** | Pointer to **string** | Value of the HTTP header. Omit to keep the stored value; required when switching from another authentication type. | [optional] 

## Methods

### NewInputNotificationIntegrationWebhookSettingUpdateAuthentication

`func NewInputNotificationIntegrationWebhookSettingUpdateAuthentication(type_ string, ) *InputNotificationIntegrationWebhookSettingUpdateAuthentication`

NewInputNotificationIntegrationWebhookSettingUpdateAuthentication instantiates a new InputNotificationIntegrationWebhookSettingUpdateAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputNotificationIntegrationWebhookSettingUpdateAuthenticationWithDefaults

`func NewInputNotificationIntegrationWebhookSettingUpdateAuthenticationWithDefaults() *InputNotificationIntegrationWebhookSettingUpdateAuthentication`

NewInputNotificationIntegrationWebhookSettingUpdateAuthenticationWithDefaults instantiates a new InputNotificationIntegrationWebhookSettingUpdateAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputNotificationIntegrationWebhookSettingUpdateAuthentication) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputNotificationIntegrationWebhookSettingUpdateAuthentication) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputNotificationIntegrationWebhookSettingUpdateAuthentication) SetType(v string)`

SetType sets Type field to given value.


### GetToken

`func (o *InputNotificationIntegrationWebhookSettingUpdateAuthentication) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *InputNotificationIntegrationWebhookSettingUpdateAuthentication) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *InputNotificationIntegrationWebhookSettingUpdateAuthentication) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *InputNotificationIntegrationWebhookSettingUpdateAuthentication) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUsername

`func (o *InputNotificationIntegrationWebhookSettingUpdateAuthentication) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *InputNotificationIntegrationWebhookSettingUpdateAuthentication) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *InputNotificationIntegrationWebhookSettingUpdateAuthentication) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *InputNotificationIntegrationWebhookSettingUpdateAuthentication) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *InputNotificationIntegrationWebhookSettingUpdateAuthentication) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InputNotificationIntegrationWebhookSettingUpdateAuthentication) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InputNotificationIntegrationWebhookSettingUpdateAuthentication) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *InputNotificationIntegrationWebhookSettingUpdateAuthentication) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetKey

`func (o *InputNotificationIntegrationWebhookSettingUpdateAuthentication) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *InputNotificationIntegrationWebhookSettingUpdateAuthentication) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *InputNotificationIntegrationWebhookSettingUpdateAuthentication) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *InputNotificationIntegrationWebhookSettingUpdateAuthentication) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *InputNotificationIntegrationWebhookSettingUpdateAuthentication) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InputNotificationIntegrationWebhookSettingUpdateAuthentication) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InputNotificationIntegrationWebhookSettingUpdateAuthentication) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *InputNotificationIntegrationWebhookSettingUpdateAuthentication) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


