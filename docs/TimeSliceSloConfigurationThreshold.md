# TimeSliceSloConfigurationThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | **string** | Comparison operator between the query value and the threshold | 
**Value** | **float32** | Threshold value compared against the query signal | 

## Methods

### NewTimeSliceSloConfigurationThreshold

`func NewTimeSliceSloConfigurationThreshold(operator string, value float32, ) *TimeSliceSloConfigurationThreshold`

NewTimeSliceSloConfigurationThreshold instantiates a new TimeSliceSloConfigurationThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSliceSloConfigurationThresholdWithDefaults

`func NewTimeSliceSloConfigurationThresholdWithDefaults() *TimeSliceSloConfigurationThreshold`

NewTimeSliceSloConfigurationThresholdWithDefaults instantiates a new TimeSliceSloConfigurationThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *TimeSliceSloConfigurationThreshold) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *TimeSliceSloConfigurationThreshold) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *TimeSliceSloConfigurationThreshold) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *TimeSliceSloConfigurationThreshold) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TimeSliceSloConfigurationThreshold) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TimeSliceSloConfigurationThreshold) SetValue(v float32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


