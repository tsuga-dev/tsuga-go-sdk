# CreateCustomRoleRequestAssets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Database** | Pointer to **string** | Access granted on this resource type. | [optional] 
**Inventory** | Pointer to **string** | Access granted on this resource type. | [optional] 
**Queue** | Pointer to **string** | Access granted on this resource type. | [optional] 
**Service** | Pointer to **string** | Access granted on this resource type. | [optional] 
**AuditLog** | Pointer to **string** | Access granted on this resource type. | [optional] 
**AllowedDomains** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**AuditLogStreamDestination** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**ClusterIntegrationConfig** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**CustomUsageTag** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**DatabricksConnection** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**GithubConnection** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**LlmProvider** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**RetentionPolicy** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**SamlConfig** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**SensitiveDataScannerRule** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**ServiceTarget** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**TeamRemapConfig** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**AdaptiveNotificationRule** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**CloudAccount** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**DashboardFolder** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**MonitorFolder** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**DashboardReport** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**DashboardShare** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**Dashboard** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**IngestionApiKey** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**Investigation** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**LogRoute** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**MetricRoute** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**Monitor** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**Notebook** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**NotificationIntegration** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**NotificationRule** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**NotificationSilence** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**OperationApiKey** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**QualityReport** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**RumPublicToken** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**SandboxConfiguration** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**Slo** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**TagPolicy** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**UserInvitation** | Pointer to **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | [optional] 
**Team** | **string** | Access granted on this resource type. &#x60;write&#x60; implies &#x60;read&#x60;. | 

## Methods

### NewCreateCustomRoleRequestAssets

`func NewCreateCustomRoleRequestAssets(team string, ) *CreateCustomRoleRequestAssets`

NewCreateCustomRoleRequestAssets instantiates a new CreateCustomRoleRequestAssets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomRoleRequestAssetsWithDefaults

`func NewCreateCustomRoleRequestAssetsWithDefaults() *CreateCustomRoleRequestAssets`

NewCreateCustomRoleRequestAssetsWithDefaults instantiates a new CreateCustomRoleRequestAssets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabase

`func (o *CreateCustomRoleRequestAssets) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *CreateCustomRoleRequestAssets) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *CreateCustomRoleRequestAssets) SetDatabase(v string)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *CreateCustomRoleRequestAssets) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetInventory

`func (o *CreateCustomRoleRequestAssets) GetInventory() string`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *CreateCustomRoleRequestAssets) GetInventoryOk() (*string, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *CreateCustomRoleRequestAssets) SetInventory(v string)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *CreateCustomRoleRequestAssets) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetQueue

`func (o *CreateCustomRoleRequestAssets) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *CreateCustomRoleRequestAssets) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *CreateCustomRoleRequestAssets) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *CreateCustomRoleRequestAssets) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### GetService

`func (o *CreateCustomRoleRequestAssets) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CreateCustomRoleRequestAssets) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CreateCustomRoleRequestAssets) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *CreateCustomRoleRequestAssets) HasService() bool`

HasService returns a boolean if a field has been set.

### GetAuditLog

`func (o *CreateCustomRoleRequestAssets) GetAuditLog() string`

GetAuditLog returns the AuditLog field if non-nil, zero value otherwise.

### GetAuditLogOk

`func (o *CreateCustomRoleRequestAssets) GetAuditLogOk() (*string, bool)`

GetAuditLogOk returns a tuple with the AuditLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLog

`func (o *CreateCustomRoleRequestAssets) SetAuditLog(v string)`

SetAuditLog sets AuditLog field to given value.

### HasAuditLog

`func (o *CreateCustomRoleRequestAssets) HasAuditLog() bool`

HasAuditLog returns a boolean if a field has been set.

### GetAllowedDomains

`func (o *CreateCustomRoleRequestAssets) GetAllowedDomains() string`

GetAllowedDomains returns the AllowedDomains field if non-nil, zero value otherwise.

### GetAllowedDomainsOk

`func (o *CreateCustomRoleRequestAssets) GetAllowedDomainsOk() (*string, bool)`

GetAllowedDomainsOk returns a tuple with the AllowedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDomains

`func (o *CreateCustomRoleRequestAssets) SetAllowedDomains(v string)`

SetAllowedDomains sets AllowedDomains field to given value.

### HasAllowedDomains

`func (o *CreateCustomRoleRequestAssets) HasAllowedDomains() bool`

HasAllowedDomains returns a boolean if a field has been set.

### GetAuditLogStreamDestination

`func (o *CreateCustomRoleRequestAssets) GetAuditLogStreamDestination() string`

GetAuditLogStreamDestination returns the AuditLogStreamDestination field if non-nil, zero value otherwise.

### GetAuditLogStreamDestinationOk

`func (o *CreateCustomRoleRequestAssets) GetAuditLogStreamDestinationOk() (*string, bool)`

GetAuditLogStreamDestinationOk returns a tuple with the AuditLogStreamDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogStreamDestination

`func (o *CreateCustomRoleRequestAssets) SetAuditLogStreamDestination(v string)`

SetAuditLogStreamDestination sets AuditLogStreamDestination field to given value.

### HasAuditLogStreamDestination

`func (o *CreateCustomRoleRequestAssets) HasAuditLogStreamDestination() bool`

HasAuditLogStreamDestination returns a boolean if a field has been set.

### GetClusterIntegrationConfig

`func (o *CreateCustomRoleRequestAssets) GetClusterIntegrationConfig() string`

GetClusterIntegrationConfig returns the ClusterIntegrationConfig field if non-nil, zero value otherwise.

### GetClusterIntegrationConfigOk

`func (o *CreateCustomRoleRequestAssets) GetClusterIntegrationConfigOk() (*string, bool)`

GetClusterIntegrationConfigOk returns a tuple with the ClusterIntegrationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIntegrationConfig

`func (o *CreateCustomRoleRequestAssets) SetClusterIntegrationConfig(v string)`

SetClusterIntegrationConfig sets ClusterIntegrationConfig field to given value.

### HasClusterIntegrationConfig

`func (o *CreateCustomRoleRequestAssets) HasClusterIntegrationConfig() bool`

HasClusterIntegrationConfig returns a boolean if a field has been set.

### GetCustomUsageTag

`func (o *CreateCustomRoleRequestAssets) GetCustomUsageTag() string`

GetCustomUsageTag returns the CustomUsageTag field if non-nil, zero value otherwise.

### GetCustomUsageTagOk

`func (o *CreateCustomRoleRequestAssets) GetCustomUsageTagOk() (*string, bool)`

GetCustomUsageTagOk returns a tuple with the CustomUsageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUsageTag

`func (o *CreateCustomRoleRequestAssets) SetCustomUsageTag(v string)`

SetCustomUsageTag sets CustomUsageTag field to given value.

### HasCustomUsageTag

`func (o *CreateCustomRoleRequestAssets) HasCustomUsageTag() bool`

HasCustomUsageTag returns a boolean if a field has been set.

### GetDatabricksConnection

`func (o *CreateCustomRoleRequestAssets) GetDatabricksConnection() string`

GetDatabricksConnection returns the DatabricksConnection field if non-nil, zero value otherwise.

### GetDatabricksConnectionOk

`func (o *CreateCustomRoleRequestAssets) GetDatabricksConnectionOk() (*string, bool)`

GetDatabricksConnectionOk returns a tuple with the DatabricksConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabricksConnection

`func (o *CreateCustomRoleRequestAssets) SetDatabricksConnection(v string)`

SetDatabricksConnection sets DatabricksConnection field to given value.

### HasDatabricksConnection

`func (o *CreateCustomRoleRequestAssets) HasDatabricksConnection() bool`

HasDatabricksConnection returns a boolean if a field has been set.

### GetGithubConnection

`func (o *CreateCustomRoleRequestAssets) GetGithubConnection() string`

GetGithubConnection returns the GithubConnection field if non-nil, zero value otherwise.

### GetGithubConnectionOk

`func (o *CreateCustomRoleRequestAssets) GetGithubConnectionOk() (*string, bool)`

GetGithubConnectionOk returns a tuple with the GithubConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubConnection

`func (o *CreateCustomRoleRequestAssets) SetGithubConnection(v string)`

SetGithubConnection sets GithubConnection field to given value.

### HasGithubConnection

`func (o *CreateCustomRoleRequestAssets) HasGithubConnection() bool`

HasGithubConnection returns a boolean if a field has been set.

### GetLlmProvider

`func (o *CreateCustomRoleRequestAssets) GetLlmProvider() string`

GetLlmProvider returns the LlmProvider field if non-nil, zero value otherwise.

### GetLlmProviderOk

`func (o *CreateCustomRoleRequestAssets) GetLlmProviderOk() (*string, bool)`

GetLlmProviderOk returns a tuple with the LlmProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlmProvider

`func (o *CreateCustomRoleRequestAssets) SetLlmProvider(v string)`

SetLlmProvider sets LlmProvider field to given value.

### HasLlmProvider

`func (o *CreateCustomRoleRequestAssets) HasLlmProvider() bool`

HasLlmProvider returns a boolean if a field has been set.

### GetRetentionPolicy

`func (o *CreateCustomRoleRequestAssets) GetRetentionPolicy() string`

GetRetentionPolicy returns the RetentionPolicy field if non-nil, zero value otherwise.

### GetRetentionPolicyOk

`func (o *CreateCustomRoleRequestAssets) GetRetentionPolicyOk() (*string, bool)`

GetRetentionPolicyOk returns a tuple with the RetentionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicy

`func (o *CreateCustomRoleRequestAssets) SetRetentionPolicy(v string)`

SetRetentionPolicy sets RetentionPolicy field to given value.

### HasRetentionPolicy

`func (o *CreateCustomRoleRequestAssets) HasRetentionPolicy() bool`

HasRetentionPolicy returns a boolean if a field has been set.

### GetSamlConfig

`func (o *CreateCustomRoleRequestAssets) GetSamlConfig() string`

GetSamlConfig returns the SamlConfig field if non-nil, zero value otherwise.

### GetSamlConfigOk

`func (o *CreateCustomRoleRequestAssets) GetSamlConfigOk() (*string, bool)`

GetSamlConfigOk returns a tuple with the SamlConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlConfig

`func (o *CreateCustomRoleRequestAssets) SetSamlConfig(v string)`

SetSamlConfig sets SamlConfig field to given value.

### HasSamlConfig

`func (o *CreateCustomRoleRequestAssets) HasSamlConfig() bool`

HasSamlConfig returns a boolean if a field has been set.

### GetSensitiveDataScannerRule

`func (o *CreateCustomRoleRequestAssets) GetSensitiveDataScannerRule() string`

GetSensitiveDataScannerRule returns the SensitiveDataScannerRule field if non-nil, zero value otherwise.

### GetSensitiveDataScannerRuleOk

`func (o *CreateCustomRoleRequestAssets) GetSensitiveDataScannerRuleOk() (*string, bool)`

GetSensitiveDataScannerRuleOk returns a tuple with the SensitiveDataScannerRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitiveDataScannerRule

`func (o *CreateCustomRoleRequestAssets) SetSensitiveDataScannerRule(v string)`

SetSensitiveDataScannerRule sets SensitiveDataScannerRule field to given value.

### HasSensitiveDataScannerRule

`func (o *CreateCustomRoleRequestAssets) HasSensitiveDataScannerRule() bool`

HasSensitiveDataScannerRule returns a boolean if a field has been set.

### GetServiceTarget

`func (o *CreateCustomRoleRequestAssets) GetServiceTarget() string`

GetServiceTarget returns the ServiceTarget field if non-nil, zero value otherwise.

### GetServiceTargetOk

`func (o *CreateCustomRoleRequestAssets) GetServiceTargetOk() (*string, bool)`

GetServiceTargetOk returns a tuple with the ServiceTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTarget

`func (o *CreateCustomRoleRequestAssets) SetServiceTarget(v string)`

SetServiceTarget sets ServiceTarget field to given value.

### HasServiceTarget

`func (o *CreateCustomRoleRequestAssets) HasServiceTarget() bool`

HasServiceTarget returns a boolean if a field has been set.

### GetTeamRemapConfig

`func (o *CreateCustomRoleRequestAssets) GetTeamRemapConfig() string`

GetTeamRemapConfig returns the TeamRemapConfig field if non-nil, zero value otherwise.

### GetTeamRemapConfigOk

`func (o *CreateCustomRoleRequestAssets) GetTeamRemapConfigOk() (*string, bool)`

GetTeamRemapConfigOk returns a tuple with the TeamRemapConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamRemapConfig

`func (o *CreateCustomRoleRequestAssets) SetTeamRemapConfig(v string)`

SetTeamRemapConfig sets TeamRemapConfig field to given value.

### HasTeamRemapConfig

`func (o *CreateCustomRoleRequestAssets) HasTeamRemapConfig() bool`

HasTeamRemapConfig returns a boolean if a field has been set.

### GetAdaptiveNotificationRule

`func (o *CreateCustomRoleRequestAssets) GetAdaptiveNotificationRule() string`

GetAdaptiveNotificationRule returns the AdaptiveNotificationRule field if non-nil, zero value otherwise.

### GetAdaptiveNotificationRuleOk

`func (o *CreateCustomRoleRequestAssets) GetAdaptiveNotificationRuleOk() (*string, bool)`

GetAdaptiveNotificationRuleOk returns a tuple with the AdaptiveNotificationRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdaptiveNotificationRule

`func (o *CreateCustomRoleRequestAssets) SetAdaptiveNotificationRule(v string)`

SetAdaptiveNotificationRule sets AdaptiveNotificationRule field to given value.

### HasAdaptiveNotificationRule

`func (o *CreateCustomRoleRequestAssets) HasAdaptiveNotificationRule() bool`

HasAdaptiveNotificationRule returns a boolean if a field has been set.

### GetCloudAccount

`func (o *CreateCustomRoleRequestAssets) GetCloudAccount() string`

GetCloudAccount returns the CloudAccount field if non-nil, zero value otherwise.

### GetCloudAccountOk

`func (o *CreateCustomRoleRequestAssets) GetCloudAccountOk() (*string, bool)`

GetCloudAccountOk returns a tuple with the CloudAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAccount

`func (o *CreateCustomRoleRequestAssets) SetCloudAccount(v string)`

SetCloudAccount sets CloudAccount field to given value.

### HasCloudAccount

`func (o *CreateCustomRoleRequestAssets) HasCloudAccount() bool`

HasCloudAccount returns a boolean if a field has been set.

### GetDashboardFolder

`func (o *CreateCustomRoleRequestAssets) GetDashboardFolder() string`

GetDashboardFolder returns the DashboardFolder field if non-nil, zero value otherwise.

### GetDashboardFolderOk

`func (o *CreateCustomRoleRequestAssets) GetDashboardFolderOk() (*string, bool)`

GetDashboardFolderOk returns a tuple with the DashboardFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardFolder

`func (o *CreateCustomRoleRequestAssets) SetDashboardFolder(v string)`

SetDashboardFolder sets DashboardFolder field to given value.

### HasDashboardFolder

`func (o *CreateCustomRoleRequestAssets) HasDashboardFolder() bool`

HasDashboardFolder returns a boolean if a field has been set.

### GetMonitorFolder

`func (o *CreateCustomRoleRequestAssets) GetMonitorFolder() string`

GetMonitorFolder returns the MonitorFolder field if non-nil, zero value otherwise.

### GetMonitorFolderOk

`func (o *CreateCustomRoleRequestAssets) GetMonitorFolderOk() (*string, bool)`

GetMonitorFolderOk returns a tuple with the MonitorFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorFolder

`func (o *CreateCustomRoleRequestAssets) SetMonitorFolder(v string)`

SetMonitorFolder sets MonitorFolder field to given value.

### HasMonitorFolder

`func (o *CreateCustomRoleRequestAssets) HasMonitorFolder() bool`

HasMonitorFolder returns a boolean if a field has been set.

### GetDashboardReport

`func (o *CreateCustomRoleRequestAssets) GetDashboardReport() string`

GetDashboardReport returns the DashboardReport field if non-nil, zero value otherwise.

### GetDashboardReportOk

`func (o *CreateCustomRoleRequestAssets) GetDashboardReportOk() (*string, bool)`

GetDashboardReportOk returns a tuple with the DashboardReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardReport

`func (o *CreateCustomRoleRequestAssets) SetDashboardReport(v string)`

SetDashboardReport sets DashboardReport field to given value.

### HasDashboardReport

`func (o *CreateCustomRoleRequestAssets) HasDashboardReport() bool`

HasDashboardReport returns a boolean if a field has been set.

### GetDashboardShare

`func (o *CreateCustomRoleRequestAssets) GetDashboardShare() string`

GetDashboardShare returns the DashboardShare field if non-nil, zero value otherwise.

### GetDashboardShareOk

`func (o *CreateCustomRoleRequestAssets) GetDashboardShareOk() (*string, bool)`

GetDashboardShareOk returns a tuple with the DashboardShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardShare

`func (o *CreateCustomRoleRequestAssets) SetDashboardShare(v string)`

SetDashboardShare sets DashboardShare field to given value.

### HasDashboardShare

`func (o *CreateCustomRoleRequestAssets) HasDashboardShare() bool`

HasDashboardShare returns a boolean if a field has been set.

### GetDashboard

`func (o *CreateCustomRoleRequestAssets) GetDashboard() string`

GetDashboard returns the Dashboard field if non-nil, zero value otherwise.

### GetDashboardOk

`func (o *CreateCustomRoleRequestAssets) GetDashboardOk() (*string, bool)`

GetDashboardOk returns a tuple with the Dashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboard

`func (o *CreateCustomRoleRequestAssets) SetDashboard(v string)`

SetDashboard sets Dashboard field to given value.

### HasDashboard

`func (o *CreateCustomRoleRequestAssets) HasDashboard() bool`

HasDashboard returns a boolean if a field has been set.

### GetIngestionApiKey

`func (o *CreateCustomRoleRequestAssets) GetIngestionApiKey() string`

GetIngestionApiKey returns the IngestionApiKey field if non-nil, zero value otherwise.

### GetIngestionApiKeyOk

`func (o *CreateCustomRoleRequestAssets) GetIngestionApiKeyOk() (*string, bool)`

GetIngestionApiKeyOk returns a tuple with the IngestionApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestionApiKey

`func (o *CreateCustomRoleRequestAssets) SetIngestionApiKey(v string)`

SetIngestionApiKey sets IngestionApiKey field to given value.

### HasIngestionApiKey

`func (o *CreateCustomRoleRequestAssets) HasIngestionApiKey() bool`

HasIngestionApiKey returns a boolean if a field has been set.

### GetInvestigation

`func (o *CreateCustomRoleRequestAssets) GetInvestigation() string`

GetInvestigation returns the Investigation field if non-nil, zero value otherwise.

### GetInvestigationOk

`func (o *CreateCustomRoleRequestAssets) GetInvestigationOk() (*string, bool)`

GetInvestigationOk returns a tuple with the Investigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigation

`func (o *CreateCustomRoleRequestAssets) SetInvestigation(v string)`

SetInvestigation sets Investigation field to given value.

### HasInvestigation

`func (o *CreateCustomRoleRequestAssets) HasInvestigation() bool`

HasInvestigation returns a boolean if a field has been set.

### GetLogRoute

`func (o *CreateCustomRoleRequestAssets) GetLogRoute() string`

GetLogRoute returns the LogRoute field if non-nil, zero value otherwise.

### GetLogRouteOk

`func (o *CreateCustomRoleRequestAssets) GetLogRouteOk() (*string, bool)`

GetLogRouteOk returns a tuple with the LogRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRoute

`func (o *CreateCustomRoleRequestAssets) SetLogRoute(v string)`

SetLogRoute sets LogRoute field to given value.

### HasLogRoute

`func (o *CreateCustomRoleRequestAssets) HasLogRoute() bool`

HasLogRoute returns a boolean if a field has been set.

### GetMetricRoute

`func (o *CreateCustomRoleRequestAssets) GetMetricRoute() string`

GetMetricRoute returns the MetricRoute field if non-nil, zero value otherwise.

### GetMetricRouteOk

`func (o *CreateCustomRoleRequestAssets) GetMetricRouteOk() (*string, bool)`

GetMetricRouteOk returns a tuple with the MetricRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricRoute

`func (o *CreateCustomRoleRequestAssets) SetMetricRoute(v string)`

SetMetricRoute sets MetricRoute field to given value.

### HasMetricRoute

`func (o *CreateCustomRoleRequestAssets) HasMetricRoute() bool`

HasMetricRoute returns a boolean if a field has been set.

### GetMonitor

`func (o *CreateCustomRoleRequestAssets) GetMonitor() string`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *CreateCustomRoleRequestAssets) GetMonitorOk() (*string, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *CreateCustomRoleRequestAssets) SetMonitor(v string)`

SetMonitor sets Monitor field to given value.

### HasMonitor

`func (o *CreateCustomRoleRequestAssets) HasMonitor() bool`

HasMonitor returns a boolean if a field has been set.

### GetNotebook

`func (o *CreateCustomRoleRequestAssets) GetNotebook() string`

GetNotebook returns the Notebook field if non-nil, zero value otherwise.

### GetNotebookOk

`func (o *CreateCustomRoleRequestAssets) GetNotebookOk() (*string, bool)`

GetNotebookOk returns a tuple with the Notebook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotebook

`func (o *CreateCustomRoleRequestAssets) SetNotebook(v string)`

SetNotebook sets Notebook field to given value.

### HasNotebook

`func (o *CreateCustomRoleRequestAssets) HasNotebook() bool`

HasNotebook returns a boolean if a field has been set.

### GetNotificationIntegration

`func (o *CreateCustomRoleRequestAssets) GetNotificationIntegration() string`

GetNotificationIntegration returns the NotificationIntegration field if non-nil, zero value otherwise.

### GetNotificationIntegrationOk

`func (o *CreateCustomRoleRequestAssets) GetNotificationIntegrationOk() (*string, bool)`

GetNotificationIntegrationOk returns a tuple with the NotificationIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationIntegration

`func (o *CreateCustomRoleRequestAssets) SetNotificationIntegration(v string)`

SetNotificationIntegration sets NotificationIntegration field to given value.

### HasNotificationIntegration

`func (o *CreateCustomRoleRequestAssets) HasNotificationIntegration() bool`

HasNotificationIntegration returns a boolean if a field has been set.

### GetNotificationRule

`func (o *CreateCustomRoleRequestAssets) GetNotificationRule() string`

GetNotificationRule returns the NotificationRule field if non-nil, zero value otherwise.

### GetNotificationRuleOk

`func (o *CreateCustomRoleRequestAssets) GetNotificationRuleOk() (*string, bool)`

GetNotificationRuleOk returns a tuple with the NotificationRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationRule

`func (o *CreateCustomRoleRequestAssets) SetNotificationRule(v string)`

SetNotificationRule sets NotificationRule field to given value.

### HasNotificationRule

`func (o *CreateCustomRoleRequestAssets) HasNotificationRule() bool`

HasNotificationRule returns a boolean if a field has been set.

### GetNotificationSilence

`func (o *CreateCustomRoleRequestAssets) GetNotificationSilence() string`

GetNotificationSilence returns the NotificationSilence field if non-nil, zero value otherwise.

### GetNotificationSilenceOk

`func (o *CreateCustomRoleRequestAssets) GetNotificationSilenceOk() (*string, bool)`

GetNotificationSilenceOk returns a tuple with the NotificationSilence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSilence

`func (o *CreateCustomRoleRequestAssets) SetNotificationSilence(v string)`

SetNotificationSilence sets NotificationSilence field to given value.

### HasNotificationSilence

`func (o *CreateCustomRoleRequestAssets) HasNotificationSilence() bool`

HasNotificationSilence returns a boolean if a field has been set.

### GetOperationApiKey

`func (o *CreateCustomRoleRequestAssets) GetOperationApiKey() string`

GetOperationApiKey returns the OperationApiKey field if non-nil, zero value otherwise.

### GetOperationApiKeyOk

`func (o *CreateCustomRoleRequestAssets) GetOperationApiKeyOk() (*string, bool)`

GetOperationApiKeyOk returns a tuple with the OperationApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationApiKey

`func (o *CreateCustomRoleRequestAssets) SetOperationApiKey(v string)`

SetOperationApiKey sets OperationApiKey field to given value.

### HasOperationApiKey

`func (o *CreateCustomRoleRequestAssets) HasOperationApiKey() bool`

HasOperationApiKey returns a boolean if a field has been set.

### GetQualityReport

`func (o *CreateCustomRoleRequestAssets) GetQualityReport() string`

GetQualityReport returns the QualityReport field if non-nil, zero value otherwise.

### GetQualityReportOk

`func (o *CreateCustomRoleRequestAssets) GetQualityReportOk() (*string, bool)`

GetQualityReportOk returns a tuple with the QualityReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityReport

`func (o *CreateCustomRoleRequestAssets) SetQualityReport(v string)`

SetQualityReport sets QualityReport field to given value.

### HasQualityReport

`func (o *CreateCustomRoleRequestAssets) HasQualityReport() bool`

HasQualityReport returns a boolean if a field has been set.

### GetRumPublicToken

`func (o *CreateCustomRoleRequestAssets) GetRumPublicToken() string`

GetRumPublicToken returns the RumPublicToken field if non-nil, zero value otherwise.

### GetRumPublicTokenOk

`func (o *CreateCustomRoleRequestAssets) GetRumPublicTokenOk() (*string, bool)`

GetRumPublicTokenOk returns a tuple with the RumPublicToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRumPublicToken

`func (o *CreateCustomRoleRequestAssets) SetRumPublicToken(v string)`

SetRumPublicToken sets RumPublicToken field to given value.

### HasRumPublicToken

`func (o *CreateCustomRoleRequestAssets) HasRumPublicToken() bool`

HasRumPublicToken returns a boolean if a field has been set.

### GetSandboxConfiguration

`func (o *CreateCustomRoleRequestAssets) GetSandboxConfiguration() string`

GetSandboxConfiguration returns the SandboxConfiguration field if non-nil, zero value otherwise.

### GetSandboxConfigurationOk

`func (o *CreateCustomRoleRequestAssets) GetSandboxConfigurationOk() (*string, bool)`

GetSandboxConfigurationOk returns a tuple with the SandboxConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandboxConfiguration

`func (o *CreateCustomRoleRequestAssets) SetSandboxConfiguration(v string)`

SetSandboxConfiguration sets SandboxConfiguration field to given value.

### HasSandboxConfiguration

`func (o *CreateCustomRoleRequestAssets) HasSandboxConfiguration() bool`

HasSandboxConfiguration returns a boolean if a field has been set.

### GetSlo

`func (o *CreateCustomRoleRequestAssets) GetSlo() string`

GetSlo returns the Slo field if non-nil, zero value otherwise.

### GetSloOk

`func (o *CreateCustomRoleRequestAssets) GetSloOk() (*string, bool)`

GetSloOk returns a tuple with the Slo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlo

`func (o *CreateCustomRoleRequestAssets) SetSlo(v string)`

SetSlo sets Slo field to given value.

### HasSlo

`func (o *CreateCustomRoleRequestAssets) HasSlo() bool`

HasSlo returns a boolean if a field has been set.

### GetTagPolicy

`func (o *CreateCustomRoleRequestAssets) GetTagPolicy() string`

GetTagPolicy returns the TagPolicy field if non-nil, zero value otherwise.

### GetTagPolicyOk

`func (o *CreateCustomRoleRequestAssets) GetTagPolicyOk() (*string, bool)`

GetTagPolicyOk returns a tuple with the TagPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagPolicy

`func (o *CreateCustomRoleRequestAssets) SetTagPolicy(v string)`

SetTagPolicy sets TagPolicy field to given value.

### HasTagPolicy

`func (o *CreateCustomRoleRequestAssets) HasTagPolicy() bool`

HasTagPolicy returns a boolean if a field has been set.

### GetUserInvitation

`func (o *CreateCustomRoleRequestAssets) GetUserInvitation() string`

GetUserInvitation returns the UserInvitation field if non-nil, zero value otherwise.

### GetUserInvitationOk

`func (o *CreateCustomRoleRequestAssets) GetUserInvitationOk() (*string, bool)`

GetUserInvitationOk returns a tuple with the UserInvitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInvitation

`func (o *CreateCustomRoleRequestAssets) SetUserInvitation(v string)`

SetUserInvitation sets UserInvitation field to given value.

### HasUserInvitation

`func (o *CreateCustomRoleRequestAssets) HasUserInvitation() bool`

HasUserInvitation returns a boolean if a field has been set.

### GetTeam

`func (o *CreateCustomRoleRequestAssets) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *CreateCustomRoleRequestAssets) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *CreateCustomRoleRequestAssets) SetTeam(v string)`

SetTeam sets Team field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


