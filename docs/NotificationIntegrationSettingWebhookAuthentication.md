# NotificationIntegrationSettingWebhookAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Authenticate with a bearer token in the &#x60;Authorization&#x60; header. | 
**TokenLastChars** | **string** | Redacted bearer token: only its last 4 characters, masked as &#x60;****1234&#x60;. | 
**Username** | **string** | Basic authentication username. | 
**PasswordLastChars** | **string** | Redacted Basic authentication password: only its last 4 characters, masked as &#x60;****1234&#x60;. | 
**Key** | **string** | Name of the HTTP header sent with each request. | 
**ValueLastChars** | **string** | Redacted header value: only its last 4 characters, masked as &#x60;****1234&#x60;. | 

## Methods

### NewNotificationIntegrationSettingWebhookAuthentication

`func NewNotificationIntegrationSettingWebhookAuthentication(type_ string, tokenLastChars string, username string, passwordLastChars string, key string, valueLastChars string, ) *NotificationIntegrationSettingWebhookAuthentication`

NewNotificationIntegrationSettingWebhookAuthentication instantiates a new NotificationIntegrationSettingWebhookAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationIntegrationSettingWebhookAuthenticationWithDefaults

`func NewNotificationIntegrationSettingWebhookAuthenticationWithDefaults() *NotificationIntegrationSettingWebhookAuthentication`

NewNotificationIntegrationSettingWebhookAuthenticationWithDefaults instantiates a new NotificationIntegrationSettingWebhookAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NotificationIntegrationSettingWebhookAuthentication) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationIntegrationSettingWebhookAuthentication) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationIntegrationSettingWebhookAuthentication) SetType(v string)`

SetType sets Type field to given value.


### GetTokenLastChars

`func (o *NotificationIntegrationSettingWebhookAuthentication) GetTokenLastChars() string`

GetTokenLastChars returns the TokenLastChars field if non-nil, zero value otherwise.

### GetTokenLastCharsOk

`func (o *NotificationIntegrationSettingWebhookAuthentication) GetTokenLastCharsOk() (*string, bool)`

GetTokenLastCharsOk returns a tuple with the TokenLastChars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenLastChars

`func (o *NotificationIntegrationSettingWebhookAuthentication) SetTokenLastChars(v string)`

SetTokenLastChars sets TokenLastChars field to given value.


### GetUsername

`func (o *NotificationIntegrationSettingWebhookAuthentication) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *NotificationIntegrationSettingWebhookAuthentication) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *NotificationIntegrationSettingWebhookAuthentication) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPasswordLastChars

`func (o *NotificationIntegrationSettingWebhookAuthentication) GetPasswordLastChars() string`

GetPasswordLastChars returns the PasswordLastChars field if non-nil, zero value otherwise.

### GetPasswordLastCharsOk

`func (o *NotificationIntegrationSettingWebhookAuthentication) GetPasswordLastCharsOk() (*string, bool)`

GetPasswordLastCharsOk returns a tuple with the PasswordLastChars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLastChars

`func (o *NotificationIntegrationSettingWebhookAuthentication) SetPasswordLastChars(v string)`

SetPasswordLastChars sets PasswordLastChars field to given value.


### GetKey

`func (o *NotificationIntegrationSettingWebhookAuthentication) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *NotificationIntegrationSettingWebhookAuthentication) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *NotificationIntegrationSettingWebhookAuthentication) SetKey(v string)`

SetKey sets Key field to given value.


### GetValueLastChars

`func (o *NotificationIntegrationSettingWebhookAuthentication) GetValueLastChars() string`

GetValueLastChars returns the ValueLastChars field if non-nil, zero value otherwise.

### GetValueLastCharsOk

`func (o *NotificationIntegrationSettingWebhookAuthentication) GetValueLastCharsOk() (*string, bool)`

GetValueLastCharsOk returns a tuple with the ValueLastChars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueLastChars

`func (o *NotificationIntegrationSettingWebhookAuthentication) SetValueLastChars(v string)`

SetValueLastChars sets ValueLastChars field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


