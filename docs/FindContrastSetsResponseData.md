# FindContrastSetsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContrastSets** | [**[]ContrastSet**](ContrastSet.md) | One entry per attribute that separates the two groups, strongest first. Empty when nothing cleared the gates, which means the target and the baseline look alike on the attributes tested. | 
**TargetGroupCount** | **float32** | Number of spans matched by the target group filter over its time range. Not a sample: it is a plain count over the whole group. | 
**BaselineGroupCount** | **float32** | Number of spans matched by the baseline group filter over its time range. Not a sample: it is a plain count over the whole group. | 
**FailedAttributeCount** | **float32** | Number of candidate attributes Tsuga dropped because a query for them failed. Those attributes were not tested at all, so a high count next to an empty &#x60;contrastSets&#x60; means the comparison was incomplete, not that the two groups are alike. | 

## Methods

### NewFindContrastSetsResponseData

`func NewFindContrastSetsResponseData(contrastSets []ContrastSet, targetGroupCount float32, baselineGroupCount float32, failedAttributeCount float32, ) *FindContrastSetsResponseData`

NewFindContrastSetsResponseData instantiates a new FindContrastSetsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindContrastSetsResponseDataWithDefaults

`func NewFindContrastSetsResponseDataWithDefaults() *FindContrastSetsResponseData`

NewFindContrastSetsResponseDataWithDefaults instantiates a new FindContrastSetsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContrastSets

`func (o *FindContrastSetsResponseData) GetContrastSets() []ContrastSet`

GetContrastSets returns the ContrastSets field if non-nil, zero value otherwise.

### GetContrastSetsOk

`func (o *FindContrastSetsResponseData) GetContrastSetsOk() (*[]ContrastSet, bool)`

GetContrastSetsOk returns a tuple with the ContrastSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContrastSets

`func (o *FindContrastSetsResponseData) SetContrastSets(v []ContrastSet)`

SetContrastSets sets ContrastSets field to given value.


### GetTargetGroupCount

`func (o *FindContrastSetsResponseData) GetTargetGroupCount() float32`

GetTargetGroupCount returns the TargetGroupCount field if non-nil, zero value otherwise.

### GetTargetGroupCountOk

`func (o *FindContrastSetsResponseData) GetTargetGroupCountOk() (*float32, bool)`

GetTargetGroupCountOk returns a tuple with the TargetGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupCount

`func (o *FindContrastSetsResponseData) SetTargetGroupCount(v float32)`

SetTargetGroupCount sets TargetGroupCount field to given value.


### GetBaselineGroupCount

`func (o *FindContrastSetsResponseData) GetBaselineGroupCount() float32`

GetBaselineGroupCount returns the BaselineGroupCount field if non-nil, zero value otherwise.

### GetBaselineGroupCountOk

`func (o *FindContrastSetsResponseData) GetBaselineGroupCountOk() (*float32, bool)`

GetBaselineGroupCountOk returns a tuple with the BaselineGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineGroupCount

`func (o *FindContrastSetsResponseData) SetBaselineGroupCount(v float32)`

SetBaselineGroupCount sets BaselineGroupCount field to given value.


### GetFailedAttributeCount

`func (o *FindContrastSetsResponseData) GetFailedAttributeCount() float32`

GetFailedAttributeCount returns the FailedAttributeCount field if non-nil, zero value otherwise.

### GetFailedAttributeCountOk

`func (o *FindContrastSetsResponseData) GetFailedAttributeCountOk() (*float32, bool)`

GetFailedAttributeCountOk returns a tuple with the FailedAttributeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAttributeCount

`func (o *FindContrastSetsResponseData) SetFailedAttributeCount(v float32)`

SetFailedAttributeCount sets FailedAttributeCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


