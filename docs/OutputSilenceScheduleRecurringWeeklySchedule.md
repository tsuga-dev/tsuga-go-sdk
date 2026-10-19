# OutputSilenceScheduleRecurringWeeklySchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Monday** | Pointer to [**[]SilenceTimeRange**](SilenceTimeRange.md) | Silence windows for this day. Each window uses the schedule time zone, and windows on the same day cannot overlap. | [optional] 
**Tuesday** | Pointer to [**[]SilenceTimeRange**](SilenceTimeRange.md) | Silence windows for this day. Each window uses the schedule time zone, and windows on the same day cannot overlap. | [optional] 
**Wednesday** | Pointer to [**[]SilenceTimeRange**](SilenceTimeRange.md) | Silence windows for this day. Each window uses the schedule time zone, and windows on the same day cannot overlap. | [optional] 
**Thursday** | Pointer to [**[]SilenceTimeRange**](SilenceTimeRange.md) | Silence windows for this day. Each window uses the schedule time zone, and windows on the same day cannot overlap. | [optional] 
**Friday** | Pointer to [**[]SilenceTimeRange**](SilenceTimeRange.md) | Silence windows for this day. Each window uses the schedule time zone, and windows on the same day cannot overlap. | [optional] 
**Saturday** | Pointer to [**[]SilenceTimeRange**](SilenceTimeRange.md) | Silence windows for this day. Each window uses the schedule time zone, and windows on the same day cannot overlap. | [optional] 
**Sunday** | Pointer to [**[]SilenceTimeRange**](SilenceTimeRange.md) | Silence windows for this day. Each window uses the schedule time zone, and windows on the same day cannot overlap. | [optional] 

## Methods

### NewOutputSilenceScheduleRecurringWeeklySchedule

`func NewOutputSilenceScheduleRecurringWeeklySchedule() *OutputSilenceScheduleRecurringWeeklySchedule`

NewOutputSilenceScheduleRecurringWeeklySchedule instantiates a new OutputSilenceScheduleRecurringWeeklySchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputSilenceScheduleRecurringWeeklyScheduleWithDefaults

`func NewOutputSilenceScheduleRecurringWeeklyScheduleWithDefaults() *OutputSilenceScheduleRecurringWeeklySchedule`

NewOutputSilenceScheduleRecurringWeeklyScheduleWithDefaults instantiates a new OutputSilenceScheduleRecurringWeeklySchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonday

`func (o *OutputSilenceScheduleRecurringWeeklySchedule) GetMonday() []SilenceTimeRange`

GetMonday returns the Monday field if non-nil, zero value otherwise.

### GetMondayOk

`func (o *OutputSilenceScheduleRecurringWeeklySchedule) GetMondayOk() (*[]SilenceTimeRange, bool)`

GetMondayOk returns a tuple with the Monday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonday

`func (o *OutputSilenceScheduleRecurringWeeklySchedule) SetMonday(v []SilenceTimeRange)`

SetMonday sets Monday field to given value.

### HasMonday

`func (o *OutputSilenceScheduleRecurringWeeklySchedule) HasMonday() bool`

HasMonday returns a boolean if a field has been set.

### GetTuesday

`func (o *OutputSilenceScheduleRecurringWeeklySchedule) GetTuesday() []SilenceTimeRange`

GetTuesday returns the Tuesday field if non-nil, zero value otherwise.

### GetTuesdayOk

`func (o *OutputSilenceScheduleRecurringWeeklySchedule) GetTuesdayOk() (*[]SilenceTimeRange, bool)`

GetTuesdayOk returns a tuple with the Tuesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuesday

`func (o *OutputSilenceScheduleRecurringWeeklySchedule) SetTuesday(v []SilenceTimeRange)`

SetTuesday sets Tuesday field to given value.

### HasTuesday

`func (o *OutputSilenceScheduleRecurringWeeklySchedule) HasTuesday() bool`

HasTuesday returns a boolean if a field has been set.

### GetWednesday

`func (o *OutputSilenceScheduleRecurringWeeklySchedule) GetWednesday() []SilenceTimeRange`

GetWednesday returns the Wednesday field if non-nil, zero value otherwise.

### GetWednesdayOk

`func (o *OutputSilenceScheduleRecurringWeeklySchedule) GetWednesdayOk() (*[]SilenceTimeRange, bool)`

GetWednesdayOk returns a tuple with the Wednesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWednesday

`func (o *OutputSilenceScheduleRecurringWeeklySchedule) SetWednesday(v []SilenceTimeRange)`

SetWednesday sets Wednesday field to given value.

### HasWednesday

`func (o *OutputSilenceScheduleRecurringWeeklySchedule) HasWednesday() bool`

HasWednesday returns a boolean if a field has been set.

### GetThursday

`func (o *OutputSilenceScheduleRecurringWeeklySchedule) GetThursday() []SilenceTimeRange`

GetThursday returns the Thursday field if non-nil, zero value otherwise.

### GetThursdayOk

`func (o *OutputSilenceScheduleRecurringWeeklySchedule) GetThursdayOk() (*[]SilenceTimeRange, bool)`

GetThursdayOk returns a tuple with the Thursday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThursday

`func (o *OutputSilenceScheduleRecurringWeeklySchedule) SetThursday(v []SilenceTimeRange)`

SetThursday sets Thursday field to given value.

### HasThursday

`func (o *OutputSilenceScheduleRecurringWeeklySchedule) HasThursday() bool`

HasThursday returns a boolean if a field has been set.

### GetFriday

`func (o *OutputSilenceScheduleRecurringWeeklySchedule) GetFriday() []SilenceTimeRange`

GetFriday returns the Friday field if non-nil, zero value otherwise.

### GetFridayOk

`func (o *OutputSilenceScheduleRecurringWeeklySchedule) GetFridayOk() (*[]SilenceTimeRange, bool)`

GetFridayOk returns a tuple with the Friday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriday

`func (o *OutputSilenceScheduleRecurringWeeklySchedule) SetFriday(v []SilenceTimeRange)`

SetFriday sets Friday field to given value.

### HasFriday

`func (o *OutputSilenceScheduleRecurringWeeklySchedule) HasFriday() bool`

HasFriday returns a boolean if a field has been set.

### GetSaturday

`func (o *OutputSilenceScheduleRecurringWeeklySchedule) GetSaturday() []SilenceTimeRange`

GetSaturday returns the Saturday field if non-nil, zero value otherwise.

### GetSaturdayOk

`func (o *OutputSilenceScheduleRecurringWeeklySchedule) GetSaturdayOk() (*[]SilenceTimeRange, bool)`

GetSaturdayOk returns a tuple with the Saturday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturday

`func (o *OutputSilenceScheduleRecurringWeeklySchedule) SetSaturday(v []SilenceTimeRange)`

SetSaturday sets Saturday field to given value.

### HasSaturday

`func (o *OutputSilenceScheduleRecurringWeeklySchedule) HasSaturday() bool`

HasSaturday returns a boolean if a field has been set.

### GetSunday

`func (o *OutputSilenceScheduleRecurringWeeklySchedule) GetSunday() []SilenceTimeRange`

GetSunday returns the Sunday field if non-nil, zero value otherwise.

### GetSundayOk

`func (o *OutputSilenceScheduleRecurringWeeklySchedule) GetSundayOk() (*[]SilenceTimeRange, bool)`

GetSundayOk returns a tuple with the Sunday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSunday

`func (o *OutputSilenceScheduleRecurringWeeklySchedule) SetSunday(v []SilenceTimeRange)`

SetSunday sets Sunday field to given value.

### HasSunday

`func (o *OutputSilenceScheduleRecurringWeeklySchedule) HasSunday() bool`

HasSunday returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


