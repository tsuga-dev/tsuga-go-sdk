# MySqlConnectionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**User** | **string** |  | 
**Password** | **string** |  | 
**Host** | **string** |  | 
**Port** | **float32** |  | 
**DbName** | **string** |  | 
**SslRequire** | **bool** |  | [default to true]

## Methods

### NewMySqlConnectionInput

`func NewMySqlConnectionInput(type_ string, user string, password string, host string, port float32, dbName string, sslRequire bool, ) *MySqlConnectionInput`

NewMySqlConnectionInput instantiates a new MySqlConnectionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMySqlConnectionInputWithDefaults

`func NewMySqlConnectionInputWithDefaults() *MySqlConnectionInput`

NewMySqlConnectionInputWithDefaults instantiates a new MySqlConnectionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MySqlConnectionInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MySqlConnectionInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MySqlConnectionInput) SetType(v string)`

SetType sets Type field to given value.


### GetUser

`func (o *MySqlConnectionInput) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MySqlConnectionInput) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *MySqlConnectionInput) SetUser(v string)`

SetUser sets User field to given value.


### GetPassword

`func (o *MySqlConnectionInput) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MySqlConnectionInput) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MySqlConnectionInput) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetHost

`func (o *MySqlConnectionInput) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *MySqlConnectionInput) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *MySqlConnectionInput) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *MySqlConnectionInput) GetPort() float32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *MySqlConnectionInput) GetPortOk() (*float32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *MySqlConnectionInput) SetPort(v float32)`

SetPort sets Port field to given value.


### GetDbName

`func (o *MySqlConnectionInput) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *MySqlConnectionInput) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *MySqlConnectionInput) SetDbName(v string)`

SetDbName sets DbName field to given value.


### GetSslRequire

`func (o *MySqlConnectionInput) GetSslRequire() bool`

GetSslRequire returns the SslRequire field if non-nil, zero value otherwise.

### GetSslRequireOk

`func (o *MySqlConnectionInput) GetSslRequireOk() (*bool, bool)`

GetSslRequireOk returns a tuple with the SslRequire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslRequire

`func (o *MySqlConnectionInput) SetSslRequire(v bool)`

SetSslRequire sets SslRequire field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


