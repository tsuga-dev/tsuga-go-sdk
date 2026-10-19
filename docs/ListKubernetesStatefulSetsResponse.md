# ListKubernetesStatefulSetsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**[]ListKubernetesStatefulSetsResponseDataInner**](ListKubernetesStatefulSetsResponseDataInner.md) |  | 

## Methods

### NewListKubernetesStatefulSetsResponse

`func NewListKubernetesStatefulSetsResponse(requestId string, data []ListKubernetesStatefulSetsResponseDataInner, ) *ListKubernetesStatefulSetsResponse`

NewListKubernetesStatefulSetsResponse instantiates a new ListKubernetesStatefulSetsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListKubernetesStatefulSetsResponseWithDefaults

`func NewListKubernetesStatefulSetsResponseWithDefaults() *ListKubernetesStatefulSetsResponse`

NewListKubernetesStatefulSetsResponseWithDefaults instantiates a new ListKubernetesStatefulSetsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ListKubernetesStatefulSetsResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ListKubernetesStatefulSetsResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ListKubernetesStatefulSetsResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *ListKubernetesStatefulSetsResponse) GetData() []ListKubernetesStatefulSetsResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListKubernetesStatefulSetsResponse) GetDataOk() (*[]ListKubernetesStatefulSetsResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListKubernetesStatefulSetsResponse) SetData(v []ListKubernetesStatefulSetsResponseDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


