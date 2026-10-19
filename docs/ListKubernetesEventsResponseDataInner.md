# ListKubernetesEventsResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **float32** | Last time the event was observed, in milliseconds since the Unix epoch. | 
**Type** | Pointer to **string** | Event type, &#x60;Normal&#x60; or &#x60;Warning&#x60;. | [optional] 
**Reason** | Pointer to **string** | Short, machine-readable reason for the event, such as &#x60;BackOff&#x60; or &#x60;FailedScheduling&#x60;. | [optional] 
**Message** | Pointer to **string** | Human-readable description of the event. | [optional] 
**ResourceName** | Pointer to **string** | Name of the Kubernetes resource the event is about. | [optional] 
**ResourceKind** | Pointer to **string** | Kind of the Kubernetes resource the event is about, such as &#x60;Pod&#x60; or &#x60;Node&#x60;. | [optional] 
**Namespace** | Pointer to **string** | Kubernetes namespace the event was recorded in. | [optional] 
**Cluster** | Pointer to **string** | Name of the Kubernetes cluster that emitted the event. | [optional] 
**Count** | Pointer to **int32** | Number of times the event occurred. | [optional] 

## Methods

### NewListKubernetesEventsResponseDataInner

`func NewListKubernetesEventsResponseDataInner(timestamp float32, ) *ListKubernetesEventsResponseDataInner`

NewListKubernetesEventsResponseDataInner instantiates a new ListKubernetesEventsResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListKubernetesEventsResponseDataInnerWithDefaults

`func NewListKubernetesEventsResponseDataInnerWithDefaults() *ListKubernetesEventsResponseDataInner`

NewListKubernetesEventsResponseDataInnerWithDefaults instantiates a new ListKubernetesEventsResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *ListKubernetesEventsResponseDataInner) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ListKubernetesEventsResponseDataInner) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ListKubernetesEventsResponseDataInner) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.


### GetType

`func (o *ListKubernetesEventsResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListKubernetesEventsResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListKubernetesEventsResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListKubernetesEventsResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetReason

`func (o *ListKubernetesEventsResponseDataInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ListKubernetesEventsResponseDataInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ListKubernetesEventsResponseDataInner) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ListKubernetesEventsResponseDataInner) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetMessage

`func (o *ListKubernetesEventsResponseDataInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListKubernetesEventsResponseDataInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListKubernetesEventsResponseDataInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ListKubernetesEventsResponseDataInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResourceName

`func (o *ListKubernetesEventsResponseDataInner) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ListKubernetesEventsResponseDataInner) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *ListKubernetesEventsResponseDataInner) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *ListKubernetesEventsResponseDataInner) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetResourceKind

`func (o *ListKubernetesEventsResponseDataInner) GetResourceKind() string`

GetResourceKind returns the ResourceKind field if non-nil, zero value otherwise.

### GetResourceKindOk

`func (o *ListKubernetesEventsResponseDataInner) GetResourceKindOk() (*string, bool)`

GetResourceKindOk returns a tuple with the ResourceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKind

`func (o *ListKubernetesEventsResponseDataInner) SetResourceKind(v string)`

SetResourceKind sets ResourceKind field to given value.

### HasResourceKind

`func (o *ListKubernetesEventsResponseDataInner) HasResourceKind() bool`

HasResourceKind returns a boolean if a field has been set.

### GetNamespace

`func (o *ListKubernetesEventsResponseDataInner) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ListKubernetesEventsResponseDataInner) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ListKubernetesEventsResponseDataInner) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ListKubernetesEventsResponseDataInner) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetCluster

`func (o *ListKubernetesEventsResponseDataInner) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ListKubernetesEventsResponseDataInner) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ListKubernetesEventsResponseDataInner) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ListKubernetesEventsResponseDataInner) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCount

`func (o *ListKubernetesEventsResponseDataInner) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListKubernetesEventsResponseDataInner) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListKubernetesEventsResponseDataInner) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListKubernetesEventsResponseDataInner) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


