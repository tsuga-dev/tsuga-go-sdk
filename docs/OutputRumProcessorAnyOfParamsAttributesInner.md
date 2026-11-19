# OutputRumProcessorAnyOfParamsAttributesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginAttribute** | **string** | Attribute name to map to the target attribute | 
**KeepOrigin** | Pointer to **bool** | Whether the source attribute is preserved after mapping. | [optional] 
**OverrideTarget** | Pointer to **bool** | Whether to overwrite the target attribute when it already exists. | [optional] 
**TargetAttribute** | **string** | Attribute name that will receive the mapped value | 

## Methods

### NewOutputRumProcessorAnyOfParamsAttributesInner

`func NewOutputRumProcessorAnyOfParamsAttributesInner(originAttribute string, targetAttribute string, ) *OutputRumProcessorAnyOfParamsAttributesInner`

NewOutputRumProcessorAnyOfParamsAttributesInner instantiates a new OutputRumProcessorAnyOfParamsAttributesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputRumProcessorAnyOfParamsAttributesInnerWithDefaults

`func NewOutputRumProcessorAnyOfParamsAttributesInnerWithDefaults() *OutputRumProcessorAnyOfParamsAttributesInner`

NewOutputRumProcessorAnyOfParamsAttributesInnerWithDefaults instantiates a new OutputRumProcessorAnyOfParamsAttributesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginAttribute

`func (o *OutputRumProcessorAnyOfParamsAttributesInner) GetOriginAttribute() string`

GetOriginAttribute returns the OriginAttribute field if non-nil, zero value otherwise.

### GetOriginAttributeOk

`func (o *OutputRumProcessorAnyOfParamsAttributesInner) GetOriginAttributeOk() (*string, bool)`

GetOriginAttributeOk returns a tuple with the OriginAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAttribute

`func (o *OutputRumProcessorAnyOfParamsAttributesInner) SetOriginAttribute(v string)`

SetOriginAttribute sets OriginAttribute field to given value.


### GetKeepOrigin

`func (o *OutputRumProcessorAnyOfParamsAttributesInner) GetKeepOrigin() bool`

GetKeepOrigin returns the KeepOrigin field if non-nil, zero value otherwise.

### GetKeepOriginOk

`func (o *OutputRumProcessorAnyOfParamsAttributesInner) GetKeepOriginOk() (*bool, bool)`

GetKeepOriginOk returns a tuple with the KeepOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepOrigin

`func (o *OutputRumProcessorAnyOfParamsAttributesInner) SetKeepOrigin(v bool)`

SetKeepOrigin sets KeepOrigin field to given value.

### HasKeepOrigin

`func (o *OutputRumProcessorAnyOfParamsAttributesInner) HasKeepOrigin() bool`

HasKeepOrigin returns a boolean if a field has been set.

### GetOverrideTarget

`func (o *OutputRumProcessorAnyOfParamsAttributesInner) GetOverrideTarget() bool`

GetOverrideTarget returns the OverrideTarget field if non-nil, zero value otherwise.

### GetOverrideTargetOk

`func (o *OutputRumProcessorAnyOfParamsAttributesInner) GetOverrideTargetOk() (*bool, bool)`

GetOverrideTargetOk returns a tuple with the OverrideTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideTarget

`func (o *OutputRumProcessorAnyOfParamsAttributesInner) SetOverrideTarget(v bool)`

SetOverrideTarget sets OverrideTarget field to given value.

### HasOverrideTarget

`func (o *OutputRumProcessorAnyOfParamsAttributesInner) HasOverrideTarget() bool`

HasOverrideTarget returns a boolean if a field has been set.

### GetTargetAttribute

`func (o *OutputRumProcessorAnyOfParamsAttributesInner) GetTargetAttribute() string`

GetTargetAttribute returns the TargetAttribute field if non-nil, zero value otherwise.

### GetTargetAttributeOk

`func (o *OutputRumProcessorAnyOfParamsAttributesInner) GetTargetAttributeOk() (*string, bool)`

GetTargetAttributeOk returns a tuple with the TargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttribute

`func (o *OutputRumProcessorAnyOfParamsAttributesInner) SetTargetAttribute(v string)`

SetTargetAttribute sets TargetAttribute field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


