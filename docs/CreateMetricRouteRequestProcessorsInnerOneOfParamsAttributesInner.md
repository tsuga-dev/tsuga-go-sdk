# CreateMetricRouteRequestProcessorsInnerOneOfParamsAttributesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginAttribute** | **string** | Attribute name to map to the target attribute | 
**KeepOrigin** | Pointer to **bool** | Preserve the source attribute after mapping (defaults to false) | [optional] 
**OverrideTarget** | Pointer to **bool** | Overwrite the target attribute when it already exists (defaults to true) | [optional] 
**TargetAttribute** | **string** | Attribute name that will receive the mapped value | 

## Methods

### NewCreateMetricRouteRequestProcessorsInnerOneOfParamsAttributesInner

`func NewCreateMetricRouteRequestProcessorsInnerOneOfParamsAttributesInner(originAttribute string, targetAttribute string, ) *CreateMetricRouteRequestProcessorsInnerOneOfParamsAttributesInner`

NewCreateMetricRouteRequestProcessorsInnerOneOfParamsAttributesInner instantiates a new CreateMetricRouteRequestProcessorsInnerOneOfParamsAttributesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMetricRouteRequestProcessorsInnerOneOfParamsAttributesInnerWithDefaults

`func NewCreateMetricRouteRequestProcessorsInnerOneOfParamsAttributesInnerWithDefaults() *CreateMetricRouteRequestProcessorsInnerOneOfParamsAttributesInner`

NewCreateMetricRouteRequestProcessorsInnerOneOfParamsAttributesInnerWithDefaults instantiates a new CreateMetricRouteRequestProcessorsInnerOneOfParamsAttributesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginAttribute

`func (o *CreateMetricRouteRequestProcessorsInnerOneOfParamsAttributesInner) GetOriginAttribute() string`

GetOriginAttribute returns the OriginAttribute field if non-nil, zero value otherwise.

### GetOriginAttributeOk

`func (o *CreateMetricRouteRequestProcessorsInnerOneOfParamsAttributesInner) GetOriginAttributeOk() (*string, bool)`

GetOriginAttributeOk returns a tuple with the OriginAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAttribute

`func (o *CreateMetricRouteRequestProcessorsInnerOneOfParamsAttributesInner) SetOriginAttribute(v string)`

SetOriginAttribute sets OriginAttribute field to given value.


### GetKeepOrigin

`func (o *CreateMetricRouteRequestProcessorsInnerOneOfParamsAttributesInner) GetKeepOrigin() bool`

GetKeepOrigin returns the KeepOrigin field if non-nil, zero value otherwise.

### GetKeepOriginOk

`func (o *CreateMetricRouteRequestProcessorsInnerOneOfParamsAttributesInner) GetKeepOriginOk() (*bool, bool)`

GetKeepOriginOk returns a tuple with the KeepOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepOrigin

`func (o *CreateMetricRouteRequestProcessorsInnerOneOfParamsAttributesInner) SetKeepOrigin(v bool)`

SetKeepOrigin sets KeepOrigin field to given value.

### HasKeepOrigin

`func (o *CreateMetricRouteRequestProcessorsInnerOneOfParamsAttributesInner) HasKeepOrigin() bool`

HasKeepOrigin returns a boolean if a field has been set.

### GetOverrideTarget

`func (o *CreateMetricRouteRequestProcessorsInnerOneOfParamsAttributesInner) GetOverrideTarget() bool`

GetOverrideTarget returns the OverrideTarget field if non-nil, zero value otherwise.

### GetOverrideTargetOk

`func (o *CreateMetricRouteRequestProcessorsInnerOneOfParamsAttributesInner) GetOverrideTargetOk() (*bool, bool)`

GetOverrideTargetOk returns a tuple with the OverrideTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideTarget

`func (o *CreateMetricRouteRequestProcessorsInnerOneOfParamsAttributesInner) SetOverrideTarget(v bool)`

SetOverrideTarget sets OverrideTarget field to given value.

### HasOverrideTarget

`func (o *CreateMetricRouteRequestProcessorsInnerOneOfParamsAttributesInner) HasOverrideTarget() bool`

HasOverrideTarget returns a boolean if a field has been set.

### GetTargetAttribute

`func (o *CreateMetricRouteRequestProcessorsInnerOneOfParamsAttributesInner) GetTargetAttribute() string`

GetTargetAttribute returns the TargetAttribute field if non-nil, zero value otherwise.

### GetTargetAttributeOk

`func (o *CreateMetricRouteRequestProcessorsInnerOneOfParamsAttributesInner) GetTargetAttributeOk() (*string, bool)`

GetTargetAttributeOk returns a tuple with the TargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttribute

`func (o *CreateMetricRouteRequestProcessorsInnerOneOfParamsAttributesInner) SetTargetAttribute(v string)`

SetTargetAttribute sets TargetAttribute field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


