# OsTypeCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the osType.  | 
**Description** | Pointer to **string** | The description of the osType.   | [optional] 
**Platform** | **string** | The platform of the osType.   | 
**Code** | **string** | The code morph uses as an identifier  | 
**Category** | Pointer to **string** | The category of the osType.  | [optional] 
**Vendor** | Pointer to **string** | The vendor of the osType.  | [optional] 
**OsName** | Pointer to **string** | The osName of the osType.  | [optional] 
**OsVersion** | Pointer to **string** | The osVersion of the osType.  | [optional] 
**OsCodename** | Pointer to **string** | The osCodename of the osType.  | [optional] 
**OsFamily** | Pointer to **string** | The family of the osType.  | [optional] 
**BitCount** | **int64** | The bitCount/architecture of the osType.  | 
**CloudInitVersion** | Pointer to **string** | The version of CloudInit being used.  | [optional] 
**InstallAgent** | Pointer to **bool** | Whether the morpheus agent is installed.  | [optional] 

## Methods

### NewOsTypeCreate

`func NewOsTypeCreate(name string, platform string, code string, bitCount int64, ) *OsTypeCreate`

NewOsTypeCreate instantiates a new OsTypeCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsTypeCreateWithDefaults

`func NewOsTypeCreateWithDefaults() *OsTypeCreate`

NewOsTypeCreateWithDefaults instantiates a new OsTypeCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OsTypeCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OsTypeCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OsTypeCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *OsTypeCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OsTypeCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OsTypeCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OsTypeCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPlatform

`func (o *OsTypeCreate) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *OsTypeCreate) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *OsTypeCreate) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetCode

`func (o *OsTypeCreate) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OsTypeCreate) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OsTypeCreate) SetCode(v string)`

SetCode sets Code field to given value.


### GetCategory

`func (o *OsTypeCreate) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *OsTypeCreate) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *OsTypeCreate) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *OsTypeCreate) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetVendor

`func (o *OsTypeCreate) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *OsTypeCreate) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *OsTypeCreate) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *OsTypeCreate) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetOsName

`func (o *OsTypeCreate) GetOsName() string`

GetOsName returns the OsName field if non-nil, zero value otherwise.

### GetOsNameOk

`func (o *OsTypeCreate) GetOsNameOk() (*string, bool)`

GetOsNameOk returns a tuple with the OsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsName

`func (o *OsTypeCreate) SetOsName(v string)`

SetOsName sets OsName field to given value.

### HasOsName

`func (o *OsTypeCreate) HasOsName() bool`

HasOsName returns a boolean if a field has been set.

### GetOsVersion

`func (o *OsTypeCreate) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *OsTypeCreate) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *OsTypeCreate) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *OsTypeCreate) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetOsCodename

`func (o *OsTypeCreate) GetOsCodename() string`

GetOsCodename returns the OsCodename field if non-nil, zero value otherwise.

### GetOsCodenameOk

`func (o *OsTypeCreate) GetOsCodenameOk() (*string, bool)`

GetOsCodenameOk returns a tuple with the OsCodename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsCodename

`func (o *OsTypeCreate) SetOsCodename(v string)`

SetOsCodename sets OsCodename field to given value.

### HasOsCodename

`func (o *OsTypeCreate) HasOsCodename() bool`

HasOsCodename returns a boolean if a field has been set.

### GetOsFamily

`func (o *OsTypeCreate) GetOsFamily() string`

GetOsFamily returns the OsFamily field if non-nil, zero value otherwise.

### GetOsFamilyOk

`func (o *OsTypeCreate) GetOsFamilyOk() (*string, bool)`

GetOsFamilyOk returns a tuple with the OsFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsFamily

`func (o *OsTypeCreate) SetOsFamily(v string)`

SetOsFamily sets OsFamily field to given value.

### HasOsFamily

`func (o *OsTypeCreate) HasOsFamily() bool`

HasOsFamily returns a boolean if a field has been set.

### GetBitCount

`func (o *OsTypeCreate) GetBitCount() int64`

GetBitCount returns the BitCount field if non-nil, zero value otherwise.

### GetBitCountOk

`func (o *OsTypeCreate) GetBitCountOk() (*int64, bool)`

GetBitCountOk returns a tuple with the BitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitCount

`func (o *OsTypeCreate) SetBitCount(v int64)`

SetBitCount sets BitCount field to given value.


### GetCloudInitVersion

`func (o *OsTypeCreate) GetCloudInitVersion() string`

GetCloudInitVersion returns the CloudInitVersion field if non-nil, zero value otherwise.

### GetCloudInitVersionOk

`func (o *OsTypeCreate) GetCloudInitVersionOk() (*string, bool)`

GetCloudInitVersionOk returns a tuple with the CloudInitVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitVersion

`func (o *OsTypeCreate) SetCloudInitVersion(v string)`

SetCloudInitVersion sets CloudInitVersion field to given value.

### HasCloudInitVersion

`func (o *OsTypeCreate) HasCloudInitVersion() bool`

HasCloudInitVersion returns a boolean if a field has been set.

### GetInstallAgent

`func (o *OsTypeCreate) GetInstallAgent() bool`

GetInstallAgent returns the InstallAgent field if non-nil, zero value otherwise.

### GetInstallAgentOk

`func (o *OsTypeCreate) GetInstallAgentOk() (*bool, bool)`

GetInstallAgentOk returns a tuple with the InstallAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAgent

`func (o *OsTypeCreate) SetInstallAgent(v bool)`

SetInstallAgent sets InstallAgent field to given value.

### HasInstallAgent

`func (o *OsTypeCreate) HasInstallAgent() bool`

HasInstallAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


