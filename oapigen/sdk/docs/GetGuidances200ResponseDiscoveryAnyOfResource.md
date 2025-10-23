# GetGuidances200ResponseDiscoveryAnyOfResource

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
**ParentServer** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Owner** | Pointer to [**ListActivity200ResponseAllOfActivityInnerActivityInnerUser**](ListActivity200ResponseAllOfActivityInnerActivityInnerUser.md) |  | [optional] 
**Zone** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Plan** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**ComputeServerType** | Pointer to [**ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType**](ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType.md) |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ZoneId** | Pointer to **int64** |  | [optional] 
**SiteId** | Pointer to **int64** |  | [optional] 
**ResourcePoolId** | Pointer to **int64** |  | [optional] 
**FolderId** | Pointer to **int64** |  | [optional] 
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
**Stats** | Pointer to [**ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceStats**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceStats.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**StatusDate** | Pointer to **NullableTime** |  | [optional] 
**StatusPercent** | Pointer to **NullableString** |  | [optional] 
**StatusEta** | Pointer to **NullableString** |  | [optional] 
**PowerState** | Pointer to **string** |  | [optional] 
**AgentInstalled** | Pointer to **bool** |  | [optional] 
**LastAgentUpdate** | Pointer to **time.Time** |  | [optional] 
**AgentVersion** | Pointer to **string** |  | [optional] 
**MaxCores** | Pointer to **int64** |  | [optional] 
**CoresPerSocket** | Pointer to **int64** |  | [optional] 
**MaxMemory** | Pointer to **int64** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 
**MaxCpu** | Pointer to **NullableString** |  | [optional] 
**HourlyPrice** | Pointer to **float32** |  | [optional] 
**SourceImage** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**ServerOs** | Pointer to [**ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceServerOs**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceServerOs.md) |  | [optional] 
**Volumes** | Pointer to [**[]ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceVolumesInner**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceVolumesInner.md) |  | [optional] 
**Controllers** | Pointer to [**[]ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceControllersInner**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceControllersInner.md) |  | [optional] 
**Interfaces** | Pointer to [**[]GetGuidances200ResponseDiscoveryAnyOfResourceInterfacesInner**](GetGuidances200ResponseDiscoveryAnyOfResourceInterfacesInner.md) |  | [optional] 
**Labels** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Tags** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**TagCompliant** | Pointer to **NullableString** |  | [optional] 
**Containers** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewGetGuidances200ResponseDiscoveryAnyOfResource

`func NewGetGuidances200ResponseDiscoveryAnyOfResource() *GetGuidances200ResponseDiscoveryAnyOfResource`

NewGetGuidances200ResponseDiscoveryAnyOfResource instantiates a new GetGuidances200ResponseDiscoveryAnyOfResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGuidances200ResponseDiscoveryAnyOfResourceWithDefaults

`func NewGetGuidances200ResponseDiscoveryAnyOfResourceWithDefaults() *GetGuidances200ResponseDiscoveryAnyOfResource`

NewGetGuidances200ResponseDiscoveryAnyOfResourceWithDefaults instantiates a new GetGuidances200ResponseDiscoveryAnyOfResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetExternalId

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetInternalId

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalUniqueId

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetExternalUniqueId() string`

GetExternalUniqueId returns the ExternalUniqueId field if non-nil, zero value otherwise.

### GetExternalUniqueIdOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetExternalUniqueIdOk() (*string, bool)`

GetExternalUniqueIdOk returns a tuple with the ExternalUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUniqueId

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetExternalUniqueId(v string)`

SetExternalUniqueId sets ExternalUniqueId field to given value.

### HasExternalUniqueId

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasExternalUniqueId() bool`

HasExternalUniqueId returns a boolean if a field has been set.

### GetName

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalName

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetExternalName() string`

GetExternalName returns the ExternalName field if non-nil, zero value otherwise.

### GetExternalNameOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetExternalNameOk() (*string, bool)`

GetExternalNameOk returns a tuple with the ExternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalName

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetExternalName(v string)`

SetExternalName sets ExternalName field to given value.

### HasExternalName

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasExternalName() bool`

HasExternalName returns a boolean if a field has been set.

### GetHostname

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetParentServer

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetParentServer() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetParentServer returns the ParentServer field if non-nil, zero value otherwise.

### GetParentServerOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetParentServerOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetParentServerOk returns a tuple with the ParentServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentServer

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetParentServer(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetParentServer sets ParentServer field to given value.

### HasParentServer

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasParentServer() bool`

HasParentServer returns a boolean if a field has been set.

### GetAccountId

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccount

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetAccount() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetAccountOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetAccount(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetOwner() ListActivity200ResponseAllOfActivityInnerActivityInnerUser`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetOwnerOk() (*ListActivity200ResponseAllOfActivityInnerActivityInnerUser, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetOwner(v ListActivity200ResponseAllOfActivityInnerActivityInnerUser)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetZone

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetZone() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetZoneOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetZone(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetPlan

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetPlan() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetPlanOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetPlan(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetComputeServerType

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetComputeServerType() ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType`

GetComputeServerType returns the ComputeServerType field if non-nil, zero value otherwise.

### GetComputeServerTypeOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetComputeServerTypeOk() (*ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType, bool)`

GetComputeServerTypeOk returns a tuple with the ComputeServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServerType

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetComputeServerType(v ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType)`

SetComputeServerType sets ComputeServerType field to given value.

### HasComputeServerType

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasComputeServerType() bool`

HasComputeServerType returns a boolean if a field has been set.

### GetVisibility

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetZoneId

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetSiteId

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetResourcePoolId() int64`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetResourcePoolIdOk() (*int64, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetResourcePoolId(v int64)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetFolderId

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetFolderId() int64`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetFolderIdOk() (*int64, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetFolderId(v int64)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetSshHost

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetSshHost() string`

GetSshHost returns the SshHost field if non-nil, zero value otherwise.

### GetSshHostOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetSshHostOk() (*string, bool)`

GetSshHostOk returns a tuple with the SshHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHost

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetSshHost(v string)`

SetSshHost sets SshHost field to given value.

### HasSshHost

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasSshHost() bool`

HasSshHost returns a boolean if a field has been set.

### GetSshPort

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetExternalIp

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### GetInternalIp

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### GetVolumeId

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.

### SetVolumeIdNil

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetVolumeIdNil(b bool)`

 SetVolumeIdNil sets the value for VolumeId to be an explicit nil

### UnsetVolumeId
`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) UnsetVolumeId()`

UnsetVolumeId ensures that no value is present for VolumeId, not even an explicit nil
### GetPlatform

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPlatformVersion

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetPlatformVersion() string`

GetPlatformVersion returns the PlatformVersion field if non-nil, zero value otherwise.

### GetPlatformVersionOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetPlatformVersionOk() (*string, bool)`

GetPlatformVersionOk returns a tuple with the PlatformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformVersion

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetPlatformVersion(v string)`

SetPlatformVersion sets PlatformVersion field to given value.

### HasPlatformVersion

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasPlatformVersion() bool`

HasPlatformVersion returns a boolean if a field has been set.

### GetSshUsername

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetSshPasswordHash

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetSshPasswordHash() string`

GetSshPasswordHash returns the SshPasswordHash field if non-nil, zero value otherwise.

### GetSshPasswordHashOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetSshPasswordHashOk() (*string, bool)`

GetSshPasswordHashOk returns a tuple with the SshPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPasswordHash

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetSshPasswordHash(v string)`

SetSshPasswordHash sets SshPasswordHash field to given value.

### HasSshPasswordHash

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasSshPasswordHash() bool`

HasSshPasswordHash returns a boolean if a field has been set.

### GetOsDevice

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetOsDevice() string`

GetOsDevice returns the OsDevice field if non-nil, zero value otherwise.

### GetOsDeviceOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetOsDeviceOk() (*string, bool)`

GetOsDeviceOk returns a tuple with the OsDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDevice

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetOsDevice(v string)`

SetOsDevice sets OsDevice field to given value.

### HasOsDevice

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasOsDevice() bool`

HasOsDevice returns a boolean if a field has been set.

### GetOsType

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetDataDevice

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetDataDevice() string`

GetDataDevice returns the DataDevice field if non-nil, zero value otherwise.

### GetDataDeviceOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetDataDeviceOk() (*string, bool)`

GetDataDeviceOk returns a tuple with the DataDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDevice

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetDataDevice(v string)`

SetDataDevice sets DataDevice field to given value.

### HasDataDevice

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasDataDevice() bool`

HasDataDevice returns a boolean if a field has been set.

### GetLvmEnabled

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetLvmEnabled() bool`

GetLvmEnabled returns the LvmEnabled field if non-nil, zero value otherwise.

### GetLvmEnabledOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetLvmEnabledOk() (*bool, bool)`

GetLvmEnabledOk returns a tuple with the LvmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLvmEnabled

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetLvmEnabled(v bool)`

SetLvmEnabled sets LvmEnabled field to given value.

### HasLvmEnabled

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasLvmEnabled() bool`

HasLvmEnabled returns a boolean if a field has been set.

### GetApiKey

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetSoftwareRaid

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetSoftwareRaid() bool`

GetSoftwareRaid returns the SoftwareRaid field if non-nil, zero value otherwise.

### GetSoftwareRaidOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetSoftwareRaidOk() (*bool, bool)`

GetSoftwareRaidOk returns a tuple with the SoftwareRaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareRaid

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetSoftwareRaid(v bool)`

SetSoftwareRaid sets SoftwareRaid field to given value.

### HasSoftwareRaid

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasSoftwareRaid() bool`

HasSoftwareRaid returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetStats

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetStats() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetStatsOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetStats(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetStatus

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetErrorMessage

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetStatusDate

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### SetStatusDateNil

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetStatusDateNil(b bool)`

 SetStatusDateNil sets the value for StatusDate to be an explicit nil

### UnsetStatusDate
`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) UnsetStatusDate()`

UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
### GetStatusPercent

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetStatusPercent() string`

GetStatusPercent returns the StatusPercent field if non-nil, zero value otherwise.

### GetStatusPercentOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetStatusPercentOk() (*string, bool)`

GetStatusPercentOk returns a tuple with the StatusPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusPercent

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetStatusPercent(v string)`

SetStatusPercent sets StatusPercent field to given value.

### HasStatusPercent

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasStatusPercent() bool`

HasStatusPercent returns a boolean if a field has been set.

### SetStatusPercentNil

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetStatusPercentNil(b bool)`

 SetStatusPercentNil sets the value for StatusPercent to be an explicit nil

### UnsetStatusPercent
`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) UnsetStatusPercent()`

UnsetStatusPercent ensures that no value is present for StatusPercent, not even an explicit nil
### GetStatusEta

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetStatusEta() string`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetStatusEtaOk() (*string, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetStatusEta(v string)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### SetStatusEtaNil

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetStatusEtaNil(b bool)`

 SetStatusEtaNil sets the value for StatusEta to be an explicit nil

### UnsetStatusEta
`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) UnsetStatusEta()`

UnsetStatusEta ensures that no value is present for StatusEta, not even an explicit nil
### GetPowerState

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetPowerState() string`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetPowerStateOk() (*string, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetPowerState(v string)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetAgentInstalled

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetAgentInstalled() bool`

GetAgentInstalled returns the AgentInstalled field if non-nil, zero value otherwise.

### GetAgentInstalledOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetAgentInstalledOk() (*bool, bool)`

GetAgentInstalledOk returns a tuple with the AgentInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentInstalled

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetAgentInstalled(v bool)`

SetAgentInstalled sets AgentInstalled field to given value.

### HasAgentInstalled

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasAgentInstalled() bool`

HasAgentInstalled returns a boolean if a field has been set.

### GetLastAgentUpdate

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetLastAgentUpdate() time.Time`

GetLastAgentUpdate returns the LastAgentUpdate field if non-nil, zero value otherwise.

### GetLastAgentUpdateOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetLastAgentUpdateOk() (*time.Time, bool)`

GetLastAgentUpdateOk returns a tuple with the LastAgentUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAgentUpdate

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetLastAgentUpdate(v time.Time)`

SetLastAgentUpdate sets LastAgentUpdate field to given value.

### HasLastAgentUpdate

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasLastAgentUpdate() bool`

HasLastAgentUpdate returns a boolean if a field has been set.

### GetAgentVersion

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### GetMaxCores

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetMaxCores() int64`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetMaxCoresOk() (*int64, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetMaxCores(v int64)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetCoresPerSocket() int64`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetCoresPerSocketOk() (*int64, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetCoresPerSocket(v int64)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### GetMaxMemory

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxStorage

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMaxCpu

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetMaxCpu() string`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetMaxCpuOk() (*string, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetMaxCpu(v string)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### SetMaxCpuNil

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetMaxCpuNil(b bool)`

 SetMaxCpuNil sets the value for MaxCpu to be an explicit nil

### UnsetMaxCpu
`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) UnsetMaxCpu()`

UnsetMaxCpu ensures that no value is present for MaxCpu, not even an explicit nil
### GetHourlyPrice

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetHourlyPrice() float32`

GetHourlyPrice returns the HourlyPrice field if non-nil, zero value otherwise.

### GetHourlyPriceOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetHourlyPriceOk() (*float32, bool)`

GetHourlyPriceOk returns a tuple with the HourlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyPrice

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetHourlyPrice(v float32)`

SetHourlyPrice sets HourlyPrice field to given value.

### HasHourlyPrice

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasHourlyPrice() bool`

HasHourlyPrice returns a boolean if a field has been set.

### GetSourceImage

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetSourceImage() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetSourceImage returns the SourceImage field if non-nil, zero value otherwise.

### GetSourceImageOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetSourceImageOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetSourceImageOk returns a tuple with the SourceImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceImage

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetSourceImage(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetSourceImage sets SourceImage field to given value.

### HasSourceImage

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasSourceImage() bool`

HasSourceImage returns a boolean if a field has been set.

### GetServerOs

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetServerOs() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceServerOs`

GetServerOs returns the ServerOs field if non-nil, zero value otherwise.

### GetServerOsOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetServerOsOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceServerOs, bool)`

GetServerOsOk returns a tuple with the ServerOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerOs

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetServerOs(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceServerOs)`

SetServerOs sets ServerOs field to given value.

### HasServerOs

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasServerOs() bool`

HasServerOs returns a boolean if a field has been set.

### GetVolumes

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetVolumes() []ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetVolumesOk() (*[]ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetVolumes(v []ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetControllers

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetControllers() []ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceControllersInner`

GetControllers returns the Controllers field if non-nil, zero value otherwise.

### GetControllersOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetControllersOk() (*[]ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceControllersInner, bool)`

GetControllersOk returns a tuple with the Controllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllers

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetControllers(v []ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceControllersInner)`

SetControllers sets Controllers field to given value.

### HasControllers

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasControllers() bool`

HasControllers returns a boolean if a field has been set.

### GetInterfaces

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetInterfaces() []GetGuidances200ResponseDiscoveryAnyOfResourceInterfacesInner`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetInterfacesOk() (*[]GetGuidances200ResponseDiscoveryAnyOfResourceInterfacesInner, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetInterfaces(v []GetGuidances200ResponseDiscoveryAnyOfResourceInterfacesInner)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetLabels

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetLabels() []map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetLabelsOk() (*[]map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetLabels(v []map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTags

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetTags() []map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetTagsOk() (*[]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetTags(v []map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEnabled

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTagCompliant

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetTagCompliant() string`

GetTagCompliant returns the TagCompliant field if non-nil, zero value otherwise.

### GetTagCompliantOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetTagCompliantOk() (*string, bool)`

GetTagCompliantOk returns a tuple with the TagCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagCompliant

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetTagCompliant(v string)`

SetTagCompliant sets TagCompliant field to given value.

### HasTagCompliant

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasTagCompliant() bool`

HasTagCompliant returns a boolean if a field has been set.

### SetTagCompliantNil

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetTagCompliantNil(b bool)`

 SetTagCompliantNil sets the value for TagCompliant to be an explicit nil

### UnsetTagCompliant
`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) UnsetTagCompliant()`

UnsetTagCompliant ensures that no value is present for TagCompliant, not even an explicit nil
### GetContainers

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetContainers() []int64`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) GetContainersOk() (*[]int64, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) SetContainers(v []int64)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *GetGuidances200ResponseDiscoveryAnyOfResource) HasContainers() bool`

HasContainers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


