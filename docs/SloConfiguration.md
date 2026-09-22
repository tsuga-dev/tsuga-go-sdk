# SloConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Configuration kind. Set to &#x60;event&#x60; for SLOs that compute SLI from good and total query formulas over the rolling timeframe. | 
**DataSource** | **string** | Telemetry source queried by this aggregation: &#x60;logs&#x60;, &#x60;metrics&#x60;, &#x60;traces&#x60;, or &#x60;rum&#x60;. | 
**GoodQuery** | [**SloQueryFormula1**](SloQueryFormula1.md) |  | 
**TotalQuery** | [**SloQueryFormula1**](SloQueryFormula1.md) |  | 
**GroupByFields** | Pointer to [**[]AggregationGroupBy1**](AggregationGroupBy1.md) | Nested grouping levels applied to the results, outermost first (e.g. group by service, then by level within each service). Each level splits results further, so the response contains one result per unique combination of group values. | [optional] 
**NoDataBehavior** | **string** | How to treat time buckets with no data: good (count as meeting the SLO), bad (count as breaching it), or ignore (exclude from the error budget) | 
**Query** | [**SloQueryFormula1**](SloQueryFormula1.md) |  | 
**SliceSizeMinutes** | **int32** | Length of each time-SLO evaluation slice in minutes. Required for &#x60;type: time&#x60;. | 
**Threshold** | [**TimeSliceSloConfigurationThreshold**](TimeSliceSloConfigurationThreshold.md) |  | 

## Methods

### NewSloConfiguration

`func NewSloConfiguration(type_ string, dataSource string, goodQuery SloQueryFormula1, totalQuery SloQueryFormula1, noDataBehavior string, query SloQueryFormula1, sliceSizeMinutes int32, threshold TimeSliceSloConfigurationThreshold, ) *SloConfiguration`

NewSloConfiguration instantiates a new SloConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSloConfigurationWithDefaults

`func NewSloConfigurationWithDefaults() *SloConfiguration`

NewSloConfigurationWithDefaults instantiates a new SloConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SloConfiguration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SloConfiguration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SloConfiguration) SetType(v string)`

SetType sets Type field to given value.


### GetDataSource

`func (o *SloConfiguration) GetDataSource() string`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *SloConfiguration) GetDataSourceOk() (*string, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *SloConfiguration) SetDataSource(v string)`

SetDataSource sets DataSource field to given value.


### GetGoodQuery

`func (o *SloConfiguration) GetGoodQuery() SloQueryFormula1`

GetGoodQuery returns the GoodQuery field if non-nil, zero value otherwise.

### GetGoodQueryOk

`func (o *SloConfiguration) GetGoodQueryOk() (*SloQueryFormula1, bool)`

GetGoodQueryOk returns a tuple with the GoodQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodQuery

`func (o *SloConfiguration) SetGoodQuery(v SloQueryFormula1)`

SetGoodQuery sets GoodQuery field to given value.


### GetTotalQuery

`func (o *SloConfiguration) GetTotalQuery() SloQueryFormula1`

GetTotalQuery returns the TotalQuery field if non-nil, zero value otherwise.

### GetTotalQueryOk

`func (o *SloConfiguration) GetTotalQueryOk() (*SloQueryFormula1, bool)`

GetTotalQueryOk returns a tuple with the TotalQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalQuery

`func (o *SloConfiguration) SetTotalQuery(v SloQueryFormula1)`

SetTotalQuery sets TotalQuery field to given value.


### GetGroupByFields

`func (o *SloConfiguration) GetGroupByFields() []AggregationGroupBy1`

GetGroupByFields returns the GroupByFields field if non-nil, zero value otherwise.

### GetGroupByFieldsOk

`func (o *SloConfiguration) GetGroupByFieldsOk() (*[]AggregationGroupBy1, bool)`

GetGroupByFieldsOk returns a tuple with the GroupByFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByFields

`func (o *SloConfiguration) SetGroupByFields(v []AggregationGroupBy1)`

SetGroupByFields sets GroupByFields field to given value.

### HasGroupByFields

`func (o *SloConfiguration) HasGroupByFields() bool`

HasGroupByFields returns a boolean if a field has been set.

### GetNoDataBehavior

`func (o *SloConfiguration) GetNoDataBehavior() string`

GetNoDataBehavior returns the NoDataBehavior field if non-nil, zero value otherwise.

### GetNoDataBehaviorOk

`func (o *SloConfiguration) GetNoDataBehaviorOk() (*string, bool)`

GetNoDataBehaviorOk returns a tuple with the NoDataBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDataBehavior

`func (o *SloConfiguration) SetNoDataBehavior(v string)`

SetNoDataBehavior sets NoDataBehavior field to given value.


### GetQuery

`func (o *SloConfiguration) GetQuery() SloQueryFormula1`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SloConfiguration) GetQueryOk() (*SloQueryFormula1, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SloConfiguration) SetQuery(v SloQueryFormula1)`

SetQuery sets Query field to given value.


### GetSliceSizeMinutes

`func (o *SloConfiguration) GetSliceSizeMinutes() int32`

GetSliceSizeMinutes returns the SliceSizeMinutes field if non-nil, zero value otherwise.

### GetSliceSizeMinutesOk

`func (o *SloConfiguration) GetSliceSizeMinutesOk() (*int32, bool)`

GetSliceSizeMinutesOk returns a tuple with the SliceSizeMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceSizeMinutes

`func (o *SloConfiguration) SetSliceSizeMinutes(v int32)`

SetSliceSizeMinutes sets SliceSizeMinutes field to given value.


### GetThreshold

`func (o *SloConfiguration) GetThreshold() TimeSliceSloConfigurationThreshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *SloConfiguration) GetThresholdOk() (*TimeSliceSloConfigurationThreshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *SloConfiguration) SetThreshold(v TimeSliceSloConfigurationThreshold)`

SetThreshold sets Threshold field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


