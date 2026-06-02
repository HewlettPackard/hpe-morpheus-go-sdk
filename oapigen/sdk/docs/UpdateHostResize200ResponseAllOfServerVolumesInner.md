# UpdateHostResize200ResponseAllOfServerVolumesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Controller** | Pointer to [**UpdateHostResize200ResponseAllOfServerVolumesInnerController**](UpdateHostResize200ResponseAllOfServerVolumesInnerController.md) |  | [optional] 
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
**Type** | Pointer to [**UpdateHostResize200ResponseAllOfServerVolumesInnerType**](UpdateHostResize200ResponseAllOfServerVolumesInnerType.md) |  | [optional] 
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
**Zone** | Pointer to [**UpdateHostResize200ResponseAllOfServerVolumesInnerZone**](UpdateHostResize200ResponseAllOfServerVolumesInnerZone.md) |  | [optional] 
**ZoneId** | Pointer to **int64** |  | [optional] 
**Datastore** | Pointer to [**UpdateHostResize200ResponseAllOfServerVolumesInnerDatastore**](UpdateHostResize200ResponseAllOfServerVolumesInnerDatastore.md) |  | [optional] 
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
**Account** | Pointer to [**UpdateHostResize200ResponseAllOfServerVolumesInnerAccount**](UpdateHostResize200ResponseAllOfServerVolumesInnerAccount.md) |  | [optional] 
**Owner** | Pointer to [**UpdateHostResize200ResponseAllOfServerVolumesInnerOwner**](UpdateHostResize200ResponseAllOfServerVolumesInnerOwner.md) |  | [optional] 

## Methods

### NewUpdateHostResize200ResponseAllOfServerVolumesInner

`func NewUpdateHostResize200ResponseAllOfServerVolumesInner() *UpdateHostResize200ResponseAllOfServerVolumesInner`

NewUpdateHostResize200ResponseAllOfServerVolumesInner instantiates a new UpdateHostResize200ResponseAllOfServerVolumesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetController

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetController() UpdateHostResize200ResponseAllOfServerVolumesInnerController`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetControllerOk() (*UpdateHostResize200ResponseAllOfServerVolumesInnerController, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetController(v UpdateHostResize200ResponseAllOfServerVolumesInnerController)`

SetController sets Controller field to given value.

### HasController

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasController() bool`

HasController returns a boolean if a field has been set.

### GetControllerId

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetControllerId() int64`

GetControllerId returns the ControllerId field if non-nil, zero value otherwise.

### GetControllerIdOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetControllerIdOk() (*int64, bool)`

GetControllerIdOk returns a tuple with the ControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerId

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetControllerId(v int64)`

SetControllerId sets ControllerId field to given value.

### HasControllerId

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasControllerId() bool`

HasControllerId returns a boolean if a field has been set.

### SetControllerIdNil

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetControllerIdNil(b bool)`

 SetControllerIdNil sets the value for ControllerId to be an explicit nil

### UnsetControllerId
`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) UnsetControllerId()`

UnsetControllerId ensures that no value is present for ControllerId, not even an explicit nil
### GetControllerMountPoint

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetControllerMountPoint() string`

GetControllerMountPoint returns the ControllerMountPoint field if non-nil, zero value otherwise.

### GetControllerMountPointOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetControllerMountPointOk() (*string, bool)`

GetControllerMountPointOk returns a tuple with the ControllerMountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerMountPoint

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetControllerMountPoint(v string)`

SetControllerMountPoint sets ControllerMountPoint field to given value.

### HasControllerMountPoint

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasControllerMountPoint() bool`

HasControllerMountPoint returns a boolean if a field has been set.

### SetControllerMountPointNil

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetControllerMountPointNil(b bool)`

 SetControllerMountPointNil sets the value for ControllerMountPoint to be an explicit nil

### UnsetControllerMountPoint
`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) UnsetControllerMountPoint()`

UnsetControllerMountPoint ensures that no value is present for ControllerMountPoint, not even an explicit nil
### GetResizeable

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetResizeable() bool`

GetResizeable returns the Resizeable field if non-nil, zero value otherwise.

### GetResizeableOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetResizeableOk() (*bool, bool)`

GetResizeableOk returns a tuple with the Resizeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeable

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetResizeable(v bool)`

SetResizeable sets Resizeable field to given value.

### HasResizeable

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasResizeable() bool`

HasResizeable returns a boolean if a field has been set.

### SetResizeableNil

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetResizeableNil(b bool)`

 SetResizeableNil sets the value for Resizeable to be an explicit nil

### UnsetResizeable
`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) UnsetResizeable()`

UnsetResizeable ensures that no value is present for Resizeable, not even an explicit nil
### GetRootVolume

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetRootVolume() bool`

GetRootVolume returns the RootVolume field if non-nil, zero value otherwise.

### GetRootVolumeOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetRootVolumeOk() (*bool, bool)`

GetRootVolumeOk returns a tuple with the RootVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolume

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetRootVolume(v bool)`

SetRootVolume sets RootVolume field to given value.

### HasRootVolume

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasRootVolume() bool`

HasRootVolume returns a boolean if a field has been set.

### GetUnitNumber

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetUnitNumber() string`

GetUnitNumber returns the UnitNumber field if non-nil, zero value otherwise.

### GetUnitNumberOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetUnitNumberOk() (*string, bool)`

GetUnitNumberOk returns a tuple with the UnitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitNumber

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetUnitNumber(v string)`

SetUnitNumber sets UnitNumber field to given value.

### HasUnitNumber

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasUnitNumber() bool`

HasUnitNumber returns a boolean if a field has been set.

### SetUnitNumberNil

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetUnitNumberNil(b bool)`

 SetUnitNumberNil sets the value for UnitNumber to be an explicit nil

### UnsetUnitNumber
`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) UnsetUnitNumber()`

UnsetUnitNumber ensures that no value is present for UnitNumber, not even an explicit nil
### GetDeviceName

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceDisplayName

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetDeviceDisplayName() string`

GetDeviceDisplayName returns the DeviceDisplayName field if non-nil, zero value otherwise.

### GetDeviceDisplayNameOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetDeviceDisplayNameOk() (*string, bool)`

GetDeviceDisplayNameOk returns a tuple with the DeviceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDisplayName

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetDeviceDisplayName(v string)`

SetDeviceDisplayName sets DeviceDisplayName field to given value.

### HasDeviceDisplayName

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasDeviceDisplayName() bool`

HasDeviceDisplayName returns a boolean if a field has been set.

### GetVolumeName

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.

### GetVolumePath

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetVolumePath() string`

GetVolumePath returns the VolumePath field if non-nil, zero value otherwise.

### GetVolumePathOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetVolumePathOk() (*string, bool)`

GetVolumePathOk returns a tuple with the VolumePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumePath

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetVolumePath(v string)`

SetVolumePath sets VolumePath field to given value.

### HasVolumePath

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasVolumePath() bool`

HasVolumePath returns a boolean if a field has been set.

### GetVolumeType

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetRefType

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetDiskMode

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetDiskMode() string`

GetDiskMode returns the DiskMode field if non-nil, zero value otherwise.

### GetDiskModeOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetDiskModeOk() (*string, bool)`

GetDiskModeOk returns a tuple with the DiskMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskMode

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetDiskMode(v string)`

SetDiskMode sets DiskMode field to given value.

### HasDiskMode

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasDiskMode() bool`

HasDiskMode returns a boolean if a field has been set.

### GetDiskType

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetDiskType() string`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetDiskTypeOk() (*string, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetDiskType(v string)`

SetDiskType sets DiskType field to given value.

### HasDiskType

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasDiskType() bool`

HasDiskType returns a boolean if a field has been set.

### GetType

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetType() UpdateHostResize200ResponseAllOfServerVolumesInnerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetTypeOk() (*UpdateHostResize200ResponseAllOfServerVolumesInnerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetType(v UpdateHostResize200ResponseAllOfServerVolumesInnerType)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypeId

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetCategory

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetConfigurableIOPS

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetConfigurableIOPS() bool`

GetConfigurableIOPS returns the ConfigurableIOPS field if non-nil, zero value otherwise.

### GetConfigurableIOPSOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetConfigurableIOPSOk() (*bool, bool)`

GetConfigurableIOPSOk returns a tuple with the ConfigurableIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurableIOPS

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetConfigurableIOPS(v bool)`

SetConfigurableIOPS sets ConfigurableIOPS field to given value.

### HasConfigurableIOPS

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasConfigurableIOPS() bool`

HasConfigurableIOPS returns a boolean if a field has been set.

### GetMaxStorage

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetUsedStorage

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetUsedStorage() int64`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetUsedStorageOk() (*int64, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetUsedStorage(v int64)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetMaxIOPS

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetMaxIOPS() string`

GetMaxIOPS returns the MaxIOPS field if non-nil, zero value otherwise.

### GetMaxIOPSOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetMaxIOPSOk() (*string, bool)`

GetMaxIOPSOk returns a tuple with the MaxIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIOPS

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetMaxIOPS(v string)`

SetMaxIOPS sets MaxIOPS field to given value.

### HasMaxIOPS

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasMaxIOPS() bool`

HasMaxIOPS returns a boolean if a field has been set.

### SetMaxIOPSNil

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetMaxIOPSNil(b bool)`

 SetMaxIOPSNil sets the value for MaxIOPS to be an explicit nil

### UnsetMaxIOPS
`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) UnsetMaxIOPS()`

UnsetMaxIOPS ensures that no value is present for MaxIOPS, not even an explicit nil
### GetUuid

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetActive

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetReadOnly

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetRemovable

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetRemovable() bool`

GetRemovable returns the Removable field if non-nil, zero value otherwise.

### GetRemovableOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetRemovableOk() (*bool, bool)`

GetRemovableOk returns a tuple with the Removable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovable

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetRemovable(v bool)`

SetRemovable sets Removable field to given value.

### HasRemovable

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasRemovable() bool`

HasRemovable returns a boolean if a field has been set.

### GetPoolName

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetZone

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetZone() UpdateHostResize200ResponseAllOfServerVolumesInnerZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetZoneOk() (*UpdateHostResize200ResponseAllOfServerVolumesInnerZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetZone(v UpdateHostResize200ResponseAllOfServerVolumesInnerZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetZoneId

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetDatastore

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetDatastore() UpdateHostResize200ResponseAllOfServerVolumesInnerDatastore`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetDatastoreOk() (*UpdateHostResize200ResponseAllOfServerVolumesInnerDatastore, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetDatastore(v UpdateHostResize200ResponseAllOfServerVolumesInnerDatastore)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetDatastoreId

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetDatastoreId() int64`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetDatastoreIdOk() (*int64, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetDatastoreId(v int64)`

SetDatastoreId sets DatastoreId field to given value.

### HasDatastoreId

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasDatastoreId() bool`

HasDatastoreId returns a boolean if a field has been set.

### SetDatastoreIdNil

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetDatastoreIdNil(b bool)`

 SetDatastoreIdNil sets the value for DatastoreId to be an explicit nil

### UnsetDatastoreId
`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) UnsetDatastoreId()`

UnsetDatastoreId ensures that no value is present for DatastoreId, not even an explicit nil
### GetDatastoreOption

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetDatastoreOption() string`

GetDatastoreOption returns the DatastoreOption field if non-nil, zero value otherwise.

### GetDatastoreOptionOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetDatastoreOptionOk() (*string, bool)`

GetDatastoreOptionOk returns a tuple with the DatastoreOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreOption

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetDatastoreOption(v string)`

SetDatastoreOption sets DatastoreOption field to given value.

### HasDatastoreOption

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasDatastoreOption() bool`

HasDatastoreOption returns a boolean if a field has been set.

### GetStorageGroup

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetStorageGroup() string`

GetStorageGroup returns the StorageGroup field if non-nil, zero value otherwise.

### GetStorageGroupOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetStorageGroupOk() (*string, bool)`

GetStorageGroupOk returns a tuple with the StorageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageGroup

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetStorageGroup(v string)`

SetStorageGroup sets StorageGroup field to given value.

### HasStorageGroup

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasStorageGroup() bool`

HasStorageGroup returns a boolean if a field has been set.

### SetStorageGroupNil

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetStorageGroupNil(b bool)`

 SetStorageGroupNil sets the value for StorageGroup to be an explicit nil

### UnsetStorageGroup
`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) UnsetStorageGroup()`

UnsetStorageGroup ensures that no value is present for StorageGroup, not even an explicit nil
### GetNamespace

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetStorageServer

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetStorageServer() map[string]interface{}`

GetStorageServer returns the StorageServer field if non-nil, zero value otherwise.

### GetStorageServerOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetStorageServerOk() (*map[string]interface{}, bool)`

GetStorageServerOk returns a tuple with the StorageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServer

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetStorageServer(v map[string]interface{})`

SetStorageServer sets StorageServer field to given value.

### HasStorageServer

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasStorageServer() bool`

HasStorageServer returns a boolean if a field has been set.

### GetSource

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUniqueId

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### SetUniqueIdNil

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetUniqueIdNil(b bool)`

 SetUniqueIdNil sets the value for UniqueId to be an explicit nil

### UnsetUniqueId
`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) UnsetUniqueId()`

UnsetUniqueId ensures that no value is present for UniqueId, not even an explicit nil
### GetInternalId

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetProvisionType

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetProvisionType() string`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetProvisionTypeOk() (*string, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetProvisionType(v string)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### SetProvisionTypeNil

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetProvisionTypeNil(b bool)`

 SetProvisionTypeNil sets the value for ProvisionType to be an explicit nil

### UnsetProvisionType
`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) UnsetProvisionType()`

UnsetProvisionType ensures that no value is present for ProvisionType, not even an explicit nil
### GetCopyType

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetCopyType() string`

GetCopyType returns the CopyType field if non-nil, zero value otherwise.

### GetCopyTypeOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetCopyTypeOk() (*string, bool)`

GetCopyTypeOk returns a tuple with the CopyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyType

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetCopyType(v string)`

SetCopyType sets CopyType field to given value.

### HasCopyType

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasCopyType() bool`

HasCopyType returns a boolean if a field has been set.

### SetCopyTypeNil

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetCopyTypeNil(b bool)`

 SetCopyTypeNil sets the value for CopyType to be an explicit nil

### UnsetCopyType
`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) UnsetCopyType()`

UnsetCopyType ensures that no value is present for CopyType, not even an explicit nil
### GetFiberWwn

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetFiberWwn() string`

GetFiberWwn returns the FiberWwn field if non-nil, zero value otherwise.

### GetFiberWwnOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetFiberWwnOk() (*string, bool)`

GetFiberWwnOk returns a tuple with the FiberWwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiberWwn

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetFiberWwn(v string)`

SetFiberWwn sets FiberWwn field to given value.

### HasFiberWwn

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasFiberWwn() bool`

HasFiberWwn returns a boolean if a field has been set.

### SetFiberWwnNil

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetFiberWwnNil(b bool)`

 SetFiberWwnNil sets the value for FiberWwn to be an explicit nil

### UnsetFiberWwn
`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) UnsetFiberWwn()`

UnsetFiberWwn ensures that no value is present for FiberWwn, not even an explicit nil
### GetFileName

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetClaimName

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetClaimName() string`

GetClaimName returns the ClaimName field if non-nil, zero value otherwise.

### GetClaimNameOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetClaimNameOk() (*string, bool)`

GetClaimNameOk returns a tuple with the ClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimName

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetClaimName(v string)`

SetClaimName sets ClaimName field to given value.

### HasClaimName

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasClaimName() bool`

HasClaimName returns a boolean if a field has been set.

### SetClaimNameNil

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetClaimNameNil(b bool)`

 SetClaimNameNil sets the value for ClaimName to be an explicit nil

### UnsetClaimName
`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) UnsetClaimName()`

UnsetClaimName ensures that no value is present for ClaimName, not even an explicit nil
### GetSharePath

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetSharePath() string`

GetSharePath returns the SharePath field if non-nil, zero value otherwise.

### GetSharePathOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetSharePathOk() (*string, bool)`

GetSharePathOk returns a tuple with the SharePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePath

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetSharePath(v string)`

SetSharePath sets SharePath field to given value.

### HasSharePath

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasSharePath() bool`

HasSharePath returns a boolean if a field has been set.

### SetSharePathNil

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetSharePathNil(b bool)`

 SetSharePathNil sets the value for SharePath to be an explicit nil

### UnsetSharePath
`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) UnsetSharePath()`

UnsetSharePath ensures that no value is present for SharePath, not even an explicit nil
### GetSourceId

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceImage

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetSourceImage() string`

GetSourceImage returns the SourceImage field if non-nil, zero value otherwise.

### GetSourceImageOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetSourceImageOk() (*string, bool)`

GetSourceImageOk returns a tuple with the SourceImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceImage

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetSourceImage(v string)`

SetSourceImage sets SourceImage field to given value.

### HasSourceImage

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasSourceImage() bool`

HasSourceImage returns a boolean if a field has been set.

### GetImageType

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetImageTypeOk() (*string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetImageType(v string)`

SetImageType sets ImageType field to given value.

### HasImageType

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasImageType() bool`

HasImageType returns a boolean if a field has been set.

### GetOnline

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetRawData

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetRawData() string`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetRawDataOk() (*string, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetRawData(v string)`

SetRawData sets RawData field to given value.

### HasRawData

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasRawData() bool`

HasRawData returns a boolean if a field has been set.

### GetCreateForMultiAttach

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetCreateForMultiAttach() bool`

GetCreateForMultiAttach returns the CreateForMultiAttach field if non-nil, zero value otherwise.

### GetCreateForMultiAttachOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetCreateForMultiAttachOk() (*bool, bool)`

GetCreateForMultiAttachOk returns a tuple with the CreateForMultiAttach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateForMultiAttach

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetCreateForMultiAttach(v bool)`

SetCreateForMultiAttach sets CreateForMultiAttach field to given value.

### HasCreateForMultiAttach

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasCreateForMultiAttach() bool`

HasCreateForMultiAttach returns a boolean if a field has been set.

### GetIsMultiAttach

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetIsMultiAttach() bool`

GetIsMultiAttach returns the IsMultiAttach field if non-nil, zero value otherwise.

### GetIsMultiAttachOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetIsMultiAttachOk() (*bool, bool)`

GetIsMultiAttachOk returns a tuple with the IsMultiAttach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiAttach

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetIsMultiAttach(v bool)`

SetIsMultiAttach sets IsMultiAttach field to given value.

### HasIsMultiAttach

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasIsMultiAttach() bool`

HasIsMultiAttach returns a boolean if a field has been set.

### GetStorageProfile

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetStorageProfile() string`

GetStorageProfile returns the StorageProfile field if non-nil, zero value otherwise.

### GetStorageProfileOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetStorageProfileOk() (*string, bool)`

GetStorageProfileOk returns a tuple with the StorageProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProfile

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetStorageProfile(v string)`

SetStorageProfile sets StorageProfile field to given value.

### HasStorageProfile

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasStorageProfile() bool`

HasStorageProfile returns a boolean if a field has been set.

### SetStorageProfileNil

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetStorageProfileNil(b bool)`

 SetStorageProfileNil sets the value for StorageProfile to be an explicit nil

### UnsetStorageProfile
`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) UnsetStorageProfile()`

UnsetStorageProfile ensures that no value is present for StorageProfile, not even an explicit nil
### GetAccount

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetAccount() UpdateHostResize200ResponseAllOfServerVolumesInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetAccountOk() (*UpdateHostResize200ResponseAllOfServerVolumesInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetAccount(v UpdateHostResize200ResponseAllOfServerVolumesInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetOwner() UpdateHostResize200ResponseAllOfServerVolumesInnerOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) GetOwnerOk() (*UpdateHostResize200ResponseAllOfServerVolumesInnerOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) SetOwner(v UpdateHostResize200ResponseAllOfServerVolumesInnerOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *UpdateHostResize200ResponseAllOfServerVolumesInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


