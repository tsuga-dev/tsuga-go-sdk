# UpdateDashboardGraphRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**DescriptionAlign** | Pointer to **NullableString** | Flex alignment keyword used for widget layout | [optional] 
**DescriptionJustifyContent** | Pointer to **NullableString** | Flex alignment keyword used for widget layout | [optional] 
**Visualization** | Pointer to [**GraphVisualization**](GraphVisualization.md) |  | [optional] 
**Layout** | Pointer to [**NullableGraphLayout**](GraphLayout.md) |  | [optional] 

## Methods

### NewUpdateDashboardGraphRequest

`func NewUpdateDashboardGraphRequest() *UpdateDashboardGraphRequest`

NewUpdateDashboardGraphRequest instantiates a new UpdateDashboardGraphRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDashboardGraphRequestWithDefaults

`func NewUpdateDashboardGraphRequestWithDefaults() *UpdateDashboardGraphRequest`

NewUpdateDashboardGraphRequestWithDefaults instantiates a new UpdateDashboardGraphRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateDashboardGraphRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDashboardGraphRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDashboardGraphRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateDashboardGraphRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateDashboardGraphRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateDashboardGraphRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *UpdateDashboardGraphRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateDashboardGraphRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateDashboardGraphRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateDashboardGraphRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateDashboardGraphRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateDashboardGraphRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDescriptionAlign

`func (o *UpdateDashboardGraphRequest) GetDescriptionAlign() string`

GetDescriptionAlign returns the DescriptionAlign field if non-nil, zero value otherwise.

### GetDescriptionAlignOk

`func (o *UpdateDashboardGraphRequest) GetDescriptionAlignOk() (*string, bool)`

GetDescriptionAlignOk returns a tuple with the DescriptionAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionAlign

`func (o *UpdateDashboardGraphRequest) SetDescriptionAlign(v string)`

SetDescriptionAlign sets DescriptionAlign field to given value.

### HasDescriptionAlign

`func (o *UpdateDashboardGraphRequest) HasDescriptionAlign() bool`

HasDescriptionAlign returns a boolean if a field has been set.

### SetDescriptionAlignNil

`func (o *UpdateDashboardGraphRequest) SetDescriptionAlignNil(b bool)`

 SetDescriptionAlignNil sets the value for DescriptionAlign to be an explicit nil

### UnsetDescriptionAlign
`func (o *UpdateDashboardGraphRequest) UnsetDescriptionAlign()`

UnsetDescriptionAlign ensures that no value is present for DescriptionAlign, not even an explicit nil
### GetDescriptionJustifyContent

`func (o *UpdateDashboardGraphRequest) GetDescriptionJustifyContent() string`

GetDescriptionJustifyContent returns the DescriptionJustifyContent field if non-nil, zero value otherwise.

### GetDescriptionJustifyContentOk

`func (o *UpdateDashboardGraphRequest) GetDescriptionJustifyContentOk() (*string, bool)`

GetDescriptionJustifyContentOk returns a tuple with the DescriptionJustifyContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionJustifyContent

`func (o *UpdateDashboardGraphRequest) SetDescriptionJustifyContent(v string)`

SetDescriptionJustifyContent sets DescriptionJustifyContent field to given value.

### HasDescriptionJustifyContent

`func (o *UpdateDashboardGraphRequest) HasDescriptionJustifyContent() bool`

HasDescriptionJustifyContent returns a boolean if a field has been set.

### SetDescriptionJustifyContentNil

`func (o *UpdateDashboardGraphRequest) SetDescriptionJustifyContentNil(b bool)`

 SetDescriptionJustifyContentNil sets the value for DescriptionJustifyContent to be an explicit nil

### UnsetDescriptionJustifyContent
`func (o *UpdateDashboardGraphRequest) UnsetDescriptionJustifyContent()`

UnsetDescriptionJustifyContent ensures that no value is present for DescriptionJustifyContent, not even an explicit nil
### GetVisualization

`func (o *UpdateDashboardGraphRequest) GetVisualization() GraphVisualization`

GetVisualization returns the Visualization field if non-nil, zero value otherwise.

### GetVisualizationOk

`func (o *UpdateDashboardGraphRequest) GetVisualizationOk() (*GraphVisualization, bool)`

GetVisualizationOk returns a tuple with the Visualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualization

`func (o *UpdateDashboardGraphRequest) SetVisualization(v GraphVisualization)`

SetVisualization sets Visualization field to given value.

### HasVisualization

`func (o *UpdateDashboardGraphRequest) HasVisualization() bool`

HasVisualization returns a boolean if a field has been set.

### GetLayout

`func (o *UpdateDashboardGraphRequest) GetLayout() GraphLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *UpdateDashboardGraphRequest) GetLayoutOk() (*GraphLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *UpdateDashboardGraphRequest) SetLayout(v GraphLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *UpdateDashboardGraphRequest) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### SetLayoutNil

`func (o *UpdateDashboardGraphRequest) SetLayoutNil(b bool)`

 SetLayoutNil sets the value for Layout to be an explicit nil

### UnsetLayout
`func (o *UpdateDashboardGraphRequest) UnsetLayout()`

UnsetLayout ensures that no value is present for Layout, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


