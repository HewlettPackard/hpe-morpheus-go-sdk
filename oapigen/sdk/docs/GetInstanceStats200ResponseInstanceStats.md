# GetInstanceStats200ResponseInstanceStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxCpu** | Pointer to **float32** | Total CPUs for all instances | [optional] 
**MaxCores** | Pointer to **float32** | Total cores for all instances | [optional] 
**CpuUsageAverage** | Pointer to **float32** | Average current CPU Usage across all instances | [optional] 
**CpuUsagePeak** | Pointer to **float32** | Max current CPU Usage across all instances | [optional] 
**UsedMemory** | Pointer to **float32** | Total used memory across all instances | [optional] 
**MaxMemory** | Pointer to **float32** | Total memory across all instances | [optional] 
**UsedStorage** | Pointer to **float32** | Total used storage total across all instances | [optional] 
**MaxStorage** | Pointer to **float32** | Total storage across all instances | [optional] 
**Running** | Pointer to **float32** | Total number of running instances | [optional] 
**Total** | Pointer to **float32** | Total number of instances | [optional] 
**TotalContainers** | Pointer to **float32** | Total number of containers across all instances | [optional] 

## Methods

### NewGetInstanceStats200ResponseInstanceStats

`func NewGetInstanceStats200ResponseInstanceStats() *GetInstanceStats200ResponseInstanceStats`

NewGetInstanceStats200ResponseInstanceStats instantiates a new GetInstanceStats200ResponseInstanceStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInstanceStats200ResponseInstanceStatsWithDefaults

`func NewGetInstanceStats200ResponseInstanceStatsWithDefaults() *GetInstanceStats200ResponseInstanceStats`

NewGetInstanceStats200ResponseInstanceStatsWithDefaults instantiates a new GetInstanceStats200ResponseInstanceStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxCpu

`func (o *GetInstanceStats200ResponseInstanceStats) GetMaxCpu() float32`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *GetInstanceStats200ResponseInstanceStats) GetMaxCpuOk() (*float32, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *GetInstanceStats200ResponseInstanceStats) SetMaxCpu(v float32)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *GetInstanceStats200ResponseInstanceStats) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetMaxCores

`func (o *GetInstanceStats200ResponseInstanceStats) GetMaxCores() float32`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *GetInstanceStats200ResponseInstanceStats) GetMaxCoresOk() (*float32, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *GetInstanceStats200ResponseInstanceStats) SetMaxCores(v float32)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *GetInstanceStats200ResponseInstanceStats) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetCpuUsageAverage

`func (o *GetInstanceStats200ResponseInstanceStats) GetCpuUsageAverage() float32`

GetCpuUsageAverage returns the CpuUsageAverage field if non-nil, zero value otherwise.

### GetCpuUsageAverageOk

`func (o *GetInstanceStats200ResponseInstanceStats) GetCpuUsageAverageOk() (*float32, bool)`

GetCpuUsageAverageOk returns a tuple with the CpuUsageAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsageAverage

`func (o *GetInstanceStats200ResponseInstanceStats) SetCpuUsageAverage(v float32)`

SetCpuUsageAverage sets CpuUsageAverage field to given value.

### HasCpuUsageAverage

`func (o *GetInstanceStats200ResponseInstanceStats) HasCpuUsageAverage() bool`

HasCpuUsageAverage returns a boolean if a field has been set.

### GetCpuUsagePeak

`func (o *GetInstanceStats200ResponseInstanceStats) GetCpuUsagePeak() float32`

GetCpuUsagePeak returns the CpuUsagePeak field if non-nil, zero value otherwise.

### GetCpuUsagePeakOk

`func (o *GetInstanceStats200ResponseInstanceStats) GetCpuUsagePeakOk() (*float32, bool)`

GetCpuUsagePeakOk returns a tuple with the CpuUsagePeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsagePeak

`func (o *GetInstanceStats200ResponseInstanceStats) SetCpuUsagePeak(v float32)`

SetCpuUsagePeak sets CpuUsagePeak field to given value.

### HasCpuUsagePeak

`func (o *GetInstanceStats200ResponseInstanceStats) HasCpuUsagePeak() bool`

HasCpuUsagePeak returns a boolean if a field has been set.

### GetUsedMemory

`func (o *GetInstanceStats200ResponseInstanceStats) GetUsedMemory() float32`

GetUsedMemory returns the UsedMemory field if non-nil, zero value otherwise.

### GetUsedMemoryOk

`func (o *GetInstanceStats200ResponseInstanceStats) GetUsedMemoryOk() (*float32, bool)`

GetUsedMemoryOk returns a tuple with the UsedMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedMemory

`func (o *GetInstanceStats200ResponseInstanceStats) SetUsedMemory(v float32)`

SetUsedMemory sets UsedMemory field to given value.

### HasUsedMemory

`func (o *GetInstanceStats200ResponseInstanceStats) HasUsedMemory() bool`

HasUsedMemory returns a boolean if a field has been set.

### GetMaxMemory

`func (o *GetInstanceStats200ResponseInstanceStats) GetMaxMemory() float32`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *GetInstanceStats200ResponseInstanceStats) GetMaxMemoryOk() (*float32, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *GetInstanceStats200ResponseInstanceStats) SetMaxMemory(v float32)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *GetInstanceStats200ResponseInstanceStats) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetUsedStorage

`func (o *GetInstanceStats200ResponseInstanceStats) GetUsedStorage() float32`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *GetInstanceStats200ResponseInstanceStats) GetUsedStorageOk() (*float32, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *GetInstanceStats200ResponseInstanceStats) SetUsedStorage(v float32)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *GetInstanceStats200ResponseInstanceStats) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### GetMaxStorage

`func (o *GetInstanceStats200ResponseInstanceStats) GetMaxStorage() float32`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *GetInstanceStats200ResponseInstanceStats) GetMaxStorageOk() (*float32, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *GetInstanceStats200ResponseInstanceStats) SetMaxStorage(v float32)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *GetInstanceStats200ResponseInstanceStats) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetRunning

`func (o *GetInstanceStats200ResponseInstanceStats) GetRunning() float32`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *GetInstanceStats200ResponseInstanceStats) GetRunningOk() (*float32, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *GetInstanceStats200ResponseInstanceStats) SetRunning(v float32)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *GetInstanceStats200ResponseInstanceStats) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetTotal

`func (o *GetInstanceStats200ResponseInstanceStats) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetInstanceStats200ResponseInstanceStats) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetInstanceStats200ResponseInstanceStats) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetInstanceStats200ResponseInstanceStats) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTotalContainers

`func (o *GetInstanceStats200ResponseInstanceStats) GetTotalContainers() float32`

GetTotalContainers returns the TotalContainers field if non-nil, zero value otherwise.

### GetTotalContainersOk

`func (o *GetInstanceStats200ResponseInstanceStats) GetTotalContainersOk() (*float32, bool)`

GetTotalContainersOk returns a tuple with the TotalContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalContainers

`func (o *GetInstanceStats200ResponseInstanceStats) SetTotalContainers(v float32)`

SetTotalContainers sets TotalContainers field to given value.

### HasTotalContainers

`func (o *GetInstanceStats200ResponseInstanceStats) HasTotalContainers() bool`

HasTotalContainers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


