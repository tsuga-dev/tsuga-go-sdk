# AggregationGroupBy1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | **[]string** | Telemetry attribute name to group aggregation results by. Currently limited to one field per grouping level. | 
**Limit** | **float32** | Maximum number of distinct values of &#x60;fields&#x60; to return for this grouping level, ranked by aggregate value. Groups beyond the limit are truncated from the response, not merged into an \&quot;other\&quot; group. | 
**SortOrder** | Pointer to **string** | Sort direction applied to groups: ascending or descending. | [optional] 
**ReplaceNullWith** | Pointer to **string** | Value used to group documents that have no value for a grouped field. | [optional] 

## Methods

### NewAggregationGroupBy1

`func NewAggregationGroupBy1(fields []string, limit float32, ) *AggregationGroupBy1`

NewAggregationGroupBy1 instantiates a new AggregationGroupBy1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregationGroupBy1WithDefaults

`func NewAggregationGroupBy1WithDefaults() *AggregationGroupBy1`

NewAggregationGroupBy1WithDefaults instantiates a new AggregationGroupBy1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *AggregationGroupBy1) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AggregationGroupBy1) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AggregationGroupBy1) SetFields(v []string)`

SetFields sets Fields field to given value.


### GetLimit

`func (o *AggregationGroupBy1) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AggregationGroupBy1) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AggregationGroupBy1) SetLimit(v float32)`

SetLimit sets Limit field to given value.


### GetSortOrder

`func (o *AggregationGroupBy1) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *AggregationGroupBy1) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *AggregationGroupBy1) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *AggregationGroupBy1) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetReplaceNullWith

`func (o *AggregationGroupBy1) GetReplaceNullWith() string`

GetReplaceNullWith returns the ReplaceNullWith field if non-nil, zero value otherwise.

### GetReplaceNullWithOk

`func (o *AggregationGroupBy1) GetReplaceNullWithOk() (*string, bool)`

GetReplaceNullWithOk returns a tuple with the ReplaceNullWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceNullWith

`func (o *AggregationGroupBy1) SetReplaceNullWith(v string)`

SetReplaceNullWith sets ReplaceNullWith field to given value.

### HasReplaceNullWith

`func (o *AggregationGroupBy1) HasReplaceNullWith() bool`

HasReplaceNullWith returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


