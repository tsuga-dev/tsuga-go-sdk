# UpdateIngestionApiKey200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**IngestionApiKey**](IngestionApiKey.md) |  | 

## Methods

### NewUpdateIngestionApiKey200Response

`func NewUpdateIngestionApiKey200Response(requestId string, data IngestionApiKey, ) *UpdateIngestionApiKey200Response`

NewUpdateIngestionApiKey200Response instantiates a new UpdateIngestionApiKey200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIngestionApiKey200ResponseWithDefaults

`func NewUpdateIngestionApiKey200ResponseWithDefaults() *UpdateIngestionApiKey200Response`

NewUpdateIngestionApiKey200ResponseWithDefaults instantiates a new UpdateIngestionApiKey200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *UpdateIngestionApiKey200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *UpdateIngestionApiKey200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *UpdateIngestionApiKey200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *UpdateIngestionApiKey200Response) GetData() IngestionApiKey`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateIngestionApiKey200Response) GetDataOk() (*IngestionApiKey, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateIngestionApiKey200Response) SetData(v IngestionApiKey)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


