# OutputRumProcessorAnyOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the processor within the route. Used to preserve processor identity during updates. | 
**Description** | Pointer to **string** | Optional user-provided processor note. | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Key/value tags applied to the resource. Use them to organize resources and to satisfy tag policies. | [optional] 
**UpdatedAt** | Pointer to **string** | Timestamp when Tsuga last updated the processor. | [optional] 
**Params** | [**OutputRumProcessorAnyOf1Params**](OutputRumProcessorAnyOf1Params.md) |  | 
**Type** | **string** | Split processor. It sends each RUM event through the first matching conditional branch. | 

## Methods

### NewOutputRumProcessorAnyOf1

`func NewOutputRumProcessorAnyOf1(id string, params OutputRumProcessorAnyOf1Params, type_ string, ) *OutputRumProcessorAnyOf1`

NewOutputRumProcessorAnyOf1 instantiates a new OutputRumProcessorAnyOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputRumProcessorAnyOf1WithDefaults

`func NewOutputRumProcessorAnyOf1WithDefaults() *OutputRumProcessorAnyOf1`

NewOutputRumProcessorAnyOf1WithDefaults instantiates a new OutputRumProcessorAnyOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OutputRumProcessorAnyOf1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutputRumProcessorAnyOf1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OutputRumProcessorAnyOf1) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *OutputRumProcessorAnyOf1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OutputRumProcessorAnyOf1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OutputRumProcessorAnyOf1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OutputRumProcessorAnyOf1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *OutputRumProcessorAnyOf1) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OutputRumProcessorAnyOf1) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OutputRumProcessorAnyOf1) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OutputRumProcessorAnyOf1) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OutputRumProcessorAnyOf1) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OutputRumProcessorAnyOf1) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OutputRumProcessorAnyOf1) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OutputRumProcessorAnyOf1) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetParams

`func (o *OutputRumProcessorAnyOf1) GetParams() OutputRumProcessorAnyOf1Params`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *OutputRumProcessorAnyOf1) GetParamsOk() (*OutputRumProcessorAnyOf1Params, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *OutputRumProcessorAnyOf1) SetParams(v OutputRumProcessorAnyOf1Params)`

SetParams sets Params field to given value.


### GetType

`func (o *OutputRumProcessorAnyOf1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OutputRumProcessorAnyOf1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OutputRumProcessorAnyOf1) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


