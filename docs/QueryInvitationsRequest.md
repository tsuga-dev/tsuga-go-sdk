# QueryInvitationsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | Maximum number of items to return in this page. Omit to use the public API default of 100. | [optional] [default to 100]
**Offset** | Pointer to **int32** | Zero-based index of the first matching item to return. Increase it with &#x60;limit&#x60; to request later pages. If &#x60;limit&#x60; is provided without &#x60;offset&#x60;, the offset defaults to 0. | [optional] 
**Filters** | Pointer to [**QueryInvitationsRequestFilters**](QueryInvitationsRequestFilters.md) |  | [optional] 
**Sort** | Pointer to [**QueryInvitationsRequestSort**](QueryInvitationsRequestSort.md) |  | [optional] 

## Methods

### NewQueryInvitationsRequest

`func NewQueryInvitationsRequest() *QueryInvitationsRequest`

NewQueryInvitationsRequest instantiates a new QueryInvitationsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryInvitationsRequestWithDefaults

`func NewQueryInvitationsRequestWithDefaults() *QueryInvitationsRequest`

NewQueryInvitationsRequestWithDefaults instantiates a new QueryInvitationsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *QueryInvitationsRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *QueryInvitationsRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *QueryInvitationsRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *QueryInvitationsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *QueryInvitationsRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *QueryInvitationsRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *QueryInvitationsRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *QueryInvitationsRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetFilters

`func (o *QueryInvitationsRequest) GetFilters() QueryInvitationsRequestFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *QueryInvitationsRequest) GetFiltersOk() (*QueryInvitationsRequestFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *QueryInvitationsRequest) SetFilters(v QueryInvitationsRequestFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *QueryInvitationsRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSort

`func (o *QueryInvitationsRequest) GetSort() QueryInvitationsRequestSort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *QueryInvitationsRequest) GetSortOk() (*QueryInvitationsRequestSort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *QueryInvitationsRequest) SetSort(v QueryInvitationsRequestSort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *QueryInvitationsRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


