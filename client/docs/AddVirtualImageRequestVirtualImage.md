# AddVirtualImageRequestVirtualImage

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
**Tags** | Pointer to [**[]AddVirtualImageRequestVirtualImageTagsInner**](AddVirtualImageRequestVirtualImageTagsInner.md) | Metadata tags, Array of objects having a name and value | [optional] 
**Url** | Pointer to **string** | Image File URL, a virtual image file will be created by fetching the specified URL | [optional] 

## Methods

### NewAddVirtualImageRequestVirtualImage

`func NewAddVirtualImageRequestVirtualImage() *AddVirtualImageRequestVirtualImage`

NewAddVirtualImageRequestVirtualImage instantiates a new AddVirtualImageRequestVirtualImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddVirtualImageRequestVirtualImageWithDefaults

`func NewAddVirtualImageRequestVirtualImageWithDefaults() *AddVirtualImageRequestVirtualImage`

NewAddVirtualImageRequestVirtualImageWithDefaults instantiates a new AddVirtualImageRequestVirtualImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddVirtualImageRequestVirtualImage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddVirtualImageRequestVirtualImage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddVirtualImageRequestVirtualImage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddVirtualImageRequestVirtualImage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *AddVirtualImageRequestVirtualImage) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddVirtualImageRequestVirtualImage) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddVirtualImageRequestVirtualImage) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddVirtualImageRequestVirtualImage) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetImageType

`func (o *AddVirtualImageRequestVirtualImage) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *AddVirtualImageRequestVirtualImage) GetImageTypeOk() (*string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *AddVirtualImageRequestVirtualImage) SetImageType(v string)`

SetImageType sets ImageType field to given value.

### HasImageType

`func (o *AddVirtualImageRequestVirtualImage) HasImageType() bool`

HasImageType returns a boolean if a field has been set.

### GetStorageProvider

`func (o *AddVirtualImageRequestVirtualImage) GetStorageProvider() AddVirtualImageRequestVirtualImageStorageProvider`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *AddVirtualImageRequestVirtualImage) GetStorageProviderOk() (*AddVirtualImageRequestVirtualImageStorageProvider, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *AddVirtualImageRequestVirtualImage) SetStorageProvider(v AddVirtualImageRequestVirtualImageStorageProvider)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *AddVirtualImageRequestVirtualImage) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### GetIsCloudInit

`func (o *AddVirtualImageRequestVirtualImage) GetIsCloudInit() bool`

GetIsCloudInit returns the IsCloudInit field if non-nil, zero value otherwise.

### GetIsCloudInitOk

`func (o *AddVirtualImageRequestVirtualImage) GetIsCloudInitOk() (*bool, bool)`

GetIsCloudInitOk returns a tuple with the IsCloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudInit

`func (o *AddVirtualImageRequestVirtualImage) SetIsCloudInit(v bool)`

SetIsCloudInit sets IsCloudInit field to given value.

### HasIsCloudInit

`func (o *AddVirtualImageRequestVirtualImage) HasIsCloudInit() bool`

HasIsCloudInit returns a boolean if a field has been set.

### GetUserData

`func (o *AddVirtualImageRequestVirtualImage) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *AddVirtualImageRequestVirtualImage) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *AddVirtualImageRequestVirtualImage) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *AddVirtualImageRequestVirtualImage) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetInstallAgent

`func (o *AddVirtualImageRequestVirtualImage) GetInstallAgent() bool`

GetInstallAgent returns the InstallAgent field if non-nil, zero value otherwise.

### GetInstallAgentOk

`func (o *AddVirtualImageRequestVirtualImage) GetInstallAgentOk() (*bool, bool)`

GetInstallAgentOk returns a tuple with the InstallAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAgent

`func (o *AddVirtualImageRequestVirtualImage) SetInstallAgent(v bool)`

SetInstallAgent sets InstallAgent field to given value.

### HasInstallAgent

`func (o *AddVirtualImageRequestVirtualImage) HasInstallAgent() bool`

HasInstallAgent returns a boolean if a field has been set.

### GetSshUsername

`func (o *AddVirtualImageRequestVirtualImage) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *AddVirtualImageRequestVirtualImage) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *AddVirtualImageRequestVirtualImage) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *AddVirtualImageRequestVirtualImage) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *AddVirtualImageRequestVirtualImage) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *AddVirtualImageRequestVirtualImage) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *AddVirtualImageRequestVirtualImage) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *AddVirtualImageRequestVirtualImage) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetSshKey

`func (o *AddVirtualImageRequestVirtualImage) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *AddVirtualImageRequestVirtualImage) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *AddVirtualImageRequestVirtualImage) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *AddVirtualImageRequestVirtualImage) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.

### GetOsType

`func (o *AddVirtualImageRequestVirtualImage) GetOsType() AddVirtualImageRequestVirtualImageOsType`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *AddVirtualImageRequestVirtualImage) GetOsTypeOk() (*AddVirtualImageRequestVirtualImageOsType, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *AddVirtualImageRequestVirtualImage) SetOsType(v AddVirtualImageRequestVirtualImageOsType)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *AddVirtualImageRequestVirtualImage) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetVisibility

`func (o *AddVirtualImageRequestVirtualImage) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AddVirtualImageRequestVirtualImage) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AddVirtualImageRequestVirtualImage) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AddVirtualImageRequestVirtualImage) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAccounts

`func (o *AddVirtualImageRequestVirtualImage) GetAccounts() []int64`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *AddVirtualImageRequestVirtualImage) GetAccountsOk() (*[]int64, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *AddVirtualImageRequestVirtualImage) SetAccounts(v []int64)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *AddVirtualImageRequestVirtualImage) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetIsAutoJoinDomain

`func (o *AddVirtualImageRequestVirtualImage) GetIsAutoJoinDomain() bool`

GetIsAutoJoinDomain returns the IsAutoJoinDomain field if non-nil, zero value otherwise.

### GetIsAutoJoinDomainOk

`func (o *AddVirtualImageRequestVirtualImage) GetIsAutoJoinDomainOk() (*bool, bool)`

GetIsAutoJoinDomainOk returns a tuple with the IsAutoJoinDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoJoinDomain

`func (o *AddVirtualImageRequestVirtualImage) SetIsAutoJoinDomain(v bool)`

SetIsAutoJoinDomain sets IsAutoJoinDomain field to given value.

### HasIsAutoJoinDomain

`func (o *AddVirtualImageRequestVirtualImage) HasIsAutoJoinDomain() bool`

HasIsAutoJoinDomain returns a boolean if a field has been set.

### GetVirtioSupported

`func (o *AddVirtualImageRequestVirtualImage) GetVirtioSupported() bool`

GetVirtioSupported returns the VirtioSupported field if non-nil, zero value otherwise.

### GetVirtioSupportedOk

`func (o *AddVirtualImageRequestVirtualImage) GetVirtioSupportedOk() (*bool, bool)`

GetVirtioSupportedOk returns a tuple with the VirtioSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtioSupported

`func (o *AddVirtualImageRequestVirtualImage) SetVirtioSupported(v bool)`

SetVirtioSupported sets VirtioSupported field to given value.

### HasVirtioSupported

`func (o *AddVirtualImageRequestVirtualImage) HasVirtioSupported() bool`

HasVirtioSupported returns a boolean if a field has been set.

### GetVmToolsInstalled

`func (o *AddVirtualImageRequestVirtualImage) GetVmToolsInstalled() bool`

GetVmToolsInstalled returns the VmToolsInstalled field if non-nil, zero value otherwise.

### GetVmToolsInstalledOk

`func (o *AddVirtualImageRequestVirtualImage) GetVmToolsInstalledOk() (*bool, bool)`

GetVmToolsInstalledOk returns a tuple with the VmToolsInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmToolsInstalled

`func (o *AddVirtualImageRequestVirtualImage) SetVmToolsInstalled(v bool)`

SetVmToolsInstalled sets VmToolsInstalled field to given value.

### HasVmToolsInstalled

`func (o *AddVirtualImageRequestVirtualImage) HasVmToolsInstalled() bool`

HasVmToolsInstalled returns a boolean if a field has been set.

### GetIsForceCustomization

`func (o *AddVirtualImageRequestVirtualImage) GetIsForceCustomization() bool`

GetIsForceCustomization returns the IsForceCustomization field if non-nil, zero value otherwise.

### GetIsForceCustomizationOk

`func (o *AddVirtualImageRequestVirtualImage) GetIsForceCustomizationOk() (*bool, bool)`

GetIsForceCustomizationOk returns a tuple with the IsForceCustomization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForceCustomization

`func (o *AddVirtualImageRequestVirtualImage) SetIsForceCustomization(v bool)`

SetIsForceCustomization sets IsForceCustomization field to given value.

### HasIsForceCustomization

`func (o *AddVirtualImageRequestVirtualImage) HasIsForceCustomization() bool`

HasIsForceCustomization returns a boolean if a field has been set.

### GetTrialVersion

`func (o *AddVirtualImageRequestVirtualImage) GetTrialVersion() bool`

GetTrialVersion returns the TrialVersion field if non-nil, zero value otherwise.

### GetTrialVersionOk

`func (o *AddVirtualImageRequestVirtualImage) GetTrialVersionOk() (*bool, bool)`

GetTrialVersionOk returns a tuple with the TrialVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialVersion

`func (o *AddVirtualImageRequestVirtualImage) SetTrialVersion(v bool)`

SetTrialVersion sets TrialVersion field to given value.

### HasTrialVersion

`func (o *AddVirtualImageRequestVirtualImage) HasTrialVersion() bool`

HasTrialVersion returns a boolean if a field has been set.

### GetIsSysprep

`func (o *AddVirtualImageRequestVirtualImage) GetIsSysprep() bool`

GetIsSysprep returns the IsSysprep field if non-nil, zero value otherwise.

### GetIsSysprepOk

`func (o *AddVirtualImageRequestVirtualImage) GetIsSysprepOk() (*bool, bool)`

GetIsSysprepOk returns a tuple with the IsSysprep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSysprep

`func (o *AddVirtualImageRequestVirtualImage) SetIsSysprep(v bool)`

SetIsSysprep sets IsSysprep field to given value.

### HasIsSysprep

`func (o *AddVirtualImageRequestVirtualImage) HasIsSysprep() bool`

HasIsSysprep returns a boolean if a field has been set.

### GetConfig

`func (o *AddVirtualImageRequestVirtualImage) GetConfig() AddVirtualImageRequestVirtualImageConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddVirtualImageRequestVirtualImage) GetConfigOk() (*AddVirtualImageRequestVirtualImageConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddVirtualImageRequestVirtualImage) SetConfig(v AddVirtualImageRequestVirtualImageConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddVirtualImageRequestVirtualImage) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetTags

`func (o *AddVirtualImageRequestVirtualImage) GetTags() []AddVirtualImageRequestVirtualImageTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AddVirtualImageRequestVirtualImage) GetTagsOk() (*[]AddVirtualImageRequestVirtualImageTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AddVirtualImageRequestVirtualImage) SetTags(v []AddVirtualImageRequestVirtualImageTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AddVirtualImageRequestVirtualImage) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUrl

`func (o *AddVirtualImageRequestVirtualImage) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AddVirtualImageRequestVirtualImage) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AddVirtualImageRequestVirtualImage) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AddVirtualImageRequestVirtualImage) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


