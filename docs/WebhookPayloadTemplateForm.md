# WebhookPayloadTemplateForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Value** | [**map[string]WebhookPayloadTemplateFormValueValue**](WebhookPayloadTemplateFormValueValue.md) |  | 

## Methods

### NewWebhookPayloadTemplateForm

`func NewWebhookPayloadTemplateForm(type_ string, value map[string]WebhookPayloadTemplateFormValueValue, ) *WebhookPayloadTemplateForm`

NewWebhookPayloadTemplateForm instantiates a new WebhookPayloadTemplateForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadTemplateFormWithDefaults

`func NewWebhookPayloadTemplateFormWithDefaults() *WebhookPayloadTemplateForm`

NewWebhookPayloadTemplateFormWithDefaults instantiates a new WebhookPayloadTemplateForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WebhookPayloadTemplateForm) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookPayloadTemplateForm) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookPayloadTemplateForm) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *WebhookPayloadTemplateForm) GetValue() map[string]WebhookPayloadTemplateFormValueValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WebhookPayloadTemplateForm) GetValueOk() (*map[string]WebhookPayloadTemplateFormValueValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WebhookPayloadTemplateForm) SetValue(v map[string]WebhookPayloadTemplateFormValueValue)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


