# ProcessorParseGeoIp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceAttribute** | **string** | Attribute containing the IP address to look up | 
**TargetAttribute** | **string** | Target map attribute where geolocation details will be written | 
**RedactSource** | **bool** | Whether to replace a public IP address in the source attribute with &#x60;[REDACTED_IP]&#x60;. Private IP addresses and non-IP values are left unchanged. | 
**Subtype** | **string** | Resolves an IP address attribute to geolocation details. | 

## Methods

### NewProcessorParseGeoIp

`func NewProcessorParseGeoIp(sourceAttribute string, targetAttribute string, redactSource bool, subtype string, ) *ProcessorParseGeoIp`

NewProcessorParseGeoIp instantiates a new ProcessorParseGeoIp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorParseGeoIpWithDefaults

`func NewProcessorParseGeoIpWithDefaults() *ProcessorParseGeoIp`

NewProcessorParseGeoIpWithDefaults instantiates a new ProcessorParseGeoIp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceAttribute

`func (o *ProcessorParseGeoIp) GetSourceAttribute() string`

GetSourceAttribute returns the SourceAttribute field if non-nil, zero value otherwise.

### GetSourceAttributeOk

`func (o *ProcessorParseGeoIp) GetSourceAttributeOk() (*string, bool)`

GetSourceAttributeOk returns a tuple with the SourceAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAttribute

`func (o *ProcessorParseGeoIp) SetSourceAttribute(v string)`

SetSourceAttribute sets SourceAttribute field to given value.


### GetTargetAttribute

`func (o *ProcessorParseGeoIp) GetTargetAttribute() string`

GetTargetAttribute returns the TargetAttribute field if non-nil, zero value otherwise.

### GetTargetAttributeOk

`func (o *ProcessorParseGeoIp) GetTargetAttributeOk() (*string, bool)`

GetTargetAttributeOk returns a tuple with the TargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttribute

`func (o *ProcessorParseGeoIp) SetTargetAttribute(v string)`

SetTargetAttribute sets TargetAttribute field to given value.


### GetRedactSource

`func (o *ProcessorParseGeoIp) GetRedactSource() bool`

GetRedactSource returns the RedactSource field if non-nil, zero value otherwise.

### GetRedactSourceOk

`func (o *ProcessorParseGeoIp) GetRedactSourceOk() (*bool, bool)`

GetRedactSourceOk returns a tuple with the RedactSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactSource

`func (o *ProcessorParseGeoIp) SetRedactSource(v bool)`

SetRedactSource sets RedactSource field to given value.


### GetSubtype

`func (o *ProcessorParseGeoIp) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *ProcessorParseGeoIp) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *ProcessorParseGeoIp) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


