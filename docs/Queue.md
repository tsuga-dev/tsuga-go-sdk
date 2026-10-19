# Queue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Env** | **string** |  | 
**MessagingSystem** | **string** |  | 
**PublishPerSecond** | **float32** | Messages published per second. | 
**ConsumePerSecond** | **float32** | Messages consumed per second. | 
**ConsumerErrorsPerSecond** | **float32** | Messages per second whose consumer span has an error status. | 
**P90ProcessingMs** | **float32** | p90 duration of consumer spans. | 

## Methods

### NewQueue

`func NewQueue(name string, env string, messagingSystem string, publishPerSecond float32, consumePerSecond float32, consumerErrorsPerSecond float32, p90ProcessingMs float32, ) *Queue`

NewQueue instantiates a new Queue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueWithDefaults

`func NewQueueWithDefaults() *Queue`

NewQueueWithDefaults instantiates a new Queue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Queue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Queue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Queue) SetName(v string)`

SetName sets Name field to given value.


### GetEnv

`func (o *Queue) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *Queue) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *Queue) SetEnv(v string)`

SetEnv sets Env field to given value.


### GetMessagingSystem

`func (o *Queue) GetMessagingSystem() string`

GetMessagingSystem returns the MessagingSystem field if non-nil, zero value otherwise.

### GetMessagingSystemOk

`func (o *Queue) GetMessagingSystemOk() (*string, bool)`

GetMessagingSystemOk returns a tuple with the MessagingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingSystem

`func (o *Queue) SetMessagingSystem(v string)`

SetMessagingSystem sets MessagingSystem field to given value.


### GetPublishPerSecond

`func (o *Queue) GetPublishPerSecond() float32`

GetPublishPerSecond returns the PublishPerSecond field if non-nil, zero value otherwise.

### GetPublishPerSecondOk

`func (o *Queue) GetPublishPerSecondOk() (*float32, bool)`

GetPublishPerSecondOk returns a tuple with the PublishPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishPerSecond

`func (o *Queue) SetPublishPerSecond(v float32)`

SetPublishPerSecond sets PublishPerSecond field to given value.


### GetConsumePerSecond

`func (o *Queue) GetConsumePerSecond() float32`

GetConsumePerSecond returns the ConsumePerSecond field if non-nil, zero value otherwise.

### GetConsumePerSecondOk

`func (o *Queue) GetConsumePerSecondOk() (*float32, bool)`

GetConsumePerSecondOk returns a tuple with the ConsumePerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumePerSecond

`func (o *Queue) SetConsumePerSecond(v float32)`

SetConsumePerSecond sets ConsumePerSecond field to given value.


### GetConsumerErrorsPerSecond

`func (o *Queue) GetConsumerErrorsPerSecond() float32`

GetConsumerErrorsPerSecond returns the ConsumerErrorsPerSecond field if non-nil, zero value otherwise.

### GetConsumerErrorsPerSecondOk

`func (o *Queue) GetConsumerErrorsPerSecondOk() (*float32, bool)`

GetConsumerErrorsPerSecondOk returns a tuple with the ConsumerErrorsPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerErrorsPerSecond

`func (o *Queue) SetConsumerErrorsPerSecond(v float32)`

SetConsumerErrorsPerSecond sets ConsumerErrorsPerSecond field to given value.


### GetP90ProcessingMs

`func (o *Queue) GetP90ProcessingMs() float32`

GetP90ProcessingMs returns the P90ProcessingMs field if non-nil, zero value otherwise.

### GetP90ProcessingMsOk

`func (o *Queue) GetP90ProcessingMsOk() (*float32, bool)`

GetP90ProcessingMsOk returns a tuple with the P90ProcessingMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP90ProcessingMs

`func (o *Queue) SetP90ProcessingMs(v float32)`

SetP90ProcessingMs sets P90ProcessingMs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


