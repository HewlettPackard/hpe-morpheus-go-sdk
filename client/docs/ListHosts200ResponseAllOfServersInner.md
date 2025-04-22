# ListHosts200ResponseAllOfServersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**ExternalUniqueId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ExternalName** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**ParentServer** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Owner** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Zone** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Plan** | Pointer to [**ListHosts200ResponseAllOfServersInnerPlan**](ListHosts200ResponseAllOfServersInnerPlan.md) |  | [optional] 
**ComputeServerType** | Pointer to [**ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType**](ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType.md) |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ZoneId** | Pointer to **int64** |  | [optional] 
**SiteId** | Pointer to **int64** |  | [optional] 
**ResourcePoolId** | Pointer to **int64** |  | [optional] 
**FolderId** | Pointer to **int64** |  | [optional] 
**SshHost** | Pointer to **string** |  | [optional] 
**SshPort** | Pointer to **int64** |  | [optional] 
**ExternalIp** | Pointer to **string** |  | [optional] 
**InternalIp** | Pointer to **string** |  | [optional] 
**VolumeId** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**PlatformVersion** | Pointer to **string** |  | [optional] 
**SshUsername** | Pointer to **string** |  | [optional] 
**SshPassword** | Pointer to **string** |  | [optional] 
**SshPasswordHash** | Pointer to **string** |  | [optional] 
**SshKeyPair** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerContainer**](GetAlerts200ResponseAllOfChecksInnerContainer.md) |  | [optional] 
**OsDevice** | Pointer to **string** |  | [optional] 
**OsType** | Pointer to **string** |  | [optional] 
**DataDevice** | Pointer to **string** |  | [optional] 
**LvmEnabled** | Pointer to **bool** |  | [optional] 
**ApiKey** | Pointer to **string** |  | [optional] 
**SoftwareRaid** | Pointer to **bool** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Stats** | Pointer to [**ListHosts200ResponseAllOfServersInnerStats**](ListHosts200ResponseAllOfServersInnerStats.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**StatusDate** | Pointer to **time.Time** |  | [optional] 
**StatusPercent** | Pointer to **string** |  | [optional] 
**StatusEta** | Pointer to **string** |  | [optional] 
**PowerState** | Pointer to **string** |  | [optional] 
**AgentInstalled** | Pointer to **bool** |  | [optional] 
**LastAgentUpdate** | Pointer to **string** |  | [optional] 
**AgentVersion** | Pointer to **string** |  | [optional] 
**MaxCores** | Pointer to **int64** |  | [optional] 
**CoresPerSocket** | Pointer to **int64** |  | [optional] 
**MaxMemory** | Pointer to **int64** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 
**MaxCpu** | Pointer to **int64** |  | [optional] 
**MaxGpus** | Pointer to **int64** |  | [optional] 
**ManageInternalFirewall** | Pointer to **bool** |  | [optional] 
**EnableLogs** | Pointer to **bool** |  | [optional] 
**HourlyCost** | Pointer to **float32** |  | [optional] 
**HourlyPrice** | Pointer to **float32** |  | [optional] 
**SourceImage** | Pointer to [**ListPriceSets200ResponseAllOfPriceSetsInnerPricesInnerVolumeType**](ListPriceSets200ResponseAllOfPriceSetsInnerPricesInnerVolumeType.md) |  | [optional] 
**ServerOs** | Pointer to [**ListHosts200ResponseAllOfServersInnerServerOs**](ListHosts200ResponseAllOfServersInnerServerOs.md) |  | [optional] 
**Volumes** | Pointer to [**[]ListHosts200ResponseAllOfServersInnerVolumesInner**](ListHosts200ResponseAllOfServersInnerVolumesInner.md) |  | [optional] 
**Controllers** | Pointer to [**[]ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceControllersInner**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceControllersInner.md) |  | [optional] 
**Interfaces** | Pointer to [**[]ListHosts200ResponseAllOfServersInnerInterfacesInner**](ListHosts200ResponseAllOfServersInnerInterfacesInner.md) |  | [optional] 
**Labels** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Tags** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**TagCompliant** | Pointer to **string** |  | [optional] 
**Containers** | Pointer to **[]int64** |  | [optional] 
**Config** | Pointer to [**ListHosts200ResponseAllOfServersInnerConfig**](ListHosts200ResponseAllOfServersInnerConfig.md) |  | [optional] 
**GuestConsolePreferred** | Pointer to **bool** |  | [optional] 
**GuestConsoleType** | Pointer to **string** |  | [optional] 
**GuestConsoleUsername** | Pointer to **string** |  | [optional] 
**GuestConsolePassword** | Pointer to **string** |  | [optional] 
**GuestConsolePasswordHash** | Pointer to **string** |  | [optional] 
**GuestConsolePort** | Pointer to **string** |  | [optional] 

## Methods

### NewListHosts200ResponseAllOfServersInner

`func NewListHosts200ResponseAllOfServersInner() *ListHosts200ResponseAllOfServersInner`

NewListHosts200ResponseAllOfServersInner instantiates a new ListHosts200ResponseAllOfServersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHosts200ResponseAllOfServersInnerWithDefaults

`func NewListHosts200ResponseAllOfServersInnerWithDefaults() *ListHosts200ResponseAllOfServersInner`

NewListHosts200ResponseAllOfServersInnerWithDefaults instantiates a new ListHosts200ResponseAllOfServersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListHosts200ResponseAllOfServersInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListHosts200ResponseAllOfServersInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListHosts200ResponseAllOfServersInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListHosts200ResponseAllOfServersInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ListHosts200ResponseAllOfServersInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ListHosts200ResponseAllOfServersInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ListHosts200ResponseAllOfServersInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ListHosts200ResponseAllOfServersInner) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetExternalId

`func (o *ListHosts200ResponseAllOfServersInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListHosts200ResponseAllOfServersInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListHosts200ResponseAllOfServersInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListHosts200ResponseAllOfServersInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetInternalId

`func (o *ListHosts200ResponseAllOfServersInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *ListHosts200ResponseAllOfServersInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *ListHosts200ResponseAllOfServersInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *ListHosts200ResponseAllOfServersInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalUniqueId

`func (o *ListHosts200ResponseAllOfServersInner) GetExternalUniqueId() string`

GetExternalUniqueId returns the ExternalUniqueId field if non-nil, zero value otherwise.

### GetExternalUniqueIdOk

`func (o *ListHosts200ResponseAllOfServersInner) GetExternalUniqueIdOk() (*string, bool)`

GetExternalUniqueIdOk returns a tuple with the ExternalUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUniqueId

`func (o *ListHosts200ResponseAllOfServersInner) SetExternalUniqueId(v string)`

SetExternalUniqueId sets ExternalUniqueId field to given value.

### HasExternalUniqueId

`func (o *ListHosts200ResponseAllOfServersInner) HasExternalUniqueId() bool`

HasExternalUniqueId returns a boolean if a field has been set.

### GetName

`func (o *ListHosts200ResponseAllOfServersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListHosts200ResponseAllOfServersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListHosts200ResponseAllOfServersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListHosts200ResponseAllOfServersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalName

`func (o *ListHosts200ResponseAllOfServersInner) GetExternalName() string`

GetExternalName returns the ExternalName field if non-nil, zero value otherwise.

### GetExternalNameOk

`func (o *ListHosts200ResponseAllOfServersInner) GetExternalNameOk() (*string, bool)`

GetExternalNameOk returns a tuple with the ExternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalName

`func (o *ListHosts200ResponseAllOfServersInner) SetExternalName(v string)`

SetExternalName sets ExternalName field to given value.

### HasExternalName

`func (o *ListHosts200ResponseAllOfServersInner) HasExternalName() bool`

HasExternalName returns a boolean if a field has been set.

### GetHostname

`func (o *ListHosts200ResponseAllOfServersInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ListHosts200ResponseAllOfServersInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ListHosts200ResponseAllOfServersInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ListHosts200ResponseAllOfServersInner) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetParentServer

`func (o *ListHosts200ResponseAllOfServersInner) GetParentServer() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetParentServer returns the ParentServer field if non-nil, zero value otherwise.

### GetParentServerOk

`func (o *ListHosts200ResponseAllOfServersInner) GetParentServerOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetParentServerOk returns a tuple with the ParentServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentServer

`func (o *ListHosts200ResponseAllOfServersInner) SetParentServer(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetParentServer sets ParentServer field to given value.

### HasParentServer

`func (o *ListHosts200ResponseAllOfServersInner) HasParentServer() bool`

HasParentServer returns a boolean if a field has been set.

### GetAccountId

`func (o *ListHosts200ResponseAllOfServersInner) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListHosts200ResponseAllOfServersInner) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListHosts200ResponseAllOfServersInner) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListHosts200ResponseAllOfServersInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccount

`func (o *ListHosts200ResponseAllOfServersInner) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListHosts200ResponseAllOfServersInner) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListHosts200ResponseAllOfServersInner) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListHosts200ResponseAllOfServersInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *ListHosts200ResponseAllOfServersInner) GetOwner() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListHosts200ResponseAllOfServersInner) GetOwnerOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListHosts200ResponseAllOfServersInner) SetOwner(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListHosts200ResponseAllOfServersInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetZone

`func (o *ListHosts200ResponseAllOfServersInner) GetZone() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ListHosts200ResponseAllOfServersInner) GetZoneOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ListHosts200ResponseAllOfServersInner) SetZone(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ListHosts200ResponseAllOfServersInner) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetPlan

`func (o *ListHosts200ResponseAllOfServersInner) GetPlan() ListHosts200ResponseAllOfServersInnerPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *ListHosts200ResponseAllOfServersInner) GetPlanOk() (*ListHosts200ResponseAllOfServersInnerPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *ListHosts200ResponseAllOfServersInner) SetPlan(v ListHosts200ResponseAllOfServersInnerPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *ListHosts200ResponseAllOfServersInner) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetComputeServerType

`func (o *ListHosts200ResponseAllOfServersInner) GetComputeServerType() ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType`

GetComputeServerType returns the ComputeServerType field if non-nil, zero value otherwise.

### GetComputeServerTypeOk

`func (o *ListHosts200ResponseAllOfServersInner) GetComputeServerTypeOk() (*ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType, bool)`

GetComputeServerTypeOk returns a tuple with the ComputeServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServerType

`func (o *ListHosts200ResponseAllOfServersInner) SetComputeServerType(v ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType)`

SetComputeServerType sets ComputeServerType field to given value.

### HasComputeServerType

`func (o *ListHosts200ResponseAllOfServersInner) HasComputeServerType() bool`

HasComputeServerType returns a boolean if a field has been set.

### GetVisibility

`func (o *ListHosts200ResponseAllOfServersInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListHosts200ResponseAllOfServersInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListHosts200ResponseAllOfServersInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListHosts200ResponseAllOfServersInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *ListHosts200ResponseAllOfServersInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListHosts200ResponseAllOfServersInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListHosts200ResponseAllOfServersInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListHosts200ResponseAllOfServersInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetZoneId

`func (o *ListHosts200ResponseAllOfServersInner) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ListHosts200ResponseAllOfServersInner) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ListHosts200ResponseAllOfServersInner) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *ListHosts200ResponseAllOfServersInner) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetSiteId

`func (o *ListHosts200ResponseAllOfServersInner) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ListHosts200ResponseAllOfServersInner) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ListHosts200ResponseAllOfServersInner) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ListHosts200ResponseAllOfServersInner) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *ListHosts200ResponseAllOfServersInner) GetResourcePoolId() int64`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *ListHosts200ResponseAllOfServersInner) GetResourcePoolIdOk() (*int64, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *ListHosts200ResponseAllOfServersInner) SetResourcePoolId(v int64)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *ListHosts200ResponseAllOfServersInner) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetFolderId

`func (o *ListHosts200ResponseAllOfServersInner) GetFolderId() int64`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *ListHosts200ResponseAllOfServersInner) GetFolderIdOk() (*int64, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *ListHosts200ResponseAllOfServersInner) SetFolderId(v int64)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *ListHosts200ResponseAllOfServersInner) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetSshHost

`func (o *ListHosts200ResponseAllOfServersInner) GetSshHost() string`

GetSshHost returns the SshHost field if non-nil, zero value otherwise.

### GetSshHostOk

`func (o *ListHosts200ResponseAllOfServersInner) GetSshHostOk() (*string, bool)`

GetSshHostOk returns a tuple with the SshHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHost

`func (o *ListHosts200ResponseAllOfServersInner) SetSshHost(v string)`

SetSshHost sets SshHost field to given value.

### HasSshHost

`func (o *ListHosts200ResponseAllOfServersInner) HasSshHost() bool`

HasSshHost returns a boolean if a field has been set.

### GetSshPort

`func (o *ListHosts200ResponseAllOfServersInner) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *ListHosts200ResponseAllOfServersInner) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *ListHosts200ResponseAllOfServersInner) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *ListHosts200ResponseAllOfServersInner) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetExternalIp

`func (o *ListHosts200ResponseAllOfServersInner) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *ListHosts200ResponseAllOfServersInner) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *ListHosts200ResponseAllOfServersInner) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *ListHosts200ResponseAllOfServersInner) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### GetInternalIp

`func (o *ListHosts200ResponseAllOfServersInner) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *ListHosts200ResponseAllOfServersInner) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *ListHosts200ResponseAllOfServersInner) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *ListHosts200ResponseAllOfServersInner) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### GetVolumeId

`func (o *ListHosts200ResponseAllOfServersInner) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *ListHosts200ResponseAllOfServersInner) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *ListHosts200ResponseAllOfServersInner) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *ListHosts200ResponseAllOfServersInner) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.

### GetPlatform

`func (o *ListHosts200ResponseAllOfServersInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListHosts200ResponseAllOfServersInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListHosts200ResponseAllOfServersInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ListHosts200ResponseAllOfServersInner) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPlatformVersion

`func (o *ListHosts200ResponseAllOfServersInner) GetPlatformVersion() string`

GetPlatformVersion returns the PlatformVersion field if non-nil, zero value otherwise.

### GetPlatformVersionOk

`func (o *ListHosts200ResponseAllOfServersInner) GetPlatformVersionOk() (*string, bool)`

GetPlatformVersionOk returns a tuple with the PlatformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformVersion

`func (o *ListHosts200ResponseAllOfServersInner) SetPlatformVersion(v string)`

SetPlatformVersion sets PlatformVersion field to given value.

### HasPlatformVersion

`func (o *ListHosts200ResponseAllOfServersInner) HasPlatformVersion() bool`

HasPlatformVersion returns a boolean if a field has been set.

### GetSshUsername

`func (o *ListHosts200ResponseAllOfServersInner) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *ListHosts200ResponseAllOfServersInner) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *ListHosts200ResponseAllOfServersInner) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *ListHosts200ResponseAllOfServersInner) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *ListHosts200ResponseAllOfServersInner) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *ListHosts200ResponseAllOfServersInner) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *ListHosts200ResponseAllOfServersInner) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *ListHosts200ResponseAllOfServersInner) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetSshPasswordHash

`func (o *ListHosts200ResponseAllOfServersInner) GetSshPasswordHash() string`

GetSshPasswordHash returns the SshPasswordHash field if non-nil, zero value otherwise.

### GetSshPasswordHashOk

`func (o *ListHosts200ResponseAllOfServersInner) GetSshPasswordHashOk() (*string, bool)`

GetSshPasswordHashOk returns a tuple with the SshPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPasswordHash

`func (o *ListHosts200ResponseAllOfServersInner) SetSshPasswordHash(v string)`

SetSshPasswordHash sets SshPasswordHash field to given value.

### HasSshPasswordHash

`func (o *ListHosts200ResponseAllOfServersInner) HasSshPasswordHash() bool`

HasSshPasswordHash returns a boolean if a field has been set.

### GetSshKeyPair

`func (o *ListHosts200ResponseAllOfServersInner) GetSshKeyPair() GetAlerts200ResponseAllOfChecksInnerContainer`

GetSshKeyPair returns the SshKeyPair field if non-nil, zero value otherwise.

### GetSshKeyPairOk

`func (o *ListHosts200ResponseAllOfServersInner) GetSshKeyPairOk() (*GetAlerts200ResponseAllOfChecksInnerContainer, bool)`

GetSshKeyPairOk returns a tuple with the SshKeyPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyPair

`func (o *ListHosts200ResponseAllOfServersInner) SetSshKeyPair(v GetAlerts200ResponseAllOfChecksInnerContainer)`

SetSshKeyPair sets SshKeyPair field to given value.

### HasSshKeyPair

`func (o *ListHosts200ResponseAllOfServersInner) HasSshKeyPair() bool`

HasSshKeyPair returns a boolean if a field has been set.

### GetOsDevice

`func (o *ListHosts200ResponseAllOfServersInner) GetOsDevice() string`

GetOsDevice returns the OsDevice field if non-nil, zero value otherwise.

### GetOsDeviceOk

`func (o *ListHosts200ResponseAllOfServersInner) GetOsDeviceOk() (*string, bool)`

GetOsDeviceOk returns a tuple with the OsDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDevice

`func (o *ListHosts200ResponseAllOfServersInner) SetOsDevice(v string)`

SetOsDevice sets OsDevice field to given value.

### HasOsDevice

`func (o *ListHosts200ResponseAllOfServersInner) HasOsDevice() bool`

HasOsDevice returns a boolean if a field has been set.

### GetOsType

`func (o *ListHosts200ResponseAllOfServersInner) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *ListHosts200ResponseAllOfServersInner) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *ListHosts200ResponseAllOfServersInner) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *ListHosts200ResponseAllOfServersInner) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetDataDevice

`func (o *ListHosts200ResponseAllOfServersInner) GetDataDevice() string`

GetDataDevice returns the DataDevice field if non-nil, zero value otherwise.

### GetDataDeviceOk

`func (o *ListHosts200ResponseAllOfServersInner) GetDataDeviceOk() (*string, bool)`

GetDataDeviceOk returns a tuple with the DataDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDevice

`func (o *ListHosts200ResponseAllOfServersInner) SetDataDevice(v string)`

SetDataDevice sets DataDevice field to given value.

### HasDataDevice

`func (o *ListHosts200ResponseAllOfServersInner) HasDataDevice() bool`

HasDataDevice returns a boolean if a field has been set.

### GetLvmEnabled

`func (o *ListHosts200ResponseAllOfServersInner) GetLvmEnabled() bool`

GetLvmEnabled returns the LvmEnabled field if non-nil, zero value otherwise.

### GetLvmEnabledOk

`func (o *ListHosts200ResponseAllOfServersInner) GetLvmEnabledOk() (*bool, bool)`

GetLvmEnabledOk returns a tuple with the LvmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLvmEnabled

`func (o *ListHosts200ResponseAllOfServersInner) SetLvmEnabled(v bool)`

SetLvmEnabled sets LvmEnabled field to given value.

### HasLvmEnabled

`func (o *ListHosts200ResponseAllOfServersInner) HasLvmEnabled() bool`

HasLvmEnabled returns a boolean if a field has been set.

### GetApiKey

`func (o *ListHosts200ResponseAllOfServersInner) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ListHosts200ResponseAllOfServersInner) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ListHosts200ResponseAllOfServersInner) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *ListHosts200ResponseAllOfServersInner) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetSoftwareRaid

`func (o *ListHosts200ResponseAllOfServersInner) GetSoftwareRaid() bool`

GetSoftwareRaid returns the SoftwareRaid field if non-nil, zero value otherwise.

### GetSoftwareRaidOk

`func (o *ListHosts200ResponseAllOfServersInner) GetSoftwareRaidOk() (*bool, bool)`

GetSoftwareRaidOk returns a tuple with the SoftwareRaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareRaid

`func (o *ListHosts200ResponseAllOfServersInner) SetSoftwareRaid(v bool)`

SetSoftwareRaid sets SoftwareRaid field to given value.

### HasSoftwareRaid

`func (o *ListHosts200ResponseAllOfServersInner) HasSoftwareRaid() bool`

HasSoftwareRaid returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListHosts200ResponseAllOfServersInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListHosts200ResponseAllOfServersInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListHosts200ResponseAllOfServersInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListHosts200ResponseAllOfServersInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListHosts200ResponseAllOfServersInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListHosts200ResponseAllOfServersInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListHosts200ResponseAllOfServersInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListHosts200ResponseAllOfServersInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetStats

`func (o *ListHosts200ResponseAllOfServersInner) GetStats() ListHosts200ResponseAllOfServersInnerStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ListHosts200ResponseAllOfServersInner) GetStatsOk() (*ListHosts200ResponseAllOfServersInnerStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ListHosts200ResponseAllOfServersInner) SetStats(v ListHosts200ResponseAllOfServersInnerStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ListHosts200ResponseAllOfServersInner) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetStatus

`func (o *ListHosts200ResponseAllOfServersInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListHosts200ResponseAllOfServersInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListHosts200ResponseAllOfServersInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListHosts200ResponseAllOfServersInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ListHosts200ResponseAllOfServersInner) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ListHosts200ResponseAllOfServersInner) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ListHosts200ResponseAllOfServersInner) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ListHosts200ResponseAllOfServersInner) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ListHosts200ResponseAllOfServersInner) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ListHosts200ResponseAllOfServersInner) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ListHosts200ResponseAllOfServersInner) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ListHosts200ResponseAllOfServersInner) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetStatusDate

`func (o *ListHosts200ResponseAllOfServersInner) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *ListHosts200ResponseAllOfServersInner) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *ListHosts200ResponseAllOfServersInner) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *ListHosts200ResponseAllOfServersInner) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetStatusPercent

`func (o *ListHosts200ResponseAllOfServersInner) GetStatusPercent() string`

GetStatusPercent returns the StatusPercent field if non-nil, zero value otherwise.

### GetStatusPercentOk

`func (o *ListHosts200ResponseAllOfServersInner) GetStatusPercentOk() (*string, bool)`

GetStatusPercentOk returns a tuple with the StatusPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusPercent

`func (o *ListHosts200ResponseAllOfServersInner) SetStatusPercent(v string)`

SetStatusPercent sets StatusPercent field to given value.

### HasStatusPercent

`func (o *ListHosts200ResponseAllOfServersInner) HasStatusPercent() bool`

HasStatusPercent returns a boolean if a field has been set.

### GetStatusEta

`func (o *ListHosts200ResponseAllOfServersInner) GetStatusEta() string`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *ListHosts200ResponseAllOfServersInner) GetStatusEtaOk() (*string, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *ListHosts200ResponseAllOfServersInner) SetStatusEta(v string)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *ListHosts200ResponseAllOfServersInner) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### GetPowerState

`func (o *ListHosts200ResponseAllOfServersInner) GetPowerState() string`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *ListHosts200ResponseAllOfServersInner) GetPowerStateOk() (*string, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *ListHosts200ResponseAllOfServersInner) SetPowerState(v string)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *ListHosts200ResponseAllOfServersInner) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetAgentInstalled

`func (o *ListHosts200ResponseAllOfServersInner) GetAgentInstalled() bool`

GetAgentInstalled returns the AgentInstalled field if non-nil, zero value otherwise.

### GetAgentInstalledOk

`func (o *ListHosts200ResponseAllOfServersInner) GetAgentInstalledOk() (*bool, bool)`

GetAgentInstalledOk returns a tuple with the AgentInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentInstalled

`func (o *ListHosts200ResponseAllOfServersInner) SetAgentInstalled(v bool)`

SetAgentInstalled sets AgentInstalled field to given value.

### HasAgentInstalled

`func (o *ListHosts200ResponseAllOfServersInner) HasAgentInstalled() bool`

HasAgentInstalled returns a boolean if a field has been set.

### GetLastAgentUpdate

`func (o *ListHosts200ResponseAllOfServersInner) GetLastAgentUpdate() string`

GetLastAgentUpdate returns the LastAgentUpdate field if non-nil, zero value otherwise.

### GetLastAgentUpdateOk

`func (o *ListHosts200ResponseAllOfServersInner) GetLastAgentUpdateOk() (*string, bool)`

GetLastAgentUpdateOk returns a tuple with the LastAgentUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAgentUpdate

`func (o *ListHosts200ResponseAllOfServersInner) SetLastAgentUpdate(v string)`

SetLastAgentUpdate sets LastAgentUpdate field to given value.

### HasLastAgentUpdate

`func (o *ListHosts200ResponseAllOfServersInner) HasLastAgentUpdate() bool`

HasLastAgentUpdate returns a boolean if a field has been set.

### GetAgentVersion

`func (o *ListHosts200ResponseAllOfServersInner) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *ListHosts200ResponseAllOfServersInner) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *ListHosts200ResponseAllOfServersInner) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *ListHosts200ResponseAllOfServersInner) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### GetMaxCores

`func (o *ListHosts200ResponseAllOfServersInner) GetMaxCores() int64`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *ListHosts200ResponseAllOfServersInner) GetMaxCoresOk() (*int64, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *ListHosts200ResponseAllOfServersInner) SetMaxCores(v int64)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *ListHosts200ResponseAllOfServersInner) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *ListHosts200ResponseAllOfServersInner) GetCoresPerSocket() int64`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *ListHosts200ResponseAllOfServersInner) GetCoresPerSocketOk() (*int64, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *ListHosts200ResponseAllOfServersInner) SetCoresPerSocket(v int64)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *ListHosts200ResponseAllOfServersInner) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### GetMaxMemory

`func (o *ListHosts200ResponseAllOfServersInner) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *ListHosts200ResponseAllOfServersInner) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *ListHosts200ResponseAllOfServersInner) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *ListHosts200ResponseAllOfServersInner) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxStorage

`func (o *ListHosts200ResponseAllOfServersInner) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ListHosts200ResponseAllOfServersInner) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ListHosts200ResponseAllOfServersInner) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *ListHosts200ResponseAllOfServersInner) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMaxCpu

`func (o *ListHosts200ResponseAllOfServersInner) GetMaxCpu() int64`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *ListHosts200ResponseAllOfServersInner) GetMaxCpuOk() (*int64, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *ListHosts200ResponseAllOfServersInner) SetMaxCpu(v int64)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *ListHosts200ResponseAllOfServersInner) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetMaxGpus

`func (o *ListHosts200ResponseAllOfServersInner) GetMaxGpus() int64`

GetMaxGpus returns the MaxGpus field if non-nil, zero value otherwise.

### GetMaxGpusOk

`func (o *ListHosts200ResponseAllOfServersInner) GetMaxGpusOk() (*int64, bool)`

GetMaxGpusOk returns a tuple with the MaxGpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGpus

`func (o *ListHosts200ResponseAllOfServersInner) SetMaxGpus(v int64)`

SetMaxGpus sets MaxGpus field to given value.

### HasMaxGpus

`func (o *ListHosts200ResponseAllOfServersInner) HasMaxGpus() bool`

HasMaxGpus returns a boolean if a field has been set.

### GetManageInternalFirewall

`func (o *ListHosts200ResponseAllOfServersInner) GetManageInternalFirewall() bool`

GetManageInternalFirewall returns the ManageInternalFirewall field if non-nil, zero value otherwise.

### GetManageInternalFirewallOk

`func (o *ListHosts200ResponseAllOfServersInner) GetManageInternalFirewallOk() (*bool, bool)`

GetManageInternalFirewallOk returns a tuple with the ManageInternalFirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageInternalFirewall

`func (o *ListHosts200ResponseAllOfServersInner) SetManageInternalFirewall(v bool)`

SetManageInternalFirewall sets ManageInternalFirewall field to given value.

### HasManageInternalFirewall

`func (o *ListHosts200ResponseAllOfServersInner) HasManageInternalFirewall() bool`

HasManageInternalFirewall returns a boolean if a field has been set.

### GetEnableLogs

`func (o *ListHosts200ResponseAllOfServersInner) GetEnableLogs() bool`

GetEnableLogs returns the EnableLogs field if non-nil, zero value otherwise.

### GetEnableLogsOk

`func (o *ListHosts200ResponseAllOfServersInner) GetEnableLogsOk() (*bool, bool)`

GetEnableLogsOk returns a tuple with the EnableLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLogs

`func (o *ListHosts200ResponseAllOfServersInner) SetEnableLogs(v bool)`

SetEnableLogs sets EnableLogs field to given value.

### HasEnableLogs

`func (o *ListHosts200ResponseAllOfServersInner) HasEnableLogs() bool`

HasEnableLogs returns a boolean if a field has been set.

### GetHourlyCost

`func (o *ListHosts200ResponseAllOfServersInner) GetHourlyCost() float32`

GetHourlyCost returns the HourlyCost field if non-nil, zero value otherwise.

### GetHourlyCostOk

`func (o *ListHosts200ResponseAllOfServersInner) GetHourlyCostOk() (*float32, bool)`

GetHourlyCostOk returns a tuple with the HourlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyCost

`func (o *ListHosts200ResponseAllOfServersInner) SetHourlyCost(v float32)`

SetHourlyCost sets HourlyCost field to given value.

### HasHourlyCost

`func (o *ListHosts200ResponseAllOfServersInner) HasHourlyCost() bool`

HasHourlyCost returns a boolean if a field has been set.

### GetHourlyPrice

`func (o *ListHosts200ResponseAllOfServersInner) GetHourlyPrice() float32`

GetHourlyPrice returns the HourlyPrice field if non-nil, zero value otherwise.

### GetHourlyPriceOk

`func (o *ListHosts200ResponseAllOfServersInner) GetHourlyPriceOk() (*float32, bool)`

GetHourlyPriceOk returns a tuple with the HourlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyPrice

`func (o *ListHosts200ResponseAllOfServersInner) SetHourlyPrice(v float32)`

SetHourlyPrice sets HourlyPrice field to given value.

### HasHourlyPrice

`func (o *ListHosts200ResponseAllOfServersInner) HasHourlyPrice() bool`

HasHourlyPrice returns a boolean if a field has been set.

### GetSourceImage

`func (o *ListHosts200ResponseAllOfServersInner) GetSourceImage() ListPriceSets200ResponseAllOfPriceSetsInnerPricesInnerVolumeType`

GetSourceImage returns the SourceImage field if non-nil, zero value otherwise.

### GetSourceImageOk

`func (o *ListHosts200ResponseAllOfServersInner) GetSourceImageOk() (*ListPriceSets200ResponseAllOfPriceSetsInnerPricesInnerVolumeType, bool)`

GetSourceImageOk returns a tuple with the SourceImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceImage

`func (o *ListHosts200ResponseAllOfServersInner) SetSourceImage(v ListPriceSets200ResponseAllOfPriceSetsInnerPricesInnerVolumeType)`

SetSourceImage sets SourceImage field to given value.

### HasSourceImage

`func (o *ListHosts200ResponseAllOfServersInner) HasSourceImage() bool`

HasSourceImage returns a boolean if a field has been set.

### GetServerOs

`func (o *ListHosts200ResponseAllOfServersInner) GetServerOs() ListHosts200ResponseAllOfServersInnerServerOs`

GetServerOs returns the ServerOs field if non-nil, zero value otherwise.

### GetServerOsOk

`func (o *ListHosts200ResponseAllOfServersInner) GetServerOsOk() (*ListHosts200ResponseAllOfServersInnerServerOs, bool)`

GetServerOsOk returns a tuple with the ServerOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerOs

`func (o *ListHosts200ResponseAllOfServersInner) SetServerOs(v ListHosts200ResponseAllOfServersInnerServerOs)`

SetServerOs sets ServerOs field to given value.

### HasServerOs

`func (o *ListHosts200ResponseAllOfServersInner) HasServerOs() bool`

HasServerOs returns a boolean if a field has been set.

### GetVolumes

`func (o *ListHosts200ResponseAllOfServersInner) GetVolumes() []ListHosts200ResponseAllOfServersInnerVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ListHosts200ResponseAllOfServersInner) GetVolumesOk() (*[]ListHosts200ResponseAllOfServersInnerVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ListHosts200ResponseAllOfServersInner) SetVolumes(v []ListHosts200ResponseAllOfServersInnerVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *ListHosts200ResponseAllOfServersInner) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetControllers

`func (o *ListHosts200ResponseAllOfServersInner) GetControllers() []ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceControllersInner`

GetControllers returns the Controllers field if non-nil, zero value otherwise.

### GetControllersOk

`func (o *ListHosts200ResponseAllOfServersInner) GetControllersOk() (*[]ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceControllersInner, bool)`

GetControllersOk returns a tuple with the Controllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllers

`func (o *ListHosts200ResponseAllOfServersInner) SetControllers(v []ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceControllersInner)`

SetControllers sets Controllers field to given value.

### HasControllers

`func (o *ListHosts200ResponseAllOfServersInner) HasControllers() bool`

HasControllers returns a boolean if a field has been set.

### GetInterfaces

`func (o *ListHosts200ResponseAllOfServersInner) GetInterfaces() []ListHosts200ResponseAllOfServersInnerInterfacesInner`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *ListHosts200ResponseAllOfServersInner) GetInterfacesOk() (*[]ListHosts200ResponseAllOfServersInnerInterfacesInner, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *ListHosts200ResponseAllOfServersInner) SetInterfaces(v []ListHosts200ResponseAllOfServersInnerInterfacesInner)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *ListHosts200ResponseAllOfServersInner) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetLabels

`func (o *ListHosts200ResponseAllOfServersInner) GetLabels() []map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListHosts200ResponseAllOfServersInner) GetLabelsOk() (*[]map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListHosts200ResponseAllOfServersInner) SetLabels(v []map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListHosts200ResponseAllOfServersInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTags

`func (o *ListHosts200ResponseAllOfServersInner) GetTags() []map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ListHosts200ResponseAllOfServersInner) GetTagsOk() (*[]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ListHosts200ResponseAllOfServersInner) SetTags(v []map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *ListHosts200ResponseAllOfServersInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEnabled

`func (o *ListHosts200ResponseAllOfServersInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListHosts200ResponseAllOfServersInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListHosts200ResponseAllOfServersInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListHosts200ResponseAllOfServersInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTagCompliant

`func (o *ListHosts200ResponseAllOfServersInner) GetTagCompliant() string`

GetTagCompliant returns the TagCompliant field if non-nil, zero value otherwise.

### GetTagCompliantOk

`func (o *ListHosts200ResponseAllOfServersInner) GetTagCompliantOk() (*string, bool)`

GetTagCompliantOk returns a tuple with the TagCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagCompliant

`func (o *ListHosts200ResponseAllOfServersInner) SetTagCompliant(v string)`

SetTagCompliant sets TagCompliant field to given value.

### HasTagCompliant

`func (o *ListHosts200ResponseAllOfServersInner) HasTagCompliant() bool`

HasTagCompliant returns a boolean if a field has been set.

### GetContainers

`func (o *ListHosts200ResponseAllOfServersInner) GetContainers() []int64`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *ListHosts200ResponseAllOfServersInner) GetContainersOk() (*[]int64, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *ListHosts200ResponseAllOfServersInner) SetContainers(v []int64)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *ListHosts200ResponseAllOfServersInner) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetConfig

`func (o *ListHosts200ResponseAllOfServersInner) GetConfig() ListHosts200ResponseAllOfServersInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListHosts200ResponseAllOfServersInner) GetConfigOk() (*ListHosts200ResponseAllOfServersInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListHosts200ResponseAllOfServersInner) SetConfig(v ListHosts200ResponseAllOfServersInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListHosts200ResponseAllOfServersInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetGuestConsolePreferred

`func (o *ListHosts200ResponseAllOfServersInner) GetGuestConsolePreferred() bool`

GetGuestConsolePreferred returns the GuestConsolePreferred field if non-nil, zero value otherwise.

### GetGuestConsolePreferredOk

`func (o *ListHosts200ResponseAllOfServersInner) GetGuestConsolePreferredOk() (*bool, bool)`

GetGuestConsolePreferredOk returns a tuple with the GuestConsolePreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsolePreferred

`func (o *ListHosts200ResponseAllOfServersInner) SetGuestConsolePreferred(v bool)`

SetGuestConsolePreferred sets GuestConsolePreferred field to given value.

### HasGuestConsolePreferred

`func (o *ListHosts200ResponseAllOfServersInner) HasGuestConsolePreferred() bool`

HasGuestConsolePreferred returns a boolean if a field has been set.

### GetGuestConsoleType

`func (o *ListHosts200ResponseAllOfServersInner) GetGuestConsoleType() string`

GetGuestConsoleType returns the GuestConsoleType field if non-nil, zero value otherwise.

### GetGuestConsoleTypeOk

`func (o *ListHosts200ResponseAllOfServersInner) GetGuestConsoleTypeOk() (*string, bool)`

GetGuestConsoleTypeOk returns a tuple with the GuestConsoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleType

`func (o *ListHosts200ResponseAllOfServersInner) SetGuestConsoleType(v string)`

SetGuestConsoleType sets GuestConsoleType field to given value.

### HasGuestConsoleType

`func (o *ListHosts200ResponseAllOfServersInner) HasGuestConsoleType() bool`

HasGuestConsoleType returns a boolean if a field has been set.

### GetGuestConsoleUsername

`func (o *ListHosts200ResponseAllOfServersInner) GetGuestConsoleUsername() string`

GetGuestConsoleUsername returns the GuestConsoleUsername field if non-nil, zero value otherwise.

### GetGuestConsoleUsernameOk

`func (o *ListHosts200ResponseAllOfServersInner) GetGuestConsoleUsernameOk() (*string, bool)`

GetGuestConsoleUsernameOk returns a tuple with the GuestConsoleUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleUsername

`func (o *ListHosts200ResponseAllOfServersInner) SetGuestConsoleUsername(v string)`

SetGuestConsoleUsername sets GuestConsoleUsername field to given value.

### HasGuestConsoleUsername

`func (o *ListHosts200ResponseAllOfServersInner) HasGuestConsoleUsername() bool`

HasGuestConsoleUsername returns a boolean if a field has been set.

### GetGuestConsolePassword

`func (o *ListHosts200ResponseAllOfServersInner) GetGuestConsolePassword() string`

GetGuestConsolePassword returns the GuestConsolePassword field if non-nil, zero value otherwise.

### GetGuestConsolePasswordOk

`func (o *ListHosts200ResponseAllOfServersInner) GetGuestConsolePasswordOk() (*string, bool)`

GetGuestConsolePasswordOk returns a tuple with the GuestConsolePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsolePassword

`func (o *ListHosts200ResponseAllOfServersInner) SetGuestConsolePassword(v string)`

SetGuestConsolePassword sets GuestConsolePassword field to given value.

### HasGuestConsolePassword

`func (o *ListHosts200ResponseAllOfServersInner) HasGuestConsolePassword() bool`

HasGuestConsolePassword returns a boolean if a field has been set.

### GetGuestConsolePasswordHash

`func (o *ListHosts200ResponseAllOfServersInner) GetGuestConsolePasswordHash() string`

GetGuestConsolePasswordHash returns the GuestConsolePasswordHash field if non-nil, zero value otherwise.

### GetGuestConsolePasswordHashOk

`func (o *ListHosts200ResponseAllOfServersInner) GetGuestConsolePasswordHashOk() (*string, bool)`

GetGuestConsolePasswordHashOk returns a tuple with the GuestConsolePasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsolePasswordHash

`func (o *ListHosts200ResponseAllOfServersInner) SetGuestConsolePasswordHash(v string)`

SetGuestConsolePasswordHash sets GuestConsolePasswordHash field to given value.

### HasGuestConsolePasswordHash

`func (o *ListHosts200ResponseAllOfServersInner) HasGuestConsolePasswordHash() bool`

HasGuestConsolePasswordHash returns a boolean if a field has been set.

### GetGuestConsolePort

`func (o *ListHosts200ResponseAllOfServersInner) GetGuestConsolePort() string`

GetGuestConsolePort returns the GuestConsolePort field if non-nil, zero value otherwise.

### GetGuestConsolePortOk

`func (o *ListHosts200ResponseAllOfServersInner) GetGuestConsolePortOk() (*string, bool)`

GetGuestConsolePortOk returns a tuple with the GuestConsolePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsolePort

`func (o *ListHosts200ResponseAllOfServersInner) SetGuestConsolePort(v string)`

SetGuestConsolePort sets GuestConsolePort field to given value.

### HasGuestConsolePort

`func (o *ListHosts200ResponseAllOfServersInner) HasGuestConsolePort() bool`

HasGuestConsolePort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


