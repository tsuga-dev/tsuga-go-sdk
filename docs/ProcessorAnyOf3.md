# ProcessorAnyOf3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the processor | 
**Description** | Pointer to **string** |  | [optional] 
**Example** | Pointer to [**ProcessorAnyOfExample**](ProcessorAnyOfExample.md) |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | List of key/value tags applied to the resource | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Params** | [**ProcessorAnyOf3Params**](ProcessorAnyOf3Params.md) |  | 
**Type** | **string** |  | 

## Methods

### NewProcessorAnyOf3

`func NewProcessorAnyOf3(id string, params ProcessorAnyOf3Params, type_ string, ) *ProcessorAnyOf3`

NewProcessorAnyOf3 instantiates a new ProcessorAnyOf3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorAnyOf3WithDefaults

`func NewProcessorAnyOf3WithDefaults() *ProcessorAnyOf3`

NewProcessorAnyOf3WithDefaults instantiates a new ProcessorAnyOf3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProcessorAnyOf3) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProcessorAnyOf3) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProcessorAnyOf3) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *ProcessorAnyOf3) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProcessorAnyOf3) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProcessorAnyOf3) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProcessorAnyOf3) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExample

`func (o *ProcessorAnyOf3) GetExample() ProcessorAnyOfExample`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *ProcessorAnyOf3) GetExampleOk() (*ProcessorAnyOfExample, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *ProcessorAnyOf3) SetExample(v ProcessorAnyOfExample)`

SetExample sets Example field to given value.

### HasExample

`func (o *ProcessorAnyOf3) HasExample() bool`

HasExample returns a boolean if a field has been set.

### GetTags

`func (o *ProcessorAnyOf3) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProcessorAnyOf3) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProcessorAnyOf3) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProcessorAnyOf3) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ProcessorAnyOf3) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProcessorAnyOf3) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProcessorAnyOf3) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProcessorAnyOf3) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetParams

`func (o *ProcessorAnyOf3) GetParams() ProcessorAnyOf3Params`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ProcessorAnyOf3) GetParamsOk() (*ProcessorAnyOf3Params, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ProcessorAnyOf3) SetParams(v ProcessorAnyOf3Params)`

SetParams sets Params field to given value.


### GetType

`func (o *ProcessorAnyOf3) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProcessorAnyOf3) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProcessorAnyOf3) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


