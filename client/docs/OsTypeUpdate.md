# OsTypeUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

## Methods

### NewOsTypeUpdate

`func NewOsTypeUpdate() *OsTypeUpdate`

NewOsTypeUpdate instantiates a new OsTypeUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsTypeUpdateWithDefaults

`func NewOsTypeUpdateWithDefaults() *OsTypeUpdate`

NewOsTypeUpdateWithDefaults instantiates a new OsTypeUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OsTypeUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OsTypeUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OsTypeUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OsTypeUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *OsTypeUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OsTypeUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OsTypeUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OsTypeUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPlatform

`func (o *OsTypeUpdate) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *OsTypeUpdate) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *OsTypeUpdate) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *OsTypeUpdate) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetCategory

`func (o *OsTypeUpdate) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *OsTypeUpdate) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *OsTypeUpdate) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *OsTypeUpdate) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetVendor

`func (o *OsTypeUpdate) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *OsTypeUpdate) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *OsTypeUpdate) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *OsTypeUpdate) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetOsName

`func (o *OsTypeUpdate) GetOsName() string`

GetOsName returns the OsName field if non-nil, zero value otherwise.

### GetOsNameOk

`func (o *OsTypeUpdate) GetOsNameOk() (*string, bool)`

GetOsNameOk returns a tuple with the OsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsName

`func (o *OsTypeUpdate) SetOsName(v string)`

SetOsName sets OsName field to given value.

### HasOsName

`func (o *OsTypeUpdate) HasOsName() bool`

HasOsName returns a boolean if a field has been set.

### GetOsVersion

`func (o *OsTypeUpdate) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *OsTypeUpdate) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *OsTypeUpdate) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *OsTypeUpdate) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetOsCodename

`func (o *OsTypeUpdate) GetOsCodename() string`

GetOsCodename returns the OsCodename field if non-nil, zero value otherwise.

### GetOsCodenameOk

`func (o *OsTypeUpdate) GetOsCodenameOk() (*string, bool)`

GetOsCodenameOk returns a tuple with the OsCodename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsCodename

`func (o *OsTypeUpdate) SetOsCodename(v string)`

SetOsCodename sets OsCodename field to given value.

### HasOsCodename

`func (o *OsTypeUpdate) HasOsCodename() bool`

HasOsCodename returns a boolean if a field has been set.

### GetOsFamily

`func (o *OsTypeUpdate) GetOsFamily() string`

GetOsFamily returns the OsFamily field if non-nil, zero value otherwise.

### GetOsFamilyOk

`func (o *OsTypeUpdate) GetOsFamilyOk() (*string, bool)`

GetOsFamilyOk returns a tuple with the OsFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsFamily

`func (o *OsTypeUpdate) SetOsFamily(v string)`

SetOsFamily sets OsFamily field to given value.

### HasOsFamily

`func (o *OsTypeUpdate) HasOsFamily() bool`

HasOsFamily returns a boolean if a field has been set.

### GetBitCount

`func (o *OsTypeUpdate) GetBitCount() int64`

GetBitCount returns the BitCount field if non-nil, zero value otherwise.

### GetBitCountOk

`func (o *OsTypeUpdate) GetBitCountOk() (*int64, bool)`

GetBitCountOk returns a tuple with the BitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitCount

`func (o *OsTypeUpdate) SetBitCount(v int64)`

SetBitCount sets BitCount field to given value.

### HasBitCount

`func (o *OsTypeUpdate) HasBitCount() bool`

HasBitCount returns a boolean if a field has been set.

### GetCloudInitVersion

`func (o *OsTypeUpdate) GetCloudInitVersion() string`

GetCloudInitVersion returns the CloudInitVersion field if non-nil, zero value otherwise.

### GetCloudInitVersionOk

`func (o *OsTypeUpdate) GetCloudInitVersionOk() (*string, bool)`

GetCloudInitVersionOk returns a tuple with the CloudInitVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitVersion

`func (o *OsTypeUpdate) SetCloudInitVersion(v string)`

SetCloudInitVersion sets CloudInitVersion field to given value.

### HasCloudInitVersion

`func (o *OsTypeUpdate) HasCloudInitVersion() bool`

HasCloudInitVersion returns a boolean if a field has been set.

### GetInstallAgent

`func (o *OsTypeUpdate) GetInstallAgent() bool`

GetInstallAgent returns the InstallAgent field if non-nil, zero value otherwise.

### GetInstallAgentOk

`func (o *OsTypeUpdate) GetInstallAgentOk() (*bool, bool)`

GetInstallAgentOk returns a tuple with the InstallAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAgent

`func (o *OsTypeUpdate) SetInstallAgent(v bool)`

SetInstallAgent sets InstallAgent field to given value.

### HasInstallAgent

`func (o *OsTypeUpdate) HasInstallAgent() bool`

HasInstallAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


