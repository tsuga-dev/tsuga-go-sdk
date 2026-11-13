# TimelineSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | **string** | Comparator used to evaluate each time bucket value against &#x60;value&#x60; | 
**Value** | **float32** | Bound of the section, compared with &#x60;operator&#x60; | 
**Color** | **string** | Color of the time buckets in this section | 
**Name** | Pointer to **string** | Name of the state, shown in the legend | [optional] 

## Methods

### NewTimelineSection

`func NewTimelineSection(operator string, value float32, color string, ) *TimelineSection`

NewTimelineSection instantiates a new TimelineSection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineSectionWithDefaults

`func NewTimelineSectionWithDefaults() *TimelineSection`

NewTimelineSectionWithDefaults instantiates a new TimelineSection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *TimelineSection) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *TimelineSection) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *TimelineSection) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *TimelineSection) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TimelineSection) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TimelineSection) SetValue(v float32)`

SetValue sets Value field to given value.


### GetColor

`func (o *TimelineSection) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *TimelineSection) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *TimelineSection) SetColor(v string)`

SetColor sets Color field to given value.


### GetName

`func (o *TimelineSection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TimelineSection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TimelineSection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TimelineSection) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


