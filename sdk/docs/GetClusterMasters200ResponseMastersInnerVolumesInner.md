# GetClusterMasters200ResponseMastersInnerVolumesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ControllerId** | Pointer to **string** |  | [optional] 
**ControllerMountPoint** | Pointer to **string** |  | [optional] 
**Resizeable** | Pointer to **bool** |  | [optional] 
**PlanResizable** | Pointer to **bool** |  | [optional] 
**RootVolume** | Pointer to **bool** |  | [optional] 
**UnitNumber** | Pointer to **string** |  | [optional] 
**TypeId** | Pointer to **int64** |  | [optional] 
**ConfigurableIOPS** | Pointer to **bool** |  | [optional] 
**DatastoreId** | Pointer to **string** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 
**DisplayOrder** | Pointer to **int64** |  | [optional] 
**MaxIOPS** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewGetClusterMasters200ResponseMastersInnerVolumesInner

`func NewGetClusterMasters200ResponseMastersInnerVolumesInner() *GetClusterMasters200ResponseMastersInnerVolumesInner`

NewGetClusterMasters200ResponseMastersInnerVolumesInner instantiates a new GetClusterMasters200ResponseMastersInnerVolumesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClusterMasters200ResponseMastersInnerVolumesInnerWithDefaults

`func NewGetClusterMasters200ResponseMastersInnerVolumesInnerWithDefaults() *GetClusterMasters200ResponseMastersInnerVolumesInner`

NewGetClusterMasters200ResponseMastersInnerVolumesInnerWithDefaults instantiates a new GetClusterMasters200ResponseMastersInnerVolumesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetControllerId

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetControllerId() string`

GetControllerId returns the ControllerId field if non-nil, zero value otherwise.

### GetControllerIdOk

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetControllerIdOk() (*string, bool)`

GetControllerIdOk returns a tuple with the ControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerId

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) SetControllerId(v string)`

SetControllerId sets ControllerId field to given value.

### HasControllerId

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) HasControllerId() bool`

HasControllerId returns a boolean if a field has been set.

### GetControllerMountPoint

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetControllerMountPoint() string`

GetControllerMountPoint returns the ControllerMountPoint field if non-nil, zero value otherwise.

### GetControllerMountPointOk

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetControllerMountPointOk() (*string, bool)`

GetControllerMountPointOk returns a tuple with the ControllerMountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerMountPoint

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) SetControllerMountPoint(v string)`

SetControllerMountPoint sets ControllerMountPoint field to given value.

### HasControllerMountPoint

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) HasControllerMountPoint() bool`

HasControllerMountPoint returns a boolean if a field has been set.

### GetResizeable

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetResizeable() bool`

GetResizeable returns the Resizeable field if non-nil, zero value otherwise.

### GetResizeableOk

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetResizeableOk() (*bool, bool)`

GetResizeableOk returns a tuple with the Resizeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeable

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) SetResizeable(v bool)`

SetResizeable sets Resizeable field to given value.

### HasResizeable

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) HasResizeable() bool`

HasResizeable returns a boolean if a field has been set.

### GetPlanResizable

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetPlanResizable() bool`

GetPlanResizable returns the PlanResizable field if non-nil, zero value otherwise.

### GetPlanResizableOk

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetPlanResizableOk() (*bool, bool)`

GetPlanResizableOk returns a tuple with the PlanResizable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanResizable

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) SetPlanResizable(v bool)`

SetPlanResizable sets PlanResizable field to given value.

### HasPlanResizable

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) HasPlanResizable() bool`

HasPlanResizable returns a boolean if a field has been set.

### GetRootVolume

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetRootVolume() bool`

GetRootVolume returns the RootVolume field if non-nil, zero value otherwise.

### GetRootVolumeOk

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetRootVolumeOk() (*bool, bool)`

GetRootVolumeOk returns a tuple with the RootVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolume

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) SetRootVolume(v bool)`

SetRootVolume sets RootVolume field to given value.

### HasRootVolume

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) HasRootVolume() bool`

HasRootVolume returns a boolean if a field has been set.

### GetUnitNumber

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetUnitNumber() string`

GetUnitNumber returns the UnitNumber field if non-nil, zero value otherwise.

### GetUnitNumberOk

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetUnitNumberOk() (*string, bool)`

GetUnitNumberOk returns a tuple with the UnitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitNumber

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) SetUnitNumber(v string)`

SetUnitNumber sets UnitNumber field to given value.

### HasUnitNumber

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) HasUnitNumber() bool`

HasUnitNumber returns a boolean if a field has been set.

### GetTypeId

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetConfigurableIOPS

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetConfigurableIOPS() bool`

GetConfigurableIOPS returns the ConfigurableIOPS field if non-nil, zero value otherwise.

### GetConfigurableIOPSOk

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetConfigurableIOPSOk() (*bool, bool)`

GetConfigurableIOPSOk returns a tuple with the ConfigurableIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurableIOPS

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) SetConfigurableIOPS(v bool)`

SetConfigurableIOPS sets ConfigurableIOPS field to given value.

### HasConfigurableIOPS

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) HasConfigurableIOPS() bool`

HasConfigurableIOPS returns a boolean if a field has been set.

### GetDatastoreId

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetDatastoreId() string`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetDatastoreIdOk() (*string, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) SetDatastoreId(v string)`

SetDatastoreId sets DatastoreId field to given value.

### HasDatastoreId

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) HasDatastoreId() bool`

HasDatastoreId returns a boolean if a field has been set.

### GetMaxStorage

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetMaxIOPS

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetMaxIOPS() string`

GetMaxIOPS returns the MaxIOPS field if non-nil, zero value otherwise.

### GetMaxIOPSOk

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetMaxIOPSOk() (*string, bool)`

GetMaxIOPSOk returns a tuple with the MaxIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIOPS

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) SetMaxIOPS(v string)`

SetMaxIOPS sets MaxIOPS field to given value.

### HasMaxIOPS

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) HasMaxIOPS() bool`

HasMaxIOPS returns a boolean if a field has been set.

### GetUuid

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


