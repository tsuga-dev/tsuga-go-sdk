# CreateTeam200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**Team**](Team.md) |  | 

## Methods

### NewCreateTeam200Response

`func NewCreateTeam200Response(requestId string, data Team, ) *CreateTeam200Response`

NewCreateTeam200Response instantiates a new CreateTeam200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTeam200ResponseWithDefaults

`func NewCreateTeam200ResponseWithDefaults() *CreateTeam200Response`

NewCreateTeam200ResponseWithDefaults instantiates a new CreateTeam200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *CreateTeam200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateTeam200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateTeam200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *CreateTeam200Response) GetData() Team`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateTeam200Response) GetDataOk() (*Team, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateTeam200Response) SetData(v Team)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


