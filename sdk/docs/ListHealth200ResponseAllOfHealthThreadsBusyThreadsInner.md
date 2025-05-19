# ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreadId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**CpuTime** | Pointer to **int64** |  | [optional] 
**BlockedTime** | Pointer to **int64** |  | [optional] 
**LockName** | Pointer to **string** |  | [optional] 
**LockOwnerId** | Pointer to **int64** |  | [optional] 
**LockOwnerName** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**WaitedCount** | Pointer to **int64** |  | [optional] 
**WaitedTime** | Pointer to **int64** |  | [optional] 
**IsInNative** | Pointer to **bool** |  | [optional] 
**IsSuspended** | Pointer to **bool** |  | [optional] 
**LockedMonitors** | Pointer to **[]map[string]interface{}** |  | [optional] 
**LockedSynchronizers** | Pointer to **[]map[string]interface{}** |  | [optional] 
**LockInfo** | Pointer to **string** |  | [optional] 
**CurrentLines** | Pointer to **string** |  | [optional] 
**CpuPercent** | Pointer to **float32** |  | [optional] 

## Methods

### NewListHealth200ResponseAllOfHealthThreadsBusyThreadsInner

`func NewListHealth200ResponseAllOfHealthThreadsBusyThreadsInner() *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner`

NewListHealth200ResponseAllOfHealthThreadsBusyThreadsInner instantiates a new ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHealth200ResponseAllOfHealthThreadsBusyThreadsInnerWithDefaults

`func NewListHealth200ResponseAllOfHealthThreadsBusyThreadsInnerWithDefaults() *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner`

NewListHealth200ResponseAllOfHealthThreadsBusyThreadsInnerWithDefaults instantiates a new ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreadId

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetThreadId() int64`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetThreadIdOk() (*int64, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) SetThreadId(v int64)`

SetThreadId sets ThreadId field to given value.

### HasThreadId

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) HasThreadId() bool`

HasThreadId returns a boolean if a field has been set.

### GetName

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCpuTime

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetCpuTime() int64`

GetCpuTime returns the CpuTime field if non-nil, zero value otherwise.

### GetCpuTimeOk

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetCpuTimeOk() (*int64, bool)`

GetCpuTimeOk returns a tuple with the CpuTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuTime

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) SetCpuTime(v int64)`

SetCpuTime sets CpuTime field to given value.

### HasCpuTime

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) HasCpuTime() bool`

HasCpuTime returns a boolean if a field has been set.

### GetBlockedTime

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetBlockedTime() int64`

GetBlockedTime returns the BlockedTime field if non-nil, zero value otherwise.

### GetBlockedTimeOk

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetBlockedTimeOk() (*int64, bool)`

GetBlockedTimeOk returns a tuple with the BlockedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedTime

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) SetBlockedTime(v int64)`

SetBlockedTime sets BlockedTime field to given value.

### HasBlockedTime

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) HasBlockedTime() bool`

HasBlockedTime returns a boolean if a field has been set.

### GetLockName

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetLockName() string`

GetLockName returns the LockName field if non-nil, zero value otherwise.

### GetLockNameOk

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetLockNameOk() (*string, bool)`

GetLockNameOk returns a tuple with the LockName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockName

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) SetLockName(v string)`

SetLockName sets LockName field to given value.

### HasLockName

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) HasLockName() bool`

HasLockName returns a boolean if a field has been set.

### GetLockOwnerId

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetLockOwnerId() int64`

GetLockOwnerId returns the LockOwnerId field if non-nil, zero value otherwise.

### GetLockOwnerIdOk

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetLockOwnerIdOk() (*int64, bool)`

GetLockOwnerIdOk returns a tuple with the LockOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockOwnerId

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) SetLockOwnerId(v int64)`

SetLockOwnerId sets LockOwnerId field to given value.

### HasLockOwnerId

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) HasLockOwnerId() bool`

HasLockOwnerId returns a boolean if a field has been set.

### GetLockOwnerName

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetLockOwnerName() string`

GetLockOwnerName returns the LockOwnerName field if non-nil, zero value otherwise.

### GetLockOwnerNameOk

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetLockOwnerNameOk() (*string, bool)`

GetLockOwnerNameOk returns a tuple with the LockOwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockOwnerName

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) SetLockOwnerName(v string)`

SetLockOwnerName sets LockOwnerName field to given value.

### HasLockOwnerName

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) HasLockOwnerName() bool`

HasLockOwnerName returns a boolean if a field has been set.

### GetState

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetWaitedCount

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetWaitedCount() int64`

GetWaitedCount returns the WaitedCount field if non-nil, zero value otherwise.

### GetWaitedCountOk

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetWaitedCountOk() (*int64, bool)`

GetWaitedCountOk returns a tuple with the WaitedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitedCount

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) SetWaitedCount(v int64)`

SetWaitedCount sets WaitedCount field to given value.

### HasWaitedCount

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) HasWaitedCount() bool`

HasWaitedCount returns a boolean if a field has been set.

### GetWaitedTime

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetWaitedTime() int64`

GetWaitedTime returns the WaitedTime field if non-nil, zero value otherwise.

### GetWaitedTimeOk

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetWaitedTimeOk() (*int64, bool)`

GetWaitedTimeOk returns a tuple with the WaitedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitedTime

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) SetWaitedTime(v int64)`

SetWaitedTime sets WaitedTime field to given value.

### HasWaitedTime

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) HasWaitedTime() bool`

HasWaitedTime returns a boolean if a field has been set.

### GetIsInNative

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetIsInNative() bool`

GetIsInNative returns the IsInNative field if non-nil, zero value otherwise.

### GetIsInNativeOk

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetIsInNativeOk() (*bool, bool)`

GetIsInNativeOk returns a tuple with the IsInNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInNative

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) SetIsInNative(v bool)`

SetIsInNative sets IsInNative field to given value.

### HasIsInNative

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) HasIsInNative() bool`

HasIsInNative returns a boolean if a field has been set.

### GetIsSuspended

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetIsSuspended() bool`

GetIsSuspended returns the IsSuspended field if non-nil, zero value otherwise.

### GetIsSuspendedOk

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetIsSuspendedOk() (*bool, bool)`

GetIsSuspendedOk returns a tuple with the IsSuspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuspended

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) SetIsSuspended(v bool)`

SetIsSuspended sets IsSuspended field to given value.

### HasIsSuspended

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) HasIsSuspended() bool`

HasIsSuspended returns a boolean if a field has been set.

### GetLockedMonitors

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetLockedMonitors() []map[string]interface{}`

GetLockedMonitors returns the LockedMonitors field if non-nil, zero value otherwise.

### GetLockedMonitorsOk

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetLockedMonitorsOk() (*[]map[string]interface{}, bool)`

GetLockedMonitorsOk returns a tuple with the LockedMonitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedMonitors

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) SetLockedMonitors(v []map[string]interface{})`

SetLockedMonitors sets LockedMonitors field to given value.

### HasLockedMonitors

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) HasLockedMonitors() bool`

HasLockedMonitors returns a boolean if a field has been set.

### GetLockedSynchronizers

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetLockedSynchronizers() []map[string]interface{}`

GetLockedSynchronizers returns the LockedSynchronizers field if non-nil, zero value otherwise.

### GetLockedSynchronizersOk

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetLockedSynchronizersOk() (*[]map[string]interface{}, bool)`

GetLockedSynchronizersOk returns a tuple with the LockedSynchronizers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedSynchronizers

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) SetLockedSynchronizers(v []map[string]interface{})`

SetLockedSynchronizers sets LockedSynchronizers field to given value.

### HasLockedSynchronizers

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) HasLockedSynchronizers() bool`

HasLockedSynchronizers returns a boolean if a field has been set.

### GetLockInfo

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetLockInfo() string`

GetLockInfo returns the LockInfo field if non-nil, zero value otherwise.

### GetLockInfoOk

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetLockInfoOk() (*string, bool)`

GetLockInfoOk returns a tuple with the LockInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockInfo

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) SetLockInfo(v string)`

SetLockInfo sets LockInfo field to given value.

### HasLockInfo

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) HasLockInfo() bool`

HasLockInfo returns a boolean if a field has been set.

### GetCurrentLines

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetCurrentLines() string`

GetCurrentLines returns the CurrentLines field if non-nil, zero value otherwise.

### GetCurrentLinesOk

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetCurrentLinesOk() (*string, bool)`

GetCurrentLinesOk returns a tuple with the CurrentLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentLines

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) SetCurrentLines(v string)`

SetCurrentLines sets CurrentLines field to given value.

### HasCurrentLines

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) HasCurrentLines() bool`

HasCurrentLines returns a boolean if a field has been set.

### GetCpuPercent

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetCpuPercent() float32`

GetCpuPercent returns the CpuPercent field if non-nil, zero value otherwise.

### GetCpuPercentOk

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) GetCpuPercentOk() (*float32, bool)`

GetCpuPercentOk returns a tuple with the CpuPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuPercent

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) SetCpuPercent(v float32)`

SetCpuPercent sets CpuPercent field to given value.

### HasCpuPercent

`func (o *ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner) HasCpuPercent() bool`

HasCpuPercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


