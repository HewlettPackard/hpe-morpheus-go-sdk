# ListInstanceServicePlans200ResponsePlansInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **int32** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 
**MaxMemory** | Pointer to **int32** |  | [optional] 
**MaxCpu** | Pointer to **int32** |  | [optional] 
**MaxCores** | Pointer to **int32** |  | [optional] 
**CustomCpu** | Pointer to **bool** |  | [optional] 
**CustomMaxMemory** | Pointer to **bool** |  | [optional] 
**CustomMaxStorage** | Pointer to **bool** |  | [optional] 
**CustomMaxDataStorage** | Pointer to **bool** |  | [optional] 
**CustomCoresPerSocket** | Pointer to **bool** |  | [optional] 
**CoresPerSocket** | Pointer to **int32** |  | [optional] 
**StorageTypes** | Pointer to [**[]ListInstanceServicePlans200ResponsePlansInnerStorageTypesInner**](ListInstanceServicePlans200ResponsePlansInnerStorageTypesInner.md) |  | [optional] 
**RootStorageTypes** | Pointer to [**[]ListInstanceServicePlans200ResponsePlansInnerStorageTypesInner**](ListInstanceServicePlans200ResponsePlansInnerStorageTypesInner.md) |  | [optional] 
**AddVolumes** | Pointer to **bool** |  | [optional] 
**CustomizeVolume** | Pointer to **bool** |  | [optional] 
**RootDiskCustomizable** | Pointer to **bool** |  | [optional] 
**NoDisks** | Pointer to **bool** |  | [optional] 
**HasDatastore** | Pointer to **bool** |  | [optional] 
**MinDisk** | Pointer to **int32** |  | [optional] 
**MaxDisk** | Pointer to **string** |  | [optional] 
**LvmSupported** | Pointer to **bool** |  | [optional] 
**Datastores** | Pointer to [**ListInstanceServicePlans200ResponsePlansInnerDatastores**](ListInstanceServicePlans200ResponsePlansInnerDatastores.md) |  | [optional] 
**SupportsAutoDatastore** | Pointer to **bool** |  | [optional] 
**AutoOptions** | Pointer to [**[]ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner**](ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner.md) |  | [optional] 
**CpuOptions** | Pointer to **[]map[string]interface{}** |  | [optional] 
**CoreOptions** | Pointer to **[]map[string]interface{}** |  | [optional] 
**MemoryOptions** | Pointer to **[]map[string]interface{}** |  | [optional] 
**RootCustomSizeOptions** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomSizeOptions** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomCores** | Pointer to **bool** |  | [optional] 
**MaxDisks** | Pointer to **string** |  | [optional] 
**MemorySizeType** | Pointer to **string** |  | [optional] 

## Methods

### NewListInstanceServicePlans200ResponsePlansInner

`func NewListInstanceServicePlans200ResponsePlansInner() *ListInstanceServicePlans200ResponsePlansInner`

NewListInstanceServicePlans200ResponsePlansInner instantiates a new ListInstanceServicePlans200ResponsePlansInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInstanceServicePlans200ResponsePlansInnerWithDefaults

`func NewListInstanceServicePlans200ResponsePlansInnerWithDefaults() *ListInstanceServicePlans200ResponsePlansInner`

NewListInstanceServicePlans200ResponsePlansInnerWithDefaults instantiates a new ListInstanceServicePlans200ResponsePlansInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCode

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMaxStorage

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMaxMemory

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetMaxMemory() int32`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetMaxMemoryOk() (*int32, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetMaxMemory(v int32)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxCpu

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetMaxCpu() int32`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetMaxCpuOk() (*int32, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetMaxCpu(v int32)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetMaxCores

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetMaxCores() int32`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetMaxCoresOk() (*int32, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetMaxCores(v int32)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetCustomCpu

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetCustomCpu() bool`

GetCustomCpu returns the CustomCpu field if non-nil, zero value otherwise.

### GetCustomCpuOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetCustomCpuOk() (*bool, bool)`

GetCustomCpuOk returns a tuple with the CustomCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCpu

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetCustomCpu(v bool)`

SetCustomCpu sets CustomCpu field to given value.

### HasCustomCpu

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasCustomCpu() bool`

HasCustomCpu returns a boolean if a field has been set.

### GetCustomMaxMemory

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetCustomMaxMemory() bool`

GetCustomMaxMemory returns the CustomMaxMemory field if non-nil, zero value otherwise.

### GetCustomMaxMemoryOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetCustomMaxMemoryOk() (*bool, bool)`

GetCustomMaxMemoryOk returns a tuple with the CustomMaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxMemory

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetCustomMaxMemory(v bool)`

SetCustomMaxMemory sets CustomMaxMemory field to given value.

### HasCustomMaxMemory

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasCustomMaxMemory() bool`

HasCustomMaxMemory returns a boolean if a field has been set.

### GetCustomMaxStorage

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetCustomMaxStorage() bool`

GetCustomMaxStorage returns the CustomMaxStorage field if non-nil, zero value otherwise.

### GetCustomMaxStorageOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetCustomMaxStorageOk() (*bool, bool)`

GetCustomMaxStorageOk returns a tuple with the CustomMaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxStorage

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetCustomMaxStorage(v bool)`

SetCustomMaxStorage sets CustomMaxStorage field to given value.

### HasCustomMaxStorage

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasCustomMaxStorage() bool`

HasCustomMaxStorage returns a boolean if a field has been set.

### GetCustomMaxDataStorage

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetCustomMaxDataStorage() bool`

GetCustomMaxDataStorage returns the CustomMaxDataStorage field if non-nil, zero value otherwise.

### GetCustomMaxDataStorageOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetCustomMaxDataStorageOk() (*bool, bool)`

GetCustomMaxDataStorageOk returns a tuple with the CustomMaxDataStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxDataStorage

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetCustomMaxDataStorage(v bool)`

SetCustomMaxDataStorage sets CustomMaxDataStorage field to given value.

### HasCustomMaxDataStorage

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasCustomMaxDataStorage() bool`

HasCustomMaxDataStorage returns a boolean if a field has been set.

### GetCustomCoresPerSocket

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetCustomCoresPerSocket() bool`

GetCustomCoresPerSocket returns the CustomCoresPerSocket field if non-nil, zero value otherwise.

### GetCustomCoresPerSocketOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetCustomCoresPerSocketOk() (*bool, bool)`

GetCustomCoresPerSocketOk returns a tuple with the CustomCoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCoresPerSocket

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetCustomCoresPerSocket(v bool)`

SetCustomCoresPerSocket sets CustomCoresPerSocket field to given value.

### HasCustomCoresPerSocket

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasCustomCoresPerSocket() bool`

HasCustomCoresPerSocket returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetCoresPerSocket() int32`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetCoresPerSocketOk() (*int32, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetCoresPerSocket(v int32)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### GetStorageTypes

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetStorageTypes() []ListInstanceServicePlans200ResponsePlansInnerStorageTypesInner`

GetStorageTypes returns the StorageTypes field if non-nil, zero value otherwise.

### GetStorageTypesOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetStorageTypesOk() (*[]ListInstanceServicePlans200ResponsePlansInnerStorageTypesInner, bool)`

GetStorageTypesOk returns a tuple with the StorageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageTypes

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetStorageTypes(v []ListInstanceServicePlans200ResponsePlansInnerStorageTypesInner)`

SetStorageTypes sets StorageTypes field to given value.

### HasStorageTypes

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasStorageTypes() bool`

HasStorageTypes returns a boolean if a field has been set.

### GetRootStorageTypes

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetRootStorageTypes() []ListInstanceServicePlans200ResponsePlansInnerStorageTypesInner`

GetRootStorageTypes returns the RootStorageTypes field if non-nil, zero value otherwise.

### GetRootStorageTypesOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetRootStorageTypesOk() (*[]ListInstanceServicePlans200ResponsePlansInnerStorageTypesInner, bool)`

GetRootStorageTypesOk returns a tuple with the RootStorageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootStorageTypes

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetRootStorageTypes(v []ListInstanceServicePlans200ResponsePlansInnerStorageTypesInner)`

SetRootStorageTypes sets RootStorageTypes field to given value.

### HasRootStorageTypes

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasRootStorageTypes() bool`

HasRootStorageTypes returns a boolean if a field has been set.

### GetAddVolumes

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetAddVolumes() bool`

GetAddVolumes returns the AddVolumes field if non-nil, zero value otherwise.

### GetAddVolumesOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetAddVolumesOk() (*bool, bool)`

GetAddVolumesOk returns a tuple with the AddVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddVolumes

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetAddVolumes(v bool)`

SetAddVolumes sets AddVolumes field to given value.

### HasAddVolumes

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasAddVolumes() bool`

HasAddVolumes returns a boolean if a field has been set.

### GetCustomizeVolume

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetCustomizeVolume() bool`

GetCustomizeVolume returns the CustomizeVolume field if non-nil, zero value otherwise.

### GetCustomizeVolumeOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetCustomizeVolumeOk() (*bool, bool)`

GetCustomizeVolumeOk returns a tuple with the CustomizeVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizeVolume

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetCustomizeVolume(v bool)`

SetCustomizeVolume sets CustomizeVolume field to given value.

### HasCustomizeVolume

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasCustomizeVolume() bool`

HasCustomizeVolume returns a boolean if a field has been set.

### GetRootDiskCustomizable

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetRootDiskCustomizable() bool`

GetRootDiskCustomizable returns the RootDiskCustomizable field if non-nil, zero value otherwise.

### GetRootDiskCustomizableOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetRootDiskCustomizableOk() (*bool, bool)`

GetRootDiskCustomizableOk returns a tuple with the RootDiskCustomizable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskCustomizable

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetRootDiskCustomizable(v bool)`

SetRootDiskCustomizable sets RootDiskCustomizable field to given value.

### HasRootDiskCustomizable

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasRootDiskCustomizable() bool`

HasRootDiskCustomizable returns a boolean if a field has been set.

### GetNoDisks

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetNoDisks() bool`

GetNoDisks returns the NoDisks field if non-nil, zero value otherwise.

### GetNoDisksOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetNoDisksOk() (*bool, bool)`

GetNoDisksOk returns a tuple with the NoDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDisks

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetNoDisks(v bool)`

SetNoDisks sets NoDisks field to given value.

### HasNoDisks

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasNoDisks() bool`

HasNoDisks returns a boolean if a field has been set.

### GetHasDatastore

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetHasDatastore() bool`

GetHasDatastore returns the HasDatastore field if non-nil, zero value otherwise.

### GetHasDatastoreOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetHasDatastoreOk() (*bool, bool)`

GetHasDatastoreOk returns a tuple with the HasDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDatastore

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetHasDatastore(v bool)`

SetHasDatastore sets HasDatastore field to given value.

### HasHasDatastore

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasHasDatastore() bool`

HasHasDatastore returns a boolean if a field has been set.

### GetMinDisk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetMinDisk() int32`

GetMinDisk returns the MinDisk field if non-nil, zero value otherwise.

### GetMinDiskOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetMinDiskOk() (*int32, bool)`

GetMinDiskOk returns a tuple with the MinDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDisk

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetMinDisk(v int32)`

SetMinDisk sets MinDisk field to given value.

### HasMinDisk

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasMinDisk() bool`

HasMinDisk returns a boolean if a field has been set.

### GetMaxDisk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetMaxDisk() string`

GetMaxDisk returns the MaxDisk field if non-nil, zero value otherwise.

### GetMaxDiskOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetMaxDiskOk() (*string, bool)`

GetMaxDiskOk returns a tuple with the MaxDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDisk

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetMaxDisk(v string)`

SetMaxDisk sets MaxDisk field to given value.

### HasMaxDisk

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasMaxDisk() bool`

HasMaxDisk returns a boolean if a field has been set.

### GetLvmSupported

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetLvmSupported() bool`

GetLvmSupported returns the LvmSupported field if non-nil, zero value otherwise.

### GetLvmSupportedOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetLvmSupportedOk() (*bool, bool)`

GetLvmSupportedOk returns a tuple with the LvmSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLvmSupported

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetLvmSupported(v bool)`

SetLvmSupported sets LvmSupported field to given value.

### HasLvmSupported

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasLvmSupported() bool`

HasLvmSupported returns a boolean if a field has been set.

### GetDatastores

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetDatastores() ListInstanceServicePlans200ResponsePlansInnerDatastores`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetDatastoresOk() (*ListInstanceServicePlans200ResponsePlansInnerDatastores, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetDatastores(v ListInstanceServicePlans200ResponsePlansInnerDatastores)`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.

### GetSupportsAutoDatastore

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetSupportsAutoDatastore() bool`

GetSupportsAutoDatastore returns the SupportsAutoDatastore field if non-nil, zero value otherwise.

### GetSupportsAutoDatastoreOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetSupportsAutoDatastoreOk() (*bool, bool)`

GetSupportsAutoDatastoreOk returns a tuple with the SupportsAutoDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsAutoDatastore

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetSupportsAutoDatastore(v bool)`

SetSupportsAutoDatastore sets SupportsAutoDatastore field to given value.

### HasSupportsAutoDatastore

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasSupportsAutoDatastore() bool`

HasSupportsAutoDatastore returns a boolean if a field has been set.

### GetAutoOptions

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetAutoOptions() []ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner`

GetAutoOptions returns the AutoOptions field if non-nil, zero value otherwise.

### GetAutoOptionsOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetAutoOptionsOk() (*[]ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner, bool)`

GetAutoOptionsOk returns a tuple with the AutoOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoOptions

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetAutoOptions(v []ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner)`

SetAutoOptions sets AutoOptions field to given value.

### HasAutoOptions

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasAutoOptions() bool`

HasAutoOptions returns a boolean if a field has been set.

### GetCpuOptions

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetCpuOptions() []map[string]interface{}`

GetCpuOptions returns the CpuOptions field if non-nil, zero value otherwise.

### GetCpuOptionsOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetCpuOptionsOk() (*[]map[string]interface{}, bool)`

GetCpuOptionsOk returns a tuple with the CpuOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuOptions

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetCpuOptions(v []map[string]interface{})`

SetCpuOptions sets CpuOptions field to given value.

### HasCpuOptions

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasCpuOptions() bool`

HasCpuOptions returns a boolean if a field has been set.

### GetCoreOptions

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetCoreOptions() []map[string]interface{}`

GetCoreOptions returns the CoreOptions field if non-nil, zero value otherwise.

### GetCoreOptionsOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetCoreOptionsOk() (*[]map[string]interface{}, bool)`

GetCoreOptionsOk returns a tuple with the CoreOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreOptions

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetCoreOptions(v []map[string]interface{})`

SetCoreOptions sets CoreOptions field to given value.

### HasCoreOptions

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasCoreOptions() bool`

HasCoreOptions returns a boolean if a field has been set.

### GetMemoryOptions

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetMemoryOptions() []map[string]interface{}`

GetMemoryOptions returns the MemoryOptions field if non-nil, zero value otherwise.

### GetMemoryOptionsOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetMemoryOptionsOk() (*[]map[string]interface{}, bool)`

GetMemoryOptionsOk returns a tuple with the MemoryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryOptions

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetMemoryOptions(v []map[string]interface{})`

SetMemoryOptions sets MemoryOptions field to given value.

### HasMemoryOptions

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasMemoryOptions() bool`

HasMemoryOptions returns a boolean if a field has been set.

### GetRootCustomSizeOptions

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetRootCustomSizeOptions() map[string]interface{}`

GetRootCustomSizeOptions returns the RootCustomSizeOptions field if non-nil, zero value otherwise.

### GetRootCustomSizeOptionsOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetRootCustomSizeOptionsOk() (*map[string]interface{}, bool)`

GetRootCustomSizeOptionsOk returns a tuple with the RootCustomSizeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCustomSizeOptions

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetRootCustomSizeOptions(v map[string]interface{})`

SetRootCustomSizeOptions sets RootCustomSizeOptions field to given value.

### HasRootCustomSizeOptions

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasRootCustomSizeOptions() bool`

HasRootCustomSizeOptions returns a boolean if a field has been set.

### GetCustomSizeOptions

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetCustomSizeOptions() map[string]interface{}`

GetCustomSizeOptions returns the CustomSizeOptions field if non-nil, zero value otherwise.

### GetCustomSizeOptionsOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetCustomSizeOptionsOk() (*map[string]interface{}, bool)`

GetCustomSizeOptionsOk returns a tuple with the CustomSizeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSizeOptions

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetCustomSizeOptions(v map[string]interface{})`

SetCustomSizeOptions sets CustomSizeOptions field to given value.

### HasCustomSizeOptions

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasCustomSizeOptions() bool`

HasCustomSizeOptions returns a boolean if a field has been set.

### GetCustomCores

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetCustomCores() bool`

GetCustomCores returns the CustomCores field if non-nil, zero value otherwise.

### GetCustomCoresOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetCustomCoresOk() (*bool, bool)`

GetCustomCoresOk returns a tuple with the CustomCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCores

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetCustomCores(v bool)`

SetCustomCores sets CustomCores field to given value.

### HasCustomCores

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasCustomCores() bool`

HasCustomCores returns a boolean if a field has been set.

### GetMaxDisks

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetMaxDisks() string`

GetMaxDisks returns the MaxDisks field if non-nil, zero value otherwise.

### GetMaxDisksOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetMaxDisksOk() (*string, bool)`

GetMaxDisksOk returns a tuple with the MaxDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDisks

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetMaxDisks(v string)`

SetMaxDisks sets MaxDisks field to given value.

### HasMaxDisks

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasMaxDisks() bool`

HasMaxDisks returns a boolean if a field has been set.

### GetMemorySizeType

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetMemorySizeType() string`

GetMemorySizeType returns the MemorySizeType field if non-nil, zero value otherwise.

### GetMemorySizeTypeOk

`func (o *ListInstanceServicePlans200ResponsePlansInner) GetMemorySizeTypeOk() (*string, bool)`

GetMemorySizeTypeOk returns a tuple with the MemorySizeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySizeType

`func (o *ListInstanceServicePlans200ResponsePlansInner) SetMemorySizeType(v string)`

SetMemorySizeType sets MemorySizeType field to given value.

### HasMemorySizeType

`func (o *ListInstanceServicePlans200ResponsePlansInner) HasMemorySizeType() bool`

HasMemorySizeType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


