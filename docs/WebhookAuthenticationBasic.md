# WebhookAuthenticationBasic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Authenticate with HTTP Basic credentials. | 
**Username** | **string** | Basic authentication username. | 
**PasswordLastChars** | **string** | Redacted Basic authentication password: only its last 4 characters, masked as &#x60;****1234&#x60;. | 

## Methods

### NewWebhookAuthenticationBasic

`func NewWebhookAuthenticationBasic(type_ string, username string, passwordLastChars string, ) *WebhookAuthenticationBasic`

NewWebhookAuthenticationBasic instantiates a new WebhookAuthenticationBasic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookAuthenticationBasicWithDefaults

`func NewWebhookAuthenticationBasicWithDefaults() *WebhookAuthenticationBasic`

NewWebhookAuthenticationBasicWithDefaults instantiates a new WebhookAuthenticationBasic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WebhookAuthenticationBasic) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookAuthenticationBasic) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookAuthenticationBasic) SetType(v string)`

SetType sets Type field to given value.


### GetUsername

`func (o *WebhookAuthenticationBasic) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *WebhookAuthenticationBasic) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *WebhookAuthenticationBasic) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPasswordLastChars

`func (o *WebhookAuthenticationBasic) GetPasswordLastChars() string`

GetPasswordLastChars returns the PasswordLastChars field if non-nil, zero value otherwise.

### GetPasswordLastCharsOk

`func (o *WebhookAuthenticationBasic) GetPasswordLastCharsOk() (*string, bool)`

GetPasswordLastCharsOk returns a tuple with the PasswordLastChars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLastChars

`func (o *WebhookAuthenticationBasic) SetPasswordLastChars(v string)`

SetPasswordLastChars sets PasswordLastChars field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


