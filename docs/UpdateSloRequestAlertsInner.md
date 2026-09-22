# UpdateSloRequestAlertsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifier of an existing alert to update. Omit to create a new alert on this SLO. | [optional] 
**Priority** | **float32** | Alert priority (1 is highest, 5 is lowest) | 
**Configuration** | [**UpdateSloRequestAlertsInnerConfiguration**](UpdateSloRequestAlertsInnerConfiguration.md) |  | 

## Methods

### NewUpdateSloRequestAlertsInner

`func NewUpdateSloRequestAlertsInner(priority float32, configuration UpdateSloRequestAlertsInnerConfiguration, ) *UpdateSloRequestAlertsInner`

NewUpdateSloRequestAlertsInner instantiates a new UpdateSloRequestAlertsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSloRequestAlertsInnerWithDefaults

`func NewUpdateSloRequestAlertsInnerWithDefaults() *UpdateSloRequestAlertsInner`

NewUpdateSloRequestAlertsInnerWithDefaults instantiates a new UpdateSloRequestAlertsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateSloRequestAlertsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateSloRequestAlertsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateSloRequestAlertsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateSloRequestAlertsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPriority

`func (o *UpdateSloRequestAlertsInner) GetPriority() float32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *UpdateSloRequestAlertsInner) GetPriorityOk() (*float32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *UpdateSloRequestAlertsInner) SetPriority(v float32)`

SetPriority sets Priority field to given value.


### GetConfiguration

`func (o *UpdateSloRequestAlertsInner) GetConfiguration() UpdateSloRequestAlertsInnerConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *UpdateSloRequestAlertsInner) GetConfigurationOk() (*UpdateSloRequestAlertsInnerConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *UpdateSloRequestAlertsInner) SetConfiguration(v UpdateSloRequestAlertsInnerConfiguration)`

SetConfiguration sets Configuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


