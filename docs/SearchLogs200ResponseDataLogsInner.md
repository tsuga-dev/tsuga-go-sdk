# SearchLogs200ResponseDataLogsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**SearchLogs200ResponseDataLogsInnerContext**](SearchLogs200ResponseDataLogsInnerContext.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Timestamp** | **float32** |  | 
**Level** | **string** |  | 

## Methods

### NewSearchLogs200ResponseDataLogsInner

`func NewSearchLogs200ResponseDataLogsInner(timestamp float32, level string, ) *SearchLogs200ResponseDataLogsInner`

NewSearchLogs200ResponseDataLogsInner instantiates a new SearchLogs200ResponseDataLogsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchLogs200ResponseDataLogsInnerWithDefaults

`func NewSearchLogs200ResponseDataLogsInnerWithDefaults() *SearchLogs200ResponseDataLogsInner`

NewSearchLogs200ResponseDataLogsInnerWithDefaults instantiates a new SearchLogs200ResponseDataLogsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *SearchLogs200ResponseDataLogsInner) GetContext() SearchLogs200ResponseDataLogsInnerContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *SearchLogs200ResponseDataLogsInner) GetContextOk() (*SearchLogs200ResponseDataLogsInnerContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *SearchLogs200ResponseDataLogsInner) SetContext(v SearchLogs200ResponseDataLogsInnerContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *SearchLogs200ResponseDataLogsInner) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetMessage

`func (o *SearchLogs200ResponseDataLogsInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SearchLogs200ResponseDataLogsInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SearchLogs200ResponseDataLogsInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SearchLogs200ResponseDataLogsInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTimestamp

`func (o *SearchLogs200ResponseDataLogsInner) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SearchLogs200ResponseDataLogsInner) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SearchLogs200ResponseDataLogsInner) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.


### GetLevel

`func (o *SearchLogs200ResponseDataLogsInner) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *SearchLogs200ResponseDataLogsInner) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *SearchLogs200ResponseDataLogsInner) SetLevel(v string)`

SetLevel sets Level field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


