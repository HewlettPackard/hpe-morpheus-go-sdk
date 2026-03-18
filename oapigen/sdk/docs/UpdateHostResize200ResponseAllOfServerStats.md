# UpdateHostResize200ResponseAllOfServerStats

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

### NewUpdateHostResize200ResponseAllOfServerStats

`func NewUpdateHostResize200ResponseAllOfServerStats() *UpdateHostResize200ResponseAllOfServerStats`

NewUpdateHostResize200ResponseAllOfServerStats instantiates a new UpdateHostResize200ResponseAllOfServerStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateHostResize200ResponseAllOfServerStatsWithDefaults

`func NewUpdateHostResize200ResponseAllOfServerStatsWithDefaults() *UpdateHostResize200ResponseAllOfServerStats`

NewUpdateHostResize200ResponseAllOfServerStatsWithDefaults instantiates a new UpdateHostResize200ResponseAllOfServerStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *UpdateHostResize200ResponseAllOfServerStats) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *UpdateHostResize200ResponseAllOfServerStats) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetFreeSwap

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetFreeSwap() int64`

GetFreeSwap returns the FreeSwap field if non-nil, zero value otherwise.

### GetFreeSwapOk

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetFreeSwapOk() (*int64, bool)`

GetFreeSwapOk returns a tuple with the FreeSwap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSwap

`func (o *UpdateHostResize200ResponseAllOfServerStats) SetFreeSwap(v int64)`

SetFreeSwap sets FreeSwap field to given value.

### HasFreeSwap

`func (o *UpdateHostResize200ResponseAllOfServerStats) HasFreeSwap() bool`

HasFreeSwap returns a boolean if a field has been set.

### GetUsedSwap

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetUsedSwap() int64`

GetUsedSwap returns the UsedSwap field if non-nil, zero value otherwise.

### GetUsedSwapOk

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetUsedSwapOk() (*int64, bool)`

GetUsedSwapOk returns a tuple with the UsedSwap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSwap

`func (o *UpdateHostResize200ResponseAllOfServerStats) SetUsedSwap(v int64)`

SetUsedSwap sets UsedSwap field to given value.

### HasUsedSwap

`func (o *UpdateHostResize200ResponseAllOfServerStats) HasUsedSwap() bool`

HasUsedSwap returns a boolean if a field has been set.

### GetCpuIdleTime

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetCpuIdleTime() int64`

GetCpuIdleTime returns the CpuIdleTime field if non-nil, zero value otherwise.

### GetCpuIdleTimeOk

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetCpuIdleTimeOk() (*int64, bool)`

GetCpuIdleTimeOk returns a tuple with the CpuIdleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuIdleTime

`func (o *UpdateHostResize200ResponseAllOfServerStats) SetCpuIdleTime(v int64)`

SetCpuIdleTime sets CpuIdleTime field to given value.

### HasCpuIdleTime

`func (o *UpdateHostResize200ResponseAllOfServerStats) HasCpuIdleTime() bool`

HasCpuIdleTime returns a boolean if a field has been set.

### GetCpuSystemTime

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetCpuSystemTime() int64`

GetCpuSystemTime returns the CpuSystemTime field if non-nil, zero value otherwise.

### GetCpuSystemTimeOk

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetCpuSystemTimeOk() (*int64, bool)`

GetCpuSystemTimeOk returns a tuple with the CpuSystemTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuSystemTime

`func (o *UpdateHostResize200ResponseAllOfServerStats) SetCpuSystemTime(v int64)`

SetCpuSystemTime sets CpuSystemTime field to given value.

### HasCpuSystemTime

`func (o *UpdateHostResize200ResponseAllOfServerStats) HasCpuSystemTime() bool`

HasCpuSystemTime returns a boolean if a field has been set.

### GetCpuUserTime

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetCpuUserTime() int64`

GetCpuUserTime returns the CpuUserTime field if non-nil, zero value otherwise.

### GetCpuUserTimeOk

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetCpuUserTimeOk() (*int64, bool)`

GetCpuUserTimeOk returns a tuple with the CpuUserTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUserTime

`func (o *UpdateHostResize200ResponseAllOfServerStats) SetCpuUserTime(v int64)`

SetCpuUserTime sets CpuUserTime field to given value.

### HasCpuUserTime

`func (o *UpdateHostResize200ResponseAllOfServerStats) HasCpuUserTime() bool`

HasCpuUserTime returns a boolean if a field has been set.

### GetCpuTotalTime

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetCpuTotalTime() int64`

GetCpuTotalTime returns the CpuTotalTime field if non-nil, zero value otherwise.

### GetCpuTotalTimeOk

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetCpuTotalTimeOk() (*int64, bool)`

GetCpuTotalTimeOk returns a tuple with the CpuTotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuTotalTime

`func (o *UpdateHostResize200ResponseAllOfServerStats) SetCpuTotalTime(v int64)`

SetCpuTotalTime sets CpuTotalTime field to given value.

### HasCpuTotalTime

`func (o *UpdateHostResize200ResponseAllOfServerStats) HasCpuTotalTime() bool`

HasCpuTotalTime returns a boolean if a field has been set.

### GetMaxMemory

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *UpdateHostResize200ResponseAllOfServerStats) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *UpdateHostResize200ResponseAllOfServerStats) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetUsedMemory

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetUsedMemory() int64`

GetUsedMemory returns the UsedMemory field if non-nil, zero value otherwise.

### GetUsedMemoryOk

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetUsedMemoryOk() (*int64, bool)`

GetUsedMemoryOk returns a tuple with the UsedMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedMemory

`func (o *UpdateHostResize200ResponseAllOfServerStats) SetUsedMemory(v int64)`

SetUsedMemory sets UsedMemory field to given value.

### HasUsedMemory

`func (o *UpdateHostResize200ResponseAllOfServerStats) HasUsedMemory() bool`

HasUsedMemory returns a boolean if a field has been set.

### GetMaxStorage

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *UpdateHostResize200ResponseAllOfServerStats) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *UpdateHostResize200ResponseAllOfServerStats) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetUsedStorage

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetUsedStorage() int64`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetUsedStorageOk() (*int64, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *UpdateHostResize200ResponseAllOfServerStats) SetUsedStorage(v int64)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *UpdateHostResize200ResponseAllOfServerStats) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### GetReservedStorage

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetReservedStorage() int64`

GetReservedStorage returns the ReservedStorage field if non-nil, zero value otherwise.

### GetReservedStorageOk

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetReservedStorageOk() (*int64, bool)`

GetReservedStorageOk returns a tuple with the ReservedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedStorage

`func (o *UpdateHostResize200ResponseAllOfServerStats) SetReservedStorage(v int64)`

SetReservedStorage sets ReservedStorage field to given value.

### HasReservedStorage

`func (o *UpdateHostResize200ResponseAllOfServerStats) HasReservedStorage() bool`

HasReservedStorage returns a boolean if a field has been set.

### GetCpuUsage

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetCpuUsage() float32`

GetCpuUsage returns the CpuUsage field if non-nil, zero value otherwise.

### GetCpuUsageOk

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetCpuUsageOk() (*float32, bool)`

GetCpuUsageOk returns a tuple with the CpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsage

`func (o *UpdateHostResize200ResponseAllOfServerStats) SetCpuUsage(v float32)`

SetCpuUsage sets CpuUsage field to given value.

### HasCpuUsage

`func (o *UpdateHostResize200ResponseAllOfServerStats) HasCpuUsage() bool`

HasCpuUsage returns a boolean if a field has been set.

### GetFreeMemory

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetFreeMemory() int64`

GetFreeMemory returns the FreeMemory field if non-nil, zero value otherwise.

### GetFreeMemoryOk

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetFreeMemoryOk() (*int64, bool)`

GetFreeMemoryOk returns a tuple with the FreeMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeMemory

`func (o *UpdateHostResize200ResponseAllOfServerStats) SetFreeMemory(v int64)`

SetFreeMemory sets FreeMemory field to given value.

### HasFreeMemory

`func (o *UpdateHostResize200ResponseAllOfServerStats) HasFreeMemory() bool`

HasFreeMemory returns a boolean if a field has been set.

### GetNetTxUsage

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetNetTxUsage() int64`

GetNetTxUsage returns the NetTxUsage field if non-nil, zero value otherwise.

### GetNetTxUsageOk

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetNetTxUsageOk() (*int64, bool)`

GetNetTxUsageOk returns a tuple with the NetTxUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetTxUsage

`func (o *UpdateHostResize200ResponseAllOfServerStats) SetNetTxUsage(v int64)`

SetNetTxUsage sets NetTxUsage field to given value.

### HasNetTxUsage

`func (o *UpdateHostResize200ResponseAllOfServerStats) HasNetTxUsage() bool`

HasNetTxUsage returns a boolean if a field has been set.

### GetNetRxUsage

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetNetRxUsage() int64`

GetNetRxUsage returns the NetRxUsage field if non-nil, zero value otherwise.

### GetNetRxUsageOk

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetNetRxUsageOk() (*int64, bool)`

GetNetRxUsageOk returns a tuple with the NetRxUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetRxUsage

`func (o *UpdateHostResize200ResponseAllOfServerStats) SetNetRxUsage(v int64)`

SetNetRxUsage sets NetRxUsage field to given value.

### HasNetRxUsage

`func (o *UpdateHostResize200ResponseAllOfServerStats) HasNetRxUsage() bool`

HasNetRxUsage returns a boolean if a field has been set.

### GetNetworkBandwidth

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetNetworkBandwidth() int64`

GetNetworkBandwidth returns the NetworkBandwidth field if non-nil, zero value otherwise.

### GetNetworkBandwidthOk

`func (o *UpdateHostResize200ResponseAllOfServerStats) GetNetworkBandwidthOk() (*int64, bool)`

GetNetworkBandwidthOk returns a tuple with the NetworkBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkBandwidth

`func (o *UpdateHostResize200ResponseAllOfServerStats) SetNetworkBandwidth(v int64)`

SetNetworkBandwidth sets NetworkBandwidth field to given value.

### HasNetworkBandwidth

`func (o *UpdateHostResize200ResponseAllOfServerStats) HasNetworkBandwidth() bool`

HasNetworkBandwidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


