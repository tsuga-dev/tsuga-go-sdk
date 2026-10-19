# InputSilenceScheduleRecurring

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Silence active on a recurring weekly schedule. | 
**WeeklySchedule** | [**InputSilenceScheduleRecurringWeeklySchedule**](InputSilenceScheduleRecurringWeeklySchedule.md) |  | 
**TimeZone** | Pointer to **string** | IANA timezone identifier (e.g. \&quot;America/New_York\&quot;). Defaults to UTC if omitted. | [optional] 

## Methods

### NewInputSilenceScheduleRecurring

`func NewInputSilenceScheduleRecurring(type_ string, weeklySchedule InputSilenceScheduleRecurringWeeklySchedule, ) *InputSilenceScheduleRecurring`

NewInputSilenceScheduleRecurring instantiates a new InputSilenceScheduleRecurring object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputSilenceScheduleRecurringWithDefaults

`func NewInputSilenceScheduleRecurringWithDefaults() *InputSilenceScheduleRecurring`

NewInputSilenceScheduleRecurringWithDefaults instantiates a new InputSilenceScheduleRecurring object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputSilenceScheduleRecurring) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputSilenceScheduleRecurring) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputSilenceScheduleRecurring) SetType(v string)`

SetType sets Type field to given value.


### GetWeeklySchedule

`func (o *InputSilenceScheduleRecurring) GetWeeklySchedule() InputSilenceScheduleRecurringWeeklySchedule`

GetWeeklySchedule returns the WeeklySchedule field if non-nil, zero value otherwise.

### GetWeeklyScheduleOk

`func (o *InputSilenceScheduleRecurring) GetWeeklyScheduleOk() (*InputSilenceScheduleRecurringWeeklySchedule, bool)`

GetWeeklyScheduleOk returns a tuple with the WeeklySchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeklySchedule

`func (o *InputSilenceScheduleRecurring) SetWeeklySchedule(v InputSilenceScheduleRecurringWeeklySchedule)`

SetWeeklySchedule sets WeeklySchedule field to given value.


### GetTimeZone

`func (o *InputSilenceScheduleRecurring) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *InputSilenceScheduleRecurring) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *InputSilenceScheduleRecurring) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *InputSilenceScheduleRecurring) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


