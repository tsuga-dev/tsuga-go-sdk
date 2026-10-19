# QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancerType** | Pointer to **string** |  | [optional] 
**DnsName** | Pointer to **string** |  | [optional] 
**VirtualNetworkId** | Pointer to **string** |  | [optional] 
**SubnetIds** | **[]string** |  | 
**Scheme** | Pointer to **string** |  | [optional] 
**Listeners** | Pointer to [**QueryInventoryResourcesResponseDataInnerAllOfAnyOf17AttributesListeners**](QueryInventoryResourcesResponseDataInnerAllOfAnyOf17AttributesListeners.md) |  | [optional] 
**Targets** | Pointer to [**QueryInventoryResourcesResponseDataInnerAllOfAnyOf17AttributesTargets**](QueryInventoryResourcesResponseDataInnerAllOfAnyOf17AttributesTargets.md) |  | [optional] 
**AzureLoadBalancingRules** | Pointer to [**[]QueryInventoryResourcesResponseDataInnerAllOfAnyOf17AttributesAzureLoadBalancingRulesInner**](QueryInventoryResourcesResponseDataInnerAllOfAnyOf17AttributesAzureLoadBalancingRulesInner.md) |  | [optional] 

## Methods

### NewQueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes

`func NewQueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes(subnetIds []string, ) *QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes`

NewQueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes instantiates a new QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryInventoryResourcesResponseDataInnerAllOfAnyOf17AttributesWithDefaults

`func NewQueryInventoryResourcesResponseDataInnerAllOfAnyOf17AttributesWithDefaults() *QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes`

NewQueryInventoryResourcesResponseDataInnerAllOfAnyOf17AttributesWithDefaults instantiates a new QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancerType

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes) GetLoadBalancerType() string`

GetLoadBalancerType returns the LoadBalancerType field if non-nil, zero value otherwise.

### GetLoadBalancerTypeOk

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes) GetLoadBalancerTypeOk() (*string, bool)`

GetLoadBalancerTypeOk returns a tuple with the LoadBalancerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerType

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes) SetLoadBalancerType(v string)`

SetLoadBalancerType sets LoadBalancerType field to given value.

### HasLoadBalancerType

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes) HasLoadBalancerType() bool`

HasLoadBalancerType returns a boolean if a field has been set.

### GetDnsName

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetVirtualNetworkId

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes) GetVirtualNetworkId() string`

GetVirtualNetworkId returns the VirtualNetworkId field if non-nil, zero value otherwise.

### GetVirtualNetworkIdOk

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes) GetVirtualNetworkIdOk() (*string, bool)`

GetVirtualNetworkIdOk returns a tuple with the VirtualNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetworkId

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes) SetVirtualNetworkId(v string)`

SetVirtualNetworkId sets VirtualNetworkId field to given value.

### HasVirtualNetworkId

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes) HasVirtualNetworkId() bool`

HasVirtualNetworkId returns a boolean if a field has been set.

### GetSubnetIds

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes) GetSubnetIds() []string`

GetSubnetIds returns the SubnetIds field if non-nil, zero value otherwise.

### GetSubnetIdsOk

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes) GetSubnetIdsOk() (*[]string, bool)`

GetSubnetIdsOk returns a tuple with the SubnetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetIds

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes) SetSubnetIds(v []string)`

SetSubnetIds sets SubnetIds field to given value.


### GetScheme

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetListeners

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes) GetListeners() QueryInventoryResourcesResponseDataInnerAllOfAnyOf17AttributesListeners`

GetListeners returns the Listeners field if non-nil, zero value otherwise.

### GetListenersOk

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes) GetListenersOk() (*QueryInventoryResourcesResponseDataInnerAllOfAnyOf17AttributesListeners, bool)`

GetListenersOk returns a tuple with the Listeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListeners

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes) SetListeners(v QueryInventoryResourcesResponseDataInnerAllOfAnyOf17AttributesListeners)`

SetListeners sets Listeners field to given value.

### HasListeners

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes) HasListeners() bool`

HasListeners returns a boolean if a field has been set.

### GetTargets

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes) GetTargets() QueryInventoryResourcesResponseDataInnerAllOfAnyOf17AttributesTargets`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes) GetTargetsOk() (*QueryInventoryResourcesResponseDataInnerAllOfAnyOf17AttributesTargets, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes) SetTargets(v QueryInventoryResourcesResponseDataInnerAllOfAnyOf17AttributesTargets)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetAzureLoadBalancingRules

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes) GetAzureLoadBalancingRules() []QueryInventoryResourcesResponseDataInnerAllOfAnyOf17AttributesAzureLoadBalancingRulesInner`

GetAzureLoadBalancingRules returns the AzureLoadBalancingRules field if non-nil, zero value otherwise.

### GetAzureLoadBalancingRulesOk

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes) GetAzureLoadBalancingRulesOk() (*[]QueryInventoryResourcesResponseDataInnerAllOfAnyOf17AttributesAzureLoadBalancingRulesInner, bool)`

GetAzureLoadBalancingRulesOk returns a tuple with the AzureLoadBalancingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureLoadBalancingRules

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes) SetAzureLoadBalancingRules(v []QueryInventoryResourcesResponseDataInnerAllOfAnyOf17AttributesAzureLoadBalancingRulesInner)`

SetAzureLoadBalancingRules sets AzureLoadBalancingRules field to given value.

### HasAzureLoadBalancingRules

`func (o *QueryInventoryResourcesResponseDataInnerAllOfAnyOf17Attributes) HasAzureLoadBalancingRules() bool`

HasAzureLoadBalancingRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


