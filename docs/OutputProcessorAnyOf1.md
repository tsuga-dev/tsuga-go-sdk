# OutputProcessorAnyOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the processor within the route. Used to preserve processor identity during updates. | 
**Description** | Pointer to **string** | Optional user-provided processor note. | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Key/value tags applied to the resource. Use them to organize resources and to satisfy tag policies. | [optional] 
**UpdatedAt** | Pointer to **string** | Timestamp when Tsuga last updated the processor. | [optional] 
**Type** | **string** | Parse processor. It extracts structured fields from a text attribute. | 
**Params** | [**OutputProcessorAnyOf1Params**](OutputProcessorAnyOf1Params.md) |  | 

## Methods

### NewOutputProcessorAnyOf1

`func NewOutputProcessorAnyOf1(id string, type_ string, params OutputProcessorAnyOf1Params, ) *OutputProcessorAnyOf1`

NewOutputProcessorAnyOf1 instantiates a new OutputProcessorAnyOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputProcessorAnyOf1WithDefaults

`func NewOutputProcessorAnyOf1WithDefaults() *OutputProcessorAnyOf1`

NewOutputProcessorAnyOf1WithDefaults instantiates a new OutputProcessorAnyOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OutputProcessorAnyOf1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutputProcessorAnyOf1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OutputProcessorAnyOf1) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *OutputProcessorAnyOf1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OutputProcessorAnyOf1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OutputProcessorAnyOf1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OutputProcessorAnyOf1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *OutputProcessorAnyOf1) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OutputProcessorAnyOf1) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OutputProcessorAnyOf1) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OutputProcessorAnyOf1) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OutputProcessorAnyOf1) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OutputProcessorAnyOf1) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OutputProcessorAnyOf1) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OutputProcessorAnyOf1) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetType

`func (o *OutputProcessorAnyOf1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OutputProcessorAnyOf1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OutputProcessorAnyOf1) SetType(v string)`

SetType sets Type field to given value.


### GetParams

`func (o *OutputProcessorAnyOf1) GetParams() OutputProcessorAnyOf1Params`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *OutputProcessorAnyOf1) GetParamsOk() (*OutputProcessorAnyOf1Params, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *OutputProcessorAnyOf1) SetParams(v OutputProcessorAnyOf1Params)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


