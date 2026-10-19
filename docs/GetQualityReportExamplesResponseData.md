# GetQualityReportExamplesResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Examples** | [**[]QualityReportRowExamplesInner**](QualityReportRowExamplesInner.md) | Requested page of the examples stored for this row. Empty when the rule evaluation produced none, or when the offset is past the last one. | 

## Methods

### NewGetQualityReportExamplesResponseData

`func NewGetQualityReportExamplesResponseData(examples []QualityReportRowExamplesInner, ) *GetQualityReportExamplesResponseData`

NewGetQualityReportExamplesResponseData instantiates a new GetQualityReportExamplesResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetQualityReportExamplesResponseDataWithDefaults

`func NewGetQualityReportExamplesResponseDataWithDefaults() *GetQualityReportExamplesResponseData`

NewGetQualityReportExamplesResponseDataWithDefaults instantiates a new GetQualityReportExamplesResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExamples

`func (o *GetQualityReportExamplesResponseData) GetExamples() []QualityReportRowExamplesInner`

GetExamples returns the Examples field if non-nil, zero value otherwise.

### GetExamplesOk

`func (o *GetQualityReportExamplesResponseData) GetExamplesOk() (*[]QualityReportRowExamplesInner, bool)`

GetExamplesOk returns a tuple with the Examples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExamples

`func (o *GetQualityReportExamplesResponseData) SetExamples(v []QualityReportRowExamplesInner)`

SetExamples sets Examples field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


