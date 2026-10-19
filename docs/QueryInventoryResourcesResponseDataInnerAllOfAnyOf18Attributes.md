# QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **float32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**TargetGroupArn** | Pointer to **string** |  | [optional] 
**VirtualNetworkId** | Pointer to **string** |  | [optional] 
**HealthCheckProtocol** | Pointer to **string** |  | [optional] 
**HealthCheckPort** | Pointer to **string** |  | [optional] 
**HealthCheckPath** | Pointer to **string** |  | [optional] 
**HealthCheckIntervalSeconds** | Pointer to **float32** |  | [optional] 
**HealthCheckTimeoutSeconds** | Pointer to **float32** |  | [optional] 
**HealthyThresholdCount** | Pointer to **float32** |  | [optional] 
**UnhealthyThresholdCount** | Pointer to **float32** |  | [optional] 
**Targets** | Pointer to [**[]QueryInventoryResourcesResponseDataInnerAllOfAnyOf18AttributesTargetsInner**](QueryInventoryResourcesResponseDataInnerAllOfAnyOf18AttributesTargetsInner.md) |  | [optional] 
**TargetCount** | Pointer to **float32** |  | [optional] 
**HealthyTargetCount** | Pointer to **float32** |  | [optional] 
**LoadBalancerArns** | Pointer to **[]string** |  | [optional] 

## Methods

### NewQueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes

`func NewQueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes() *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes`

NewQueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes instantiates a new QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryInventoryResourcesResponseDataInnerAllOfAnyOf18AttributesWithDefaults

`func NewQueryInventoryResourcesResponseDataInnerAllOfAnyOf18AttributesWithDefaults() *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes`

NewQueryInventoryResourcesResponseDataInnerAllOfAnyOf18AttributesWithDefaults instantiates a new QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPort

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetPort() float32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetPortOk() (*float32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) SetPort(v float32)`

SetPort sets Port field to given value.

### HasPort

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetType

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTargetGroupArn

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetTargetGroupArn() string`

GetTargetGroupArn returns the TargetGroupArn field if non-nil, zero value otherwise.

### GetTargetGroupArnOk

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetTargetGroupArnOk() (*string, bool)`

GetTargetGroupArnOk returns a tuple with the TargetGroupArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupArn

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) SetTargetGroupArn(v string)`

SetTargetGroupArn sets TargetGroupArn field to given value.

### HasTargetGroupArn

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) HasTargetGroupArn() bool`

HasTargetGroupArn returns a boolean if a field has been set.

### GetVirtualNetworkId

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetVirtualNetworkId() string`

GetVirtualNetworkId returns the VirtualNetworkId field if non-nil, zero value otherwise.

### GetVirtualNetworkIdOk

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetVirtualNetworkIdOk() (*string, bool)`

GetVirtualNetworkIdOk returns a tuple with the VirtualNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetworkId

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) SetVirtualNetworkId(v string)`

SetVirtualNetworkId sets VirtualNetworkId field to given value.

### HasVirtualNetworkId

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) HasVirtualNetworkId() bool`

HasVirtualNetworkId returns a boolean if a field has been set.

### GetHealthCheckProtocol

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetHealthCheckProtocol() string`

GetHealthCheckProtocol returns the HealthCheckProtocol field if non-nil, zero value otherwise.

### GetHealthCheckProtocolOk

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetHealthCheckProtocolOk() (*string, bool)`

GetHealthCheckProtocolOk returns a tuple with the HealthCheckProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckProtocol

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) SetHealthCheckProtocol(v string)`

SetHealthCheckProtocol sets HealthCheckProtocol field to given value.

### HasHealthCheckProtocol

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) HasHealthCheckProtocol() bool`

HasHealthCheckProtocol returns a boolean if a field has been set.

### GetHealthCheckPort

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetHealthCheckPort() string`

GetHealthCheckPort returns the HealthCheckPort field if non-nil, zero value otherwise.

### GetHealthCheckPortOk

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetHealthCheckPortOk() (*string, bool)`

GetHealthCheckPortOk returns a tuple with the HealthCheckPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckPort

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) SetHealthCheckPort(v string)`

SetHealthCheckPort sets HealthCheckPort field to given value.

### HasHealthCheckPort

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) HasHealthCheckPort() bool`

HasHealthCheckPort returns a boolean if a field has been set.

### GetHealthCheckPath

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetHealthCheckPath() string`

GetHealthCheckPath returns the HealthCheckPath field if non-nil, zero value otherwise.

### GetHealthCheckPathOk

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetHealthCheckPathOk() (*string, bool)`

GetHealthCheckPathOk returns a tuple with the HealthCheckPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckPath

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) SetHealthCheckPath(v string)`

SetHealthCheckPath sets HealthCheckPath field to given value.

### HasHealthCheckPath

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) HasHealthCheckPath() bool`

HasHealthCheckPath returns a boolean if a field has been set.

### GetHealthCheckIntervalSeconds

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetHealthCheckIntervalSeconds() float32`

GetHealthCheckIntervalSeconds returns the HealthCheckIntervalSeconds field if non-nil, zero value otherwise.

### GetHealthCheckIntervalSecondsOk

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetHealthCheckIntervalSecondsOk() (*float32, bool)`

GetHealthCheckIntervalSecondsOk returns a tuple with the HealthCheckIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckIntervalSeconds

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) SetHealthCheckIntervalSeconds(v float32)`

SetHealthCheckIntervalSeconds sets HealthCheckIntervalSeconds field to given value.

### HasHealthCheckIntervalSeconds

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) HasHealthCheckIntervalSeconds() bool`

HasHealthCheckIntervalSeconds returns a boolean if a field has been set.

### GetHealthCheckTimeoutSeconds

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetHealthCheckTimeoutSeconds() float32`

GetHealthCheckTimeoutSeconds returns the HealthCheckTimeoutSeconds field if non-nil, zero value otherwise.

### GetHealthCheckTimeoutSecondsOk

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetHealthCheckTimeoutSecondsOk() (*float32, bool)`

GetHealthCheckTimeoutSecondsOk returns a tuple with the HealthCheckTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckTimeoutSeconds

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) SetHealthCheckTimeoutSeconds(v float32)`

SetHealthCheckTimeoutSeconds sets HealthCheckTimeoutSeconds field to given value.

### HasHealthCheckTimeoutSeconds

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) HasHealthCheckTimeoutSeconds() bool`

HasHealthCheckTimeoutSeconds returns a boolean if a field has been set.

### GetHealthyThresholdCount

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetHealthyThresholdCount() float32`

GetHealthyThresholdCount returns the HealthyThresholdCount field if non-nil, zero value otherwise.

### GetHealthyThresholdCountOk

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetHealthyThresholdCountOk() (*float32, bool)`

GetHealthyThresholdCountOk returns a tuple with the HealthyThresholdCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthyThresholdCount

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) SetHealthyThresholdCount(v float32)`

SetHealthyThresholdCount sets HealthyThresholdCount field to given value.

### HasHealthyThresholdCount

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) HasHealthyThresholdCount() bool`

HasHealthyThresholdCount returns a boolean if a field has been set.

### GetUnhealthyThresholdCount

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetUnhealthyThresholdCount() float32`

GetUnhealthyThresholdCount returns the UnhealthyThresholdCount field if non-nil, zero value otherwise.

### GetUnhealthyThresholdCountOk

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetUnhealthyThresholdCountOk() (*float32, bool)`

GetUnhealthyThresholdCountOk returns a tuple with the UnhealthyThresholdCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnhealthyThresholdCount

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) SetUnhealthyThresholdCount(v float32)`

SetUnhealthyThresholdCount sets UnhealthyThresholdCount field to given value.

### HasUnhealthyThresholdCount

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) HasUnhealthyThresholdCount() bool`

HasUnhealthyThresholdCount returns a boolean if a field has been set.

### GetTargets

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetTargets() []QueryInventoryResourcesResponseDataInnerAllOfAnyOf18AttributesTargetsInner`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetTargetsOk() (*[]QueryInventoryResourcesResponseDataInnerAllOfAnyOf18AttributesTargetsInner, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) SetTargets(v []QueryInventoryResourcesResponseDataInnerAllOfAnyOf18AttributesTargetsInner)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetTargetCount

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetTargetCount() float32`

GetTargetCount returns the TargetCount field if non-nil, zero value otherwise.

### GetTargetCountOk

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetTargetCountOk() (*float32, bool)`

GetTargetCountOk returns a tuple with the TargetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCount

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) SetTargetCount(v float32)`

SetTargetCount sets TargetCount field to given value.

### HasTargetCount

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) HasTargetCount() bool`

HasTargetCount returns a boolean if a field has been set.

### GetHealthyTargetCount

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetHealthyTargetCount() float32`

GetHealthyTargetCount returns the HealthyTargetCount field if non-nil, zero value otherwise.

### GetHealthyTargetCountOk

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetHealthyTargetCountOk() (*float32, bool)`

GetHealthyTargetCountOk returns a tuple with the HealthyTargetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthyTargetCount

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) SetHealthyTargetCount(v float32)`

SetHealthyTargetCount sets HealthyTargetCount field to given value.

### HasHealthyTargetCount

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) HasHealthyTargetCount() bool`

HasHealthyTargetCount returns a boolean if a field has been set.

### GetLoadBalancerArns

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetLoadBalancerArns() []string`

GetLoadBalancerArns returns the LoadBalancerArns field if non-nil, zero value otherwise.

### GetLoadBalancerArnsOk

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) GetLoadBalancerArnsOk() (*[]string, bool)`

GetLoadBalancerArnsOk returns a tuple with the LoadBalancerArns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerArns

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) SetLoadBalancerArns(v []string)`

SetLoadBalancerArns sets LoadBalancerArns field to given value.

### HasLoadBalancerArns

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf18Attributes) HasLoadBalancerArns() bool`

HasLoadBalancerArns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


