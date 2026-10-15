# ListAdaptiveNotificationRulesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**[]AdaptiveRule**](AdaptiveRule.md) |  | 
**Metadata** | Pointer to [**ResponseMetadata**](ResponseMetadata.md) |  | [optional] 

## Methods

### NewListAdaptiveNotificationRulesResponse

`func NewListAdaptiveNotificationRulesResponse(requestId string, data []AdaptiveRule, ) *ListAdaptiveNotificationRulesResponse`

NewListAdaptiveNotificationRulesResponse instantiates a new ListAdaptiveNotificationRulesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAdaptiveNotificationRulesResponseWithDefaults

`func NewListAdaptiveNotificationRulesResponseWithDefaults() *ListAdaptiveNotificationRulesResponse`

NewListAdaptiveNotificationRulesResponseWithDefaults instantiates a new ListAdaptiveNotificationRulesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ListAdaptiveNotificationRulesResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ListAdaptiveNotificationRulesResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ListAdaptiveNotificationRulesResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *ListAdaptiveNotificationRulesResponse) GetData() []AdaptiveRule`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListAdaptiveNotificationRulesResponse) GetDataOk() (*[]AdaptiveRule, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListAdaptiveNotificationRulesResponse) SetData(v []AdaptiveRule)`

SetData sets Data field to given value.


### GetMetadata

`func (o *ListAdaptiveNotificationRulesResponse) GetMetadata() ResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListAdaptiveNotificationRulesResponse) GetMetadataOk() (*ResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListAdaptiveNotificationRulesResponse) SetMetadata(v ResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ListAdaptiveNotificationRulesResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


