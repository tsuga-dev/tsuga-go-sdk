# FindContrastSetsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetGroup** | [**FindContrastSetsRequestTargetGroup**](FindContrastSetsRequestTargetGroup.md) |  | 
**BaselineGroup** | [**FindContrastSetsRequestBaselineGroup**](FindContrastSetsRequestBaselineGroup.md) |  | 
**CandidateAttrs** | Pointer to **[]string** | Span attribute dot-paths to test, such as &#x60;context.service.name&#x60; or &#x60;span_attributes.http.route&#x60;. At most 20. Attributes referenced by either group filter are always excluded, along with their known sibling attributes for the same field, so the condition being contrasted does not come back as a finding under one of those names. Omit or pass an empty array to let Tsuga sample both groups and test the attributes, most frequent in the target first, that take more than one value across the two samples. When none varies, nothing is tested and &#x60;contrastSets&#x60; is empty. | [optional] 
**TopK** | Pointer to **int32** | Maximum number of findings to return, strongest first. A finding is one attribute/value pair, and findings are grouped into one &#x60;contrastSets&#x60; entry per attribute, so a result carries at most this many &#x60;values&#x60; rows across possibly fewer entries. Valid values are 1 through 100. Defaults to 10. | [optional] 
**MinLift** | Pointer to **float32** | Minimum lift a value must reach to be reported as a finding. Lift is the value&#39;s coverage of the target group divided by its coverage of the baseline, so &#x60;2&#x60; keeps only values at least twice as frequent in the target. Raise it to cut weak findings. Defaults to 1.5. | [optional] 
**MinSupport** | Pointer to **float32** | Support floor as a percentage: a value is reported as a finding only when it covers at least this much of the target group or of the baseline group. Valid values are 0 through 100. Raise it to drop findings that explain only a handful of spans. Defaults to 0.5. | [optional] 

## Methods

### NewFindContrastSetsRequest

`func NewFindContrastSetsRequest(targetGroup FindContrastSetsRequestTargetGroup, baselineGroup FindContrastSetsRequestBaselineGroup, ) *FindContrastSetsRequest`

NewFindContrastSetsRequest instantiates a new FindContrastSetsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindContrastSetsRequestWithDefaults

`func NewFindContrastSetsRequestWithDefaults() *FindContrastSetsRequest`

NewFindContrastSetsRequestWithDefaults instantiates a new FindContrastSetsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetGroup

`func (o *FindContrastSetsRequest) GetTargetGroup() FindContrastSetsRequestTargetGroup`

GetTargetGroup returns the TargetGroup field if non-nil, zero value otherwise.

### GetTargetGroupOk

`func (o *FindContrastSetsRequest) GetTargetGroupOk() (*FindContrastSetsRequestTargetGroup, bool)`

GetTargetGroupOk returns a tuple with the TargetGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroup

`func (o *FindContrastSetsRequest) SetTargetGroup(v FindContrastSetsRequestTargetGroup)`

SetTargetGroup sets TargetGroup field to given value.


### GetBaselineGroup

`func (o *FindContrastSetsRequest) GetBaselineGroup() FindContrastSetsRequestBaselineGroup`

GetBaselineGroup returns the BaselineGroup field if non-nil, zero value otherwise.

### GetBaselineGroupOk

`func (o *FindContrastSetsRequest) GetBaselineGroupOk() (*FindContrastSetsRequestBaselineGroup, bool)`

GetBaselineGroupOk returns a tuple with the BaselineGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineGroup

`func (o *FindContrastSetsRequest) SetBaselineGroup(v FindContrastSetsRequestBaselineGroup)`

SetBaselineGroup sets BaselineGroup field to given value.


### GetCandidateAttrs

`func (o *FindContrastSetsRequest) GetCandidateAttrs() []*string`

GetCandidateAttrs returns the CandidateAttrs field if non-nil, zero value otherwise.

### GetCandidateAttrsOk

`func (o *FindContrastSetsRequest) GetCandidateAttrsOk() (*[]*string, bool)`

GetCandidateAttrsOk returns a tuple with the CandidateAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidateAttrs

`func (o *FindContrastSetsRequest) SetCandidateAttrs(v []*string)`

SetCandidateAttrs sets CandidateAttrs field to given value.

### HasCandidateAttrs

`func (o *FindContrastSetsRequest) HasCandidateAttrs() bool`

HasCandidateAttrs returns a boolean if a field has been set.

### GetTopK

`func (o *FindContrastSetsRequest) GetTopK() int32`

GetTopK returns the TopK field if non-nil, zero value otherwise.

### GetTopKOk

`func (o *FindContrastSetsRequest) GetTopKOk() (*int32, bool)`

GetTopKOk returns a tuple with the TopK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopK

`func (o *FindContrastSetsRequest) SetTopK(v int32)`

SetTopK sets TopK field to given value.

### HasTopK

`func (o *FindContrastSetsRequest) HasTopK() bool`

HasTopK returns a boolean if a field has been set.

### GetMinLift

`func (o *FindContrastSetsRequest) GetMinLift() float32`

GetMinLift returns the MinLift field if non-nil, zero value otherwise.

### GetMinLiftOk

`func (o *FindContrastSetsRequest) GetMinLiftOk() (*float32, bool)`

GetMinLiftOk returns a tuple with the MinLift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLift

`func (o *FindContrastSetsRequest) SetMinLift(v float32)`

SetMinLift sets MinLift field to given value.

### HasMinLift

`func (o *FindContrastSetsRequest) HasMinLift() bool`

HasMinLift returns a boolean if a field has been set.

### GetMinSupport

`func (o *FindContrastSetsRequest) GetMinSupport() float32`

GetMinSupport returns the MinSupport field if non-nil, zero value otherwise.

### GetMinSupportOk

`func (o *FindContrastSetsRequest) GetMinSupportOk() (*float32, bool)`

GetMinSupportOk returns a tuple with the MinSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSupport

`func (o *FindContrastSetsRequest) SetMinSupport(v float32)`

SetMinSupport sets MinSupport field to given value.

### HasMinSupport

`func (o *FindContrastSetsRequest) HasMinSupport() bool`

HasMinSupport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


