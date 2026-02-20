# CreateTagPolicy200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**TagPolicy**](TagPolicy.md) |  | 

## Methods

### NewCreateTagPolicy200Response

`func NewCreateTagPolicy200Response(requestId string, data TagPolicy, ) *CreateTagPolicy200Response`

NewCreateTagPolicy200Response instantiates a new CreateTagPolicy200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTagPolicy200ResponseWithDefaults

`func NewCreateTagPolicy200ResponseWithDefaults() *CreateTagPolicy200Response`

NewCreateTagPolicy200ResponseWithDefaults instantiates a new CreateTagPolicy200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *CreateTagPolicy200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateTagPolicy200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateTagPolicy200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *CreateTagPolicy200Response) GetData() TagPolicy`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateTagPolicy200Response) GetDataOk() (*TagPolicy, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateTagPolicy200Response) SetData(v TagPolicy)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


