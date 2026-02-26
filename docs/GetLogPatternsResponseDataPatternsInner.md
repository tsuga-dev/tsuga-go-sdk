# GetLogPatternsResponseDataPatternsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pattern** | [**[]TokenOutput**](TokenOutput.md) |  | 
**Size** | **int32** | Number of logs that matched the pattern | 
**Groups** | [**[]GetLogPatternsResponseDataPatternsInnerGroupsInner**](GetLogPatternsResponseDataPatternsInnerGroupsInner.md) |  | 

## Methods

### NewGetLogPatternsResponseDataPatternsInner

`func NewGetLogPatternsResponseDataPatternsInner(pattern []TokenOutput, size int32, groups []GetLogPatternsResponseDataPatternsInnerGroupsInner, ) *GetLogPatternsResponseDataPatternsInner`

NewGetLogPatternsResponseDataPatternsInner instantiates a new GetLogPatternsResponseDataPatternsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLogPatternsResponseDataPatternsInnerWithDefaults

`func NewGetLogPatternsResponseDataPatternsInnerWithDefaults() *GetLogPatternsResponseDataPatternsInner`

NewGetLogPatternsResponseDataPatternsInnerWithDefaults instantiates a new GetLogPatternsResponseDataPatternsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPattern

`func (o *GetLogPatternsResponseDataPatternsInner) GetPattern() []TokenOutput`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *GetLogPatternsResponseDataPatternsInner) GetPatternOk() (*[]TokenOutput, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *GetLogPatternsResponseDataPatternsInner) SetPattern(v []TokenOutput)`

SetPattern sets Pattern field to given value.


### GetSize

`func (o *GetLogPatternsResponseDataPatternsInner) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetLogPatternsResponseDataPatternsInner) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetLogPatternsResponseDataPatternsInner) SetSize(v int32)`

SetSize sets Size field to given value.


### GetGroups

`func (o *GetLogPatternsResponseDataPatternsInner) GetGroups() []GetLogPatternsResponseDataPatternsInnerGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *GetLogPatternsResponseDataPatternsInner) GetGroupsOk() (*[]GetLogPatternsResponseDataPatternsInnerGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *GetLogPatternsResponseDataPatternsInner) SetGroups(v []GetLogPatternsResponseDataPatternsInnerGroupsInner)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


