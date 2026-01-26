# CreateIngestionApiKey200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**CreateIngestionApiKey200ResponseData**](CreateIngestionApiKey200ResponseData.md) |  | 

## Methods

### NewCreateIngestionApiKey200Response

`func NewCreateIngestionApiKey200Response(requestId string, data CreateIngestionApiKey200ResponseData, ) *CreateIngestionApiKey200Response`

NewCreateIngestionApiKey200Response instantiates a new CreateIngestionApiKey200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIngestionApiKey200ResponseWithDefaults

`func NewCreateIngestionApiKey200ResponseWithDefaults() *CreateIngestionApiKey200Response`

NewCreateIngestionApiKey200ResponseWithDefaults instantiates a new CreateIngestionApiKey200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *CreateIngestionApiKey200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateIngestionApiKey200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateIngestionApiKey200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *CreateIngestionApiKey200Response) GetData() CreateIngestionApiKey200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateIngestionApiKey200Response) GetDataOk() (*CreateIngestionApiKey200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateIngestionApiKey200Response) SetData(v CreateIngestionApiKey200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


