# ListClusterWorkers200ResponseAllOfWorkersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **NullableString** |  | [optional] 
**ExternalUniqueId** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ExternalName** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Owner** | Pointer to [**ListActivity200ResponseAllOfActivityInnerActivityInnerUser**](ListActivity200ResponseAllOfActivityInnerActivityInnerUser.md) |  | [optional] 
**Zone** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Plan** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**ComputeServerType** | Pointer to [**ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType**](ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType.md) |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ZoneId** | Pointer to **int64** |  | [optional] 
**SiteId** | Pointer to **int64** |  | [optional] 
**ResourcePoolId** | Pointer to **int64** |  | [optional] 
**FolderId** | Pointer to **NullableString** |  | [optional] 
**SshHost** | Pointer to **string** |  | [optional] 
**SshPort** | Pointer to **int64** |  | [optional] 
**ExternalIp** | Pointer to **string** |  | [optional] 
**InternalIp** | Pointer to **string** |  | [optional] 
**VolumeId** | Pointer to **NullableString** |  | [optional] 
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
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**StatusDate** | Pointer to **NullableString** |  | [optional] 
**StatusPercent** | Pointer to **NullableString** |  | [optional] 
**StatusEta** | Pointer to **NullableString** |  | [optional] 
**PowerState** | Pointer to **string** |  | [optional] 
**AgentInstalled** | Pointer to **bool** |  | [optional] 
**LastAgentUpdate** | Pointer to **time.Time** |  | [optional] 
**AgentVersion** | Pointer to **string** |  | [optional] 
**MaxCores** | Pointer to **int64** |  | [optional] 
**CoresPerSocket** | Pointer to **NullableString** |  | [optional] 
**MaxMemory** | Pointer to **int64** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 
**MaxCpu** | Pointer to **int64** |  | [optional] 
**HourlyPrice** | Pointer to **float32** |  | [optional] 
**SourceImage** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**ServerOs** | Pointer to **NullableString** |  | [optional] 
**Volumes** | Pointer to [**[]GetClusterMasters200ResponseMastersInnerVolumesInner**](GetClusterMasters200ResponseMastersInnerVolumesInner.md) |  | [optional] 
**Controllers** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Interfaces** | Pointer to [**[]ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner**](ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner.md) |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to [**[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**TagCompliant** | Pointer to **NullableString** |  | [optional] 
**Containers** | Pointer to **[]int64** |  | [optional] 
**GuestConsolePreferred** | Pointer to **bool** |  | [optional] 
**GuestConsoleType** | Pointer to **NullableString** |  | [optional] 
**GuestConsoleUsername** | Pointer to **NullableString** |  | [optional] 
**GuestConsolePassword** | Pointer to **NullableString** |  | [optional] 
**GuestConsolePasswordHash** | Pointer to **NullableString** |  | [optional] 
**GuestConsolePort** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewListClusterWorkers200ResponseAllOfWorkersInner

`func NewListClusterWorkers200ResponseAllOfWorkersInner() *ListClusterWorkers200ResponseAllOfWorkersInner`

NewListClusterWorkers200ResponseAllOfWorkersInner instantiates a new ListClusterWorkers200ResponseAllOfWorkersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterWorkers200ResponseAllOfWorkersInnerWithDefaults

`func NewListClusterWorkers200ResponseAllOfWorkersInnerWithDefaults() *ListClusterWorkers200ResponseAllOfWorkersInner`

NewListClusterWorkers200ResponseAllOfWorkersInnerWithDefaults instantiates a new ListClusterWorkers200ResponseAllOfWorkersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetExternalId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetInternalId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalUniqueId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetExternalUniqueId() string`

GetExternalUniqueId returns the ExternalUniqueId field if non-nil, zero value otherwise.

### GetExternalUniqueIdOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetExternalUniqueIdOk() (*string, bool)`

GetExternalUniqueIdOk returns a tuple with the ExternalUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUniqueId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetExternalUniqueId(v string)`

SetExternalUniqueId sets ExternalUniqueId field to given value.

### HasExternalUniqueId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasExternalUniqueId() bool`

HasExternalUniqueId returns a boolean if a field has been set.

### SetExternalUniqueIdNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetExternalUniqueIdNil(b bool)`

 SetExternalUniqueIdNil sets the value for ExternalUniqueId to be an explicit nil

### UnsetExternalUniqueId
`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) UnsetExternalUniqueId()`

UnsetExternalUniqueId ensures that no value is present for ExternalUniqueId, not even an explicit nil
### GetName

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalName

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetExternalName() string`

GetExternalName returns the ExternalName field if non-nil, zero value otherwise.

### GetExternalNameOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetExternalNameOk() (*string, bool)`

GetExternalNameOk returns a tuple with the ExternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalName

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetExternalName(v string)`

SetExternalName sets ExternalName field to given value.

### HasExternalName

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasExternalName() bool`

HasExternalName returns a boolean if a field has been set.

### GetHostname

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetAccountId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccount

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetAccount() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetAccountOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetAccount(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetOwner() ListActivity200ResponseAllOfActivityInnerActivityInnerUser`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetOwnerOk() (*ListActivity200ResponseAllOfActivityInnerActivityInnerUser, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetOwner(v ListActivity200ResponseAllOfActivityInnerActivityInnerUser)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetZone

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetZone() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetZoneOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetZone(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetPlan

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetPlan() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetPlanOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetPlan(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetComputeServerType

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetComputeServerType() ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType`

GetComputeServerType returns the ComputeServerType field if non-nil, zero value otherwise.

### GetComputeServerTypeOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetComputeServerTypeOk() (*ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType, bool)`

GetComputeServerTypeOk returns a tuple with the ComputeServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServerType

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetComputeServerType(v ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType)`

SetComputeServerType sets ComputeServerType field to given value.

### HasComputeServerType

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasComputeServerType() bool`

HasComputeServerType returns a boolean if a field has been set.

### GetVisibility

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetZoneId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetSiteId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetResourcePoolId() int64`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetResourcePoolIdOk() (*int64, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetResourcePoolId(v int64)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetFolderId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetSshHost

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetSshHost() string`

GetSshHost returns the SshHost field if non-nil, zero value otherwise.

### GetSshHostOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetSshHostOk() (*string, bool)`

GetSshHostOk returns a tuple with the SshHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHost

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetSshHost(v string)`

SetSshHost sets SshHost field to given value.

### HasSshHost

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasSshHost() bool`

HasSshHost returns a boolean if a field has been set.

### GetSshPort

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetExternalIp

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### GetInternalIp

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### GetVolumeId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.

### SetVolumeIdNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetVolumeIdNil(b bool)`

 SetVolumeIdNil sets the value for VolumeId to be an explicit nil

### UnsetVolumeId
`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) UnsetVolumeId()`

UnsetVolumeId ensures that no value is present for VolumeId, not even an explicit nil
### GetPlatform

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPlatformVersion

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetPlatformVersion() string`

GetPlatformVersion returns the PlatformVersion field if non-nil, zero value otherwise.

### GetPlatformVersionOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetPlatformVersionOk() (*string, bool)`

GetPlatformVersionOk returns a tuple with the PlatformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformVersion

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetPlatformVersion(v string)`

SetPlatformVersion sets PlatformVersion field to given value.

### HasPlatformVersion

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasPlatformVersion() bool`

HasPlatformVersion returns a boolean if a field has been set.

### GetSshUsername

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetSshPasswordHash

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetSshPasswordHash() string`

GetSshPasswordHash returns the SshPasswordHash field if non-nil, zero value otherwise.

### GetSshPasswordHashOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetSshPasswordHashOk() (*string, bool)`

GetSshPasswordHashOk returns a tuple with the SshPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPasswordHash

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetSshPasswordHash(v string)`

SetSshPasswordHash sets SshPasswordHash field to given value.

### HasSshPasswordHash

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasSshPasswordHash() bool`

HasSshPasswordHash returns a boolean if a field has been set.

### GetOsDevice

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetOsDevice() string`

GetOsDevice returns the OsDevice field if non-nil, zero value otherwise.

### GetOsDeviceOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetOsDeviceOk() (*string, bool)`

GetOsDeviceOk returns a tuple with the OsDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDevice

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetOsDevice(v string)`

SetOsDevice sets OsDevice field to given value.

### HasOsDevice

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasOsDevice() bool`

HasOsDevice returns a boolean if a field has been set.

### GetOsType

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetDataDevice

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetDataDevice() string`

GetDataDevice returns the DataDevice field if non-nil, zero value otherwise.

### GetDataDeviceOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetDataDeviceOk() (*string, bool)`

GetDataDeviceOk returns a tuple with the DataDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDevice

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetDataDevice(v string)`

SetDataDevice sets DataDevice field to given value.

### HasDataDevice

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasDataDevice() bool`

HasDataDevice returns a boolean if a field has been set.

### GetLvmEnabled

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetLvmEnabled() bool`

GetLvmEnabled returns the LvmEnabled field if non-nil, zero value otherwise.

### GetLvmEnabledOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetLvmEnabledOk() (*bool, bool)`

GetLvmEnabledOk returns a tuple with the LvmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLvmEnabled

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetLvmEnabled(v bool)`

SetLvmEnabled sets LvmEnabled field to given value.

### HasLvmEnabled

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasLvmEnabled() bool`

HasLvmEnabled returns a boolean if a field has been set.

### GetApiKey

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetSoftwareRaid

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetSoftwareRaid() bool`

GetSoftwareRaid returns the SoftwareRaid field if non-nil, zero value otherwise.

### GetSoftwareRaidOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetSoftwareRaidOk() (*bool, bool)`

GetSoftwareRaidOk returns a tuple with the SoftwareRaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareRaid

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetSoftwareRaid(v bool)`

SetSoftwareRaid sets SoftwareRaid field to given value.

### HasSoftwareRaid

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasSoftwareRaid() bool`

HasSoftwareRaid returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetStats

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetStats() GetClusterMasters200ResponseMastersInnerStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetStatsOk() (*GetClusterMasters200ResponseMastersInnerStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetStats(v GetClusterMasters200ResponseMastersInnerStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetStatus

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetErrorMessage

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetStatusDate

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetStatusDate() string`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetStatusDateOk() (*string, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetStatusDate(v string)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### SetStatusDateNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetStatusDateNil(b bool)`

 SetStatusDateNil sets the value for StatusDate to be an explicit nil

### UnsetStatusDate
`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) UnsetStatusDate()`

UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
### GetStatusPercent

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetStatusPercent() string`

GetStatusPercent returns the StatusPercent field if non-nil, zero value otherwise.

### GetStatusPercentOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetStatusPercentOk() (*string, bool)`

GetStatusPercentOk returns a tuple with the StatusPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusPercent

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetStatusPercent(v string)`

SetStatusPercent sets StatusPercent field to given value.

### HasStatusPercent

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasStatusPercent() bool`

HasStatusPercent returns a boolean if a field has been set.

### SetStatusPercentNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetStatusPercentNil(b bool)`

 SetStatusPercentNil sets the value for StatusPercent to be an explicit nil

### UnsetStatusPercent
`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) UnsetStatusPercent()`

UnsetStatusPercent ensures that no value is present for StatusPercent, not even an explicit nil
### GetStatusEta

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetStatusEta() string`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetStatusEtaOk() (*string, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetStatusEta(v string)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### SetStatusEtaNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetStatusEtaNil(b bool)`

 SetStatusEtaNil sets the value for StatusEta to be an explicit nil

### UnsetStatusEta
`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) UnsetStatusEta()`

UnsetStatusEta ensures that no value is present for StatusEta, not even an explicit nil
### GetPowerState

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetPowerState() string`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetPowerStateOk() (*string, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetPowerState(v string)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetAgentInstalled

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetAgentInstalled() bool`

GetAgentInstalled returns the AgentInstalled field if non-nil, zero value otherwise.

### GetAgentInstalledOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetAgentInstalledOk() (*bool, bool)`

GetAgentInstalledOk returns a tuple with the AgentInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentInstalled

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetAgentInstalled(v bool)`

SetAgentInstalled sets AgentInstalled field to given value.

### HasAgentInstalled

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasAgentInstalled() bool`

HasAgentInstalled returns a boolean if a field has been set.

### GetLastAgentUpdate

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetLastAgentUpdate() time.Time`

GetLastAgentUpdate returns the LastAgentUpdate field if non-nil, zero value otherwise.

### GetLastAgentUpdateOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetLastAgentUpdateOk() (*time.Time, bool)`

GetLastAgentUpdateOk returns a tuple with the LastAgentUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAgentUpdate

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetLastAgentUpdate(v time.Time)`

SetLastAgentUpdate sets LastAgentUpdate field to given value.

### HasLastAgentUpdate

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasLastAgentUpdate() bool`

HasLastAgentUpdate returns a boolean if a field has been set.

### GetAgentVersion

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### GetMaxCores

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetMaxCores() int64`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetMaxCoresOk() (*int64, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetMaxCores(v int64)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetCoresPerSocket() string`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetCoresPerSocketOk() (*string, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetCoresPerSocket(v string)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### SetCoresPerSocketNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetCoresPerSocketNil(b bool)`

 SetCoresPerSocketNil sets the value for CoresPerSocket to be an explicit nil

### UnsetCoresPerSocket
`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) UnsetCoresPerSocket()`

UnsetCoresPerSocket ensures that no value is present for CoresPerSocket, not even an explicit nil
### GetMaxMemory

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxStorage

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMaxCpu

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetMaxCpu() int64`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetMaxCpuOk() (*int64, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetMaxCpu(v int64)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetHourlyPrice

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetHourlyPrice() float32`

GetHourlyPrice returns the HourlyPrice field if non-nil, zero value otherwise.

### GetHourlyPriceOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetHourlyPriceOk() (*float32, bool)`

GetHourlyPriceOk returns a tuple with the HourlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyPrice

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetHourlyPrice(v float32)`

SetHourlyPrice sets HourlyPrice field to given value.

### HasHourlyPrice

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasHourlyPrice() bool`

HasHourlyPrice returns a boolean if a field has been set.

### GetSourceImage

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetSourceImage() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetSourceImage returns the SourceImage field if non-nil, zero value otherwise.

### GetSourceImageOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetSourceImageOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetSourceImageOk returns a tuple with the SourceImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceImage

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetSourceImage(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetSourceImage sets SourceImage field to given value.

### HasSourceImage

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasSourceImage() bool`

HasSourceImage returns a boolean if a field has been set.

### GetServerOs

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetServerOs() string`

GetServerOs returns the ServerOs field if non-nil, zero value otherwise.

### GetServerOsOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetServerOsOk() (*string, bool)`

GetServerOsOk returns a tuple with the ServerOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerOs

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetServerOs(v string)`

SetServerOs sets ServerOs field to given value.

### HasServerOs

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasServerOs() bool`

HasServerOs returns a boolean if a field has been set.

### SetServerOsNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetServerOsNil(b bool)`

 SetServerOsNil sets the value for ServerOs to be an explicit nil

### UnsetServerOs
`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) UnsetServerOs()`

UnsetServerOs ensures that no value is present for ServerOs, not even an explicit nil
### GetVolumes

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetVolumes() []GetClusterMasters200ResponseMastersInnerVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetVolumesOk() (*[]GetClusterMasters200ResponseMastersInnerVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetVolumes(v []GetClusterMasters200ResponseMastersInnerVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetControllers

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetControllers() []map[string]interface{}`

GetControllers returns the Controllers field if non-nil, zero value otherwise.

### GetControllersOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetControllersOk() (*[]map[string]interface{}, bool)`

GetControllersOk returns a tuple with the Controllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllers

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetControllers(v []map[string]interface{})`

SetControllers sets Controllers field to given value.

### HasControllers

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasControllers() bool`

HasControllers returns a boolean if a field has been set.

### GetInterfaces

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetInterfaces() []ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetInterfacesOk() (*[]ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetInterfaces(v []ListClusterWorkers200ResponseAllOfWorkersInnerInterfacesInner)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetLabels

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTags

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetTags() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetTagsOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetTags(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEnabled

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTagCompliant

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetTagCompliant() string`

GetTagCompliant returns the TagCompliant field if non-nil, zero value otherwise.

### GetTagCompliantOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetTagCompliantOk() (*string, bool)`

GetTagCompliantOk returns a tuple with the TagCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagCompliant

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetTagCompliant(v string)`

SetTagCompliant sets TagCompliant field to given value.

### HasTagCompliant

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasTagCompliant() bool`

HasTagCompliant returns a boolean if a field has been set.

### SetTagCompliantNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetTagCompliantNil(b bool)`

 SetTagCompliantNil sets the value for TagCompliant to be an explicit nil

### UnsetTagCompliant
`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) UnsetTagCompliant()`

UnsetTagCompliant ensures that no value is present for TagCompliant, not even an explicit nil
### GetContainers

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetContainers() []int64`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetContainersOk() (*[]int64, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetContainers(v []int64)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetGuestConsolePreferred

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetGuestConsolePreferred() bool`

GetGuestConsolePreferred returns the GuestConsolePreferred field if non-nil, zero value otherwise.

### GetGuestConsolePreferredOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetGuestConsolePreferredOk() (*bool, bool)`

GetGuestConsolePreferredOk returns a tuple with the GuestConsolePreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsolePreferred

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetGuestConsolePreferred(v bool)`

SetGuestConsolePreferred sets GuestConsolePreferred field to given value.

### HasGuestConsolePreferred

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasGuestConsolePreferred() bool`

HasGuestConsolePreferred returns a boolean if a field has been set.

### GetGuestConsoleType

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetGuestConsoleType() string`

GetGuestConsoleType returns the GuestConsoleType field if non-nil, zero value otherwise.

### GetGuestConsoleTypeOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetGuestConsoleTypeOk() (*string, bool)`

GetGuestConsoleTypeOk returns a tuple with the GuestConsoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleType

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetGuestConsoleType(v string)`

SetGuestConsoleType sets GuestConsoleType field to given value.

### HasGuestConsoleType

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasGuestConsoleType() bool`

HasGuestConsoleType returns a boolean if a field has been set.

### SetGuestConsoleTypeNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetGuestConsoleTypeNil(b bool)`

 SetGuestConsoleTypeNil sets the value for GuestConsoleType to be an explicit nil

### UnsetGuestConsoleType
`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) UnsetGuestConsoleType()`

UnsetGuestConsoleType ensures that no value is present for GuestConsoleType, not even an explicit nil
### GetGuestConsoleUsername

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetGuestConsoleUsername() string`

GetGuestConsoleUsername returns the GuestConsoleUsername field if non-nil, zero value otherwise.

### GetGuestConsoleUsernameOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetGuestConsoleUsernameOk() (*string, bool)`

GetGuestConsoleUsernameOk returns a tuple with the GuestConsoleUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleUsername

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetGuestConsoleUsername(v string)`

SetGuestConsoleUsername sets GuestConsoleUsername field to given value.

### HasGuestConsoleUsername

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasGuestConsoleUsername() bool`

HasGuestConsoleUsername returns a boolean if a field has been set.

### SetGuestConsoleUsernameNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetGuestConsoleUsernameNil(b bool)`

 SetGuestConsoleUsernameNil sets the value for GuestConsoleUsername to be an explicit nil

### UnsetGuestConsoleUsername
`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) UnsetGuestConsoleUsername()`

UnsetGuestConsoleUsername ensures that no value is present for GuestConsoleUsername, not even an explicit nil
### GetGuestConsolePassword

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetGuestConsolePassword() string`

GetGuestConsolePassword returns the GuestConsolePassword field if non-nil, zero value otherwise.

### GetGuestConsolePasswordOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetGuestConsolePasswordOk() (*string, bool)`

GetGuestConsolePasswordOk returns a tuple with the GuestConsolePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsolePassword

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetGuestConsolePassword(v string)`

SetGuestConsolePassword sets GuestConsolePassword field to given value.

### HasGuestConsolePassword

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasGuestConsolePassword() bool`

HasGuestConsolePassword returns a boolean if a field has been set.

### SetGuestConsolePasswordNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetGuestConsolePasswordNil(b bool)`

 SetGuestConsolePasswordNil sets the value for GuestConsolePassword to be an explicit nil

### UnsetGuestConsolePassword
`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) UnsetGuestConsolePassword()`

UnsetGuestConsolePassword ensures that no value is present for GuestConsolePassword, not even an explicit nil
### GetGuestConsolePasswordHash

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetGuestConsolePasswordHash() string`

GetGuestConsolePasswordHash returns the GuestConsolePasswordHash field if non-nil, zero value otherwise.

### GetGuestConsolePasswordHashOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetGuestConsolePasswordHashOk() (*string, bool)`

GetGuestConsolePasswordHashOk returns a tuple with the GuestConsolePasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsolePasswordHash

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetGuestConsolePasswordHash(v string)`

SetGuestConsolePasswordHash sets GuestConsolePasswordHash field to given value.

### HasGuestConsolePasswordHash

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasGuestConsolePasswordHash() bool`

HasGuestConsolePasswordHash returns a boolean if a field has been set.

### SetGuestConsolePasswordHashNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetGuestConsolePasswordHashNil(b bool)`

 SetGuestConsolePasswordHashNil sets the value for GuestConsolePasswordHash to be an explicit nil

### UnsetGuestConsolePasswordHash
`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) UnsetGuestConsolePasswordHash()`

UnsetGuestConsolePasswordHash ensures that no value is present for GuestConsolePasswordHash, not even an explicit nil
### GetGuestConsolePort

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetGuestConsolePort() string`

GetGuestConsolePort returns the GuestConsolePort field if non-nil, zero value otherwise.

### GetGuestConsolePortOk

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) GetGuestConsolePortOk() (*string, bool)`

GetGuestConsolePortOk returns a tuple with the GuestConsolePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsolePort

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetGuestConsolePort(v string)`

SetGuestConsolePort sets GuestConsolePort field to given value.

### HasGuestConsolePort

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) HasGuestConsolePort() bool`

HasGuestConsolePort returns a boolean if a field has been set.

### SetGuestConsolePortNil

`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) SetGuestConsolePortNil(b bool)`

 SetGuestConsolePortNil sets the value for GuestConsolePort to be an explicit nil

### UnsetGuestConsolePort
`func (o *ListClusterWorkers200ResponseAllOfWorkersInner) UnsetGuestConsolePort()`

UnsetGuestConsolePort ensures that no value is present for GuestConsolePort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


