# CreateSloRequestAlertsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Priority** | **float32** | Alert priority (1 is highest, 5 is lowest) | 
**Configuration** | [**UpdateSloRequestAlertsInnerConfiguration**](UpdateSloRequestAlertsInnerConfiguration.md) |  | 

## Methods

### NewCreateSloRequestAlertsInner

`func NewCreateSloRequestAlertsInner(priority float32, configuration UpdateSloRequestAlertsInnerConfiguration, ) *CreateSloRequestAlertsInner`

NewCreateSloRequestAlertsInner instantiates a new CreateSloRequestAlertsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSloRequestAlertsInnerWithDefaults

`func NewCreateSloRequestAlertsInnerWithDefaults() *CreateSloRequestAlertsInner`

NewCreateSloRequestAlertsInnerWithDefaults instantiates a new CreateSloRequestAlertsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriority

`func (o *CreateSloRequestAlertsInner) GetPriority() float32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CreateSloRequestAlertsInner) GetPriorityOk() (*float32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CreateSloRequestAlertsInner) SetPriority(v float32)`

SetPriority sets Priority field to given value.


### GetConfiguration

`func (o *CreateSloRequestAlertsInner) GetConfiguration() UpdateSloRequestAlertsInnerConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *CreateSloRequestAlertsInner) GetConfigurationOk() (*UpdateSloRequestAlertsInnerConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *CreateSloRequestAlertsInner) SetConfiguration(v UpdateSloRequestAlertsInnerConfiguration)`

SetConfiguration sets Configuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


