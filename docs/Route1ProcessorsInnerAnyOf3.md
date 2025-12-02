# Route1ProcessorsInnerAnyOf3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the processor | 
**Name** | **string** | Display name of the processor | 
**Description** | Pointer to **string** |  | [optional] 
**Example** | Pointer to [**Route1ProcessorsInnerAnyOfExample**](Route1ProcessorsInnerAnyOfExample.md) |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | List of key/value tags applied to the resource | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Params** | [**Route1ProcessorsInnerAnyOf3Params**](Route1ProcessorsInnerAnyOf3Params.md) |  | 
**Type** | **string** |  | 

## Methods

### NewRoute1ProcessorsInnerAnyOf3

`func NewRoute1ProcessorsInnerAnyOf3(id string, name string, params Route1ProcessorsInnerAnyOf3Params, type_ string, ) *Route1ProcessorsInnerAnyOf3`

NewRoute1ProcessorsInnerAnyOf3 instantiates a new Route1ProcessorsInnerAnyOf3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoute1ProcessorsInnerAnyOf3WithDefaults

`func NewRoute1ProcessorsInnerAnyOf3WithDefaults() *Route1ProcessorsInnerAnyOf3`

NewRoute1ProcessorsInnerAnyOf3WithDefaults instantiates a new Route1ProcessorsInnerAnyOf3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Route1ProcessorsInnerAnyOf3) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Route1ProcessorsInnerAnyOf3) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Route1ProcessorsInnerAnyOf3) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Route1ProcessorsInnerAnyOf3) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Route1ProcessorsInnerAnyOf3) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Route1ProcessorsInnerAnyOf3) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Route1ProcessorsInnerAnyOf3) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Route1ProcessorsInnerAnyOf3) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Route1ProcessorsInnerAnyOf3) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Route1ProcessorsInnerAnyOf3) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExample

`func (o *Route1ProcessorsInnerAnyOf3) GetExample() Route1ProcessorsInnerAnyOfExample`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *Route1ProcessorsInnerAnyOf3) GetExampleOk() (*Route1ProcessorsInnerAnyOfExample, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *Route1ProcessorsInnerAnyOf3) SetExample(v Route1ProcessorsInnerAnyOfExample)`

SetExample sets Example field to given value.

### HasExample

`func (o *Route1ProcessorsInnerAnyOf3) HasExample() bool`

HasExample returns a boolean if a field has been set.

### GetTags

`func (o *Route1ProcessorsInnerAnyOf3) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Route1ProcessorsInnerAnyOf3) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Route1ProcessorsInnerAnyOf3) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Route1ProcessorsInnerAnyOf3) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Route1ProcessorsInnerAnyOf3) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Route1ProcessorsInnerAnyOf3) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Route1ProcessorsInnerAnyOf3) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Route1ProcessorsInnerAnyOf3) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetParams

`func (o *Route1ProcessorsInnerAnyOf3) GetParams() Route1ProcessorsInnerAnyOf3Params`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *Route1ProcessorsInnerAnyOf3) GetParamsOk() (*Route1ProcessorsInnerAnyOf3Params, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *Route1ProcessorsInnerAnyOf3) SetParams(v Route1ProcessorsInnerAnyOf3Params)`

SetParams sets Params field to given value.


### GetType

`func (o *Route1ProcessorsInnerAnyOf3) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Route1ProcessorsInnerAnyOf3) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Route1ProcessorsInnerAnyOf3) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


