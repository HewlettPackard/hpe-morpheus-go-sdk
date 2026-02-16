# InstanceContainerServer2ServerOs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**OsFamily** | Pointer to **string** |  | [optional] 
**OsVersion** | Pointer to **string** |  | [optional] 
**BitCount** | Pointer to **int64** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 

## Methods

### NewInstanceContainerServer2ServerOs

`func NewInstanceContainerServer2ServerOs() *InstanceContainerServer2ServerOs`

NewInstanceContainerServer2ServerOs instantiates a new InstanceContainerServer2ServerOs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceContainerServer2ServerOsWithDefaults

`func NewInstanceContainerServer2ServerOsWithDefaults() *InstanceContainerServer2ServerOs`

NewInstanceContainerServer2ServerOsWithDefaults instantiates a new InstanceContainerServer2ServerOs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceContainerServer2ServerOs) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceContainerServer2ServerOs) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceContainerServer2ServerOs) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceContainerServer2ServerOs) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *InstanceContainerServer2ServerOs) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InstanceContainerServer2ServerOs) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InstanceContainerServer2ServerOs) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *InstanceContainerServer2ServerOs) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *InstanceContainerServer2ServerOs) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceContainerServer2ServerOs) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceContainerServer2ServerOs) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceContainerServer2ServerOs) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InstanceContainerServer2ServerOs) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InstanceContainerServer2ServerOs) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InstanceContainerServer2ServerOs) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InstanceContainerServer2ServerOs) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InstanceContainerServer2ServerOs) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InstanceContainerServer2ServerOs) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetVendor

`func (o *InstanceContainerServer2ServerOs) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *InstanceContainerServer2ServerOs) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *InstanceContainerServer2ServerOs) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *InstanceContainerServer2ServerOs) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetCategory

`func (o *InstanceContainerServer2ServerOs) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InstanceContainerServer2ServerOs) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InstanceContainerServer2ServerOs) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *InstanceContainerServer2ServerOs) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetOsFamily

`func (o *InstanceContainerServer2ServerOs) GetOsFamily() string`

GetOsFamily returns the OsFamily field if non-nil, zero value otherwise.

### GetOsFamilyOk

`func (o *InstanceContainerServer2ServerOs) GetOsFamilyOk() (*string, bool)`

GetOsFamilyOk returns a tuple with the OsFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsFamily

`func (o *InstanceContainerServer2ServerOs) SetOsFamily(v string)`

SetOsFamily sets OsFamily field to given value.

### HasOsFamily

`func (o *InstanceContainerServer2ServerOs) HasOsFamily() bool`

HasOsFamily returns a boolean if a field has been set.

### GetOsVersion

`func (o *InstanceContainerServer2ServerOs) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *InstanceContainerServer2ServerOs) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *InstanceContainerServer2ServerOs) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *InstanceContainerServer2ServerOs) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetBitCount

`func (o *InstanceContainerServer2ServerOs) GetBitCount() int64`

GetBitCount returns the BitCount field if non-nil, zero value otherwise.

### GetBitCountOk

`func (o *InstanceContainerServer2ServerOs) GetBitCountOk() (*int64, bool)`

GetBitCountOk returns a tuple with the BitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitCount

`func (o *InstanceContainerServer2ServerOs) SetBitCount(v int64)`

SetBitCount sets BitCount field to given value.

### HasBitCount

`func (o *InstanceContainerServer2ServerOs) HasBitCount() bool`

HasBitCount returns a boolean if a field has been set.

### GetPlatform

`func (o *InstanceContainerServer2ServerOs) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *InstanceContainerServer2ServerOs) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *InstanceContainerServer2ServerOs) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *InstanceContainerServer2ServerOs) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


