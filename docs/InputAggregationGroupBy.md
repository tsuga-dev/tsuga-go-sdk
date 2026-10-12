# InputAggregationGroupBy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | **[]string** | Telemetry attribute name to group aggregation results by. Currently limited to one field per grouping level. | 
**Limit** | **float32** | Maximum number of distinct values of &#x60;fields&#x60; to return for this grouping level, ranked by aggregate value. Groups beyond the limit are truncated from the response, not merged into an \&quot;other\&quot; group. | 
**SortOrder** | Pointer to **string** | Sort direction: ascending or descending. | [optional] 
**ReplaceNullWith** | Pointer to **string** | Value used to group documents that have no value for a grouped field. | [optional] 

## Methods

### NewInputAggregationGroupBy

`func NewInputAggregationGroupBy(fields []string, limit float32, ) *InputAggregationGroupBy`

NewInputAggregationGroupBy instantiates a new InputAggregationGroupBy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputAggregationGroupByWithDefaults

`func NewInputAggregationGroupByWithDefaults() *InputAggregationGroupBy`

NewInputAggregationGroupByWithDefaults instantiates a new InputAggregationGroupBy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *InputAggregationGroupBy) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *InputAggregationGroupBy) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *InputAggregationGroupBy) SetFields(v []string)`

SetFields sets Fields field to given value.


### GetLimit

`func (o *InputAggregationGroupBy) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *InputAggregationGroupBy) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *InputAggregationGroupBy) SetLimit(v float32)`

SetLimit sets Limit field to given value.


### GetSortOrder

`func (o *InputAggregationGroupBy) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *InputAggregationGroupBy) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *InputAggregationGroupBy) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *InputAggregationGroupBy) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetReplaceNullWith

`func (o *InputAggregationGroupBy) GetReplaceNullWith() string`

GetReplaceNullWith returns the ReplaceNullWith field if non-nil, zero value otherwise.

### GetReplaceNullWithOk

`func (o *InputAggregationGroupBy) GetReplaceNullWithOk() (*string, bool)`

GetReplaceNullWithOk returns a tuple with the ReplaceNullWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceNullWith

`func (o *InputAggregationGroupBy) SetReplaceNullWith(v string)`

SetReplaceNullWith sets ReplaceNullWith field to given value.

### HasReplaceNullWith

`func (o *InputAggregationGroupBy) HasReplaceNullWith() bool`

HasReplaceNullWith returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


