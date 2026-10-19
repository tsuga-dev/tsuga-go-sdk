# InputSilenceScheduleRecurringWeeklySchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Monday** | Pointer to [**[]InputSilenceTimeRange**](InputSilenceTimeRange.md) | Silence windows for this day. Each window uses the schedule time zone, and windows on the same day cannot overlap. | [optional] 
**Tuesday** | Pointer to [**[]InputSilenceTimeRange**](InputSilenceTimeRange.md) | Silence windows for this day. Each window uses the schedule time zone, and windows on the same day cannot overlap. | [optional] 
**Wednesday** | Pointer to [**[]InputSilenceTimeRange**](InputSilenceTimeRange.md) | Silence windows for this day. Each window uses the schedule time zone, and windows on the same day cannot overlap. | [optional] 
**Thursday** | Pointer to [**[]InputSilenceTimeRange**](InputSilenceTimeRange.md) | Silence windows for this day. Each window uses the schedule time zone, and windows on the same day cannot overlap. | [optional] 
**Friday** | Pointer to [**[]InputSilenceTimeRange**](InputSilenceTimeRange.md) | Silence windows for this day. Each window uses the schedule time zone, and windows on the same day cannot overlap. | [optional] 
**Saturday** | Pointer to [**[]InputSilenceTimeRange**](InputSilenceTimeRange.md) | Silence windows for this day. Each window uses the schedule time zone, and windows on the same day cannot overlap. | [optional] 
**Sunday** | Pointer to [**[]InputSilenceTimeRange**](InputSilenceTimeRange.md) | Silence windows for this day. Each window uses the schedule time zone, and windows on the same day cannot overlap. | [optional] 

## Methods

### NewInputSilenceScheduleRecurringWeeklySchedule

`func NewInputSilenceScheduleRecurringWeeklySchedule() *InputSilenceScheduleRecurringWeeklySchedule`

NewInputSilenceScheduleRecurringWeeklySchedule instantiates a new InputSilenceScheduleRecurringWeeklySchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputSilenceScheduleRecurringWeeklyScheduleWithDefaults

`func NewInputSilenceScheduleRecurringWeeklyScheduleWithDefaults() *InputSilenceScheduleRecurringWeeklySchedule`

NewInputSilenceScheduleRecurringWeeklyScheduleWithDefaults instantiates a new InputSilenceScheduleRecurringWeeklySchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonday

`func (o *InputSilenceScheduleRecurringWeeklySchedule) GetMonday() []InputSilenceTimeRange`

GetMonday returns the Monday field if non-nil, zero value otherwise.

### GetMondayOk

`func (o *InputSilenceScheduleRecurringWeeklySchedule) GetMondayOk() (*[]InputSilenceTimeRange, bool)`

GetMondayOk returns a tuple with the Monday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonday

`func (o *InputSilenceScheduleRecurringWeeklySchedule) SetMonday(v []InputSilenceTimeRange)`

SetMonday sets Monday field to given value.

### HasMonday

`func (o *InputSilenceScheduleRecurringWeeklySchedule) HasMonday() bool`

HasMonday returns a boolean if a field has been set.

### GetTuesday

`func (o *InputSilenceScheduleRecurringWeeklySchedule) GetTuesday() []InputSilenceTimeRange`

GetTuesday returns the Tuesday field if non-nil, zero value otherwise.

### GetTuesdayOk

`func (o *InputSilenceScheduleRecurringWeeklySchedule) GetTuesdayOk() (*[]InputSilenceTimeRange, bool)`

GetTuesdayOk returns a tuple with the Tuesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuesday

`func (o *InputSilenceScheduleRecurringWeeklySchedule) SetTuesday(v []InputSilenceTimeRange)`

SetTuesday sets Tuesday field to given value.

### HasTuesday

`func (o *InputSilenceScheduleRecurringWeeklySchedule) HasTuesday() bool`

HasTuesday returns a boolean if a field has been set.

### GetWednesday

`func (o *InputSilenceScheduleRecurringWeeklySchedule) GetWednesday() []InputSilenceTimeRange`

GetWednesday returns the Wednesday field if non-nil, zero value otherwise.

### GetWednesdayOk

`func (o *InputSilenceScheduleRecurringWeeklySchedule) GetWednesdayOk() (*[]InputSilenceTimeRange, bool)`

GetWednesdayOk returns a tuple with the Wednesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWednesday

`func (o *InputSilenceScheduleRecurringWeeklySchedule) SetWednesday(v []InputSilenceTimeRange)`

SetWednesday sets Wednesday field to given value.

### HasWednesday

`func (o *InputSilenceScheduleRecurringWeeklySchedule) HasWednesday() bool`

HasWednesday returns a boolean if a field has been set.

### GetThursday

`func (o *InputSilenceScheduleRecurringWeeklySchedule) GetThursday() []InputSilenceTimeRange`

GetThursday returns the Thursday field if non-nil, zero value otherwise.

### GetThursdayOk

`func (o *InputSilenceScheduleRecurringWeeklySchedule) GetThursdayOk() (*[]InputSilenceTimeRange, bool)`

GetThursdayOk returns a tuple with the Thursday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThursday

`func (o *InputSilenceScheduleRecurringWeeklySchedule) SetThursday(v []InputSilenceTimeRange)`

SetThursday sets Thursday field to given value.

### HasThursday

`func (o *InputSilenceScheduleRecurringWeeklySchedule) HasThursday() bool`

HasThursday returns a boolean if a field has been set.

### GetFriday

`func (o *InputSilenceScheduleRecurringWeeklySchedule) GetFriday() []InputSilenceTimeRange`

GetFriday returns the Friday field if non-nil, zero value otherwise.

### GetFridayOk

`func (o *InputSilenceScheduleRecurringWeeklySchedule) GetFridayOk() (*[]InputSilenceTimeRange, bool)`

GetFridayOk returns a tuple with the Friday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriday

`func (o *InputSilenceScheduleRecurringWeeklySchedule) SetFriday(v []InputSilenceTimeRange)`

SetFriday sets Friday field to given value.

### HasFriday

`func (o *InputSilenceScheduleRecurringWeeklySchedule) HasFriday() bool`

HasFriday returns a boolean if a field has been set.

### GetSaturday

`func (o *InputSilenceScheduleRecurringWeeklySchedule) GetSaturday() []InputSilenceTimeRange`

GetSaturday returns the Saturday field if non-nil, zero value otherwise.

### GetSaturdayOk

`func (o *InputSilenceScheduleRecurringWeeklySchedule) GetSaturdayOk() (*[]InputSilenceTimeRange, bool)`

GetSaturdayOk returns a tuple with the Saturday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturday

`func (o *InputSilenceScheduleRecurringWeeklySchedule) SetSaturday(v []InputSilenceTimeRange)`

SetSaturday sets Saturday field to given value.

### HasSaturday

`func (o *InputSilenceScheduleRecurringWeeklySchedule) HasSaturday() bool`

HasSaturday returns a boolean if a field has been set.

### GetSunday

`func (o *InputSilenceScheduleRecurringWeeklySchedule) GetSunday() []InputSilenceTimeRange`

GetSunday returns the Sunday field if non-nil, zero value otherwise.

### GetSundayOk

`func (o *InputSilenceScheduleRecurringWeeklySchedule) GetSundayOk() (*[]InputSilenceTimeRange, bool)`

GetSundayOk returns a tuple with the Sunday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSunday

`func (o *InputSilenceScheduleRecurringWeeklySchedule) SetSunday(v []InputSilenceTimeRange)`

SetSunday sets Sunday field to given value.

### HasSunday

`func (o *InputSilenceScheduleRecurringWeeklySchedule) HasSunday() bool`

HasSunday returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


