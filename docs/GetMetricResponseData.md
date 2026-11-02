# GetMetricResponseData

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
**SeriesCardinality** | Pointer to **float32** | Estimated number of distinct series for this metric across the most recent data generations kept by the metric catalog. Omitted when the catalog has no statistics for the metric yet. | [optional] 
**AttributeCardinalities** | Pointer to **map[string]float32** | Estimated number of distinct values per attribute over the same window as &#x60;seriesCardinality&#x60;, keyed like &#x60;attributes&#x60;. Omitted when the catalog has no statistics for the metric yet. | [optional] 

## Methods

### NewGetMetricResponseData

`func NewGetMetricResponseData(name string, type_ string, ) *GetMetricResponseData`

NewGetMetricResponseData instantiates a new GetMetricResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMetricResponseDataWithDefaults

`func NewGetMetricResponseDataWithDefaults() *GetMetricResponseData`

NewGetMetricResponseDataWithDefaults instantiates a new GetMetricResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetMetricResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetMetricResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetMetricResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *GetMetricResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetMetricResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetMetricResponseData) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *GetMetricResponseData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetMetricResponseData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetMetricResponseData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetMetricResponseData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUnit

`func (o *GetMetricResponseData) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *GetMetricResponseData) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *GetMetricResponseData) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *GetMetricResponseData) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetTemporality

`func (o *GetMetricResponseData) GetTemporality() string`

GetTemporality returns the Temporality field if non-nil, zero value otherwise.

### GetTemporalityOk

`func (o *GetMetricResponseData) GetTemporalityOk() (*string, bool)`

GetTemporalityOk returns a tuple with the Temporality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporality

`func (o *GetMetricResponseData) SetTemporality(v string)`

SetTemporality sets Temporality field to given value.

### HasTemporality

`func (o *GetMetricResponseData) HasTemporality() bool`

HasTemporality returns a boolean if a field has been set.

### GetCapabilities

`func (o *GetMetricResponseData) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *GetMetricResponseData) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *GetMetricResponseData) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *GetMetricResponseData) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetAttributes

`func (o *GetMetricResponseData) GetAttributes() []string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GetMetricResponseData) GetAttributesOk() (*[]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GetMetricResponseData) SetAttributes(v []string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GetMetricResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetSeriesCardinality

`func (o *GetMetricResponseData) GetSeriesCardinality() float32`

GetSeriesCardinality returns the SeriesCardinality field if non-nil, zero value otherwise.

### GetSeriesCardinalityOk

`func (o *GetMetricResponseData) GetSeriesCardinalityOk() (*float32, bool)`

GetSeriesCardinalityOk returns a tuple with the SeriesCardinality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesCardinality

`func (o *GetMetricResponseData) SetSeriesCardinality(v float32)`

SetSeriesCardinality sets SeriesCardinality field to given value.

### HasSeriesCardinality

`func (o *GetMetricResponseData) HasSeriesCardinality() bool`

HasSeriesCardinality returns a boolean if a field has been set.

### GetAttributeCardinalities

`func (o *GetMetricResponseData) GetAttributeCardinalities() map[string]float32`

GetAttributeCardinalities returns the AttributeCardinalities field if non-nil, zero value otherwise.

### GetAttributeCardinalitiesOk

`func (o *GetMetricResponseData) GetAttributeCardinalitiesOk() (*map[string]float32, bool)`

GetAttributeCardinalitiesOk returns a tuple with the AttributeCardinalities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeCardinalities

`func (o *GetMetricResponseData) SetAttributeCardinalities(v map[string]float32)`

SetAttributeCardinalities sets AttributeCardinalities field to given value.

### HasAttributeCardinalities

`func (o *GetMetricResponseData) HasAttributeCardinalities() bool`

HasAttributeCardinalities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


