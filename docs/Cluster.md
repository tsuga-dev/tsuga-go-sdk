# Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique Tsuga cluster ID. Tsuga assigns it when the cluster record is created. Use this value as &#x60;clusterId&#x60; with public API requests that accept &#x60;clusterId&#x60;. | 
**Name** | **string** | Cluster name stored in Tsuga. Usable as a display label when &#x60;friendlyName&#x60; is not set. | 
**FriendlyName** | Pointer to **string** | Human-readable display name for the cluster. Present when a friendly name is configured. | [optional] 
**ClusterType** | **string** | Cloud provider selected for the Tsuga cluster. | 
**ClusterRegion** | **string** | Cloud region selected for the cluster using the provider region identifier. | 

## Methods

### NewCluster

`func NewCluster(id string, name string, clusterType string, clusterRegion string, ) *Cluster`

NewCluster instantiates a new Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWithDefaults

`func NewClusterWithDefaults() *Cluster`

NewClusterWithDefaults instantiates a new Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Cluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cluster) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Cluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cluster) SetName(v string)`

SetName sets Name field to given value.


### GetFriendlyName

`func (o *Cluster) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *Cluster) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *Cluster) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *Cluster) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### GetClusterType

`func (o *Cluster) GetClusterType() string`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *Cluster) GetClusterTypeOk() (*string, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *Cluster) SetClusterType(v string)`

SetClusterType sets ClusterType field to given value.


### GetClusterRegion

`func (o *Cluster) GetClusterRegion() string`

GetClusterRegion returns the ClusterRegion field if non-nil, zero value otherwise.

### GetClusterRegionOk

`func (o *Cluster) GetClusterRegionOk() (*string, bool)`

GetClusterRegionOk returns a tuple with the ClusterRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterRegion

`func (o *Cluster) SetClusterRegion(v string)`

SetClusterRegion sets ClusterRegion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


