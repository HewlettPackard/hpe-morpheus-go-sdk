# OsType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** | The name of the osType.  | [optional] 
**Description** | Pointer to **string** | The description of the osType.   | [optional] 
**Platform** | Pointer to **string** | The platform of the osType.   | [optional] 
**Category** | Pointer to **string** | The category of the osType.  | [optional] 
**Vendor** | Pointer to **string** | The vendor of the osType.  | [optional] 
**OsName** | Pointer to **string** | The osName of the osType.  | [optional] 
**OsVersion** | Pointer to **string** | The osVersion of the osType.  | [optional] 
**OsCodename** | Pointer to **string** | The osCodename of the osType.  | [optional] 
**OsFamily** | Pointer to **string** | The family of the osType.  | [optional] 
**BitCount** | Pointer to **int64** | The bitCount/architecture of the osType.  | [optional] 
**CloudInitVersion** | Pointer to **string** | The version of CloudInit being used.  | [optional] 
**InstallAgent** | Pointer to **bool** | Whether the morpheus agent is installed.  | [optional] 
**Images** | Pointer to [**[]ListOsTypes200ResponseAllOfOsTypesInnerImagesInner**](ListOsTypes200ResponseAllOfOsTypesInnerImagesInner.md) |  | [optional] 

## Methods

### NewOsType

`func NewOsType() *OsType`

NewOsType instantiates a new OsType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsTypeWithDefaults

`func NewOsTypeWithDefaults() *OsType`

NewOsTypeWithDefaults instantiates a new OsType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OsType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OsType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OsType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OsType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OsType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OsType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OsType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OsType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *OsType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OsType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OsType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OsType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPlatform

`func (o *OsType) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *OsType) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *OsType) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *OsType) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetCategory

`func (o *OsType) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *OsType) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *OsType) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *OsType) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetVendor

`func (o *OsType) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *OsType) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *OsType) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *OsType) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetOsName

`func (o *OsType) GetOsName() string`

GetOsName returns the OsName field if non-nil, zero value otherwise.

### GetOsNameOk

`func (o *OsType) GetOsNameOk() (*string, bool)`

GetOsNameOk returns a tuple with the OsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsName

`func (o *OsType) SetOsName(v string)`

SetOsName sets OsName field to given value.

### HasOsName

`func (o *OsType) HasOsName() bool`

HasOsName returns a boolean if a field has been set.

### GetOsVersion

`func (o *OsType) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *OsType) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *OsType) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *OsType) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetOsCodename

`func (o *OsType) GetOsCodename() string`

GetOsCodename returns the OsCodename field if non-nil, zero value otherwise.

### GetOsCodenameOk

`func (o *OsType) GetOsCodenameOk() (*string, bool)`

GetOsCodenameOk returns a tuple with the OsCodename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsCodename

`func (o *OsType) SetOsCodename(v string)`

SetOsCodename sets OsCodename field to given value.

### HasOsCodename

`func (o *OsType) HasOsCodename() bool`

HasOsCodename returns a boolean if a field has been set.

### GetOsFamily

`func (o *OsType) GetOsFamily() string`

GetOsFamily returns the OsFamily field if non-nil, zero value otherwise.

### GetOsFamilyOk

`func (o *OsType) GetOsFamilyOk() (*string, bool)`

GetOsFamilyOk returns a tuple with the OsFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsFamily

`func (o *OsType) SetOsFamily(v string)`

SetOsFamily sets OsFamily field to given value.

### HasOsFamily

`func (o *OsType) HasOsFamily() bool`

HasOsFamily returns a boolean if a field has been set.

### GetBitCount

`func (o *OsType) GetBitCount() int64`

GetBitCount returns the BitCount field if non-nil, zero value otherwise.

### GetBitCountOk

`func (o *OsType) GetBitCountOk() (*int64, bool)`

GetBitCountOk returns a tuple with the BitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitCount

`func (o *OsType) SetBitCount(v int64)`

SetBitCount sets BitCount field to given value.

### HasBitCount

`func (o *OsType) HasBitCount() bool`

HasBitCount returns a boolean if a field has been set.

### GetCloudInitVersion

`func (o *OsType) GetCloudInitVersion() string`

GetCloudInitVersion returns the CloudInitVersion field if non-nil, zero value otherwise.

### GetCloudInitVersionOk

`func (o *OsType) GetCloudInitVersionOk() (*string, bool)`

GetCloudInitVersionOk returns a tuple with the CloudInitVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitVersion

`func (o *OsType) SetCloudInitVersion(v string)`

SetCloudInitVersion sets CloudInitVersion field to given value.

### HasCloudInitVersion

`func (o *OsType) HasCloudInitVersion() bool`

HasCloudInitVersion returns a boolean if a field has been set.

### GetInstallAgent

`func (o *OsType) GetInstallAgent() bool`

GetInstallAgent returns the InstallAgent field if non-nil, zero value otherwise.

### GetInstallAgentOk

`func (o *OsType) GetInstallAgentOk() (*bool, bool)`

GetInstallAgentOk returns a tuple with the InstallAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAgent

`func (o *OsType) SetInstallAgent(v bool)`

SetInstallAgent sets InstallAgent field to given value.

### HasInstallAgent

`func (o *OsType) HasInstallAgent() bool`

HasInstallAgent returns a boolean if a field has been set.

### GetImages

`func (o *OsType) GetImages() []ListOsTypes200ResponseAllOfOsTypesInnerImagesInner`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *OsType) GetImagesOk() (*[]ListOsTypes200ResponseAllOfOsTypesInnerImagesInner, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *OsType) SetImages(v []ListOsTypes200ResponseAllOfOsTypesInnerImagesInner)`

SetImages sets Images field to given value.

### HasImages

`func (o *OsType) HasImages() bool`

HasImages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


