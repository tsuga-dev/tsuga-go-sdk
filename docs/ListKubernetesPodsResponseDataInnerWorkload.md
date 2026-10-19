# ListKubernetesPodsResponseDataInnerWorkload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Kind of Kubernetes workload associated with the pod. Derived from workload attributes in pod telemetry. | 
**Name** | **string** | Name of the workload associated with the pod. Derived from workload attributes in pod telemetry. | 

## Methods

### NewListKubernetesPodsResponseDataInnerWorkload

`func NewListKubernetesPodsResponseDataInnerWorkload(kind string, name string, ) *ListKubernetesPodsResponseDataInnerWorkload`

NewListKubernetesPodsResponseDataInnerWorkload instantiates a new ListKubernetesPodsResponseDataInnerWorkload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListKubernetesPodsResponseDataInnerWorkloadWithDefaults

`func NewListKubernetesPodsResponseDataInnerWorkloadWithDefaults() *ListKubernetesPodsResponseDataInnerWorkload`

NewListKubernetesPodsResponseDataInnerWorkloadWithDefaults instantiates a new ListKubernetesPodsResponseDataInnerWorkload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ListKubernetesPodsResponseDataInnerWorkload) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ListKubernetesPodsResponseDataInnerWorkload) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ListKubernetesPodsResponseDataInnerWorkload) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *ListKubernetesPodsResponseDataInnerWorkload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListKubernetesPodsResponseDataInnerWorkload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListKubernetesPodsResponseDataInnerWorkload) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


