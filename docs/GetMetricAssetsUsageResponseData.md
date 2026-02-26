# GetMetricAssetsUsageResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | [**[]GetMetricAssetsUsageResponseDataAssetsInner**](GetMetricAssetsUsageResponseDataAssetsInner.md) | List of assets (dashboards and monitors) using this metric | 

## Methods

### NewGetMetricAssetsUsageResponseData

`func NewGetMetricAssetsUsageResponseData(assets []GetMetricAssetsUsageResponseDataAssetsInner, ) *GetMetricAssetsUsageResponseData`

NewGetMetricAssetsUsageResponseData instantiates a new GetMetricAssetsUsageResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMetricAssetsUsageResponseDataWithDefaults

`func NewGetMetricAssetsUsageResponseDataWithDefaults() *GetMetricAssetsUsageResponseData`

NewGetMetricAssetsUsageResponseDataWithDefaults instantiates a new GetMetricAssetsUsageResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *GetMetricAssetsUsageResponseData) GetAssets() []GetMetricAssetsUsageResponseDataAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *GetMetricAssetsUsageResponseData) GetAssetsOk() (*[]GetMetricAssetsUsageResponseDataAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *GetMetricAssetsUsageResponseData) SetAssets(v []GetMetricAssetsUsageResponseDataAssetsInner)`

SetAssets sets Assets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


