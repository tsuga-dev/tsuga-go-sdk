# DeleteIngestionApiKey200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**DeleteIngestionApiKey200ResponseData**](DeleteIngestionApiKey200ResponseData.md) |  | 

## Methods

### NewDeleteIngestionApiKey200Response

`func NewDeleteIngestionApiKey200Response(requestId string, data DeleteIngestionApiKey200ResponseData, ) *DeleteIngestionApiKey200Response`

NewDeleteIngestionApiKey200Response instantiates a new DeleteIngestionApiKey200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteIngestionApiKey200ResponseWithDefaults

`func NewDeleteIngestionApiKey200ResponseWithDefaults() *DeleteIngestionApiKey200Response`

NewDeleteIngestionApiKey200ResponseWithDefaults instantiates a new DeleteIngestionApiKey200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *DeleteIngestionApiKey200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DeleteIngestionApiKey200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DeleteIngestionApiKey200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *DeleteIngestionApiKey200Response) GetData() DeleteIngestionApiKey200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeleteIngestionApiKey200Response) GetDataOk() (*DeleteIngestionApiKey200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeleteIngestionApiKey200Response) SetData(v DeleteIngestionApiKey200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


