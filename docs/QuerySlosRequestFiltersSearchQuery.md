# QuerySlosRequestFiltersSearchQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** | Substring matched case-insensitively against the SLO name or SLO ID. | 
**Exclude** | Pointer to **bool** | If true, exclude rows matching this value instead of including them | [optional] 

## Methods

### NewQuerySlosRequestFiltersSearchQuery

`func NewQuerySlosRequestFiltersSearchQuery(value string, ) *QuerySlosRequestFiltersSearchQuery`

NewQuerySlosRequestFiltersSearchQuery instantiates a new QuerySlosRequestFiltersSearchQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuerySlosRequestFiltersSearchQueryWithDefaults

`func NewQuerySlosRequestFiltersSearchQueryWithDefaults() *QuerySlosRequestFiltersSearchQuery`

NewQuerySlosRequestFiltersSearchQueryWithDefaults instantiates a new QuerySlosRequestFiltersSearchQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *QuerySlosRequestFiltersSearchQuery) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *QuerySlosRequestFiltersSearchQuery) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *QuerySlosRequestFiltersSearchQuery) SetValue(v string)`

SetValue sets Value field to given value.


### GetExclude

`func (o *QuerySlosRequestFiltersSearchQuery) GetExclude() bool`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *QuerySlosRequestFiltersSearchQuery) GetExcludeOk() (*bool, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *QuerySlosRequestFiltersSearchQuery) SetExclude(v bool)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *QuerySlosRequestFiltersSearchQuery) HasExclude() bool`

HasExclude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


