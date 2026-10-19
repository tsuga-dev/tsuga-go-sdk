# QualityReportRowExamplesInnerLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | **string** | In-app route that opens the example in Tsuga. Set by the rule evaluator when the example has a product surface for inspection. | 
**Search** | Pointer to **interface{}** | Route search payload for the in-app link. Set when the destination needs query-string state such as a search query, time range, or table filter. | [optional] 
**Params** | Pointer to **interface{}** | Route params for the in-app link. Set when the destination route needs path parameters such as an asset identifier. | [optional] 

## Methods

### NewQualityReportRowExamplesInnerLink

`func NewQualityReportRowExamplesInnerLink(to string, ) *QualityReportRowExamplesInnerLink`

NewQualityReportRowExamplesInnerLink instantiates a new QualityReportRowExamplesInnerLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQualityReportRowExamplesInnerLinkWithDefaults

`func NewQualityReportRowExamplesInnerLinkWithDefaults() *QualityReportRowExamplesInnerLink`

NewQualityReportRowExamplesInnerLinkWithDefaults instantiates a new QualityReportRowExamplesInnerLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *QualityReportRowExamplesInnerLink) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *QualityReportRowExamplesInnerLink) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *QualityReportRowExamplesInnerLink) SetTo(v string)`

SetTo sets To field to given value.


### GetSearch

`func (o *QualityReportRowExamplesInnerLink) GetSearch() interface{}`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *QualityReportRowExamplesInnerLink) GetSearchOk() (*interface{}, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *QualityReportRowExamplesInnerLink) SetSearch(v interface{})`

SetSearch sets Search field to given value.

### HasSearch

`func (o *QualityReportRowExamplesInnerLink) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### SetSearchNil

`func (o *QualityReportRowExamplesInnerLink) SetSearchNil(b bool)`

 SetSearchNil sets the value for Search to be an explicit nil

### UnsetSearch
`func (o *QualityReportRowExamplesInnerLink) UnsetSearch()`

UnsetSearch ensures that no value is present for Search, not even an explicit nil
### GetParams

`func (o *QualityReportRowExamplesInnerLink) GetParams() interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *QualityReportRowExamplesInnerLink) GetParamsOk() (*interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *QualityReportRowExamplesInnerLink) SetParams(v interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *QualityReportRowExamplesInnerLink) HasParams() bool`

HasParams returns a boolean if a field has been set.

### SetParamsNil

`func (o *QualityReportRowExamplesInnerLink) SetParamsNil(b bool)`

 SetParamsNil sets the value for Params to be an explicit nil

### UnsetParams
`func (o *QualityReportRowExamplesInnerLink) UnsetParams()`

UnsetParams ensures that no value is present for Params, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


