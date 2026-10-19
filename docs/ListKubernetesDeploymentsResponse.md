# ListKubernetesDeploymentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**[]ListKubernetesDeploymentsResponseDataInner**](ListKubernetesDeploymentsResponseDataInner.md) |  | 

## Methods

### NewListKubernetesDeploymentsResponse

`func NewListKubernetesDeploymentsResponse(requestId string, data []ListKubernetesDeploymentsResponseDataInner, ) *ListKubernetesDeploymentsResponse`

NewListKubernetesDeploymentsResponse instantiates a new ListKubernetesDeploymentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListKubernetesDeploymentsResponseWithDefaults

`func NewListKubernetesDeploymentsResponseWithDefaults() *ListKubernetesDeploymentsResponse`

NewListKubernetesDeploymentsResponseWithDefaults instantiates a new ListKubernetesDeploymentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ListKubernetesDeploymentsResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ListKubernetesDeploymentsResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ListKubernetesDeploymentsResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *ListKubernetesDeploymentsResponse) GetData() []ListKubernetesDeploymentsResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListKubernetesDeploymentsResponse) GetDataOk() (*[]ListKubernetesDeploymentsResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListKubernetesDeploymentsResponse) SetData(v []ListKubernetesDeploymentsResponseDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


