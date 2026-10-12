# InputGraph

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the graph widget inside the dashboard. Assign it when the graph is created; use it to target the graph in later update requests. | 
**Name** | Pointer to **string** | Display name of the graph widget. | [optional] 
**Description** | Pointer to **string** | Text shown with the graph widget. | [optional] 
**DescriptionAlign** | Pointer to **string** | Flex alignment keyword used for widget layout | [optional] 
**DescriptionJustifyContent** | Pointer to **string** | Flex alignment keyword used for widget layout | [optional] 
**Visualization** | [**UpdateDashboardGraphRequestVisualization**](UpdateDashboardGraphRequestVisualization.md) |  | 
**Layout** | Pointer to [**GraphLayout**](GraphLayout.md) |  | [optional] 

## Methods

### NewInputGraph

`func NewInputGraph(id string, visualization UpdateDashboardGraphRequestVisualization, ) *InputGraph`

NewInputGraph instantiates a new InputGraph object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputGraphWithDefaults

`func NewInputGraphWithDefaults() *InputGraph`

NewInputGraphWithDefaults instantiates a new InputGraph object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InputGraph) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InputGraph) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InputGraph) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *InputGraph) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputGraph) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputGraph) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InputGraph) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InputGraph) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputGraph) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputGraph) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InputGraph) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescriptionAlign

`func (o *InputGraph) GetDescriptionAlign() string`

GetDescriptionAlign returns the DescriptionAlign field if non-nil, zero value otherwise.

### GetDescriptionAlignOk

`func (o *InputGraph) GetDescriptionAlignOk() (*string, bool)`

GetDescriptionAlignOk returns a tuple with the DescriptionAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionAlign

`func (o *InputGraph) SetDescriptionAlign(v string)`

SetDescriptionAlign sets DescriptionAlign field to given value.

### HasDescriptionAlign

`func (o *InputGraph) HasDescriptionAlign() bool`

HasDescriptionAlign returns a boolean if a field has been set.

### GetDescriptionJustifyContent

`func (o *InputGraph) GetDescriptionJustifyContent() string`

GetDescriptionJustifyContent returns the DescriptionJustifyContent field if non-nil, zero value otherwise.

### GetDescriptionJustifyContentOk

`func (o *InputGraph) GetDescriptionJustifyContentOk() (*string, bool)`

GetDescriptionJustifyContentOk returns a tuple with the DescriptionJustifyContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionJustifyContent

`func (o *InputGraph) SetDescriptionJustifyContent(v string)`

SetDescriptionJustifyContent sets DescriptionJustifyContent field to given value.

### HasDescriptionJustifyContent

`func (o *InputGraph) HasDescriptionJustifyContent() bool`

HasDescriptionJustifyContent returns a boolean if a field has been set.

### GetVisualization

`func (o *InputGraph) GetVisualization() UpdateDashboardGraphRequestVisualization`

GetVisualization returns the Visualization field if non-nil, zero value otherwise.

### GetVisualizationOk

`func (o *InputGraph) GetVisualizationOk() (*UpdateDashboardGraphRequestVisualization, bool)`

GetVisualizationOk returns a tuple with the Visualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualization

`func (o *InputGraph) SetVisualization(v UpdateDashboardGraphRequestVisualization)`

SetVisualization sets Visualization field to given value.


### GetLayout

`func (o *InputGraph) GetLayout() GraphLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *InputGraph) GetLayoutOk() (*GraphLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *InputGraph) SetLayout(v GraphLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *InputGraph) HasLayout() bool`

HasLayout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


