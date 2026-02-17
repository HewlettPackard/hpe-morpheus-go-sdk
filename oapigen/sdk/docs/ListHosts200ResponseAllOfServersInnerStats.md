# ListHosts200ResponseAllOfServersInnerStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **time.Time** |  | [optional] 
**FreeSwap** | Pointer to **int64** |  | [optional] 
**UsedSwap** | Pointer to **int64** |  | [optional] 
**CpuIdleTime** | Pointer to **int64** |  | [optional] 
**CpuSystemTime** | Pointer to **int64** |  | [optional] 
**CpuUserTime** | Pointer to **int64** |  | [optional] 
**CpuTotalTime** | Pointer to **int64** |  | [optional] 
**MaxMemory** | Pointer to **int64** |  | [optional] 
**UsedMemory** | Pointer to **int64** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 
**UsedStorage** | Pointer to **int64** |  | [optional] 
**ReservedStorage** | Pointer to **int64** |  | [optional] 
**CpuUsage** | Pointer to **float32** |  | [optional] 
**FreeMemory** | Pointer to **int64** |  | [optional] 
**NetTxUsage** | Pointer to **int64** |  | [optional] 
**NetRxUsage** | Pointer to **int64** |  | [optional] 
**NetworkBandwidth** | Pointer to **int64** |  | [optional] 

## Methods

### NewListHosts200ResponseAllOfServersInnerStats

`func NewListHosts200ResponseAllOfServersInnerStats() *ListHosts200ResponseAllOfServersInnerStats`

NewListHosts200ResponseAllOfServersInnerStats instantiates a new ListHosts200ResponseAllOfServersInnerStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHosts200ResponseAllOfServersInnerStatsWithDefaults

`func NewListHosts200ResponseAllOfServersInnerStatsWithDefaults() *ListHosts200ResponseAllOfServersInnerStats`

NewListHosts200ResponseAllOfServersInnerStatsWithDefaults instantiates a new ListHosts200ResponseAllOfServersInnerStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *ListHosts200ResponseAllOfServersInnerStats) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *ListHosts200ResponseAllOfServersInnerStats) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetFreeSwap

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetFreeSwap() int64`

GetFreeSwap returns the FreeSwap field if non-nil, zero value otherwise.

### GetFreeSwapOk

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetFreeSwapOk() (*int64, bool)`

GetFreeSwapOk returns a tuple with the FreeSwap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSwap

`func (o *ListHosts200ResponseAllOfServersInnerStats) SetFreeSwap(v int64)`

SetFreeSwap sets FreeSwap field to given value.

### HasFreeSwap

`func (o *ListHosts200ResponseAllOfServersInnerStats) HasFreeSwap() bool`

HasFreeSwap returns a boolean if a field has been set.

### GetUsedSwap

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetUsedSwap() int64`

GetUsedSwap returns the UsedSwap field if non-nil, zero value otherwise.

### GetUsedSwapOk

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetUsedSwapOk() (*int64, bool)`

GetUsedSwapOk returns a tuple with the UsedSwap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSwap

`func (o *ListHosts200ResponseAllOfServersInnerStats) SetUsedSwap(v int64)`

SetUsedSwap sets UsedSwap field to given value.

### HasUsedSwap

`func (o *ListHosts200ResponseAllOfServersInnerStats) HasUsedSwap() bool`

HasUsedSwap returns a boolean if a field has been set.

### GetCpuIdleTime

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetCpuIdleTime() int64`

GetCpuIdleTime returns the CpuIdleTime field if non-nil, zero value otherwise.

### GetCpuIdleTimeOk

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetCpuIdleTimeOk() (*int64, bool)`

GetCpuIdleTimeOk returns a tuple with the CpuIdleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuIdleTime

`func (o *ListHosts200ResponseAllOfServersInnerStats) SetCpuIdleTime(v int64)`

SetCpuIdleTime sets CpuIdleTime field to given value.

### HasCpuIdleTime

`func (o *ListHosts200ResponseAllOfServersInnerStats) HasCpuIdleTime() bool`

HasCpuIdleTime returns a boolean if a field has been set.

### GetCpuSystemTime

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetCpuSystemTime() int64`

GetCpuSystemTime returns the CpuSystemTime field if non-nil, zero value otherwise.

### GetCpuSystemTimeOk

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetCpuSystemTimeOk() (*int64, bool)`

GetCpuSystemTimeOk returns a tuple with the CpuSystemTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuSystemTime

`func (o *ListHosts200ResponseAllOfServersInnerStats) SetCpuSystemTime(v int64)`

SetCpuSystemTime sets CpuSystemTime field to given value.

### HasCpuSystemTime

`func (o *ListHosts200ResponseAllOfServersInnerStats) HasCpuSystemTime() bool`

HasCpuSystemTime returns a boolean if a field has been set.

### GetCpuUserTime

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetCpuUserTime() int64`

GetCpuUserTime returns the CpuUserTime field if non-nil, zero value otherwise.

### GetCpuUserTimeOk

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetCpuUserTimeOk() (*int64, bool)`

GetCpuUserTimeOk returns a tuple with the CpuUserTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUserTime

`func (o *ListHosts200ResponseAllOfServersInnerStats) SetCpuUserTime(v int64)`

SetCpuUserTime sets CpuUserTime field to given value.

### HasCpuUserTime

`func (o *ListHosts200ResponseAllOfServersInnerStats) HasCpuUserTime() bool`

HasCpuUserTime returns a boolean if a field has been set.

### GetCpuTotalTime

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetCpuTotalTime() int64`

GetCpuTotalTime returns the CpuTotalTime field if non-nil, zero value otherwise.

### GetCpuTotalTimeOk

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetCpuTotalTimeOk() (*int64, bool)`

GetCpuTotalTimeOk returns a tuple with the CpuTotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuTotalTime

`func (o *ListHosts200ResponseAllOfServersInnerStats) SetCpuTotalTime(v int64)`

SetCpuTotalTime sets CpuTotalTime field to given value.

### HasCpuTotalTime

`func (o *ListHosts200ResponseAllOfServersInnerStats) HasCpuTotalTime() bool`

HasCpuTotalTime returns a boolean if a field has been set.

### GetMaxMemory

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *ListHosts200ResponseAllOfServersInnerStats) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *ListHosts200ResponseAllOfServersInnerStats) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetUsedMemory

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetUsedMemory() int64`

GetUsedMemory returns the UsedMemory field if non-nil, zero value otherwise.

### GetUsedMemoryOk

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetUsedMemoryOk() (*int64, bool)`

GetUsedMemoryOk returns a tuple with the UsedMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedMemory

`func (o *ListHosts200ResponseAllOfServersInnerStats) SetUsedMemory(v int64)`

SetUsedMemory sets UsedMemory field to given value.

### HasUsedMemory

`func (o *ListHosts200ResponseAllOfServersInnerStats) HasUsedMemory() bool`

HasUsedMemory returns a boolean if a field has been set.

### GetMaxStorage

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ListHosts200ResponseAllOfServersInnerStats) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *ListHosts200ResponseAllOfServersInnerStats) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetUsedStorage

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetUsedStorage() int64`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetUsedStorageOk() (*int64, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *ListHosts200ResponseAllOfServersInnerStats) SetUsedStorage(v int64)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *ListHosts200ResponseAllOfServersInnerStats) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### GetReservedStorage

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetReservedStorage() int64`

GetReservedStorage returns the ReservedStorage field if non-nil, zero value otherwise.

### GetReservedStorageOk

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetReservedStorageOk() (*int64, bool)`

GetReservedStorageOk returns a tuple with the ReservedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedStorage

`func (o *ListHosts200ResponseAllOfServersInnerStats) SetReservedStorage(v int64)`

SetReservedStorage sets ReservedStorage field to given value.

### HasReservedStorage

`func (o *ListHosts200ResponseAllOfServersInnerStats) HasReservedStorage() bool`

HasReservedStorage returns a boolean if a field has been set.

### GetCpuUsage

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetCpuUsage() float32`

GetCpuUsage returns the CpuUsage field if non-nil, zero value otherwise.

### GetCpuUsageOk

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetCpuUsageOk() (*float32, bool)`

GetCpuUsageOk returns a tuple with the CpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsage

`func (o *ListHosts200ResponseAllOfServersInnerStats) SetCpuUsage(v float32)`

SetCpuUsage sets CpuUsage field to given value.

### HasCpuUsage

`func (o *ListHosts200ResponseAllOfServersInnerStats) HasCpuUsage() bool`

HasCpuUsage returns a boolean if a field has been set.

### GetFreeMemory

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetFreeMemory() int64`

GetFreeMemory returns the FreeMemory field if non-nil, zero value otherwise.

### GetFreeMemoryOk

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetFreeMemoryOk() (*int64, bool)`

GetFreeMemoryOk returns a tuple with the FreeMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeMemory

`func (o *ListHosts200ResponseAllOfServersInnerStats) SetFreeMemory(v int64)`

SetFreeMemory sets FreeMemory field to given value.

### HasFreeMemory

`func (o *ListHosts200ResponseAllOfServersInnerStats) HasFreeMemory() bool`

HasFreeMemory returns a boolean if a field has been set.

### GetNetTxUsage

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetNetTxUsage() int64`

GetNetTxUsage returns the NetTxUsage field if non-nil, zero value otherwise.

### GetNetTxUsageOk

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetNetTxUsageOk() (*int64, bool)`

GetNetTxUsageOk returns a tuple with the NetTxUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetTxUsage

`func (o *ListHosts200ResponseAllOfServersInnerStats) SetNetTxUsage(v int64)`

SetNetTxUsage sets NetTxUsage field to given value.

### HasNetTxUsage

`func (o *ListHosts200ResponseAllOfServersInnerStats) HasNetTxUsage() bool`

HasNetTxUsage returns a boolean if a field has been set.

### GetNetRxUsage

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetNetRxUsage() int64`

GetNetRxUsage returns the NetRxUsage field if non-nil, zero value otherwise.

### GetNetRxUsageOk

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetNetRxUsageOk() (*int64, bool)`

GetNetRxUsageOk returns a tuple with the NetRxUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetRxUsage

`func (o *ListHosts200ResponseAllOfServersInnerStats) SetNetRxUsage(v int64)`

SetNetRxUsage sets NetRxUsage field to given value.

### HasNetRxUsage

`func (o *ListHosts200ResponseAllOfServersInnerStats) HasNetRxUsage() bool`

HasNetRxUsage returns a boolean if a field has been set.

### GetNetworkBandwidth

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetNetworkBandwidth() int64`

GetNetworkBandwidth returns the NetworkBandwidth field if non-nil, zero value otherwise.

### GetNetworkBandwidthOk

`func (o *ListHosts200ResponseAllOfServersInnerStats) GetNetworkBandwidthOk() (*int64, bool)`

GetNetworkBandwidthOk returns a tuple with the NetworkBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkBandwidth

`func (o *ListHosts200ResponseAllOfServersInnerStats) SetNetworkBandwidth(v int64)`

SetNetworkBandwidth sets NetworkBandwidth field to given value.

### HasNetworkBandwidth

`func (o *ListHosts200ResponseAllOfServersInnerStats) HasNetworkBandwidth() bool`

HasNetworkBandwidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


