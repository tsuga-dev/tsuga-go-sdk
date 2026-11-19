# InputRumProcessor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the processor within the route. Preserve it across updates for existing processors. | 
**Description** | Pointer to **string** | Optional processor note. | [optional] 
**Tags** | Pointer to [**[]InputTag**](InputTag.md) | Key/value tags to apply to the resource. Tag policies may require specific keys or values. | [optional] 
**UpdatedAt** | Pointer to **string** | Backend-maintained timestamp for this processor. Omit this field on create and update requests; Tsuga sets or preserves it based on processor changes. | [optional] 
**Type** | **string** | Map processor. It moves, copies, or normalizes existing event attributes. | 
**Params** | [**InputRumProcessorAnyOfParams**](InputRumProcessorAnyOfParams.md) |  | 

## Methods

### NewInputRumProcessor

`func NewInputRumProcessor(id string, type_ string, params InputRumProcessorAnyOfParams, ) *InputRumProcessor`

NewInputRumProcessor instantiates a new InputRumProcessor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputRumProcessorWithDefaults

`func NewInputRumProcessorWithDefaults() *InputRumProcessor`

NewInputRumProcessorWithDefaults instantiates a new InputRumProcessor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InputRumProcessor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InputRumProcessor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InputRumProcessor) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *InputRumProcessor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputRumProcessor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputRumProcessor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InputRumProcessor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *InputRumProcessor) GetTags() []InputTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InputRumProcessor) GetTagsOk() (*[]InputTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InputRumProcessor) SetTags(v []InputTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InputRumProcessor) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InputRumProcessor) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InputRumProcessor) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InputRumProcessor) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InputRumProcessor) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetType

`func (o *InputRumProcessor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputRumProcessor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputRumProcessor) SetType(v string)`

SetType sets Type field to given value.


### GetParams

`func (o *InputRumProcessor) GetParams() InputRumProcessorAnyOfParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *InputRumProcessor) GetParamsOk() (*InputRumProcessorAnyOfParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *InputRumProcessor) SetParams(v InputRumProcessorAnyOfParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


