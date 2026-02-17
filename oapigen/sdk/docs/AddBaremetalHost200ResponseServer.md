# AddBaremetalHost200ResponseServer

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

### NewAddBaremetalHost200ResponseServer

`func NewAddBaremetalHost200ResponseServer() *AddBaremetalHost200ResponseServer`

NewAddBaremetalHost200ResponseServer instantiates a new AddBaremetalHost200ResponseServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBaremetalHost200ResponseServerWithDefaults

`func NewAddBaremetalHost200ResponseServerWithDefaults() *AddBaremetalHost200ResponseServer`

NewAddBaremetalHost200ResponseServerWithDefaults instantiates a new AddBaremetalHost200ResponseServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddBaremetalHost200ResponseServer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddBaremetalHost200ResponseServer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddBaremetalHost200ResponseServer) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddBaremetalHost200ResponseServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *AddBaremetalHost200ResponseServer) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AddBaremetalHost200ResponseServer) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AddBaremetalHost200ResponseServer) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AddBaremetalHost200ResponseServer) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetExternalId

`func (o *AddBaremetalHost200ResponseServer) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AddBaremetalHost200ResponseServer) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AddBaremetalHost200ResponseServer) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AddBaremetalHost200ResponseServer) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *AddBaremetalHost200ResponseServer) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *AddBaremetalHost200ResponseServer) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetInternalId

`func (o *AddBaremetalHost200ResponseServer) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *AddBaremetalHost200ResponseServer) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *AddBaremetalHost200ResponseServer) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *AddBaremetalHost200ResponseServer) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *AddBaremetalHost200ResponseServer) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *AddBaremetalHost200ResponseServer) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalUniqueId

`func (o *AddBaremetalHost200ResponseServer) GetExternalUniqueId() string`

GetExternalUniqueId returns the ExternalUniqueId field if non-nil, zero value otherwise.

### GetExternalUniqueIdOk

`func (o *AddBaremetalHost200ResponseServer) GetExternalUniqueIdOk() (*string, bool)`

GetExternalUniqueIdOk returns a tuple with the ExternalUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUniqueId

`func (o *AddBaremetalHost200ResponseServer) SetExternalUniqueId(v string)`

SetExternalUniqueId sets ExternalUniqueId field to given value.

### HasExternalUniqueId

`func (o *AddBaremetalHost200ResponseServer) HasExternalUniqueId() bool`

HasExternalUniqueId returns a boolean if a field has been set.

### SetExternalUniqueIdNil

`func (o *AddBaremetalHost200ResponseServer) SetExternalUniqueIdNil(b bool)`

 SetExternalUniqueIdNil sets the value for ExternalUniqueId to be an explicit nil

### UnsetExternalUniqueId
`func (o *AddBaremetalHost200ResponseServer) UnsetExternalUniqueId()`

UnsetExternalUniqueId ensures that no value is present for ExternalUniqueId, not even an explicit nil
### GetName

`func (o *AddBaremetalHost200ResponseServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddBaremetalHost200ResponseServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddBaremetalHost200ResponseServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddBaremetalHost200ResponseServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalName

`func (o *AddBaremetalHost200ResponseServer) GetExternalName() string`

GetExternalName returns the ExternalName field if non-nil, zero value otherwise.

### GetExternalNameOk

`func (o *AddBaremetalHost200ResponseServer) GetExternalNameOk() (*string, bool)`

GetExternalNameOk returns a tuple with the ExternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalName

`func (o *AddBaremetalHost200ResponseServer) SetExternalName(v string)`

SetExternalName sets ExternalName field to given value.

### HasExternalName

`func (o *AddBaremetalHost200ResponseServer) HasExternalName() bool`

HasExternalName returns a boolean if a field has been set.

### GetHostname

`func (o *AddBaremetalHost200ResponseServer) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *AddBaremetalHost200ResponseServer) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *AddBaremetalHost200ResponseServer) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *AddBaremetalHost200ResponseServer) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetParentServer

`func (o *AddBaremetalHost200ResponseServer) GetParentServer() AddBaremetalHost200ResponseServerParentServer`

GetParentServer returns the ParentServer field if non-nil, zero value otherwise.

### GetParentServerOk

`func (o *AddBaremetalHost200ResponseServer) GetParentServerOk() (*AddBaremetalHost200ResponseServerParentServer, bool)`

GetParentServerOk returns a tuple with the ParentServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentServer

`func (o *AddBaremetalHost200ResponseServer) SetParentServer(v AddBaremetalHost200ResponseServerParentServer)`

SetParentServer sets ParentServer field to given value.

### HasParentServer

`func (o *AddBaremetalHost200ResponseServer) HasParentServer() bool`

HasParentServer returns a boolean if a field has been set.

### GetAccountId

`func (o *AddBaremetalHost200ResponseServer) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AddBaremetalHost200ResponseServer) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AddBaremetalHost200ResponseServer) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AddBaremetalHost200ResponseServer) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccount

`func (o *AddBaremetalHost200ResponseServer) GetAccount() AddBaremetalHost200ResponseServerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AddBaremetalHost200ResponseServer) GetAccountOk() (*AddBaremetalHost200ResponseServerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AddBaremetalHost200ResponseServer) SetAccount(v AddBaremetalHost200ResponseServerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AddBaremetalHost200ResponseServer) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *AddBaremetalHost200ResponseServer) GetOwner() AddBaremetalHost200ResponseServerOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AddBaremetalHost200ResponseServer) GetOwnerOk() (*AddBaremetalHost200ResponseServerOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AddBaremetalHost200ResponseServer) SetOwner(v AddBaremetalHost200ResponseServerOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AddBaremetalHost200ResponseServer) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetZone

`func (o *AddBaremetalHost200ResponseServer) GetZone() AddBaremetalHost200ResponseServerZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *AddBaremetalHost200ResponseServer) GetZoneOk() (*AddBaremetalHost200ResponseServerZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *AddBaremetalHost200ResponseServer) SetZone(v AddBaremetalHost200ResponseServerZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *AddBaremetalHost200ResponseServer) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetPlan

`func (o *AddBaremetalHost200ResponseServer) GetPlan() AddBaremetalHost200ResponseServerPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *AddBaremetalHost200ResponseServer) GetPlanOk() (*AddBaremetalHost200ResponseServerPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *AddBaremetalHost200ResponseServer) SetPlan(v AddBaremetalHost200ResponseServerPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *AddBaremetalHost200ResponseServer) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetComputeServerType

`func (o *AddBaremetalHost200ResponseServer) GetComputeServerType() AddBaremetalHost200ResponseServerComputeServerType`

GetComputeServerType returns the ComputeServerType field if non-nil, zero value otherwise.

### GetComputeServerTypeOk

`func (o *AddBaremetalHost200ResponseServer) GetComputeServerTypeOk() (*AddBaremetalHost200ResponseServerComputeServerType, bool)`

GetComputeServerTypeOk returns a tuple with the ComputeServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServerType

`func (o *AddBaremetalHost200ResponseServer) SetComputeServerType(v AddBaremetalHost200ResponseServerComputeServerType)`

SetComputeServerType sets ComputeServerType field to given value.

### HasComputeServerType

`func (o *AddBaremetalHost200ResponseServer) HasComputeServerType() bool`

HasComputeServerType returns a boolean if a field has been set.

### GetVisibility

`func (o *AddBaremetalHost200ResponseServer) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AddBaremetalHost200ResponseServer) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AddBaremetalHost200ResponseServer) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AddBaremetalHost200ResponseServer) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *AddBaremetalHost200ResponseServer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddBaremetalHost200ResponseServer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddBaremetalHost200ResponseServer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddBaremetalHost200ResponseServer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AddBaremetalHost200ResponseServer) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AddBaremetalHost200ResponseServer) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetZoneId

`func (o *AddBaremetalHost200ResponseServer) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *AddBaremetalHost200ResponseServer) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *AddBaremetalHost200ResponseServer) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *AddBaremetalHost200ResponseServer) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetSiteId

`func (o *AddBaremetalHost200ResponseServer) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *AddBaremetalHost200ResponseServer) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *AddBaremetalHost200ResponseServer) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *AddBaremetalHost200ResponseServer) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *AddBaremetalHost200ResponseServer) GetResourcePoolId() int64`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *AddBaremetalHost200ResponseServer) GetResourcePoolIdOk() (*int64, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *AddBaremetalHost200ResponseServer) SetResourcePoolId(v int64)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *AddBaremetalHost200ResponseServer) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### SetResourcePoolIdNil

`func (o *AddBaremetalHost200ResponseServer) SetResourcePoolIdNil(b bool)`

 SetResourcePoolIdNil sets the value for ResourcePoolId to be an explicit nil

### UnsetResourcePoolId
`func (o *AddBaremetalHost200ResponseServer) UnsetResourcePoolId()`

UnsetResourcePoolId ensures that no value is present for ResourcePoolId, not even an explicit nil
### GetFolderId

`func (o *AddBaremetalHost200ResponseServer) GetFolderId() int64`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *AddBaremetalHost200ResponseServer) GetFolderIdOk() (*int64, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *AddBaremetalHost200ResponseServer) SetFolderId(v int64)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *AddBaremetalHost200ResponseServer) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *AddBaremetalHost200ResponseServer) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *AddBaremetalHost200ResponseServer) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetSshHost

`func (o *AddBaremetalHost200ResponseServer) GetSshHost() string`

GetSshHost returns the SshHost field if non-nil, zero value otherwise.

### GetSshHostOk

`func (o *AddBaremetalHost200ResponseServer) GetSshHostOk() (*string, bool)`

GetSshHostOk returns a tuple with the SshHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHost

`func (o *AddBaremetalHost200ResponseServer) SetSshHost(v string)`

SetSshHost sets SshHost field to given value.

### HasSshHost

`func (o *AddBaremetalHost200ResponseServer) HasSshHost() bool`

HasSshHost returns a boolean if a field has been set.

### SetSshHostNil

`func (o *AddBaremetalHost200ResponseServer) SetSshHostNil(b bool)`

 SetSshHostNil sets the value for SshHost to be an explicit nil

### UnsetSshHost
`func (o *AddBaremetalHost200ResponseServer) UnsetSshHost()`

UnsetSshHost ensures that no value is present for SshHost, not even an explicit nil
### GetSshPort

`func (o *AddBaremetalHost200ResponseServer) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *AddBaremetalHost200ResponseServer) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *AddBaremetalHost200ResponseServer) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *AddBaremetalHost200ResponseServer) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetExternalIp

`func (o *AddBaremetalHost200ResponseServer) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *AddBaremetalHost200ResponseServer) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *AddBaremetalHost200ResponseServer) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *AddBaremetalHost200ResponseServer) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### SetExternalIpNil

`func (o *AddBaremetalHost200ResponseServer) SetExternalIpNil(b bool)`

 SetExternalIpNil sets the value for ExternalIp to be an explicit nil

### UnsetExternalIp
`func (o *AddBaremetalHost200ResponseServer) UnsetExternalIp()`

UnsetExternalIp ensures that no value is present for ExternalIp, not even an explicit nil
### GetInternalIp

`func (o *AddBaremetalHost200ResponseServer) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *AddBaremetalHost200ResponseServer) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *AddBaremetalHost200ResponseServer) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *AddBaremetalHost200ResponseServer) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### SetInternalIpNil

`func (o *AddBaremetalHost200ResponseServer) SetInternalIpNil(b bool)`

 SetInternalIpNil sets the value for InternalIp to be an explicit nil

### UnsetInternalIp
`func (o *AddBaremetalHost200ResponseServer) UnsetInternalIp()`

UnsetInternalIp ensures that no value is present for InternalIp, not even an explicit nil
### GetVolumeId

`func (o *AddBaremetalHost200ResponseServer) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *AddBaremetalHost200ResponseServer) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *AddBaremetalHost200ResponseServer) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *AddBaremetalHost200ResponseServer) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.

### SetVolumeIdNil

`func (o *AddBaremetalHost200ResponseServer) SetVolumeIdNil(b bool)`

 SetVolumeIdNil sets the value for VolumeId to be an explicit nil

### UnsetVolumeId
`func (o *AddBaremetalHost200ResponseServer) UnsetVolumeId()`

UnsetVolumeId ensures that no value is present for VolumeId, not even an explicit nil
### GetPlatform

`func (o *AddBaremetalHost200ResponseServer) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *AddBaremetalHost200ResponseServer) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *AddBaremetalHost200ResponseServer) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *AddBaremetalHost200ResponseServer) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *AddBaremetalHost200ResponseServer) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *AddBaremetalHost200ResponseServer) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetPlatformVersion

`func (o *AddBaremetalHost200ResponseServer) GetPlatformVersion() string`

GetPlatformVersion returns the PlatformVersion field if non-nil, zero value otherwise.

### GetPlatformVersionOk

`func (o *AddBaremetalHost200ResponseServer) GetPlatformVersionOk() (*string, bool)`

GetPlatformVersionOk returns a tuple with the PlatformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformVersion

`func (o *AddBaremetalHost200ResponseServer) SetPlatformVersion(v string)`

SetPlatformVersion sets PlatformVersion field to given value.

### HasPlatformVersion

`func (o *AddBaremetalHost200ResponseServer) HasPlatformVersion() bool`

HasPlatformVersion returns a boolean if a field has been set.

### SetPlatformVersionNil

`func (o *AddBaremetalHost200ResponseServer) SetPlatformVersionNil(b bool)`

 SetPlatformVersionNil sets the value for PlatformVersion to be an explicit nil

### UnsetPlatformVersion
`func (o *AddBaremetalHost200ResponseServer) UnsetPlatformVersion()`

UnsetPlatformVersion ensures that no value is present for PlatformVersion, not even an explicit nil
### GetSshUsername

`func (o *AddBaremetalHost200ResponseServer) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *AddBaremetalHost200ResponseServer) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *AddBaremetalHost200ResponseServer) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *AddBaremetalHost200ResponseServer) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### SetSshUsernameNil

`func (o *AddBaremetalHost200ResponseServer) SetSshUsernameNil(b bool)`

 SetSshUsernameNil sets the value for SshUsername to be an explicit nil

### UnsetSshUsername
`func (o *AddBaremetalHost200ResponseServer) UnsetSshUsername()`

UnsetSshUsername ensures that no value is present for SshUsername, not even an explicit nil
### GetSshPassword

`func (o *AddBaremetalHost200ResponseServer) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *AddBaremetalHost200ResponseServer) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *AddBaremetalHost200ResponseServer) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *AddBaremetalHost200ResponseServer) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### SetSshPasswordNil

`func (o *AddBaremetalHost200ResponseServer) SetSshPasswordNil(b bool)`

 SetSshPasswordNil sets the value for SshPassword to be an explicit nil

### UnsetSshPassword
`func (o *AddBaremetalHost200ResponseServer) UnsetSshPassword()`

UnsetSshPassword ensures that no value is present for SshPassword, not even an explicit nil
### GetSshPasswordHash

`func (o *AddBaremetalHost200ResponseServer) GetSshPasswordHash() string`

GetSshPasswordHash returns the SshPasswordHash field if non-nil, zero value otherwise.

### GetSshPasswordHashOk

`func (o *AddBaremetalHost200ResponseServer) GetSshPasswordHashOk() (*string, bool)`

GetSshPasswordHashOk returns a tuple with the SshPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPasswordHash

`func (o *AddBaremetalHost200ResponseServer) SetSshPasswordHash(v string)`

SetSshPasswordHash sets SshPasswordHash field to given value.

### HasSshPasswordHash

`func (o *AddBaremetalHost200ResponseServer) HasSshPasswordHash() bool`

HasSshPasswordHash returns a boolean if a field has been set.

### SetSshPasswordHashNil

`func (o *AddBaremetalHost200ResponseServer) SetSshPasswordHashNil(b bool)`

 SetSshPasswordHashNil sets the value for SshPasswordHash to be an explicit nil

### UnsetSshPasswordHash
`func (o *AddBaremetalHost200ResponseServer) UnsetSshPasswordHash()`

UnsetSshPasswordHash ensures that no value is present for SshPasswordHash, not even an explicit nil
### GetSshKeyPair

`func (o *AddBaremetalHost200ResponseServer) GetSshKeyPair() AddBaremetalHost200ResponseServerSshKeyPair`

GetSshKeyPair returns the SshKeyPair field if non-nil, zero value otherwise.

### GetSshKeyPairOk

`func (o *AddBaremetalHost200ResponseServer) GetSshKeyPairOk() (*AddBaremetalHost200ResponseServerSshKeyPair, bool)`

GetSshKeyPairOk returns a tuple with the SshKeyPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyPair

`func (o *AddBaremetalHost200ResponseServer) SetSshKeyPair(v AddBaremetalHost200ResponseServerSshKeyPair)`

SetSshKeyPair sets SshKeyPair field to given value.

### HasSshKeyPair

`func (o *AddBaremetalHost200ResponseServer) HasSshKeyPair() bool`

HasSshKeyPair returns a boolean if a field has been set.

### GetOsDevice

`func (o *AddBaremetalHost200ResponseServer) GetOsDevice() string`

GetOsDevice returns the OsDevice field if non-nil, zero value otherwise.

### GetOsDeviceOk

`func (o *AddBaremetalHost200ResponseServer) GetOsDeviceOk() (*string, bool)`

GetOsDeviceOk returns a tuple with the OsDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDevice

`func (o *AddBaremetalHost200ResponseServer) SetOsDevice(v string)`

SetOsDevice sets OsDevice field to given value.

### HasOsDevice

`func (o *AddBaremetalHost200ResponseServer) HasOsDevice() bool`

HasOsDevice returns a boolean if a field has been set.

### GetOsType

`func (o *AddBaremetalHost200ResponseServer) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *AddBaremetalHost200ResponseServer) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *AddBaremetalHost200ResponseServer) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *AddBaremetalHost200ResponseServer) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetDataDevice

`func (o *AddBaremetalHost200ResponseServer) GetDataDevice() string`

GetDataDevice returns the DataDevice field if non-nil, zero value otherwise.

### GetDataDeviceOk

`func (o *AddBaremetalHost200ResponseServer) GetDataDeviceOk() (*string, bool)`

GetDataDeviceOk returns a tuple with the DataDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDevice

`func (o *AddBaremetalHost200ResponseServer) SetDataDevice(v string)`

SetDataDevice sets DataDevice field to given value.

### HasDataDevice

`func (o *AddBaremetalHost200ResponseServer) HasDataDevice() bool`

HasDataDevice returns a boolean if a field has been set.

### GetLvmEnabled

`func (o *AddBaremetalHost200ResponseServer) GetLvmEnabled() bool`

GetLvmEnabled returns the LvmEnabled field if non-nil, zero value otherwise.

### GetLvmEnabledOk

`func (o *AddBaremetalHost200ResponseServer) GetLvmEnabledOk() (*bool, bool)`

GetLvmEnabledOk returns a tuple with the LvmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLvmEnabled

`func (o *AddBaremetalHost200ResponseServer) SetLvmEnabled(v bool)`

SetLvmEnabled sets LvmEnabled field to given value.

### HasLvmEnabled

`func (o *AddBaremetalHost200ResponseServer) HasLvmEnabled() bool`

HasLvmEnabled returns a boolean if a field has been set.

### GetApiKey

`func (o *AddBaremetalHost200ResponseServer) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *AddBaremetalHost200ResponseServer) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *AddBaremetalHost200ResponseServer) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *AddBaremetalHost200ResponseServer) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetSoftwareRaid

`func (o *AddBaremetalHost200ResponseServer) GetSoftwareRaid() bool`

GetSoftwareRaid returns the SoftwareRaid field if non-nil, zero value otherwise.

### GetSoftwareRaidOk

`func (o *AddBaremetalHost200ResponseServer) GetSoftwareRaidOk() (*bool, bool)`

GetSoftwareRaidOk returns a tuple with the SoftwareRaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareRaid

`func (o *AddBaremetalHost200ResponseServer) SetSoftwareRaid(v bool)`

SetSoftwareRaid sets SoftwareRaid field to given value.

### HasSoftwareRaid

`func (o *AddBaremetalHost200ResponseServer) HasSoftwareRaid() bool`

HasSoftwareRaid returns a boolean if a field has been set.

### GetDateCreated

`func (o *AddBaremetalHost200ResponseServer) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddBaremetalHost200ResponseServer) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddBaremetalHost200ResponseServer) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddBaremetalHost200ResponseServer) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddBaremetalHost200ResponseServer) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddBaremetalHost200ResponseServer) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddBaremetalHost200ResponseServer) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddBaremetalHost200ResponseServer) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetStats

`func (o *AddBaremetalHost200ResponseServer) GetStats() AddBaremetalHost200ResponseServerStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *AddBaremetalHost200ResponseServer) GetStatsOk() (*AddBaremetalHost200ResponseServerStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *AddBaremetalHost200ResponseServer) SetStats(v AddBaremetalHost200ResponseServerStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *AddBaremetalHost200ResponseServer) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetStatus

`func (o *AddBaremetalHost200ResponseServer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddBaremetalHost200ResponseServer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddBaremetalHost200ResponseServer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddBaremetalHost200ResponseServer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *AddBaremetalHost200ResponseServer) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *AddBaremetalHost200ResponseServer) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *AddBaremetalHost200ResponseServer) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *AddBaremetalHost200ResponseServer) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *AddBaremetalHost200ResponseServer) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *AddBaremetalHost200ResponseServer) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetErrorMessage

`func (o *AddBaremetalHost200ResponseServer) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AddBaremetalHost200ResponseServer) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AddBaremetalHost200ResponseServer) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AddBaremetalHost200ResponseServer) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *AddBaremetalHost200ResponseServer) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *AddBaremetalHost200ResponseServer) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetStatusDate

`func (o *AddBaremetalHost200ResponseServer) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *AddBaremetalHost200ResponseServer) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *AddBaremetalHost200ResponseServer) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *AddBaremetalHost200ResponseServer) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### SetStatusDateNil

`func (o *AddBaremetalHost200ResponseServer) SetStatusDateNil(b bool)`

 SetStatusDateNil sets the value for StatusDate to be an explicit nil

### UnsetStatusDate
`func (o *AddBaremetalHost200ResponseServer) UnsetStatusDate()`

UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
### GetStatusPercent

`func (o *AddBaremetalHost200ResponseServer) GetStatusPercent() string`

GetStatusPercent returns the StatusPercent field if non-nil, zero value otherwise.

### GetStatusPercentOk

`func (o *AddBaremetalHost200ResponseServer) GetStatusPercentOk() (*string, bool)`

GetStatusPercentOk returns a tuple with the StatusPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusPercent

`func (o *AddBaremetalHost200ResponseServer) SetStatusPercent(v string)`

SetStatusPercent sets StatusPercent field to given value.

### HasStatusPercent

`func (o *AddBaremetalHost200ResponseServer) HasStatusPercent() bool`

HasStatusPercent returns a boolean if a field has been set.

### SetStatusPercentNil

`func (o *AddBaremetalHost200ResponseServer) SetStatusPercentNil(b bool)`

 SetStatusPercentNil sets the value for StatusPercent to be an explicit nil

### UnsetStatusPercent
`func (o *AddBaremetalHost200ResponseServer) UnsetStatusPercent()`

UnsetStatusPercent ensures that no value is present for StatusPercent, not even an explicit nil
### GetStatusEta

`func (o *AddBaremetalHost200ResponseServer) GetStatusEta() string`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *AddBaremetalHost200ResponseServer) GetStatusEtaOk() (*string, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *AddBaremetalHost200ResponseServer) SetStatusEta(v string)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *AddBaremetalHost200ResponseServer) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### SetStatusEtaNil

`func (o *AddBaremetalHost200ResponseServer) SetStatusEtaNil(b bool)`

 SetStatusEtaNil sets the value for StatusEta to be an explicit nil

### UnsetStatusEta
`func (o *AddBaremetalHost200ResponseServer) UnsetStatusEta()`

UnsetStatusEta ensures that no value is present for StatusEta, not even an explicit nil
### GetPowerState

`func (o *AddBaremetalHost200ResponseServer) GetPowerState() string`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *AddBaremetalHost200ResponseServer) GetPowerStateOk() (*string, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *AddBaremetalHost200ResponseServer) SetPowerState(v string)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *AddBaremetalHost200ResponseServer) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetAgentInstalled

`func (o *AddBaremetalHost200ResponseServer) GetAgentInstalled() bool`

GetAgentInstalled returns the AgentInstalled field if non-nil, zero value otherwise.

### GetAgentInstalledOk

`func (o *AddBaremetalHost200ResponseServer) GetAgentInstalledOk() (*bool, bool)`

GetAgentInstalledOk returns a tuple with the AgentInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentInstalled

`func (o *AddBaremetalHost200ResponseServer) SetAgentInstalled(v bool)`

SetAgentInstalled sets AgentInstalled field to given value.

### HasAgentInstalled

`func (o *AddBaremetalHost200ResponseServer) HasAgentInstalled() bool`

HasAgentInstalled returns a boolean if a field has been set.

### GetLastAgentUpdate

`func (o *AddBaremetalHost200ResponseServer) GetLastAgentUpdate() string`

GetLastAgentUpdate returns the LastAgentUpdate field if non-nil, zero value otherwise.

### GetLastAgentUpdateOk

`func (o *AddBaremetalHost200ResponseServer) GetLastAgentUpdateOk() (*string, bool)`

GetLastAgentUpdateOk returns a tuple with the LastAgentUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAgentUpdate

`func (o *AddBaremetalHost200ResponseServer) SetLastAgentUpdate(v string)`

SetLastAgentUpdate sets LastAgentUpdate field to given value.

### HasLastAgentUpdate

`func (o *AddBaremetalHost200ResponseServer) HasLastAgentUpdate() bool`

HasLastAgentUpdate returns a boolean if a field has been set.

### SetLastAgentUpdateNil

`func (o *AddBaremetalHost200ResponseServer) SetLastAgentUpdateNil(b bool)`

 SetLastAgentUpdateNil sets the value for LastAgentUpdate to be an explicit nil

### UnsetLastAgentUpdate
`func (o *AddBaremetalHost200ResponseServer) UnsetLastAgentUpdate()`

UnsetLastAgentUpdate ensures that no value is present for LastAgentUpdate, not even an explicit nil
### GetAgentVersion

`func (o *AddBaremetalHost200ResponseServer) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *AddBaremetalHost200ResponseServer) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *AddBaremetalHost200ResponseServer) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *AddBaremetalHost200ResponseServer) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### SetAgentVersionNil

`func (o *AddBaremetalHost200ResponseServer) SetAgentVersionNil(b bool)`

 SetAgentVersionNil sets the value for AgentVersion to be an explicit nil

### UnsetAgentVersion
`func (o *AddBaremetalHost200ResponseServer) UnsetAgentVersion()`

UnsetAgentVersion ensures that no value is present for AgentVersion, not even an explicit nil
### GetMaxCores

`func (o *AddBaremetalHost200ResponseServer) GetMaxCores() int64`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *AddBaremetalHost200ResponseServer) GetMaxCoresOk() (*int64, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *AddBaremetalHost200ResponseServer) SetMaxCores(v int64)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *AddBaremetalHost200ResponseServer) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *AddBaremetalHost200ResponseServer) GetCoresPerSocket() int64`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *AddBaremetalHost200ResponseServer) GetCoresPerSocketOk() (*int64, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *AddBaremetalHost200ResponseServer) SetCoresPerSocket(v int64)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *AddBaremetalHost200ResponseServer) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### SetCoresPerSocketNil

`func (o *AddBaremetalHost200ResponseServer) SetCoresPerSocketNil(b bool)`

 SetCoresPerSocketNil sets the value for CoresPerSocket to be an explicit nil

### UnsetCoresPerSocket
`func (o *AddBaremetalHost200ResponseServer) UnsetCoresPerSocket()`

UnsetCoresPerSocket ensures that no value is present for CoresPerSocket, not even an explicit nil
### GetMaxMemory

`func (o *AddBaremetalHost200ResponseServer) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *AddBaremetalHost200ResponseServer) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *AddBaremetalHost200ResponseServer) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *AddBaremetalHost200ResponseServer) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxStorage

`func (o *AddBaremetalHost200ResponseServer) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *AddBaremetalHost200ResponseServer) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *AddBaremetalHost200ResponseServer) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *AddBaremetalHost200ResponseServer) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMaxCpu

`func (o *AddBaremetalHost200ResponseServer) GetMaxCpu() int64`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *AddBaremetalHost200ResponseServer) GetMaxCpuOk() (*int64, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *AddBaremetalHost200ResponseServer) SetMaxCpu(v int64)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *AddBaremetalHost200ResponseServer) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### SetMaxCpuNil

`func (o *AddBaremetalHost200ResponseServer) SetMaxCpuNil(b bool)`

 SetMaxCpuNil sets the value for MaxCpu to be an explicit nil

### UnsetMaxCpu
`func (o *AddBaremetalHost200ResponseServer) UnsetMaxCpu()`

UnsetMaxCpu ensures that no value is present for MaxCpu, not even an explicit nil
### GetMaxGpus

`func (o *AddBaremetalHost200ResponseServer) GetMaxGpus() int64`

GetMaxGpus returns the MaxGpus field if non-nil, zero value otherwise.

### GetMaxGpusOk

`func (o *AddBaremetalHost200ResponseServer) GetMaxGpusOk() (*int64, bool)`

GetMaxGpusOk returns a tuple with the MaxGpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGpus

`func (o *AddBaremetalHost200ResponseServer) SetMaxGpus(v int64)`

SetMaxGpus sets MaxGpus field to given value.

### HasMaxGpus

`func (o *AddBaremetalHost200ResponseServer) HasMaxGpus() bool`

HasMaxGpus returns a boolean if a field has been set.

### SetMaxGpusNil

`func (o *AddBaremetalHost200ResponseServer) SetMaxGpusNil(b bool)`

 SetMaxGpusNil sets the value for MaxGpus to be an explicit nil

### UnsetMaxGpus
`func (o *AddBaremetalHost200ResponseServer) UnsetMaxGpus()`

UnsetMaxGpus ensures that no value is present for MaxGpus, not even an explicit nil
### GetManageInternalFirewall

`func (o *AddBaremetalHost200ResponseServer) GetManageInternalFirewall() bool`

GetManageInternalFirewall returns the ManageInternalFirewall field if non-nil, zero value otherwise.

### GetManageInternalFirewallOk

`func (o *AddBaremetalHost200ResponseServer) GetManageInternalFirewallOk() (*bool, bool)`

GetManageInternalFirewallOk returns a tuple with the ManageInternalFirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageInternalFirewall

`func (o *AddBaremetalHost200ResponseServer) SetManageInternalFirewall(v bool)`

SetManageInternalFirewall sets ManageInternalFirewall field to given value.

### HasManageInternalFirewall

`func (o *AddBaremetalHost200ResponseServer) HasManageInternalFirewall() bool`

HasManageInternalFirewall returns a boolean if a field has been set.

### GetEnableLogs

`func (o *AddBaremetalHost200ResponseServer) GetEnableLogs() bool`

GetEnableLogs returns the EnableLogs field if non-nil, zero value otherwise.

### GetEnableLogsOk

`func (o *AddBaremetalHost200ResponseServer) GetEnableLogsOk() (*bool, bool)`

GetEnableLogsOk returns a tuple with the EnableLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLogs

`func (o *AddBaremetalHost200ResponseServer) SetEnableLogs(v bool)`

SetEnableLogs sets EnableLogs field to given value.

### HasEnableLogs

`func (o *AddBaremetalHost200ResponseServer) HasEnableLogs() bool`

HasEnableLogs returns a boolean if a field has been set.

### GetHourlyCost

`func (o *AddBaremetalHost200ResponseServer) GetHourlyCost() float32`

GetHourlyCost returns the HourlyCost field if non-nil, zero value otherwise.

### GetHourlyCostOk

`func (o *AddBaremetalHost200ResponseServer) GetHourlyCostOk() (*float32, bool)`

GetHourlyCostOk returns a tuple with the HourlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyCost

`func (o *AddBaremetalHost200ResponseServer) SetHourlyCost(v float32)`

SetHourlyCost sets HourlyCost field to given value.

### HasHourlyCost

`func (o *AddBaremetalHost200ResponseServer) HasHourlyCost() bool`

HasHourlyCost returns a boolean if a field has been set.

### GetHourlyPrice

`func (o *AddBaremetalHost200ResponseServer) GetHourlyPrice() float32`

GetHourlyPrice returns the HourlyPrice field if non-nil, zero value otherwise.

### GetHourlyPriceOk

`func (o *AddBaremetalHost200ResponseServer) GetHourlyPriceOk() (*float32, bool)`

GetHourlyPriceOk returns a tuple with the HourlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyPrice

`func (o *AddBaremetalHost200ResponseServer) SetHourlyPrice(v float32)`

SetHourlyPrice sets HourlyPrice field to given value.

### HasHourlyPrice

`func (o *AddBaremetalHost200ResponseServer) HasHourlyPrice() bool`

HasHourlyPrice returns a boolean if a field has been set.

### GetSourceImage

`func (o *AddBaremetalHost200ResponseServer) GetSourceImage() AddBaremetalHost200ResponseServerSourceImage`

GetSourceImage returns the SourceImage field if non-nil, zero value otherwise.

### GetSourceImageOk

`func (o *AddBaremetalHost200ResponseServer) GetSourceImageOk() (*AddBaremetalHost200ResponseServerSourceImage, bool)`

GetSourceImageOk returns a tuple with the SourceImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceImage

`func (o *AddBaremetalHost200ResponseServer) SetSourceImage(v AddBaremetalHost200ResponseServerSourceImage)`

SetSourceImage sets SourceImage field to given value.

### HasSourceImage

`func (o *AddBaremetalHost200ResponseServer) HasSourceImage() bool`

HasSourceImage returns a boolean if a field has been set.

### GetServerOs

`func (o *AddBaremetalHost200ResponseServer) GetServerOs() AddBaremetalHost200ResponseServerServerOs`

GetServerOs returns the ServerOs field if non-nil, zero value otherwise.

### GetServerOsOk

`func (o *AddBaremetalHost200ResponseServer) GetServerOsOk() (*AddBaremetalHost200ResponseServerServerOs, bool)`

GetServerOsOk returns a tuple with the ServerOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerOs

`func (o *AddBaremetalHost200ResponseServer) SetServerOs(v AddBaremetalHost200ResponseServerServerOs)`

SetServerOs sets ServerOs field to given value.

### HasServerOs

`func (o *AddBaremetalHost200ResponseServer) HasServerOs() bool`

HasServerOs returns a boolean if a field has been set.

### GetVolumes

`func (o *AddBaremetalHost200ResponseServer) GetVolumes() []AddBaremetalHost200ResponseServerVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *AddBaremetalHost200ResponseServer) GetVolumesOk() (*[]AddBaremetalHost200ResponseServerVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *AddBaremetalHost200ResponseServer) SetVolumes(v []AddBaremetalHost200ResponseServerVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *AddBaremetalHost200ResponseServer) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetControllers

`func (o *AddBaremetalHost200ResponseServer) GetControllers() []AddBaremetalHost200ResponseServerControllersInner`

GetControllers returns the Controllers field if non-nil, zero value otherwise.

### GetControllersOk

`func (o *AddBaremetalHost200ResponseServer) GetControllersOk() (*[]AddBaremetalHost200ResponseServerControllersInner, bool)`

GetControllersOk returns a tuple with the Controllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllers

`func (o *AddBaremetalHost200ResponseServer) SetControllers(v []AddBaremetalHost200ResponseServerControllersInner)`

SetControllers sets Controllers field to given value.

### HasControllers

`func (o *AddBaremetalHost200ResponseServer) HasControllers() bool`

HasControllers returns a boolean if a field has been set.

### GetInterfaces

`func (o *AddBaremetalHost200ResponseServer) GetInterfaces() []AddBaremetalHost200ResponseServerInterfacesInner`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *AddBaremetalHost200ResponseServer) GetInterfacesOk() (*[]AddBaremetalHost200ResponseServerInterfacesInner, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *AddBaremetalHost200ResponseServer) SetInterfaces(v []AddBaremetalHost200ResponseServerInterfacesInner)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *AddBaremetalHost200ResponseServer) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetLabels

`func (o *AddBaremetalHost200ResponseServer) GetLabels() []map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddBaremetalHost200ResponseServer) GetLabelsOk() (*[]map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddBaremetalHost200ResponseServer) SetLabels(v []map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddBaremetalHost200ResponseServer) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *AddBaremetalHost200ResponseServer) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *AddBaremetalHost200ResponseServer) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetTags

`func (o *AddBaremetalHost200ResponseServer) GetTags() []map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AddBaremetalHost200ResponseServer) GetTagsOk() (*[]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AddBaremetalHost200ResponseServer) SetTags(v []map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *AddBaremetalHost200ResponseServer) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *AddBaremetalHost200ResponseServer) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *AddBaremetalHost200ResponseServer) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetEnabled

`func (o *AddBaremetalHost200ResponseServer) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddBaremetalHost200ResponseServer) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddBaremetalHost200ResponseServer) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddBaremetalHost200ResponseServer) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTagCompliant

`func (o *AddBaremetalHost200ResponseServer) GetTagCompliant() string`

GetTagCompliant returns the TagCompliant field if non-nil, zero value otherwise.

### GetTagCompliantOk

`func (o *AddBaremetalHost200ResponseServer) GetTagCompliantOk() (*string, bool)`

GetTagCompliantOk returns a tuple with the TagCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagCompliant

`func (o *AddBaremetalHost200ResponseServer) SetTagCompliant(v string)`

SetTagCompliant sets TagCompliant field to given value.

### HasTagCompliant

`func (o *AddBaremetalHost200ResponseServer) HasTagCompliant() bool`

HasTagCompliant returns a boolean if a field has been set.

### SetTagCompliantNil

`func (o *AddBaremetalHost200ResponseServer) SetTagCompliantNil(b bool)`

 SetTagCompliantNil sets the value for TagCompliant to be an explicit nil

### UnsetTagCompliant
`func (o *AddBaremetalHost200ResponseServer) UnsetTagCompliant()`

UnsetTagCompliant ensures that no value is present for TagCompliant, not even an explicit nil
### GetContainers

`func (o *AddBaremetalHost200ResponseServer) GetContainers() []int64`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *AddBaremetalHost200ResponseServer) GetContainersOk() (*[]int64, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *AddBaremetalHost200ResponseServer) SetContainers(v []int64)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *AddBaremetalHost200ResponseServer) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetConfig

`func (o *AddBaremetalHost200ResponseServer) GetConfig() AddBaremetalHost200ResponseServerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddBaremetalHost200ResponseServer) GetConfigOk() (*AddBaremetalHost200ResponseServerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddBaremetalHost200ResponseServer) SetConfig(v AddBaremetalHost200ResponseServerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddBaremetalHost200ResponseServer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetInstance

`func (o *AddBaremetalHost200ResponseServer) GetInstance() AddBaremetalHost200ResponseServerInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *AddBaremetalHost200ResponseServer) GetInstanceOk() (*AddBaremetalHost200ResponseServerInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *AddBaremetalHost200ResponseServer) SetInstance(v AddBaremetalHost200ResponseServerInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *AddBaremetalHost200ResponseServer) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetGuestConsolePreferred

`func (o *AddBaremetalHost200ResponseServer) GetGuestConsolePreferred() bool`

GetGuestConsolePreferred returns the GuestConsolePreferred field if non-nil, zero value otherwise.

### GetGuestConsolePreferredOk

`func (o *AddBaremetalHost200ResponseServer) GetGuestConsolePreferredOk() (*bool, bool)`

GetGuestConsolePreferredOk returns a tuple with the GuestConsolePreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsolePreferred

`func (o *AddBaremetalHost200ResponseServer) SetGuestConsolePreferred(v bool)`

SetGuestConsolePreferred sets GuestConsolePreferred field to given value.

### HasGuestConsolePreferred

`func (o *AddBaremetalHost200ResponseServer) HasGuestConsolePreferred() bool`

HasGuestConsolePreferred returns a boolean if a field has been set.

### GetGuestConsoleType

`func (o *AddBaremetalHost200ResponseServer) GetGuestConsoleType() string`

GetGuestConsoleType returns the GuestConsoleType field if non-nil, zero value otherwise.

### GetGuestConsoleTypeOk

`func (o *AddBaremetalHost200ResponseServer) GetGuestConsoleTypeOk() (*string, bool)`

GetGuestConsoleTypeOk returns a tuple with the GuestConsoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleType

`func (o *AddBaremetalHost200ResponseServer) SetGuestConsoleType(v string)`

SetGuestConsoleType sets GuestConsoleType field to given value.

### HasGuestConsoleType

`func (o *AddBaremetalHost200ResponseServer) HasGuestConsoleType() bool`

HasGuestConsoleType returns a boolean if a field has been set.

### SetGuestConsoleTypeNil

`func (o *AddBaremetalHost200ResponseServer) SetGuestConsoleTypeNil(b bool)`

 SetGuestConsoleTypeNil sets the value for GuestConsoleType to be an explicit nil

### UnsetGuestConsoleType
`func (o *AddBaremetalHost200ResponseServer) UnsetGuestConsoleType()`

UnsetGuestConsoleType ensures that no value is present for GuestConsoleType, not even an explicit nil
### GetGuestConsoleUsername

`func (o *AddBaremetalHost200ResponseServer) GetGuestConsoleUsername() string`

GetGuestConsoleUsername returns the GuestConsoleUsername field if non-nil, zero value otherwise.

### GetGuestConsoleUsernameOk

`func (o *AddBaremetalHost200ResponseServer) GetGuestConsoleUsernameOk() (*string, bool)`

GetGuestConsoleUsernameOk returns a tuple with the GuestConsoleUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleUsername

`func (o *AddBaremetalHost200ResponseServer) SetGuestConsoleUsername(v string)`

SetGuestConsoleUsername sets GuestConsoleUsername field to given value.

### HasGuestConsoleUsername

`func (o *AddBaremetalHost200ResponseServer) HasGuestConsoleUsername() bool`

HasGuestConsoleUsername returns a boolean if a field has been set.

### SetGuestConsoleUsernameNil

`func (o *AddBaremetalHost200ResponseServer) SetGuestConsoleUsernameNil(b bool)`

 SetGuestConsoleUsernameNil sets the value for GuestConsoleUsername to be an explicit nil

### UnsetGuestConsoleUsername
`func (o *AddBaremetalHost200ResponseServer) UnsetGuestConsoleUsername()`

UnsetGuestConsoleUsername ensures that no value is present for GuestConsoleUsername, not even an explicit nil
### GetGuestConsolePassword

`func (o *AddBaremetalHost200ResponseServer) GetGuestConsolePassword() string`

GetGuestConsolePassword returns the GuestConsolePassword field if non-nil, zero value otherwise.

### GetGuestConsolePasswordOk

`func (o *AddBaremetalHost200ResponseServer) GetGuestConsolePasswordOk() (*string, bool)`

GetGuestConsolePasswordOk returns a tuple with the GuestConsolePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsolePassword

`func (o *AddBaremetalHost200ResponseServer) SetGuestConsolePassword(v string)`

SetGuestConsolePassword sets GuestConsolePassword field to given value.

### HasGuestConsolePassword

`func (o *AddBaremetalHost200ResponseServer) HasGuestConsolePassword() bool`

HasGuestConsolePassword returns a boolean if a field has been set.

### SetGuestConsolePasswordNil

`func (o *AddBaremetalHost200ResponseServer) SetGuestConsolePasswordNil(b bool)`

 SetGuestConsolePasswordNil sets the value for GuestConsolePassword to be an explicit nil

### UnsetGuestConsolePassword
`func (o *AddBaremetalHost200ResponseServer) UnsetGuestConsolePassword()`

UnsetGuestConsolePassword ensures that no value is present for GuestConsolePassword, not even an explicit nil
### GetGuestConsolePasswordHash

`func (o *AddBaremetalHost200ResponseServer) GetGuestConsolePasswordHash() string`

GetGuestConsolePasswordHash returns the GuestConsolePasswordHash field if non-nil, zero value otherwise.

### GetGuestConsolePasswordHashOk

`func (o *AddBaremetalHost200ResponseServer) GetGuestConsolePasswordHashOk() (*string, bool)`

GetGuestConsolePasswordHashOk returns a tuple with the GuestConsolePasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsolePasswordHash

`func (o *AddBaremetalHost200ResponseServer) SetGuestConsolePasswordHash(v string)`

SetGuestConsolePasswordHash sets GuestConsolePasswordHash field to given value.

### HasGuestConsolePasswordHash

`func (o *AddBaremetalHost200ResponseServer) HasGuestConsolePasswordHash() bool`

HasGuestConsolePasswordHash returns a boolean if a field has been set.

### SetGuestConsolePasswordHashNil

`func (o *AddBaremetalHost200ResponseServer) SetGuestConsolePasswordHashNil(b bool)`

 SetGuestConsolePasswordHashNil sets the value for GuestConsolePasswordHash to be an explicit nil

### UnsetGuestConsolePasswordHash
`func (o *AddBaremetalHost200ResponseServer) UnsetGuestConsolePasswordHash()`

UnsetGuestConsolePasswordHash ensures that no value is present for GuestConsolePasswordHash, not even an explicit nil
### GetGuestConsolePort

`func (o *AddBaremetalHost200ResponseServer) GetGuestConsolePort() string`

GetGuestConsolePort returns the GuestConsolePort field if non-nil, zero value otherwise.

### GetGuestConsolePortOk

`func (o *AddBaremetalHost200ResponseServer) GetGuestConsolePortOk() (*string, bool)`

GetGuestConsolePortOk returns a tuple with the GuestConsolePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsolePort

`func (o *AddBaremetalHost200ResponseServer) SetGuestConsolePort(v string)`

SetGuestConsolePort sets GuestConsolePort field to given value.

### HasGuestConsolePort

`func (o *AddBaremetalHost200ResponseServer) HasGuestConsolePort() bool`

HasGuestConsolePort returns a boolean if a field has been set.

### SetGuestConsolePortNil

`func (o *AddBaremetalHost200ResponseServer) SetGuestConsolePortNil(b bool)`

 SetGuestConsolePortNil sets the value for GuestConsolePort to be an explicit nil

### UnsetGuestConsolePort
`func (o *AddBaremetalHost200ResponseServer) UnsetGuestConsolePort()`

UnsetGuestConsolePort ensures that no value is present for GuestConsolePort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


