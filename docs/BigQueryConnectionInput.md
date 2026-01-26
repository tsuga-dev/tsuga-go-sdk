# BigQueryConnectionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Auth** | [**BigQueryConnectionInputAuth**](BigQueryConnectionInputAuth.md) |  | 

## Methods

### NewBigQueryConnectionInput

`func NewBigQueryConnectionInput(type_ string, auth BigQueryConnectionInputAuth, ) *BigQueryConnectionInput`

NewBigQueryConnectionInput instantiates a new BigQueryConnectionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBigQueryConnectionInputWithDefaults

`func NewBigQueryConnectionInputWithDefaults() *BigQueryConnectionInput`

NewBigQueryConnectionInputWithDefaults instantiates a new BigQueryConnectionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BigQueryConnectionInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BigQueryConnectionInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BigQueryConnectionInput) SetType(v string)`

SetType sets Type field to given value.


### GetAuth

`func (o *BigQueryConnectionInput) GetAuth() BigQueryConnectionInputAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *BigQueryConnectionInput) GetAuthOk() (*BigQueryConnectionInputAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *BigQueryConnectionInput) SetAuth(v BigQueryConnectionInputAuth)`

SetAuth sets Auth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


