# Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ServiceUrl** | Pointer to **string** |  | [optional] 
**ServiceHost** | Pointer to **string** |  | [optional] 
**ServicePath** | Pointer to **string** |  | [optional] 
**ServiceHostname** | Pointer to **string** |  | [optional] 
**ServicePort** | Pointer to **int64** |  | [optional] 
**ServiceUsername** | Pointer to **string** |  | [optional] 
**ServicePassword** | Pointer to **string** |  | [optional] 
**ServicePasswordHash** | Pointer to **string** |  | [optional] 
**ServiceToken** | Pointer to **string** |  | [optional] 
**ServiceTokenHash** | Pointer to **string** |  | [optional] 
**ServiceAccess** | Pointer to **string** |  | [optional] 
**ServiceAccessHash** | Pointer to **string** |  | [optional] 
**ServiceCert** | Pointer to **string** |  | [optional] 
**ServiceCertHash** | Pointer to **string** |  | [optional] 
**ServiceVersion** | Pointer to **string** |  | [optional] 
**SearchDomains** | Pointer to **string** |  | [optional] 
**EnableInternalDns** | Pointer to **bool** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**DatacenterId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusDate** | Pointer to **time.Time** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**InventoryLevel** | Pointer to **string** |  | [optional] 
**LastSync** | Pointer to **time.Time** |  | [optional] 
**NextRunDate** | Pointer to **time.Time** |  | [optional] 
**LastSyncDuration** | Pointer to **int64** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Managed** | Pointer to **bool** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**AutoRecoverPowerState** | Pointer to **bool** | Automatically Power on VMs | [optional] [default to false]
**UseAgent** | Pointer to **string** | Use the Agent to relay communications for the Kubernetes API instead of direct. | [optional] 
**ProvisionComplete** | Pointer to **bool** | Changes from false to true once provisioning is finished. | [optional] 
**ServiceEntry** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to [**ListActivity200ResponseAllOfActivityInnerActivityInnerUser**](ListActivity200ResponseAllOfActivityInnerActivityInnerUser.md) |  | [optional] 
**UserGroup** | Pointer to **string** |  | [optional] 
**Layout** | Pointer to [**ListClusters200ResponseAllOfClustersInnerLayout**](ListClusters200ResponseAllOfClustersInnerLayout.md) |  | [optional] 
**Owner** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Servers** | Pointer to [**[]ListClusters200ResponseAllOfClustersInnerServersInner**](ListClusters200ResponseAllOfClustersInnerServersInner.md) |  | [optional] 
**Accounts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Integrations** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Site** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Type** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Zone** | Pointer to [**ListClusters200ResponseAllOfClustersInnerZone**](ListClusters200ResponseAllOfClustersInnerZone.md) |  | [optional] 
**WorkerStats** | Pointer to [**ListClusters200ResponseAllOfClustersInnerWorkerStats**](ListClusters200ResponseAllOfClustersInnerWorkerStats.md) |  | [optional] 
**ContainersCount** | Pointer to **int64** |  | [optional] 
**DeploymentsCount** | Pointer to **int64** |  | [optional] 
**PodsCount** | Pointer to **int64** |  | [optional] 
**JobsCount** | Pointer to **int64** |  | [optional] 
**VolumesCount** | Pointer to **int64** |  | [optional] 
**NamespacesCount** | Pointer to **int64** |  | [optional] 
**WorkersCount** | Pointer to **int64** |  | [optional] 
**ServicesCount** | Pointer to **int64** |  | [optional] 
**Permissions** | Pointer to [**AddCluster200ResponseAllOfClusterPermissions**](AddCluster200ResponseAllOfClusterPermissions.md) |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCluster

`func NewCluster() *Cluster`

NewCluster instantiates a new Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWithDefaults

`func NewClusterWithDefaults() *Cluster`

NewClusterWithDefaults instantiates a new Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Cluster) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cluster) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cluster) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Cluster) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *Cluster) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Cluster) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Cluster) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Cluster) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *Cluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Cluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *Cluster) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Cluster) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Cluster) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Cluster) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCategory

`func (o *Cluster) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Cluster) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Cluster) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Cluster) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetVisibility

`func (o *Cluster) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Cluster) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Cluster) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *Cluster) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *Cluster) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Cluster) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Cluster) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Cluster) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocation

`func (o *Cluster) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Cluster) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Cluster) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Cluster) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetEnabled

`func (o *Cluster) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Cluster) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Cluster) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Cluster) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetServiceUrl

`func (o *Cluster) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *Cluster) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *Cluster) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *Cluster) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.

### GetServiceHost

`func (o *Cluster) GetServiceHost() string`

GetServiceHost returns the ServiceHost field if non-nil, zero value otherwise.

### GetServiceHostOk

`func (o *Cluster) GetServiceHostOk() (*string, bool)`

GetServiceHostOk returns a tuple with the ServiceHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHost

`func (o *Cluster) SetServiceHost(v string)`

SetServiceHost sets ServiceHost field to given value.

### HasServiceHost

`func (o *Cluster) HasServiceHost() bool`

HasServiceHost returns a boolean if a field has been set.

### GetServicePath

`func (o *Cluster) GetServicePath() string`

GetServicePath returns the ServicePath field if non-nil, zero value otherwise.

### GetServicePathOk

`func (o *Cluster) GetServicePathOk() (*string, bool)`

GetServicePathOk returns a tuple with the ServicePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePath

`func (o *Cluster) SetServicePath(v string)`

SetServicePath sets ServicePath field to given value.

### HasServicePath

`func (o *Cluster) HasServicePath() bool`

HasServicePath returns a boolean if a field has been set.

### GetServiceHostname

`func (o *Cluster) GetServiceHostname() string`

GetServiceHostname returns the ServiceHostname field if non-nil, zero value otherwise.

### GetServiceHostnameOk

`func (o *Cluster) GetServiceHostnameOk() (*string, bool)`

GetServiceHostnameOk returns a tuple with the ServiceHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHostname

`func (o *Cluster) SetServiceHostname(v string)`

SetServiceHostname sets ServiceHostname field to given value.

### HasServiceHostname

`func (o *Cluster) HasServiceHostname() bool`

HasServiceHostname returns a boolean if a field has been set.

### GetServicePort

`func (o *Cluster) GetServicePort() int64`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *Cluster) GetServicePortOk() (*int64, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *Cluster) SetServicePort(v int64)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *Cluster) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### GetServiceUsername

`func (o *Cluster) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *Cluster) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *Cluster) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *Cluster) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### GetServicePassword

`func (o *Cluster) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *Cluster) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *Cluster) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *Cluster) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### GetServicePasswordHash

`func (o *Cluster) GetServicePasswordHash() string`

GetServicePasswordHash returns the ServicePasswordHash field if non-nil, zero value otherwise.

### GetServicePasswordHashOk

`func (o *Cluster) GetServicePasswordHashOk() (*string, bool)`

GetServicePasswordHashOk returns a tuple with the ServicePasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePasswordHash

`func (o *Cluster) SetServicePasswordHash(v string)`

SetServicePasswordHash sets ServicePasswordHash field to given value.

### HasServicePasswordHash

`func (o *Cluster) HasServicePasswordHash() bool`

HasServicePasswordHash returns a boolean if a field has been set.

### GetServiceToken

`func (o *Cluster) GetServiceToken() string`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *Cluster) GetServiceTokenOk() (*string, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *Cluster) SetServiceToken(v string)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *Cluster) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### GetServiceTokenHash

`func (o *Cluster) GetServiceTokenHash() string`

GetServiceTokenHash returns the ServiceTokenHash field if non-nil, zero value otherwise.

### GetServiceTokenHashOk

`func (o *Cluster) GetServiceTokenHashOk() (*string, bool)`

GetServiceTokenHashOk returns a tuple with the ServiceTokenHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTokenHash

`func (o *Cluster) SetServiceTokenHash(v string)`

SetServiceTokenHash sets ServiceTokenHash field to given value.

### HasServiceTokenHash

`func (o *Cluster) HasServiceTokenHash() bool`

HasServiceTokenHash returns a boolean if a field has been set.

### GetServiceAccess

`func (o *Cluster) GetServiceAccess() string`

GetServiceAccess returns the ServiceAccess field if non-nil, zero value otherwise.

### GetServiceAccessOk

`func (o *Cluster) GetServiceAccessOk() (*string, bool)`

GetServiceAccessOk returns a tuple with the ServiceAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccess

`func (o *Cluster) SetServiceAccess(v string)`

SetServiceAccess sets ServiceAccess field to given value.

### HasServiceAccess

`func (o *Cluster) HasServiceAccess() bool`

HasServiceAccess returns a boolean if a field has been set.

### GetServiceAccessHash

`func (o *Cluster) GetServiceAccessHash() string`

GetServiceAccessHash returns the ServiceAccessHash field if non-nil, zero value otherwise.

### GetServiceAccessHashOk

`func (o *Cluster) GetServiceAccessHashOk() (*string, bool)`

GetServiceAccessHashOk returns a tuple with the ServiceAccessHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccessHash

`func (o *Cluster) SetServiceAccessHash(v string)`

SetServiceAccessHash sets ServiceAccessHash field to given value.

### HasServiceAccessHash

`func (o *Cluster) HasServiceAccessHash() bool`

HasServiceAccessHash returns a boolean if a field has been set.

### GetServiceCert

`func (o *Cluster) GetServiceCert() string`

GetServiceCert returns the ServiceCert field if non-nil, zero value otherwise.

### GetServiceCertOk

`func (o *Cluster) GetServiceCertOk() (*string, bool)`

GetServiceCertOk returns a tuple with the ServiceCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCert

`func (o *Cluster) SetServiceCert(v string)`

SetServiceCert sets ServiceCert field to given value.

### HasServiceCert

`func (o *Cluster) HasServiceCert() bool`

HasServiceCert returns a boolean if a field has been set.

### GetServiceCertHash

`func (o *Cluster) GetServiceCertHash() string`

GetServiceCertHash returns the ServiceCertHash field if non-nil, zero value otherwise.

### GetServiceCertHashOk

`func (o *Cluster) GetServiceCertHashOk() (*string, bool)`

GetServiceCertHashOk returns a tuple with the ServiceCertHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCertHash

`func (o *Cluster) SetServiceCertHash(v string)`

SetServiceCertHash sets ServiceCertHash field to given value.

### HasServiceCertHash

`func (o *Cluster) HasServiceCertHash() bool`

HasServiceCertHash returns a boolean if a field has been set.

### GetServiceVersion

`func (o *Cluster) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *Cluster) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *Cluster) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *Cluster) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetSearchDomains

`func (o *Cluster) GetSearchDomains() string`

GetSearchDomains returns the SearchDomains field if non-nil, zero value otherwise.

### GetSearchDomainsOk

`func (o *Cluster) GetSearchDomainsOk() (*string, bool)`

GetSearchDomainsOk returns a tuple with the SearchDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchDomains

`func (o *Cluster) SetSearchDomains(v string)`

SetSearchDomains sets SearchDomains field to given value.

### HasSearchDomains

`func (o *Cluster) HasSearchDomains() bool`

HasSearchDomains returns a boolean if a field has been set.

### GetEnableInternalDns

`func (o *Cluster) GetEnableInternalDns() bool`

GetEnableInternalDns returns the EnableInternalDns field if non-nil, zero value otherwise.

### GetEnableInternalDnsOk

`func (o *Cluster) GetEnableInternalDnsOk() (*bool, bool)`

GetEnableInternalDnsOk returns a tuple with the EnableInternalDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInternalDns

`func (o *Cluster) SetEnableInternalDns(v bool)`

SetEnableInternalDns sets EnableInternalDns field to given value.

### HasEnableInternalDns

`func (o *Cluster) HasEnableInternalDns() bool`

HasEnableInternalDns returns a boolean if a field has been set.

### GetInternalId

`func (o *Cluster) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *Cluster) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *Cluster) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *Cluster) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *Cluster) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Cluster) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Cluster) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Cluster) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetDatacenterId

`func (o *Cluster) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *Cluster) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *Cluster) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.

### HasDatacenterId

`func (o *Cluster) HasDatacenterId() bool`

HasDatacenterId returns a boolean if a field has been set.

### GetStatus

`func (o *Cluster) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Cluster) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Cluster) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Cluster) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDate

`func (o *Cluster) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *Cluster) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *Cluster) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *Cluster) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetStatusMessage

`func (o *Cluster) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *Cluster) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *Cluster) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *Cluster) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetInventoryLevel

`func (o *Cluster) GetInventoryLevel() string`

GetInventoryLevel returns the InventoryLevel field if non-nil, zero value otherwise.

### GetInventoryLevelOk

`func (o *Cluster) GetInventoryLevelOk() (*string, bool)`

GetInventoryLevelOk returns a tuple with the InventoryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryLevel

`func (o *Cluster) SetInventoryLevel(v string)`

SetInventoryLevel sets InventoryLevel field to given value.

### HasInventoryLevel

`func (o *Cluster) HasInventoryLevel() bool`

HasInventoryLevel returns a boolean if a field has been set.

### GetLastSync

`func (o *Cluster) GetLastSync() time.Time`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *Cluster) GetLastSyncOk() (*time.Time, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *Cluster) SetLastSync(v time.Time)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *Cluster) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetNextRunDate

`func (o *Cluster) GetNextRunDate() time.Time`

GetNextRunDate returns the NextRunDate field if non-nil, zero value otherwise.

### GetNextRunDateOk

`func (o *Cluster) GetNextRunDateOk() (*time.Time, bool)`

GetNextRunDateOk returns a tuple with the NextRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunDate

`func (o *Cluster) SetNextRunDate(v time.Time)`

SetNextRunDate sets NextRunDate field to given value.

### HasNextRunDate

`func (o *Cluster) HasNextRunDate() bool`

HasNextRunDate returns a boolean if a field has been set.

### GetLastSyncDuration

`func (o *Cluster) GetLastSyncDuration() int64`

GetLastSyncDuration returns the LastSyncDuration field if non-nil, zero value otherwise.

### GetLastSyncDurationOk

`func (o *Cluster) GetLastSyncDurationOk() (*int64, bool)`

GetLastSyncDurationOk returns a tuple with the LastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDuration

`func (o *Cluster) SetLastSyncDuration(v int64)`

SetLastSyncDuration sets LastSyncDuration field to given value.

### HasLastSyncDuration

`func (o *Cluster) HasLastSyncDuration() bool`

HasLastSyncDuration returns a boolean if a field has been set.

### GetDateCreated

`func (o *Cluster) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Cluster) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Cluster) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Cluster) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Cluster) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Cluster) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Cluster) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Cluster) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetManaged

`func (o *Cluster) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *Cluster) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *Cluster) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *Cluster) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetLabels

`func (o *Cluster) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Cluster) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Cluster) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Cluster) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAutoRecoverPowerState

`func (o *Cluster) GetAutoRecoverPowerState() bool`

GetAutoRecoverPowerState returns the AutoRecoverPowerState field if non-nil, zero value otherwise.

### GetAutoRecoverPowerStateOk

`func (o *Cluster) GetAutoRecoverPowerStateOk() (*bool, bool)`

GetAutoRecoverPowerStateOk returns a tuple with the AutoRecoverPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRecoverPowerState

`func (o *Cluster) SetAutoRecoverPowerState(v bool)`

SetAutoRecoverPowerState sets AutoRecoverPowerState field to given value.

### HasAutoRecoverPowerState

`func (o *Cluster) HasAutoRecoverPowerState() bool`

HasAutoRecoverPowerState returns a boolean if a field has been set.

### GetUseAgent

`func (o *Cluster) GetUseAgent() string`

GetUseAgent returns the UseAgent field if non-nil, zero value otherwise.

### GetUseAgentOk

`func (o *Cluster) GetUseAgentOk() (*string, bool)`

GetUseAgentOk returns a tuple with the UseAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAgent

`func (o *Cluster) SetUseAgent(v string)`

SetUseAgent sets UseAgent field to given value.

### HasUseAgent

`func (o *Cluster) HasUseAgent() bool`

HasUseAgent returns a boolean if a field has been set.

### GetProvisionComplete

`func (o *Cluster) GetProvisionComplete() bool`

GetProvisionComplete returns the ProvisionComplete field if non-nil, zero value otherwise.

### GetProvisionCompleteOk

`func (o *Cluster) GetProvisionCompleteOk() (*bool, bool)`

GetProvisionCompleteOk returns a tuple with the ProvisionComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionComplete

`func (o *Cluster) SetProvisionComplete(v bool)`

SetProvisionComplete sets ProvisionComplete field to given value.

### HasProvisionComplete

`func (o *Cluster) HasProvisionComplete() bool`

HasProvisionComplete returns a boolean if a field has been set.

### GetServiceEntry

`func (o *Cluster) GetServiceEntry() string`

GetServiceEntry returns the ServiceEntry field if non-nil, zero value otherwise.

### GetServiceEntryOk

`func (o *Cluster) GetServiceEntryOk() (*string, bool)`

GetServiceEntryOk returns a tuple with the ServiceEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEntry

`func (o *Cluster) SetServiceEntry(v string)`

SetServiceEntry sets ServiceEntry field to given value.

### HasServiceEntry

`func (o *Cluster) HasServiceEntry() bool`

HasServiceEntry returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Cluster) GetCreatedBy() ListActivity200ResponseAllOfActivityInnerActivityInnerUser`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Cluster) GetCreatedByOk() (*ListActivity200ResponseAllOfActivityInnerActivityInnerUser, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Cluster) SetCreatedBy(v ListActivity200ResponseAllOfActivityInnerActivityInnerUser)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Cluster) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUserGroup

`func (o *Cluster) GetUserGroup() string`

GetUserGroup returns the UserGroup field if non-nil, zero value otherwise.

### GetUserGroupOk

`func (o *Cluster) GetUserGroupOk() (*string, bool)`

GetUserGroupOk returns a tuple with the UserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroup

`func (o *Cluster) SetUserGroup(v string)`

SetUserGroup sets UserGroup field to given value.

### HasUserGroup

`func (o *Cluster) HasUserGroup() bool`

HasUserGroup returns a boolean if a field has been set.

### GetLayout

`func (o *Cluster) GetLayout() ListClusters200ResponseAllOfClustersInnerLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *Cluster) GetLayoutOk() (*ListClusters200ResponseAllOfClustersInnerLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *Cluster) SetLayout(v ListClusters200ResponseAllOfClustersInnerLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *Cluster) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetOwner

`func (o *Cluster) GetOwner() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Cluster) GetOwnerOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Cluster) SetOwner(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Cluster) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetServers

`func (o *Cluster) GetServers() []ListClusters200ResponseAllOfClustersInnerServersInner`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *Cluster) GetServersOk() (*[]ListClusters200ResponseAllOfClustersInnerServersInner, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *Cluster) SetServers(v []ListClusters200ResponseAllOfClustersInnerServersInner)`

SetServers sets Servers field to given value.

### HasServers

`func (o *Cluster) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetAccounts

`func (o *Cluster) GetAccounts() []map[string]interface{}`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *Cluster) GetAccountsOk() (*[]map[string]interface{}, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *Cluster) SetAccounts(v []map[string]interface{})`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *Cluster) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetIntegrations

`func (o *Cluster) GetIntegrations() []map[string]interface{}`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *Cluster) GetIntegrationsOk() (*[]map[string]interface{}, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *Cluster) SetIntegrations(v []map[string]interface{})`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *Cluster) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.

### GetSite

`func (o *Cluster) GetSite() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *Cluster) GetSiteOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *Cluster) SetSite(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetSite sets Site field to given value.

### HasSite

`func (o *Cluster) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetType

`func (o *Cluster) GetType() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Cluster) GetTypeOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Cluster) SetType(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetType sets Type field to given value.

### HasType

`func (o *Cluster) HasType() bool`

HasType returns a boolean if a field has been set.

### GetZone

`func (o *Cluster) GetZone() ListClusters200ResponseAllOfClustersInnerZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *Cluster) GetZoneOk() (*ListClusters200ResponseAllOfClustersInnerZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *Cluster) SetZone(v ListClusters200ResponseAllOfClustersInnerZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *Cluster) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetWorkerStats

`func (o *Cluster) GetWorkerStats() ListClusters200ResponseAllOfClustersInnerWorkerStats`

GetWorkerStats returns the WorkerStats field if non-nil, zero value otherwise.

### GetWorkerStatsOk

`func (o *Cluster) GetWorkerStatsOk() (*ListClusters200ResponseAllOfClustersInnerWorkerStats, bool)`

GetWorkerStatsOk returns a tuple with the WorkerStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerStats

`func (o *Cluster) SetWorkerStats(v ListClusters200ResponseAllOfClustersInnerWorkerStats)`

SetWorkerStats sets WorkerStats field to given value.

### HasWorkerStats

`func (o *Cluster) HasWorkerStats() bool`

HasWorkerStats returns a boolean if a field has been set.

### GetContainersCount

`func (o *Cluster) GetContainersCount() int64`

GetContainersCount returns the ContainersCount field if non-nil, zero value otherwise.

### GetContainersCountOk

`func (o *Cluster) GetContainersCountOk() (*int64, bool)`

GetContainersCountOk returns a tuple with the ContainersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainersCount

`func (o *Cluster) SetContainersCount(v int64)`

SetContainersCount sets ContainersCount field to given value.

### HasContainersCount

`func (o *Cluster) HasContainersCount() bool`

HasContainersCount returns a boolean if a field has been set.

### GetDeploymentsCount

`func (o *Cluster) GetDeploymentsCount() int64`

GetDeploymentsCount returns the DeploymentsCount field if non-nil, zero value otherwise.

### GetDeploymentsCountOk

`func (o *Cluster) GetDeploymentsCountOk() (*int64, bool)`

GetDeploymentsCountOk returns a tuple with the DeploymentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentsCount

`func (o *Cluster) SetDeploymentsCount(v int64)`

SetDeploymentsCount sets DeploymentsCount field to given value.

### HasDeploymentsCount

`func (o *Cluster) HasDeploymentsCount() bool`

HasDeploymentsCount returns a boolean if a field has been set.

### GetPodsCount

`func (o *Cluster) GetPodsCount() int64`

GetPodsCount returns the PodsCount field if non-nil, zero value otherwise.

### GetPodsCountOk

`func (o *Cluster) GetPodsCountOk() (*int64, bool)`

GetPodsCountOk returns a tuple with the PodsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodsCount

`func (o *Cluster) SetPodsCount(v int64)`

SetPodsCount sets PodsCount field to given value.

### HasPodsCount

`func (o *Cluster) HasPodsCount() bool`

HasPodsCount returns a boolean if a field has been set.

### GetJobsCount

`func (o *Cluster) GetJobsCount() int64`

GetJobsCount returns the JobsCount field if non-nil, zero value otherwise.

### GetJobsCountOk

`func (o *Cluster) GetJobsCountOk() (*int64, bool)`

GetJobsCountOk returns a tuple with the JobsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsCount

`func (o *Cluster) SetJobsCount(v int64)`

SetJobsCount sets JobsCount field to given value.

### HasJobsCount

`func (o *Cluster) HasJobsCount() bool`

HasJobsCount returns a boolean if a field has been set.

### GetVolumesCount

`func (o *Cluster) GetVolumesCount() int64`

GetVolumesCount returns the VolumesCount field if non-nil, zero value otherwise.

### GetVolumesCountOk

`func (o *Cluster) GetVolumesCountOk() (*int64, bool)`

GetVolumesCountOk returns a tuple with the VolumesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumesCount

`func (o *Cluster) SetVolumesCount(v int64)`

SetVolumesCount sets VolumesCount field to given value.

### HasVolumesCount

`func (o *Cluster) HasVolumesCount() bool`

HasVolumesCount returns a boolean if a field has been set.

### GetNamespacesCount

`func (o *Cluster) GetNamespacesCount() int64`

GetNamespacesCount returns the NamespacesCount field if non-nil, zero value otherwise.

### GetNamespacesCountOk

`func (o *Cluster) GetNamespacesCountOk() (*int64, bool)`

GetNamespacesCountOk returns a tuple with the NamespacesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespacesCount

`func (o *Cluster) SetNamespacesCount(v int64)`

SetNamespacesCount sets NamespacesCount field to given value.

### HasNamespacesCount

`func (o *Cluster) HasNamespacesCount() bool`

HasNamespacesCount returns a boolean if a field has been set.

### GetWorkersCount

`func (o *Cluster) GetWorkersCount() int64`

GetWorkersCount returns the WorkersCount field if non-nil, zero value otherwise.

### GetWorkersCountOk

`func (o *Cluster) GetWorkersCountOk() (*int64, bool)`

GetWorkersCountOk returns a tuple with the WorkersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkersCount

`func (o *Cluster) SetWorkersCount(v int64)`

SetWorkersCount sets WorkersCount field to given value.

### HasWorkersCount

`func (o *Cluster) HasWorkersCount() bool`

HasWorkersCount returns a boolean if a field has been set.

### GetServicesCount

`func (o *Cluster) GetServicesCount() int64`

GetServicesCount returns the ServicesCount field if non-nil, zero value otherwise.

### GetServicesCountOk

`func (o *Cluster) GetServicesCountOk() (*int64, bool)`

GetServicesCountOk returns a tuple with the ServicesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesCount

`func (o *Cluster) SetServicesCount(v int64)`

SetServicesCount sets ServicesCount field to given value.

### HasServicesCount

`func (o *Cluster) HasServicesCount() bool`

HasServicesCount returns a boolean if a field has been set.

### GetPermissions

`func (o *Cluster) GetPermissions() AddCluster200ResponseAllOfClusterPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *Cluster) GetPermissionsOk() (*AddCluster200ResponseAllOfClusterPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *Cluster) SetPermissions(v AddCluster200ResponseAllOfClusterPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *Cluster) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetConfig

`func (o *Cluster) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Cluster) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Cluster) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Cluster) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


