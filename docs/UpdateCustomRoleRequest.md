# UpdateCustomRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable custom role name, unique across the account. | 
**Description** | Pointer to **NullableString** | Free-form description of the custom role. Omit on update to keep the current description, or pass null to remove it. | [optional] 
**TeamScope** | [**UpdateCustomRoleRequestTeamScope**](UpdateCustomRoleRequestTeamScope.md) |  | 
**Assets** | [**CreateCustomRoleRequestAssets**](CreateCustomRoleRequestAssets.md) |  | 
**Telemetry** | **[]string** | Telemetry signals the role can read. An empty array grants no telemetry access. | 
**TelemetryQuery** | Pointer to [**[]CreateCustomRoleRequestTelemetryQueryInner**](CreateCustomRoleRequestTelemetryQueryInner.md) | Additional restriction applied cluster-side to every telemetry query the role holder runs. Each entry keeps only the data points whose &#x60;key&#x60; matches one of &#x60;values&#x60;, and entries are combined with AND. | [optional] 

## Methods

### NewUpdateCustomRoleRequest

`func NewUpdateCustomRoleRequest(name string, teamScope UpdateCustomRoleRequestTeamScope, assets CreateCustomRoleRequestAssets, telemetry []string, ) *UpdateCustomRoleRequest`

NewUpdateCustomRoleRequest instantiates a new UpdateCustomRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCustomRoleRequestWithDefaults

`func NewUpdateCustomRoleRequestWithDefaults() *UpdateCustomRoleRequest`

NewUpdateCustomRoleRequestWithDefaults instantiates a new UpdateCustomRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateCustomRoleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCustomRoleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCustomRoleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateCustomRoleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCustomRoleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCustomRoleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCustomRoleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateCustomRoleRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateCustomRoleRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTeamScope

`func (o *UpdateCustomRoleRequest) GetTeamScope() UpdateCustomRoleRequestTeamScope`

GetTeamScope returns the TeamScope field if non-nil, zero value otherwise.

### GetTeamScopeOk

`func (o *UpdateCustomRoleRequest) GetTeamScopeOk() (*UpdateCustomRoleRequestTeamScope, bool)`

GetTeamScopeOk returns a tuple with the TeamScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamScope

`func (o *UpdateCustomRoleRequest) SetTeamScope(v UpdateCustomRoleRequestTeamScope)`

SetTeamScope sets TeamScope field to given value.


### GetAssets

`func (o *UpdateCustomRoleRequest) GetAssets() CreateCustomRoleRequestAssets`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *UpdateCustomRoleRequest) GetAssetsOk() (*CreateCustomRoleRequestAssets, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *UpdateCustomRoleRequest) SetAssets(v CreateCustomRoleRequestAssets)`

SetAssets sets Assets field to given value.


### GetTelemetry

`func (o *UpdateCustomRoleRequest) GetTelemetry() []string`

GetTelemetry returns the Telemetry field if non-nil, zero value otherwise.

### GetTelemetryOk

`func (o *UpdateCustomRoleRequest) GetTelemetryOk() (*[]string, bool)`

GetTelemetryOk returns a tuple with the Telemetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelemetry

`func (o *UpdateCustomRoleRequest) SetTelemetry(v []string)`

SetTelemetry sets Telemetry field to given value.


### GetTelemetryQuery

`func (o *UpdateCustomRoleRequest) GetTelemetryQuery() []CreateCustomRoleRequestTelemetryQueryInner`

GetTelemetryQuery returns the TelemetryQuery field if non-nil, zero value otherwise.

### GetTelemetryQueryOk

`func (o *UpdateCustomRoleRequest) GetTelemetryQueryOk() (*[]CreateCustomRoleRequestTelemetryQueryInner, bool)`

GetTelemetryQueryOk returns a tuple with the TelemetryQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelemetryQuery

`func (o *UpdateCustomRoleRequest) SetTelemetryQuery(v []CreateCustomRoleRequestTelemetryQueryInner)`

SetTelemetryQuery sets TelemetryQuery field to given value.

### HasTelemetryQuery

`func (o *UpdateCustomRoleRequest) HasTelemetryQuery() bool`

HasTelemetryQuery returns a boolean if a field has been set.

### SetTelemetryQueryNil

`func (o *UpdateCustomRoleRequest) SetTelemetryQueryNil(b bool)`

 SetTelemetryQueryNil sets the value for TelemetryQuery to be an explicit nil

### UnsetTelemetryQuery
`func (o *UpdateCustomRoleRequest) UnsetTelemetryQuery()`

UnsetTelemetryQuery ensures that no value is present for TelemetryQuery, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


