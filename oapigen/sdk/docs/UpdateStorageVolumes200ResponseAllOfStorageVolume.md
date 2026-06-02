# UpdateStorageVolumes200ResponseAllOfStorageVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Controller** | Pointer to [**UpdateStorageVolumes200ResponseAllOfStorageVolumeController**](UpdateStorageVolumes200ResponseAllOfStorageVolumeController.md) |  | [optional] 
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
**Type** | Pointer to [**UpdateStorageVolumes200ResponseAllOfStorageVolumeType**](UpdateStorageVolumes200ResponseAllOfStorageVolumeType.md) |  | [optional] 
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
**Zone** | Pointer to [**UpdateStorageVolumes200ResponseAllOfStorageVolumeZone**](UpdateStorageVolumes200ResponseAllOfStorageVolumeZone.md) |  | [optional] 
**ZoneId** | Pointer to **int64** |  | [optional] 
**Datastore** | Pointer to [**UpdateStorageVolumes200ResponseAllOfStorageVolumeDatastore**](UpdateStorageVolumes200ResponseAllOfStorageVolumeDatastore.md) |  | [optional] 
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
**Account** | Pointer to [**UpdateStorageVolumes200ResponseAllOfStorageVolumeAccount**](UpdateStorageVolumes200ResponseAllOfStorageVolumeAccount.md) |  | [optional] 
**Owner** | Pointer to [**UpdateStorageVolumes200ResponseAllOfStorageVolumeOwner**](UpdateStorageVolumes200ResponseAllOfStorageVolumeOwner.md) |  | [optional] 

## Methods

### NewUpdateStorageVolumes200ResponseAllOfStorageVolume

`func NewUpdateStorageVolumes200ResponseAllOfStorageVolume() *UpdateStorageVolumes200ResponseAllOfStorageVolume`

NewUpdateStorageVolumes200ResponseAllOfStorageVolume instantiates a new UpdateStorageVolumes200ResponseAllOfStorageVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetController

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetController() UpdateStorageVolumes200ResponseAllOfStorageVolumeController`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetControllerOk() (*UpdateStorageVolumes200ResponseAllOfStorageVolumeController, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetController(v UpdateStorageVolumes200ResponseAllOfStorageVolumeController)`

SetController sets Controller field to given value.

### HasController

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasController() bool`

HasController returns a boolean if a field has been set.

### GetControllerId

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetControllerId() int64`

GetControllerId returns the ControllerId field if non-nil, zero value otherwise.

### GetControllerIdOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetControllerIdOk() (*int64, bool)`

GetControllerIdOk returns a tuple with the ControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerId

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetControllerId(v int64)`

SetControllerId sets ControllerId field to given value.

### HasControllerId

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasControllerId() bool`

HasControllerId returns a boolean if a field has been set.

### SetControllerIdNil

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetControllerIdNil(b bool)`

 SetControllerIdNil sets the value for ControllerId to be an explicit nil

### UnsetControllerId
`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) UnsetControllerId()`

UnsetControllerId ensures that no value is present for ControllerId, not even an explicit nil
### GetControllerMountPoint

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetControllerMountPoint() string`

GetControllerMountPoint returns the ControllerMountPoint field if non-nil, zero value otherwise.

### GetControllerMountPointOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetControllerMountPointOk() (*string, bool)`

GetControllerMountPointOk returns a tuple with the ControllerMountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerMountPoint

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetControllerMountPoint(v string)`

SetControllerMountPoint sets ControllerMountPoint field to given value.

### HasControllerMountPoint

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasControllerMountPoint() bool`

HasControllerMountPoint returns a boolean if a field has been set.

### SetControllerMountPointNil

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetControllerMountPointNil(b bool)`

 SetControllerMountPointNil sets the value for ControllerMountPoint to be an explicit nil

### UnsetControllerMountPoint
`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) UnsetControllerMountPoint()`

UnsetControllerMountPoint ensures that no value is present for ControllerMountPoint, not even an explicit nil
### GetResizeable

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetResizeable() bool`

GetResizeable returns the Resizeable field if non-nil, zero value otherwise.

### GetResizeableOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetResizeableOk() (*bool, bool)`

GetResizeableOk returns a tuple with the Resizeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeable

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetResizeable(v bool)`

SetResizeable sets Resizeable field to given value.

### HasResizeable

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasResizeable() bool`

HasResizeable returns a boolean if a field has been set.

### SetResizeableNil

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetResizeableNil(b bool)`

 SetResizeableNil sets the value for Resizeable to be an explicit nil

### UnsetResizeable
`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) UnsetResizeable()`

UnsetResizeable ensures that no value is present for Resizeable, not even an explicit nil
### GetRootVolume

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetRootVolume() bool`

GetRootVolume returns the RootVolume field if non-nil, zero value otherwise.

### GetRootVolumeOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetRootVolumeOk() (*bool, bool)`

GetRootVolumeOk returns a tuple with the RootVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolume

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetRootVolume(v bool)`

SetRootVolume sets RootVolume field to given value.

### HasRootVolume

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasRootVolume() bool`

HasRootVolume returns a boolean if a field has been set.

### GetUnitNumber

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetUnitNumber() string`

GetUnitNumber returns the UnitNumber field if non-nil, zero value otherwise.

### GetUnitNumberOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetUnitNumberOk() (*string, bool)`

GetUnitNumberOk returns a tuple with the UnitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitNumber

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetUnitNumber(v string)`

SetUnitNumber sets UnitNumber field to given value.

### HasUnitNumber

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasUnitNumber() bool`

HasUnitNumber returns a boolean if a field has been set.

### SetUnitNumberNil

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetUnitNumberNil(b bool)`

 SetUnitNumberNil sets the value for UnitNumber to be an explicit nil

### UnsetUnitNumber
`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) UnsetUnitNumber()`

UnsetUnitNumber ensures that no value is present for UnitNumber, not even an explicit nil
### GetDeviceName

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceDisplayName

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetDeviceDisplayName() string`

GetDeviceDisplayName returns the DeviceDisplayName field if non-nil, zero value otherwise.

### GetDeviceDisplayNameOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetDeviceDisplayNameOk() (*string, bool)`

GetDeviceDisplayNameOk returns a tuple with the DeviceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDisplayName

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetDeviceDisplayName(v string)`

SetDeviceDisplayName sets DeviceDisplayName field to given value.

### HasDeviceDisplayName

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasDeviceDisplayName() bool`

HasDeviceDisplayName returns a boolean if a field has been set.

### GetVolumeName

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.

### GetVolumePath

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetVolumePath() string`

GetVolumePath returns the VolumePath field if non-nil, zero value otherwise.

### GetVolumePathOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetVolumePathOk() (*string, bool)`

GetVolumePathOk returns a tuple with the VolumePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumePath

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetVolumePath(v string)`

SetVolumePath sets VolumePath field to given value.

### HasVolumePath

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasVolumePath() bool`

HasVolumePath returns a boolean if a field has been set.

### GetVolumeType

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetRefType

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetDiskMode

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetDiskMode() string`

GetDiskMode returns the DiskMode field if non-nil, zero value otherwise.

### GetDiskModeOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetDiskModeOk() (*string, bool)`

GetDiskModeOk returns a tuple with the DiskMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskMode

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetDiskMode(v string)`

SetDiskMode sets DiskMode field to given value.

### HasDiskMode

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasDiskMode() bool`

HasDiskMode returns a boolean if a field has been set.

### GetDiskType

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetDiskType() string`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetDiskTypeOk() (*string, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetDiskType(v string)`

SetDiskType sets DiskType field to given value.

### HasDiskType

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasDiskType() bool`

HasDiskType returns a boolean if a field has been set.

### GetType

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetType() UpdateStorageVolumes200ResponseAllOfStorageVolumeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetTypeOk() (*UpdateStorageVolumes200ResponseAllOfStorageVolumeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetType(v UpdateStorageVolumes200ResponseAllOfStorageVolumeType)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypeId

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetCategory

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetConfigurableIOPS

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetConfigurableIOPS() bool`

GetConfigurableIOPS returns the ConfigurableIOPS field if non-nil, zero value otherwise.

### GetConfigurableIOPSOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetConfigurableIOPSOk() (*bool, bool)`

GetConfigurableIOPSOk returns a tuple with the ConfigurableIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurableIOPS

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetConfigurableIOPS(v bool)`

SetConfigurableIOPS sets ConfigurableIOPS field to given value.

### HasConfigurableIOPS

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasConfigurableIOPS() bool`

HasConfigurableIOPS returns a boolean if a field has been set.

### GetMaxStorage

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetUsedStorage

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetUsedStorage() int64`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetUsedStorageOk() (*int64, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetUsedStorage(v int64)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetMaxIOPS

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetMaxIOPS() string`

GetMaxIOPS returns the MaxIOPS field if non-nil, zero value otherwise.

### GetMaxIOPSOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetMaxIOPSOk() (*string, bool)`

GetMaxIOPSOk returns a tuple with the MaxIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIOPS

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetMaxIOPS(v string)`

SetMaxIOPS sets MaxIOPS field to given value.

### HasMaxIOPS

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasMaxIOPS() bool`

HasMaxIOPS returns a boolean if a field has been set.

### SetMaxIOPSNil

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetMaxIOPSNil(b bool)`

 SetMaxIOPSNil sets the value for MaxIOPS to be an explicit nil

### UnsetMaxIOPS
`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) UnsetMaxIOPS()`

UnsetMaxIOPS ensures that no value is present for MaxIOPS, not even an explicit nil
### GetUuid

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetActive

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetReadOnly

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetRemovable

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetRemovable() bool`

GetRemovable returns the Removable field if non-nil, zero value otherwise.

### GetRemovableOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetRemovableOk() (*bool, bool)`

GetRemovableOk returns a tuple with the Removable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovable

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetRemovable(v bool)`

SetRemovable sets Removable field to given value.

### HasRemovable

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasRemovable() bool`

HasRemovable returns a boolean if a field has been set.

### GetPoolName

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetZone

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetZone() UpdateStorageVolumes200ResponseAllOfStorageVolumeZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetZoneOk() (*UpdateStorageVolumes200ResponseAllOfStorageVolumeZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetZone(v UpdateStorageVolumes200ResponseAllOfStorageVolumeZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetZoneId

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetDatastore

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetDatastore() UpdateStorageVolumes200ResponseAllOfStorageVolumeDatastore`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetDatastoreOk() (*UpdateStorageVolumes200ResponseAllOfStorageVolumeDatastore, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetDatastore(v UpdateStorageVolumes200ResponseAllOfStorageVolumeDatastore)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetDatastoreId

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetDatastoreId() int64`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetDatastoreIdOk() (*int64, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetDatastoreId(v int64)`

SetDatastoreId sets DatastoreId field to given value.

### HasDatastoreId

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasDatastoreId() bool`

HasDatastoreId returns a boolean if a field has been set.

### SetDatastoreIdNil

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetDatastoreIdNil(b bool)`

 SetDatastoreIdNil sets the value for DatastoreId to be an explicit nil

### UnsetDatastoreId
`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) UnsetDatastoreId()`

UnsetDatastoreId ensures that no value is present for DatastoreId, not even an explicit nil
### GetDatastoreOption

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetDatastoreOption() string`

GetDatastoreOption returns the DatastoreOption field if non-nil, zero value otherwise.

### GetDatastoreOptionOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetDatastoreOptionOk() (*string, bool)`

GetDatastoreOptionOk returns a tuple with the DatastoreOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreOption

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetDatastoreOption(v string)`

SetDatastoreOption sets DatastoreOption field to given value.

### HasDatastoreOption

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasDatastoreOption() bool`

HasDatastoreOption returns a boolean if a field has been set.

### GetStorageGroup

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetStorageGroup() string`

GetStorageGroup returns the StorageGroup field if non-nil, zero value otherwise.

### GetStorageGroupOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetStorageGroupOk() (*string, bool)`

GetStorageGroupOk returns a tuple with the StorageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageGroup

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetStorageGroup(v string)`

SetStorageGroup sets StorageGroup field to given value.

### HasStorageGroup

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasStorageGroup() bool`

HasStorageGroup returns a boolean if a field has been set.

### SetStorageGroupNil

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetStorageGroupNil(b bool)`

 SetStorageGroupNil sets the value for StorageGroup to be an explicit nil

### UnsetStorageGroup
`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) UnsetStorageGroup()`

UnsetStorageGroup ensures that no value is present for StorageGroup, not even an explicit nil
### GetNamespace

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetStorageServer

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetStorageServer() map[string]interface{}`

GetStorageServer returns the StorageServer field if non-nil, zero value otherwise.

### GetStorageServerOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetStorageServerOk() (*map[string]interface{}, bool)`

GetStorageServerOk returns a tuple with the StorageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServer

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetStorageServer(v map[string]interface{})`

SetStorageServer sets StorageServer field to given value.

### HasStorageServer

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasStorageServer() bool`

HasStorageServer returns a boolean if a field has been set.

### GetSource

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUniqueId

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### SetUniqueIdNil

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetUniqueIdNil(b bool)`

 SetUniqueIdNil sets the value for UniqueId to be an explicit nil

### UnsetUniqueId
`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) UnsetUniqueId()`

UnsetUniqueId ensures that no value is present for UniqueId, not even an explicit nil
### GetInternalId

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetProvisionType

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetProvisionType() string`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetProvisionTypeOk() (*string, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetProvisionType(v string)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### SetProvisionTypeNil

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetProvisionTypeNil(b bool)`

 SetProvisionTypeNil sets the value for ProvisionType to be an explicit nil

### UnsetProvisionType
`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) UnsetProvisionType()`

UnsetProvisionType ensures that no value is present for ProvisionType, not even an explicit nil
### GetCopyType

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetCopyType() string`

GetCopyType returns the CopyType field if non-nil, zero value otherwise.

### GetCopyTypeOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetCopyTypeOk() (*string, bool)`

GetCopyTypeOk returns a tuple with the CopyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyType

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetCopyType(v string)`

SetCopyType sets CopyType field to given value.

### HasCopyType

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasCopyType() bool`

HasCopyType returns a boolean if a field has been set.

### SetCopyTypeNil

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetCopyTypeNil(b bool)`

 SetCopyTypeNil sets the value for CopyType to be an explicit nil

### UnsetCopyType
`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) UnsetCopyType()`

UnsetCopyType ensures that no value is present for CopyType, not even an explicit nil
### GetFiberWwn

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetFiberWwn() string`

GetFiberWwn returns the FiberWwn field if non-nil, zero value otherwise.

### GetFiberWwnOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetFiberWwnOk() (*string, bool)`

GetFiberWwnOk returns a tuple with the FiberWwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiberWwn

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetFiberWwn(v string)`

SetFiberWwn sets FiberWwn field to given value.

### HasFiberWwn

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasFiberWwn() bool`

HasFiberWwn returns a boolean if a field has been set.

### SetFiberWwnNil

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetFiberWwnNil(b bool)`

 SetFiberWwnNil sets the value for FiberWwn to be an explicit nil

### UnsetFiberWwn
`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) UnsetFiberWwn()`

UnsetFiberWwn ensures that no value is present for FiberWwn, not even an explicit nil
### GetFileName

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetClaimName

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetClaimName() string`

GetClaimName returns the ClaimName field if non-nil, zero value otherwise.

### GetClaimNameOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetClaimNameOk() (*string, bool)`

GetClaimNameOk returns a tuple with the ClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimName

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetClaimName(v string)`

SetClaimName sets ClaimName field to given value.

### HasClaimName

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasClaimName() bool`

HasClaimName returns a boolean if a field has been set.

### SetClaimNameNil

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetClaimNameNil(b bool)`

 SetClaimNameNil sets the value for ClaimName to be an explicit nil

### UnsetClaimName
`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) UnsetClaimName()`

UnsetClaimName ensures that no value is present for ClaimName, not even an explicit nil
### GetSharePath

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetSharePath() string`

GetSharePath returns the SharePath field if non-nil, zero value otherwise.

### GetSharePathOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetSharePathOk() (*string, bool)`

GetSharePathOk returns a tuple with the SharePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePath

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetSharePath(v string)`

SetSharePath sets SharePath field to given value.

### HasSharePath

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasSharePath() bool`

HasSharePath returns a boolean if a field has been set.

### SetSharePathNil

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetSharePathNil(b bool)`

 SetSharePathNil sets the value for SharePath to be an explicit nil

### UnsetSharePath
`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) UnsetSharePath()`

UnsetSharePath ensures that no value is present for SharePath, not even an explicit nil
### GetSourceId

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceImage

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetSourceImage() string`

GetSourceImage returns the SourceImage field if non-nil, zero value otherwise.

### GetSourceImageOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetSourceImageOk() (*string, bool)`

GetSourceImageOk returns a tuple with the SourceImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceImage

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetSourceImage(v string)`

SetSourceImage sets SourceImage field to given value.

### HasSourceImage

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasSourceImage() bool`

HasSourceImage returns a boolean if a field has been set.

### GetImageType

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetImageTypeOk() (*string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetImageType(v string)`

SetImageType sets ImageType field to given value.

### HasImageType

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasImageType() bool`

HasImageType returns a boolean if a field has been set.

### GetOnline

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetRawData

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetRawData() string`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetRawDataOk() (*string, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetRawData(v string)`

SetRawData sets RawData field to given value.

### HasRawData

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasRawData() bool`

HasRawData returns a boolean if a field has been set.

### GetCreateForMultiAttach

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetCreateForMultiAttach() bool`

GetCreateForMultiAttach returns the CreateForMultiAttach field if non-nil, zero value otherwise.

### GetCreateForMultiAttachOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetCreateForMultiAttachOk() (*bool, bool)`

GetCreateForMultiAttachOk returns a tuple with the CreateForMultiAttach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateForMultiAttach

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetCreateForMultiAttach(v bool)`

SetCreateForMultiAttach sets CreateForMultiAttach field to given value.

### HasCreateForMultiAttach

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasCreateForMultiAttach() bool`

HasCreateForMultiAttach returns a boolean if a field has been set.

### GetIsMultiAttach

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetIsMultiAttach() bool`

GetIsMultiAttach returns the IsMultiAttach field if non-nil, zero value otherwise.

### GetIsMultiAttachOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetIsMultiAttachOk() (*bool, bool)`

GetIsMultiAttachOk returns a tuple with the IsMultiAttach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiAttach

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetIsMultiAttach(v bool)`

SetIsMultiAttach sets IsMultiAttach field to given value.

### HasIsMultiAttach

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasIsMultiAttach() bool`

HasIsMultiAttach returns a boolean if a field has been set.

### GetStorageProfile

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetStorageProfile() string`

GetStorageProfile returns the StorageProfile field if non-nil, zero value otherwise.

### GetStorageProfileOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetStorageProfileOk() (*string, bool)`

GetStorageProfileOk returns a tuple with the StorageProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProfile

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetStorageProfile(v string)`

SetStorageProfile sets StorageProfile field to given value.

### HasStorageProfile

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasStorageProfile() bool`

HasStorageProfile returns a boolean if a field has been set.

### SetStorageProfileNil

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetStorageProfileNil(b bool)`

 SetStorageProfileNil sets the value for StorageProfile to be an explicit nil

### UnsetStorageProfile
`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) UnsetStorageProfile()`

UnsetStorageProfile ensures that no value is present for StorageProfile, not even an explicit nil
### GetAccount

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetAccount() UpdateStorageVolumes200ResponseAllOfStorageVolumeAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetAccountOk() (*UpdateStorageVolumes200ResponseAllOfStorageVolumeAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetAccount(v UpdateStorageVolumes200ResponseAllOfStorageVolumeAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetOwner() UpdateStorageVolumes200ResponseAllOfStorageVolumeOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) GetOwnerOk() (*UpdateStorageVolumes200ResponseAllOfStorageVolumeOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) SetOwner(v UpdateStorageVolumes200ResponseAllOfStorageVolumeOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *UpdateStorageVolumes200ResponseAllOfStorageVolume) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


