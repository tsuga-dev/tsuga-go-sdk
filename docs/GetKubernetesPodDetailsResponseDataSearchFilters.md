# GetKubernetesPodDetailsResponseDataSearchFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | **string** | TQL log search filter scoped to this pod. Built from cluster, namespace, and pod identifiers. | 
**Traces** | **string** | TQL trace search filter scoped to this pod. Built from cluster, namespace, and pod identifiers. | 

## Methods

### NewGetKubernetesPodDetailsResponseDataSearchFilters

`func NewGetKubernetesPodDetailsResponseDataSearchFilters(logs string, traces string, ) *GetKubernetesPodDetailsResponseDataSearchFilters`

NewGetKubernetesPodDetailsResponseDataSearchFilters instantiates a new GetKubernetesPodDetailsResponseDataSearchFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKubernetesPodDetailsResponseDataSearchFiltersWithDefaults

`func NewGetKubernetesPodDetailsResponseDataSearchFiltersWithDefaults() *GetKubernetesPodDetailsResponseDataSearchFilters`

NewGetKubernetesPodDetailsResponseDataSearchFiltersWithDefaults instantiates a new GetKubernetesPodDetailsResponseDataSearchFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *GetKubernetesPodDetailsResponseDataSearchFilters) GetLogs() string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *GetKubernetesPodDetailsResponseDataSearchFilters) GetLogsOk() (*string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *GetKubernetesPodDetailsResponseDataSearchFilters) SetLogs(v string)`

SetLogs sets Logs field to given value.


### GetTraces

`func (o *GetKubernetesPodDetailsResponseDataSearchFilters) GetTraces() string`

GetTraces returns the Traces field if non-nil, zero value otherwise.

### GetTracesOk

`func (o *GetKubernetesPodDetailsResponseDataSearchFilters) GetTracesOk() (*string, bool)`

GetTracesOk returns a tuple with the Traces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraces

`func (o *GetKubernetesPodDetailsResponseDataSearchFilters) SetTraces(v string)`

SetTraces sets Traces field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


