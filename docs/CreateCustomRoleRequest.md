# CreateCustomRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable custom role name. | 
**Description** | Pointer to **NullableString** | Free-form description of the custom role. Omit on update to keep the current description, or pass null to remove it. | [optional] 
**TeamScope** | [**CreateCustomRoleRequestTeamScope**](CreateCustomRoleRequestTeamScope.md) |  | 
**Assets** | [**CreateCustomRoleRequestAssets**](CreateCustomRoleRequestAssets.md) |  | 
**Telemetry** | **[]string** | Telemetry signals the role can read. An empty array grants no telemetry access. | 
**TelemetryQuery** | Pointer to [**[]CreateCustomRoleRequestTelemetryQueryInner**](CreateCustomRoleRequestTelemetryQueryInner.md) | Additional restriction applied cluster-side to every telemetry query the role holder runs. Each entry keeps only the data points whose &#x60;key&#x60; matches one of &#x60;values&#x60;, and entries are combined with AND. | [optional] 

## Methods

### NewCreateCustomRoleRequest

`func NewCreateCustomRoleRequest(name string, teamScope CreateCustomRoleRequestTeamScope, assets CreateCustomRoleRequestAssets, telemetry []string, ) *CreateCustomRoleRequest`

NewCreateCustomRoleRequest instantiates a new CreateCustomRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomRoleRequestWithDefaults

`func NewCreateCustomRoleRequestWithDefaults() *CreateCustomRoleRequest`

NewCreateCustomRoleRequestWithDefaults instantiates a new CreateCustomRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateCustomRoleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCustomRoleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCustomRoleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateCustomRoleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCustomRoleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCustomRoleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateCustomRoleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateCustomRoleRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateCustomRoleRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTeamScope

`func (o *CreateCustomRoleRequest) GetTeamScope() CreateCustomRoleRequestTeamScope`

GetTeamScope returns the TeamScope field if non-nil, zero value otherwise.

### GetTeamScopeOk

`func (o *CreateCustomRoleRequest) GetTeamScopeOk() (*CreateCustomRoleRequestTeamScope, bool)`

GetTeamScopeOk returns a tuple with the TeamScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamScope

`func (o *CreateCustomRoleRequest) SetTeamScope(v CreateCustomRoleRequestTeamScope)`

SetTeamScope sets TeamScope field to given value.


### GetAssets

`func (o *CreateCustomRoleRequest) GetAssets() CreateCustomRoleRequestAssets`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *CreateCustomRoleRequest) GetAssetsOk() (*CreateCustomRoleRequestAssets, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *CreateCustomRoleRequest) SetAssets(v CreateCustomRoleRequestAssets)`

SetAssets sets Assets field to given value.


### GetTelemetry

`func (o *CreateCustomRoleRequest) GetTelemetry() []string`

GetTelemetry returns the Telemetry field if non-nil, zero value otherwise.

### GetTelemetryOk

`func (o *CreateCustomRoleRequest) GetTelemetryOk() (*[]string, bool)`

GetTelemetryOk returns a tuple with the Telemetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelemetry

`func (o *CreateCustomRoleRequest) SetTelemetry(v []string)`

SetTelemetry sets Telemetry field to given value.


### GetTelemetryQuery

`func (o *CreateCustomRoleRequest) GetTelemetryQuery() []CreateCustomRoleRequestTelemetryQueryInner`

GetTelemetryQuery returns the TelemetryQuery field if non-nil, zero value otherwise.

### GetTelemetryQueryOk

`func (o *CreateCustomRoleRequest) GetTelemetryQueryOk() (*[]CreateCustomRoleRequestTelemetryQueryInner, bool)`

GetTelemetryQueryOk returns a tuple with the TelemetryQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelemetryQuery

`func (o *CreateCustomRoleRequest) SetTelemetryQuery(v []CreateCustomRoleRequestTelemetryQueryInner)`

SetTelemetryQuery sets TelemetryQuery field to given value.

### HasTelemetryQuery

`func (o *CreateCustomRoleRequest) HasTelemetryQuery() bool`

HasTelemetryQuery returns a boolean if a field has been set.

### SetTelemetryQueryNil

`func (o *CreateCustomRoleRequest) SetTelemetryQueryNil(b bool)`

 SetTelemetryQueryNil sets the value for TelemetryQuery to be an explicit nil

### UnsetTelemetryQuery
`func (o *CreateCustomRoleRequest) UnsetTelemetryQuery()`

UnsetTelemetryQuery ensures that no value is present for TelemetryQuery, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


