# OutputSilenceScheduleRecurring

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Silence active on a recurring weekly schedule. | 
**WeeklySchedule** | [**OutputSilenceScheduleRecurringWeeklySchedule**](OutputSilenceScheduleRecurringWeeklySchedule.md) |  | 
**TimeZone** | Pointer to **string** | IANA timezone identifier (e.g. \&quot;America/New_York\&quot;). Defaults to UTC if omitted. | [optional] 

## Methods

### NewOutputSilenceScheduleRecurring

`func NewOutputSilenceScheduleRecurring(type_ string, weeklySchedule OutputSilenceScheduleRecurringWeeklySchedule, ) *OutputSilenceScheduleRecurring`

NewOutputSilenceScheduleRecurring instantiates a new OutputSilenceScheduleRecurring object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputSilenceScheduleRecurringWithDefaults

`func NewOutputSilenceScheduleRecurringWithDefaults() *OutputSilenceScheduleRecurring`

NewOutputSilenceScheduleRecurringWithDefaults instantiates a new OutputSilenceScheduleRecurring object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OutputSilenceScheduleRecurring) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OutputSilenceScheduleRecurring) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OutputSilenceScheduleRecurring) SetType(v string)`

SetType sets Type field to given value.


### GetWeeklySchedule

`func (o *OutputSilenceScheduleRecurring) GetWeeklySchedule() OutputSilenceScheduleRecurringWeeklySchedule`

GetWeeklySchedule returns the WeeklySchedule field if non-nil, zero value otherwise.

### GetWeeklyScheduleOk

`func (o *OutputSilenceScheduleRecurring) GetWeeklyScheduleOk() (*OutputSilenceScheduleRecurringWeeklySchedule, bool)`

GetWeeklyScheduleOk returns a tuple with the WeeklySchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeklySchedule

`func (o *OutputSilenceScheduleRecurring) SetWeeklySchedule(v OutputSilenceScheduleRecurringWeeklySchedule)`

SetWeeklySchedule sets WeeklySchedule field to given value.


### GetTimeZone

`func (o *OutputSilenceScheduleRecurring) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *OutputSilenceScheduleRecurring) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *OutputSilenceScheduleRecurring) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *OutputSilenceScheduleRecurring) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


