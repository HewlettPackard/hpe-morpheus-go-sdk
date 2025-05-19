# LicenseFeatures

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dashboard** | Pointer to **bool** |  | [optional] 
**Guidance** | Pointer to **bool** |  | [optional] 
**Discovery** | Pointer to **bool** |  | [optional] 
**Analytics** | Pointer to **bool** |  | [optional] 
**Scheduling** | Pointer to **bool** |  | [optional] 
**Approvals** | Pointer to **bool** |  | [optional] 
**Usage** | Pointer to **bool** |  | [optional] 
**Activity** | Pointer to **bool** |  | [optional] 
**Instances** | Pointer to **bool** |  | [optional] 
**Apps** | Pointer to **bool** |  | [optional] 
**Templates** | Pointer to **bool** |  | [optional] 
**Automation** | Pointer to **bool** |  | [optional] 
**VirtualImages** | Pointer to **bool** |  | [optional] 
**Library** | Pointer to **bool** |  | [optional] 
**Migrations** | Pointer to **bool** |  | [optional] 
**Deployments** | Pointer to **bool** |  | [optional] 
**Groups** | Pointer to **bool** |  | [optional] 
**Clouds** | Pointer to **bool** |  | [optional] 
**Hosts** | Pointer to **bool** |  | [optional] 
**Network** | Pointer to **bool** |  | [optional] 
**LoadBalancers** | Pointer to **bool** |  | [optional] 
**Storage** | Pointer to **bool** |  | [optional] 
**KeyPairs** | Pointer to **bool** |  | [optional] 
**SslCertificates** | Pointer to **bool** |  | [optional] 
**Boot** | Pointer to **bool** |  | [optional] 
**Backups** | Pointer to **bool** |  | [optional] 
**Cypher** | Pointer to **bool** |  | [optional] 
**Archives** | Pointer to **bool** |  | [optional] 
**ImageBuilder** | Pointer to **bool** |  | [optional] 
**Tenants** | Pointer to **bool** |  | [optional] 
**Plans** | Pointer to **bool** |  | [optional] 
**Pricing** | Pointer to **bool** |  | [optional] 
**Users** | Pointer to **bool** |  | [optional] 
**UserGroups** | Pointer to **bool** |  | [optional] 
**Monitoring** | Pointer to **bool** |  | [optional] 
**Logging** | Pointer to **bool** |  | [optional] 
**MonitoringServices** | Pointer to **bool** |  | [optional] 
**LoggingServices** | Pointer to **bool** |  | [optional] 
**BackupServices** | Pointer to **bool** |  | [optional] 
**DnsServices** | Pointer to **bool** |  | [optional] 
**CodeService** | Pointer to **bool** |  | [optional] 
**BuildServices** | Pointer to **bool** |  | [optional] 
**LoadBalancerServices** | Pointer to **bool** |  | [optional] 
**IpamServices** | Pointer to **bool** |  | [optional] 
**ApprovalServices** | Pointer to **bool** |  | [optional] 
**CmdbServices** | Pointer to **bool** |  | [optional] 
**DeploymentServices** | Pointer to **bool** |  | [optional] 
**AutomationServices** | Pointer to **bool** |  | [optional] 
**ServiceDiscoveryServices** | Pointer to **bool** |  | [optional] 
**IdentityServices** | Pointer to **bool** |  | [optional] 
**TrustServices** | Pointer to **bool** |  | [optional] 
**SecurityServices** | Pointer to **bool** |  | [optional] 

## Methods

### NewLicenseFeatures

`func NewLicenseFeatures() *LicenseFeatures`

NewLicenseFeatures instantiates a new LicenseFeatures object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseFeaturesWithDefaults

`func NewLicenseFeaturesWithDefaults() *LicenseFeatures`

NewLicenseFeaturesWithDefaults instantiates a new LicenseFeatures object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDashboard

`func (o *LicenseFeatures) GetDashboard() bool`

GetDashboard returns the Dashboard field if non-nil, zero value otherwise.

### GetDashboardOk

`func (o *LicenseFeatures) GetDashboardOk() (*bool, bool)`

GetDashboardOk returns a tuple with the Dashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboard

`func (o *LicenseFeatures) SetDashboard(v bool)`

SetDashboard sets Dashboard field to given value.

### HasDashboard

`func (o *LicenseFeatures) HasDashboard() bool`

HasDashboard returns a boolean if a field has been set.

### GetGuidance

`func (o *LicenseFeatures) GetGuidance() bool`

GetGuidance returns the Guidance field if non-nil, zero value otherwise.

### GetGuidanceOk

`func (o *LicenseFeatures) GetGuidanceOk() (*bool, bool)`

GetGuidanceOk returns a tuple with the Guidance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuidance

`func (o *LicenseFeatures) SetGuidance(v bool)`

SetGuidance sets Guidance field to given value.

### HasGuidance

`func (o *LicenseFeatures) HasGuidance() bool`

HasGuidance returns a boolean if a field has been set.

### GetDiscovery

`func (o *LicenseFeatures) GetDiscovery() bool`

GetDiscovery returns the Discovery field if non-nil, zero value otherwise.

### GetDiscoveryOk

`func (o *LicenseFeatures) GetDiscoveryOk() (*bool, bool)`

GetDiscoveryOk returns a tuple with the Discovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovery

`func (o *LicenseFeatures) SetDiscovery(v bool)`

SetDiscovery sets Discovery field to given value.

### HasDiscovery

`func (o *LicenseFeatures) HasDiscovery() bool`

HasDiscovery returns a boolean if a field has been set.

### GetAnalytics

`func (o *LicenseFeatures) GetAnalytics() bool`

GetAnalytics returns the Analytics field if non-nil, zero value otherwise.

### GetAnalyticsOk

`func (o *LicenseFeatures) GetAnalyticsOk() (*bool, bool)`

GetAnalyticsOk returns a tuple with the Analytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalytics

`func (o *LicenseFeatures) SetAnalytics(v bool)`

SetAnalytics sets Analytics field to given value.

### HasAnalytics

`func (o *LicenseFeatures) HasAnalytics() bool`

HasAnalytics returns a boolean if a field has been set.

### GetScheduling

`func (o *LicenseFeatures) GetScheduling() bool`

GetScheduling returns the Scheduling field if non-nil, zero value otherwise.

### GetSchedulingOk

`func (o *LicenseFeatures) GetSchedulingOk() (*bool, bool)`

GetSchedulingOk returns a tuple with the Scheduling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduling

`func (o *LicenseFeatures) SetScheduling(v bool)`

SetScheduling sets Scheduling field to given value.

### HasScheduling

`func (o *LicenseFeatures) HasScheduling() bool`

HasScheduling returns a boolean if a field has been set.

### GetApprovals

`func (o *LicenseFeatures) GetApprovals() bool`

GetApprovals returns the Approvals field if non-nil, zero value otherwise.

### GetApprovalsOk

`func (o *LicenseFeatures) GetApprovalsOk() (*bool, bool)`

GetApprovalsOk returns a tuple with the Approvals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovals

`func (o *LicenseFeatures) SetApprovals(v bool)`

SetApprovals sets Approvals field to given value.

### HasApprovals

`func (o *LicenseFeatures) HasApprovals() bool`

HasApprovals returns a boolean if a field has been set.

### GetUsage

`func (o *LicenseFeatures) GetUsage() bool`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *LicenseFeatures) GetUsageOk() (*bool, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *LicenseFeatures) SetUsage(v bool)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *LicenseFeatures) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetActivity

`func (o *LicenseFeatures) GetActivity() bool`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *LicenseFeatures) GetActivityOk() (*bool, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *LicenseFeatures) SetActivity(v bool)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *LicenseFeatures) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetInstances

`func (o *LicenseFeatures) GetInstances() bool`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *LicenseFeatures) GetInstancesOk() (*bool, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *LicenseFeatures) SetInstances(v bool)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *LicenseFeatures) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetApps

`func (o *LicenseFeatures) GetApps() bool`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *LicenseFeatures) GetAppsOk() (*bool, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *LicenseFeatures) SetApps(v bool)`

SetApps sets Apps field to given value.

### HasApps

`func (o *LicenseFeatures) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetTemplates

`func (o *LicenseFeatures) GetTemplates() bool`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *LicenseFeatures) GetTemplatesOk() (*bool, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *LicenseFeatures) SetTemplates(v bool)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *LicenseFeatures) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.

### GetAutomation

`func (o *LicenseFeatures) GetAutomation() bool`

GetAutomation returns the Automation field if non-nil, zero value otherwise.

### GetAutomationOk

`func (o *LicenseFeatures) GetAutomationOk() (*bool, bool)`

GetAutomationOk returns a tuple with the Automation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomation

`func (o *LicenseFeatures) SetAutomation(v bool)`

SetAutomation sets Automation field to given value.

### HasAutomation

`func (o *LicenseFeatures) HasAutomation() bool`

HasAutomation returns a boolean if a field has been set.

### GetVirtualImages

`func (o *LicenseFeatures) GetVirtualImages() bool`

GetVirtualImages returns the VirtualImages field if non-nil, zero value otherwise.

### GetVirtualImagesOk

`func (o *LicenseFeatures) GetVirtualImagesOk() (*bool, bool)`

GetVirtualImagesOk returns a tuple with the VirtualImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImages

`func (o *LicenseFeatures) SetVirtualImages(v bool)`

SetVirtualImages sets VirtualImages field to given value.

### HasVirtualImages

`func (o *LicenseFeatures) HasVirtualImages() bool`

HasVirtualImages returns a boolean if a field has been set.

### GetLibrary

`func (o *LicenseFeatures) GetLibrary() bool`

GetLibrary returns the Library field if non-nil, zero value otherwise.

### GetLibraryOk

`func (o *LicenseFeatures) GetLibraryOk() (*bool, bool)`

GetLibraryOk returns a tuple with the Library field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibrary

`func (o *LicenseFeatures) SetLibrary(v bool)`

SetLibrary sets Library field to given value.

### HasLibrary

`func (o *LicenseFeatures) HasLibrary() bool`

HasLibrary returns a boolean if a field has been set.

### GetMigrations

`func (o *LicenseFeatures) GetMigrations() bool`

GetMigrations returns the Migrations field if non-nil, zero value otherwise.

### GetMigrationsOk

`func (o *LicenseFeatures) GetMigrationsOk() (*bool, bool)`

GetMigrationsOk returns a tuple with the Migrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrations

`func (o *LicenseFeatures) SetMigrations(v bool)`

SetMigrations sets Migrations field to given value.

### HasMigrations

`func (o *LicenseFeatures) HasMigrations() bool`

HasMigrations returns a boolean if a field has been set.

### GetDeployments

`func (o *LicenseFeatures) GetDeployments() bool`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *LicenseFeatures) GetDeploymentsOk() (*bool, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *LicenseFeatures) SetDeployments(v bool)`

SetDeployments sets Deployments field to given value.

### HasDeployments

`func (o *LicenseFeatures) HasDeployments() bool`

HasDeployments returns a boolean if a field has been set.

### GetGroups

`func (o *LicenseFeatures) GetGroups() bool`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *LicenseFeatures) GetGroupsOk() (*bool, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *LicenseFeatures) SetGroups(v bool)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *LicenseFeatures) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetClouds

`func (o *LicenseFeatures) GetClouds() bool`

GetClouds returns the Clouds field if non-nil, zero value otherwise.

### GetCloudsOk

`func (o *LicenseFeatures) GetCloudsOk() (*bool, bool)`

GetCloudsOk returns a tuple with the Clouds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClouds

`func (o *LicenseFeatures) SetClouds(v bool)`

SetClouds sets Clouds field to given value.

### HasClouds

`func (o *LicenseFeatures) HasClouds() bool`

HasClouds returns a boolean if a field has been set.

### GetHosts

`func (o *LicenseFeatures) GetHosts() bool`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *LicenseFeatures) GetHostsOk() (*bool, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *LicenseFeatures) SetHosts(v bool)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *LicenseFeatures) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetNetwork

`func (o *LicenseFeatures) GetNetwork() bool`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *LicenseFeatures) GetNetworkOk() (*bool, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *LicenseFeatures) SetNetwork(v bool)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *LicenseFeatures) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetLoadBalancers

`func (o *LicenseFeatures) GetLoadBalancers() bool`

GetLoadBalancers returns the LoadBalancers field if non-nil, zero value otherwise.

### GetLoadBalancersOk

`func (o *LicenseFeatures) GetLoadBalancersOk() (*bool, bool)`

GetLoadBalancersOk returns a tuple with the LoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancers

`func (o *LicenseFeatures) SetLoadBalancers(v bool)`

SetLoadBalancers sets LoadBalancers field to given value.

### HasLoadBalancers

`func (o *LicenseFeatures) HasLoadBalancers() bool`

HasLoadBalancers returns a boolean if a field has been set.

### GetStorage

`func (o *LicenseFeatures) GetStorage() bool`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *LicenseFeatures) GetStorageOk() (*bool, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *LicenseFeatures) SetStorage(v bool)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *LicenseFeatures) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetKeyPairs

`func (o *LicenseFeatures) GetKeyPairs() bool`

GetKeyPairs returns the KeyPairs field if non-nil, zero value otherwise.

### GetKeyPairsOk

`func (o *LicenseFeatures) GetKeyPairsOk() (*bool, bool)`

GetKeyPairsOk returns a tuple with the KeyPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPairs

`func (o *LicenseFeatures) SetKeyPairs(v bool)`

SetKeyPairs sets KeyPairs field to given value.

### HasKeyPairs

`func (o *LicenseFeatures) HasKeyPairs() bool`

HasKeyPairs returns a boolean if a field has been set.

### GetSslCertificates

`func (o *LicenseFeatures) GetSslCertificates() bool`

GetSslCertificates returns the SslCertificates field if non-nil, zero value otherwise.

### GetSslCertificatesOk

`func (o *LicenseFeatures) GetSslCertificatesOk() (*bool, bool)`

GetSslCertificatesOk returns a tuple with the SslCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificates

`func (o *LicenseFeatures) SetSslCertificates(v bool)`

SetSslCertificates sets SslCertificates field to given value.

### HasSslCertificates

`func (o *LicenseFeatures) HasSslCertificates() bool`

HasSslCertificates returns a boolean if a field has been set.

### GetBoot

`func (o *LicenseFeatures) GetBoot() bool`

GetBoot returns the Boot field if non-nil, zero value otherwise.

### GetBootOk

`func (o *LicenseFeatures) GetBootOk() (*bool, bool)`

GetBootOk returns a tuple with the Boot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoot

`func (o *LicenseFeatures) SetBoot(v bool)`

SetBoot sets Boot field to given value.

### HasBoot

`func (o *LicenseFeatures) HasBoot() bool`

HasBoot returns a boolean if a field has been set.

### GetBackups

`func (o *LicenseFeatures) GetBackups() bool`

GetBackups returns the Backups field if non-nil, zero value otherwise.

### GetBackupsOk

`func (o *LicenseFeatures) GetBackupsOk() (*bool, bool)`

GetBackupsOk returns a tuple with the Backups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackups

`func (o *LicenseFeatures) SetBackups(v bool)`

SetBackups sets Backups field to given value.

### HasBackups

`func (o *LicenseFeatures) HasBackups() bool`

HasBackups returns a boolean if a field has been set.

### GetCypher

`func (o *LicenseFeatures) GetCypher() bool`

GetCypher returns the Cypher field if non-nil, zero value otherwise.

### GetCypherOk

`func (o *LicenseFeatures) GetCypherOk() (*bool, bool)`

GetCypherOk returns a tuple with the Cypher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCypher

`func (o *LicenseFeatures) SetCypher(v bool)`

SetCypher sets Cypher field to given value.

### HasCypher

`func (o *LicenseFeatures) HasCypher() bool`

HasCypher returns a boolean if a field has been set.

### GetArchives

`func (o *LicenseFeatures) GetArchives() bool`

GetArchives returns the Archives field if non-nil, zero value otherwise.

### GetArchivesOk

`func (o *LicenseFeatures) GetArchivesOk() (*bool, bool)`

GetArchivesOk returns a tuple with the Archives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchives

`func (o *LicenseFeatures) SetArchives(v bool)`

SetArchives sets Archives field to given value.

### HasArchives

`func (o *LicenseFeatures) HasArchives() bool`

HasArchives returns a boolean if a field has been set.

### GetImageBuilder

`func (o *LicenseFeatures) GetImageBuilder() bool`

GetImageBuilder returns the ImageBuilder field if non-nil, zero value otherwise.

### GetImageBuilderOk

`func (o *LicenseFeatures) GetImageBuilderOk() (*bool, bool)`

GetImageBuilderOk returns a tuple with the ImageBuilder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBuilder

`func (o *LicenseFeatures) SetImageBuilder(v bool)`

SetImageBuilder sets ImageBuilder field to given value.

### HasImageBuilder

`func (o *LicenseFeatures) HasImageBuilder() bool`

HasImageBuilder returns a boolean if a field has been set.

### GetTenants

`func (o *LicenseFeatures) GetTenants() bool`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *LicenseFeatures) GetTenantsOk() (*bool, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *LicenseFeatures) SetTenants(v bool)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *LicenseFeatures) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetPlans

`func (o *LicenseFeatures) GetPlans() bool`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *LicenseFeatures) GetPlansOk() (*bool, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *LicenseFeatures) SetPlans(v bool)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *LicenseFeatures) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### GetPricing

`func (o *LicenseFeatures) GetPricing() bool`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *LicenseFeatures) GetPricingOk() (*bool, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *LicenseFeatures) SetPricing(v bool)`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *LicenseFeatures) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### GetUsers

`func (o *LicenseFeatures) GetUsers() bool`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *LicenseFeatures) GetUsersOk() (*bool, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *LicenseFeatures) SetUsers(v bool)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *LicenseFeatures) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetUserGroups

`func (o *LicenseFeatures) GetUserGroups() bool`

GetUserGroups returns the UserGroups field if non-nil, zero value otherwise.

### GetUserGroupsOk

`func (o *LicenseFeatures) GetUserGroupsOk() (*bool, bool)`

GetUserGroupsOk returns a tuple with the UserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroups

`func (o *LicenseFeatures) SetUserGroups(v bool)`

SetUserGroups sets UserGroups field to given value.

### HasUserGroups

`func (o *LicenseFeatures) HasUserGroups() bool`

HasUserGroups returns a boolean if a field has been set.

### GetMonitoring

`func (o *LicenseFeatures) GetMonitoring() bool`

GetMonitoring returns the Monitoring field if non-nil, zero value otherwise.

### GetMonitoringOk

`func (o *LicenseFeatures) GetMonitoringOk() (*bool, bool)`

GetMonitoringOk returns a tuple with the Monitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoring

`func (o *LicenseFeatures) SetMonitoring(v bool)`

SetMonitoring sets Monitoring field to given value.

### HasMonitoring

`func (o *LicenseFeatures) HasMonitoring() bool`

HasMonitoring returns a boolean if a field has been set.

### GetLogging

`func (o *LicenseFeatures) GetLogging() bool`

GetLogging returns the Logging field if non-nil, zero value otherwise.

### GetLoggingOk

`func (o *LicenseFeatures) GetLoggingOk() (*bool, bool)`

GetLoggingOk returns a tuple with the Logging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogging

`func (o *LicenseFeatures) SetLogging(v bool)`

SetLogging sets Logging field to given value.

### HasLogging

`func (o *LicenseFeatures) HasLogging() bool`

HasLogging returns a boolean if a field has been set.

### GetMonitoringServices

`func (o *LicenseFeatures) GetMonitoringServices() bool`

GetMonitoringServices returns the MonitoringServices field if non-nil, zero value otherwise.

### GetMonitoringServicesOk

`func (o *LicenseFeatures) GetMonitoringServicesOk() (*bool, bool)`

GetMonitoringServicesOk returns a tuple with the MonitoringServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringServices

`func (o *LicenseFeatures) SetMonitoringServices(v bool)`

SetMonitoringServices sets MonitoringServices field to given value.

### HasMonitoringServices

`func (o *LicenseFeatures) HasMonitoringServices() bool`

HasMonitoringServices returns a boolean if a field has been set.

### GetLoggingServices

`func (o *LicenseFeatures) GetLoggingServices() bool`

GetLoggingServices returns the LoggingServices field if non-nil, zero value otherwise.

### GetLoggingServicesOk

`func (o *LicenseFeatures) GetLoggingServicesOk() (*bool, bool)`

GetLoggingServicesOk returns a tuple with the LoggingServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingServices

`func (o *LicenseFeatures) SetLoggingServices(v bool)`

SetLoggingServices sets LoggingServices field to given value.

### HasLoggingServices

`func (o *LicenseFeatures) HasLoggingServices() bool`

HasLoggingServices returns a boolean if a field has been set.

### GetBackupServices

`func (o *LicenseFeatures) GetBackupServices() bool`

GetBackupServices returns the BackupServices field if non-nil, zero value otherwise.

### GetBackupServicesOk

`func (o *LicenseFeatures) GetBackupServicesOk() (*bool, bool)`

GetBackupServicesOk returns a tuple with the BackupServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupServices

`func (o *LicenseFeatures) SetBackupServices(v bool)`

SetBackupServices sets BackupServices field to given value.

### HasBackupServices

`func (o *LicenseFeatures) HasBackupServices() bool`

HasBackupServices returns a boolean if a field has been set.

### GetDnsServices

`func (o *LicenseFeatures) GetDnsServices() bool`

GetDnsServices returns the DnsServices field if non-nil, zero value otherwise.

### GetDnsServicesOk

`func (o *LicenseFeatures) GetDnsServicesOk() (*bool, bool)`

GetDnsServicesOk returns a tuple with the DnsServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServices

`func (o *LicenseFeatures) SetDnsServices(v bool)`

SetDnsServices sets DnsServices field to given value.

### HasDnsServices

`func (o *LicenseFeatures) HasDnsServices() bool`

HasDnsServices returns a boolean if a field has been set.

### GetCodeService

`func (o *LicenseFeatures) GetCodeService() bool`

GetCodeService returns the CodeService field if non-nil, zero value otherwise.

### GetCodeServiceOk

`func (o *LicenseFeatures) GetCodeServiceOk() (*bool, bool)`

GetCodeServiceOk returns a tuple with the CodeService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeService

`func (o *LicenseFeatures) SetCodeService(v bool)`

SetCodeService sets CodeService field to given value.

### HasCodeService

`func (o *LicenseFeatures) HasCodeService() bool`

HasCodeService returns a boolean if a field has been set.

### GetBuildServices

`func (o *LicenseFeatures) GetBuildServices() bool`

GetBuildServices returns the BuildServices field if non-nil, zero value otherwise.

### GetBuildServicesOk

`func (o *LicenseFeatures) GetBuildServicesOk() (*bool, bool)`

GetBuildServicesOk returns a tuple with the BuildServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildServices

`func (o *LicenseFeatures) SetBuildServices(v bool)`

SetBuildServices sets BuildServices field to given value.

### HasBuildServices

`func (o *LicenseFeatures) HasBuildServices() bool`

HasBuildServices returns a boolean if a field has been set.

### GetLoadBalancerServices

`func (o *LicenseFeatures) GetLoadBalancerServices() bool`

GetLoadBalancerServices returns the LoadBalancerServices field if non-nil, zero value otherwise.

### GetLoadBalancerServicesOk

`func (o *LicenseFeatures) GetLoadBalancerServicesOk() (*bool, bool)`

GetLoadBalancerServicesOk returns a tuple with the LoadBalancerServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerServices

`func (o *LicenseFeatures) SetLoadBalancerServices(v bool)`

SetLoadBalancerServices sets LoadBalancerServices field to given value.

### HasLoadBalancerServices

`func (o *LicenseFeatures) HasLoadBalancerServices() bool`

HasLoadBalancerServices returns a boolean if a field has been set.

### GetIpamServices

`func (o *LicenseFeatures) GetIpamServices() bool`

GetIpamServices returns the IpamServices field if non-nil, zero value otherwise.

### GetIpamServicesOk

`func (o *LicenseFeatures) GetIpamServicesOk() (*bool, bool)`

GetIpamServicesOk returns a tuple with the IpamServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpamServices

`func (o *LicenseFeatures) SetIpamServices(v bool)`

SetIpamServices sets IpamServices field to given value.

### HasIpamServices

`func (o *LicenseFeatures) HasIpamServices() bool`

HasIpamServices returns a boolean if a field has been set.

### GetApprovalServices

`func (o *LicenseFeatures) GetApprovalServices() bool`

GetApprovalServices returns the ApprovalServices field if non-nil, zero value otherwise.

### GetApprovalServicesOk

`func (o *LicenseFeatures) GetApprovalServicesOk() (*bool, bool)`

GetApprovalServicesOk returns a tuple with the ApprovalServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalServices

`func (o *LicenseFeatures) SetApprovalServices(v bool)`

SetApprovalServices sets ApprovalServices field to given value.

### HasApprovalServices

`func (o *LicenseFeatures) HasApprovalServices() bool`

HasApprovalServices returns a boolean if a field has been set.

### GetCmdbServices

`func (o *LicenseFeatures) GetCmdbServices() bool`

GetCmdbServices returns the CmdbServices field if non-nil, zero value otherwise.

### GetCmdbServicesOk

`func (o *LicenseFeatures) GetCmdbServicesOk() (*bool, bool)`

GetCmdbServicesOk returns a tuple with the CmdbServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmdbServices

`func (o *LicenseFeatures) SetCmdbServices(v bool)`

SetCmdbServices sets CmdbServices field to given value.

### HasCmdbServices

`func (o *LicenseFeatures) HasCmdbServices() bool`

HasCmdbServices returns a boolean if a field has been set.

### GetDeploymentServices

`func (o *LicenseFeatures) GetDeploymentServices() bool`

GetDeploymentServices returns the DeploymentServices field if non-nil, zero value otherwise.

### GetDeploymentServicesOk

`func (o *LicenseFeatures) GetDeploymentServicesOk() (*bool, bool)`

GetDeploymentServicesOk returns a tuple with the DeploymentServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentServices

`func (o *LicenseFeatures) SetDeploymentServices(v bool)`

SetDeploymentServices sets DeploymentServices field to given value.

### HasDeploymentServices

`func (o *LicenseFeatures) HasDeploymentServices() bool`

HasDeploymentServices returns a boolean if a field has been set.

### GetAutomationServices

`func (o *LicenseFeatures) GetAutomationServices() bool`

GetAutomationServices returns the AutomationServices field if non-nil, zero value otherwise.

### GetAutomationServicesOk

`func (o *LicenseFeatures) GetAutomationServicesOk() (*bool, bool)`

GetAutomationServicesOk returns a tuple with the AutomationServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomationServices

`func (o *LicenseFeatures) SetAutomationServices(v bool)`

SetAutomationServices sets AutomationServices field to given value.

### HasAutomationServices

`func (o *LicenseFeatures) HasAutomationServices() bool`

HasAutomationServices returns a boolean if a field has been set.

### GetServiceDiscoveryServices

`func (o *LicenseFeatures) GetServiceDiscoveryServices() bool`

GetServiceDiscoveryServices returns the ServiceDiscoveryServices field if non-nil, zero value otherwise.

### GetServiceDiscoveryServicesOk

`func (o *LicenseFeatures) GetServiceDiscoveryServicesOk() (*bool, bool)`

GetServiceDiscoveryServicesOk returns a tuple with the ServiceDiscoveryServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDiscoveryServices

`func (o *LicenseFeatures) SetServiceDiscoveryServices(v bool)`

SetServiceDiscoveryServices sets ServiceDiscoveryServices field to given value.

### HasServiceDiscoveryServices

`func (o *LicenseFeatures) HasServiceDiscoveryServices() bool`

HasServiceDiscoveryServices returns a boolean if a field has been set.

### GetIdentityServices

`func (o *LicenseFeatures) GetIdentityServices() bool`

GetIdentityServices returns the IdentityServices field if non-nil, zero value otherwise.

### GetIdentityServicesOk

`func (o *LicenseFeatures) GetIdentityServicesOk() (*bool, bool)`

GetIdentityServicesOk returns a tuple with the IdentityServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityServices

`func (o *LicenseFeatures) SetIdentityServices(v bool)`

SetIdentityServices sets IdentityServices field to given value.

### HasIdentityServices

`func (o *LicenseFeatures) HasIdentityServices() bool`

HasIdentityServices returns a boolean if a field has been set.

### GetTrustServices

`func (o *LicenseFeatures) GetTrustServices() bool`

GetTrustServices returns the TrustServices field if non-nil, zero value otherwise.

### GetTrustServicesOk

`func (o *LicenseFeatures) GetTrustServicesOk() (*bool, bool)`

GetTrustServicesOk returns a tuple with the TrustServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustServices

`func (o *LicenseFeatures) SetTrustServices(v bool)`

SetTrustServices sets TrustServices field to given value.

### HasTrustServices

`func (o *LicenseFeatures) HasTrustServices() bool`

HasTrustServices returns a boolean if a field has been set.

### GetSecurityServices

`func (o *LicenseFeatures) GetSecurityServices() bool`

GetSecurityServices returns the SecurityServices field if non-nil, zero value otherwise.

### GetSecurityServicesOk

`func (o *LicenseFeatures) GetSecurityServicesOk() (*bool, bool)`

GetSecurityServicesOk returns a tuple with the SecurityServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityServices

`func (o *LicenseFeatures) SetSecurityServices(v bool)`

SetSecurityServices sets SecurityServices field to given value.

### HasSecurityServices

`func (o *LicenseFeatures) HasSecurityServices() bool`

HasSecurityServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


