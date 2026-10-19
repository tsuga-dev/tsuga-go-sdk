# GetInterestingFieldsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | [**[]GetInterestingFieldsResponseDataLogsInner**](GetInterestingFieldsResponseDataLogsInner.md) | Top log fields for the requested scope, ranked by score and mined from readable logs-sourced dashboards and monitors. | 
**Traces** | [**[]GetInterestingFieldsResponseDataLogsInner**](GetInterestingFieldsResponseDataLogsInner.md) | Top trace fields for the requested scope, ranked by score and mined from readable traces-sourced dashboards and monitors. | 
**Metrics** | [**[]GetInterestingFieldsResponseDataLogsInner**](GetInterestingFieldsResponseDataLogsInner.md) | Top metric names and metric-attribute fields for the requested scope, ranked by score and mined from readable metrics-sourced dashboards and monitors. | 
**DashboardsMatched** | **float32** | Number of dashboards that matched the requested scope and that the caller can read. Dashboards the caller cannot read are removed before field rankings are returned. | 
**MonitorsMatched** | **float32** | Number of monitors that matched the requested scope and that the caller can read. Monitors the caller cannot read are removed before field rankings are returned. | 

## Methods

### NewGetInterestingFieldsResponseData

`func NewGetInterestingFieldsResponseData(logs []GetInterestingFieldsResponseDataLogsInner, traces []GetInterestingFieldsResponseDataLogsInner, metrics []GetInterestingFieldsResponseDataLogsInner, dashboardsMatched float32, monitorsMatched float32, ) *GetInterestingFieldsResponseData`

NewGetInterestingFieldsResponseData instantiates a new GetInterestingFieldsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInterestingFieldsResponseDataWithDefaults

`func NewGetInterestingFieldsResponseDataWithDefaults() *GetInterestingFieldsResponseData`

NewGetInterestingFieldsResponseDataWithDefaults instantiates a new GetInterestingFieldsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *GetInterestingFieldsResponseData) GetLogs() []GetInterestingFieldsResponseDataLogsInner`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *GetInterestingFieldsResponseData) GetLogsOk() (*[]GetInterestingFieldsResponseDataLogsInner, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *GetInterestingFieldsResponseData) SetLogs(v []GetInterestingFieldsResponseDataLogsInner)`

SetLogs sets Logs field to given value.


### GetTraces

`func (o *GetInterestingFieldsResponseData) GetTraces() []GetInterestingFieldsResponseDataLogsInner`

GetTraces returns the Traces field if non-nil, zero value otherwise.

### GetTracesOk

`func (o *GetInterestingFieldsResponseData) GetTracesOk() (*[]GetInterestingFieldsResponseDataLogsInner, bool)`

GetTracesOk returns a tuple with the Traces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraces

`func (o *GetInterestingFieldsResponseData) SetTraces(v []GetInterestingFieldsResponseDataLogsInner)`

SetTraces sets Traces field to given value.


### GetMetrics

`func (o *GetInterestingFieldsResponseData) GetMetrics() []GetInterestingFieldsResponseDataLogsInner`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GetInterestingFieldsResponseData) GetMetricsOk() (*[]GetInterestingFieldsResponseDataLogsInner, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GetInterestingFieldsResponseData) SetMetrics(v []GetInterestingFieldsResponseDataLogsInner)`

SetMetrics sets Metrics field to given value.


### GetDashboardsMatched

`func (o *GetInterestingFieldsResponseData) GetDashboardsMatched() float32`

GetDashboardsMatched returns the DashboardsMatched field if non-nil, zero value otherwise.

### GetDashboardsMatchedOk

`func (o *GetInterestingFieldsResponseData) GetDashboardsMatchedOk() (*float32, bool)`

GetDashboardsMatchedOk returns a tuple with the DashboardsMatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardsMatched

`func (o *GetInterestingFieldsResponseData) SetDashboardsMatched(v float32)`

SetDashboardsMatched sets DashboardsMatched field to given value.


### GetMonitorsMatched

`func (o *GetInterestingFieldsResponseData) GetMonitorsMatched() float32`

GetMonitorsMatched returns the MonitorsMatched field if non-nil, zero value otherwise.

### GetMonitorsMatchedOk

`func (o *GetInterestingFieldsResponseData) GetMonitorsMatchedOk() (*float32, bool)`

GetMonitorsMatchedOk returns a tuple with the MonitorsMatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorsMatched

`func (o *GetInterestingFieldsResponseData) SetMonitorsMatched(v float32)`

SetMonitorsMatched sets MonitorsMatched field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


