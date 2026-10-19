# GetInterestingFieldsResponseDataLogsInnerUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupBy** | **float32** | Number of distinct readable matching artifacts that use this field as a group-by or split dimension. | 
**Filter** | **float32** | Number of distinct readable matching artifacts that filter on this field in widget queries, monitor queries, or dashboard-wide filters. | 
**Measure** | **float32** | Number of distinct readable matching artifacts that aggregate this field as &#x60;aggregate.field&#x60;. | 
**Column** | **float32** | Number of distinct readable matching artifacts that pin this field as a list-widget column. | 

## Methods

### NewGetInterestingFieldsResponseDataLogsInnerUsage

`func NewGetInterestingFieldsResponseDataLogsInnerUsage(groupBy float32, filter float32, measure float32, column float32, ) *GetInterestingFieldsResponseDataLogsInnerUsage`

NewGetInterestingFieldsResponseDataLogsInnerUsage instantiates a new GetInterestingFieldsResponseDataLogsInnerUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInterestingFieldsResponseDataLogsInnerUsageWithDefaults

`func NewGetInterestingFieldsResponseDataLogsInnerUsageWithDefaults() *GetInterestingFieldsResponseDataLogsInnerUsage`

NewGetInterestingFieldsResponseDataLogsInnerUsageWithDefaults instantiates a new GetInterestingFieldsResponseDataLogsInnerUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupBy

`func (o *GetInterestingFieldsResponseDataLogsInnerUsage) GetGroupBy() float32`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *GetInterestingFieldsResponseDataLogsInnerUsage) GetGroupByOk() (*float32, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *GetInterestingFieldsResponseDataLogsInnerUsage) SetGroupBy(v float32)`

SetGroupBy sets GroupBy field to given value.


### GetFilter

`func (o *GetInterestingFieldsResponseDataLogsInnerUsage) GetFilter() float32`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *GetInterestingFieldsResponseDataLogsInnerUsage) GetFilterOk() (*float32, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *GetInterestingFieldsResponseDataLogsInnerUsage) SetFilter(v float32)`

SetFilter sets Filter field to given value.


### GetMeasure

`func (o *GetInterestingFieldsResponseDataLogsInnerUsage) GetMeasure() float32`

GetMeasure returns the Measure field if non-nil, zero value otherwise.

### GetMeasureOk

`func (o *GetInterestingFieldsResponseDataLogsInnerUsage) GetMeasureOk() (*float32, bool)`

GetMeasureOk returns a tuple with the Measure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasure

`func (o *GetInterestingFieldsResponseDataLogsInnerUsage) SetMeasure(v float32)`

SetMeasure sets Measure field to given value.


### GetColumn

`func (o *GetInterestingFieldsResponseDataLogsInnerUsage) GetColumn() float32`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *GetInterestingFieldsResponseDataLogsInnerUsage) GetColumnOk() (*float32, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *GetInterestingFieldsResponseDataLogsInnerUsage) SetColumn(v float32)`

SetColumn sets Column field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


