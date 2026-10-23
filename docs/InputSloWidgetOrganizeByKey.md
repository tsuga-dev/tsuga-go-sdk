# InputSloWidgetOrganizeByKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Groups the results by the value of a tag key or group-by field | 
**Key** | **string** | Tag key or group-by field to group by, such as &#x60;env&#x60; or &#x60;context.team&#x60;. A result falls under each value it has for this key: its group-by value when its SLO groups by this field, and its SLO&#39;s tag values for this key. Results with no value for the key are not shown. | 

## Methods

### NewInputSloWidgetOrganizeByKey

`func NewInputSloWidgetOrganizeByKey(type_ string, key string, ) *InputSloWidgetOrganizeByKey`

NewInputSloWidgetOrganizeByKey instantiates a new InputSloWidgetOrganizeByKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputSloWidgetOrganizeByKeyWithDefaults

`func NewInputSloWidgetOrganizeByKeyWithDefaults() *InputSloWidgetOrganizeByKey`

NewInputSloWidgetOrganizeByKeyWithDefaults instantiates a new InputSloWidgetOrganizeByKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputSloWidgetOrganizeByKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputSloWidgetOrganizeByKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputSloWidgetOrganizeByKey) SetType(v string)`

SetType sets Type field to given value.


### GetKey

`func (o *InputSloWidgetOrganizeByKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *InputSloWidgetOrganizeByKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *InputSloWidgetOrganizeByKey) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


