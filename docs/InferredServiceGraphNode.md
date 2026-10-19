# InferredServiceGraphNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Service or inferred dependency name shown in the graph. | 
**Env** | Pointer to **string** | Environment observed from spans. | [optional] 
**Endpoint** | Pointer to **string** | Inferred service endpoint address | [optional] 
**ResourceName** | Pointer to **string** | Inferred resource name, e.g. table or queue name | [optional] 
**NodeType** | **string** | Discriminator for this node variant. Always &#x60;inferred&#x60; for a non-database dependency detected from client-span attributes (e.g. a queue or external HTTP call) that could not be matched to a service catalog entry. Database dependencies use the &#x60;inferred_db&#x60; variant. | 
**ServiceType** | **string** | Inferred dependency type from span metadata: http, rpc, messaging, genai, or object_store. | 

## Methods

### NewInferredServiceGraphNode

`func NewInferredServiceGraphNode(name string, nodeType string, serviceType string, ) *InferredServiceGraphNode`

NewInferredServiceGraphNode instantiates a new InferredServiceGraphNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInferredServiceGraphNodeWithDefaults

`func NewInferredServiceGraphNodeWithDefaults() *InferredServiceGraphNode`

NewInferredServiceGraphNodeWithDefaults instantiates a new InferredServiceGraphNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InferredServiceGraphNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InferredServiceGraphNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InferredServiceGraphNode) SetName(v string)`

SetName sets Name field to given value.


### GetEnv

`func (o *InferredServiceGraphNode) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *InferredServiceGraphNode) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *InferredServiceGraphNode) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *InferredServiceGraphNode) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetEndpoint

`func (o *InferredServiceGraphNode) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *InferredServiceGraphNode) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *InferredServiceGraphNode) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *InferredServiceGraphNode) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetResourceName

`func (o *InferredServiceGraphNode) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *InferredServiceGraphNode) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *InferredServiceGraphNode) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *InferredServiceGraphNode) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetNodeType

`func (o *InferredServiceGraphNode) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *InferredServiceGraphNode) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *InferredServiceGraphNode) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.


### GetServiceType

`func (o *InferredServiceGraphNode) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *InferredServiceGraphNode) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *InferredServiceGraphNode) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


