# UpdateOsTypeRequestOsType

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

### NewUpdateOsTypeRequestOsType

`func NewUpdateOsTypeRequestOsType() *UpdateOsTypeRequestOsType`

NewUpdateOsTypeRequestOsType instantiates a new UpdateOsTypeRequestOsType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOsTypeRequestOsTypeWithDefaults

`func NewUpdateOsTypeRequestOsTypeWithDefaults() *UpdateOsTypeRequestOsType`

NewUpdateOsTypeRequestOsTypeWithDefaults instantiates a new UpdateOsTypeRequestOsType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateOsTypeRequestOsType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateOsTypeRequestOsType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateOsTypeRequestOsType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateOsTypeRequestOsType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateOsTypeRequestOsType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateOsTypeRequestOsType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateOsTypeRequestOsType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateOsTypeRequestOsType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPlatform

`func (o *UpdateOsTypeRequestOsType) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *UpdateOsTypeRequestOsType) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *UpdateOsTypeRequestOsType) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *UpdateOsTypeRequestOsType) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetCategory

`func (o *UpdateOsTypeRequestOsType) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *UpdateOsTypeRequestOsType) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *UpdateOsTypeRequestOsType) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *UpdateOsTypeRequestOsType) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetVendor

`func (o *UpdateOsTypeRequestOsType) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *UpdateOsTypeRequestOsType) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *UpdateOsTypeRequestOsType) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *UpdateOsTypeRequestOsType) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetOsName

`func (o *UpdateOsTypeRequestOsType) GetOsName() string`

GetOsName returns the OsName field if non-nil, zero value otherwise.

### GetOsNameOk

`func (o *UpdateOsTypeRequestOsType) GetOsNameOk() (*string, bool)`

GetOsNameOk returns a tuple with the OsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsName

`func (o *UpdateOsTypeRequestOsType) SetOsName(v string)`

SetOsName sets OsName field to given value.

### HasOsName

`func (o *UpdateOsTypeRequestOsType) HasOsName() bool`

HasOsName returns a boolean if a field has been set.

### GetOsVersion

`func (o *UpdateOsTypeRequestOsType) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *UpdateOsTypeRequestOsType) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *UpdateOsTypeRequestOsType) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *UpdateOsTypeRequestOsType) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetOsCodename

`func (o *UpdateOsTypeRequestOsType) GetOsCodename() string`

GetOsCodename returns the OsCodename field if non-nil, zero value otherwise.

### GetOsCodenameOk

`func (o *UpdateOsTypeRequestOsType) GetOsCodenameOk() (*string, bool)`

GetOsCodenameOk returns a tuple with the OsCodename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsCodename

`func (o *UpdateOsTypeRequestOsType) SetOsCodename(v string)`

SetOsCodename sets OsCodename field to given value.

### HasOsCodename

`func (o *UpdateOsTypeRequestOsType) HasOsCodename() bool`

HasOsCodename returns a boolean if a field has been set.

### GetOsFamily

`func (o *UpdateOsTypeRequestOsType) GetOsFamily() string`

GetOsFamily returns the OsFamily field if non-nil, zero value otherwise.

### GetOsFamilyOk

`func (o *UpdateOsTypeRequestOsType) GetOsFamilyOk() (*string, bool)`

GetOsFamilyOk returns a tuple with the OsFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsFamily

`func (o *UpdateOsTypeRequestOsType) SetOsFamily(v string)`

SetOsFamily sets OsFamily field to given value.

### HasOsFamily

`func (o *UpdateOsTypeRequestOsType) HasOsFamily() bool`

HasOsFamily returns a boolean if a field has been set.

### GetBitCount

`func (o *UpdateOsTypeRequestOsType) GetBitCount() int64`

GetBitCount returns the BitCount field if non-nil, zero value otherwise.

### GetBitCountOk

`func (o *UpdateOsTypeRequestOsType) GetBitCountOk() (*int64, bool)`

GetBitCountOk returns a tuple with the BitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitCount

`func (o *UpdateOsTypeRequestOsType) SetBitCount(v int64)`

SetBitCount sets BitCount field to given value.

### HasBitCount

`func (o *UpdateOsTypeRequestOsType) HasBitCount() bool`

HasBitCount returns a boolean if a field has been set.

### GetCloudInitVersion

`func (o *UpdateOsTypeRequestOsType) GetCloudInitVersion() string`

GetCloudInitVersion returns the CloudInitVersion field if non-nil, zero value otherwise.

### GetCloudInitVersionOk

`func (o *UpdateOsTypeRequestOsType) GetCloudInitVersionOk() (*string, bool)`

GetCloudInitVersionOk returns a tuple with the CloudInitVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitVersion

`func (o *UpdateOsTypeRequestOsType) SetCloudInitVersion(v string)`

SetCloudInitVersion sets CloudInitVersion field to given value.

### HasCloudInitVersion

`func (o *UpdateOsTypeRequestOsType) HasCloudInitVersion() bool`

HasCloudInitVersion returns a boolean if a field has been set.

### GetInstallAgent

`func (o *UpdateOsTypeRequestOsType) GetInstallAgent() bool`

GetInstallAgent returns the InstallAgent field if non-nil, zero value otherwise.

### GetInstallAgentOk

`func (o *UpdateOsTypeRequestOsType) GetInstallAgentOk() (*bool, bool)`

GetInstallAgentOk returns a tuple with the InstallAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAgent

`func (o *UpdateOsTypeRequestOsType) SetInstallAgent(v bool)`

SetInstallAgent sets InstallAgent field to given value.

### HasInstallAgent

`func (o *UpdateOsTypeRequestOsType) HasInstallAgent() bool`

HasInstallAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


