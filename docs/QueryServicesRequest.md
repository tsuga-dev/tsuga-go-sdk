# QueryServicesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**QueryServicesRequestFilters**](QueryServicesRequestFilters.md) |  | [optional] 
**Sort** | Pointer to [**QueryServicesRequestSort**](QueryServicesRequestSort.md) |  | [optional] 
**StatsWindow** | Pointer to **string** | Lookback used to select the services returned, which are those with trace span metrics in it, and to compute the trace request rate and trace error rate fields and the span-call totals that pick each service&#39;s primary environment. Defaults to &#x60;1h&#x60;. | [optional] 
**Limit** | Pointer to **int32** | Maximum number of items to return in this page. Omit to use the public API default of 100. | [optional] [default to 100]
**Offset** | Pointer to **int32** | Zero-based index of the first matching item to return. Increase it with &#x60;limit&#x60; to request later pages. If &#x60;limit&#x60; is provided without &#x60;offset&#x60;, the offset defaults to 0. | [optional] 
**ClusterId** | Pointer to **string** | Telemetry cluster to target. Omit to use the first cluster in the organization; this does not query all clusters. | [optional] 

## Methods

### NewQueryServicesRequest

`func NewQueryServicesRequest() *QueryServicesRequest`

NewQueryServicesRequest instantiates a new QueryServicesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryServicesRequestWithDefaults

`func NewQueryServicesRequestWithDefaults() *QueryServicesRequest`

NewQueryServicesRequestWithDefaults instantiates a new QueryServicesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *QueryServicesRequest) GetFilters() QueryServicesRequestFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *QueryServicesRequest) GetFiltersOk() (*QueryServicesRequestFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *QueryServicesRequest) SetFilters(v QueryServicesRequestFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *QueryServicesRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSort

`func (o *QueryServicesRequest) GetSort() QueryServicesRequestSort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *QueryServicesRequest) GetSortOk() (*QueryServicesRequestSort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *QueryServicesRequest) SetSort(v QueryServicesRequestSort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *QueryServicesRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetStatsWindow

`func (o *QueryServicesRequest) GetStatsWindow() string`

GetStatsWindow returns the StatsWindow field if non-nil, zero value otherwise.

### GetStatsWindowOk

`func (o *QueryServicesRequest) GetStatsWindowOk() (*string, bool)`

GetStatsWindowOk returns a tuple with the StatsWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsWindow

`func (o *QueryServicesRequest) SetStatsWindow(v string)`

SetStatsWindow sets StatsWindow field to given value.

### HasStatsWindow

`func (o *QueryServicesRequest) HasStatsWindow() bool`

HasStatsWindow returns a boolean if a field has been set.

### GetLimit

`func (o *QueryServicesRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *QueryServicesRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *QueryServicesRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *QueryServicesRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *QueryServicesRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *QueryServicesRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *QueryServicesRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *QueryServicesRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetClusterId

`func (o *QueryServicesRequest) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *QueryServicesRequest) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *QueryServicesRequest) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *QueryServicesRequest) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


