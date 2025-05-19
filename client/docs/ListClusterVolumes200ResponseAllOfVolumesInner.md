# ListClusterVolumes200ResponseAllOfVolumesInner

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

### NewListClusterVolumes200ResponseAllOfVolumesInner

`func NewListClusterVolumes200ResponseAllOfVolumesInner() *ListClusterVolumes200ResponseAllOfVolumesInner`

NewListClusterVolumes200ResponseAllOfVolumesInner instantiates a new ListClusterVolumes200ResponseAllOfVolumesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterVolumes200ResponseAllOfVolumesInnerWithDefaults

`func NewListClusterVolumes200ResponseAllOfVolumesInnerWithDefaults() *ListClusterVolumes200ResponseAllOfVolumesInner`

NewListClusterVolumes200ResponseAllOfVolumesInnerWithDefaults instantiates a new ListClusterVolumes200ResponseAllOfVolumesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetController

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetController() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetControllerOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetController(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetController sets Controller field to given value.

### HasController

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasController() bool`

HasController returns a boolean if a field has been set.

### GetControllerId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetControllerId() int64`

GetControllerId returns the ControllerId field if non-nil, zero value otherwise.

### GetControllerIdOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetControllerIdOk() (*int64, bool)`

GetControllerIdOk returns a tuple with the ControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetControllerId(v int64)`

SetControllerId sets ControllerId field to given value.

### HasControllerId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasControllerId() bool`

HasControllerId returns a boolean if a field has been set.

### GetControllerMountPoint

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetControllerMountPoint() string`

GetControllerMountPoint returns the ControllerMountPoint field if non-nil, zero value otherwise.

### GetControllerMountPointOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetControllerMountPointOk() (*string, bool)`

GetControllerMountPointOk returns a tuple with the ControllerMountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerMountPoint

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetControllerMountPoint(v string)`

SetControllerMountPoint sets ControllerMountPoint field to given value.

### HasControllerMountPoint

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasControllerMountPoint() bool`

HasControllerMountPoint returns a boolean if a field has been set.

### GetResizeable

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetResizeable() bool`

GetResizeable returns the Resizeable field if non-nil, zero value otherwise.

### GetResizeableOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetResizeableOk() (*bool, bool)`

GetResizeableOk returns a tuple with the Resizeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeable

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetResizeable(v bool)`

SetResizeable sets Resizeable field to given value.

### HasResizeable

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasResizeable() bool`

HasResizeable returns a boolean if a field has been set.

### GetRootVolume

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetRootVolume() bool`

GetRootVolume returns the RootVolume field if non-nil, zero value otherwise.

### GetRootVolumeOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetRootVolumeOk() (*bool, bool)`

GetRootVolumeOk returns a tuple with the RootVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolume

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetRootVolume(v bool)`

SetRootVolume sets RootVolume field to given value.

### HasRootVolume

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasRootVolume() bool`

HasRootVolume returns a boolean if a field has been set.

### GetUnitNumber

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetUnitNumber() string`

GetUnitNumber returns the UnitNumber field if non-nil, zero value otherwise.

### GetUnitNumberOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetUnitNumberOk() (*string, bool)`

GetUnitNumberOk returns a tuple with the UnitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitNumber

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetUnitNumber(v string)`

SetUnitNumber sets UnitNumber field to given value.

### HasUnitNumber

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasUnitNumber() bool`

HasUnitNumber returns a boolean if a field has been set.

### GetDeviceName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceDisplayName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetDeviceDisplayName() string`

GetDeviceDisplayName returns the DeviceDisplayName field if non-nil, zero value otherwise.

### GetDeviceDisplayNameOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetDeviceDisplayNameOk() (*string, bool)`

GetDeviceDisplayNameOk returns a tuple with the DeviceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDisplayName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetDeviceDisplayName(v string)`

SetDeviceDisplayName sets DeviceDisplayName field to given value.

### HasDeviceDisplayName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasDeviceDisplayName() bool`

HasDeviceDisplayName returns a boolean if a field has been set.

### GetVolumeName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.

### GetVolumePath

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetVolumePath() string`

GetVolumePath returns the VolumePath field if non-nil, zero value otherwise.

### GetVolumePathOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetVolumePathOk() (*string, bool)`

GetVolumePathOk returns a tuple with the VolumePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumePath

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetVolumePath(v string)`

SetVolumePath sets VolumePath field to given value.

### HasVolumePath

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasVolumePath() bool`

HasVolumePath returns a boolean if a field has been set.

### GetVolumeType

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetRefType

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetDiskMode

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetDiskMode() string`

GetDiskMode returns the DiskMode field if non-nil, zero value otherwise.

### GetDiskModeOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetDiskModeOk() (*string, bool)`

GetDiskModeOk returns a tuple with the DiskMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskMode

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetDiskMode(v string)`

SetDiskMode sets DiskMode field to given value.

### HasDiskMode

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasDiskMode() bool`

HasDiskMode returns a boolean if a field has been set.

### GetDiskType

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetDiskType() string`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetDiskTypeOk() (*string, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetDiskType(v string)`

SetDiskType sets DiskType field to given value.

### HasDiskType

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasDiskType() bool`

HasDiskType returns a boolean if a field has been set.

### GetType

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetType sets Type field to given value.

### HasType

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypeId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetCategory

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetStatus

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetConfigurableIOPS

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetConfigurableIOPS() bool`

GetConfigurableIOPS returns the ConfigurableIOPS field if non-nil, zero value otherwise.

### GetConfigurableIOPSOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetConfigurableIOPSOk() (*bool, bool)`

GetConfigurableIOPSOk returns a tuple with the ConfigurableIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurableIOPS

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetConfigurableIOPS(v bool)`

SetConfigurableIOPS sets ConfigurableIOPS field to given value.

### HasConfigurableIOPS

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasConfigurableIOPS() bool`

HasConfigurableIOPS returns a boolean if a field has been set.

### GetMaxStorage

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetUsedStorage

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetUsedStorage() int64`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetUsedStorageOk() (*int64, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetUsedStorage(v int64)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetMaxIOPS

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetMaxIOPS() string`

GetMaxIOPS returns the MaxIOPS field if non-nil, zero value otherwise.

### GetMaxIOPSOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetMaxIOPSOk() (*string, bool)`

GetMaxIOPSOk returns a tuple with the MaxIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIOPS

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetMaxIOPS(v string)`

SetMaxIOPS sets MaxIOPS field to given value.

### HasMaxIOPS

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasMaxIOPS() bool`

HasMaxIOPS returns a boolean if a field has been set.

### GetUuid

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetActive

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetReadOnly

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetRemovable

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetRemovable() bool`

GetRemovable returns the Removable field if non-nil, zero value otherwise.

### GetRemovableOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetRemovableOk() (*bool, bool)`

GetRemovableOk returns a tuple with the Removable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovable

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetRemovable(v bool)`

SetRemovable sets Removable field to given value.

### HasRemovable

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasRemovable() bool`

HasRemovable returns a boolean if a field has been set.

### GetPoolName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetZone

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetZone() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetZoneOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetZone(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetZoneId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetDatastore

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetDatastore() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetDatastoreOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetDatastore(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetDatastoreId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetDatastoreId() int64`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetDatastoreIdOk() (*int64, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetDatastoreId(v int64)`

SetDatastoreId sets DatastoreId field to given value.

### HasDatastoreId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasDatastoreId() bool`

HasDatastoreId returns a boolean if a field has been set.

### GetDatastoreOption

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetDatastoreOption() string`

GetDatastoreOption returns the DatastoreOption field if non-nil, zero value otherwise.

### GetDatastoreOptionOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetDatastoreOptionOk() (*string, bool)`

GetDatastoreOptionOk returns a tuple with the DatastoreOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreOption

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetDatastoreOption(v string)`

SetDatastoreOption sets DatastoreOption field to given value.

### HasDatastoreOption

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasDatastoreOption() bool`

HasDatastoreOption returns a boolean if a field has been set.

### GetStorageGroup

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetStorageGroup() string`

GetStorageGroup returns the StorageGroup field if non-nil, zero value otherwise.

### GetStorageGroupOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetStorageGroupOk() (*string, bool)`

GetStorageGroupOk returns a tuple with the StorageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageGroup

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetStorageGroup(v string)`

SetStorageGroup sets StorageGroup field to given value.

### HasStorageGroup

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasStorageGroup() bool`

HasStorageGroup returns a boolean if a field has been set.

### GetNamespace

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetStorageServer

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetStorageServer() string`

GetStorageServer returns the StorageServer field if non-nil, zero value otherwise.

### GetStorageServerOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetStorageServerOk() (*string, bool)`

GetStorageServerOk returns a tuple with the StorageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServer

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetStorageServer(v string)`

SetStorageServer sets StorageServer field to given value.

### HasStorageServer

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasStorageServer() bool`

HasStorageServer returns a boolean if a field has been set.

### GetSource

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUniqueId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetInternalId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetProvisionType

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetProvisionType() string`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetProvisionTypeOk() (*string, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetProvisionType(v string)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetCopyType

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetCopyType() string`

GetCopyType returns the CopyType field if non-nil, zero value otherwise.

### GetCopyTypeOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetCopyTypeOk() (*string, bool)`

GetCopyTypeOk returns a tuple with the CopyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyType

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetCopyType(v string)`

SetCopyType sets CopyType field to given value.

### HasCopyType

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasCopyType() bool`

HasCopyType returns a boolean if a field has been set.

### GetFiberWwn

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetFiberWwn() string`

GetFiberWwn returns the FiberWwn field if non-nil, zero value otherwise.

### GetFiberWwnOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetFiberWwnOk() (*string, bool)`

GetFiberWwnOk returns a tuple with the FiberWwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiberWwn

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetFiberWwn(v string)`

SetFiberWwn sets FiberWwn field to given value.

### HasFiberWwn

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasFiberWwn() bool`

HasFiberWwn returns a boolean if a field has been set.

### GetFileName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetClaimName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetClaimName() string`

GetClaimName returns the ClaimName field if non-nil, zero value otherwise.

### GetClaimNameOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetClaimNameOk() (*string, bool)`

GetClaimNameOk returns a tuple with the ClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetClaimName(v string)`

SetClaimName sets ClaimName field to given value.

### HasClaimName

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasClaimName() bool`

HasClaimName returns a boolean if a field has been set.

### GetSharePath

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetSharePath() string`

GetSharePath returns the SharePath field if non-nil, zero value otherwise.

### GetSharePathOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetSharePathOk() (*string, bool)`

GetSharePathOk returns a tuple with the SharePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePath

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetSharePath(v string)`

SetSharePath sets SharePath field to given value.

### HasSharePath

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasSharePath() bool`

HasSharePath returns a boolean if a field has been set.

### GetSourceId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceImage

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetSourceImage() string`

GetSourceImage returns the SourceImage field if non-nil, zero value otherwise.

### GetSourceImageOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetSourceImageOk() (*string, bool)`

GetSourceImageOk returns a tuple with the SourceImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceImage

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetSourceImage(v string)`

SetSourceImage sets SourceImage field to given value.

### HasSourceImage

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasSourceImage() bool`

HasSourceImage returns a boolean if a field has been set.

### GetImageType

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetImageTypeOk() (*string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetImageType(v string)`

SetImageType sets ImageType field to given value.

### HasImageType

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasImageType() bool`

HasImageType returns a boolean if a field has been set.

### GetOnline

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetRawData

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetRawData() string`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetRawDataOk() (*string, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetRawData(v string)`

SetRawData sets RawData field to given value.

### HasRawData

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasRawData() bool`

HasRawData returns a boolean if a field has been set.

### GetAccount

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetOwner() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) GetOwnerOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) SetOwner(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListClusterVolumes200ResponseAllOfVolumesInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


