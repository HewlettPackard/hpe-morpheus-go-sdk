# GuidanceVmwareSizingResourceStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **string** |  | [optional] 
**FreeMemory** | Pointer to **NullableInt64** |  | [optional] 
**UsedMemory** | Pointer to **NullableInt64** |  | [optional] 
**FreeSwap** | Pointer to **NullableInt64** |  | [optional] 
**UsedSwap** | Pointer to **NullableInt64** |  | [optional] 
**CpuIdleTime** | Pointer to **NullableInt64** |  | [optional] 
**CpuSystemTime** | Pointer to **NullableInt64** |  | [optional] 
**CpuUserTime** | Pointer to **NullableInt64** |  | [optional] 
**CpuTotalTime** | Pointer to **NullableInt64** |  | [optional] 
**CpuUsage** | Pointer to **NullableFloat32** |  | [optional] 
**UsedStorage** | Pointer to **NullableInt64** |  | [optional] 
**ReservedStorage** | Pointer to **NullableInt64** |  | [optional] 
**MaxStorage** | Pointer to **NullableInt64** |  | [optional] 
**NetTxUsage** | Pointer to **int64** |  | [optional] 
**NetRxUsage** | Pointer to **int64** |  | [optional] 
**NetworkBandwidth** | Pointer to **int64** |  | [optional] 

## Methods

### NewGuidanceVmwareSizingResourceStats

`func NewGuidanceVmwareSizingResourceStats() *GuidanceVmwareSizingResourceStats`

NewGuidanceVmwareSizingResourceStats instantiates a new GuidanceVmwareSizingResourceStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuidanceVmwareSizingResourceStatsWithDefaults

`func NewGuidanceVmwareSizingResourceStatsWithDefaults() *GuidanceVmwareSizingResourceStats`

NewGuidanceVmwareSizingResourceStatsWithDefaults instantiates a new GuidanceVmwareSizingResourceStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *GuidanceVmwareSizingResourceStats) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GuidanceVmwareSizingResourceStats) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GuidanceVmwareSizingResourceStats) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GuidanceVmwareSizingResourceStats) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetFreeMemory

`func (o *GuidanceVmwareSizingResourceStats) GetFreeMemory() int64`

GetFreeMemory returns the FreeMemory field if non-nil, zero value otherwise.

### GetFreeMemoryOk

`func (o *GuidanceVmwareSizingResourceStats) GetFreeMemoryOk() (*int64, bool)`

GetFreeMemoryOk returns a tuple with the FreeMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeMemory

`func (o *GuidanceVmwareSizingResourceStats) SetFreeMemory(v int64)`

SetFreeMemory sets FreeMemory field to given value.

### HasFreeMemory

`func (o *GuidanceVmwareSizingResourceStats) HasFreeMemory() bool`

HasFreeMemory returns a boolean if a field has been set.

### SetFreeMemoryNil

`func (o *GuidanceVmwareSizingResourceStats) SetFreeMemoryNil(b bool)`

 SetFreeMemoryNil sets the value for FreeMemory to be an explicit nil

### UnsetFreeMemory
`func (o *GuidanceVmwareSizingResourceStats) UnsetFreeMemory()`

UnsetFreeMemory ensures that no value is present for FreeMemory, not even an explicit nil
### GetUsedMemory

`func (o *GuidanceVmwareSizingResourceStats) GetUsedMemory() int64`

GetUsedMemory returns the UsedMemory field if non-nil, zero value otherwise.

### GetUsedMemoryOk

`func (o *GuidanceVmwareSizingResourceStats) GetUsedMemoryOk() (*int64, bool)`

GetUsedMemoryOk returns a tuple with the UsedMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedMemory

`func (o *GuidanceVmwareSizingResourceStats) SetUsedMemory(v int64)`

SetUsedMemory sets UsedMemory field to given value.

### HasUsedMemory

`func (o *GuidanceVmwareSizingResourceStats) HasUsedMemory() bool`

HasUsedMemory returns a boolean if a field has been set.

### SetUsedMemoryNil

`func (o *GuidanceVmwareSizingResourceStats) SetUsedMemoryNil(b bool)`

 SetUsedMemoryNil sets the value for UsedMemory to be an explicit nil

### UnsetUsedMemory
`func (o *GuidanceVmwareSizingResourceStats) UnsetUsedMemory()`

UnsetUsedMemory ensures that no value is present for UsedMemory, not even an explicit nil
### GetFreeSwap

`func (o *GuidanceVmwareSizingResourceStats) GetFreeSwap() int64`

GetFreeSwap returns the FreeSwap field if non-nil, zero value otherwise.

### GetFreeSwapOk

`func (o *GuidanceVmwareSizingResourceStats) GetFreeSwapOk() (*int64, bool)`

GetFreeSwapOk returns a tuple with the FreeSwap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSwap

`func (o *GuidanceVmwareSizingResourceStats) SetFreeSwap(v int64)`

SetFreeSwap sets FreeSwap field to given value.

### HasFreeSwap

`func (o *GuidanceVmwareSizingResourceStats) HasFreeSwap() bool`

HasFreeSwap returns a boolean if a field has been set.

### SetFreeSwapNil

`func (o *GuidanceVmwareSizingResourceStats) SetFreeSwapNil(b bool)`

 SetFreeSwapNil sets the value for FreeSwap to be an explicit nil

### UnsetFreeSwap
`func (o *GuidanceVmwareSizingResourceStats) UnsetFreeSwap()`

UnsetFreeSwap ensures that no value is present for FreeSwap, not even an explicit nil
### GetUsedSwap

`func (o *GuidanceVmwareSizingResourceStats) GetUsedSwap() int64`

GetUsedSwap returns the UsedSwap field if non-nil, zero value otherwise.

### GetUsedSwapOk

`func (o *GuidanceVmwareSizingResourceStats) GetUsedSwapOk() (*int64, bool)`

GetUsedSwapOk returns a tuple with the UsedSwap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSwap

`func (o *GuidanceVmwareSizingResourceStats) SetUsedSwap(v int64)`

SetUsedSwap sets UsedSwap field to given value.

### HasUsedSwap

`func (o *GuidanceVmwareSizingResourceStats) HasUsedSwap() bool`

HasUsedSwap returns a boolean if a field has been set.

### SetUsedSwapNil

`func (o *GuidanceVmwareSizingResourceStats) SetUsedSwapNil(b bool)`

 SetUsedSwapNil sets the value for UsedSwap to be an explicit nil

### UnsetUsedSwap
`func (o *GuidanceVmwareSizingResourceStats) UnsetUsedSwap()`

UnsetUsedSwap ensures that no value is present for UsedSwap, not even an explicit nil
### GetCpuIdleTime

`func (o *GuidanceVmwareSizingResourceStats) GetCpuIdleTime() int64`

GetCpuIdleTime returns the CpuIdleTime field if non-nil, zero value otherwise.

### GetCpuIdleTimeOk

`func (o *GuidanceVmwareSizingResourceStats) GetCpuIdleTimeOk() (*int64, bool)`

GetCpuIdleTimeOk returns a tuple with the CpuIdleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuIdleTime

`func (o *GuidanceVmwareSizingResourceStats) SetCpuIdleTime(v int64)`

SetCpuIdleTime sets CpuIdleTime field to given value.

### HasCpuIdleTime

`func (o *GuidanceVmwareSizingResourceStats) HasCpuIdleTime() bool`

HasCpuIdleTime returns a boolean if a field has been set.

### SetCpuIdleTimeNil

`func (o *GuidanceVmwareSizingResourceStats) SetCpuIdleTimeNil(b bool)`

 SetCpuIdleTimeNil sets the value for CpuIdleTime to be an explicit nil

### UnsetCpuIdleTime
`func (o *GuidanceVmwareSizingResourceStats) UnsetCpuIdleTime()`

UnsetCpuIdleTime ensures that no value is present for CpuIdleTime, not even an explicit nil
### GetCpuSystemTime

`func (o *GuidanceVmwareSizingResourceStats) GetCpuSystemTime() int64`

GetCpuSystemTime returns the CpuSystemTime field if non-nil, zero value otherwise.

### GetCpuSystemTimeOk

`func (o *GuidanceVmwareSizingResourceStats) GetCpuSystemTimeOk() (*int64, bool)`

GetCpuSystemTimeOk returns a tuple with the CpuSystemTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuSystemTime

`func (o *GuidanceVmwareSizingResourceStats) SetCpuSystemTime(v int64)`

SetCpuSystemTime sets CpuSystemTime field to given value.

### HasCpuSystemTime

`func (o *GuidanceVmwareSizingResourceStats) HasCpuSystemTime() bool`

HasCpuSystemTime returns a boolean if a field has been set.

### SetCpuSystemTimeNil

`func (o *GuidanceVmwareSizingResourceStats) SetCpuSystemTimeNil(b bool)`

 SetCpuSystemTimeNil sets the value for CpuSystemTime to be an explicit nil

### UnsetCpuSystemTime
`func (o *GuidanceVmwareSizingResourceStats) UnsetCpuSystemTime()`

UnsetCpuSystemTime ensures that no value is present for CpuSystemTime, not even an explicit nil
### GetCpuUserTime

`func (o *GuidanceVmwareSizingResourceStats) GetCpuUserTime() int64`

GetCpuUserTime returns the CpuUserTime field if non-nil, zero value otherwise.

### GetCpuUserTimeOk

`func (o *GuidanceVmwareSizingResourceStats) GetCpuUserTimeOk() (*int64, bool)`

GetCpuUserTimeOk returns a tuple with the CpuUserTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUserTime

`func (o *GuidanceVmwareSizingResourceStats) SetCpuUserTime(v int64)`

SetCpuUserTime sets CpuUserTime field to given value.

### HasCpuUserTime

`func (o *GuidanceVmwareSizingResourceStats) HasCpuUserTime() bool`

HasCpuUserTime returns a boolean if a field has been set.

### SetCpuUserTimeNil

`func (o *GuidanceVmwareSizingResourceStats) SetCpuUserTimeNil(b bool)`

 SetCpuUserTimeNil sets the value for CpuUserTime to be an explicit nil

### UnsetCpuUserTime
`func (o *GuidanceVmwareSizingResourceStats) UnsetCpuUserTime()`

UnsetCpuUserTime ensures that no value is present for CpuUserTime, not even an explicit nil
### GetCpuTotalTime

`func (o *GuidanceVmwareSizingResourceStats) GetCpuTotalTime() int64`

GetCpuTotalTime returns the CpuTotalTime field if non-nil, zero value otherwise.

### GetCpuTotalTimeOk

`func (o *GuidanceVmwareSizingResourceStats) GetCpuTotalTimeOk() (*int64, bool)`

GetCpuTotalTimeOk returns a tuple with the CpuTotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuTotalTime

`func (o *GuidanceVmwareSizingResourceStats) SetCpuTotalTime(v int64)`

SetCpuTotalTime sets CpuTotalTime field to given value.

### HasCpuTotalTime

`func (o *GuidanceVmwareSizingResourceStats) HasCpuTotalTime() bool`

HasCpuTotalTime returns a boolean if a field has been set.

### SetCpuTotalTimeNil

`func (o *GuidanceVmwareSizingResourceStats) SetCpuTotalTimeNil(b bool)`

 SetCpuTotalTimeNil sets the value for CpuTotalTime to be an explicit nil

### UnsetCpuTotalTime
`func (o *GuidanceVmwareSizingResourceStats) UnsetCpuTotalTime()`

UnsetCpuTotalTime ensures that no value is present for CpuTotalTime, not even an explicit nil
### GetCpuUsage

`func (o *GuidanceVmwareSizingResourceStats) GetCpuUsage() float32`

GetCpuUsage returns the CpuUsage field if non-nil, zero value otherwise.

### GetCpuUsageOk

`func (o *GuidanceVmwareSizingResourceStats) GetCpuUsageOk() (*float32, bool)`

GetCpuUsageOk returns a tuple with the CpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsage

`func (o *GuidanceVmwareSizingResourceStats) SetCpuUsage(v float32)`

SetCpuUsage sets CpuUsage field to given value.

### HasCpuUsage

`func (o *GuidanceVmwareSizingResourceStats) HasCpuUsage() bool`

HasCpuUsage returns a boolean if a field has been set.

### SetCpuUsageNil

`func (o *GuidanceVmwareSizingResourceStats) SetCpuUsageNil(b bool)`

 SetCpuUsageNil sets the value for CpuUsage to be an explicit nil

### UnsetCpuUsage
`func (o *GuidanceVmwareSizingResourceStats) UnsetCpuUsage()`

UnsetCpuUsage ensures that no value is present for CpuUsage, not even an explicit nil
### GetUsedStorage

`func (o *GuidanceVmwareSizingResourceStats) GetUsedStorage() int64`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *GuidanceVmwareSizingResourceStats) GetUsedStorageOk() (*int64, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *GuidanceVmwareSizingResourceStats) SetUsedStorage(v int64)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *GuidanceVmwareSizingResourceStats) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### SetUsedStorageNil

`func (o *GuidanceVmwareSizingResourceStats) SetUsedStorageNil(b bool)`

 SetUsedStorageNil sets the value for UsedStorage to be an explicit nil

### UnsetUsedStorage
`func (o *GuidanceVmwareSizingResourceStats) UnsetUsedStorage()`

UnsetUsedStorage ensures that no value is present for UsedStorage, not even an explicit nil
### GetReservedStorage

`func (o *GuidanceVmwareSizingResourceStats) GetReservedStorage() int64`

GetReservedStorage returns the ReservedStorage field if non-nil, zero value otherwise.

### GetReservedStorageOk

`func (o *GuidanceVmwareSizingResourceStats) GetReservedStorageOk() (*int64, bool)`

GetReservedStorageOk returns a tuple with the ReservedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedStorage

`func (o *GuidanceVmwareSizingResourceStats) SetReservedStorage(v int64)`

SetReservedStorage sets ReservedStorage field to given value.

### HasReservedStorage

`func (o *GuidanceVmwareSizingResourceStats) HasReservedStorage() bool`

HasReservedStorage returns a boolean if a field has been set.

### SetReservedStorageNil

`func (o *GuidanceVmwareSizingResourceStats) SetReservedStorageNil(b bool)`

 SetReservedStorageNil sets the value for ReservedStorage to be an explicit nil

### UnsetReservedStorage
`func (o *GuidanceVmwareSizingResourceStats) UnsetReservedStorage()`

UnsetReservedStorage ensures that no value is present for ReservedStorage, not even an explicit nil
### GetMaxStorage

`func (o *GuidanceVmwareSizingResourceStats) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *GuidanceVmwareSizingResourceStats) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *GuidanceVmwareSizingResourceStats) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *GuidanceVmwareSizingResourceStats) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### SetMaxStorageNil

`func (o *GuidanceVmwareSizingResourceStats) SetMaxStorageNil(b bool)`

 SetMaxStorageNil sets the value for MaxStorage to be an explicit nil

### UnsetMaxStorage
`func (o *GuidanceVmwareSizingResourceStats) UnsetMaxStorage()`

UnsetMaxStorage ensures that no value is present for MaxStorage, not even an explicit nil
### GetNetTxUsage

`func (o *GuidanceVmwareSizingResourceStats) GetNetTxUsage() int64`

GetNetTxUsage returns the NetTxUsage field if non-nil, zero value otherwise.

### GetNetTxUsageOk

`func (o *GuidanceVmwareSizingResourceStats) GetNetTxUsageOk() (*int64, bool)`

GetNetTxUsageOk returns a tuple with the NetTxUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetTxUsage

`func (o *GuidanceVmwareSizingResourceStats) SetNetTxUsage(v int64)`

SetNetTxUsage sets NetTxUsage field to given value.

### HasNetTxUsage

`func (o *GuidanceVmwareSizingResourceStats) HasNetTxUsage() bool`

HasNetTxUsage returns a boolean if a field has been set.

### GetNetRxUsage

`func (o *GuidanceVmwareSizingResourceStats) GetNetRxUsage() int64`

GetNetRxUsage returns the NetRxUsage field if non-nil, zero value otherwise.

### GetNetRxUsageOk

`func (o *GuidanceVmwareSizingResourceStats) GetNetRxUsageOk() (*int64, bool)`

GetNetRxUsageOk returns a tuple with the NetRxUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetRxUsage

`func (o *GuidanceVmwareSizingResourceStats) SetNetRxUsage(v int64)`

SetNetRxUsage sets NetRxUsage field to given value.

### HasNetRxUsage

`func (o *GuidanceVmwareSizingResourceStats) HasNetRxUsage() bool`

HasNetRxUsage returns a boolean if a field has been set.

### GetNetworkBandwidth

`func (o *GuidanceVmwareSizingResourceStats) GetNetworkBandwidth() int64`

GetNetworkBandwidth returns the NetworkBandwidth field if non-nil, zero value otherwise.

### GetNetworkBandwidthOk

`func (o *GuidanceVmwareSizingResourceStats) GetNetworkBandwidthOk() (*int64, bool)`

GetNetworkBandwidthOk returns a tuple with the NetworkBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkBandwidth

`func (o *GuidanceVmwareSizingResourceStats) SetNetworkBandwidth(v int64)`

SetNetworkBandwidth sets NetworkBandwidth field to given value.

### HasNetworkBandwidth

`func (o *GuidanceVmwareSizingResourceStats) HasNetworkBandwidth() bool`

HasNetworkBandwidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


