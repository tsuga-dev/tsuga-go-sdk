# GetKubernetesPodDetailsResponseDataHealthMetricsCpu

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Usage** | **string** | Metric name callers can use to query pod CPU usage. This is not the current CPU value. | 
**Limit** | **string** | Metric name callers can use to query configured container CPU limits for the pod. This is not the current CPU limit value. | 
**Request** | **string** | Metric name callers can use to query configured container CPU requests for the pod. This is not the current CPU request value. | 

## Methods

### NewGetKubernetesPodDetailsResponseDataHealthMetricsCpu

`func NewGetKubernetesPodDetailsResponseDataHealthMetricsCpu(usage string, limit string, request string, ) *GetKubernetesPodDetailsResponseDataHealthMetricsCpu`

NewGetKubernetesPodDetailsResponseDataHealthMetricsCpu instantiates a new GetKubernetesPodDetailsResponseDataHealthMetricsCpu object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKubernetesPodDetailsResponseDataHealthMetricsCpuWithDefaults

`func NewGetKubernetesPodDetailsResponseDataHealthMetricsCpuWithDefaults() *GetKubernetesPodDetailsResponseDataHealthMetricsCpu`

NewGetKubernetesPodDetailsResponseDataHealthMetricsCpuWithDefaults instantiates a new GetKubernetesPodDetailsResponseDataHealthMetricsCpu object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsage

`func (o *GetKubernetesPodDetailsResponseDataHealthMetricsCpu) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *GetKubernetesPodDetailsResponseDataHealthMetricsCpu) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *GetKubernetesPodDetailsResponseDataHealthMetricsCpu) SetUsage(v string)`

SetUsage sets Usage field to given value.


### GetLimit

`func (o *GetKubernetesPodDetailsResponseDataHealthMetricsCpu) GetLimit() string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetKubernetesPodDetailsResponseDataHealthMetricsCpu) GetLimitOk() (*string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetKubernetesPodDetailsResponseDataHealthMetricsCpu) SetLimit(v string)`

SetLimit sets Limit field to given value.


### GetRequest

`func (o *GetKubernetesPodDetailsResponseDataHealthMetricsCpu) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *GetKubernetesPodDetailsResponseDataHealthMetricsCpu) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *GetKubernetesPodDetailsResponseDataHealthMetricsCpu) SetRequest(v string)`

SetRequest sets Request field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


