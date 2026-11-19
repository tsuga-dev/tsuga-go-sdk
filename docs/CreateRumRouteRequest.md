# CreateRumRouteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable RUM route name shown in Tsuga. | 
**Description** | Pointer to **NullableString** | Message included in notifications triggered by this monitor. Optional on create. | [optional] 
**IsEnabled** | **bool** | Set to true for the RUM route to process newly ingested matching RUM events. Set to false to save it without running it. | 
**Query** | **string** | Tsuga query that selects which newly ingested RUM events enter the RUM route processor chain. | 
**Tags** | Pointer to [**[]InputTag**](InputTag.md) | Key/value tags to apply to the resource. Tag policies may require specific keys or values. | [optional] 
**Owner** | **string** | ID of the team that will own the RUM route, determining which team members can view and manage it and which tag policies apply. Only an org admin or an admin of this team may change fields other than &#x60;processors&#x60;, or delete the route. | 
**Processors** | [**[]InputRumProcessor**](InputRumProcessor.md) | Processors applied in array order to RUM events that match the RUM route, each transforming the output of the previous one. A &#x60;split&#x60; processor forks the chain into branches and cannot be followed by another processor at the same level. | 

## Methods

### NewCreateRumRouteRequest

`func NewCreateRumRouteRequest(name string, isEnabled bool, query string, owner string, processors []InputRumProcessor, ) *CreateRumRouteRequest`

NewCreateRumRouteRequest instantiates a new CreateRumRouteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRumRouteRequestWithDefaults

`func NewCreateRumRouteRequestWithDefaults() *CreateRumRouteRequest`

NewCreateRumRouteRequestWithDefaults instantiates a new CreateRumRouteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateRumRouteRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRumRouteRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRumRouteRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateRumRouteRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateRumRouteRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateRumRouteRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateRumRouteRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateRumRouteRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateRumRouteRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsEnabled

`func (o *CreateRumRouteRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *CreateRumRouteRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *CreateRumRouteRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetQuery

`func (o *CreateRumRouteRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CreateRumRouteRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CreateRumRouteRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetTags

`func (o *CreateRumRouteRequest) GetTags() []InputTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateRumRouteRequest) GetTagsOk() (*[]InputTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateRumRouteRequest) SetTags(v []InputTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateRumRouteRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOwner

`func (o *CreateRumRouteRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateRumRouteRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateRumRouteRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetProcessors

`func (o *CreateRumRouteRequest) GetProcessors() []InputRumProcessor`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *CreateRumRouteRequest) GetProcessorsOk() (*[]InputRumProcessor, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *CreateRumRouteRequest) SetProcessors(v []InputRumProcessor)`

SetProcessors sets Processors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


