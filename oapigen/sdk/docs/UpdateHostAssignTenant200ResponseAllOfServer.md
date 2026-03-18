# UpdateHostAssignTenant200ResponseAllOfServer

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

### NewUpdateHostAssignTenant200ResponseAllOfServer

`func NewUpdateHostAssignTenant200ResponseAllOfServer() *UpdateHostAssignTenant200ResponseAllOfServer`

NewUpdateHostAssignTenant200ResponseAllOfServer instantiates a new UpdateHostAssignTenant200ResponseAllOfServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateHostAssignTenant200ResponseAllOfServerWithDefaults

`func NewUpdateHostAssignTenant200ResponseAllOfServerWithDefaults() *UpdateHostAssignTenant200ResponseAllOfServer`

NewUpdateHostAssignTenant200ResponseAllOfServerWithDefaults instantiates a new UpdateHostAssignTenant200ResponseAllOfServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetExternalId

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *UpdateHostAssignTenant200ResponseAllOfServer) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetInternalId

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *UpdateHostAssignTenant200ResponseAllOfServer) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalUniqueId

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetExternalUniqueId() string`

GetExternalUniqueId returns the ExternalUniqueId field if non-nil, zero value otherwise.

### GetExternalUniqueIdOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetExternalUniqueIdOk() (*string, bool)`

GetExternalUniqueIdOk returns a tuple with the ExternalUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUniqueId

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetExternalUniqueId(v string)`

SetExternalUniqueId sets ExternalUniqueId field to given value.

### HasExternalUniqueId

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasExternalUniqueId() bool`

HasExternalUniqueId returns a boolean if a field has been set.

### SetExternalUniqueIdNil

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetExternalUniqueIdNil(b bool)`

 SetExternalUniqueIdNil sets the value for ExternalUniqueId to be an explicit nil

### UnsetExternalUniqueId
`func (o *UpdateHostAssignTenant200ResponseAllOfServer) UnsetExternalUniqueId()`

UnsetExternalUniqueId ensures that no value is present for ExternalUniqueId, not even an explicit nil
### GetName

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalName

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetExternalName() string`

GetExternalName returns the ExternalName field if non-nil, zero value otherwise.

### GetExternalNameOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetExternalNameOk() (*string, bool)`

GetExternalNameOk returns a tuple with the ExternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalName

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetExternalName(v string)`

SetExternalName sets ExternalName field to given value.

### HasExternalName

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasExternalName() bool`

HasExternalName returns a boolean if a field has been set.

### GetHostname

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetParentServer

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetParentServer() AddBaremetalHost200ResponseServerParentServer`

GetParentServer returns the ParentServer field if non-nil, zero value otherwise.

### GetParentServerOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetParentServerOk() (*AddBaremetalHost200ResponseServerParentServer, bool)`

GetParentServerOk returns a tuple with the ParentServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentServer

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetParentServer(v AddBaremetalHost200ResponseServerParentServer)`

SetParentServer sets ParentServer field to given value.

### HasParentServer

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasParentServer() bool`

HasParentServer returns a boolean if a field has been set.

### GetAccountId

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccount

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetAccount() AddBaremetalHost200ResponseServerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetAccountOk() (*AddBaremetalHost200ResponseServerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetAccount(v AddBaremetalHost200ResponseServerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetOwner() AddBaremetalHost200ResponseServerOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetOwnerOk() (*AddBaremetalHost200ResponseServerOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetOwner(v AddBaremetalHost200ResponseServerOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetZone

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetZone() AddBaremetalHost200ResponseServerZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetZoneOk() (*AddBaremetalHost200ResponseServerZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetZone(v AddBaremetalHost200ResponseServerZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetPlan

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetPlan() AddBaremetalHost200ResponseServerPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetPlanOk() (*AddBaremetalHost200ResponseServerPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetPlan(v AddBaremetalHost200ResponseServerPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetComputeServerType

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetComputeServerType() AddBaremetalHost200ResponseServerComputeServerType`

GetComputeServerType returns the ComputeServerType field if non-nil, zero value otherwise.

### GetComputeServerTypeOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetComputeServerTypeOk() (*AddBaremetalHost200ResponseServerComputeServerType, bool)`

GetComputeServerTypeOk returns a tuple with the ComputeServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServerType

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetComputeServerType(v AddBaremetalHost200ResponseServerComputeServerType)`

SetComputeServerType sets ComputeServerType field to given value.

### HasComputeServerType

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasComputeServerType() bool`

HasComputeServerType returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateHostAssignTenant200ResponseAllOfServer) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetZoneId

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetSiteId

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetResourcePoolId() int64`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetResourcePoolIdOk() (*int64, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetResourcePoolId(v int64)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### SetResourcePoolIdNil

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetResourcePoolIdNil(b bool)`

 SetResourcePoolIdNil sets the value for ResourcePoolId to be an explicit nil

### UnsetResourcePoolId
`func (o *UpdateHostAssignTenant200ResponseAllOfServer) UnsetResourcePoolId()`

UnsetResourcePoolId ensures that no value is present for ResourcePoolId, not even an explicit nil
### GetFolderId

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetFolderId() int64`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetFolderIdOk() (*int64, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetFolderId(v int64)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *UpdateHostAssignTenant200ResponseAllOfServer) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetSshHost

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetSshHost() string`

GetSshHost returns the SshHost field if non-nil, zero value otherwise.

### GetSshHostOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetSshHostOk() (*string, bool)`

GetSshHostOk returns a tuple with the SshHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHost

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetSshHost(v string)`

SetSshHost sets SshHost field to given value.

### HasSshHost

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasSshHost() bool`

HasSshHost returns a boolean if a field has been set.

### SetSshHostNil

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetSshHostNil(b bool)`

 SetSshHostNil sets the value for SshHost to be an explicit nil

### UnsetSshHost
`func (o *UpdateHostAssignTenant200ResponseAllOfServer) UnsetSshHost()`

UnsetSshHost ensures that no value is present for SshHost, not even an explicit nil
### GetSshPort

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetExternalIp

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### SetExternalIpNil

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetExternalIpNil(b bool)`

 SetExternalIpNil sets the value for ExternalIp to be an explicit nil

### UnsetExternalIp
`func (o *UpdateHostAssignTenant200ResponseAllOfServer) UnsetExternalIp()`

UnsetExternalIp ensures that no value is present for ExternalIp, not even an explicit nil
### GetInternalIp

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### SetInternalIpNil

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetInternalIpNil(b bool)`

 SetInternalIpNil sets the value for InternalIp to be an explicit nil

### UnsetInternalIp
`func (o *UpdateHostAssignTenant200ResponseAllOfServer) UnsetInternalIp()`

UnsetInternalIp ensures that no value is present for InternalIp, not even an explicit nil
### GetVolumeId

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.

### SetVolumeIdNil

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetVolumeIdNil(b bool)`

 SetVolumeIdNil sets the value for VolumeId to be an explicit nil

### UnsetVolumeId
`func (o *UpdateHostAssignTenant200ResponseAllOfServer) UnsetVolumeId()`

UnsetVolumeId ensures that no value is present for VolumeId, not even an explicit nil
### GetPlatform

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *UpdateHostAssignTenant200ResponseAllOfServer) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetPlatformVersion

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetPlatformVersion() string`

GetPlatformVersion returns the PlatformVersion field if non-nil, zero value otherwise.

### GetPlatformVersionOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetPlatformVersionOk() (*string, bool)`

GetPlatformVersionOk returns a tuple with the PlatformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformVersion

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetPlatformVersion(v string)`

SetPlatformVersion sets PlatformVersion field to given value.

### HasPlatformVersion

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasPlatformVersion() bool`

HasPlatformVersion returns a boolean if a field has been set.

### SetPlatformVersionNil

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetPlatformVersionNil(b bool)`

 SetPlatformVersionNil sets the value for PlatformVersion to be an explicit nil

### UnsetPlatformVersion
`func (o *UpdateHostAssignTenant200ResponseAllOfServer) UnsetPlatformVersion()`

UnsetPlatformVersion ensures that no value is present for PlatformVersion, not even an explicit nil
### GetSshUsername

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### SetSshUsernameNil

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetSshUsernameNil(b bool)`

 SetSshUsernameNil sets the value for SshUsername to be an explicit nil

### UnsetSshUsername
`func (o *UpdateHostAssignTenant200ResponseAllOfServer) UnsetSshUsername()`

UnsetSshUsername ensures that no value is present for SshUsername, not even an explicit nil
### GetSshPassword

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### SetSshPasswordNil

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetSshPasswordNil(b bool)`

 SetSshPasswordNil sets the value for SshPassword to be an explicit nil

### UnsetSshPassword
`func (o *UpdateHostAssignTenant200ResponseAllOfServer) UnsetSshPassword()`

UnsetSshPassword ensures that no value is present for SshPassword, not even an explicit nil
### GetSshPasswordHash

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetSshPasswordHash() string`

GetSshPasswordHash returns the SshPasswordHash field if non-nil, zero value otherwise.

### GetSshPasswordHashOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetSshPasswordHashOk() (*string, bool)`

GetSshPasswordHashOk returns a tuple with the SshPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPasswordHash

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetSshPasswordHash(v string)`

SetSshPasswordHash sets SshPasswordHash field to given value.

### HasSshPasswordHash

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasSshPasswordHash() bool`

HasSshPasswordHash returns a boolean if a field has been set.

### SetSshPasswordHashNil

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetSshPasswordHashNil(b bool)`

 SetSshPasswordHashNil sets the value for SshPasswordHash to be an explicit nil

### UnsetSshPasswordHash
`func (o *UpdateHostAssignTenant200ResponseAllOfServer) UnsetSshPasswordHash()`

UnsetSshPasswordHash ensures that no value is present for SshPasswordHash, not even an explicit nil
### GetSshKeyPair

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetSshKeyPair() AddBaremetalHost200ResponseServerSshKeyPair`

GetSshKeyPair returns the SshKeyPair field if non-nil, zero value otherwise.

### GetSshKeyPairOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetSshKeyPairOk() (*AddBaremetalHost200ResponseServerSshKeyPair, bool)`

GetSshKeyPairOk returns a tuple with the SshKeyPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyPair

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetSshKeyPair(v AddBaremetalHost200ResponseServerSshKeyPair)`

SetSshKeyPair sets SshKeyPair field to given value.

### HasSshKeyPair

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasSshKeyPair() bool`

HasSshKeyPair returns a boolean if a field has been set.

### GetOsDevice

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetOsDevice() string`

GetOsDevice returns the OsDevice field if non-nil, zero value otherwise.

### GetOsDeviceOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetOsDeviceOk() (*string, bool)`

GetOsDeviceOk returns a tuple with the OsDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDevice

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetOsDevice(v string)`

SetOsDevice sets OsDevice field to given value.

### HasOsDevice

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasOsDevice() bool`

HasOsDevice returns a boolean if a field has been set.

### GetOsType

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetDataDevice

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetDataDevice() string`

GetDataDevice returns the DataDevice field if non-nil, zero value otherwise.

### GetDataDeviceOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetDataDeviceOk() (*string, bool)`

GetDataDeviceOk returns a tuple with the DataDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDevice

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetDataDevice(v string)`

SetDataDevice sets DataDevice field to given value.

### HasDataDevice

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasDataDevice() bool`

HasDataDevice returns a boolean if a field has been set.

### GetLvmEnabled

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetLvmEnabled() bool`

GetLvmEnabled returns the LvmEnabled field if non-nil, zero value otherwise.

### GetLvmEnabledOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetLvmEnabledOk() (*bool, bool)`

GetLvmEnabledOk returns a tuple with the LvmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLvmEnabled

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetLvmEnabled(v bool)`

SetLvmEnabled sets LvmEnabled field to given value.

### HasLvmEnabled

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasLvmEnabled() bool`

HasLvmEnabled returns a boolean if a field has been set.

### GetApiKey

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetSoftwareRaid

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetSoftwareRaid() bool`

GetSoftwareRaid returns the SoftwareRaid field if non-nil, zero value otherwise.

### GetSoftwareRaidOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetSoftwareRaidOk() (*bool, bool)`

GetSoftwareRaidOk returns a tuple with the SoftwareRaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareRaid

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetSoftwareRaid(v bool)`

SetSoftwareRaid sets SoftwareRaid field to given value.

### HasSoftwareRaid

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasSoftwareRaid() bool`

HasSoftwareRaid returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetStats

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetStats() AddBaremetalHost200ResponseServerStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetStatsOk() (*AddBaremetalHost200ResponseServerStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetStats(v AddBaremetalHost200ResponseServerStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *UpdateHostAssignTenant200ResponseAllOfServer) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetErrorMessage

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *UpdateHostAssignTenant200ResponseAllOfServer) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetStatusDate

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### SetStatusDateNil

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetStatusDateNil(b bool)`

 SetStatusDateNil sets the value for StatusDate to be an explicit nil

### UnsetStatusDate
`func (o *UpdateHostAssignTenant200ResponseAllOfServer) UnsetStatusDate()`

UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
### GetStatusPercent

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetStatusPercent() string`

GetStatusPercent returns the StatusPercent field if non-nil, zero value otherwise.

### GetStatusPercentOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetStatusPercentOk() (*string, bool)`

GetStatusPercentOk returns a tuple with the StatusPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusPercent

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetStatusPercent(v string)`

SetStatusPercent sets StatusPercent field to given value.

### HasStatusPercent

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasStatusPercent() bool`

HasStatusPercent returns a boolean if a field has been set.

### SetStatusPercentNil

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetStatusPercentNil(b bool)`

 SetStatusPercentNil sets the value for StatusPercent to be an explicit nil

### UnsetStatusPercent
`func (o *UpdateHostAssignTenant200ResponseAllOfServer) UnsetStatusPercent()`

UnsetStatusPercent ensures that no value is present for StatusPercent, not even an explicit nil
### GetStatusEta

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetStatusEta() string`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetStatusEtaOk() (*string, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetStatusEta(v string)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### SetStatusEtaNil

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetStatusEtaNil(b bool)`

 SetStatusEtaNil sets the value for StatusEta to be an explicit nil

### UnsetStatusEta
`func (o *UpdateHostAssignTenant200ResponseAllOfServer) UnsetStatusEta()`

UnsetStatusEta ensures that no value is present for StatusEta, not even an explicit nil
### GetPowerState

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetPowerState() string`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetPowerStateOk() (*string, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetPowerState(v string)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetAgentInstalled

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetAgentInstalled() bool`

GetAgentInstalled returns the AgentInstalled field if non-nil, zero value otherwise.

### GetAgentInstalledOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetAgentInstalledOk() (*bool, bool)`

GetAgentInstalledOk returns a tuple with the AgentInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentInstalled

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetAgentInstalled(v bool)`

SetAgentInstalled sets AgentInstalled field to given value.

### HasAgentInstalled

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasAgentInstalled() bool`

HasAgentInstalled returns a boolean if a field has been set.

### GetLastAgentUpdate

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetLastAgentUpdate() string`

GetLastAgentUpdate returns the LastAgentUpdate field if non-nil, zero value otherwise.

### GetLastAgentUpdateOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetLastAgentUpdateOk() (*string, bool)`

GetLastAgentUpdateOk returns a tuple with the LastAgentUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAgentUpdate

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetLastAgentUpdate(v string)`

SetLastAgentUpdate sets LastAgentUpdate field to given value.

### HasLastAgentUpdate

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasLastAgentUpdate() bool`

HasLastAgentUpdate returns a boolean if a field has been set.

### SetLastAgentUpdateNil

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetLastAgentUpdateNil(b bool)`

 SetLastAgentUpdateNil sets the value for LastAgentUpdate to be an explicit nil

### UnsetLastAgentUpdate
`func (o *UpdateHostAssignTenant200ResponseAllOfServer) UnsetLastAgentUpdate()`

UnsetLastAgentUpdate ensures that no value is present for LastAgentUpdate, not even an explicit nil
### GetAgentVersion

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### SetAgentVersionNil

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetAgentVersionNil(b bool)`

 SetAgentVersionNil sets the value for AgentVersion to be an explicit nil

### UnsetAgentVersion
`func (o *UpdateHostAssignTenant200ResponseAllOfServer) UnsetAgentVersion()`

UnsetAgentVersion ensures that no value is present for AgentVersion, not even an explicit nil
### GetMaxCores

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetMaxCores() int64`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetMaxCoresOk() (*int64, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetMaxCores(v int64)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetCoresPerSocket() int64`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetCoresPerSocketOk() (*int64, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetCoresPerSocket(v int64)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### SetCoresPerSocketNil

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetCoresPerSocketNil(b bool)`

 SetCoresPerSocketNil sets the value for CoresPerSocket to be an explicit nil

### UnsetCoresPerSocket
`func (o *UpdateHostAssignTenant200ResponseAllOfServer) UnsetCoresPerSocket()`

UnsetCoresPerSocket ensures that no value is present for CoresPerSocket, not even an explicit nil
### GetMaxMemory

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxStorage

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMaxCpu

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetMaxCpu() int64`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetMaxCpuOk() (*int64, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetMaxCpu(v int64)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### SetMaxCpuNil

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetMaxCpuNil(b bool)`

 SetMaxCpuNil sets the value for MaxCpu to be an explicit nil

### UnsetMaxCpu
`func (o *UpdateHostAssignTenant200ResponseAllOfServer) UnsetMaxCpu()`

UnsetMaxCpu ensures that no value is present for MaxCpu, not even an explicit nil
### GetMaxGpus

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetMaxGpus() int64`

GetMaxGpus returns the MaxGpus field if non-nil, zero value otherwise.

### GetMaxGpusOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetMaxGpusOk() (*int64, bool)`

GetMaxGpusOk returns a tuple with the MaxGpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGpus

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetMaxGpus(v int64)`

SetMaxGpus sets MaxGpus field to given value.

### HasMaxGpus

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasMaxGpus() bool`

HasMaxGpus returns a boolean if a field has been set.

### SetMaxGpusNil

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetMaxGpusNil(b bool)`

 SetMaxGpusNil sets the value for MaxGpus to be an explicit nil

### UnsetMaxGpus
`func (o *UpdateHostAssignTenant200ResponseAllOfServer) UnsetMaxGpus()`

UnsetMaxGpus ensures that no value is present for MaxGpus, not even an explicit nil
### GetManageInternalFirewall

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetManageInternalFirewall() bool`

GetManageInternalFirewall returns the ManageInternalFirewall field if non-nil, zero value otherwise.

### GetManageInternalFirewallOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetManageInternalFirewallOk() (*bool, bool)`

GetManageInternalFirewallOk returns a tuple with the ManageInternalFirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageInternalFirewall

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetManageInternalFirewall(v bool)`

SetManageInternalFirewall sets ManageInternalFirewall field to given value.

### HasManageInternalFirewall

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasManageInternalFirewall() bool`

HasManageInternalFirewall returns a boolean if a field has been set.

### GetEnableLogs

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetEnableLogs() bool`

GetEnableLogs returns the EnableLogs field if non-nil, zero value otherwise.

### GetEnableLogsOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetEnableLogsOk() (*bool, bool)`

GetEnableLogsOk returns a tuple with the EnableLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLogs

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetEnableLogs(v bool)`

SetEnableLogs sets EnableLogs field to given value.

### HasEnableLogs

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasEnableLogs() bool`

HasEnableLogs returns a boolean if a field has been set.

### GetHourlyCost

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetHourlyCost() float32`

GetHourlyCost returns the HourlyCost field if non-nil, zero value otherwise.

### GetHourlyCostOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetHourlyCostOk() (*float32, bool)`

GetHourlyCostOk returns a tuple with the HourlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyCost

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetHourlyCost(v float32)`

SetHourlyCost sets HourlyCost field to given value.

### HasHourlyCost

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasHourlyCost() bool`

HasHourlyCost returns a boolean if a field has been set.

### GetHourlyPrice

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetHourlyPrice() float32`

GetHourlyPrice returns the HourlyPrice field if non-nil, zero value otherwise.

### GetHourlyPriceOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetHourlyPriceOk() (*float32, bool)`

GetHourlyPriceOk returns a tuple with the HourlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyPrice

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetHourlyPrice(v float32)`

SetHourlyPrice sets HourlyPrice field to given value.

### HasHourlyPrice

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasHourlyPrice() bool`

HasHourlyPrice returns a boolean if a field has been set.

### GetSourceImage

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetSourceImage() AddBaremetalHost200ResponseServerSourceImage`

GetSourceImage returns the SourceImage field if non-nil, zero value otherwise.

### GetSourceImageOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetSourceImageOk() (*AddBaremetalHost200ResponseServerSourceImage, bool)`

GetSourceImageOk returns a tuple with the SourceImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceImage

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetSourceImage(v AddBaremetalHost200ResponseServerSourceImage)`

SetSourceImage sets SourceImage field to given value.

### HasSourceImage

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasSourceImage() bool`

HasSourceImage returns a boolean if a field has been set.

### GetServerOs

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetServerOs() AddBaremetalHost200ResponseServerServerOs`

GetServerOs returns the ServerOs field if non-nil, zero value otherwise.

### GetServerOsOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetServerOsOk() (*AddBaremetalHost200ResponseServerServerOs, bool)`

GetServerOsOk returns a tuple with the ServerOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerOs

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetServerOs(v AddBaremetalHost200ResponseServerServerOs)`

SetServerOs sets ServerOs field to given value.

### HasServerOs

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasServerOs() bool`

HasServerOs returns a boolean if a field has been set.

### GetVolumes

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetVolumes() []AddBaremetalHost200ResponseServerVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetVolumesOk() (*[]AddBaremetalHost200ResponseServerVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetVolumes(v []AddBaremetalHost200ResponseServerVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetControllers

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetControllers() []AddBaremetalHost200ResponseServerControllersInner`

GetControllers returns the Controllers field if non-nil, zero value otherwise.

### GetControllersOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetControllersOk() (*[]AddBaremetalHost200ResponseServerControllersInner, bool)`

GetControllersOk returns a tuple with the Controllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllers

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetControllers(v []AddBaremetalHost200ResponseServerControllersInner)`

SetControllers sets Controllers field to given value.

### HasControllers

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasControllers() bool`

HasControllers returns a boolean if a field has been set.

### GetInterfaces

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetInterfaces() []AddBaremetalHost200ResponseServerInterfacesInner`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetInterfacesOk() (*[]AddBaremetalHost200ResponseServerInterfacesInner, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetInterfaces(v []AddBaremetalHost200ResponseServerInterfacesInner)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetLabels() []map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetLabelsOk() (*[]map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetLabels(v []map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *UpdateHostAssignTenant200ResponseAllOfServer) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetTags

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetTags() []map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetTagsOk() (*[]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetTags(v []map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *UpdateHostAssignTenant200ResponseAllOfServer) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetEnabled

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTagCompliant

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetTagCompliant() string`

GetTagCompliant returns the TagCompliant field if non-nil, zero value otherwise.

### GetTagCompliantOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetTagCompliantOk() (*string, bool)`

GetTagCompliantOk returns a tuple with the TagCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagCompliant

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetTagCompliant(v string)`

SetTagCompliant sets TagCompliant field to given value.

### HasTagCompliant

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasTagCompliant() bool`

HasTagCompliant returns a boolean if a field has been set.

### SetTagCompliantNil

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetTagCompliantNil(b bool)`

 SetTagCompliantNil sets the value for TagCompliant to be an explicit nil

### UnsetTagCompliant
`func (o *UpdateHostAssignTenant200ResponseAllOfServer) UnsetTagCompliant()`

UnsetTagCompliant ensures that no value is present for TagCompliant, not even an explicit nil
### GetContainers

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetContainers() []int64`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetContainersOk() (*[]int64, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetContainers(v []int64)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetConfig() AddBaremetalHost200ResponseServerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetConfigOk() (*AddBaremetalHost200ResponseServerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetConfig(v AddBaremetalHost200ResponseServerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetInstance

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetInstance() AddBaremetalHost200ResponseServerInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetInstanceOk() (*AddBaremetalHost200ResponseServerInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetInstance(v AddBaremetalHost200ResponseServerInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetGuestConsolePreferred

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetGuestConsolePreferred() bool`

GetGuestConsolePreferred returns the GuestConsolePreferred field if non-nil, zero value otherwise.

### GetGuestConsolePreferredOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetGuestConsolePreferredOk() (*bool, bool)`

GetGuestConsolePreferredOk returns a tuple with the GuestConsolePreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsolePreferred

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetGuestConsolePreferred(v bool)`

SetGuestConsolePreferred sets GuestConsolePreferred field to given value.

### HasGuestConsolePreferred

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasGuestConsolePreferred() bool`

HasGuestConsolePreferred returns a boolean if a field has been set.

### GetGuestConsoleType

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetGuestConsoleType() string`

GetGuestConsoleType returns the GuestConsoleType field if non-nil, zero value otherwise.

### GetGuestConsoleTypeOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetGuestConsoleTypeOk() (*string, bool)`

GetGuestConsoleTypeOk returns a tuple with the GuestConsoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleType

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetGuestConsoleType(v string)`

SetGuestConsoleType sets GuestConsoleType field to given value.

### HasGuestConsoleType

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasGuestConsoleType() bool`

HasGuestConsoleType returns a boolean if a field has been set.

### SetGuestConsoleTypeNil

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetGuestConsoleTypeNil(b bool)`

 SetGuestConsoleTypeNil sets the value for GuestConsoleType to be an explicit nil

### UnsetGuestConsoleType
`func (o *UpdateHostAssignTenant200ResponseAllOfServer) UnsetGuestConsoleType()`

UnsetGuestConsoleType ensures that no value is present for GuestConsoleType, not even an explicit nil
### GetGuestConsoleUsername

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetGuestConsoleUsername() string`

GetGuestConsoleUsername returns the GuestConsoleUsername field if non-nil, zero value otherwise.

### GetGuestConsoleUsernameOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetGuestConsoleUsernameOk() (*string, bool)`

GetGuestConsoleUsernameOk returns a tuple with the GuestConsoleUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleUsername

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetGuestConsoleUsername(v string)`

SetGuestConsoleUsername sets GuestConsoleUsername field to given value.

### HasGuestConsoleUsername

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasGuestConsoleUsername() bool`

HasGuestConsoleUsername returns a boolean if a field has been set.

### SetGuestConsoleUsernameNil

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetGuestConsoleUsernameNil(b bool)`

 SetGuestConsoleUsernameNil sets the value for GuestConsoleUsername to be an explicit nil

### UnsetGuestConsoleUsername
`func (o *UpdateHostAssignTenant200ResponseAllOfServer) UnsetGuestConsoleUsername()`

UnsetGuestConsoleUsername ensures that no value is present for GuestConsoleUsername, not even an explicit nil
### GetGuestConsolePassword

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetGuestConsolePassword() string`

GetGuestConsolePassword returns the GuestConsolePassword field if non-nil, zero value otherwise.

### GetGuestConsolePasswordOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetGuestConsolePasswordOk() (*string, bool)`

GetGuestConsolePasswordOk returns a tuple with the GuestConsolePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsolePassword

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetGuestConsolePassword(v string)`

SetGuestConsolePassword sets GuestConsolePassword field to given value.

### HasGuestConsolePassword

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasGuestConsolePassword() bool`

HasGuestConsolePassword returns a boolean if a field has been set.

### SetGuestConsolePasswordNil

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetGuestConsolePasswordNil(b bool)`

 SetGuestConsolePasswordNil sets the value for GuestConsolePassword to be an explicit nil

### UnsetGuestConsolePassword
`func (o *UpdateHostAssignTenant200ResponseAllOfServer) UnsetGuestConsolePassword()`

UnsetGuestConsolePassword ensures that no value is present for GuestConsolePassword, not even an explicit nil
### GetGuestConsolePasswordHash

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetGuestConsolePasswordHash() string`

GetGuestConsolePasswordHash returns the GuestConsolePasswordHash field if non-nil, zero value otherwise.

### GetGuestConsolePasswordHashOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetGuestConsolePasswordHashOk() (*string, bool)`

GetGuestConsolePasswordHashOk returns a tuple with the GuestConsolePasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsolePasswordHash

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetGuestConsolePasswordHash(v string)`

SetGuestConsolePasswordHash sets GuestConsolePasswordHash field to given value.

### HasGuestConsolePasswordHash

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasGuestConsolePasswordHash() bool`

HasGuestConsolePasswordHash returns a boolean if a field has been set.

### SetGuestConsolePasswordHashNil

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetGuestConsolePasswordHashNil(b bool)`

 SetGuestConsolePasswordHashNil sets the value for GuestConsolePasswordHash to be an explicit nil

### UnsetGuestConsolePasswordHash
`func (o *UpdateHostAssignTenant200ResponseAllOfServer) UnsetGuestConsolePasswordHash()`

UnsetGuestConsolePasswordHash ensures that no value is present for GuestConsolePasswordHash, not even an explicit nil
### GetGuestConsolePort

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetGuestConsolePort() string`

GetGuestConsolePort returns the GuestConsolePort field if non-nil, zero value otherwise.

### GetGuestConsolePortOk

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) GetGuestConsolePortOk() (*string, bool)`

GetGuestConsolePortOk returns a tuple with the GuestConsolePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsolePort

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetGuestConsolePort(v string)`

SetGuestConsolePort sets GuestConsolePort field to given value.

### HasGuestConsolePort

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) HasGuestConsolePort() bool`

HasGuestConsolePort returns a boolean if a field has been set.

### SetGuestConsolePortNil

`func (o *UpdateHostAssignTenant200ResponseAllOfServer) SetGuestConsolePortNil(b bool)`

 SetGuestConsolePortNil sets the value for GuestConsolePort to be an explicit nil

### UnsetGuestConsolePort
`func (o *UpdateHostAssignTenant200ResponseAllOfServer) UnsetGuestConsolePort()`

UnsetGuestConsolePort ensures that no value is present for GuestConsolePort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


