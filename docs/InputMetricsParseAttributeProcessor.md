# InputMetricsParseAttributeProcessor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the processor within the route. Preserve it across updates for existing processors. | 
**Description** | Pointer to **string** | Optional processor note. | [optional] 
**Tags** | Pointer to [**[]InputTag**](InputTag.md) | Key/value tags to apply to the resource. Tag policies may require specific keys or values. | [optional] 
**UpdatedAt** | Pointer to **string** | Backend-maintained timestamp for this processor. Omit this field on create and update requests; Tsuga sets or preserves it based on processor changes. | [optional] 
**Type** | **string** | Parse processor. It extracts structured fields from a text attribute. | 
**Params** | [**InputMetricsParseAttributeProcessorParams**](InputMetricsParseAttributeProcessorParams.md) |  | 

## Methods

### NewInputMetricsParseAttributeProcessor

`func NewInputMetricsParseAttributeProcessor(id string, type_ string, params InputMetricsParseAttributeProcessorParams, ) *InputMetricsParseAttributeProcessor`

NewInputMetricsParseAttributeProcessor instantiates a new InputMetricsParseAttributeProcessor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputMetricsParseAttributeProcessorWithDefaults

`func NewInputMetricsParseAttributeProcessorWithDefaults() *InputMetricsParseAttributeProcessor`

NewInputMetricsParseAttributeProcessorWithDefaults instantiates a new InputMetricsParseAttributeProcessor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InputMetricsParseAttributeProcessor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InputMetricsParseAttributeProcessor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InputMetricsParseAttributeProcessor) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *InputMetricsParseAttributeProcessor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputMetricsParseAttributeProcessor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputMetricsParseAttributeProcessor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InputMetricsParseAttributeProcessor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *InputMetricsParseAttributeProcessor) GetTags() []InputTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InputMetricsParseAttributeProcessor) GetTagsOk() (*[]InputTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InputMetricsParseAttributeProcessor) SetTags(v []InputTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InputMetricsParseAttributeProcessor) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InputMetricsParseAttributeProcessor) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InputMetricsParseAttributeProcessor) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InputMetricsParseAttributeProcessor) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InputMetricsParseAttributeProcessor) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetType

`func (o *InputMetricsParseAttributeProcessor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputMetricsParseAttributeProcessor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputMetricsParseAttributeProcessor) SetType(v string)`

SetType sets Type field to given value.


### GetParams

`func (o *InputMetricsParseAttributeProcessor) GetParams() InputMetricsParseAttributeProcessorParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *InputMetricsParseAttributeProcessor) GetParamsOk() (*InputMetricsParseAttributeProcessorParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *InputMetricsParseAttributeProcessor) SetParams(v InputMetricsParseAttributeProcessorParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


