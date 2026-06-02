# GetApp200ResponseAppStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsedMemory** | Pointer to **int64** |  | [optional] 
**MaxMemory** | Pointer to **int64** |  | [optional] 
**UsedStorage** | Pointer to **int64** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 
**Running** | Pointer to **int64** |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 
**CpuUsage** | Pointer to **float32** |  | [optional] 
**InstanceCount** | Pointer to **int64** |  | [optional] 
**InstanceDayCount** | Pointer to **[]int64** |  | [optional] 
**InstanceDayCountTotal** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetApp200ResponseAppStats

`func NewGetApp200ResponseAppStats() *GetApp200ResponseAppStats`

NewGetApp200ResponseAppStats instantiates a new GetApp200ResponseAppStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetUsedMemory

`func (o *GetApp200ResponseAppStats) GetUsedMemory() int64`

GetUsedMemory returns the UsedMemory field if non-nil, zero value otherwise.

### GetUsedMemoryOk

`func (o *GetApp200ResponseAppStats) GetUsedMemoryOk() (*int64, bool)`

GetUsedMemoryOk returns a tuple with the UsedMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedMemory

`func (o *GetApp200ResponseAppStats) SetUsedMemory(v int64)`

SetUsedMemory sets UsedMemory field to given value.

### HasUsedMemory

`func (o *GetApp200ResponseAppStats) HasUsedMemory() bool`

HasUsedMemory returns a boolean if a field has been set.

### GetMaxMemory

`func (o *GetApp200ResponseAppStats) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *GetApp200ResponseAppStats) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *GetApp200ResponseAppStats) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *GetApp200ResponseAppStats) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetUsedStorage

`func (o *GetApp200ResponseAppStats) GetUsedStorage() int64`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *GetApp200ResponseAppStats) GetUsedStorageOk() (*int64, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *GetApp200ResponseAppStats) SetUsedStorage(v int64)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *GetApp200ResponseAppStats) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### GetMaxStorage

`func (o *GetApp200ResponseAppStats) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *GetApp200ResponseAppStats) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *GetApp200ResponseAppStats) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *GetApp200ResponseAppStats) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetRunning

`func (o *GetApp200ResponseAppStats) GetRunning() int64`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *GetApp200ResponseAppStats) GetRunningOk() (*int64, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *GetApp200ResponseAppStats) SetRunning(v int64)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *GetApp200ResponseAppStats) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetTotal

`func (o *GetApp200ResponseAppStats) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetApp200ResponseAppStats) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetApp200ResponseAppStats) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetApp200ResponseAppStats) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetCpuUsage

`func (o *GetApp200ResponseAppStats) GetCpuUsage() float32`

GetCpuUsage returns the CpuUsage field if non-nil, zero value otherwise.

### GetCpuUsageOk

`func (o *GetApp200ResponseAppStats) GetCpuUsageOk() (*float32, bool)`

GetCpuUsageOk returns a tuple with the CpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsage

`func (o *GetApp200ResponseAppStats) SetCpuUsage(v float32)`

SetCpuUsage sets CpuUsage field to given value.

### HasCpuUsage

`func (o *GetApp200ResponseAppStats) HasCpuUsage() bool`

HasCpuUsage returns a boolean if a field has been set.

### GetInstanceCount

`func (o *GetApp200ResponseAppStats) GetInstanceCount() int64`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *GetApp200ResponseAppStats) GetInstanceCountOk() (*int64, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *GetApp200ResponseAppStats) SetInstanceCount(v int64)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *GetApp200ResponseAppStats) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

### GetInstanceDayCount

`func (o *GetApp200ResponseAppStats) GetInstanceDayCount() []int64`

GetInstanceDayCount returns the InstanceDayCount field if non-nil, zero value otherwise.

### GetInstanceDayCountOk

`func (o *GetApp200ResponseAppStats) GetInstanceDayCountOk() (*[]int64, bool)`

GetInstanceDayCountOk returns a tuple with the InstanceDayCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceDayCount

`func (o *GetApp200ResponseAppStats) SetInstanceDayCount(v []int64)`

SetInstanceDayCount sets InstanceDayCount field to given value.

### HasInstanceDayCount

`func (o *GetApp200ResponseAppStats) HasInstanceDayCount() bool`

HasInstanceDayCount returns a boolean if a field has been set.

### GetInstanceDayCountTotal

`func (o *GetApp200ResponseAppStats) GetInstanceDayCountTotal() int64`

GetInstanceDayCountTotal returns the InstanceDayCountTotal field if non-nil, zero value otherwise.

### GetInstanceDayCountTotalOk

`func (o *GetApp200ResponseAppStats) GetInstanceDayCountTotalOk() (*int64, bool)`

GetInstanceDayCountTotalOk returns a tuple with the InstanceDayCountTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceDayCountTotal

`func (o *GetApp200ResponseAppStats) SetInstanceDayCountTotal(v int64)`

SetInstanceDayCountTotal sets InstanceDayCountTotal field to given value.

### HasInstanceDayCountTotal

`func (o *GetApp200ResponseAppStats) HasInstanceDayCountTotal() bool`

HasInstanceDayCountTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


