# Graph1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the graph widget | 
**Name** | Pointer to **string** | Display name of the graph widget | [optional] 
**Description** | Pointer to **string** | Description of the graph widget | [optional] 
**Visualization** | [**Graph1Visualization**](Graph1Visualization.md) |  | 
**Layout** | Pointer to [**GraphLayout**](GraphLayout.md) |  | [optional] 

## Methods

### NewGraph1

`func NewGraph1(id string, visualization Graph1Visualization, ) *Graph1`

NewGraph1 instantiates a new Graph1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraph1WithDefaults

`func NewGraph1WithDefaults() *Graph1`

NewGraph1WithDefaults instantiates a new Graph1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Graph1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Graph1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Graph1) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Graph1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Graph1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Graph1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Graph1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Graph1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Graph1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Graph1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Graph1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVisualization

`func (o *Graph1) GetVisualization() Graph1Visualization`

GetVisualization returns the Visualization field if non-nil, zero value otherwise.

### GetVisualizationOk

`func (o *Graph1) GetVisualizationOk() (*Graph1Visualization, bool)`

GetVisualizationOk returns a tuple with the Visualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualization

`func (o *Graph1) SetVisualization(v Graph1Visualization)`

SetVisualization sets Visualization field to given value.


### GetLayout

`func (o *Graph1) GetLayout() GraphLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *Graph1) GetLayoutOk() (*GraphLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *Graph1) SetLayout(v GraphLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *Graph1) HasLayout() bool`

HasLayout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


