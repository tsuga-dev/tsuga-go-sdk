# FindContrastSetsRequestBaselineGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | **string** | Tsuga query selecting the baseline spans. Terms are ANDed by whitespace, as in &#x60;context.service.name:checkout status_code:error&#x60;. | 
**TimeRange** | [**InputTimeRange**](InputTimeRange.md) |  | 

## Methods

### NewFindContrastSetsRequestBaselineGroup

`func NewFindContrastSetsRequestBaselineGroup(filter string, timeRange InputTimeRange, ) *FindContrastSetsRequestBaselineGroup`

NewFindContrastSetsRequestBaselineGroup instantiates a new FindContrastSetsRequestBaselineGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindContrastSetsRequestBaselineGroupWithDefaults

`func NewFindContrastSetsRequestBaselineGroupWithDefaults() *FindContrastSetsRequestBaselineGroup`

NewFindContrastSetsRequestBaselineGroupWithDefaults instantiates a new FindContrastSetsRequestBaselineGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *FindContrastSetsRequestBaselineGroup) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *FindContrastSetsRequestBaselineGroup) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *FindContrastSetsRequestBaselineGroup) SetFilter(v string)`

SetFilter sets Filter field to given value.


### GetTimeRange

`func (o *FindContrastSetsRequestBaselineGroup) GetTimeRange() InputTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *FindContrastSetsRequestBaselineGroup) GetTimeRangeOk() (*InputTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *FindContrastSetsRequestBaselineGroup) SetTimeRange(v InputTimeRange)`

SetTimeRange sets TimeRange field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


