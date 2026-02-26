# GetService200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**Service**](Service.md) |  | 

## Methods

### NewGetService200Response

`func NewGetService200Response(requestId string, data Service, ) *GetService200Response`

NewGetService200Response instantiates a new GetService200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetService200ResponseWithDefaults

`func NewGetService200ResponseWithDefaults() *GetService200Response`

NewGetService200ResponseWithDefaults instantiates a new GetService200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *GetService200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetService200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetService200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *GetService200Response) GetData() Service`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetService200Response) GetDataOk() (*Service, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetService200Response) SetData(v Service)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


