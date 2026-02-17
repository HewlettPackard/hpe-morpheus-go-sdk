# GetHost200ResponseServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**InternalId** | Pointer to **NullableString** |  | [optional] 
**ExternalUniqueId** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ExternalName** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**ParentServer** | Pointer to [**AddBaremetalHost200ResponseServerParentServer**](AddBaremetalHost200ResponseServerParentServer.md) |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**AddBaremetalHost200ResponseServerAccount**](AddBaremetalHost200ResponseServerAccount.md) |  | [optional] 
**Owner** | Pointer to [**AddBaremetalHost200ResponseServerOwner**](AddBaremetalHost200ResponseServerOwner.md) |  | [optional] 
**Zone** | Pointer to [**AddBaremetalHost200ResponseServerZone**](AddBaremetalHost200ResponseServerZone.md) |  | [optional] 
**Plan** | Pointer to [**AddBaremetalHost200ResponseServerPlan**](AddBaremetalHost200ResponseServerPlan.md) |  | [optional] 
**ComputeServerType** | Pointer to [**AddBaremetalHost200ResponseServerComputeServerType**](AddBaremetalHost200ResponseServerComputeServerType.md) |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ZoneId** | Pointer to **int64** |  | [optional] 
**SiteId** | Pointer to **int64** |  | [optional] 
**ResourcePoolId** | Pointer to **NullableInt64** |  | [optional] 
**FolderId** | Pointer to **NullableInt64** |  | [optional] 
**SshHost** | Pointer to **NullableString** |  | [optional] 
**SshPort** | Pointer to **int64** |  | [optional] 
**ExternalIp** | Pointer to **NullableString** |  | [optional] 
**InternalIp** | Pointer to **NullableString** |  | [optional] 
**VolumeId** | Pointer to **NullableString** |  | [optional] 
**Platform** | Pointer to **NullableString** |  | [optional] 
**PlatformVersion** | Pointer to **NullableString** |  | [optional] 
**SshUsername** | Pointer to **NullableString** |  | [optional] 
**SshPassword** | Pointer to **NullableString** |  | [optional] 
**SshPasswordHash** | Pointer to **NullableString** |  | [optional] 
**SshKeyPair** | Pointer to [**AddBaremetalHost200ResponseServerSshKeyPair**](AddBaremetalHost200ResponseServerSshKeyPair.md) |  | [optional] 
**OsDevice** | Pointer to **string** |  | [optional] 
**OsType** | Pointer to **string** |  | [optional] 
**DataDevice** | Pointer to **string** |  | [optional] 
**LvmEnabled** | Pointer to **bool** |  | [optional] 
**ApiKey** | Pointer to **string** |  | [optional] 
**SoftwareRaid** | Pointer to **bool** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Stats** | Pointer to [**AddBaremetalHost200ResponseServerStats**](AddBaremetalHost200ResponseServerStats.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**StatusDate** | Pointer to **NullableTime** |  | [optional] 
**StatusPercent** | Pointer to **NullableString** |  | [optional] 
**StatusEta** | Pointer to **NullableString** |  | [optional] 
**PowerState** | Pointer to **string** |  | [optional] 
**AgentInstalled** | Pointer to **bool** |  | [optional] 
**LastAgentUpdate** | Pointer to **NullableString** |  | [optional] 
**AgentVersion** | Pointer to **NullableString** |  | [optional] 
**MaxCores** | Pointer to **int64** |  | [optional] 
**CoresPerSocket** | Pointer to **NullableInt64** |  | [optional] 
**MaxMemory** | Pointer to **int64** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 
**MaxCpu** | Pointer to **NullableInt64** |  | [optional] 
**MaxGpus** | Pointer to **NullableInt64** |  | [optional] 
**ManageInternalFirewall** | Pointer to **bool** |  | [optional] 
**EnableLogs** | Pointer to **bool** |  | [optional] 
**HourlyCost** | Pointer to **float32** |  | [optional] 
**HourlyPrice** | Pointer to **float32** |  | [optional] 
**SourceImage** | Pointer to [**AddBaremetalHost200ResponseServerSourceImage**](AddBaremetalHost200ResponseServerSourceImage.md) |  | [optional] 
**ServerOs** | Pointer to [**AddBaremetalHost200ResponseServerServerOs**](AddBaremetalHost200ResponseServerServerOs.md) |  | [optional] 
**Volumes** | Pointer to [**[]AddBaremetalHost200ResponseServerVolumesInner**](AddBaremetalHost200ResponseServerVolumesInner.md) |  | [optional] 
**Controllers** | Pointer to [**[]AddBaremetalHost200ResponseServerControllersInner**](AddBaremetalHost200ResponseServerControllersInner.md) |  | [optional] 
**Interfaces** | Pointer to [**[]AddBaremetalHost200ResponseServerInterfacesInner**](AddBaremetalHost200ResponseServerInterfacesInner.md) |  | [optional] 
**Labels** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Tags** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**TagCompliant** | Pointer to **NullableString** |  | [optional] 
**Containers** | Pointer to **[]int64** |  | [optional] 
**Config** | Pointer to [**AddBaremetalHost200ResponseServerConfig**](AddBaremetalHost200ResponseServerConfig.md) |  | [optional] 
**Instance** | Pointer to [**AddBaremetalHost200ResponseServerInstance**](AddBaremetalHost200ResponseServerInstance.md) |  | [optional] 
**GuestConsolePreferred** | Pointer to **bool** |  | [optional] 
**GuestConsoleType** | Pointer to **NullableString** |  | [optional] 
**GuestConsoleUsername** | Pointer to **NullableString** |  | [optional] 
**GuestConsolePassword** | Pointer to **NullableString** |  | [optional] 
**GuestConsolePasswordHash** | Pointer to **NullableString** |  | [optional] 
**GuestConsolePort** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetHost200ResponseServer

`func NewGetHost200ResponseServer() *GetHost200ResponseServer`

NewGetHost200ResponseServer instantiates a new GetHost200ResponseServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetHost200ResponseServerWithDefaults

`func NewGetHost200ResponseServerWithDefaults() *GetHost200ResponseServer`

NewGetHost200ResponseServerWithDefaults instantiates a new GetHost200ResponseServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetHost200ResponseServer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetHost200ResponseServer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetHost200ResponseServer) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetHost200ResponseServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *GetHost200ResponseServer) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetHost200ResponseServer) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetHost200ResponseServer) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetHost200ResponseServer) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetExternalId

`func (o *GetHost200ResponseServer) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetHost200ResponseServer) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetHost200ResponseServer) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetHost200ResponseServer) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *GetHost200ResponseServer) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *GetHost200ResponseServer) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetInternalId

`func (o *GetHost200ResponseServer) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *GetHost200ResponseServer) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *GetHost200ResponseServer) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *GetHost200ResponseServer) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *GetHost200ResponseServer) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *GetHost200ResponseServer) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalUniqueId

`func (o *GetHost200ResponseServer) GetExternalUniqueId() string`

GetExternalUniqueId returns the ExternalUniqueId field if non-nil, zero value otherwise.

### GetExternalUniqueIdOk

`func (o *GetHost200ResponseServer) GetExternalUniqueIdOk() (*string, bool)`

GetExternalUniqueIdOk returns a tuple with the ExternalUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUniqueId

`func (o *GetHost200ResponseServer) SetExternalUniqueId(v string)`

SetExternalUniqueId sets ExternalUniqueId field to given value.

### HasExternalUniqueId

`func (o *GetHost200ResponseServer) HasExternalUniqueId() bool`

HasExternalUniqueId returns a boolean if a field has been set.

### SetExternalUniqueIdNil

`func (o *GetHost200ResponseServer) SetExternalUniqueIdNil(b bool)`

 SetExternalUniqueIdNil sets the value for ExternalUniqueId to be an explicit nil

### UnsetExternalUniqueId
`func (o *GetHost200ResponseServer) UnsetExternalUniqueId()`

UnsetExternalUniqueId ensures that no value is present for ExternalUniqueId, not even an explicit nil
### GetName

`func (o *GetHost200ResponseServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetHost200ResponseServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetHost200ResponseServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetHost200ResponseServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalName

`func (o *GetHost200ResponseServer) GetExternalName() string`

GetExternalName returns the ExternalName field if non-nil, zero value otherwise.

### GetExternalNameOk

`func (o *GetHost200ResponseServer) GetExternalNameOk() (*string, bool)`

GetExternalNameOk returns a tuple with the ExternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalName

`func (o *GetHost200ResponseServer) SetExternalName(v string)`

SetExternalName sets ExternalName field to given value.

### HasExternalName

`func (o *GetHost200ResponseServer) HasExternalName() bool`

HasExternalName returns a boolean if a field has been set.

### GetHostname

`func (o *GetHost200ResponseServer) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *GetHost200ResponseServer) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *GetHost200ResponseServer) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *GetHost200ResponseServer) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetParentServer

`func (o *GetHost200ResponseServer) GetParentServer() AddBaremetalHost200ResponseServerParentServer`

GetParentServer returns the ParentServer field if non-nil, zero value otherwise.

### GetParentServerOk

`func (o *GetHost200ResponseServer) GetParentServerOk() (*AddBaremetalHost200ResponseServerParentServer, bool)`

GetParentServerOk returns a tuple with the ParentServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentServer

`func (o *GetHost200ResponseServer) SetParentServer(v AddBaremetalHost200ResponseServerParentServer)`

SetParentServer sets ParentServer field to given value.

### HasParentServer

`func (o *GetHost200ResponseServer) HasParentServer() bool`

HasParentServer returns a boolean if a field has been set.

### GetAccountId

`func (o *GetHost200ResponseServer) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetHost200ResponseServer) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetHost200ResponseServer) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetHost200ResponseServer) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccount

`func (o *GetHost200ResponseServer) GetAccount() AddBaremetalHost200ResponseServerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetHost200ResponseServer) GetAccountOk() (*AddBaremetalHost200ResponseServerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetHost200ResponseServer) SetAccount(v AddBaremetalHost200ResponseServerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetHost200ResponseServer) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *GetHost200ResponseServer) GetOwner() AddBaremetalHost200ResponseServerOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetHost200ResponseServer) GetOwnerOk() (*AddBaremetalHost200ResponseServerOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetHost200ResponseServer) SetOwner(v AddBaremetalHost200ResponseServerOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetHost200ResponseServer) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetZone

`func (o *GetHost200ResponseServer) GetZone() AddBaremetalHost200ResponseServerZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetHost200ResponseServer) GetZoneOk() (*AddBaremetalHost200ResponseServerZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetHost200ResponseServer) SetZone(v AddBaremetalHost200ResponseServerZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetHost200ResponseServer) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetPlan

`func (o *GetHost200ResponseServer) GetPlan() AddBaremetalHost200ResponseServerPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *GetHost200ResponseServer) GetPlanOk() (*AddBaremetalHost200ResponseServerPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *GetHost200ResponseServer) SetPlan(v AddBaremetalHost200ResponseServerPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *GetHost200ResponseServer) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetComputeServerType

`func (o *GetHost200ResponseServer) GetComputeServerType() AddBaremetalHost200ResponseServerComputeServerType`

GetComputeServerType returns the ComputeServerType field if non-nil, zero value otherwise.

### GetComputeServerTypeOk

`func (o *GetHost200ResponseServer) GetComputeServerTypeOk() (*AddBaremetalHost200ResponseServerComputeServerType, bool)`

GetComputeServerTypeOk returns a tuple with the ComputeServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServerType

`func (o *GetHost200ResponseServer) SetComputeServerType(v AddBaremetalHost200ResponseServerComputeServerType)`

SetComputeServerType sets ComputeServerType field to given value.

### HasComputeServerType

`func (o *GetHost200ResponseServer) HasComputeServerType() bool`

HasComputeServerType returns a boolean if a field has been set.

### GetVisibility

`func (o *GetHost200ResponseServer) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetHost200ResponseServer) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetHost200ResponseServer) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetHost200ResponseServer) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *GetHost200ResponseServer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetHost200ResponseServer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetHost200ResponseServer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetHost200ResponseServer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetHost200ResponseServer) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetHost200ResponseServer) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetZoneId

`func (o *GetHost200ResponseServer) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *GetHost200ResponseServer) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *GetHost200ResponseServer) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *GetHost200ResponseServer) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetSiteId

`func (o *GetHost200ResponseServer) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *GetHost200ResponseServer) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *GetHost200ResponseServer) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *GetHost200ResponseServer) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *GetHost200ResponseServer) GetResourcePoolId() int64`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *GetHost200ResponseServer) GetResourcePoolIdOk() (*int64, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *GetHost200ResponseServer) SetResourcePoolId(v int64)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *GetHost200ResponseServer) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### SetResourcePoolIdNil

`func (o *GetHost200ResponseServer) SetResourcePoolIdNil(b bool)`

 SetResourcePoolIdNil sets the value for ResourcePoolId to be an explicit nil

### UnsetResourcePoolId
`func (o *GetHost200ResponseServer) UnsetResourcePoolId()`

UnsetResourcePoolId ensures that no value is present for ResourcePoolId, not even an explicit nil
### GetFolderId

`func (o *GetHost200ResponseServer) GetFolderId() int64`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *GetHost200ResponseServer) GetFolderIdOk() (*int64, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *GetHost200ResponseServer) SetFolderId(v int64)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *GetHost200ResponseServer) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *GetHost200ResponseServer) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *GetHost200ResponseServer) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetSshHost

`func (o *GetHost200ResponseServer) GetSshHost() string`

GetSshHost returns the SshHost field if non-nil, zero value otherwise.

### GetSshHostOk

`func (o *GetHost200ResponseServer) GetSshHostOk() (*string, bool)`

GetSshHostOk returns a tuple with the SshHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHost

`func (o *GetHost200ResponseServer) SetSshHost(v string)`

SetSshHost sets SshHost field to given value.

### HasSshHost

`func (o *GetHost200ResponseServer) HasSshHost() bool`

HasSshHost returns a boolean if a field has been set.

### SetSshHostNil

`func (o *GetHost200ResponseServer) SetSshHostNil(b bool)`

 SetSshHostNil sets the value for SshHost to be an explicit nil

### UnsetSshHost
`func (o *GetHost200ResponseServer) UnsetSshHost()`

UnsetSshHost ensures that no value is present for SshHost, not even an explicit nil
### GetSshPort

`func (o *GetHost200ResponseServer) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *GetHost200ResponseServer) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *GetHost200ResponseServer) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *GetHost200ResponseServer) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetExternalIp

`func (o *GetHost200ResponseServer) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *GetHost200ResponseServer) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *GetHost200ResponseServer) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *GetHost200ResponseServer) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### SetExternalIpNil

`func (o *GetHost200ResponseServer) SetExternalIpNil(b bool)`

 SetExternalIpNil sets the value for ExternalIp to be an explicit nil

### UnsetExternalIp
`func (o *GetHost200ResponseServer) UnsetExternalIp()`

UnsetExternalIp ensures that no value is present for ExternalIp, not even an explicit nil
### GetInternalIp

`func (o *GetHost200ResponseServer) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *GetHost200ResponseServer) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *GetHost200ResponseServer) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *GetHost200ResponseServer) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### SetInternalIpNil

`func (o *GetHost200ResponseServer) SetInternalIpNil(b bool)`

 SetInternalIpNil sets the value for InternalIp to be an explicit nil

### UnsetInternalIp
`func (o *GetHost200ResponseServer) UnsetInternalIp()`

UnsetInternalIp ensures that no value is present for InternalIp, not even an explicit nil
### GetVolumeId

`func (o *GetHost200ResponseServer) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *GetHost200ResponseServer) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *GetHost200ResponseServer) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *GetHost200ResponseServer) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.

### SetVolumeIdNil

`func (o *GetHost200ResponseServer) SetVolumeIdNil(b bool)`

 SetVolumeIdNil sets the value for VolumeId to be an explicit nil

### UnsetVolumeId
`func (o *GetHost200ResponseServer) UnsetVolumeId()`

UnsetVolumeId ensures that no value is present for VolumeId, not even an explicit nil
### GetPlatform

`func (o *GetHost200ResponseServer) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetHost200ResponseServer) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetHost200ResponseServer) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetHost200ResponseServer) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *GetHost200ResponseServer) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *GetHost200ResponseServer) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetPlatformVersion

`func (o *GetHost200ResponseServer) GetPlatformVersion() string`

GetPlatformVersion returns the PlatformVersion field if non-nil, zero value otherwise.

### GetPlatformVersionOk

`func (o *GetHost200ResponseServer) GetPlatformVersionOk() (*string, bool)`

GetPlatformVersionOk returns a tuple with the PlatformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformVersion

`func (o *GetHost200ResponseServer) SetPlatformVersion(v string)`

SetPlatformVersion sets PlatformVersion field to given value.

### HasPlatformVersion

`func (o *GetHost200ResponseServer) HasPlatformVersion() bool`

HasPlatformVersion returns a boolean if a field has been set.

### SetPlatformVersionNil

`func (o *GetHost200ResponseServer) SetPlatformVersionNil(b bool)`

 SetPlatformVersionNil sets the value for PlatformVersion to be an explicit nil

### UnsetPlatformVersion
`func (o *GetHost200ResponseServer) UnsetPlatformVersion()`

UnsetPlatformVersion ensures that no value is present for PlatformVersion, not even an explicit nil
### GetSshUsername

`func (o *GetHost200ResponseServer) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *GetHost200ResponseServer) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *GetHost200ResponseServer) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *GetHost200ResponseServer) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### SetSshUsernameNil

`func (o *GetHost200ResponseServer) SetSshUsernameNil(b bool)`

 SetSshUsernameNil sets the value for SshUsername to be an explicit nil

### UnsetSshUsername
`func (o *GetHost200ResponseServer) UnsetSshUsername()`

UnsetSshUsername ensures that no value is present for SshUsername, not even an explicit nil
### GetSshPassword

`func (o *GetHost200ResponseServer) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *GetHost200ResponseServer) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *GetHost200ResponseServer) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *GetHost200ResponseServer) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### SetSshPasswordNil

`func (o *GetHost200ResponseServer) SetSshPasswordNil(b bool)`

 SetSshPasswordNil sets the value for SshPassword to be an explicit nil

### UnsetSshPassword
`func (o *GetHost200ResponseServer) UnsetSshPassword()`

UnsetSshPassword ensures that no value is present for SshPassword, not even an explicit nil
### GetSshPasswordHash

`func (o *GetHost200ResponseServer) GetSshPasswordHash() string`

GetSshPasswordHash returns the SshPasswordHash field if non-nil, zero value otherwise.

### GetSshPasswordHashOk

`func (o *GetHost200ResponseServer) GetSshPasswordHashOk() (*string, bool)`

GetSshPasswordHashOk returns a tuple with the SshPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPasswordHash

`func (o *GetHost200ResponseServer) SetSshPasswordHash(v string)`

SetSshPasswordHash sets SshPasswordHash field to given value.

### HasSshPasswordHash

`func (o *GetHost200ResponseServer) HasSshPasswordHash() bool`

HasSshPasswordHash returns a boolean if a field has been set.

### SetSshPasswordHashNil

`func (o *GetHost200ResponseServer) SetSshPasswordHashNil(b bool)`

 SetSshPasswordHashNil sets the value for SshPasswordHash to be an explicit nil

### UnsetSshPasswordHash
`func (o *GetHost200ResponseServer) UnsetSshPasswordHash()`

UnsetSshPasswordHash ensures that no value is present for SshPasswordHash, not even an explicit nil
### GetSshKeyPair

`func (o *GetHost200ResponseServer) GetSshKeyPair() AddBaremetalHost200ResponseServerSshKeyPair`

GetSshKeyPair returns the SshKeyPair field if non-nil, zero value otherwise.

### GetSshKeyPairOk

`func (o *GetHost200ResponseServer) GetSshKeyPairOk() (*AddBaremetalHost200ResponseServerSshKeyPair, bool)`

GetSshKeyPairOk returns a tuple with the SshKeyPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyPair

`func (o *GetHost200ResponseServer) SetSshKeyPair(v AddBaremetalHost200ResponseServerSshKeyPair)`

SetSshKeyPair sets SshKeyPair field to given value.

### HasSshKeyPair

`func (o *GetHost200ResponseServer) HasSshKeyPair() bool`

HasSshKeyPair returns a boolean if a field has been set.

### GetOsDevice

`func (o *GetHost200ResponseServer) GetOsDevice() string`

GetOsDevice returns the OsDevice field if non-nil, zero value otherwise.

### GetOsDeviceOk

`func (o *GetHost200ResponseServer) GetOsDeviceOk() (*string, bool)`

GetOsDeviceOk returns a tuple with the OsDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDevice

`func (o *GetHost200ResponseServer) SetOsDevice(v string)`

SetOsDevice sets OsDevice field to given value.

### HasOsDevice

`func (o *GetHost200ResponseServer) HasOsDevice() bool`

HasOsDevice returns a boolean if a field has been set.

### GetOsType

`func (o *GetHost200ResponseServer) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *GetHost200ResponseServer) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *GetHost200ResponseServer) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *GetHost200ResponseServer) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetDataDevice

`func (o *GetHost200ResponseServer) GetDataDevice() string`

GetDataDevice returns the DataDevice field if non-nil, zero value otherwise.

### GetDataDeviceOk

`func (o *GetHost200ResponseServer) GetDataDeviceOk() (*string, bool)`

GetDataDeviceOk returns a tuple with the DataDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDevice

`func (o *GetHost200ResponseServer) SetDataDevice(v string)`

SetDataDevice sets DataDevice field to given value.

### HasDataDevice

`func (o *GetHost200ResponseServer) HasDataDevice() bool`

HasDataDevice returns a boolean if a field has been set.

### GetLvmEnabled

`func (o *GetHost200ResponseServer) GetLvmEnabled() bool`

GetLvmEnabled returns the LvmEnabled field if non-nil, zero value otherwise.

### GetLvmEnabledOk

`func (o *GetHost200ResponseServer) GetLvmEnabledOk() (*bool, bool)`

GetLvmEnabledOk returns a tuple with the LvmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLvmEnabled

`func (o *GetHost200ResponseServer) SetLvmEnabled(v bool)`

SetLvmEnabled sets LvmEnabled field to given value.

### HasLvmEnabled

`func (o *GetHost200ResponseServer) HasLvmEnabled() bool`

HasLvmEnabled returns a boolean if a field has been set.

### GetApiKey

`func (o *GetHost200ResponseServer) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *GetHost200ResponseServer) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *GetHost200ResponseServer) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *GetHost200ResponseServer) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetSoftwareRaid

`func (o *GetHost200ResponseServer) GetSoftwareRaid() bool`

GetSoftwareRaid returns the SoftwareRaid field if non-nil, zero value otherwise.

### GetSoftwareRaidOk

`func (o *GetHost200ResponseServer) GetSoftwareRaidOk() (*bool, bool)`

GetSoftwareRaidOk returns a tuple with the SoftwareRaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareRaid

`func (o *GetHost200ResponseServer) SetSoftwareRaid(v bool)`

SetSoftwareRaid sets SoftwareRaid field to given value.

### HasSoftwareRaid

`func (o *GetHost200ResponseServer) HasSoftwareRaid() bool`

HasSoftwareRaid returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetHost200ResponseServer) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetHost200ResponseServer) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetHost200ResponseServer) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetHost200ResponseServer) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetHost200ResponseServer) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetHost200ResponseServer) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetHost200ResponseServer) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetHost200ResponseServer) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetStats

`func (o *GetHost200ResponseServer) GetStats() AddBaremetalHost200ResponseServerStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *GetHost200ResponseServer) GetStatsOk() (*AddBaremetalHost200ResponseServerStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *GetHost200ResponseServer) SetStats(v AddBaremetalHost200ResponseServerStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *GetHost200ResponseServer) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetStatus

`func (o *GetHost200ResponseServer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetHost200ResponseServer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetHost200ResponseServer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetHost200ResponseServer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *GetHost200ResponseServer) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetHost200ResponseServer) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetHost200ResponseServer) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GetHost200ResponseServer) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *GetHost200ResponseServer) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *GetHost200ResponseServer) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetErrorMessage

`func (o *GetHost200ResponseServer) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *GetHost200ResponseServer) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *GetHost200ResponseServer) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *GetHost200ResponseServer) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *GetHost200ResponseServer) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *GetHost200ResponseServer) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetStatusDate

`func (o *GetHost200ResponseServer) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *GetHost200ResponseServer) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *GetHost200ResponseServer) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *GetHost200ResponseServer) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### SetStatusDateNil

`func (o *GetHost200ResponseServer) SetStatusDateNil(b bool)`

 SetStatusDateNil sets the value for StatusDate to be an explicit nil

### UnsetStatusDate
`func (o *GetHost200ResponseServer) UnsetStatusDate()`

UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
### GetStatusPercent

`func (o *GetHost200ResponseServer) GetStatusPercent() string`

GetStatusPercent returns the StatusPercent field if non-nil, zero value otherwise.

### GetStatusPercentOk

`func (o *GetHost200ResponseServer) GetStatusPercentOk() (*string, bool)`

GetStatusPercentOk returns a tuple with the StatusPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusPercent

`func (o *GetHost200ResponseServer) SetStatusPercent(v string)`

SetStatusPercent sets StatusPercent field to given value.

### HasStatusPercent

`func (o *GetHost200ResponseServer) HasStatusPercent() bool`

HasStatusPercent returns a boolean if a field has been set.

### SetStatusPercentNil

`func (o *GetHost200ResponseServer) SetStatusPercentNil(b bool)`

 SetStatusPercentNil sets the value for StatusPercent to be an explicit nil

### UnsetStatusPercent
`func (o *GetHost200ResponseServer) UnsetStatusPercent()`

UnsetStatusPercent ensures that no value is present for StatusPercent, not even an explicit nil
### GetStatusEta

`func (o *GetHost200ResponseServer) GetStatusEta() string`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *GetHost200ResponseServer) GetStatusEtaOk() (*string, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *GetHost200ResponseServer) SetStatusEta(v string)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *GetHost200ResponseServer) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### SetStatusEtaNil

`func (o *GetHost200ResponseServer) SetStatusEtaNil(b bool)`

 SetStatusEtaNil sets the value for StatusEta to be an explicit nil

### UnsetStatusEta
`func (o *GetHost200ResponseServer) UnsetStatusEta()`

UnsetStatusEta ensures that no value is present for StatusEta, not even an explicit nil
### GetPowerState

`func (o *GetHost200ResponseServer) GetPowerState() string`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *GetHost200ResponseServer) GetPowerStateOk() (*string, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *GetHost200ResponseServer) SetPowerState(v string)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *GetHost200ResponseServer) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetAgentInstalled

`func (o *GetHost200ResponseServer) GetAgentInstalled() bool`

GetAgentInstalled returns the AgentInstalled field if non-nil, zero value otherwise.

### GetAgentInstalledOk

`func (o *GetHost200ResponseServer) GetAgentInstalledOk() (*bool, bool)`

GetAgentInstalledOk returns a tuple with the AgentInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentInstalled

`func (o *GetHost200ResponseServer) SetAgentInstalled(v bool)`

SetAgentInstalled sets AgentInstalled field to given value.

### HasAgentInstalled

`func (o *GetHost200ResponseServer) HasAgentInstalled() bool`

HasAgentInstalled returns a boolean if a field has been set.

### GetLastAgentUpdate

`func (o *GetHost200ResponseServer) GetLastAgentUpdate() string`

GetLastAgentUpdate returns the LastAgentUpdate field if non-nil, zero value otherwise.

### GetLastAgentUpdateOk

`func (o *GetHost200ResponseServer) GetLastAgentUpdateOk() (*string, bool)`

GetLastAgentUpdateOk returns a tuple with the LastAgentUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAgentUpdate

`func (o *GetHost200ResponseServer) SetLastAgentUpdate(v string)`

SetLastAgentUpdate sets LastAgentUpdate field to given value.

### HasLastAgentUpdate

`func (o *GetHost200ResponseServer) HasLastAgentUpdate() bool`

HasLastAgentUpdate returns a boolean if a field has been set.

### SetLastAgentUpdateNil

`func (o *GetHost200ResponseServer) SetLastAgentUpdateNil(b bool)`

 SetLastAgentUpdateNil sets the value for LastAgentUpdate to be an explicit nil

### UnsetLastAgentUpdate
`func (o *GetHost200ResponseServer) UnsetLastAgentUpdate()`

UnsetLastAgentUpdate ensures that no value is present for LastAgentUpdate, not even an explicit nil
### GetAgentVersion

`func (o *GetHost200ResponseServer) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *GetHost200ResponseServer) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *GetHost200ResponseServer) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *GetHost200ResponseServer) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### SetAgentVersionNil

`func (o *GetHost200ResponseServer) SetAgentVersionNil(b bool)`

 SetAgentVersionNil sets the value for AgentVersion to be an explicit nil

### UnsetAgentVersion
`func (o *GetHost200ResponseServer) UnsetAgentVersion()`

UnsetAgentVersion ensures that no value is present for AgentVersion, not even an explicit nil
### GetMaxCores

`func (o *GetHost200ResponseServer) GetMaxCores() int64`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *GetHost200ResponseServer) GetMaxCoresOk() (*int64, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *GetHost200ResponseServer) SetMaxCores(v int64)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *GetHost200ResponseServer) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *GetHost200ResponseServer) GetCoresPerSocket() int64`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *GetHost200ResponseServer) GetCoresPerSocketOk() (*int64, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *GetHost200ResponseServer) SetCoresPerSocket(v int64)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *GetHost200ResponseServer) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### SetCoresPerSocketNil

`func (o *GetHost200ResponseServer) SetCoresPerSocketNil(b bool)`

 SetCoresPerSocketNil sets the value for CoresPerSocket to be an explicit nil

### UnsetCoresPerSocket
`func (o *GetHost200ResponseServer) UnsetCoresPerSocket()`

UnsetCoresPerSocket ensures that no value is present for CoresPerSocket, not even an explicit nil
### GetMaxMemory

`func (o *GetHost200ResponseServer) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *GetHost200ResponseServer) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *GetHost200ResponseServer) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *GetHost200ResponseServer) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxStorage

`func (o *GetHost200ResponseServer) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *GetHost200ResponseServer) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *GetHost200ResponseServer) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *GetHost200ResponseServer) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMaxCpu

`func (o *GetHost200ResponseServer) GetMaxCpu() int64`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *GetHost200ResponseServer) GetMaxCpuOk() (*int64, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *GetHost200ResponseServer) SetMaxCpu(v int64)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *GetHost200ResponseServer) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### SetMaxCpuNil

`func (o *GetHost200ResponseServer) SetMaxCpuNil(b bool)`

 SetMaxCpuNil sets the value for MaxCpu to be an explicit nil

### UnsetMaxCpu
`func (o *GetHost200ResponseServer) UnsetMaxCpu()`

UnsetMaxCpu ensures that no value is present for MaxCpu, not even an explicit nil
### GetMaxGpus

`func (o *GetHost200ResponseServer) GetMaxGpus() int64`

GetMaxGpus returns the MaxGpus field if non-nil, zero value otherwise.

### GetMaxGpusOk

`func (o *GetHost200ResponseServer) GetMaxGpusOk() (*int64, bool)`

GetMaxGpusOk returns a tuple with the MaxGpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGpus

`func (o *GetHost200ResponseServer) SetMaxGpus(v int64)`

SetMaxGpus sets MaxGpus field to given value.

### HasMaxGpus

`func (o *GetHost200ResponseServer) HasMaxGpus() bool`

HasMaxGpus returns a boolean if a field has been set.

### SetMaxGpusNil

`func (o *GetHost200ResponseServer) SetMaxGpusNil(b bool)`

 SetMaxGpusNil sets the value for MaxGpus to be an explicit nil

### UnsetMaxGpus
`func (o *GetHost200ResponseServer) UnsetMaxGpus()`

UnsetMaxGpus ensures that no value is present for MaxGpus, not even an explicit nil
### GetManageInternalFirewall

`func (o *GetHost200ResponseServer) GetManageInternalFirewall() bool`

GetManageInternalFirewall returns the ManageInternalFirewall field if non-nil, zero value otherwise.

### GetManageInternalFirewallOk

`func (o *GetHost200ResponseServer) GetManageInternalFirewallOk() (*bool, bool)`

GetManageInternalFirewallOk returns a tuple with the ManageInternalFirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageInternalFirewall

`func (o *GetHost200ResponseServer) SetManageInternalFirewall(v bool)`

SetManageInternalFirewall sets ManageInternalFirewall field to given value.

### HasManageInternalFirewall

`func (o *GetHost200ResponseServer) HasManageInternalFirewall() bool`

HasManageInternalFirewall returns a boolean if a field has been set.

### GetEnableLogs

`func (o *GetHost200ResponseServer) GetEnableLogs() bool`

GetEnableLogs returns the EnableLogs field if non-nil, zero value otherwise.

### GetEnableLogsOk

`func (o *GetHost200ResponseServer) GetEnableLogsOk() (*bool, bool)`

GetEnableLogsOk returns a tuple with the EnableLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLogs

`func (o *GetHost200ResponseServer) SetEnableLogs(v bool)`

SetEnableLogs sets EnableLogs field to given value.

### HasEnableLogs

`func (o *GetHost200ResponseServer) HasEnableLogs() bool`

HasEnableLogs returns a boolean if a field has been set.

### GetHourlyCost

`func (o *GetHost200ResponseServer) GetHourlyCost() float32`

GetHourlyCost returns the HourlyCost field if non-nil, zero value otherwise.

### GetHourlyCostOk

`func (o *GetHost200ResponseServer) GetHourlyCostOk() (*float32, bool)`

GetHourlyCostOk returns a tuple with the HourlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyCost

`func (o *GetHost200ResponseServer) SetHourlyCost(v float32)`

SetHourlyCost sets HourlyCost field to given value.

### HasHourlyCost

`func (o *GetHost200ResponseServer) HasHourlyCost() bool`

HasHourlyCost returns a boolean if a field has been set.

### GetHourlyPrice

`func (o *GetHost200ResponseServer) GetHourlyPrice() float32`

GetHourlyPrice returns the HourlyPrice field if non-nil, zero value otherwise.

### GetHourlyPriceOk

`func (o *GetHost200ResponseServer) GetHourlyPriceOk() (*float32, bool)`

GetHourlyPriceOk returns a tuple with the HourlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyPrice

`func (o *GetHost200ResponseServer) SetHourlyPrice(v float32)`

SetHourlyPrice sets HourlyPrice field to given value.

### HasHourlyPrice

`func (o *GetHost200ResponseServer) HasHourlyPrice() bool`

HasHourlyPrice returns a boolean if a field has been set.

### GetSourceImage

`func (o *GetHost200ResponseServer) GetSourceImage() AddBaremetalHost200ResponseServerSourceImage`

GetSourceImage returns the SourceImage field if non-nil, zero value otherwise.

### GetSourceImageOk

`func (o *GetHost200ResponseServer) GetSourceImageOk() (*AddBaremetalHost200ResponseServerSourceImage, bool)`

GetSourceImageOk returns a tuple with the SourceImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceImage

`func (o *GetHost200ResponseServer) SetSourceImage(v AddBaremetalHost200ResponseServerSourceImage)`

SetSourceImage sets SourceImage field to given value.

### HasSourceImage

`func (o *GetHost200ResponseServer) HasSourceImage() bool`

HasSourceImage returns a boolean if a field has been set.

### GetServerOs

`func (o *GetHost200ResponseServer) GetServerOs() AddBaremetalHost200ResponseServerServerOs`

GetServerOs returns the ServerOs field if non-nil, zero value otherwise.

### GetServerOsOk

`func (o *GetHost200ResponseServer) GetServerOsOk() (*AddBaremetalHost200ResponseServerServerOs, bool)`

GetServerOsOk returns a tuple with the ServerOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerOs

`func (o *GetHost200ResponseServer) SetServerOs(v AddBaremetalHost200ResponseServerServerOs)`

SetServerOs sets ServerOs field to given value.

### HasServerOs

`func (o *GetHost200ResponseServer) HasServerOs() bool`

HasServerOs returns a boolean if a field has been set.

### GetVolumes

`func (o *GetHost200ResponseServer) GetVolumes() []AddBaremetalHost200ResponseServerVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *GetHost200ResponseServer) GetVolumesOk() (*[]AddBaremetalHost200ResponseServerVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *GetHost200ResponseServer) SetVolumes(v []AddBaremetalHost200ResponseServerVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *GetHost200ResponseServer) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetControllers

`func (o *GetHost200ResponseServer) GetControllers() []AddBaremetalHost200ResponseServerControllersInner`

GetControllers returns the Controllers field if non-nil, zero value otherwise.

### GetControllersOk

`func (o *GetHost200ResponseServer) GetControllersOk() (*[]AddBaremetalHost200ResponseServerControllersInner, bool)`

GetControllersOk returns a tuple with the Controllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllers

`func (o *GetHost200ResponseServer) SetControllers(v []AddBaremetalHost200ResponseServerControllersInner)`

SetControllers sets Controllers field to given value.

### HasControllers

`func (o *GetHost200ResponseServer) HasControllers() bool`

HasControllers returns a boolean if a field has been set.

### GetInterfaces

`func (o *GetHost200ResponseServer) GetInterfaces() []AddBaremetalHost200ResponseServerInterfacesInner`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *GetHost200ResponseServer) GetInterfacesOk() (*[]AddBaremetalHost200ResponseServerInterfacesInner, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *GetHost200ResponseServer) SetInterfaces(v []AddBaremetalHost200ResponseServerInterfacesInner)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *GetHost200ResponseServer) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetLabels

`func (o *GetHost200ResponseServer) GetLabels() []map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetHost200ResponseServer) GetLabelsOk() (*[]map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetHost200ResponseServer) SetLabels(v []map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetHost200ResponseServer) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *GetHost200ResponseServer) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *GetHost200ResponseServer) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetTags

`func (o *GetHost200ResponseServer) GetTags() []map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetHost200ResponseServer) GetTagsOk() (*[]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetHost200ResponseServer) SetTags(v []map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetHost200ResponseServer) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *GetHost200ResponseServer) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *GetHost200ResponseServer) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetEnabled

`func (o *GetHost200ResponseServer) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetHost200ResponseServer) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetHost200ResponseServer) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetHost200ResponseServer) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTagCompliant

`func (o *GetHost200ResponseServer) GetTagCompliant() string`

GetTagCompliant returns the TagCompliant field if non-nil, zero value otherwise.

### GetTagCompliantOk

`func (o *GetHost200ResponseServer) GetTagCompliantOk() (*string, bool)`

GetTagCompliantOk returns a tuple with the TagCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagCompliant

`func (o *GetHost200ResponseServer) SetTagCompliant(v string)`

SetTagCompliant sets TagCompliant field to given value.

### HasTagCompliant

`func (o *GetHost200ResponseServer) HasTagCompliant() bool`

HasTagCompliant returns a boolean if a field has been set.

### SetTagCompliantNil

`func (o *GetHost200ResponseServer) SetTagCompliantNil(b bool)`

 SetTagCompliantNil sets the value for TagCompliant to be an explicit nil

### UnsetTagCompliant
`func (o *GetHost200ResponseServer) UnsetTagCompliant()`

UnsetTagCompliant ensures that no value is present for TagCompliant, not even an explicit nil
### GetContainers

`func (o *GetHost200ResponseServer) GetContainers() []int64`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *GetHost200ResponseServer) GetContainersOk() (*[]int64, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *GetHost200ResponseServer) SetContainers(v []int64)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *GetHost200ResponseServer) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetConfig

`func (o *GetHost200ResponseServer) GetConfig() AddBaremetalHost200ResponseServerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetHost200ResponseServer) GetConfigOk() (*AddBaremetalHost200ResponseServerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetHost200ResponseServer) SetConfig(v AddBaremetalHost200ResponseServerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetHost200ResponseServer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetInstance

`func (o *GetHost200ResponseServer) GetInstance() AddBaremetalHost200ResponseServerInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *GetHost200ResponseServer) GetInstanceOk() (*AddBaremetalHost200ResponseServerInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *GetHost200ResponseServer) SetInstance(v AddBaremetalHost200ResponseServerInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *GetHost200ResponseServer) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetGuestConsolePreferred

`func (o *GetHost200ResponseServer) GetGuestConsolePreferred() bool`

GetGuestConsolePreferred returns the GuestConsolePreferred field if non-nil, zero value otherwise.

### GetGuestConsolePreferredOk

`func (o *GetHost200ResponseServer) GetGuestConsolePreferredOk() (*bool, bool)`

GetGuestConsolePreferredOk returns a tuple with the GuestConsolePreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsolePreferred

`func (o *GetHost200ResponseServer) SetGuestConsolePreferred(v bool)`

SetGuestConsolePreferred sets GuestConsolePreferred field to given value.

### HasGuestConsolePreferred

`func (o *GetHost200ResponseServer) HasGuestConsolePreferred() bool`

HasGuestConsolePreferred returns a boolean if a field has been set.

### GetGuestConsoleType

`func (o *GetHost200ResponseServer) GetGuestConsoleType() string`

GetGuestConsoleType returns the GuestConsoleType field if non-nil, zero value otherwise.

### GetGuestConsoleTypeOk

`func (o *GetHost200ResponseServer) GetGuestConsoleTypeOk() (*string, bool)`

GetGuestConsoleTypeOk returns a tuple with the GuestConsoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleType

`func (o *GetHost200ResponseServer) SetGuestConsoleType(v string)`

SetGuestConsoleType sets GuestConsoleType field to given value.

### HasGuestConsoleType

`func (o *GetHost200ResponseServer) HasGuestConsoleType() bool`

HasGuestConsoleType returns a boolean if a field has been set.

### SetGuestConsoleTypeNil

`func (o *GetHost200ResponseServer) SetGuestConsoleTypeNil(b bool)`

 SetGuestConsoleTypeNil sets the value for GuestConsoleType to be an explicit nil

### UnsetGuestConsoleType
`func (o *GetHost200ResponseServer) UnsetGuestConsoleType()`

UnsetGuestConsoleType ensures that no value is present for GuestConsoleType, not even an explicit nil
### GetGuestConsoleUsername

`func (o *GetHost200ResponseServer) GetGuestConsoleUsername() string`

GetGuestConsoleUsername returns the GuestConsoleUsername field if non-nil, zero value otherwise.

### GetGuestConsoleUsernameOk

`func (o *GetHost200ResponseServer) GetGuestConsoleUsernameOk() (*string, bool)`

GetGuestConsoleUsernameOk returns a tuple with the GuestConsoleUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleUsername

`func (o *GetHost200ResponseServer) SetGuestConsoleUsername(v string)`

SetGuestConsoleUsername sets GuestConsoleUsername field to given value.

### HasGuestConsoleUsername

`func (o *GetHost200ResponseServer) HasGuestConsoleUsername() bool`

HasGuestConsoleUsername returns a boolean if a field has been set.

### SetGuestConsoleUsernameNil

`func (o *GetHost200ResponseServer) SetGuestConsoleUsernameNil(b bool)`

 SetGuestConsoleUsernameNil sets the value for GuestConsoleUsername to be an explicit nil

### UnsetGuestConsoleUsername
`func (o *GetHost200ResponseServer) UnsetGuestConsoleUsername()`

UnsetGuestConsoleUsername ensures that no value is present for GuestConsoleUsername, not even an explicit nil
### GetGuestConsolePassword

`func (o *GetHost200ResponseServer) GetGuestConsolePassword() string`

GetGuestConsolePassword returns the GuestConsolePassword field if non-nil, zero value otherwise.

### GetGuestConsolePasswordOk

`func (o *GetHost200ResponseServer) GetGuestConsolePasswordOk() (*string, bool)`

GetGuestConsolePasswordOk returns a tuple with the GuestConsolePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsolePassword

`func (o *GetHost200ResponseServer) SetGuestConsolePassword(v string)`

SetGuestConsolePassword sets GuestConsolePassword field to given value.

### HasGuestConsolePassword

`func (o *GetHost200ResponseServer) HasGuestConsolePassword() bool`

HasGuestConsolePassword returns a boolean if a field has been set.

### SetGuestConsolePasswordNil

`func (o *GetHost200ResponseServer) SetGuestConsolePasswordNil(b bool)`

 SetGuestConsolePasswordNil sets the value for GuestConsolePassword to be an explicit nil

### UnsetGuestConsolePassword
`func (o *GetHost200ResponseServer) UnsetGuestConsolePassword()`

UnsetGuestConsolePassword ensures that no value is present for GuestConsolePassword, not even an explicit nil
### GetGuestConsolePasswordHash

`func (o *GetHost200ResponseServer) GetGuestConsolePasswordHash() string`

GetGuestConsolePasswordHash returns the GuestConsolePasswordHash field if non-nil, zero value otherwise.

### GetGuestConsolePasswordHashOk

`func (o *GetHost200ResponseServer) GetGuestConsolePasswordHashOk() (*string, bool)`

GetGuestConsolePasswordHashOk returns a tuple with the GuestConsolePasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsolePasswordHash

`func (o *GetHost200ResponseServer) SetGuestConsolePasswordHash(v string)`

SetGuestConsolePasswordHash sets GuestConsolePasswordHash field to given value.

### HasGuestConsolePasswordHash

`func (o *GetHost200ResponseServer) HasGuestConsolePasswordHash() bool`

HasGuestConsolePasswordHash returns a boolean if a field has been set.

### SetGuestConsolePasswordHashNil

`func (o *GetHost200ResponseServer) SetGuestConsolePasswordHashNil(b bool)`

 SetGuestConsolePasswordHashNil sets the value for GuestConsolePasswordHash to be an explicit nil

### UnsetGuestConsolePasswordHash
`func (o *GetHost200ResponseServer) UnsetGuestConsolePasswordHash()`

UnsetGuestConsolePasswordHash ensures that no value is present for GuestConsolePasswordHash, not even an explicit nil
### GetGuestConsolePort

`func (o *GetHost200ResponseServer) GetGuestConsolePort() string`

GetGuestConsolePort returns the GuestConsolePort field if non-nil, zero value otherwise.

### GetGuestConsolePortOk

`func (o *GetHost200ResponseServer) GetGuestConsolePortOk() (*string, bool)`

GetGuestConsolePortOk returns a tuple with the GuestConsolePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsolePort

`func (o *GetHost200ResponseServer) SetGuestConsolePort(v string)`

SetGuestConsolePort sets GuestConsolePort field to given value.

### HasGuestConsolePort

`func (o *GetHost200ResponseServer) HasGuestConsolePort() bool`

HasGuestConsolePort returns a boolean if a field has been set.

### SetGuestConsolePortNil

`func (o *GetHost200ResponseServer) SetGuestConsolePortNil(b bool)`

 SetGuestConsolePortNil sets the value for GuestConsolePort to be an explicit nil

### UnsetGuestConsolePort
`func (o *GetHost200ResponseServer) UnsetGuestConsolePort()`

UnsetGuestConsolePort ensures that no value is present for GuestConsolePort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


