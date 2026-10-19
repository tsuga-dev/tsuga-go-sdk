# ObservedServiceGraphNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Service or inferred dependency name shown in the graph. | 
**Env** | Pointer to **string** | Environment observed from spans. | [optional] 
**NodeType** | **string** | Discriminator for this node variant. Always &#x60;observed&#x60; for a graph node matched to a service catalog entry by name and env. | 
**Id** | **NullableString** | Service ID, or null when multiple namespaces match the observed service name. | 
**Namespace** | **NullableString** | Service namespace, or null when multiple namespaces match the observed service name. | 
**ErrorRate** | Pointer to **float32** | Ratio of entry requests with at least one error to total entry requests observed on this service. Omitted when no entry requests were observed. Computed from the raw sampled counts, so trace sampling does not affect it. | [optional] 
**EntryErrorRequestCount** | **float32** | Estimated number of entry requests on this service where at least one entry span had error status, projected from the sampled traces onto the full population. | 
**EntryRequestCount** | **float32** | Estimated number of entry requests on this service, projected from the sampled traces onto the full population. Observed as unique trace IDs that have at least one entry span on this service. An entry span is a consumer-kind span, a root span (no parent), or any span whose parent belongs to a different service. | 

## Methods

### NewObservedServiceGraphNode

`func NewObservedServiceGraphNode(name string, nodeType string, id NullableString, namespace NullableString, entryErrorRequestCount float32, entryRequestCount float32, ) *ObservedServiceGraphNode`

NewObservedServiceGraphNode instantiates a new ObservedServiceGraphNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObservedServiceGraphNodeWithDefaults

`func NewObservedServiceGraphNodeWithDefaults() *ObservedServiceGraphNode`

NewObservedServiceGraphNodeWithDefaults instantiates a new ObservedServiceGraphNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ObservedServiceGraphNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObservedServiceGraphNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObservedServiceGraphNode) SetName(v string)`

SetName sets Name field to given value.


### GetEnv

`func (o *ObservedServiceGraphNode) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *ObservedServiceGraphNode) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *ObservedServiceGraphNode) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *ObservedServiceGraphNode) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetNodeType

`func (o *ObservedServiceGraphNode) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *ObservedServiceGraphNode) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *ObservedServiceGraphNode) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.


### GetId

`func (o *ObservedServiceGraphNode) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObservedServiceGraphNode) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObservedServiceGraphNode) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *ObservedServiceGraphNode) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ObservedServiceGraphNode) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetNamespace

`func (o *ObservedServiceGraphNode) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ObservedServiceGraphNode) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ObservedServiceGraphNode) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### SetNamespaceNil

`func (o *ObservedServiceGraphNode) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *ObservedServiceGraphNode) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetErrorRate

`func (o *ObservedServiceGraphNode) GetErrorRate() float32`

GetErrorRate returns the ErrorRate field if non-nil, zero value otherwise.

### GetErrorRateOk

`func (o *ObservedServiceGraphNode) GetErrorRateOk() (*float32, bool)`

GetErrorRateOk returns a tuple with the ErrorRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRate

`func (o *ObservedServiceGraphNode) SetErrorRate(v float32)`

SetErrorRate sets ErrorRate field to given value.

### HasErrorRate

`func (o *ObservedServiceGraphNode) HasErrorRate() bool`

HasErrorRate returns a boolean if a field has been set.

### GetEntryErrorRequestCount

`func (o *ObservedServiceGraphNode) GetEntryErrorRequestCount() float32`

GetEntryErrorRequestCount returns the EntryErrorRequestCount field if non-nil, zero value otherwise.

### GetEntryErrorRequestCountOk

`func (o *ObservedServiceGraphNode) GetEntryErrorRequestCountOk() (*float32, bool)`

GetEntryErrorRequestCountOk returns a tuple with the EntryErrorRequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryErrorRequestCount

`func (o *ObservedServiceGraphNode) SetEntryErrorRequestCount(v float32)`

SetEntryErrorRequestCount sets EntryErrorRequestCount field to given value.


### GetEntryRequestCount

`func (o *ObservedServiceGraphNode) GetEntryRequestCount() float32`

GetEntryRequestCount returns the EntryRequestCount field if non-nil, zero value otherwise.

### GetEntryRequestCountOk

`func (o *ObservedServiceGraphNode) GetEntryRequestCountOk() (*float32, bool)`

GetEntryRequestCountOk returns a tuple with the EntryRequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryRequestCount

`func (o *ObservedServiceGraphNode) SetEntryRequestCount(v float32)`

SetEntryRequestCount sets EntryRequestCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


