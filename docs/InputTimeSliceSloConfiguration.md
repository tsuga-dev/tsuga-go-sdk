# InputTimeSliceSloConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Configuration kind. Set to &#x60;time&#x60; for SLOs that evaluate a query against a threshold in fixed-size time slices. | 
**DataSource** | **string** | Telemetry source queried by this aggregation: &#x60;logs&#x60;, &#x60;metrics&#x60;, &#x60;traces&#x60;, or &#x60;rum&#x60;. | 
**Query** | [**InputSloQueryFormula**](InputSloQueryFormula.md) |  | 
**SliceSizeMinutes** | **int32** | Length of each time-SLO evaluation slice in minutes. Required for &#x60;type: time&#x60;. Valid input is 30 through 1440. | 
**Threshold** | [**TimeSliceSloConfigurationThreshold**](TimeSliceSloConfigurationThreshold.md) |  | 
**Normalizer** | Pointer to [**InputNormalizer**](InputNormalizer.md) |  | [optional] 
**GroupByFields** | Pointer to [**[]InputAggregationGroupBy**](InputAggregationGroupBy.md) | Nested grouping levels applied to aggregation results, outermost first (e.g. group by service, then by level within each service). Each level splits results further, so the response contains one result per unique combination of group values instead of one aggregated total. Defaults to an empty array (one ungrouped result) when omitted. | [optional] 
**NoDataBehavior** | **string** | How to treat time buckets with no data: good (count as meeting the SLO), bad (count as breaching it), or ignore (exclude from the error budget) | 

## Methods

### NewInputTimeSliceSloConfiguration

`func NewInputTimeSliceSloConfiguration(type_ string, dataSource string, query InputSloQueryFormula, sliceSizeMinutes int32, threshold TimeSliceSloConfigurationThreshold, noDataBehavior string, ) *InputTimeSliceSloConfiguration`

NewInputTimeSliceSloConfiguration instantiates a new InputTimeSliceSloConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputTimeSliceSloConfigurationWithDefaults

`func NewInputTimeSliceSloConfigurationWithDefaults() *InputTimeSliceSloConfiguration`

NewInputTimeSliceSloConfigurationWithDefaults instantiates a new InputTimeSliceSloConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputTimeSliceSloConfiguration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputTimeSliceSloConfiguration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputTimeSliceSloConfiguration) SetType(v string)`

SetType sets Type field to given value.


### GetDataSource

`func (o *InputTimeSliceSloConfiguration) GetDataSource() string`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *InputTimeSliceSloConfiguration) GetDataSourceOk() (*string, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *InputTimeSliceSloConfiguration) SetDataSource(v string)`

SetDataSource sets DataSource field to given value.


### GetQuery

`func (o *InputTimeSliceSloConfiguration) GetQuery() InputSloQueryFormula`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *InputTimeSliceSloConfiguration) GetQueryOk() (*InputSloQueryFormula, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *InputTimeSliceSloConfiguration) SetQuery(v InputSloQueryFormula)`

SetQuery sets Query field to given value.


### GetSliceSizeMinutes

`func (o *InputTimeSliceSloConfiguration) GetSliceSizeMinutes() int32`

GetSliceSizeMinutes returns the SliceSizeMinutes field if non-nil, zero value otherwise.

### GetSliceSizeMinutesOk

`func (o *InputTimeSliceSloConfiguration) GetSliceSizeMinutesOk() (*int32, bool)`

GetSliceSizeMinutesOk returns a tuple with the SliceSizeMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceSizeMinutes

`func (o *InputTimeSliceSloConfiguration) SetSliceSizeMinutes(v int32)`

SetSliceSizeMinutes sets SliceSizeMinutes field to given value.


### GetThreshold

`func (o *InputTimeSliceSloConfiguration) GetThreshold() TimeSliceSloConfigurationThreshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *InputTimeSliceSloConfiguration) GetThresholdOk() (*TimeSliceSloConfigurationThreshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *InputTimeSliceSloConfiguration) SetThreshold(v TimeSliceSloConfigurationThreshold)`

SetThreshold sets Threshold field to given value.


### GetNormalizer

`func (o *InputTimeSliceSloConfiguration) GetNormalizer() InputNormalizer`

GetNormalizer returns the Normalizer field if non-nil, zero value otherwise.

### GetNormalizerOk

`func (o *InputTimeSliceSloConfiguration) GetNormalizerOk() (*InputNormalizer, bool)`

GetNormalizerOk returns a tuple with the Normalizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizer

`func (o *InputTimeSliceSloConfiguration) SetNormalizer(v InputNormalizer)`

SetNormalizer sets Normalizer field to given value.

### HasNormalizer

`func (o *InputTimeSliceSloConfiguration) HasNormalizer() bool`

HasNormalizer returns a boolean if a field has been set.

### GetGroupByFields

`func (o *InputTimeSliceSloConfiguration) GetGroupByFields() []InputAggregationGroupBy`

GetGroupByFields returns the GroupByFields field if non-nil, zero value otherwise.

### GetGroupByFieldsOk

`func (o *InputTimeSliceSloConfiguration) GetGroupByFieldsOk() (*[]InputAggregationGroupBy, bool)`

GetGroupByFieldsOk returns a tuple with the GroupByFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByFields

`func (o *InputTimeSliceSloConfiguration) SetGroupByFields(v []InputAggregationGroupBy)`

SetGroupByFields sets GroupByFields field to given value.

### HasGroupByFields

`func (o *InputTimeSliceSloConfiguration) HasGroupByFields() bool`

HasGroupByFields returns a boolean if a field has been set.

### GetNoDataBehavior

`func (o *InputTimeSliceSloConfiguration) GetNoDataBehavior() string`

GetNoDataBehavior returns the NoDataBehavior field if non-nil, zero value otherwise.

### GetNoDataBehaviorOk

`func (o *InputTimeSliceSloConfiguration) GetNoDataBehaviorOk() (*string, bool)`

GetNoDataBehaviorOk returns a tuple with the NoDataBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDataBehavior

`func (o *InputTimeSliceSloConfiguration) SetNoDataBehavior(v string)`

SetNoDataBehavior sets NoDataBehavior field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


