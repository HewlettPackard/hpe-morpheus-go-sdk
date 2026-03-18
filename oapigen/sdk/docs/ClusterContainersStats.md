# ClusterContainersStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **time.Time** |  | [optional] 
**Running** | Pointer to **bool** |  | [optional] 
**UserCpuUsage** | Pointer to **float32** |  | [optional] 
**SystemCpuUsage** | Pointer to **float32** |  | [optional] 
**UsedMemory** | Pointer to **int64** |  | [optional] 
**MaxMemory** | Pointer to **int64** |  | [optional] 
**CacheMemory** | Pointer to **int64** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 
**UsedStorage** | Pointer to **int64** |  | [optional] 
**ReadIOPS** | Pointer to **int64** |  | [optional] 
**WriteIOPS** | Pointer to **int64** |  | [optional] 
**TotalIOPS** | Pointer to **int64** |  | [optional] 
**NetTxUsage** | Pointer to **int64** |  | [optional] 
**NetRxUsage** | Pointer to **int64** |  | [optional] 

## Methods

### NewClusterContainersStats

`func NewClusterContainersStats() *ClusterContainersStats`

NewClusterContainersStats instantiates a new ClusterContainersStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterContainersStatsWithDefaults

`func NewClusterContainersStatsWithDefaults() *ClusterContainersStats`

NewClusterContainersStatsWithDefaults instantiates a new ClusterContainersStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *ClusterContainersStats) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *ClusterContainersStats) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *ClusterContainersStats) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *ClusterContainersStats) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetRunning

`func (o *ClusterContainersStats) GetRunning() bool`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *ClusterContainersStats) GetRunningOk() (*bool, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *ClusterContainersStats) SetRunning(v bool)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *ClusterContainersStats) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetUserCpuUsage

`func (o *ClusterContainersStats) GetUserCpuUsage() float32`

GetUserCpuUsage returns the UserCpuUsage field if non-nil, zero value otherwise.

### GetUserCpuUsageOk

`func (o *ClusterContainersStats) GetUserCpuUsageOk() (*float32, bool)`

GetUserCpuUsageOk returns a tuple with the UserCpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCpuUsage

`func (o *ClusterContainersStats) SetUserCpuUsage(v float32)`

SetUserCpuUsage sets UserCpuUsage field to given value.

### HasUserCpuUsage

`func (o *ClusterContainersStats) HasUserCpuUsage() bool`

HasUserCpuUsage returns a boolean if a field has been set.

### GetSystemCpuUsage

`func (o *ClusterContainersStats) GetSystemCpuUsage() float32`

GetSystemCpuUsage returns the SystemCpuUsage field if non-nil, zero value otherwise.

### GetSystemCpuUsageOk

`func (o *ClusterContainersStats) GetSystemCpuUsageOk() (*float32, bool)`

GetSystemCpuUsageOk returns a tuple with the SystemCpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemCpuUsage

`func (o *ClusterContainersStats) SetSystemCpuUsage(v float32)`

SetSystemCpuUsage sets SystemCpuUsage field to given value.

### HasSystemCpuUsage

`func (o *ClusterContainersStats) HasSystemCpuUsage() bool`

HasSystemCpuUsage returns a boolean if a field has been set.

### GetUsedMemory

`func (o *ClusterContainersStats) GetUsedMemory() int64`

GetUsedMemory returns the UsedMemory field if non-nil, zero value otherwise.

### GetUsedMemoryOk

`func (o *ClusterContainersStats) GetUsedMemoryOk() (*int64, bool)`

GetUsedMemoryOk returns a tuple with the UsedMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedMemory

`func (o *ClusterContainersStats) SetUsedMemory(v int64)`

SetUsedMemory sets UsedMemory field to given value.

### HasUsedMemory

`func (o *ClusterContainersStats) HasUsedMemory() bool`

HasUsedMemory returns a boolean if a field has been set.

### GetMaxMemory

`func (o *ClusterContainersStats) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *ClusterContainersStats) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *ClusterContainersStats) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *ClusterContainersStats) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetCacheMemory

`func (o *ClusterContainersStats) GetCacheMemory() int64`

GetCacheMemory returns the CacheMemory field if non-nil, zero value otherwise.

### GetCacheMemoryOk

`func (o *ClusterContainersStats) GetCacheMemoryOk() (*int64, bool)`

GetCacheMemoryOk returns a tuple with the CacheMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheMemory

`func (o *ClusterContainersStats) SetCacheMemory(v int64)`

SetCacheMemory sets CacheMemory field to given value.

### HasCacheMemory

`func (o *ClusterContainersStats) HasCacheMemory() bool`

HasCacheMemory returns a boolean if a field has been set.

### GetMaxStorage

`func (o *ClusterContainersStats) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ClusterContainersStats) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ClusterContainersStats) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *ClusterContainersStats) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetUsedStorage

`func (o *ClusterContainersStats) GetUsedStorage() int64`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *ClusterContainersStats) GetUsedStorageOk() (*int64, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *ClusterContainersStats) SetUsedStorage(v int64)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *ClusterContainersStats) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### GetReadIOPS

`func (o *ClusterContainersStats) GetReadIOPS() int64`

GetReadIOPS returns the ReadIOPS field if non-nil, zero value otherwise.

### GetReadIOPSOk

`func (o *ClusterContainersStats) GetReadIOPSOk() (*int64, bool)`

GetReadIOPSOk returns a tuple with the ReadIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadIOPS

`func (o *ClusterContainersStats) SetReadIOPS(v int64)`

SetReadIOPS sets ReadIOPS field to given value.

### HasReadIOPS

`func (o *ClusterContainersStats) HasReadIOPS() bool`

HasReadIOPS returns a boolean if a field has been set.

### GetWriteIOPS

`func (o *ClusterContainersStats) GetWriteIOPS() int64`

GetWriteIOPS returns the WriteIOPS field if non-nil, zero value otherwise.

### GetWriteIOPSOk

`func (o *ClusterContainersStats) GetWriteIOPSOk() (*int64, bool)`

GetWriteIOPSOk returns a tuple with the WriteIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteIOPS

`func (o *ClusterContainersStats) SetWriteIOPS(v int64)`

SetWriteIOPS sets WriteIOPS field to given value.

### HasWriteIOPS

`func (o *ClusterContainersStats) HasWriteIOPS() bool`

HasWriteIOPS returns a boolean if a field has been set.

### GetTotalIOPS

`func (o *ClusterContainersStats) GetTotalIOPS() int64`

GetTotalIOPS returns the TotalIOPS field if non-nil, zero value otherwise.

### GetTotalIOPSOk

`func (o *ClusterContainersStats) GetTotalIOPSOk() (*int64, bool)`

GetTotalIOPSOk returns a tuple with the TotalIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIOPS

`func (o *ClusterContainersStats) SetTotalIOPS(v int64)`

SetTotalIOPS sets TotalIOPS field to given value.

### HasTotalIOPS

`func (o *ClusterContainersStats) HasTotalIOPS() bool`

HasTotalIOPS returns a boolean if a field has been set.

### GetNetTxUsage

`func (o *ClusterContainersStats) GetNetTxUsage() int64`

GetNetTxUsage returns the NetTxUsage field if non-nil, zero value otherwise.

### GetNetTxUsageOk

`func (o *ClusterContainersStats) GetNetTxUsageOk() (*int64, bool)`

GetNetTxUsageOk returns a tuple with the NetTxUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetTxUsage

`func (o *ClusterContainersStats) SetNetTxUsage(v int64)`

SetNetTxUsage sets NetTxUsage field to given value.

### HasNetTxUsage

`func (o *ClusterContainersStats) HasNetTxUsage() bool`

HasNetTxUsage returns a boolean if a field has been set.

### GetNetRxUsage

`func (o *ClusterContainersStats) GetNetRxUsage() int64`

GetNetRxUsage returns the NetRxUsage field if non-nil, zero value otherwise.

### GetNetRxUsageOk

`func (o *ClusterContainersStats) GetNetRxUsageOk() (*int64, bool)`

GetNetRxUsageOk returns a tuple with the NetRxUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetRxUsage

`func (o *ClusterContainersStats) SetNetRxUsage(v int64)`

SetNetRxUsage sets NetRxUsage field to given value.

### HasNetRxUsage

`func (o *ClusterContainersStats) HasNetRxUsage() bool`

HasNetRxUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


