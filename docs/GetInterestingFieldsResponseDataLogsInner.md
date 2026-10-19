# GetInterestingFieldsResponseDataLogsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Field name to use in queries. For &#x60;kind: attribute&#x60;, this is an attribute dot path. For &#x60;kind: metric&#x60;, this is either a metric name or a &#x60;metricName.context.attribute&#x60; path when the metric query used a metric attribute. | 
**Kind** | **string** | Whether the field is a telemetry attribute or a metric name. &#x60;metric&#x60; appears only in the metrics bucket, for fields referenced as the aggregated field of a metrics-source query. | 
**InferredType** | **string** | Value type inferred from how the field is used: &#x60;numeric&#x60; when aggregated with sum/average/percentile, referenced by numeric or range filters, or a metric name; &#x60;categorical&#x60; when only grouped by or matched with equality/text filters; &#x60;unknown&#x60; otherwise. Evidence-based — a numeric field used only in equality filters reads as categorical. | 
**Score** | **float32** | Relative relevance score computed by Tsuga from how strongly readable dashboards and monitors use this field for the requested scope, discounted when the field is common across the whole organization. Higher values mean stronger evidence that the field is useful for this scope. Scores are intended for ranking, not for comparison across unrelated scopes. | 
**Artifacts** | **float32** | Number of distinct readable dashboards and monitors that reference this field and match the requested scope. Tsuga counts an artifact once per field even if it uses the field in multiple widgets or queries. | 
**GlobalArtifacts** | **float32** | Number of organization-wide dashboards and monitors that reference this field, regardless of the requested scope or caller permissions. Tsuga uses this anonymous baseline to down-rank fields that are common everywhere. | 
**Usage** | [**GetInterestingFieldsResponseDataLogsInnerUsage**](GetInterestingFieldsResponseDataLogsInnerUsage.md) |  | 

## Methods

### NewGetInterestingFieldsResponseDataLogsInner

`func NewGetInterestingFieldsResponseDataLogsInner(field string, kind string, inferredType string, score float32, artifacts float32, globalArtifacts float32, usage GetInterestingFieldsResponseDataLogsInnerUsage, ) *GetInterestingFieldsResponseDataLogsInner`

NewGetInterestingFieldsResponseDataLogsInner instantiates a new GetInterestingFieldsResponseDataLogsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInterestingFieldsResponseDataLogsInnerWithDefaults

`func NewGetInterestingFieldsResponseDataLogsInnerWithDefaults() *GetInterestingFieldsResponseDataLogsInner`

NewGetInterestingFieldsResponseDataLogsInnerWithDefaults instantiates a new GetInterestingFieldsResponseDataLogsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *GetInterestingFieldsResponseDataLogsInner) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *GetInterestingFieldsResponseDataLogsInner) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *GetInterestingFieldsResponseDataLogsInner) SetField(v string)`

SetField sets Field field to given value.


### GetKind

`func (o *GetInterestingFieldsResponseDataLogsInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetInterestingFieldsResponseDataLogsInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetInterestingFieldsResponseDataLogsInner) SetKind(v string)`

SetKind sets Kind field to given value.


### GetInferredType

`func (o *GetInterestingFieldsResponseDataLogsInner) GetInferredType() string`

GetInferredType returns the InferredType field if non-nil, zero value otherwise.

### GetInferredTypeOk

`func (o *GetInterestingFieldsResponseDataLogsInner) GetInferredTypeOk() (*string, bool)`

GetInferredTypeOk returns a tuple with the InferredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferredType

`func (o *GetInterestingFieldsResponseDataLogsInner) SetInferredType(v string)`

SetInferredType sets InferredType field to given value.


### GetScore

`func (o *GetInterestingFieldsResponseDataLogsInner) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *GetInterestingFieldsResponseDataLogsInner) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *GetInterestingFieldsResponseDataLogsInner) SetScore(v float32)`

SetScore sets Score field to given value.


### GetArtifacts

`func (o *GetInterestingFieldsResponseDataLogsInner) GetArtifacts() float32`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *GetInterestingFieldsResponseDataLogsInner) GetArtifactsOk() (*float32, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *GetInterestingFieldsResponseDataLogsInner) SetArtifacts(v float32)`

SetArtifacts sets Artifacts field to given value.


### GetGlobalArtifacts

`func (o *GetInterestingFieldsResponseDataLogsInner) GetGlobalArtifacts() float32`

GetGlobalArtifacts returns the GlobalArtifacts field if non-nil, zero value otherwise.

### GetGlobalArtifactsOk

`func (o *GetInterestingFieldsResponseDataLogsInner) GetGlobalArtifactsOk() (*float32, bool)`

GetGlobalArtifactsOk returns a tuple with the GlobalArtifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalArtifacts

`func (o *GetInterestingFieldsResponseDataLogsInner) SetGlobalArtifacts(v float32)`

SetGlobalArtifacts sets GlobalArtifacts field to given value.


### GetUsage

`func (o *GetInterestingFieldsResponseDataLogsInner) GetUsage() GetInterestingFieldsResponseDataLogsInnerUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *GetInterestingFieldsResponseDataLogsInner) GetUsageOk() (*GetInterestingFieldsResponseDataLogsInnerUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *GetInterestingFieldsResponseDataLogsInner) SetUsage(v GetInterestingFieldsResponseDataLogsInnerUsage)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


