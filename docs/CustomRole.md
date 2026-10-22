# CustomRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Tsuga-generated custom role ID assigned at creation. | 
**Name** | **string** | Human-readable custom role name, unique across the account. | 
**Description** | Pointer to **string** | Free-form description of the custom role. Absent when not set. | [optional] 
**TeamScope** | [**CustomRoleTeamScope**](CustomRoleTeamScope.md) |  | 
**Assets** | [**CreateCustomRoleRequestAssets**](CreateCustomRoleRequestAssets.md) |  | 
**Telemetry** | **[]string** | Telemetry signals the role can read. An empty array grants no telemetry access. | 
**TelemetryQuery** | Pointer to [**[]CustomRoleTelemetryQueryInner**](CustomRoleTelemetryQueryInner.md) | Additional restriction applied cluster-side to every telemetry query the role holder runs. Each entry keeps only the data points whose &#x60;key&#x60; matches one of &#x60;values&#x60;, and entries are combined with AND. | [optional] 

## Methods

### NewCustomRole

`func NewCustomRole(id string, name string, teamScope CustomRoleTeamScope, assets CreateCustomRoleRequestAssets, telemetry []string, ) *CustomRole`

NewCustomRole instantiates a new CustomRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomRoleWithDefaults

`func NewCustomRoleWithDefaults() *CustomRole`

NewCustomRoleWithDefaults instantiates a new CustomRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomRole) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CustomRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomRole) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CustomRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTeamScope

`func (o *CustomRole) GetTeamScope() CustomRoleTeamScope`

GetTeamScope returns the TeamScope field if non-nil, zero value otherwise.

### GetTeamScopeOk

`func (o *CustomRole) GetTeamScopeOk() (*CustomRoleTeamScope, bool)`

GetTeamScopeOk returns a tuple with the TeamScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamScope

`func (o *CustomRole) SetTeamScope(v CustomRoleTeamScope)`

SetTeamScope sets TeamScope field to given value.


### GetAssets

`func (o *CustomRole) GetAssets() CreateCustomRoleRequestAssets`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *CustomRole) GetAssetsOk() (*CreateCustomRoleRequestAssets, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *CustomRole) SetAssets(v CreateCustomRoleRequestAssets)`

SetAssets sets Assets field to given value.


### GetTelemetry

`func (o *CustomRole) GetTelemetry() []string`

GetTelemetry returns the Telemetry field if non-nil, zero value otherwise.

### GetTelemetryOk

`func (o *CustomRole) GetTelemetryOk() (*[]string, bool)`

GetTelemetryOk returns a tuple with the Telemetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelemetry

`func (o *CustomRole) SetTelemetry(v []string)`

SetTelemetry sets Telemetry field to given value.


### GetTelemetryQuery

`func (o *CustomRole) GetTelemetryQuery() []CustomRoleTelemetryQueryInner`

GetTelemetryQuery returns the TelemetryQuery field if non-nil, zero value otherwise.

### GetTelemetryQueryOk

`func (o *CustomRole) GetTelemetryQueryOk() (*[]CustomRoleTelemetryQueryInner, bool)`

GetTelemetryQueryOk returns a tuple with the TelemetryQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelemetryQuery

`func (o *CustomRole) SetTelemetryQuery(v []CustomRoleTelemetryQueryInner)`

SetTelemetryQuery sets TelemetryQuery field to given value.

### HasTelemetryQuery

`func (o *CustomRole) HasTelemetryQuery() bool`

HasTelemetryQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


