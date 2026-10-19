# InputProcessorAnyOf3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the processor within the route. Preserve it across updates for existing processors. | 
**Description** | Pointer to **string** | Optional processor note. | [optional] 
**Tags** | Pointer to [**[]InputTag**](InputTag.md) | Key/value tags to apply to the resource. Tag policies may require specific keys or values. | [optional] 
**UpdatedAt** | Pointer to **string** | Backend-maintained timestamp for this processor. Omit this field on create and update requests; Tsuga sets or preserves it based on processor changes. | [optional] 
**Params** | [**InputProcessorAnyOf3Params**](InputProcessorAnyOf3Params.md) |  | 
**Type** | **string** | Split processor. It sends each log through the first matching conditional branch. | 

## Methods

### NewInputProcessorAnyOf3

`func NewInputProcessorAnyOf3(id string, params InputProcessorAnyOf3Params, type_ string, ) *InputProcessorAnyOf3`

NewInputProcessorAnyOf3 instantiates a new InputProcessorAnyOf3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputProcessorAnyOf3WithDefaults

`func NewInputProcessorAnyOf3WithDefaults() *InputProcessorAnyOf3`

NewInputProcessorAnyOf3WithDefaults instantiates a new InputProcessorAnyOf3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InputProcessorAnyOf3) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InputProcessorAnyOf3) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InputProcessorAnyOf3) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *InputProcessorAnyOf3) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputProcessorAnyOf3) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputProcessorAnyOf3) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InputProcessorAnyOf3) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *InputProcessorAnyOf3) GetTags() []InputTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InputProcessorAnyOf3) GetTagsOk() (*[]InputTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InputProcessorAnyOf3) SetTags(v []InputTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InputProcessorAnyOf3) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InputProcessorAnyOf3) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InputProcessorAnyOf3) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InputProcessorAnyOf3) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InputProcessorAnyOf3) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetParams

`func (o *InputProcessorAnyOf3) GetParams() InputProcessorAnyOf3Params`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *InputProcessorAnyOf3) GetParamsOk() (*InputProcessorAnyOf3Params, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *InputProcessorAnyOf3) SetParams(v InputProcessorAnyOf3Params)`

SetParams sets Params field to given value.


### GetType

`func (o *InputProcessorAnyOf3) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputProcessorAnyOf3) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputProcessorAnyOf3) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


