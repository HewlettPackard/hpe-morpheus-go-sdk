# AddScaleThresholds200ResponseAllOfScaleThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**AutoUp** | Pointer to **bool** |  | [optional] 
**AutoDown** | Pointer to **bool** |  | [optional] 
**MinCount** | Pointer to **int64** |  | [optional] 
**MaxCount** | Pointer to **int64** |  | [optional] 
**ScaleIncrement** | Pointer to **int64** |  | [optional] 
**CpuEnabled** | Pointer to **bool** |  | [optional] 
**MinCpu** | Pointer to **int64** |  | [optional] 
**MaxCpu** | Pointer to **int64** |  | [optional] 
**MemoryEnabled** | Pointer to **bool** |  | [optional] 
**MinMemory** | Pointer to **int64** |  | [optional] 
**MaxMemory** | Pointer to **int64** |  | [optional] 
**DiskEnabled** | Pointer to **bool** |  | [optional] 
**MinDisk** | Pointer to **int64** |  | [optional] 
**MaxDisk** | Pointer to **int64** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAddScaleThresholds200ResponseAllOfScaleThreshold

`func NewAddScaleThresholds200ResponseAllOfScaleThreshold() *AddScaleThresholds200ResponseAllOfScaleThreshold`

NewAddScaleThresholds200ResponseAllOfScaleThreshold instantiates a new AddScaleThresholds200ResponseAllOfScaleThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddScaleThresholds200ResponseAllOfScaleThresholdWithDefaults

`func NewAddScaleThresholds200ResponseAllOfScaleThresholdWithDefaults() *AddScaleThresholds200ResponseAllOfScaleThreshold`

NewAddScaleThresholds200ResponseAllOfScaleThresholdWithDefaults instantiates a new AddScaleThresholds200ResponseAllOfScaleThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAutoUp

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetAutoUp() bool`

GetAutoUp returns the AutoUp field if non-nil, zero value otherwise.

### GetAutoUpOk

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetAutoUpOk() (*bool, bool)`

GetAutoUpOk returns a tuple with the AutoUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUp

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) SetAutoUp(v bool)`

SetAutoUp sets AutoUp field to given value.

### HasAutoUp

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) HasAutoUp() bool`

HasAutoUp returns a boolean if a field has been set.

### GetAutoDown

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetAutoDown() bool`

GetAutoDown returns the AutoDown field if non-nil, zero value otherwise.

### GetAutoDownOk

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetAutoDownOk() (*bool, bool)`

GetAutoDownOk returns a tuple with the AutoDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDown

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) SetAutoDown(v bool)`

SetAutoDown sets AutoDown field to given value.

### HasAutoDown

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) HasAutoDown() bool`

HasAutoDown returns a boolean if a field has been set.

### GetMinCount

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetMinCount() int64`

GetMinCount returns the MinCount field if non-nil, zero value otherwise.

### GetMinCountOk

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetMinCountOk() (*int64, bool)`

GetMinCountOk returns a tuple with the MinCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCount

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) SetMinCount(v int64)`

SetMinCount sets MinCount field to given value.

### HasMinCount

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) HasMinCount() bool`

HasMinCount returns a boolean if a field has been set.

### GetMaxCount

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetMaxCount() int64`

GetMaxCount returns the MaxCount field if non-nil, zero value otherwise.

### GetMaxCountOk

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetMaxCountOk() (*int64, bool)`

GetMaxCountOk returns a tuple with the MaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCount

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) SetMaxCount(v int64)`

SetMaxCount sets MaxCount field to given value.

### HasMaxCount

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) HasMaxCount() bool`

HasMaxCount returns a boolean if a field has been set.

### GetScaleIncrement

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetScaleIncrement() int64`

GetScaleIncrement returns the ScaleIncrement field if non-nil, zero value otherwise.

### GetScaleIncrementOk

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetScaleIncrementOk() (*int64, bool)`

GetScaleIncrementOk returns a tuple with the ScaleIncrement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleIncrement

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) SetScaleIncrement(v int64)`

SetScaleIncrement sets ScaleIncrement field to given value.

### HasScaleIncrement

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) HasScaleIncrement() bool`

HasScaleIncrement returns a boolean if a field has been set.

### GetCpuEnabled

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetCpuEnabled() bool`

GetCpuEnabled returns the CpuEnabled field if non-nil, zero value otherwise.

### GetCpuEnabledOk

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetCpuEnabledOk() (*bool, bool)`

GetCpuEnabledOk returns a tuple with the CpuEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuEnabled

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) SetCpuEnabled(v bool)`

SetCpuEnabled sets CpuEnabled field to given value.

### HasCpuEnabled

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) HasCpuEnabled() bool`

HasCpuEnabled returns a boolean if a field has been set.

### GetMinCpu

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetMinCpu() int64`

GetMinCpu returns the MinCpu field if non-nil, zero value otherwise.

### GetMinCpuOk

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetMinCpuOk() (*int64, bool)`

GetMinCpuOk returns a tuple with the MinCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCpu

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) SetMinCpu(v int64)`

SetMinCpu sets MinCpu field to given value.

### HasMinCpu

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) HasMinCpu() bool`

HasMinCpu returns a boolean if a field has been set.

### GetMaxCpu

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetMaxCpu() int64`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetMaxCpuOk() (*int64, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) SetMaxCpu(v int64)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetMemoryEnabled

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetMemoryEnabled() bool`

GetMemoryEnabled returns the MemoryEnabled field if non-nil, zero value otherwise.

### GetMemoryEnabledOk

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetMemoryEnabledOk() (*bool, bool)`

GetMemoryEnabledOk returns a tuple with the MemoryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryEnabled

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) SetMemoryEnabled(v bool)`

SetMemoryEnabled sets MemoryEnabled field to given value.

### HasMemoryEnabled

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) HasMemoryEnabled() bool`

HasMemoryEnabled returns a boolean if a field has been set.

### GetMinMemory

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetMinMemory() int64`

GetMinMemory returns the MinMemory field if non-nil, zero value otherwise.

### GetMinMemoryOk

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetMinMemoryOk() (*int64, bool)`

GetMinMemoryOk returns a tuple with the MinMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMemory

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) SetMinMemory(v int64)`

SetMinMemory sets MinMemory field to given value.

### HasMinMemory

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) HasMinMemory() bool`

HasMinMemory returns a boolean if a field has been set.

### GetMaxMemory

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetDiskEnabled

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetDiskEnabled() bool`

GetDiskEnabled returns the DiskEnabled field if non-nil, zero value otherwise.

### GetDiskEnabledOk

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetDiskEnabledOk() (*bool, bool)`

GetDiskEnabledOk returns a tuple with the DiskEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEnabled

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) SetDiskEnabled(v bool)`

SetDiskEnabled sets DiskEnabled field to given value.

### HasDiskEnabled

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) HasDiskEnabled() bool`

HasDiskEnabled returns a boolean if a field has been set.

### GetMinDisk

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetMinDisk() int64`

GetMinDisk returns the MinDisk field if non-nil, zero value otherwise.

### GetMinDiskOk

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetMinDiskOk() (*int64, bool)`

GetMinDiskOk returns a tuple with the MinDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDisk

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) SetMinDisk(v int64)`

SetMinDisk sets MinDisk field to given value.

### HasMinDisk

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) HasMinDisk() bool`

HasMinDisk returns a boolean if a field has been set.

### GetMaxDisk

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetMaxDisk() int64`

GetMaxDisk returns the MaxDisk field if non-nil, zero value otherwise.

### GetMaxDiskOk

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetMaxDiskOk() (*int64, bool)`

GetMaxDiskOk returns a tuple with the MaxDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDisk

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) SetMaxDisk(v int64)`

SetMaxDisk sets MaxDisk field to given value.

### HasMaxDisk

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) HasMaxDisk() bool`

HasMaxDisk returns a boolean if a field has been set.

### GetDateCreated

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddScaleThresholds200ResponseAllOfScaleThreshold) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


