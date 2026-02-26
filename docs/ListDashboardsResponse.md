# ListDashboardsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**[]Dashboard**](Dashboard.md) |  | 

## Methods

### NewListDashboardsResponse

`func NewListDashboardsResponse(requestId string, data []Dashboard, ) *ListDashboardsResponse`

NewListDashboardsResponse instantiates a new ListDashboardsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDashboardsResponseWithDefaults

`func NewListDashboardsResponseWithDefaults() *ListDashboardsResponse`

NewListDashboardsResponseWithDefaults instantiates a new ListDashboardsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ListDashboardsResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ListDashboardsResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ListDashboardsResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *ListDashboardsResponse) GetData() []Dashboard`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListDashboardsResponse) GetDataOk() (*[]Dashboard, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListDashboardsResponse) SetData(v []Dashboard)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


