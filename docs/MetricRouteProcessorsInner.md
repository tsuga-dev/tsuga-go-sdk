# MetricRouteProcessorsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the processor within the route. Used to preserve processor identity during updates. | 
**Description** | Pointer to **string** | Optional user-provided processor note. | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Key/value tags applied to the resource. Use them to organize resources and to satisfy tag policies. | [optional] 
**UpdatedAt** | Pointer to **string** | Timestamp when Tsuga last updated the processor. | [optional] 
**Type** | **string** | Map processor. It moves, copies, or normalizes existing log attributes. | 
**Params** | [**MetricRouteProcessorsInnerOneOf2Params**](MetricRouteProcessorsInnerOneOf2Params.md) |  | 

## Methods

### NewMetricRouteProcessorsInner

`func NewMetricRouteProcessorsInner(id string, type_ string, params MetricRouteProcessorsInnerOneOf2Params, ) *MetricRouteProcessorsInner`

NewMetricRouteProcessorsInner instantiates a new MetricRouteProcessorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricRouteProcessorsInnerWithDefaults

`func NewMetricRouteProcessorsInnerWithDefaults() *MetricRouteProcessorsInner`

NewMetricRouteProcessorsInnerWithDefaults instantiates a new MetricRouteProcessorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetricRouteProcessorsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetricRouteProcessorsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetricRouteProcessorsInner) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *MetricRouteProcessorsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricRouteProcessorsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricRouteProcessorsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricRouteProcessorsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *MetricRouteProcessorsInner) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MetricRouteProcessorsInner) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MetricRouteProcessorsInner) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MetricRouteProcessorsInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MetricRouteProcessorsInner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MetricRouteProcessorsInner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MetricRouteProcessorsInner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MetricRouteProcessorsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetType

`func (o *MetricRouteProcessorsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricRouteProcessorsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricRouteProcessorsInner) SetType(v string)`

SetType sets Type field to given value.


### GetParams

`func (o *MetricRouteProcessorsInner) GetParams() MetricRouteProcessorsInnerOneOf2Params`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *MetricRouteProcessorsInner) GetParamsOk() (*MetricRouteProcessorsInnerOneOf2Params, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *MetricRouteProcessorsInner) SetParams(v MetricRouteProcessorsInnerOneOf2Params)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


