# ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ControllerId** | Pointer to **int64** |  | [optional] 
**ControllerMountPoint** | Pointer to **string** |  | [optional] 
**Resizeable** | Pointer to **bool** |  | [optional] 
**RootVolume** | Pointer to **bool** |  | [optional] 
**UnitNumber** | Pointer to **string** |  | [optional] 
**DeviceName** | Pointer to **string** |  | [optional] 
**DeviceDisplayName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**TypeId** | Pointer to **int64** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**ConfigurableIOPS** | Pointer to **bool** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 
**DisplayOrder** | Pointer to **int64** |  | [optional] 
**MaxIOPS** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Zone** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**ZoneId** | Pointer to **int64** |  | [optional] 
**Datastore** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**DatastoreId** | Pointer to **int64** |  | [optional] 
**StorageGroup** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**StorageServer** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 

## Methods

### NewListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner

`func NewListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner() *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner`

NewListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner instantiates a new ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerWithDefaults

`func NewListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerWithDefaults() *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner`

NewListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerWithDefaults instantiates a new ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetControllerId

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetControllerId() int64`

GetControllerId returns the ControllerId field if non-nil, zero value otherwise.

### GetControllerIdOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetControllerIdOk() (*int64, bool)`

GetControllerIdOk returns a tuple with the ControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerId

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetControllerId(v int64)`

SetControllerId sets ControllerId field to given value.

### HasControllerId

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasControllerId() bool`

HasControllerId returns a boolean if a field has been set.

### GetControllerMountPoint

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetControllerMountPoint() string`

GetControllerMountPoint returns the ControllerMountPoint field if non-nil, zero value otherwise.

### GetControllerMountPointOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetControllerMountPointOk() (*string, bool)`

GetControllerMountPointOk returns a tuple with the ControllerMountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerMountPoint

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetControllerMountPoint(v string)`

SetControllerMountPoint sets ControllerMountPoint field to given value.

### HasControllerMountPoint

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasControllerMountPoint() bool`

HasControllerMountPoint returns a boolean if a field has been set.

### GetResizeable

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetResizeable() bool`

GetResizeable returns the Resizeable field if non-nil, zero value otherwise.

### GetResizeableOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetResizeableOk() (*bool, bool)`

GetResizeableOk returns a tuple with the Resizeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeable

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetResizeable(v bool)`

SetResizeable sets Resizeable field to given value.

### HasResizeable

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasResizeable() bool`

HasResizeable returns a boolean if a field has been set.

### GetRootVolume

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetRootVolume() bool`

GetRootVolume returns the RootVolume field if non-nil, zero value otherwise.

### GetRootVolumeOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetRootVolumeOk() (*bool, bool)`

GetRootVolumeOk returns a tuple with the RootVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolume

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetRootVolume(v bool)`

SetRootVolume sets RootVolume field to given value.

### HasRootVolume

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasRootVolume() bool`

HasRootVolume returns a boolean if a field has been set.

### GetUnitNumber

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetUnitNumber() string`

GetUnitNumber returns the UnitNumber field if non-nil, zero value otherwise.

### GetUnitNumberOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetUnitNumberOk() (*string, bool)`

GetUnitNumberOk returns a tuple with the UnitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitNumber

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetUnitNumber(v string)`

SetUnitNumber sets UnitNumber field to given value.

### HasUnitNumber

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasUnitNumber() bool`

HasUnitNumber returns a boolean if a field has been set.

### GetDeviceName

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceDisplayName

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetDeviceDisplayName() string`

GetDeviceDisplayName returns the DeviceDisplayName field if non-nil, zero value otherwise.

### GetDeviceDisplayNameOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetDeviceDisplayNameOk() (*string, bool)`

GetDeviceDisplayNameOk returns a tuple with the DeviceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDisplayName

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetDeviceDisplayName(v string)`

SetDeviceDisplayName sets DeviceDisplayName field to given value.

### HasDeviceDisplayName

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasDeviceDisplayName() bool`

HasDeviceDisplayName returns a boolean if a field has been set.

### GetType

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetType sets Type field to given value.

### HasType

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypeId

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetCategory

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetStatus

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetConfigurableIOPS

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetConfigurableIOPS() bool`

GetConfigurableIOPS returns the ConfigurableIOPS field if non-nil, zero value otherwise.

### GetConfigurableIOPSOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetConfigurableIOPSOk() (*bool, bool)`

GetConfigurableIOPSOk returns a tuple with the ConfigurableIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurableIOPS

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetConfigurableIOPS(v bool)`

SetConfigurableIOPS sets ConfigurableIOPS field to given value.

### HasConfigurableIOPS

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasConfigurableIOPS() bool`

HasConfigurableIOPS returns a boolean if a field has been set.

### GetMaxStorage

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetMaxIOPS

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetMaxIOPS() string`

GetMaxIOPS returns the MaxIOPS field if non-nil, zero value otherwise.

### GetMaxIOPSOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetMaxIOPSOk() (*string, bool)`

GetMaxIOPSOk returns a tuple with the MaxIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIOPS

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetMaxIOPS(v string)`

SetMaxIOPS sets MaxIOPS field to given value.

### HasMaxIOPS

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasMaxIOPS() bool`

HasMaxIOPS returns a boolean if a field has been set.

### GetUuid

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetActive

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetZone

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetZone() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetZoneOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetZone(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetZoneId

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetDatastore

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetDatastore() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetDatastoreOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetDatastore(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetDatastoreId

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetDatastoreId() int64`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetDatastoreIdOk() (*int64, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetDatastoreId(v int64)`

SetDatastoreId sets DatastoreId field to given value.

### HasDatastoreId

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasDatastoreId() bool`

HasDatastoreId returns a boolean if a field has been set.

### GetStorageGroup

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetStorageGroup() string`

GetStorageGroup returns the StorageGroup field if non-nil, zero value otherwise.

### GetStorageGroupOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetStorageGroupOk() (*string, bool)`

GetStorageGroupOk returns a tuple with the StorageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageGroup

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetStorageGroup(v string)`

SetStorageGroup sets StorageGroup field to given value.

### HasStorageGroup

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasStorageGroup() bool`

HasStorageGroup returns a boolean if a field has been set.

### GetNamespace

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetStorageServer

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetStorageServer() string`

GetStorageServer returns the StorageServer field if non-nil, zero value otherwise.

### GetStorageServerOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetStorageServerOk() (*string, bool)`

GetStorageServerOk returns a tuple with the StorageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServer

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetStorageServer(v string)`

SetStorageServer sets StorageServer field to given value.

### HasStorageServer

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasStorageServer() bool`

HasStorageServer returns a boolean if a field has been set.

### GetSource

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetOwner

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetOwner() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetOwnerOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetOwner(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


