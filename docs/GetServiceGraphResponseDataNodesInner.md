# GetServiceGraphResponseDataNodesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Service or inferred dependency name shown in the graph. | 
**Env** | Pointer to **string** | Environment observed from spans. | [optional] 
**NodeType** | **string** | Discriminator for this node variant. Always &#x60;inferred&#x60; for a non-database dependency detected from client-span attributes (e.g. a queue or external HTTP call) that could not be matched to a service catalog entry. Database dependencies use the &#x60;inferred_db&#x60; variant. | 
**Id** | **NullableString** | Service ID, or null when multiple namespaces match the observed service name. | 
**Namespace** | **NullableString** | Service namespace, or null when multiple namespaces match the observed service name. | 
**ErrorRate** | Pointer to **float32** | Ratio of entry requests with at least one error to total entry requests observed on this service. Omitted when no entry requests were observed. Computed from the raw sampled counts, so trace sampling does not affect it. | [optional] 
**EntryErrorRequestCount** | **float32** | Estimated number of entry requests on this service where at least one entry span had error status, projected from the sampled traces onto the full population. | 
**EntryRequestCount** | **float32** | Estimated number of entry requests on this service, projected from the sampled traces onto the full population. Observed as unique trace IDs that have at least one entry span on this service. An entry span is a consumer-kind span, a root span (no parent), or any span whose parent belongs to a different service. | 
**Endpoint** | Pointer to **string** | Inferred service endpoint address | [optional] 
**ResourceName** | Pointer to **string** | Inferred resource name, e.g. table or queue name | [optional] 
**ServiceType** | **string** | Inferred dependency type from span metadata: http, rpc, messaging, genai, or object_store. | 

## Methods

### NewGetServiceGraphResponseDataNodesInner

`func NewGetServiceGraphResponseDataNodesInner(name string, nodeType string, id NullableString, namespace NullableString, entryErrorRequestCount float32, entryRequestCount float32, serviceType string, ) *GetServiceGraphResponseDataNodesInner`

NewGetServiceGraphResponseDataNodesInner instantiates a new GetServiceGraphResponseDataNodesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServiceGraphResponseDataNodesInnerWithDefaults

`func NewGetServiceGraphResponseDataNodesInnerWithDefaults() *GetServiceGraphResponseDataNodesInner`

NewGetServiceGraphResponseDataNodesInnerWithDefaults instantiates a new GetServiceGraphResponseDataNodesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetServiceGraphResponseDataNodesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetServiceGraphResponseDataNodesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetServiceGraphResponseDataNodesInner) SetName(v string)`

SetName sets Name field to given value.


### GetEnv

`func (o *GetServiceGraphResponseDataNodesInner) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *GetServiceGraphResponseDataNodesInner) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *GetServiceGraphResponseDataNodesInner) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *GetServiceGraphResponseDataNodesInner) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetNodeType

`func (o *GetServiceGraphResponseDataNodesInner) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *GetServiceGraphResponseDataNodesInner) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *GetServiceGraphResponseDataNodesInner) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.


### GetId

`func (o *GetServiceGraphResponseDataNodesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetServiceGraphResponseDataNodesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetServiceGraphResponseDataNodesInner) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *GetServiceGraphResponseDataNodesInner) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GetServiceGraphResponseDataNodesInner) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetNamespace

`func (o *GetServiceGraphResponseDataNodesInner) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *GetServiceGraphResponseDataNodesInner) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *GetServiceGraphResponseDataNodesInner) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### SetNamespaceNil

`func (o *GetServiceGraphResponseDataNodesInner) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *GetServiceGraphResponseDataNodesInner) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetErrorRate

`func (o *GetServiceGraphResponseDataNodesInner) GetErrorRate() float32`

GetErrorRate returns the ErrorRate field if non-nil, zero value otherwise.

### GetErrorRateOk

`func (o *GetServiceGraphResponseDataNodesInner) GetErrorRateOk() (*float32, bool)`

GetErrorRateOk returns a tuple with the ErrorRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRate

`func (o *GetServiceGraphResponseDataNodesInner) SetErrorRate(v float32)`

SetErrorRate sets ErrorRate field to given value.

### HasErrorRate

`func (o *GetServiceGraphResponseDataNodesInner) HasErrorRate() bool`

HasErrorRate returns a boolean if a field has been set.

### GetEntryErrorRequestCount

`func (o *GetServiceGraphResponseDataNodesInner) GetEntryErrorRequestCount() float32`

GetEntryErrorRequestCount returns the EntryErrorRequestCount field if non-nil, zero value otherwise.

### GetEntryErrorRequestCountOk

`func (o *GetServiceGraphResponseDataNodesInner) GetEntryErrorRequestCountOk() (*float32, bool)`

GetEntryErrorRequestCountOk returns a tuple with the EntryErrorRequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryErrorRequestCount

`func (o *GetServiceGraphResponseDataNodesInner) SetEntryErrorRequestCount(v float32)`

SetEntryErrorRequestCount sets EntryErrorRequestCount field to given value.


### GetEntryRequestCount

`func (o *GetServiceGraphResponseDataNodesInner) GetEntryRequestCount() float32`

GetEntryRequestCount returns the EntryRequestCount field if non-nil, zero value otherwise.

### GetEntryRequestCountOk

`func (o *GetServiceGraphResponseDataNodesInner) GetEntryRequestCountOk() (*float32, bool)`

GetEntryRequestCountOk returns a tuple with the EntryRequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryRequestCount

`func (o *GetServiceGraphResponseDataNodesInner) SetEntryRequestCount(v float32)`

SetEntryRequestCount sets EntryRequestCount field to given value.


### GetEndpoint

`func (o *GetServiceGraphResponseDataNodesInner) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *GetServiceGraphResponseDataNodesInner) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *GetServiceGraphResponseDataNodesInner) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *GetServiceGraphResponseDataNodesInner) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetResourceName

`func (o *GetServiceGraphResponseDataNodesInner) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *GetServiceGraphResponseDataNodesInner) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *GetServiceGraphResponseDataNodesInner) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *GetServiceGraphResponseDataNodesInner) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetServiceType

`func (o *GetServiceGraphResponseDataNodesInner) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *GetServiceGraphResponseDataNodesInner) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *GetServiceGraphResponseDataNodesInner) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


