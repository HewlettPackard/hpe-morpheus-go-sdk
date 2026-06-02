# AppStatsStats

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

### NewAppStatsStats

`func NewAppStatsStats() *AppStatsStats`

NewAppStatsStats instantiates a new AppStatsStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetUsedMemory

`func (o *AppStatsStats) GetUsedMemory() int64`

GetUsedMemory returns the UsedMemory field if non-nil, zero value otherwise.

### GetUsedMemoryOk

`func (o *AppStatsStats) GetUsedMemoryOk() (*int64, bool)`

GetUsedMemoryOk returns a tuple with the UsedMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedMemory

`func (o *AppStatsStats) SetUsedMemory(v int64)`

SetUsedMemory sets UsedMemory field to given value.

### HasUsedMemory

`func (o *AppStatsStats) HasUsedMemory() bool`

HasUsedMemory returns a boolean if a field has been set.

### GetMaxMemory

`func (o *AppStatsStats) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *AppStatsStats) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *AppStatsStats) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *AppStatsStats) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetUsedStorage

`func (o *AppStatsStats) GetUsedStorage() int64`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *AppStatsStats) GetUsedStorageOk() (*int64, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *AppStatsStats) SetUsedStorage(v int64)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *AppStatsStats) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### GetMaxStorage

`func (o *AppStatsStats) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *AppStatsStats) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *AppStatsStats) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *AppStatsStats) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetRunning

`func (o *AppStatsStats) GetRunning() int64`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *AppStatsStats) GetRunningOk() (*int64, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *AppStatsStats) SetRunning(v int64)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *AppStatsStats) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetTotal

`func (o *AppStatsStats) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AppStatsStats) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AppStatsStats) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AppStatsStats) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetCpuUsage

`func (o *AppStatsStats) GetCpuUsage() float32`

GetCpuUsage returns the CpuUsage field if non-nil, zero value otherwise.

### GetCpuUsageOk

`func (o *AppStatsStats) GetCpuUsageOk() (*float32, bool)`

GetCpuUsageOk returns a tuple with the CpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsage

`func (o *AppStatsStats) SetCpuUsage(v float32)`

SetCpuUsage sets CpuUsage field to given value.

### HasCpuUsage

`func (o *AppStatsStats) HasCpuUsage() bool`

HasCpuUsage returns a boolean if a field has been set.

### GetInstanceCount

`func (o *AppStatsStats) GetInstanceCount() int64`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *AppStatsStats) GetInstanceCountOk() (*int64, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *AppStatsStats) SetInstanceCount(v int64)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *AppStatsStats) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

### GetInstanceDayCount

`func (o *AppStatsStats) GetInstanceDayCount() []int64`

GetInstanceDayCount returns the InstanceDayCount field if non-nil, zero value otherwise.

### GetInstanceDayCountOk

`func (o *AppStatsStats) GetInstanceDayCountOk() (*[]int64, bool)`

GetInstanceDayCountOk returns a tuple with the InstanceDayCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceDayCount

`func (o *AppStatsStats) SetInstanceDayCount(v []int64)`

SetInstanceDayCount sets InstanceDayCount field to given value.

### HasInstanceDayCount

`func (o *AppStatsStats) HasInstanceDayCount() bool`

HasInstanceDayCount returns a boolean if a field has been set.

### GetInstanceDayCountTotal

`func (o *AppStatsStats) GetInstanceDayCountTotal() int64`

GetInstanceDayCountTotal returns the InstanceDayCountTotal field if non-nil, zero value otherwise.

### GetInstanceDayCountTotalOk

`func (o *AppStatsStats) GetInstanceDayCountTotalOk() (*int64, bool)`

GetInstanceDayCountTotalOk returns a tuple with the InstanceDayCountTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceDayCountTotal

`func (o *AppStatsStats) SetInstanceDayCountTotal(v int64)`

SetInstanceDayCountTotal sets InstanceDayCountTotal field to given value.

### HasInstanceDayCountTotal

`func (o *AppStatsStats) HasInstanceDayCountTotal() bool`

HasInstanceDayCountTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


