# GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoUp** | Pointer to **bool** | Auto Upscale | [optional] [default to false]
**AutoDown** | Pointer to **bool** | Auto Downscale | [optional] [default to false]
**MinCount** | Pointer to **int32** | The minimum number of nodes to scale down to | [optional] 
**MaxCount** | Pointer to **int32** | The maximum number of nodes to scale up to | [optional] 
**ScaleIncrement** | Pointer to **int64** | The number of nodes that are added or removed at one time when scaling up or down | [optional] [default to 1]
**CpuEnabled** | Pointer to **bool** | Enable CPU Threshold | [optional] [default to false]
**MinCpu** | Pointer to **float64** | Min CPU (%) | [optional] [default to 0]
**MaxCpu** | Pointer to **float64** | Max CPU (%) | [optional] [default to 0]
**MemoryEnabled** | Pointer to **bool** | Enable Memory Threshold | [optional] [default to false]
**MinMemory** | Pointer to **float64** | Min Memory (%) | [optional] [default to 0]
**MaxMemory** | Pointer to **float64** | Max Memory (%) | [optional] [default to 0]
**DiskEnabled** | Pointer to **bool** | Enable Disk Threshold | [optional] [default to false]
**MinDisk** | Pointer to **float64** | Min Disk (%) | [optional] [default to 0]
**MaxDisk** | Pointer to **float64** | Max Disk (%) | [optional] [default to 0]

## Methods

### NewGetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold

`func NewGetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold() *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold`

NewGetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold instantiates a new GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInstanceThreshold200ResponseInstanceSchedulesInnerThresholdWithDefaults

`func NewGetInstanceThreshold200ResponseInstanceSchedulesInnerThresholdWithDefaults() *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold`

NewGetInstanceThreshold200ResponseInstanceSchedulesInnerThresholdWithDefaults instantiates a new GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoUp

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) GetAutoUp() bool`

GetAutoUp returns the AutoUp field if non-nil, zero value otherwise.

### GetAutoUpOk

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) GetAutoUpOk() (*bool, bool)`

GetAutoUpOk returns a tuple with the AutoUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUp

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) SetAutoUp(v bool)`

SetAutoUp sets AutoUp field to given value.

### HasAutoUp

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) HasAutoUp() bool`

HasAutoUp returns a boolean if a field has been set.

### GetAutoDown

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) GetAutoDown() bool`

GetAutoDown returns the AutoDown field if non-nil, zero value otherwise.

### GetAutoDownOk

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) GetAutoDownOk() (*bool, bool)`

GetAutoDownOk returns a tuple with the AutoDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDown

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) SetAutoDown(v bool)`

SetAutoDown sets AutoDown field to given value.

### HasAutoDown

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) HasAutoDown() bool`

HasAutoDown returns a boolean if a field has been set.

### GetMinCount

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) GetMinCount() int32`

GetMinCount returns the MinCount field if non-nil, zero value otherwise.

### GetMinCountOk

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) GetMinCountOk() (*int32, bool)`

GetMinCountOk returns a tuple with the MinCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCount

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) SetMinCount(v int32)`

SetMinCount sets MinCount field to given value.

### HasMinCount

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) HasMinCount() bool`

HasMinCount returns a boolean if a field has been set.

### GetMaxCount

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) GetMaxCount() int32`

GetMaxCount returns the MaxCount field if non-nil, zero value otherwise.

### GetMaxCountOk

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) GetMaxCountOk() (*int32, bool)`

GetMaxCountOk returns a tuple with the MaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCount

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) SetMaxCount(v int32)`

SetMaxCount sets MaxCount field to given value.

### HasMaxCount

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) HasMaxCount() bool`

HasMaxCount returns a boolean if a field has been set.

### GetScaleIncrement

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) GetScaleIncrement() int64`

GetScaleIncrement returns the ScaleIncrement field if non-nil, zero value otherwise.

### GetScaleIncrementOk

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) GetScaleIncrementOk() (*int64, bool)`

GetScaleIncrementOk returns a tuple with the ScaleIncrement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleIncrement

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) SetScaleIncrement(v int64)`

SetScaleIncrement sets ScaleIncrement field to given value.

### HasScaleIncrement

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) HasScaleIncrement() bool`

HasScaleIncrement returns a boolean if a field has been set.

### GetCpuEnabled

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) GetCpuEnabled() bool`

GetCpuEnabled returns the CpuEnabled field if non-nil, zero value otherwise.

### GetCpuEnabledOk

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) GetCpuEnabledOk() (*bool, bool)`

GetCpuEnabledOk returns a tuple with the CpuEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuEnabled

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) SetCpuEnabled(v bool)`

SetCpuEnabled sets CpuEnabled field to given value.

### HasCpuEnabled

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) HasCpuEnabled() bool`

HasCpuEnabled returns a boolean if a field has been set.

### GetMinCpu

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) GetMinCpu() float64`

GetMinCpu returns the MinCpu field if non-nil, zero value otherwise.

### GetMinCpuOk

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) GetMinCpuOk() (*float64, bool)`

GetMinCpuOk returns a tuple with the MinCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCpu

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) SetMinCpu(v float64)`

SetMinCpu sets MinCpu field to given value.

### HasMinCpu

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) HasMinCpu() bool`

HasMinCpu returns a boolean if a field has been set.

### GetMaxCpu

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) GetMaxCpu() float64`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) GetMaxCpuOk() (*float64, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) SetMaxCpu(v float64)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetMemoryEnabled

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) GetMemoryEnabled() bool`

GetMemoryEnabled returns the MemoryEnabled field if non-nil, zero value otherwise.

### GetMemoryEnabledOk

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) GetMemoryEnabledOk() (*bool, bool)`

GetMemoryEnabledOk returns a tuple with the MemoryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryEnabled

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) SetMemoryEnabled(v bool)`

SetMemoryEnabled sets MemoryEnabled field to given value.

### HasMemoryEnabled

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) HasMemoryEnabled() bool`

HasMemoryEnabled returns a boolean if a field has been set.

### GetMinMemory

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) GetMinMemory() float64`

GetMinMemory returns the MinMemory field if non-nil, zero value otherwise.

### GetMinMemoryOk

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) GetMinMemoryOk() (*float64, bool)`

GetMinMemoryOk returns a tuple with the MinMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMemory

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) SetMinMemory(v float64)`

SetMinMemory sets MinMemory field to given value.

### HasMinMemory

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) HasMinMemory() bool`

HasMinMemory returns a boolean if a field has been set.

### GetMaxMemory

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) GetMaxMemory() float64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) GetMaxMemoryOk() (*float64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) SetMaxMemory(v float64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetDiskEnabled

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) GetDiskEnabled() bool`

GetDiskEnabled returns the DiskEnabled field if non-nil, zero value otherwise.

### GetDiskEnabledOk

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) GetDiskEnabledOk() (*bool, bool)`

GetDiskEnabledOk returns a tuple with the DiskEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEnabled

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) SetDiskEnabled(v bool)`

SetDiskEnabled sets DiskEnabled field to given value.

### HasDiskEnabled

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) HasDiskEnabled() bool`

HasDiskEnabled returns a boolean if a field has been set.

### GetMinDisk

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) GetMinDisk() float64`

GetMinDisk returns the MinDisk field if non-nil, zero value otherwise.

### GetMinDiskOk

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) GetMinDiskOk() (*float64, bool)`

GetMinDiskOk returns a tuple with the MinDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDisk

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) SetMinDisk(v float64)`

SetMinDisk sets MinDisk field to given value.

### HasMinDisk

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) HasMinDisk() bool`

HasMinDisk returns a boolean if a field has been set.

### GetMaxDisk

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) GetMaxDisk() float64`

GetMaxDisk returns the MaxDisk field if non-nil, zero value otherwise.

### GetMaxDiskOk

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) GetMaxDiskOk() (*float64, bool)`

GetMaxDiskOk returns a tuple with the MaxDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDisk

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) SetMaxDisk(v float64)`

SetMaxDisk sets MaxDisk field to given value.

### HasMaxDisk

`func (o *GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold) HasMaxDisk() bool`

HasMaxDisk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


