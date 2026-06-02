# HealthThreads

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**ThreadList** | Pointer to **[]map[string]interface{}** |  | [optional] 
**BusyThreads** | Pointer to [**[]HealthThreadsBusyThreadsInner**](HealthThreadsBusyThreadsInner.md) |  | [optional] 
**BlockedThreads** | Pointer to **[]map[string]interface{}** |  | [optional] 
**RunningThreads** | Pointer to **[]map[string]interface{}** |  | [optional] 
**TotalCpuTime** | Pointer to **int64** |  | [optional] 
**TotalThreads** | Pointer to **int64** |  | [optional] 
**RunningWebThreads** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewHealthThreads

`func NewHealthThreads() *HealthThreads`

NewHealthThreads instantiates a new HealthThreads object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetSuccess

`func (o *HealthThreads) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *HealthThreads) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *HealthThreads) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *HealthThreads) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetThreadList

`func (o *HealthThreads) GetThreadList() []map[string]interface{}`

GetThreadList returns the ThreadList field if non-nil, zero value otherwise.

### GetThreadListOk

`func (o *HealthThreads) GetThreadListOk() (*[]map[string]interface{}, bool)`

GetThreadListOk returns a tuple with the ThreadList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadList

`func (o *HealthThreads) SetThreadList(v []map[string]interface{})`

SetThreadList sets ThreadList field to given value.

### HasThreadList

`func (o *HealthThreads) HasThreadList() bool`

HasThreadList returns a boolean if a field has been set.

### SetThreadListNil

`func (o *HealthThreads) SetThreadListNil(b bool)`

 SetThreadListNil sets the value for ThreadList to be an explicit nil

### UnsetThreadList
`func (o *HealthThreads) UnsetThreadList()`

UnsetThreadList ensures that no value is present for ThreadList, not even an explicit nil
### GetBusyThreads

`func (o *HealthThreads) GetBusyThreads() []HealthThreadsBusyThreadsInner`

GetBusyThreads returns the BusyThreads field if non-nil, zero value otherwise.

### GetBusyThreadsOk

`func (o *HealthThreads) GetBusyThreadsOk() (*[]HealthThreadsBusyThreadsInner, bool)`

GetBusyThreadsOk returns a tuple with the BusyThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusyThreads

`func (o *HealthThreads) SetBusyThreads(v []HealthThreadsBusyThreadsInner)`

SetBusyThreads sets BusyThreads field to given value.

### HasBusyThreads

`func (o *HealthThreads) HasBusyThreads() bool`

HasBusyThreads returns a boolean if a field has been set.

### GetBlockedThreads

`func (o *HealthThreads) GetBlockedThreads() []map[string]interface{}`

GetBlockedThreads returns the BlockedThreads field if non-nil, zero value otherwise.

### GetBlockedThreadsOk

`func (o *HealthThreads) GetBlockedThreadsOk() (*[]map[string]interface{}, bool)`

GetBlockedThreadsOk returns a tuple with the BlockedThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedThreads

`func (o *HealthThreads) SetBlockedThreads(v []map[string]interface{})`

SetBlockedThreads sets BlockedThreads field to given value.

### HasBlockedThreads

`func (o *HealthThreads) HasBlockedThreads() bool`

HasBlockedThreads returns a boolean if a field has been set.

### SetBlockedThreadsNil

`func (o *HealthThreads) SetBlockedThreadsNil(b bool)`

 SetBlockedThreadsNil sets the value for BlockedThreads to be an explicit nil

### UnsetBlockedThreads
`func (o *HealthThreads) UnsetBlockedThreads()`

UnsetBlockedThreads ensures that no value is present for BlockedThreads, not even an explicit nil
### GetRunningThreads

`func (o *HealthThreads) GetRunningThreads() []map[string]interface{}`

GetRunningThreads returns the RunningThreads field if non-nil, zero value otherwise.

### GetRunningThreadsOk

`func (o *HealthThreads) GetRunningThreadsOk() (*[]map[string]interface{}, bool)`

GetRunningThreadsOk returns a tuple with the RunningThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningThreads

`func (o *HealthThreads) SetRunningThreads(v []map[string]interface{})`

SetRunningThreads sets RunningThreads field to given value.

### HasRunningThreads

`func (o *HealthThreads) HasRunningThreads() bool`

HasRunningThreads returns a boolean if a field has been set.

### SetRunningThreadsNil

`func (o *HealthThreads) SetRunningThreadsNil(b bool)`

 SetRunningThreadsNil sets the value for RunningThreads to be an explicit nil

### UnsetRunningThreads
`func (o *HealthThreads) UnsetRunningThreads()`

UnsetRunningThreads ensures that no value is present for RunningThreads, not even an explicit nil
### GetTotalCpuTime

`func (o *HealthThreads) GetTotalCpuTime() int64`

GetTotalCpuTime returns the TotalCpuTime field if non-nil, zero value otherwise.

### GetTotalCpuTimeOk

`func (o *HealthThreads) GetTotalCpuTimeOk() (*int64, bool)`

GetTotalCpuTimeOk returns a tuple with the TotalCpuTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCpuTime

`func (o *HealthThreads) SetTotalCpuTime(v int64)`

SetTotalCpuTime sets TotalCpuTime field to given value.

### HasTotalCpuTime

`func (o *HealthThreads) HasTotalCpuTime() bool`

HasTotalCpuTime returns a boolean if a field has been set.

### GetTotalThreads

`func (o *HealthThreads) GetTotalThreads() int64`

GetTotalThreads returns the TotalThreads field if non-nil, zero value otherwise.

### GetTotalThreadsOk

`func (o *HealthThreads) GetTotalThreadsOk() (*int64, bool)`

GetTotalThreadsOk returns a tuple with the TotalThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalThreads

`func (o *HealthThreads) SetTotalThreads(v int64)`

SetTotalThreads sets TotalThreads field to given value.

### HasTotalThreads

`func (o *HealthThreads) HasTotalThreads() bool`

HasTotalThreads returns a boolean if a field has been set.

### GetRunningWebThreads

`func (o *HealthThreads) GetRunningWebThreads() int64`

GetRunningWebThreads returns the RunningWebThreads field if non-nil, zero value otherwise.

### GetRunningWebThreadsOk

`func (o *HealthThreads) GetRunningWebThreadsOk() (*int64, bool)`

GetRunningWebThreadsOk returns a tuple with the RunningWebThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningWebThreads

`func (o *HealthThreads) SetRunningWebThreads(v int64)`

SetRunningWebThreads sets RunningWebThreads field to given value.

### HasRunningWebThreads

`func (o *HealthThreads) HasRunningWebThreads() bool`

HasRunningWebThreads returns a boolean if a field has been set.

### GetStatus

`func (o *HealthThreads) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthThreads) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthThreads) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HealthThreads) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


