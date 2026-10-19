# GetKubernetesPodManifestResponseDataConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | **map[string]interface{}** | Kubernetes PodSpec from the latest observed pod manifest. Fields outside &#x60;object.spec&#x60; are stripped. | 
**LastObservedAt** | **string** | Timestamp when Tsuga last observed the pod manifest used for configuration. | 

## Methods

### NewGetKubernetesPodManifestResponseDataConfiguration

`func NewGetKubernetesPodManifestResponseDataConfiguration(spec map[string]*interface{}, lastObservedAt string, ) *GetKubernetesPodManifestResponseDataConfiguration`

NewGetKubernetesPodManifestResponseDataConfiguration instantiates a new GetKubernetesPodManifestResponseDataConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKubernetesPodManifestResponseDataConfigurationWithDefaults

`func NewGetKubernetesPodManifestResponseDataConfigurationWithDefaults() *GetKubernetesPodManifestResponseDataConfiguration`

NewGetKubernetesPodManifestResponseDataConfigurationWithDefaults instantiates a new GetKubernetesPodManifestResponseDataConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *GetKubernetesPodManifestResponseDataConfiguration) GetSpec() map[string]*interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *GetKubernetesPodManifestResponseDataConfiguration) GetSpecOk() (*map[string]*interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *GetKubernetesPodManifestResponseDataConfiguration) SetSpec(v map[string]*interface{})`

SetSpec sets Spec field to given value.


### GetLastObservedAt

`func (o *GetKubernetesPodManifestResponseDataConfiguration) GetLastObservedAt() string`

GetLastObservedAt returns the LastObservedAt field if non-nil, zero value otherwise.

### GetLastObservedAtOk

`func (o *GetKubernetesPodManifestResponseDataConfiguration) GetLastObservedAtOk() (*string, bool)`

GetLastObservedAtOk returns a tuple with the LastObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastObservedAt

`func (o *GetKubernetesPodManifestResponseDataConfiguration) SetLastObservedAt(v string)`

SetLastObservedAt sets LastObservedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


