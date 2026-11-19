# InputWebhookAuthenticationCustomHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Authenticate with a custom HTTP header. | 
**Key** | **string** | Name of the HTTP header sent with each request. | 
**Value** | **string** | Value of the HTTP header. | 

## Methods

### NewInputWebhookAuthenticationCustomHeader

`func NewInputWebhookAuthenticationCustomHeader(type_ string, key string, value string, ) *InputWebhookAuthenticationCustomHeader`

NewInputWebhookAuthenticationCustomHeader instantiates a new InputWebhookAuthenticationCustomHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputWebhookAuthenticationCustomHeaderWithDefaults

`func NewInputWebhookAuthenticationCustomHeaderWithDefaults() *InputWebhookAuthenticationCustomHeader`

NewInputWebhookAuthenticationCustomHeaderWithDefaults instantiates a new InputWebhookAuthenticationCustomHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputWebhookAuthenticationCustomHeader) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputWebhookAuthenticationCustomHeader) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputWebhookAuthenticationCustomHeader) SetType(v string)`

SetType sets Type field to given value.


### GetKey

`func (o *InputWebhookAuthenticationCustomHeader) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *InputWebhookAuthenticationCustomHeader) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *InputWebhookAuthenticationCustomHeader) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *InputWebhookAuthenticationCustomHeader) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InputWebhookAuthenticationCustomHeader) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InputWebhookAuthenticationCustomHeader) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


