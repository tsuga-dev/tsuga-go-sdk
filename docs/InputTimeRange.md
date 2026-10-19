# InputTimeRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **int64** | Start of the query window as a Unix timestamp in seconds | 
**To** | **int64** | End of the query window as a Unix timestamp in seconds | 

## Methods

### NewInputTimeRange

`func NewInputTimeRange(from int64, to int64, ) *InputTimeRange`

NewInputTimeRange instantiates a new InputTimeRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputTimeRangeWithDefaults

`func NewInputTimeRangeWithDefaults() *InputTimeRange`

NewInputTimeRangeWithDefaults instantiates a new InputTimeRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *InputTimeRange) GetFrom() int64`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *InputTimeRange) GetFromOk() (*int64, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *InputTimeRange) SetFrom(v int64)`

SetFrom sets From field to given value.


### GetTo

`func (o *InputTimeRange) GetTo() int64`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *InputTimeRange) GetToOk() (*int64, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *InputTimeRange) SetTo(v int64)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


