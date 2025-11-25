# ListDashboards200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**[]ListDashboards200ResponseDataInner**](ListDashboards200ResponseDataInner.md) |  | 

## Methods

### NewListDashboards200Response

`func NewListDashboards200Response(requestId string, data []ListDashboards200ResponseDataInner, ) *ListDashboards200Response`

NewListDashboards200Response instantiates a new ListDashboards200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDashboards200ResponseWithDefaults

`func NewListDashboards200ResponseWithDefaults() *ListDashboards200Response`

NewListDashboards200ResponseWithDefaults instantiates a new ListDashboards200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ListDashboards200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ListDashboards200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ListDashboards200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *ListDashboards200Response) GetData() []ListDashboards200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListDashboards200Response) GetDataOk() (*[]ListDashboards200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListDashboards200Response) SetData(v []ListDashboards200ResponseDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


