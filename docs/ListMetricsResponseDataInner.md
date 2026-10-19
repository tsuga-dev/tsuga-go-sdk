# ListMetricsResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Metric name observed by Tsuga. Use this value in metric queries and aggregations. | 
**Type** | **string** | Metric instrument type, set by the reporting instrumentation. &#x60;counter&#x60; accumulates a monotonically increasing total (aggregate with &#x60;sum&#x60; or &#x60;increase&#x60;); &#x60;gauge&#x60; reports a point-in-time value that can go up or down (aggregate with &#x60;average&#x60;, &#x60;max&#x60;, &#x60;min&#x60;, or &#x60;last&#x60;); &#x60;histogram&#x60;, &#x60;composite_histogram&#x60;, and &#x60;summary&#x60; report a distribution of values, from which percentiles can be computed. Tsuga displays &#x60;composite_histogram&#x60; metrics as Histogram in the app. | 
**Description** | Pointer to **string** | Human-readable description of what the metric measures, set by the reporting instrumentation. Omitted when the instrumentation did not provide one. | [optional] 
**Unit** | Pointer to **string** | Metric unit reported by instrumentation or inferred metadata. | [optional] 
**Temporality** | Pointer to **string** | Metric temporality. &#x60;delta&#x60; reports changes per collection interval, &#x60;cumulative&#x60; reports running totals, and &#x60;none&#x60; means temporality is not applicable or unknown. | [optional] 
**Capabilities** | Pointer to **[]string** | Aggregate function types (such as &#x60;max&#x60;, &#x60;min&#x60;, &#x60;sum&#x60;) supported when querying or graphing this metric. An empty or omitted array means all aggregate types are supported. | [optional] 
**Attributes** | Pointer to **[]string** | Metric attribute names available as dimensions for filtering or grouping. | [optional] 

## Methods

### NewListMetricsResponseDataInner

`func NewListMetricsResponseDataInner(name string, type_ string, ) *ListMetricsResponseDataInner`

NewListMetricsResponseDataInner instantiates a new ListMetricsResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMetricsResponseDataInnerWithDefaults

`func NewListMetricsResponseDataInnerWithDefaults() *ListMetricsResponseDataInner`

NewListMetricsResponseDataInnerWithDefaults instantiates a new ListMetricsResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ListMetricsResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListMetricsResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListMetricsResponseDataInner) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ListMetricsResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListMetricsResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListMetricsResponseDataInner) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *ListMetricsResponseDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListMetricsResponseDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListMetricsResponseDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListMetricsResponseDataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUnit

`func (o *ListMetricsResponseDataInner) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ListMetricsResponseDataInner) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ListMetricsResponseDataInner) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *ListMetricsResponseDataInner) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetTemporality

`func (o *ListMetricsResponseDataInner) GetTemporality() string`

GetTemporality returns the Temporality field if non-nil, zero value otherwise.

### GetTemporalityOk

`func (o *ListMetricsResponseDataInner) GetTemporalityOk() (*string, bool)`

GetTemporalityOk returns a tuple with the Temporality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporality

`func (o *ListMetricsResponseDataInner) SetTemporality(v string)`

SetTemporality sets Temporality field to given value.

### HasTemporality

`func (o *ListMetricsResponseDataInner) HasTemporality() bool`

HasTemporality returns a boolean if a field has been set.

### GetCapabilities

`func (o *ListMetricsResponseDataInner) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ListMetricsResponseDataInner) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ListMetricsResponseDataInner) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *ListMetricsResponseDataInner) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetAttributes

`func (o *ListMetricsResponseDataInner) GetAttributes() []string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ListMetricsResponseDataInner) GetAttributesOk() (*[]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ListMetricsResponseDataInner) SetAttributes(v []string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ListMetricsResponseDataInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


