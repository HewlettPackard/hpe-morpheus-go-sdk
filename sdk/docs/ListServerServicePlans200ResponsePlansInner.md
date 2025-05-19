# ListServerServicePlans200ResponsePlansInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **int64** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 
**MaxMemory** | Pointer to **int64** |  | [optional] 
**MaxCpu** | Pointer to **int64** |  | [optional] 
**MaxCores** | Pointer to **int64** |  | [optional] 
**MaxDataStorage** | Pointer to **int64** |  | [optional] 
**CustomCpu** | Pointer to **bool** |  | [optional] 
**CustomMaxMemory** | Pointer to **bool** |  | [optional] 
**CustomMaxStorage** | Pointer to **bool** |  | [optional] 
**CustomMaxDataStorage** | Pointer to **bool** |  | [optional] 
**CustomCoresPerSocket** | Pointer to **bool** |  | [optional] 
**CoresPerSocket** | Pointer to **int64** |  | [optional] 
**StorageTypes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**RootStorageTypes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**AddVolumes** | Pointer to **bool** |  | [optional] 
**CustomizeVolume** | Pointer to **bool** |  | [optional] 
**RootDiskCustomizable** | Pointer to **bool** |  | [optional] 
**HostDiskMode** | Pointer to **string** |  | [optional] 
**HasDatastore** | Pointer to **string** |  | [optional] 
**LvmSupported** | Pointer to **string** |  | [optional] 
**MinDisk** | Pointer to **string** |  | [optional] 
**MaxDisk** | Pointer to **string** |  | [optional] 
**Datastores** | Pointer to [**ListServerServicePlans200ResponsePlansInnerDatastores**](ListServerServicePlans200ResponsePlansInnerDatastores.md) |  | [optional] 
**SupportsAutoDatastore** | Pointer to **bool** |  | [optional] 
**AutoOptions** | Pointer to [**[]ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner**](ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner.md) |  | [optional] 
**CpuOptions** | Pointer to **[]map[string]interface{}** |  | [optional] 
**MemoryOptions** | Pointer to **[]map[string]interface{}** |  | [optional] 
**RootCustomSizeOptions** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomSizeOptions** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomCores** | Pointer to **bool** |  | [optional] 
**MaxDisks** | Pointer to **string** |  | [optional] 
**MemorySizeType** | Pointer to **string** |  | [optional] 

## Methods

### NewListServerServicePlans200ResponsePlansInner

`func NewListServerServicePlans200ResponsePlansInner() *ListServerServicePlans200ResponsePlansInner`

NewListServerServicePlans200ResponsePlansInner instantiates a new ListServerServicePlans200ResponsePlansInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServerServicePlans200ResponsePlansInnerWithDefaults

`func NewListServerServicePlans200ResponsePlansInnerWithDefaults() *ListServerServicePlans200ResponsePlansInner`

NewListServerServicePlans200ResponsePlansInnerWithDefaults instantiates a new ListServerServicePlans200ResponsePlansInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListServerServicePlans200ResponsePlansInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListServerServicePlans200ResponsePlansInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListServerServicePlans200ResponsePlansInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListServerServicePlans200ResponsePlansInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListServerServicePlans200ResponsePlansInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListServerServicePlans200ResponsePlansInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *ListServerServicePlans200ResponsePlansInner) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListServerServicePlans200ResponsePlansInner) SetValue(v int64)`

SetValue sets Value field to given value.

### HasValue

`func (o *ListServerServicePlans200ResponsePlansInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCode

`func (o *ListServerServicePlans200ResponsePlansInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListServerServicePlans200ResponsePlansInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListServerServicePlans200ResponsePlansInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMaxStorage

`func (o *ListServerServicePlans200ResponsePlansInner) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ListServerServicePlans200ResponsePlansInner) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *ListServerServicePlans200ResponsePlansInner) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMaxMemory

`func (o *ListServerServicePlans200ResponsePlansInner) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *ListServerServicePlans200ResponsePlansInner) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *ListServerServicePlans200ResponsePlansInner) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxCpu

`func (o *ListServerServicePlans200ResponsePlansInner) GetMaxCpu() int64`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetMaxCpuOk() (*int64, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *ListServerServicePlans200ResponsePlansInner) SetMaxCpu(v int64)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *ListServerServicePlans200ResponsePlansInner) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetMaxCores

`func (o *ListServerServicePlans200ResponsePlansInner) GetMaxCores() int64`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetMaxCoresOk() (*int64, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *ListServerServicePlans200ResponsePlansInner) SetMaxCores(v int64)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *ListServerServicePlans200ResponsePlansInner) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetMaxDataStorage

`func (o *ListServerServicePlans200ResponsePlansInner) GetMaxDataStorage() int64`

GetMaxDataStorage returns the MaxDataStorage field if non-nil, zero value otherwise.

### GetMaxDataStorageOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetMaxDataStorageOk() (*int64, bool)`

GetMaxDataStorageOk returns a tuple with the MaxDataStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDataStorage

`func (o *ListServerServicePlans200ResponsePlansInner) SetMaxDataStorage(v int64)`

SetMaxDataStorage sets MaxDataStorage field to given value.

### HasMaxDataStorage

`func (o *ListServerServicePlans200ResponsePlansInner) HasMaxDataStorage() bool`

HasMaxDataStorage returns a boolean if a field has been set.

### GetCustomCpu

`func (o *ListServerServicePlans200ResponsePlansInner) GetCustomCpu() bool`

GetCustomCpu returns the CustomCpu field if non-nil, zero value otherwise.

### GetCustomCpuOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetCustomCpuOk() (*bool, bool)`

GetCustomCpuOk returns a tuple with the CustomCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCpu

`func (o *ListServerServicePlans200ResponsePlansInner) SetCustomCpu(v bool)`

SetCustomCpu sets CustomCpu field to given value.

### HasCustomCpu

`func (o *ListServerServicePlans200ResponsePlansInner) HasCustomCpu() bool`

HasCustomCpu returns a boolean if a field has been set.

### GetCustomMaxMemory

`func (o *ListServerServicePlans200ResponsePlansInner) GetCustomMaxMemory() bool`

GetCustomMaxMemory returns the CustomMaxMemory field if non-nil, zero value otherwise.

### GetCustomMaxMemoryOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetCustomMaxMemoryOk() (*bool, bool)`

GetCustomMaxMemoryOk returns a tuple with the CustomMaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxMemory

`func (o *ListServerServicePlans200ResponsePlansInner) SetCustomMaxMemory(v bool)`

SetCustomMaxMemory sets CustomMaxMemory field to given value.

### HasCustomMaxMemory

`func (o *ListServerServicePlans200ResponsePlansInner) HasCustomMaxMemory() bool`

HasCustomMaxMemory returns a boolean if a field has been set.

### GetCustomMaxStorage

`func (o *ListServerServicePlans200ResponsePlansInner) GetCustomMaxStorage() bool`

GetCustomMaxStorage returns the CustomMaxStorage field if non-nil, zero value otherwise.

### GetCustomMaxStorageOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetCustomMaxStorageOk() (*bool, bool)`

GetCustomMaxStorageOk returns a tuple with the CustomMaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxStorage

`func (o *ListServerServicePlans200ResponsePlansInner) SetCustomMaxStorage(v bool)`

SetCustomMaxStorage sets CustomMaxStorage field to given value.

### HasCustomMaxStorage

`func (o *ListServerServicePlans200ResponsePlansInner) HasCustomMaxStorage() bool`

HasCustomMaxStorage returns a boolean if a field has been set.

### GetCustomMaxDataStorage

`func (o *ListServerServicePlans200ResponsePlansInner) GetCustomMaxDataStorage() bool`

GetCustomMaxDataStorage returns the CustomMaxDataStorage field if non-nil, zero value otherwise.

### GetCustomMaxDataStorageOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetCustomMaxDataStorageOk() (*bool, bool)`

GetCustomMaxDataStorageOk returns a tuple with the CustomMaxDataStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxDataStorage

`func (o *ListServerServicePlans200ResponsePlansInner) SetCustomMaxDataStorage(v bool)`

SetCustomMaxDataStorage sets CustomMaxDataStorage field to given value.

### HasCustomMaxDataStorage

`func (o *ListServerServicePlans200ResponsePlansInner) HasCustomMaxDataStorage() bool`

HasCustomMaxDataStorage returns a boolean if a field has been set.

### GetCustomCoresPerSocket

`func (o *ListServerServicePlans200ResponsePlansInner) GetCustomCoresPerSocket() bool`

GetCustomCoresPerSocket returns the CustomCoresPerSocket field if non-nil, zero value otherwise.

### GetCustomCoresPerSocketOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetCustomCoresPerSocketOk() (*bool, bool)`

GetCustomCoresPerSocketOk returns a tuple with the CustomCoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCoresPerSocket

`func (o *ListServerServicePlans200ResponsePlansInner) SetCustomCoresPerSocket(v bool)`

SetCustomCoresPerSocket sets CustomCoresPerSocket field to given value.

### HasCustomCoresPerSocket

`func (o *ListServerServicePlans200ResponsePlansInner) HasCustomCoresPerSocket() bool`

HasCustomCoresPerSocket returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *ListServerServicePlans200ResponsePlansInner) GetCoresPerSocket() int64`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetCoresPerSocketOk() (*int64, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *ListServerServicePlans200ResponsePlansInner) SetCoresPerSocket(v int64)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *ListServerServicePlans200ResponsePlansInner) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### GetStorageTypes

`func (o *ListServerServicePlans200ResponsePlansInner) GetStorageTypes() []map[string]interface{}`

GetStorageTypes returns the StorageTypes field if non-nil, zero value otherwise.

### GetStorageTypesOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetStorageTypesOk() (*[]map[string]interface{}, bool)`

GetStorageTypesOk returns a tuple with the StorageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageTypes

`func (o *ListServerServicePlans200ResponsePlansInner) SetStorageTypes(v []map[string]interface{})`

SetStorageTypes sets StorageTypes field to given value.

### HasStorageTypes

`func (o *ListServerServicePlans200ResponsePlansInner) HasStorageTypes() bool`

HasStorageTypes returns a boolean if a field has been set.

### GetRootStorageTypes

`func (o *ListServerServicePlans200ResponsePlansInner) GetRootStorageTypes() []map[string]interface{}`

GetRootStorageTypes returns the RootStorageTypes field if non-nil, zero value otherwise.

### GetRootStorageTypesOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetRootStorageTypesOk() (*[]map[string]interface{}, bool)`

GetRootStorageTypesOk returns a tuple with the RootStorageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootStorageTypes

`func (o *ListServerServicePlans200ResponsePlansInner) SetRootStorageTypes(v []map[string]interface{})`

SetRootStorageTypes sets RootStorageTypes field to given value.

### HasRootStorageTypes

`func (o *ListServerServicePlans200ResponsePlansInner) HasRootStorageTypes() bool`

HasRootStorageTypes returns a boolean if a field has been set.

### GetAddVolumes

`func (o *ListServerServicePlans200ResponsePlansInner) GetAddVolumes() bool`

GetAddVolumes returns the AddVolumes field if non-nil, zero value otherwise.

### GetAddVolumesOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetAddVolumesOk() (*bool, bool)`

GetAddVolumesOk returns a tuple with the AddVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddVolumes

`func (o *ListServerServicePlans200ResponsePlansInner) SetAddVolumes(v bool)`

SetAddVolumes sets AddVolumes field to given value.

### HasAddVolumes

`func (o *ListServerServicePlans200ResponsePlansInner) HasAddVolumes() bool`

HasAddVolumes returns a boolean if a field has been set.

### GetCustomizeVolume

`func (o *ListServerServicePlans200ResponsePlansInner) GetCustomizeVolume() bool`

GetCustomizeVolume returns the CustomizeVolume field if non-nil, zero value otherwise.

### GetCustomizeVolumeOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetCustomizeVolumeOk() (*bool, bool)`

GetCustomizeVolumeOk returns a tuple with the CustomizeVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizeVolume

`func (o *ListServerServicePlans200ResponsePlansInner) SetCustomizeVolume(v bool)`

SetCustomizeVolume sets CustomizeVolume field to given value.

### HasCustomizeVolume

`func (o *ListServerServicePlans200ResponsePlansInner) HasCustomizeVolume() bool`

HasCustomizeVolume returns a boolean if a field has been set.

### GetRootDiskCustomizable

`func (o *ListServerServicePlans200ResponsePlansInner) GetRootDiskCustomizable() bool`

GetRootDiskCustomizable returns the RootDiskCustomizable field if non-nil, zero value otherwise.

### GetRootDiskCustomizableOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetRootDiskCustomizableOk() (*bool, bool)`

GetRootDiskCustomizableOk returns a tuple with the RootDiskCustomizable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDiskCustomizable

`func (o *ListServerServicePlans200ResponsePlansInner) SetRootDiskCustomizable(v bool)`

SetRootDiskCustomizable sets RootDiskCustomizable field to given value.

### HasRootDiskCustomizable

`func (o *ListServerServicePlans200ResponsePlansInner) HasRootDiskCustomizable() bool`

HasRootDiskCustomizable returns a boolean if a field has been set.

### GetHostDiskMode

`func (o *ListServerServicePlans200ResponsePlansInner) GetHostDiskMode() string`

GetHostDiskMode returns the HostDiskMode field if non-nil, zero value otherwise.

### GetHostDiskModeOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetHostDiskModeOk() (*string, bool)`

GetHostDiskModeOk returns a tuple with the HostDiskMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostDiskMode

`func (o *ListServerServicePlans200ResponsePlansInner) SetHostDiskMode(v string)`

SetHostDiskMode sets HostDiskMode field to given value.

### HasHostDiskMode

`func (o *ListServerServicePlans200ResponsePlansInner) HasHostDiskMode() bool`

HasHostDiskMode returns a boolean if a field has been set.

### GetHasDatastore

`func (o *ListServerServicePlans200ResponsePlansInner) GetHasDatastore() string`

GetHasDatastore returns the HasDatastore field if non-nil, zero value otherwise.

### GetHasDatastoreOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetHasDatastoreOk() (*string, bool)`

GetHasDatastoreOk returns a tuple with the HasDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDatastore

`func (o *ListServerServicePlans200ResponsePlansInner) SetHasDatastore(v string)`

SetHasDatastore sets HasDatastore field to given value.

### HasHasDatastore

`func (o *ListServerServicePlans200ResponsePlansInner) HasHasDatastore() bool`

HasHasDatastore returns a boolean if a field has been set.

### GetLvmSupported

`func (o *ListServerServicePlans200ResponsePlansInner) GetLvmSupported() string`

GetLvmSupported returns the LvmSupported field if non-nil, zero value otherwise.

### GetLvmSupportedOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetLvmSupportedOk() (*string, bool)`

GetLvmSupportedOk returns a tuple with the LvmSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLvmSupported

`func (o *ListServerServicePlans200ResponsePlansInner) SetLvmSupported(v string)`

SetLvmSupported sets LvmSupported field to given value.

### HasLvmSupported

`func (o *ListServerServicePlans200ResponsePlansInner) HasLvmSupported() bool`

HasLvmSupported returns a boolean if a field has been set.

### GetMinDisk

`func (o *ListServerServicePlans200ResponsePlansInner) GetMinDisk() string`

GetMinDisk returns the MinDisk field if non-nil, zero value otherwise.

### GetMinDiskOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetMinDiskOk() (*string, bool)`

GetMinDiskOk returns a tuple with the MinDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDisk

`func (o *ListServerServicePlans200ResponsePlansInner) SetMinDisk(v string)`

SetMinDisk sets MinDisk field to given value.

### HasMinDisk

`func (o *ListServerServicePlans200ResponsePlansInner) HasMinDisk() bool`

HasMinDisk returns a boolean if a field has been set.

### GetMaxDisk

`func (o *ListServerServicePlans200ResponsePlansInner) GetMaxDisk() string`

GetMaxDisk returns the MaxDisk field if non-nil, zero value otherwise.

### GetMaxDiskOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetMaxDiskOk() (*string, bool)`

GetMaxDiskOk returns a tuple with the MaxDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDisk

`func (o *ListServerServicePlans200ResponsePlansInner) SetMaxDisk(v string)`

SetMaxDisk sets MaxDisk field to given value.

### HasMaxDisk

`func (o *ListServerServicePlans200ResponsePlansInner) HasMaxDisk() bool`

HasMaxDisk returns a boolean if a field has been set.

### GetDatastores

`func (o *ListServerServicePlans200ResponsePlansInner) GetDatastores() ListServerServicePlans200ResponsePlansInnerDatastores`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetDatastoresOk() (*ListServerServicePlans200ResponsePlansInnerDatastores, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *ListServerServicePlans200ResponsePlansInner) SetDatastores(v ListServerServicePlans200ResponsePlansInnerDatastores)`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *ListServerServicePlans200ResponsePlansInner) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.

### GetSupportsAutoDatastore

`func (o *ListServerServicePlans200ResponsePlansInner) GetSupportsAutoDatastore() bool`

GetSupportsAutoDatastore returns the SupportsAutoDatastore field if non-nil, zero value otherwise.

### GetSupportsAutoDatastoreOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetSupportsAutoDatastoreOk() (*bool, bool)`

GetSupportsAutoDatastoreOk returns a tuple with the SupportsAutoDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsAutoDatastore

`func (o *ListServerServicePlans200ResponsePlansInner) SetSupportsAutoDatastore(v bool)`

SetSupportsAutoDatastore sets SupportsAutoDatastore field to given value.

### HasSupportsAutoDatastore

`func (o *ListServerServicePlans200ResponsePlansInner) HasSupportsAutoDatastore() bool`

HasSupportsAutoDatastore returns a boolean if a field has been set.

### GetAutoOptions

`func (o *ListServerServicePlans200ResponsePlansInner) GetAutoOptions() []ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner`

GetAutoOptions returns the AutoOptions field if non-nil, zero value otherwise.

### GetAutoOptionsOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetAutoOptionsOk() (*[]ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner, bool)`

GetAutoOptionsOk returns a tuple with the AutoOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoOptions

`func (o *ListServerServicePlans200ResponsePlansInner) SetAutoOptions(v []ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner)`

SetAutoOptions sets AutoOptions field to given value.

### HasAutoOptions

`func (o *ListServerServicePlans200ResponsePlansInner) HasAutoOptions() bool`

HasAutoOptions returns a boolean if a field has been set.

### GetCpuOptions

`func (o *ListServerServicePlans200ResponsePlansInner) GetCpuOptions() []map[string]interface{}`

GetCpuOptions returns the CpuOptions field if non-nil, zero value otherwise.

### GetCpuOptionsOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetCpuOptionsOk() (*[]map[string]interface{}, bool)`

GetCpuOptionsOk returns a tuple with the CpuOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuOptions

`func (o *ListServerServicePlans200ResponsePlansInner) SetCpuOptions(v []map[string]interface{})`

SetCpuOptions sets CpuOptions field to given value.

### HasCpuOptions

`func (o *ListServerServicePlans200ResponsePlansInner) HasCpuOptions() bool`

HasCpuOptions returns a boolean if a field has been set.

### GetMemoryOptions

`func (o *ListServerServicePlans200ResponsePlansInner) GetMemoryOptions() []map[string]interface{}`

GetMemoryOptions returns the MemoryOptions field if non-nil, zero value otherwise.

### GetMemoryOptionsOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetMemoryOptionsOk() (*[]map[string]interface{}, bool)`

GetMemoryOptionsOk returns a tuple with the MemoryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryOptions

`func (o *ListServerServicePlans200ResponsePlansInner) SetMemoryOptions(v []map[string]interface{})`

SetMemoryOptions sets MemoryOptions field to given value.

### HasMemoryOptions

`func (o *ListServerServicePlans200ResponsePlansInner) HasMemoryOptions() bool`

HasMemoryOptions returns a boolean if a field has been set.

### GetRootCustomSizeOptions

`func (o *ListServerServicePlans200ResponsePlansInner) GetRootCustomSizeOptions() map[string]interface{}`

GetRootCustomSizeOptions returns the RootCustomSizeOptions field if non-nil, zero value otherwise.

### GetRootCustomSizeOptionsOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetRootCustomSizeOptionsOk() (*map[string]interface{}, bool)`

GetRootCustomSizeOptionsOk returns a tuple with the RootCustomSizeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCustomSizeOptions

`func (o *ListServerServicePlans200ResponsePlansInner) SetRootCustomSizeOptions(v map[string]interface{})`

SetRootCustomSizeOptions sets RootCustomSizeOptions field to given value.

### HasRootCustomSizeOptions

`func (o *ListServerServicePlans200ResponsePlansInner) HasRootCustomSizeOptions() bool`

HasRootCustomSizeOptions returns a boolean if a field has been set.

### GetCustomSizeOptions

`func (o *ListServerServicePlans200ResponsePlansInner) GetCustomSizeOptions() map[string]interface{}`

GetCustomSizeOptions returns the CustomSizeOptions field if non-nil, zero value otherwise.

### GetCustomSizeOptionsOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetCustomSizeOptionsOk() (*map[string]interface{}, bool)`

GetCustomSizeOptionsOk returns a tuple with the CustomSizeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSizeOptions

`func (o *ListServerServicePlans200ResponsePlansInner) SetCustomSizeOptions(v map[string]interface{})`

SetCustomSizeOptions sets CustomSizeOptions field to given value.

### HasCustomSizeOptions

`func (o *ListServerServicePlans200ResponsePlansInner) HasCustomSizeOptions() bool`

HasCustomSizeOptions returns a boolean if a field has been set.

### GetCustomCores

`func (o *ListServerServicePlans200ResponsePlansInner) GetCustomCores() bool`

GetCustomCores returns the CustomCores field if non-nil, zero value otherwise.

### GetCustomCoresOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetCustomCoresOk() (*bool, bool)`

GetCustomCoresOk returns a tuple with the CustomCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCores

`func (o *ListServerServicePlans200ResponsePlansInner) SetCustomCores(v bool)`

SetCustomCores sets CustomCores field to given value.

### HasCustomCores

`func (o *ListServerServicePlans200ResponsePlansInner) HasCustomCores() bool`

HasCustomCores returns a boolean if a field has been set.

### GetMaxDisks

`func (o *ListServerServicePlans200ResponsePlansInner) GetMaxDisks() string`

GetMaxDisks returns the MaxDisks field if non-nil, zero value otherwise.

### GetMaxDisksOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetMaxDisksOk() (*string, bool)`

GetMaxDisksOk returns a tuple with the MaxDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDisks

`func (o *ListServerServicePlans200ResponsePlansInner) SetMaxDisks(v string)`

SetMaxDisks sets MaxDisks field to given value.

### HasMaxDisks

`func (o *ListServerServicePlans200ResponsePlansInner) HasMaxDisks() bool`

HasMaxDisks returns a boolean if a field has been set.

### GetMemorySizeType

`func (o *ListServerServicePlans200ResponsePlansInner) GetMemorySizeType() string`

GetMemorySizeType returns the MemorySizeType field if non-nil, zero value otherwise.

### GetMemorySizeTypeOk

`func (o *ListServerServicePlans200ResponsePlansInner) GetMemorySizeTypeOk() (*string, bool)`

GetMemorySizeTypeOk returns a tuple with the MemorySizeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySizeType

`func (o *ListServerServicePlans200ResponsePlansInner) SetMemorySizeType(v string)`

SetMemorySizeType sets MemorySizeType field to given value.

### HasMemorySizeType

`func (o *ListServerServicePlans200ResponsePlansInner) HasMemorySizeType() bool`

HasMemorySizeType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


