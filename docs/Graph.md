# Graph

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the graph widget | 
**Name** | Pointer to **string** | Display name of the graph widget | [optional] 
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


