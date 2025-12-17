# CreateRouteRequestProcessorsInnerAnyOf1

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
**Params** | [**CreateRouteRequestProcessorsInnerAnyOf1Params**](CreateRouteRequestProcessorsInnerAnyOf1Params.md) |  | 

## Methods

### NewCreateRouteRequestProcessorsInnerAnyOf1

`func NewCreateRouteRequestProcessorsInnerAnyOf1(id string, name string, type_ string, params CreateRouteRequestProcessorsInnerAnyOf1Params, ) *CreateRouteRequestProcessorsInnerAnyOf1`

NewCreateRouteRequestProcessorsInnerAnyOf1 instantiates a new CreateRouteRequestProcessorsInnerAnyOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRouteRequestProcessorsInnerAnyOf1WithDefaults

`func NewCreateRouteRequestProcessorsInnerAnyOf1WithDefaults() *CreateRouteRequestProcessorsInnerAnyOf1`

NewCreateRouteRequestProcessorsInnerAnyOf1WithDefaults instantiates a new CreateRouteRequestProcessorsInnerAnyOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateRouteRequestProcessorsInnerAnyOf1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateRouteRequestProcessorsInnerAnyOf1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateRouteRequestProcessorsInnerAnyOf1) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CreateRouteRequestProcessorsInnerAnyOf1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRouteRequestProcessorsInnerAnyOf1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRouteRequestProcessorsInnerAnyOf1) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateRouteRequestProcessorsInnerAnyOf1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateRouteRequestProcessorsInnerAnyOf1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateRouteRequestProcessorsInnerAnyOf1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateRouteRequestProcessorsInnerAnyOf1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExample

`func (o *CreateRouteRequestProcessorsInnerAnyOf1) GetExample() CreateRouteRequestProcessorsInnerAnyOfExample`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *CreateRouteRequestProcessorsInnerAnyOf1) GetExampleOk() (*CreateRouteRequestProcessorsInnerAnyOfExample, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *CreateRouteRequestProcessorsInnerAnyOf1) SetExample(v CreateRouteRequestProcessorsInnerAnyOfExample)`

SetExample sets Example field to given value.

### HasExample

`func (o *CreateRouteRequestProcessorsInnerAnyOf1) HasExample() bool`

HasExample returns a boolean if a field has been set.

### GetTags

`func (o *CreateRouteRequestProcessorsInnerAnyOf1) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateRouteRequestProcessorsInnerAnyOf1) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateRouteRequestProcessorsInnerAnyOf1) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateRouteRequestProcessorsInnerAnyOf1) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CreateRouteRequestProcessorsInnerAnyOf1) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateRouteRequestProcessorsInnerAnyOf1) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateRouteRequestProcessorsInnerAnyOf1) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CreateRouteRequestProcessorsInnerAnyOf1) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetType

`func (o *CreateRouteRequestProcessorsInnerAnyOf1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateRouteRequestProcessorsInnerAnyOf1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateRouteRequestProcessorsInnerAnyOf1) SetType(v string)`

SetType sets Type field to given value.


### GetParams

`func (o *CreateRouteRequestProcessorsInnerAnyOf1) GetParams() CreateRouteRequestProcessorsInnerAnyOf1Params`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CreateRouteRequestProcessorsInnerAnyOf1) GetParamsOk() (*CreateRouteRequestProcessorsInnerAnyOf1Params, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CreateRouteRequestProcessorsInnerAnyOf1) SetParams(v CreateRouteRequestProcessorsInnerAnyOf1Params)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


