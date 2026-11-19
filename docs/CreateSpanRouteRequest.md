# CreateSpanRouteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable span route name shown in Tsuga. | 
**Description** | Pointer to **NullableString** | Message included in notifications triggered by this monitor. Optional on create. | [optional] 
**IsEnabled** | **bool** | Set to true for the span route to process newly ingested matching spans. Set to false to save it without running it. | 
**Query** | **string** | Tsuga query that selects which newly ingested spans enter the span route processor chain. | 
**Tags** | Pointer to [**[]InputTag**](InputTag.md) | Key/value tags to apply to the resource. Tag policies may require specific keys or values. | [optional] 
**Owner** | **string** | ID of the team that will own the span route, determining which team members can view and manage it and which tag policies apply. Only an org admin or an admin of this team may change fields other than &#x60;processors&#x60;, or delete the route. | 
**Processors** | [**[]InputSpansProcessor**](InputSpansProcessor.md) | Processors applied in array order to spans that match the span route, each transforming the output of the previous one. A &#x60;split&#x60; processor forks the chain into branches and cannot be followed by another processor at the same level. | 

## Methods

### NewCreateSpanRouteRequest

`func NewCreateSpanRouteRequest(name string, isEnabled bool, query string, owner string, processors []InputSpansProcessor, ) *CreateSpanRouteRequest`

NewCreateSpanRouteRequest instantiates a new CreateSpanRouteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSpanRouteRequestWithDefaults

`func NewCreateSpanRouteRequestWithDefaults() *CreateSpanRouteRequest`

NewCreateSpanRouteRequestWithDefaults instantiates a new CreateSpanRouteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateSpanRouteRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSpanRouteRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSpanRouteRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateSpanRouteRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSpanRouteRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSpanRouteRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateSpanRouteRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateSpanRouteRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateSpanRouteRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsEnabled

`func (o *CreateSpanRouteRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *CreateSpanRouteRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *CreateSpanRouteRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetQuery

`func (o *CreateSpanRouteRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CreateSpanRouteRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CreateSpanRouteRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetTags

`func (o *CreateSpanRouteRequest) GetTags() []InputTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateSpanRouteRequest) GetTagsOk() (*[]InputTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateSpanRouteRequest) SetTags(v []InputTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateSpanRouteRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOwner

`func (o *CreateSpanRouteRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateSpanRouteRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateSpanRouteRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetProcessors

`func (o *CreateSpanRouteRequest) GetProcessors() []InputSpansProcessor`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *CreateSpanRouteRequest) GetProcessorsOk() (*[]InputSpansProcessor, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *CreateSpanRouteRequest) SetProcessors(v []InputSpansProcessor)`

SetProcessors sets Processors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


