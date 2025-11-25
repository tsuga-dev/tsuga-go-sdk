# DeleteDashboard200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**DeleteDashboard200ResponseData**](DeleteDashboard200ResponseData.md) |  | 

## Methods

### NewDeleteDashboard200Response

`func NewDeleteDashboard200Response(requestId string, data DeleteDashboard200ResponseData, ) *DeleteDashboard200Response`

NewDeleteDashboard200Response instantiates a new DeleteDashboard200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteDashboard200ResponseWithDefaults

`func NewDeleteDashboard200ResponseWithDefaults() *DeleteDashboard200Response`

NewDeleteDashboard200ResponseWithDefaults instantiates a new DeleteDashboard200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *DeleteDashboard200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DeleteDashboard200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DeleteDashboard200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *DeleteDashboard200Response) GetData() DeleteDashboard200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeleteDashboard200Response) GetDataOk() (*DeleteDashboard200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeleteDashboard200Response) SetData(v DeleteDashboard200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


