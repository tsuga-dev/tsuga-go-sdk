# InputProcessorAnyOf1ParamsOneOf3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceAttribute** | **string** | Attribute containing the IP address to look up | 
**TargetAttribute** | **string** | Target map attribute where geolocation details will be written | 
**RedactSource** | **bool** | Whether to replace a public IP address in the source attribute with &#x60;[REDACTED_IP]&#x60;. Private IP addresses and non-IP values are left unchanged. | 
**Subtype** | **string** | Resolves an IP address attribute to geolocation details. | 

## Methods

### NewInputProcessorAnyOf1ParamsOneOf3

`func NewInputProcessorAnyOf1ParamsOneOf3(sourceAttribute string, targetAttribute string, redactSource bool, subtype string, ) *InputProcessorAnyOf1ParamsOneOf3`

NewInputProcessorAnyOf1ParamsOneOf3 instantiates a new InputProcessorAnyOf1ParamsOneOf3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputProcessorAnyOf1ParamsOneOf3WithDefaults

`func NewInputProcessorAnyOf1ParamsOneOf3WithDefaults() *InputProcessorAnyOf1ParamsOneOf3`

NewInputProcessorAnyOf1ParamsOneOf3WithDefaults instantiates a new InputProcessorAnyOf1ParamsOneOf3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceAttribute

`func (o *InputProcessorAnyOf1ParamsOneOf3) GetSourceAttribute() string`

GetSourceAttribute returns the SourceAttribute field if non-nil, zero value otherwise.

### GetSourceAttributeOk

`func (o *InputProcessorAnyOf1ParamsOneOf3) GetSourceAttributeOk() (*string, bool)`

GetSourceAttributeOk returns a tuple with the SourceAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAttribute

`func (o *InputProcessorAnyOf1ParamsOneOf3) SetSourceAttribute(v string)`

SetSourceAttribute sets SourceAttribute field to given value.


### GetTargetAttribute

`func (o *InputProcessorAnyOf1ParamsOneOf3) GetTargetAttribute() string`

GetTargetAttribute returns the TargetAttribute field if non-nil, zero value otherwise.

### GetTargetAttributeOk

`func (o *InputProcessorAnyOf1ParamsOneOf3) GetTargetAttributeOk() (*string, bool)`

GetTargetAttributeOk returns a tuple with the TargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttribute

`func (o *InputProcessorAnyOf1ParamsOneOf3) SetTargetAttribute(v string)`

SetTargetAttribute sets TargetAttribute field to given value.


### GetRedactSource

`func (o *InputProcessorAnyOf1ParamsOneOf3) GetRedactSource() bool`

GetRedactSource returns the RedactSource field if non-nil, zero value otherwise.

### GetRedactSourceOk

`func (o *InputProcessorAnyOf1ParamsOneOf3) GetRedactSourceOk() (*bool, bool)`

GetRedactSourceOk returns a tuple with the RedactSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactSource

`func (o *InputProcessorAnyOf1ParamsOneOf3) SetRedactSource(v bool)`

SetRedactSource sets RedactSource field to given value.


### GetSubtype

`func (o *InputProcessorAnyOf1ParamsOneOf3) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *InputProcessorAnyOf1ParamsOneOf3) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *InputProcessorAnyOf1ParamsOneOf3) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


