# StorageVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Controller** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**ControllerId** | Pointer to **int64** |  | [optional] 
**ControllerMountPoint** | Pointer to **string** |  | [optional] 
**Resizeable** | Pointer to **bool** |  | [optional] 
**RootVolume** | Pointer to **bool** |  | [optional] 
**UnitNumber** | Pointer to **string** |  | [optional] 
**DeviceName** | Pointer to **string** |  | [optional] 
**DeviceDisplayName** | Pointer to **string** |  | [optional] 
**VolumeName** | Pointer to **string** |  | [optional] 
**VolumePath** | Pointer to **string** |  | [optional] 
**VolumeType** | Pointer to **string** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **int64** |  | [optional] 
**DiskMode** | Pointer to **string** |  | [optional] 
**DiskType** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**TypeId** | Pointer to **int64** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**ConfigurableIOPS** | Pointer to **bool** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 
**UsedStorage** | Pointer to **int64** |  | [optional] 
**DisplayOrder** | Pointer to **int64** |  | [optional] 
**MaxIOPS** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**Removable** | Pointer to **bool** |  | [optional] 
**PoolName** | Pointer to **string** |  | [optional] 
**Zone** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**ZoneId** | Pointer to **int64** |  | [optional] 
**Datastore** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**DatastoreId** | Pointer to **int64** |  | [optional] 
**DatastoreOption** | Pointer to **string** |  | [optional] 
**StorageGroup** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**StorageServer** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**UniqueId** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**ProvisionType** | Pointer to **string** |  | [optional] 
**CopyType** | Pointer to **string** |  | [optional] 
**FiberWwn** | Pointer to **string** |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**ClaimName** | Pointer to **string** |  | [optional] 
**SharePath** | Pointer to **string** |  | [optional] 
**SourceId** | Pointer to **string** |  | [optional] 
**SourceImage** | Pointer to **string** |  | [optional] 
**ImageType** | Pointer to **string** |  | [optional] 
**Online** | Pointer to **bool** |  | [optional] 
**RawData** | Pointer to **string** |  | [optional] 
**Account** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Owner** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 

## Methods

### NewStorageVolume

`func NewStorageVolume() *StorageVolume`

NewStorageVolume instantiates a new StorageVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageVolumeWithDefaults

`func NewStorageVolumeWithDefaults() *StorageVolume`

NewStorageVolumeWithDefaults instantiates a new StorageVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorageVolume) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageVolume) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageVolume) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *StorageVolume) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *StorageVolume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageVolume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageVolume) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageVolume) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *StorageVolume) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StorageVolume) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StorageVolume) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StorageVolume) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetController

`func (o *StorageVolume) GetController() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *StorageVolume) GetControllerOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *StorageVolume) SetController(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetController sets Controller field to given value.

### HasController

`func (o *StorageVolume) HasController() bool`

HasController returns a boolean if a field has been set.

### GetControllerId

`func (o *StorageVolume) GetControllerId() int64`

GetControllerId returns the ControllerId field if non-nil, zero value otherwise.

### GetControllerIdOk

`func (o *StorageVolume) GetControllerIdOk() (*int64, bool)`

GetControllerIdOk returns a tuple with the ControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerId

`func (o *StorageVolume) SetControllerId(v int64)`

SetControllerId sets ControllerId field to given value.

### HasControllerId

`func (o *StorageVolume) HasControllerId() bool`

HasControllerId returns a boolean if a field has been set.

### GetControllerMountPoint

`func (o *StorageVolume) GetControllerMountPoint() string`

GetControllerMountPoint returns the ControllerMountPoint field if non-nil, zero value otherwise.

### GetControllerMountPointOk

`func (o *StorageVolume) GetControllerMountPointOk() (*string, bool)`

GetControllerMountPointOk returns a tuple with the ControllerMountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerMountPoint

`func (o *StorageVolume) SetControllerMountPoint(v string)`

SetControllerMountPoint sets ControllerMountPoint field to given value.

### HasControllerMountPoint

`func (o *StorageVolume) HasControllerMountPoint() bool`

HasControllerMountPoint returns a boolean if a field has been set.

### GetResizeable

`func (o *StorageVolume) GetResizeable() bool`

GetResizeable returns the Resizeable field if non-nil, zero value otherwise.

### GetResizeableOk

`func (o *StorageVolume) GetResizeableOk() (*bool, bool)`

GetResizeableOk returns a tuple with the Resizeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeable

`func (o *StorageVolume) SetResizeable(v bool)`

SetResizeable sets Resizeable field to given value.

### HasResizeable

`func (o *StorageVolume) HasResizeable() bool`

HasResizeable returns a boolean if a field has been set.

### GetRootVolume

`func (o *StorageVolume) GetRootVolume() bool`

GetRootVolume returns the RootVolume field if non-nil, zero value otherwise.

### GetRootVolumeOk

`func (o *StorageVolume) GetRootVolumeOk() (*bool, bool)`

GetRootVolumeOk returns a tuple with the RootVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolume

`func (o *StorageVolume) SetRootVolume(v bool)`

SetRootVolume sets RootVolume field to given value.

### HasRootVolume

`func (o *StorageVolume) HasRootVolume() bool`

HasRootVolume returns a boolean if a field has been set.

### GetUnitNumber

`func (o *StorageVolume) GetUnitNumber() string`

GetUnitNumber returns the UnitNumber field if non-nil, zero value otherwise.

### GetUnitNumberOk

`func (o *StorageVolume) GetUnitNumberOk() (*string, bool)`

GetUnitNumberOk returns a tuple with the UnitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitNumber

`func (o *StorageVolume) SetUnitNumber(v string)`

SetUnitNumber sets UnitNumber field to given value.

### HasUnitNumber

`func (o *StorageVolume) HasUnitNumber() bool`

HasUnitNumber returns a boolean if a field has been set.

### GetDeviceName

`func (o *StorageVolume) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *StorageVolume) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *StorageVolume) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *StorageVolume) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceDisplayName

`func (o *StorageVolume) GetDeviceDisplayName() string`

GetDeviceDisplayName returns the DeviceDisplayName field if non-nil, zero value otherwise.

### GetDeviceDisplayNameOk

`func (o *StorageVolume) GetDeviceDisplayNameOk() (*string, bool)`

GetDeviceDisplayNameOk returns a tuple with the DeviceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDisplayName

`func (o *StorageVolume) SetDeviceDisplayName(v string)`

SetDeviceDisplayName sets DeviceDisplayName field to given value.

### HasDeviceDisplayName

`func (o *StorageVolume) HasDeviceDisplayName() bool`

HasDeviceDisplayName returns a boolean if a field has been set.

### GetVolumeName

`func (o *StorageVolume) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *StorageVolume) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *StorageVolume) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *StorageVolume) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.

### GetVolumePath

`func (o *StorageVolume) GetVolumePath() string`

GetVolumePath returns the VolumePath field if non-nil, zero value otherwise.

### GetVolumePathOk

`func (o *StorageVolume) GetVolumePathOk() (*string, bool)`

GetVolumePathOk returns a tuple with the VolumePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumePath

`func (o *StorageVolume) SetVolumePath(v string)`

SetVolumePath sets VolumePath field to given value.

### HasVolumePath

`func (o *StorageVolume) HasVolumePath() bool`

HasVolumePath returns a boolean if a field has been set.

### GetVolumeType

`func (o *StorageVolume) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *StorageVolume) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *StorageVolume) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *StorageVolume) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetRefType

`func (o *StorageVolume) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *StorageVolume) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *StorageVolume) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *StorageVolume) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *StorageVolume) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *StorageVolume) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *StorageVolume) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *StorageVolume) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetDiskMode

`func (o *StorageVolume) GetDiskMode() string`

GetDiskMode returns the DiskMode field if non-nil, zero value otherwise.

### GetDiskModeOk

`func (o *StorageVolume) GetDiskModeOk() (*string, bool)`

GetDiskModeOk returns a tuple with the DiskMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskMode

`func (o *StorageVolume) SetDiskMode(v string)`

SetDiskMode sets DiskMode field to given value.

### HasDiskMode

`func (o *StorageVolume) HasDiskMode() bool`

HasDiskMode returns a boolean if a field has been set.

### GetDiskType

`func (o *StorageVolume) GetDiskType() string`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *StorageVolume) GetDiskTypeOk() (*string, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *StorageVolume) SetDiskType(v string)`

SetDiskType sets DiskType field to given value.

### HasDiskType

`func (o *StorageVolume) HasDiskType() bool`

HasDiskType returns a boolean if a field has been set.

### GetType

`func (o *StorageVolume) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageVolume) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageVolume) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetType sets Type field to given value.

### HasType

`func (o *StorageVolume) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypeId

`func (o *StorageVolume) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *StorageVolume) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *StorageVolume) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *StorageVolume) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetCategory

`func (o *StorageVolume) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *StorageVolume) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *StorageVolume) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *StorageVolume) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetStatus

`func (o *StorageVolume) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StorageVolume) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StorageVolume) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StorageVolume) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *StorageVolume) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *StorageVolume) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *StorageVolume) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *StorageVolume) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetConfigurableIOPS

`func (o *StorageVolume) GetConfigurableIOPS() bool`

GetConfigurableIOPS returns the ConfigurableIOPS field if non-nil, zero value otherwise.

### GetConfigurableIOPSOk

`func (o *StorageVolume) GetConfigurableIOPSOk() (*bool, bool)`

GetConfigurableIOPSOk returns a tuple with the ConfigurableIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurableIOPS

`func (o *StorageVolume) SetConfigurableIOPS(v bool)`

SetConfigurableIOPS sets ConfigurableIOPS field to given value.

### HasConfigurableIOPS

`func (o *StorageVolume) HasConfigurableIOPS() bool`

HasConfigurableIOPS returns a boolean if a field has been set.

### GetMaxStorage

`func (o *StorageVolume) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *StorageVolume) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *StorageVolume) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *StorageVolume) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetUsedStorage

`func (o *StorageVolume) GetUsedStorage() int64`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *StorageVolume) GetUsedStorageOk() (*int64, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *StorageVolume) SetUsedStorage(v int64)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *StorageVolume) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *StorageVolume) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *StorageVolume) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *StorageVolume) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *StorageVolume) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetMaxIOPS

`func (o *StorageVolume) GetMaxIOPS() string`

GetMaxIOPS returns the MaxIOPS field if non-nil, zero value otherwise.

### GetMaxIOPSOk

`func (o *StorageVolume) GetMaxIOPSOk() (*string, bool)`

GetMaxIOPSOk returns a tuple with the MaxIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIOPS

`func (o *StorageVolume) SetMaxIOPS(v string)`

SetMaxIOPS sets MaxIOPS field to given value.

### HasMaxIOPS

`func (o *StorageVolume) HasMaxIOPS() bool`

HasMaxIOPS returns a boolean if a field has been set.

### GetUuid

`func (o *StorageVolume) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageVolume) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageVolume) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageVolume) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetActive

`func (o *StorageVolume) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *StorageVolume) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *StorageVolume) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *StorageVolume) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetReadOnly

`func (o *StorageVolume) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *StorageVolume) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *StorageVolume) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *StorageVolume) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetRemovable

`func (o *StorageVolume) GetRemovable() bool`

GetRemovable returns the Removable field if non-nil, zero value otherwise.

### GetRemovableOk

`func (o *StorageVolume) GetRemovableOk() (*bool, bool)`

GetRemovableOk returns a tuple with the Removable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovable

`func (o *StorageVolume) SetRemovable(v bool)`

SetRemovable sets Removable field to given value.

### HasRemovable

`func (o *StorageVolume) HasRemovable() bool`

HasRemovable returns a boolean if a field has been set.

### GetPoolName

`func (o *StorageVolume) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *StorageVolume) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *StorageVolume) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *StorageVolume) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetZone

`func (o *StorageVolume) GetZone() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *StorageVolume) GetZoneOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *StorageVolume) SetZone(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetZone sets Zone field to given value.

### HasZone

`func (o *StorageVolume) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetZoneId

`func (o *StorageVolume) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *StorageVolume) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *StorageVolume) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *StorageVolume) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetDatastore

`func (o *StorageVolume) GetDatastore() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *StorageVolume) GetDatastoreOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *StorageVolume) SetDatastore(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *StorageVolume) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetDatastoreId

`func (o *StorageVolume) GetDatastoreId() int64`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *StorageVolume) GetDatastoreIdOk() (*int64, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *StorageVolume) SetDatastoreId(v int64)`

SetDatastoreId sets DatastoreId field to given value.

### HasDatastoreId

`func (o *StorageVolume) HasDatastoreId() bool`

HasDatastoreId returns a boolean if a field has been set.

### GetDatastoreOption

`func (o *StorageVolume) GetDatastoreOption() string`

GetDatastoreOption returns the DatastoreOption field if non-nil, zero value otherwise.

### GetDatastoreOptionOk

`func (o *StorageVolume) GetDatastoreOptionOk() (*string, bool)`

GetDatastoreOptionOk returns a tuple with the DatastoreOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreOption

`func (o *StorageVolume) SetDatastoreOption(v string)`

SetDatastoreOption sets DatastoreOption field to given value.

### HasDatastoreOption

`func (o *StorageVolume) HasDatastoreOption() bool`

HasDatastoreOption returns a boolean if a field has been set.

### GetStorageGroup

`func (o *StorageVolume) GetStorageGroup() string`

GetStorageGroup returns the StorageGroup field if non-nil, zero value otherwise.

### GetStorageGroupOk

`func (o *StorageVolume) GetStorageGroupOk() (*string, bool)`

GetStorageGroupOk returns a tuple with the StorageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageGroup

`func (o *StorageVolume) SetStorageGroup(v string)`

SetStorageGroup sets StorageGroup field to given value.

### HasStorageGroup

`func (o *StorageVolume) HasStorageGroup() bool`

HasStorageGroup returns a boolean if a field has been set.

### GetNamespace

`func (o *StorageVolume) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *StorageVolume) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *StorageVolume) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *StorageVolume) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetStorageServer

`func (o *StorageVolume) GetStorageServer() string`

GetStorageServer returns the StorageServer field if non-nil, zero value otherwise.

### GetStorageServerOk

`func (o *StorageVolume) GetStorageServerOk() (*string, bool)`

GetStorageServerOk returns a tuple with the StorageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServer

`func (o *StorageVolume) SetStorageServer(v string)`

SetStorageServer sets StorageServer field to given value.

### HasStorageServer

`func (o *StorageVolume) HasStorageServer() bool`

HasStorageServer returns a boolean if a field has been set.

### GetSource

`func (o *StorageVolume) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *StorageVolume) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *StorageVolume) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *StorageVolume) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUniqueId

`func (o *StorageVolume) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *StorageVolume) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *StorageVolume) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *StorageVolume) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetInternalId

`func (o *StorageVolume) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *StorageVolume) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *StorageVolume) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *StorageVolume) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *StorageVolume) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *StorageVolume) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *StorageVolume) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *StorageVolume) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetProvisionType

`func (o *StorageVolume) GetProvisionType() string`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *StorageVolume) GetProvisionTypeOk() (*string, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *StorageVolume) SetProvisionType(v string)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *StorageVolume) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetCopyType

`func (o *StorageVolume) GetCopyType() string`

GetCopyType returns the CopyType field if non-nil, zero value otherwise.

### GetCopyTypeOk

`func (o *StorageVolume) GetCopyTypeOk() (*string, bool)`

GetCopyTypeOk returns a tuple with the CopyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyType

`func (o *StorageVolume) SetCopyType(v string)`

SetCopyType sets CopyType field to given value.

### HasCopyType

`func (o *StorageVolume) HasCopyType() bool`

HasCopyType returns a boolean if a field has been set.

### GetFiberWwn

`func (o *StorageVolume) GetFiberWwn() string`

GetFiberWwn returns the FiberWwn field if non-nil, zero value otherwise.

### GetFiberWwnOk

`func (o *StorageVolume) GetFiberWwnOk() (*string, bool)`

GetFiberWwnOk returns a tuple with the FiberWwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiberWwn

`func (o *StorageVolume) SetFiberWwn(v string)`

SetFiberWwn sets FiberWwn field to given value.

### HasFiberWwn

`func (o *StorageVolume) HasFiberWwn() bool`

HasFiberWwn returns a boolean if a field has been set.

### GetFileName

`func (o *StorageVolume) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *StorageVolume) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *StorageVolume) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *StorageVolume) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetClaimName

`func (o *StorageVolume) GetClaimName() string`

GetClaimName returns the ClaimName field if non-nil, zero value otherwise.

### GetClaimNameOk

`func (o *StorageVolume) GetClaimNameOk() (*string, bool)`

GetClaimNameOk returns a tuple with the ClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimName

`func (o *StorageVolume) SetClaimName(v string)`

SetClaimName sets ClaimName field to given value.

### HasClaimName

`func (o *StorageVolume) HasClaimName() bool`

HasClaimName returns a boolean if a field has been set.

### GetSharePath

`func (o *StorageVolume) GetSharePath() string`

GetSharePath returns the SharePath field if non-nil, zero value otherwise.

### GetSharePathOk

`func (o *StorageVolume) GetSharePathOk() (*string, bool)`

GetSharePathOk returns a tuple with the SharePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePath

`func (o *StorageVolume) SetSharePath(v string)`

SetSharePath sets SharePath field to given value.

### HasSharePath

`func (o *StorageVolume) HasSharePath() bool`

HasSharePath returns a boolean if a field has been set.

### GetSourceId

`func (o *StorageVolume) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *StorageVolume) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *StorageVolume) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *StorageVolume) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceImage

`func (o *StorageVolume) GetSourceImage() string`

GetSourceImage returns the SourceImage field if non-nil, zero value otherwise.

### GetSourceImageOk

`func (o *StorageVolume) GetSourceImageOk() (*string, bool)`

GetSourceImageOk returns a tuple with the SourceImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceImage

`func (o *StorageVolume) SetSourceImage(v string)`

SetSourceImage sets SourceImage field to given value.

### HasSourceImage

`func (o *StorageVolume) HasSourceImage() bool`

HasSourceImage returns a boolean if a field has been set.

### GetImageType

`func (o *StorageVolume) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *StorageVolume) GetImageTypeOk() (*string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *StorageVolume) SetImageType(v string)`

SetImageType sets ImageType field to given value.

### HasImageType

`func (o *StorageVolume) HasImageType() bool`

HasImageType returns a boolean if a field has been set.

### GetOnline

`func (o *StorageVolume) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *StorageVolume) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *StorageVolume) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *StorageVolume) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetRawData

`func (o *StorageVolume) GetRawData() string`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *StorageVolume) GetRawDataOk() (*string, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *StorageVolume) SetRawData(v string)`

SetRawData sets RawData field to given value.

### HasRawData

`func (o *StorageVolume) HasRawData() bool`

HasRawData returns a boolean if a field has been set.

### GetAccount

`func (o *StorageVolume) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *StorageVolume) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *StorageVolume) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *StorageVolume) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *StorageVolume) GetOwner() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *StorageVolume) GetOwnerOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *StorageVolume) SetOwner(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *StorageVolume) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


