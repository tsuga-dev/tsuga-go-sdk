# QueryDashboardsRequestFiltersWidgetQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** | Substring matched case-insensitively against every text value in the widget queries (filters, aggregate types such as &#x60;percentile&#x60;, fields, metric names and PromQL), in the group-by fields and in the dashboard-wide filters applied to every widget. Display settings, such as the widget type, legend or units, and note text are not matched. Omit the whole &#x60;widgetQuery&#x60; filter to skip it. | 
**Exclude** | Pointer to **bool** | If true, exclude dashboards with a matching widget or dashboard-wide filter instead of including them | [optional] 

## Methods

### NewQueryDashboardsRequestFiltersWidgetQuery

`func NewQueryDashboardsRequestFiltersWidgetQuery(value string, ) *QueryDashboardsRequestFiltersWidgetQuery`

NewQueryDashboardsRequestFiltersWidgetQuery instantiates a new QueryDashboardsRequestFiltersWidgetQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryDashboardsRequestFiltersWidgetQueryWithDefaults

`func NewQueryDashboardsRequestFiltersWidgetQueryWithDefaults() *QueryDashboardsRequestFiltersWidgetQuery`

NewQueryDashboardsRequestFiltersWidgetQueryWithDefaults instantiates a new QueryDashboardsRequestFiltersWidgetQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *QueryDashboardsRequestFiltersWidgetQuery) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *QueryDashboardsRequestFiltersWidgetQuery) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *QueryDashboardsRequestFiltersWidgetQuery) SetValue(v string)`

SetValue sets Value field to given value.


### GetExclude

`func (o *QueryDashboardsRequestFiltersWidgetQuery) GetExclude() bool`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *QueryDashboardsRequestFiltersWidgetQuery) GetExcludeOk() (*bool, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *QueryDashboardsRequestFiltersWidgetQuery) SetExclude(v bool)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *QueryDashboardsRequestFiltersWidgetQuery) HasExclude() bool`

HasExclude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


