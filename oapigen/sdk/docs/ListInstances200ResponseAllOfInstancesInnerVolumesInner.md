# ListInstances200ResponseAllOfInstancesInnerVolumesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ControllerId** | Pointer to **NullableInt64** |  | [optional] 
**DatastoreId** | Pointer to **NullableString** |  | [optional] 
**DisplayOrder** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**MaxIOPS** | Pointer to **NullableString** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ShortName** | Pointer to **string** |  | [optional] 
**Resizeable** | Pointer to **bool** |  | [optional] 
**PlanResizable** | Pointer to **bool** |  | [optional] 
**RootVolume** | Pointer to **bool** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**StorageType** | Pointer to **int64** |  | [optional] 
**UnitNumber** | Pointer to **NullableString** |  | [optional] 
**ControllerMountPoint** | Pointer to **NullableString** |  | [optional] 
**CreateForMultiAttach** | Pointer to **bool** |  | [optional] 
**StorageProfile** | Pointer to **NullableString** | Storage Profile Code for the volume storage profile assignment. eg. &#x60;\&quot;kvm-cache-none\&quot;&#x60; or &#x60;\&quot;kvm-cache-directsync\&quot;&#x60;. Use &#x60;/api/provision-types?code&#x3D;kvm&#x60; to see the available &#x60;storageProfiles&#x60; for HVM and KVM. | [optional] 

## Methods

### NewListInstances200ResponseAllOfInstancesInnerVolumesInner

`func NewListInstances200ResponseAllOfInstancesInnerVolumesInner() *ListInstances200ResponseAllOfInstancesInnerVolumesInner`

NewListInstances200ResponseAllOfInstancesInnerVolumesInner instantiates a new ListInstances200ResponseAllOfInstancesInnerVolumesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInstances200ResponseAllOfInstancesInnerVolumesInnerWithDefaults

`func NewListInstances200ResponseAllOfInstancesInnerVolumesInnerWithDefaults() *ListInstances200ResponseAllOfInstancesInnerVolumesInner`

NewListInstances200ResponseAllOfInstancesInnerVolumesInnerWithDefaults instantiates a new ListInstances200ResponseAllOfInstancesInnerVolumesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControllerId

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetControllerId() int64`

GetControllerId returns the ControllerId field if non-nil, zero value otherwise.

### GetControllerIdOk

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetControllerIdOk() (*int64, bool)`

GetControllerIdOk returns a tuple with the ControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerId

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) SetControllerId(v int64)`

SetControllerId sets ControllerId field to given value.

### HasControllerId

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) HasControllerId() bool`

HasControllerId returns a boolean if a field has been set.

### SetControllerIdNil

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) SetControllerIdNil(b bool)`

 SetControllerIdNil sets the value for ControllerId to be an explicit nil

### UnsetControllerId
`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) UnsetControllerId()`

UnsetControllerId ensures that no value is present for ControllerId, not even an explicit nil
### GetDatastoreId

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetDatastoreId() string`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetDatastoreIdOk() (*string, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) SetDatastoreId(v string)`

SetDatastoreId sets DatastoreId field to given value.

### HasDatastoreId

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) HasDatastoreId() bool`

HasDatastoreId returns a boolean if a field has been set.

### SetDatastoreIdNil

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) SetDatastoreIdNil(b bool)`

 SetDatastoreIdNil sets the value for DatastoreId to be an explicit nil

### UnsetDatastoreId
`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) UnsetDatastoreId()`

UnsetDatastoreId ensures that no value is present for DatastoreId, not even an explicit nil
### GetDisplayOrder

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetId

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetMaxIOPS

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetMaxIOPS() string`

GetMaxIOPS returns the MaxIOPS field if non-nil, zero value otherwise.

### GetMaxIOPSOk

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetMaxIOPSOk() (*string, bool)`

GetMaxIOPSOk returns a tuple with the MaxIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIOPS

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) SetMaxIOPS(v string)`

SetMaxIOPS sets MaxIOPS field to given value.

### HasMaxIOPS

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) HasMaxIOPS() bool`

HasMaxIOPS returns a boolean if a field has been set.

### SetMaxIOPSNil

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) SetMaxIOPSNil(b bool)`

 SetMaxIOPSNil sets the value for MaxIOPS to be an explicit nil

### UnsetMaxIOPS
`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) UnsetMaxIOPS()`

UnsetMaxIOPS ensures that no value is present for MaxIOPS, not even an explicit nil
### GetMaxStorage

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetName

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShortName

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetResizeable

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetResizeable() bool`

GetResizeable returns the Resizeable field if non-nil, zero value otherwise.

### GetResizeableOk

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetResizeableOk() (*bool, bool)`

GetResizeableOk returns a tuple with the Resizeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeable

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) SetResizeable(v bool)`

SetResizeable sets Resizeable field to given value.

### HasResizeable

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) HasResizeable() bool`

HasResizeable returns a boolean if a field has been set.

### GetPlanResizable

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetPlanResizable() bool`

GetPlanResizable returns the PlanResizable field if non-nil, zero value otherwise.

### GetPlanResizableOk

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetPlanResizableOk() (*bool, bool)`

GetPlanResizableOk returns a tuple with the PlanResizable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanResizable

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) SetPlanResizable(v bool)`

SetPlanResizable sets PlanResizable field to given value.

### HasPlanResizable

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) HasPlanResizable() bool`

HasPlanResizable returns a boolean if a field has been set.

### GetRootVolume

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetRootVolume() bool`

GetRootVolume returns the RootVolume field if non-nil, zero value otherwise.

### GetRootVolumeOk

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetRootVolumeOk() (*bool, bool)`

GetRootVolumeOk returns a tuple with the RootVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolume

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) SetRootVolume(v bool)`

SetRootVolume sets RootVolume field to given value.

### HasRootVolume

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) HasRootVolume() bool`

HasRootVolume returns a boolean if a field has been set.

### GetSize

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStorageType

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetStorageType() int64`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetStorageTypeOk() (*int64, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) SetStorageType(v int64)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetUnitNumber

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetUnitNumber() string`

GetUnitNumber returns the UnitNumber field if non-nil, zero value otherwise.

### GetUnitNumberOk

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetUnitNumberOk() (*string, bool)`

GetUnitNumberOk returns a tuple with the UnitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitNumber

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) SetUnitNumber(v string)`

SetUnitNumber sets UnitNumber field to given value.

### HasUnitNumber

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) HasUnitNumber() bool`

HasUnitNumber returns a boolean if a field has been set.

### SetUnitNumberNil

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) SetUnitNumberNil(b bool)`

 SetUnitNumberNil sets the value for UnitNumber to be an explicit nil

### UnsetUnitNumber
`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) UnsetUnitNumber()`

UnsetUnitNumber ensures that no value is present for UnitNumber, not even an explicit nil
### GetControllerMountPoint

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetControllerMountPoint() string`

GetControllerMountPoint returns the ControllerMountPoint field if non-nil, zero value otherwise.

### GetControllerMountPointOk

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetControllerMountPointOk() (*string, bool)`

GetControllerMountPointOk returns a tuple with the ControllerMountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerMountPoint

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) SetControllerMountPoint(v string)`

SetControllerMountPoint sets ControllerMountPoint field to given value.

### HasControllerMountPoint

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) HasControllerMountPoint() bool`

HasControllerMountPoint returns a boolean if a field has been set.

### SetControllerMountPointNil

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) SetControllerMountPointNil(b bool)`

 SetControllerMountPointNil sets the value for ControllerMountPoint to be an explicit nil

### UnsetControllerMountPoint
`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) UnsetControllerMountPoint()`

UnsetControllerMountPoint ensures that no value is present for ControllerMountPoint, not even an explicit nil
### GetCreateForMultiAttach

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetCreateForMultiAttach() bool`

GetCreateForMultiAttach returns the CreateForMultiAttach field if non-nil, zero value otherwise.

### GetCreateForMultiAttachOk

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetCreateForMultiAttachOk() (*bool, bool)`

GetCreateForMultiAttachOk returns a tuple with the CreateForMultiAttach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateForMultiAttach

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) SetCreateForMultiAttach(v bool)`

SetCreateForMultiAttach sets CreateForMultiAttach field to given value.

### HasCreateForMultiAttach

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) HasCreateForMultiAttach() bool`

HasCreateForMultiAttach returns a boolean if a field has been set.

### GetStorageProfile

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetStorageProfile() string`

GetStorageProfile returns the StorageProfile field if non-nil, zero value otherwise.

### GetStorageProfileOk

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) GetStorageProfileOk() (*string, bool)`

GetStorageProfileOk returns a tuple with the StorageProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProfile

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) SetStorageProfile(v string)`

SetStorageProfile sets StorageProfile field to given value.

### HasStorageProfile

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) HasStorageProfile() bool`

HasStorageProfile returns a boolean if a field has been set.

### SetStorageProfileNil

`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) SetStorageProfileNil(b bool)`

 SetStorageProfileNil sets the value for StorageProfile to be an explicit nil

### UnsetStorageProfile
`func (o *ListInstances200ResponseAllOfInstancesInnerVolumesInner) UnsetStorageProfile()`

UnsetStorageProfile ensures that no value is present for StorageProfile, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


