# ListDashboards200ResponseDataInnerGraphsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Stable identifier of the graph widget | 
**Name** | Pointer to **string** | Display name of the graph widget | [optional] 
**Visualization** | [**ListDashboards200ResponseDataInnerGraphsInnerVisualization**](ListDashboards200ResponseDataInnerGraphsInnerVisualization.md) |  | 
**Layout** | Pointer to [**ListDashboards200ResponseDataInnerGraphsInnerLayout**](ListDashboards200ResponseDataInnerGraphsInnerLayout.md) |  | [optional] 

## Methods

### NewListDashboards200ResponseDataInnerGraphsInner

`func NewListDashboards200ResponseDataInnerGraphsInner(id string, visualization ListDashboards200ResponseDataInnerGraphsInnerVisualization, ) *ListDashboards200ResponseDataInnerGraphsInner`

NewListDashboards200ResponseDataInnerGraphsInner instantiates a new ListDashboards200ResponseDataInnerGraphsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDashboards200ResponseDataInnerGraphsInnerWithDefaults

`func NewListDashboards200ResponseDataInnerGraphsInnerWithDefaults() *ListDashboards200ResponseDataInnerGraphsInner`

NewListDashboards200ResponseDataInnerGraphsInnerWithDefaults instantiates a new ListDashboards200ResponseDataInnerGraphsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListDashboards200ResponseDataInnerGraphsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListDashboards200ResponseDataInnerGraphsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListDashboards200ResponseDataInnerGraphsInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ListDashboards200ResponseDataInnerGraphsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListDashboards200ResponseDataInnerGraphsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListDashboards200ResponseDataInnerGraphsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListDashboards200ResponseDataInnerGraphsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVisualization

`func (o *ListDashboards200ResponseDataInnerGraphsInner) GetVisualization() ListDashboards200ResponseDataInnerGraphsInnerVisualization`

GetVisualization returns the Visualization field if non-nil, zero value otherwise.

### GetVisualizationOk

`func (o *ListDashboards200ResponseDataInnerGraphsInner) GetVisualizationOk() (*ListDashboards200ResponseDataInnerGraphsInnerVisualization, bool)`

GetVisualizationOk returns a tuple with the Visualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualization

`func (o *ListDashboards200ResponseDataInnerGraphsInner) SetVisualization(v ListDashboards200ResponseDataInnerGraphsInnerVisualization)`

SetVisualization sets Visualization field to given value.


### GetLayout

`func (o *ListDashboards200ResponseDataInnerGraphsInner) GetLayout() ListDashboards200ResponseDataInnerGraphsInnerLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *ListDashboards200ResponseDataInnerGraphsInner) GetLayoutOk() (*ListDashboards200ResponseDataInnerGraphsInnerLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *ListDashboards200ResponseDataInnerGraphsInner) SetLayout(v ListDashboards200ResponseDataInnerGraphsInnerLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *ListDashboards200ResponseDataInnerGraphsInner) HasLayout() bool`

HasLayout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


