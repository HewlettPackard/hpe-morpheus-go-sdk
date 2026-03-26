# CreateOsTypeImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OsType** | **int64** | id of osType | 
**VirtualImage** | **int64** | id of virtualImage | 
**ProvisionType** | Pointer to **NullableInt64** | id of provisionType | [optional] 
**Zone** | Pointer to **NullableInt64** | id of cloud/zone | [optional] 

## Methods

### NewCreateOsTypeImage

`func NewCreateOsTypeImage(osType int64, virtualImage int64, ) *CreateOsTypeImage`

NewCreateOsTypeImage instantiates a new CreateOsTypeImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOsTypeImageWithDefaults

`func NewCreateOsTypeImageWithDefaults() *CreateOsTypeImage`

NewCreateOsTypeImageWithDefaults instantiates a new CreateOsTypeImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOsType

`func (o *CreateOsTypeImage) GetOsType() int64`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *CreateOsTypeImage) GetOsTypeOk() (*int64, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *CreateOsTypeImage) SetOsType(v int64)`

SetOsType sets OsType field to given value.


### GetVirtualImage

`func (o *CreateOsTypeImage) GetVirtualImage() int64`

GetVirtualImage returns the VirtualImage field if non-nil, zero value otherwise.

### GetVirtualImageOk

`func (o *CreateOsTypeImage) GetVirtualImageOk() (*int64, bool)`

GetVirtualImageOk returns a tuple with the VirtualImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImage

`func (o *CreateOsTypeImage) SetVirtualImage(v int64)`

SetVirtualImage sets VirtualImage field to given value.


### GetProvisionType

`func (o *CreateOsTypeImage) GetProvisionType() int64`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *CreateOsTypeImage) GetProvisionTypeOk() (*int64, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *CreateOsTypeImage) SetProvisionType(v int64)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *CreateOsTypeImage) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### SetProvisionTypeNil

`func (o *CreateOsTypeImage) SetProvisionTypeNil(b bool)`

 SetProvisionTypeNil sets the value for ProvisionType to be an explicit nil

### UnsetProvisionType
`func (o *CreateOsTypeImage) UnsetProvisionType()`

UnsetProvisionType ensures that no value is present for ProvisionType, not even an explicit nil
### GetZone

`func (o *CreateOsTypeImage) GetZone() int64`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *CreateOsTypeImage) GetZoneOk() (*int64, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *CreateOsTypeImage) SetZone(v int64)`

SetZone sets Zone field to given value.

### HasZone

`func (o *CreateOsTypeImage) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *CreateOsTypeImage) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *CreateOsTypeImage) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


