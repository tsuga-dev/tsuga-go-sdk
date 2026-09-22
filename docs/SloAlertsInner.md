# SloAlertsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Tsuga-generated SLO alert ID assigned when the alert is created. | 
**SloId** | **string** | Parent SLO ID for this alert. | 
**Priority** | **float32** | Alert priority (1 is highest, 5 is lowest) | 
**Configuration** | [**SloAlertsInnerConfiguration**](SloAlertsInnerConfiguration.md) |  | 

## Methods

### NewSloAlertsInner

`func NewSloAlertsInner(id string, sloId string, priority float32, configuration SloAlertsInnerConfiguration, ) *SloAlertsInner`

NewSloAlertsInner instantiates a new SloAlertsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSloAlertsInnerWithDefaults

`func NewSloAlertsInnerWithDefaults() *SloAlertsInner`

NewSloAlertsInnerWithDefaults instantiates a new SloAlertsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SloAlertsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SloAlertsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SloAlertsInner) SetId(v string)`

SetId sets Id field to given value.


### GetSloId

`func (o *SloAlertsInner) GetSloId() string`

GetSloId returns the SloId field if non-nil, zero value otherwise.

### GetSloIdOk

`func (o *SloAlertsInner) GetSloIdOk() (*string, bool)`

GetSloIdOk returns a tuple with the SloId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloId

`func (o *SloAlertsInner) SetSloId(v string)`

SetSloId sets SloId field to given value.


### GetPriority

`func (o *SloAlertsInner) GetPriority() float32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SloAlertsInner) GetPriorityOk() (*float32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SloAlertsInner) SetPriority(v float32)`

SetPriority sets Priority field to given value.


### GetConfiguration

`func (o *SloAlertsInner) GetConfiguration() SloAlertsInnerConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *SloAlertsInner) GetConfigurationOk() (*SloAlertsInnerConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *SloAlertsInner) SetConfiguration(v SloAlertsInnerConfiguration)`

SetConfiguration sets Configuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


