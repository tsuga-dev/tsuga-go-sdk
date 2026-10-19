# GetKubernetesPodDetailsResponseDataHealthMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | [**GetKubernetesPodDetailsResponseDataHealthMetricsCpu**](GetKubernetesPodDetailsResponseDataHealthMetricsCpu.md) |  | 
**Memory** | [**GetKubernetesPodDetailsResponseDataHealthMetricsMemory**](GetKubernetesPodDetailsResponseDataHealthMetricsMemory.md) |  | 

## Methods

### NewGetKubernetesPodDetailsResponseDataHealthMetrics

`func NewGetKubernetesPodDetailsResponseDataHealthMetrics(cpu GetKubernetesPodDetailsResponseDataHealthMetricsCpu, memory GetKubernetesPodDetailsResponseDataHealthMetricsMemory, ) *GetKubernetesPodDetailsResponseDataHealthMetrics`

NewGetKubernetesPodDetailsResponseDataHealthMetrics instantiates a new GetKubernetesPodDetailsResponseDataHealthMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKubernetesPodDetailsResponseDataHealthMetricsWithDefaults

`func NewGetKubernetesPodDetailsResponseDataHealthMetricsWithDefaults() *GetKubernetesPodDetailsResponseDataHealthMetrics`

NewGetKubernetesPodDetailsResponseDataHealthMetricsWithDefaults instantiates a new GetKubernetesPodDetailsResponseDataHealthMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *GetKubernetesPodDetailsResponseDataHealthMetrics) GetCpu() GetKubernetesPodDetailsResponseDataHealthMetricsCpu`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *GetKubernetesPodDetailsResponseDataHealthMetrics) GetCpuOk() (*GetKubernetesPodDetailsResponseDataHealthMetricsCpu, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *GetKubernetesPodDetailsResponseDataHealthMetrics) SetCpu(v GetKubernetesPodDetailsResponseDataHealthMetricsCpu)`

SetCpu sets Cpu field to given value.


### GetMemory

`func (o *GetKubernetesPodDetailsResponseDataHealthMetrics) GetMemory() GetKubernetesPodDetailsResponseDataHealthMetricsMemory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *GetKubernetesPodDetailsResponseDataHealthMetrics) GetMemoryOk() (*GetKubernetesPodDetailsResponseDataHealthMetricsMemory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *GetKubernetesPodDetailsResponseDataHealthMetrics) SetMemory(v GetKubernetesPodDetailsResponseDataHealthMetricsMemory)`

SetMemory sets Memory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


