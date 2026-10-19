# UpdateMetricRouteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable metrics route name shown in Tsuga. | 
**Description** | Pointer to **NullableString** | Optional free-text note about the metrics route&#39;s purpose. This metadata is not evaluated when matching or processing metrics. | [optional] 
**IsEnabled** | **bool** | Set to true for the route to process newly ingested matching metrics. Set to false to save it without running it. | 
**MetricNameRegex** | **string** | Regex matched against metric names to decide which metrics enter the route processor chain. Supported syntax is the Rust regex subset used for metric names: literals, &#x60;.&#x60;, character classes like &#x60;[a-z]&#x60;, escaped classes like &#x60;\\d&#x60;, &#x60;\\w&#x60;, &#x60;\\s&#x60;, anchors, groups, alternation, and repetitions like &#x60;*&#x60;, &#x60;+&#x60;, &#x60;?&#x60;, &#x60;{2}&#x60;, &#x60;{2,5}&#x60;. Escape regex metacharacters (&#x60;\\ . + * ? ( ) | [ ] { } ^ $&#x60;) to match them literally. Inline flags, look-around, backreferences, atomic groups, and possessive quantifiers are rejected. | 
**Tags** | Pointer to [**[]InputTag**](InputTag.md) | Key/value tags to apply to the resource. Tag policies may require specific keys or values. | [optional] 
**Owner** | **string** | ID of the team that will own the metrics route, determining which team members can view and manage it and which tag policies apply. Only an org admin or an admin of this team may change fields other than &#x60;processors&#x60;, or delete the route. | 
**Processors** | [**[]UpdateMetricRouteRequestProcessorsInner**](UpdateMetricRouteRequestProcessorsInner.md) | Processors applied in array order to metrics that match the route, each transforming the output of the previous one. Up to 50 are accepted. Supported processors are create format-string, create category, map attributes, and parse grok. | 

## Methods

### NewUpdateMetricRouteRequest

`func NewUpdateMetricRouteRequest(name string, isEnabled bool, metricNameRegex string, owner string, processors []UpdateMetricRouteRequestProcessorsInner, ) *UpdateMetricRouteRequest`

NewUpdateMetricRouteRequest instantiates a new UpdateMetricRouteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMetricRouteRequestWithDefaults

`func NewUpdateMetricRouteRequestWithDefaults() *UpdateMetricRouteRequest`

NewUpdateMetricRouteRequestWithDefaults instantiates a new UpdateMetricRouteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateMetricRouteRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateMetricRouteRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateMetricRouteRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateMetricRouteRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateMetricRouteRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateMetricRouteRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateMetricRouteRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateMetricRouteRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateMetricRouteRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsEnabled

`func (o *UpdateMetricRouteRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *UpdateMetricRouteRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *UpdateMetricRouteRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetMetricNameRegex

`func (o *UpdateMetricRouteRequest) GetMetricNameRegex() string`

GetMetricNameRegex returns the MetricNameRegex field if non-nil, zero value otherwise.

### GetMetricNameRegexOk

`func (o *UpdateMetricRouteRequest) GetMetricNameRegexOk() (*string, bool)`

GetMetricNameRegexOk returns a tuple with the MetricNameRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricNameRegex

`func (o *UpdateMetricRouteRequest) SetMetricNameRegex(v string)`

SetMetricNameRegex sets MetricNameRegex field to given value.


### GetTags

`func (o *UpdateMetricRouteRequest) GetTags() []InputTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateMetricRouteRequest) GetTagsOk() (*[]InputTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateMetricRouteRequest) SetTags(v []InputTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateMetricRouteRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOwner

`func (o *UpdateMetricRouteRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateMetricRouteRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateMetricRouteRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetProcessors

`func (o *UpdateMetricRouteRequest) GetProcessors() []UpdateMetricRouteRequestProcessorsInner`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *UpdateMetricRouteRequest) GetProcessorsOk() (*[]UpdateMetricRouteRequestProcessorsInner, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *UpdateMetricRouteRequest) SetProcessors(v []UpdateMetricRouteRequestProcessorsInner)`

SetProcessors sets Processors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


