# ContrastSetOtherValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** | The attribute value this row is about. | 
**TargetSupport** | **float32** | Percentage of target spans carrying this value, from 0 to 100. | 
**BaselineSupport** | **float32** | Percentage of baseline spans carrying this value, from 0 to 100. | 

## Methods

### NewContrastSetOtherValue

`func NewContrastSetOtherValue(value string, targetSupport float32, baselineSupport float32, ) *ContrastSetOtherValue`

NewContrastSetOtherValue instantiates a new ContrastSetOtherValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContrastSetOtherValueWithDefaults

`func NewContrastSetOtherValueWithDefaults() *ContrastSetOtherValue`

NewContrastSetOtherValueWithDefaults instantiates a new ContrastSetOtherValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ContrastSetOtherValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ContrastSetOtherValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ContrastSetOtherValue) SetValue(v string)`

SetValue sets Value field to given value.


### GetTargetSupport

`func (o *ContrastSetOtherValue) GetTargetSupport() float32`

GetTargetSupport returns the TargetSupport field if non-nil, zero value otherwise.

### GetTargetSupportOk

`func (o *ContrastSetOtherValue) GetTargetSupportOk() (*float32, bool)`

GetTargetSupportOk returns a tuple with the TargetSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSupport

`func (o *ContrastSetOtherValue) SetTargetSupport(v float32)`

SetTargetSupport sets TargetSupport field to given value.


### GetBaselineSupport

`func (o *ContrastSetOtherValue) GetBaselineSupport() float32`

GetBaselineSupport returns the BaselineSupport field if non-nil, zero value otherwise.

### GetBaselineSupportOk

`func (o *ContrastSetOtherValue) GetBaselineSupportOk() (*float32, bool)`

GetBaselineSupportOk returns a tuple with the BaselineSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineSupport

`func (o *ContrastSetOtherValue) SetBaselineSupport(v float32)`

SetBaselineSupport sets BaselineSupport field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


