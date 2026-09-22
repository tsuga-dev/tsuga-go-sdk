# CreateSloResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**Slo**](Slo.md) |  | 

## Methods

### NewCreateSloResponse

`func NewCreateSloResponse(requestId string, data Slo, ) *CreateSloResponse`

NewCreateSloResponse instantiates a new CreateSloResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSloResponseWithDefaults

`func NewCreateSloResponseWithDefaults() *CreateSloResponse`

NewCreateSloResponseWithDefaults instantiates a new CreateSloResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *CreateSloResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateSloResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateSloResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *CreateSloResponse) GetData() Slo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateSloResponse) GetDataOk() (*Slo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateSloResponse) SetData(v Slo)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


