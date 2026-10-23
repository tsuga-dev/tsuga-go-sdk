# OngoingSloAlertSloAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Tsuga-generated SLO alert ID assigned when the alert is created. | 
**Priority** | **int32** | Alert priority (1 is highest, 5 is lowest) | 
**Configuration** | [**SloAlertsInnerConfiguration**](SloAlertsInnerConfiguration.md) |  | 

## Methods

### NewOngoingSloAlertSloAlert

`func NewOngoingSloAlertSloAlert(id string, priority int32, configuration SloAlertsInnerConfiguration, ) *OngoingSloAlertSloAlert`

NewOngoingSloAlertSloAlert instantiates a new OngoingSloAlertSloAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOngoingSloAlertSloAlertWithDefaults

`func NewOngoingSloAlertSloAlertWithDefaults() *OngoingSloAlertSloAlert`

NewOngoingSloAlertSloAlertWithDefaults instantiates a new OngoingSloAlertSloAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OngoingSloAlertSloAlert) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OngoingSloAlertSloAlert) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OngoingSloAlertSloAlert) SetId(v string)`

SetId sets Id field to given value.


### GetPriority

`func (o *OngoingSloAlertSloAlert) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *OngoingSloAlertSloAlert) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *OngoingSloAlertSloAlert) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetConfiguration

`func (o *OngoingSloAlertSloAlert) GetConfiguration() SloAlertsInnerConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *OngoingSloAlertSloAlert) GetConfigurationOk() (*SloAlertsInnerConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *OngoingSloAlertSloAlert) SetConfiguration(v SloAlertsInnerConfiguration)`

SetConfiguration sets Configuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


