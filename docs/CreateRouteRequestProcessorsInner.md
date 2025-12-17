# CreateRouteRequestProcessorsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the processor | 
**Name** | **string** | Display name of the processor | 
**Description** | Pointer to **string** |  | [optional] 
**Example** | Pointer to [**CreateRouteRequestProcessorsInnerAnyOfExample**](CreateRouteRequestProcessorsInnerAnyOfExample.md) |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | List of key/value tags applied to the resource | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Type** | **string** |  | 
**Params** | [**CreateRouteRequestProcessorsInnerAnyOf3Params**](CreateRouteRequestProcessorsInnerAnyOf3Params.md) |  | 

## Methods

### NewCreateRouteRequestProcessorsInner

`func NewCreateRouteRequestProcessorsInner(id string, name string, type_ string, params CreateRouteRequestProcessorsInnerAnyOf3Params, ) *CreateRouteRequestProcessorsInner`

NewCreateRouteRequestProcessorsInner instantiates a new CreateRouteRequestProcessorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRouteRequestProcessorsInnerWithDefaults

`func NewCreateRouteRequestProcessorsInnerWithDefaults() *CreateRouteRequestProcessorsInner`

NewCreateRouteRequestProcessorsInnerWithDefaults instantiates a new CreateRouteRequestProcessorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateRouteRequestProcessorsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateRouteRequestProcessorsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateRouteRequestProcessorsInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CreateRouteRequestProcessorsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRouteRequestProcessorsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRouteRequestProcessorsInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateRouteRequestProcessorsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateRouteRequestProcessorsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateRouteRequestProcessorsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateRouteRequestProcessorsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExample

`func (o *CreateRouteRequestProcessorsInner) GetExample() CreateRouteRequestProcessorsInnerAnyOfExample`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *CreateRouteRequestProcessorsInner) GetExampleOk() (*CreateRouteRequestProcessorsInnerAnyOfExample, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *CreateRouteRequestProcessorsInner) SetExample(v CreateRouteRequestProcessorsInnerAnyOfExample)`

SetExample sets Example field to given value.

### HasExample

`func (o *CreateRouteRequestProcessorsInner) HasExample() bool`

HasExample returns a boolean if a field has been set.

### GetTags

`func (o *CreateRouteRequestProcessorsInner) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateRouteRequestProcessorsInner) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateRouteRequestProcessorsInner) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateRouteRequestProcessorsInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CreateRouteRequestProcessorsInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateRouteRequestProcessorsInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateRouteRequestProcessorsInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CreateRouteRequestProcessorsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetType

`func (o *CreateRouteRequestProcessorsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateRouteRequestProcessorsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateRouteRequestProcessorsInner) SetType(v string)`

SetType sets Type field to given value.


### GetParams

`func (o *CreateRouteRequestProcessorsInner) GetParams() CreateRouteRequestProcessorsInnerAnyOf3Params`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CreateRouteRequestProcessorsInner) GetParamsOk() (*CreateRouteRequestProcessorsInnerAnyOf3Params, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CreateRouteRequestProcessorsInner) SetParams(v CreateRouteRequestProcessorsInnerAnyOf3Params)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


