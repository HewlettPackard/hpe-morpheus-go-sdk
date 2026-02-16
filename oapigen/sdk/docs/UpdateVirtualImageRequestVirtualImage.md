# UpdateVirtualImageRequestVirtualImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the virtual image | [optional] 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**ImageType** | Pointer to **string** | Code of image type. eg. vmware, ami, etc. | [optional] 
**StorageProvider** | Pointer to [**UpdateVirtualImageRequestVirtualImageStorageProvider**](UpdateVirtualImageRequestVirtualImageStorageProvider.md) |  | [optional] 
**IsCloudInit** | Pointer to **bool** | Cloud Init Enabled? | [optional] [default to false]
**UserData** | Pointer to **NullableString** | Cloud-Init User Data, a bash script | [optional] 
**Uefi** | Pointer to **bool** | UEFI enabled? | [optional] 
**FipsEnabled** | Pointer to **bool** | FIPS enabled? | [optional] 
**InstallAgent** | Pointer to **bool** | Install Agent? | [optional] [default to false]
**SshUsername** | Pointer to **NullableString** | SSH Username | [optional] 
**SshPassword** | Pointer to **NullableString** | SSH Password | [optional] 
**SshKey** | Pointer to **NullableString** | SSH Key | [optional] 
**OsType** | Pointer to **NullableInt64** | A Map containing the id of the OS Type. This can also be passed as a string (code or name) instead. | [optional] 
**Visibility** | Pointer to **string** | private or public | [optional] [default to "private"]
**Accounts** | Pointer to **[]int64** |  | [optional] 
**IsAutoJoinDomain** | Pointer to **bool** | Auto Join Domain? | [optional] [default to false]
**VirtioSupported** | Pointer to **bool** | VirtIO Drivers Loaded? | [optional] [default to true]
**VmToolsInstalled** | Pointer to **bool** | VM Tools Installed? | [optional] [default to true]
**IsForceCustomization** | Pointer to **bool** | Force Guest Customization? | [optional] [default to false]
**TrialVersion** | Pointer to **bool** | Trial Version | [optional] [default to false]
**IsSysprep** | Pointer to **bool** | Sysprep Enabled? | [optional] [default to false]
**Config** | Pointer to [**UpdateVirtualImageRequestVirtualImageConfig**](UpdateVirtualImageRequestVirtualImageConfig.md) |  | [optional] 
**Tags** | Pointer to [**[]UpdateVirtualImageRequestVirtualImageTagsInner**](UpdateVirtualImageRequestVirtualImageTagsInner.md) | Metadata tags, Array of objects having a name and value, this adds or updates the specified tags and removes any tags not specified. | [optional] 
**AddTags** | Pointer to [**[]UpdateVirtualImageRequestVirtualImageAddTagsInner**](UpdateVirtualImageRequestVirtualImageAddTagsInner.md) | Add or update value of Metadata tags, Array of objects having a name and value. | [optional] 
**RemoveTags** | Pointer to [**[]UpdateVirtualImageRequestVirtualImageRemoveTagsInner**](UpdateVirtualImageRequestVirtualImageRemoveTagsInner.md) | Remove Metadata tags, Array of objects having a name and an optional value. If value is passed, it must match to be removed. | [optional] 
**MinRamGB** | Pointer to **NullableInt64** |  | [optional] 
**MinDisk** | Pointer to **NullableInt64** |  | [optional] 
**MinDiskGB** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewUpdateVirtualImageRequestVirtualImage

`func NewUpdateVirtualImageRequestVirtualImage() *UpdateVirtualImageRequestVirtualImage`

NewUpdateVirtualImageRequestVirtualImage instantiates a new UpdateVirtualImageRequestVirtualImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVirtualImageRequestVirtualImageWithDefaults

`func NewUpdateVirtualImageRequestVirtualImageWithDefaults() *UpdateVirtualImageRequestVirtualImage`

NewUpdateVirtualImageRequestVirtualImageWithDefaults instantiates a new UpdateVirtualImageRequestVirtualImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateVirtualImageRequestVirtualImage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateVirtualImageRequestVirtualImage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateVirtualImageRequestVirtualImage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateVirtualImageRequestVirtualImage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateVirtualImageRequestVirtualImage) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateVirtualImageRequestVirtualImage) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateVirtualImageRequestVirtualImage) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateVirtualImageRequestVirtualImage) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *UpdateVirtualImageRequestVirtualImage) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *UpdateVirtualImageRequestVirtualImage) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetImageType

`func (o *UpdateVirtualImageRequestVirtualImage) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *UpdateVirtualImageRequestVirtualImage) GetImageTypeOk() (*string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *UpdateVirtualImageRequestVirtualImage) SetImageType(v string)`

SetImageType sets ImageType field to given value.

### HasImageType

`func (o *UpdateVirtualImageRequestVirtualImage) HasImageType() bool`

HasImageType returns a boolean if a field has been set.

### GetStorageProvider

`func (o *UpdateVirtualImageRequestVirtualImage) GetStorageProvider() UpdateVirtualImageRequestVirtualImageStorageProvider`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *UpdateVirtualImageRequestVirtualImage) GetStorageProviderOk() (*UpdateVirtualImageRequestVirtualImageStorageProvider, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *UpdateVirtualImageRequestVirtualImage) SetStorageProvider(v UpdateVirtualImageRequestVirtualImageStorageProvider)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *UpdateVirtualImageRequestVirtualImage) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### GetIsCloudInit

`func (o *UpdateVirtualImageRequestVirtualImage) GetIsCloudInit() bool`

GetIsCloudInit returns the IsCloudInit field if non-nil, zero value otherwise.

### GetIsCloudInitOk

`func (o *UpdateVirtualImageRequestVirtualImage) GetIsCloudInitOk() (*bool, bool)`

GetIsCloudInitOk returns a tuple with the IsCloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudInit

`func (o *UpdateVirtualImageRequestVirtualImage) SetIsCloudInit(v bool)`

SetIsCloudInit sets IsCloudInit field to given value.

### HasIsCloudInit

`func (o *UpdateVirtualImageRequestVirtualImage) HasIsCloudInit() bool`

HasIsCloudInit returns a boolean if a field has been set.

### GetUserData

`func (o *UpdateVirtualImageRequestVirtualImage) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *UpdateVirtualImageRequestVirtualImage) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *UpdateVirtualImageRequestVirtualImage) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *UpdateVirtualImageRequestVirtualImage) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### SetUserDataNil

`func (o *UpdateVirtualImageRequestVirtualImage) SetUserDataNil(b bool)`

 SetUserDataNil sets the value for UserData to be an explicit nil

### UnsetUserData
`func (o *UpdateVirtualImageRequestVirtualImage) UnsetUserData()`

UnsetUserData ensures that no value is present for UserData, not even an explicit nil
### GetUefi

`func (o *UpdateVirtualImageRequestVirtualImage) GetUefi() bool`

GetUefi returns the Uefi field if non-nil, zero value otherwise.

### GetUefiOk

`func (o *UpdateVirtualImageRequestVirtualImage) GetUefiOk() (*bool, bool)`

GetUefiOk returns a tuple with the Uefi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUefi

`func (o *UpdateVirtualImageRequestVirtualImage) SetUefi(v bool)`

SetUefi sets Uefi field to given value.

### HasUefi

`func (o *UpdateVirtualImageRequestVirtualImage) HasUefi() bool`

HasUefi returns a boolean if a field has been set.

### GetFipsEnabled

`func (o *UpdateVirtualImageRequestVirtualImage) GetFipsEnabled() bool`

GetFipsEnabled returns the FipsEnabled field if non-nil, zero value otherwise.

### GetFipsEnabledOk

`func (o *UpdateVirtualImageRequestVirtualImage) GetFipsEnabledOk() (*bool, bool)`

GetFipsEnabledOk returns a tuple with the FipsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFipsEnabled

`func (o *UpdateVirtualImageRequestVirtualImage) SetFipsEnabled(v bool)`

SetFipsEnabled sets FipsEnabled field to given value.

### HasFipsEnabled

`func (o *UpdateVirtualImageRequestVirtualImage) HasFipsEnabled() bool`

HasFipsEnabled returns a boolean if a field has been set.

### GetInstallAgent

`func (o *UpdateVirtualImageRequestVirtualImage) GetInstallAgent() bool`

GetInstallAgent returns the InstallAgent field if non-nil, zero value otherwise.

### GetInstallAgentOk

`func (o *UpdateVirtualImageRequestVirtualImage) GetInstallAgentOk() (*bool, bool)`

GetInstallAgentOk returns a tuple with the InstallAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAgent

`func (o *UpdateVirtualImageRequestVirtualImage) SetInstallAgent(v bool)`

SetInstallAgent sets InstallAgent field to given value.

### HasInstallAgent

`func (o *UpdateVirtualImageRequestVirtualImage) HasInstallAgent() bool`

HasInstallAgent returns a boolean if a field has been set.

### GetSshUsername

`func (o *UpdateVirtualImageRequestVirtualImage) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *UpdateVirtualImageRequestVirtualImage) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *UpdateVirtualImageRequestVirtualImage) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *UpdateVirtualImageRequestVirtualImage) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### SetSshUsernameNil

`func (o *UpdateVirtualImageRequestVirtualImage) SetSshUsernameNil(b bool)`

 SetSshUsernameNil sets the value for SshUsername to be an explicit nil

### UnsetSshUsername
`func (o *UpdateVirtualImageRequestVirtualImage) UnsetSshUsername()`

UnsetSshUsername ensures that no value is present for SshUsername, not even an explicit nil
### GetSshPassword

`func (o *UpdateVirtualImageRequestVirtualImage) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *UpdateVirtualImageRequestVirtualImage) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *UpdateVirtualImageRequestVirtualImage) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *UpdateVirtualImageRequestVirtualImage) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### SetSshPasswordNil

`func (o *UpdateVirtualImageRequestVirtualImage) SetSshPasswordNil(b bool)`

 SetSshPasswordNil sets the value for SshPassword to be an explicit nil

### UnsetSshPassword
`func (o *UpdateVirtualImageRequestVirtualImage) UnsetSshPassword()`

UnsetSshPassword ensures that no value is present for SshPassword, not even an explicit nil
### GetSshKey

`func (o *UpdateVirtualImageRequestVirtualImage) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *UpdateVirtualImageRequestVirtualImage) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *UpdateVirtualImageRequestVirtualImage) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *UpdateVirtualImageRequestVirtualImage) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.

### SetSshKeyNil

`func (o *UpdateVirtualImageRequestVirtualImage) SetSshKeyNil(b bool)`

 SetSshKeyNil sets the value for SshKey to be an explicit nil

### UnsetSshKey
`func (o *UpdateVirtualImageRequestVirtualImage) UnsetSshKey()`

UnsetSshKey ensures that no value is present for SshKey, not even an explicit nil
### GetOsType

`func (o *UpdateVirtualImageRequestVirtualImage) GetOsType() int64`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *UpdateVirtualImageRequestVirtualImage) GetOsTypeOk() (*int64, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *UpdateVirtualImageRequestVirtualImage) SetOsType(v int64)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *UpdateVirtualImageRequestVirtualImage) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### SetOsTypeNil

`func (o *UpdateVirtualImageRequestVirtualImage) SetOsTypeNil(b bool)`

 SetOsTypeNil sets the value for OsType to be an explicit nil

### UnsetOsType
`func (o *UpdateVirtualImageRequestVirtualImage) UnsetOsType()`

UnsetOsType ensures that no value is present for OsType, not even an explicit nil
### GetVisibility

`func (o *UpdateVirtualImageRequestVirtualImage) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateVirtualImageRequestVirtualImage) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateVirtualImageRequestVirtualImage) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateVirtualImageRequestVirtualImage) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAccounts

`func (o *UpdateVirtualImageRequestVirtualImage) GetAccounts() []int64`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *UpdateVirtualImageRequestVirtualImage) GetAccountsOk() (*[]int64, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *UpdateVirtualImageRequestVirtualImage) SetAccounts(v []int64)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *UpdateVirtualImageRequestVirtualImage) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetIsAutoJoinDomain

`func (o *UpdateVirtualImageRequestVirtualImage) GetIsAutoJoinDomain() bool`

GetIsAutoJoinDomain returns the IsAutoJoinDomain field if non-nil, zero value otherwise.

### GetIsAutoJoinDomainOk

`func (o *UpdateVirtualImageRequestVirtualImage) GetIsAutoJoinDomainOk() (*bool, bool)`

GetIsAutoJoinDomainOk returns a tuple with the IsAutoJoinDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoJoinDomain

`func (o *UpdateVirtualImageRequestVirtualImage) SetIsAutoJoinDomain(v bool)`

SetIsAutoJoinDomain sets IsAutoJoinDomain field to given value.

### HasIsAutoJoinDomain

`func (o *UpdateVirtualImageRequestVirtualImage) HasIsAutoJoinDomain() bool`

HasIsAutoJoinDomain returns a boolean if a field has been set.

### GetVirtioSupported

`func (o *UpdateVirtualImageRequestVirtualImage) GetVirtioSupported() bool`

GetVirtioSupported returns the VirtioSupported field if non-nil, zero value otherwise.

### GetVirtioSupportedOk

`func (o *UpdateVirtualImageRequestVirtualImage) GetVirtioSupportedOk() (*bool, bool)`

GetVirtioSupportedOk returns a tuple with the VirtioSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtioSupported

`func (o *UpdateVirtualImageRequestVirtualImage) SetVirtioSupported(v bool)`

SetVirtioSupported sets VirtioSupported field to given value.

### HasVirtioSupported

`func (o *UpdateVirtualImageRequestVirtualImage) HasVirtioSupported() bool`

HasVirtioSupported returns a boolean if a field has been set.

### GetVmToolsInstalled

`func (o *UpdateVirtualImageRequestVirtualImage) GetVmToolsInstalled() bool`

GetVmToolsInstalled returns the VmToolsInstalled field if non-nil, zero value otherwise.

### GetVmToolsInstalledOk

`func (o *UpdateVirtualImageRequestVirtualImage) GetVmToolsInstalledOk() (*bool, bool)`

GetVmToolsInstalledOk returns a tuple with the VmToolsInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmToolsInstalled

`func (o *UpdateVirtualImageRequestVirtualImage) SetVmToolsInstalled(v bool)`

SetVmToolsInstalled sets VmToolsInstalled field to given value.

### HasVmToolsInstalled

`func (o *UpdateVirtualImageRequestVirtualImage) HasVmToolsInstalled() bool`

HasVmToolsInstalled returns a boolean if a field has been set.

### GetIsForceCustomization

`func (o *UpdateVirtualImageRequestVirtualImage) GetIsForceCustomization() bool`

GetIsForceCustomization returns the IsForceCustomization field if non-nil, zero value otherwise.

### GetIsForceCustomizationOk

`func (o *UpdateVirtualImageRequestVirtualImage) GetIsForceCustomizationOk() (*bool, bool)`

GetIsForceCustomizationOk returns a tuple with the IsForceCustomization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForceCustomization

`func (o *UpdateVirtualImageRequestVirtualImage) SetIsForceCustomization(v bool)`

SetIsForceCustomization sets IsForceCustomization field to given value.

### HasIsForceCustomization

`func (o *UpdateVirtualImageRequestVirtualImage) HasIsForceCustomization() bool`

HasIsForceCustomization returns a boolean if a field has been set.

### GetTrialVersion

`func (o *UpdateVirtualImageRequestVirtualImage) GetTrialVersion() bool`

GetTrialVersion returns the TrialVersion field if non-nil, zero value otherwise.

### GetTrialVersionOk

`func (o *UpdateVirtualImageRequestVirtualImage) GetTrialVersionOk() (*bool, bool)`

GetTrialVersionOk returns a tuple with the TrialVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialVersion

`func (o *UpdateVirtualImageRequestVirtualImage) SetTrialVersion(v bool)`

SetTrialVersion sets TrialVersion field to given value.

### HasTrialVersion

`func (o *UpdateVirtualImageRequestVirtualImage) HasTrialVersion() bool`

HasTrialVersion returns a boolean if a field has been set.

### GetIsSysprep

`func (o *UpdateVirtualImageRequestVirtualImage) GetIsSysprep() bool`

GetIsSysprep returns the IsSysprep field if non-nil, zero value otherwise.

### GetIsSysprepOk

`func (o *UpdateVirtualImageRequestVirtualImage) GetIsSysprepOk() (*bool, bool)`

GetIsSysprepOk returns a tuple with the IsSysprep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSysprep

`func (o *UpdateVirtualImageRequestVirtualImage) SetIsSysprep(v bool)`

SetIsSysprep sets IsSysprep field to given value.

### HasIsSysprep

`func (o *UpdateVirtualImageRequestVirtualImage) HasIsSysprep() bool`

HasIsSysprep returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateVirtualImageRequestVirtualImage) GetConfig() UpdateVirtualImageRequestVirtualImageConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateVirtualImageRequestVirtualImage) GetConfigOk() (*UpdateVirtualImageRequestVirtualImageConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateVirtualImageRequestVirtualImage) SetConfig(v UpdateVirtualImageRequestVirtualImageConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateVirtualImageRequestVirtualImage) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetTags

`func (o *UpdateVirtualImageRequestVirtualImage) GetTags() []UpdateVirtualImageRequestVirtualImageTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateVirtualImageRequestVirtualImage) GetTagsOk() (*[]UpdateVirtualImageRequestVirtualImageTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateVirtualImageRequestVirtualImage) SetTags(v []UpdateVirtualImageRequestVirtualImageTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateVirtualImageRequestVirtualImage) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAddTags

`func (o *UpdateVirtualImageRequestVirtualImage) GetAddTags() []UpdateVirtualImageRequestVirtualImageAddTagsInner`

GetAddTags returns the AddTags field if non-nil, zero value otherwise.

### GetAddTagsOk

`func (o *UpdateVirtualImageRequestVirtualImage) GetAddTagsOk() (*[]UpdateVirtualImageRequestVirtualImageAddTagsInner, bool)`

GetAddTagsOk returns a tuple with the AddTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTags

`func (o *UpdateVirtualImageRequestVirtualImage) SetAddTags(v []UpdateVirtualImageRequestVirtualImageAddTagsInner)`

SetAddTags sets AddTags field to given value.

### HasAddTags

`func (o *UpdateVirtualImageRequestVirtualImage) HasAddTags() bool`

HasAddTags returns a boolean if a field has been set.

### GetRemoveTags

`func (o *UpdateVirtualImageRequestVirtualImage) GetRemoveTags() []UpdateVirtualImageRequestVirtualImageRemoveTagsInner`

GetRemoveTags returns the RemoveTags field if non-nil, zero value otherwise.

### GetRemoveTagsOk

`func (o *UpdateVirtualImageRequestVirtualImage) GetRemoveTagsOk() (*[]UpdateVirtualImageRequestVirtualImageRemoveTagsInner, bool)`

GetRemoveTagsOk returns a tuple with the RemoveTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveTags

`func (o *UpdateVirtualImageRequestVirtualImage) SetRemoveTags(v []UpdateVirtualImageRequestVirtualImageRemoveTagsInner)`

SetRemoveTags sets RemoveTags field to given value.

### HasRemoveTags

`func (o *UpdateVirtualImageRequestVirtualImage) HasRemoveTags() bool`

HasRemoveTags returns a boolean if a field has been set.

### GetMinRamGB

`func (o *UpdateVirtualImageRequestVirtualImage) GetMinRamGB() int64`

GetMinRamGB returns the MinRamGB field if non-nil, zero value otherwise.

### GetMinRamGBOk

`func (o *UpdateVirtualImageRequestVirtualImage) GetMinRamGBOk() (*int64, bool)`

GetMinRamGBOk returns a tuple with the MinRamGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRamGB

`func (o *UpdateVirtualImageRequestVirtualImage) SetMinRamGB(v int64)`

SetMinRamGB sets MinRamGB field to given value.

### HasMinRamGB

`func (o *UpdateVirtualImageRequestVirtualImage) HasMinRamGB() bool`

HasMinRamGB returns a boolean if a field has been set.

### SetMinRamGBNil

`func (o *UpdateVirtualImageRequestVirtualImage) SetMinRamGBNil(b bool)`

 SetMinRamGBNil sets the value for MinRamGB to be an explicit nil

### UnsetMinRamGB
`func (o *UpdateVirtualImageRequestVirtualImage) UnsetMinRamGB()`

UnsetMinRamGB ensures that no value is present for MinRamGB, not even an explicit nil
### GetMinDisk

`func (o *UpdateVirtualImageRequestVirtualImage) GetMinDisk() int64`

GetMinDisk returns the MinDisk field if non-nil, zero value otherwise.

### GetMinDiskOk

`func (o *UpdateVirtualImageRequestVirtualImage) GetMinDiskOk() (*int64, bool)`

GetMinDiskOk returns a tuple with the MinDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDisk

`func (o *UpdateVirtualImageRequestVirtualImage) SetMinDisk(v int64)`

SetMinDisk sets MinDisk field to given value.

### HasMinDisk

`func (o *UpdateVirtualImageRequestVirtualImage) HasMinDisk() bool`

HasMinDisk returns a boolean if a field has been set.

### SetMinDiskNil

`func (o *UpdateVirtualImageRequestVirtualImage) SetMinDiskNil(b bool)`

 SetMinDiskNil sets the value for MinDisk to be an explicit nil

### UnsetMinDisk
`func (o *UpdateVirtualImageRequestVirtualImage) UnsetMinDisk()`

UnsetMinDisk ensures that no value is present for MinDisk, not even an explicit nil
### GetMinDiskGB

`func (o *UpdateVirtualImageRequestVirtualImage) GetMinDiskGB() int64`

GetMinDiskGB returns the MinDiskGB field if non-nil, zero value otherwise.

### GetMinDiskGBOk

`func (o *UpdateVirtualImageRequestVirtualImage) GetMinDiskGBOk() (*int64, bool)`

GetMinDiskGBOk returns a tuple with the MinDiskGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDiskGB

`func (o *UpdateVirtualImageRequestVirtualImage) SetMinDiskGB(v int64)`

SetMinDiskGB sets MinDiskGB field to given value.

### HasMinDiskGB

`func (o *UpdateVirtualImageRequestVirtualImage) HasMinDiskGB() bool`

HasMinDiskGB returns a boolean if a field has been set.

### SetMinDiskGBNil

`func (o *UpdateVirtualImageRequestVirtualImage) SetMinDiskGBNil(b bool)`

 SetMinDiskGBNil sets the value for MinDiskGB to be an explicit nil

### UnsetMinDiskGB
`func (o *UpdateVirtualImageRequestVirtualImage) UnsetMinDiskGB()`

UnsetMinDiskGB ensures that no value is present for MinDiskGB, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


