# Graph

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the graph widget inside the dashboard. Set by the caller when the graph is created; use it to target the graph in later update requests. | 
**Name** | Pointer to **string** | Display name of the graph widget. | [optional] 
**Description** | Pointer to **string** | Optional text shown with the graph widget. Maximum length is 800 characters. | [optional] 
**DescriptionAlign** | Pointer to **string** | Flex alignment keyword used for widget layout | [optional] 
**DescriptionJustifyContent** | Pointer to **string** | Flex alignment keyword used for widget layout | [optional] 
**Visualization** | [**GraphVisualization**](GraphVisualization.md) |  | 
**Layout** | Pointer to [**GraphLayout**](GraphLayout.md) |  | [optional] 

## Methods

### NewGraph

`func NewGraph(id string, visualization GraphVisualization, ) *Graph`

NewGraph instantiates a new Graph object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphWithDefaults

`func NewGraphWithDefaults() *Graph`

NewGraphWithDefaults instantiates a new Graph object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Graph) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Graph) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Graph) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Graph) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Graph) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Graph) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Graph) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Graph) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Graph) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Graph) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Graph) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescriptionAlign

`func (o *Graph) GetDescriptionAlign() string`

GetDescriptionAlign returns the DescriptionAlign field if non-nil, zero value otherwise.

### GetDescriptionAlignOk

`func (o *Graph) GetDescriptionAlignOk() (*string, bool)`

GetDescriptionAlignOk returns a tuple with the DescriptionAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionAlign

`func (o *Graph) SetDescriptionAlign(v string)`

SetDescriptionAlign sets DescriptionAlign field to given value.

### HasDescriptionAlign

`func (o *Graph) HasDescriptionAlign() bool`

HasDescriptionAlign returns a boolean if a field has been set.

### GetDescriptionJustifyContent

`func (o *Graph) GetDescriptionJustifyContent() string`

GetDescriptionJustifyContent returns the DescriptionJustifyContent field if non-nil, zero value otherwise.

### GetDescriptionJustifyContentOk

`func (o *Graph) GetDescriptionJustifyContentOk() (*string, bool)`

GetDescriptionJustifyContentOk returns a tuple with the DescriptionJustifyContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionJustifyContent

`func (o *Graph) SetDescriptionJustifyContent(v string)`

SetDescriptionJustifyContent sets DescriptionJustifyContent field to given value.

### HasDescriptionJustifyContent

`func (o *Graph) HasDescriptionJustifyContent() bool`

HasDescriptionJustifyContent returns a boolean if a field has been set.

### GetVisualization

`func (o *Graph) GetVisualization() GraphVisualization`

GetVisualization returns the Visualization field if non-nil, zero value otherwise.

### GetVisualizationOk

`func (o *Graph) GetVisualizationOk() (*GraphVisualization, bool)`

GetVisualizationOk returns a tuple with the Visualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualization

`func (o *Graph) SetVisualization(v GraphVisualization)`

SetVisualization sets Visualization field to given value.


### GetLayout

`func (o *Graph) GetLayout() GraphLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *Graph) GetLayoutOk() (*GraphLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *Graph) SetLayout(v GraphLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *Graph) HasLayout() bool`

HasLayout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


