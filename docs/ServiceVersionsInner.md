# ServiceVersionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | Service version value observed in telemetry. The emitting service sets it through &#x60;context.service.version&#x60;. Returned when Tsuga has observed a version for the service. | 
**FirstSeenAt** | **string** | First time this version was observed for the service. Tsuga sets it from telemetry timestamps and refreshes observed versions on each service catalog pass. | 
**LastSeenAt** | **string** | Most recent time this version was observed for the service. Tsuga sets it from telemetry timestamps and refreshes observed versions on each service catalog pass. | 
**Faulty** | Pointer to **bool** | Whether Tsuga detected this version as associated with faulty behavior. | [optional] 
**FaultyLatency** | Pointer to **bool** | Whether the faulty behavior Tsuga detected for this version is a latency degradation. Only set by faulty deployment detection when it reports on latency; absent when the detection did not evaluate latency. | [optional] 
**DetectionTriggered** | Pointer to **bool** | Whether faulty deployment detection has already been triggered for this version. Set when the detection job is launched to avoid re-triggering it. Absent on versions predating this field. | [optional] 

## Methods

### NewServiceVersionsInner

`func NewServiceVersionsInner(version string, firstSeenAt string, lastSeenAt string, ) *ServiceVersionsInner`

NewServiceVersionsInner instantiates a new ServiceVersionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceVersionsInnerWithDefaults

`func NewServiceVersionsInnerWithDefaults() *ServiceVersionsInner`

NewServiceVersionsInnerWithDefaults instantiates a new ServiceVersionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ServiceVersionsInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ServiceVersionsInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ServiceVersionsInner) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetFirstSeenAt

`func (o *ServiceVersionsInner) GetFirstSeenAt() string`

GetFirstSeenAt returns the FirstSeenAt field if non-nil, zero value otherwise.

### GetFirstSeenAtOk

`func (o *ServiceVersionsInner) GetFirstSeenAtOk() (*string, bool)`

GetFirstSeenAtOk returns a tuple with the FirstSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenAt

`func (o *ServiceVersionsInner) SetFirstSeenAt(v string)`

SetFirstSeenAt sets FirstSeenAt field to given value.


### GetLastSeenAt

`func (o *ServiceVersionsInner) GetLastSeenAt() string`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *ServiceVersionsInner) GetLastSeenAtOk() (*string, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *ServiceVersionsInner) SetLastSeenAt(v string)`

SetLastSeenAt sets LastSeenAt field to given value.


### GetFaulty

`func (o *ServiceVersionsInner) GetFaulty() bool`

GetFaulty returns the Faulty field if non-nil, zero value otherwise.

### GetFaultyOk

`func (o *ServiceVersionsInner) GetFaultyOk() (*bool, bool)`

GetFaultyOk returns a tuple with the Faulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaulty

`func (o *ServiceVersionsInner) SetFaulty(v bool)`

SetFaulty sets Faulty field to given value.

### HasFaulty

`func (o *ServiceVersionsInner) HasFaulty() bool`

HasFaulty returns a boolean if a field has been set.

### GetFaultyLatency

`func (o *ServiceVersionsInner) GetFaultyLatency() bool`

GetFaultyLatency returns the FaultyLatency field if non-nil, zero value otherwise.

### GetFaultyLatencyOk

`func (o *ServiceVersionsInner) GetFaultyLatencyOk() (*bool, bool)`

GetFaultyLatencyOk returns a tuple with the FaultyLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultyLatency

`func (o *ServiceVersionsInner) SetFaultyLatency(v bool)`

SetFaultyLatency sets FaultyLatency field to given value.

### HasFaultyLatency

`func (o *ServiceVersionsInner) HasFaultyLatency() bool`

HasFaultyLatency returns a boolean if a field has been set.

### GetDetectionTriggered

`func (o *ServiceVersionsInner) GetDetectionTriggered() bool`

GetDetectionTriggered returns the DetectionTriggered field if non-nil, zero value otherwise.

### GetDetectionTriggeredOk

`func (o *ServiceVersionsInner) GetDetectionTriggeredOk() (*bool, bool)`

GetDetectionTriggeredOk returns a tuple with the DetectionTriggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectionTriggered

`func (o *ServiceVersionsInner) SetDetectionTriggered(v bool)`

SetDetectionTriggered sets DetectionTriggered field to given value.

### HasDetectionTriggered

`func (o *ServiceVersionsInner) HasDetectionTriggered() bool`

HasDetectionTriggered returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


