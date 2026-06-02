# ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner

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

### NewListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner

`func NewListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner() *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner`

NewListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner instantiates a new ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

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

### SetDescriptionNil

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetController

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetController() ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerController`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetControllerOk() (*ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerController, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetController(v ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerController)`

SetController sets Controller field to given value.

### HasController

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasController() bool`

HasController returns a boolean if a field has been set.

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

### SetControllerIdNil

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetControllerIdNil(b bool)`

 SetControllerIdNil sets the value for ControllerId to be an explicit nil

### UnsetControllerId
`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) UnsetControllerId()`

UnsetControllerId ensures that no value is present for ControllerId, not even an explicit nil
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

### SetControllerMountPointNil

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetControllerMountPointNil(b bool)`

 SetControllerMountPointNil sets the value for ControllerMountPoint to be an explicit nil

### UnsetControllerMountPoint
`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) UnsetControllerMountPoint()`

UnsetControllerMountPoint ensures that no value is present for ControllerMountPoint, not even an explicit nil
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

### SetResizeableNil

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetResizeableNil(b bool)`

 SetResizeableNil sets the value for Resizeable to be an explicit nil

### UnsetResizeable
`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) UnsetResizeable()`

UnsetResizeable ensures that no value is present for Resizeable, not even an explicit nil
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

### SetUnitNumberNil

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetUnitNumberNil(b bool)`

 SetUnitNumberNil sets the value for UnitNumber to be an explicit nil

### UnsetUnitNumber
`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) UnsetUnitNumber()`

UnsetUnitNumber ensures that no value is present for UnitNumber, not even an explicit nil
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

### GetVolumeName

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.

### GetVolumePath

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetVolumePath() string`

GetVolumePath returns the VolumePath field if non-nil, zero value otherwise.

### GetVolumePathOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetVolumePathOk() (*string, bool)`

GetVolumePathOk returns a tuple with the VolumePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumePath

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetVolumePath(v string)`

SetVolumePath sets VolumePath field to given value.

### HasVolumePath

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasVolumePath() bool`

HasVolumePath returns a boolean if a field has been set.

### GetVolumeType

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetRefType

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetDiskMode

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetDiskMode() string`

GetDiskMode returns the DiskMode field if non-nil, zero value otherwise.

### GetDiskModeOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetDiskModeOk() (*string, bool)`

GetDiskModeOk returns a tuple with the DiskMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskMode

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetDiskMode(v string)`

SetDiskMode sets DiskMode field to given value.

### HasDiskMode

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasDiskMode() bool`

HasDiskMode returns a boolean if a field has been set.

### GetDiskType

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetDiskType() string`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetDiskTypeOk() (*string, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetDiskType(v string)`

SetDiskType sets DiskType field to given value.

### HasDiskType

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasDiskType() bool`

HasDiskType returns a boolean if a field has been set.

### GetType

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetType() ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetTypeOk() (*ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetType(v ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerType)`

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

### SetStatusMessageNil

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
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

### GetUsedStorage

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetUsedStorage() int64`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetUsedStorageOk() (*int64, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetUsedStorage(v int64)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

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

### SetMaxIOPSNil

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetMaxIOPSNil(b bool)`

 SetMaxIOPSNil sets the value for MaxIOPS to be an explicit nil

### UnsetMaxIOPS
`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) UnsetMaxIOPS()`

UnsetMaxIOPS ensures that no value is present for MaxIOPS, not even an explicit nil
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

### GetReadOnly

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetRemovable

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetRemovable() bool`

GetRemovable returns the Removable field if non-nil, zero value otherwise.

### GetRemovableOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetRemovableOk() (*bool, bool)`

GetRemovableOk returns a tuple with the Removable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovable

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetRemovable(v bool)`

SetRemovable sets Removable field to given value.

### HasRemovable

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasRemovable() bool`

HasRemovable returns a boolean if a field has been set.

### GetPoolName

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetZone

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetZone() ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetZoneOk() (*ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetZone(v ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerZone)`

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

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetDatastore() ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerDatastore`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetDatastoreOk() (*ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerDatastore, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetDatastore(v ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerDatastore)`

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

### SetDatastoreIdNil

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetDatastoreIdNil(b bool)`

 SetDatastoreIdNil sets the value for DatastoreId to be an explicit nil

### UnsetDatastoreId
`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) UnsetDatastoreId()`

UnsetDatastoreId ensures that no value is present for DatastoreId, not even an explicit nil
### GetDatastoreOption

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetDatastoreOption() string`

GetDatastoreOption returns the DatastoreOption field if non-nil, zero value otherwise.

### GetDatastoreOptionOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetDatastoreOptionOk() (*string, bool)`

GetDatastoreOptionOk returns a tuple with the DatastoreOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreOption

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetDatastoreOption(v string)`

SetDatastoreOption sets DatastoreOption field to given value.

### HasDatastoreOption

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasDatastoreOption() bool`

HasDatastoreOption returns a boolean if a field has been set.

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

### SetStorageGroupNil

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetStorageGroupNil(b bool)`

 SetStorageGroupNil sets the value for StorageGroup to be an explicit nil

### UnsetStorageGroup
`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) UnsetStorageGroup()`

UnsetStorageGroup ensures that no value is present for StorageGroup, not even an explicit nil
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

### SetNamespaceNil

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetStorageServer

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetStorageServer() map[string]interface{}`

GetStorageServer returns the StorageServer field if non-nil, zero value otherwise.

### GetStorageServerOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetStorageServerOk() (*map[string]interface{}, bool)`

GetStorageServerOk returns a tuple with the StorageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServer

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetStorageServer(v map[string]interface{})`

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

### GetUniqueId

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### SetUniqueIdNil

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetUniqueIdNil(b bool)`

 SetUniqueIdNil sets the value for UniqueId to be an explicit nil

### UnsetUniqueId
`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) UnsetUniqueId()`

UnsetUniqueId ensures that no value is present for UniqueId, not even an explicit nil
### GetInternalId

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetProvisionType

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetProvisionType() string`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetProvisionTypeOk() (*string, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetProvisionType(v string)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### SetProvisionTypeNil

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetProvisionTypeNil(b bool)`

 SetProvisionTypeNil sets the value for ProvisionType to be an explicit nil

### UnsetProvisionType
`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) UnsetProvisionType()`

UnsetProvisionType ensures that no value is present for ProvisionType, not even an explicit nil
### GetCopyType

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetCopyType() string`

GetCopyType returns the CopyType field if non-nil, zero value otherwise.

### GetCopyTypeOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetCopyTypeOk() (*string, bool)`

GetCopyTypeOk returns a tuple with the CopyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyType

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetCopyType(v string)`

SetCopyType sets CopyType field to given value.

### HasCopyType

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasCopyType() bool`

HasCopyType returns a boolean if a field has been set.

### SetCopyTypeNil

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetCopyTypeNil(b bool)`

 SetCopyTypeNil sets the value for CopyType to be an explicit nil

### UnsetCopyType
`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) UnsetCopyType()`

UnsetCopyType ensures that no value is present for CopyType, not even an explicit nil
### GetFiberWwn

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetFiberWwn() string`

GetFiberWwn returns the FiberWwn field if non-nil, zero value otherwise.

### GetFiberWwnOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetFiberWwnOk() (*string, bool)`

GetFiberWwnOk returns a tuple with the FiberWwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiberWwn

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetFiberWwn(v string)`

SetFiberWwn sets FiberWwn field to given value.

### HasFiberWwn

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasFiberWwn() bool`

HasFiberWwn returns a boolean if a field has been set.

### SetFiberWwnNil

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetFiberWwnNil(b bool)`

 SetFiberWwnNil sets the value for FiberWwn to be an explicit nil

### UnsetFiberWwn
`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) UnsetFiberWwn()`

UnsetFiberWwn ensures that no value is present for FiberWwn, not even an explicit nil
### GetFileName

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetClaimName

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetClaimName() string`

GetClaimName returns the ClaimName field if non-nil, zero value otherwise.

### GetClaimNameOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetClaimNameOk() (*string, bool)`

GetClaimNameOk returns a tuple with the ClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimName

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetClaimName(v string)`

SetClaimName sets ClaimName field to given value.

### HasClaimName

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasClaimName() bool`

HasClaimName returns a boolean if a field has been set.

### SetClaimNameNil

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetClaimNameNil(b bool)`

 SetClaimNameNil sets the value for ClaimName to be an explicit nil

### UnsetClaimName
`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) UnsetClaimName()`

UnsetClaimName ensures that no value is present for ClaimName, not even an explicit nil
### GetSharePath

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetSharePath() string`

GetSharePath returns the SharePath field if non-nil, zero value otherwise.

### GetSharePathOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetSharePathOk() (*string, bool)`

GetSharePathOk returns a tuple with the SharePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePath

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetSharePath(v string)`

SetSharePath sets SharePath field to given value.

### HasSharePath

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasSharePath() bool`

HasSharePath returns a boolean if a field has been set.

### SetSharePathNil

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetSharePathNil(b bool)`

 SetSharePathNil sets the value for SharePath to be an explicit nil

### UnsetSharePath
`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) UnsetSharePath()`

UnsetSharePath ensures that no value is present for SharePath, not even an explicit nil
### GetSourceId

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceImage

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetSourceImage() string`

GetSourceImage returns the SourceImage field if non-nil, zero value otherwise.

### GetSourceImageOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetSourceImageOk() (*string, bool)`

GetSourceImageOk returns a tuple with the SourceImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceImage

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetSourceImage(v string)`

SetSourceImage sets SourceImage field to given value.

### HasSourceImage

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasSourceImage() bool`

HasSourceImage returns a boolean if a field has been set.

### GetImageType

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetImageTypeOk() (*string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetImageType(v string)`

SetImageType sets ImageType field to given value.

### HasImageType

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasImageType() bool`

HasImageType returns a boolean if a field has been set.

### GetOnline

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetRawData

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetRawData() string`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetRawDataOk() (*string, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetRawData(v string)`

SetRawData sets RawData field to given value.

### HasRawData

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasRawData() bool`

HasRawData returns a boolean if a field has been set.

### GetCreateForMultiAttach

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetCreateForMultiAttach() bool`

GetCreateForMultiAttach returns the CreateForMultiAttach field if non-nil, zero value otherwise.

### GetCreateForMultiAttachOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetCreateForMultiAttachOk() (*bool, bool)`

GetCreateForMultiAttachOk returns a tuple with the CreateForMultiAttach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateForMultiAttach

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetCreateForMultiAttach(v bool)`

SetCreateForMultiAttach sets CreateForMultiAttach field to given value.

### HasCreateForMultiAttach

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasCreateForMultiAttach() bool`

HasCreateForMultiAttach returns a boolean if a field has been set.

### GetIsMultiAttach

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetIsMultiAttach() bool`

GetIsMultiAttach returns the IsMultiAttach field if non-nil, zero value otherwise.

### GetIsMultiAttachOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetIsMultiAttachOk() (*bool, bool)`

GetIsMultiAttachOk returns a tuple with the IsMultiAttach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiAttach

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetIsMultiAttach(v bool)`

SetIsMultiAttach sets IsMultiAttach field to given value.

### HasIsMultiAttach

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasIsMultiAttach() bool`

HasIsMultiAttach returns a boolean if a field has been set.

### GetStorageProfile

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetStorageProfile() string`

GetStorageProfile returns the StorageProfile field if non-nil, zero value otherwise.

### GetStorageProfileOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetStorageProfileOk() (*string, bool)`

GetStorageProfileOk returns a tuple with the StorageProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProfile

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetStorageProfile(v string)`

SetStorageProfile sets StorageProfile field to given value.

### HasStorageProfile

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasStorageProfile() bool`

HasStorageProfile returns a boolean if a field has been set.

### SetStorageProfileNil

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetStorageProfileNil(b bool)`

 SetStorageProfileNil sets the value for StorageProfile to be an explicit nil

### UnsetStorageProfile
`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) UnsetStorageProfile()`

UnsetStorageProfile ensures that no value is present for StorageProfile, not even an explicit nil
### GetAccount

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetAccount() ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetAccountOk() (*ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetAccount(v ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetOwner() ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetOwnerOk() (*ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetOwner(v ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


