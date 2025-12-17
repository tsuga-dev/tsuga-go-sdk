# MonitorConfigurationMetricQueriesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Filter** | **string** |  | 
**Aggregate** | [**MetricAggregate**](MetricAggregate.md) |  | 
**ValueIfNoData** | Pointer to **string** |  | [optional] 

## Methods

### NewMonitorConfigurationMetricQueriesInner

`func NewMonitorConfigurationMetricQueriesInner(name string, filter string, aggregate MetricAggregate, ) *MonitorConfigurationMetricQueriesInner`

NewMonitorConfigurationMetricQueriesInner instantiates a new MonitorConfigurationMetricQueriesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorConfigurationMetricQueriesInnerWithDefaults

`func NewMonitorConfigurationMetricQueriesInnerWithDefaults() *MonitorConfigurationMetricQueriesInner`

NewMonitorConfigurationMetricQueriesInnerWithDefaults instantiates a new MonitorConfigurationMetricQueriesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MonitorConfigurationMetricQueriesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonitorConfigurationMetricQueriesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MonitorConfigurationMetricQueriesInner) SetName(v string)`

SetName sets Name field to given value.


### GetFilter

`func (o *MonitorConfigurationMetricQueriesInner) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MonitorConfigurationMetricQueriesInner) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *MonitorConfigurationMetricQueriesInner) SetFilter(v string)`

SetFilter sets Filter field to given value.


### GetAggregate

`func (o *MonitorConfigurationMetricQueriesInner) GetAggregate() MetricAggregate`

GetAggregate returns the Aggregate field if non-nil, zero value otherwise.

### GetAggregateOk

`func (o *MonitorConfigurationMetricQueriesInner) GetAggregateOk() (*MetricAggregate, bool)`

GetAggregateOk returns a tuple with the Aggregate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregate

`func (o *MonitorConfigurationMetricQueriesInner) SetAggregate(v MetricAggregate)`

SetAggregate sets Aggregate field to given value.


### GetValueIfNoData

`func (o *MonitorConfigurationMetricQueriesInner) GetValueIfNoData() string`

GetValueIfNoData returns the ValueIfNoData field if non-nil, zero value otherwise.

### GetValueIfNoDataOk

`func (o *MonitorConfigurationMetricQueriesInner) GetValueIfNoDataOk() (*string, bool)`

GetValueIfNoDataOk returns a tuple with the ValueIfNoData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueIfNoData

`func (o *MonitorConfigurationMetricQueriesInner) SetValueIfNoData(v string)`

SetValueIfNoData sets ValueIfNoData field to given value.

### HasValueIfNoData

`func (o *MonitorConfigurationMetricQueriesInner) HasValueIfNoData() bool`

HasValueIfNoData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


