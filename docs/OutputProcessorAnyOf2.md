# OutputProcessorAnyOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the processor within the route. Used to preserve processor identity during updates. | 
**Description** | Pointer to **string** | Optional user-provided processor note. | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Key/value tags applied to the resource. Use them to organize resources and to satisfy tag policies. | [optional] 
**UpdatedAt** | Pointer to **string** | Timestamp when Tsuga last updated the processor. | [optional] 
**Type** | **string** | Create processor. It writes a new attribute from a format string, formula, or category rule. | 
**Params** | [**OutputProcessorAnyOf2Params**](OutputProcessorAnyOf2Params.md) |  | 

## Methods

### NewOutputProcessorAnyOf2

`func NewOutputProcessorAnyOf2(id string, type_ string, params OutputProcessorAnyOf2Params, ) *OutputProcessorAnyOf2`

NewOutputProcessorAnyOf2 instantiates a new OutputProcessorAnyOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputProcessorAnyOf2WithDefaults

`func NewOutputProcessorAnyOf2WithDefaults() *OutputProcessorAnyOf2`

NewOutputProcessorAnyOf2WithDefaults instantiates a new OutputProcessorAnyOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OutputProcessorAnyOf2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutputProcessorAnyOf2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OutputProcessorAnyOf2) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *OutputProcessorAnyOf2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OutputProcessorAnyOf2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OutputProcessorAnyOf2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OutputProcessorAnyOf2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *OutputProcessorAnyOf2) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OutputProcessorAnyOf2) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OutputProcessorAnyOf2) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OutputProcessorAnyOf2) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OutputProcessorAnyOf2) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OutputProcessorAnyOf2) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OutputProcessorAnyOf2) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OutputProcessorAnyOf2) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetType

`func (o *OutputProcessorAnyOf2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OutputProcessorAnyOf2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OutputProcessorAnyOf2) SetType(v string)`

SetType sets Type field to given value.


### GetParams

`func (o *OutputProcessorAnyOf2) GetParams() OutputProcessorAnyOf2Params`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *OutputProcessorAnyOf2) GetParamsOk() (*OutputProcessorAnyOf2Params, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *OutputProcessorAnyOf2) SetParams(v OutputProcessorAnyOf2Params)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


