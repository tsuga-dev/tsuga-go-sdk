# GetTraceLatencySummaryResponseDataServicesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Opaque stable key identifying this contributing service within the response. Tsuga assigns it while building the summary; &#x60;ranges[].contributions[].serviceKey&#x60; and &#x60;serviceTotals[].serviceKey&#x60; reference this value. | 
**Name** | **string** | Observed service name that contributed leaf-span time to the trace. Comes from trace service context and is returned for every contributing service. | 
**Env** | Pointer to **string** | Observed service environment for this contributing service. Comes from trace service context and is absent when the spans do not carry an environment. | [optional] 
**Id** | Pointer to **string** | Service catalog ID for this contributing service. Tsuga sets it when exactly one catalog entry matches the observed service and environment; absent when the service is not cataloged or matches ambiguously. | [optional] 
**Namespace** | Pointer to **string** | Service namespace from the matched service catalog entry. Returned only when Tsuga can resolve the contributing service to one catalog entry with a namespace. | [optional] 

## Methods

### NewGetTraceLatencySummaryResponseDataServicesInner

`func NewGetTraceLatencySummaryResponseDataServicesInner(key string, name string, ) *GetTraceLatencySummaryResponseDataServicesInner`

NewGetTraceLatencySummaryResponseDataServicesInner instantiates a new GetTraceLatencySummaryResponseDataServicesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTraceLatencySummaryResponseDataServicesInnerWithDefaults

`func NewGetTraceLatencySummaryResponseDataServicesInnerWithDefaults() *GetTraceLatencySummaryResponseDataServicesInner`

NewGetTraceLatencySummaryResponseDataServicesInnerWithDefaults instantiates a new GetTraceLatencySummaryResponseDataServicesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *GetTraceLatencySummaryResponseDataServicesInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GetTraceLatencySummaryResponseDataServicesInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GetTraceLatencySummaryResponseDataServicesInner) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *GetTraceLatencySummaryResponseDataServicesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetTraceLatencySummaryResponseDataServicesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetTraceLatencySummaryResponseDataServicesInner) SetName(v string)`

SetName sets Name field to given value.


### GetEnv

`func (o *GetTraceLatencySummaryResponseDataServicesInner) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *GetTraceLatencySummaryResponseDataServicesInner) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *GetTraceLatencySummaryResponseDataServicesInner) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *GetTraceLatencySummaryResponseDataServicesInner) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetId

`func (o *GetTraceLatencySummaryResponseDataServicesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetTraceLatencySummaryResponseDataServicesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetTraceLatencySummaryResponseDataServicesInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetTraceLatencySummaryResponseDataServicesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNamespace

`func (o *GetTraceLatencySummaryResponseDataServicesInner) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *GetTraceLatencySummaryResponseDataServicesInner) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *GetTraceLatencySummaryResponseDataServicesInner) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *GetTraceLatencySummaryResponseDataServicesInner) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


