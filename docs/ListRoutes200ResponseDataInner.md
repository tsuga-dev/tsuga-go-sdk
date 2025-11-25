# ListRoutes200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the log route | 
**Name** | **string** | Human readable name shown for the route | 
**Description** | Pointer to **string** |  | [optional] 
**IsEnabled** | **bool** |  | 
**Query** | **string** | Query that selects which logs should enter the route | 
**Tags** | Pointer to [**[]ListDashboards200ResponseDataInnerTagsInner**](ListDashboards200ResponseDataInnerTagsInner.md) | List of key/value tags applied to the resource | [optional] 
**Owner** | **string** | Team ID owning and managing the route | 
**Processors** | [**[]ListRoutes200ResponseDataInnerProcessorsInner**](ListRoutes200ResponseDataInnerProcessorsInner.md) | Ordered processors applied to logs that match the route | 

## Methods

### NewListRoutes200ResponseDataInner

`func NewListRoutes200ResponseDataInner(id string, name string, isEnabled bool, query string, owner string, processors []ListRoutes200ResponseDataInnerProcessorsInner, ) *ListRoutes200ResponseDataInner`

NewListRoutes200ResponseDataInner instantiates a new ListRoutes200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRoutes200ResponseDataInnerWithDefaults

`func NewListRoutes200ResponseDataInnerWithDefaults() *ListRoutes200ResponseDataInner`

NewListRoutes200ResponseDataInnerWithDefaults instantiates a new ListRoutes200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListRoutes200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListRoutes200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListRoutes200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ListRoutes200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListRoutes200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListRoutes200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ListRoutes200ResponseDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListRoutes200ResponseDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListRoutes200ResponseDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListRoutes200ResponseDataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsEnabled

`func (o *ListRoutes200ResponseDataInner) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ListRoutes200ResponseDataInner) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ListRoutes200ResponseDataInner) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetQuery

`func (o *ListRoutes200ResponseDataInner) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ListRoutes200ResponseDataInner) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ListRoutes200ResponseDataInner) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetTags

`func (o *ListRoutes200ResponseDataInner) GetTags() []ListDashboards200ResponseDataInnerTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ListRoutes200ResponseDataInner) GetTagsOk() (*[]ListDashboards200ResponseDataInnerTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ListRoutes200ResponseDataInner) SetTags(v []ListDashboards200ResponseDataInnerTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ListRoutes200ResponseDataInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOwner

`func (o *ListRoutes200ResponseDataInner) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListRoutes200ResponseDataInner) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListRoutes200ResponseDataInner) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetProcessors

`func (o *ListRoutes200ResponseDataInner) GetProcessors() []ListRoutes200ResponseDataInnerProcessorsInner`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *ListRoutes200ResponseDataInner) GetProcessorsOk() (*[]ListRoutes200ResponseDataInnerProcessorsInner, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *ListRoutes200ResponseDataInner) SetProcessors(v []ListRoutes200ResponseDataInnerProcessorsInner)`

SetProcessors sets Processors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


