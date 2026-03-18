# GetContainer200ResponseContainerStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **string** |  | [optional] 
**Running** | Pointer to **bool** |  | [optional] 
**UserCpuUsage** | Pointer to **float32** |  | [optional] 
**SystemCpuUsage** | Pointer to **float32** |  | [optional] 
**UsedMemory** | Pointer to **int32** |  | [optional] 
**MaxMemory** | Pointer to **int32** |  | [optional] 
**CacheMemory** | Pointer to **int32** |  | [optional] 
**MaxStorage** | Pointer to **int32** |  | [optional] 
**UsedStorage** | Pointer to **int32** |  | [optional] 
**ReadIOPS** | Pointer to **int32** |  | [optional] 
**WriteIOPS** | Pointer to **float32** |  | [optional] 
**TotalIOPS** | Pointer to **float32** |  | [optional] 
**NetTxUsage** | Pointer to **int32** |  | [optional] 
**NetRxUsage** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetContainer200ResponseContainerStats

`func NewGetContainer200ResponseContainerStats() *GetContainer200ResponseContainerStats`

NewGetContainer200ResponseContainerStats instantiates a new GetContainer200ResponseContainerStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContainer200ResponseContainerStatsWithDefaults

`func NewGetContainer200ResponseContainerStatsWithDefaults() *GetContainer200ResponseContainerStats`

NewGetContainer200ResponseContainerStatsWithDefaults instantiates a new GetContainer200ResponseContainerStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *GetContainer200ResponseContainerStats) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetContainer200ResponseContainerStats) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetContainer200ResponseContainerStats) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetContainer200ResponseContainerStats) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetRunning

`func (o *GetContainer200ResponseContainerStats) GetRunning() bool`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *GetContainer200ResponseContainerStats) GetRunningOk() (*bool, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *GetContainer200ResponseContainerStats) SetRunning(v bool)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *GetContainer200ResponseContainerStats) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetUserCpuUsage

`func (o *GetContainer200ResponseContainerStats) GetUserCpuUsage() float32`

GetUserCpuUsage returns the UserCpuUsage field if non-nil, zero value otherwise.

### GetUserCpuUsageOk

`func (o *GetContainer200ResponseContainerStats) GetUserCpuUsageOk() (*float32, bool)`

GetUserCpuUsageOk returns a tuple with the UserCpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCpuUsage

`func (o *GetContainer200ResponseContainerStats) SetUserCpuUsage(v float32)`

SetUserCpuUsage sets UserCpuUsage field to given value.

### HasUserCpuUsage

`func (o *GetContainer200ResponseContainerStats) HasUserCpuUsage() bool`

HasUserCpuUsage returns a boolean if a field has been set.

### GetSystemCpuUsage

`func (o *GetContainer200ResponseContainerStats) GetSystemCpuUsage() float32`

GetSystemCpuUsage returns the SystemCpuUsage field if non-nil, zero value otherwise.

### GetSystemCpuUsageOk

`func (o *GetContainer200ResponseContainerStats) GetSystemCpuUsageOk() (*float32, bool)`

GetSystemCpuUsageOk returns a tuple with the SystemCpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemCpuUsage

`func (o *GetContainer200ResponseContainerStats) SetSystemCpuUsage(v float32)`

SetSystemCpuUsage sets SystemCpuUsage field to given value.

### HasSystemCpuUsage

`func (o *GetContainer200ResponseContainerStats) HasSystemCpuUsage() bool`

HasSystemCpuUsage returns a boolean if a field has been set.

### GetUsedMemory

`func (o *GetContainer200ResponseContainerStats) GetUsedMemory() int32`

GetUsedMemory returns the UsedMemory field if non-nil, zero value otherwise.

### GetUsedMemoryOk

`func (o *GetContainer200ResponseContainerStats) GetUsedMemoryOk() (*int32, bool)`

GetUsedMemoryOk returns a tuple with the UsedMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedMemory

`func (o *GetContainer200ResponseContainerStats) SetUsedMemory(v int32)`

SetUsedMemory sets UsedMemory field to given value.

### HasUsedMemory

`func (o *GetContainer200ResponseContainerStats) HasUsedMemory() bool`

HasUsedMemory returns a boolean if a field has been set.

### GetMaxMemory

`func (o *GetContainer200ResponseContainerStats) GetMaxMemory() int32`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *GetContainer200ResponseContainerStats) GetMaxMemoryOk() (*int32, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *GetContainer200ResponseContainerStats) SetMaxMemory(v int32)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *GetContainer200ResponseContainerStats) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetCacheMemory

`func (o *GetContainer200ResponseContainerStats) GetCacheMemory() int32`

GetCacheMemory returns the CacheMemory field if non-nil, zero value otherwise.

### GetCacheMemoryOk

`func (o *GetContainer200ResponseContainerStats) GetCacheMemoryOk() (*int32, bool)`

GetCacheMemoryOk returns a tuple with the CacheMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheMemory

`func (o *GetContainer200ResponseContainerStats) SetCacheMemory(v int32)`

SetCacheMemory sets CacheMemory field to given value.

### HasCacheMemory

`func (o *GetContainer200ResponseContainerStats) HasCacheMemory() bool`

HasCacheMemory returns a boolean if a field has been set.

### GetMaxStorage

`func (o *GetContainer200ResponseContainerStats) GetMaxStorage() int32`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *GetContainer200ResponseContainerStats) GetMaxStorageOk() (*int32, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *GetContainer200ResponseContainerStats) SetMaxStorage(v int32)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *GetContainer200ResponseContainerStats) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetUsedStorage

`func (o *GetContainer200ResponseContainerStats) GetUsedStorage() int32`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *GetContainer200ResponseContainerStats) GetUsedStorageOk() (*int32, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *GetContainer200ResponseContainerStats) SetUsedStorage(v int32)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *GetContainer200ResponseContainerStats) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### GetReadIOPS

`func (o *GetContainer200ResponseContainerStats) GetReadIOPS() int32`

GetReadIOPS returns the ReadIOPS field if non-nil, zero value otherwise.

### GetReadIOPSOk

`func (o *GetContainer200ResponseContainerStats) GetReadIOPSOk() (*int32, bool)`

GetReadIOPSOk returns a tuple with the ReadIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadIOPS

`func (o *GetContainer200ResponseContainerStats) SetReadIOPS(v int32)`

SetReadIOPS sets ReadIOPS field to given value.

### HasReadIOPS

`func (o *GetContainer200ResponseContainerStats) HasReadIOPS() bool`

HasReadIOPS returns a boolean if a field has been set.

### GetWriteIOPS

`func (o *GetContainer200ResponseContainerStats) GetWriteIOPS() float32`

GetWriteIOPS returns the WriteIOPS field if non-nil, zero value otherwise.

### GetWriteIOPSOk

`func (o *GetContainer200ResponseContainerStats) GetWriteIOPSOk() (*float32, bool)`

GetWriteIOPSOk returns a tuple with the WriteIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteIOPS

`func (o *GetContainer200ResponseContainerStats) SetWriteIOPS(v float32)`

SetWriteIOPS sets WriteIOPS field to given value.

### HasWriteIOPS

`func (o *GetContainer200ResponseContainerStats) HasWriteIOPS() bool`

HasWriteIOPS returns a boolean if a field has been set.

### GetTotalIOPS

`func (o *GetContainer200ResponseContainerStats) GetTotalIOPS() float32`

GetTotalIOPS returns the TotalIOPS field if non-nil, zero value otherwise.

### GetTotalIOPSOk

`func (o *GetContainer200ResponseContainerStats) GetTotalIOPSOk() (*float32, bool)`

GetTotalIOPSOk returns a tuple with the TotalIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIOPS

`func (o *GetContainer200ResponseContainerStats) SetTotalIOPS(v float32)`

SetTotalIOPS sets TotalIOPS field to given value.

### HasTotalIOPS

`func (o *GetContainer200ResponseContainerStats) HasTotalIOPS() bool`

HasTotalIOPS returns a boolean if a field has been set.

### GetNetTxUsage

`func (o *GetContainer200ResponseContainerStats) GetNetTxUsage() int32`

GetNetTxUsage returns the NetTxUsage field if non-nil, zero value otherwise.

### GetNetTxUsageOk

`func (o *GetContainer200ResponseContainerStats) GetNetTxUsageOk() (*int32, bool)`

GetNetTxUsageOk returns a tuple with the NetTxUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetTxUsage

`func (o *GetContainer200ResponseContainerStats) SetNetTxUsage(v int32)`

SetNetTxUsage sets NetTxUsage field to given value.

### HasNetTxUsage

`func (o *GetContainer200ResponseContainerStats) HasNetTxUsage() bool`

HasNetTxUsage returns a boolean if a field has been set.

### GetNetRxUsage

`func (o *GetContainer200ResponseContainerStats) GetNetRxUsage() int32`

GetNetRxUsage returns the NetRxUsage field if non-nil, zero value otherwise.

### GetNetRxUsageOk

`func (o *GetContainer200ResponseContainerStats) GetNetRxUsageOk() (*int32, bool)`

GetNetRxUsageOk returns a tuple with the NetRxUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetRxUsage

`func (o *GetContainer200ResponseContainerStats) SetNetRxUsage(v int32)`

SetNetRxUsage sets NetRxUsage field to given value.

### HasNetRxUsage

`func (o *GetContainer200ResponseContainerStats) HasNetRxUsage() bool`

HasNetRxUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


