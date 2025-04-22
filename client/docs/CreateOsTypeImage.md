# CreateOsTypeImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OsType** | **int32** | id of osType | 
**VirtualImage** | **int32** | id of virtualImage | 
**ProvisionType** | Pointer to **int32** | id of provisionType | [optional] 
**Zone** | Pointer to **int32** | id of cloud/zone | [optional] 

## Methods

### NewCreateOsTypeImage

`func NewCreateOsTypeImage(osType int32, virtualImage int32, ) *CreateOsTypeImage`

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

`func (o *CreateOsTypeImage) GetOsType() int32`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *CreateOsTypeImage) GetOsTypeOk() (*int32, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *CreateOsTypeImage) SetOsType(v int32)`

SetOsType sets OsType field to given value.


### GetVirtualImage

`func (o *CreateOsTypeImage) GetVirtualImage() int32`

GetVirtualImage returns the VirtualImage field if non-nil, zero value otherwise.

### GetVirtualImageOk

`func (o *CreateOsTypeImage) GetVirtualImageOk() (*int32, bool)`

GetVirtualImageOk returns a tuple with the VirtualImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImage

`func (o *CreateOsTypeImage) SetVirtualImage(v int32)`

SetVirtualImage sets VirtualImage field to given value.


### GetProvisionType

`func (o *CreateOsTypeImage) GetProvisionType() int32`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *CreateOsTypeImage) GetProvisionTypeOk() (*int32, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *CreateOsTypeImage) SetProvisionType(v int32)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *CreateOsTypeImage) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetZone

`func (o *CreateOsTypeImage) GetZone() int32`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *CreateOsTypeImage) GetZoneOk() (*int32, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *CreateOsTypeImage) SetZone(v int32)`

SetZone sets Zone field to given value.

### HasZone

`func (o *CreateOsTypeImage) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


