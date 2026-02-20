# MonitorConfigurationCertificateExpiry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**WarnBeforeInDays** | **int32** |  | 
**CloudAccounts** | Pointer to **[]string** |  | [optional] 
**AggregationAlertLogic** | **string** |  | 
**NoDataBehavior** | **string** |  | 

## Methods

### NewMonitorConfigurationCertificateExpiry

`func NewMonitorConfigurationCertificateExpiry(type_ string, warnBeforeInDays int32, aggregationAlertLogic string, noDataBehavior string, ) *MonitorConfigurationCertificateExpiry`

NewMonitorConfigurationCertificateExpiry instantiates a new MonitorConfigurationCertificateExpiry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorConfigurationCertificateExpiryWithDefaults

`func NewMonitorConfigurationCertificateExpiryWithDefaults() *MonitorConfigurationCertificateExpiry`

NewMonitorConfigurationCertificateExpiryWithDefaults instantiates a new MonitorConfigurationCertificateExpiry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MonitorConfigurationCertificateExpiry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonitorConfigurationCertificateExpiry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonitorConfigurationCertificateExpiry) SetType(v string)`

SetType sets Type field to given value.


### GetWarnBeforeInDays

`func (o *MonitorConfigurationCertificateExpiry) GetWarnBeforeInDays() int32`

GetWarnBeforeInDays returns the WarnBeforeInDays field if non-nil, zero value otherwise.

### GetWarnBeforeInDaysOk

`func (o *MonitorConfigurationCertificateExpiry) GetWarnBeforeInDaysOk() (*int32, bool)`

GetWarnBeforeInDaysOk returns a tuple with the WarnBeforeInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnBeforeInDays

`func (o *MonitorConfigurationCertificateExpiry) SetWarnBeforeInDays(v int32)`

SetWarnBeforeInDays sets WarnBeforeInDays field to given value.


### GetCloudAccounts

`func (o *MonitorConfigurationCertificateExpiry) GetCloudAccounts() []string`

GetCloudAccounts returns the CloudAccounts field if non-nil, zero value otherwise.

### GetCloudAccountsOk

`func (o *MonitorConfigurationCertificateExpiry) GetCloudAccountsOk() (*[]string, bool)`

GetCloudAccountsOk returns a tuple with the CloudAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAccounts

`func (o *MonitorConfigurationCertificateExpiry) SetCloudAccounts(v []string)`

SetCloudAccounts sets CloudAccounts field to given value.

### HasCloudAccounts

`func (o *MonitorConfigurationCertificateExpiry) HasCloudAccounts() bool`

HasCloudAccounts returns a boolean if a field has been set.

### GetAggregationAlertLogic

`func (o *MonitorConfigurationCertificateExpiry) GetAggregationAlertLogic() string`

GetAggregationAlertLogic returns the AggregationAlertLogic field if non-nil, zero value otherwise.

### GetAggregationAlertLogicOk

`func (o *MonitorConfigurationCertificateExpiry) GetAggregationAlertLogicOk() (*string, bool)`

GetAggregationAlertLogicOk returns a tuple with the AggregationAlertLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationAlertLogic

`func (o *MonitorConfigurationCertificateExpiry) SetAggregationAlertLogic(v string)`

SetAggregationAlertLogic sets AggregationAlertLogic field to given value.


### GetNoDataBehavior

`func (o *MonitorConfigurationCertificateExpiry) GetNoDataBehavior() string`

GetNoDataBehavior returns the NoDataBehavior field if non-nil, zero value otherwise.

### GetNoDataBehaviorOk

`func (o *MonitorConfigurationCertificateExpiry) GetNoDataBehaviorOk() (*string, bool)`

GetNoDataBehaviorOk returns a tuple with the NoDataBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDataBehavior

`func (o *MonitorConfigurationCertificateExpiry) SetNoDataBehavior(v string)`

SetNoDataBehavior sets NoDataBehavior field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


