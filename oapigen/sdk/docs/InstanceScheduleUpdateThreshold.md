# InstanceScheduleUpdateThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceThresholdId** | Pointer to **int64** | Source Scale Threshold to apply as a template. All threshold settings with be copied from this threshold, and can be overridden by also passing each setting explicitly. | [optional] 
**AutoUp** | Pointer to **bool** | Auto Upscale | [optional] 
**AutoDown** | Pointer to **bool** | Auto Downscale | [optional] 
**MinCount** | Pointer to **int32** | The minimum number of nodes to scale down to | [optional] 
**MaxCount** | Pointer to **int32** | The maximum number of nodes to scale up to | [optional] 
**CpuEnabled** | Pointer to **bool** | Enable CPU Threshold | [optional] 
**MinCpu** | Pointer to **float64** | Min CPU (%) | [optional] 
**MaxCpu** | Pointer to **float64** | Max CPU (%) | [optional] 
**MemoryEnabled** | Pointer to **bool** | Enable Memory Threshold | [optional] 
**MinMemory** | Pointer to **float64** | Min Memory (%) | [optional] 
**MaxMemory** | Pointer to **float64** | Max Memory (%) | [optional] 
**DiskEnabled** | Pointer to **bool** | Enable Disk Threshold | [optional] 
**MinDisk** | Pointer to **float64** | Min Disk (%) | [optional] 
**MaxDisk** | Pointer to **float64** | Max Disk (%) | [optional] 

## Methods

### NewInstanceScheduleUpdateThreshold

`func NewInstanceScheduleUpdateThreshold() *InstanceScheduleUpdateThreshold`

NewInstanceScheduleUpdateThreshold instantiates a new InstanceScheduleUpdateThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceScheduleUpdateThresholdWithDefaults

`func NewInstanceScheduleUpdateThresholdWithDefaults() *InstanceScheduleUpdateThreshold`

NewInstanceScheduleUpdateThresholdWithDefaults instantiates a new InstanceScheduleUpdateThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceThresholdId

`func (o *InstanceScheduleUpdateThreshold) GetSourceThresholdId() int64`

GetSourceThresholdId returns the SourceThresholdId field if non-nil, zero value otherwise.

### GetSourceThresholdIdOk

`func (o *InstanceScheduleUpdateThreshold) GetSourceThresholdIdOk() (*int64, bool)`

GetSourceThresholdIdOk returns a tuple with the SourceThresholdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceThresholdId

`func (o *InstanceScheduleUpdateThreshold) SetSourceThresholdId(v int64)`

SetSourceThresholdId sets SourceThresholdId field to given value.

### HasSourceThresholdId

`func (o *InstanceScheduleUpdateThreshold) HasSourceThresholdId() bool`

HasSourceThresholdId returns a boolean if a field has been set.

### GetAutoUp

`func (o *InstanceScheduleUpdateThreshold) GetAutoUp() bool`

GetAutoUp returns the AutoUp field if non-nil, zero value otherwise.

### GetAutoUpOk

`func (o *InstanceScheduleUpdateThreshold) GetAutoUpOk() (*bool, bool)`

GetAutoUpOk returns a tuple with the AutoUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUp

`func (o *InstanceScheduleUpdateThreshold) SetAutoUp(v bool)`

SetAutoUp sets AutoUp field to given value.

### HasAutoUp

`func (o *InstanceScheduleUpdateThreshold) HasAutoUp() bool`

HasAutoUp returns a boolean if a field has been set.

### GetAutoDown

`func (o *InstanceScheduleUpdateThreshold) GetAutoDown() bool`

GetAutoDown returns the AutoDown field if non-nil, zero value otherwise.

### GetAutoDownOk

`func (o *InstanceScheduleUpdateThreshold) GetAutoDownOk() (*bool, bool)`

GetAutoDownOk returns a tuple with the AutoDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDown

`func (o *InstanceScheduleUpdateThreshold) SetAutoDown(v bool)`

SetAutoDown sets AutoDown field to given value.

### HasAutoDown

`func (o *InstanceScheduleUpdateThreshold) HasAutoDown() bool`

HasAutoDown returns a boolean if a field has been set.

### GetMinCount

`func (o *InstanceScheduleUpdateThreshold) GetMinCount() int32`

GetMinCount returns the MinCount field if non-nil, zero value otherwise.

### GetMinCountOk

`func (o *InstanceScheduleUpdateThreshold) GetMinCountOk() (*int32, bool)`

GetMinCountOk returns a tuple with the MinCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCount

`func (o *InstanceScheduleUpdateThreshold) SetMinCount(v int32)`

SetMinCount sets MinCount field to given value.

### HasMinCount

`func (o *InstanceScheduleUpdateThreshold) HasMinCount() bool`

HasMinCount returns a boolean if a field has been set.

### GetMaxCount

`func (o *InstanceScheduleUpdateThreshold) GetMaxCount() int32`

GetMaxCount returns the MaxCount field if non-nil, zero value otherwise.

### GetMaxCountOk

`func (o *InstanceScheduleUpdateThreshold) GetMaxCountOk() (*int32, bool)`

GetMaxCountOk returns a tuple with the MaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCount

`func (o *InstanceScheduleUpdateThreshold) SetMaxCount(v int32)`

SetMaxCount sets MaxCount field to given value.

### HasMaxCount

`func (o *InstanceScheduleUpdateThreshold) HasMaxCount() bool`

HasMaxCount returns a boolean if a field has been set.

### GetCpuEnabled

`func (o *InstanceScheduleUpdateThreshold) GetCpuEnabled() bool`

GetCpuEnabled returns the CpuEnabled field if non-nil, zero value otherwise.

### GetCpuEnabledOk

`func (o *InstanceScheduleUpdateThreshold) GetCpuEnabledOk() (*bool, bool)`

GetCpuEnabledOk returns a tuple with the CpuEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuEnabled

`func (o *InstanceScheduleUpdateThreshold) SetCpuEnabled(v bool)`

SetCpuEnabled sets CpuEnabled field to given value.

### HasCpuEnabled

`func (o *InstanceScheduleUpdateThreshold) HasCpuEnabled() bool`

HasCpuEnabled returns a boolean if a field has been set.

### GetMinCpu

`func (o *InstanceScheduleUpdateThreshold) GetMinCpu() float64`

GetMinCpu returns the MinCpu field if non-nil, zero value otherwise.

### GetMinCpuOk

`func (o *InstanceScheduleUpdateThreshold) GetMinCpuOk() (*float64, bool)`

GetMinCpuOk returns a tuple with the MinCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCpu

`func (o *InstanceScheduleUpdateThreshold) SetMinCpu(v float64)`

SetMinCpu sets MinCpu field to given value.

### HasMinCpu

`func (o *InstanceScheduleUpdateThreshold) HasMinCpu() bool`

HasMinCpu returns a boolean if a field has been set.

### GetMaxCpu

`func (o *InstanceScheduleUpdateThreshold) GetMaxCpu() float64`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *InstanceScheduleUpdateThreshold) GetMaxCpuOk() (*float64, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *InstanceScheduleUpdateThreshold) SetMaxCpu(v float64)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *InstanceScheduleUpdateThreshold) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetMemoryEnabled

`func (o *InstanceScheduleUpdateThreshold) GetMemoryEnabled() bool`

GetMemoryEnabled returns the MemoryEnabled field if non-nil, zero value otherwise.

### GetMemoryEnabledOk

`func (o *InstanceScheduleUpdateThreshold) GetMemoryEnabledOk() (*bool, bool)`

GetMemoryEnabledOk returns a tuple with the MemoryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryEnabled

`func (o *InstanceScheduleUpdateThreshold) SetMemoryEnabled(v bool)`

SetMemoryEnabled sets MemoryEnabled field to given value.

### HasMemoryEnabled

`func (o *InstanceScheduleUpdateThreshold) HasMemoryEnabled() bool`

HasMemoryEnabled returns a boolean if a field has been set.

### GetMinMemory

`func (o *InstanceScheduleUpdateThreshold) GetMinMemory() float64`

GetMinMemory returns the MinMemory field if non-nil, zero value otherwise.

### GetMinMemoryOk

`func (o *InstanceScheduleUpdateThreshold) GetMinMemoryOk() (*float64, bool)`

GetMinMemoryOk returns a tuple with the MinMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMemory

`func (o *InstanceScheduleUpdateThreshold) SetMinMemory(v float64)`

SetMinMemory sets MinMemory field to given value.

### HasMinMemory

`func (o *InstanceScheduleUpdateThreshold) HasMinMemory() bool`

HasMinMemory returns a boolean if a field has been set.

### GetMaxMemory

`func (o *InstanceScheduleUpdateThreshold) GetMaxMemory() float64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *InstanceScheduleUpdateThreshold) GetMaxMemoryOk() (*float64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *InstanceScheduleUpdateThreshold) SetMaxMemory(v float64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *InstanceScheduleUpdateThreshold) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetDiskEnabled

`func (o *InstanceScheduleUpdateThreshold) GetDiskEnabled() bool`

GetDiskEnabled returns the DiskEnabled field if non-nil, zero value otherwise.

### GetDiskEnabledOk

`func (o *InstanceScheduleUpdateThreshold) GetDiskEnabledOk() (*bool, bool)`

GetDiskEnabledOk returns a tuple with the DiskEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEnabled

`func (o *InstanceScheduleUpdateThreshold) SetDiskEnabled(v bool)`

SetDiskEnabled sets DiskEnabled field to given value.

### HasDiskEnabled

`func (o *InstanceScheduleUpdateThreshold) HasDiskEnabled() bool`

HasDiskEnabled returns a boolean if a field has been set.

### GetMinDisk

`func (o *InstanceScheduleUpdateThreshold) GetMinDisk() float64`

GetMinDisk returns the MinDisk field if non-nil, zero value otherwise.

### GetMinDiskOk

`func (o *InstanceScheduleUpdateThreshold) GetMinDiskOk() (*float64, bool)`

GetMinDiskOk returns a tuple with the MinDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDisk

`func (o *InstanceScheduleUpdateThreshold) SetMinDisk(v float64)`

SetMinDisk sets MinDisk field to given value.

### HasMinDisk

`func (o *InstanceScheduleUpdateThreshold) HasMinDisk() bool`

HasMinDisk returns a boolean if a field has been set.

### GetMaxDisk

`func (o *InstanceScheduleUpdateThreshold) GetMaxDisk() float64`

GetMaxDisk returns the MaxDisk field if non-nil, zero value otherwise.

### GetMaxDiskOk

`func (o *InstanceScheduleUpdateThreshold) GetMaxDiskOk() (*float64, bool)`

GetMaxDiskOk returns a tuple with the MaxDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDisk

`func (o *InstanceScheduleUpdateThreshold) SetMaxDisk(v float64)`

SetMaxDisk sets MaxDisk field to given value.

### HasMaxDisk

`func (o *InstanceScheduleUpdateThreshold) HasMaxDisk() bool`

HasMaxDisk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


