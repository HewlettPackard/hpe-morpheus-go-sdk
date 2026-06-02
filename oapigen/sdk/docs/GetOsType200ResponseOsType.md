# GetOsType200ResponseOsType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | The name of the osType.  | [optional] 
**Description** | Pointer to **NullableString** | The description of the osType.   | [optional] 
**Platform** | Pointer to **string** | The platform of the osType.   | [optional] 
**Owner** | Pointer to [**GetOsType200ResponseOsTypeOwner**](GetOsType200ResponseOsTypeOwner.md) |  | [optional] 
**Category** | Pointer to **NullableString** | The category of the osType.  | [optional] 
**Vendor** | Pointer to **NullableString** | The vendor of the osType.  | [optional] 
**OsName** | Pointer to **NullableString** | The osName of the osType.  | [optional] 
**OsVersion** | Pointer to **NullableString** | The osVersion of the osType.  | [optional] 
**OsCodename** | Pointer to **NullableString** | The osCodename of the osType.  | [optional] 
**OsFamily** | Pointer to **NullableString** | The family of the osType.  | [optional] 
**BitCount** | Pointer to **int64** | The bitCount/architecture of the osType.  | [optional] 
**CloudInitVersion** | Pointer to **string** | The version of CloudInit being used.  | [optional] 
**InstallAgent** | Pointer to **NullableBool** | Whether the morpheus agent is installed.  | [optional] 
**Images** | Pointer to [**[]GetOsType200ResponseOsTypeImagesInner**](GetOsType200ResponseOsTypeImagesInner.md) |  | [optional] 

## Methods

### NewGetOsType200ResponseOsType

`func NewGetOsType200ResponseOsType() *GetOsType200ResponseOsType`

NewGetOsType200ResponseOsType instantiates a new GetOsType200ResponseOsType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetOsType200ResponseOsType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOsType200ResponseOsType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOsType200ResponseOsType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetOsType200ResponseOsType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *GetOsType200ResponseOsType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetOsType200ResponseOsType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetOsType200ResponseOsType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetOsType200ResponseOsType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *GetOsType200ResponseOsType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetOsType200ResponseOsType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetOsType200ResponseOsType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetOsType200ResponseOsType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetOsType200ResponseOsType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetOsType200ResponseOsType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetOsType200ResponseOsType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetOsType200ResponseOsType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetOsType200ResponseOsType) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetOsType200ResponseOsType) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPlatform

`func (o *GetOsType200ResponseOsType) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetOsType200ResponseOsType) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetOsType200ResponseOsType) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetOsType200ResponseOsType) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetOwner

`func (o *GetOsType200ResponseOsType) GetOwner() GetOsType200ResponseOsTypeOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetOsType200ResponseOsType) GetOwnerOk() (*GetOsType200ResponseOsTypeOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetOsType200ResponseOsType) SetOwner(v GetOsType200ResponseOsTypeOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetOsType200ResponseOsType) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetCategory

`func (o *GetOsType200ResponseOsType) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetOsType200ResponseOsType) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetOsType200ResponseOsType) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetOsType200ResponseOsType) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *GetOsType200ResponseOsType) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *GetOsType200ResponseOsType) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetVendor

`func (o *GetOsType200ResponseOsType) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *GetOsType200ResponseOsType) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *GetOsType200ResponseOsType) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *GetOsType200ResponseOsType) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### SetVendorNil

`func (o *GetOsType200ResponseOsType) SetVendorNil(b bool)`

 SetVendorNil sets the value for Vendor to be an explicit nil

### UnsetVendor
`func (o *GetOsType200ResponseOsType) UnsetVendor()`

UnsetVendor ensures that no value is present for Vendor, not even an explicit nil
### GetOsName

`func (o *GetOsType200ResponseOsType) GetOsName() string`

GetOsName returns the OsName field if non-nil, zero value otherwise.

### GetOsNameOk

`func (o *GetOsType200ResponseOsType) GetOsNameOk() (*string, bool)`

GetOsNameOk returns a tuple with the OsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsName

`func (o *GetOsType200ResponseOsType) SetOsName(v string)`

SetOsName sets OsName field to given value.

### HasOsName

`func (o *GetOsType200ResponseOsType) HasOsName() bool`

HasOsName returns a boolean if a field has been set.

### SetOsNameNil

`func (o *GetOsType200ResponseOsType) SetOsNameNil(b bool)`

 SetOsNameNil sets the value for OsName to be an explicit nil

### UnsetOsName
`func (o *GetOsType200ResponseOsType) UnsetOsName()`

UnsetOsName ensures that no value is present for OsName, not even an explicit nil
### GetOsVersion

`func (o *GetOsType200ResponseOsType) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *GetOsType200ResponseOsType) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *GetOsType200ResponseOsType) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *GetOsType200ResponseOsType) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### SetOsVersionNil

`func (o *GetOsType200ResponseOsType) SetOsVersionNil(b bool)`

 SetOsVersionNil sets the value for OsVersion to be an explicit nil

### UnsetOsVersion
`func (o *GetOsType200ResponseOsType) UnsetOsVersion()`

UnsetOsVersion ensures that no value is present for OsVersion, not even an explicit nil
### GetOsCodename

`func (o *GetOsType200ResponseOsType) GetOsCodename() string`

GetOsCodename returns the OsCodename field if non-nil, zero value otherwise.

### GetOsCodenameOk

`func (o *GetOsType200ResponseOsType) GetOsCodenameOk() (*string, bool)`

GetOsCodenameOk returns a tuple with the OsCodename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsCodename

`func (o *GetOsType200ResponseOsType) SetOsCodename(v string)`

SetOsCodename sets OsCodename field to given value.

### HasOsCodename

`func (o *GetOsType200ResponseOsType) HasOsCodename() bool`

HasOsCodename returns a boolean if a field has been set.

### SetOsCodenameNil

`func (o *GetOsType200ResponseOsType) SetOsCodenameNil(b bool)`

 SetOsCodenameNil sets the value for OsCodename to be an explicit nil

### UnsetOsCodename
`func (o *GetOsType200ResponseOsType) UnsetOsCodename()`

UnsetOsCodename ensures that no value is present for OsCodename, not even an explicit nil
### GetOsFamily

`func (o *GetOsType200ResponseOsType) GetOsFamily() string`

GetOsFamily returns the OsFamily field if non-nil, zero value otherwise.

### GetOsFamilyOk

`func (o *GetOsType200ResponseOsType) GetOsFamilyOk() (*string, bool)`

GetOsFamilyOk returns a tuple with the OsFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsFamily

`func (o *GetOsType200ResponseOsType) SetOsFamily(v string)`

SetOsFamily sets OsFamily field to given value.

### HasOsFamily

`func (o *GetOsType200ResponseOsType) HasOsFamily() bool`

HasOsFamily returns a boolean if a field has been set.

### SetOsFamilyNil

`func (o *GetOsType200ResponseOsType) SetOsFamilyNil(b bool)`

 SetOsFamilyNil sets the value for OsFamily to be an explicit nil

### UnsetOsFamily
`func (o *GetOsType200ResponseOsType) UnsetOsFamily()`

UnsetOsFamily ensures that no value is present for OsFamily, not even an explicit nil
### GetBitCount

`func (o *GetOsType200ResponseOsType) GetBitCount() int64`

GetBitCount returns the BitCount field if non-nil, zero value otherwise.

### GetBitCountOk

`func (o *GetOsType200ResponseOsType) GetBitCountOk() (*int64, bool)`

GetBitCountOk returns a tuple with the BitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitCount

`func (o *GetOsType200ResponseOsType) SetBitCount(v int64)`

SetBitCount sets BitCount field to given value.

### HasBitCount

`func (o *GetOsType200ResponseOsType) HasBitCount() bool`

HasBitCount returns a boolean if a field has been set.

### GetCloudInitVersion

`func (o *GetOsType200ResponseOsType) GetCloudInitVersion() string`

GetCloudInitVersion returns the CloudInitVersion field if non-nil, zero value otherwise.

### GetCloudInitVersionOk

`func (o *GetOsType200ResponseOsType) GetCloudInitVersionOk() (*string, bool)`

GetCloudInitVersionOk returns a tuple with the CloudInitVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitVersion

`func (o *GetOsType200ResponseOsType) SetCloudInitVersion(v string)`

SetCloudInitVersion sets CloudInitVersion field to given value.

### HasCloudInitVersion

`func (o *GetOsType200ResponseOsType) HasCloudInitVersion() bool`

HasCloudInitVersion returns a boolean if a field has been set.

### GetInstallAgent

`func (o *GetOsType200ResponseOsType) GetInstallAgent() bool`

GetInstallAgent returns the InstallAgent field if non-nil, zero value otherwise.

### GetInstallAgentOk

`func (o *GetOsType200ResponseOsType) GetInstallAgentOk() (*bool, bool)`

GetInstallAgentOk returns a tuple with the InstallAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAgent

`func (o *GetOsType200ResponseOsType) SetInstallAgent(v bool)`

SetInstallAgent sets InstallAgent field to given value.

### HasInstallAgent

`func (o *GetOsType200ResponseOsType) HasInstallAgent() bool`

HasInstallAgent returns a boolean if a field has been set.

### SetInstallAgentNil

`func (o *GetOsType200ResponseOsType) SetInstallAgentNil(b bool)`

 SetInstallAgentNil sets the value for InstallAgent to be an explicit nil

### UnsetInstallAgent
`func (o *GetOsType200ResponseOsType) UnsetInstallAgent()`

UnsetInstallAgent ensures that no value is present for InstallAgent, not even an explicit nil
### GetImages

`func (o *GetOsType200ResponseOsType) GetImages() []GetOsType200ResponseOsTypeImagesInner`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *GetOsType200ResponseOsType) GetImagesOk() (*[]GetOsType200ResponseOsTypeImagesInner, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *GetOsType200ResponseOsType) SetImages(v []GetOsType200ResponseOsTypeImagesInner)`

SetImages sets Images field to given value.

### HasImages

`func (o *GetOsType200ResponseOsType) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *GetOsType200ResponseOsType) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *GetOsType200ResponseOsType) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


