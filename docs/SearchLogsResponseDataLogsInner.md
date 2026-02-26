# SearchLogsResponseDataLogsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**SearchLogsResponseDataLogsInnerContext**](SearchLogsResponseDataLogsInnerContext.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Timestamp** | **float32** |  | 
**Level** | **string** |  | 

## Methods

### NewSearchLogsResponseDataLogsInner

`func NewSearchLogsResponseDataLogsInner(timestamp float32, level string, ) *SearchLogsResponseDataLogsInner`

NewSearchLogsResponseDataLogsInner instantiates a new SearchLogsResponseDataLogsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchLogsResponseDataLogsInnerWithDefaults

`func NewSearchLogsResponseDataLogsInnerWithDefaults() *SearchLogsResponseDataLogsInner`

NewSearchLogsResponseDataLogsInnerWithDefaults instantiates a new SearchLogsResponseDataLogsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *SearchLogsResponseDataLogsInner) GetContext() SearchLogsResponseDataLogsInnerContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *SearchLogsResponseDataLogsInner) GetContextOk() (*SearchLogsResponseDataLogsInnerContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *SearchLogsResponseDataLogsInner) SetContext(v SearchLogsResponseDataLogsInnerContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *SearchLogsResponseDataLogsInner) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetMessage

`func (o *SearchLogsResponseDataLogsInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SearchLogsResponseDataLogsInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SearchLogsResponseDataLogsInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SearchLogsResponseDataLogsInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTimestamp

`func (o *SearchLogsResponseDataLogsInner) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SearchLogsResponseDataLogsInner) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SearchLogsResponseDataLogsInner) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.


### GetLevel

`func (o *SearchLogsResponseDataLogsInner) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *SearchLogsResponseDataLogsInner) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *SearchLogsResponseDataLogsInner) SetLevel(v string)`

SetLevel sets Level field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


