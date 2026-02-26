# TimeRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **int64** | Start timestamp in seconds | 
**To** | **int64** | End timestamp in seconds | 

## Methods

### NewTimeRange

`func NewTimeRange(from int64, to int64, ) *TimeRange`

NewTimeRange instantiates a new TimeRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeRangeWithDefaults

`func NewTimeRangeWithDefaults() *TimeRange`

NewTimeRangeWithDefaults instantiates a new TimeRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *TimeRange) GetFrom() int64`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *TimeRange) GetFromOk() (*int64, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *TimeRange) SetFrom(v int64)`

SetFrom sets From field to given value.


### GetTo

`func (o *TimeRange) GetTo() int64`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *TimeRange) GetToOk() (*int64, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *TimeRange) SetTo(v int64)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


