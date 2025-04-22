# UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ZoneId** | Pointer to **int64** |  | [optional] 
**SiteId** | Pointer to **int64** |  | [optional] 
**SshHost** | Pointer to **string** |  | [optional] 
**SshPort** | Pointer to **int64** |  | [optional] 
**ExternalIp** | Pointer to **string** |  | [optional] 
**InternalIp** | Pointer to **string** |  | [optional] 
**VolumeId** | Pointer to **string** |  | [optional] 
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
**CapacityInfo** | Pointer to [**UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerCapacityInfo**](UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerCapacityInfo.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**LastStats** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ComputeServerType** | Pointer to [**UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerComputeServerType**](UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerComputeServerType.md) |  | [optional] 
**Interfaces** | Pointer to [**[]UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerInterfacesInner**](UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerInterfacesInner.md) |  | [optional] 
**Zone** | Pointer to [**UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone**](UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone.md) |  | [optional] 

## Methods

### NewUpdateInstanceNetworkInterface200ResponseAllOfOneOfServer

`func NewUpdateInstanceNetworkInterface200ResponseAllOfOneOfServer() *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer`

NewUpdateInstanceNetworkInterface200ResponseAllOfOneOfServer instantiates a new UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInstanceNetworkInterface200ResponseAllOfOneOfServerWithDefaults

`func NewUpdateInstanceNetworkInterface200ResponseAllOfOneOfServerWithDefaults() *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer`

NewUpdateInstanceNetworkInterface200ResponseAllOfOneOfServerWithDefaults instantiates a new UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExternalId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetAccountId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetName

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetZoneId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetSiteId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSshHost

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetSshHost() string`

GetSshHost returns the SshHost field if non-nil, zero value otherwise.

### GetSshHostOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetSshHostOk() (*string, bool)`

GetSshHostOk returns a tuple with the SshHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHost

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) SetSshHost(v string)`

SetSshHost sets SshHost field to given value.

### HasSshHost

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) HasSshHost() bool`

HasSshHost returns a boolean if a field has been set.

### GetSshPort

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetExternalIp

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### GetInternalIp

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### GetVolumeId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.

### GetPlatform

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPlatformVersion

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetPlatformVersion() string`

GetPlatformVersion returns the PlatformVersion field if non-nil, zero value otherwise.

### GetPlatformVersionOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetPlatformVersionOk() (*string, bool)`

GetPlatformVersionOk returns a tuple with the PlatformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformVersion

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) SetPlatformVersion(v string)`

SetPlatformVersion sets PlatformVersion field to given value.

### HasPlatformVersion

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) HasPlatformVersion() bool`

HasPlatformVersion returns a boolean if a field has been set.

### GetSshUsername

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetOsDevice

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetOsDevice() string`

GetOsDevice returns the OsDevice field if non-nil, zero value otherwise.

### GetOsDeviceOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetOsDeviceOk() (*string, bool)`

GetOsDeviceOk returns a tuple with the OsDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDevice

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) SetOsDevice(v string)`

SetOsDevice sets OsDevice field to given value.

### HasOsDevice

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) HasOsDevice() bool`

HasOsDevice returns a boolean if a field has been set.

### GetDataDevice

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetDataDevice() string`

GetDataDevice returns the DataDevice field if non-nil, zero value otherwise.

### GetDataDeviceOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetDataDeviceOk() (*string, bool)`

GetDataDeviceOk returns a tuple with the DataDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDevice

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) SetDataDevice(v string)`

SetDataDevice sets DataDevice field to given value.

### HasDataDevice

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) HasDataDevice() bool`

HasDataDevice returns a boolean if a field has been set.

### GetLvmEnabled

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetLvmEnabled() bool`

GetLvmEnabled returns the LvmEnabled field if non-nil, zero value otherwise.

### GetLvmEnabledOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetLvmEnabledOk() (*bool, bool)`

GetLvmEnabledOk returns a tuple with the LvmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLvmEnabled

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) SetLvmEnabled(v bool)`

SetLvmEnabled sets LvmEnabled field to given value.

### HasLvmEnabled

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) HasLvmEnabled() bool`

HasLvmEnabled returns a boolean if a field has been set.

### GetApiKey

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetSoftwareRaid

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetSoftwareRaid() bool`

GetSoftwareRaid returns the SoftwareRaid field if non-nil, zero value otherwise.

### GetSoftwareRaidOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetSoftwareRaidOk() (*bool, bool)`

GetSoftwareRaidOk returns a tuple with the SoftwareRaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareRaid

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) SetSoftwareRaid(v bool)`

SetSoftwareRaid sets SoftwareRaid field to given value.

### HasSoftwareRaid

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) HasSoftwareRaid() bool`

HasSoftwareRaid returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) SetConfig(v string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCapacityInfo

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetCapacityInfo() UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerCapacityInfo`

GetCapacityInfo returns the CapacityInfo field if non-nil, zero value otherwise.

### GetCapacityInfoOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetCapacityInfoOk() (*UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerCapacityInfo, bool)`

GetCapacityInfoOk returns a tuple with the CapacityInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityInfo

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) SetCapacityInfo(v UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerCapacityInfo)`

SetCapacityInfo sets CapacityInfo field to given value.

### HasCapacityInfo

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) HasCapacityInfo() bool`

HasCapacityInfo returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastStats

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetLastStats() string`

GetLastStats returns the LastStats field if non-nil, zero value otherwise.

### GetLastStatsOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetLastStatsOk() (*string, bool)`

GetLastStatsOk returns a tuple with the LastStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStats

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) SetLastStats(v string)`

SetLastStats sets LastStats field to given value.

### HasLastStats

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) HasLastStats() bool`

HasLastStats returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetComputeServerType

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetComputeServerType() UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerComputeServerType`

GetComputeServerType returns the ComputeServerType field if non-nil, zero value otherwise.

### GetComputeServerTypeOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetComputeServerTypeOk() (*UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerComputeServerType, bool)`

GetComputeServerTypeOk returns a tuple with the ComputeServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServerType

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) SetComputeServerType(v UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerComputeServerType)`

SetComputeServerType sets ComputeServerType field to given value.

### HasComputeServerType

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) HasComputeServerType() bool`

HasComputeServerType returns a boolean if a field has been set.

### GetInterfaces

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetInterfaces() []UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerInterfacesInner`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetInterfacesOk() (*[]UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerInterfacesInner, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) SetInterfaces(v []UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerInterfacesInner)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetZone

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetZone() UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) GetZoneOk() (*UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) SetZone(v UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


