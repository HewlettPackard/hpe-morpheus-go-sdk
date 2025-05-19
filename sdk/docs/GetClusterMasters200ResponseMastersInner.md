# GetClusterMasters200ResponseMastersInner

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

### NewGetClusterMasters200ResponseMastersInner

`func NewGetClusterMasters200ResponseMastersInner() *GetClusterMasters200ResponseMastersInner`

NewGetClusterMasters200ResponseMastersInner instantiates a new GetClusterMasters200ResponseMastersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClusterMasters200ResponseMastersInnerWithDefaults

`func NewGetClusterMasters200ResponseMastersInnerWithDefaults() *GetClusterMasters200ResponseMastersInner`

NewGetClusterMasters200ResponseMastersInnerWithDefaults instantiates a new GetClusterMasters200ResponseMastersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetClusterMasters200ResponseMastersInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetClusterMasters200ResponseMastersInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetClusterMasters200ResponseMastersInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetClusterMasters200ResponseMastersInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *GetClusterMasters200ResponseMastersInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetClusterMasters200ResponseMastersInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetClusterMasters200ResponseMastersInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetClusterMasters200ResponseMastersInner) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetExternalId

`func (o *GetClusterMasters200ResponseMastersInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetClusterMasters200ResponseMastersInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetClusterMasters200ResponseMastersInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetClusterMasters200ResponseMastersInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetInternalId

`func (o *GetClusterMasters200ResponseMastersInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *GetClusterMasters200ResponseMastersInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *GetClusterMasters200ResponseMastersInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *GetClusterMasters200ResponseMastersInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalUniqueId

`func (o *GetClusterMasters200ResponseMastersInner) GetExternalUniqueId() string`

GetExternalUniqueId returns the ExternalUniqueId field if non-nil, zero value otherwise.

### GetExternalUniqueIdOk

`func (o *GetClusterMasters200ResponseMastersInner) GetExternalUniqueIdOk() (*string, bool)`

GetExternalUniqueIdOk returns a tuple with the ExternalUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUniqueId

`func (o *GetClusterMasters200ResponseMastersInner) SetExternalUniqueId(v string)`

SetExternalUniqueId sets ExternalUniqueId field to given value.

### HasExternalUniqueId

`func (o *GetClusterMasters200ResponseMastersInner) HasExternalUniqueId() bool`

HasExternalUniqueId returns a boolean if a field has been set.

### GetName

`func (o *GetClusterMasters200ResponseMastersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetClusterMasters200ResponseMastersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetClusterMasters200ResponseMastersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetClusterMasters200ResponseMastersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalName

`func (o *GetClusterMasters200ResponseMastersInner) GetExternalName() string`

GetExternalName returns the ExternalName field if non-nil, zero value otherwise.

### GetExternalNameOk

`func (o *GetClusterMasters200ResponseMastersInner) GetExternalNameOk() (*string, bool)`

GetExternalNameOk returns a tuple with the ExternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalName

`func (o *GetClusterMasters200ResponseMastersInner) SetExternalName(v string)`

SetExternalName sets ExternalName field to given value.

### HasExternalName

`func (o *GetClusterMasters200ResponseMastersInner) HasExternalName() bool`

HasExternalName returns a boolean if a field has been set.

### GetHostname

`func (o *GetClusterMasters200ResponseMastersInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *GetClusterMasters200ResponseMastersInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *GetClusterMasters200ResponseMastersInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *GetClusterMasters200ResponseMastersInner) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetAccountId

`func (o *GetClusterMasters200ResponseMastersInner) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetClusterMasters200ResponseMastersInner) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetClusterMasters200ResponseMastersInner) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetClusterMasters200ResponseMastersInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccount

`func (o *GetClusterMasters200ResponseMastersInner) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetClusterMasters200ResponseMastersInner) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetClusterMasters200ResponseMastersInner) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetClusterMasters200ResponseMastersInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *GetClusterMasters200ResponseMastersInner) GetOwner() ListActivity200ResponseAllOfActivityInnerActivityInnerUser`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetClusterMasters200ResponseMastersInner) GetOwnerOk() (*ListActivity200ResponseAllOfActivityInnerActivityInnerUser, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetClusterMasters200ResponseMastersInner) SetOwner(v ListActivity200ResponseAllOfActivityInnerActivityInnerUser)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetClusterMasters200ResponseMastersInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetZone

`func (o *GetClusterMasters200ResponseMastersInner) GetZone() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetClusterMasters200ResponseMastersInner) GetZoneOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetClusterMasters200ResponseMastersInner) SetZone(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetClusterMasters200ResponseMastersInner) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetPlan

`func (o *GetClusterMasters200ResponseMastersInner) GetPlan() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *GetClusterMasters200ResponseMastersInner) GetPlanOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *GetClusterMasters200ResponseMastersInner) SetPlan(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *GetClusterMasters200ResponseMastersInner) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetComputeServerType

`func (o *GetClusterMasters200ResponseMastersInner) GetComputeServerType() ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType`

GetComputeServerType returns the ComputeServerType field if non-nil, zero value otherwise.

### GetComputeServerTypeOk

`func (o *GetClusterMasters200ResponseMastersInner) GetComputeServerTypeOk() (*ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType, bool)`

GetComputeServerTypeOk returns a tuple with the ComputeServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServerType

`func (o *GetClusterMasters200ResponseMastersInner) SetComputeServerType(v ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType)`

SetComputeServerType sets ComputeServerType field to given value.

### HasComputeServerType

`func (o *GetClusterMasters200ResponseMastersInner) HasComputeServerType() bool`

HasComputeServerType returns a boolean if a field has been set.

### GetVisibility

`func (o *GetClusterMasters200ResponseMastersInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetClusterMasters200ResponseMastersInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetClusterMasters200ResponseMastersInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetClusterMasters200ResponseMastersInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *GetClusterMasters200ResponseMastersInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetClusterMasters200ResponseMastersInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetClusterMasters200ResponseMastersInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetClusterMasters200ResponseMastersInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetZoneId

`func (o *GetClusterMasters200ResponseMastersInner) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *GetClusterMasters200ResponseMastersInner) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *GetClusterMasters200ResponseMastersInner) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *GetClusterMasters200ResponseMastersInner) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetSiteId

`func (o *GetClusterMasters200ResponseMastersInner) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *GetClusterMasters200ResponseMastersInner) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *GetClusterMasters200ResponseMastersInner) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *GetClusterMasters200ResponseMastersInner) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *GetClusterMasters200ResponseMastersInner) GetResourcePoolId() int64`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *GetClusterMasters200ResponseMastersInner) GetResourcePoolIdOk() (*int64, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *GetClusterMasters200ResponseMastersInner) SetResourcePoolId(v int64)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *GetClusterMasters200ResponseMastersInner) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetFolderId

`func (o *GetClusterMasters200ResponseMastersInner) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *GetClusterMasters200ResponseMastersInner) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *GetClusterMasters200ResponseMastersInner) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *GetClusterMasters200ResponseMastersInner) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetSshHost

`func (o *GetClusterMasters200ResponseMastersInner) GetSshHost() string`

GetSshHost returns the SshHost field if non-nil, zero value otherwise.

### GetSshHostOk

`func (o *GetClusterMasters200ResponseMastersInner) GetSshHostOk() (*string, bool)`

GetSshHostOk returns a tuple with the SshHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHost

`func (o *GetClusterMasters200ResponseMastersInner) SetSshHost(v string)`

SetSshHost sets SshHost field to given value.

### HasSshHost

`func (o *GetClusterMasters200ResponseMastersInner) HasSshHost() bool`

HasSshHost returns a boolean if a field has been set.

### GetSshPort

`func (o *GetClusterMasters200ResponseMastersInner) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *GetClusterMasters200ResponseMastersInner) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *GetClusterMasters200ResponseMastersInner) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *GetClusterMasters200ResponseMastersInner) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetExternalIp

`func (o *GetClusterMasters200ResponseMastersInner) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *GetClusterMasters200ResponseMastersInner) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *GetClusterMasters200ResponseMastersInner) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *GetClusterMasters200ResponseMastersInner) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### GetInternalIp

`func (o *GetClusterMasters200ResponseMastersInner) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *GetClusterMasters200ResponseMastersInner) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *GetClusterMasters200ResponseMastersInner) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *GetClusterMasters200ResponseMastersInner) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### GetVolumeId

`func (o *GetClusterMasters200ResponseMastersInner) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *GetClusterMasters200ResponseMastersInner) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *GetClusterMasters200ResponseMastersInner) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *GetClusterMasters200ResponseMastersInner) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.

### GetPlatform

`func (o *GetClusterMasters200ResponseMastersInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetClusterMasters200ResponseMastersInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetClusterMasters200ResponseMastersInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetClusterMasters200ResponseMastersInner) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPlatformVersion

`func (o *GetClusterMasters200ResponseMastersInner) GetPlatformVersion() string`

GetPlatformVersion returns the PlatformVersion field if non-nil, zero value otherwise.

### GetPlatformVersionOk

`func (o *GetClusterMasters200ResponseMastersInner) GetPlatformVersionOk() (*string, bool)`

GetPlatformVersionOk returns a tuple with the PlatformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformVersion

`func (o *GetClusterMasters200ResponseMastersInner) SetPlatformVersion(v string)`

SetPlatformVersion sets PlatformVersion field to given value.

### HasPlatformVersion

`func (o *GetClusterMasters200ResponseMastersInner) HasPlatformVersion() bool`

HasPlatformVersion returns a boolean if a field has been set.

### GetSshUsername

`func (o *GetClusterMasters200ResponseMastersInner) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *GetClusterMasters200ResponseMastersInner) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *GetClusterMasters200ResponseMastersInner) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *GetClusterMasters200ResponseMastersInner) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *GetClusterMasters200ResponseMastersInner) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *GetClusterMasters200ResponseMastersInner) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *GetClusterMasters200ResponseMastersInner) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *GetClusterMasters200ResponseMastersInner) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetSshPasswordHash

`func (o *GetClusterMasters200ResponseMastersInner) GetSshPasswordHash() string`

GetSshPasswordHash returns the SshPasswordHash field if non-nil, zero value otherwise.

### GetSshPasswordHashOk

`func (o *GetClusterMasters200ResponseMastersInner) GetSshPasswordHashOk() (*string, bool)`

GetSshPasswordHashOk returns a tuple with the SshPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPasswordHash

`func (o *GetClusterMasters200ResponseMastersInner) SetSshPasswordHash(v string)`

SetSshPasswordHash sets SshPasswordHash field to given value.

### HasSshPasswordHash

`func (o *GetClusterMasters200ResponseMastersInner) HasSshPasswordHash() bool`

HasSshPasswordHash returns a boolean if a field has been set.

### GetOsDevice

`func (o *GetClusterMasters200ResponseMastersInner) GetOsDevice() string`

GetOsDevice returns the OsDevice field if non-nil, zero value otherwise.

### GetOsDeviceOk

`func (o *GetClusterMasters200ResponseMastersInner) GetOsDeviceOk() (*string, bool)`

GetOsDeviceOk returns a tuple with the OsDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDevice

`func (o *GetClusterMasters200ResponseMastersInner) SetOsDevice(v string)`

SetOsDevice sets OsDevice field to given value.

### HasOsDevice

`func (o *GetClusterMasters200ResponseMastersInner) HasOsDevice() bool`

HasOsDevice returns a boolean if a field has been set.

### GetOsType

`func (o *GetClusterMasters200ResponseMastersInner) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *GetClusterMasters200ResponseMastersInner) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *GetClusterMasters200ResponseMastersInner) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *GetClusterMasters200ResponseMastersInner) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetDataDevice

`func (o *GetClusterMasters200ResponseMastersInner) GetDataDevice() string`

GetDataDevice returns the DataDevice field if non-nil, zero value otherwise.

### GetDataDeviceOk

`func (o *GetClusterMasters200ResponseMastersInner) GetDataDeviceOk() (*string, bool)`

GetDataDeviceOk returns a tuple with the DataDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDevice

`func (o *GetClusterMasters200ResponseMastersInner) SetDataDevice(v string)`

SetDataDevice sets DataDevice field to given value.

### HasDataDevice

`func (o *GetClusterMasters200ResponseMastersInner) HasDataDevice() bool`

HasDataDevice returns a boolean if a field has been set.

### GetLvmEnabled

`func (o *GetClusterMasters200ResponseMastersInner) GetLvmEnabled() bool`

GetLvmEnabled returns the LvmEnabled field if non-nil, zero value otherwise.

### GetLvmEnabledOk

`func (o *GetClusterMasters200ResponseMastersInner) GetLvmEnabledOk() (*bool, bool)`

GetLvmEnabledOk returns a tuple with the LvmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLvmEnabled

`func (o *GetClusterMasters200ResponseMastersInner) SetLvmEnabled(v bool)`

SetLvmEnabled sets LvmEnabled field to given value.

### HasLvmEnabled

`func (o *GetClusterMasters200ResponseMastersInner) HasLvmEnabled() bool`

HasLvmEnabled returns a boolean if a field has been set.

### GetApiKey

`func (o *GetClusterMasters200ResponseMastersInner) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *GetClusterMasters200ResponseMastersInner) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *GetClusterMasters200ResponseMastersInner) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *GetClusterMasters200ResponseMastersInner) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetSoftwareRaid

`func (o *GetClusterMasters200ResponseMastersInner) GetSoftwareRaid() bool`

GetSoftwareRaid returns the SoftwareRaid field if non-nil, zero value otherwise.

### GetSoftwareRaidOk

`func (o *GetClusterMasters200ResponseMastersInner) GetSoftwareRaidOk() (*bool, bool)`

GetSoftwareRaidOk returns a tuple with the SoftwareRaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareRaid

`func (o *GetClusterMasters200ResponseMastersInner) SetSoftwareRaid(v bool)`

SetSoftwareRaid sets SoftwareRaid field to given value.

### HasSoftwareRaid

`func (o *GetClusterMasters200ResponseMastersInner) HasSoftwareRaid() bool`

HasSoftwareRaid returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetClusterMasters200ResponseMastersInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetClusterMasters200ResponseMastersInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetClusterMasters200ResponseMastersInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetClusterMasters200ResponseMastersInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetClusterMasters200ResponseMastersInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetClusterMasters200ResponseMastersInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetClusterMasters200ResponseMastersInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetClusterMasters200ResponseMastersInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetStats

`func (o *GetClusterMasters200ResponseMastersInner) GetStats() GetClusterMasters200ResponseMastersInnerStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *GetClusterMasters200ResponseMastersInner) GetStatsOk() (*GetClusterMasters200ResponseMastersInnerStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *GetClusterMasters200ResponseMastersInner) SetStats(v GetClusterMasters200ResponseMastersInnerStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *GetClusterMasters200ResponseMastersInner) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetStatus

`func (o *GetClusterMasters200ResponseMastersInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetClusterMasters200ResponseMastersInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetClusterMasters200ResponseMastersInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetClusterMasters200ResponseMastersInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *GetClusterMasters200ResponseMastersInner) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetClusterMasters200ResponseMastersInner) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetClusterMasters200ResponseMastersInner) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GetClusterMasters200ResponseMastersInner) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetErrorMessage

`func (o *GetClusterMasters200ResponseMastersInner) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *GetClusterMasters200ResponseMastersInner) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *GetClusterMasters200ResponseMastersInner) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *GetClusterMasters200ResponseMastersInner) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetStatusDate

`func (o *GetClusterMasters200ResponseMastersInner) GetStatusDate() string`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *GetClusterMasters200ResponseMastersInner) GetStatusDateOk() (*string, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *GetClusterMasters200ResponseMastersInner) SetStatusDate(v string)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *GetClusterMasters200ResponseMastersInner) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetStatusPercent

`func (o *GetClusterMasters200ResponseMastersInner) GetStatusPercent() string`

GetStatusPercent returns the StatusPercent field if non-nil, zero value otherwise.

### GetStatusPercentOk

`func (o *GetClusterMasters200ResponseMastersInner) GetStatusPercentOk() (*string, bool)`

GetStatusPercentOk returns a tuple with the StatusPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusPercent

`func (o *GetClusterMasters200ResponseMastersInner) SetStatusPercent(v string)`

SetStatusPercent sets StatusPercent field to given value.

### HasStatusPercent

`func (o *GetClusterMasters200ResponseMastersInner) HasStatusPercent() bool`

HasStatusPercent returns a boolean if a field has been set.

### GetStatusEta

`func (o *GetClusterMasters200ResponseMastersInner) GetStatusEta() string`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *GetClusterMasters200ResponseMastersInner) GetStatusEtaOk() (*string, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *GetClusterMasters200ResponseMastersInner) SetStatusEta(v string)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *GetClusterMasters200ResponseMastersInner) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### GetPowerState

`func (o *GetClusterMasters200ResponseMastersInner) GetPowerState() string`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *GetClusterMasters200ResponseMastersInner) GetPowerStateOk() (*string, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *GetClusterMasters200ResponseMastersInner) SetPowerState(v string)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *GetClusterMasters200ResponseMastersInner) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetAgentInstalled

`func (o *GetClusterMasters200ResponseMastersInner) GetAgentInstalled() bool`

GetAgentInstalled returns the AgentInstalled field if non-nil, zero value otherwise.

### GetAgentInstalledOk

`func (o *GetClusterMasters200ResponseMastersInner) GetAgentInstalledOk() (*bool, bool)`

GetAgentInstalledOk returns a tuple with the AgentInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentInstalled

`func (o *GetClusterMasters200ResponseMastersInner) SetAgentInstalled(v bool)`

SetAgentInstalled sets AgentInstalled field to given value.

### HasAgentInstalled

`func (o *GetClusterMasters200ResponseMastersInner) HasAgentInstalled() bool`

HasAgentInstalled returns a boolean if a field has been set.

### GetLastAgentUpdate

`func (o *GetClusterMasters200ResponseMastersInner) GetLastAgentUpdate() time.Time`

GetLastAgentUpdate returns the LastAgentUpdate field if non-nil, zero value otherwise.

### GetLastAgentUpdateOk

`func (o *GetClusterMasters200ResponseMastersInner) GetLastAgentUpdateOk() (*time.Time, bool)`

GetLastAgentUpdateOk returns a tuple with the LastAgentUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAgentUpdate

`func (o *GetClusterMasters200ResponseMastersInner) SetLastAgentUpdate(v time.Time)`

SetLastAgentUpdate sets LastAgentUpdate field to given value.

### HasLastAgentUpdate

`func (o *GetClusterMasters200ResponseMastersInner) HasLastAgentUpdate() bool`

HasLastAgentUpdate returns a boolean if a field has been set.

### GetAgentVersion

`func (o *GetClusterMasters200ResponseMastersInner) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *GetClusterMasters200ResponseMastersInner) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *GetClusterMasters200ResponseMastersInner) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *GetClusterMasters200ResponseMastersInner) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### GetMaxCores

`func (o *GetClusterMasters200ResponseMastersInner) GetMaxCores() int64`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *GetClusterMasters200ResponseMastersInner) GetMaxCoresOk() (*int64, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *GetClusterMasters200ResponseMastersInner) SetMaxCores(v int64)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *GetClusterMasters200ResponseMastersInner) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *GetClusterMasters200ResponseMastersInner) GetCoresPerSocket() string`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *GetClusterMasters200ResponseMastersInner) GetCoresPerSocketOk() (*string, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *GetClusterMasters200ResponseMastersInner) SetCoresPerSocket(v string)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *GetClusterMasters200ResponseMastersInner) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### GetMaxMemory

`func (o *GetClusterMasters200ResponseMastersInner) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *GetClusterMasters200ResponseMastersInner) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *GetClusterMasters200ResponseMastersInner) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *GetClusterMasters200ResponseMastersInner) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxStorage

`func (o *GetClusterMasters200ResponseMastersInner) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *GetClusterMasters200ResponseMastersInner) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *GetClusterMasters200ResponseMastersInner) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *GetClusterMasters200ResponseMastersInner) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMaxCpu

`func (o *GetClusterMasters200ResponseMastersInner) GetMaxCpu() int64`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *GetClusterMasters200ResponseMastersInner) GetMaxCpuOk() (*int64, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *GetClusterMasters200ResponseMastersInner) SetMaxCpu(v int64)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *GetClusterMasters200ResponseMastersInner) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetHourlyPrice

`func (o *GetClusterMasters200ResponseMastersInner) GetHourlyPrice() float32`

GetHourlyPrice returns the HourlyPrice field if non-nil, zero value otherwise.

### GetHourlyPriceOk

`func (o *GetClusterMasters200ResponseMastersInner) GetHourlyPriceOk() (*float32, bool)`

GetHourlyPriceOk returns a tuple with the HourlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyPrice

`func (o *GetClusterMasters200ResponseMastersInner) SetHourlyPrice(v float32)`

SetHourlyPrice sets HourlyPrice field to given value.

### HasHourlyPrice

`func (o *GetClusterMasters200ResponseMastersInner) HasHourlyPrice() bool`

HasHourlyPrice returns a boolean if a field has been set.

### GetSourceImage

`func (o *GetClusterMasters200ResponseMastersInner) GetSourceImage() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetSourceImage returns the SourceImage field if non-nil, zero value otherwise.

### GetSourceImageOk

`func (o *GetClusterMasters200ResponseMastersInner) GetSourceImageOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetSourceImageOk returns a tuple with the SourceImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceImage

`func (o *GetClusterMasters200ResponseMastersInner) SetSourceImage(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetSourceImage sets SourceImage field to given value.

### HasSourceImage

`func (o *GetClusterMasters200ResponseMastersInner) HasSourceImage() bool`

HasSourceImage returns a boolean if a field has been set.

### GetServerOs

`func (o *GetClusterMasters200ResponseMastersInner) GetServerOs() string`

GetServerOs returns the ServerOs field if non-nil, zero value otherwise.

### GetServerOsOk

`func (o *GetClusterMasters200ResponseMastersInner) GetServerOsOk() (*string, bool)`

GetServerOsOk returns a tuple with the ServerOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerOs

`func (o *GetClusterMasters200ResponseMastersInner) SetServerOs(v string)`

SetServerOs sets ServerOs field to given value.

### HasServerOs

`func (o *GetClusterMasters200ResponseMastersInner) HasServerOs() bool`

HasServerOs returns a boolean if a field has been set.

### GetVolumes

`func (o *GetClusterMasters200ResponseMastersInner) GetVolumes() []GetClusterMasters200ResponseMastersInnerVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *GetClusterMasters200ResponseMastersInner) GetVolumesOk() (*[]GetClusterMasters200ResponseMastersInnerVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *GetClusterMasters200ResponseMastersInner) SetVolumes(v []GetClusterMasters200ResponseMastersInnerVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *GetClusterMasters200ResponseMastersInner) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetControllers

`func (o *GetClusterMasters200ResponseMastersInner) GetControllers() []map[string]interface{}`

GetControllers returns the Controllers field if non-nil, zero value otherwise.

### GetControllersOk

`func (o *GetClusterMasters200ResponseMastersInner) GetControllersOk() (*[]map[string]interface{}, bool)`

GetControllersOk returns a tuple with the Controllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllers

`func (o *GetClusterMasters200ResponseMastersInner) SetControllers(v []map[string]interface{})`

SetControllers sets Controllers field to given value.

### HasControllers

`func (o *GetClusterMasters200ResponseMastersInner) HasControllers() bool`

HasControllers returns a boolean if a field has been set.

### GetInterfaces

`func (o *GetClusterMasters200ResponseMastersInner) GetInterfaces() []GetClusterMasters200ResponseMastersInnerInterfacesInner`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *GetClusterMasters200ResponseMastersInner) GetInterfacesOk() (*[]GetClusterMasters200ResponseMastersInnerInterfacesInner, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *GetClusterMasters200ResponseMastersInner) SetInterfaces(v []GetClusterMasters200ResponseMastersInnerInterfacesInner)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *GetClusterMasters200ResponseMastersInner) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetLabels

`func (o *GetClusterMasters200ResponseMastersInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetClusterMasters200ResponseMastersInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetClusterMasters200ResponseMastersInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetClusterMasters200ResponseMastersInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTags

`func (o *GetClusterMasters200ResponseMastersInner) GetTags() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetClusterMasters200ResponseMastersInner) GetTagsOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetClusterMasters200ResponseMastersInner) SetTags(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetClusterMasters200ResponseMastersInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEnabled

`func (o *GetClusterMasters200ResponseMastersInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetClusterMasters200ResponseMastersInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetClusterMasters200ResponseMastersInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetClusterMasters200ResponseMastersInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTagCompliant

`func (o *GetClusterMasters200ResponseMastersInner) GetTagCompliant() string`

GetTagCompliant returns the TagCompliant field if non-nil, zero value otherwise.

### GetTagCompliantOk

`func (o *GetClusterMasters200ResponseMastersInner) GetTagCompliantOk() (*string, bool)`

GetTagCompliantOk returns a tuple with the TagCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagCompliant

`func (o *GetClusterMasters200ResponseMastersInner) SetTagCompliant(v string)`

SetTagCompliant sets TagCompliant field to given value.

### HasTagCompliant

`func (o *GetClusterMasters200ResponseMastersInner) HasTagCompliant() bool`

HasTagCompliant returns a boolean if a field has been set.

### GetContainers

`func (o *GetClusterMasters200ResponseMastersInner) GetContainers() []int64`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *GetClusterMasters200ResponseMastersInner) GetContainersOk() (*[]int64, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *GetClusterMasters200ResponseMastersInner) SetContainers(v []int64)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *GetClusterMasters200ResponseMastersInner) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetGuestConsolePreferred

`func (o *GetClusterMasters200ResponseMastersInner) GetGuestConsolePreferred() bool`

GetGuestConsolePreferred returns the GuestConsolePreferred field if non-nil, zero value otherwise.

### GetGuestConsolePreferredOk

`func (o *GetClusterMasters200ResponseMastersInner) GetGuestConsolePreferredOk() (*bool, bool)`

GetGuestConsolePreferredOk returns a tuple with the GuestConsolePreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsolePreferred

`func (o *GetClusterMasters200ResponseMastersInner) SetGuestConsolePreferred(v bool)`

SetGuestConsolePreferred sets GuestConsolePreferred field to given value.

### HasGuestConsolePreferred

`func (o *GetClusterMasters200ResponseMastersInner) HasGuestConsolePreferred() bool`

HasGuestConsolePreferred returns a boolean if a field has been set.

### GetGuestConsoleType

`func (o *GetClusterMasters200ResponseMastersInner) GetGuestConsoleType() string`

GetGuestConsoleType returns the GuestConsoleType field if non-nil, zero value otherwise.

### GetGuestConsoleTypeOk

`func (o *GetClusterMasters200ResponseMastersInner) GetGuestConsoleTypeOk() (*string, bool)`

GetGuestConsoleTypeOk returns a tuple with the GuestConsoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleType

`func (o *GetClusterMasters200ResponseMastersInner) SetGuestConsoleType(v string)`

SetGuestConsoleType sets GuestConsoleType field to given value.

### HasGuestConsoleType

`func (o *GetClusterMasters200ResponseMastersInner) HasGuestConsoleType() bool`

HasGuestConsoleType returns a boolean if a field has been set.

### GetGuestConsoleUsername

`func (o *GetClusterMasters200ResponseMastersInner) GetGuestConsoleUsername() string`

GetGuestConsoleUsername returns the GuestConsoleUsername field if non-nil, zero value otherwise.

### GetGuestConsoleUsernameOk

`func (o *GetClusterMasters200ResponseMastersInner) GetGuestConsoleUsernameOk() (*string, bool)`

GetGuestConsoleUsernameOk returns a tuple with the GuestConsoleUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleUsername

`func (o *GetClusterMasters200ResponseMastersInner) SetGuestConsoleUsername(v string)`

SetGuestConsoleUsername sets GuestConsoleUsername field to given value.

### HasGuestConsoleUsername

`func (o *GetClusterMasters200ResponseMastersInner) HasGuestConsoleUsername() bool`

HasGuestConsoleUsername returns a boolean if a field has been set.

### GetGuestConsolePassword

`func (o *GetClusterMasters200ResponseMastersInner) GetGuestConsolePassword() string`

GetGuestConsolePassword returns the GuestConsolePassword field if non-nil, zero value otherwise.

### GetGuestConsolePasswordOk

`func (o *GetClusterMasters200ResponseMastersInner) GetGuestConsolePasswordOk() (*string, bool)`

GetGuestConsolePasswordOk returns a tuple with the GuestConsolePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsolePassword

`func (o *GetClusterMasters200ResponseMastersInner) SetGuestConsolePassword(v string)`

SetGuestConsolePassword sets GuestConsolePassword field to given value.

### HasGuestConsolePassword

`func (o *GetClusterMasters200ResponseMastersInner) HasGuestConsolePassword() bool`

HasGuestConsolePassword returns a boolean if a field has been set.

### GetGuestConsolePasswordHash

`func (o *GetClusterMasters200ResponseMastersInner) GetGuestConsolePasswordHash() string`

GetGuestConsolePasswordHash returns the GuestConsolePasswordHash field if non-nil, zero value otherwise.

### GetGuestConsolePasswordHashOk

`func (o *GetClusterMasters200ResponseMastersInner) GetGuestConsolePasswordHashOk() (*string, bool)`

GetGuestConsolePasswordHashOk returns a tuple with the GuestConsolePasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsolePasswordHash

`func (o *GetClusterMasters200ResponseMastersInner) SetGuestConsolePasswordHash(v string)`

SetGuestConsolePasswordHash sets GuestConsolePasswordHash field to given value.

### HasGuestConsolePasswordHash

`func (o *GetClusterMasters200ResponseMastersInner) HasGuestConsolePasswordHash() bool`

HasGuestConsolePasswordHash returns a boolean if a field has been set.

### GetGuestConsolePort

`func (o *GetClusterMasters200ResponseMastersInner) GetGuestConsolePort() string`

GetGuestConsolePort returns the GuestConsolePort field if non-nil, zero value otherwise.

### GetGuestConsolePortOk

`func (o *GetClusterMasters200ResponseMastersInner) GetGuestConsolePortOk() (*string, bool)`

GetGuestConsolePortOk returns a tuple with the GuestConsolePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsolePort

`func (o *GetClusterMasters200ResponseMastersInner) SetGuestConsolePort(v string)`

SetGuestConsolePort sets GuestConsolePort field to given value.

### HasGuestConsolePort

`func (o *GetClusterMasters200ResponseMastersInner) HasGuestConsolePort() bool`

HasGuestConsolePort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


