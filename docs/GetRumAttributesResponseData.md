# GetRumAttributesResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | **[]string** | Every RUM attribute dot-path that can be queried. Use these names as-is in RUM search queries and RUM aggregations. | 

## Methods

### NewGetRumAttributesResponseData

`func NewGetRumAttributesResponseData(attributes []string, ) *GetRumAttributesResponseData`

NewGetRumAttributesResponseData instantiates a new GetRumAttributesResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRumAttributesResponseDataWithDefaults

`func NewGetRumAttributesResponseDataWithDefaults() *GetRumAttributesResponseData`

NewGetRumAttributesResponseDataWithDefaults instantiates a new GetRumAttributesResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *GetRumAttributesResponseData) GetAttributes() []string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GetRumAttributesResponseData) GetAttributesOk() (*[]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GetRumAttributesResponseData) SetAttributes(v []string)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


