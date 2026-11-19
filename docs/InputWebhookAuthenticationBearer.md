# InputWebhookAuthenticationBearer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Authenticate with a bearer token in the &#x60;Authorization&#x60; header. | 
**Token** | **string** | Bearer token sent in the &#x60;Authorization&#x60; header. | 

## Methods

### NewInputWebhookAuthenticationBearer

`func NewInputWebhookAuthenticationBearer(type_ string, token string, ) *InputWebhookAuthenticationBearer`

NewInputWebhookAuthenticationBearer instantiates a new InputWebhookAuthenticationBearer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputWebhookAuthenticationBearerWithDefaults

`func NewInputWebhookAuthenticationBearerWithDefaults() *InputWebhookAuthenticationBearer`

NewInputWebhookAuthenticationBearerWithDefaults instantiates a new InputWebhookAuthenticationBearer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputWebhookAuthenticationBearer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputWebhookAuthenticationBearer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputWebhookAuthenticationBearer) SetType(v string)`

SetType sets Type field to given value.


### GetToken

`func (o *InputWebhookAuthenticationBearer) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *InputWebhookAuthenticationBearer) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *InputWebhookAuthenticationBearer) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


