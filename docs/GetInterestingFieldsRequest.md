# GetInterestingFieldsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | [**[]GetInterestingFieldsRequestScopeInner**](GetInterestingFieldsRequestScopeInner.md) | Scope entries that define the population to rank fields for, drawn from positive equality pairs in a TQL query, such as service and environment. Tsuga matches entries softly by default: an artifact&#39;s contribution is weighted by the fraction of entries it matches. | 

## Methods

### NewGetInterestingFieldsRequest

`func NewGetInterestingFieldsRequest(scope []GetInterestingFieldsRequestScopeInner, ) *GetInterestingFieldsRequest`

NewGetInterestingFieldsRequest instantiates a new GetInterestingFieldsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInterestingFieldsRequestWithDefaults

`func NewGetInterestingFieldsRequestWithDefaults() *GetInterestingFieldsRequest`

NewGetInterestingFieldsRequestWithDefaults instantiates a new GetInterestingFieldsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *GetInterestingFieldsRequest) GetScope() []GetInterestingFieldsRequestScopeInner`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GetInterestingFieldsRequest) GetScopeOk() (*[]GetInterestingFieldsRequestScopeInner, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GetInterestingFieldsRequest) SetScope(v []GetInterestingFieldsRequestScopeInner)`

SetScope sets Scope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


