# MetricRouteProcessorsInnerOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the processor within the route. Used to preserve processor identity during updates. | 
**Description** | Pointer to **string** | Optional user-provided processor note. | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Key/value tags applied to the resource. Use them to organize resources and to satisfy tag policies. | [optional] 
**UpdatedAt** | Pointer to **string** | Timestamp when Tsuga last updated the processor. | [optional] 
**Type** | **string** | Parse processor. It extracts structured fields from a text attribute. | 
**Params** | [**MetricRouteProcessorsInnerOneOf1Params**](MetricRouteProcessorsInnerOneOf1Params.md) |  | 

## Methods

### NewMetricRouteProcessorsInnerOneOf1

`func NewMetricRouteProcessorsInnerOneOf1(id string, type_ string, params MetricRouteProcessorsInnerOneOf1Params, ) *MetricRouteProcessorsInnerOneOf1`

NewMetricRouteProcessorsInnerOneOf1 instantiates a new MetricRouteProcessorsInnerOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricRouteProcessorsInnerOneOf1WithDefaults

`func NewMetricRouteProcessorsInnerOneOf1WithDefaults() *MetricRouteProcessorsInnerOneOf1`

NewMetricRouteProcessorsInnerOneOf1WithDefaults instantiates a new MetricRouteProcessorsInnerOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetricRouteProcessorsInnerOneOf1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetricRouteProcessorsInnerOneOf1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetricRouteProcessorsInnerOneOf1) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *MetricRouteProcessorsInnerOneOf1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricRouteProcessorsInnerOneOf1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricRouteProcessorsInnerOneOf1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricRouteProcessorsInnerOneOf1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *MetricRouteProcessorsInnerOneOf1) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MetricRouteProcessorsInnerOneOf1) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MetricRouteProcessorsInnerOneOf1) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MetricRouteProcessorsInnerOneOf1) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MetricRouteProcessorsInnerOneOf1) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MetricRouteProcessorsInnerOneOf1) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MetricRouteProcessorsInnerOneOf1) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MetricRouteProcessorsInnerOneOf1) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetType

`func (o *MetricRouteProcessorsInnerOneOf1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricRouteProcessorsInnerOneOf1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricRouteProcessorsInnerOneOf1) SetType(v string)`

SetType sets Type field to given value.


### GetParams

`func (o *MetricRouteProcessorsInnerOneOf1) GetParams() MetricRouteProcessorsInnerOneOf1Params`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *MetricRouteProcessorsInnerOneOf1) GetParamsOk() (*MetricRouteProcessorsInnerOneOf1Params, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *MetricRouteProcessorsInnerOneOf1) SetParams(v MetricRouteProcessorsInnerOneOf1Params)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


