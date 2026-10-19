# ContrastSetValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** | The attribute value this row is about. | 
**TargetSupport** | **float32** | Percentage of target spans carrying this value, from 0 to 100. | 
**BaselineSupport** | **float32** | Percentage of baseline spans carrying this value, from 0 to 100. | 
**Lift** | Pointer to **float32** | &#x60;targetSupport / baselineSupport&#x60;: how many times more frequent the value is in the target group. Absent when the value never appears in the baseline, which makes the lift unbounded. | [optional] 
**PValue** | **float32** | Probability that a difference this large would occur if the value were unrelated to the split. Smaller is stronger. | 

## Methods

### NewContrastSetValue

`func NewContrastSetValue(value string, targetSupport float32, baselineSupport float32, pValue float32, ) *ContrastSetValue`

NewContrastSetValue instantiates a new ContrastSetValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContrastSetValueWithDefaults

`func NewContrastSetValueWithDefaults() *ContrastSetValue`

NewContrastSetValueWithDefaults instantiates a new ContrastSetValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ContrastSetValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ContrastSetValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ContrastSetValue) SetValue(v string)`

SetValue sets Value field to given value.


### GetTargetSupport

`func (o *ContrastSetValue) GetTargetSupport() float32`

GetTargetSupport returns the TargetSupport field if non-nil, zero value otherwise.

### GetTargetSupportOk

`func (o *ContrastSetValue) GetTargetSupportOk() (*float32, bool)`

GetTargetSupportOk returns a tuple with the TargetSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSupport

`func (o *ContrastSetValue) SetTargetSupport(v float32)`

SetTargetSupport sets TargetSupport field to given value.


### GetBaselineSupport

`func (o *ContrastSetValue) GetBaselineSupport() float32`

GetBaselineSupport returns the BaselineSupport field if non-nil, zero value otherwise.

### GetBaselineSupportOk

`func (o *ContrastSetValue) GetBaselineSupportOk() (*float32, bool)`

GetBaselineSupportOk returns a tuple with the BaselineSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineSupport

`func (o *ContrastSetValue) SetBaselineSupport(v float32)`

SetBaselineSupport sets BaselineSupport field to given value.


### GetLift

`func (o *ContrastSetValue) GetLift() float32`

GetLift returns the Lift field if non-nil, zero value otherwise.

### GetLiftOk

`func (o *ContrastSetValue) GetLiftOk() (*float32, bool)`

GetLiftOk returns a tuple with the Lift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLift

`func (o *ContrastSetValue) SetLift(v float32)`

SetLift sets Lift field to given value.

### HasLift

`func (o *ContrastSetValue) HasLift() bool`

HasLift returns a boolean if a field has been set.

### GetPValue

`func (o *ContrastSetValue) GetPValue() float32`

GetPValue returns the PValue field if non-nil, zero value otherwise.

### GetPValueOk

`func (o *ContrastSetValue) GetPValueOk() (*float32, bool)`

GetPValueOk returns a tuple with the PValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPValue

`func (o *ContrastSetValue) SetPValue(v float32)`

SetPValue sets PValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


