# NetworkInterfaceUpdateSuccessServer

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
**CapacityInfo** | Pointer to [**NetworkInterfaceUpdateSuccessServerCapacityInfo**](NetworkInterfaceUpdateSuccessServerCapacityInfo.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**LastStats** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ComputeServerType** | Pointer to [**NetworkInterfaceUpdateSuccessServerComputeServerType**](NetworkInterfaceUpdateSuccessServerComputeServerType.md) |  | [optional] 
**Interfaces** | Pointer to [**[]NetworkInterfaceUpdateSuccessServerInterfacesInner**](NetworkInterfaceUpdateSuccessServerInterfacesInner.md) |  | [optional] 
**Zone** | Pointer to [**NetworkInterfaceUpdateSuccessServerZone**](NetworkInterfaceUpdateSuccessServerZone.md) |  | [optional] 

## Methods

### NewNetworkInterfaceUpdateSuccessServer

`func NewNetworkInterfaceUpdateSuccessServer() *NetworkInterfaceUpdateSuccessServer`

NewNetworkInterfaceUpdateSuccessServer instantiates a new NetworkInterfaceUpdateSuccessServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *NetworkInterfaceUpdateSuccessServer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkInterfaceUpdateSuccessServer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkInterfaceUpdateSuccessServer) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkInterfaceUpdateSuccessServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExternalId

`func (o *NetworkInterfaceUpdateSuccessServer) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *NetworkInterfaceUpdateSuccessServer) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *NetworkInterfaceUpdateSuccessServer) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *NetworkInterfaceUpdateSuccessServer) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetAccountId

`func (o *NetworkInterfaceUpdateSuccessServer) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *NetworkInterfaceUpdateSuccessServer) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *NetworkInterfaceUpdateSuccessServer) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *NetworkInterfaceUpdateSuccessServer) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetName

`func (o *NetworkInterfaceUpdateSuccessServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkInterfaceUpdateSuccessServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkInterfaceUpdateSuccessServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkInterfaceUpdateSuccessServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *NetworkInterfaceUpdateSuccessServer) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NetworkInterfaceUpdateSuccessServer) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NetworkInterfaceUpdateSuccessServer) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NetworkInterfaceUpdateSuccessServer) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetVisibility

`func (o *NetworkInterfaceUpdateSuccessServer) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *NetworkInterfaceUpdateSuccessServer) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *NetworkInterfaceUpdateSuccessServer) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *NetworkInterfaceUpdateSuccessServer) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *NetworkInterfaceUpdateSuccessServer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkInterfaceUpdateSuccessServer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkInterfaceUpdateSuccessServer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkInterfaceUpdateSuccessServer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *NetworkInterfaceUpdateSuccessServer) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *NetworkInterfaceUpdateSuccessServer) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetZoneId

`func (o *NetworkInterfaceUpdateSuccessServer) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *NetworkInterfaceUpdateSuccessServer) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *NetworkInterfaceUpdateSuccessServer) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *NetworkInterfaceUpdateSuccessServer) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetSiteId

`func (o *NetworkInterfaceUpdateSuccessServer) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *NetworkInterfaceUpdateSuccessServer) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *NetworkInterfaceUpdateSuccessServer) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *NetworkInterfaceUpdateSuccessServer) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSshHost

`func (o *NetworkInterfaceUpdateSuccessServer) GetSshHost() string`

GetSshHost returns the SshHost field if non-nil, zero value otherwise.

### GetSshHostOk

`func (o *NetworkInterfaceUpdateSuccessServer) GetSshHostOk() (*string, bool)`

GetSshHostOk returns a tuple with the SshHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHost

`func (o *NetworkInterfaceUpdateSuccessServer) SetSshHost(v string)`

SetSshHost sets SshHost field to given value.

### HasSshHost

`func (o *NetworkInterfaceUpdateSuccessServer) HasSshHost() bool`

HasSshHost returns a boolean if a field has been set.

### GetSshPort

`func (o *NetworkInterfaceUpdateSuccessServer) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *NetworkInterfaceUpdateSuccessServer) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *NetworkInterfaceUpdateSuccessServer) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *NetworkInterfaceUpdateSuccessServer) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetExternalIp

`func (o *NetworkInterfaceUpdateSuccessServer) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *NetworkInterfaceUpdateSuccessServer) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *NetworkInterfaceUpdateSuccessServer) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *NetworkInterfaceUpdateSuccessServer) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### GetInternalIp

`func (o *NetworkInterfaceUpdateSuccessServer) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *NetworkInterfaceUpdateSuccessServer) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *NetworkInterfaceUpdateSuccessServer) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *NetworkInterfaceUpdateSuccessServer) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### GetVolumeId

`func (o *NetworkInterfaceUpdateSuccessServer) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *NetworkInterfaceUpdateSuccessServer) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *NetworkInterfaceUpdateSuccessServer) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *NetworkInterfaceUpdateSuccessServer) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.

### SetVolumeIdNil

`func (o *NetworkInterfaceUpdateSuccessServer) SetVolumeIdNil(b bool)`

 SetVolumeIdNil sets the value for VolumeId to be an explicit nil

### UnsetVolumeId
`func (o *NetworkInterfaceUpdateSuccessServer) UnsetVolumeId()`

UnsetVolumeId ensures that no value is present for VolumeId, not even an explicit nil
### GetPlatform

`func (o *NetworkInterfaceUpdateSuccessServer) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *NetworkInterfaceUpdateSuccessServer) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *NetworkInterfaceUpdateSuccessServer) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *NetworkInterfaceUpdateSuccessServer) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPlatformVersion

`func (o *NetworkInterfaceUpdateSuccessServer) GetPlatformVersion() string`

GetPlatformVersion returns the PlatformVersion field if non-nil, zero value otherwise.

### GetPlatformVersionOk

`func (o *NetworkInterfaceUpdateSuccessServer) GetPlatformVersionOk() (*string, bool)`

GetPlatformVersionOk returns a tuple with the PlatformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformVersion

`func (o *NetworkInterfaceUpdateSuccessServer) SetPlatformVersion(v string)`

SetPlatformVersion sets PlatformVersion field to given value.

### HasPlatformVersion

`func (o *NetworkInterfaceUpdateSuccessServer) HasPlatformVersion() bool`

HasPlatformVersion returns a boolean if a field has been set.

### GetSshUsername

`func (o *NetworkInterfaceUpdateSuccessServer) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *NetworkInterfaceUpdateSuccessServer) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *NetworkInterfaceUpdateSuccessServer) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *NetworkInterfaceUpdateSuccessServer) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *NetworkInterfaceUpdateSuccessServer) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *NetworkInterfaceUpdateSuccessServer) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *NetworkInterfaceUpdateSuccessServer) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *NetworkInterfaceUpdateSuccessServer) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetOsDevice

`func (o *NetworkInterfaceUpdateSuccessServer) GetOsDevice() string`

GetOsDevice returns the OsDevice field if non-nil, zero value otherwise.

### GetOsDeviceOk

`func (o *NetworkInterfaceUpdateSuccessServer) GetOsDeviceOk() (*string, bool)`

GetOsDeviceOk returns a tuple with the OsDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDevice

`func (o *NetworkInterfaceUpdateSuccessServer) SetOsDevice(v string)`

SetOsDevice sets OsDevice field to given value.

### HasOsDevice

`func (o *NetworkInterfaceUpdateSuccessServer) HasOsDevice() bool`

HasOsDevice returns a boolean if a field has been set.

### GetDataDevice

`func (o *NetworkInterfaceUpdateSuccessServer) GetDataDevice() string`

GetDataDevice returns the DataDevice field if non-nil, zero value otherwise.

### GetDataDeviceOk

`func (o *NetworkInterfaceUpdateSuccessServer) GetDataDeviceOk() (*string, bool)`

GetDataDeviceOk returns a tuple with the DataDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDevice

`func (o *NetworkInterfaceUpdateSuccessServer) SetDataDevice(v string)`

SetDataDevice sets DataDevice field to given value.

### HasDataDevice

`func (o *NetworkInterfaceUpdateSuccessServer) HasDataDevice() bool`

HasDataDevice returns a boolean if a field has been set.

### GetLvmEnabled

`func (o *NetworkInterfaceUpdateSuccessServer) GetLvmEnabled() bool`

GetLvmEnabled returns the LvmEnabled field if non-nil, zero value otherwise.

### GetLvmEnabledOk

`func (o *NetworkInterfaceUpdateSuccessServer) GetLvmEnabledOk() (*bool, bool)`

GetLvmEnabledOk returns a tuple with the LvmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLvmEnabled

`func (o *NetworkInterfaceUpdateSuccessServer) SetLvmEnabled(v bool)`

SetLvmEnabled sets LvmEnabled field to given value.

### HasLvmEnabled

`func (o *NetworkInterfaceUpdateSuccessServer) HasLvmEnabled() bool`

HasLvmEnabled returns a boolean if a field has been set.

### GetApiKey

`func (o *NetworkInterfaceUpdateSuccessServer) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *NetworkInterfaceUpdateSuccessServer) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *NetworkInterfaceUpdateSuccessServer) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *NetworkInterfaceUpdateSuccessServer) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetSoftwareRaid

`func (o *NetworkInterfaceUpdateSuccessServer) GetSoftwareRaid() bool`

GetSoftwareRaid returns the SoftwareRaid field if non-nil, zero value otherwise.

### GetSoftwareRaidOk

`func (o *NetworkInterfaceUpdateSuccessServer) GetSoftwareRaidOk() (*bool, bool)`

GetSoftwareRaidOk returns a tuple with the SoftwareRaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareRaid

`func (o *NetworkInterfaceUpdateSuccessServer) SetSoftwareRaid(v bool)`

SetSoftwareRaid sets SoftwareRaid field to given value.

### HasSoftwareRaid

`func (o *NetworkInterfaceUpdateSuccessServer) HasSoftwareRaid() bool`

HasSoftwareRaid returns a boolean if a field has been set.

### GetConfig

`func (o *NetworkInterfaceUpdateSuccessServer) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NetworkInterfaceUpdateSuccessServer) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NetworkInterfaceUpdateSuccessServer) SetConfig(v string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *NetworkInterfaceUpdateSuccessServer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCapacityInfo

`func (o *NetworkInterfaceUpdateSuccessServer) GetCapacityInfo() NetworkInterfaceUpdateSuccessServerCapacityInfo`

GetCapacityInfo returns the CapacityInfo field if non-nil, zero value otherwise.

### GetCapacityInfoOk

`func (o *NetworkInterfaceUpdateSuccessServer) GetCapacityInfoOk() (*NetworkInterfaceUpdateSuccessServerCapacityInfo, bool)`

GetCapacityInfoOk returns a tuple with the CapacityInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityInfo

`func (o *NetworkInterfaceUpdateSuccessServer) SetCapacityInfo(v NetworkInterfaceUpdateSuccessServerCapacityInfo)`

SetCapacityInfo sets CapacityInfo field to given value.

### HasCapacityInfo

`func (o *NetworkInterfaceUpdateSuccessServer) HasCapacityInfo() bool`

HasCapacityInfo returns a boolean if a field has been set.

### GetDateCreated

`func (o *NetworkInterfaceUpdateSuccessServer) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *NetworkInterfaceUpdateSuccessServer) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *NetworkInterfaceUpdateSuccessServer) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *NetworkInterfaceUpdateSuccessServer) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *NetworkInterfaceUpdateSuccessServer) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *NetworkInterfaceUpdateSuccessServer) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *NetworkInterfaceUpdateSuccessServer) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *NetworkInterfaceUpdateSuccessServer) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastStats

`func (o *NetworkInterfaceUpdateSuccessServer) GetLastStats() string`

GetLastStats returns the LastStats field if non-nil, zero value otherwise.

### GetLastStatsOk

`func (o *NetworkInterfaceUpdateSuccessServer) GetLastStatsOk() (*string, bool)`

GetLastStatsOk returns a tuple with the LastStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStats

`func (o *NetworkInterfaceUpdateSuccessServer) SetLastStats(v string)`

SetLastStats sets LastStats field to given value.

### HasLastStats

`func (o *NetworkInterfaceUpdateSuccessServer) HasLastStats() bool`

HasLastStats returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkInterfaceUpdateSuccessServer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkInterfaceUpdateSuccessServer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkInterfaceUpdateSuccessServer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkInterfaceUpdateSuccessServer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetComputeServerType

`func (o *NetworkInterfaceUpdateSuccessServer) GetComputeServerType() NetworkInterfaceUpdateSuccessServerComputeServerType`

GetComputeServerType returns the ComputeServerType field if non-nil, zero value otherwise.

### GetComputeServerTypeOk

`func (o *NetworkInterfaceUpdateSuccessServer) GetComputeServerTypeOk() (*NetworkInterfaceUpdateSuccessServerComputeServerType, bool)`

GetComputeServerTypeOk returns a tuple with the ComputeServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServerType

`func (o *NetworkInterfaceUpdateSuccessServer) SetComputeServerType(v NetworkInterfaceUpdateSuccessServerComputeServerType)`

SetComputeServerType sets ComputeServerType field to given value.

### HasComputeServerType

`func (o *NetworkInterfaceUpdateSuccessServer) HasComputeServerType() bool`

HasComputeServerType returns a boolean if a field has been set.

### GetInterfaces

`func (o *NetworkInterfaceUpdateSuccessServer) GetInterfaces() []NetworkInterfaceUpdateSuccessServerInterfacesInner`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *NetworkInterfaceUpdateSuccessServer) GetInterfacesOk() (*[]NetworkInterfaceUpdateSuccessServerInterfacesInner, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *NetworkInterfaceUpdateSuccessServer) SetInterfaces(v []NetworkInterfaceUpdateSuccessServerInterfacesInner)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *NetworkInterfaceUpdateSuccessServer) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetZone

`func (o *NetworkInterfaceUpdateSuccessServer) GetZone() NetworkInterfaceUpdateSuccessServerZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *NetworkInterfaceUpdateSuccessServer) GetZoneOk() (*NetworkInterfaceUpdateSuccessServerZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *NetworkInterfaceUpdateSuccessServer) SetZone(v NetworkInterfaceUpdateSuccessServerZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *NetworkInterfaceUpdateSuccessServer) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


