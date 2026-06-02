# InputGraphVisualizationListLogPatterns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Displays log patterns clustered from logs matching the query | 
**Query** | **string** | Tsuga query that selects logs to cluster into patterns | 
**Layout** | Pointer to **string** | Layout used to render log patterns | [optional] 

## Methods

### NewInputGraphVisualizationListLogPatterns

`func NewInputGraphVisualizationListLogPatterns(type_ string, query string, ) *InputGraphVisualizationListLogPatterns`

NewInputGraphVisualizationListLogPatterns instantiates a new InputGraphVisualizationListLogPatterns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputGraphVisualizationListLogPatternsWithDefaults

`func NewInputGraphVisualizationListLogPatternsWithDefaults() *InputGraphVisualizationListLogPatterns`

NewInputGraphVisualizationListLogPatternsWithDefaults instantiates a new InputGraphVisualizationListLogPatterns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputGraphVisualizationListLogPatterns) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputGraphVisualizationListLogPatterns) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputGraphVisualizationListLogPatterns) SetType(v string)`

SetType sets Type field to given value.


### GetQuery

`func (o *InputGraphVisualizationListLogPatterns) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *InputGraphVisualizationListLogPatterns) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *InputGraphVisualizationListLogPatterns) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetLayout

`func (o *InputGraphVisualizationListLogPatterns) GetLayout() string`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *InputGraphVisualizationListLogPatterns) GetLayoutOk() (*string, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *InputGraphVisualizationListLogPatterns) SetLayout(v string)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *InputGraphVisualizationListLogPatterns) HasLayout() bool`

HasLayout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


