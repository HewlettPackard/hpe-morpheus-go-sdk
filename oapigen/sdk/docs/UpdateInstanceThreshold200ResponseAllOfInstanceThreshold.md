# UpdateInstanceThreshold200ResponseAllOfInstanceThreshold

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
**MinNetwork** | Pointer to **NullableString** |  | [optional] 
**NetworkEnabled** | Pointer to **bool** |  | [optional] 
**IopsEnabled** | Pointer to **bool** |  | [optional] 
**MinIops** | Pointer to **NullableString** |  | [optional] 
**MaxIops** | Pointer to **NullableString** |  | [optional] 
**Comment** | Pointer to **NullableString** |  | [optional] 
**ZoneId** | Pointer to **NullableInt64** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUpdateInstanceThreshold200ResponseAllOfInstanceThreshold

`func NewUpdateInstanceThreshold200ResponseAllOfInstanceThreshold() *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold`

NewUpdateInstanceThreshold200ResponseAllOfInstanceThreshold instantiates a new UpdateInstanceThreshold200ResponseAllOfInstanceThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAutoUp

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetAutoUp() bool`

GetAutoUp returns the AutoUp field if non-nil, zero value otherwise.

### GetAutoUpOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetAutoUpOk() (*bool, bool)`

GetAutoUpOk returns a tuple with the AutoUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUp

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) SetAutoUp(v bool)`

SetAutoUp sets AutoUp field to given value.

### HasAutoUp

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) HasAutoUp() bool`

HasAutoUp returns a boolean if a field has been set.

### GetAutoDown

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetAutoDown() bool`

GetAutoDown returns the AutoDown field if non-nil, zero value otherwise.

### GetAutoDownOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetAutoDownOk() (*bool, bool)`

GetAutoDownOk returns a tuple with the AutoDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDown

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) SetAutoDown(v bool)`

SetAutoDown sets AutoDown field to given value.

### HasAutoDown

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) HasAutoDown() bool`

HasAutoDown returns a boolean if a field has been set.

### GetMinCount

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetMinCount() int64`

GetMinCount returns the MinCount field if non-nil, zero value otherwise.

### GetMinCountOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetMinCountOk() (*int64, bool)`

GetMinCountOk returns a tuple with the MinCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCount

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) SetMinCount(v int64)`

SetMinCount sets MinCount field to given value.

### HasMinCount

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) HasMinCount() bool`

HasMinCount returns a boolean if a field has been set.

### GetMaxCount

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetMaxCount() int64`

GetMaxCount returns the MaxCount field if non-nil, zero value otherwise.

### GetMaxCountOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetMaxCountOk() (*int64, bool)`

GetMaxCountOk returns a tuple with the MaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCount

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) SetMaxCount(v int64)`

SetMaxCount sets MaxCount field to given value.

### HasMaxCount

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) HasMaxCount() bool`

HasMaxCount returns a boolean if a field has been set.

### GetScaleIncrement

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetScaleIncrement() int64`

GetScaleIncrement returns the ScaleIncrement field if non-nil, zero value otherwise.

### GetScaleIncrementOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetScaleIncrementOk() (*int64, bool)`

GetScaleIncrementOk returns a tuple with the ScaleIncrement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleIncrement

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) SetScaleIncrement(v int64)`

SetScaleIncrement sets ScaleIncrement field to given value.

### HasScaleIncrement

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) HasScaleIncrement() bool`

HasScaleIncrement returns a boolean if a field has been set.

### GetCpuEnabled

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetCpuEnabled() bool`

GetCpuEnabled returns the CpuEnabled field if non-nil, zero value otherwise.

### GetCpuEnabledOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetCpuEnabledOk() (*bool, bool)`

GetCpuEnabledOk returns a tuple with the CpuEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuEnabled

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) SetCpuEnabled(v bool)`

SetCpuEnabled sets CpuEnabled field to given value.

### HasCpuEnabled

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) HasCpuEnabled() bool`

HasCpuEnabled returns a boolean if a field has been set.

### GetMinCpu

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetMinCpu() int64`

GetMinCpu returns the MinCpu field if non-nil, zero value otherwise.

### GetMinCpuOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetMinCpuOk() (*int64, bool)`

GetMinCpuOk returns a tuple with the MinCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCpu

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) SetMinCpu(v int64)`

SetMinCpu sets MinCpu field to given value.

### HasMinCpu

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) HasMinCpu() bool`

HasMinCpu returns a boolean if a field has been set.

### GetMaxCpu

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetMaxCpu() int64`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetMaxCpuOk() (*int64, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) SetMaxCpu(v int64)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetMemoryEnabled

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetMemoryEnabled() bool`

GetMemoryEnabled returns the MemoryEnabled field if non-nil, zero value otherwise.

### GetMemoryEnabledOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetMemoryEnabledOk() (*bool, bool)`

GetMemoryEnabledOk returns a tuple with the MemoryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryEnabled

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) SetMemoryEnabled(v bool)`

SetMemoryEnabled sets MemoryEnabled field to given value.

### HasMemoryEnabled

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) HasMemoryEnabled() bool`

HasMemoryEnabled returns a boolean if a field has been set.

### GetMinMemory

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetMinMemory() int64`

GetMinMemory returns the MinMemory field if non-nil, zero value otherwise.

### GetMinMemoryOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetMinMemoryOk() (*int64, bool)`

GetMinMemoryOk returns a tuple with the MinMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMemory

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) SetMinMemory(v int64)`

SetMinMemory sets MinMemory field to given value.

### HasMinMemory

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) HasMinMemory() bool`

HasMinMemory returns a boolean if a field has been set.

### GetMaxMemory

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetDiskEnabled

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetDiskEnabled() bool`

GetDiskEnabled returns the DiskEnabled field if non-nil, zero value otherwise.

### GetDiskEnabledOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetDiskEnabledOk() (*bool, bool)`

GetDiskEnabledOk returns a tuple with the DiskEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEnabled

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) SetDiskEnabled(v bool)`

SetDiskEnabled sets DiskEnabled field to given value.

### HasDiskEnabled

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) HasDiskEnabled() bool`

HasDiskEnabled returns a boolean if a field has been set.

### GetMinDisk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetMinDisk() int64`

GetMinDisk returns the MinDisk field if non-nil, zero value otherwise.

### GetMinDiskOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetMinDiskOk() (*int64, bool)`

GetMinDiskOk returns a tuple with the MinDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDisk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) SetMinDisk(v int64)`

SetMinDisk sets MinDisk field to given value.

### HasMinDisk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) HasMinDisk() bool`

HasMinDisk returns a boolean if a field has been set.

### GetMaxDisk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetMaxDisk() int64`

GetMaxDisk returns the MaxDisk field if non-nil, zero value otherwise.

### GetMaxDiskOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetMaxDiskOk() (*int64, bool)`

GetMaxDiskOk returns a tuple with the MaxDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDisk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) SetMaxDisk(v int64)`

SetMaxDisk sets MaxDisk field to given value.

### HasMaxDisk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) HasMaxDisk() bool`

HasMaxDisk returns a boolean if a field has been set.

### GetMinNetwork

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetMinNetwork() string`

GetMinNetwork returns the MinNetwork field if non-nil, zero value otherwise.

### GetMinNetworkOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetMinNetworkOk() (*string, bool)`

GetMinNetworkOk returns a tuple with the MinNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNetwork

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) SetMinNetwork(v string)`

SetMinNetwork sets MinNetwork field to given value.

### HasMinNetwork

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) HasMinNetwork() bool`

HasMinNetwork returns a boolean if a field has been set.

### SetMinNetworkNil

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) SetMinNetworkNil(b bool)`

 SetMinNetworkNil sets the value for MinNetwork to be an explicit nil

### UnsetMinNetwork
`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) UnsetMinNetwork()`

UnsetMinNetwork ensures that no value is present for MinNetwork, not even an explicit nil
### GetNetworkEnabled

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetNetworkEnabled() bool`

GetNetworkEnabled returns the NetworkEnabled field if non-nil, zero value otherwise.

### GetNetworkEnabledOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetNetworkEnabledOk() (*bool, bool)`

GetNetworkEnabledOk returns a tuple with the NetworkEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEnabled

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) SetNetworkEnabled(v bool)`

SetNetworkEnabled sets NetworkEnabled field to given value.

### HasNetworkEnabled

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) HasNetworkEnabled() bool`

HasNetworkEnabled returns a boolean if a field has been set.

### GetIopsEnabled

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetIopsEnabled() bool`

GetIopsEnabled returns the IopsEnabled field if non-nil, zero value otherwise.

### GetIopsEnabledOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetIopsEnabledOk() (*bool, bool)`

GetIopsEnabledOk returns a tuple with the IopsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsEnabled

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) SetIopsEnabled(v bool)`

SetIopsEnabled sets IopsEnabled field to given value.

### HasIopsEnabled

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) HasIopsEnabled() bool`

HasIopsEnabled returns a boolean if a field has been set.

### GetMinIops

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetMinIops() string`

GetMinIops returns the MinIops field if non-nil, zero value otherwise.

### GetMinIopsOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetMinIopsOk() (*string, bool)`

GetMinIopsOk returns a tuple with the MinIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinIops

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) SetMinIops(v string)`

SetMinIops sets MinIops field to given value.

### HasMinIops

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) HasMinIops() bool`

HasMinIops returns a boolean if a field has been set.

### SetMinIopsNil

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) SetMinIopsNil(b bool)`

 SetMinIopsNil sets the value for MinIops to be an explicit nil

### UnsetMinIops
`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) UnsetMinIops()`

UnsetMinIops ensures that no value is present for MinIops, not even an explicit nil
### GetMaxIops

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetMaxIops() string`

GetMaxIops returns the MaxIops field if non-nil, zero value otherwise.

### GetMaxIopsOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetMaxIopsOk() (*string, bool)`

GetMaxIopsOk returns a tuple with the MaxIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIops

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) SetMaxIops(v string)`

SetMaxIops sets MaxIops field to given value.

### HasMaxIops

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) HasMaxIops() bool`

HasMaxIops returns a boolean if a field has been set.

### SetMaxIopsNil

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) SetMaxIopsNil(b bool)`

 SetMaxIopsNil sets the value for MaxIops to be an explicit nil

### UnsetMaxIops
`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) UnsetMaxIops()`

UnsetMaxIops ensures that no value is present for MaxIops, not even an explicit nil
### GetComment

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetZoneId

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### SetZoneIdNil

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) SetZoneIdNil(b bool)`

 SetZoneIdNil sets the value for ZoneId to be an explicit nil

### UnsetZoneId
`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) UnsetZoneId()`

UnsetZoneId ensures that no value is present for ZoneId, not even an explicit nil
### GetDateCreated

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceThreshold) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


