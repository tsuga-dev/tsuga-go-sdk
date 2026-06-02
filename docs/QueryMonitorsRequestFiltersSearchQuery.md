# QueryMonitorsRequestFiltersSearchQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** |  | 
**Exclude** | Pointer to **bool** | If true, exclude monitors matching this value instead of including them | [optional] 

## Methods

### NewQueryMonitorsRequestFiltersSearchQuery

`func NewQueryMonitorsRequestFiltersSearchQuery(value string, ) *QueryMonitorsRequestFiltersSearchQuery`

NewQueryMonitorsRequestFiltersSearchQuery instantiates a new QueryMonitorsRequestFiltersSearchQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryMonitorsRequestFiltersSearchQueryWithDefaults

`func NewQueryMonitorsRequestFiltersSearchQueryWithDefaults() *QueryMonitorsRequestFiltersSearchQuery`

NewQueryMonitorsRequestFiltersSearchQueryWithDefaults instantiates a new QueryMonitorsRequestFiltersSearchQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *QueryMonitorsRequestFiltersSearchQuery) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *QueryMonitorsRequestFiltersSearchQuery) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *QueryMonitorsRequestFiltersSearchQuery) SetValue(v string)`

SetValue sets Value field to given value.


### GetExclude

`func (o *QueryMonitorsRequestFiltersSearchQuery) GetExclude() bool`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *QueryMonitorsRequestFiltersSearchQuery) GetExcludeOk() (*bool, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *QueryMonitorsRequestFiltersSearchQuery) SetExclude(v bool)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *QueryMonitorsRequestFiltersSearchQuery) HasExclude() bool`

HasExclude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


