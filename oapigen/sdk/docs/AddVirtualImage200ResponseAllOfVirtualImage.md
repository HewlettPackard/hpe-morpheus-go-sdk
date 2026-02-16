# AddVirtualImage200ResponseAllOfVirtualImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** | A name for the virtual image | [optional] 
**Description** | Pointer to **NullableString** | A description for the virtual image | [optional] 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**OwnerId** | Pointer to **int64** | Owner of the image | [optional] 
**Tenant** | Pointer to [**AddVirtualImage200ResponseAllOfVirtualImageTenant**](AddVirtualImage200ResponseAllOfVirtualImageTenant.md) |  | [optional] 
**ImageType** | Pointer to **string** | Code of image type. eg. vmware, ami, etc. | [optional] 
**UserUploaded** | Pointer to **bool** | Is uploaded by an user? | [optional] 
**UserDefined** | Pointer to **bool** | Is defined by an user? | [optional] 
**SystemImage** | Pointer to **bool** | Is created by system? | [optional] 
**IsCloudInit** | Pointer to **bool** | Cloud Init Enabled? | [optional] 
**SshUsername** | Pointer to **NullableString** | SSH Username | [optional] 
**SshPassword** | Pointer to **NullableString** |  | [optional] 
**SshPasswordHash** | Pointer to **NullableString** |  | [optional] 
**SshKey** | Pointer to **NullableString** | SSH Key | [optional] 
**OsType** | Pointer to [**AddVirtualImage200ResponseAllOfVirtualImageOsType**](AddVirtualImage200ResponseAllOfVirtualImageOsType.md) |  | [optional] 
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
**StorageProvider** | Pointer to [**AddVirtualImage200ResponseAllOfVirtualImageStorageProvider**](AddVirtualImage200ResponseAllOfVirtualImageStorageProvider.md) |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**Visibility** | Pointer to **string** | private or public | [optional] 
**Accounts** | Pointer to [**[]AddVirtualImage200ResponseAllOfVirtualImageAccountsInner**](AddVirtualImage200ResponseAllOfVirtualImageAccountsInner.md) |  | [optional] 
**Config** | Pointer to [**AddVirtualImage200ResponseAllOfVirtualImageConfig**](AddVirtualImage200ResponseAllOfVirtualImageConfig.md) |  | [optional] 
**Volumes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**StorageControllers** | Pointer to [**[]AddVirtualImage200ResponseAllOfVirtualImageStorageControllersInner**](AddVirtualImage200ResponseAllOfVirtualImageStorageControllersInner.md) |  | [optional] 
**NetworkInterfaces** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Tags** | Pointer to [**[]AddVirtualImage200ResponseAllOfVirtualImageTagsInner**](AddVirtualImage200ResponseAllOfVirtualImageTagsInner.md) | Metadata tags, Array of objects having a name and value | [optional] 
**Locations** | Pointer to [**[]AddVirtualImage200ResponseAllOfVirtualImageLocationsInner**](AddVirtualImage200ResponseAllOfVirtualImageLocationsInner.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewAddVirtualImage200ResponseAllOfVirtualImage

`func NewAddVirtualImage200ResponseAllOfVirtualImage() *AddVirtualImage200ResponseAllOfVirtualImage`

NewAddVirtualImage200ResponseAllOfVirtualImage instantiates a new AddVirtualImage200ResponseAllOfVirtualImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddVirtualImage200ResponseAllOfVirtualImageWithDefaults

`func NewAddVirtualImage200ResponseAllOfVirtualImageWithDefaults() *AddVirtualImage200ResponseAllOfVirtualImage`

NewAddVirtualImage200ResponseAllOfVirtualImageWithDefaults instantiates a new AddVirtualImage200ResponseAllOfVirtualImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AddVirtualImage200ResponseAllOfVirtualImage) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLabels

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetOwnerId

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetTenant

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetTenant() AddVirtualImage200ResponseAllOfVirtualImageTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetTenantOk() (*AddVirtualImage200ResponseAllOfVirtualImageTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetTenant(v AddVirtualImage200ResponseAllOfVirtualImageTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetImageType

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetImageTypeOk() (*string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetImageType(v string)`

SetImageType sets ImageType field to given value.

### HasImageType

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasImageType() bool`

HasImageType returns a boolean if a field has been set.

### GetUserUploaded

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetUserUploaded() bool`

GetUserUploaded returns the UserUploaded field if non-nil, zero value otherwise.

### GetUserUploadedOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetUserUploadedOk() (*bool, bool)`

GetUserUploadedOk returns a tuple with the UserUploaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUploaded

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetUserUploaded(v bool)`

SetUserUploaded sets UserUploaded field to given value.

### HasUserUploaded

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasUserUploaded() bool`

HasUserUploaded returns a boolean if a field has been set.

### GetUserDefined

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetUserDefined() bool`

GetUserDefined returns the UserDefined field if non-nil, zero value otherwise.

### GetUserDefinedOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetUserDefinedOk() (*bool, bool)`

GetUserDefinedOk returns a tuple with the UserDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefined

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetUserDefined(v bool)`

SetUserDefined sets UserDefined field to given value.

### HasUserDefined

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasUserDefined() bool`

HasUserDefined returns a boolean if a field has been set.

### GetSystemImage

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetSystemImage() bool`

GetSystemImage returns the SystemImage field if non-nil, zero value otherwise.

### GetSystemImageOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetSystemImageOk() (*bool, bool)`

GetSystemImageOk returns a tuple with the SystemImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemImage

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetSystemImage(v bool)`

SetSystemImage sets SystemImage field to given value.

### HasSystemImage

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasSystemImage() bool`

HasSystemImage returns a boolean if a field has been set.

### GetIsCloudInit

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetIsCloudInit() bool`

GetIsCloudInit returns the IsCloudInit field if non-nil, zero value otherwise.

### GetIsCloudInitOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetIsCloudInitOk() (*bool, bool)`

GetIsCloudInitOk returns a tuple with the IsCloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudInit

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetIsCloudInit(v bool)`

SetIsCloudInit sets IsCloudInit field to given value.

### HasIsCloudInit

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasIsCloudInit() bool`

HasIsCloudInit returns a boolean if a field has been set.

### GetSshUsername

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### SetSshUsernameNil

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetSshUsernameNil(b bool)`

 SetSshUsernameNil sets the value for SshUsername to be an explicit nil

### UnsetSshUsername
`func (o *AddVirtualImage200ResponseAllOfVirtualImage) UnsetSshUsername()`

UnsetSshUsername ensures that no value is present for SshUsername, not even an explicit nil
### GetSshPassword

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### SetSshPasswordNil

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetSshPasswordNil(b bool)`

 SetSshPasswordNil sets the value for SshPassword to be an explicit nil

### UnsetSshPassword
`func (o *AddVirtualImage200ResponseAllOfVirtualImage) UnsetSshPassword()`

UnsetSshPassword ensures that no value is present for SshPassword, not even an explicit nil
### GetSshPasswordHash

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetSshPasswordHash() string`

GetSshPasswordHash returns the SshPasswordHash field if non-nil, zero value otherwise.

### GetSshPasswordHashOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetSshPasswordHashOk() (*string, bool)`

GetSshPasswordHashOk returns a tuple with the SshPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPasswordHash

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetSshPasswordHash(v string)`

SetSshPasswordHash sets SshPasswordHash field to given value.

### HasSshPasswordHash

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasSshPasswordHash() bool`

HasSshPasswordHash returns a boolean if a field has been set.

### SetSshPasswordHashNil

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetSshPasswordHashNil(b bool)`

 SetSshPasswordHashNil sets the value for SshPasswordHash to be an explicit nil

### UnsetSshPasswordHash
`func (o *AddVirtualImage200ResponseAllOfVirtualImage) UnsetSshPasswordHash()`

UnsetSshPasswordHash ensures that no value is present for SshPasswordHash, not even an explicit nil
### GetSshKey

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.

### SetSshKeyNil

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetSshKeyNil(b bool)`

 SetSshKeyNil sets the value for SshKey to be an explicit nil

### UnsetSshKey
`func (o *AddVirtualImage200ResponseAllOfVirtualImage) UnsetSshKey()`

UnsetSshKey ensures that no value is present for SshKey, not even an explicit nil
### GetOsType

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetOsType() AddVirtualImage200ResponseAllOfVirtualImageOsType`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetOsTypeOk() (*AddVirtualImage200ResponseAllOfVirtualImageOsType, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetOsType(v AddVirtualImage200ResponseAllOfVirtualImageOsType)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetMinRam

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetMinRam() int64`

GetMinRam returns the MinRam field if non-nil, zero value otherwise.

### GetMinRamOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetMinRamOk() (*int64, bool)`

GetMinRamOk returns a tuple with the MinRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRam

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetMinRam(v int64)`

SetMinRam sets MinRam field to given value.

### HasMinRam

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasMinRam() bool`

HasMinRam returns a boolean if a field has been set.

### SetMinRamNil

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetMinRamNil(b bool)`

 SetMinRamNil sets the value for MinRam to be an explicit nil

### UnsetMinRam
`func (o *AddVirtualImage200ResponseAllOfVirtualImage) UnsetMinRam()`

UnsetMinRam ensures that no value is present for MinRam, not even an explicit nil
### GetMinRamGB

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetMinRamGB() float64`

GetMinRamGB returns the MinRamGB field if non-nil, zero value otherwise.

### GetMinRamGBOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetMinRamGBOk() (*float64, bool)`

GetMinRamGBOk returns a tuple with the MinRamGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRamGB

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetMinRamGB(v float64)`

SetMinRamGB sets MinRamGB field to given value.

### HasMinRamGB

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasMinRamGB() bool`

HasMinRamGB returns a boolean if a field has been set.

### SetMinRamGBNil

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetMinRamGBNil(b bool)`

 SetMinRamGBNil sets the value for MinRamGB to be an explicit nil

### UnsetMinRamGB
`func (o *AddVirtualImage200ResponseAllOfVirtualImage) UnsetMinRamGB()`

UnsetMinRamGB ensures that no value is present for MinRamGB, not even an explicit nil
### GetMinDisk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetMinDisk() int64`

GetMinDisk returns the MinDisk field if non-nil, zero value otherwise.

### GetMinDiskOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetMinDiskOk() (*int64, bool)`

GetMinDiskOk returns a tuple with the MinDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDisk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetMinDisk(v int64)`

SetMinDisk sets MinDisk field to given value.

### HasMinDisk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasMinDisk() bool`

HasMinDisk returns a boolean if a field has been set.

### SetMinDiskNil

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetMinDiskNil(b bool)`

 SetMinDiskNil sets the value for MinDisk to be an explicit nil

### UnsetMinDisk
`func (o *AddVirtualImage200ResponseAllOfVirtualImage) UnsetMinDisk()`

UnsetMinDisk ensures that no value is present for MinDisk, not even an explicit nil
### GetMinDiskGB

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetMinDiskGB() string`

GetMinDiskGB returns the MinDiskGB field if non-nil, zero value otherwise.

### GetMinDiskGBOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetMinDiskGBOk() (*string, bool)`

GetMinDiskGBOk returns a tuple with the MinDiskGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDiskGB

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetMinDiskGB(v string)`

SetMinDiskGB sets MinDiskGB field to given value.

### HasMinDiskGB

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasMinDiskGB() bool`

HasMinDiskGB returns a boolean if a field has been set.

### SetMinDiskGBNil

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetMinDiskGBNil(b bool)`

 SetMinDiskGBNil sets the value for MinDiskGB to be an explicit nil

### UnsetMinDiskGB
`func (o *AddVirtualImage200ResponseAllOfVirtualImage) UnsetMinDiskGB()`

UnsetMinDiskGB ensures that no value is present for MinDiskGB, not even an explicit nil
### GetRawSize

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetRawSize() int64`

GetRawSize returns the RawSize field if non-nil, zero value otherwise.

### GetRawSizeOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetRawSizeOk() (*int64, bool)`

GetRawSizeOk returns a tuple with the RawSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSize

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetRawSize(v int64)`

SetRawSize sets RawSize field to given value.

### HasRawSize

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasRawSize() bool`

HasRawSize returns a boolean if a field has been set.

### SetRawSizeNil

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetRawSizeNil(b bool)`

 SetRawSizeNil sets the value for RawSize to be an explicit nil

### UnsetRawSize
`func (o *AddVirtualImage200ResponseAllOfVirtualImage) UnsetRawSize()`

UnsetRawSize ensures that no value is present for RawSize, not even an explicit nil
### GetRawSizeGB

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetRawSizeGB() float32`

GetRawSizeGB returns the RawSizeGB field if non-nil, zero value otherwise.

### GetRawSizeGBOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetRawSizeGBOk() (*float32, bool)`

GetRawSizeGBOk returns a tuple with the RawSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSizeGB

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetRawSizeGB(v float32)`

SetRawSizeGB sets RawSizeGB field to given value.

### HasRawSizeGB

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasRawSizeGB() bool`

HasRawSizeGB returns a boolean if a field has been set.

### SetRawSizeGBNil

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetRawSizeGBNil(b bool)`

 SetRawSizeGBNil sets the value for RawSizeGB to be an explicit nil

### UnsetRawSizeGB
`func (o *AddVirtualImage200ResponseAllOfVirtualImage) UnsetRawSizeGB()`

UnsetRawSizeGB ensures that no value is present for RawSizeGB, not even an explicit nil
### GetTrialVersion

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetTrialVersion() bool`

GetTrialVersion returns the TrialVersion field if non-nil, zero value otherwise.

### GetTrialVersionOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetTrialVersionOk() (*bool, bool)`

GetTrialVersionOk returns a tuple with the TrialVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialVersion

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetTrialVersion(v bool)`

SetTrialVersion sets TrialVersion field to given value.

### HasTrialVersion

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasTrialVersion() bool`

HasTrialVersion returns a boolean if a field has been set.

### GetVirtioSupported

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetVirtioSupported() bool`

GetVirtioSupported returns the VirtioSupported field if non-nil, zero value otherwise.

### GetVirtioSupportedOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetVirtioSupportedOk() (*bool, bool)`

GetVirtioSupportedOk returns a tuple with the VirtioSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtioSupported

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetVirtioSupported(v bool)`

SetVirtioSupported sets VirtioSupported field to given value.

### HasVirtioSupported

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasVirtioSupported() bool`

HasVirtioSupported returns a boolean if a field has been set.

### GetUefi

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetUefi() bool`

GetUefi returns the Uefi field if non-nil, zero value otherwise.

### GetUefiOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetUefiOk() (*bool, bool)`

GetUefiOk returns a tuple with the Uefi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUefi

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetUefi(v bool)`

SetUefi sets Uefi field to given value.

### HasUefi

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasUefi() bool`

HasUefi returns a boolean if a field has been set.

### SetUefiNil

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetUefiNil(b bool)`

 SetUefiNil sets the value for Uefi to be an explicit nil

### UnsetUefi
`func (o *AddVirtualImage200ResponseAllOfVirtualImage) UnsetUefi()`

UnsetUefi ensures that no value is present for Uefi, not even an explicit nil
### GetIsAutoJoinDomain

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetIsAutoJoinDomain() bool`

GetIsAutoJoinDomain returns the IsAutoJoinDomain field if non-nil, zero value otherwise.

### GetIsAutoJoinDomainOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetIsAutoJoinDomainOk() (*bool, bool)`

GetIsAutoJoinDomainOk returns a tuple with the IsAutoJoinDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoJoinDomain

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetIsAutoJoinDomain(v bool)`

SetIsAutoJoinDomain sets IsAutoJoinDomain field to given value.

### HasIsAutoJoinDomain

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasIsAutoJoinDomain() bool`

HasIsAutoJoinDomain returns a boolean if a field has been set.

### GetVmToolsInstalled

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetVmToolsInstalled() bool`

GetVmToolsInstalled returns the VmToolsInstalled field if non-nil, zero value otherwise.

### GetVmToolsInstalledOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetVmToolsInstalledOk() (*bool, bool)`

GetVmToolsInstalledOk returns a tuple with the VmToolsInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmToolsInstalled

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetVmToolsInstalled(v bool)`

SetVmToolsInstalled sets VmToolsInstalled field to given value.

### HasVmToolsInstalled

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasVmToolsInstalled() bool`

HasVmToolsInstalled returns a boolean if a field has been set.

### GetInstallAgent

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetInstallAgent() bool`

GetInstallAgent returns the InstallAgent field if non-nil, zero value otherwise.

### GetInstallAgentOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetInstallAgentOk() (*bool, bool)`

GetInstallAgentOk returns a tuple with the InstallAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAgent

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetInstallAgent(v bool)`

SetInstallAgent sets InstallAgent field to given value.

### HasInstallAgent

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasInstallAgent() bool`

HasInstallAgent returns a boolean if a field has been set.

### GetIsForceCustomization

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetIsForceCustomization() bool`

GetIsForceCustomization returns the IsForceCustomization field if non-nil, zero value otherwise.

### GetIsForceCustomizationOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetIsForceCustomizationOk() (*bool, bool)`

GetIsForceCustomizationOk returns a tuple with the IsForceCustomization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForceCustomization

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetIsForceCustomization(v bool)`

SetIsForceCustomization sets IsForceCustomization field to given value.

### HasIsForceCustomization

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasIsForceCustomization() bool`

HasIsForceCustomization returns a boolean if a field has been set.

### GetIsSysprep

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetIsSysprep() bool`

GetIsSysprep returns the IsSysprep field if non-nil, zero value otherwise.

### GetIsSysprepOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetIsSysprepOk() (*bool, bool)`

GetIsSysprepOk returns a tuple with the IsSysprep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSysprep

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetIsSysprep(v bool)`

SetIsSysprep sets IsSysprep field to given value.

### HasIsSysprep

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasIsSysprep() bool`

HasIsSysprep returns a boolean if a field has been set.

### GetFipsEnabled

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetFipsEnabled() bool`

GetFipsEnabled returns the FipsEnabled field if non-nil, zero value otherwise.

### GetFipsEnabledOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetFipsEnabledOk() (*bool, bool)`

GetFipsEnabledOk returns a tuple with the FipsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFipsEnabled

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetFipsEnabled(v bool)`

SetFipsEnabled sets FipsEnabled field to given value.

### HasFipsEnabled

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasFipsEnabled() bool`

HasFipsEnabled returns a boolean if a field has been set.

### GetUserData

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### SetUserDataNil

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetUserDataNil(b bool)`

 SetUserDataNil sets the value for UserData to be an explicit nil

### UnsetUserData
`func (o *AddVirtualImage200ResponseAllOfVirtualImage) UnsetUserData()`

UnsetUserData ensures that no value is present for UserData, not even an explicit nil
### GetConsoleKeymap

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetConsoleKeymap() string`

GetConsoleKeymap returns the ConsoleKeymap field if non-nil, zero value otherwise.

### GetConsoleKeymapOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetConsoleKeymapOk() (*string, bool)`

GetConsoleKeymapOk returns a tuple with the ConsoleKeymap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleKeymap

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetConsoleKeymap(v string)`

SetConsoleKeymap sets ConsoleKeymap field to given value.

### HasConsoleKeymap

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasConsoleKeymap() bool`

HasConsoleKeymap returns a boolean if a field has been set.

### SetConsoleKeymapNil

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetConsoleKeymapNil(b bool)`

 SetConsoleKeymapNil sets the value for ConsoleKeymap to be an explicit nil

### UnsetConsoleKeymap
`func (o *AddVirtualImage200ResponseAllOfVirtualImage) UnsetConsoleKeymap()`

UnsetConsoleKeymap ensures that no value is present for ConsoleKeymap, not even an explicit nil
### GetStorageProvider

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetStorageProvider() AddVirtualImage200ResponseAllOfVirtualImageStorageProvider`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetStorageProviderOk() (*AddVirtualImage200ResponseAllOfVirtualImageStorageProvider, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetStorageProvider(v AddVirtualImage200ResponseAllOfVirtualImageStorageProvider)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### GetExternalId

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *AddVirtualImage200ResponseAllOfVirtualImage) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetVisibility

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAccounts

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetAccounts() []AddVirtualImage200ResponseAllOfVirtualImageAccountsInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetAccountsOk() (*[]AddVirtualImage200ResponseAllOfVirtualImageAccountsInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetAccounts(v []AddVirtualImage200ResponseAllOfVirtualImageAccountsInner)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetConfig

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetConfig() AddVirtualImage200ResponseAllOfVirtualImageConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetConfigOk() (*AddVirtualImage200ResponseAllOfVirtualImageConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetConfig(v AddVirtualImage200ResponseAllOfVirtualImageConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetVolumes

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetVolumes() []map[string]interface{}`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetVolumesOk() (*[]map[string]interface{}, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetVolumes(v []map[string]interface{})`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetStorageControllers

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetStorageControllers() []AddVirtualImage200ResponseAllOfVirtualImageStorageControllersInner`

GetStorageControllers returns the StorageControllers field if non-nil, zero value otherwise.

### GetStorageControllersOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetStorageControllersOk() (*[]AddVirtualImage200ResponseAllOfVirtualImageStorageControllersInner, bool)`

GetStorageControllersOk returns a tuple with the StorageControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllers

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetStorageControllers(v []AddVirtualImage200ResponseAllOfVirtualImageStorageControllersInner)`

SetStorageControllers sets StorageControllers field to given value.

### HasStorageControllers

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasStorageControllers() bool`

HasStorageControllers returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetNetworkInterfaces() []map[string]interface{}`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetNetworkInterfacesOk() (*[]map[string]interface{}, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetNetworkInterfaces(v []map[string]interface{})`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetTags

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetTags() []AddVirtualImage200ResponseAllOfVirtualImageTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetTagsOk() (*[]AddVirtualImage200ResponseAllOfVirtualImageTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetTags(v []AddVirtualImage200ResponseAllOfVirtualImageTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLocations

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetLocations() []AddVirtualImage200ResponseAllOfVirtualImageLocationsInner`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetLocationsOk() (*[]AddVirtualImage200ResponseAllOfVirtualImageLocationsInner, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetLocations(v []AddVirtualImage200ResponseAllOfVirtualImageLocationsInner)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetDateCreated

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetStatus

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddVirtualImage200ResponseAllOfVirtualImage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


