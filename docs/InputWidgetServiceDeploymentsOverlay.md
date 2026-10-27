# InputWidgetServiceDeploymentsOverlay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Draws the deployments contained in the time range over the chart | 
**Teams** | Pointer to **[]string** | IDs of the teams whose services are drawn. The set is resolved when the widget renders, so a service joining one of these teams later is picked up automatically. Every service is drawn when omitted. | [optional] 
**Services** | Pointer to **[]string** | OpenTelemetry service names to draw, narrowing &#x60;teams&#x60; down to an explicit list. Every service matching &#x60;teams&#x60; is drawn when omitted. | [optional] 
**Env** | Pointer to **string** | Environment whose deployments are drawn. Deployments from every environment are drawn when omitted. | [optional] 

## Methods

### NewInputWidgetServiceDeploymentsOverlay

`func NewInputWidgetServiceDeploymentsOverlay(type_ string, ) *InputWidgetServiceDeploymentsOverlay`

NewInputWidgetServiceDeploymentsOverlay instantiates a new InputWidgetServiceDeploymentsOverlay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputWidgetServiceDeploymentsOverlayWithDefaults

`func NewInputWidgetServiceDeploymentsOverlayWithDefaults() *InputWidgetServiceDeploymentsOverlay`

NewInputWidgetServiceDeploymentsOverlayWithDefaults instantiates a new InputWidgetServiceDeploymentsOverlay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputWidgetServiceDeploymentsOverlay) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputWidgetServiceDeploymentsOverlay) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputWidgetServiceDeploymentsOverlay) SetType(v string)`

SetType sets Type field to given value.


### GetTeams

`func (o *InputWidgetServiceDeploymentsOverlay) GetTeams() []string`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *InputWidgetServiceDeploymentsOverlay) GetTeamsOk() (*[]string, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *InputWidgetServiceDeploymentsOverlay) SetTeams(v []string)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *InputWidgetServiceDeploymentsOverlay) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetServices

`func (o *InputWidgetServiceDeploymentsOverlay) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *InputWidgetServiceDeploymentsOverlay) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *InputWidgetServiceDeploymentsOverlay) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *InputWidgetServiceDeploymentsOverlay) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetEnv

`func (o *InputWidgetServiceDeploymentsOverlay) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *InputWidgetServiceDeploymentsOverlay) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *InputWidgetServiceDeploymentsOverlay) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *InputWidgetServiceDeploymentsOverlay) HasEnv() bool`

HasEnv returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


