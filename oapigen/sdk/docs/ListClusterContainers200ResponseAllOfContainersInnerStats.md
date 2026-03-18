# ListClusterContainers200ResponseAllOfContainersInnerStats

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

### NewListClusterContainers200ResponseAllOfContainersInnerStats

`func NewListClusterContainers200ResponseAllOfContainersInnerStats() *ListClusterContainers200ResponseAllOfContainersInnerStats`

NewListClusterContainers200ResponseAllOfContainersInnerStats instantiates a new ListClusterContainers200ResponseAllOfContainersInnerStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterContainers200ResponseAllOfContainersInnerStatsWithDefaults

`func NewListClusterContainers200ResponseAllOfContainersInnerStatsWithDefaults() *ListClusterContainers200ResponseAllOfContainersInnerStats`

NewListClusterContainers200ResponseAllOfContainersInnerStatsWithDefaults instantiates a new ListClusterContainers200ResponseAllOfContainersInnerStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetRunning

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) GetRunning() bool`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) GetRunningOk() (*bool, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) SetRunning(v bool)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetUserCpuUsage

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) GetUserCpuUsage() float32`

GetUserCpuUsage returns the UserCpuUsage field if non-nil, zero value otherwise.

### GetUserCpuUsageOk

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) GetUserCpuUsageOk() (*float32, bool)`

GetUserCpuUsageOk returns a tuple with the UserCpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCpuUsage

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) SetUserCpuUsage(v float32)`

SetUserCpuUsage sets UserCpuUsage field to given value.

### HasUserCpuUsage

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) HasUserCpuUsage() bool`

HasUserCpuUsage returns a boolean if a field has been set.

### GetSystemCpuUsage

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) GetSystemCpuUsage() float32`

GetSystemCpuUsage returns the SystemCpuUsage field if non-nil, zero value otherwise.

### GetSystemCpuUsageOk

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) GetSystemCpuUsageOk() (*float32, bool)`

GetSystemCpuUsageOk returns a tuple with the SystemCpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemCpuUsage

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) SetSystemCpuUsage(v float32)`

SetSystemCpuUsage sets SystemCpuUsage field to given value.

### HasSystemCpuUsage

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) HasSystemCpuUsage() bool`

HasSystemCpuUsage returns a boolean if a field has been set.

### GetUsedMemory

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) GetUsedMemory() int64`

GetUsedMemory returns the UsedMemory field if non-nil, zero value otherwise.

### GetUsedMemoryOk

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) GetUsedMemoryOk() (*int64, bool)`

GetUsedMemoryOk returns a tuple with the UsedMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedMemory

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) SetUsedMemory(v int64)`

SetUsedMemory sets UsedMemory field to given value.

### HasUsedMemory

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) HasUsedMemory() bool`

HasUsedMemory returns a boolean if a field has been set.

### GetMaxMemory

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetCacheMemory

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) GetCacheMemory() int64`

GetCacheMemory returns the CacheMemory field if non-nil, zero value otherwise.

### GetCacheMemoryOk

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) GetCacheMemoryOk() (*int64, bool)`

GetCacheMemoryOk returns a tuple with the CacheMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheMemory

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) SetCacheMemory(v int64)`

SetCacheMemory sets CacheMemory field to given value.

### HasCacheMemory

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) HasCacheMemory() bool`

HasCacheMemory returns a boolean if a field has been set.

### GetMaxStorage

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetUsedStorage

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) GetUsedStorage() int64`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) GetUsedStorageOk() (*int64, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) SetUsedStorage(v int64)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### GetReadIOPS

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) GetReadIOPS() int64`

GetReadIOPS returns the ReadIOPS field if non-nil, zero value otherwise.

### GetReadIOPSOk

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) GetReadIOPSOk() (*int64, bool)`

GetReadIOPSOk returns a tuple with the ReadIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadIOPS

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) SetReadIOPS(v int64)`

SetReadIOPS sets ReadIOPS field to given value.

### HasReadIOPS

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) HasReadIOPS() bool`

HasReadIOPS returns a boolean if a field has been set.

### GetWriteIOPS

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) GetWriteIOPS() int64`

GetWriteIOPS returns the WriteIOPS field if non-nil, zero value otherwise.

### GetWriteIOPSOk

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) GetWriteIOPSOk() (*int64, bool)`

GetWriteIOPSOk returns a tuple with the WriteIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteIOPS

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) SetWriteIOPS(v int64)`

SetWriteIOPS sets WriteIOPS field to given value.

### HasWriteIOPS

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) HasWriteIOPS() bool`

HasWriteIOPS returns a boolean if a field has been set.

### GetTotalIOPS

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) GetTotalIOPS() int64`

GetTotalIOPS returns the TotalIOPS field if non-nil, zero value otherwise.

### GetTotalIOPSOk

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) GetTotalIOPSOk() (*int64, bool)`

GetTotalIOPSOk returns a tuple with the TotalIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIOPS

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) SetTotalIOPS(v int64)`

SetTotalIOPS sets TotalIOPS field to given value.

### HasTotalIOPS

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) HasTotalIOPS() bool`

HasTotalIOPS returns a boolean if a field has been set.

### GetNetTxUsage

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) GetNetTxUsage() int64`

GetNetTxUsage returns the NetTxUsage field if non-nil, zero value otherwise.

### GetNetTxUsageOk

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) GetNetTxUsageOk() (*int64, bool)`

GetNetTxUsageOk returns a tuple with the NetTxUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetTxUsage

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) SetNetTxUsage(v int64)`

SetNetTxUsage sets NetTxUsage field to given value.

### HasNetTxUsage

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) HasNetTxUsage() bool`

HasNetTxUsage returns a boolean if a field has been set.

### GetNetRxUsage

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) GetNetRxUsage() int64`

GetNetRxUsage returns the NetRxUsage field if non-nil, zero value otherwise.

### GetNetRxUsageOk

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) GetNetRxUsageOk() (*int64, bool)`

GetNetRxUsageOk returns a tuple with the NetRxUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetRxUsage

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) SetNetRxUsage(v int64)`

SetNetRxUsage sets NetRxUsage field to given value.

### HasNetRxUsage

`func (o *ListClusterContainers200ResponseAllOfContainersInnerStats) HasNetRxUsage() bool`

HasNetRxUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


