# InstanceCreateSuccessInstanceVolumesInner

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

## Methods

### NewInstanceCreateSuccessInstanceVolumesInner

`func NewInstanceCreateSuccessInstanceVolumesInner() *InstanceCreateSuccessInstanceVolumesInner`

NewInstanceCreateSuccessInstanceVolumesInner instantiates a new InstanceCreateSuccessInstanceVolumesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceCreateSuccessInstanceVolumesInnerWithDefaults

`func NewInstanceCreateSuccessInstanceVolumesInnerWithDefaults() *InstanceCreateSuccessInstanceVolumesInner`

NewInstanceCreateSuccessInstanceVolumesInnerWithDefaults instantiates a new InstanceCreateSuccessInstanceVolumesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControllerId

`func (o *InstanceCreateSuccessInstanceVolumesInner) GetControllerId() int64`

GetControllerId returns the ControllerId field if non-nil, zero value otherwise.

### GetControllerIdOk

`func (o *InstanceCreateSuccessInstanceVolumesInner) GetControllerIdOk() (*int64, bool)`

GetControllerIdOk returns a tuple with the ControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerId

`func (o *InstanceCreateSuccessInstanceVolumesInner) SetControllerId(v int64)`

SetControllerId sets ControllerId field to given value.

### HasControllerId

`func (o *InstanceCreateSuccessInstanceVolumesInner) HasControllerId() bool`

HasControllerId returns a boolean if a field has been set.

### SetControllerIdNil

`func (o *InstanceCreateSuccessInstanceVolumesInner) SetControllerIdNil(b bool)`

 SetControllerIdNil sets the value for ControllerId to be an explicit nil

### UnsetControllerId
`func (o *InstanceCreateSuccessInstanceVolumesInner) UnsetControllerId()`

UnsetControllerId ensures that no value is present for ControllerId, not even an explicit nil
### GetDatastoreId

`func (o *InstanceCreateSuccessInstanceVolumesInner) GetDatastoreId() string`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *InstanceCreateSuccessInstanceVolumesInner) GetDatastoreIdOk() (*string, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *InstanceCreateSuccessInstanceVolumesInner) SetDatastoreId(v string)`

SetDatastoreId sets DatastoreId field to given value.

### HasDatastoreId

`func (o *InstanceCreateSuccessInstanceVolumesInner) HasDatastoreId() bool`

HasDatastoreId returns a boolean if a field has been set.

### SetDatastoreIdNil

`func (o *InstanceCreateSuccessInstanceVolumesInner) SetDatastoreIdNil(b bool)`

 SetDatastoreIdNil sets the value for DatastoreId to be an explicit nil

### UnsetDatastoreId
`func (o *InstanceCreateSuccessInstanceVolumesInner) UnsetDatastoreId()`

UnsetDatastoreId ensures that no value is present for DatastoreId, not even an explicit nil
### GetDisplayOrder

`func (o *InstanceCreateSuccessInstanceVolumesInner) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *InstanceCreateSuccessInstanceVolumesInner) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *InstanceCreateSuccessInstanceVolumesInner) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *InstanceCreateSuccessInstanceVolumesInner) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetId

`func (o *InstanceCreateSuccessInstanceVolumesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceCreateSuccessInstanceVolumesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceCreateSuccessInstanceVolumesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceCreateSuccessInstanceVolumesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *InstanceCreateSuccessInstanceVolumesInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *InstanceCreateSuccessInstanceVolumesInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *InstanceCreateSuccessInstanceVolumesInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *InstanceCreateSuccessInstanceVolumesInner) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetMaxIOPS

`func (o *InstanceCreateSuccessInstanceVolumesInner) GetMaxIOPS() string`

GetMaxIOPS returns the MaxIOPS field if non-nil, zero value otherwise.

### GetMaxIOPSOk

`func (o *InstanceCreateSuccessInstanceVolumesInner) GetMaxIOPSOk() (*string, bool)`

GetMaxIOPSOk returns a tuple with the MaxIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIOPS

`func (o *InstanceCreateSuccessInstanceVolumesInner) SetMaxIOPS(v string)`

SetMaxIOPS sets MaxIOPS field to given value.

### HasMaxIOPS

`func (o *InstanceCreateSuccessInstanceVolumesInner) HasMaxIOPS() bool`

HasMaxIOPS returns a boolean if a field has been set.

### SetMaxIOPSNil

`func (o *InstanceCreateSuccessInstanceVolumesInner) SetMaxIOPSNil(b bool)`

 SetMaxIOPSNil sets the value for MaxIOPS to be an explicit nil

### UnsetMaxIOPS
`func (o *InstanceCreateSuccessInstanceVolumesInner) UnsetMaxIOPS()`

UnsetMaxIOPS ensures that no value is present for MaxIOPS, not even an explicit nil
### GetMaxStorage

`func (o *InstanceCreateSuccessInstanceVolumesInner) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *InstanceCreateSuccessInstanceVolumesInner) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *InstanceCreateSuccessInstanceVolumesInner) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *InstanceCreateSuccessInstanceVolumesInner) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetName

`func (o *InstanceCreateSuccessInstanceVolumesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceCreateSuccessInstanceVolumesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceCreateSuccessInstanceVolumesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceCreateSuccessInstanceVolumesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShortName

`func (o *InstanceCreateSuccessInstanceVolumesInner) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *InstanceCreateSuccessInstanceVolumesInner) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *InstanceCreateSuccessInstanceVolumesInner) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *InstanceCreateSuccessInstanceVolumesInner) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetResizeable

`func (o *InstanceCreateSuccessInstanceVolumesInner) GetResizeable() bool`

GetResizeable returns the Resizeable field if non-nil, zero value otherwise.

### GetResizeableOk

`func (o *InstanceCreateSuccessInstanceVolumesInner) GetResizeableOk() (*bool, bool)`

GetResizeableOk returns a tuple with the Resizeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeable

`func (o *InstanceCreateSuccessInstanceVolumesInner) SetResizeable(v bool)`

SetResizeable sets Resizeable field to given value.

### HasResizeable

`func (o *InstanceCreateSuccessInstanceVolumesInner) HasResizeable() bool`

HasResizeable returns a boolean if a field has been set.

### GetPlanResizable

`func (o *InstanceCreateSuccessInstanceVolumesInner) GetPlanResizable() bool`

GetPlanResizable returns the PlanResizable field if non-nil, zero value otherwise.

### GetPlanResizableOk

`func (o *InstanceCreateSuccessInstanceVolumesInner) GetPlanResizableOk() (*bool, bool)`

GetPlanResizableOk returns a tuple with the PlanResizable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanResizable

`func (o *InstanceCreateSuccessInstanceVolumesInner) SetPlanResizable(v bool)`

SetPlanResizable sets PlanResizable field to given value.

### HasPlanResizable

`func (o *InstanceCreateSuccessInstanceVolumesInner) HasPlanResizable() bool`

HasPlanResizable returns a boolean if a field has been set.

### GetRootVolume

`func (o *InstanceCreateSuccessInstanceVolumesInner) GetRootVolume() bool`

GetRootVolume returns the RootVolume field if non-nil, zero value otherwise.

### GetRootVolumeOk

`func (o *InstanceCreateSuccessInstanceVolumesInner) GetRootVolumeOk() (*bool, bool)`

GetRootVolumeOk returns a tuple with the RootVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolume

`func (o *InstanceCreateSuccessInstanceVolumesInner) SetRootVolume(v bool)`

SetRootVolume sets RootVolume field to given value.

### HasRootVolume

`func (o *InstanceCreateSuccessInstanceVolumesInner) HasRootVolume() bool`

HasRootVolume returns a boolean if a field has been set.

### GetSize

`func (o *InstanceCreateSuccessInstanceVolumesInner) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *InstanceCreateSuccessInstanceVolumesInner) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *InstanceCreateSuccessInstanceVolumesInner) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *InstanceCreateSuccessInstanceVolumesInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStorageType

`func (o *InstanceCreateSuccessInstanceVolumesInner) GetStorageType() int64`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *InstanceCreateSuccessInstanceVolumesInner) GetStorageTypeOk() (*int64, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *InstanceCreateSuccessInstanceVolumesInner) SetStorageType(v int64)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *InstanceCreateSuccessInstanceVolumesInner) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetUnitNumber

`func (o *InstanceCreateSuccessInstanceVolumesInner) GetUnitNumber() string`

GetUnitNumber returns the UnitNumber field if non-nil, zero value otherwise.

### GetUnitNumberOk

`func (o *InstanceCreateSuccessInstanceVolumesInner) GetUnitNumberOk() (*string, bool)`

GetUnitNumberOk returns a tuple with the UnitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitNumber

`func (o *InstanceCreateSuccessInstanceVolumesInner) SetUnitNumber(v string)`

SetUnitNumber sets UnitNumber field to given value.

### HasUnitNumber

`func (o *InstanceCreateSuccessInstanceVolumesInner) HasUnitNumber() bool`

HasUnitNumber returns a boolean if a field has been set.

### SetUnitNumberNil

`func (o *InstanceCreateSuccessInstanceVolumesInner) SetUnitNumberNil(b bool)`

 SetUnitNumberNil sets the value for UnitNumber to be an explicit nil

### UnsetUnitNumber
`func (o *InstanceCreateSuccessInstanceVolumesInner) UnsetUnitNumber()`

UnsetUnitNumber ensures that no value is present for UnitNumber, not even an explicit nil
### GetControllerMountPoint

`func (o *InstanceCreateSuccessInstanceVolumesInner) GetControllerMountPoint() string`

GetControllerMountPoint returns the ControllerMountPoint field if non-nil, zero value otherwise.

### GetControllerMountPointOk

`func (o *InstanceCreateSuccessInstanceVolumesInner) GetControllerMountPointOk() (*string, bool)`

GetControllerMountPointOk returns a tuple with the ControllerMountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerMountPoint

`func (o *InstanceCreateSuccessInstanceVolumesInner) SetControllerMountPoint(v string)`

SetControllerMountPoint sets ControllerMountPoint field to given value.

### HasControllerMountPoint

`func (o *InstanceCreateSuccessInstanceVolumesInner) HasControllerMountPoint() bool`

HasControllerMountPoint returns a boolean if a field has been set.

### SetControllerMountPointNil

`func (o *InstanceCreateSuccessInstanceVolumesInner) SetControllerMountPointNil(b bool)`

 SetControllerMountPointNil sets the value for ControllerMountPoint to be an explicit nil

### UnsetControllerMountPoint
`func (o *InstanceCreateSuccessInstanceVolumesInner) UnsetControllerMountPoint()`

UnsetControllerMountPoint ensures that no value is present for ControllerMountPoint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


