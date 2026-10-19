# MetricRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Tsuga-generated metrics route ID assigned when the route is created. | 
**Name** | **string** | Human-readable metrics route name shown in Tsuga. | 
**Description** | Pointer to **string** | Optional free-text note about the metrics route&#39;s purpose, set by the caller. This metadata is not evaluated when matching or processing metrics. | [optional] 
**IsEnabled** | **bool** | Whether this route is active for newly ingested metrics. Disabled routes remain saved but do not process metrics. | 
**MetricNameRegex** | **string** | Regex matched against metric names to decide which metrics enter the route processor chain. Supported syntax is the Rust regex subset used for metric names: literals, &#x60;.&#x60;, character classes like &#x60;[a-z]&#x60;, escaped classes like &#x60;\\d&#x60;, &#x60;\\w&#x60;, &#x60;\\s&#x60;, anchors, groups, alternation, and repetitions like &#x60;*&#x60;, &#x60;+&#x60;, &#x60;?&#x60;, &#x60;{2}&#x60;, &#x60;{2,5}&#x60;. Escape regex metacharacters (&#x60;\\ . + * ? ( ) | [ ] { } ^ $&#x60;) to match them literally. Inline flags, look-around, backreferences, atomic groups, and possessive quantifiers are rejected. | 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Key/value tags applied to the resource. Use them to organize resources and to satisfy tag policies. | [optional] 
**Owner** | **string** | ID of the team that owns the metrics route. Determines which team members can view and manage it. Only an org admin or an admin of this team may change fields other than &#x60;processors&#x60;, or delete the route. | 
**Processors** | [**[]MetricRouteProcessorsInner**](MetricRouteProcessorsInner.md) | Processors applied in array order to metrics that match the route, each transforming the output of the previous one. | 

## Methods

### NewMetricRoute

`func NewMetricRoute(id string, name string, isEnabled bool, metricNameRegex string, owner string, processors []MetricRouteProcessorsInner, ) *MetricRoute`

NewMetricRoute instantiates a new MetricRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricRouteWithDefaults

`func NewMetricRouteWithDefaults() *MetricRoute`

NewMetricRouteWithDefaults instantiates a new MetricRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetricRoute) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetricRoute) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetricRoute) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *MetricRoute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricRoute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricRoute) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *MetricRoute) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricRoute) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricRoute) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricRoute) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsEnabled

`func (o *MetricRoute) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *MetricRoute) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *MetricRoute) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetMetricNameRegex

`func (o *MetricRoute) GetMetricNameRegex() string`

GetMetricNameRegex returns the MetricNameRegex field if non-nil, zero value otherwise.

### GetMetricNameRegexOk

`func (o *MetricRoute) GetMetricNameRegexOk() (*string, bool)`

GetMetricNameRegexOk returns a tuple with the MetricNameRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricNameRegex

`func (o *MetricRoute) SetMetricNameRegex(v string)`

SetMetricNameRegex sets MetricNameRegex field to given value.


### GetTags

`func (o *MetricRoute) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MetricRoute) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MetricRoute) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MetricRoute) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOwner

`func (o *MetricRoute) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *MetricRoute) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *MetricRoute) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetProcessors

`func (o *MetricRoute) GetProcessors() []MetricRouteProcessorsInner`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *MetricRoute) GetProcessorsOk() (*[]MetricRouteProcessorsInner, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *MetricRoute) SetProcessors(v []MetricRouteProcessorsInner)`

SetProcessors sets Processors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


