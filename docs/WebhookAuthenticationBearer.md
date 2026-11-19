# WebhookAuthenticationBearer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Authenticate with a bearer token in the &#x60;Authorization&#x60; header. | 
**TokenLastChars** | **string** | Redacted bearer token: only its last 4 characters, masked as &#x60;****1234&#x60;. | 

## Methods

### NewWebhookAuthenticationBearer

`func NewWebhookAuthenticationBearer(type_ string, tokenLastChars string, ) *WebhookAuthenticationBearer`

NewWebhookAuthenticationBearer instantiates a new WebhookAuthenticationBearer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookAuthenticationBearerWithDefaults

`func NewWebhookAuthenticationBearerWithDefaults() *WebhookAuthenticationBearer`

NewWebhookAuthenticationBearerWithDefaults instantiates a new WebhookAuthenticationBearer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WebhookAuthenticationBearer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookAuthenticationBearer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookAuthenticationBearer) SetType(v string)`

SetType sets Type field to given value.


### GetTokenLastChars

`func (o *WebhookAuthenticationBearer) GetTokenLastChars() string`

GetTokenLastChars returns the TokenLastChars field if non-nil, zero value otherwise.

### GetTokenLastCharsOk

`func (o *WebhookAuthenticationBearer) GetTokenLastCharsOk() (*string, bool)`

GetTokenLastCharsOk returns a tuple with the TokenLastChars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenLastChars

`func (o *WebhookAuthenticationBearer) SetTokenLastChars(v string)`

SetTokenLastChars sets TokenLastChars field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


