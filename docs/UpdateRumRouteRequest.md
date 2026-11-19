# UpdateRumRouteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable RUM route name shown in Tsuga. | 
**Description** | Pointer to **NullableString** | Free-text note about the RUM route&#39;s purpose. This metadata is not evaluated when matching or processing RUM events. | [optional] 
**IsEnabled** | **bool** | Set to true for the RUM route to process newly ingested matching RUM events. Set to false to save it without running it. | 
**Query** | **string** | Tsuga query that selects which newly ingested RUM events enter the RUM route processor chain. | 
**Tags** | Pointer to [**[]InputTag**](InputTag.md) | Key/value tags to apply to the resource. Tag policies may require specific keys or values. | [optional] 
**Owner** | **string** | ID of the team that will own the RUM route, determining which team members can view and manage it and which tag policies apply. Only an org admin or an admin of this team may change fields other than &#x60;processors&#x60;, or delete the route. | 
**Processors** | [**[]InputRumProcessor**](InputRumProcessor.md) | Processors applied in array order to RUM events that match the RUM route, each transforming the output of the previous one. A &#x60;split&#x60; processor forks the chain into branches and cannot be followed by another processor at the same level. | 

## Methods

### NewUpdateRumRouteRequest

`func NewUpdateRumRouteRequest(name string, isEnabled bool, query string, owner string, processors []InputRumProcessor, ) *UpdateRumRouteRequest`

NewUpdateRumRouteRequest instantiates a new UpdateRumRouteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRumRouteRequestWithDefaults

`func NewUpdateRumRouteRequestWithDefaults() *UpdateRumRouteRequest`

NewUpdateRumRouteRequestWithDefaults instantiates a new UpdateRumRouteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateRumRouteRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateRumRouteRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateRumRouteRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateRumRouteRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateRumRouteRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateRumRouteRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateRumRouteRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateRumRouteRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateRumRouteRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsEnabled

`func (o *UpdateRumRouteRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *UpdateRumRouteRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *UpdateRumRouteRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetQuery

`func (o *UpdateRumRouteRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *UpdateRumRouteRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *UpdateRumRouteRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetTags

`func (o *UpdateRumRouteRequest) GetTags() []InputTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateRumRouteRequest) GetTagsOk() (*[]InputTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateRumRouteRequest) SetTags(v []InputTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateRumRouteRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOwner

`func (o *UpdateRumRouteRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateRumRouteRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateRumRouteRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetProcessors

`func (o *UpdateRumRouteRequest) GetProcessors() []InputRumProcessor`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *UpdateRumRouteRequest) GetProcessorsOk() (*[]InputRumProcessor, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *UpdateRumRouteRequest) SetProcessors(v []InputRumProcessor)`

SetProcessors sets Processors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


