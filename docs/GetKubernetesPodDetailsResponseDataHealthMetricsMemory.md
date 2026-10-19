# GetKubernetesPodDetailsResponseDataHealthMetricsMemory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Usage** | **string** | Metric name callers can use to query pod memory usage. This is not the current memory value. | 
**Limit** | **string** | Metric name callers can use to query configured container memory limits for the pod. This is not the current memory limit value. | 
**Request** | **string** | Metric name callers can use to query configured container memory requests for the pod. This is not the current memory request value. | 

## Methods

### NewGetKubernetesPodDetailsResponseDataHealthMetricsMemory

`func NewGetKubernetesPodDetailsResponseDataHealthMetricsMemory(usage string, limit string, request string, ) *GetKubernetesPodDetailsResponseDataHealthMetricsMemory`

NewGetKubernetesPodDetailsResponseDataHealthMetricsMemory instantiates a new GetKubernetesPodDetailsResponseDataHealthMetricsMemory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKubernetesPodDetailsResponseDataHealthMetricsMemoryWithDefaults

`func NewGetKubernetesPodDetailsResponseDataHealthMetricsMemoryWithDefaults() *GetKubernetesPodDetailsResponseDataHealthMetricsMemory`

NewGetKubernetesPodDetailsResponseDataHealthMetricsMemoryWithDefaults instantiates a new GetKubernetesPodDetailsResponseDataHealthMetricsMemory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsage

`func (o *GetKubernetesPodDetailsResponseDataHealthMetricsMemory) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *GetKubernetesPodDetailsResponseDataHealthMetricsMemory) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *GetKubernetesPodDetailsResponseDataHealthMetricsMemory) SetUsage(v string)`

SetUsage sets Usage field to given value.


### GetLimit

`func (o *GetKubernetesPodDetailsResponseDataHealthMetricsMemory) GetLimit() string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetKubernetesPodDetailsResponseDataHealthMetricsMemory) GetLimitOk() (*string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetKubernetesPodDetailsResponseDataHealthMetricsMemory) SetLimit(v string)`

SetLimit sets Limit field to given value.


### GetRequest

`func (o *GetKubernetesPodDetailsResponseDataHealthMetricsMemory) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *GetKubernetesPodDetailsResponseDataHealthMetricsMemory) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *GetKubernetesPodDetailsResponseDataHealthMetricsMemory) SetRequest(v string)`

SetRequest sets Request field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


