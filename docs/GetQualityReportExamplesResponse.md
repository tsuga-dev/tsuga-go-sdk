# GetQualityReportExamplesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Identifier used to trace the lifecycle of this API request | 
**Data** | [**GetQualityReportExamplesResponseData**](GetQualityReportExamplesResponseData.md) |  | 
**Metadata** | Pointer to [**ResponseMetadata**](ResponseMetadata.md) |  | [optional] 

## Methods

### NewGetQualityReportExamplesResponse

`func NewGetQualityReportExamplesResponse(requestId string, data GetQualityReportExamplesResponseData, ) *GetQualityReportExamplesResponse`

NewGetQualityReportExamplesResponse instantiates a new GetQualityReportExamplesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetQualityReportExamplesResponseWithDefaults

`func NewGetQualityReportExamplesResponseWithDefaults() *GetQualityReportExamplesResponse`

NewGetQualityReportExamplesResponseWithDefaults instantiates a new GetQualityReportExamplesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *GetQualityReportExamplesResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetQualityReportExamplesResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetQualityReportExamplesResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetData

`func (o *GetQualityReportExamplesResponse) GetData() GetQualityReportExamplesResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetQualityReportExamplesResponse) GetDataOk() (*GetQualityReportExamplesResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetQualityReportExamplesResponse) SetData(v GetQualityReportExamplesResponseData)`

SetData sets Data field to given value.


### GetMetadata

`func (o *GetQualityReportExamplesResponse) GetMetadata() ResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetQualityReportExamplesResponse) GetMetadataOk() (*ResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetQualityReportExamplesResponse) SetMetadata(v ResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetQualityReportExamplesResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


