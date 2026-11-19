# InputWebhookAuthenticationUpdateBasic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Authenticate with HTTP Basic credentials. | 
**Username** | Pointer to **string** | Basic authentication username. Omit to keep the stored value; required when switching from another authentication type. | [optional] 
**Password** | Pointer to **string** | Basic authentication password. Omit to keep the stored value; required when switching from another authentication type. | [optional] 

## Methods

### NewInputWebhookAuthenticationUpdateBasic

`func NewInputWebhookAuthenticationUpdateBasic(type_ string, ) *InputWebhookAuthenticationUpdateBasic`

NewInputWebhookAuthenticationUpdateBasic instantiates a new InputWebhookAuthenticationUpdateBasic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputWebhookAuthenticationUpdateBasicWithDefaults

`func NewInputWebhookAuthenticationUpdateBasicWithDefaults() *InputWebhookAuthenticationUpdateBasic`

NewInputWebhookAuthenticationUpdateBasicWithDefaults instantiates a new InputWebhookAuthenticationUpdateBasic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputWebhookAuthenticationUpdateBasic) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputWebhookAuthenticationUpdateBasic) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputWebhookAuthenticationUpdateBasic) SetType(v string)`

SetType sets Type field to given value.


### GetUsername

`func (o *InputWebhookAuthenticationUpdateBasic) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *InputWebhookAuthenticationUpdateBasic) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *InputWebhookAuthenticationUpdateBasic) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *InputWebhookAuthenticationUpdateBasic) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *InputWebhookAuthenticationUpdateBasic) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InputWebhookAuthenticationUpdateBasic) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InputWebhookAuthenticationUpdateBasic) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *InputWebhookAuthenticationUpdateBasic) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


