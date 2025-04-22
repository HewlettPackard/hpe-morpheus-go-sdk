# ListVirtualImages200ResponseAllOfVirtualImagesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**OwnerId** | Pointer to **int64** |  | [optional] 
**Tenant** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**ImageType** | Pointer to **string** |  | [optional] 
**UserUploaded** | Pointer to **bool** |  | [optional] 
**UserDefined** | Pointer to **bool** |  | [optional] 
**SystemImage** | Pointer to **bool** |  | [optional] 
**IsCloudInit** | Pointer to **bool** |  | [optional] 
**SshUsername** | Pointer to **string** |  | [optional] 
**SshPassword** | Pointer to **string** |  | [optional] 
**SshPasswordHash** | Pointer to **string** |  | [optional] 
**SshKey** | Pointer to **string** |  | [optional] 
**OsType** | Pointer to [**ListVirtualImages200ResponseAllOfVirtualImagesInnerOsType**](ListVirtualImages200ResponseAllOfVirtualImagesInnerOsType.md) |  | [optional] 
**MinRam** | Pointer to **int64** |  | [optional] 
**MinRamGB** | Pointer to **int64** |  | [optional] 
**MinDisk** | Pointer to **string** |  | [optional] 
**MinDiskGB** | Pointer to **string** |  | [optional] 
**RawSize** | Pointer to **int64** |  | [optional] 
**RawSizeGB** | Pointer to **float32** |  | [optional] 
**TrialVersion** | Pointer to **bool** |  | [optional] 
**VirtioSupported** | Pointer to **bool** |  | [optional] 
**Uefi** | Pointer to **string** |  | [optional] 
**IsAutoJoinDomain** | Pointer to **bool** |  | [optional] 
**VmToolsInstalled** | Pointer to **bool** |  | [optional] 
**InstallAgent** | Pointer to **bool** |  | [optional] 
**IsForceCustomization** | Pointer to **bool** |  | [optional] 
**IsSysprep** | Pointer to **bool** |  | [optional] 
**FipsEnabled** | Pointer to **bool** |  | [optional] 
**UserData** | Pointer to **string** |  | [optional] 
**ConsoleKeymap** | Pointer to **string** |  | [optional] 
**StorageProvider** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Accounts** | Pointer to [**[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Config** | Pointer to [**ListVirtualImages200ResponseAllOfVirtualImagesInnerConfig**](ListVirtualImages200ResponseAllOfVirtualImagesInnerConfig.md) |  | [optional] 
**Volumes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**StorageControllers** | Pointer to **[]map[string]interface{}** |  | [optional] 
**NetworkInterfaces** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Tags** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Locations** | Pointer to **[]map[string]interface{}** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewListVirtualImages200ResponseAllOfVirtualImagesInner

`func NewListVirtualImages200ResponseAllOfVirtualImagesInner() *ListVirtualImages200ResponseAllOfVirtualImagesInner`

NewListVirtualImages200ResponseAllOfVirtualImagesInner instantiates a new ListVirtualImages200ResponseAllOfVirtualImagesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVirtualImages200ResponseAllOfVirtualImagesInnerWithDefaults

`func NewListVirtualImages200ResponseAllOfVirtualImagesInnerWithDefaults() *ListVirtualImages200ResponseAllOfVirtualImagesInner`

NewListVirtualImages200ResponseAllOfVirtualImagesInnerWithDefaults instantiates a new ListVirtualImages200ResponseAllOfVirtualImagesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetOwnerId

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetTenant

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetTenant() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetTenantOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetTenant(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetImageType

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetImageTypeOk() (*string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetImageType(v string)`

SetImageType sets ImageType field to given value.

### HasImageType

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasImageType() bool`

HasImageType returns a boolean if a field has been set.

### GetUserUploaded

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetUserUploaded() bool`

GetUserUploaded returns the UserUploaded field if non-nil, zero value otherwise.

### GetUserUploadedOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetUserUploadedOk() (*bool, bool)`

GetUserUploadedOk returns a tuple with the UserUploaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUploaded

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetUserUploaded(v bool)`

SetUserUploaded sets UserUploaded field to given value.

### HasUserUploaded

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasUserUploaded() bool`

HasUserUploaded returns a boolean if a field has been set.

### GetUserDefined

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetUserDefined() bool`

GetUserDefined returns the UserDefined field if non-nil, zero value otherwise.

### GetUserDefinedOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetUserDefinedOk() (*bool, bool)`

GetUserDefinedOk returns a tuple with the UserDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefined

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetUserDefined(v bool)`

SetUserDefined sets UserDefined field to given value.

### HasUserDefined

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasUserDefined() bool`

HasUserDefined returns a boolean if a field has been set.

### GetSystemImage

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetSystemImage() bool`

GetSystemImage returns the SystemImage field if non-nil, zero value otherwise.

### GetSystemImageOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetSystemImageOk() (*bool, bool)`

GetSystemImageOk returns a tuple with the SystemImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemImage

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetSystemImage(v bool)`

SetSystemImage sets SystemImage field to given value.

### HasSystemImage

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasSystemImage() bool`

HasSystemImage returns a boolean if a field has been set.

### GetIsCloudInit

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetIsCloudInit() bool`

GetIsCloudInit returns the IsCloudInit field if non-nil, zero value otherwise.

### GetIsCloudInitOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetIsCloudInitOk() (*bool, bool)`

GetIsCloudInitOk returns a tuple with the IsCloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudInit

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetIsCloudInit(v bool)`

SetIsCloudInit sets IsCloudInit field to given value.

### HasIsCloudInit

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasIsCloudInit() bool`

HasIsCloudInit returns a boolean if a field has been set.

### GetSshUsername

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetSshPasswordHash

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetSshPasswordHash() string`

GetSshPasswordHash returns the SshPasswordHash field if non-nil, zero value otherwise.

### GetSshPasswordHashOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetSshPasswordHashOk() (*string, bool)`

GetSshPasswordHashOk returns a tuple with the SshPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPasswordHash

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetSshPasswordHash(v string)`

SetSshPasswordHash sets SshPasswordHash field to given value.

### HasSshPasswordHash

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasSshPasswordHash() bool`

HasSshPasswordHash returns a boolean if a field has been set.

### GetSshKey

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.

### GetOsType

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetOsType() ListVirtualImages200ResponseAllOfVirtualImagesInnerOsType`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetOsTypeOk() (*ListVirtualImages200ResponseAllOfVirtualImagesInnerOsType, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetOsType(v ListVirtualImages200ResponseAllOfVirtualImagesInnerOsType)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetMinRam

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetMinRam() int64`

GetMinRam returns the MinRam field if non-nil, zero value otherwise.

### GetMinRamOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetMinRamOk() (*int64, bool)`

GetMinRamOk returns a tuple with the MinRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRam

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetMinRam(v int64)`

SetMinRam sets MinRam field to given value.

### HasMinRam

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasMinRam() bool`

HasMinRam returns a boolean if a field has been set.

### GetMinRamGB

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetMinRamGB() int64`

GetMinRamGB returns the MinRamGB field if non-nil, zero value otherwise.

### GetMinRamGBOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetMinRamGBOk() (*int64, bool)`

GetMinRamGBOk returns a tuple with the MinRamGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRamGB

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetMinRamGB(v int64)`

SetMinRamGB sets MinRamGB field to given value.

### HasMinRamGB

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasMinRamGB() bool`

HasMinRamGB returns a boolean if a field has been set.

### GetMinDisk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetMinDisk() string`

GetMinDisk returns the MinDisk field if non-nil, zero value otherwise.

### GetMinDiskOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetMinDiskOk() (*string, bool)`

GetMinDiskOk returns a tuple with the MinDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDisk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetMinDisk(v string)`

SetMinDisk sets MinDisk field to given value.

### HasMinDisk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasMinDisk() bool`

HasMinDisk returns a boolean if a field has been set.

### GetMinDiskGB

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetMinDiskGB() string`

GetMinDiskGB returns the MinDiskGB field if non-nil, zero value otherwise.

### GetMinDiskGBOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetMinDiskGBOk() (*string, bool)`

GetMinDiskGBOk returns a tuple with the MinDiskGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDiskGB

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetMinDiskGB(v string)`

SetMinDiskGB sets MinDiskGB field to given value.

### HasMinDiskGB

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasMinDiskGB() bool`

HasMinDiskGB returns a boolean if a field has been set.

### GetRawSize

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetRawSize() int64`

GetRawSize returns the RawSize field if non-nil, zero value otherwise.

### GetRawSizeOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetRawSizeOk() (*int64, bool)`

GetRawSizeOk returns a tuple with the RawSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSize

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetRawSize(v int64)`

SetRawSize sets RawSize field to given value.

### HasRawSize

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasRawSize() bool`

HasRawSize returns a boolean if a field has been set.

### GetRawSizeGB

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetRawSizeGB() float32`

GetRawSizeGB returns the RawSizeGB field if non-nil, zero value otherwise.

### GetRawSizeGBOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetRawSizeGBOk() (*float32, bool)`

GetRawSizeGBOk returns a tuple with the RawSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSizeGB

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetRawSizeGB(v float32)`

SetRawSizeGB sets RawSizeGB field to given value.

### HasRawSizeGB

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasRawSizeGB() bool`

HasRawSizeGB returns a boolean if a field has been set.

### GetTrialVersion

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetTrialVersion() bool`

GetTrialVersion returns the TrialVersion field if non-nil, zero value otherwise.

### GetTrialVersionOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetTrialVersionOk() (*bool, bool)`

GetTrialVersionOk returns a tuple with the TrialVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialVersion

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetTrialVersion(v bool)`

SetTrialVersion sets TrialVersion field to given value.

### HasTrialVersion

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasTrialVersion() bool`

HasTrialVersion returns a boolean if a field has been set.

### GetVirtioSupported

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetVirtioSupported() bool`

GetVirtioSupported returns the VirtioSupported field if non-nil, zero value otherwise.

### GetVirtioSupportedOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetVirtioSupportedOk() (*bool, bool)`

GetVirtioSupportedOk returns a tuple with the VirtioSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtioSupported

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetVirtioSupported(v bool)`

SetVirtioSupported sets VirtioSupported field to given value.

### HasVirtioSupported

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasVirtioSupported() bool`

HasVirtioSupported returns a boolean if a field has been set.

### GetUefi

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetUefi() string`

GetUefi returns the Uefi field if non-nil, zero value otherwise.

### GetUefiOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetUefiOk() (*string, bool)`

GetUefiOk returns a tuple with the Uefi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUefi

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetUefi(v string)`

SetUefi sets Uefi field to given value.

### HasUefi

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasUefi() bool`

HasUefi returns a boolean if a field has been set.

### GetIsAutoJoinDomain

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetIsAutoJoinDomain() bool`

GetIsAutoJoinDomain returns the IsAutoJoinDomain field if non-nil, zero value otherwise.

### GetIsAutoJoinDomainOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetIsAutoJoinDomainOk() (*bool, bool)`

GetIsAutoJoinDomainOk returns a tuple with the IsAutoJoinDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoJoinDomain

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetIsAutoJoinDomain(v bool)`

SetIsAutoJoinDomain sets IsAutoJoinDomain field to given value.

### HasIsAutoJoinDomain

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasIsAutoJoinDomain() bool`

HasIsAutoJoinDomain returns a boolean if a field has been set.

### GetVmToolsInstalled

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetVmToolsInstalled() bool`

GetVmToolsInstalled returns the VmToolsInstalled field if non-nil, zero value otherwise.

### GetVmToolsInstalledOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetVmToolsInstalledOk() (*bool, bool)`

GetVmToolsInstalledOk returns a tuple with the VmToolsInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmToolsInstalled

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetVmToolsInstalled(v bool)`

SetVmToolsInstalled sets VmToolsInstalled field to given value.

### HasVmToolsInstalled

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasVmToolsInstalled() bool`

HasVmToolsInstalled returns a boolean if a field has been set.

### GetInstallAgent

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetInstallAgent() bool`

GetInstallAgent returns the InstallAgent field if non-nil, zero value otherwise.

### GetInstallAgentOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetInstallAgentOk() (*bool, bool)`

GetInstallAgentOk returns a tuple with the InstallAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAgent

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetInstallAgent(v bool)`

SetInstallAgent sets InstallAgent field to given value.

### HasInstallAgent

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasInstallAgent() bool`

HasInstallAgent returns a boolean if a field has been set.

### GetIsForceCustomization

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetIsForceCustomization() bool`

GetIsForceCustomization returns the IsForceCustomization field if non-nil, zero value otherwise.

### GetIsForceCustomizationOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetIsForceCustomizationOk() (*bool, bool)`

GetIsForceCustomizationOk returns a tuple with the IsForceCustomization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForceCustomization

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetIsForceCustomization(v bool)`

SetIsForceCustomization sets IsForceCustomization field to given value.

### HasIsForceCustomization

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasIsForceCustomization() bool`

HasIsForceCustomization returns a boolean if a field has been set.

### GetIsSysprep

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetIsSysprep() bool`

GetIsSysprep returns the IsSysprep field if non-nil, zero value otherwise.

### GetIsSysprepOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetIsSysprepOk() (*bool, bool)`

GetIsSysprepOk returns a tuple with the IsSysprep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSysprep

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetIsSysprep(v bool)`

SetIsSysprep sets IsSysprep field to given value.

### HasIsSysprep

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasIsSysprep() bool`

HasIsSysprep returns a boolean if a field has been set.

### GetFipsEnabled

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetFipsEnabled() bool`

GetFipsEnabled returns the FipsEnabled field if non-nil, zero value otherwise.

### GetFipsEnabledOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetFipsEnabledOk() (*bool, bool)`

GetFipsEnabledOk returns a tuple with the FipsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFipsEnabled

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetFipsEnabled(v bool)`

SetFipsEnabled sets FipsEnabled field to given value.

### HasFipsEnabled

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasFipsEnabled() bool`

HasFipsEnabled returns a boolean if a field has been set.

### GetUserData

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetConsoleKeymap

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetConsoleKeymap() string`

GetConsoleKeymap returns the ConsoleKeymap field if non-nil, zero value otherwise.

### GetConsoleKeymapOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetConsoleKeymapOk() (*string, bool)`

GetConsoleKeymapOk returns a tuple with the ConsoleKeymap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleKeymap

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetConsoleKeymap(v string)`

SetConsoleKeymap sets ConsoleKeymap field to given value.

### HasConsoleKeymap

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasConsoleKeymap() bool`

HasConsoleKeymap returns a boolean if a field has been set.

### GetStorageProvider

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetStorageProvider() string`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetStorageProviderOk() (*string, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetStorageProvider(v string)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### GetExternalId

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetVisibility

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAccounts

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetAccounts() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetAccountsOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetAccounts(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetConfig

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetConfig() ListVirtualImages200ResponseAllOfVirtualImagesInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetConfigOk() (*ListVirtualImages200ResponseAllOfVirtualImagesInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetConfig(v ListVirtualImages200ResponseAllOfVirtualImagesInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetVolumes

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetVolumes() []map[string]interface{}`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetVolumesOk() (*[]map[string]interface{}, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetVolumes(v []map[string]interface{})`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetStorageControllers

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetStorageControllers() []map[string]interface{}`

GetStorageControllers returns the StorageControllers field if non-nil, zero value otherwise.

### GetStorageControllersOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetStorageControllersOk() (*[]map[string]interface{}, bool)`

GetStorageControllersOk returns a tuple with the StorageControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllers

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetStorageControllers(v []map[string]interface{})`

SetStorageControllers sets StorageControllers field to given value.

### HasStorageControllers

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasStorageControllers() bool`

HasStorageControllers returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetNetworkInterfaces() []map[string]interface{}`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetNetworkInterfacesOk() (*[]map[string]interface{}, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetNetworkInterfaces(v []map[string]interface{})`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetTags

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetTags() []map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetTagsOk() (*[]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetTags(v []map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLocations

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetLocations() []map[string]interface{}`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetLocationsOk() (*[]map[string]interface{}, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetLocations(v []map[string]interface{})`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetStatus

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListVirtualImages200ResponseAllOfVirtualImagesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


