# InputProcessorAnyOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the processor within the route. Preserve it across updates for existing processors. | 
**Description** | Pointer to **string** | Optional processor note. | [optional] 
**Tags** | Pointer to [**[]InputTag**](InputTag.md) | Key/value tags to apply to the resource. Tag policies may require specific keys or values. | [optional] 
**UpdatedAt** | Pointer to **string** | Backend-maintained timestamp for this processor. Omit this field on create and update requests; Tsuga sets or preserves it based on processor changes. | [optional] 
**Type** | **string** | Create processor. It writes a new attribute from a format string, formula, or category rule. | 
**Params** | [**InputProcessorAnyOf2Params**](InputProcessorAnyOf2Params.md) |  | 

## Methods

### NewInputProcessorAnyOf2

`func NewInputProcessorAnyOf2(id string, type_ string, params InputProcessorAnyOf2Params, ) *InputProcessorAnyOf2`

NewInputProcessorAnyOf2 instantiates a new InputProcessorAnyOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputProcessorAnyOf2WithDefaults

`func NewInputProcessorAnyOf2WithDefaults() *InputProcessorAnyOf2`

NewInputProcessorAnyOf2WithDefaults instantiates a new InputProcessorAnyOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InputProcessorAnyOf2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InputProcessorAnyOf2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InputProcessorAnyOf2) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *InputProcessorAnyOf2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputProcessorAnyOf2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputProcessorAnyOf2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InputProcessorAnyOf2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *InputProcessorAnyOf2) GetTags() []InputTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InputProcessorAnyOf2) GetTagsOk() (*[]InputTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InputProcessorAnyOf2) SetTags(v []InputTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InputProcessorAnyOf2) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InputProcessorAnyOf2) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InputProcessorAnyOf2) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InputProcessorAnyOf2) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InputProcessorAnyOf2) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetType

`func (o *InputProcessorAnyOf2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputProcessorAnyOf2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputProcessorAnyOf2) SetType(v string)`

SetType sets Type field to given value.


### GetParams

`func (o *InputProcessorAnyOf2) GetParams() InputProcessorAnyOf2Params`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *InputProcessorAnyOf2) GetParamsOk() (*InputProcessorAnyOf2Params, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *InputProcessorAnyOf2) SetParams(v InputProcessorAnyOf2Params)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


