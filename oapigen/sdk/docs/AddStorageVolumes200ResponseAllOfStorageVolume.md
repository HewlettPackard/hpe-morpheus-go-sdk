# AddStorageVolumes200ResponseAllOfStorageVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Controller** | Pointer to [**AddStorageVolumes200ResponseAllOfStorageVolumeController**](AddStorageVolumes200ResponseAllOfStorageVolumeController.md) |  | [optional] 
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
**Type** | Pointer to [**AddStorageVolumes200ResponseAllOfStorageVolumeType**](AddStorageVolumes200ResponseAllOfStorageVolumeType.md) |  | [optional] 
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
**Zone** | Pointer to [**AddStorageVolumes200ResponseAllOfStorageVolumeZone**](AddStorageVolumes200ResponseAllOfStorageVolumeZone.md) |  | [optional] 
**ZoneId** | Pointer to **int64** |  | [optional] 
**Datastore** | Pointer to [**AddStorageVolumes200ResponseAllOfStorageVolumeDatastore**](AddStorageVolumes200ResponseAllOfStorageVolumeDatastore.md) |  | [optional] 
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
**CreateForMultiAttach** | Pointer to **bool** |  | [optional] 
**IsMultiAttach** | Pointer to **bool** |  | [optional] 
**StorageProfile** | Pointer to **NullableString** | Storage Profile Code for the volume storage profile assignment. eg. &#x60;\&quot;kvm-cache-none\&quot;&#x60; or &#x60;\&quot;kvm-cache-directsync\&quot;&#x60;. Use &#x60;/api/provision-types?code&#x3D;kvm&#x60; to see the available &#x60;storageProfiles&#x60; for HVM and KVM. | [optional] 
**Account** | Pointer to [**AddStorageVolumes200ResponseAllOfStorageVolumeAccount**](AddStorageVolumes200ResponseAllOfStorageVolumeAccount.md) |  | [optional] 
**Owner** | Pointer to [**AddStorageVolumes200ResponseAllOfStorageVolumeOwner**](AddStorageVolumes200ResponseAllOfStorageVolumeOwner.md) |  | [optional] 

## Methods

### NewAddStorageVolumes200ResponseAllOfStorageVolume

`func NewAddStorageVolumes200ResponseAllOfStorageVolume() *AddStorageVolumes200ResponseAllOfStorageVolume`

NewAddStorageVolumes200ResponseAllOfStorageVolume instantiates a new AddStorageVolumes200ResponseAllOfStorageVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetController

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetController() AddStorageVolumes200ResponseAllOfStorageVolumeController`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetControllerOk() (*AddStorageVolumes200ResponseAllOfStorageVolumeController, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetController(v AddStorageVolumes200ResponseAllOfStorageVolumeController)`

SetController sets Controller field to given value.

### HasController

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasController() bool`

HasController returns a boolean if a field has been set.

### GetControllerId

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetControllerId() int64`

GetControllerId returns the ControllerId field if non-nil, zero value otherwise.

### GetControllerIdOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetControllerIdOk() (*int64, bool)`

GetControllerIdOk returns a tuple with the ControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerId

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetControllerId(v int64)`

SetControllerId sets ControllerId field to given value.

### HasControllerId

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasControllerId() bool`

HasControllerId returns a boolean if a field has been set.

### SetControllerIdNil

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetControllerIdNil(b bool)`

 SetControllerIdNil sets the value for ControllerId to be an explicit nil

### UnsetControllerId
`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) UnsetControllerId()`

UnsetControllerId ensures that no value is present for ControllerId, not even an explicit nil
### GetControllerMountPoint

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetControllerMountPoint() string`

GetControllerMountPoint returns the ControllerMountPoint field if non-nil, zero value otherwise.

### GetControllerMountPointOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetControllerMountPointOk() (*string, bool)`

GetControllerMountPointOk returns a tuple with the ControllerMountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerMountPoint

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetControllerMountPoint(v string)`

SetControllerMountPoint sets ControllerMountPoint field to given value.

### HasControllerMountPoint

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasControllerMountPoint() bool`

HasControllerMountPoint returns a boolean if a field has been set.

### SetControllerMountPointNil

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetControllerMountPointNil(b bool)`

 SetControllerMountPointNil sets the value for ControllerMountPoint to be an explicit nil

### UnsetControllerMountPoint
`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) UnsetControllerMountPoint()`

UnsetControllerMountPoint ensures that no value is present for ControllerMountPoint, not even an explicit nil
### GetResizeable

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetResizeable() bool`

GetResizeable returns the Resizeable field if non-nil, zero value otherwise.

### GetResizeableOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetResizeableOk() (*bool, bool)`

GetResizeableOk returns a tuple with the Resizeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeable

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetResizeable(v bool)`

SetResizeable sets Resizeable field to given value.

### HasResizeable

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasResizeable() bool`

HasResizeable returns a boolean if a field has been set.

### SetResizeableNil

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetResizeableNil(b bool)`

 SetResizeableNil sets the value for Resizeable to be an explicit nil

### UnsetResizeable
`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) UnsetResizeable()`

UnsetResizeable ensures that no value is present for Resizeable, not even an explicit nil
### GetRootVolume

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetRootVolume() bool`

GetRootVolume returns the RootVolume field if non-nil, zero value otherwise.

### GetRootVolumeOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetRootVolumeOk() (*bool, bool)`

GetRootVolumeOk returns a tuple with the RootVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolume

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetRootVolume(v bool)`

SetRootVolume sets RootVolume field to given value.

### HasRootVolume

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasRootVolume() bool`

HasRootVolume returns a boolean if a field has been set.

### GetUnitNumber

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetUnitNumber() string`

GetUnitNumber returns the UnitNumber field if non-nil, zero value otherwise.

### GetUnitNumberOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetUnitNumberOk() (*string, bool)`

GetUnitNumberOk returns a tuple with the UnitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitNumber

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetUnitNumber(v string)`

SetUnitNumber sets UnitNumber field to given value.

### HasUnitNumber

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasUnitNumber() bool`

HasUnitNumber returns a boolean if a field has been set.

### SetUnitNumberNil

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetUnitNumberNil(b bool)`

 SetUnitNumberNil sets the value for UnitNumber to be an explicit nil

### UnsetUnitNumber
`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) UnsetUnitNumber()`

UnsetUnitNumber ensures that no value is present for UnitNumber, not even an explicit nil
### GetDeviceName

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceDisplayName

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetDeviceDisplayName() string`

GetDeviceDisplayName returns the DeviceDisplayName field if non-nil, zero value otherwise.

### GetDeviceDisplayNameOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetDeviceDisplayNameOk() (*string, bool)`

GetDeviceDisplayNameOk returns a tuple with the DeviceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDisplayName

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetDeviceDisplayName(v string)`

SetDeviceDisplayName sets DeviceDisplayName field to given value.

### HasDeviceDisplayName

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasDeviceDisplayName() bool`

HasDeviceDisplayName returns a boolean if a field has been set.

### GetVolumeName

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.

### GetVolumePath

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetVolumePath() string`

GetVolumePath returns the VolumePath field if non-nil, zero value otherwise.

### GetVolumePathOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetVolumePathOk() (*string, bool)`

GetVolumePathOk returns a tuple with the VolumePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumePath

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetVolumePath(v string)`

SetVolumePath sets VolumePath field to given value.

### HasVolumePath

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasVolumePath() bool`

HasVolumePath returns a boolean if a field has been set.

### GetVolumeType

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetRefType

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetDiskMode

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetDiskMode() string`

GetDiskMode returns the DiskMode field if non-nil, zero value otherwise.

### GetDiskModeOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetDiskModeOk() (*string, bool)`

GetDiskModeOk returns a tuple with the DiskMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskMode

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetDiskMode(v string)`

SetDiskMode sets DiskMode field to given value.

### HasDiskMode

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasDiskMode() bool`

HasDiskMode returns a boolean if a field has been set.

### GetDiskType

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetDiskType() string`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetDiskTypeOk() (*string, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetDiskType(v string)`

SetDiskType sets DiskType field to given value.

### HasDiskType

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasDiskType() bool`

HasDiskType returns a boolean if a field has been set.

### GetType

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetType() AddStorageVolumes200ResponseAllOfStorageVolumeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetTypeOk() (*AddStorageVolumes200ResponseAllOfStorageVolumeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetType(v AddStorageVolumes200ResponseAllOfStorageVolumeType)`

SetType sets Type field to given value.

### HasType

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypeId

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetCategory

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetStatus

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetConfigurableIOPS

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetConfigurableIOPS() bool`

GetConfigurableIOPS returns the ConfigurableIOPS field if non-nil, zero value otherwise.

### GetConfigurableIOPSOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetConfigurableIOPSOk() (*bool, bool)`

GetConfigurableIOPSOk returns a tuple with the ConfigurableIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurableIOPS

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetConfigurableIOPS(v bool)`

SetConfigurableIOPS sets ConfigurableIOPS field to given value.

### HasConfigurableIOPS

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasConfigurableIOPS() bool`

HasConfigurableIOPS returns a boolean if a field has been set.

### GetMaxStorage

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetUsedStorage

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetUsedStorage() int64`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetUsedStorageOk() (*int64, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetUsedStorage(v int64)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetMaxIOPS

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetMaxIOPS() string`

GetMaxIOPS returns the MaxIOPS field if non-nil, zero value otherwise.

### GetMaxIOPSOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetMaxIOPSOk() (*string, bool)`

GetMaxIOPSOk returns a tuple with the MaxIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIOPS

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetMaxIOPS(v string)`

SetMaxIOPS sets MaxIOPS field to given value.

### HasMaxIOPS

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasMaxIOPS() bool`

HasMaxIOPS returns a boolean if a field has been set.

### SetMaxIOPSNil

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetMaxIOPSNil(b bool)`

 SetMaxIOPSNil sets the value for MaxIOPS to be an explicit nil

### UnsetMaxIOPS
`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) UnsetMaxIOPS()`

UnsetMaxIOPS ensures that no value is present for MaxIOPS, not even an explicit nil
### GetUuid

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetActive

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetReadOnly

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetRemovable

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetRemovable() bool`

GetRemovable returns the Removable field if non-nil, zero value otherwise.

### GetRemovableOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetRemovableOk() (*bool, bool)`

GetRemovableOk returns a tuple with the Removable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovable

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetRemovable(v bool)`

SetRemovable sets Removable field to given value.

### HasRemovable

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasRemovable() bool`

HasRemovable returns a boolean if a field has been set.

### GetPoolName

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetZone

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetZone() AddStorageVolumes200ResponseAllOfStorageVolumeZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetZoneOk() (*AddStorageVolumes200ResponseAllOfStorageVolumeZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetZone(v AddStorageVolumes200ResponseAllOfStorageVolumeZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetZoneId

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetDatastore

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetDatastore() AddStorageVolumes200ResponseAllOfStorageVolumeDatastore`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetDatastoreOk() (*AddStorageVolumes200ResponseAllOfStorageVolumeDatastore, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetDatastore(v AddStorageVolumes200ResponseAllOfStorageVolumeDatastore)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetDatastoreId

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetDatastoreId() int64`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetDatastoreIdOk() (*int64, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetDatastoreId(v int64)`

SetDatastoreId sets DatastoreId field to given value.

### HasDatastoreId

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasDatastoreId() bool`

HasDatastoreId returns a boolean if a field has been set.

### SetDatastoreIdNil

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetDatastoreIdNil(b bool)`

 SetDatastoreIdNil sets the value for DatastoreId to be an explicit nil

### UnsetDatastoreId
`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) UnsetDatastoreId()`

UnsetDatastoreId ensures that no value is present for DatastoreId, not even an explicit nil
### GetDatastoreOption

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetDatastoreOption() string`

GetDatastoreOption returns the DatastoreOption field if non-nil, zero value otherwise.

### GetDatastoreOptionOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetDatastoreOptionOk() (*string, bool)`

GetDatastoreOptionOk returns a tuple with the DatastoreOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreOption

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetDatastoreOption(v string)`

SetDatastoreOption sets DatastoreOption field to given value.

### HasDatastoreOption

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasDatastoreOption() bool`

HasDatastoreOption returns a boolean if a field has been set.

### GetStorageGroup

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetStorageGroup() string`

GetStorageGroup returns the StorageGroup field if non-nil, zero value otherwise.

### GetStorageGroupOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetStorageGroupOk() (*string, bool)`

GetStorageGroupOk returns a tuple with the StorageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageGroup

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetStorageGroup(v string)`

SetStorageGroup sets StorageGroup field to given value.

### HasStorageGroup

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasStorageGroup() bool`

HasStorageGroup returns a boolean if a field has been set.

### SetStorageGroupNil

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetStorageGroupNil(b bool)`

 SetStorageGroupNil sets the value for StorageGroup to be an explicit nil

### UnsetStorageGroup
`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) UnsetStorageGroup()`

UnsetStorageGroup ensures that no value is present for StorageGroup, not even an explicit nil
### GetNamespace

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetStorageServer

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetStorageServer() map[string]interface{}`

GetStorageServer returns the StorageServer field if non-nil, zero value otherwise.

### GetStorageServerOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetStorageServerOk() (*map[string]interface{}, bool)`

GetStorageServerOk returns a tuple with the StorageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServer

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetStorageServer(v map[string]interface{})`

SetStorageServer sets StorageServer field to given value.

### HasStorageServer

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasStorageServer() bool`

HasStorageServer returns a boolean if a field has been set.

### GetSource

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUniqueId

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### SetUniqueIdNil

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetUniqueIdNil(b bool)`

 SetUniqueIdNil sets the value for UniqueId to be an explicit nil

### UnsetUniqueId
`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) UnsetUniqueId()`

UnsetUniqueId ensures that no value is present for UniqueId, not even an explicit nil
### GetInternalId

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetProvisionType

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetProvisionType() string`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetProvisionTypeOk() (*string, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetProvisionType(v string)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### SetProvisionTypeNil

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetProvisionTypeNil(b bool)`

 SetProvisionTypeNil sets the value for ProvisionType to be an explicit nil

### UnsetProvisionType
`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) UnsetProvisionType()`

UnsetProvisionType ensures that no value is present for ProvisionType, not even an explicit nil
### GetCopyType

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetCopyType() string`

GetCopyType returns the CopyType field if non-nil, zero value otherwise.

### GetCopyTypeOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetCopyTypeOk() (*string, bool)`

GetCopyTypeOk returns a tuple with the CopyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyType

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetCopyType(v string)`

SetCopyType sets CopyType field to given value.

### HasCopyType

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasCopyType() bool`

HasCopyType returns a boolean if a field has been set.

### SetCopyTypeNil

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetCopyTypeNil(b bool)`

 SetCopyTypeNil sets the value for CopyType to be an explicit nil

### UnsetCopyType
`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) UnsetCopyType()`

UnsetCopyType ensures that no value is present for CopyType, not even an explicit nil
### GetFiberWwn

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetFiberWwn() string`

GetFiberWwn returns the FiberWwn field if non-nil, zero value otherwise.

### GetFiberWwnOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetFiberWwnOk() (*string, bool)`

GetFiberWwnOk returns a tuple with the FiberWwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiberWwn

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetFiberWwn(v string)`

SetFiberWwn sets FiberWwn field to given value.

### HasFiberWwn

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasFiberWwn() bool`

HasFiberWwn returns a boolean if a field has been set.

### SetFiberWwnNil

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetFiberWwnNil(b bool)`

 SetFiberWwnNil sets the value for FiberWwn to be an explicit nil

### UnsetFiberWwn
`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) UnsetFiberWwn()`

UnsetFiberWwn ensures that no value is present for FiberWwn, not even an explicit nil
### GetFileName

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetClaimName

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetClaimName() string`

GetClaimName returns the ClaimName field if non-nil, zero value otherwise.

### GetClaimNameOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetClaimNameOk() (*string, bool)`

GetClaimNameOk returns a tuple with the ClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimName

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetClaimName(v string)`

SetClaimName sets ClaimName field to given value.

### HasClaimName

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasClaimName() bool`

HasClaimName returns a boolean if a field has been set.

### SetClaimNameNil

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetClaimNameNil(b bool)`

 SetClaimNameNil sets the value for ClaimName to be an explicit nil

### UnsetClaimName
`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) UnsetClaimName()`

UnsetClaimName ensures that no value is present for ClaimName, not even an explicit nil
### GetSharePath

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetSharePath() string`

GetSharePath returns the SharePath field if non-nil, zero value otherwise.

### GetSharePathOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetSharePathOk() (*string, bool)`

GetSharePathOk returns a tuple with the SharePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePath

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetSharePath(v string)`

SetSharePath sets SharePath field to given value.

### HasSharePath

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasSharePath() bool`

HasSharePath returns a boolean if a field has been set.

### SetSharePathNil

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetSharePathNil(b bool)`

 SetSharePathNil sets the value for SharePath to be an explicit nil

### UnsetSharePath
`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) UnsetSharePath()`

UnsetSharePath ensures that no value is present for SharePath, not even an explicit nil
### GetSourceId

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceImage

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetSourceImage() string`

GetSourceImage returns the SourceImage field if non-nil, zero value otherwise.

### GetSourceImageOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetSourceImageOk() (*string, bool)`

GetSourceImageOk returns a tuple with the SourceImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceImage

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetSourceImage(v string)`

SetSourceImage sets SourceImage field to given value.

### HasSourceImage

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasSourceImage() bool`

HasSourceImage returns a boolean if a field has been set.

### GetImageType

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetImageTypeOk() (*string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetImageType(v string)`

SetImageType sets ImageType field to given value.

### HasImageType

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasImageType() bool`

HasImageType returns a boolean if a field has been set.

### GetOnline

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetRawData

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetRawData() string`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetRawDataOk() (*string, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetRawData(v string)`

SetRawData sets RawData field to given value.

### HasRawData

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasRawData() bool`

HasRawData returns a boolean if a field has been set.

### GetCreateForMultiAttach

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetCreateForMultiAttach() bool`

GetCreateForMultiAttach returns the CreateForMultiAttach field if non-nil, zero value otherwise.

### GetCreateForMultiAttachOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetCreateForMultiAttachOk() (*bool, bool)`

GetCreateForMultiAttachOk returns a tuple with the CreateForMultiAttach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateForMultiAttach

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetCreateForMultiAttach(v bool)`

SetCreateForMultiAttach sets CreateForMultiAttach field to given value.

### HasCreateForMultiAttach

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasCreateForMultiAttach() bool`

HasCreateForMultiAttach returns a boolean if a field has been set.

### GetIsMultiAttach

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetIsMultiAttach() bool`

GetIsMultiAttach returns the IsMultiAttach field if non-nil, zero value otherwise.

### GetIsMultiAttachOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetIsMultiAttachOk() (*bool, bool)`

GetIsMultiAttachOk returns a tuple with the IsMultiAttach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiAttach

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetIsMultiAttach(v bool)`

SetIsMultiAttach sets IsMultiAttach field to given value.

### HasIsMultiAttach

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasIsMultiAttach() bool`

HasIsMultiAttach returns a boolean if a field has been set.

### GetStorageProfile

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetStorageProfile() string`

GetStorageProfile returns the StorageProfile field if non-nil, zero value otherwise.

### GetStorageProfileOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetStorageProfileOk() (*string, bool)`

GetStorageProfileOk returns a tuple with the StorageProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProfile

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetStorageProfile(v string)`

SetStorageProfile sets StorageProfile field to given value.

### HasStorageProfile

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasStorageProfile() bool`

HasStorageProfile returns a boolean if a field has been set.

### SetStorageProfileNil

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetStorageProfileNil(b bool)`

 SetStorageProfileNil sets the value for StorageProfile to be an explicit nil

### UnsetStorageProfile
`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) UnsetStorageProfile()`

UnsetStorageProfile ensures that no value is present for StorageProfile, not even an explicit nil
### GetAccount

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetAccount() AddStorageVolumes200ResponseAllOfStorageVolumeAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetAccountOk() (*AddStorageVolumes200ResponseAllOfStorageVolumeAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetAccount(v AddStorageVolumes200ResponseAllOfStorageVolumeAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetOwner() AddStorageVolumes200ResponseAllOfStorageVolumeOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) GetOwnerOk() (*AddStorageVolumes200ResponseAllOfStorageVolumeOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) SetOwner(v AddStorageVolumes200ResponseAllOfStorageVolumeOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AddStorageVolumes200ResponseAllOfStorageVolume) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


