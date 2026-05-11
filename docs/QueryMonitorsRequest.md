# QueryMonitorsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | The maximum number of items to return | [optional] 
**Offset** | Pointer to **int32** | The offset of the first item to return | [optional] 
**Filters** | Pointer to [**QueryMonitorsRequestFilters**](QueryMonitorsRequestFilters.md) |  | [optional] 
**Sort** | Pointer to [**QueryMonitorsRequestSort**](QueryMonitorsRequestSort.md) |  | [optional] 

## Methods

### NewQueryMonitorsRequest

`func NewQueryMonitorsRequest() *QueryMonitorsRequest`

NewQueryMonitorsRequest instantiates a new QueryMonitorsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryMonitorsRequestWithDefaults

`func NewQueryMonitorsRequestWithDefaults() *QueryMonitorsRequest`

NewQueryMonitorsRequestWithDefaults instantiates a new QueryMonitorsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *QueryMonitorsRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *QueryMonitorsRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *QueryMonitorsRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *QueryMonitorsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *QueryMonitorsRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *QueryMonitorsRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *QueryMonitorsRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *QueryMonitorsRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetFilters

`func (o *QueryMonitorsRequest) GetFilters() QueryMonitorsRequestFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *QueryMonitorsRequest) GetFiltersOk() (*QueryMonitorsRequestFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *QueryMonitorsRequest) SetFilters(v QueryMonitorsRequestFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *QueryMonitorsRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSort

`func (o *QueryMonitorsRequest) GetSort() QueryMonitorsRequestSort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *QueryMonitorsRequest) GetSortOk() (*QueryMonitorsRequestSort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *QueryMonitorsRequest) SetSort(v QueryMonitorsRequestSort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *QueryMonitorsRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


