# ServiceVersionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** |  | 
**FirstSeenAt** | **time.Time** |  | 
**LastSeenAt** | **time.Time** |  | 
**Faulty** | Pointer to **bool** |  | [optional] 

## Methods

### NewServiceVersionsInner

`func NewServiceVersionsInner(version string, firstSeenAt time.Time, lastSeenAt time.Time, ) *ServiceVersionsInner`

NewServiceVersionsInner instantiates a new ServiceVersionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceVersionsInnerWithDefaults

`func NewServiceVersionsInnerWithDefaults() *ServiceVersionsInner`

NewServiceVersionsInnerWithDefaults instantiates a new ServiceVersionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ServiceVersionsInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ServiceVersionsInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ServiceVersionsInner) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetFirstSeenAt

`func (o *ServiceVersionsInner) GetFirstSeenAt() time.Time`

GetFirstSeenAt returns the FirstSeenAt field if non-nil, zero value otherwise.

### GetFirstSeenAtOk

`func (o *ServiceVersionsInner) GetFirstSeenAtOk() (*time.Time, bool)`

GetFirstSeenAtOk returns a tuple with the FirstSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenAt

`func (o *ServiceVersionsInner) SetFirstSeenAt(v time.Time)`

SetFirstSeenAt sets FirstSeenAt field to given value.


### GetLastSeenAt

`func (o *ServiceVersionsInner) GetLastSeenAt() time.Time`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *ServiceVersionsInner) GetLastSeenAtOk() (*time.Time, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *ServiceVersionsInner) SetLastSeenAt(v time.Time)`

SetLastSeenAt sets LastSeenAt field to given value.


### GetFaulty

`func (o *ServiceVersionsInner) GetFaulty() bool`

GetFaulty returns the Faulty field if non-nil, zero value otherwise.

### GetFaultyOk

`func (o *ServiceVersionsInner) GetFaultyOk() (*bool, bool)`

GetFaultyOk returns a tuple with the Faulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaulty

`func (o *ServiceVersionsInner) SetFaulty(v bool)`

SetFaulty sets Faulty field to given value.

### HasFaulty

`func (o *ServiceVersionsInner) HasFaulty() bool`

HasFaulty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


