# PostgresConnectionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**User** | **string** |  | 
**Password** | **string** |  | 
**Host** | **string** |  | 
**Port** | **int32** |  | 
**DbName** | **string** |  | 
**SslRequire** | **bool** |  | [default to true]

## Methods

### NewPostgresConnectionInput

`func NewPostgresConnectionInput(type_ string, user string, password string, host string, port int32, dbName string, sslRequire bool, ) *PostgresConnectionInput`

NewPostgresConnectionInput instantiates a new PostgresConnectionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresConnectionInputWithDefaults

`func NewPostgresConnectionInputWithDefaults() *PostgresConnectionInput`

NewPostgresConnectionInputWithDefaults instantiates a new PostgresConnectionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PostgresConnectionInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostgresConnectionInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostgresConnectionInput) SetType(v string)`

SetType sets Type field to given value.


### GetUser

`func (o *PostgresConnectionInput) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PostgresConnectionInput) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PostgresConnectionInput) SetUser(v string)`

SetUser sets User field to given value.


### GetPassword

`func (o *PostgresConnectionInput) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PostgresConnectionInput) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PostgresConnectionInput) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetHost

`func (o *PostgresConnectionInput) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *PostgresConnectionInput) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *PostgresConnectionInput) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *PostgresConnectionInput) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PostgresConnectionInput) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PostgresConnectionInput) SetPort(v int32)`

SetPort sets Port field to given value.


### GetDbName

`func (o *PostgresConnectionInput) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *PostgresConnectionInput) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *PostgresConnectionInput) SetDbName(v string)`

SetDbName sets DbName field to given value.


### GetSslRequire

`func (o *PostgresConnectionInput) GetSslRequire() bool`

GetSslRequire returns the SslRequire field if non-nil, zero value otherwise.

### GetSslRequireOk

`func (o *PostgresConnectionInput) GetSslRequireOk() (*bool, bool)`

GetSslRequireOk returns a tuple with the SslRequire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslRequire

`func (o *PostgresConnectionInput) SetSslRequire(v bool)`

SetSslRequire sets SslRequire field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


