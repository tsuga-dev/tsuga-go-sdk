# CreateNotificationSilenceRequestSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**StartTime** | **time.Time** |  | 
**EndTime** | **time.Time** |  | 
**WeeklySchedule** | [**SilenceScheduleRecurringWeeklySchedule**](SilenceScheduleRecurringWeeklySchedule.md) |  | 

## Methods

### NewCreateNotificationSilenceRequestSchedule

`func NewCreateNotificationSilenceRequestSchedule(type_ string, startTime time.Time, endTime time.Time, weeklySchedule SilenceScheduleRecurringWeeklySchedule, ) *CreateNotificationSilenceRequestSchedule`

NewCreateNotificationSilenceRequestSchedule instantiates a new CreateNotificationSilenceRequestSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNotificationSilenceRequestScheduleWithDefaults

`func NewCreateNotificationSilenceRequestScheduleWithDefaults() *CreateNotificationSilenceRequestSchedule`

NewCreateNotificationSilenceRequestScheduleWithDefaults instantiates a new CreateNotificationSilenceRequestSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateNotificationSilenceRequestSchedule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateNotificationSilenceRequestSchedule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateNotificationSilenceRequestSchedule) SetType(v string)`

SetType sets Type field to given value.


### GetStartTime

`func (o *CreateNotificationSilenceRequestSchedule) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CreateNotificationSilenceRequestSchedule) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CreateNotificationSilenceRequestSchedule) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *CreateNotificationSilenceRequestSchedule) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CreateNotificationSilenceRequestSchedule) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *CreateNotificationSilenceRequestSchedule) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### GetWeeklySchedule

`func (o *CreateNotificationSilenceRequestSchedule) GetWeeklySchedule() SilenceScheduleRecurringWeeklySchedule`

GetWeeklySchedule returns the WeeklySchedule field if non-nil, zero value otherwise.

### GetWeeklyScheduleOk

`func (o *CreateNotificationSilenceRequestSchedule) GetWeeklyScheduleOk() (*SilenceScheduleRecurringWeeklySchedule, bool)`

GetWeeklyScheduleOk returns a tuple with the WeeklySchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeklySchedule

`func (o *CreateNotificationSilenceRequestSchedule) SetWeeklySchedule(v SilenceScheduleRecurringWeeklySchedule)`

SetWeeklySchedule sets WeeklySchedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


