# InputSilenceTimeRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | **string** | Local start time for this recurring silence window on this day, using the schedule time zone. | 
**EndTime** | **string** | Local end time for this recurring silence window on this day. It must be after &#x60;startTime&#x60;, and ranges on the same day cannot overlap. | 

## Methods

### NewInputSilenceTimeRange

`func NewInputSilenceTimeRange(startTime string, endTime string, ) *InputSilenceTimeRange`

NewInputSilenceTimeRange instantiates a new InputSilenceTimeRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputSilenceTimeRangeWithDefaults

`func NewInputSilenceTimeRangeWithDefaults() *InputSilenceTimeRange`

NewInputSilenceTimeRangeWithDefaults instantiates a new InputSilenceTimeRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *InputSilenceTimeRange) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *InputSilenceTimeRange) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *InputSilenceTimeRange) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *InputSilenceTimeRange) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *InputSilenceTimeRange) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *InputSilenceTimeRange) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


