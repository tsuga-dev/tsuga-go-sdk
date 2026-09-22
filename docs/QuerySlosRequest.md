# QuerySlosRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | Maximum number of items to return in this page. Omit to use the public API default of 100. | [optional] [default to 100]
**Offset** | Pointer to **int32** | Zero-based index of the first matching item to return. Increase it with &#x60;limit&#x60; to request later pages. If &#x60;limit&#x60; is provided without &#x60;offset&#x60;, the offset defaults to 0. | [optional] 
**Filters** | Pointer to [**QuerySlosRequestFilters**](QuerySlosRequestFilters.md) |  | [optional] 
**Sort** | Pointer to [**QuerySlosRequestSort**](QuerySlosRequestSort.md) |  | [optional] 

## Methods

### NewQuerySlosRequest

`func NewQuerySlosRequest() *QuerySlosRequest`

NewQuerySlosRequest instantiates a new QuerySlosRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuerySlosRequestWithDefaults

`func NewQuerySlosRequestWithDefaults() *QuerySlosRequest`

NewQuerySlosRequestWithDefaults instantiates a new QuerySlosRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *QuerySlosRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *QuerySlosRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *QuerySlosRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *QuerySlosRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *QuerySlosRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *QuerySlosRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *QuerySlosRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *QuerySlosRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetFilters

`func (o *QuerySlosRequest) GetFilters() QuerySlosRequestFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *QuerySlosRequest) GetFiltersOk() (*QuerySlosRequestFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *QuerySlosRequest) SetFilters(v QuerySlosRequestFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *QuerySlosRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSort

`func (o *QuerySlosRequest) GetSort() QuerySlosRequestSort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *QuerySlosRequest) GetSortOk() (*QuerySlosRequestSort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *QuerySlosRequest) SetSort(v QuerySlosRequestSort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *QuerySlosRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


