# InputMonitorConfigurationMetricGroupByFieldsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | **[]string** | Attribute names the monitor splits results by, evaluating each resulting group independently. Empty or omitted means the monitor evaluates a single ungrouped series. | 
**Limit** | **float32** | Configured limit for this group by. Warning! This setting is currently ignored. Monitor evaluation will instead apply a fixed limit of 100 groups per field. | 
**SortOrder** | Pointer to **string** | Sort direction applied to groups: ascending or descending. | [optional] 
**ReplaceNullWith** | Pointer to **string** | Value used to group documents that have no value for a grouped field. | [optional] 

## Methods

### NewInputMonitorConfigurationMetricGroupByFieldsInner

`func NewInputMonitorConfigurationMetricGroupByFieldsInner(fields []string, limit float32, ) *InputMonitorConfigurationMetricGroupByFieldsInner`

NewInputMonitorConfigurationMetricGroupByFieldsInner instantiates a new InputMonitorConfigurationMetricGroupByFieldsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputMonitorConfigurationMetricGroupByFieldsInnerWithDefaults

`func NewInputMonitorConfigurationMetricGroupByFieldsInnerWithDefaults() *InputMonitorConfigurationMetricGroupByFieldsInner`

NewInputMonitorConfigurationMetricGroupByFieldsInnerWithDefaults instantiates a new InputMonitorConfigurationMetricGroupByFieldsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *InputMonitorConfigurationMetricGroupByFieldsInner) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *InputMonitorConfigurationMetricGroupByFieldsInner) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *InputMonitorConfigurationMetricGroupByFieldsInner) SetFields(v []string)`

SetFields sets Fields field to given value.


### GetLimit

`func (o *InputMonitorConfigurationMetricGroupByFieldsInner) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *InputMonitorConfigurationMetricGroupByFieldsInner) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *InputMonitorConfigurationMetricGroupByFieldsInner) SetLimit(v float32)`

SetLimit sets Limit field to given value.


### GetSortOrder

`func (o *InputMonitorConfigurationMetricGroupByFieldsInner) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *InputMonitorConfigurationMetricGroupByFieldsInner) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *InputMonitorConfigurationMetricGroupByFieldsInner) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *InputMonitorConfigurationMetricGroupByFieldsInner) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetReplaceNullWith

`func (o *InputMonitorConfigurationMetricGroupByFieldsInner) GetReplaceNullWith() string`

GetReplaceNullWith returns the ReplaceNullWith field if non-nil, zero value otherwise.

### GetReplaceNullWithOk

`func (o *InputMonitorConfigurationMetricGroupByFieldsInner) GetReplaceNullWithOk() (*string, bool)`

GetReplaceNullWithOk returns a tuple with the ReplaceNullWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceNullWith

`func (o *InputMonitorConfigurationMetricGroupByFieldsInner) SetReplaceNullWith(v string)`

SetReplaceNullWith sets ReplaceNullWith field to given value.

### HasReplaceNullWith

`func (o *InputMonitorConfigurationMetricGroupByFieldsInner) HasReplaceNullWith() bool`

HasReplaceNullWith returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


