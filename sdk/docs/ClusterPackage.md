# ClusterPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Account** | Pointer to **int64** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**RepeatInstall** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**PackageType** | Pointer to **string** |  | [optional] 
**PackageVersion** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**IconPath** | Pointer to **string** |  | [optional] 
**ImagePath** | Pointer to **string** |  | [optional] 
**DarkImagePath** | Pointer to **string** |  | [optional] 
**SpecTemplates** | Pointer to [**[]ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 

## Methods

### NewClusterPackage

`func NewClusterPackage() *ClusterPackage`

NewClusterPackage instantiates a new ClusterPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterPackageWithDefaults

`func NewClusterPackageWithDefaults() *ClusterPackage`

NewClusterPackageWithDefaults instantiates a new ClusterPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterPackage) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterPackage) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterPackage) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterPackage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ClusterPackage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterPackage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterPackage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterPackage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ClusterPackage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterPackage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterPackage) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClusterPackage) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccount

`func (o *ClusterPackage) GetAccount() int64`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ClusterPackage) GetAccountOk() (*int64, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ClusterPackage) SetAccount(v int64)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ClusterPackage) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCode

`func (o *ClusterPackage) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ClusterPackage) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ClusterPackage) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ClusterPackage) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetRepeatInstall

`func (o *ClusterPackage) GetRepeatInstall() bool`

GetRepeatInstall returns the RepeatInstall field if non-nil, zero value otherwise.

### GetRepeatInstallOk

`func (o *ClusterPackage) GetRepeatInstallOk() (*bool, bool)`

GetRepeatInstallOk returns a tuple with the RepeatInstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatInstall

`func (o *ClusterPackage) SetRepeatInstall(v bool)`

SetRepeatInstall sets RepeatInstall field to given value.

### HasRepeatInstall

`func (o *ClusterPackage) HasRepeatInstall() bool`

HasRepeatInstall returns a boolean if a field has been set.

### GetType

`func (o *ClusterPackage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterPackage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterPackage) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClusterPackage) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPackageType

`func (o *ClusterPackage) GetPackageType() string`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *ClusterPackage) GetPackageTypeOk() (*string, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *ClusterPackage) SetPackageType(v string)`

SetPackageType sets PackageType field to given value.

### HasPackageType

`func (o *ClusterPackage) HasPackageType() bool`

HasPackageType returns a boolean if a field has been set.

### GetPackageVersion

`func (o *ClusterPackage) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *ClusterPackage) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *ClusterPackage) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.

### HasPackageVersion

`func (o *ClusterPackage) HasPackageVersion() bool`

HasPackageVersion returns a boolean if a field has been set.

### GetEnabled

`func (o *ClusterPackage) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ClusterPackage) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ClusterPackage) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ClusterPackage) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIconPath

`func (o *ClusterPackage) GetIconPath() string`

GetIconPath returns the IconPath field if non-nil, zero value otherwise.

### GetIconPathOk

`func (o *ClusterPackage) GetIconPathOk() (*string, bool)`

GetIconPathOk returns a tuple with the IconPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconPath

`func (o *ClusterPackage) SetIconPath(v string)`

SetIconPath sets IconPath field to given value.

### HasIconPath

`func (o *ClusterPackage) HasIconPath() bool`

HasIconPath returns a boolean if a field has been set.

### GetImagePath

`func (o *ClusterPackage) GetImagePath() string`

GetImagePath returns the ImagePath field if non-nil, zero value otherwise.

### GetImagePathOk

`func (o *ClusterPackage) GetImagePathOk() (*string, bool)`

GetImagePathOk returns a tuple with the ImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePath

`func (o *ClusterPackage) SetImagePath(v string)`

SetImagePath sets ImagePath field to given value.

### HasImagePath

`func (o *ClusterPackage) HasImagePath() bool`

HasImagePath returns a boolean if a field has been set.

### GetDarkImagePath

`func (o *ClusterPackage) GetDarkImagePath() string`

GetDarkImagePath returns the DarkImagePath field if non-nil, zero value otherwise.

### GetDarkImagePathOk

`func (o *ClusterPackage) GetDarkImagePathOk() (*string, bool)`

GetDarkImagePathOk returns a tuple with the DarkImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkImagePath

`func (o *ClusterPackage) SetDarkImagePath(v string)`

SetDarkImagePath sets DarkImagePath field to given value.

### HasDarkImagePath

`func (o *ClusterPackage) HasDarkImagePath() bool`

HasDarkImagePath returns a boolean if a field has been set.

### GetSpecTemplates

`func (o *ClusterPackage) GetSpecTemplates() []ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetSpecTemplates returns the SpecTemplates field if non-nil, zero value otherwise.

### GetSpecTemplatesOk

`func (o *ClusterPackage) GetSpecTemplatesOk() (*[]ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetSpecTemplatesOk returns a tuple with the SpecTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecTemplates

`func (o *ClusterPackage) SetSpecTemplates(v []ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetSpecTemplates sets SpecTemplates field to given value.

### HasSpecTemplates

`func (o *ClusterPackage) HasSpecTemplates() bool`

HasSpecTemplates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


