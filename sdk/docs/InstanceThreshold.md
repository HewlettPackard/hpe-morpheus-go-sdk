# InstanceThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
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
**MinNetwork** | Pointer to **string** |  | [optional] 
**NetworkEnabled** | Pointer to **bool** |  | [optional] 
**IopsEnabled** | Pointer to **bool** |  | [optional] 
**MinIops** | Pointer to **string** |  | [optional] 
**MaxIops** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**ZoneId** | Pointer to **int64** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewInstanceThreshold

`func NewInstanceThreshold() *InstanceThreshold`

NewInstanceThreshold instantiates a new InstanceThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceThresholdWithDefaults

`func NewInstanceThresholdWithDefaults() *InstanceThreshold`

NewInstanceThresholdWithDefaults instantiates a new InstanceThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceThreshold) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceThreshold) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceThreshold) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceThreshold) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAutoUp

`func (o *InstanceThreshold) GetAutoUp() bool`

GetAutoUp returns the AutoUp field if non-nil, zero value otherwise.

### GetAutoUpOk

`func (o *InstanceThreshold) GetAutoUpOk() (*bool, bool)`

GetAutoUpOk returns a tuple with the AutoUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUp

`func (o *InstanceThreshold) SetAutoUp(v bool)`

SetAutoUp sets AutoUp field to given value.

### HasAutoUp

`func (o *InstanceThreshold) HasAutoUp() bool`

HasAutoUp returns a boolean if a field has been set.

### GetAutoDown

`func (o *InstanceThreshold) GetAutoDown() bool`

GetAutoDown returns the AutoDown field if non-nil, zero value otherwise.

### GetAutoDownOk

`func (o *InstanceThreshold) GetAutoDownOk() (*bool, bool)`

GetAutoDownOk returns a tuple with the AutoDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDown

`func (o *InstanceThreshold) SetAutoDown(v bool)`

SetAutoDown sets AutoDown field to given value.

### HasAutoDown

`func (o *InstanceThreshold) HasAutoDown() bool`

HasAutoDown returns a boolean if a field has been set.

### GetMinCount

`func (o *InstanceThreshold) GetMinCount() int64`

GetMinCount returns the MinCount field if non-nil, zero value otherwise.

### GetMinCountOk

`func (o *InstanceThreshold) GetMinCountOk() (*int64, bool)`

GetMinCountOk returns a tuple with the MinCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCount

`func (o *InstanceThreshold) SetMinCount(v int64)`

SetMinCount sets MinCount field to given value.

### HasMinCount

`func (o *InstanceThreshold) HasMinCount() bool`

HasMinCount returns a boolean if a field has been set.

### GetMaxCount

`func (o *InstanceThreshold) GetMaxCount() int64`

GetMaxCount returns the MaxCount field if non-nil, zero value otherwise.

### GetMaxCountOk

`func (o *InstanceThreshold) GetMaxCountOk() (*int64, bool)`

GetMaxCountOk returns a tuple with the MaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCount

`func (o *InstanceThreshold) SetMaxCount(v int64)`

SetMaxCount sets MaxCount field to given value.

### HasMaxCount

`func (o *InstanceThreshold) HasMaxCount() bool`

HasMaxCount returns a boolean if a field has been set.

### GetScaleIncrement

`func (o *InstanceThreshold) GetScaleIncrement() int64`

GetScaleIncrement returns the ScaleIncrement field if non-nil, zero value otherwise.

### GetScaleIncrementOk

`func (o *InstanceThreshold) GetScaleIncrementOk() (*int64, bool)`

GetScaleIncrementOk returns a tuple with the ScaleIncrement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleIncrement

`func (o *InstanceThreshold) SetScaleIncrement(v int64)`

SetScaleIncrement sets ScaleIncrement field to given value.

### HasScaleIncrement

`func (o *InstanceThreshold) HasScaleIncrement() bool`

HasScaleIncrement returns a boolean if a field has been set.

### GetCpuEnabled

`func (o *InstanceThreshold) GetCpuEnabled() bool`

GetCpuEnabled returns the CpuEnabled field if non-nil, zero value otherwise.

### GetCpuEnabledOk

`func (o *InstanceThreshold) GetCpuEnabledOk() (*bool, bool)`

GetCpuEnabledOk returns a tuple with the CpuEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuEnabled

`func (o *InstanceThreshold) SetCpuEnabled(v bool)`

SetCpuEnabled sets CpuEnabled field to given value.

### HasCpuEnabled

`func (o *InstanceThreshold) HasCpuEnabled() bool`

HasCpuEnabled returns a boolean if a field has been set.

### GetMinCpu

`func (o *InstanceThreshold) GetMinCpu() int64`

GetMinCpu returns the MinCpu field if non-nil, zero value otherwise.

### GetMinCpuOk

`func (o *InstanceThreshold) GetMinCpuOk() (*int64, bool)`

GetMinCpuOk returns a tuple with the MinCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCpu

`func (o *InstanceThreshold) SetMinCpu(v int64)`

SetMinCpu sets MinCpu field to given value.

### HasMinCpu

`func (o *InstanceThreshold) HasMinCpu() bool`

HasMinCpu returns a boolean if a field has been set.

### GetMaxCpu

`func (o *InstanceThreshold) GetMaxCpu() int64`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *InstanceThreshold) GetMaxCpuOk() (*int64, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *InstanceThreshold) SetMaxCpu(v int64)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *InstanceThreshold) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetMemoryEnabled

`func (o *InstanceThreshold) GetMemoryEnabled() bool`

GetMemoryEnabled returns the MemoryEnabled field if non-nil, zero value otherwise.

### GetMemoryEnabledOk

`func (o *InstanceThreshold) GetMemoryEnabledOk() (*bool, bool)`

GetMemoryEnabledOk returns a tuple with the MemoryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryEnabled

`func (o *InstanceThreshold) SetMemoryEnabled(v bool)`

SetMemoryEnabled sets MemoryEnabled field to given value.

### HasMemoryEnabled

`func (o *InstanceThreshold) HasMemoryEnabled() bool`

HasMemoryEnabled returns a boolean if a field has been set.

### GetMinMemory

`func (o *InstanceThreshold) GetMinMemory() int64`

GetMinMemory returns the MinMemory field if non-nil, zero value otherwise.

### GetMinMemoryOk

`func (o *InstanceThreshold) GetMinMemoryOk() (*int64, bool)`

GetMinMemoryOk returns a tuple with the MinMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMemory

`func (o *InstanceThreshold) SetMinMemory(v int64)`

SetMinMemory sets MinMemory field to given value.

### HasMinMemory

`func (o *InstanceThreshold) HasMinMemory() bool`

HasMinMemory returns a boolean if a field has been set.

### GetMaxMemory

`func (o *InstanceThreshold) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *InstanceThreshold) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *InstanceThreshold) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *InstanceThreshold) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetDiskEnabled

`func (o *InstanceThreshold) GetDiskEnabled() bool`

GetDiskEnabled returns the DiskEnabled field if non-nil, zero value otherwise.

### GetDiskEnabledOk

`func (o *InstanceThreshold) GetDiskEnabledOk() (*bool, bool)`

GetDiskEnabledOk returns a tuple with the DiskEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEnabled

`func (o *InstanceThreshold) SetDiskEnabled(v bool)`

SetDiskEnabled sets DiskEnabled field to given value.

### HasDiskEnabled

`func (o *InstanceThreshold) HasDiskEnabled() bool`

HasDiskEnabled returns a boolean if a field has been set.

### GetMinDisk

`func (o *InstanceThreshold) GetMinDisk() int64`

GetMinDisk returns the MinDisk field if non-nil, zero value otherwise.

### GetMinDiskOk

`func (o *InstanceThreshold) GetMinDiskOk() (*int64, bool)`

GetMinDiskOk returns a tuple with the MinDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDisk

`func (o *InstanceThreshold) SetMinDisk(v int64)`

SetMinDisk sets MinDisk field to given value.

### HasMinDisk

`func (o *InstanceThreshold) HasMinDisk() bool`

HasMinDisk returns a boolean if a field has been set.

### GetMaxDisk

`func (o *InstanceThreshold) GetMaxDisk() int64`

GetMaxDisk returns the MaxDisk field if non-nil, zero value otherwise.

### GetMaxDiskOk

`func (o *InstanceThreshold) GetMaxDiskOk() (*int64, bool)`

GetMaxDiskOk returns a tuple with the MaxDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDisk

`func (o *InstanceThreshold) SetMaxDisk(v int64)`

SetMaxDisk sets MaxDisk field to given value.

### HasMaxDisk

`func (o *InstanceThreshold) HasMaxDisk() bool`

HasMaxDisk returns a boolean if a field has been set.

### GetMinNetwork

`func (o *InstanceThreshold) GetMinNetwork() string`

GetMinNetwork returns the MinNetwork field if non-nil, zero value otherwise.

### GetMinNetworkOk

`func (o *InstanceThreshold) GetMinNetworkOk() (*string, bool)`

GetMinNetworkOk returns a tuple with the MinNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNetwork

`func (o *InstanceThreshold) SetMinNetwork(v string)`

SetMinNetwork sets MinNetwork field to given value.

### HasMinNetwork

`func (o *InstanceThreshold) HasMinNetwork() bool`

HasMinNetwork returns a boolean if a field has been set.

### GetNetworkEnabled

`func (o *InstanceThreshold) GetNetworkEnabled() bool`

GetNetworkEnabled returns the NetworkEnabled field if non-nil, zero value otherwise.

### GetNetworkEnabledOk

`func (o *InstanceThreshold) GetNetworkEnabledOk() (*bool, bool)`

GetNetworkEnabledOk returns a tuple with the NetworkEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEnabled

`func (o *InstanceThreshold) SetNetworkEnabled(v bool)`

SetNetworkEnabled sets NetworkEnabled field to given value.

### HasNetworkEnabled

`func (o *InstanceThreshold) HasNetworkEnabled() bool`

HasNetworkEnabled returns a boolean if a field has been set.

### GetIopsEnabled

`func (o *InstanceThreshold) GetIopsEnabled() bool`

GetIopsEnabled returns the IopsEnabled field if non-nil, zero value otherwise.

### GetIopsEnabledOk

`func (o *InstanceThreshold) GetIopsEnabledOk() (*bool, bool)`

GetIopsEnabledOk returns a tuple with the IopsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsEnabled

`func (o *InstanceThreshold) SetIopsEnabled(v bool)`

SetIopsEnabled sets IopsEnabled field to given value.

### HasIopsEnabled

`func (o *InstanceThreshold) HasIopsEnabled() bool`

HasIopsEnabled returns a boolean if a field has been set.

### GetMinIops

`func (o *InstanceThreshold) GetMinIops() string`

GetMinIops returns the MinIops field if non-nil, zero value otherwise.

### GetMinIopsOk

`func (o *InstanceThreshold) GetMinIopsOk() (*string, bool)`

GetMinIopsOk returns a tuple with the MinIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinIops

`func (o *InstanceThreshold) SetMinIops(v string)`

SetMinIops sets MinIops field to given value.

### HasMinIops

`func (o *InstanceThreshold) HasMinIops() bool`

HasMinIops returns a boolean if a field has been set.

### GetMaxIops

`func (o *InstanceThreshold) GetMaxIops() string`

GetMaxIops returns the MaxIops field if non-nil, zero value otherwise.

### GetMaxIopsOk

`func (o *InstanceThreshold) GetMaxIopsOk() (*string, bool)`

GetMaxIopsOk returns a tuple with the MaxIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIops

`func (o *InstanceThreshold) SetMaxIops(v string)`

SetMaxIops sets MaxIops field to given value.

### HasMaxIops

`func (o *InstanceThreshold) HasMaxIops() bool`

HasMaxIops returns a boolean if a field has been set.

### GetComment

`func (o *InstanceThreshold) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *InstanceThreshold) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *InstanceThreshold) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *InstanceThreshold) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetZoneId

`func (o *InstanceThreshold) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *InstanceThreshold) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *InstanceThreshold) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *InstanceThreshold) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetDateCreated

`func (o *InstanceThreshold) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *InstanceThreshold) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *InstanceThreshold) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *InstanceThreshold) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *InstanceThreshold) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *InstanceThreshold) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *InstanceThreshold) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *InstanceThreshold) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


