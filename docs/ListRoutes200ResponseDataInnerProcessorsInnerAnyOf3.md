# ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the processor | 
**Name** | **string** | Display name of the processor | 
**Description** | Pointer to **string** |  | [optional] 
**Example** | Pointer to [**ListRoutes200ResponseDataInnerProcessorsInnerAnyOfExample**](ListRoutes200ResponseDataInnerProcessorsInnerAnyOfExample.md) |  | [optional] 
**Tags** | Pointer to [**[]ListDashboards200ResponseDataInnerTagsInner**](ListDashboards200ResponseDataInnerTagsInner.md) | List of key/value tags applied to the resource | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Params** | [**ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3Params**](ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3Params.md) |  | 
**Type** | **string** |  | 

## Methods

### NewListRoutes200ResponseDataInnerProcessorsInnerAnyOf3

`func NewListRoutes200ResponseDataInnerProcessorsInnerAnyOf3(id string, name string, params ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3Params, type_ string, ) *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3`

NewListRoutes200ResponseDataInnerProcessorsInnerAnyOf3 instantiates a new ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRoutes200ResponseDataInnerProcessorsInnerAnyOf3WithDefaults

`func NewListRoutes200ResponseDataInnerProcessorsInnerAnyOf3WithDefaults() *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3`

NewListRoutes200ResponseDataInnerProcessorsInnerAnyOf3WithDefaults instantiates a new ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExample

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3) GetExample() ListRoutes200ResponseDataInnerProcessorsInnerAnyOfExample`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3) GetExampleOk() (*ListRoutes200ResponseDataInnerProcessorsInnerAnyOfExample, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3) SetExample(v ListRoutes200ResponseDataInnerProcessorsInnerAnyOfExample)`

SetExample sets Example field to given value.

### HasExample

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3) HasExample() bool`

HasExample returns a boolean if a field has been set.

### GetTags

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3) GetTags() []ListDashboards200ResponseDataInnerTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3) GetTagsOk() (*[]ListDashboards200ResponseDataInnerTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3) SetTags(v []ListDashboards200ResponseDataInnerTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetParams

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3) GetParams() ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3Params`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3) GetParamsOk() (*ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3Params, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3) SetParams(v ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3Params)`

SetParams sets Params field to given value.


### GetType

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListRoutes200ResponseDataInnerProcessorsInnerAnyOf3) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


