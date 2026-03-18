# GetVDIPools200ResponseVdiPoolConfigVolumesInner

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

### NewGetVDIPools200ResponseVdiPoolConfigVolumesInner

`func NewGetVDIPools200ResponseVdiPoolConfigVolumesInner() *GetVDIPools200ResponseVdiPoolConfigVolumesInner`

NewGetVDIPools200ResponseVdiPoolConfigVolumesInner instantiates a new GetVDIPools200ResponseVdiPoolConfigVolumesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVDIPools200ResponseVdiPoolConfigVolumesInnerWithDefaults

`func NewGetVDIPools200ResponseVdiPoolConfigVolumesInnerWithDefaults() *GetVDIPools200ResponseVdiPoolConfigVolumesInner`

NewGetVDIPools200ResponseVdiPoolConfigVolumesInnerWithDefaults instantiates a new GetVDIPools200ResponseVdiPoolConfigVolumesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumeCustomizable

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) GetVolumeCustomizable() bool`

GetVolumeCustomizable returns the VolumeCustomizable field if non-nil, zero value otherwise.

### GetVolumeCustomizableOk

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) GetVolumeCustomizableOk() (*bool, bool)`

GetVolumeCustomizableOk returns a tuple with the VolumeCustomizable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeCustomizable

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) SetVolumeCustomizable(v bool)`

SetVolumeCustomizable sets VolumeCustomizable field to given value.

### HasVolumeCustomizable

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) HasVolumeCustomizable() bool`

HasVolumeCustomizable returns a boolean if a field has been set.

### GetVId

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) GetVId() int64`

GetVId returns the VId field if non-nil, zero value otherwise.

### GetVIdOk

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) GetVIdOk() (*int64, bool)`

GetVIdOk returns a tuple with the VId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVId

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) SetVId(v int64)`

SetVId sets VId field to given value.

### HasVId

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) HasVId() bool`

HasVId returns a boolean if a field has been set.

### GetReadonlyName

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) GetReadonlyName() bool`

GetReadonlyName returns the ReadonlyName field if non-nil, zero value otherwise.

### GetReadonlyNameOk

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) GetReadonlyNameOk() (*bool, bool)`

GetReadonlyNameOk returns a tuple with the ReadonlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonlyName

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) SetReadonlyName(v bool)`

SetReadonlyName sets ReadonlyName field to given value.

### HasReadonlyName

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) HasReadonlyName() bool`

HasReadonlyName returns a boolean if a field has been set.

### GetSize

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetMaxIOPS

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) GetMaxIOPS() string`

GetMaxIOPS returns the MaxIOPS field if non-nil, zero value otherwise.

### GetMaxIOPSOk

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) GetMaxIOPSOk() (*string, bool)`

GetMaxIOPSOk returns a tuple with the MaxIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIOPS

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) SetMaxIOPS(v string)`

SetMaxIOPS sets MaxIOPS field to given value.

### HasMaxIOPS

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) HasMaxIOPS() bool`

HasMaxIOPS returns a boolean if a field has been set.

### SetMaxIOPSNil

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) SetMaxIOPSNil(b bool)`

 SetMaxIOPSNil sets the value for MaxIOPS to be an explicit nil

### UnsetMaxIOPS
`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) UnsetMaxIOPS()`

UnsetMaxIOPS ensures that no value is present for MaxIOPS, not even an explicit nil
### GetName

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRootVolume

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) GetRootVolume() bool`

GetRootVolume returns the RootVolume field if non-nil, zero value otherwise.

### GetRootVolumeOk

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) GetRootVolumeOk() (*bool, bool)`

GetRootVolumeOk returns a tuple with the RootVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolume

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) SetRootVolume(v bool)`

SetRootVolume sets RootVolume field to given value.

### HasRootVolume

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) HasRootVolume() bool`

HasRootVolume returns a boolean if a field has been set.

### GetStorageType

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) GetStorageType() int64`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) GetStorageTypeOk() (*int64, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) SetStorageType(v int64)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetDatastoreId

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) GetDatastoreId() string`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) GetDatastoreIdOk() (*string, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) SetDatastoreId(v string)`

SetDatastoreId sets DatastoreId field to given value.

### HasDatastoreId

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) HasDatastoreId() bool`

HasDatastoreId returns a boolean if a field has been set.

### GetMaxStorage

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *GetVDIPools200ResponseVdiPoolConfigVolumesInner) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


