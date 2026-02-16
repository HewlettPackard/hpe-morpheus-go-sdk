# InstanceCreateVolumesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The id for the LV configuration being created. | [optional] [default to -1]
**RootVolume** | Pointer to **bool** | If set to false then a non-root LV will be created. | [optional] [default to true]
**Name** | Pointer to **string** | Name/type of the LV being created. | [optional] [default to "root"]
**Size** | Pointer to **int64** | Size of the LV to be created in GBs.  Uses default from service plan. | [optional] 
**SizeId** | Pointer to **NullableInt64** | Can be used to select pre-existing LV choices from Morpheus. | [optional] 
**StorageType** | Pointer to **NullableInt64** | Identifier for LV type | [optional] 
**DatastoreId** | Pointer to [**InstanceCreateVolumesInnerDatastoreId**](InstanceCreateVolumesInnerDatastoreId.md) |  | [optional] 
**ControllerMountPoint** | Pointer to **string** | The controller mount point specification for this volume in the format: &#x60;\&quot;id:busNumber:typeId:unitNumber\&quot;&#x60; For new storage controllers the id is passed as -1, so an example value would be: &#x60;\&quot;-1:1:6:0\&quot;&#x60; which translates to id: -1 (new), busNumber: 1, storage controller type id: 6 (SCSI VMware Paravirtual), unit number: 0. The current list of storage controllers is returned for instances and servers for determining existing id values. Use &#x60;/api/provision-types?code&#x3D;vmware&#x60; to see the available &#x60;controllerTypes&#x60; for vmware. | [optional] 

## Methods

### NewInstanceCreateVolumesInner

`func NewInstanceCreateVolumesInner() *InstanceCreateVolumesInner`

NewInstanceCreateVolumesInner instantiates a new InstanceCreateVolumesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceCreateVolumesInnerWithDefaults

`func NewInstanceCreateVolumesInnerWithDefaults() *InstanceCreateVolumesInner`

NewInstanceCreateVolumesInnerWithDefaults instantiates a new InstanceCreateVolumesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceCreateVolumesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceCreateVolumesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceCreateVolumesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceCreateVolumesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRootVolume

`func (o *InstanceCreateVolumesInner) GetRootVolume() bool`

GetRootVolume returns the RootVolume field if non-nil, zero value otherwise.

### GetRootVolumeOk

`func (o *InstanceCreateVolumesInner) GetRootVolumeOk() (*bool, bool)`

GetRootVolumeOk returns a tuple with the RootVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolume

`func (o *InstanceCreateVolumesInner) SetRootVolume(v bool)`

SetRootVolume sets RootVolume field to given value.

### HasRootVolume

`func (o *InstanceCreateVolumesInner) HasRootVolume() bool`

HasRootVolume returns a boolean if a field has been set.

### GetName

`func (o *InstanceCreateVolumesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceCreateVolumesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceCreateVolumesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceCreateVolumesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *InstanceCreateVolumesInner) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *InstanceCreateVolumesInner) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *InstanceCreateVolumesInner) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *InstanceCreateVolumesInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSizeId

`func (o *InstanceCreateVolumesInner) GetSizeId() int64`

GetSizeId returns the SizeId field if non-nil, zero value otherwise.

### GetSizeIdOk

`func (o *InstanceCreateVolumesInner) GetSizeIdOk() (*int64, bool)`

GetSizeIdOk returns a tuple with the SizeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeId

`func (o *InstanceCreateVolumesInner) SetSizeId(v int64)`

SetSizeId sets SizeId field to given value.

### HasSizeId

`func (o *InstanceCreateVolumesInner) HasSizeId() bool`

HasSizeId returns a boolean if a field has been set.

### SetSizeIdNil

`func (o *InstanceCreateVolumesInner) SetSizeIdNil(b bool)`

 SetSizeIdNil sets the value for SizeId to be an explicit nil

### UnsetSizeId
`func (o *InstanceCreateVolumesInner) UnsetSizeId()`

UnsetSizeId ensures that no value is present for SizeId, not even an explicit nil
### GetStorageType

`func (o *InstanceCreateVolumesInner) GetStorageType() int64`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *InstanceCreateVolumesInner) GetStorageTypeOk() (*int64, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *InstanceCreateVolumesInner) SetStorageType(v int64)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *InstanceCreateVolumesInner) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### SetStorageTypeNil

`func (o *InstanceCreateVolumesInner) SetStorageTypeNil(b bool)`

 SetStorageTypeNil sets the value for StorageType to be an explicit nil

### UnsetStorageType
`func (o *InstanceCreateVolumesInner) UnsetStorageType()`

UnsetStorageType ensures that no value is present for StorageType, not even an explicit nil
### GetDatastoreId

`func (o *InstanceCreateVolumesInner) GetDatastoreId() InstanceCreateVolumesInnerDatastoreId`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *InstanceCreateVolumesInner) GetDatastoreIdOk() (*InstanceCreateVolumesInnerDatastoreId, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *InstanceCreateVolumesInner) SetDatastoreId(v InstanceCreateVolumesInnerDatastoreId)`

SetDatastoreId sets DatastoreId field to given value.

### HasDatastoreId

`func (o *InstanceCreateVolumesInner) HasDatastoreId() bool`

HasDatastoreId returns a boolean if a field has been set.

### GetControllerMountPoint

`func (o *InstanceCreateVolumesInner) GetControllerMountPoint() string`

GetControllerMountPoint returns the ControllerMountPoint field if non-nil, zero value otherwise.

### GetControllerMountPointOk

`func (o *InstanceCreateVolumesInner) GetControllerMountPointOk() (*string, bool)`

GetControllerMountPointOk returns a tuple with the ControllerMountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerMountPoint

`func (o *InstanceCreateVolumesInner) SetControllerMountPoint(v string)`

SetControllerMountPoint sets ControllerMountPoint field to given value.

### HasControllerMountPoint

`func (o *InstanceCreateVolumesInner) HasControllerMountPoint() bool`

HasControllerMountPoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


