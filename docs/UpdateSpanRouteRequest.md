# UpdateSpanRouteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable span route name shown in Tsuga. | 
**Description** | Pointer to **NullableString** | Free-text note about the span route&#39;s purpose. This metadata is not evaluated when matching or processing spans. | [optional] 
**IsEnabled** | **bool** | Set to true for the span route to process newly ingested matching spans. Set to false to save it without running it. | 
**Query** | **string** | Tsuga query that selects which newly ingested spans enter the span route processor chain. | 
**Tags** | Pointer to [**[]InputTag**](InputTag.md) | Key/value tags to apply to the resource. Tag policies may require specific keys or values. | [optional] 
**Owner** | **string** | ID of the team that will own the span route, determining which team members can view and manage it and which tag policies apply. Only an org admin or an admin of this team may change fields other than &#x60;processors&#x60;, or delete the route. | 
**Processors** | [**[]InputSpansProcessor**](InputSpansProcessor.md) | Processors applied in array order to spans that match the span route, each transforming the output of the previous one. A &#x60;split&#x60; processor forks the chain into branches and cannot be followed by another processor at the same level. | 

## Methods

### NewUpdateSpanRouteRequest

`func NewUpdateSpanRouteRequest(name string, isEnabled bool, query string, owner string, processors []InputSpansProcessor, ) *UpdateSpanRouteRequest`

NewUpdateSpanRouteRequest instantiates a new UpdateSpanRouteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSpanRouteRequestWithDefaults

`func NewUpdateSpanRouteRequestWithDefaults() *UpdateSpanRouteRequest`

NewUpdateSpanRouteRequestWithDefaults instantiates a new UpdateSpanRouteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateSpanRouteRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSpanRouteRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSpanRouteRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateSpanRouteRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateSpanRouteRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateSpanRouteRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateSpanRouteRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateSpanRouteRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateSpanRouteRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsEnabled

`func (o *UpdateSpanRouteRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *UpdateSpanRouteRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *UpdateSpanRouteRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetQuery

`func (o *UpdateSpanRouteRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *UpdateSpanRouteRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *UpdateSpanRouteRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetTags

`func (o *UpdateSpanRouteRequest) GetTags() []InputTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateSpanRouteRequest) GetTagsOk() (*[]InputTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateSpanRouteRequest) SetTags(v []InputTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateSpanRouteRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOwner

`func (o *UpdateSpanRouteRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateSpanRouteRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateSpanRouteRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetProcessors

`func (o *UpdateSpanRouteRequest) GetProcessors() []InputSpansProcessor`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *UpdateSpanRouteRequest) GetProcessorsOk() (*[]InputSpansProcessor, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *UpdateSpanRouteRequest) SetProcessors(v []InputSpansProcessor)`

SetProcessors sets Processors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


