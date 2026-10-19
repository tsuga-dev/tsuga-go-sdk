# CreateFeedbackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | Free-text feedback content, e.g. what went wrong, what was confusing, or a suggestion. Required and must be non-empty. | 

## Methods

### NewCreateFeedbackRequest

`func NewCreateFeedbackRequest(content string, ) *CreateFeedbackRequest`

NewCreateFeedbackRequest instantiates a new CreateFeedbackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFeedbackRequestWithDefaults

`func NewCreateFeedbackRequestWithDefaults() *CreateFeedbackRequest`

NewCreateFeedbackRequestWithDefaults instantiates a new CreateFeedbackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CreateFeedbackRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CreateFeedbackRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CreateFeedbackRequest) SetContent(v string)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


