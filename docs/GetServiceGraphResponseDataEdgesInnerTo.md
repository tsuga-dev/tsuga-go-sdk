# GetServiceGraphResponseDataEdgesInnerTo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Env** | Pointer to **string** | Environment of the target node | [optional] 
**Name** | **string** | Name of the target node | 
**DatabaseSystem** | Pointer to **string** | Database system of the target when it is an inferred database node (&#x60;inferred_db&#x60;). Together with &#x60;name&#x60; and &#x60;env&#x60; it identifies the target node, distinguishing a database from a same-named service or a same-named database on another system. | [optional] 

## Methods

### NewGetServiceGraphResponseDataEdgesInnerTo

`func NewGetServiceGraphResponseDataEdgesInnerTo(name string, ) *GetServiceGraphResponseDataEdgesInnerTo`

NewGetServiceGraphResponseDataEdgesInnerTo instantiates a new GetServiceGraphResponseDataEdgesInnerTo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServiceGraphResponseDataEdgesInnerToWithDefaults

`func NewGetServiceGraphResponseDataEdgesInnerToWithDefaults() *GetServiceGraphResponseDataEdgesInnerTo`

NewGetServiceGraphResponseDataEdgesInnerToWithDefaults instantiates a new GetServiceGraphResponseDataEdgesInnerTo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnv

`func (o *GetServiceGraphResponseDataEdgesInnerTo) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *GetServiceGraphResponseDataEdgesInnerTo) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *GetServiceGraphResponseDataEdgesInnerTo) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *GetServiceGraphResponseDataEdgesInnerTo) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetName

`func (o *GetServiceGraphResponseDataEdgesInnerTo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetServiceGraphResponseDataEdgesInnerTo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetServiceGraphResponseDataEdgesInnerTo) SetName(v string)`

SetName sets Name field to given value.


### GetDatabaseSystem

`func (o *GetServiceGraphResponseDataEdgesInnerTo) GetDatabaseSystem() string`

GetDatabaseSystem returns the DatabaseSystem field if non-nil, zero value otherwise.

### GetDatabaseSystemOk

`func (o *GetServiceGraphResponseDataEdgesInnerTo) GetDatabaseSystemOk() (*string, bool)`

GetDatabaseSystemOk returns a tuple with the DatabaseSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseSystem

`func (o *GetServiceGraphResponseDataEdgesInnerTo) SetDatabaseSystem(v string)`

SetDatabaseSystem sets DatabaseSystem field to given value.

### HasDatabaseSystem

`func (o *GetServiceGraphResponseDataEdgesInnerTo) HasDatabaseSystem() bool`

HasDatabaseSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


