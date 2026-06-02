# QueryDashboardsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | The maximum number of items to return | [optional] 
**Offset** | Pointer to **int32** | The offset of the first item to return | [optional] 
**Filters** | Pointer to [**QueryDashboardsRequestFilters**](QueryDashboardsRequestFilters.md) |  | [optional] 
**Sort** | Pointer to [**QueryDashboardsRequestSort**](QueryDashboardsRequestSort.md) |  | [optional] 

## Methods

### NewQueryDashboardsRequest

`func NewQueryDashboardsRequest() *QueryDashboardsRequest`

NewQueryDashboardsRequest instantiates a new QueryDashboardsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryDashboardsRequestWithDefaults

`func NewQueryDashboardsRequestWithDefaults() *QueryDashboardsRequest`

NewQueryDashboardsRequestWithDefaults instantiates a new QueryDashboardsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *QueryDashboardsRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *QueryDashboardsRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *QueryDashboardsRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *QueryDashboardsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *QueryDashboardsRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *QueryDashboardsRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *QueryDashboardsRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *QueryDashboardsRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetFilters

`func (o *QueryDashboardsRequest) GetFilters() QueryDashboardsRequestFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *QueryDashboardsRequest) GetFiltersOk() (*QueryDashboardsRequestFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *QueryDashboardsRequest) SetFilters(v QueryDashboardsRequestFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *QueryDashboardsRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSort

`func (o *QueryDashboardsRequest) GetSort() QueryDashboardsRequestSort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *QueryDashboardsRequest) GetSortOk() (*QueryDashboardsRequestSort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *QueryDashboardsRequest) SetSort(v QueryDashboardsRequestSort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *QueryDashboardsRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


