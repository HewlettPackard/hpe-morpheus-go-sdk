# VirtualImageUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the virtual image | [optional] 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**ImageType** | Pointer to **string** | Code of image type. eg. vmware, ami, etc. | [optional] 
**StorageProvider** | Pointer to [**AddVirtualImageRequestVirtualImageStorageProvider**](AddVirtualImageRequestVirtualImageStorageProvider.md) |  | [optional] 
**IsCloudInit** | Pointer to **bool** | Cloud Init Enabled? | [optional] [default to false]
**UserData** | Pointer to **string** | Cloud-Init User Data, a bash script | [optional] 
**InstallAgent** | Pointer to **bool** | Install Agent? | [optional] [default to false]
**SshUsername** | Pointer to **string** | SSH Username | [optional] 
**SshPassword** | Pointer to **string** | SSH Password | [optional] 
**SshKey** | Pointer to **string** | SSH Key | [optional] 
**OsType** | Pointer to [**AddVirtualImageRequestVirtualImageOsType**](AddVirtualImageRequestVirtualImageOsType.md) |  | [optional] 
**Visibility** | Pointer to **string** | private or public | [optional] [default to "private"]
**Accounts** | Pointer to **[]int64** |  | [optional] 
**IsAutoJoinDomain** | Pointer to **bool** | Auto Join Domain? | [optional] [default to false]
**VirtioSupported** | Pointer to **bool** | VirtIO Drivers Loaded? | [optional] [default to true]
**VmToolsInstalled** | Pointer to **bool** | VM Tools Installed? | [optional] [default to true]
**IsForceCustomization** | Pointer to **bool** | Force Guest Customization? | [optional] [default to false]
**TrialVersion** | Pointer to **bool** | Trial Version | [optional] [default to false]
**IsSysprep** | Pointer to **bool** | Sysprep Enabled? | [optional] [default to false]
**Config** | Pointer to [**AddVirtualImageRequestVirtualImageConfig**](AddVirtualImageRequestVirtualImageConfig.md) |  | [optional] 
**Tags** | Pointer to [**[]AddVirtualImageRequestVirtualImageTagsInner**](AddVirtualImageRequestVirtualImageTagsInner.md) | Metadata tags, Array of objects having a name and value, this adds or updates the specified tags and removes any tags not specified. | [optional] 
**AddTags** | Pointer to [**[]AddVirtualImageRequestVirtualImageTagsInner**](AddVirtualImageRequestVirtualImageTagsInner.md) | Add or update value of Metadata tags, Array of objects having a name and value. | [optional] 
**RemoveTags** | Pointer to [**[]UpdateVirtualImageRequestVirtualImageRemoveTagsInner**](UpdateVirtualImageRequestVirtualImageRemoveTagsInner.md) | Remove Metadata tags, Array of objects having a name and an optional value. If value is passed, it must match to be removed. | [optional] 

## Methods

### NewVirtualImageUpdate

`func NewVirtualImageUpdate() *VirtualImageUpdate`

NewVirtualImageUpdate instantiates a new VirtualImageUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualImageUpdateWithDefaults

`func NewVirtualImageUpdateWithDefaults() *VirtualImageUpdate`

NewVirtualImageUpdateWithDefaults instantiates a new VirtualImageUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VirtualImageUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualImageUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualImageUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualImageUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *VirtualImageUpdate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *VirtualImageUpdate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *VirtualImageUpdate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *VirtualImageUpdate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetImageType

`func (o *VirtualImageUpdate) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *VirtualImageUpdate) GetImageTypeOk() (*string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *VirtualImageUpdate) SetImageType(v string)`

SetImageType sets ImageType field to given value.

### HasImageType

`func (o *VirtualImageUpdate) HasImageType() bool`

HasImageType returns a boolean if a field has been set.

### GetStorageProvider

`func (o *VirtualImageUpdate) GetStorageProvider() AddVirtualImageRequestVirtualImageStorageProvider`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *VirtualImageUpdate) GetStorageProviderOk() (*AddVirtualImageRequestVirtualImageStorageProvider, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *VirtualImageUpdate) SetStorageProvider(v AddVirtualImageRequestVirtualImageStorageProvider)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *VirtualImageUpdate) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### GetIsCloudInit

`func (o *VirtualImageUpdate) GetIsCloudInit() bool`

GetIsCloudInit returns the IsCloudInit field if non-nil, zero value otherwise.

### GetIsCloudInitOk

`func (o *VirtualImageUpdate) GetIsCloudInitOk() (*bool, bool)`

GetIsCloudInitOk returns a tuple with the IsCloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudInit

`func (o *VirtualImageUpdate) SetIsCloudInit(v bool)`

SetIsCloudInit sets IsCloudInit field to given value.

### HasIsCloudInit

`func (o *VirtualImageUpdate) HasIsCloudInit() bool`

HasIsCloudInit returns a boolean if a field has been set.

### GetUserData

`func (o *VirtualImageUpdate) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *VirtualImageUpdate) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *VirtualImageUpdate) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *VirtualImageUpdate) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetInstallAgent

`func (o *VirtualImageUpdate) GetInstallAgent() bool`

GetInstallAgent returns the InstallAgent field if non-nil, zero value otherwise.

### GetInstallAgentOk

`func (o *VirtualImageUpdate) GetInstallAgentOk() (*bool, bool)`

GetInstallAgentOk returns a tuple with the InstallAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAgent

`func (o *VirtualImageUpdate) SetInstallAgent(v bool)`

SetInstallAgent sets InstallAgent field to given value.

### HasInstallAgent

`func (o *VirtualImageUpdate) HasInstallAgent() bool`

HasInstallAgent returns a boolean if a field has been set.

### GetSshUsername

`func (o *VirtualImageUpdate) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *VirtualImageUpdate) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *VirtualImageUpdate) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *VirtualImageUpdate) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *VirtualImageUpdate) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *VirtualImageUpdate) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *VirtualImageUpdate) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *VirtualImageUpdate) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetSshKey

`func (o *VirtualImageUpdate) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *VirtualImageUpdate) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *VirtualImageUpdate) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *VirtualImageUpdate) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.

### GetOsType

`func (o *VirtualImageUpdate) GetOsType() AddVirtualImageRequestVirtualImageOsType`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *VirtualImageUpdate) GetOsTypeOk() (*AddVirtualImageRequestVirtualImageOsType, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *VirtualImageUpdate) SetOsType(v AddVirtualImageRequestVirtualImageOsType)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *VirtualImageUpdate) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetVisibility

`func (o *VirtualImageUpdate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *VirtualImageUpdate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *VirtualImageUpdate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *VirtualImageUpdate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAccounts

`func (o *VirtualImageUpdate) GetAccounts() []int64`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *VirtualImageUpdate) GetAccountsOk() (*[]int64, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *VirtualImageUpdate) SetAccounts(v []int64)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *VirtualImageUpdate) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetIsAutoJoinDomain

`func (o *VirtualImageUpdate) GetIsAutoJoinDomain() bool`

GetIsAutoJoinDomain returns the IsAutoJoinDomain field if non-nil, zero value otherwise.

### GetIsAutoJoinDomainOk

`func (o *VirtualImageUpdate) GetIsAutoJoinDomainOk() (*bool, bool)`

GetIsAutoJoinDomainOk returns a tuple with the IsAutoJoinDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoJoinDomain

`func (o *VirtualImageUpdate) SetIsAutoJoinDomain(v bool)`

SetIsAutoJoinDomain sets IsAutoJoinDomain field to given value.

### HasIsAutoJoinDomain

`func (o *VirtualImageUpdate) HasIsAutoJoinDomain() bool`

HasIsAutoJoinDomain returns a boolean if a field has been set.

### GetVirtioSupported

`func (o *VirtualImageUpdate) GetVirtioSupported() bool`

GetVirtioSupported returns the VirtioSupported field if non-nil, zero value otherwise.

### GetVirtioSupportedOk

`func (o *VirtualImageUpdate) GetVirtioSupportedOk() (*bool, bool)`

GetVirtioSupportedOk returns a tuple with the VirtioSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtioSupported

`func (o *VirtualImageUpdate) SetVirtioSupported(v bool)`

SetVirtioSupported sets VirtioSupported field to given value.

### HasVirtioSupported

`func (o *VirtualImageUpdate) HasVirtioSupported() bool`

HasVirtioSupported returns a boolean if a field has been set.

### GetVmToolsInstalled

`func (o *VirtualImageUpdate) GetVmToolsInstalled() bool`

GetVmToolsInstalled returns the VmToolsInstalled field if non-nil, zero value otherwise.

### GetVmToolsInstalledOk

`func (o *VirtualImageUpdate) GetVmToolsInstalledOk() (*bool, bool)`

GetVmToolsInstalledOk returns a tuple with the VmToolsInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmToolsInstalled

`func (o *VirtualImageUpdate) SetVmToolsInstalled(v bool)`

SetVmToolsInstalled sets VmToolsInstalled field to given value.

### HasVmToolsInstalled

`func (o *VirtualImageUpdate) HasVmToolsInstalled() bool`

HasVmToolsInstalled returns a boolean if a field has been set.

### GetIsForceCustomization

`func (o *VirtualImageUpdate) GetIsForceCustomization() bool`

GetIsForceCustomization returns the IsForceCustomization field if non-nil, zero value otherwise.

### GetIsForceCustomizationOk

`func (o *VirtualImageUpdate) GetIsForceCustomizationOk() (*bool, bool)`

GetIsForceCustomizationOk returns a tuple with the IsForceCustomization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForceCustomization

`func (o *VirtualImageUpdate) SetIsForceCustomization(v bool)`

SetIsForceCustomization sets IsForceCustomization field to given value.

### HasIsForceCustomization

`func (o *VirtualImageUpdate) HasIsForceCustomization() bool`

HasIsForceCustomization returns a boolean if a field has been set.

### GetTrialVersion

`func (o *VirtualImageUpdate) GetTrialVersion() bool`

GetTrialVersion returns the TrialVersion field if non-nil, zero value otherwise.

### GetTrialVersionOk

`func (o *VirtualImageUpdate) GetTrialVersionOk() (*bool, bool)`

GetTrialVersionOk returns a tuple with the TrialVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialVersion

`func (o *VirtualImageUpdate) SetTrialVersion(v bool)`

SetTrialVersion sets TrialVersion field to given value.

### HasTrialVersion

`func (o *VirtualImageUpdate) HasTrialVersion() bool`

HasTrialVersion returns a boolean if a field has been set.

### GetIsSysprep

`func (o *VirtualImageUpdate) GetIsSysprep() bool`

GetIsSysprep returns the IsSysprep field if non-nil, zero value otherwise.

### GetIsSysprepOk

`func (o *VirtualImageUpdate) GetIsSysprepOk() (*bool, bool)`

GetIsSysprepOk returns a tuple with the IsSysprep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSysprep

`func (o *VirtualImageUpdate) SetIsSysprep(v bool)`

SetIsSysprep sets IsSysprep field to given value.

### HasIsSysprep

`func (o *VirtualImageUpdate) HasIsSysprep() bool`

HasIsSysprep returns a boolean if a field has been set.

### GetConfig

`func (o *VirtualImageUpdate) GetConfig() AddVirtualImageRequestVirtualImageConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *VirtualImageUpdate) GetConfigOk() (*AddVirtualImageRequestVirtualImageConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *VirtualImageUpdate) SetConfig(v AddVirtualImageRequestVirtualImageConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *VirtualImageUpdate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetTags

`func (o *VirtualImageUpdate) GetTags() []AddVirtualImageRequestVirtualImageTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VirtualImageUpdate) GetTagsOk() (*[]AddVirtualImageRequestVirtualImageTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VirtualImageUpdate) SetTags(v []AddVirtualImageRequestVirtualImageTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VirtualImageUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAddTags

`func (o *VirtualImageUpdate) GetAddTags() []AddVirtualImageRequestVirtualImageTagsInner`

GetAddTags returns the AddTags field if non-nil, zero value otherwise.

### GetAddTagsOk

`func (o *VirtualImageUpdate) GetAddTagsOk() (*[]AddVirtualImageRequestVirtualImageTagsInner, bool)`

GetAddTagsOk returns a tuple with the AddTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTags

`func (o *VirtualImageUpdate) SetAddTags(v []AddVirtualImageRequestVirtualImageTagsInner)`

SetAddTags sets AddTags field to given value.

### HasAddTags

`func (o *VirtualImageUpdate) HasAddTags() bool`

HasAddTags returns a boolean if a field has been set.

### GetRemoveTags

`func (o *VirtualImageUpdate) GetRemoveTags() []UpdateVirtualImageRequestVirtualImageRemoveTagsInner`

GetRemoveTags returns the RemoveTags field if non-nil, zero value otherwise.

### GetRemoveTagsOk

`func (o *VirtualImageUpdate) GetRemoveTagsOk() (*[]UpdateVirtualImageRequestVirtualImageRemoveTagsInner, bool)`

GetRemoveTagsOk returns a tuple with the RemoveTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveTags

`func (o *VirtualImageUpdate) SetRemoveTags(v []UpdateVirtualImageRequestVirtualImageRemoveTagsInner)`

SetRemoveTags sets RemoveTags field to given value.

### HasRemoveTags

`func (o *VirtualImageUpdate) HasRemoveTags() bool`

HasRemoveTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


