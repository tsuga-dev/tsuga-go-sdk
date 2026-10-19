# ServiceCloudPlatformsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | **string** | Observed cloud or orchestration platform. Tsuga reads the &#x60;cloud.platform&#x60; resource attribute, such as &#x60;aws_eks&#x60; or &#x60;aws_lambda&#x60;, and adds &#x60;kubernetes&#x60; when spans carry &#x60;k8s.cluster.name&#x60; and &#x60;aws_ecs&#x60; when they carry &#x60;aws.ecs.task.arn&#x60;. The emitting service or collector sets these resource attributes. | 
**LastSeenAt** | **string** | Most recent time this platform was observed. Tsuga sets it from telemetry timestamps during service metadata refreshes. | 

## Methods

### NewServiceCloudPlatformsInner

`func NewServiceCloudPlatformsInner(platform string, lastSeenAt string, ) *ServiceCloudPlatformsInner`

NewServiceCloudPlatformsInner instantiates a new ServiceCloudPlatformsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceCloudPlatformsInnerWithDefaults

`func NewServiceCloudPlatformsInnerWithDefaults() *ServiceCloudPlatformsInner`

NewServiceCloudPlatformsInnerWithDefaults instantiates a new ServiceCloudPlatformsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *ServiceCloudPlatformsInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ServiceCloudPlatformsInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ServiceCloudPlatformsInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetLastSeenAt

`func (o *ServiceCloudPlatformsInner) GetLastSeenAt() string`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *ServiceCloudPlatformsInner) GetLastSeenAtOk() (*string, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *ServiceCloudPlatformsInner) SetLastSeenAt(v string)`

SetLastSeenAt sets LastSeenAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


