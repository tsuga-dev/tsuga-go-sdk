# CategoricalContrastSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The attribute is compared value by value. | 
**Attr** | **string** | Span attribute dot-path this contrast set is about, such as &#x60;context.env&#x60;. | 
**Values** | [**[]ContrastSetValue**](ContrastSetValue.md) | Values over-represented in the target group that cleared the lift, support, and significance gates, strongest first. These are the findings. | 
**OtherValues** | [**[]ContrastSetOtherValue**](ContrastSetOtherValue.md) | The attribute&#39;s other frequent values in the target group, by target percentage descending, as context for reading &#x60;values&#x60;. Never overlaps &#x60;values&#x60;. | 

## Methods

### NewCategoricalContrastSet

`func NewCategoricalContrastSet(type_ string, attr string, values []ContrastSetValue, otherValues []ContrastSetOtherValue, ) *CategoricalContrastSet`

NewCategoricalContrastSet instantiates a new CategoricalContrastSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoricalContrastSetWithDefaults

`func NewCategoricalContrastSetWithDefaults() *CategoricalContrastSet`

NewCategoricalContrastSetWithDefaults instantiates a new CategoricalContrastSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CategoricalContrastSet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CategoricalContrastSet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CategoricalContrastSet) SetType(v string)`

SetType sets Type field to given value.


### GetAttr

`func (o *CategoricalContrastSet) GetAttr() string`

GetAttr returns the Attr field if non-nil, zero value otherwise.

### GetAttrOk

`func (o *CategoricalContrastSet) GetAttrOk() (*string, bool)`

GetAttrOk returns a tuple with the Attr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttr

`func (o *CategoricalContrastSet) SetAttr(v string)`

SetAttr sets Attr field to given value.


### GetValues

`func (o *CategoricalContrastSet) GetValues() []ContrastSetValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *CategoricalContrastSet) GetValuesOk() (*[]ContrastSetValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *CategoricalContrastSet) SetValues(v []ContrastSetValue)`

SetValues sets Values field to given value.


### GetOtherValues

`func (o *CategoricalContrastSet) GetOtherValues() []ContrastSetOtherValue`

GetOtherValues returns the OtherValues field if non-nil, zero value otherwise.

### GetOtherValuesOk

`func (o *CategoricalContrastSet) GetOtherValuesOk() (*[]ContrastSetOtherValue, bool)`

GetOtherValuesOk returns a tuple with the OtherValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherValues

`func (o *CategoricalContrastSet) SetOtherValues(v []ContrastSetOtherValue)`

SetOtherValues sets OtherValues field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


