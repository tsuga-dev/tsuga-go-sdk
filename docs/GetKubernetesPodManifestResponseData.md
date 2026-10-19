# GetKubernetesPodManifestResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to [**GetKubernetesPodManifestResponseDataConfiguration**](GetKubernetesPodManifestResponseDataConfiguration.md) |  | [optional] 
**Labels** | **map[string]string** | Kubernetes labels on the pod from the latest observed pod manifest. Key-value strings. | 
**Annotations** | **map[string]string** | Kubernetes annotations on the pod from the latest observed pod manifest. Key-value strings. | 
**DeletedAt** | Pointer to **string** | Timestamp when Tsuga last observed the pod deletion event. Present when the latest observed pod event is deleted. | [optional] 

## Methods

### NewGetKubernetesPodManifestResponseData

`func NewGetKubernetesPodManifestResponseData(labels map[string]string, annotations map[string]string, ) *GetKubernetesPodManifestResponseData`

NewGetKubernetesPodManifestResponseData instantiates a new GetKubernetesPodManifestResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKubernetesPodManifestResponseDataWithDefaults

`func NewGetKubernetesPodManifestResponseDataWithDefaults() *GetKubernetesPodManifestResponseData`

NewGetKubernetesPodManifestResponseDataWithDefaults instantiates a new GetKubernetesPodManifestResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *GetKubernetesPodManifestResponseData) GetConfiguration() GetKubernetesPodManifestResponseDataConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GetKubernetesPodManifestResponseData) GetConfigurationOk() (*GetKubernetesPodManifestResponseDataConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GetKubernetesPodManifestResponseData) SetConfiguration(v GetKubernetesPodManifestResponseDataConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *GetKubernetesPodManifestResponseData) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetLabels

`func (o *GetKubernetesPodManifestResponseData) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetKubernetesPodManifestResponseData) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetKubernetesPodManifestResponseData) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.


### GetAnnotations

`func (o *GetKubernetesPodManifestResponseData) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *GetKubernetesPodManifestResponseData) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *GetKubernetesPodManifestResponseData) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.


### GetDeletedAt

`func (o *GetKubernetesPodManifestResponseData) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *GetKubernetesPodManifestResponseData) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *GetKubernetesPodManifestResponseData) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *GetKubernetesPodManifestResponseData) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


