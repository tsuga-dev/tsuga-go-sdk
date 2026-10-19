# QueryOngoingAlertsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertsPage** | [**[]QueryOngoingAlertsResponseDataAlertsPageInner**](QueryOngoingAlertsResponseDataAlertsPageInner.md) | Page of currently open alerts, ordered by the requested sort. | 

## Methods

### NewQueryOngoingAlertsResponseData

`func NewQueryOngoingAlertsResponseData(alertsPage []QueryOngoingAlertsResponseDataAlertsPageInner, ) *QueryOngoingAlertsResponseData`

NewQueryOngoingAlertsResponseData instantiates a new QueryOngoingAlertsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryOngoingAlertsResponseDataWithDefaults

`func NewQueryOngoingAlertsResponseDataWithDefaults() *QueryOngoingAlertsResponseData`

NewQueryOngoingAlertsResponseDataWithDefaults instantiates a new QueryOngoingAlertsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertsPage

`func (o *QueryOngoingAlertsResponseData) GetAlertsPage() []QueryOngoingAlertsResponseDataAlertsPageInner`

GetAlertsPage returns the AlertsPage field if non-nil, zero value otherwise.

### GetAlertsPageOk

`func (o *QueryOngoingAlertsResponseData) GetAlertsPageOk() (*[]QueryOngoingAlertsResponseDataAlertsPageInner, bool)`

GetAlertsPageOk returns a tuple with the AlertsPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertsPage

`func (o *QueryOngoingAlertsResponseData) SetAlertsPage(v []QueryOngoingAlertsResponseDataAlertsPageInner)`

SetAlertsPage sets AlertsPage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


