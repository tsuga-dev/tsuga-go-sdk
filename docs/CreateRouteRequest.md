# CreateRouteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable log route name shown in Tsuga. | 
**Description** | Pointer to **NullableString** | Message included in notifications triggered by this monitor. Optional on create. | [optional] 
**IsEnabled** | **bool** | Set to true for the log route to process newly ingested matching logs. Set to false to save it without running it. | 
**Query** | **string** | Tsuga query that selects which newly ingested logs enter the log route processor chain. | 
**Tags** | Pointer to [**[]InputTag**](InputTag.md) | Key/value tags to apply to the resource. Tag policies may require specific keys or values. | [optional] 
**Owner** | **string** | ID of the team that will own the log route, determining which team members can view and manage it and which tag policies apply. Only an org admin or an admin of this team may change fields other than &#x60;processors&#x60;, or delete the route. | 
**Processors** | [**[]InputProcessor**](InputProcessor.md) | Processors applied in array order to logs that match the log route, each transforming the output of the previous one. Up to 50 are accepted. A &#x60;split&#x60; processor forks the chain into branches and cannot be followed by another processor at the same level. | 

## Methods

### NewCreateRouteRequest

`func NewCreateRouteRequest(name string, isEnabled bool, query string, owner string, processors []InputProcessor, ) *CreateRouteRequest`

NewCreateRouteRequest instantiates a new CreateRouteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRouteRequestWithDefaults

`func NewCreateRouteRequestWithDefaults() *CreateRouteRequest`

NewCreateRouteRequestWithDefaults instantiates a new CreateRouteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateRouteRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRouteRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRouteRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateRouteRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateRouteRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateRouteRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateRouteRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateRouteRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateRouteRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsEnabled

`func (o *CreateRouteRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *CreateRouteRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *CreateRouteRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetQuery

`func (o *CreateRouteRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CreateRouteRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CreateRouteRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetTags

`func (o *CreateRouteRequest) GetTags() []InputTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateRouteRequest) GetTagsOk() (*[]InputTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateRouteRequest) SetTags(v []InputTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateRouteRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOwner

`func (o *CreateRouteRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateRouteRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateRouteRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetProcessors

`func (o *CreateRouteRequest) GetProcessors() []InputProcessor`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *CreateRouteRequest) GetProcessorsOk() (*[]InputProcessor, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *CreateRouteRequest) SetProcessors(v []InputProcessor)`

SetProcessors sets Processors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


