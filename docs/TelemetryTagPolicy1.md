# TelemetryTagPolicy1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Policy applies to telemetry data points during ingestion or evaluation. | 
**AssetTypes** | **[]string** | Telemetry signals covered by the policy. An empty array means logs, metrics, and traces. | 
**ShouldInsertWarning** | **bool** | When true, Tsuga annotates non-compliant telemetry with tag policy warnings instead of only measuring violations. | 
**DropSample** | Pointer to **float32** | Percentage of non-compliant telemetry datapoints to drop. Omitted means no sampling drop is configured. | [optional] 

## Methods

### NewTelemetryTagPolicy1

`func NewTelemetryTagPolicy1(type_ string, assetTypes []string, shouldInsertWarning bool, ) *TelemetryTagPolicy1`

NewTelemetryTagPolicy1 instantiates a new TelemetryTagPolicy1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryTagPolicy1WithDefaults

`func NewTelemetryTagPolicy1WithDefaults() *TelemetryTagPolicy1`

NewTelemetryTagPolicy1WithDefaults instantiates a new TelemetryTagPolicy1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryTagPolicy1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryTagPolicy1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryTagPolicy1) SetType(v string)`

SetType sets Type field to given value.


### GetAssetTypes

`func (o *TelemetryTagPolicy1) GetAssetTypes() []string`

GetAssetTypes returns the AssetTypes field if non-nil, zero value otherwise.

### GetAssetTypesOk

`func (o *TelemetryTagPolicy1) GetAssetTypesOk() (*[]string, bool)`

GetAssetTypesOk returns a tuple with the AssetTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTypes

`func (o *TelemetryTagPolicy1) SetAssetTypes(v []string)`

SetAssetTypes sets AssetTypes field to given value.


### GetShouldInsertWarning

`func (o *TelemetryTagPolicy1) GetShouldInsertWarning() bool`

GetShouldInsertWarning returns the ShouldInsertWarning field if non-nil, zero value otherwise.

### GetShouldInsertWarningOk

`func (o *TelemetryTagPolicy1) GetShouldInsertWarningOk() (*bool, bool)`

GetShouldInsertWarningOk returns a tuple with the ShouldInsertWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldInsertWarning

`func (o *TelemetryTagPolicy1) SetShouldInsertWarning(v bool)`

SetShouldInsertWarning sets ShouldInsertWarning field to given value.


### GetDropSample

`func (o *TelemetryTagPolicy1) GetDropSample() float32`

GetDropSample returns the DropSample field if non-nil, zero value otherwise.

### GetDropSampleOk

`func (o *TelemetryTagPolicy1) GetDropSampleOk() (*float32, bool)`

GetDropSampleOk returns a tuple with the DropSample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropSample

`func (o *TelemetryTagPolicy1) SetDropSample(v float32)`

SetDropSample sets DropSample field to given value.

### HasDropSample

`func (o *TelemetryTagPolicy1) HasDropSample() bool`

HasDropSample returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


