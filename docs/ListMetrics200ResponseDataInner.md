# ListMetrics200ResponseDataInner

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

### NewListMetrics200ResponseDataInner

`func NewListMetrics200ResponseDataInner(name string, type_ string, ) *ListMetrics200ResponseDataInner`

NewListMetrics200ResponseDataInner instantiates a new ListMetrics200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMetrics200ResponseDataInnerWithDefaults

`func NewListMetrics200ResponseDataInnerWithDefaults() *ListMetrics200ResponseDataInner`

NewListMetrics200ResponseDataInnerWithDefaults instantiates a new ListMetrics200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ListMetrics200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListMetrics200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListMetrics200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ListMetrics200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListMetrics200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListMetrics200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *ListMetrics200ResponseDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListMetrics200ResponseDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListMetrics200ResponseDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListMetrics200ResponseDataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUnit

`func (o *ListMetrics200ResponseDataInner) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ListMetrics200ResponseDataInner) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ListMetrics200ResponseDataInner) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *ListMetrics200ResponseDataInner) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetTemporality

`func (o *ListMetrics200ResponseDataInner) GetTemporality() string`

GetTemporality returns the Temporality field if non-nil, zero value otherwise.

### GetTemporalityOk

`func (o *ListMetrics200ResponseDataInner) GetTemporalityOk() (*string, bool)`

GetTemporalityOk returns a tuple with the Temporality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporality

`func (o *ListMetrics200ResponseDataInner) SetTemporality(v string)`

SetTemporality sets Temporality field to given value.

### HasTemporality

`func (o *ListMetrics200ResponseDataInner) HasTemporality() bool`

HasTemporality returns a boolean if a field has been set.

### GetCapabilities

`func (o *ListMetrics200ResponseDataInner) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ListMetrics200ResponseDataInner) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ListMetrics200ResponseDataInner) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *ListMetrics200ResponseDataInner) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetAttributes

`func (o *ListMetrics200ResponseDataInner) GetAttributes() []string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ListMetrics200ResponseDataInner) GetAttributesOk() (*[]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ListMetrics200ResponseDataInner) SetAttributes(v []string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ListMetrics200ResponseDataInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


