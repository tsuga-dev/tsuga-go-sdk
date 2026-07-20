# UpdateDashboardRequestFiltersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Dashboard-wide filter key, usually a telemetry attribute name. | 
**Values** | **[]string** | Allowed values for this dashboard-wide filter. Set by the dashboard author. | 
**Exclude** | Pointer to **bool** | If true, widgets exclude telemetry whose value matches this filter instead of including it (is-not). | [optional] 

## Methods

### NewUpdateDashboardRequestFiltersInner

`func NewUpdateDashboardRequestFiltersInner(key string, values []string, ) *UpdateDashboardRequestFiltersInner`

NewUpdateDashboardRequestFiltersInner instantiates a new UpdateDashboardRequestFiltersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDashboardRequestFiltersInnerWithDefaults

`func NewUpdateDashboardRequestFiltersInnerWithDefaults() *UpdateDashboardRequestFiltersInner`

NewUpdateDashboardRequestFiltersInnerWithDefaults instantiates a new UpdateDashboardRequestFiltersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *UpdateDashboardRequestFiltersInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateDashboardRequestFiltersInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateDashboardRequestFiltersInner) SetKey(v string)`

SetKey sets Key field to given value.


### GetValues

`func (o *UpdateDashboardRequestFiltersInner) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *UpdateDashboardRequestFiltersInner) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *UpdateDashboardRequestFiltersInner) SetValues(v []string)`

SetValues sets Values field to given value.


### GetExclude

`func (o *UpdateDashboardRequestFiltersInner) GetExclude() bool`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *UpdateDashboardRequestFiltersInner) GetExcludeOk() (*bool, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *UpdateDashboardRequestFiltersInner) SetExclude(v bool)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *UpdateDashboardRequestFiltersInner) HasExclude() bool`

HasExclude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


