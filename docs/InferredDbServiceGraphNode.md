# InferredDbServiceGraphNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Service or inferred dependency name shown in the graph. | 
**Env** | Pointer to **string** | Environment observed from spans. | [optional] 
**Endpoint** | Pointer to **string** | Inferred service endpoint address | [optional] 
**ResourceName** | Pointer to **string** | Inferred resource name, e.g. table or queue name | [optional] 
**NodeType** | **string** | Discriminator for this node variant. Always &#x60;inferred_db&#x60; for a database dependency detected from client-span attributes that could not be matched to a service catalog entry. | 
**ServiceType** | **string** | Inferred dependency type. Always &#x60;db&#x60; for this node variant. | 

## Methods

### NewInferredDbServiceGraphNode

`func NewInferredDbServiceGraphNode(name string, nodeType string, serviceType string, ) *InferredDbServiceGraphNode`

NewInferredDbServiceGraphNode instantiates a new InferredDbServiceGraphNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInferredDbServiceGraphNodeWithDefaults

`func NewInferredDbServiceGraphNodeWithDefaults() *InferredDbServiceGraphNode`

NewInferredDbServiceGraphNodeWithDefaults instantiates a new InferredDbServiceGraphNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InferredDbServiceGraphNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InferredDbServiceGraphNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InferredDbServiceGraphNode) SetName(v string)`

SetName sets Name field to given value.


### GetEnv

`func (o *InferredDbServiceGraphNode) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *InferredDbServiceGraphNode) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *InferredDbServiceGraphNode) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *InferredDbServiceGraphNode) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetEndpoint

`func (o *InferredDbServiceGraphNode) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *InferredDbServiceGraphNode) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *InferredDbServiceGraphNode) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *InferredDbServiceGraphNode) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetResourceName

`func (o *InferredDbServiceGraphNode) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *InferredDbServiceGraphNode) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *InferredDbServiceGraphNode) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *InferredDbServiceGraphNode) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetNodeType

`func (o *InferredDbServiceGraphNode) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *InferredDbServiceGraphNode) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *InferredDbServiceGraphNode) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.


### GetServiceType

`func (o *InferredDbServiceGraphNode) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *InferredDbServiceGraphNode) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *InferredDbServiceGraphNode) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


