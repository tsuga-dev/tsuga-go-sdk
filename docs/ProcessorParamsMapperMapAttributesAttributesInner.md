# ProcessorParamsMapperMapAttributesAttributesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginAttribute** | **string** | Attribute name to map to the target attribute | 
**KeepOrigin** | Pointer to **bool** | Preserve the source attribute after mapping (defaults to false) | [optional] 
**OverrideTarget** | Pointer to **bool** | Overwrite the target attribute when it already exists (defaults to true) | [optional] 
**TargetAttribute** | **string** | Attribute name that will receive the mapped value | 

## Methods

### NewProcessorParamsMapperMapAttributesAttributesInner

`func NewProcessorParamsMapperMapAttributesAttributesInner(originAttribute string, targetAttribute string, ) *ProcessorParamsMapperMapAttributesAttributesInner`

NewProcessorParamsMapperMapAttributesAttributesInner instantiates a new ProcessorParamsMapperMapAttributesAttributesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorParamsMapperMapAttributesAttributesInnerWithDefaults

`func NewProcessorParamsMapperMapAttributesAttributesInnerWithDefaults() *ProcessorParamsMapperMapAttributesAttributesInner`

NewProcessorParamsMapperMapAttributesAttributesInnerWithDefaults instantiates a new ProcessorParamsMapperMapAttributesAttributesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginAttribute

`func (o *ProcessorParamsMapperMapAttributesAttributesInner) GetOriginAttribute() string`

GetOriginAttribute returns the OriginAttribute field if non-nil, zero value otherwise.

### GetOriginAttributeOk

`func (o *ProcessorParamsMapperMapAttributesAttributesInner) GetOriginAttributeOk() (*string, bool)`

GetOriginAttributeOk returns a tuple with the OriginAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAttribute

`func (o *ProcessorParamsMapperMapAttributesAttributesInner) SetOriginAttribute(v string)`

SetOriginAttribute sets OriginAttribute field to given value.


### GetKeepOrigin

`func (o *ProcessorParamsMapperMapAttributesAttributesInner) GetKeepOrigin() bool`

GetKeepOrigin returns the KeepOrigin field if non-nil, zero value otherwise.

### GetKeepOriginOk

`func (o *ProcessorParamsMapperMapAttributesAttributesInner) GetKeepOriginOk() (*bool, bool)`

GetKeepOriginOk returns a tuple with the KeepOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepOrigin

`func (o *ProcessorParamsMapperMapAttributesAttributesInner) SetKeepOrigin(v bool)`

SetKeepOrigin sets KeepOrigin field to given value.

### HasKeepOrigin

`func (o *ProcessorParamsMapperMapAttributesAttributesInner) HasKeepOrigin() bool`

HasKeepOrigin returns a boolean if a field has been set.

### GetOverrideTarget

`func (o *ProcessorParamsMapperMapAttributesAttributesInner) GetOverrideTarget() bool`

GetOverrideTarget returns the OverrideTarget field if non-nil, zero value otherwise.

### GetOverrideTargetOk

`func (o *ProcessorParamsMapperMapAttributesAttributesInner) GetOverrideTargetOk() (*bool, bool)`

GetOverrideTargetOk returns a tuple with the OverrideTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideTarget

`func (o *ProcessorParamsMapperMapAttributesAttributesInner) SetOverrideTarget(v bool)`

SetOverrideTarget sets OverrideTarget field to given value.

### HasOverrideTarget

`func (o *ProcessorParamsMapperMapAttributesAttributesInner) HasOverrideTarget() bool`

HasOverrideTarget returns a boolean if a field has been set.

### GetTargetAttribute

`func (o *ProcessorParamsMapperMapAttributesAttributesInner) GetTargetAttribute() string`

GetTargetAttribute returns the TargetAttribute field if non-nil, zero value otherwise.

### GetTargetAttributeOk

`func (o *ProcessorParamsMapperMapAttributesAttributesInner) GetTargetAttributeOk() (*string, bool)`

GetTargetAttributeOk returns a tuple with the TargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttribute

`func (o *ProcessorParamsMapperMapAttributesAttributesInner) SetTargetAttribute(v string)`

SetTargetAttribute sets TargetAttribute field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


