# MonitorConfigurationLogQueriesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Filter** | **string** |  | 
**Aggregate** | [**LogAggregate**](LogAggregate.md) |  | 
**ValueIfNoData** | Pointer to **string** |  | [optional] 

## Methods

### NewMonitorConfigurationLogQueriesInner

`func NewMonitorConfigurationLogQueriesInner(name string, filter string, aggregate LogAggregate, ) *MonitorConfigurationLogQueriesInner`

NewMonitorConfigurationLogQueriesInner instantiates a new MonitorConfigurationLogQueriesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorConfigurationLogQueriesInnerWithDefaults

`func NewMonitorConfigurationLogQueriesInnerWithDefaults() *MonitorConfigurationLogQueriesInner`

NewMonitorConfigurationLogQueriesInnerWithDefaults instantiates a new MonitorConfigurationLogQueriesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MonitorConfigurationLogQueriesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonitorConfigurationLogQueriesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MonitorConfigurationLogQueriesInner) SetName(v string)`

SetName sets Name field to given value.


### GetFilter

`func (o *MonitorConfigurationLogQueriesInner) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MonitorConfigurationLogQueriesInner) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *MonitorConfigurationLogQueriesInner) SetFilter(v string)`

SetFilter sets Filter field to given value.


### GetAggregate

`func (o *MonitorConfigurationLogQueriesInner) GetAggregate() LogAggregate`

GetAggregate returns the Aggregate field if non-nil, zero value otherwise.

### GetAggregateOk

`func (o *MonitorConfigurationLogQueriesInner) GetAggregateOk() (*LogAggregate, bool)`

GetAggregateOk returns a tuple with the Aggregate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregate

`func (o *MonitorConfigurationLogQueriesInner) SetAggregate(v LogAggregate)`

SetAggregate sets Aggregate field to given value.


### GetValueIfNoData

`func (o *MonitorConfigurationLogQueriesInner) GetValueIfNoData() string`

GetValueIfNoData returns the ValueIfNoData field if non-nil, zero value otherwise.

### GetValueIfNoDataOk

`func (o *MonitorConfigurationLogQueriesInner) GetValueIfNoDataOk() (*string, bool)`

GetValueIfNoDataOk returns a tuple with the ValueIfNoData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueIfNoData

`func (o *MonitorConfigurationLogQueriesInner) SetValueIfNoData(v string)`

SetValueIfNoData sets ValueIfNoData field to given value.

### HasValueIfNoData

`func (o *MonitorConfigurationLogQueriesInner) HasValueIfNoData() bool`

HasValueIfNoData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


