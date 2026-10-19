# CreateCloudAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**CloudAccount**](CloudAccount.md) |  | 

## Methods

### NewCreateCloudAccountResponse

`func NewCreateCloudAccountResponse(requestId string, data CloudAccount, ) *CreateCloudAccountResponse`

NewCreateCloudAccountResponse instantiates a new CreateCloudAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCloudAccountResponseWithDefaults

`func NewCreateCloudAccountResponseWithDefaults() *CreateCloudAccountResponse`

NewCreateCloudAccountResponseWithDefaults instantiates a new CreateCloudAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *CreateCloudAccountResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateCloudAccountResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateCloudAccountResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *CreateCloudAccountResponse) GetData() CloudAccount`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateCloudAccountResponse) GetDataOk() (*CloudAccount, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateCloudAccountResponse) SetData(v CloudAccount)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


