# HealthMemory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**MaxMemory** | Pointer to **int64** |  | [optional] 
**TotalMemory** | Pointer to **int64** |  | [optional] 
**FreeMemory** | Pointer to **int64** |  | [optional] 
**UsedMemory** | Pointer to **int64** |  | [optional] 
**SystemMemory** | Pointer to **int64** |  | [optional] 
**CommittedMemory** | Pointer to **int64** |  | [optional] 
**SystemFreeMemory** | Pointer to **int64** |  | [optional] 
**SystemSwap** | Pointer to **int64** |  | [optional] 
**SystemFreeSwap** | Pointer to **int64** |  | [optional] 
**SwapPercent** | Pointer to **int64** |  | [optional] 
**MemoryPercent** | Pointer to **float32** |  | [optional] 
**SystemMemoryPercent** | Pointer to **float32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewHealthMemory

`func NewHealthMemory() *HealthMemory`

NewHealthMemory instantiates a new HealthMemory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthMemoryWithDefaults

`func NewHealthMemoryWithDefaults() *HealthMemory`

NewHealthMemoryWithDefaults instantiates a new HealthMemory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *HealthMemory) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *HealthMemory) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *HealthMemory) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *HealthMemory) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMaxMemory

`func (o *HealthMemory) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *HealthMemory) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *HealthMemory) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *HealthMemory) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetTotalMemory

`func (o *HealthMemory) GetTotalMemory() int64`

GetTotalMemory returns the TotalMemory field if non-nil, zero value otherwise.

### GetTotalMemoryOk

`func (o *HealthMemory) GetTotalMemoryOk() (*int64, bool)`

GetTotalMemoryOk returns a tuple with the TotalMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemory

`func (o *HealthMemory) SetTotalMemory(v int64)`

SetTotalMemory sets TotalMemory field to given value.

### HasTotalMemory

`func (o *HealthMemory) HasTotalMemory() bool`

HasTotalMemory returns a boolean if a field has been set.

### GetFreeMemory

`func (o *HealthMemory) GetFreeMemory() int64`

GetFreeMemory returns the FreeMemory field if non-nil, zero value otherwise.

### GetFreeMemoryOk

`func (o *HealthMemory) GetFreeMemoryOk() (*int64, bool)`

GetFreeMemoryOk returns a tuple with the FreeMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeMemory

`func (o *HealthMemory) SetFreeMemory(v int64)`

SetFreeMemory sets FreeMemory field to given value.

### HasFreeMemory

`func (o *HealthMemory) HasFreeMemory() bool`

HasFreeMemory returns a boolean if a field has been set.

### GetUsedMemory

`func (o *HealthMemory) GetUsedMemory() int64`

GetUsedMemory returns the UsedMemory field if non-nil, zero value otherwise.

### GetUsedMemoryOk

`func (o *HealthMemory) GetUsedMemoryOk() (*int64, bool)`

GetUsedMemoryOk returns a tuple with the UsedMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedMemory

`func (o *HealthMemory) SetUsedMemory(v int64)`

SetUsedMemory sets UsedMemory field to given value.

### HasUsedMemory

`func (o *HealthMemory) HasUsedMemory() bool`

HasUsedMemory returns a boolean if a field has been set.

### GetSystemMemory

`func (o *HealthMemory) GetSystemMemory() int64`

GetSystemMemory returns the SystemMemory field if non-nil, zero value otherwise.

### GetSystemMemoryOk

`func (o *HealthMemory) GetSystemMemoryOk() (*int64, bool)`

GetSystemMemoryOk returns a tuple with the SystemMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemMemory

`func (o *HealthMemory) SetSystemMemory(v int64)`

SetSystemMemory sets SystemMemory field to given value.

### HasSystemMemory

`func (o *HealthMemory) HasSystemMemory() bool`

HasSystemMemory returns a boolean if a field has been set.

### GetCommittedMemory

`func (o *HealthMemory) GetCommittedMemory() int64`

GetCommittedMemory returns the CommittedMemory field if non-nil, zero value otherwise.

### GetCommittedMemoryOk

`func (o *HealthMemory) GetCommittedMemoryOk() (*int64, bool)`

GetCommittedMemoryOk returns a tuple with the CommittedMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommittedMemory

`func (o *HealthMemory) SetCommittedMemory(v int64)`

SetCommittedMemory sets CommittedMemory field to given value.

### HasCommittedMemory

`func (o *HealthMemory) HasCommittedMemory() bool`

HasCommittedMemory returns a boolean if a field has been set.

### GetSystemFreeMemory

`func (o *HealthMemory) GetSystemFreeMemory() int64`

GetSystemFreeMemory returns the SystemFreeMemory field if non-nil, zero value otherwise.

### GetSystemFreeMemoryOk

`func (o *HealthMemory) GetSystemFreeMemoryOk() (*int64, bool)`

GetSystemFreeMemoryOk returns a tuple with the SystemFreeMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemFreeMemory

`func (o *HealthMemory) SetSystemFreeMemory(v int64)`

SetSystemFreeMemory sets SystemFreeMemory field to given value.

### HasSystemFreeMemory

`func (o *HealthMemory) HasSystemFreeMemory() bool`

HasSystemFreeMemory returns a boolean if a field has been set.

### GetSystemSwap

`func (o *HealthMemory) GetSystemSwap() int64`

GetSystemSwap returns the SystemSwap field if non-nil, zero value otherwise.

### GetSystemSwapOk

`func (o *HealthMemory) GetSystemSwapOk() (*int64, bool)`

GetSystemSwapOk returns a tuple with the SystemSwap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemSwap

`func (o *HealthMemory) SetSystemSwap(v int64)`

SetSystemSwap sets SystemSwap field to given value.

### HasSystemSwap

`func (o *HealthMemory) HasSystemSwap() bool`

HasSystemSwap returns a boolean if a field has been set.

### GetSystemFreeSwap

`func (o *HealthMemory) GetSystemFreeSwap() int64`

GetSystemFreeSwap returns the SystemFreeSwap field if non-nil, zero value otherwise.

### GetSystemFreeSwapOk

`func (o *HealthMemory) GetSystemFreeSwapOk() (*int64, bool)`

GetSystemFreeSwapOk returns a tuple with the SystemFreeSwap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemFreeSwap

`func (o *HealthMemory) SetSystemFreeSwap(v int64)`

SetSystemFreeSwap sets SystemFreeSwap field to given value.

### HasSystemFreeSwap

`func (o *HealthMemory) HasSystemFreeSwap() bool`

HasSystemFreeSwap returns a boolean if a field has been set.

### GetSwapPercent

`func (o *HealthMemory) GetSwapPercent() int64`

GetSwapPercent returns the SwapPercent field if non-nil, zero value otherwise.

### GetSwapPercentOk

`func (o *HealthMemory) GetSwapPercentOk() (*int64, bool)`

GetSwapPercentOk returns a tuple with the SwapPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapPercent

`func (o *HealthMemory) SetSwapPercent(v int64)`

SetSwapPercent sets SwapPercent field to given value.

### HasSwapPercent

`func (o *HealthMemory) HasSwapPercent() bool`

HasSwapPercent returns a boolean if a field has been set.

### GetMemoryPercent

`func (o *HealthMemory) GetMemoryPercent() float32`

GetMemoryPercent returns the MemoryPercent field if non-nil, zero value otherwise.

### GetMemoryPercentOk

`func (o *HealthMemory) GetMemoryPercentOk() (*float32, bool)`

GetMemoryPercentOk returns a tuple with the MemoryPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPercent

`func (o *HealthMemory) SetMemoryPercent(v float32)`

SetMemoryPercent sets MemoryPercent field to given value.

### HasMemoryPercent

`func (o *HealthMemory) HasMemoryPercent() bool`

HasMemoryPercent returns a boolean if a field has been set.

### GetSystemMemoryPercent

`func (o *HealthMemory) GetSystemMemoryPercent() float32`

GetSystemMemoryPercent returns the SystemMemoryPercent field if non-nil, zero value otherwise.

### GetSystemMemoryPercentOk

`func (o *HealthMemory) GetSystemMemoryPercentOk() (*float32, bool)`

GetSystemMemoryPercentOk returns a tuple with the SystemMemoryPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemMemoryPercent

`func (o *HealthMemory) SetSystemMemoryPercent(v float32)`

SetSystemMemoryPercent sets SystemMemoryPercent field to given value.

### HasSystemMemoryPercent

`func (o *HealthMemory) HasSystemMemoryPercent() bool`

HasSystemMemoryPercent returns a boolean if a field has been set.

### GetStatus

`func (o *HealthMemory) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthMemory) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthMemory) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HealthMemory) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


