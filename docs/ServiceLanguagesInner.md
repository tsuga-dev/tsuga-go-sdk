# ServiceLanguagesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Language** | **string** | Programming language observed for this service. The OpenTelemetry SDK sets it through telemetry metadata. Returned when language metadata is available for the service. | 
**LastSeenAt** | **string** | Most recent time this language was observed. Tsuga sets it from telemetry timestamps during service metadata refreshes. | 

## Methods

### NewServiceLanguagesInner

`func NewServiceLanguagesInner(language string, lastSeenAt string, ) *ServiceLanguagesInner`

NewServiceLanguagesInner instantiates a new ServiceLanguagesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceLanguagesInnerWithDefaults

`func NewServiceLanguagesInnerWithDefaults() *ServiceLanguagesInner`

NewServiceLanguagesInnerWithDefaults instantiates a new ServiceLanguagesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguage

`func (o *ServiceLanguagesInner) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ServiceLanguagesInner) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ServiceLanguagesInner) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetLastSeenAt

`func (o *ServiceLanguagesInner) GetLastSeenAt() string`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *ServiceLanguagesInner) GetLastSeenAtOk() (*string, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *ServiceLanguagesInner) SetLastSeenAt(v string)`

SetLastSeenAt sets LastSeenAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


