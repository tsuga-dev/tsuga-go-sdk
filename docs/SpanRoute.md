# SpanRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Tsuga-generated span route ID assigned when the route is created. | 
**Name** | **string** | Human-readable span route name shown in Tsuga. | 
**Description** | Pointer to **string** | Free-text note about the span route&#39;s purpose, set by the caller. This metadata is not evaluated when matching or processing spans. | [optional] 
**IsEnabled** | **bool** | Whether this route is active for newly ingested spans. Disabled routes remain saved but do not process spans. | 
**Query** | **string** | Tsuga query that selects which newly ingested spans enter the route processor chain. | 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Key/value tags applied to the resource. Use them to organize resources and to satisfy tag policies. | [optional] 
**Owner** | **string** | ID of the team that owns the span route. Determines which team members can view and manage it. Only an org admin or an admin of this team may change fields other than &#x60;processors&#x60;, or delete the route. | 
**Processors** | [**[]OutputSpansProcessor**](OutputSpansProcessor.md) | Processors applied in array order to spans that match the route, each transforming the output of the previous one. | 

## Methods

### NewSpanRoute

`func NewSpanRoute(id string, name string, isEnabled bool, query string, owner string, processors []OutputSpansProcessor, ) *SpanRoute`

NewSpanRoute instantiates a new SpanRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanRouteWithDefaults

`func NewSpanRouteWithDefaults() *SpanRoute`

NewSpanRouteWithDefaults instantiates a new SpanRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SpanRoute) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpanRoute) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SpanRoute) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SpanRoute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpanRoute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpanRoute) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SpanRoute) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SpanRoute) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SpanRoute) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SpanRoute) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsEnabled

`func (o *SpanRoute) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *SpanRoute) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *SpanRoute) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetQuery

`func (o *SpanRoute) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SpanRoute) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SpanRoute) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetTags

`func (o *SpanRoute) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SpanRoute) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SpanRoute) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SpanRoute) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOwner

`func (o *SpanRoute) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *SpanRoute) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *SpanRoute) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetProcessors

`func (o *SpanRoute) GetProcessors() []OutputSpansProcessor`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *SpanRoute) GetProcessorsOk() (*[]OutputSpansProcessor, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *SpanRoute) SetProcessors(v []OutputSpansProcessor)`

SetProcessors sets Processors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


