# AddBaremetalHost200ResponseServerVolumesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Controller** | Pointer to [**AddBaremetalHost200ResponseServerVolumesInnerController**](AddBaremetalHost200ResponseServerVolumesInnerController.md) |  | [optional] 
**ControllerId** | Pointer to **NullableInt64** |  | [optional] 
**ControllerMountPoint** | Pointer to **NullableString** |  | [optional] 
**Resizeable** | Pointer to **NullableBool** |  | [optional] 
**RootVolume** | Pointer to **bool** |  | [optional] 
**UnitNumber** | Pointer to **NullableString** |  | [optional] 
**DeviceName** | Pointer to **string** |  | [optional] 
**DeviceDisplayName** | Pointer to **string** |  | [optional] 
**VolumeName** | Pointer to **string** |  | [optional] 
**VolumePath** | Pointer to **string** |  | [optional] 
**VolumeType** | Pointer to **string** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **int64** |  | [optional] 
**DiskMode** | Pointer to **string** |  | [optional] 
**DiskType** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**AddBaremetalHost200ResponseServerVolumesInnerType**](AddBaremetalHost200ResponseServerVolumesInnerType.md) |  | [optional] 
**TypeId** | Pointer to **int64** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**ConfigurableIOPS** | Pointer to **bool** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 
**UsedStorage** | Pointer to **int64** |  | [optional] 
**DisplayOrder** | Pointer to **int64** |  | [optional] 
**MaxIOPS** | Pointer to **NullableString** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**Removable** | Pointer to **bool** |  | [optional] 
**PoolName** | Pointer to **string** |  | [optional] 
**Zone** | Pointer to [**AddBaremetalHost200ResponseServerVolumesInnerZone**](AddBaremetalHost200ResponseServerVolumesInnerZone.md) |  | [optional] 
**ZoneId** | Pointer to **int64** |  | [optional] 
**Datastore** | Pointer to [**AddBaremetalHost200ResponseServerVolumesInnerDatastore**](AddBaremetalHost200ResponseServerVolumesInnerDatastore.md) |  | [optional] 
**DatastoreId** | Pointer to **NullableInt64** |  | [optional] 
**DatastoreOption** | Pointer to **string** |  | [optional] 
**StorageGroup** | Pointer to **NullableString** |  | [optional] 
**Namespace** | Pointer to **NullableString** |  | [optional] 
**StorageServer** | Pointer to **map[string]interface{}** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**UniqueId** | Pointer to **NullableString** |  | [optional] 
**InternalId** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**ProvisionType** | Pointer to **NullableString** |  | [optional] 
**CopyType** | Pointer to **NullableString** |  | [optional] 
**FiberWwn** | Pointer to **NullableString** |  | [optional] 
**FileName** | Pointer to **NullableString** |  | [optional] 
**ClaimName** | Pointer to **NullableString** |  | [optional] 
**SharePath** | Pointer to **NullableString** |  | [optional] 
**SourceId** | Pointer to **string** |  | [optional] 
**SourceImage** | Pointer to **string** |  | [optional] 
**ImageType** | Pointer to **string** |  | [optional] 
**Online** | Pointer to **bool** |  | [optional] 
**RawData** | Pointer to **string** |  | [optional] 
**Account** | Pointer to [**AddBaremetalHost200ResponseServerVolumesInnerAccount**](AddBaremetalHost200ResponseServerVolumesInnerAccount.md) |  | [optional] 
**Owner** | Pointer to [**AddBaremetalHost200ResponseServerVolumesInnerOwner**](AddBaremetalHost200ResponseServerVolumesInnerOwner.md) |  | [optional] 

## Methods

### NewAddBaremetalHost200ResponseServerVolumesInner

`func NewAddBaremetalHost200ResponseServerVolumesInner() *AddBaremetalHost200ResponseServerVolumesInner`

NewAddBaremetalHost200ResponseServerVolumesInner instantiates a new AddBaremetalHost200ResponseServerVolumesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBaremetalHost200ResponseServerVolumesInnerWithDefaults

`func NewAddBaremetalHost200ResponseServerVolumesInnerWithDefaults() *AddBaremetalHost200ResponseServerVolumesInner`

NewAddBaremetalHost200ResponseServerVolumesInnerWithDefaults instantiates a new AddBaremetalHost200ResponseServerVolumesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AddBaremetalHost200ResponseServerVolumesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetController

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetController() AddBaremetalHost200ResponseServerVolumesInnerController`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetControllerOk() (*AddBaremetalHost200ResponseServerVolumesInnerController, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetController(v AddBaremetalHost200ResponseServerVolumesInnerController)`

SetController sets Controller field to given value.

### HasController

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasController() bool`

HasController returns a boolean if a field has been set.

### GetControllerId

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetControllerId() int64`

GetControllerId returns the ControllerId field if non-nil, zero value otherwise.

### GetControllerIdOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetControllerIdOk() (*int64, bool)`

GetControllerIdOk returns a tuple with the ControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerId

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetControllerId(v int64)`

SetControllerId sets ControllerId field to given value.

### HasControllerId

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasControllerId() bool`

HasControllerId returns a boolean if a field has been set.

### SetControllerIdNil

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetControllerIdNil(b bool)`

 SetControllerIdNil sets the value for ControllerId to be an explicit nil

### UnsetControllerId
`func (o *AddBaremetalHost200ResponseServerVolumesInner) UnsetControllerId()`

UnsetControllerId ensures that no value is present for ControllerId, not even an explicit nil
### GetControllerMountPoint

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetControllerMountPoint() string`

GetControllerMountPoint returns the ControllerMountPoint field if non-nil, zero value otherwise.

### GetControllerMountPointOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetControllerMountPointOk() (*string, bool)`

GetControllerMountPointOk returns a tuple with the ControllerMountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerMountPoint

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetControllerMountPoint(v string)`

SetControllerMountPoint sets ControllerMountPoint field to given value.

### HasControllerMountPoint

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasControllerMountPoint() bool`

HasControllerMountPoint returns a boolean if a field has been set.

### SetControllerMountPointNil

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetControllerMountPointNil(b bool)`

 SetControllerMountPointNil sets the value for ControllerMountPoint to be an explicit nil

### UnsetControllerMountPoint
`func (o *AddBaremetalHost200ResponseServerVolumesInner) UnsetControllerMountPoint()`

UnsetControllerMountPoint ensures that no value is present for ControllerMountPoint, not even an explicit nil
### GetResizeable

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetResizeable() bool`

GetResizeable returns the Resizeable field if non-nil, zero value otherwise.

### GetResizeableOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetResizeableOk() (*bool, bool)`

GetResizeableOk returns a tuple with the Resizeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeable

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetResizeable(v bool)`

SetResizeable sets Resizeable field to given value.

### HasResizeable

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasResizeable() bool`

HasResizeable returns a boolean if a field has been set.

### SetResizeableNil

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetResizeableNil(b bool)`

 SetResizeableNil sets the value for Resizeable to be an explicit nil

### UnsetResizeable
`func (o *AddBaremetalHost200ResponseServerVolumesInner) UnsetResizeable()`

UnsetResizeable ensures that no value is present for Resizeable, not even an explicit nil
### GetRootVolume

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetRootVolume() bool`

GetRootVolume returns the RootVolume field if non-nil, zero value otherwise.

### GetRootVolumeOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetRootVolumeOk() (*bool, bool)`

GetRootVolumeOk returns a tuple with the RootVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolume

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetRootVolume(v bool)`

SetRootVolume sets RootVolume field to given value.

### HasRootVolume

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasRootVolume() bool`

HasRootVolume returns a boolean if a field has been set.

### GetUnitNumber

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetUnitNumber() string`

GetUnitNumber returns the UnitNumber field if non-nil, zero value otherwise.

### GetUnitNumberOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetUnitNumberOk() (*string, bool)`

GetUnitNumberOk returns a tuple with the UnitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitNumber

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetUnitNumber(v string)`

SetUnitNumber sets UnitNumber field to given value.

### HasUnitNumber

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasUnitNumber() bool`

HasUnitNumber returns a boolean if a field has been set.

### SetUnitNumberNil

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetUnitNumberNil(b bool)`

 SetUnitNumberNil sets the value for UnitNumber to be an explicit nil

### UnsetUnitNumber
`func (o *AddBaremetalHost200ResponseServerVolumesInner) UnsetUnitNumber()`

UnsetUnitNumber ensures that no value is present for UnitNumber, not even an explicit nil
### GetDeviceName

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceDisplayName

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetDeviceDisplayName() string`

GetDeviceDisplayName returns the DeviceDisplayName field if non-nil, zero value otherwise.

### GetDeviceDisplayNameOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetDeviceDisplayNameOk() (*string, bool)`

GetDeviceDisplayNameOk returns a tuple with the DeviceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDisplayName

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetDeviceDisplayName(v string)`

SetDeviceDisplayName sets DeviceDisplayName field to given value.

### HasDeviceDisplayName

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasDeviceDisplayName() bool`

HasDeviceDisplayName returns a boolean if a field has been set.

### GetVolumeName

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.

### GetVolumePath

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetVolumePath() string`

GetVolumePath returns the VolumePath field if non-nil, zero value otherwise.

### GetVolumePathOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetVolumePathOk() (*string, bool)`

GetVolumePathOk returns a tuple with the VolumePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumePath

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetVolumePath(v string)`

SetVolumePath sets VolumePath field to given value.

### HasVolumePath

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasVolumePath() bool`

HasVolumePath returns a boolean if a field has been set.

### GetVolumeType

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetRefType

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetDiskMode

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetDiskMode() string`

GetDiskMode returns the DiskMode field if non-nil, zero value otherwise.

### GetDiskModeOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetDiskModeOk() (*string, bool)`

GetDiskModeOk returns a tuple with the DiskMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskMode

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetDiskMode(v string)`

SetDiskMode sets DiskMode field to given value.

### HasDiskMode

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasDiskMode() bool`

HasDiskMode returns a boolean if a field has been set.

### GetDiskType

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetDiskType() string`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetDiskTypeOk() (*string, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetDiskType(v string)`

SetDiskType sets DiskType field to given value.

### HasDiskType

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasDiskType() bool`

HasDiskType returns a boolean if a field has been set.

### GetType

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetType() AddBaremetalHost200ResponseServerVolumesInnerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetTypeOk() (*AddBaremetalHost200ResponseServerVolumesInnerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetType(v AddBaremetalHost200ResponseServerVolumesInnerType)`

SetType sets Type field to given value.

### HasType

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypeId

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetCategory

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetStatus

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *AddBaremetalHost200ResponseServerVolumesInner) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetConfigurableIOPS

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetConfigurableIOPS() bool`

GetConfigurableIOPS returns the ConfigurableIOPS field if non-nil, zero value otherwise.

### GetConfigurableIOPSOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetConfigurableIOPSOk() (*bool, bool)`

GetConfigurableIOPSOk returns a tuple with the ConfigurableIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurableIOPS

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetConfigurableIOPS(v bool)`

SetConfigurableIOPS sets ConfigurableIOPS field to given value.

### HasConfigurableIOPS

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasConfigurableIOPS() bool`

HasConfigurableIOPS returns a boolean if a field has been set.

### GetMaxStorage

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetUsedStorage

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetUsedStorage() int64`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetUsedStorageOk() (*int64, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetUsedStorage(v int64)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetMaxIOPS

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetMaxIOPS() string`

GetMaxIOPS returns the MaxIOPS field if non-nil, zero value otherwise.

### GetMaxIOPSOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetMaxIOPSOk() (*string, bool)`

GetMaxIOPSOk returns a tuple with the MaxIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIOPS

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetMaxIOPS(v string)`

SetMaxIOPS sets MaxIOPS field to given value.

### HasMaxIOPS

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasMaxIOPS() bool`

HasMaxIOPS returns a boolean if a field has been set.

### SetMaxIOPSNil

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetMaxIOPSNil(b bool)`

 SetMaxIOPSNil sets the value for MaxIOPS to be an explicit nil

### UnsetMaxIOPS
`func (o *AddBaremetalHost200ResponseServerVolumesInner) UnsetMaxIOPS()`

UnsetMaxIOPS ensures that no value is present for MaxIOPS, not even an explicit nil
### GetUuid

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetActive

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetReadOnly

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetRemovable

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetRemovable() bool`

GetRemovable returns the Removable field if non-nil, zero value otherwise.

### GetRemovableOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetRemovableOk() (*bool, bool)`

GetRemovableOk returns a tuple with the Removable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovable

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetRemovable(v bool)`

SetRemovable sets Removable field to given value.

### HasRemovable

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasRemovable() bool`

HasRemovable returns a boolean if a field has been set.

### GetPoolName

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetZone

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetZone() AddBaremetalHost200ResponseServerVolumesInnerZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetZoneOk() (*AddBaremetalHost200ResponseServerVolumesInnerZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetZone(v AddBaremetalHost200ResponseServerVolumesInnerZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetZoneId

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetDatastore

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetDatastore() AddBaremetalHost200ResponseServerVolumesInnerDatastore`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetDatastoreOk() (*AddBaremetalHost200ResponseServerVolumesInnerDatastore, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetDatastore(v AddBaremetalHost200ResponseServerVolumesInnerDatastore)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetDatastoreId

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetDatastoreId() int64`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetDatastoreIdOk() (*int64, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetDatastoreId(v int64)`

SetDatastoreId sets DatastoreId field to given value.

### HasDatastoreId

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasDatastoreId() bool`

HasDatastoreId returns a boolean if a field has been set.

### SetDatastoreIdNil

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetDatastoreIdNil(b bool)`

 SetDatastoreIdNil sets the value for DatastoreId to be an explicit nil

### UnsetDatastoreId
`func (o *AddBaremetalHost200ResponseServerVolumesInner) UnsetDatastoreId()`

UnsetDatastoreId ensures that no value is present for DatastoreId, not even an explicit nil
### GetDatastoreOption

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetDatastoreOption() string`

GetDatastoreOption returns the DatastoreOption field if non-nil, zero value otherwise.

### GetDatastoreOptionOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetDatastoreOptionOk() (*string, bool)`

GetDatastoreOptionOk returns a tuple with the DatastoreOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreOption

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetDatastoreOption(v string)`

SetDatastoreOption sets DatastoreOption field to given value.

### HasDatastoreOption

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasDatastoreOption() bool`

HasDatastoreOption returns a boolean if a field has been set.

### GetStorageGroup

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetStorageGroup() string`

GetStorageGroup returns the StorageGroup field if non-nil, zero value otherwise.

### GetStorageGroupOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetStorageGroupOk() (*string, bool)`

GetStorageGroupOk returns a tuple with the StorageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageGroup

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetStorageGroup(v string)`

SetStorageGroup sets StorageGroup field to given value.

### HasStorageGroup

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasStorageGroup() bool`

HasStorageGroup returns a boolean if a field has been set.

### SetStorageGroupNil

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetStorageGroupNil(b bool)`

 SetStorageGroupNil sets the value for StorageGroup to be an explicit nil

### UnsetStorageGroup
`func (o *AddBaremetalHost200ResponseServerVolumesInner) UnsetStorageGroup()`

UnsetStorageGroup ensures that no value is present for StorageGroup, not even an explicit nil
### GetNamespace

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *AddBaremetalHost200ResponseServerVolumesInner) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetStorageServer

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetStorageServer() map[string]interface{}`

GetStorageServer returns the StorageServer field if non-nil, zero value otherwise.

### GetStorageServerOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetStorageServerOk() (*map[string]interface{}, bool)`

GetStorageServerOk returns a tuple with the StorageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServer

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetStorageServer(v map[string]interface{})`

SetStorageServer sets StorageServer field to given value.

### HasStorageServer

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasStorageServer() bool`

HasStorageServer returns a boolean if a field has been set.

### GetSource

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUniqueId

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### SetUniqueIdNil

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetUniqueIdNil(b bool)`

 SetUniqueIdNil sets the value for UniqueId to be an explicit nil

### UnsetUniqueId
`func (o *AddBaremetalHost200ResponseServerVolumesInner) UnsetUniqueId()`

UnsetUniqueId ensures that no value is present for UniqueId, not even an explicit nil
### GetInternalId

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *AddBaremetalHost200ResponseServerVolumesInner) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetProvisionType

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetProvisionType() string`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetProvisionTypeOk() (*string, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetProvisionType(v string)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### SetProvisionTypeNil

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetProvisionTypeNil(b bool)`

 SetProvisionTypeNil sets the value for ProvisionType to be an explicit nil

### UnsetProvisionType
`func (o *AddBaremetalHost200ResponseServerVolumesInner) UnsetProvisionType()`

UnsetProvisionType ensures that no value is present for ProvisionType, not even an explicit nil
### GetCopyType

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetCopyType() string`

GetCopyType returns the CopyType field if non-nil, zero value otherwise.

### GetCopyTypeOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetCopyTypeOk() (*string, bool)`

GetCopyTypeOk returns a tuple with the CopyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyType

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetCopyType(v string)`

SetCopyType sets CopyType field to given value.

### HasCopyType

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasCopyType() bool`

HasCopyType returns a boolean if a field has been set.

### SetCopyTypeNil

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetCopyTypeNil(b bool)`

 SetCopyTypeNil sets the value for CopyType to be an explicit nil

### UnsetCopyType
`func (o *AddBaremetalHost200ResponseServerVolumesInner) UnsetCopyType()`

UnsetCopyType ensures that no value is present for CopyType, not even an explicit nil
### GetFiberWwn

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetFiberWwn() string`

GetFiberWwn returns the FiberWwn field if non-nil, zero value otherwise.

### GetFiberWwnOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetFiberWwnOk() (*string, bool)`

GetFiberWwnOk returns a tuple with the FiberWwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiberWwn

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetFiberWwn(v string)`

SetFiberWwn sets FiberWwn field to given value.

### HasFiberWwn

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasFiberWwn() bool`

HasFiberWwn returns a boolean if a field has been set.

### SetFiberWwnNil

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetFiberWwnNil(b bool)`

 SetFiberWwnNil sets the value for FiberWwn to be an explicit nil

### UnsetFiberWwn
`func (o *AddBaremetalHost200ResponseServerVolumesInner) UnsetFiberWwn()`

UnsetFiberWwn ensures that no value is present for FiberWwn, not even an explicit nil
### GetFileName

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *AddBaremetalHost200ResponseServerVolumesInner) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetClaimName

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetClaimName() string`

GetClaimName returns the ClaimName field if non-nil, zero value otherwise.

### GetClaimNameOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetClaimNameOk() (*string, bool)`

GetClaimNameOk returns a tuple with the ClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimName

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetClaimName(v string)`

SetClaimName sets ClaimName field to given value.

### HasClaimName

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasClaimName() bool`

HasClaimName returns a boolean if a field has been set.

### SetClaimNameNil

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetClaimNameNil(b bool)`

 SetClaimNameNil sets the value for ClaimName to be an explicit nil

### UnsetClaimName
`func (o *AddBaremetalHost200ResponseServerVolumesInner) UnsetClaimName()`

UnsetClaimName ensures that no value is present for ClaimName, not even an explicit nil
### GetSharePath

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetSharePath() string`

GetSharePath returns the SharePath field if non-nil, zero value otherwise.

### GetSharePathOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetSharePathOk() (*string, bool)`

GetSharePathOk returns a tuple with the SharePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePath

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetSharePath(v string)`

SetSharePath sets SharePath field to given value.

### HasSharePath

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasSharePath() bool`

HasSharePath returns a boolean if a field has been set.

### SetSharePathNil

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetSharePathNil(b bool)`

 SetSharePathNil sets the value for SharePath to be an explicit nil

### UnsetSharePath
`func (o *AddBaremetalHost200ResponseServerVolumesInner) UnsetSharePath()`

UnsetSharePath ensures that no value is present for SharePath, not even an explicit nil
### GetSourceId

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceImage

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetSourceImage() string`

GetSourceImage returns the SourceImage field if non-nil, zero value otherwise.

### GetSourceImageOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetSourceImageOk() (*string, bool)`

GetSourceImageOk returns a tuple with the SourceImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceImage

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetSourceImage(v string)`

SetSourceImage sets SourceImage field to given value.

### HasSourceImage

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasSourceImage() bool`

HasSourceImage returns a boolean if a field has been set.

### GetImageType

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetImageTypeOk() (*string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetImageType(v string)`

SetImageType sets ImageType field to given value.

### HasImageType

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasImageType() bool`

HasImageType returns a boolean if a field has been set.

### GetOnline

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetRawData

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetRawData() string`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetRawDataOk() (*string, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetRawData(v string)`

SetRawData sets RawData field to given value.

### HasRawData

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasRawData() bool`

HasRawData returns a boolean if a field has been set.

### GetAccount

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetAccount() AddBaremetalHost200ResponseServerVolumesInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetAccountOk() (*AddBaremetalHost200ResponseServerVolumesInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetAccount(v AddBaremetalHost200ResponseServerVolumesInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetOwner() AddBaremetalHost200ResponseServerVolumesInnerOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AddBaremetalHost200ResponseServerVolumesInner) GetOwnerOk() (*AddBaremetalHost200ResponseServerVolumesInnerOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AddBaremetalHost200ResponseServerVolumesInner) SetOwner(v AddBaremetalHost200ResponseServerVolumesInnerOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AddBaremetalHost200ResponseServerVolumesInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


