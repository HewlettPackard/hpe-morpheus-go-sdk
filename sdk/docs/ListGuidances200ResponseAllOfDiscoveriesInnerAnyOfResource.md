# ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource

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
**Owner** | Pointer to [**ListActivity200ResponseAllOfActivityInnerActivityInnerUser**](ListActivity200ResponseAllOfActivityInnerActivityInnerUser.md) |  | [optional] 
**Zone** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Plan** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
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
**StatusMessage** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**StatusDate** | Pointer to **time.Time** |  | [optional] 
**StatusPercent** | Pointer to **string** |  | [optional] 
**StatusEta** | Pointer to **string** |  | [optional] 
**PowerState** | Pointer to **string** |  | [optional] 
**AgentInstalled** | Pointer to **bool** |  | [optional] 
**LastAgentUpdate** | Pointer to **time.Time** |  | [optional] 
**AgentVersion** | Pointer to **string** |  | [optional] 
**MaxCores** | Pointer to **int64** |  | [optional] 
**CoresPerSocket** | Pointer to **int64** |  | [optional] 
**MaxMemory** | Pointer to **int64** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 
**MaxCpu** | Pointer to **string** |  | [optional] 
**HourlyPrice** | Pointer to **float32** |  | [optional] 
**SourceImage** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**ServerOs** | Pointer to [**ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceServerOs**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceServerOs.md) |  | [optional] 
**Volumes** | Pointer to [**[]ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceVolumesInner**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceVolumesInner.md) |  | [optional] 
**Controllers** | Pointer to [**[]ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceControllersInner**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceControllersInner.md) |  | [optional] 
**Interfaces** | Pointer to [**[]ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner.md) |  | [optional] 
**Labels** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Tags** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**TagCompliant** | Pointer to **string** |  | [optional] 
**Containers** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource

`func NewListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource() *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource`

NewListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource instantiates a new ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceWithDefaults

`func NewListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceWithDefaults() *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource`

NewListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceWithDefaults instantiates a new ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetExternalId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetInternalId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalUniqueId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetExternalUniqueId() string`

GetExternalUniqueId returns the ExternalUniqueId field if non-nil, zero value otherwise.

### GetExternalUniqueIdOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetExternalUniqueIdOk() (*string, bool)`

GetExternalUniqueIdOk returns a tuple with the ExternalUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUniqueId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetExternalUniqueId(v string)`

SetExternalUniqueId sets ExternalUniqueId field to given value.

### HasExternalUniqueId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasExternalUniqueId() bool`

HasExternalUniqueId returns a boolean if a field has been set.

### GetName

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalName

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetExternalName() string`

GetExternalName returns the ExternalName field if non-nil, zero value otherwise.

### GetExternalNameOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetExternalNameOk() (*string, bool)`

GetExternalNameOk returns a tuple with the ExternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalName

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetExternalName(v string)`

SetExternalName sets ExternalName field to given value.

### HasExternalName

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasExternalName() bool`

HasExternalName returns a boolean if a field has been set.

### GetHostname

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetParentServer

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetParentServer() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetParentServer returns the ParentServer field if non-nil, zero value otherwise.

### GetParentServerOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetParentServerOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetParentServerOk returns a tuple with the ParentServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentServer

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetParentServer(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetParentServer sets ParentServer field to given value.

### HasParentServer

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasParentServer() bool`

HasParentServer returns a boolean if a field has been set.

### GetAccountId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccount

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetOwner() ListActivity200ResponseAllOfActivityInnerActivityInnerUser`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetOwnerOk() (*ListActivity200ResponseAllOfActivityInnerActivityInnerUser, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetOwner(v ListActivity200ResponseAllOfActivityInnerActivityInnerUser)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetZone

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetZone() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetZoneOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetZone(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetPlan

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetPlan() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetPlanOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetPlan(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetComputeServerType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetComputeServerType() ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType`

GetComputeServerType returns the ComputeServerType field if non-nil, zero value otherwise.

### GetComputeServerTypeOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetComputeServerTypeOk() (*ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType, bool)`

GetComputeServerTypeOk returns a tuple with the ComputeServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServerType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetComputeServerType(v ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType)`

SetComputeServerType sets ComputeServerType field to given value.

### HasComputeServerType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasComputeServerType() bool`

HasComputeServerType returns a boolean if a field has been set.

### GetVisibility

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetZoneId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetSiteId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetResourcePoolId() int64`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetResourcePoolIdOk() (*int64, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetResourcePoolId(v int64)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetFolderId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetFolderId() int64`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetFolderIdOk() (*int64, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetFolderId(v int64)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetSshHost

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetSshHost() string`

GetSshHost returns the SshHost field if non-nil, zero value otherwise.

### GetSshHostOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetSshHostOk() (*string, bool)`

GetSshHostOk returns a tuple with the SshHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHost

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetSshHost(v string)`

SetSshHost sets SshHost field to given value.

### HasSshHost

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasSshHost() bool`

HasSshHost returns a boolean if a field has been set.

### GetSshPort

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetExternalIp

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### GetInternalIp

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### GetVolumeId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.

### GetPlatform

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPlatformVersion

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetPlatformVersion() string`

GetPlatformVersion returns the PlatformVersion field if non-nil, zero value otherwise.

### GetPlatformVersionOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetPlatformVersionOk() (*string, bool)`

GetPlatformVersionOk returns a tuple with the PlatformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformVersion

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetPlatformVersion(v string)`

SetPlatformVersion sets PlatformVersion field to given value.

### HasPlatformVersion

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasPlatformVersion() bool`

HasPlatformVersion returns a boolean if a field has been set.

### GetSshUsername

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetSshPasswordHash

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetSshPasswordHash() string`

GetSshPasswordHash returns the SshPasswordHash field if non-nil, zero value otherwise.

### GetSshPasswordHashOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetSshPasswordHashOk() (*string, bool)`

GetSshPasswordHashOk returns a tuple with the SshPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPasswordHash

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetSshPasswordHash(v string)`

SetSshPasswordHash sets SshPasswordHash field to given value.

### HasSshPasswordHash

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasSshPasswordHash() bool`

HasSshPasswordHash returns a boolean if a field has been set.

### GetOsDevice

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetOsDevice() string`

GetOsDevice returns the OsDevice field if non-nil, zero value otherwise.

### GetOsDeviceOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetOsDeviceOk() (*string, bool)`

GetOsDeviceOk returns a tuple with the OsDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDevice

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetOsDevice(v string)`

SetOsDevice sets OsDevice field to given value.

### HasOsDevice

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasOsDevice() bool`

HasOsDevice returns a boolean if a field has been set.

### GetOsType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetDataDevice

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetDataDevice() string`

GetDataDevice returns the DataDevice field if non-nil, zero value otherwise.

### GetDataDeviceOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetDataDeviceOk() (*string, bool)`

GetDataDeviceOk returns a tuple with the DataDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDevice

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetDataDevice(v string)`

SetDataDevice sets DataDevice field to given value.

### HasDataDevice

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasDataDevice() bool`

HasDataDevice returns a boolean if a field has been set.

### GetLvmEnabled

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetLvmEnabled() bool`

GetLvmEnabled returns the LvmEnabled field if non-nil, zero value otherwise.

### GetLvmEnabledOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetLvmEnabledOk() (*bool, bool)`

GetLvmEnabledOk returns a tuple with the LvmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLvmEnabled

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetLvmEnabled(v bool)`

SetLvmEnabled sets LvmEnabled field to given value.

### HasLvmEnabled

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasLvmEnabled() bool`

HasLvmEnabled returns a boolean if a field has been set.

### GetApiKey

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetSoftwareRaid

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetSoftwareRaid() bool`

GetSoftwareRaid returns the SoftwareRaid field if non-nil, zero value otherwise.

### GetSoftwareRaidOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetSoftwareRaidOk() (*bool, bool)`

GetSoftwareRaidOk returns a tuple with the SoftwareRaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareRaid

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetSoftwareRaid(v bool)`

SetSoftwareRaid sets SoftwareRaid field to given value.

### HasSoftwareRaid

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasSoftwareRaid() bool`

HasSoftwareRaid returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetStats

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetStats() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetStatsOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetStats(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetStatus

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetStatusDate

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetStatusPercent

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetStatusPercent() string`

GetStatusPercent returns the StatusPercent field if non-nil, zero value otherwise.

### GetStatusPercentOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetStatusPercentOk() (*string, bool)`

GetStatusPercentOk returns a tuple with the StatusPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusPercent

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetStatusPercent(v string)`

SetStatusPercent sets StatusPercent field to given value.

### HasStatusPercent

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasStatusPercent() bool`

HasStatusPercent returns a boolean if a field has been set.

### GetStatusEta

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetStatusEta() string`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetStatusEtaOk() (*string, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetStatusEta(v string)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### GetPowerState

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetPowerState() string`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetPowerStateOk() (*string, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetPowerState(v string)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetAgentInstalled

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetAgentInstalled() bool`

GetAgentInstalled returns the AgentInstalled field if non-nil, zero value otherwise.

### GetAgentInstalledOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetAgentInstalledOk() (*bool, bool)`

GetAgentInstalledOk returns a tuple with the AgentInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentInstalled

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetAgentInstalled(v bool)`

SetAgentInstalled sets AgentInstalled field to given value.

### HasAgentInstalled

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasAgentInstalled() bool`

HasAgentInstalled returns a boolean if a field has been set.

### GetLastAgentUpdate

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetLastAgentUpdate() time.Time`

GetLastAgentUpdate returns the LastAgentUpdate field if non-nil, zero value otherwise.

### GetLastAgentUpdateOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetLastAgentUpdateOk() (*time.Time, bool)`

GetLastAgentUpdateOk returns a tuple with the LastAgentUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAgentUpdate

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetLastAgentUpdate(v time.Time)`

SetLastAgentUpdate sets LastAgentUpdate field to given value.

### HasLastAgentUpdate

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasLastAgentUpdate() bool`

HasLastAgentUpdate returns a boolean if a field has been set.

### GetAgentVersion

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### GetMaxCores

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetMaxCores() int64`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetMaxCoresOk() (*int64, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetMaxCores(v int64)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetCoresPerSocket() int64`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetCoresPerSocketOk() (*int64, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetCoresPerSocket(v int64)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### GetMaxMemory

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxStorage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMaxCpu

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetMaxCpu() string`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetMaxCpuOk() (*string, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetMaxCpu(v string)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetHourlyPrice

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetHourlyPrice() float32`

GetHourlyPrice returns the HourlyPrice field if non-nil, zero value otherwise.

### GetHourlyPriceOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetHourlyPriceOk() (*float32, bool)`

GetHourlyPriceOk returns a tuple with the HourlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyPrice

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetHourlyPrice(v float32)`

SetHourlyPrice sets HourlyPrice field to given value.

### HasHourlyPrice

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasHourlyPrice() bool`

HasHourlyPrice returns a boolean if a field has been set.

### GetSourceImage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetSourceImage() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetSourceImage returns the SourceImage field if non-nil, zero value otherwise.

### GetSourceImageOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetSourceImageOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetSourceImageOk returns a tuple with the SourceImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceImage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetSourceImage(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetSourceImage sets SourceImage field to given value.

### HasSourceImage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasSourceImage() bool`

HasSourceImage returns a boolean if a field has been set.

### GetServerOs

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetServerOs() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceServerOs`

GetServerOs returns the ServerOs field if non-nil, zero value otherwise.

### GetServerOsOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetServerOsOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceServerOs, bool)`

GetServerOsOk returns a tuple with the ServerOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerOs

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetServerOs(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceServerOs)`

SetServerOs sets ServerOs field to given value.

### HasServerOs

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasServerOs() bool`

HasServerOs returns a boolean if a field has been set.

### GetVolumes

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetVolumes() []ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetVolumesOk() (*[]ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetVolumes(v []ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetControllers

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetControllers() []ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceControllersInner`

GetControllers returns the Controllers field if non-nil, zero value otherwise.

### GetControllersOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetControllersOk() (*[]ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceControllersInner, bool)`

GetControllersOk returns a tuple with the Controllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllers

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetControllers(v []ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceControllersInner)`

SetControllers sets Controllers field to given value.

### HasControllers

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasControllers() bool`

HasControllers returns a boolean if a field has been set.

### GetInterfaces

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetInterfaces() []ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetInterfacesOk() (*[]ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetInterfaces(v []ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceInterfacesInner)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetLabels

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetLabels() []map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetLabelsOk() (*[]map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetLabels(v []map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTags

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetTags() []map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetTagsOk() (*[]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetTags(v []map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEnabled

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTagCompliant

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetTagCompliant() string`

GetTagCompliant returns the TagCompliant field if non-nil, zero value otherwise.

### GetTagCompliantOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetTagCompliantOk() (*string, bool)`

GetTagCompliantOk returns a tuple with the TagCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagCompliant

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetTagCompliant(v string)`

SetTagCompliant sets TagCompliant field to given value.

### HasTagCompliant

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasTagCompliant() bool`

HasTagCompliant returns a boolean if a field has been set.

### GetContainers

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetContainers() []int64`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) GetContainersOk() (*[]int64, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) SetContainers(v []int64)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource) HasContainers() bool`

HasContainers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


