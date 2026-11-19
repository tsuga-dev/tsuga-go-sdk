# RumRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Tsuga-generated RUM route ID assigned when the route is created. | 
**Name** | **string** | Human-readable RUM route name shown in Tsuga. | 
**Description** | Pointer to **string** | Free-text note about the RUM route&#39;s purpose, set by the caller. This metadata is not evaluated when matching or processing RUM events. | [optional] 
**IsEnabled** | **bool** | Whether this route is active for newly ingested RUM events. Disabled routes remain saved but do not process RUM events. | 
**Query** | **string** | Tsuga query that selects which newly ingested RUM events enter the route processor chain. | 
**Rank** | **float32** | Tsuga-assigned evaluation order, unique across all RUM routes and dense (no gaps). | 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Key/value tags applied to the resource. Use them to organize resources and to satisfy tag policies. | [optional] 
**Owner** | **string** | ID of the team that owns the RUM route. Determines which team members can view and manage it. Only an org admin or an admin of this team may change fields other than &#x60;processors&#x60;, or delete the route. | 
**Processors** | [**[]OutputRumProcessor**](OutputRumProcessor.md) | Processors applied in array order to RUM events that match the route, each transforming the output of the previous one. | 

## Methods

### NewRumRoute

`func NewRumRoute(id string, name string, isEnabled bool, query string, rank float32, owner string, processors []OutputRumProcessor, ) *RumRoute`

NewRumRoute instantiates a new RumRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRumRouteWithDefaults

`func NewRumRouteWithDefaults() *RumRoute`

NewRumRouteWithDefaults instantiates a new RumRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RumRoute) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RumRoute) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RumRoute) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *RumRoute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RumRoute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RumRoute) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RumRoute) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RumRoute) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RumRoute) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RumRoute) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsEnabled

`func (o *RumRoute) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *RumRoute) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *RumRoute) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetQuery

`func (o *RumRoute) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *RumRoute) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *RumRoute) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetRank

`func (o *RumRoute) GetRank() float32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *RumRoute) GetRankOk() (*float32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *RumRoute) SetRank(v float32)`

SetRank sets Rank field to given value.


### GetTags

`func (o *RumRoute) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RumRoute) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RumRoute) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RumRoute) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOwner

`func (o *RumRoute) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *RumRoute) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *RumRoute) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetProcessors

`func (o *RumRoute) GetProcessors() []OutputRumProcessor`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *RumRoute) GetProcessorsOk() (*[]OutputRumProcessor, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *RumRoute) SetProcessors(v []OutputRumProcessor)`

SetProcessors sets Processors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


