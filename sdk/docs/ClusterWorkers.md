# ClusterWorkers

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
**AccountId** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Owner** | Pointer to [**ListActivity200ResponseAllOfActivityInnerActivityInnerUser**](ListActivity200ResponseAllOfActivityInnerActivityInnerUser.md) |  | [optional] 
**Zone** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Plan** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**ComputeServerType** | Pointer to [**ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType**](ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType.md) |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ZoneId** | Pointer to **int64** |  | [optional] 
**SiteId** | Pointer to **int64** |  | [optional] 
**ResourcePoolId** | Pointer to **int64** |  | [optional] 
**FolderId** | Pointer to **string** |  | [optional] 
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
**OsDevice** | Pointer to **string** |  | [optional] 
**OsType** | Pointer to **string** |  | [optional] 
**DataDevice** | Pointer to **string** |  | [optional] 
**LvmEnabled** | Pointer to **bool** |  | [optional] 
**ApiKey** | Pointer to **string** |  | [optional] 
**SoftwareRaid** | Pointer to **bool** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Stats** | Pointer to [**GetClusterMasters200ResponseMastersInnerStats**](GetClusterMasters200ResponseMastersInnerStats.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**StatusDate** | Pointer to **string** |  | [optional] 
**StatusPercent** | Pointer to **string** |  | [optional] 
**StatusEta** | Pointer to **string** |  | [optional] 
**PowerState** | Pointer to **string** |  | [optional] 
**AgentInstalled** | Pointer to **bool** |  | [optional] 
**LastAgentUpdate** | Pointer to **time.Time** |  | [optional] 
**AgentVersion** | Pointer to **string** |  | [optional] 
**MaxCores** | Pointer to **int64** |  | [optional] 
**CoresPerSocket** | Pointer to **string** |  | [optional] 
**MaxMemory** | Pointer to **int64** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 
**MaxCpu** | Pointer to **int64** |  | [optional] 
**HourlyPrice** | Pointer to **float32** |  | [optional] 
**SourceImage** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**ServerOs** | Pointer to **string** |  | [optional] 
**Volumes** | Pointer to [**[]GetClusterMasters200ResponseMastersInnerVolumesInner**](GetClusterMasters200ResponseMastersInnerVolumesInner.md) |  | [optional] 
**Controllers** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Interfaces** | Pointer to [**[]GetClusterMasters200ResponseMastersInnerInterfacesInner**](GetClusterMasters200ResponseMastersInnerInterfacesInner.md) |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to [**[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**TagCompliant** | Pointer to **string** |  | [optional] 
**Containers** | Pointer to **[]int64** |  | [optional] 
**GuestConsolePreferred** | Pointer to **bool** |  | [optional] 
**GuestConsoleType** | Pointer to **string** |  | [optional] 
**GuestConsoleUsername** | Pointer to **string** |  | [optional] 
**GuestConsolePassword** | Pointer to **string** |  | [optional] 
**GuestConsolePasswordHash** | Pointer to **string** |  | [optional] 
**GuestConsolePort** | Pointer to **string** |  | [optional] 

## Methods

### NewClusterWorkers

`func NewClusterWorkers() *ClusterWorkers`

NewClusterWorkers instantiates a new ClusterWorkers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWorkersWithDefaults

`func NewClusterWorkersWithDefaults() *ClusterWorkers`

NewClusterWorkersWithDefaults instantiates a new ClusterWorkers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterWorkers) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterWorkers) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterWorkers) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterWorkers) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ClusterWorkers) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ClusterWorkers) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ClusterWorkers) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ClusterWorkers) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetExternalId

`func (o *ClusterWorkers) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ClusterWorkers) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ClusterWorkers) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ClusterWorkers) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetInternalId

`func (o *ClusterWorkers) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *ClusterWorkers) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *ClusterWorkers) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *ClusterWorkers) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalUniqueId

`func (o *ClusterWorkers) GetExternalUniqueId() string`

GetExternalUniqueId returns the ExternalUniqueId field if non-nil, zero value otherwise.

### GetExternalUniqueIdOk

`func (o *ClusterWorkers) GetExternalUniqueIdOk() (*string, bool)`

GetExternalUniqueIdOk returns a tuple with the ExternalUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUniqueId

`func (o *ClusterWorkers) SetExternalUniqueId(v string)`

SetExternalUniqueId sets ExternalUniqueId field to given value.

### HasExternalUniqueId

`func (o *ClusterWorkers) HasExternalUniqueId() bool`

HasExternalUniqueId returns a boolean if a field has been set.

### GetName

`func (o *ClusterWorkers) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterWorkers) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterWorkers) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterWorkers) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalName

`func (o *ClusterWorkers) GetExternalName() string`

GetExternalName returns the ExternalName field if non-nil, zero value otherwise.

### GetExternalNameOk

`func (o *ClusterWorkers) GetExternalNameOk() (*string, bool)`

GetExternalNameOk returns a tuple with the ExternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalName

`func (o *ClusterWorkers) SetExternalName(v string)`

SetExternalName sets ExternalName field to given value.

### HasExternalName

`func (o *ClusterWorkers) HasExternalName() bool`

HasExternalName returns a boolean if a field has been set.

### GetHostname

`func (o *ClusterWorkers) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ClusterWorkers) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ClusterWorkers) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ClusterWorkers) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetAccountId

`func (o *ClusterWorkers) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ClusterWorkers) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ClusterWorkers) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ClusterWorkers) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccount

`func (o *ClusterWorkers) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ClusterWorkers) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ClusterWorkers) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ClusterWorkers) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *ClusterWorkers) GetOwner() ListActivity200ResponseAllOfActivityInnerActivityInnerUser`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ClusterWorkers) GetOwnerOk() (*ListActivity200ResponseAllOfActivityInnerActivityInnerUser, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ClusterWorkers) SetOwner(v ListActivity200ResponseAllOfActivityInnerActivityInnerUser)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ClusterWorkers) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetZone

`func (o *ClusterWorkers) GetZone() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ClusterWorkers) GetZoneOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ClusterWorkers) SetZone(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ClusterWorkers) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetPlan

`func (o *ClusterWorkers) GetPlan() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *ClusterWorkers) GetPlanOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *ClusterWorkers) SetPlan(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *ClusterWorkers) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetComputeServerType

`func (o *ClusterWorkers) GetComputeServerType() ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType`

GetComputeServerType returns the ComputeServerType field if non-nil, zero value otherwise.

### GetComputeServerTypeOk

`func (o *ClusterWorkers) GetComputeServerTypeOk() (*ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType, bool)`

GetComputeServerTypeOk returns a tuple with the ComputeServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServerType

`func (o *ClusterWorkers) SetComputeServerType(v ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType)`

SetComputeServerType sets ComputeServerType field to given value.

### HasComputeServerType

`func (o *ClusterWorkers) HasComputeServerType() bool`

HasComputeServerType returns a boolean if a field has been set.

### GetVisibility

`func (o *ClusterWorkers) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ClusterWorkers) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ClusterWorkers) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ClusterWorkers) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *ClusterWorkers) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterWorkers) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterWorkers) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClusterWorkers) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetZoneId

`func (o *ClusterWorkers) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ClusterWorkers) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ClusterWorkers) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *ClusterWorkers) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetSiteId

`func (o *ClusterWorkers) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ClusterWorkers) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ClusterWorkers) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ClusterWorkers) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *ClusterWorkers) GetResourcePoolId() int64`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *ClusterWorkers) GetResourcePoolIdOk() (*int64, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *ClusterWorkers) SetResourcePoolId(v int64)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *ClusterWorkers) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetFolderId

`func (o *ClusterWorkers) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *ClusterWorkers) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *ClusterWorkers) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *ClusterWorkers) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetSshHost

`func (o *ClusterWorkers) GetSshHost() string`

GetSshHost returns the SshHost field if non-nil, zero value otherwise.

### GetSshHostOk

`func (o *ClusterWorkers) GetSshHostOk() (*string, bool)`

GetSshHostOk returns a tuple with the SshHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHost

`func (o *ClusterWorkers) SetSshHost(v string)`

SetSshHost sets SshHost field to given value.

### HasSshHost

`func (o *ClusterWorkers) HasSshHost() bool`

HasSshHost returns a boolean if a field has been set.

### GetSshPort

`func (o *ClusterWorkers) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *ClusterWorkers) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *ClusterWorkers) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *ClusterWorkers) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetExternalIp

`func (o *ClusterWorkers) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *ClusterWorkers) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *ClusterWorkers) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *ClusterWorkers) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### GetInternalIp

`func (o *ClusterWorkers) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *ClusterWorkers) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *ClusterWorkers) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *ClusterWorkers) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### GetVolumeId

`func (o *ClusterWorkers) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *ClusterWorkers) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *ClusterWorkers) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *ClusterWorkers) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.

### GetPlatform

`func (o *ClusterWorkers) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ClusterWorkers) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ClusterWorkers) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ClusterWorkers) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPlatformVersion

`func (o *ClusterWorkers) GetPlatformVersion() string`

GetPlatformVersion returns the PlatformVersion field if non-nil, zero value otherwise.

### GetPlatformVersionOk

`func (o *ClusterWorkers) GetPlatformVersionOk() (*string, bool)`

GetPlatformVersionOk returns a tuple with the PlatformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformVersion

`func (o *ClusterWorkers) SetPlatformVersion(v string)`

SetPlatformVersion sets PlatformVersion field to given value.

### HasPlatformVersion

`func (o *ClusterWorkers) HasPlatformVersion() bool`

HasPlatformVersion returns a boolean if a field has been set.

### GetSshUsername

`func (o *ClusterWorkers) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *ClusterWorkers) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *ClusterWorkers) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *ClusterWorkers) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *ClusterWorkers) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *ClusterWorkers) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *ClusterWorkers) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *ClusterWorkers) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetSshPasswordHash

`func (o *ClusterWorkers) GetSshPasswordHash() string`

GetSshPasswordHash returns the SshPasswordHash field if non-nil, zero value otherwise.

### GetSshPasswordHashOk

`func (o *ClusterWorkers) GetSshPasswordHashOk() (*string, bool)`

GetSshPasswordHashOk returns a tuple with the SshPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPasswordHash

`func (o *ClusterWorkers) SetSshPasswordHash(v string)`

SetSshPasswordHash sets SshPasswordHash field to given value.

### HasSshPasswordHash

`func (o *ClusterWorkers) HasSshPasswordHash() bool`

HasSshPasswordHash returns a boolean if a field has been set.

### GetOsDevice

`func (o *ClusterWorkers) GetOsDevice() string`

GetOsDevice returns the OsDevice field if non-nil, zero value otherwise.

### GetOsDeviceOk

`func (o *ClusterWorkers) GetOsDeviceOk() (*string, bool)`

GetOsDeviceOk returns a tuple with the OsDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDevice

`func (o *ClusterWorkers) SetOsDevice(v string)`

SetOsDevice sets OsDevice field to given value.

### HasOsDevice

`func (o *ClusterWorkers) HasOsDevice() bool`

HasOsDevice returns a boolean if a field has been set.

### GetOsType

`func (o *ClusterWorkers) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *ClusterWorkers) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *ClusterWorkers) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *ClusterWorkers) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetDataDevice

`func (o *ClusterWorkers) GetDataDevice() string`

GetDataDevice returns the DataDevice field if non-nil, zero value otherwise.

### GetDataDeviceOk

`func (o *ClusterWorkers) GetDataDeviceOk() (*string, bool)`

GetDataDeviceOk returns a tuple with the DataDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDevice

`func (o *ClusterWorkers) SetDataDevice(v string)`

SetDataDevice sets DataDevice field to given value.

### HasDataDevice

`func (o *ClusterWorkers) HasDataDevice() bool`

HasDataDevice returns a boolean if a field has been set.

### GetLvmEnabled

`func (o *ClusterWorkers) GetLvmEnabled() bool`

GetLvmEnabled returns the LvmEnabled field if non-nil, zero value otherwise.

### GetLvmEnabledOk

`func (o *ClusterWorkers) GetLvmEnabledOk() (*bool, bool)`

GetLvmEnabledOk returns a tuple with the LvmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLvmEnabled

`func (o *ClusterWorkers) SetLvmEnabled(v bool)`

SetLvmEnabled sets LvmEnabled field to given value.

### HasLvmEnabled

`func (o *ClusterWorkers) HasLvmEnabled() bool`

HasLvmEnabled returns a boolean if a field has been set.

### GetApiKey

`func (o *ClusterWorkers) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ClusterWorkers) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ClusterWorkers) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *ClusterWorkers) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetSoftwareRaid

`func (o *ClusterWorkers) GetSoftwareRaid() bool`

GetSoftwareRaid returns the SoftwareRaid field if non-nil, zero value otherwise.

### GetSoftwareRaidOk

`func (o *ClusterWorkers) GetSoftwareRaidOk() (*bool, bool)`

GetSoftwareRaidOk returns a tuple with the SoftwareRaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareRaid

`func (o *ClusterWorkers) SetSoftwareRaid(v bool)`

SetSoftwareRaid sets SoftwareRaid field to given value.

### HasSoftwareRaid

`func (o *ClusterWorkers) HasSoftwareRaid() bool`

HasSoftwareRaid returns a boolean if a field has been set.

### GetDateCreated

`func (o *ClusterWorkers) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ClusterWorkers) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ClusterWorkers) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ClusterWorkers) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ClusterWorkers) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ClusterWorkers) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ClusterWorkers) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ClusterWorkers) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetStats

`func (o *ClusterWorkers) GetStats() GetClusterMasters200ResponseMastersInnerStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ClusterWorkers) GetStatsOk() (*GetClusterMasters200ResponseMastersInnerStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ClusterWorkers) SetStats(v GetClusterMasters200ResponseMastersInnerStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ClusterWorkers) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterWorkers) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterWorkers) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterWorkers) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterWorkers) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ClusterWorkers) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ClusterWorkers) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ClusterWorkers) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ClusterWorkers) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ClusterWorkers) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ClusterWorkers) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ClusterWorkers) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ClusterWorkers) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetStatusDate

`func (o *ClusterWorkers) GetStatusDate() string`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *ClusterWorkers) GetStatusDateOk() (*string, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *ClusterWorkers) SetStatusDate(v string)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *ClusterWorkers) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetStatusPercent

`func (o *ClusterWorkers) GetStatusPercent() string`

GetStatusPercent returns the StatusPercent field if non-nil, zero value otherwise.

### GetStatusPercentOk

`func (o *ClusterWorkers) GetStatusPercentOk() (*string, bool)`

GetStatusPercentOk returns a tuple with the StatusPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusPercent

`func (o *ClusterWorkers) SetStatusPercent(v string)`

SetStatusPercent sets StatusPercent field to given value.

### HasStatusPercent

`func (o *ClusterWorkers) HasStatusPercent() bool`

HasStatusPercent returns a boolean if a field has been set.

### GetStatusEta

`func (o *ClusterWorkers) GetStatusEta() string`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *ClusterWorkers) GetStatusEtaOk() (*string, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *ClusterWorkers) SetStatusEta(v string)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *ClusterWorkers) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### GetPowerState

`func (o *ClusterWorkers) GetPowerState() string`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *ClusterWorkers) GetPowerStateOk() (*string, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *ClusterWorkers) SetPowerState(v string)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *ClusterWorkers) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetAgentInstalled

`func (o *ClusterWorkers) GetAgentInstalled() bool`

GetAgentInstalled returns the AgentInstalled field if non-nil, zero value otherwise.

### GetAgentInstalledOk

`func (o *ClusterWorkers) GetAgentInstalledOk() (*bool, bool)`

GetAgentInstalledOk returns a tuple with the AgentInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentInstalled

`func (o *ClusterWorkers) SetAgentInstalled(v bool)`

SetAgentInstalled sets AgentInstalled field to given value.

### HasAgentInstalled

`func (o *ClusterWorkers) HasAgentInstalled() bool`

HasAgentInstalled returns a boolean if a field has been set.

### GetLastAgentUpdate

`func (o *ClusterWorkers) GetLastAgentUpdate() time.Time`

GetLastAgentUpdate returns the LastAgentUpdate field if non-nil, zero value otherwise.

### GetLastAgentUpdateOk

`func (o *ClusterWorkers) GetLastAgentUpdateOk() (*time.Time, bool)`

GetLastAgentUpdateOk returns a tuple with the LastAgentUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAgentUpdate

`func (o *ClusterWorkers) SetLastAgentUpdate(v time.Time)`

SetLastAgentUpdate sets LastAgentUpdate field to given value.

### HasLastAgentUpdate

`func (o *ClusterWorkers) HasLastAgentUpdate() bool`

HasLastAgentUpdate returns a boolean if a field has been set.

### GetAgentVersion

`func (o *ClusterWorkers) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *ClusterWorkers) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *ClusterWorkers) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *ClusterWorkers) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### GetMaxCores

`func (o *ClusterWorkers) GetMaxCores() int64`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *ClusterWorkers) GetMaxCoresOk() (*int64, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *ClusterWorkers) SetMaxCores(v int64)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *ClusterWorkers) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *ClusterWorkers) GetCoresPerSocket() string`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *ClusterWorkers) GetCoresPerSocketOk() (*string, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *ClusterWorkers) SetCoresPerSocket(v string)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *ClusterWorkers) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### GetMaxMemory

`func (o *ClusterWorkers) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *ClusterWorkers) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *ClusterWorkers) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *ClusterWorkers) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxStorage

`func (o *ClusterWorkers) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ClusterWorkers) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ClusterWorkers) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *ClusterWorkers) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMaxCpu

`func (o *ClusterWorkers) GetMaxCpu() int64`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *ClusterWorkers) GetMaxCpuOk() (*int64, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *ClusterWorkers) SetMaxCpu(v int64)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *ClusterWorkers) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetHourlyPrice

`func (o *ClusterWorkers) GetHourlyPrice() float32`

GetHourlyPrice returns the HourlyPrice field if non-nil, zero value otherwise.

### GetHourlyPriceOk

`func (o *ClusterWorkers) GetHourlyPriceOk() (*float32, bool)`

GetHourlyPriceOk returns a tuple with the HourlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyPrice

`func (o *ClusterWorkers) SetHourlyPrice(v float32)`

SetHourlyPrice sets HourlyPrice field to given value.

### HasHourlyPrice

`func (o *ClusterWorkers) HasHourlyPrice() bool`

HasHourlyPrice returns a boolean if a field has been set.

### GetSourceImage

`func (o *ClusterWorkers) GetSourceImage() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetSourceImage returns the SourceImage field if non-nil, zero value otherwise.

### GetSourceImageOk

`func (o *ClusterWorkers) GetSourceImageOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetSourceImageOk returns a tuple with the SourceImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceImage

`func (o *ClusterWorkers) SetSourceImage(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetSourceImage sets SourceImage field to given value.

### HasSourceImage

`func (o *ClusterWorkers) HasSourceImage() bool`

HasSourceImage returns a boolean if a field has been set.

### GetServerOs

`func (o *ClusterWorkers) GetServerOs() string`

GetServerOs returns the ServerOs field if non-nil, zero value otherwise.

### GetServerOsOk

`func (o *ClusterWorkers) GetServerOsOk() (*string, bool)`

GetServerOsOk returns a tuple with the ServerOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerOs

`func (o *ClusterWorkers) SetServerOs(v string)`

SetServerOs sets ServerOs field to given value.

### HasServerOs

`func (o *ClusterWorkers) HasServerOs() bool`

HasServerOs returns a boolean if a field has been set.

### GetVolumes

`func (o *ClusterWorkers) GetVolumes() []GetClusterMasters200ResponseMastersInnerVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ClusterWorkers) GetVolumesOk() (*[]GetClusterMasters200ResponseMastersInnerVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ClusterWorkers) SetVolumes(v []GetClusterMasters200ResponseMastersInnerVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *ClusterWorkers) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetControllers

`func (o *ClusterWorkers) GetControllers() []map[string]interface{}`

GetControllers returns the Controllers field if non-nil, zero value otherwise.

### GetControllersOk

`func (o *ClusterWorkers) GetControllersOk() (*[]map[string]interface{}, bool)`

GetControllersOk returns a tuple with the Controllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllers

`func (o *ClusterWorkers) SetControllers(v []map[string]interface{})`

SetControllers sets Controllers field to given value.

### HasControllers

`func (o *ClusterWorkers) HasControllers() bool`

HasControllers returns a boolean if a field has been set.

### GetInterfaces

`func (o *ClusterWorkers) GetInterfaces() []GetClusterMasters200ResponseMastersInnerInterfacesInner`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *ClusterWorkers) GetInterfacesOk() (*[]GetClusterMasters200ResponseMastersInnerInterfacesInner, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *ClusterWorkers) SetInterfaces(v []GetClusterMasters200ResponseMastersInnerInterfacesInner)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *ClusterWorkers) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetLabels

`func (o *ClusterWorkers) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ClusterWorkers) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ClusterWorkers) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ClusterWorkers) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTags

`func (o *ClusterWorkers) GetTags() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ClusterWorkers) GetTagsOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ClusterWorkers) SetTags(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ClusterWorkers) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEnabled

`func (o *ClusterWorkers) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ClusterWorkers) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ClusterWorkers) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ClusterWorkers) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTagCompliant

`func (o *ClusterWorkers) GetTagCompliant() string`

GetTagCompliant returns the TagCompliant field if non-nil, zero value otherwise.

### GetTagCompliantOk

`func (o *ClusterWorkers) GetTagCompliantOk() (*string, bool)`

GetTagCompliantOk returns a tuple with the TagCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagCompliant

`func (o *ClusterWorkers) SetTagCompliant(v string)`

SetTagCompliant sets TagCompliant field to given value.

### HasTagCompliant

`func (o *ClusterWorkers) HasTagCompliant() bool`

HasTagCompliant returns a boolean if a field has been set.

### GetContainers

`func (o *ClusterWorkers) GetContainers() []int64`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *ClusterWorkers) GetContainersOk() (*[]int64, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *ClusterWorkers) SetContainers(v []int64)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *ClusterWorkers) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetGuestConsolePreferred

`func (o *ClusterWorkers) GetGuestConsolePreferred() bool`

GetGuestConsolePreferred returns the GuestConsolePreferred field if non-nil, zero value otherwise.

### GetGuestConsolePreferredOk

`func (o *ClusterWorkers) GetGuestConsolePreferredOk() (*bool, bool)`

GetGuestConsolePreferredOk returns a tuple with the GuestConsolePreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsolePreferred

`func (o *ClusterWorkers) SetGuestConsolePreferred(v bool)`

SetGuestConsolePreferred sets GuestConsolePreferred field to given value.

### HasGuestConsolePreferred

`func (o *ClusterWorkers) HasGuestConsolePreferred() bool`

HasGuestConsolePreferred returns a boolean if a field has been set.

### GetGuestConsoleType

`func (o *ClusterWorkers) GetGuestConsoleType() string`

GetGuestConsoleType returns the GuestConsoleType field if non-nil, zero value otherwise.

### GetGuestConsoleTypeOk

`func (o *ClusterWorkers) GetGuestConsoleTypeOk() (*string, bool)`

GetGuestConsoleTypeOk returns a tuple with the GuestConsoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleType

`func (o *ClusterWorkers) SetGuestConsoleType(v string)`

SetGuestConsoleType sets GuestConsoleType field to given value.

### HasGuestConsoleType

`func (o *ClusterWorkers) HasGuestConsoleType() bool`

HasGuestConsoleType returns a boolean if a field has been set.

### GetGuestConsoleUsername

`func (o *ClusterWorkers) GetGuestConsoleUsername() string`

GetGuestConsoleUsername returns the GuestConsoleUsername field if non-nil, zero value otherwise.

### GetGuestConsoleUsernameOk

`func (o *ClusterWorkers) GetGuestConsoleUsernameOk() (*string, bool)`

GetGuestConsoleUsernameOk returns a tuple with the GuestConsoleUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleUsername

`func (o *ClusterWorkers) SetGuestConsoleUsername(v string)`

SetGuestConsoleUsername sets GuestConsoleUsername field to given value.

### HasGuestConsoleUsername

`func (o *ClusterWorkers) HasGuestConsoleUsername() bool`

HasGuestConsoleUsername returns a boolean if a field has been set.

### GetGuestConsolePassword

`func (o *ClusterWorkers) GetGuestConsolePassword() string`

GetGuestConsolePassword returns the GuestConsolePassword field if non-nil, zero value otherwise.

### GetGuestConsolePasswordOk

`func (o *ClusterWorkers) GetGuestConsolePasswordOk() (*string, bool)`

GetGuestConsolePasswordOk returns a tuple with the GuestConsolePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsolePassword

`func (o *ClusterWorkers) SetGuestConsolePassword(v string)`

SetGuestConsolePassword sets GuestConsolePassword field to given value.

### HasGuestConsolePassword

`func (o *ClusterWorkers) HasGuestConsolePassword() bool`

HasGuestConsolePassword returns a boolean if a field has been set.

### GetGuestConsolePasswordHash

`func (o *ClusterWorkers) GetGuestConsolePasswordHash() string`

GetGuestConsolePasswordHash returns the GuestConsolePasswordHash field if non-nil, zero value otherwise.

### GetGuestConsolePasswordHashOk

`func (o *ClusterWorkers) GetGuestConsolePasswordHashOk() (*string, bool)`

GetGuestConsolePasswordHashOk returns a tuple with the GuestConsolePasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsolePasswordHash

`func (o *ClusterWorkers) SetGuestConsolePasswordHash(v string)`

SetGuestConsolePasswordHash sets GuestConsolePasswordHash field to given value.

### HasGuestConsolePasswordHash

`func (o *ClusterWorkers) HasGuestConsolePasswordHash() bool`

HasGuestConsolePasswordHash returns a boolean if a field has been set.

### GetGuestConsolePort

`func (o *ClusterWorkers) GetGuestConsolePort() string`

GetGuestConsolePort returns the GuestConsolePort field if non-nil, zero value otherwise.

### GetGuestConsolePortOk

`func (o *ClusterWorkers) GetGuestConsolePortOk() (*string, bool)`

GetGuestConsolePortOk returns a tuple with the GuestConsolePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsolePort

`func (o *ClusterWorkers) SetGuestConsolePort(v string)`

SetGuestConsolePort sets GuestConsolePort field to given value.

### HasGuestConsolePort

`func (o *ClusterWorkers) HasGuestConsolePort() bool`

HasGuestConsolePort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


