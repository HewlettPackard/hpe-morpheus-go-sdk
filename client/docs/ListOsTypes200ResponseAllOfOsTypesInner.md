# ListOsTypes200ResponseAllOfOsTypesInner

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

### NewListOsTypes200ResponseAllOfOsTypesInner

`func NewListOsTypes200ResponseAllOfOsTypesInner() *ListOsTypes200ResponseAllOfOsTypesInner`

NewListOsTypes200ResponseAllOfOsTypesInner instantiates a new ListOsTypes200ResponseAllOfOsTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOsTypes200ResponseAllOfOsTypesInnerWithDefaults

`func NewListOsTypes200ResponseAllOfOsTypesInnerWithDefaults() *ListOsTypes200ResponseAllOfOsTypesInner`

NewListOsTypes200ResponseAllOfOsTypesInnerWithDefaults instantiates a new ListOsTypes200ResponseAllOfOsTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPlatform

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetCategory

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetVendor

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetOsName

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) GetOsName() string`

GetOsName returns the OsName field if non-nil, zero value otherwise.

### GetOsNameOk

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) GetOsNameOk() (*string, bool)`

GetOsNameOk returns a tuple with the OsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsName

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) SetOsName(v string)`

SetOsName sets OsName field to given value.

### HasOsName

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) HasOsName() bool`

HasOsName returns a boolean if a field has been set.

### GetOsVersion

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetOsCodename

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) GetOsCodename() string`

GetOsCodename returns the OsCodename field if non-nil, zero value otherwise.

### GetOsCodenameOk

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) GetOsCodenameOk() (*string, bool)`

GetOsCodenameOk returns a tuple with the OsCodename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsCodename

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) SetOsCodename(v string)`

SetOsCodename sets OsCodename field to given value.

### HasOsCodename

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) HasOsCodename() bool`

HasOsCodename returns a boolean if a field has been set.

### GetOsFamily

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) GetOsFamily() string`

GetOsFamily returns the OsFamily field if non-nil, zero value otherwise.

### GetOsFamilyOk

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) GetOsFamilyOk() (*string, bool)`

GetOsFamilyOk returns a tuple with the OsFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsFamily

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) SetOsFamily(v string)`

SetOsFamily sets OsFamily field to given value.

### HasOsFamily

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) HasOsFamily() bool`

HasOsFamily returns a boolean if a field has been set.

### GetBitCount

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) GetBitCount() int64`

GetBitCount returns the BitCount field if non-nil, zero value otherwise.

### GetBitCountOk

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) GetBitCountOk() (*int64, bool)`

GetBitCountOk returns a tuple with the BitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitCount

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) SetBitCount(v int64)`

SetBitCount sets BitCount field to given value.

### HasBitCount

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) HasBitCount() bool`

HasBitCount returns a boolean if a field has been set.

### GetCloudInitVersion

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) GetCloudInitVersion() string`

GetCloudInitVersion returns the CloudInitVersion field if non-nil, zero value otherwise.

### GetCloudInitVersionOk

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) GetCloudInitVersionOk() (*string, bool)`

GetCloudInitVersionOk returns a tuple with the CloudInitVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitVersion

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) SetCloudInitVersion(v string)`

SetCloudInitVersion sets CloudInitVersion field to given value.

### HasCloudInitVersion

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) HasCloudInitVersion() bool`

HasCloudInitVersion returns a boolean if a field has been set.

### GetInstallAgent

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) GetInstallAgent() bool`

GetInstallAgent returns the InstallAgent field if non-nil, zero value otherwise.

### GetInstallAgentOk

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) GetInstallAgentOk() (*bool, bool)`

GetInstallAgentOk returns a tuple with the InstallAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAgent

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) SetInstallAgent(v bool)`

SetInstallAgent sets InstallAgent field to given value.

### HasInstallAgent

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) HasInstallAgent() bool`

HasInstallAgent returns a boolean if a field has been set.

### GetImages

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) GetImages() []ListOsTypes200ResponseAllOfOsTypesInnerImagesInner`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) GetImagesOk() (*[]ListOsTypes200ResponseAllOfOsTypesInnerImagesInner, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) SetImages(v []ListOsTypes200ResponseAllOfOsTypesInnerImagesInner)`

SetImages sets Images field to given value.

### HasImages

`func (o *ListOsTypes200ResponseAllOfOsTypesInner) HasImages() bool`

HasImages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


