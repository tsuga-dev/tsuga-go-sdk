# GetTagPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**TagPolicy**](TagPolicy.md) |  | 

## Methods

### NewGetTagPolicyResponse

`func NewGetTagPolicyResponse(requestId string, data TagPolicy, ) *GetTagPolicyResponse`

NewGetTagPolicyResponse instantiates a new GetTagPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTagPolicyResponseWithDefaults

`func NewGetTagPolicyResponseWithDefaults() *GetTagPolicyResponse`

NewGetTagPolicyResponseWithDefaults instantiates a new GetTagPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *GetTagPolicyResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetTagPolicyResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetTagPolicyResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *GetTagPolicyResponse) GetData() TagPolicy`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetTagPolicyResponse) GetDataOk() (*TagPolicy, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetTagPolicyResponse) SetData(v TagPolicy)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


