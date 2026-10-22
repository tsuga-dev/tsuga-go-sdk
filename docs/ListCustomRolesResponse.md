# ListCustomRolesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**[]CustomRole**](CustomRole.md) |  | 

## Methods

### NewListCustomRolesResponse

`func NewListCustomRolesResponse(requestId string, data []CustomRole, ) *ListCustomRolesResponse`

NewListCustomRolesResponse instantiates a new ListCustomRolesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCustomRolesResponseWithDefaults

`func NewListCustomRolesResponseWithDefaults() *ListCustomRolesResponse`

NewListCustomRolesResponseWithDefaults instantiates a new ListCustomRolesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ListCustomRolesResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ListCustomRolesResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ListCustomRolesResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *ListCustomRolesResponse) GetData() []CustomRole`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListCustomRolesResponse) GetDataOk() (*[]CustomRole, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListCustomRolesResponse) SetData(v []CustomRole)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


