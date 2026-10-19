# InputSilenceScheduleOneTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Silence active once between the configured start and end times. | 
**StartTime** | **string** | Local start date and time for the one-time silence, interpreted in &#x60;timeZone&#x60; when provided. | 
**EndTime** | **string** | Local end date and time for the one-time silence. It must be after &#x60;startTime&#x60; and in the future. | 
**TimeZone** | Pointer to **string** | IANA timezone identifier (e.g. \&quot;America/New_York\&quot;). Defaults to UTC if omitted. | [optional] 

## Methods

### NewInputSilenceScheduleOneTime

`func NewInputSilenceScheduleOneTime(type_ string, startTime string, endTime string, ) *InputSilenceScheduleOneTime`

NewInputSilenceScheduleOneTime instantiates a new InputSilenceScheduleOneTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputSilenceScheduleOneTimeWithDefaults

`func NewInputSilenceScheduleOneTimeWithDefaults() *InputSilenceScheduleOneTime`

NewInputSilenceScheduleOneTimeWithDefaults instantiates a new InputSilenceScheduleOneTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputSilenceScheduleOneTime) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputSilenceScheduleOneTime) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputSilenceScheduleOneTime) SetType(v string)`

SetType sets Type field to given value.


### GetStartTime

`func (o *InputSilenceScheduleOneTime) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *InputSilenceScheduleOneTime) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *InputSilenceScheduleOneTime) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *InputSilenceScheduleOneTime) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *InputSilenceScheduleOneTime) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *InputSilenceScheduleOneTime) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.


### GetTimeZone

`func (o *InputSilenceScheduleOneTime) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *InputSilenceScheduleOneTime) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *InputSilenceScheduleOneTime) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *InputSilenceScheduleOneTime) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


