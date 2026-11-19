# InputWebhookAuthenticationUpdateCustomHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Authenticate with a custom HTTP header. | 
**Key** | Pointer to **string** | Name of the HTTP header sent with each request. Omit to keep the stored value; required when switching from another authentication type. | [optional] 
**Value** | Pointer to **string** | Value of the HTTP header. Omit to keep the stored value; required when switching from another authentication type. | [optional] 

## Methods

### NewInputWebhookAuthenticationUpdateCustomHeader

`func NewInputWebhookAuthenticationUpdateCustomHeader(type_ string, ) *InputWebhookAuthenticationUpdateCustomHeader`

NewInputWebhookAuthenticationUpdateCustomHeader instantiates a new InputWebhookAuthenticationUpdateCustomHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputWebhookAuthenticationUpdateCustomHeaderWithDefaults

`func NewInputWebhookAuthenticationUpdateCustomHeaderWithDefaults() *InputWebhookAuthenticationUpdateCustomHeader`

NewInputWebhookAuthenticationUpdateCustomHeaderWithDefaults instantiates a new InputWebhookAuthenticationUpdateCustomHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputWebhookAuthenticationUpdateCustomHeader) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputWebhookAuthenticationUpdateCustomHeader) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputWebhookAuthenticationUpdateCustomHeader) SetType(v string)`

SetType sets Type field to given value.


### GetKey

`func (o *InputWebhookAuthenticationUpdateCustomHeader) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *InputWebhookAuthenticationUpdateCustomHeader) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *InputWebhookAuthenticationUpdateCustomHeader) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *InputWebhookAuthenticationUpdateCustomHeader) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *InputWebhookAuthenticationUpdateCustomHeader) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InputWebhookAuthenticationUpdateCustomHeader) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InputWebhookAuthenticationUpdateCustomHeader) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *InputWebhookAuthenticationUpdateCustomHeader) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


