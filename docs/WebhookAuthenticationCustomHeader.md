# WebhookAuthenticationCustomHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Authenticate with a custom HTTP header. | 
**Key** | **string** | Name of the HTTP header sent with each request. | 
**ValueLastChars** | **string** | Redacted header value: only its last 4 characters, masked as &#x60;****1234&#x60;. | 

## Methods

### NewWebhookAuthenticationCustomHeader

`func NewWebhookAuthenticationCustomHeader(type_ string, key string, valueLastChars string, ) *WebhookAuthenticationCustomHeader`

NewWebhookAuthenticationCustomHeader instantiates a new WebhookAuthenticationCustomHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookAuthenticationCustomHeaderWithDefaults

`func NewWebhookAuthenticationCustomHeaderWithDefaults() *WebhookAuthenticationCustomHeader`

NewWebhookAuthenticationCustomHeaderWithDefaults instantiates a new WebhookAuthenticationCustomHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WebhookAuthenticationCustomHeader) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookAuthenticationCustomHeader) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookAuthenticationCustomHeader) SetType(v string)`

SetType sets Type field to given value.


### GetKey

`func (o *WebhookAuthenticationCustomHeader) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *WebhookAuthenticationCustomHeader) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *WebhookAuthenticationCustomHeader) SetKey(v string)`

SetKey sets Key field to given value.


### GetValueLastChars

`func (o *WebhookAuthenticationCustomHeader) GetValueLastChars() string`

GetValueLastChars returns the ValueLastChars field if non-nil, zero value otherwise.

### GetValueLastCharsOk

`func (o *WebhookAuthenticationCustomHeader) GetValueLastCharsOk() (*string, bool)`

GetValueLastCharsOk returns a tuple with the ValueLastChars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueLastChars

`func (o *WebhookAuthenticationCustomHeader) SetValueLastChars(v string)`

SetValueLastChars sets ValueLastChars field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


