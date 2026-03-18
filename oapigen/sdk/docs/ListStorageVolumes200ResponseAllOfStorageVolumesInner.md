# ListStorageVolumes200ResponseAllOfStorageVolumesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Controller** | Pointer to [**ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerController**](ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerController.md) |  | [optional] 
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
**Type** | Pointer to [**ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerType**](ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerType.md) |  | [optional] 
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
**Zone** | Pointer to [**ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerZone**](ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerZone.md) |  | [optional] 
**ZoneId** | Pointer to **int64** |  | [optional] 
**Datastore** | Pointer to [**ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerDatastore**](ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerDatastore.md) |  | [optional] 
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
**Account** | Pointer to [**ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerAccount**](ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerAccount.md) |  | [optional] 
**Owner** | Pointer to [**ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerOwner**](ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerOwner.md) |  | [optional] 

## Methods

### NewListStorageVolumes200ResponseAllOfStorageVolumesInner

`func NewListStorageVolumes200ResponseAllOfStorageVolumesInner() *ListStorageVolumes200ResponseAllOfStorageVolumesInner`

NewListStorageVolumes200ResponseAllOfStorageVolumesInner instantiates a new ListStorageVolumes200ResponseAllOfStorageVolumesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListStorageVolumes200ResponseAllOfStorageVolumesInnerWithDefaults

`func NewListStorageVolumes200ResponseAllOfStorageVolumesInnerWithDefaults() *ListStorageVolumes200ResponseAllOfStorageVolumesInner`

NewListStorageVolumes200ResponseAllOfStorageVolumesInnerWithDefaults instantiates a new ListStorageVolumes200ResponseAllOfStorageVolumesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetController

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetController() ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerController`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetControllerOk() (*ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerController, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetController(v ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerController)`

SetController sets Controller field to given value.

### HasController

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasController() bool`

HasController returns a boolean if a field has been set.

### GetControllerId

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetControllerId() int64`

GetControllerId returns the ControllerId field if non-nil, zero value otherwise.

### GetControllerIdOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetControllerIdOk() (*int64, bool)`

GetControllerIdOk returns a tuple with the ControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerId

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetControllerId(v int64)`

SetControllerId sets ControllerId field to given value.

### HasControllerId

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasControllerId() bool`

HasControllerId returns a boolean if a field has been set.

### SetControllerIdNil

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetControllerIdNil(b bool)`

 SetControllerIdNil sets the value for ControllerId to be an explicit nil

### UnsetControllerId
`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) UnsetControllerId()`

UnsetControllerId ensures that no value is present for ControllerId, not even an explicit nil
### GetControllerMountPoint

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetControllerMountPoint() string`

GetControllerMountPoint returns the ControllerMountPoint field if non-nil, zero value otherwise.

### GetControllerMountPointOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetControllerMountPointOk() (*string, bool)`

GetControllerMountPointOk returns a tuple with the ControllerMountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerMountPoint

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetControllerMountPoint(v string)`

SetControllerMountPoint sets ControllerMountPoint field to given value.

### HasControllerMountPoint

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasControllerMountPoint() bool`

HasControllerMountPoint returns a boolean if a field has been set.

### SetControllerMountPointNil

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetControllerMountPointNil(b bool)`

 SetControllerMountPointNil sets the value for ControllerMountPoint to be an explicit nil

### UnsetControllerMountPoint
`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) UnsetControllerMountPoint()`

UnsetControllerMountPoint ensures that no value is present for ControllerMountPoint, not even an explicit nil
### GetResizeable

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetResizeable() bool`

GetResizeable returns the Resizeable field if non-nil, zero value otherwise.

### GetResizeableOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetResizeableOk() (*bool, bool)`

GetResizeableOk returns a tuple with the Resizeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeable

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetResizeable(v bool)`

SetResizeable sets Resizeable field to given value.

### HasResizeable

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasResizeable() bool`

HasResizeable returns a boolean if a field has been set.

### SetResizeableNil

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetResizeableNil(b bool)`

 SetResizeableNil sets the value for Resizeable to be an explicit nil

### UnsetResizeable
`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) UnsetResizeable()`

UnsetResizeable ensures that no value is present for Resizeable, not even an explicit nil
### GetRootVolume

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetRootVolume() bool`

GetRootVolume returns the RootVolume field if non-nil, zero value otherwise.

### GetRootVolumeOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetRootVolumeOk() (*bool, bool)`

GetRootVolumeOk returns a tuple with the RootVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolume

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetRootVolume(v bool)`

SetRootVolume sets RootVolume field to given value.

### HasRootVolume

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasRootVolume() bool`

HasRootVolume returns a boolean if a field has been set.

### GetUnitNumber

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetUnitNumber() string`

GetUnitNumber returns the UnitNumber field if non-nil, zero value otherwise.

### GetUnitNumberOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetUnitNumberOk() (*string, bool)`

GetUnitNumberOk returns a tuple with the UnitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitNumber

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetUnitNumber(v string)`

SetUnitNumber sets UnitNumber field to given value.

### HasUnitNumber

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasUnitNumber() bool`

HasUnitNumber returns a boolean if a field has been set.

### SetUnitNumberNil

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetUnitNumberNil(b bool)`

 SetUnitNumberNil sets the value for UnitNumber to be an explicit nil

### UnsetUnitNumber
`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) UnsetUnitNumber()`

UnsetUnitNumber ensures that no value is present for UnitNumber, not even an explicit nil
### GetDeviceName

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceDisplayName

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetDeviceDisplayName() string`

GetDeviceDisplayName returns the DeviceDisplayName field if non-nil, zero value otherwise.

### GetDeviceDisplayNameOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetDeviceDisplayNameOk() (*string, bool)`

GetDeviceDisplayNameOk returns a tuple with the DeviceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDisplayName

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetDeviceDisplayName(v string)`

SetDeviceDisplayName sets DeviceDisplayName field to given value.

### HasDeviceDisplayName

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasDeviceDisplayName() bool`

HasDeviceDisplayName returns a boolean if a field has been set.

### GetVolumeName

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.

### GetVolumePath

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetVolumePath() string`

GetVolumePath returns the VolumePath field if non-nil, zero value otherwise.

### GetVolumePathOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetVolumePathOk() (*string, bool)`

GetVolumePathOk returns a tuple with the VolumePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumePath

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetVolumePath(v string)`

SetVolumePath sets VolumePath field to given value.

### HasVolumePath

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasVolumePath() bool`

HasVolumePath returns a boolean if a field has been set.

### GetVolumeType

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetRefType

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetDiskMode

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetDiskMode() string`

GetDiskMode returns the DiskMode field if non-nil, zero value otherwise.

### GetDiskModeOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetDiskModeOk() (*string, bool)`

GetDiskModeOk returns a tuple with the DiskMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskMode

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetDiskMode(v string)`

SetDiskMode sets DiskMode field to given value.

### HasDiskMode

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasDiskMode() bool`

HasDiskMode returns a boolean if a field has been set.

### GetDiskType

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetDiskType() string`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetDiskTypeOk() (*string, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetDiskType(v string)`

SetDiskType sets DiskType field to given value.

### HasDiskType

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasDiskType() bool`

HasDiskType returns a boolean if a field has been set.

### GetType

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetType() ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetTypeOk() (*ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetType(v ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerType)`

SetType sets Type field to given value.

### HasType

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypeId

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetCategory

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetStatus

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetConfigurableIOPS

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetConfigurableIOPS() bool`

GetConfigurableIOPS returns the ConfigurableIOPS field if non-nil, zero value otherwise.

### GetConfigurableIOPSOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetConfigurableIOPSOk() (*bool, bool)`

GetConfigurableIOPSOk returns a tuple with the ConfigurableIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurableIOPS

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetConfigurableIOPS(v bool)`

SetConfigurableIOPS sets ConfigurableIOPS field to given value.

### HasConfigurableIOPS

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasConfigurableIOPS() bool`

HasConfigurableIOPS returns a boolean if a field has been set.

### GetMaxStorage

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetUsedStorage

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetUsedStorage() int64`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetUsedStorageOk() (*int64, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetUsedStorage(v int64)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetMaxIOPS

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetMaxIOPS() string`

GetMaxIOPS returns the MaxIOPS field if non-nil, zero value otherwise.

### GetMaxIOPSOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetMaxIOPSOk() (*string, bool)`

GetMaxIOPSOk returns a tuple with the MaxIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIOPS

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetMaxIOPS(v string)`

SetMaxIOPS sets MaxIOPS field to given value.

### HasMaxIOPS

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasMaxIOPS() bool`

HasMaxIOPS returns a boolean if a field has been set.

### SetMaxIOPSNil

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetMaxIOPSNil(b bool)`

 SetMaxIOPSNil sets the value for MaxIOPS to be an explicit nil

### UnsetMaxIOPS
`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) UnsetMaxIOPS()`

UnsetMaxIOPS ensures that no value is present for MaxIOPS, not even an explicit nil
### GetUuid

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetActive

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetReadOnly

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetRemovable

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetRemovable() bool`

GetRemovable returns the Removable field if non-nil, zero value otherwise.

### GetRemovableOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetRemovableOk() (*bool, bool)`

GetRemovableOk returns a tuple with the Removable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovable

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetRemovable(v bool)`

SetRemovable sets Removable field to given value.

### HasRemovable

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasRemovable() bool`

HasRemovable returns a boolean if a field has been set.

### GetPoolName

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetZone

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetZone() ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetZoneOk() (*ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetZone(v ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetZoneId

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetDatastore

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetDatastore() ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerDatastore`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetDatastoreOk() (*ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerDatastore, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetDatastore(v ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerDatastore)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetDatastoreId

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetDatastoreId() int64`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetDatastoreIdOk() (*int64, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetDatastoreId(v int64)`

SetDatastoreId sets DatastoreId field to given value.

### HasDatastoreId

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasDatastoreId() bool`

HasDatastoreId returns a boolean if a field has been set.

### SetDatastoreIdNil

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetDatastoreIdNil(b bool)`

 SetDatastoreIdNil sets the value for DatastoreId to be an explicit nil

### UnsetDatastoreId
`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) UnsetDatastoreId()`

UnsetDatastoreId ensures that no value is present for DatastoreId, not even an explicit nil
### GetDatastoreOption

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetDatastoreOption() string`

GetDatastoreOption returns the DatastoreOption field if non-nil, zero value otherwise.

### GetDatastoreOptionOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetDatastoreOptionOk() (*string, bool)`

GetDatastoreOptionOk returns a tuple with the DatastoreOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreOption

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetDatastoreOption(v string)`

SetDatastoreOption sets DatastoreOption field to given value.

### HasDatastoreOption

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasDatastoreOption() bool`

HasDatastoreOption returns a boolean if a field has been set.

### GetStorageGroup

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetStorageGroup() string`

GetStorageGroup returns the StorageGroup field if non-nil, zero value otherwise.

### GetStorageGroupOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetStorageGroupOk() (*string, bool)`

GetStorageGroupOk returns a tuple with the StorageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageGroup

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetStorageGroup(v string)`

SetStorageGroup sets StorageGroup field to given value.

### HasStorageGroup

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasStorageGroup() bool`

HasStorageGroup returns a boolean if a field has been set.

### SetStorageGroupNil

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetStorageGroupNil(b bool)`

 SetStorageGroupNil sets the value for StorageGroup to be an explicit nil

### UnsetStorageGroup
`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) UnsetStorageGroup()`

UnsetStorageGroup ensures that no value is present for StorageGroup, not even an explicit nil
### GetNamespace

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetStorageServer

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetStorageServer() map[string]interface{}`

GetStorageServer returns the StorageServer field if non-nil, zero value otherwise.

### GetStorageServerOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetStorageServerOk() (*map[string]interface{}, bool)`

GetStorageServerOk returns a tuple with the StorageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServer

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetStorageServer(v map[string]interface{})`

SetStorageServer sets StorageServer field to given value.

### HasStorageServer

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasStorageServer() bool`

HasStorageServer returns a boolean if a field has been set.

### GetSource

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUniqueId

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### SetUniqueIdNil

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetUniqueIdNil(b bool)`

 SetUniqueIdNil sets the value for UniqueId to be an explicit nil

### UnsetUniqueId
`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) UnsetUniqueId()`

UnsetUniqueId ensures that no value is present for UniqueId, not even an explicit nil
### GetInternalId

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetProvisionType

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetProvisionType() string`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetProvisionTypeOk() (*string, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetProvisionType(v string)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### SetProvisionTypeNil

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetProvisionTypeNil(b bool)`

 SetProvisionTypeNil sets the value for ProvisionType to be an explicit nil

### UnsetProvisionType
`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) UnsetProvisionType()`

UnsetProvisionType ensures that no value is present for ProvisionType, not even an explicit nil
### GetCopyType

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetCopyType() string`

GetCopyType returns the CopyType field if non-nil, zero value otherwise.

### GetCopyTypeOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetCopyTypeOk() (*string, bool)`

GetCopyTypeOk returns a tuple with the CopyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyType

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetCopyType(v string)`

SetCopyType sets CopyType field to given value.

### HasCopyType

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasCopyType() bool`

HasCopyType returns a boolean if a field has been set.

### SetCopyTypeNil

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetCopyTypeNil(b bool)`

 SetCopyTypeNil sets the value for CopyType to be an explicit nil

### UnsetCopyType
`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) UnsetCopyType()`

UnsetCopyType ensures that no value is present for CopyType, not even an explicit nil
### GetFiberWwn

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetFiberWwn() string`

GetFiberWwn returns the FiberWwn field if non-nil, zero value otherwise.

### GetFiberWwnOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetFiberWwnOk() (*string, bool)`

GetFiberWwnOk returns a tuple with the FiberWwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiberWwn

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetFiberWwn(v string)`

SetFiberWwn sets FiberWwn field to given value.

### HasFiberWwn

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasFiberWwn() bool`

HasFiberWwn returns a boolean if a field has been set.

### SetFiberWwnNil

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetFiberWwnNil(b bool)`

 SetFiberWwnNil sets the value for FiberWwn to be an explicit nil

### UnsetFiberWwn
`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) UnsetFiberWwn()`

UnsetFiberWwn ensures that no value is present for FiberWwn, not even an explicit nil
### GetFileName

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetClaimName

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetClaimName() string`

GetClaimName returns the ClaimName field if non-nil, zero value otherwise.

### GetClaimNameOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetClaimNameOk() (*string, bool)`

GetClaimNameOk returns a tuple with the ClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimName

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetClaimName(v string)`

SetClaimName sets ClaimName field to given value.

### HasClaimName

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasClaimName() bool`

HasClaimName returns a boolean if a field has been set.

### SetClaimNameNil

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetClaimNameNil(b bool)`

 SetClaimNameNil sets the value for ClaimName to be an explicit nil

### UnsetClaimName
`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) UnsetClaimName()`

UnsetClaimName ensures that no value is present for ClaimName, not even an explicit nil
### GetSharePath

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetSharePath() string`

GetSharePath returns the SharePath field if non-nil, zero value otherwise.

### GetSharePathOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetSharePathOk() (*string, bool)`

GetSharePathOk returns a tuple with the SharePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePath

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetSharePath(v string)`

SetSharePath sets SharePath field to given value.

### HasSharePath

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasSharePath() bool`

HasSharePath returns a boolean if a field has been set.

### SetSharePathNil

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetSharePathNil(b bool)`

 SetSharePathNil sets the value for SharePath to be an explicit nil

### UnsetSharePath
`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) UnsetSharePath()`

UnsetSharePath ensures that no value is present for SharePath, not even an explicit nil
### GetSourceId

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceImage

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetSourceImage() string`

GetSourceImage returns the SourceImage field if non-nil, zero value otherwise.

### GetSourceImageOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetSourceImageOk() (*string, bool)`

GetSourceImageOk returns a tuple with the SourceImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceImage

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetSourceImage(v string)`

SetSourceImage sets SourceImage field to given value.

### HasSourceImage

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasSourceImage() bool`

HasSourceImage returns a boolean if a field has been set.

### GetImageType

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetImageTypeOk() (*string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetImageType(v string)`

SetImageType sets ImageType field to given value.

### HasImageType

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasImageType() bool`

HasImageType returns a boolean if a field has been set.

### GetOnline

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetRawData

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetRawData() string`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetRawDataOk() (*string, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetRawData(v string)`

SetRawData sets RawData field to given value.

### HasRawData

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasRawData() bool`

HasRawData returns a boolean if a field has been set.

### GetCreateForMultiAttach

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetCreateForMultiAttach() bool`

GetCreateForMultiAttach returns the CreateForMultiAttach field if non-nil, zero value otherwise.

### GetCreateForMultiAttachOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetCreateForMultiAttachOk() (*bool, bool)`

GetCreateForMultiAttachOk returns a tuple with the CreateForMultiAttach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateForMultiAttach

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetCreateForMultiAttach(v bool)`

SetCreateForMultiAttach sets CreateForMultiAttach field to given value.

### HasCreateForMultiAttach

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasCreateForMultiAttach() bool`

HasCreateForMultiAttach returns a boolean if a field has been set.

### GetIsMultiAttach

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetIsMultiAttach() bool`

GetIsMultiAttach returns the IsMultiAttach field if non-nil, zero value otherwise.

### GetIsMultiAttachOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetIsMultiAttachOk() (*bool, bool)`

GetIsMultiAttachOk returns a tuple with the IsMultiAttach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiAttach

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetIsMultiAttach(v bool)`

SetIsMultiAttach sets IsMultiAttach field to given value.

### HasIsMultiAttach

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasIsMultiAttach() bool`

HasIsMultiAttach returns a boolean if a field has been set.

### GetStorageProfile

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetStorageProfile() string`

GetStorageProfile returns the StorageProfile field if non-nil, zero value otherwise.

### GetStorageProfileOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetStorageProfileOk() (*string, bool)`

GetStorageProfileOk returns a tuple with the StorageProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProfile

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetStorageProfile(v string)`

SetStorageProfile sets StorageProfile field to given value.

### HasStorageProfile

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasStorageProfile() bool`

HasStorageProfile returns a boolean if a field has been set.

### SetStorageProfileNil

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetStorageProfileNil(b bool)`

 SetStorageProfileNil sets the value for StorageProfile to be an explicit nil

### UnsetStorageProfile
`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) UnsetStorageProfile()`

UnsetStorageProfile ensures that no value is present for StorageProfile, not even an explicit nil
### GetAccount

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetAccount() ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetAccountOk() (*ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetAccount(v ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetOwner() ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) GetOwnerOk() (*ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) SetOwner(v ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListStorageVolumes200ResponseAllOfStorageVolumesInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


