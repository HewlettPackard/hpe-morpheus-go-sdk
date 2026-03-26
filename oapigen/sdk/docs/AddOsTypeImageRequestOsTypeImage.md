# AddOsTypeImageRequestOsTypeImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OsType** | **int64** | id of osType | 
**VirtualImage** | **int64** | id of virtualImage | 
**ProvisionType** | Pointer to **NullableInt64** | id of provisionType | [optional] 
**Zone** | Pointer to **NullableInt64** | id of cloud/zone | [optional] 

## Methods

### NewAddOsTypeImageRequestOsTypeImage

`func NewAddOsTypeImageRequestOsTypeImage(osType int64, virtualImage int64, ) *AddOsTypeImageRequestOsTypeImage`

NewAddOsTypeImageRequestOsTypeImage instantiates a new AddOsTypeImageRequestOsTypeImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddOsTypeImageRequestOsTypeImageWithDefaults

`func NewAddOsTypeImageRequestOsTypeImageWithDefaults() *AddOsTypeImageRequestOsTypeImage`

NewAddOsTypeImageRequestOsTypeImageWithDefaults instantiates a new AddOsTypeImageRequestOsTypeImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOsType

`func (o *AddOsTypeImageRequestOsTypeImage) GetOsType() int64`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *AddOsTypeImageRequestOsTypeImage) GetOsTypeOk() (*int64, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *AddOsTypeImageRequestOsTypeImage) SetOsType(v int64)`

SetOsType sets OsType field to given value.


### GetVirtualImage

`func (o *AddOsTypeImageRequestOsTypeImage) GetVirtualImage() int64`

GetVirtualImage returns the VirtualImage field if non-nil, zero value otherwise.

### GetVirtualImageOk

`func (o *AddOsTypeImageRequestOsTypeImage) GetVirtualImageOk() (*int64, bool)`

GetVirtualImageOk returns a tuple with the VirtualImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImage

`func (o *AddOsTypeImageRequestOsTypeImage) SetVirtualImage(v int64)`

SetVirtualImage sets VirtualImage field to given value.


### GetProvisionType

`func (o *AddOsTypeImageRequestOsTypeImage) GetProvisionType() int64`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *AddOsTypeImageRequestOsTypeImage) GetProvisionTypeOk() (*int64, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *AddOsTypeImageRequestOsTypeImage) SetProvisionType(v int64)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *AddOsTypeImageRequestOsTypeImage) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### SetProvisionTypeNil

`func (o *AddOsTypeImageRequestOsTypeImage) SetProvisionTypeNil(b bool)`

 SetProvisionTypeNil sets the value for ProvisionType to be an explicit nil

### UnsetProvisionType
`func (o *AddOsTypeImageRequestOsTypeImage) UnsetProvisionType()`

UnsetProvisionType ensures that no value is present for ProvisionType, not even an explicit nil
### GetZone

`func (o *AddOsTypeImageRequestOsTypeImage) GetZone() int64`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *AddOsTypeImageRequestOsTypeImage) GetZoneOk() (*int64, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *AddOsTypeImageRequestOsTypeImage) SetZone(v int64)`

SetZone sets Zone field to given value.

### HasZone

`func (o *AddOsTypeImageRequestOsTypeImage) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *AddOsTypeImageRequestOsTypeImage) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *AddOsTypeImageRequestOsTypeImage) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


