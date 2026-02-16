# InstanceContainerServerVolume3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ControllerId** | Pointer to **NullableInt64** |  | [optional] 
**ControllerMountPoint** | Pointer to **NullableString** |  | [optional] 
**Resizeable** | Pointer to **bool** |  | [optional] 
**RootVolume** | Pointer to **bool** |  | [optional] 
**DeviceName** | Pointer to **string** |  | [optional] 
**DeviceDisplayName** | Pointer to **string** |  | [optional] 
**DiskMode** | Pointer to **NullableString** |  | [optional] 
**DiskType** | Pointer to **NullableString** |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**StorageServer** | Pointer to [**InstanceContainerServerVolumeStorageServer1**](InstanceContainerServerVolumeStorageServer1.md) |  | [optional] 
**ZoneId** | Pointer to **NullableInt64** |  | [optional] 
**Zone** | Pointer to [**InstanceContainerServerVolumeZone1**](InstanceContainerServerVolumeZone1.md) |  | [optional] 
**Datastore** | Pointer to [**InstanceContainerServerVolumeDatastore1**](InstanceContainerServerVolumeDatastore1.md) |  | [optional] 
**UnitNumber** | Pointer to **NullableString** |  | [optional] 
**TypeId** | Pointer to **int64** |  | [optional] 
**ConfigurableIOPS** | Pointer to **bool** |  | [optional] 
**DatastoreId** | Pointer to **NullableInt64** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 
**MaxIOPS** | Pointer to **NullableString** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**UniqueId** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**InternalId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInstanceContainerServerVolume3

`func NewInstanceContainerServerVolume3() *InstanceContainerServerVolume3`

NewInstanceContainerServerVolume3 instantiates a new InstanceContainerServerVolume3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceContainerServerVolume3WithDefaults

`func NewInstanceContainerServerVolume3WithDefaults() *InstanceContainerServerVolume3`

NewInstanceContainerServerVolume3WithDefaults instantiates a new InstanceContainerServerVolume3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceContainerServerVolume3) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceContainerServerVolume3) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceContainerServerVolume3) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceContainerServerVolume3) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InstanceContainerServerVolume3) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceContainerServerVolume3) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceContainerServerVolume3) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceContainerServerVolume3) HasName() bool`

HasName returns a boolean if a field has been set.

### GetControllerId

`func (o *InstanceContainerServerVolume3) GetControllerId() int64`

GetControllerId returns the ControllerId field if non-nil, zero value otherwise.

### GetControllerIdOk

`func (o *InstanceContainerServerVolume3) GetControllerIdOk() (*int64, bool)`

GetControllerIdOk returns a tuple with the ControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerId

`func (o *InstanceContainerServerVolume3) SetControllerId(v int64)`

SetControllerId sets ControllerId field to given value.

### HasControllerId

`func (o *InstanceContainerServerVolume3) HasControllerId() bool`

HasControllerId returns a boolean if a field has been set.

### SetControllerIdNil

`func (o *InstanceContainerServerVolume3) SetControllerIdNil(b bool)`

 SetControllerIdNil sets the value for ControllerId to be an explicit nil

### UnsetControllerId
`func (o *InstanceContainerServerVolume3) UnsetControllerId()`

UnsetControllerId ensures that no value is present for ControllerId, not even an explicit nil
### GetControllerMountPoint

`func (o *InstanceContainerServerVolume3) GetControllerMountPoint() string`

GetControllerMountPoint returns the ControllerMountPoint field if non-nil, zero value otherwise.

### GetControllerMountPointOk

`func (o *InstanceContainerServerVolume3) GetControllerMountPointOk() (*string, bool)`

GetControllerMountPointOk returns a tuple with the ControllerMountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerMountPoint

`func (o *InstanceContainerServerVolume3) SetControllerMountPoint(v string)`

SetControllerMountPoint sets ControllerMountPoint field to given value.

### HasControllerMountPoint

`func (o *InstanceContainerServerVolume3) HasControllerMountPoint() bool`

HasControllerMountPoint returns a boolean if a field has been set.

### SetControllerMountPointNil

`func (o *InstanceContainerServerVolume3) SetControllerMountPointNil(b bool)`

 SetControllerMountPointNil sets the value for ControllerMountPoint to be an explicit nil

### UnsetControllerMountPoint
`func (o *InstanceContainerServerVolume3) UnsetControllerMountPoint()`

UnsetControllerMountPoint ensures that no value is present for ControllerMountPoint, not even an explicit nil
### GetResizeable

`func (o *InstanceContainerServerVolume3) GetResizeable() bool`

GetResizeable returns the Resizeable field if non-nil, zero value otherwise.

### GetResizeableOk

`func (o *InstanceContainerServerVolume3) GetResizeableOk() (*bool, bool)`

GetResizeableOk returns a tuple with the Resizeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeable

`func (o *InstanceContainerServerVolume3) SetResizeable(v bool)`

SetResizeable sets Resizeable field to given value.

### HasResizeable

`func (o *InstanceContainerServerVolume3) HasResizeable() bool`

HasResizeable returns a boolean if a field has been set.

### GetRootVolume

`func (o *InstanceContainerServerVolume3) GetRootVolume() bool`

GetRootVolume returns the RootVolume field if non-nil, zero value otherwise.

### GetRootVolumeOk

`func (o *InstanceContainerServerVolume3) GetRootVolumeOk() (*bool, bool)`

GetRootVolumeOk returns a tuple with the RootVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolume

`func (o *InstanceContainerServerVolume3) SetRootVolume(v bool)`

SetRootVolume sets RootVolume field to given value.

### HasRootVolume

`func (o *InstanceContainerServerVolume3) HasRootVolume() bool`

HasRootVolume returns a boolean if a field has been set.

### GetDeviceName

`func (o *InstanceContainerServerVolume3) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *InstanceContainerServerVolume3) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *InstanceContainerServerVolume3) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *InstanceContainerServerVolume3) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceDisplayName

`func (o *InstanceContainerServerVolume3) GetDeviceDisplayName() string`

GetDeviceDisplayName returns the DeviceDisplayName field if non-nil, zero value otherwise.

### GetDeviceDisplayNameOk

`func (o *InstanceContainerServerVolume3) GetDeviceDisplayNameOk() (*string, bool)`

GetDeviceDisplayNameOk returns a tuple with the DeviceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDisplayName

`func (o *InstanceContainerServerVolume3) SetDeviceDisplayName(v string)`

SetDeviceDisplayName sets DeviceDisplayName field to given value.

### HasDeviceDisplayName

`func (o *InstanceContainerServerVolume3) HasDeviceDisplayName() bool`

HasDeviceDisplayName returns a boolean if a field has been set.

### GetDiskMode

`func (o *InstanceContainerServerVolume3) GetDiskMode() string`

GetDiskMode returns the DiskMode field if non-nil, zero value otherwise.

### GetDiskModeOk

`func (o *InstanceContainerServerVolume3) GetDiskModeOk() (*string, bool)`

GetDiskModeOk returns a tuple with the DiskMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskMode

`func (o *InstanceContainerServerVolume3) SetDiskMode(v string)`

SetDiskMode sets DiskMode field to given value.

### HasDiskMode

`func (o *InstanceContainerServerVolume3) HasDiskMode() bool`

HasDiskMode returns a boolean if a field has been set.

### SetDiskModeNil

`func (o *InstanceContainerServerVolume3) SetDiskModeNil(b bool)`

 SetDiskModeNil sets the value for DiskMode to be an explicit nil

### UnsetDiskMode
`func (o *InstanceContainerServerVolume3) UnsetDiskMode()`

UnsetDiskMode ensures that no value is present for DiskMode, not even an explicit nil
### GetDiskType

`func (o *InstanceContainerServerVolume3) GetDiskType() string`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *InstanceContainerServerVolume3) GetDiskTypeOk() (*string, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *InstanceContainerServerVolume3) SetDiskType(v string)`

SetDiskType sets DiskType field to given value.

### HasDiskType

`func (o *InstanceContainerServerVolume3) HasDiskType() bool`

HasDiskType returns a boolean if a field has been set.

### SetDiskTypeNil

`func (o *InstanceContainerServerVolume3) SetDiskTypeNil(b bool)`

 SetDiskTypeNil sets the value for DiskType to be an explicit nil

### UnsetDiskType
`func (o *InstanceContainerServerVolume3) UnsetDiskType()`

UnsetDiskType ensures that no value is present for DiskType, not even an explicit nil
### GetCategory

`func (o *InstanceContainerServerVolume3) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InstanceContainerServerVolume3) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InstanceContainerServerVolume3) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *InstanceContainerServerVolume3) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *InstanceContainerServerVolume3) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *InstanceContainerServerVolume3) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetStorageServer

`func (o *InstanceContainerServerVolume3) GetStorageServer() InstanceContainerServerVolumeStorageServer1`

GetStorageServer returns the StorageServer field if non-nil, zero value otherwise.

### GetStorageServerOk

`func (o *InstanceContainerServerVolume3) GetStorageServerOk() (*InstanceContainerServerVolumeStorageServer1, bool)`

GetStorageServerOk returns a tuple with the StorageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServer

`func (o *InstanceContainerServerVolume3) SetStorageServer(v InstanceContainerServerVolumeStorageServer1)`

SetStorageServer sets StorageServer field to given value.

### HasStorageServer

`func (o *InstanceContainerServerVolume3) HasStorageServer() bool`

HasStorageServer returns a boolean if a field has been set.

### GetZoneId

`func (o *InstanceContainerServerVolume3) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *InstanceContainerServerVolume3) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *InstanceContainerServerVolume3) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *InstanceContainerServerVolume3) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### SetZoneIdNil

`func (o *InstanceContainerServerVolume3) SetZoneIdNil(b bool)`

 SetZoneIdNil sets the value for ZoneId to be an explicit nil

### UnsetZoneId
`func (o *InstanceContainerServerVolume3) UnsetZoneId()`

UnsetZoneId ensures that no value is present for ZoneId, not even an explicit nil
### GetZone

`func (o *InstanceContainerServerVolume3) GetZone() InstanceContainerServerVolumeZone1`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *InstanceContainerServerVolume3) GetZoneOk() (*InstanceContainerServerVolumeZone1, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *InstanceContainerServerVolume3) SetZone(v InstanceContainerServerVolumeZone1)`

SetZone sets Zone field to given value.

### HasZone

`func (o *InstanceContainerServerVolume3) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetDatastore

`func (o *InstanceContainerServerVolume3) GetDatastore() InstanceContainerServerVolumeDatastore1`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *InstanceContainerServerVolume3) GetDatastoreOk() (*InstanceContainerServerVolumeDatastore1, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *InstanceContainerServerVolume3) SetDatastore(v InstanceContainerServerVolumeDatastore1)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *InstanceContainerServerVolume3) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetUnitNumber

`func (o *InstanceContainerServerVolume3) GetUnitNumber() string`

GetUnitNumber returns the UnitNumber field if non-nil, zero value otherwise.

### GetUnitNumberOk

`func (o *InstanceContainerServerVolume3) GetUnitNumberOk() (*string, bool)`

GetUnitNumberOk returns a tuple with the UnitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitNumber

`func (o *InstanceContainerServerVolume3) SetUnitNumber(v string)`

SetUnitNumber sets UnitNumber field to given value.

### HasUnitNumber

`func (o *InstanceContainerServerVolume3) HasUnitNumber() bool`

HasUnitNumber returns a boolean if a field has been set.

### SetUnitNumberNil

`func (o *InstanceContainerServerVolume3) SetUnitNumberNil(b bool)`

 SetUnitNumberNil sets the value for UnitNumber to be an explicit nil

### UnsetUnitNumber
`func (o *InstanceContainerServerVolume3) UnsetUnitNumber()`

UnsetUnitNumber ensures that no value is present for UnitNumber, not even an explicit nil
### GetTypeId

`func (o *InstanceContainerServerVolume3) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *InstanceContainerServerVolume3) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *InstanceContainerServerVolume3) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *InstanceContainerServerVolume3) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetConfigurableIOPS

`func (o *InstanceContainerServerVolume3) GetConfigurableIOPS() bool`

GetConfigurableIOPS returns the ConfigurableIOPS field if non-nil, zero value otherwise.

### GetConfigurableIOPSOk

`func (o *InstanceContainerServerVolume3) GetConfigurableIOPSOk() (*bool, bool)`

GetConfigurableIOPSOk returns a tuple with the ConfigurableIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurableIOPS

`func (o *InstanceContainerServerVolume3) SetConfigurableIOPS(v bool)`

SetConfigurableIOPS sets ConfigurableIOPS field to given value.

### HasConfigurableIOPS

`func (o *InstanceContainerServerVolume3) HasConfigurableIOPS() bool`

HasConfigurableIOPS returns a boolean if a field has been set.

### GetDatastoreId

`func (o *InstanceContainerServerVolume3) GetDatastoreId() int64`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *InstanceContainerServerVolume3) GetDatastoreIdOk() (*int64, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *InstanceContainerServerVolume3) SetDatastoreId(v int64)`

SetDatastoreId sets DatastoreId field to given value.

### HasDatastoreId

`func (o *InstanceContainerServerVolume3) HasDatastoreId() bool`

HasDatastoreId returns a boolean if a field has been set.

### SetDatastoreIdNil

`func (o *InstanceContainerServerVolume3) SetDatastoreIdNil(b bool)`

 SetDatastoreIdNil sets the value for DatastoreId to be an explicit nil

### UnsetDatastoreId
`func (o *InstanceContainerServerVolume3) UnsetDatastoreId()`

UnsetDatastoreId ensures that no value is present for DatastoreId, not even an explicit nil
### GetMaxStorage

`func (o *InstanceContainerServerVolume3) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *InstanceContainerServerVolume3) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *InstanceContainerServerVolume3) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *InstanceContainerServerVolume3) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMaxIOPS

`func (o *InstanceContainerServerVolume3) GetMaxIOPS() string`

GetMaxIOPS returns the MaxIOPS field if non-nil, zero value otherwise.

### GetMaxIOPSOk

`func (o *InstanceContainerServerVolume3) GetMaxIOPSOk() (*string, bool)`

GetMaxIOPSOk returns a tuple with the MaxIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIOPS

`func (o *InstanceContainerServerVolume3) SetMaxIOPS(v string)`

SetMaxIOPS sets MaxIOPS field to given value.

### HasMaxIOPS

`func (o *InstanceContainerServerVolume3) HasMaxIOPS() bool`

HasMaxIOPS returns a boolean if a field has been set.

### SetMaxIOPSNil

`func (o *InstanceContainerServerVolume3) SetMaxIOPSNil(b bool)`

 SetMaxIOPSNil sets the value for MaxIOPS to be an explicit nil

### UnsetMaxIOPS
`func (o *InstanceContainerServerVolume3) UnsetMaxIOPS()`

UnsetMaxIOPS ensures that no value is present for MaxIOPS, not even an explicit nil
### GetUuid

`func (o *InstanceContainerServerVolume3) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *InstanceContainerServerVolume3) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *InstanceContainerServerVolume3) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *InstanceContainerServerVolume3) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetUniqueId

`func (o *InstanceContainerServerVolume3) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *InstanceContainerServerVolume3) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *InstanceContainerServerVolume3) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *InstanceContainerServerVolume3) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetExternalId

`func (o *InstanceContainerServerVolume3) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *InstanceContainerServerVolume3) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *InstanceContainerServerVolume3) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *InstanceContainerServerVolume3) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *InstanceContainerServerVolume3) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *InstanceContainerServerVolume3) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetInternalId

`func (o *InstanceContainerServerVolume3) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *InstanceContainerServerVolume3) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *InstanceContainerServerVolume3) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *InstanceContainerServerVolume3) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *InstanceContainerServerVolume3) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *InstanceContainerServerVolume3) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


