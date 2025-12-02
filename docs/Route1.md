# Route1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human readable name shown for the route | 
**Description** | Pointer to **string** |  | [optional] 
**IsEnabled** | **bool** |  | 
**Query** | **string** | Query that selects which logs should enter the route | 
**Tags** | Pointer to [**[]Tag**](Tag.md) | List of key/value tags applied to the resource | [optional] 
**Owner** | **string** | Team ID owning and managing the route | 
**Processors** | [**[]Route1ProcessorsInner**](Route1ProcessorsInner.md) | Ordered processors applied to logs that match the route | 

## Methods

### NewRoute1

`func NewRoute1(name string, isEnabled bool, query string, owner string, processors []Route1ProcessorsInner, ) *Route1`

NewRoute1 instantiates a new Route1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoute1WithDefaults

`func NewRoute1WithDefaults() *Route1`

NewRoute1WithDefaults instantiates a new Route1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Route1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Route1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Route1) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Route1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Route1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Route1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Route1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsEnabled

`func (o *Route1) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *Route1) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *Route1) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetQuery

`func (o *Route1) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Route1) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Route1) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetTags

`func (o *Route1) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Route1) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Route1) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Route1) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOwner

`func (o *Route1) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Route1) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Route1) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetProcessors

`func (o *Route1) GetProcessors() []Route1ProcessorsInner`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *Route1) GetProcessorsOk() (*[]Route1ProcessorsInner, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *Route1) SetProcessors(v []Route1ProcessorsInner)`

SetProcessors sets Processors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


