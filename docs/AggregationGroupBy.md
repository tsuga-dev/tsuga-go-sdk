# AggregationGroupBy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | **[]string** | Telemetry attribute names used to group aggregation results. | 
**Limit** | **float32** | Maximum number of distinct values of &#x60;fields&#x60; returned for this grouping level, ranked by aggregate value. Groups beyond the limit are truncated, not merged into an \&quot;other\&quot; group. | 
**SortOrder** | Pointer to **string** | Sort direction applied to groups: ascending or descending. | [optional] 
**ReplaceNullWith** | Pointer to **string** | Value used to group documents that have no value for a grouped field. | [optional] 

## Methods

### NewAggregationGroupBy

`func NewAggregationGroupBy(fields []string, limit float32, ) *AggregationGroupBy`

NewAggregationGroupBy instantiates a new AggregationGroupBy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregationGroupByWithDefaults

`func NewAggregationGroupByWithDefaults() *AggregationGroupBy`

NewAggregationGroupByWithDefaults instantiates a new AggregationGroupBy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *AggregationGroupBy) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AggregationGroupBy) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AggregationGroupBy) SetFields(v []string)`

SetFields sets Fields field to given value.


### GetLimit

`func (o *AggregationGroupBy) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AggregationGroupBy) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AggregationGroupBy) SetLimit(v float32)`

SetLimit sets Limit field to given value.


### GetSortOrder

`func (o *AggregationGroupBy) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *AggregationGroupBy) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *AggregationGroupBy) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *AggregationGroupBy) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetReplaceNullWith

`func (o *AggregationGroupBy) GetReplaceNullWith() string`

GetReplaceNullWith returns the ReplaceNullWith field if non-nil, zero value otherwise.

### GetReplaceNullWithOk

`func (o *AggregationGroupBy) GetReplaceNullWithOk() (*string, bool)`

GetReplaceNullWithOk returns a tuple with the ReplaceNullWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceNullWith

`func (o *AggregationGroupBy) SetReplaceNullWith(v string)`

SetReplaceNullWith sets ReplaceNullWith field to given value.

### HasReplaceNullWith

`func (o *AggregationGroupBy) HasReplaceNullWith() bool`

HasReplaceNullWith returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


