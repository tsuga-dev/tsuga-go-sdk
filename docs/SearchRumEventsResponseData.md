# SearchRumEventsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]SearchRumEventsResponseDataEventsInner**](SearchRumEventsResponseDataEventsInner.md) | RUM events matching the query and time range. | 

## Methods

### NewSearchRumEventsResponseData

`func NewSearchRumEventsResponseData(events []SearchRumEventsResponseDataEventsInner, ) *SearchRumEventsResponseData`

NewSearchRumEventsResponseData instantiates a new SearchRumEventsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchRumEventsResponseDataWithDefaults

`func NewSearchRumEventsResponseDataWithDefaults() *SearchRumEventsResponseData`

NewSearchRumEventsResponseDataWithDefaults instantiates a new SearchRumEventsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *SearchRumEventsResponseData) GetEvents() []SearchRumEventsResponseDataEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *SearchRumEventsResponseData) GetEventsOk() (*[]SearchRumEventsResponseDataEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *SearchRumEventsResponseData) SetEvents(v []SearchRumEventsResponseDataEventsInner)`

SetEvents sets Events field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


