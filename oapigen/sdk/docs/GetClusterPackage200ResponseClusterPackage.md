# GetClusterPackage200ResponseClusterPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Account** | Pointer to **NullableInt64** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**RepeatInstall** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**PackageType** | Pointer to **string** |  | [optional] 
**PackageVersion** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**IconPath** | Pointer to **NullableString** |  | [optional] 
**ImagePath** | Pointer to **NullableString** |  | [optional] 
**DarkImagePath** | Pointer to **NullableString** |  | [optional] 
**SpecTemplates** | Pointer to [**[]GetClusterPackage200ResponseClusterPackageSpecTemplatesInner**](GetClusterPackage200ResponseClusterPackageSpecTemplatesInner.md) |  | [optional] 

## Methods

### NewGetClusterPackage200ResponseClusterPackage

`func NewGetClusterPackage200ResponseClusterPackage() *GetClusterPackage200ResponseClusterPackage`

NewGetClusterPackage200ResponseClusterPackage instantiates a new GetClusterPackage200ResponseClusterPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetClusterPackage200ResponseClusterPackage) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetClusterPackage200ResponseClusterPackage) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetClusterPackage200ResponseClusterPackage) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetClusterPackage200ResponseClusterPackage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetClusterPackage200ResponseClusterPackage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetClusterPackage200ResponseClusterPackage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetClusterPackage200ResponseClusterPackage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetClusterPackage200ResponseClusterPackage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetClusterPackage200ResponseClusterPackage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetClusterPackage200ResponseClusterPackage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetClusterPackage200ResponseClusterPackage) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetClusterPackage200ResponseClusterPackage) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetClusterPackage200ResponseClusterPackage) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetClusterPackage200ResponseClusterPackage) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAccount

`func (o *GetClusterPackage200ResponseClusterPackage) GetAccount() int64`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetClusterPackage200ResponseClusterPackage) GetAccountOk() (*int64, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetClusterPackage200ResponseClusterPackage) SetAccount(v int64)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetClusterPackage200ResponseClusterPackage) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *GetClusterPackage200ResponseClusterPackage) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *GetClusterPackage200ResponseClusterPackage) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetCode

`func (o *GetClusterPackage200ResponseClusterPackage) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetClusterPackage200ResponseClusterPackage) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetClusterPackage200ResponseClusterPackage) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetClusterPackage200ResponseClusterPackage) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetRepeatInstall

`func (o *GetClusterPackage200ResponseClusterPackage) GetRepeatInstall() bool`

GetRepeatInstall returns the RepeatInstall field if non-nil, zero value otherwise.

### GetRepeatInstallOk

`func (o *GetClusterPackage200ResponseClusterPackage) GetRepeatInstallOk() (*bool, bool)`

GetRepeatInstallOk returns a tuple with the RepeatInstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatInstall

`func (o *GetClusterPackage200ResponseClusterPackage) SetRepeatInstall(v bool)`

SetRepeatInstall sets RepeatInstall field to given value.

### HasRepeatInstall

`func (o *GetClusterPackage200ResponseClusterPackage) HasRepeatInstall() bool`

HasRepeatInstall returns a boolean if a field has been set.

### GetType

`func (o *GetClusterPackage200ResponseClusterPackage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetClusterPackage200ResponseClusterPackage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetClusterPackage200ResponseClusterPackage) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetClusterPackage200ResponseClusterPackage) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPackageType

`func (o *GetClusterPackage200ResponseClusterPackage) GetPackageType() string`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *GetClusterPackage200ResponseClusterPackage) GetPackageTypeOk() (*string, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *GetClusterPackage200ResponseClusterPackage) SetPackageType(v string)`

SetPackageType sets PackageType field to given value.

### HasPackageType

`func (o *GetClusterPackage200ResponseClusterPackage) HasPackageType() bool`

HasPackageType returns a boolean if a field has been set.

### GetPackageVersion

`func (o *GetClusterPackage200ResponseClusterPackage) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *GetClusterPackage200ResponseClusterPackage) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *GetClusterPackage200ResponseClusterPackage) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.

### HasPackageVersion

`func (o *GetClusterPackage200ResponseClusterPackage) HasPackageVersion() bool`

HasPackageVersion returns a boolean if a field has been set.

### GetEnabled

`func (o *GetClusterPackage200ResponseClusterPackage) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetClusterPackage200ResponseClusterPackage) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetClusterPackage200ResponseClusterPackage) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetClusterPackage200ResponseClusterPackage) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIconPath

`func (o *GetClusterPackage200ResponseClusterPackage) GetIconPath() string`

GetIconPath returns the IconPath field if non-nil, zero value otherwise.

### GetIconPathOk

`func (o *GetClusterPackage200ResponseClusterPackage) GetIconPathOk() (*string, bool)`

GetIconPathOk returns a tuple with the IconPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconPath

`func (o *GetClusterPackage200ResponseClusterPackage) SetIconPath(v string)`

SetIconPath sets IconPath field to given value.

### HasIconPath

`func (o *GetClusterPackage200ResponseClusterPackage) HasIconPath() bool`

HasIconPath returns a boolean if a field has been set.

### SetIconPathNil

`func (o *GetClusterPackage200ResponseClusterPackage) SetIconPathNil(b bool)`

 SetIconPathNil sets the value for IconPath to be an explicit nil

### UnsetIconPath
`func (o *GetClusterPackage200ResponseClusterPackage) UnsetIconPath()`

UnsetIconPath ensures that no value is present for IconPath, not even an explicit nil
### GetImagePath

`func (o *GetClusterPackage200ResponseClusterPackage) GetImagePath() string`

GetImagePath returns the ImagePath field if non-nil, zero value otherwise.

### GetImagePathOk

`func (o *GetClusterPackage200ResponseClusterPackage) GetImagePathOk() (*string, bool)`

GetImagePathOk returns a tuple with the ImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePath

`func (o *GetClusterPackage200ResponseClusterPackage) SetImagePath(v string)`

SetImagePath sets ImagePath field to given value.

### HasImagePath

`func (o *GetClusterPackage200ResponseClusterPackage) HasImagePath() bool`

HasImagePath returns a boolean if a field has been set.

### SetImagePathNil

`func (o *GetClusterPackage200ResponseClusterPackage) SetImagePathNil(b bool)`

 SetImagePathNil sets the value for ImagePath to be an explicit nil

### UnsetImagePath
`func (o *GetClusterPackage200ResponseClusterPackage) UnsetImagePath()`

UnsetImagePath ensures that no value is present for ImagePath, not even an explicit nil
### GetDarkImagePath

`func (o *GetClusterPackage200ResponseClusterPackage) GetDarkImagePath() string`

GetDarkImagePath returns the DarkImagePath field if non-nil, zero value otherwise.

### GetDarkImagePathOk

`func (o *GetClusterPackage200ResponseClusterPackage) GetDarkImagePathOk() (*string, bool)`

GetDarkImagePathOk returns a tuple with the DarkImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkImagePath

`func (o *GetClusterPackage200ResponseClusterPackage) SetDarkImagePath(v string)`

SetDarkImagePath sets DarkImagePath field to given value.

### HasDarkImagePath

`func (o *GetClusterPackage200ResponseClusterPackage) HasDarkImagePath() bool`

HasDarkImagePath returns a boolean if a field has been set.

### SetDarkImagePathNil

`func (o *GetClusterPackage200ResponseClusterPackage) SetDarkImagePathNil(b bool)`

 SetDarkImagePathNil sets the value for DarkImagePath to be an explicit nil

### UnsetDarkImagePath
`func (o *GetClusterPackage200ResponseClusterPackage) UnsetDarkImagePath()`

UnsetDarkImagePath ensures that no value is present for DarkImagePath, not even an explicit nil
### GetSpecTemplates

`func (o *GetClusterPackage200ResponseClusterPackage) GetSpecTemplates() []GetClusterPackage200ResponseClusterPackageSpecTemplatesInner`

GetSpecTemplates returns the SpecTemplates field if non-nil, zero value otherwise.

### GetSpecTemplatesOk

`func (o *GetClusterPackage200ResponseClusterPackage) GetSpecTemplatesOk() (*[]GetClusterPackage200ResponseClusterPackageSpecTemplatesInner, bool)`

GetSpecTemplatesOk returns a tuple with the SpecTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecTemplates

`func (o *GetClusterPackage200ResponseClusterPackage) SetSpecTemplates(v []GetClusterPackage200ResponseClusterPackageSpecTemplatesInner)`

SetSpecTemplates sets SpecTemplates field to given value.

### HasSpecTemplates

`func (o *GetClusterPackage200ResponseClusterPackage) HasSpecTemplates() bool`

HasSpecTemplates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


