# QualityReportRowExamplesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Tsuga asset or telemetry kind referenced by this example. Set by the rule evaluator from the failing object, record, or lookup result; examples include &#x60;monitor&#x60;, &#x60;dashboard&#x60;, &#x60;metric&#x60;, &#x60;log&#x60;, and &#x60;api-key&#x60;. | 
**Label** | **string** | Human-readable summary of the example. Set by the rule evaluator from the asset name, telemetry identifier, or lookup result so clients can display the example without following the link. | 
**Link** | [**QualityReportRowExamplesInnerLink**](QualityReportRowExamplesInnerLink.md) |  | 

## Methods

### NewQualityReportRowExamplesInner

`func NewQualityReportRowExamplesInner(type_ string, label string, link QualityReportRowExamplesInnerLink, ) *QualityReportRowExamplesInner`

NewQualityReportRowExamplesInner instantiates a new QualityReportRowExamplesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQualityReportRowExamplesInnerWithDefaults

`func NewQualityReportRowExamplesInnerWithDefaults() *QualityReportRowExamplesInner`

NewQualityReportRowExamplesInnerWithDefaults instantiates a new QualityReportRowExamplesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *QualityReportRowExamplesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QualityReportRowExamplesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QualityReportRowExamplesInner) SetType(v string)`

SetType sets Type field to given value.


### GetLabel

`func (o *QualityReportRowExamplesInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *QualityReportRowExamplesInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *QualityReportRowExamplesInner) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetLink

`func (o *QualityReportRowExamplesInner) GetLink() QualityReportRowExamplesInnerLink`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *QualityReportRowExamplesInner) GetLinkOk() (*QualityReportRowExamplesInnerLink, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *QualityReportRowExamplesInner) SetLink(v QualityReportRowExamplesInnerLink)`

SetLink sets Link field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


