# GetInstanceThreshold200ResponseInstanceThreshold

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

### NewGetInstanceThreshold200ResponseInstanceThreshold

`func NewGetInstanceThreshold200ResponseInstanceThreshold() *GetInstanceThreshold200ResponseInstanceThreshold`

NewGetInstanceThreshold200ResponseInstanceThreshold instantiates a new GetInstanceThreshold200ResponseInstanceThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInstanceThreshold200ResponseInstanceThresholdWithDefaults

`func NewGetInstanceThreshold200ResponseInstanceThresholdWithDefaults() *GetInstanceThreshold200ResponseInstanceThreshold`

NewGetInstanceThreshold200ResponseInstanceThresholdWithDefaults instantiates a new GetInstanceThreshold200ResponseInstanceThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAutoUp

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetAutoUp() bool`

GetAutoUp returns the AutoUp field if non-nil, zero value otherwise.

### GetAutoUpOk

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetAutoUpOk() (*bool, bool)`

GetAutoUpOk returns a tuple with the AutoUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUp

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) SetAutoUp(v bool)`

SetAutoUp sets AutoUp field to given value.

### HasAutoUp

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) HasAutoUp() bool`

HasAutoUp returns a boolean if a field has been set.

### GetAutoDown

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetAutoDown() bool`

GetAutoDown returns the AutoDown field if non-nil, zero value otherwise.

### GetAutoDownOk

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetAutoDownOk() (*bool, bool)`

GetAutoDownOk returns a tuple with the AutoDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDown

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) SetAutoDown(v bool)`

SetAutoDown sets AutoDown field to given value.

### HasAutoDown

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) HasAutoDown() bool`

HasAutoDown returns a boolean if a field has been set.

### GetMinCount

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetMinCount() int64`

GetMinCount returns the MinCount field if non-nil, zero value otherwise.

### GetMinCountOk

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetMinCountOk() (*int64, bool)`

GetMinCountOk returns a tuple with the MinCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCount

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) SetMinCount(v int64)`

SetMinCount sets MinCount field to given value.

### HasMinCount

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) HasMinCount() bool`

HasMinCount returns a boolean if a field has been set.

### GetMaxCount

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetMaxCount() int64`

GetMaxCount returns the MaxCount field if non-nil, zero value otherwise.

### GetMaxCountOk

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetMaxCountOk() (*int64, bool)`

GetMaxCountOk returns a tuple with the MaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCount

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) SetMaxCount(v int64)`

SetMaxCount sets MaxCount field to given value.

### HasMaxCount

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) HasMaxCount() bool`

HasMaxCount returns a boolean if a field has been set.

### GetScaleIncrement

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetScaleIncrement() int64`

GetScaleIncrement returns the ScaleIncrement field if non-nil, zero value otherwise.

### GetScaleIncrementOk

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetScaleIncrementOk() (*int64, bool)`

GetScaleIncrementOk returns a tuple with the ScaleIncrement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleIncrement

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) SetScaleIncrement(v int64)`

SetScaleIncrement sets ScaleIncrement field to given value.

### HasScaleIncrement

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) HasScaleIncrement() bool`

HasScaleIncrement returns a boolean if a field has been set.

### GetCpuEnabled

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetCpuEnabled() bool`

GetCpuEnabled returns the CpuEnabled field if non-nil, zero value otherwise.

### GetCpuEnabledOk

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetCpuEnabledOk() (*bool, bool)`

GetCpuEnabledOk returns a tuple with the CpuEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuEnabled

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) SetCpuEnabled(v bool)`

SetCpuEnabled sets CpuEnabled field to given value.

### HasCpuEnabled

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) HasCpuEnabled() bool`

HasCpuEnabled returns a boolean if a field has been set.

### GetMinCpu

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetMinCpu() int64`

GetMinCpu returns the MinCpu field if non-nil, zero value otherwise.

### GetMinCpuOk

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetMinCpuOk() (*int64, bool)`

GetMinCpuOk returns a tuple with the MinCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCpu

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) SetMinCpu(v int64)`

SetMinCpu sets MinCpu field to given value.

### HasMinCpu

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) HasMinCpu() bool`

HasMinCpu returns a boolean if a field has been set.

### GetMaxCpu

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetMaxCpu() int64`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetMaxCpuOk() (*int64, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) SetMaxCpu(v int64)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetMemoryEnabled

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetMemoryEnabled() bool`

GetMemoryEnabled returns the MemoryEnabled field if non-nil, zero value otherwise.

### GetMemoryEnabledOk

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetMemoryEnabledOk() (*bool, bool)`

GetMemoryEnabledOk returns a tuple with the MemoryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryEnabled

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) SetMemoryEnabled(v bool)`

SetMemoryEnabled sets MemoryEnabled field to given value.

### HasMemoryEnabled

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) HasMemoryEnabled() bool`

HasMemoryEnabled returns a boolean if a field has been set.

### GetMinMemory

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetMinMemory() int64`

GetMinMemory returns the MinMemory field if non-nil, zero value otherwise.

### GetMinMemoryOk

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetMinMemoryOk() (*int64, bool)`

GetMinMemoryOk returns a tuple with the MinMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMemory

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) SetMinMemory(v int64)`

SetMinMemory sets MinMemory field to given value.

### HasMinMemory

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) HasMinMemory() bool`

HasMinMemory returns a boolean if a field has been set.

### GetMaxMemory

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetDiskEnabled

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetDiskEnabled() bool`

GetDiskEnabled returns the DiskEnabled field if non-nil, zero value otherwise.

### GetDiskEnabledOk

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetDiskEnabledOk() (*bool, bool)`

GetDiskEnabledOk returns a tuple with the DiskEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEnabled

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) SetDiskEnabled(v bool)`

SetDiskEnabled sets DiskEnabled field to given value.

### HasDiskEnabled

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) HasDiskEnabled() bool`

HasDiskEnabled returns a boolean if a field has been set.

### GetMinDisk

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetMinDisk() int64`

GetMinDisk returns the MinDisk field if non-nil, zero value otherwise.

### GetMinDiskOk

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetMinDiskOk() (*int64, bool)`

GetMinDiskOk returns a tuple with the MinDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDisk

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) SetMinDisk(v int64)`

SetMinDisk sets MinDisk field to given value.

### HasMinDisk

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) HasMinDisk() bool`

HasMinDisk returns a boolean if a field has been set.

### GetMaxDisk

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetMaxDisk() int64`

GetMaxDisk returns the MaxDisk field if non-nil, zero value otherwise.

### GetMaxDiskOk

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetMaxDiskOk() (*int64, bool)`

GetMaxDiskOk returns a tuple with the MaxDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDisk

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) SetMaxDisk(v int64)`

SetMaxDisk sets MaxDisk field to given value.

### HasMaxDisk

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) HasMaxDisk() bool`

HasMaxDisk returns a boolean if a field has been set.

### GetMinNetwork

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetMinNetwork() string`

GetMinNetwork returns the MinNetwork field if non-nil, zero value otherwise.

### GetMinNetworkOk

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetMinNetworkOk() (*string, bool)`

GetMinNetworkOk returns a tuple with the MinNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNetwork

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) SetMinNetwork(v string)`

SetMinNetwork sets MinNetwork field to given value.

### HasMinNetwork

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) HasMinNetwork() bool`

HasMinNetwork returns a boolean if a field has been set.

### GetNetworkEnabled

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetNetworkEnabled() bool`

GetNetworkEnabled returns the NetworkEnabled field if non-nil, zero value otherwise.

### GetNetworkEnabledOk

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetNetworkEnabledOk() (*bool, bool)`

GetNetworkEnabledOk returns a tuple with the NetworkEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEnabled

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) SetNetworkEnabled(v bool)`

SetNetworkEnabled sets NetworkEnabled field to given value.

### HasNetworkEnabled

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) HasNetworkEnabled() bool`

HasNetworkEnabled returns a boolean if a field has been set.

### GetIopsEnabled

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetIopsEnabled() bool`

GetIopsEnabled returns the IopsEnabled field if non-nil, zero value otherwise.

### GetIopsEnabledOk

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetIopsEnabledOk() (*bool, bool)`

GetIopsEnabledOk returns a tuple with the IopsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsEnabled

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) SetIopsEnabled(v bool)`

SetIopsEnabled sets IopsEnabled field to given value.

### HasIopsEnabled

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) HasIopsEnabled() bool`

HasIopsEnabled returns a boolean if a field has been set.

### GetMinIops

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetMinIops() string`

GetMinIops returns the MinIops field if non-nil, zero value otherwise.

### GetMinIopsOk

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetMinIopsOk() (*string, bool)`

GetMinIopsOk returns a tuple with the MinIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinIops

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) SetMinIops(v string)`

SetMinIops sets MinIops field to given value.

### HasMinIops

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) HasMinIops() bool`

HasMinIops returns a boolean if a field has been set.

### GetMaxIops

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetMaxIops() string`

GetMaxIops returns the MaxIops field if non-nil, zero value otherwise.

### GetMaxIopsOk

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetMaxIopsOk() (*string, bool)`

GetMaxIopsOk returns a tuple with the MaxIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIops

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) SetMaxIops(v string)`

SetMaxIops sets MaxIops field to given value.

### HasMaxIops

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) HasMaxIops() bool`

HasMaxIops returns a boolean if a field has been set.

### GetComment

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetZoneId

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetInstanceThreshold200ResponseInstanceThreshold) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


