# UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VolumeCustomizable** | Pointer to **bool** |  | [optional] 
**VId** | Pointer to **int64** |  | [optional] 
**ReadonlyName** | Pointer to **bool** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**MaxIOPS** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RootVolume** | Pointer to **bool** |  | [optional] 
**StorageType** | Pointer to **int64** |  | [optional] 
**DatastoreId** | Pointer to **string** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 

## Methods

### NewUpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner

`func NewUpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner() *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner`

NewUpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner instantiates a new UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInnerWithDefaults

`func NewUpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInnerWithDefaults() *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner`

NewUpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInnerWithDefaults instantiates a new UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumeCustomizable

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) GetVolumeCustomizable() bool`

GetVolumeCustomizable returns the VolumeCustomizable field if non-nil, zero value otherwise.

### GetVolumeCustomizableOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) GetVolumeCustomizableOk() (*bool, bool)`

GetVolumeCustomizableOk returns a tuple with the VolumeCustomizable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeCustomizable

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) SetVolumeCustomizable(v bool)`

SetVolumeCustomizable sets VolumeCustomizable field to given value.

### HasVolumeCustomizable

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) HasVolumeCustomizable() bool`

HasVolumeCustomizable returns a boolean if a field has been set.

### GetVId

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) GetVId() int64`

GetVId returns the VId field if non-nil, zero value otherwise.

### GetVIdOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) GetVIdOk() (*int64, bool)`

GetVIdOk returns a tuple with the VId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVId

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) SetVId(v int64)`

SetVId sets VId field to given value.

### HasVId

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) HasVId() bool`

HasVId returns a boolean if a field has been set.

### GetReadonlyName

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) GetReadonlyName() bool`

GetReadonlyName returns the ReadonlyName field if non-nil, zero value otherwise.

### GetReadonlyNameOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) GetReadonlyNameOk() (*bool, bool)`

GetReadonlyNameOk returns a tuple with the ReadonlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonlyName

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) SetReadonlyName(v bool)`

SetReadonlyName sets ReadonlyName field to given value.

### HasReadonlyName

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) HasReadonlyName() bool`

HasReadonlyName returns a boolean if a field has been set.

### GetSize

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetMaxIOPS

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) GetMaxIOPS() string`

GetMaxIOPS returns the MaxIOPS field if non-nil, zero value otherwise.

### GetMaxIOPSOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) GetMaxIOPSOk() (*string, bool)`

GetMaxIOPSOk returns a tuple with the MaxIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIOPS

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) SetMaxIOPS(v string)`

SetMaxIOPS sets MaxIOPS field to given value.

### HasMaxIOPS

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) HasMaxIOPS() bool`

HasMaxIOPS returns a boolean if a field has been set.

### SetMaxIOPSNil

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) SetMaxIOPSNil(b bool)`

 SetMaxIOPSNil sets the value for MaxIOPS to be an explicit nil

### UnsetMaxIOPS
`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) UnsetMaxIOPS()`

UnsetMaxIOPS ensures that no value is present for MaxIOPS, not even an explicit nil
### GetName

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRootVolume

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) GetRootVolume() bool`

GetRootVolume returns the RootVolume field if non-nil, zero value otherwise.

### GetRootVolumeOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) GetRootVolumeOk() (*bool, bool)`

GetRootVolumeOk returns a tuple with the RootVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolume

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) SetRootVolume(v bool)`

SetRootVolume sets RootVolume field to given value.

### HasRootVolume

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) HasRootVolume() bool`

HasRootVolume returns a boolean if a field has been set.

### GetStorageType

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) GetStorageType() int64`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) GetStorageTypeOk() (*int64, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) SetStorageType(v int64)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetDatastoreId

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) GetDatastoreId() string`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) GetDatastoreIdOk() (*string, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) SetDatastoreId(v string)`

SetDatastoreId sets DatastoreId field to given value.

### HasDatastoreId

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) HasDatastoreId() bool`

HasDatastoreId returns a boolean if a field has been set.

### GetMaxStorage

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *UpdateVDIPools200ResponseAnyOfVdiPoolConfigVolumesInner) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


