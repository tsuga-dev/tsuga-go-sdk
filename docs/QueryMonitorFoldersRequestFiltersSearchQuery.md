# QueryMonitorFoldersRequestFiltersSearchQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** | Substring matched case-insensitively against monitor folder names or folder IDs. Omit the whole &#x60;searchQuery&#x60; filter to skip text filtering. | 
**Exclude** | Pointer to **bool** | If true, exclude monitor folders matching this value instead of including them | [optional] 

## Methods

### NewQueryMonitorFoldersRequestFiltersSearchQuery

`func NewQueryMonitorFoldersRequestFiltersSearchQuery(value string, ) *QueryMonitorFoldersRequestFiltersSearchQuery`

NewQueryMonitorFoldersRequestFiltersSearchQuery instantiates a new QueryMonitorFoldersRequestFiltersSearchQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryMonitorFoldersRequestFiltersSearchQueryWithDefaults

`func NewQueryMonitorFoldersRequestFiltersSearchQueryWithDefaults() *QueryMonitorFoldersRequestFiltersSearchQuery`

NewQueryMonitorFoldersRequestFiltersSearchQueryWithDefaults instantiates a new QueryMonitorFoldersRequestFiltersSearchQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *QueryMonitorFoldersRequestFiltersSearchQuery) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *QueryMonitorFoldersRequestFiltersSearchQuery) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *QueryMonitorFoldersRequestFiltersSearchQuery) SetValue(v string)`

SetValue sets Value field to given value.


### GetExclude

`func (o *QueryMonitorFoldersRequestFiltersSearchQuery) GetExclude() bool`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *QueryMonitorFoldersRequestFiltersSearchQuery) GetExcludeOk() (*bool, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *QueryMonitorFoldersRequestFiltersSearchQuery) SetExclude(v bool)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *QueryMonitorFoldersRequestFiltersSearchQuery) HasExclude() bool`

HasExclude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


