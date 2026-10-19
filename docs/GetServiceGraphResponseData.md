# GetServiceGraphResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** | Cluster ID the graph was built from | 
**Nodes** | [**[]GetServiceGraphResponseDataNodesInner**](GetServiceGraphResponseDataNodesInner.md) | Observed and inferred nodes included in the service dependency graph. | 
**Edges** | [**[]GetServiceGraphResponseDataEdgesInner**](GetServiceGraphResponseDataEdgesInner.md) | Directed dependency edges between graph nodes. | 

## Methods

### NewGetServiceGraphResponseData

`func NewGetServiceGraphResponseData(clusterId string, nodes []GetServiceGraphResponseDataNodesInner, edges []GetServiceGraphResponseDataEdgesInner, ) *GetServiceGraphResponseData`

NewGetServiceGraphResponseData instantiates a new GetServiceGraphResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServiceGraphResponseDataWithDefaults

`func NewGetServiceGraphResponseDataWithDefaults() *GetServiceGraphResponseData`

NewGetServiceGraphResponseDataWithDefaults instantiates a new GetServiceGraphResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *GetServiceGraphResponseData) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *GetServiceGraphResponseData) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *GetServiceGraphResponseData) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetNodes

`func (o *GetServiceGraphResponseData) GetNodes() []GetServiceGraphResponseDataNodesInner`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *GetServiceGraphResponseData) GetNodesOk() (*[]GetServiceGraphResponseDataNodesInner, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *GetServiceGraphResponseData) SetNodes(v []GetServiceGraphResponseDataNodesInner)`

SetNodes sets Nodes field to given value.


### GetEdges

`func (o *GetServiceGraphResponseData) GetEdges() []GetServiceGraphResponseDataEdgesInner`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *GetServiceGraphResponseData) GetEdgesOk() (*[]GetServiceGraphResponseDataEdgesInner, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *GetServiceGraphResponseData) SetEdges(v []GetServiceGraphResponseDataEdgesInner)`

SetEdges sets Edges field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


