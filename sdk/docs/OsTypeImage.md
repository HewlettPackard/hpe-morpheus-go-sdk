# OsTypeImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**VirtualImageId** | Pointer to **int64** | The id of the virtual image.  | [optional] 
**VirtualImageName** | Pointer to **string** | The name of the virtual image.   | [optional] 
**Account** | Pointer to **NullableInt64** | The account attached to the osTypeImage.   | [optional] 
**ProvisionType** | Pointer to **NullableInt64** | The Provision Type of the osTypeImage.  | [optional] 
**ComputeZoneType** | Pointer to **NullableInt64** | The zone type of the osTypeImage.  | [optional] 
**Zone** | Pointer to **NullableInt64** | The cloud that is attached to osTypeImage. | [optional] 

## Methods

### NewOsTypeImage

`func NewOsTypeImage() *OsTypeImage`

NewOsTypeImage instantiates a new OsTypeImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsTypeImageWithDefaults

`func NewOsTypeImageWithDefaults() *OsTypeImage`

NewOsTypeImageWithDefaults instantiates a new OsTypeImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OsTypeImage) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OsTypeImage) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OsTypeImage) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OsTypeImage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVirtualImageId

`func (o *OsTypeImage) GetVirtualImageId() int64`

GetVirtualImageId returns the VirtualImageId field if non-nil, zero value otherwise.

### GetVirtualImageIdOk

`func (o *OsTypeImage) GetVirtualImageIdOk() (*int64, bool)`

GetVirtualImageIdOk returns a tuple with the VirtualImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImageId

`func (o *OsTypeImage) SetVirtualImageId(v int64)`

SetVirtualImageId sets VirtualImageId field to given value.

### HasVirtualImageId

`func (o *OsTypeImage) HasVirtualImageId() bool`

HasVirtualImageId returns a boolean if a field has been set.

### GetVirtualImageName

`func (o *OsTypeImage) GetVirtualImageName() string`

GetVirtualImageName returns the VirtualImageName field if non-nil, zero value otherwise.

### GetVirtualImageNameOk

`func (o *OsTypeImage) GetVirtualImageNameOk() (*string, bool)`

GetVirtualImageNameOk returns a tuple with the VirtualImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImageName

`func (o *OsTypeImage) SetVirtualImageName(v string)`

SetVirtualImageName sets VirtualImageName field to given value.

### HasVirtualImageName

`func (o *OsTypeImage) HasVirtualImageName() bool`

HasVirtualImageName returns a boolean if a field has been set.

### GetAccount

`func (o *OsTypeImage) GetAccount() int64`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *OsTypeImage) GetAccountOk() (*int64, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *OsTypeImage) SetAccount(v int64)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *OsTypeImage) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *OsTypeImage) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *OsTypeImage) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetProvisionType

`func (o *OsTypeImage) GetProvisionType() int64`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *OsTypeImage) GetProvisionTypeOk() (*int64, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *OsTypeImage) SetProvisionType(v int64)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *OsTypeImage) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### SetProvisionTypeNil

`func (o *OsTypeImage) SetProvisionTypeNil(b bool)`

 SetProvisionTypeNil sets the value for ProvisionType to be an explicit nil

### UnsetProvisionType
`func (o *OsTypeImage) UnsetProvisionType()`

UnsetProvisionType ensures that no value is present for ProvisionType, not even an explicit nil
### GetComputeZoneType

`func (o *OsTypeImage) GetComputeZoneType() int64`

GetComputeZoneType returns the ComputeZoneType field if non-nil, zero value otherwise.

### GetComputeZoneTypeOk

`func (o *OsTypeImage) GetComputeZoneTypeOk() (*int64, bool)`

GetComputeZoneTypeOk returns a tuple with the ComputeZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeZoneType

`func (o *OsTypeImage) SetComputeZoneType(v int64)`

SetComputeZoneType sets ComputeZoneType field to given value.

### HasComputeZoneType

`func (o *OsTypeImage) HasComputeZoneType() bool`

HasComputeZoneType returns a boolean if a field has been set.

### SetComputeZoneTypeNil

`func (o *OsTypeImage) SetComputeZoneTypeNil(b bool)`

 SetComputeZoneTypeNil sets the value for ComputeZoneType to be an explicit nil

### UnsetComputeZoneType
`func (o *OsTypeImage) UnsetComputeZoneType()`

UnsetComputeZoneType ensures that no value is present for ComputeZoneType, not even an explicit nil
### GetZone

`func (o *OsTypeImage) GetZone() int64`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *OsTypeImage) GetZoneOk() (*int64, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *OsTypeImage) SetZone(v int64)`

SetZone sets Zone field to given value.

### HasZone

`func (o *OsTypeImage) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *OsTypeImage) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *OsTypeImage) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


