# GetVirtualImage200ResponseVirtualImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** | A name for the virtual image | [optional] 
**Description** | Pointer to **NullableString** | A description for the virtual image | [optional] 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**OwnerId** | Pointer to **int64** | Owner of the image | [optional] 
**Tenant** | Pointer to [**GetVirtualImage200ResponseVirtualImageTenant**](GetVirtualImage200ResponseVirtualImageTenant.md) |  | [optional] 
**ImageType** | Pointer to **string** | Code of image type. eg. vmware, ami, etc. | [optional] 
**UserUploaded** | Pointer to **bool** | Is uploaded by an user? | [optional] 
**UserDefined** | Pointer to **bool** | Is defined by an user? | [optional] 
**SystemImage** | Pointer to **bool** | Is created by system? | [optional] 
**IsCloudInit** | Pointer to **bool** | Cloud Init Enabled? | [optional] 
**SshUsername** | Pointer to **NullableString** | SSH Username | [optional] 
**SshPassword** | Pointer to **NullableString** |  | [optional] 
**SshPasswordHash** | Pointer to **NullableString** |  | [optional] 
**SshKey** | Pointer to **NullableString** | SSH Key | [optional] 
**OsType** | Pointer to [**GetVirtualImage200ResponseVirtualImageOsType**](GetVirtualImage200ResponseVirtualImageOsType.md) |  | [optional] 
**MinRam** | Pointer to **NullableInt64** | Minimum required RAM in bytes | [optional] 
**MinRamGB** | Pointer to **NullableFloat64** | Minimum required RAM in gigabytes | [optional] 
**MinDisk** | Pointer to **NullableInt64** | Minimum required disk size in bytes | [optional] 
**MinDiskGB** | Pointer to **NullableString** | Minimum required disk size in gigabytes | [optional] 
**RawSize** | Pointer to **NullableInt64** | Size of image in bytes | [optional] 
**RawSizeGB** | Pointer to **NullableFloat32** | Size of image in gigabytes | [optional] 
**TrialVersion** | Pointer to **bool** | Is Trial Version? | [optional] 
**VirtioSupported** | Pointer to **bool** | VirtIO Drivers Loaded? | [optional] 
**Uefi** | Pointer to **NullableBool** | UEFI enabled? | [optional] 
**IsAutoJoinDomain** | Pointer to **bool** | Auto Join Domain? | [optional] 
**VmToolsInstalled** | Pointer to **bool** | VM Tools Installed? | [optional] 
**InstallAgent** | Pointer to **bool** | Install Agent? | [optional] 
**IsForceCustomization** | Pointer to **bool** | Force Guest Customization? | [optional] 
**IsSysprep** | Pointer to **bool** | Sysprep Enabled? | [optional] 
**FipsEnabled** | Pointer to **bool** | FIPS enabled? | [optional] 
**UserData** | Pointer to **NullableString** | Cloud-Init User Data, a bash script | [optional] 
**ConsoleKeymap** | Pointer to **NullableString** |  | [optional] 
**StorageProvider** | Pointer to [**GetVirtualImage200ResponseVirtualImageStorageProvider**](GetVirtualImage200ResponseVirtualImageStorageProvider.md) |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**Visibility** | Pointer to **string** | private or public | [optional] 
**Accounts** | Pointer to [**[]GetVirtualImage200ResponseVirtualImageAccountsInner**](GetVirtualImage200ResponseVirtualImageAccountsInner.md) |  | [optional] 
**Config** | Pointer to [**GetVirtualImage200ResponseVirtualImageConfig**](GetVirtualImage200ResponseVirtualImageConfig.md) |  | [optional] 
**Volumes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**StorageControllers** | Pointer to [**[]GetVirtualImage200ResponseVirtualImageStorageControllersInner**](GetVirtualImage200ResponseVirtualImageStorageControllersInner.md) |  | [optional] 
**NetworkInterfaces** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Tags** | Pointer to [**[]GetVirtualImage200ResponseVirtualImageTagsInner**](GetVirtualImage200ResponseVirtualImageTagsInner.md) | Metadata tags, Array of objects having a name and value | [optional] 
**Locations** | Pointer to [**[]GetVirtualImage200ResponseVirtualImageLocationsInner**](GetVirtualImage200ResponseVirtualImageLocationsInner.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewGetVirtualImage200ResponseVirtualImage

`func NewGetVirtualImage200ResponseVirtualImage() *GetVirtualImage200ResponseVirtualImage`

NewGetVirtualImage200ResponseVirtualImage instantiates a new GetVirtualImage200ResponseVirtualImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVirtualImage200ResponseVirtualImageWithDefaults

`func NewGetVirtualImage200ResponseVirtualImageWithDefaults() *GetVirtualImage200ResponseVirtualImage`

NewGetVirtualImage200ResponseVirtualImageWithDefaults instantiates a new GetVirtualImage200ResponseVirtualImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetVirtualImage200ResponseVirtualImage) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetVirtualImage200ResponseVirtualImage) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetVirtualImage200ResponseVirtualImage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetVirtualImage200ResponseVirtualImage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetVirtualImage200ResponseVirtualImage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetVirtualImage200ResponseVirtualImage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetVirtualImage200ResponseVirtualImage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetVirtualImage200ResponseVirtualImage) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetVirtualImage200ResponseVirtualImage) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetVirtualImage200ResponseVirtualImage) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetVirtualImage200ResponseVirtualImage) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLabels

`func (o *GetVirtualImage200ResponseVirtualImage) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetVirtualImage200ResponseVirtualImage) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetVirtualImage200ResponseVirtualImage) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetOwnerId

`func (o *GetVirtualImage200ResponseVirtualImage) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *GetVirtualImage200ResponseVirtualImage) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *GetVirtualImage200ResponseVirtualImage) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetTenant

`func (o *GetVirtualImage200ResponseVirtualImage) GetTenant() GetVirtualImage200ResponseVirtualImageTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetTenantOk() (*GetVirtualImage200ResponseVirtualImageTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *GetVirtualImage200ResponseVirtualImage) SetTenant(v GetVirtualImage200ResponseVirtualImageTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *GetVirtualImage200ResponseVirtualImage) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetImageType

`func (o *GetVirtualImage200ResponseVirtualImage) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetImageTypeOk() (*string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *GetVirtualImage200ResponseVirtualImage) SetImageType(v string)`

SetImageType sets ImageType field to given value.

### HasImageType

`func (o *GetVirtualImage200ResponseVirtualImage) HasImageType() bool`

HasImageType returns a boolean if a field has been set.

### GetUserUploaded

`func (o *GetVirtualImage200ResponseVirtualImage) GetUserUploaded() bool`

GetUserUploaded returns the UserUploaded field if non-nil, zero value otherwise.

### GetUserUploadedOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetUserUploadedOk() (*bool, bool)`

GetUserUploadedOk returns a tuple with the UserUploaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUploaded

`func (o *GetVirtualImage200ResponseVirtualImage) SetUserUploaded(v bool)`

SetUserUploaded sets UserUploaded field to given value.

### HasUserUploaded

`func (o *GetVirtualImage200ResponseVirtualImage) HasUserUploaded() bool`

HasUserUploaded returns a boolean if a field has been set.

### GetUserDefined

`func (o *GetVirtualImage200ResponseVirtualImage) GetUserDefined() bool`

GetUserDefined returns the UserDefined field if non-nil, zero value otherwise.

### GetUserDefinedOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetUserDefinedOk() (*bool, bool)`

GetUserDefinedOk returns a tuple with the UserDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefined

`func (o *GetVirtualImage200ResponseVirtualImage) SetUserDefined(v bool)`

SetUserDefined sets UserDefined field to given value.

### HasUserDefined

`func (o *GetVirtualImage200ResponseVirtualImage) HasUserDefined() bool`

HasUserDefined returns a boolean if a field has been set.

### GetSystemImage

`func (o *GetVirtualImage200ResponseVirtualImage) GetSystemImage() bool`

GetSystemImage returns the SystemImage field if non-nil, zero value otherwise.

### GetSystemImageOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetSystemImageOk() (*bool, bool)`

GetSystemImageOk returns a tuple with the SystemImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemImage

`func (o *GetVirtualImage200ResponseVirtualImage) SetSystemImage(v bool)`

SetSystemImage sets SystemImage field to given value.

### HasSystemImage

`func (o *GetVirtualImage200ResponseVirtualImage) HasSystemImage() bool`

HasSystemImage returns a boolean if a field has been set.

### GetIsCloudInit

`func (o *GetVirtualImage200ResponseVirtualImage) GetIsCloudInit() bool`

GetIsCloudInit returns the IsCloudInit field if non-nil, zero value otherwise.

### GetIsCloudInitOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetIsCloudInitOk() (*bool, bool)`

GetIsCloudInitOk returns a tuple with the IsCloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudInit

`func (o *GetVirtualImage200ResponseVirtualImage) SetIsCloudInit(v bool)`

SetIsCloudInit sets IsCloudInit field to given value.

### HasIsCloudInit

`func (o *GetVirtualImage200ResponseVirtualImage) HasIsCloudInit() bool`

HasIsCloudInit returns a boolean if a field has been set.

### GetSshUsername

`func (o *GetVirtualImage200ResponseVirtualImage) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *GetVirtualImage200ResponseVirtualImage) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *GetVirtualImage200ResponseVirtualImage) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### SetSshUsernameNil

`func (o *GetVirtualImage200ResponseVirtualImage) SetSshUsernameNil(b bool)`

 SetSshUsernameNil sets the value for SshUsername to be an explicit nil

### UnsetSshUsername
`func (o *GetVirtualImage200ResponseVirtualImage) UnsetSshUsername()`

UnsetSshUsername ensures that no value is present for SshUsername, not even an explicit nil
### GetSshPassword

`func (o *GetVirtualImage200ResponseVirtualImage) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *GetVirtualImage200ResponseVirtualImage) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *GetVirtualImage200ResponseVirtualImage) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### SetSshPasswordNil

`func (o *GetVirtualImage200ResponseVirtualImage) SetSshPasswordNil(b bool)`

 SetSshPasswordNil sets the value for SshPassword to be an explicit nil

### UnsetSshPassword
`func (o *GetVirtualImage200ResponseVirtualImage) UnsetSshPassword()`

UnsetSshPassword ensures that no value is present for SshPassword, not even an explicit nil
### GetSshPasswordHash

`func (o *GetVirtualImage200ResponseVirtualImage) GetSshPasswordHash() string`

GetSshPasswordHash returns the SshPasswordHash field if non-nil, zero value otherwise.

### GetSshPasswordHashOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetSshPasswordHashOk() (*string, bool)`

GetSshPasswordHashOk returns a tuple with the SshPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPasswordHash

`func (o *GetVirtualImage200ResponseVirtualImage) SetSshPasswordHash(v string)`

SetSshPasswordHash sets SshPasswordHash field to given value.

### HasSshPasswordHash

`func (o *GetVirtualImage200ResponseVirtualImage) HasSshPasswordHash() bool`

HasSshPasswordHash returns a boolean if a field has been set.

### SetSshPasswordHashNil

`func (o *GetVirtualImage200ResponseVirtualImage) SetSshPasswordHashNil(b bool)`

 SetSshPasswordHashNil sets the value for SshPasswordHash to be an explicit nil

### UnsetSshPasswordHash
`func (o *GetVirtualImage200ResponseVirtualImage) UnsetSshPasswordHash()`

UnsetSshPasswordHash ensures that no value is present for SshPasswordHash, not even an explicit nil
### GetSshKey

`func (o *GetVirtualImage200ResponseVirtualImage) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *GetVirtualImage200ResponseVirtualImage) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *GetVirtualImage200ResponseVirtualImage) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.

### SetSshKeyNil

`func (o *GetVirtualImage200ResponseVirtualImage) SetSshKeyNil(b bool)`

 SetSshKeyNil sets the value for SshKey to be an explicit nil

### UnsetSshKey
`func (o *GetVirtualImage200ResponseVirtualImage) UnsetSshKey()`

UnsetSshKey ensures that no value is present for SshKey, not even an explicit nil
### GetOsType

`func (o *GetVirtualImage200ResponseVirtualImage) GetOsType() GetVirtualImage200ResponseVirtualImageOsType`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetOsTypeOk() (*GetVirtualImage200ResponseVirtualImageOsType, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *GetVirtualImage200ResponseVirtualImage) SetOsType(v GetVirtualImage200ResponseVirtualImageOsType)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *GetVirtualImage200ResponseVirtualImage) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetMinRam

`func (o *GetVirtualImage200ResponseVirtualImage) GetMinRam() int64`

GetMinRam returns the MinRam field if non-nil, zero value otherwise.

### GetMinRamOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetMinRamOk() (*int64, bool)`

GetMinRamOk returns a tuple with the MinRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRam

`func (o *GetVirtualImage200ResponseVirtualImage) SetMinRam(v int64)`

SetMinRam sets MinRam field to given value.

### HasMinRam

`func (o *GetVirtualImage200ResponseVirtualImage) HasMinRam() bool`

HasMinRam returns a boolean if a field has been set.

### SetMinRamNil

`func (o *GetVirtualImage200ResponseVirtualImage) SetMinRamNil(b bool)`

 SetMinRamNil sets the value for MinRam to be an explicit nil

### UnsetMinRam
`func (o *GetVirtualImage200ResponseVirtualImage) UnsetMinRam()`

UnsetMinRam ensures that no value is present for MinRam, not even an explicit nil
### GetMinRamGB

`func (o *GetVirtualImage200ResponseVirtualImage) GetMinRamGB() float64`

GetMinRamGB returns the MinRamGB field if non-nil, zero value otherwise.

### GetMinRamGBOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetMinRamGBOk() (*float64, bool)`

GetMinRamGBOk returns a tuple with the MinRamGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRamGB

`func (o *GetVirtualImage200ResponseVirtualImage) SetMinRamGB(v float64)`

SetMinRamGB sets MinRamGB field to given value.

### HasMinRamGB

`func (o *GetVirtualImage200ResponseVirtualImage) HasMinRamGB() bool`

HasMinRamGB returns a boolean if a field has been set.

### SetMinRamGBNil

`func (o *GetVirtualImage200ResponseVirtualImage) SetMinRamGBNil(b bool)`

 SetMinRamGBNil sets the value for MinRamGB to be an explicit nil

### UnsetMinRamGB
`func (o *GetVirtualImage200ResponseVirtualImage) UnsetMinRamGB()`

UnsetMinRamGB ensures that no value is present for MinRamGB, not even an explicit nil
### GetMinDisk

`func (o *GetVirtualImage200ResponseVirtualImage) GetMinDisk() int64`

GetMinDisk returns the MinDisk field if non-nil, zero value otherwise.

### GetMinDiskOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetMinDiskOk() (*int64, bool)`

GetMinDiskOk returns a tuple with the MinDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDisk

`func (o *GetVirtualImage200ResponseVirtualImage) SetMinDisk(v int64)`

SetMinDisk sets MinDisk field to given value.

### HasMinDisk

`func (o *GetVirtualImage200ResponseVirtualImage) HasMinDisk() bool`

HasMinDisk returns a boolean if a field has been set.

### SetMinDiskNil

`func (o *GetVirtualImage200ResponseVirtualImage) SetMinDiskNil(b bool)`

 SetMinDiskNil sets the value for MinDisk to be an explicit nil

### UnsetMinDisk
`func (o *GetVirtualImage200ResponseVirtualImage) UnsetMinDisk()`

UnsetMinDisk ensures that no value is present for MinDisk, not even an explicit nil
### GetMinDiskGB

`func (o *GetVirtualImage200ResponseVirtualImage) GetMinDiskGB() string`

GetMinDiskGB returns the MinDiskGB field if non-nil, zero value otherwise.

### GetMinDiskGBOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetMinDiskGBOk() (*string, bool)`

GetMinDiskGBOk returns a tuple with the MinDiskGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDiskGB

`func (o *GetVirtualImage200ResponseVirtualImage) SetMinDiskGB(v string)`

SetMinDiskGB sets MinDiskGB field to given value.

### HasMinDiskGB

`func (o *GetVirtualImage200ResponseVirtualImage) HasMinDiskGB() bool`

HasMinDiskGB returns a boolean if a field has been set.

### SetMinDiskGBNil

`func (o *GetVirtualImage200ResponseVirtualImage) SetMinDiskGBNil(b bool)`

 SetMinDiskGBNil sets the value for MinDiskGB to be an explicit nil

### UnsetMinDiskGB
`func (o *GetVirtualImage200ResponseVirtualImage) UnsetMinDiskGB()`

UnsetMinDiskGB ensures that no value is present for MinDiskGB, not even an explicit nil
### GetRawSize

`func (o *GetVirtualImage200ResponseVirtualImage) GetRawSize() int64`

GetRawSize returns the RawSize field if non-nil, zero value otherwise.

### GetRawSizeOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetRawSizeOk() (*int64, bool)`

GetRawSizeOk returns a tuple with the RawSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSize

`func (o *GetVirtualImage200ResponseVirtualImage) SetRawSize(v int64)`

SetRawSize sets RawSize field to given value.

### HasRawSize

`func (o *GetVirtualImage200ResponseVirtualImage) HasRawSize() bool`

HasRawSize returns a boolean if a field has been set.

### SetRawSizeNil

`func (o *GetVirtualImage200ResponseVirtualImage) SetRawSizeNil(b bool)`

 SetRawSizeNil sets the value for RawSize to be an explicit nil

### UnsetRawSize
`func (o *GetVirtualImage200ResponseVirtualImage) UnsetRawSize()`

UnsetRawSize ensures that no value is present for RawSize, not even an explicit nil
### GetRawSizeGB

`func (o *GetVirtualImage200ResponseVirtualImage) GetRawSizeGB() float32`

GetRawSizeGB returns the RawSizeGB field if non-nil, zero value otherwise.

### GetRawSizeGBOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetRawSizeGBOk() (*float32, bool)`

GetRawSizeGBOk returns a tuple with the RawSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSizeGB

`func (o *GetVirtualImage200ResponseVirtualImage) SetRawSizeGB(v float32)`

SetRawSizeGB sets RawSizeGB field to given value.

### HasRawSizeGB

`func (o *GetVirtualImage200ResponseVirtualImage) HasRawSizeGB() bool`

HasRawSizeGB returns a boolean if a field has been set.

### SetRawSizeGBNil

`func (o *GetVirtualImage200ResponseVirtualImage) SetRawSizeGBNil(b bool)`

 SetRawSizeGBNil sets the value for RawSizeGB to be an explicit nil

### UnsetRawSizeGB
`func (o *GetVirtualImage200ResponseVirtualImage) UnsetRawSizeGB()`

UnsetRawSizeGB ensures that no value is present for RawSizeGB, not even an explicit nil
### GetTrialVersion

`func (o *GetVirtualImage200ResponseVirtualImage) GetTrialVersion() bool`

GetTrialVersion returns the TrialVersion field if non-nil, zero value otherwise.

### GetTrialVersionOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetTrialVersionOk() (*bool, bool)`

GetTrialVersionOk returns a tuple with the TrialVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialVersion

`func (o *GetVirtualImage200ResponseVirtualImage) SetTrialVersion(v bool)`

SetTrialVersion sets TrialVersion field to given value.

### HasTrialVersion

`func (o *GetVirtualImage200ResponseVirtualImage) HasTrialVersion() bool`

HasTrialVersion returns a boolean if a field has been set.

### GetVirtioSupported

`func (o *GetVirtualImage200ResponseVirtualImage) GetVirtioSupported() bool`

GetVirtioSupported returns the VirtioSupported field if non-nil, zero value otherwise.

### GetVirtioSupportedOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetVirtioSupportedOk() (*bool, bool)`

GetVirtioSupportedOk returns a tuple with the VirtioSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtioSupported

`func (o *GetVirtualImage200ResponseVirtualImage) SetVirtioSupported(v bool)`

SetVirtioSupported sets VirtioSupported field to given value.

### HasVirtioSupported

`func (o *GetVirtualImage200ResponseVirtualImage) HasVirtioSupported() bool`

HasVirtioSupported returns a boolean if a field has been set.

### GetUefi

`func (o *GetVirtualImage200ResponseVirtualImage) GetUefi() bool`

GetUefi returns the Uefi field if non-nil, zero value otherwise.

### GetUefiOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetUefiOk() (*bool, bool)`

GetUefiOk returns a tuple with the Uefi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUefi

`func (o *GetVirtualImage200ResponseVirtualImage) SetUefi(v bool)`

SetUefi sets Uefi field to given value.

### HasUefi

`func (o *GetVirtualImage200ResponseVirtualImage) HasUefi() bool`

HasUefi returns a boolean if a field has been set.

### SetUefiNil

`func (o *GetVirtualImage200ResponseVirtualImage) SetUefiNil(b bool)`

 SetUefiNil sets the value for Uefi to be an explicit nil

### UnsetUefi
`func (o *GetVirtualImage200ResponseVirtualImage) UnsetUefi()`

UnsetUefi ensures that no value is present for Uefi, not even an explicit nil
### GetIsAutoJoinDomain

`func (o *GetVirtualImage200ResponseVirtualImage) GetIsAutoJoinDomain() bool`

GetIsAutoJoinDomain returns the IsAutoJoinDomain field if non-nil, zero value otherwise.

### GetIsAutoJoinDomainOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetIsAutoJoinDomainOk() (*bool, bool)`

GetIsAutoJoinDomainOk returns a tuple with the IsAutoJoinDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoJoinDomain

`func (o *GetVirtualImage200ResponseVirtualImage) SetIsAutoJoinDomain(v bool)`

SetIsAutoJoinDomain sets IsAutoJoinDomain field to given value.

### HasIsAutoJoinDomain

`func (o *GetVirtualImage200ResponseVirtualImage) HasIsAutoJoinDomain() bool`

HasIsAutoJoinDomain returns a boolean if a field has been set.

### GetVmToolsInstalled

`func (o *GetVirtualImage200ResponseVirtualImage) GetVmToolsInstalled() bool`

GetVmToolsInstalled returns the VmToolsInstalled field if non-nil, zero value otherwise.

### GetVmToolsInstalledOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetVmToolsInstalledOk() (*bool, bool)`

GetVmToolsInstalledOk returns a tuple with the VmToolsInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmToolsInstalled

`func (o *GetVirtualImage200ResponseVirtualImage) SetVmToolsInstalled(v bool)`

SetVmToolsInstalled sets VmToolsInstalled field to given value.

### HasVmToolsInstalled

`func (o *GetVirtualImage200ResponseVirtualImage) HasVmToolsInstalled() bool`

HasVmToolsInstalled returns a boolean if a field has been set.

### GetInstallAgent

`func (o *GetVirtualImage200ResponseVirtualImage) GetInstallAgent() bool`

GetInstallAgent returns the InstallAgent field if non-nil, zero value otherwise.

### GetInstallAgentOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetInstallAgentOk() (*bool, bool)`

GetInstallAgentOk returns a tuple with the InstallAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAgent

`func (o *GetVirtualImage200ResponseVirtualImage) SetInstallAgent(v bool)`

SetInstallAgent sets InstallAgent field to given value.

### HasInstallAgent

`func (o *GetVirtualImage200ResponseVirtualImage) HasInstallAgent() bool`

HasInstallAgent returns a boolean if a field has been set.

### GetIsForceCustomization

`func (o *GetVirtualImage200ResponseVirtualImage) GetIsForceCustomization() bool`

GetIsForceCustomization returns the IsForceCustomization field if non-nil, zero value otherwise.

### GetIsForceCustomizationOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetIsForceCustomizationOk() (*bool, bool)`

GetIsForceCustomizationOk returns a tuple with the IsForceCustomization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForceCustomization

`func (o *GetVirtualImage200ResponseVirtualImage) SetIsForceCustomization(v bool)`

SetIsForceCustomization sets IsForceCustomization field to given value.

### HasIsForceCustomization

`func (o *GetVirtualImage200ResponseVirtualImage) HasIsForceCustomization() bool`

HasIsForceCustomization returns a boolean if a field has been set.

### GetIsSysprep

`func (o *GetVirtualImage200ResponseVirtualImage) GetIsSysprep() bool`

GetIsSysprep returns the IsSysprep field if non-nil, zero value otherwise.

### GetIsSysprepOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetIsSysprepOk() (*bool, bool)`

GetIsSysprepOk returns a tuple with the IsSysprep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSysprep

`func (o *GetVirtualImage200ResponseVirtualImage) SetIsSysprep(v bool)`

SetIsSysprep sets IsSysprep field to given value.

### HasIsSysprep

`func (o *GetVirtualImage200ResponseVirtualImage) HasIsSysprep() bool`

HasIsSysprep returns a boolean if a field has been set.

### GetFipsEnabled

`func (o *GetVirtualImage200ResponseVirtualImage) GetFipsEnabled() bool`

GetFipsEnabled returns the FipsEnabled field if non-nil, zero value otherwise.

### GetFipsEnabledOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetFipsEnabledOk() (*bool, bool)`

GetFipsEnabledOk returns a tuple with the FipsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFipsEnabled

`func (o *GetVirtualImage200ResponseVirtualImage) SetFipsEnabled(v bool)`

SetFipsEnabled sets FipsEnabled field to given value.

### HasFipsEnabled

`func (o *GetVirtualImage200ResponseVirtualImage) HasFipsEnabled() bool`

HasFipsEnabled returns a boolean if a field has been set.

### GetUserData

`func (o *GetVirtualImage200ResponseVirtualImage) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *GetVirtualImage200ResponseVirtualImage) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *GetVirtualImage200ResponseVirtualImage) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### SetUserDataNil

`func (o *GetVirtualImage200ResponseVirtualImage) SetUserDataNil(b bool)`

 SetUserDataNil sets the value for UserData to be an explicit nil

### UnsetUserData
`func (o *GetVirtualImage200ResponseVirtualImage) UnsetUserData()`

UnsetUserData ensures that no value is present for UserData, not even an explicit nil
### GetConsoleKeymap

`func (o *GetVirtualImage200ResponseVirtualImage) GetConsoleKeymap() string`

GetConsoleKeymap returns the ConsoleKeymap field if non-nil, zero value otherwise.

### GetConsoleKeymapOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetConsoleKeymapOk() (*string, bool)`

GetConsoleKeymapOk returns a tuple with the ConsoleKeymap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleKeymap

`func (o *GetVirtualImage200ResponseVirtualImage) SetConsoleKeymap(v string)`

SetConsoleKeymap sets ConsoleKeymap field to given value.

### HasConsoleKeymap

`func (o *GetVirtualImage200ResponseVirtualImage) HasConsoleKeymap() bool`

HasConsoleKeymap returns a boolean if a field has been set.

### SetConsoleKeymapNil

`func (o *GetVirtualImage200ResponseVirtualImage) SetConsoleKeymapNil(b bool)`

 SetConsoleKeymapNil sets the value for ConsoleKeymap to be an explicit nil

### UnsetConsoleKeymap
`func (o *GetVirtualImage200ResponseVirtualImage) UnsetConsoleKeymap()`

UnsetConsoleKeymap ensures that no value is present for ConsoleKeymap, not even an explicit nil
### GetStorageProvider

`func (o *GetVirtualImage200ResponseVirtualImage) GetStorageProvider() GetVirtualImage200ResponseVirtualImageStorageProvider`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetStorageProviderOk() (*GetVirtualImage200ResponseVirtualImageStorageProvider, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *GetVirtualImage200ResponseVirtualImage) SetStorageProvider(v GetVirtualImage200ResponseVirtualImageStorageProvider)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *GetVirtualImage200ResponseVirtualImage) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### GetExternalId

`func (o *GetVirtualImage200ResponseVirtualImage) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetVirtualImage200ResponseVirtualImage) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetVirtualImage200ResponseVirtualImage) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *GetVirtualImage200ResponseVirtualImage) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *GetVirtualImage200ResponseVirtualImage) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetVisibility

`func (o *GetVirtualImage200ResponseVirtualImage) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetVirtualImage200ResponseVirtualImage) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetVirtualImage200ResponseVirtualImage) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAccounts

`func (o *GetVirtualImage200ResponseVirtualImage) GetAccounts() []GetVirtualImage200ResponseVirtualImageAccountsInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetAccountsOk() (*[]GetVirtualImage200ResponseVirtualImageAccountsInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *GetVirtualImage200ResponseVirtualImage) SetAccounts(v []GetVirtualImage200ResponseVirtualImageAccountsInner)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *GetVirtualImage200ResponseVirtualImage) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetConfig

`func (o *GetVirtualImage200ResponseVirtualImage) GetConfig() GetVirtualImage200ResponseVirtualImageConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetConfigOk() (*GetVirtualImage200ResponseVirtualImageConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetVirtualImage200ResponseVirtualImage) SetConfig(v GetVirtualImage200ResponseVirtualImageConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetVirtualImage200ResponseVirtualImage) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetVolumes

`func (o *GetVirtualImage200ResponseVirtualImage) GetVolumes() []map[string]interface{}`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetVolumesOk() (*[]map[string]interface{}, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *GetVirtualImage200ResponseVirtualImage) SetVolumes(v []map[string]interface{})`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *GetVirtualImage200ResponseVirtualImage) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetStorageControllers

`func (o *GetVirtualImage200ResponseVirtualImage) GetStorageControllers() []GetVirtualImage200ResponseVirtualImageStorageControllersInner`

GetStorageControllers returns the StorageControllers field if non-nil, zero value otherwise.

### GetStorageControllersOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetStorageControllersOk() (*[]GetVirtualImage200ResponseVirtualImageStorageControllersInner, bool)`

GetStorageControllersOk returns a tuple with the StorageControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllers

`func (o *GetVirtualImage200ResponseVirtualImage) SetStorageControllers(v []GetVirtualImage200ResponseVirtualImageStorageControllersInner)`

SetStorageControllers sets StorageControllers field to given value.

### HasStorageControllers

`func (o *GetVirtualImage200ResponseVirtualImage) HasStorageControllers() bool`

HasStorageControllers returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *GetVirtualImage200ResponseVirtualImage) GetNetworkInterfaces() []map[string]interface{}`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetNetworkInterfacesOk() (*[]map[string]interface{}, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *GetVirtualImage200ResponseVirtualImage) SetNetworkInterfaces(v []map[string]interface{})`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *GetVirtualImage200ResponseVirtualImage) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetTags

`func (o *GetVirtualImage200ResponseVirtualImage) GetTags() []GetVirtualImage200ResponseVirtualImageTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetTagsOk() (*[]GetVirtualImage200ResponseVirtualImageTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetVirtualImage200ResponseVirtualImage) SetTags(v []GetVirtualImage200ResponseVirtualImageTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetVirtualImage200ResponseVirtualImage) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLocations

`func (o *GetVirtualImage200ResponseVirtualImage) GetLocations() []GetVirtualImage200ResponseVirtualImageLocationsInner`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetLocationsOk() (*[]GetVirtualImage200ResponseVirtualImageLocationsInner, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *GetVirtualImage200ResponseVirtualImage) SetLocations(v []GetVirtualImage200ResponseVirtualImageLocationsInner)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *GetVirtualImage200ResponseVirtualImage) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetVirtualImage200ResponseVirtualImage) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetVirtualImage200ResponseVirtualImage) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetVirtualImage200ResponseVirtualImage) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetVirtualImage200ResponseVirtualImage) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetVirtualImage200ResponseVirtualImage) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetVirtualImage200ResponseVirtualImage) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetStatus

`func (o *GetVirtualImage200ResponseVirtualImage) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetVirtualImage200ResponseVirtualImage) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetVirtualImage200ResponseVirtualImage) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetVirtualImage200ResponseVirtualImage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


