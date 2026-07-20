# TelemetryTagPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Policy applies to telemetry data points during ingestion or evaluation. | 
**AssetTypes** | **[]string** | Telemetry signals covered by the policy. An empty array means logs, metrics, and traces. | 
**ShouldInsertWarning** | **bool** | When true, Tsuga annotates non-compliant telemetry with tag policy warnings instead of only measuring violations. | 
**DropSample** | Pointer to **float32** | Percentage of non-compliant telemetry datapoints to drop. Valid values are 0 through 100. | [optional] 

## Methods

### NewTelemetryTagPolicy

`func NewTelemetryTagPolicy(type_ string, assetTypes []string, shouldInsertWarning bool, ) *TelemetryTagPolicy`

NewTelemetryTagPolicy instantiates a new TelemetryTagPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryTagPolicyWithDefaults

`func NewTelemetryTagPolicyWithDefaults() *TelemetryTagPolicy`

NewTelemetryTagPolicyWithDefaults instantiates a new TelemetryTagPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryTagPolicy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryTagPolicy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryTagPolicy) SetType(v string)`

SetType sets Type field to given value.


### GetAssetTypes

`func (o *TelemetryTagPolicy) GetAssetTypes() []string`

GetAssetTypes returns the AssetTypes field if non-nil, zero value otherwise.

### GetAssetTypesOk

`func (o *TelemetryTagPolicy) GetAssetTypesOk() (*[]string, bool)`

GetAssetTypesOk returns a tuple with the AssetTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTypes

`func (o *TelemetryTagPolicy) SetAssetTypes(v []string)`

SetAssetTypes sets AssetTypes field to given value.


### GetShouldInsertWarning

`func (o *TelemetryTagPolicy) GetShouldInsertWarning() bool`

GetShouldInsertWarning returns the ShouldInsertWarning field if non-nil, zero value otherwise.

### GetShouldInsertWarningOk

`func (o *TelemetryTagPolicy) GetShouldInsertWarningOk() (*bool, bool)`

GetShouldInsertWarningOk returns a tuple with the ShouldInsertWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldInsertWarning

`func (o *TelemetryTagPolicy) SetShouldInsertWarning(v bool)`

SetShouldInsertWarning sets ShouldInsertWarning field to given value.


### GetDropSample

`func (o *TelemetryTagPolicy) GetDropSample() float32`

GetDropSample returns the DropSample field if non-nil, zero value otherwise.

### GetDropSampleOk

`func (o *TelemetryTagPolicy) GetDropSampleOk() (*float32, bool)`

GetDropSampleOk returns a tuple with the DropSample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropSample

`func (o *TelemetryTagPolicy) SetDropSample(v float32)`

SetDropSample sets DropSample field to given value.

### HasDropSample

`func (o *TelemetryTagPolicy) HasDropSample() bool`

HasDropSample returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


