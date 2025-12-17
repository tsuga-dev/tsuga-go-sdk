# CreateRouteRequestProcessorsInnerAnyOf3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the processor | 
**Name** | **string** | Display name of the processor | 
**Description** | Pointer to **string** |  | [optional] 
**Example** | Pointer to [**CreateRouteRequestProcessorsInnerAnyOfExample**](CreateRouteRequestProcessorsInnerAnyOfExample.md) |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | List of key/value tags applied to the resource | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Params** | [**CreateRouteRequestProcessorsInnerAnyOf3Params**](CreateRouteRequestProcessorsInnerAnyOf3Params.md) |  | 
**Type** | **string** |  | 

## Methods

### NewCreateRouteRequestProcessorsInnerAnyOf3

`func NewCreateRouteRequestProcessorsInnerAnyOf3(id string, name string, params CreateRouteRequestProcessorsInnerAnyOf3Params, type_ string, ) *CreateRouteRequestProcessorsInnerAnyOf3`

NewCreateRouteRequestProcessorsInnerAnyOf3 instantiates a new CreateRouteRequestProcessorsInnerAnyOf3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRouteRequestProcessorsInnerAnyOf3WithDefaults

`func NewCreateRouteRequestProcessorsInnerAnyOf3WithDefaults() *CreateRouteRequestProcessorsInnerAnyOf3`

NewCreateRouteRequestProcessorsInnerAnyOf3WithDefaults instantiates a new CreateRouteRequestProcessorsInnerAnyOf3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateRouteRequestProcessorsInnerAnyOf3) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateRouteRequestProcessorsInnerAnyOf3) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateRouteRequestProcessorsInnerAnyOf3) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CreateRouteRequestProcessorsInnerAnyOf3) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRouteRequestProcessorsInnerAnyOf3) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRouteRequestProcessorsInnerAnyOf3) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateRouteRequestProcessorsInnerAnyOf3) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateRouteRequestProcessorsInnerAnyOf3) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateRouteRequestProcessorsInnerAnyOf3) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateRouteRequestProcessorsInnerAnyOf3) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExample

`func (o *CreateRouteRequestProcessorsInnerAnyOf3) GetExample() CreateRouteRequestProcessorsInnerAnyOfExample`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *CreateRouteRequestProcessorsInnerAnyOf3) GetExampleOk() (*CreateRouteRequestProcessorsInnerAnyOfExample, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *CreateRouteRequestProcessorsInnerAnyOf3) SetExample(v CreateRouteRequestProcessorsInnerAnyOfExample)`

SetExample sets Example field to given value.

### HasExample

`func (o *CreateRouteRequestProcessorsInnerAnyOf3) HasExample() bool`

HasExample returns a boolean if a field has been set.

### GetTags

`func (o *CreateRouteRequestProcessorsInnerAnyOf3) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateRouteRequestProcessorsInnerAnyOf3) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateRouteRequestProcessorsInnerAnyOf3) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateRouteRequestProcessorsInnerAnyOf3) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CreateRouteRequestProcessorsInnerAnyOf3) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateRouteRequestProcessorsInnerAnyOf3) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateRouteRequestProcessorsInnerAnyOf3) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CreateRouteRequestProcessorsInnerAnyOf3) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetParams

`func (o *CreateRouteRequestProcessorsInnerAnyOf3) GetParams() CreateRouteRequestProcessorsInnerAnyOf3Params`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CreateRouteRequestProcessorsInnerAnyOf3) GetParamsOk() (*CreateRouteRequestProcessorsInnerAnyOf3Params, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CreateRouteRequestProcessorsInnerAnyOf3) SetParams(v CreateRouteRequestProcessorsInnerAnyOf3Params)`

SetParams sets Params field to given value.


### GetType

`func (o *CreateRouteRequestProcessorsInnerAnyOf3) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateRouteRequestProcessorsInnerAnyOf3) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateRouteRequestProcessorsInnerAnyOf3) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


