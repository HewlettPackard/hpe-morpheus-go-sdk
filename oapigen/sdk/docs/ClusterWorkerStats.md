# ClusterWorkerStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsedStorage** | Pointer to **int64** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 
**UsedMemory** | Pointer to **int64** |  | [optional] 
**MaxMemory** | Pointer to **int64** |  | [optional] 
**UsedCpu** | Pointer to **int64** |  | [optional] 
**CpuUsage** | Pointer to **float32** |  | [optional] 
**CpuUsagePeak** | Pointer to **float32** |  | [optional] 
**CpuUsageAvg** | Pointer to **float32** |  | [optional] 

## Methods

### NewClusterWorkerStats

`func NewClusterWorkerStats() *ClusterWorkerStats`

NewClusterWorkerStats instantiates a new ClusterWorkerStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWorkerStatsWithDefaults

`func NewClusterWorkerStatsWithDefaults() *ClusterWorkerStats`

NewClusterWorkerStatsWithDefaults instantiates a new ClusterWorkerStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsedStorage

`func (o *ClusterWorkerStats) GetUsedStorage() int64`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *ClusterWorkerStats) GetUsedStorageOk() (*int64, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *ClusterWorkerStats) SetUsedStorage(v int64)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *ClusterWorkerStats) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### GetMaxStorage

`func (o *ClusterWorkerStats) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ClusterWorkerStats) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ClusterWorkerStats) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *ClusterWorkerStats) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetUsedMemory

`func (o *ClusterWorkerStats) GetUsedMemory() int64`

GetUsedMemory returns the UsedMemory field if non-nil, zero value otherwise.

### GetUsedMemoryOk

`func (o *ClusterWorkerStats) GetUsedMemoryOk() (*int64, bool)`

GetUsedMemoryOk returns a tuple with the UsedMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedMemory

`func (o *ClusterWorkerStats) SetUsedMemory(v int64)`

SetUsedMemory sets UsedMemory field to given value.

### HasUsedMemory

`func (o *ClusterWorkerStats) HasUsedMemory() bool`

HasUsedMemory returns a boolean if a field has been set.

### GetMaxMemory

`func (o *ClusterWorkerStats) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *ClusterWorkerStats) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *ClusterWorkerStats) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *ClusterWorkerStats) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetUsedCpu

`func (o *ClusterWorkerStats) GetUsedCpu() int64`

GetUsedCpu returns the UsedCpu field if non-nil, zero value otherwise.

### GetUsedCpuOk

`func (o *ClusterWorkerStats) GetUsedCpuOk() (*int64, bool)`

GetUsedCpuOk returns a tuple with the UsedCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCpu

`func (o *ClusterWorkerStats) SetUsedCpu(v int64)`

SetUsedCpu sets UsedCpu field to given value.

### HasUsedCpu

`func (o *ClusterWorkerStats) HasUsedCpu() bool`

HasUsedCpu returns a boolean if a field has been set.

### GetCpuUsage

`func (o *ClusterWorkerStats) GetCpuUsage() float32`

GetCpuUsage returns the CpuUsage field if non-nil, zero value otherwise.

### GetCpuUsageOk

`func (o *ClusterWorkerStats) GetCpuUsageOk() (*float32, bool)`

GetCpuUsageOk returns a tuple with the CpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsage

`func (o *ClusterWorkerStats) SetCpuUsage(v float32)`

SetCpuUsage sets CpuUsage field to given value.

### HasCpuUsage

`func (o *ClusterWorkerStats) HasCpuUsage() bool`

HasCpuUsage returns a boolean if a field has been set.

### GetCpuUsagePeak

`func (o *ClusterWorkerStats) GetCpuUsagePeak() float32`

GetCpuUsagePeak returns the CpuUsagePeak field if non-nil, zero value otherwise.

### GetCpuUsagePeakOk

`func (o *ClusterWorkerStats) GetCpuUsagePeakOk() (*float32, bool)`

GetCpuUsagePeakOk returns a tuple with the CpuUsagePeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsagePeak

`func (o *ClusterWorkerStats) SetCpuUsagePeak(v float32)`

SetCpuUsagePeak sets CpuUsagePeak field to given value.

### HasCpuUsagePeak

`func (o *ClusterWorkerStats) HasCpuUsagePeak() bool`

HasCpuUsagePeak returns a boolean if a field has been set.

### GetCpuUsageAvg

`func (o *ClusterWorkerStats) GetCpuUsageAvg() float32`

GetCpuUsageAvg returns the CpuUsageAvg field if non-nil, zero value otherwise.

### GetCpuUsageAvgOk

`func (o *ClusterWorkerStats) GetCpuUsageAvgOk() (*float32, bool)`

GetCpuUsageAvgOk returns a tuple with the CpuUsageAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsageAvg

`func (o *ClusterWorkerStats) SetCpuUsageAvg(v float32)`

SetCpuUsageAvg sets CpuUsageAvg field to given value.

### HasCpuUsageAvg

`func (o *ClusterWorkerStats) HasCpuUsageAvg() bool`

HasCpuUsageAvg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


