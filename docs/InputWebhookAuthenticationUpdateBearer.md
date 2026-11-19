# InputWebhookAuthenticationUpdateBearer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Authenticate with a bearer token in the &#x60;Authorization&#x60; header. | 
**Token** | Pointer to **string** | Bearer token sent in the &#x60;Authorization&#x60; header. Omit to keep the stored value; required when switching from another authentication type. | [optional] 

## Methods

### NewInputWebhookAuthenticationUpdateBearer

`func NewInputWebhookAuthenticationUpdateBearer(type_ string, ) *InputWebhookAuthenticationUpdateBearer`

NewInputWebhookAuthenticationUpdateBearer instantiates a new InputWebhookAuthenticationUpdateBearer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputWebhookAuthenticationUpdateBearerWithDefaults

`func NewInputWebhookAuthenticationUpdateBearerWithDefaults() *InputWebhookAuthenticationUpdateBearer`

NewInputWebhookAuthenticationUpdateBearerWithDefaults instantiates a new InputWebhookAuthenticationUpdateBearer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputWebhookAuthenticationUpdateBearer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputWebhookAuthenticationUpdateBearer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputWebhookAuthenticationUpdateBearer) SetType(v string)`

SetType sets Type field to given value.


### GetToken

`func (o *InputWebhookAuthenticationUpdateBearer) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *InputWebhookAuthenticationUpdateBearer) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *InputWebhookAuthenticationUpdateBearer) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *InputWebhookAuthenticationUpdateBearer) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


