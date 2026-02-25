# GetIngestionApiKey200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**IngestionApiKey**](IngestionApiKey.md) |  | 

## Methods

### NewGetIngestionApiKey200Response

`func NewGetIngestionApiKey200Response(requestId string, data IngestionApiKey, ) *GetIngestionApiKey200Response`

NewGetIngestionApiKey200Response instantiates a new GetIngestionApiKey200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIngestionApiKey200ResponseWithDefaults

`func NewGetIngestionApiKey200ResponseWithDefaults() *GetIngestionApiKey200Response`

NewGetIngestionApiKey200ResponseWithDefaults instantiates a new GetIngestionApiKey200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *GetIngestionApiKey200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetIngestionApiKey200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetIngestionApiKey200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *GetIngestionApiKey200Response) GetData() IngestionApiKey`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetIngestionApiKey200Response) GetDataOk() (*IngestionApiKey, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetIngestionApiKey200Response) SetData(v IngestionApiKey)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


