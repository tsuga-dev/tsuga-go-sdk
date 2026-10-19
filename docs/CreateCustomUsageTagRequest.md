# CreateCustomUsageTagRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagKey** | **string** | Telemetry tag key to track as a custom usage dimension. The key must be unique within the organization. | 

## Methods

### NewCreateCustomUsageTagRequest

`func NewCreateCustomUsageTagRequest(tagKey string, ) *CreateCustomUsageTagRequest`

NewCreateCustomUsageTagRequest instantiates a new CreateCustomUsageTagRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomUsageTagRequestWithDefaults

`func NewCreateCustomUsageTagRequestWithDefaults() *CreateCustomUsageTagRequest`

NewCreateCustomUsageTagRequestWithDefaults instantiates a new CreateCustomUsageTagRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagKey

`func (o *CreateCustomUsageTagRequest) GetTagKey() string`

GetTagKey returns the TagKey field if non-nil, zero value otherwise.

### GetTagKeyOk

`func (o *CreateCustomUsageTagRequest) GetTagKeyOk() (*string, bool)`

GetTagKeyOk returns a tuple with the TagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKey

`func (o *CreateCustomUsageTagRequest) SetTagKey(v string)`

SetTagKey sets TagKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


