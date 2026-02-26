# GetRetentionPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**RetentionPolicy**](RetentionPolicy.md) |  | 

## Methods

### NewGetRetentionPolicyResponse

`func NewGetRetentionPolicyResponse(requestId string, data RetentionPolicy, ) *GetRetentionPolicyResponse`

NewGetRetentionPolicyResponse instantiates a new GetRetentionPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRetentionPolicyResponseWithDefaults

`func NewGetRetentionPolicyResponseWithDefaults() *GetRetentionPolicyResponse`

NewGetRetentionPolicyResponseWithDefaults instantiates a new GetRetentionPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *GetRetentionPolicyResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetRetentionPolicyResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetRetentionPolicyResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *GetRetentionPolicyResponse) GetData() RetentionPolicy`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetRetentionPolicyResponse) GetDataOk() (*RetentionPolicy, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetRetentionPolicyResponse) SetData(v RetentionPolicy)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


