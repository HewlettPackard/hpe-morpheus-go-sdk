# AddOsTypesRequestOsType

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

### NewAddOsTypesRequestOsType

`func NewAddOsTypesRequestOsType(name string, platform string, code string, bitCount int64, ) *AddOsTypesRequestOsType`

NewAddOsTypesRequestOsType instantiates a new AddOsTypesRequestOsType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddOsTypesRequestOsTypeWithDefaults

`func NewAddOsTypesRequestOsTypeWithDefaults() *AddOsTypesRequestOsType`

NewAddOsTypesRequestOsTypeWithDefaults instantiates a new AddOsTypesRequestOsType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddOsTypesRequestOsType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddOsTypesRequestOsType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddOsTypesRequestOsType) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AddOsTypesRequestOsType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddOsTypesRequestOsType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddOsTypesRequestOsType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddOsTypesRequestOsType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPlatform

`func (o *AddOsTypesRequestOsType) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *AddOsTypesRequestOsType) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *AddOsTypesRequestOsType) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetCode

`func (o *AddOsTypesRequestOsType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddOsTypesRequestOsType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddOsTypesRequestOsType) SetCode(v string)`

SetCode sets Code field to given value.


### GetCategory

`func (o *AddOsTypesRequestOsType) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AddOsTypesRequestOsType) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AddOsTypesRequestOsType) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AddOsTypesRequestOsType) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetVendor

`func (o *AddOsTypesRequestOsType) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *AddOsTypesRequestOsType) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *AddOsTypesRequestOsType) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *AddOsTypesRequestOsType) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetOsName

`func (o *AddOsTypesRequestOsType) GetOsName() string`

GetOsName returns the OsName field if non-nil, zero value otherwise.

### GetOsNameOk

`func (o *AddOsTypesRequestOsType) GetOsNameOk() (*string, bool)`

GetOsNameOk returns a tuple with the OsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsName

`func (o *AddOsTypesRequestOsType) SetOsName(v string)`

SetOsName sets OsName field to given value.

### HasOsName

`func (o *AddOsTypesRequestOsType) HasOsName() bool`

HasOsName returns a boolean if a field has been set.

### GetOsVersion

`func (o *AddOsTypesRequestOsType) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *AddOsTypesRequestOsType) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *AddOsTypesRequestOsType) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *AddOsTypesRequestOsType) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetOsCodename

`func (o *AddOsTypesRequestOsType) GetOsCodename() string`

GetOsCodename returns the OsCodename field if non-nil, zero value otherwise.

### GetOsCodenameOk

`func (o *AddOsTypesRequestOsType) GetOsCodenameOk() (*string, bool)`

GetOsCodenameOk returns a tuple with the OsCodename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsCodename

`func (o *AddOsTypesRequestOsType) SetOsCodename(v string)`

SetOsCodename sets OsCodename field to given value.

### HasOsCodename

`func (o *AddOsTypesRequestOsType) HasOsCodename() bool`

HasOsCodename returns a boolean if a field has been set.

### GetOsFamily

`func (o *AddOsTypesRequestOsType) GetOsFamily() string`

GetOsFamily returns the OsFamily field if non-nil, zero value otherwise.

### GetOsFamilyOk

`func (o *AddOsTypesRequestOsType) GetOsFamilyOk() (*string, bool)`

GetOsFamilyOk returns a tuple with the OsFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsFamily

`func (o *AddOsTypesRequestOsType) SetOsFamily(v string)`

SetOsFamily sets OsFamily field to given value.

### HasOsFamily

`func (o *AddOsTypesRequestOsType) HasOsFamily() bool`

HasOsFamily returns a boolean if a field has been set.

### GetBitCount

`func (o *AddOsTypesRequestOsType) GetBitCount() int64`

GetBitCount returns the BitCount field if non-nil, zero value otherwise.

### GetBitCountOk

`func (o *AddOsTypesRequestOsType) GetBitCountOk() (*int64, bool)`

GetBitCountOk returns a tuple with the BitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitCount

`func (o *AddOsTypesRequestOsType) SetBitCount(v int64)`

SetBitCount sets BitCount field to given value.


### GetCloudInitVersion

`func (o *AddOsTypesRequestOsType) GetCloudInitVersion() string`

GetCloudInitVersion returns the CloudInitVersion field if non-nil, zero value otherwise.

### GetCloudInitVersionOk

`func (o *AddOsTypesRequestOsType) GetCloudInitVersionOk() (*string, bool)`

GetCloudInitVersionOk returns a tuple with the CloudInitVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitVersion

`func (o *AddOsTypesRequestOsType) SetCloudInitVersion(v string)`

SetCloudInitVersion sets CloudInitVersion field to given value.

### HasCloudInitVersion

`func (o *AddOsTypesRequestOsType) HasCloudInitVersion() bool`

HasCloudInitVersion returns a boolean if a field has been set.

### GetInstallAgent

`func (o *AddOsTypesRequestOsType) GetInstallAgent() bool`

GetInstallAgent returns the InstallAgent field if non-nil, zero value otherwise.

### GetInstallAgentOk

`func (o *AddOsTypesRequestOsType) GetInstallAgentOk() (*bool, bool)`

GetInstallAgentOk returns a tuple with the InstallAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAgent

`func (o *AddOsTypesRequestOsType) SetInstallAgent(v bool)`

SetInstallAgent sets InstallAgent field to given value.

### HasInstallAgent

`func (o *AddOsTypesRequestOsType) HasInstallAgent() bool`

HasInstallAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


