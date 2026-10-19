# Route

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Tsuga-generated log route ID assigned when the route is created. | 
**Name** | **string** | Human-readable route name shown in Tsuga. | 
**Description** | Pointer to **string** | Optional free-text note about the route&#39;s purpose, set by the caller. This metadata is not evaluated when matching or processing logs. | [optional] 
**IsEnabled** | **bool** | Whether this route is active for newly ingested logs. Disabled routes remain saved but do not process logs. | 
**Query** | **string** | Tsuga query that selects which newly ingested logs enter the route processor chain. | 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Key/value tags applied to the resource. Use them to organize resources and to satisfy tag policies. | [optional] 
**Owner** | **string** | ID of the team that owns the route. Determines which team members can view and manage it, and which tag policies apply to it. Only an org admin or an admin of this team may change fields other than &#x60;processors&#x60;, or delete the route. | 
**Processors** | [**[]OutputProcessor**](OutputProcessor.md) | Processors applied in array order to logs that match the route, each transforming the output of the previous one. A &#x60;split&#x60; processor forks the chain into branches and cannot be followed by another processor at the same level. | 

## Methods

### NewRoute

`func NewRoute(id string, name string, isEnabled bool, query string, owner string, processors []OutputProcessor, ) *Route`

NewRoute instantiates a new Route object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteWithDefaults

`func NewRouteWithDefaults() *Route`

NewRouteWithDefaults instantiates a new Route object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Route) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Route) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Route) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Route) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Route) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Route) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Route) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Route) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Route) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Route) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsEnabled

`func (o *Route) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *Route) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *Route) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetQuery

`func (o *Route) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Route) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Route) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetTags

`func (o *Route) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Route) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Route) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Route) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOwner

`func (o *Route) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Route) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Route) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetProcessors

`func (o *Route) GetProcessors() []OutputProcessor`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *Route) GetProcessorsOk() (*[]OutputProcessor, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *Route) SetProcessors(v []OutputProcessor)`

SetProcessors sets Processors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


