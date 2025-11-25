# ListRoutes200ResponseDataInnerProcessorsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the processor | 
**Name** | **string** | Display name of the processor | 
**Description** | Pointer to **string** |  | [optional] 
**Example** | Pointer to [**ListRoutes200ResponseDataInnerProcessorsInnerAnyOfExample**](ListRoutes200ResponseDataInnerProcessorsInnerAnyOfExample.md) |  | [optional] 
**Tags** | Pointer to [**[]ListDashboards200ResponseDataInnerTagsInner**](ListDashboards200ResponseDataInnerTagsInner.md) | List of key/value tags applied to the resource | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Type** | **string** |  | 
**Params** | [**ListRoutes200ResponseDataInnerProcessorsInnerAnyOf4Params**](ListRoutes200ResponseDataInnerProcessorsInnerAnyOf4Params.md) |  | 

## Methods

### NewListRoutes200ResponseDataInnerProcessorsInner

`func NewListRoutes200ResponseDataInnerProcessorsInner(id string, name string, type_ string, params ListRoutes200ResponseDataInnerProcessorsInnerAnyOf4Params, ) *ListRoutes200ResponseDataInnerProcessorsInner`

NewListRoutes200ResponseDataInnerProcessorsInner instantiates a new ListRoutes200ResponseDataInnerProcessorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRoutes200ResponseDataInnerProcessorsInnerWithDefaults

`func NewListRoutes200ResponseDataInnerProcessorsInnerWithDefaults() *ListRoutes200ResponseDataInnerProcessorsInner`

NewListRoutes200ResponseDataInnerProcessorsInnerWithDefaults instantiates a new ListRoutes200ResponseDataInnerProcessorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListRoutes200ResponseDataInnerProcessorsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListRoutes200ResponseDataInnerProcessorsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListRoutes200ResponseDataInnerProcessorsInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ListRoutes200ResponseDataInnerProcessorsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListRoutes200ResponseDataInnerProcessorsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListRoutes200ResponseDataInnerProcessorsInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ListRoutes200ResponseDataInnerProcessorsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListRoutes200ResponseDataInnerProcessorsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListRoutes200ResponseDataInnerProcessorsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListRoutes200ResponseDataInnerProcessorsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExample

`func (o *ListRoutes200ResponseDataInnerProcessorsInner) GetExample() ListRoutes200ResponseDataInnerProcessorsInnerAnyOfExample`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *ListRoutes200ResponseDataInnerProcessorsInner) GetExampleOk() (*ListRoutes200ResponseDataInnerProcessorsInnerAnyOfExample, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *ListRoutes200ResponseDataInnerProcessorsInner) SetExample(v ListRoutes200ResponseDataInnerProcessorsInnerAnyOfExample)`

SetExample sets Example field to given value.

### HasExample

`func (o *ListRoutes200ResponseDataInnerProcessorsInner) HasExample() bool`

HasExample returns a boolean if a field has been set.

### GetTags

`func (o *ListRoutes200ResponseDataInnerProcessorsInner) GetTags() []ListDashboards200ResponseDataInnerTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ListRoutes200ResponseDataInnerProcessorsInner) GetTagsOk() (*[]ListDashboards200ResponseDataInnerTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ListRoutes200ResponseDataInnerProcessorsInner) SetTags(v []ListDashboards200ResponseDataInnerTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ListRoutes200ResponseDataInnerProcessorsInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ListRoutes200ResponseDataInnerProcessorsInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListRoutes200ResponseDataInnerProcessorsInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListRoutes200ResponseDataInnerProcessorsInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ListRoutes200ResponseDataInnerProcessorsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetType

`func (o *ListRoutes200ResponseDataInnerProcessorsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListRoutes200ResponseDataInnerProcessorsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListRoutes200ResponseDataInnerProcessorsInner) SetType(v string)`

SetType sets Type field to given value.


### GetParams

`func (o *ListRoutes200ResponseDataInnerProcessorsInner) GetParams() ListRoutes200ResponseDataInnerProcessorsInnerAnyOf4Params`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ListRoutes200ResponseDataInnerProcessorsInner) GetParamsOk() (*ListRoutes200ResponseDataInnerProcessorsInnerAnyOf4Params, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ListRoutes200ResponseDataInnerProcessorsInner) SetParams(v ListRoutes200ResponseDataInnerProcessorsInnerAnyOf4Params)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


