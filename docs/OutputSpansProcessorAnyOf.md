# OutputSpansProcessorAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the processor within the route. Used to preserve processor identity during updates. | 
**Description** | Pointer to **string** | Optional user-provided processor note. | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Key/value tags applied to the resource. Use them to organize resources and to satisfy tag policies. | [optional] 
**UpdatedAt** | Pointer to **string** | Timestamp when Tsuga last updated the processor. | [optional] 
**Params** | [**OutputSpansProcessorAnyOfParams**](OutputSpansProcessorAnyOfParams.md) |  | 
**Type** | **string** | Split processor. It sends each span through the first matching conditional branch. | 

## Methods

### NewOutputSpansProcessorAnyOf

`func NewOutputSpansProcessorAnyOf(id string, params OutputSpansProcessorAnyOfParams, type_ string, ) *OutputSpansProcessorAnyOf`

NewOutputSpansProcessorAnyOf instantiates a new OutputSpansProcessorAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputSpansProcessorAnyOfWithDefaults

`func NewOutputSpansProcessorAnyOfWithDefaults() *OutputSpansProcessorAnyOf`

NewOutputSpansProcessorAnyOfWithDefaults instantiates a new OutputSpansProcessorAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OutputSpansProcessorAnyOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutputSpansProcessorAnyOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OutputSpansProcessorAnyOf) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *OutputSpansProcessorAnyOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OutputSpansProcessorAnyOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OutputSpansProcessorAnyOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OutputSpansProcessorAnyOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *OutputSpansProcessorAnyOf) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OutputSpansProcessorAnyOf) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OutputSpansProcessorAnyOf) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OutputSpansProcessorAnyOf) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OutputSpansProcessorAnyOf) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OutputSpansProcessorAnyOf) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OutputSpansProcessorAnyOf) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OutputSpansProcessorAnyOf) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetParams

`func (o *OutputSpansProcessorAnyOf) GetParams() OutputSpansProcessorAnyOfParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *OutputSpansProcessorAnyOf) GetParamsOk() (*OutputSpansProcessorAnyOfParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *OutputSpansProcessorAnyOf) SetParams(v OutputSpansProcessorAnyOfParams)`

SetParams sets Params field to given value.


### GetType

`func (o *OutputSpansProcessorAnyOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OutputSpansProcessorAnyOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OutputSpansProcessorAnyOf) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


