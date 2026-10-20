# CreateOperationApiKeyRequestPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Investigations** | Pointer to **string** | Investigations, the records collecting the timeline and findings of an incident analysis. | [optional] 
**Dashboard** | Pointer to **string** | Dashboards and the graphs they contain. | [optional] 
**DashboardFolder** | Pointer to **string** | Folders organizing dashboards. | [optional] 
**MonitorFolder** | Pointer to **string** | Folders organizing monitors. | [optional] 
**Documentation** | Pointer to **string** | Tsuga documentation and API reference pages, which are public content rather than account data. | [optional] 
**IngestionApiKey** | Pointer to **string** | Ingestion API keys, used to send telemetry data to Tsuga. | [optional] 
**RumPublicToken** | Pointer to **string** | RUM public tokens, embedded in web and mobile applications to send real user monitoring telemetry. | [optional] 
**OperationApiKey** | Pointer to **string** | Operation API keys, used to authenticate API, Terraform, MCP, and Tsuga CLI clients. | [optional] 
**Inventory** | Pointer to **string** | Cloud resources discovered by inventory scans. | [optional] 
**CloudAccount** | Pointer to **string** | Cloud accounts Tsuga scans to discover resources. | [optional] 
**Logs** | Pointer to **string** | Ingested log events, through search and aggregation endpoints. | [optional] 
**LogRoute** | Pointer to **string** | Log routes, which select newly ingested logs and apply processors to them. | [optional] 
**Metrics** | Pointer to **string** | Ingested metrics, through query and aggregation endpoints. | [optional] 
**MetricRoute** | Pointer to **string** | Metric routes, which select newly ingested metrics by name and apply processors to them. | [optional] 
**Monitor** | Pointer to **string** | Monitors, the alert definitions evaluated against telemetry. | [optional] 
**AllowedDomains** | Pointer to **string** | Email domains whose users are allowed to join the organization. | [optional] 
**Queue** | Pointer to **string** | Message queues of the catalog, inferred from telemetry. | [optional] 
**QualityReport** | Pointer to **string** | Telemetry quality reports. | [optional] 
**Billing** | Pointer to **string** | Ingested data volume billed to the organization. | [optional] 
**Service** | Pointer to **string** | Services of the catalog, inferred from telemetry. | [optional] 
**ServiceTarget** | Pointer to **string** | Notification targets attached to a service. | [optional] 
**Slo** | Pointer to **string** | Service level objectives and the alerts attached to them. | [optional] 
**NotificationRule** | Pointer to **string** | Notification rules, which route alert transitions to notification targets. | [optional] 
**AdaptiveNotificationRule** | Pointer to **string** | Adaptive notification rules, which page the owners of the failing service when one of their monitors or SLOs fires. | [optional] 
**NotificationIntegration** | Pointer to **string** | Notification integrations, containing the credentials Tsuga uses when sending notifications. | [optional] 
**NotificationSilence** | Pointer to **string** | Notification silences, which suppress matching notification delivery. | [optional] 
**CustomUsageTag** | Pointer to **string** | Custom usage tags, the telemetry tag keys tracked as usage dimensions. | [optional] 
**RetentionPolicy** | Pointer to **string** | Retention policies applied to logs, metrics, and traces. | [optional] 
**SandboxConfiguration** | Pointer to **string** | Sandbox agent configurations of the key&#39;s teams, and the sandbox agent runs launched with them. | [optional] 
**TagPolicy** | Pointer to **string** | Tag policies, which enforce tag requirements on Tsuga-managed assets and telemetry data. | [optional] 
**Team** | Pointer to **string** | Teams, their settings and their memberships. | [optional] 
**Traces** | Pointer to **string** | Ingested spans, through search and aggregation endpoints. | [optional] 
**Rum** | Pointer to **string** | Ingested RUM events, through search and aggregation endpoints. | [optional] 
**User** | Pointer to **string** | User accounts of the organization. | [optional] 
**UserInvitation** | Pointer to **string** | Invitations for people who do not belong to the organization yet. | [optional] 

## Methods

### NewCreateOperationApiKeyRequestPermissions

`func NewCreateOperationApiKeyRequestPermissions() *CreateOperationApiKeyRequestPermissions`

NewCreateOperationApiKeyRequestPermissions instantiates a new CreateOperationApiKeyRequestPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOperationApiKeyRequestPermissionsWithDefaults

`func NewCreateOperationApiKeyRequestPermissionsWithDefaults() *CreateOperationApiKeyRequestPermissions`

NewCreateOperationApiKeyRequestPermissionsWithDefaults instantiates a new CreateOperationApiKeyRequestPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvestigations

`func (o *CreateOperationApiKeyRequestPermissions) GetInvestigations() string`

GetInvestigations returns the Investigations field if non-nil, zero value otherwise.

### GetInvestigationsOk

`func (o *CreateOperationApiKeyRequestPermissions) GetInvestigationsOk() (*string, bool)`

GetInvestigationsOk returns a tuple with the Investigations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigations

`func (o *CreateOperationApiKeyRequestPermissions) SetInvestigations(v string)`

SetInvestigations sets Investigations field to given value.

### HasInvestigations

`func (o *CreateOperationApiKeyRequestPermissions) HasInvestigations() bool`

HasInvestigations returns a boolean if a field has been set.

### GetDashboard

`func (o *CreateOperationApiKeyRequestPermissions) GetDashboard() string`

GetDashboard returns the Dashboard field if non-nil, zero value otherwise.

### GetDashboardOk

`func (o *CreateOperationApiKeyRequestPermissions) GetDashboardOk() (*string, bool)`

GetDashboardOk returns a tuple with the Dashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboard

`func (o *CreateOperationApiKeyRequestPermissions) SetDashboard(v string)`

SetDashboard sets Dashboard field to given value.

### HasDashboard

`func (o *CreateOperationApiKeyRequestPermissions) HasDashboard() bool`

HasDashboard returns a boolean if a field has been set.

### GetDashboardFolder

`func (o *CreateOperationApiKeyRequestPermissions) GetDashboardFolder() string`

GetDashboardFolder returns the DashboardFolder field if non-nil, zero value otherwise.

### GetDashboardFolderOk

`func (o *CreateOperationApiKeyRequestPermissions) GetDashboardFolderOk() (*string, bool)`

GetDashboardFolderOk returns a tuple with the DashboardFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardFolder

`func (o *CreateOperationApiKeyRequestPermissions) SetDashboardFolder(v string)`

SetDashboardFolder sets DashboardFolder field to given value.

### HasDashboardFolder

`func (o *CreateOperationApiKeyRequestPermissions) HasDashboardFolder() bool`

HasDashboardFolder returns a boolean if a field has been set.

### GetMonitorFolder

`func (o *CreateOperationApiKeyRequestPermissions) GetMonitorFolder() string`

GetMonitorFolder returns the MonitorFolder field if non-nil, zero value otherwise.

### GetMonitorFolderOk

`func (o *CreateOperationApiKeyRequestPermissions) GetMonitorFolderOk() (*string, bool)`

GetMonitorFolderOk returns a tuple with the MonitorFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorFolder

`func (o *CreateOperationApiKeyRequestPermissions) SetMonitorFolder(v string)`

SetMonitorFolder sets MonitorFolder field to given value.

### HasMonitorFolder

`func (o *CreateOperationApiKeyRequestPermissions) HasMonitorFolder() bool`

HasMonitorFolder returns a boolean if a field has been set.

### GetDocumentation

`func (o *CreateOperationApiKeyRequestPermissions) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *CreateOperationApiKeyRequestPermissions) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *CreateOperationApiKeyRequestPermissions) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *CreateOperationApiKeyRequestPermissions) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.

### GetIngestionApiKey

`func (o *CreateOperationApiKeyRequestPermissions) GetIngestionApiKey() string`

GetIngestionApiKey returns the IngestionApiKey field if non-nil, zero value otherwise.

### GetIngestionApiKeyOk

`func (o *CreateOperationApiKeyRequestPermissions) GetIngestionApiKeyOk() (*string, bool)`

GetIngestionApiKeyOk returns a tuple with the IngestionApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestionApiKey

`func (o *CreateOperationApiKeyRequestPermissions) SetIngestionApiKey(v string)`

SetIngestionApiKey sets IngestionApiKey field to given value.

### HasIngestionApiKey

`func (o *CreateOperationApiKeyRequestPermissions) HasIngestionApiKey() bool`

HasIngestionApiKey returns a boolean if a field has been set.

### GetRumPublicToken

`func (o *CreateOperationApiKeyRequestPermissions) GetRumPublicToken() string`

GetRumPublicToken returns the RumPublicToken field if non-nil, zero value otherwise.

### GetRumPublicTokenOk

`func (o *CreateOperationApiKeyRequestPermissions) GetRumPublicTokenOk() (*string, bool)`

GetRumPublicTokenOk returns a tuple with the RumPublicToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRumPublicToken

`func (o *CreateOperationApiKeyRequestPermissions) SetRumPublicToken(v string)`

SetRumPublicToken sets RumPublicToken field to given value.

### HasRumPublicToken

`func (o *CreateOperationApiKeyRequestPermissions) HasRumPublicToken() bool`

HasRumPublicToken returns a boolean if a field has been set.

### GetOperationApiKey

`func (o *CreateOperationApiKeyRequestPermissions) GetOperationApiKey() string`

GetOperationApiKey returns the OperationApiKey field if non-nil, zero value otherwise.

### GetOperationApiKeyOk

`func (o *CreateOperationApiKeyRequestPermissions) GetOperationApiKeyOk() (*string, bool)`

GetOperationApiKeyOk returns a tuple with the OperationApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationApiKey

`func (o *CreateOperationApiKeyRequestPermissions) SetOperationApiKey(v string)`

SetOperationApiKey sets OperationApiKey field to given value.

### HasOperationApiKey

`func (o *CreateOperationApiKeyRequestPermissions) HasOperationApiKey() bool`

HasOperationApiKey returns a boolean if a field has been set.

### GetInventory

`func (o *CreateOperationApiKeyRequestPermissions) GetInventory() string`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *CreateOperationApiKeyRequestPermissions) GetInventoryOk() (*string, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *CreateOperationApiKeyRequestPermissions) SetInventory(v string)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *CreateOperationApiKeyRequestPermissions) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetCloudAccount

`func (o *CreateOperationApiKeyRequestPermissions) GetCloudAccount() string`

GetCloudAccount returns the CloudAccount field if non-nil, zero value otherwise.

### GetCloudAccountOk

`func (o *CreateOperationApiKeyRequestPermissions) GetCloudAccountOk() (*string, bool)`

GetCloudAccountOk returns a tuple with the CloudAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAccount

`func (o *CreateOperationApiKeyRequestPermissions) SetCloudAccount(v string)`

SetCloudAccount sets CloudAccount field to given value.

### HasCloudAccount

`func (o *CreateOperationApiKeyRequestPermissions) HasCloudAccount() bool`

HasCloudAccount returns a boolean if a field has been set.

### GetLogs

`func (o *CreateOperationApiKeyRequestPermissions) GetLogs() string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *CreateOperationApiKeyRequestPermissions) GetLogsOk() (*string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *CreateOperationApiKeyRequestPermissions) SetLogs(v string)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *CreateOperationApiKeyRequestPermissions) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetLogRoute

`func (o *CreateOperationApiKeyRequestPermissions) GetLogRoute() string`

GetLogRoute returns the LogRoute field if non-nil, zero value otherwise.

### GetLogRouteOk

`func (o *CreateOperationApiKeyRequestPermissions) GetLogRouteOk() (*string, bool)`

GetLogRouteOk returns a tuple with the LogRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRoute

`func (o *CreateOperationApiKeyRequestPermissions) SetLogRoute(v string)`

SetLogRoute sets LogRoute field to given value.

### HasLogRoute

`func (o *CreateOperationApiKeyRequestPermissions) HasLogRoute() bool`

HasLogRoute returns a boolean if a field has been set.

### GetMetrics

`func (o *CreateOperationApiKeyRequestPermissions) GetMetrics() string`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *CreateOperationApiKeyRequestPermissions) GetMetricsOk() (*string, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *CreateOperationApiKeyRequestPermissions) SetMetrics(v string)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *CreateOperationApiKeyRequestPermissions) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetMetricRoute

`func (o *CreateOperationApiKeyRequestPermissions) GetMetricRoute() string`

GetMetricRoute returns the MetricRoute field if non-nil, zero value otherwise.

### GetMetricRouteOk

`func (o *CreateOperationApiKeyRequestPermissions) GetMetricRouteOk() (*string, bool)`

GetMetricRouteOk returns a tuple with the MetricRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricRoute

`func (o *CreateOperationApiKeyRequestPermissions) SetMetricRoute(v string)`

SetMetricRoute sets MetricRoute field to given value.

### HasMetricRoute

`func (o *CreateOperationApiKeyRequestPermissions) HasMetricRoute() bool`

HasMetricRoute returns a boolean if a field has been set.

### GetMonitor

`func (o *CreateOperationApiKeyRequestPermissions) GetMonitor() string`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *CreateOperationApiKeyRequestPermissions) GetMonitorOk() (*string, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *CreateOperationApiKeyRequestPermissions) SetMonitor(v string)`

SetMonitor sets Monitor field to given value.

### HasMonitor

`func (o *CreateOperationApiKeyRequestPermissions) HasMonitor() bool`

HasMonitor returns a boolean if a field has been set.

### GetAllowedDomains

`func (o *CreateOperationApiKeyRequestPermissions) GetAllowedDomains() string`

GetAllowedDomains returns the AllowedDomains field if non-nil, zero value otherwise.

### GetAllowedDomainsOk

`func (o *CreateOperationApiKeyRequestPermissions) GetAllowedDomainsOk() (*string, bool)`

GetAllowedDomainsOk returns a tuple with the AllowedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDomains

`func (o *CreateOperationApiKeyRequestPermissions) SetAllowedDomains(v string)`

SetAllowedDomains sets AllowedDomains field to given value.

### HasAllowedDomains

`func (o *CreateOperationApiKeyRequestPermissions) HasAllowedDomains() bool`

HasAllowedDomains returns a boolean if a field has been set.

### GetQueue

`func (o *CreateOperationApiKeyRequestPermissions) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *CreateOperationApiKeyRequestPermissions) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *CreateOperationApiKeyRequestPermissions) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *CreateOperationApiKeyRequestPermissions) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### GetQualityReport

`func (o *CreateOperationApiKeyRequestPermissions) GetQualityReport() string`

GetQualityReport returns the QualityReport field if non-nil, zero value otherwise.

### GetQualityReportOk

`func (o *CreateOperationApiKeyRequestPermissions) GetQualityReportOk() (*string, bool)`

GetQualityReportOk returns a tuple with the QualityReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityReport

`func (o *CreateOperationApiKeyRequestPermissions) SetQualityReport(v string)`

SetQualityReport sets QualityReport field to given value.

### HasQualityReport

`func (o *CreateOperationApiKeyRequestPermissions) HasQualityReport() bool`

HasQualityReport returns a boolean if a field has been set.

### GetBilling

`func (o *CreateOperationApiKeyRequestPermissions) GetBilling() string`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *CreateOperationApiKeyRequestPermissions) GetBillingOk() (*string, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *CreateOperationApiKeyRequestPermissions) SetBilling(v string)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *CreateOperationApiKeyRequestPermissions) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### GetService

`func (o *CreateOperationApiKeyRequestPermissions) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CreateOperationApiKeyRequestPermissions) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CreateOperationApiKeyRequestPermissions) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *CreateOperationApiKeyRequestPermissions) HasService() bool`

HasService returns a boolean if a field has been set.

### GetServiceTarget

`func (o *CreateOperationApiKeyRequestPermissions) GetServiceTarget() string`

GetServiceTarget returns the ServiceTarget field if non-nil, zero value otherwise.

### GetServiceTargetOk

`func (o *CreateOperationApiKeyRequestPermissions) GetServiceTargetOk() (*string, bool)`

GetServiceTargetOk returns a tuple with the ServiceTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTarget

`func (o *CreateOperationApiKeyRequestPermissions) SetServiceTarget(v string)`

SetServiceTarget sets ServiceTarget field to given value.

### HasServiceTarget

`func (o *CreateOperationApiKeyRequestPermissions) HasServiceTarget() bool`

HasServiceTarget returns a boolean if a field has been set.

### GetSlo

`func (o *CreateOperationApiKeyRequestPermissions) GetSlo() string`

GetSlo returns the Slo field if non-nil, zero value otherwise.

### GetSloOk

`func (o *CreateOperationApiKeyRequestPermissions) GetSloOk() (*string, bool)`

GetSloOk returns a tuple with the Slo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlo

`func (o *CreateOperationApiKeyRequestPermissions) SetSlo(v string)`

SetSlo sets Slo field to given value.

### HasSlo

`func (o *CreateOperationApiKeyRequestPermissions) HasSlo() bool`

HasSlo returns a boolean if a field has been set.

### GetNotificationRule

`func (o *CreateOperationApiKeyRequestPermissions) GetNotificationRule() string`

GetNotificationRule returns the NotificationRule field if non-nil, zero value otherwise.

### GetNotificationRuleOk

`func (o *CreateOperationApiKeyRequestPermissions) GetNotificationRuleOk() (*string, bool)`

GetNotificationRuleOk returns a tuple with the NotificationRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationRule

`func (o *CreateOperationApiKeyRequestPermissions) SetNotificationRule(v string)`

SetNotificationRule sets NotificationRule field to given value.

### HasNotificationRule

`func (o *CreateOperationApiKeyRequestPermissions) HasNotificationRule() bool`

HasNotificationRule returns a boolean if a field has been set.

### GetAdaptiveNotificationRule

`func (o *CreateOperationApiKeyRequestPermissions) GetAdaptiveNotificationRule() string`

GetAdaptiveNotificationRule returns the AdaptiveNotificationRule field if non-nil, zero value otherwise.

### GetAdaptiveNotificationRuleOk

`func (o *CreateOperationApiKeyRequestPermissions) GetAdaptiveNotificationRuleOk() (*string, bool)`

GetAdaptiveNotificationRuleOk returns a tuple with the AdaptiveNotificationRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdaptiveNotificationRule

`func (o *CreateOperationApiKeyRequestPermissions) SetAdaptiveNotificationRule(v string)`

SetAdaptiveNotificationRule sets AdaptiveNotificationRule field to given value.

### HasAdaptiveNotificationRule

`func (o *CreateOperationApiKeyRequestPermissions) HasAdaptiveNotificationRule() bool`

HasAdaptiveNotificationRule returns a boolean if a field has been set.

### GetNotificationIntegration

`func (o *CreateOperationApiKeyRequestPermissions) GetNotificationIntegration() string`

GetNotificationIntegration returns the NotificationIntegration field if non-nil, zero value otherwise.

### GetNotificationIntegrationOk

`func (o *CreateOperationApiKeyRequestPermissions) GetNotificationIntegrationOk() (*string, bool)`

GetNotificationIntegrationOk returns a tuple with the NotificationIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationIntegration

`func (o *CreateOperationApiKeyRequestPermissions) SetNotificationIntegration(v string)`

SetNotificationIntegration sets NotificationIntegration field to given value.

### HasNotificationIntegration

`func (o *CreateOperationApiKeyRequestPermissions) HasNotificationIntegration() bool`

HasNotificationIntegration returns a boolean if a field has been set.

### GetNotificationSilence

`func (o *CreateOperationApiKeyRequestPermissions) GetNotificationSilence() string`

GetNotificationSilence returns the NotificationSilence field if non-nil, zero value otherwise.

### GetNotificationSilenceOk

`func (o *CreateOperationApiKeyRequestPermissions) GetNotificationSilenceOk() (*string, bool)`

GetNotificationSilenceOk returns a tuple with the NotificationSilence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSilence

`func (o *CreateOperationApiKeyRequestPermissions) SetNotificationSilence(v string)`

SetNotificationSilence sets NotificationSilence field to given value.

### HasNotificationSilence

`func (o *CreateOperationApiKeyRequestPermissions) HasNotificationSilence() bool`

HasNotificationSilence returns a boolean if a field has been set.

### GetCustomUsageTag

`func (o *CreateOperationApiKeyRequestPermissions) GetCustomUsageTag() string`

GetCustomUsageTag returns the CustomUsageTag field if non-nil, zero value otherwise.

### GetCustomUsageTagOk

`func (o *CreateOperationApiKeyRequestPermissions) GetCustomUsageTagOk() (*string, bool)`

GetCustomUsageTagOk returns a tuple with the CustomUsageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUsageTag

`func (o *CreateOperationApiKeyRequestPermissions) SetCustomUsageTag(v string)`

SetCustomUsageTag sets CustomUsageTag field to given value.

### HasCustomUsageTag

`func (o *CreateOperationApiKeyRequestPermissions) HasCustomUsageTag() bool`

HasCustomUsageTag returns a boolean if a field has been set.

### GetRetentionPolicy

`func (o *CreateOperationApiKeyRequestPermissions) GetRetentionPolicy() string`

GetRetentionPolicy returns the RetentionPolicy field if non-nil, zero value otherwise.

### GetRetentionPolicyOk

`func (o *CreateOperationApiKeyRequestPermissions) GetRetentionPolicyOk() (*string, bool)`

GetRetentionPolicyOk returns a tuple with the RetentionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicy

`func (o *CreateOperationApiKeyRequestPermissions) SetRetentionPolicy(v string)`

SetRetentionPolicy sets RetentionPolicy field to given value.

### HasRetentionPolicy

`func (o *CreateOperationApiKeyRequestPermissions) HasRetentionPolicy() bool`

HasRetentionPolicy returns a boolean if a field has been set.

### GetSandboxConfiguration

`func (o *CreateOperationApiKeyRequestPermissions) GetSandboxConfiguration() string`

GetSandboxConfiguration returns the SandboxConfiguration field if non-nil, zero value otherwise.

### GetSandboxConfigurationOk

`func (o *CreateOperationApiKeyRequestPermissions) GetSandboxConfigurationOk() (*string, bool)`

GetSandboxConfigurationOk returns a tuple with the SandboxConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandboxConfiguration

`func (o *CreateOperationApiKeyRequestPermissions) SetSandboxConfiguration(v string)`

SetSandboxConfiguration sets SandboxConfiguration field to given value.

### HasSandboxConfiguration

`func (o *CreateOperationApiKeyRequestPermissions) HasSandboxConfiguration() bool`

HasSandboxConfiguration returns a boolean if a field has been set.

### GetTagPolicy

`func (o *CreateOperationApiKeyRequestPermissions) GetTagPolicy() string`

GetTagPolicy returns the TagPolicy field if non-nil, zero value otherwise.

### GetTagPolicyOk

`func (o *CreateOperationApiKeyRequestPermissions) GetTagPolicyOk() (*string, bool)`

GetTagPolicyOk returns a tuple with the TagPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagPolicy

`func (o *CreateOperationApiKeyRequestPermissions) SetTagPolicy(v string)`

SetTagPolicy sets TagPolicy field to given value.

### HasTagPolicy

`func (o *CreateOperationApiKeyRequestPermissions) HasTagPolicy() bool`

HasTagPolicy returns a boolean if a field has been set.

### GetTeam

`func (o *CreateOperationApiKeyRequestPermissions) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *CreateOperationApiKeyRequestPermissions) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *CreateOperationApiKeyRequestPermissions) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *CreateOperationApiKeyRequestPermissions) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetTraces

`func (o *CreateOperationApiKeyRequestPermissions) GetTraces() string`

GetTraces returns the Traces field if non-nil, zero value otherwise.

### GetTracesOk

`func (o *CreateOperationApiKeyRequestPermissions) GetTracesOk() (*string, bool)`

GetTracesOk returns a tuple with the Traces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraces

`func (o *CreateOperationApiKeyRequestPermissions) SetTraces(v string)`

SetTraces sets Traces field to given value.

### HasTraces

`func (o *CreateOperationApiKeyRequestPermissions) HasTraces() bool`

HasTraces returns a boolean if a field has been set.

### GetRum

`func (o *CreateOperationApiKeyRequestPermissions) GetRum() string`

GetRum returns the Rum field if non-nil, zero value otherwise.

### GetRumOk

`func (o *CreateOperationApiKeyRequestPermissions) GetRumOk() (*string, bool)`

GetRumOk returns a tuple with the Rum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRum

`func (o *CreateOperationApiKeyRequestPermissions) SetRum(v string)`

SetRum sets Rum field to given value.

### HasRum

`func (o *CreateOperationApiKeyRequestPermissions) HasRum() bool`

HasRum returns a boolean if a field has been set.

### GetUser

`func (o *CreateOperationApiKeyRequestPermissions) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CreateOperationApiKeyRequestPermissions) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CreateOperationApiKeyRequestPermissions) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *CreateOperationApiKeyRequestPermissions) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserInvitation

`func (o *CreateOperationApiKeyRequestPermissions) GetUserInvitation() string`

GetUserInvitation returns the UserInvitation field if non-nil, zero value otherwise.

### GetUserInvitationOk

`func (o *CreateOperationApiKeyRequestPermissions) GetUserInvitationOk() (*string, bool)`

GetUserInvitationOk returns a tuple with the UserInvitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInvitation

`func (o *CreateOperationApiKeyRequestPermissions) SetUserInvitation(v string)`

SetUserInvitation sets UserInvitation field to given value.

### HasUserInvitation

`func (o *CreateOperationApiKeyRequestPermissions) HasUserInvitation() bool`

HasUserInvitation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


