# LogContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Team** | Pointer to **string** |  | [optional] 

## Methods

### NewLogContext

`func NewLogContext() *LogContext`

NewLogContext instantiates a new LogContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogContextWithDefaults

`func NewLogContextWithDefaults() *LogContext`

NewLogContextWithDefaults instantiates a new LogContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeam

`func (o *LogContext) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *LogContext) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *LogContext) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *LogContext) HasTeam() bool`

HasTeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


