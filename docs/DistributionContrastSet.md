# DistributionContrastSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The attribute is numeric, such as a duration, and is compared by the shape of its values. | 
**Attr** | **string** | Span attribute dot-path this contrast set is about, such as &#x60;context.env&#x60;. | 
**Shift** | [**ContrastSetShift**](ContrastSetShift.md) |  | 

## Methods

### NewDistributionContrastSet

`func NewDistributionContrastSet(type_ string, attr string, shift ContrastSetShift, ) *DistributionContrastSet`

NewDistributionContrastSet instantiates a new DistributionContrastSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistributionContrastSetWithDefaults

`func NewDistributionContrastSetWithDefaults() *DistributionContrastSet`

NewDistributionContrastSetWithDefaults instantiates a new DistributionContrastSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DistributionContrastSet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DistributionContrastSet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DistributionContrastSet) SetType(v string)`

SetType sets Type field to given value.


### GetAttr

`func (o *DistributionContrastSet) GetAttr() string`

GetAttr returns the Attr field if non-nil, zero value otherwise.

### GetAttrOk

`func (o *DistributionContrastSet) GetAttrOk() (*string, bool)`

GetAttrOk returns a tuple with the Attr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttr

`func (o *DistributionContrastSet) SetAttr(v string)`

SetAttr sets Attr field to given value.


### GetShift

`func (o *DistributionContrastSet) GetShift() ContrastSetShift`

GetShift returns the Shift field if non-nil, zero value otherwise.

### GetShiftOk

`func (o *DistributionContrastSet) GetShiftOk() (*ContrastSetShift, bool)`

GetShiftOk returns a tuple with the Shift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShift

`func (o *DistributionContrastSet) SetShift(v ContrastSetShift)`

SetShift sets Shift field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


