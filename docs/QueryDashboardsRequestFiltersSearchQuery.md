# QueryDashboardsRequestFiltersSearchQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** |  | 
**Exclude** | Pointer to **bool** | If true, exclude dashboards matching this value instead of including them | [optional] 

## Methods

### NewQueryDashboardsRequestFiltersSearchQuery

`func NewQueryDashboardsRequestFiltersSearchQuery(value string, ) *QueryDashboardsRequestFiltersSearchQuery`

NewQueryDashboardsRequestFiltersSearchQuery instantiates a new QueryDashboardsRequestFiltersSearchQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryDashboardsRequestFiltersSearchQueryWithDefaults

`func NewQueryDashboardsRequestFiltersSearchQueryWithDefaults() *QueryDashboardsRequestFiltersSearchQuery`

NewQueryDashboardsRequestFiltersSearchQueryWithDefaults instantiates a new QueryDashboardsRequestFiltersSearchQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *QueryDashboardsRequestFiltersSearchQuery) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *QueryDashboardsRequestFiltersSearchQuery) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *QueryDashboardsRequestFiltersSearchQuery) SetValue(v string)`

SetValue sets Value field to given value.


### GetExclude

`func (o *QueryDashboardsRequestFiltersSearchQuery) GetExclude() bool`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *QueryDashboardsRequestFiltersSearchQuery) GetExcludeOk() (*bool, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *QueryDashboardsRequestFiltersSearchQuery) SetExclude(v bool)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *QueryDashboardsRequestFiltersSearchQuery) HasExclude() bool`

HasExclude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


