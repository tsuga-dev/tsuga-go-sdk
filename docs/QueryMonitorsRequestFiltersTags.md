# QueryMonitorsRequestFiltersTags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | [**[]Tag**](Tag.md) |  | 
**Exclude** | Pointer to **bool** | If true, exclude rows matching these values instead of including them | [optional] 

## Methods

### NewQueryMonitorsRequestFiltersTags

`func NewQueryMonitorsRequestFiltersTags(values []Tag, ) *QueryMonitorsRequestFiltersTags`

NewQueryMonitorsRequestFiltersTags instantiates a new QueryMonitorsRequestFiltersTags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryMonitorsRequestFiltersTagsWithDefaults

`func NewQueryMonitorsRequestFiltersTagsWithDefaults() *QueryMonitorsRequestFiltersTags`

NewQueryMonitorsRequestFiltersTagsWithDefaults instantiates a new QueryMonitorsRequestFiltersTags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *QueryMonitorsRequestFiltersTags) GetValues() []Tag`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *QueryMonitorsRequestFiltersTags) GetValuesOk() (*[]Tag, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *QueryMonitorsRequestFiltersTags) SetValues(v []Tag)`

SetValues sets Values field to given value.


### GetExclude

`func (o *QueryMonitorsRequestFiltersTags) GetExclude() bool`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *QueryMonitorsRequestFiltersTags) GetExcludeOk() (*bool, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *QueryMonitorsRequestFiltersTags) SetExclude(v bool)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *QueryMonitorsRequestFiltersTags) HasExclude() bool`

HasExclude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


