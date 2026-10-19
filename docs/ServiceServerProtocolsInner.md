# ServiceServerProtocolsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | **string** | Observed API protocol, such as &#x60;http&#x60;, &#x60;grpc&#x60;, or &#x60;graphql&#x60;. Tsuga infers it from span metadata. Returned when protocol metadata is available for the service. | 
**LastSeenAt** | **string** | Most recent time this protocol was observed. Tsuga sets it from telemetry timestamps during service metadata refreshes. | 

## Methods

### NewServiceServerProtocolsInner

`func NewServiceServerProtocolsInner(protocol string, lastSeenAt string, ) *ServiceServerProtocolsInner`

NewServiceServerProtocolsInner instantiates a new ServiceServerProtocolsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceServerProtocolsInnerWithDefaults

`func NewServiceServerProtocolsInnerWithDefaults() *ServiceServerProtocolsInner`

NewServiceServerProtocolsInnerWithDefaults instantiates a new ServiceServerProtocolsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *ServiceServerProtocolsInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ServiceServerProtocolsInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ServiceServerProtocolsInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetLastSeenAt

`func (o *ServiceServerProtocolsInner) GetLastSeenAt() string`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *ServiceServerProtocolsInner) GetLastSeenAtOk() (*string, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *ServiceServerProtocolsInner) SetLastSeenAt(v string)`

SetLastSeenAt sets LastSeenAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


