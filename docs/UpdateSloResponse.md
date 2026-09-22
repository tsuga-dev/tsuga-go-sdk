# UpdateSloResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**Slo**](Slo.md) |  | 

## Methods

### NewUpdateSloResponse

`func NewUpdateSloResponse(requestId string, data Slo, ) *UpdateSloResponse`

NewUpdateSloResponse instantiates a new UpdateSloResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSloResponseWithDefaults

`func NewUpdateSloResponseWithDefaults() *UpdateSloResponse`

NewUpdateSloResponseWithDefaults instantiates a new UpdateSloResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *UpdateSloResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *UpdateSloResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *UpdateSloResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *UpdateSloResponse) GetData() Slo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateSloResponse) GetDataOk() (*Slo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateSloResponse) SetData(v Slo)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


