# Route1ProcessorsInnerAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the processor | 
**Name** | **string** | Display name of the processor | 
**Description** | Pointer to **string** |  | [optional] 
**Example** | Pointer to [**Route1ProcessorsInnerAnyOfExample**](Route1ProcessorsInnerAnyOfExample.md) |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | List of key/value tags applied to the resource | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Type** | **string** |  | 
**Params** | [**Route1ProcessorsInnerAnyOfParams**](Route1ProcessorsInnerAnyOfParams.md) |  | 

## Methods

### NewRoute1ProcessorsInnerAnyOf

`func NewRoute1ProcessorsInnerAnyOf(id string, name string, type_ string, params Route1ProcessorsInnerAnyOfParams, ) *Route1ProcessorsInnerAnyOf`

NewRoute1ProcessorsInnerAnyOf instantiates a new Route1ProcessorsInnerAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoute1ProcessorsInnerAnyOfWithDefaults

`func NewRoute1ProcessorsInnerAnyOfWithDefaults() *Route1ProcessorsInnerAnyOf`

NewRoute1ProcessorsInnerAnyOfWithDefaults instantiates a new Route1ProcessorsInnerAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Route1ProcessorsInnerAnyOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Route1ProcessorsInnerAnyOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Route1ProcessorsInnerAnyOf) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Route1ProcessorsInnerAnyOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Route1ProcessorsInnerAnyOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Route1ProcessorsInnerAnyOf) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Route1ProcessorsInnerAnyOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Route1ProcessorsInnerAnyOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Route1ProcessorsInnerAnyOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Route1ProcessorsInnerAnyOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExample

`func (o *Route1ProcessorsInnerAnyOf) GetExample() Route1ProcessorsInnerAnyOfExample`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *Route1ProcessorsInnerAnyOf) GetExampleOk() (*Route1ProcessorsInnerAnyOfExample, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *Route1ProcessorsInnerAnyOf) SetExample(v Route1ProcessorsInnerAnyOfExample)`

SetExample sets Example field to given value.

### HasExample

`func (o *Route1ProcessorsInnerAnyOf) HasExample() bool`

HasExample returns a boolean if a field has been set.

### GetTags

`func (o *Route1ProcessorsInnerAnyOf) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Route1ProcessorsInnerAnyOf) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Route1ProcessorsInnerAnyOf) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Route1ProcessorsInnerAnyOf) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Route1ProcessorsInnerAnyOf) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Route1ProcessorsInnerAnyOf) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Route1ProcessorsInnerAnyOf) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Route1ProcessorsInnerAnyOf) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetType

`func (o *Route1ProcessorsInnerAnyOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Route1ProcessorsInnerAnyOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Route1ProcessorsInnerAnyOf) SetType(v string)`

SetType sets Type field to given value.


### GetParams

`func (o *Route1ProcessorsInnerAnyOf) GetParams() Route1ProcessorsInnerAnyOfParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *Route1ProcessorsInnerAnyOf) GetParamsOk() (*Route1ProcessorsInnerAnyOfParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *Route1ProcessorsInnerAnyOf) SetParams(v Route1ProcessorsInnerAnyOfParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


