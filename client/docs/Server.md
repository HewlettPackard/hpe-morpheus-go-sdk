# Server

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

### NewServer

`func NewServer() *Server`

NewServer instantiates a new Server object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerWithDefaults

`func NewServerWithDefaults() *Server`

NewServerWithDefaults instantiates a new Server object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Server) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Server) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Server) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Server) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *Server) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Server) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Server) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Server) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetExternalId

`func (o *Server) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Server) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Server) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Server) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetInternalId

`func (o *Server) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *Server) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *Server) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *Server) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalUniqueId

`func (o *Server) GetExternalUniqueId() string`

GetExternalUniqueId returns the ExternalUniqueId field if non-nil, zero value otherwise.

### GetExternalUniqueIdOk

`func (o *Server) GetExternalUniqueIdOk() (*string, bool)`

GetExternalUniqueIdOk returns a tuple with the ExternalUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUniqueId

`func (o *Server) SetExternalUniqueId(v string)`

SetExternalUniqueId sets ExternalUniqueId field to given value.

### HasExternalUniqueId

`func (o *Server) HasExternalUniqueId() bool`

HasExternalUniqueId returns a boolean if a field has been set.

### GetName

`func (o *Server) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Server) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Server) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Server) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalName

`func (o *Server) GetExternalName() string`

GetExternalName returns the ExternalName field if non-nil, zero value otherwise.

### GetExternalNameOk

`func (o *Server) GetExternalNameOk() (*string, bool)`

GetExternalNameOk returns a tuple with the ExternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalName

`func (o *Server) SetExternalName(v string)`

SetExternalName sets ExternalName field to given value.

### HasExternalName

`func (o *Server) HasExternalName() bool`

HasExternalName returns a boolean if a field has been set.

### GetHostname

`func (o *Server) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *Server) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *Server) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *Server) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetParentServer

`func (o *Server) GetParentServer() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetParentServer returns the ParentServer field if non-nil, zero value otherwise.

### GetParentServerOk

`func (o *Server) GetParentServerOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetParentServerOk returns a tuple with the ParentServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentServer

`func (o *Server) SetParentServer(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetParentServer sets ParentServer field to given value.

### HasParentServer

`func (o *Server) HasParentServer() bool`

HasParentServer returns a boolean if a field has been set.

### GetAccountId

`func (o *Server) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Server) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Server) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Server) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccount

`func (o *Server) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Server) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Server) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Server) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *Server) GetOwner() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Server) GetOwnerOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Server) SetOwner(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Server) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetZone

`func (o *Server) GetZone() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *Server) GetZoneOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *Server) SetZone(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetZone sets Zone field to given value.

### HasZone

`func (o *Server) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetPlan

`func (o *Server) GetPlan() ListHosts200ResponseAllOfServersInnerPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *Server) GetPlanOk() (*ListHosts200ResponseAllOfServersInnerPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *Server) SetPlan(v ListHosts200ResponseAllOfServersInnerPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *Server) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetComputeServerType

`func (o *Server) GetComputeServerType() ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType`

GetComputeServerType returns the ComputeServerType field if non-nil, zero value otherwise.

### GetComputeServerTypeOk

`func (o *Server) GetComputeServerTypeOk() (*ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType, bool)`

GetComputeServerTypeOk returns a tuple with the ComputeServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServerType

`func (o *Server) SetComputeServerType(v ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType)`

SetComputeServerType sets ComputeServerType field to given value.

### HasComputeServerType

`func (o *Server) HasComputeServerType() bool`

HasComputeServerType returns a boolean if a field has been set.

### GetVisibility

`func (o *Server) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Server) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Server) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *Server) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *Server) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Server) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Server) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Server) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetZoneId

`func (o *Server) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *Server) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *Server) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *Server) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetSiteId

`func (o *Server) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Server) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Server) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *Server) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *Server) GetResourcePoolId() int64`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *Server) GetResourcePoolIdOk() (*int64, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *Server) SetResourcePoolId(v int64)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *Server) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetFolderId

`func (o *Server) GetFolderId() int64`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *Server) GetFolderIdOk() (*int64, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *Server) SetFolderId(v int64)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *Server) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetSshHost

`func (o *Server) GetSshHost() string`

GetSshHost returns the SshHost field if non-nil, zero value otherwise.

### GetSshHostOk

`func (o *Server) GetSshHostOk() (*string, bool)`

GetSshHostOk returns a tuple with the SshHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHost

`func (o *Server) SetSshHost(v string)`

SetSshHost sets SshHost field to given value.

### HasSshHost

`func (o *Server) HasSshHost() bool`

HasSshHost returns a boolean if a field has been set.

### GetSshPort

`func (o *Server) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *Server) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *Server) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *Server) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetExternalIp

`func (o *Server) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *Server) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *Server) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *Server) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### GetInternalIp

`func (o *Server) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *Server) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *Server) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *Server) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### GetVolumeId

`func (o *Server) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *Server) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *Server) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *Server) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.

### GetPlatform

`func (o *Server) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *Server) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *Server) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *Server) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPlatformVersion

`func (o *Server) GetPlatformVersion() string`

GetPlatformVersion returns the PlatformVersion field if non-nil, zero value otherwise.

### GetPlatformVersionOk

`func (o *Server) GetPlatformVersionOk() (*string, bool)`

GetPlatformVersionOk returns a tuple with the PlatformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformVersion

`func (o *Server) SetPlatformVersion(v string)`

SetPlatformVersion sets PlatformVersion field to given value.

### HasPlatformVersion

`func (o *Server) HasPlatformVersion() bool`

HasPlatformVersion returns a boolean if a field has been set.

### GetSshUsername

`func (o *Server) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *Server) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *Server) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *Server) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *Server) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *Server) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *Server) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *Server) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetSshPasswordHash

`func (o *Server) GetSshPasswordHash() string`

GetSshPasswordHash returns the SshPasswordHash field if non-nil, zero value otherwise.

### GetSshPasswordHashOk

`func (o *Server) GetSshPasswordHashOk() (*string, bool)`

GetSshPasswordHashOk returns a tuple with the SshPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPasswordHash

`func (o *Server) SetSshPasswordHash(v string)`

SetSshPasswordHash sets SshPasswordHash field to given value.

### HasSshPasswordHash

`func (o *Server) HasSshPasswordHash() bool`

HasSshPasswordHash returns a boolean if a field has been set.

### GetSshKeyPair

`func (o *Server) GetSshKeyPair() GetAlerts200ResponseAllOfChecksInnerContainer`

GetSshKeyPair returns the SshKeyPair field if non-nil, zero value otherwise.

### GetSshKeyPairOk

`func (o *Server) GetSshKeyPairOk() (*GetAlerts200ResponseAllOfChecksInnerContainer, bool)`

GetSshKeyPairOk returns a tuple with the SshKeyPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyPair

`func (o *Server) SetSshKeyPair(v GetAlerts200ResponseAllOfChecksInnerContainer)`

SetSshKeyPair sets SshKeyPair field to given value.

### HasSshKeyPair

`func (o *Server) HasSshKeyPair() bool`

HasSshKeyPair returns a boolean if a field has been set.

### GetOsDevice

`func (o *Server) GetOsDevice() string`

GetOsDevice returns the OsDevice field if non-nil, zero value otherwise.

### GetOsDeviceOk

`func (o *Server) GetOsDeviceOk() (*string, bool)`

GetOsDeviceOk returns a tuple with the OsDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDevice

`func (o *Server) SetOsDevice(v string)`

SetOsDevice sets OsDevice field to given value.

### HasOsDevice

`func (o *Server) HasOsDevice() bool`

HasOsDevice returns a boolean if a field has been set.

### GetOsType

`func (o *Server) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *Server) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *Server) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *Server) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetDataDevice

`func (o *Server) GetDataDevice() string`

GetDataDevice returns the DataDevice field if non-nil, zero value otherwise.

### GetDataDeviceOk

`func (o *Server) GetDataDeviceOk() (*string, bool)`

GetDataDeviceOk returns a tuple with the DataDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDevice

`func (o *Server) SetDataDevice(v string)`

SetDataDevice sets DataDevice field to given value.

### HasDataDevice

`func (o *Server) HasDataDevice() bool`

HasDataDevice returns a boolean if a field has been set.

### GetLvmEnabled

`func (o *Server) GetLvmEnabled() bool`

GetLvmEnabled returns the LvmEnabled field if non-nil, zero value otherwise.

### GetLvmEnabledOk

`func (o *Server) GetLvmEnabledOk() (*bool, bool)`

GetLvmEnabledOk returns a tuple with the LvmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLvmEnabled

`func (o *Server) SetLvmEnabled(v bool)`

SetLvmEnabled sets LvmEnabled field to given value.

### HasLvmEnabled

`func (o *Server) HasLvmEnabled() bool`

HasLvmEnabled returns a boolean if a field has been set.

### GetApiKey

`func (o *Server) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *Server) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *Server) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *Server) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetSoftwareRaid

`func (o *Server) GetSoftwareRaid() bool`

GetSoftwareRaid returns the SoftwareRaid field if non-nil, zero value otherwise.

### GetSoftwareRaidOk

`func (o *Server) GetSoftwareRaidOk() (*bool, bool)`

GetSoftwareRaidOk returns a tuple with the SoftwareRaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareRaid

`func (o *Server) SetSoftwareRaid(v bool)`

SetSoftwareRaid sets SoftwareRaid field to given value.

### HasSoftwareRaid

`func (o *Server) HasSoftwareRaid() bool`

HasSoftwareRaid returns a boolean if a field has been set.

### GetDateCreated

`func (o *Server) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Server) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Server) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Server) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Server) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Server) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Server) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Server) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetStats

`func (o *Server) GetStats() ListHosts200ResponseAllOfServersInnerStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Server) GetStatsOk() (*ListHosts200ResponseAllOfServersInnerStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Server) SetStats(v ListHosts200ResponseAllOfServersInnerStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *Server) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetStatus

`func (o *Server) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Server) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Server) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Server) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *Server) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *Server) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *Server) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *Server) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetErrorMessage

`func (o *Server) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *Server) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *Server) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *Server) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetStatusDate

`func (o *Server) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *Server) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *Server) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *Server) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetStatusPercent

`func (o *Server) GetStatusPercent() string`

GetStatusPercent returns the StatusPercent field if non-nil, zero value otherwise.

### GetStatusPercentOk

`func (o *Server) GetStatusPercentOk() (*string, bool)`

GetStatusPercentOk returns a tuple with the StatusPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusPercent

`func (o *Server) SetStatusPercent(v string)`

SetStatusPercent sets StatusPercent field to given value.

### HasStatusPercent

`func (o *Server) HasStatusPercent() bool`

HasStatusPercent returns a boolean if a field has been set.

### GetStatusEta

`func (o *Server) GetStatusEta() string`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *Server) GetStatusEtaOk() (*string, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *Server) SetStatusEta(v string)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *Server) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### GetPowerState

`func (o *Server) GetPowerState() string`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *Server) GetPowerStateOk() (*string, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *Server) SetPowerState(v string)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *Server) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetAgentInstalled

`func (o *Server) GetAgentInstalled() bool`

GetAgentInstalled returns the AgentInstalled field if non-nil, zero value otherwise.

### GetAgentInstalledOk

`func (o *Server) GetAgentInstalledOk() (*bool, bool)`

GetAgentInstalledOk returns a tuple with the AgentInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentInstalled

`func (o *Server) SetAgentInstalled(v bool)`

SetAgentInstalled sets AgentInstalled field to given value.

### HasAgentInstalled

`func (o *Server) HasAgentInstalled() bool`

HasAgentInstalled returns a boolean if a field has been set.

### GetLastAgentUpdate

`func (o *Server) GetLastAgentUpdate() string`

GetLastAgentUpdate returns the LastAgentUpdate field if non-nil, zero value otherwise.

### GetLastAgentUpdateOk

`func (o *Server) GetLastAgentUpdateOk() (*string, bool)`

GetLastAgentUpdateOk returns a tuple with the LastAgentUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAgentUpdate

`func (o *Server) SetLastAgentUpdate(v string)`

SetLastAgentUpdate sets LastAgentUpdate field to given value.

### HasLastAgentUpdate

`func (o *Server) HasLastAgentUpdate() bool`

HasLastAgentUpdate returns a boolean if a field has been set.

### GetAgentVersion

`func (o *Server) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *Server) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *Server) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *Server) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### GetMaxCores

`func (o *Server) GetMaxCores() int64`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *Server) GetMaxCoresOk() (*int64, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *Server) SetMaxCores(v int64)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *Server) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *Server) GetCoresPerSocket() int64`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *Server) GetCoresPerSocketOk() (*int64, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *Server) SetCoresPerSocket(v int64)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *Server) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### GetMaxMemory

`func (o *Server) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *Server) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *Server) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *Server) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxStorage

`func (o *Server) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *Server) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *Server) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *Server) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMaxCpu

`func (o *Server) GetMaxCpu() int64`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *Server) GetMaxCpuOk() (*int64, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *Server) SetMaxCpu(v int64)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *Server) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetMaxGpus

`func (o *Server) GetMaxGpus() int64`

GetMaxGpus returns the MaxGpus field if non-nil, zero value otherwise.

### GetMaxGpusOk

`func (o *Server) GetMaxGpusOk() (*int64, bool)`

GetMaxGpusOk returns a tuple with the MaxGpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGpus

`func (o *Server) SetMaxGpus(v int64)`

SetMaxGpus sets MaxGpus field to given value.

### HasMaxGpus

`func (o *Server) HasMaxGpus() bool`

HasMaxGpus returns a boolean if a field has been set.

### GetManageInternalFirewall

`func (o *Server) GetManageInternalFirewall() bool`

GetManageInternalFirewall returns the ManageInternalFirewall field if non-nil, zero value otherwise.

### GetManageInternalFirewallOk

`func (o *Server) GetManageInternalFirewallOk() (*bool, bool)`

GetManageInternalFirewallOk returns a tuple with the ManageInternalFirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageInternalFirewall

`func (o *Server) SetManageInternalFirewall(v bool)`

SetManageInternalFirewall sets ManageInternalFirewall field to given value.

### HasManageInternalFirewall

`func (o *Server) HasManageInternalFirewall() bool`

HasManageInternalFirewall returns a boolean if a field has been set.

### GetEnableLogs

`func (o *Server) GetEnableLogs() bool`

GetEnableLogs returns the EnableLogs field if non-nil, zero value otherwise.

### GetEnableLogsOk

`func (o *Server) GetEnableLogsOk() (*bool, bool)`

GetEnableLogsOk returns a tuple with the EnableLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLogs

`func (o *Server) SetEnableLogs(v bool)`

SetEnableLogs sets EnableLogs field to given value.

### HasEnableLogs

`func (o *Server) HasEnableLogs() bool`

HasEnableLogs returns a boolean if a field has been set.

### GetHourlyCost

`func (o *Server) GetHourlyCost() float32`

GetHourlyCost returns the HourlyCost field if non-nil, zero value otherwise.

### GetHourlyCostOk

`func (o *Server) GetHourlyCostOk() (*float32, bool)`

GetHourlyCostOk returns a tuple with the HourlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyCost

`func (o *Server) SetHourlyCost(v float32)`

SetHourlyCost sets HourlyCost field to given value.

### HasHourlyCost

`func (o *Server) HasHourlyCost() bool`

HasHourlyCost returns a boolean if a field has been set.

### GetHourlyPrice

`func (o *Server) GetHourlyPrice() float32`

GetHourlyPrice returns the HourlyPrice field if non-nil, zero value otherwise.

### GetHourlyPriceOk

`func (o *Server) GetHourlyPriceOk() (*float32, bool)`

GetHourlyPriceOk returns a tuple with the HourlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyPrice

`func (o *Server) SetHourlyPrice(v float32)`

SetHourlyPrice sets HourlyPrice field to given value.

### HasHourlyPrice

`func (o *Server) HasHourlyPrice() bool`

HasHourlyPrice returns a boolean if a field has been set.

### GetSourceImage

`func (o *Server) GetSourceImage() ListPriceSets200ResponseAllOfPriceSetsInnerPricesInnerVolumeType`

GetSourceImage returns the SourceImage field if non-nil, zero value otherwise.

### GetSourceImageOk

`func (o *Server) GetSourceImageOk() (*ListPriceSets200ResponseAllOfPriceSetsInnerPricesInnerVolumeType, bool)`

GetSourceImageOk returns a tuple with the SourceImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceImage

`func (o *Server) SetSourceImage(v ListPriceSets200ResponseAllOfPriceSetsInnerPricesInnerVolumeType)`

SetSourceImage sets SourceImage field to given value.

### HasSourceImage

`func (o *Server) HasSourceImage() bool`

HasSourceImage returns a boolean if a field has been set.

### GetServerOs

`func (o *Server) GetServerOs() ListHosts200ResponseAllOfServersInnerServerOs`

GetServerOs returns the ServerOs field if non-nil, zero value otherwise.

### GetServerOsOk

`func (o *Server) GetServerOsOk() (*ListHosts200ResponseAllOfServersInnerServerOs, bool)`

GetServerOsOk returns a tuple with the ServerOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerOs

`func (o *Server) SetServerOs(v ListHosts200ResponseAllOfServersInnerServerOs)`

SetServerOs sets ServerOs field to given value.

### HasServerOs

`func (o *Server) HasServerOs() bool`

HasServerOs returns a boolean if a field has been set.

### GetVolumes

`func (o *Server) GetVolumes() []ListHosts200ResponseAllOfServersInnerVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *Server) GetVolumesOk() (*[]ListHosts200ResponseAllOfServersInnerVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *Server) SetVolumes(v []ListHosts200ResponseAllOfServersInnerVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *Server) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetControllers

`func (o *Server) GetControllers() []ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceControllersInner`

GetControllers returns the Controllers field if non-nil, zero value otherwise.

### GetControllersOk

`func (o *Server) GetControllersOk() (*[]ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceControllersInner, bool)`

GetControllersOk returns a tuple with the Controllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllers

`func (o *Server) SetControllers(v []ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceControllersInner)`

SetControllers sets Controllers field to given value.

### HasControllers

`func (o *Server) HasControllers() bool`

HasControllers returns a boolean if a field has been set.

### GetInterfaces

`func (o *Server) GetInterfaces() []ListHosts200ResponseAllOfServersInnerInterfacesInner`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *Server) GetInterfacesOk() (*[]ListHosts200ResponseAllOfServersInnerInterfacesInner, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *Server) SetInterfaces(v []ListHosts200ResponseAllOfServersInnerInterfacesInner)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *Server) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetLabels

`func (o *Server) GetLabels() []map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Server) GetLabelsOk() (*[]map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Server) SetLabels(v []map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Server) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTags

`func (o *Server) GetTags() []map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Server) GetTagsOk() (*[]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Server) SetTags(v []map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *Server) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEnabled

`func (o *Server) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Server) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Server) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Server) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTagCompliant

`func (o *Server) GetTagCompliant() string`

GetTagCompliant returns the TagCompliant field if non-nil, zero value otherwise.

### GetTagCompliantOk

`func (o *Server) GetTagCompliantOk() (*string, bool)`

GetTagCompliantOk returns a tuple with the TagCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagCompliant

`func (o *Server) SetTagCompliant(v string)`

SetTagCompliant sets TagCompliant field to given value.

### HasTagCompliant

`func (o *Server) HasTagCompliant() bool`

HasTagCompliant returns a boolean if a field has been set.

### GetContainers

`func (o *Server) GetContainers() []int64`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *Server) GetContainersOk() (*[]int64, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *Server) SetContainers(v []int64)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *Server) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetConfig

`func (o *Server) GetConfig() ListHosts200ResponseAllOfServersInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Server) GetConfigOk() (*ListHosts200ResponseAllOfServersInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Server) SetConfig(v ListHosts200ResponseAllOfServersInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Server) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetGuestConsolePreferred

`func (o *Server) GetGuestConsolePreferred() bool`

GetGuestConsolePreferred returns the GuestConsolePreferred field if non-nil, zero value otherwise.

### GetGuestConsolePreferredOk

`func (o *Server) GetGuestConsolePreferredOk() (*bool, bool)`

GetGuestConsolePreferredOk returns a tuple with the GuestConsolePreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsolePreferred

`func (o *Server) SetGuestConsolePreferred(v bool)`

SetGuestConsolePreferred sets GuestConsolePreferred field to given value.

### HasGuestConsolePreferred

`func (o *Server) HasGuestConsolePreferred() bool`

HasGuestConsolePreferred returns a boolean if a field has been set.

### GetGuestConsoleType

`func (o *Server) GetGuestConsoleType() string`

GetGuestConsoleType returns the GuestConsoleType field if non-nil, zero value otherwise.

### GetGuestConsoleTypeOk

`func (o *Server) GetGuestConsoleTypeOk() (*string, bool)`

GetGuestConsoleTypeOk returns a tuple with the GuestConsoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleType

`func (o *Server) SetGuestConsoleType(v string)`

SetGuestConsoleType sets GuestConsoleType field to given value.

### HasGuestConsoleType

`func (o *Server) HasGuestConsoleType() bool`

HasGuestConsoleType returns a boolean if a field has been set.

### GetGuestConsoleUsername

`func (o *Server) GetGuestConsoleUsername() string`

GetGuestConsoleUsername returns the GuestConsoleUsername field if non-nil, zero value otherwise.

### GetGuestConsoleUsernameOk

`func (o *Server) GetGuestConsoleUsernameOk() (*string, bool)`

GetGuestConsoleUsernameOk returns a tuple with the GuestConsoleUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleUsername

`func (o *Server) SetGuestConsoleUsername(v string)`

SetGuestConsoleUsername sets GuestConsoleUsername field to given value.

### HasGuestConsoleUsername

`func (o *Server) HasGuestConsoleUsername() bool`

HasGuestConsoleUsername returns a boolean if a field has been set.

### GetGuestConsolePassword

`func (o *Server) GetGuestConsolePassword() string`

GetGuestConsolePassword returns the GuestConsolePassword field if non-nil, zero value otherwise.

### GetGuestConsolePasswordOk

`func (o *Server) GetGuestConsolePasswordOk() (*string, bool)`

GetGuestConsolePasswordOk returns a tuple with the GuestConsolePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsolePassword

`func (o *Server) SetGuestConsolePassword(v string)`

SetGuestConsolePassword sets GuestConsolePassword field to given value.

### HasGuestConsolePassword

`func (o *Server) HasGuestConsolePassword() bool`

HasGuestConsolePassword returns a boolean if a field has been set.

### GetGuestConsolePasswordHash

`func (o *Server) GetGuestConsolePasswordHash() string`

GetGuestConsolePasswordHash returns the GuestConsolePasswordHash field if non-nil, zero value otherwise.

### GetGuestConsolePasswordHashOk

`func (o *Server) GetGuestConsolePasswordHashOk() (*string, bool)`

GetGuestConsolePasswordHashOk returns a tuple with the GuestConsolePasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsolePasswordHash

`func (o *Server) SetGuestConsolePasswordHash(v string)`

SetGuestConsolePasswordHash sets GuestConsolePasswordHash field to given value.

### HasGuestConsolePasswordHash

`func (o *Server) HasGuestConsolePasswordHash() bool`

HasGuestConsolePasswordHash returns a boolean if a field has been set.

### GetGuestConsolePort

`func (o *Server) GetGuestConsolePort() string`

GetGuestConsolePort returns the GuestConsolePort field if non-nil, zero value otherwise.

### GetGuestConsolePortOk

`func (o *Server) GetGuestConsolePortOk() (*string, bool)`

GetGuestConsolePortOk returns a tuple with the GuestConsolePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsolePort

`func (o *Server) SetGuestConsolePort(v string)`

SetGuestConsolePort sets GuestConsolePort field to given value.

### HasGuestConsolePort

`func (o *Server) HasGuestConsolePort() bool`

HasGuestConsolePort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


