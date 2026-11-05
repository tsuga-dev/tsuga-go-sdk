# DashboardSummaryFiltersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Dashboard-wide filter key, usually a telemetry attribute name. | 
**Values** | **[]string** | Allowed values for this dashboard-wide filter. | 
**Exclude** | Pointer to **bool** | If true, widgets exclude telemetry whose value matches this filter instead of including it (is-not). | [optional] 

## Methods

### NewDashboardSummaryFiltersInner

`func NewDashboardSummaryFiltersInner(key string, values []string, ) *DashboardSummaryFiltersInner`

NewDashboardSummaryFiltersInner instantiates a new DashboardSummaryFiltersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardSummaryFiltersInnerWithDefaults

`func NewDashboardSummaryFiltersInnerWithDefaults() *DashboardSummaryFiltersInner`

NewDashboardSummaryFiltersInnerWithDefaults instantiates a new DashboardSummaryFiltersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *DashboardSummaryFiltersInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DashboardSummaryFiltersInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DashboardSummaryFiltersInner) SetKey(v string)`

SetKey sets Key field to given value.


### GetValues

`func (o *DashboardSummaryFiltersInner) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *DashboardSummaryFiltersInner) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *DashboardSummaryFiltersInner) SetValues(v []string)`

SetValues sets Values field to given value.


### GetExclude

`func (o *DashboardSummaryFiltersInner) GetExclude() bool`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *DashboardSummaryFiltersInner) GetExcludeOk() (*bool, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *DashboardSummaryFiltersInner) SetExclude(v bool)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *DashboardSummaryFiltersInner) HasExclude() bool`

HasExclude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


