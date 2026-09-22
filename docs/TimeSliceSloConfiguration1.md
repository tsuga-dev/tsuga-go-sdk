# TimeSliceSloConfiguration1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Configuration kind. Set to &#x60;time&#x60; for SLOs that evaluate a query against a threshold in fixed-size time slices. | 
**DataSource** | **string** | Telemetry source queried by this aggregation: &#x60;logs&#x60;, &#x60;metrics&#x60;, &#x60;traces&#x60;, or &#x60;rum&#x60;. | 
**Query** | [**SloQueryFormula1**](SloQueryFormula1.md) |  | 
**SliceSizeMinutes** | **int32** | Length of each time-SLO evaluation slice in minutes. Required for &#x60;type: time&#x60;. | 
**Threshold** | [**TimeSliceSloConfigurationThreshold**](TimeSliceSloConfigurationThreshold.md) |  | 
**GroupByFields** | Pointer to [**[]AggregationGroupBy1**](AggregationGroupBy1.md) | Nested grouping levels applied to the results, outermost first (e.g. group by service, then by level within each service). Each level splits results further, so the response contains one result per unique combination of group values. | [optional] 
**NoDataBehavior** | **string** | How to treat time buckets with no data: good (count as meeting the SLO), bad (count as breaching it), or ignore (exclude from the error budget) | 

## Methods

### NewTimeSliceSloConfiguration1

`func NewTimeSliceSloConfiguration1(type_ string, dataSource string, query SloQueryFormula1, sliceSizeMinutes int32, threshold TimeSliceSloConfigurationThreshold, noDataBehavior string, ) *TimeSliceSloConfiguration1`

NewTimeSliceSloConfiguration1 instantiates a new TimeSliceSloConfiguration1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSliceSloConfiguration1WithDefaults

`func NewTimeSliceSloConfiguration1WithDefaults() *TimeSliceSloConfiguration1`

NewTimeSliceSloConfiguration1WithDefaults instantiates a new TimeSliceSloConfiguration1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TimeSliceSloConfiguration1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimeSliceSloConfiguration1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimeSliceSloConfiguration1) SetType(v string)`

SetType sets Type field to given value.


### GetDataSource

`func (o *TimeSliceSloConfiguration1) GetDataSource() string`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *TimeSliceSloConfiguration1) GetDataSourceOk() (*string, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *TimeSliceSloConfiguration1) SetDataSource(v string)`

SetDataSource sets DataSource field to given value.


### GetQuery

`func (o *TimeSliceSloConfiguration1) GetQuery() SloQueryFormula1`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *TimeSliceSloConfiguration1) GetQueryOk() (*SloQueryFormula1, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *TimeSliceSloConfiguration1) SetQuery(v SloQueryFormula1)`

SetQuery sets Query field to given value.


### GetSliceSizeMinutes

`func (o *TimeSliceSloConfiguration1) GetSliceSizeMinutes() int32`

GetSliceSizeMinutes returns the SliceSizeMinutes field if non-nil, zero value otherwise.

### GetSliceSizeMinutesOk

`func (o *TimeSliceSloConfiguration1) GetSliceSizeMinutesOk() (*int32, bool)`

GetSliceSizeMinutesOk returns a tuple with the SliceSizeMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceSizeMinutes

`func (o *TimeSliceSloConfiguration1) SetSliceSizeMinutes(v int32)`

SetSliceSizeMinutes sets SliceSizeMinutes field to given value.


### GetThreshold

`func (o *TimeSliceSloConfiguration1) GetThreshold() TimeSliceSloConfigurationThreshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *TimeSliceSloConfiguration1) GetThresholdOk() (*TimeSliceSloConfigurationThreshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *TimeSliceSloConfiguration1) SetThreshold(v TimeSliceSloConfigurationThreshold)`

SetThreshold sets Threshold field to given value.


### GetGroupByFields

`func (o *TimeSliceSloConfiguration1) GetGroupByFields() []AggregationGroupBy1`

GetGroupByFields returns the GroupByFields field if non-nil, zero value otherwise.

### GetGroupByFieldsOk

`func (o *TimeSliceSloConfiguration1) GetGroupByFieldsOk() (*[]AggregationGroupBy1, bool)`

GetGroupByFieldsOk returns a tuple with the GroupByFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByFields

`func (o *TimeSliceSloConfiguration1) SetGroupByFields(v []AggregationGroupBy1)`

SetGroupByFields sets GroupByFields field to given value.

### HasGroupByFields

`func (o *TimeSliceSloConfiguration1) HasGroupByFields() bool`

HasGroupByFields returns a boolean if a field has been set.

### GetNoDataBehavior

`func (o *TimeSliceSloConfiguration1) GetNoDataBehavior() string`

GetNoDataBehavior returns the NoDataBehavior field if non-nil, zero value otherwise.

### GetNoDataBehaviorOk

`func (o *TimeSliceSloConfiguration1) GetNoDataBehaviorOk() (*string, bool)`

GetNoDataBehaviorOk returns a tuple with the NoDataBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDataBehavior

`func (o *TimeSliceSloConfiguration1) SetNoDataBehavior(v string)`

SetNoDataBehavior sets NoDataBehavior field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


