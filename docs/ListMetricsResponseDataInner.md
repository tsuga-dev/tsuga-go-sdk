# ListMetricsResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the metric | 
**Type** | **string** |  | 
**Description** | Pointer to **string** | Human-readable description of the metric | [optional] 
**Unit** | Pointer to **string** | Unit of the metric | [optional] 
**Temporality** | Pointer to **string** |  | [optional] 
**Capabilities** | Pointer to **[]string** | Supported capabilities, an empty array means all capabilities are supported | [optional] 
**Attributes** | Pointer to **[]string** | Metric attributes | [optional] 

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


