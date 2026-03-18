# UpdateServerNetworkInterface200ResponseAllOfOneOfServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ZoneId** | Pointer to **int64** |  | [optional] 
**SiteId** | Pointer to **int64** |  | [optional] 
**SshHost** | Pointer to **string** |  | [optional] 
**SshPort** | Pointer to **int64** |  | [optional] 
**ExternalIp** | Pointer to **string** |  | [optional] 
**InternalIp** | Pointer to **string** |  | [optional] 
**VolumeId** | Pointer to **NullableString** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**PlatformVersion** | Pointer to **string** |  | [optional] 
**SshUsername** | Pointer to **string** |  | [optional] 
**SshPassword** | Pointer to **string** |  | [optional] 
**OsDevice** | Pointer to **string** |  | [optional] 
**DataDevice** | Pointer to **string** |  | [optional] 
**LvmEnabled** | Pointer to **bool** |  | [optional] 
**ApiKey** | Pointer to **string** |  | [optional] 
**SoftwareRaid** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to **string** |  | [optional] 
**CapacityInfo** | Pointer to [**UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo**](UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**LastStats** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ComputeServerType** | Pointer to [**UpdateServerNetworkInterface200ResponseAllOfOneOfServerComputeServerType**](UpdateServerNetworkInterface200ResponseAllOfOneOfServerComputeServerType.md) |  | [optional] 
**Interfaces** | Pointer to [**[]UpdateServerNetworkInterface200ResponseAllOfOneOfServerInterfacesInner**](UpdateServerNetworkInterface200ResponseAllOfOneOfServerInterfacesInner.md) |  | [optional] 
**Zone** | Pointer to [**UpdateServerNetworkInterface200ResponseAllOfOneOfServerZone**](UpdateServerNetworkInterface200ResponseAllOfOneOfServerZone.md) |  | [optional] 

## Methods

### NewUpdateServerNetworkInterface200ResponseAllOfOneOfServer

`func NewUpdateServerNetworkInterface200ResponseAllOfOneOfServer() *UpdateServerNetworkInterface200ResponseAllOfOneOfServer`

NewUpdateServerNetworkInterface200ResponseAllOfOneOfServer instantiates a new UpdateServerNetworkInterface200ResponseAllOfOneOfServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServerNetworkInterface200ResponseAllOfOneOfServerWithDefaults

`func NewUpdateServerNetworkInterface200ResponseAllOfOneOfServerWithDefaults() *UpdateServerNetworkInterface200ResponseAllOfOneOfServer`

NewUpdateServerNetworkInterface200ResponseAllOfOneOfServerWithDefaults instantiates a new UpdateServerNetworkInterface200ResponseAllOfOneOfServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExternalId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetAccountId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetName

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetZoneId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetSiteId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSshHost

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetSshHost() string`

GetSshHost returns the SshHost field if non-nil, zero value otherwise.

### GetSshHostOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetSshHostOk() (*string, bool)`

GetSshHostOk returns a tuple with the SshHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHost

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetSshHost(v string)`

SetSshHost sets SshHost field to given value.

### HasSshHost

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) HasSshHost() bool`

HasSshHost returns a boolean if a field has been set.

### GetSshPort

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetExternalIp

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### GetInternalIp

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### GetVolumeId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.

### SetVolumeIdNil

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetVolumeIdNil(b bool)`

 SetVolumeIdNil sets the value for VolumeId to be an explicit nil

### UnsetVolumeId
`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) UnsetVolumeId()`

UnsetVolumeId ensures that no value is present for VolumeId, not even an explicit nil
### GetPlatform

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPlatformVersion

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetPlatformVersion() string`

GetPlatformVersion returns the PlatformVersion field if non-nil, zero value otherwise.

### GetPlatformVersionOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetPlatformVersionOk() (*string, bool)`

GetPlatformVersionOk returns a tuple with the PlatformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformVersion

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetPlatformVersion(v string)`

SetPlatformVersion sets PlatformVersion field to given value.

### HasPlatformVersion

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) HasPlatformVersion() bool`

HasPlatformVersion returns a boolean if a field has been set.

### GetSshUsername

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetOsDevice

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetOsDevice() string`

GetOsDevice returns the OsDevice field if non-nil, zero value otherwise.

### GetOsDeviceOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetOsDeviceOk() (*string, bool)`

GetOsDeviceOk returns a tuple with the OsDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDevice

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetOsDevice(v string)`

SetOsDevice sets OsDevice field to given value.

### HasOsDevice

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) HasOsDevice() bool`

HasOsDevice returns a boolean if a field has been set.

### GetDataDevice

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetDataDevice() string`

GetDataDevice returns the DataDevice field if non-nil, zero value otherwise.

### GetDataDeviceOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetDataDeviceOk() (*string, bool)`

GetDataDeviceOk returns a tuple with the DataDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDevice

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetDataDevice(v string)`

SetDataDevice sets DataDevice field to given value.

### HasDataDevice

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) HasDataDevice() bool`

HasDataDevice returns a boolean if a field has been set.

### GetLvmEnabled

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetLvmEnabled() bool`

GetLvmEnabled returns the LvmEnabled field if non-nil, zero value otherwise.

### GetLvmEnabledOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetLvmEnabledOk() (*bool, bool)`

GetLvmEnabledOk returns a tuple with the LvmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLvmEnabled

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetLvmEnabled(v bool)`

SetLvmEnabled sets LvmEnabled field to given value.

### HasLvmEnabled

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) HasLvmEnabled() bool`

HasLvmEnabled returns a boolean if a field has been set.

### GetApiKey

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetSoftwareRaid

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetSoftwareRaid() bool`

GetSoftwareRaid returns the SoftwareRaid field if non-nil, zero value otherwise.

### GetSoftwareRaidOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetSoftwareRaidOk() (*bool, bool)`

GetSoftwareRaidOk returns a tuple with the SoftwareRaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareRaid

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetSoftwareRaid(v bool)`

SetSoftwareRaid sets SoftwareRaid field to given value.

### HasSoftwareRaid

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) HasSoftwareRaid() bool`

HasSoftwareRaid returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetConfig(v string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCapacityInfo

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetCapacityInfo() UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo`

GetCapacityInfo returns the CapacityInfo field if non-nil, zero value otherwise.

### GetCapacityInfoOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetCapacityInfoOk() (*UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo, bool)`

GetCapacityInfoOk returns a tuple with the CapacityInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityInfo

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetCapacityInfo(v UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo)`

SetCapacityInfo sets CapacityInfo field to given value.

### HasCapacityInfo

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) HasCapacityInfo() bool`

HasCapacityInfo returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastStats

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetLastStats() string`

GetLastStats returns the LastStats field if non-nil, zero value otherwise.

### GetLastStatsOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetLastStatsOk() (*string, bool)`

GetLastStatsOk returns a tuple with the LastStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStats

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetLastStats(v string)`

SetLastStats sets LastStats field to given value.

### HasLastStats

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) HasLastStats() bool`

HasLastStats returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetComputeServerType

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetComputeServerType() UpdateServerNetworkInterface200ResponseAllOfOneOfServerComputeServerType`

GetComputeServerType returns the ComputeServerType field if non-nil, zero value otherwise.

### GetComputeServerTypeOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetComputeServerTypeOk() (*UpdateServerNetworkInterface200ResponseAllOfOneOfServerComputeServerType, bool)`

GetComputeServerTypeOk returns a tuple with the ComputeServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServerType

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetComputeServerType(v UpdateServerNetworkInterface200ResponseAllOfOneOfServerComputeServerType)`

SetComputeServerType sets ComputeServerType field to given value.

### HasComputeServerType

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) HasComputeServerType() bool`

HasComputeServerType returns a boolean if a field has been set.

### GetInterfaces

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetInterfaces() []UpdateServerNetworkInterface200ResponseAllOfOneOfServerInterfacesInner`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetInterfacesOk() (*[]UpdateServerNetworkInterface200ResponseAllOfOneOfServerInterfacesInner, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetInterfaces(v []UpdateServerNetworkInterface200ResponseAllOfOneOfServerInterfacesInner)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetZone

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetZone() UpdateServerNetworkInterface200ResponseAllOfOneOfServerZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) GetZoneOk() (*UpdateServerNetworkInterface200ResponseAllOfOneOfServerZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) SetZone(v UpdateServerNetworkInterface200ResponseAllOfOneOfServerZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServer) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


