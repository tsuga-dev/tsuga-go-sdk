# CustomRoleTelemetryQueryInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Attribute key matched against telemetry data points, such as &#x60;env&#x60; or &#x60;service.name&#x60;. | 
**Values** | **[]string** | Values the attribute must equal for a data point to be kept. A data point matching any of them is kept. | 

## Methods

### NewCustomRoleTelemetryQueryInner

`func NewCustomRoleTelemetryQueryInner(key string, values []string, ) *CustomRoleTelemetryQueryInner`

NewCustomRoleTelemetryQueryInner instantiates a new CustomRoleTelemetryQueryInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomRoleTelemetryQueryInnerWithDefaults

`func NewCustomRoleTelemetryQueryInnerWithDefaults() *CustomRoleTelemetryQueryInner`

NewCustomRoleTelemetryQueryInnerWithDefaults instantiates a new CustomRoleTelemetryQueryInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *CustomRoleTelemetryQueryInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CustomRoleTelemetryQueryInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CustomRoleTelemetryQueryInner) SetKey(v string)`

SetKey sets Key field to given value.


### GetValues

`func (o *CustomRoleTelemetryQueryInner) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *CustomRoleTelemetryQueryInner) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *CustomRoleTelemetryQueryInner) SetValues(v []string)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


