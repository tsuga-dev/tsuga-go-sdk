# SearchLogsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | [**[]SearchLogsResponseDataLogsInner**](SearchLogsResponseDataLogsInner.md) |  | 

## Methods

### NewSearchLogsResponseData

`func NewSearchLogsResponseData(logs []SearchLogsResponseDataLogsInner, ) *SearchLogsResponseData`

NewSearchLogsResponseData instantiates a new SearchLogsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchLogsResponseDataWithDefaults

`func NewSearchLogsResponseDataWithDefaults() *SearchLogsResponseData`

NewSearchLogsResponseDataWithDefaults instantiates a new SearchLogsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *SearchLogsResponseData) GetLogs() []SearchLogsResponseDataLogsInner`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *SearchLogsResponseData) GetLogsOk() (*[]SearchLogsResponseDataLogsInner, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *SearchLogsResponseData) SetLogs(v []SearchLogsResponseDataLogsInner)`

SetLogs sets Logs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


