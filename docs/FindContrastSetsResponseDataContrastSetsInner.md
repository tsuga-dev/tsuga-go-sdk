# FindContrastSetsResponseDataContrastSetsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The attribute is compared value by value. | 
**Attr** | **string** | Span attribute dot-path this contrast set is about, such as &#x60;context.env&#x60;. | 
**Values** | [**[]ContrastSetValue**](ContrastSetValue.md) | Values over-represented in the target group that cleared the lift, support, and significance gates, strongest first. These are the findings. | 
**OtherValues** | [**[]ContrastSetOtherValue**](ContrastSetOtherValue.md) | The attribute&#39;s other frequent values in the target group, by target percentage descending, as context for reading &#x60;values&#x60;. Never overlaps &#x60;values&#x60;. | 
**Shift** | [**ContrastSetShift**](ContrastSetShift.md) |  | 

## Methods

### NewFindContrastSetsResponseDataContrastSetsInner

`func NewFindContrastSetsResponseDataContrastSetsInner(type_ string, attr string, values []ContrastSetValue, otherValues []ContrastSetOtherValue, shift ContrastSetShift, ) *FindContrastSetsResponseDataContrastSetsInner`

NewFindContrastSetsResponseDataContrastSetsInner instantiates a new FindContrastSetsResponseDataContrastSetsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindContrastSetsResponseDataContrastSetsInnerWithDefaults

`func NewFindContrastSetsResponseDataContrastSetsInnerWithDefaults() *FindContrastSetsResponseDataContrastSetsInner`

NewFindContrastSetsResponseDataContrastSetsInnerWithDefaults instantiates a new FindContrastSetsResponseDataContrastSetsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FindContrastSetsResponseDataContrastSetsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FindContrastSetsResponseDataContrastSetsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FindContrastSetsResponseDataContrastSetsInner) SetType(v string)`

SetType sets Type field to given value.


### GetAttr

`func (o *FindContrastSetsResponseDataContrastSetsInner) GetAttr() string`

GetAttr returns the Attr field if non-nil, zero value otherwise.

### GetAttrOk

`func (o *FindContrastSetsResponseDataContrastSetsInner) GetAttrOk() (*string, bool)`

GetAttrOk returns a tuple with the Attr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttr

`func (o *FindContrastSetsResponseDataContrastSetsInner) SetAttr(v string)`

SetAttr sets Attr field to given value.


### GetValues

`func (o *FindContrastSetsResponseDataContrastSetsInner) GetValues() []ContrastSetValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *FindContrastSetsResponseDataContrastSetsInner) GetValuesOk() (*[]ContrastSetValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *FindContrastSetsResponseDataContrastSetsInner) SetValues(v []ContrastSetValue)`

SetValues sets Values field to given value.


### GetOtherValues

`func (o *FindContrastSetsResponseDataContrastSetsInner) GetOtherValues() []ContrastSetOtherValue`

GetOtherValues returns the OtherValues field if non-nil, zero value otherwise.

### GetOtherValuesOk

`func (o *FindContrastSetsResponseDataContrastSetsInner) GetOtherValuesOk() (*[]ContrastSetOtherValue, bool)`

GetOtherValuesOk returns a tuple with the OtherValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherValues

`func (o *FindContrastSetsResponseDataContrastSetsInner) SetOtherValues(v []ContrastSetOtherValue)`

SetOtherValues sets OtherValues field to given value.


### GetShift

`func (o *FindContrastSetsResponseDataContrastSetsInner) GetShift() ContrastSetShift`

GetShift returns the Shift field if non-nil, zero value otherwise.

### GetShiftOk

`func (o *FindContrastSetsResponseDataContrastSetsInner) GetShiftOk() (*ContrastSetShift, bool)`

GetShiftOk returns a tuple with the Shift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShift

`func (o *FindContrastSetsResponseDataContrastSetsInner) SetShift(v ContrastSetShift)`

SetShift sets Shift field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


