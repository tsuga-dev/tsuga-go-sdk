# ContrastSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attr** | **string** | Span attribute dot-path this contrast set is about, such as &#x60;context.env&#x60;. | 
**Values** | [**[]ContrastSetValue**](ContrastSetValue.md) | Values over-represented in the target group that cleared the lift, support, and significance gates, strongest first. These are the findings. | 
**OtherValues** | [**[]ContrastSetOtherValue**](ContrastSetOtherValue.md) | The attribute&#39;s other frequent values in the target group, by target percentage descending, as context for reading &#x60;values&#x60;. Never overlaps &#x60;values&#x60;. | 

## Methods

### NewContrastSet

`func NewContrastSet(attr string, values []ContrastSetValue, otherValues []ContrastSetOtherValue, ) *ContrastSet`

NewContrastSet instantiates a new ContrastSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContrastSetWithDefaults

`func NewContrastSetWithDefaults() *ContrastSet`

NewContrastSetWithDefaults instantiates a new ContrastSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttr

`func (o *ContrastSet) GetAttr() string`

GetAttr returns the Attr field if non-nil, zero value otherwise.

### GetAttrOk

`func (o *ContrastSet) GetAttrOk() (*string, bool)`

GetAttrOk returns a tuple with the Attr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttr

`func (o *ContrastSet) SetAttr(v string)`

SetAttr sets Attr field to given value.


### GetValues

`func (o *ContrastSet) GetValues() []ContrastSetValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ContrastSet) GetValuesOk() (*[]ContrastSetValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ContrastSet) SetValues(v []ContrastSetValue)`

SetValues sets Values field to given value.


### GetOtherValues

`func (o *ContrastSet) GetOtherValues() []ContrastSetOtherValue`

GetOtherValues returns the OtherValues field if non-nil, zero value otherwise.

### GetOtherValuesOk

`func (o *ContrastSet) GetOtherValuesOk() (*[]ContrastSetOtherValue, bool)`

GetOtherValuesOk returns a tuple with the OtherValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherValues

`func (o *ContrastSet) SetOtherValues(v []ContrastSetOtherValue)`

SetOtherValues sets OtherValues field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


