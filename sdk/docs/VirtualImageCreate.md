# VirtualImageCreate

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

### NewVirtualImageCreate

`func NewVirtualImageCreate() *VirtualImageCreate`

NewVirtualImageCreate instantiates a new VirtualImageCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualImageCreateWithDefaults

`func NewVirtualImageCreateWithDefaults() *VirtualImageCreate`

NewVirtualImageCreateWithDefaults instantiates a new VirtualImageCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VirtualImageCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualImageCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualImageCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualImageCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *VirtualImageCreate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *VirtualImageCreate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *VirtualImageCreate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *VirtualImageCreate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetImageType

`func (o *VirtualImageCreate) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *VirtualImageCreate) GetImageTypeOk() (*string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *VirtualImageCreate) SetImageType(v string)`

SetImageType sets ImageType field to given value.

### HasImageType

`func (o *VirtualImageCreate) HasImageType() bool`

HasImageType returns a boolean if a field has been set.

### GetStorageProvider

`func (o *VirtualImageCreate) GetStorageProvider() AddVirtualImageRequestVirtualImageStorageProvider`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *VirtualImageCreate) GetStorageProviderOk() (*AddVirtualImageRequestVirtualImageStorageProvider, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *VirtualImageCreate) SetStorageProvider(v AddVirtualImageRequestVirtualImageStorageProvider)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *VirtualImageCreate) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### GetIsCloudInit

`func (o *VirtualImageCreate) GetIsCloudInit() bool`

GetIsCloudInit returns the IsCloudInit field if non-nil, zero value otherwise.

### GetIsCloudInitOk

`func (o *VirtualImageCreate) GetIsCloudInitOk() (*bool, bool)`

GetIsCloudInitOk returns a tuple with the IsCloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudInit

`func (o *VirtualImageCreate) SetIsCloudInit(v bool)`

SetIsCloudInit sets IsCloudInit field to given value.

### HasIsCloudInit

`func (o *VirtualImageCreate) HasIsCloudInit() bool`

HasIsCloudInit returns a boolean if a field has been set.

### GetUserData

`func (o *VirtualImageCreate) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *VirtualImageCreate) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *VirtualImageCreate) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *VirtualImageCreate) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetInstallAgent

`func (o *VirtualImageCreate) GetInstallAgent() bool`

GetInstallAgent returns the InstallAgent field if non-nil, zero value otherwise.

### GetInstallAgentOk

`func (o *VirtualImageCreate) GetInstallAgentOk() (*bool, bool)`

GetInstallAgentOk returns a tuple with the InstallAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAgent

`func (o *VirtualImageCreate) SetInstallAgent(v bool)`

SetInstallAgent sets InstallAgent field to given value.

### HasInstallAgent

`func (o *VirtualImageCreate) HasInstallAgent() bool`

HasInstallAgent returns a boolean if a field has been set.

### GetSshUsername

`func (o *VirtualImageCreate) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *VirtualImageCreate) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *VirtualImageCreate) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *VirtualImageCreate) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *VirtualImageCreate) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *VirtualImageCreate) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *VirtualImageCreate) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *VirtualImageCreate) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetSshKey

`func (o *VirtualImageCreate) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *VirtualImageCreate) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *VirtualImageCreate) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *VirtualImageCreate) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.

### GetOsType

`func (o *VirtualImageCreate) GetOsType() AddVirtualImageRequestVirtualImageOsType`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *VirtualImageCreate) GetOsTypeOk() (*AddVirtualImageRequestVirtualImageOsType, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *VirtualImageCreate) SetOsType(v AddVirtualImageRequestVirtualImageOsType)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *VirtualImageCreate) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetVisibility

`func (o *VirtualImageCreate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *VirtualImageCreate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *VirtualImageCreate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *VirtualImageCreate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAccounts

`func (o *VirtualImageCreate) GetAccounts() []int64`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *VirtualImageCreate) GetAccountsOk() (*[]int64, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *VirtualImageCreate) SetAccounts(v []int64)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *VirtualImageCreate) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetIsAutoJoinDomain

`func (o *VirtualImageCreate) GetIsAutoJoinDomain() bool`

GetIsAutoJoinDomain returns the IsAutoJoinDomain field if non-nil, zero value otherwise.

### GetIsAutoJoinDomainOk

`func (o *VirtualImageCreate) GetIsAutoJoinDomainOk() (*bool, bool)`

GetIsAutoJoinDomainOk returns a tuple with the IsAutoJoinDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoJoinDomain

`func (o *VirtualImageCreate) SetIsAutoJoinDomain(v bool)`

SetIsAutoJoinDomain sets IsAutoJoinDomain field to given value.

### HasIsAutoJoinDomain

`func (o *VirtualImageCreate) HasIsAutoJoinDomain() bool`

HasIsAutoJoinDomain returns a boolean if a field has been set.

### GetVirtioSupported

`func (o *VirtualImageCreate) GetVirtioSupported() bool`

GetVirtioSupported returns the VirtioSupported field if non-nil, zero value otherwise.

### GetVirtioSupportedOk

`func (o *VirtualImageCreate) GetVirtioSupportedOk() (*bool, bool)`

GetVirtioSupportedOk returns a tuple with the VirtioSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtioSupported

`func (o *VirtualImageCreate) SetVirtioSupported(v bool)`

SetVirtioSupported sets VirtioSupported field to given value.

### HasVirtioSupported

`func (o *VirtualImageCreate) HasVirtioSupported() bool`

HasVirtioSupported returns a boolean if a field has been set.

### GetVmToolsInstalled

`func (o *VirtualImageCreate) GetVmToolsInstalled() bool`

GetVmToolsInstalled returns the VmToolsInstalled field if non-nil, zero value otherwise.

### GetVmToolsInstalledOk

`func (o *VirtualImageCreate) GetVmToolsInstalledOk() (*bool, bool)`

GetVmToolsInstalledOk returns a tuple with the VmToolsInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmToolsInstalled

`func (o *VirtualImageCreate) SetVmToolsInstalled(v bool)`

SetVmToolsInstalled sets VmToolsInstalled field to given value.

### HasVmToolsInstalled

`func (o *VirtualImageCreate) HasVmToolsInstalled() bool`

HasVmToolsInstalled returns a boolean if a field has been set.

### GetIsForceCustomization

`func (o *VirtualImageCreate) GetIsForceCustomization() bool`

GetIsForceCustomization returns the IsForceCustomization field if non-nil, zero value otherwise.

### GetIsForceCustomizationOk

`func (o *VirtualImageCreate) GetIsForceCustomizationOk() (*bool, bool)`

GetIsForceCustomizationOk returns a tuple with the IsForceCustomization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForceCustomization

`func (o *VirtualImageCreate) SetIsForceCustomization(v bool)`

SetIsForceCustomization sets IsForceCustomization field to given value.

### HasIsForceCustomization

`func (o *VirtualImageCreate) HasIsForceCustomization() bool`

HasIsForceCustomization returns a boolean if a field has been set.

### GetTrialVersion

`func (o *VirtualImageCreate) GetTrialVersion() bool`

GetTrialVersion returns the TrialVersion field if non-nil, zero value otherwise.

### GetTrialVersionOk

`func (o *VirtualImageCreate) GetTrialVersionOk() (*bool, bool)`

GetTrialVersionOk returns a tuple with the TrialVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialVersion

`func (o *VirtualImageCreate) SetTrialVersion(v bool)`

SetTrialVersion sets TrialVersion field to given value.

### HasTrialVersion

`func (o *VirtualImageCreate) HasTrialVersion() bool`

HasTrialVersion returns a boolean if a field has been set.

### GetIsSysprep

`func (o *VirtualImageCreate) GetIsSysprep() bool`

GetIsSysprep returns the IsSysprep field if non-nil, zero value otherwise.

### GetIsSysprepOk

`func (o *VirtualImageCreate) GetIsSysprepOk() (*bool, bool)`

GetIsSysprepOk returns a tuple with the IsSysprep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSysprep

`func (o *VirtualImageCreate) SetIsSysprep(v bool)`

SetIsSysprep sets IsSysprep field to given value.

### HasIsSysprep

`func (o *VirtualImageCreate) HasIsSysprep() bool`

HasIsSysprep returns a boolean if a field has been set.

### GetConfig

`func (o *VirtualImageCreate) GetConfig() AddVirtualImageRequestVirtualImageConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *VirtualImageCreate) GetConfigOk() (*AddVirtualImageRequestVirtualImageConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *VirtualImageCreate) SetConfig(v AddVirtualImageRequestVirtualImageConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *VirtualImageCreate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetTags

`func (o *VirtualImageCreate) GetTags() []AddVirtualImageRequestVirtualImageTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VirtualImageCreate) GetTagsOk() (*[]AddVirtualImageRequestVirtualImageTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VirtualImageCreate) SetTags(v []AddVirtualImageRequestVirtualImageTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VirtualImageCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUrl

`func (o *VirtualImageCreate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VirtualImageCreate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VirtualImageCreate) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VirtualImageCreate) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


