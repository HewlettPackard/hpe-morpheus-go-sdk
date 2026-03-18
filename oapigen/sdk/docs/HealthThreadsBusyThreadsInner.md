# HealthThreadsBusyThreadsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreadId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**CpuTime** | Pointer to **int64** |  | [optional] 
**BlockedTime** | Pointer to **int64** |  | [optional] 
**LockName** | Pointer to **NullableString** |  | [optional] 
**LockOwnerId** | Pointer to **int64** |  | [optional] 
**LockOwnerName** | Pointer to **NullableString** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**WaitedCount** | Pointer to **int64** |  | [optional] 
**WaitedTime** | Pointer to **int64** |  | [optional] 
**IsInNative** | Pointer to **bool** |  | [optional] 
**IsSuspended** | Pointer to **bool** |  | [optional] 
**LockedMonitors** | Pointer to **[]map[string]interface{}** |  | [optional] 
**LockedSynchronizers** | Pointer to **[]map[string]interface{}** |  | [optional] 
**LockInfo** | Pointer to **NullableString** |  | [optional] 
**CurrentLines** | Pointer to **string** |  | [optional] 
**CpuPercent** | Pointer to **float32** |  | [optional] 

## Methods

### NewHealthThreadsBusyThreadsInner

`func NewHealthThreadsBusyThreadsInner() *HealthThreadsBusyThreadsInner`

NewHealthThreadsBusyThreadsInner instantiates a new HealthThreadsBusyThreadsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthThreadsBusyThreadsInnerWithDefaults

`func NewHealthThreadsBusyThreadsInnerWithDefaults() *HealthThreadsBusyThreadsInner`

NewHealthThreadsBusyThreadsInnerWithDefaults instantiates a new HealthThreadsBusyThreadsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreadId

`func (o *HealthThreadsBusyThreadsInner) GetThreadId() int64`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *HealthThreadsBusyThreadsInner) GetThreadIdOk() (*int64, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *HealthThreadsBusyThreadsInner) SetThreadId(v int64)`

SetThreadId sets ThreadId field to given value.

### HasThreadId

`func (o *HealthThreadsBusyThreadsInner) HasThreadId() bool`

HasThreadId returns a boolean if a field has been set.

### GetName

`func (o *HealthThreadsBusyThreadsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HealthThreadsBusyThreadsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HealthThreadsBusyThreadsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HealthThreadsBusyThreadsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCpuTime

`func (o *HealthThreadsBusyThreadsInner) GetCpuTime() int64`

GetCpuTime returns the CpuTime field if non-nil, zero value otherwise.

### GetCpuTimeOk

`func (o *HealthThreadsBusyThreadsInner) GetCpuTimeOk() (*int64, bool)`

GetCpuTimeOk returns a tuple with the CpuTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuTime

`func (o *HealthThreadsBusyThreadsInner) SetCpuTime(v int64)`

SetCpuTime sets CpuTime field to given value.

### HasCpuTime

`func (o *HealthThreadsBusyThreadsInner) HasCpuTime() bool`

HasCpuTime returns a boolean if a field has been set.

### GetBlockedTime

`func (o *HealthThreadsBusyThreadsInner) GetBlockedTime() int64`

GetBlockedTime returns the BlockedTime field if non-nil, zero value otherwise.

### GetBlockedTimeOk

`func (o *HealthThreadsBusyThreadsInner) GetBlockedTimeOk() (*int64, bool)`

GetBlockedTimeOk returns a tuple with the BlockedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedTime

`func (o *HealthThreadsBusyThreadsInner) SetBlockedTime(v int64)`

SetBlockedTime sets BlockedTime field to given value.

### HasBlockedTime

`func (o *HealthThreadsBusyThreadsInner) HasBlockedTime() bool`

HasBlockedTime returns a boolean if a field has been set.

### GetLockName

`func (o *HealthThreadsBusyThreadsInner) GetLockName() string`

GetLockName returns the LockName field if non-nil, zero value otherwise.

### GetLockNameOk

`func (o *HealthThreadsBusyThreadsInner) GetLockNameOk() (*string, bool)`

GetLockNameOk returns a tuple with the LockName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockName

`func (o *HealthThreadsBusyThreadsInner) SetLockName(v string)`

SetLockName sets LockName field to given value.

### HasLockName

`func (o *HealthThreadsBusyThreadsInner) HasLockName() bool`

HasLockName returns a boolean if a field has been set.

### SetLockNameNil

`func (o *HealthThreadsBusyThreadsInner) SetLockNameNil(b bool)`

 SetLockNameNil sets the value for LockName to be an explicit nil

### UnsetLockName
`func (o *HealthThreadsBusyThreadsInner) UnsetLockName()`

UnsetLockName ensures that no value is present for LockName, not even an explicit nil
### GetLockOwnerId

`func (o *HealthThreadsBusyThreadsInner) GetLockOwnerId() int64`

GetLockOwnerId returns the LockOwnerId field if non-nil, zero value otherwise.

### GetLockOwnerIdOk

`func (o *HealthThreadsBusyThreadsInner) GetLockOwnerIdOk() (*int64, bool)`

GetLockOwnerIdOk returns a tuple with the LockOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockOwnerId

`func (o *HealthThreadsBusyThreadsInner) SetLockOwnerId(v int64)`

SetLockOwnerId sets LockOwnerId field to given value.

### HasLockOwnerId

`func (o *HealthThreadsBusyThreadsInner) HasLockOwnerId() bool`

HasLockOwnerId returns a boolean if a field has been set.

### GetLockOwnerName

`func (o *HealthThreadsBusyThreadsInner) GetLockOwnerName() string`

GetLockOwnerName returns the LockOwnerName field if non-nil, zero value otherwise.

### GetLockOwnerNameOk

`func (o *HealthThreadsBusyThreadsInner) GetLockOwnerNameOk() (*string, bool)`

GetLockOwnerNameOk returns a tuple with the LockOwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockOwnerName

`func (o *HealthThreadsBusyThreadsInner) SetLockOwnerName(v string)`

SetLockOwnerName sets LockOwnerName field to given value.

### HasLockOwnerName

`func (o *HealthThreadsBusyThreadsInner) HasLockOwnerName() bool`

HasLockOwnerName returns a boolean if a field has been set.

### SetLockOwnerNameNil

`func (o *HealthThreadsBusyThreadsInner) SetLockOwnerNameNil(b bool)`

 SetLockOwnerNameNil sets the value for LockOwnerName to be an explicit nil

### UnsetLockOwnerName
`func (o *HealthThreadsBusyThreadsInner) UnsetLockOwnerName()`

UnsetLockOwnerName ensures that no value is present for LockOwnerName, not even an explicit nil
### GetState

`func (o *HealthThreadsBusyThreadsInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HealthThreadsBusyThreadsInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HealthThreadsBusyThreadsInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *HealthThreadsBusyThreadsInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetWaitedCount

`func (o *HealthThreadsBusyThreadsInner) GetWaitedCount() int64`

GetWaitedCount returns the WaitedCount field if non-nil, zero value otherwise.

### GetWaitedCountOk

`func (o *HealthThreadsBusyThreadsInner) GetWaitedCountOk() (*int64, bool)`

GetWaitedCountOk returns a tuple with the WaitedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitedCount

`func (o *HealthThreadsBusyThreadsInner) SetWaitedCount(v int64)`

SetWaitedCount sets WaitedCount field to given value.

### HasWaitedCount

`func (o *HealthThreadsBusyThreadsInner) HasWaitedCount() bool`

HasWaitedCount returns a boolean if a field has been set.

### GetWaitedTime

`func (o *HealthThreadsBusyThreadsInner) GetWaitedTime() int64`

GetWaitedTime returns the WaitedTime field if non-nil, zero value otherwise.

### GetWaitedTimeOk

`func (o *HealthThreadsBusyThreadsInner) GetWaitedTimeOk() (*int64, bool)`

GetWaitedTimeOk returns a tuple with the WaitedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitedTime

`func (o *HealthThreadsBusyThreadsInner) SetWaitedTime(v int64)`

SetWaitedTime sets WaitedTime field to given value.

### HasWaitedTime

`func (o *HealthThreadsBusyThreadsInner) HasWaitedTime() bool`

HasWaitedTime returns a boolean if a field has been set.

### GetIsInNative

`func (o *HealthThreadsBusyThreadsInner) GetIsInNative() bool`

GetIsInNative returns the IsInNative field if non-nil, zero value otherwise.

### GetIsInNativeOk

`func (o *HealthThreadsBusyThreadsInner) GetIsInNativeOk() (*bool, bool)`

GetIsInNativeOk returns a tuple with the IsInNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInNative

`func (o *HealthThreadsBusyThreadsInner) SetIsInNative(v bool)`

SetIsInNative sets IsInNative field to given value.

### HasIsInNative

`func (o *HealthThreadsBusyThreadsInner) HasIsInNative() bool`

HasIsInNative returns a boolean if a field has been set.

### GetIsSuspended

`func (o *HealthThreadsBusyThreadsInner) GetIsSuspended() bool`

GetIsSuspended returns the IsSuspended field if non-nil, zero value otherwise.

### GetIsSuspendedOk

`func (o *HealthThreadsBusyThreadsInner) GetIsSuspendedOk() (*bool, bool)`

GetIsSuspendedOk returns a tuple with the IsSuspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuspended

`func (o *HealthThreadsBusyThreadsInner) SetIsSuspended(v bool)`

SetIsSuspended sets IsSuspended field to given value.

### HasIsSuspended

`func (o *HealthThreadsBusyThreadsInner) HasIsSuspended() bool`

HasIsSuspended returns a boolean if a field has been set.

### GetLockedMonitors

`func (o *HealthThreadsBusyThreadsInner) GetLockedMonitors() []map[string]interface{}`

GetLockedMonitors returns the LockedMonitors field if non-nil, zero value otherwise.

### GetLockedMonitorsOk

`func (o *HealthThreadsBusyThreadsInner) GetLockedMonitorsOk() (*[]map[string]interface{}, bool)`

GetLockedMonitorsOk returns a tuple with the LockedMonitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedMonitors

`func (o *HealthThreadsBusyThreadsInner) SetLockedMonitors(v []map[string]interface{})`

SetLockedMonitors sets LockedMonitors field to given value.

### HasLockedMonitors

`func (o *HealthThreadsBusyThreadsInner) HasLockedMonitors() bool`

HasLockedMonitors returns a boolean if a field has been set.

### SetLockedMonitorsNil

`func (o *HealthThreadsBusyThreadsInner) SetLockedMonitorsNil(b bool)`

 SetLockedMonitorsNil sets the value for LockedMonitors to be an explicit nil

### UnsetLockedMonitors
`func (o *HealthThreadsBusyThreadsInner) UnsetLockedMonitors()`

UnsetLockedMonitors ensures that no value is present for LockedMonitors, not even an explicit nil
### GetLockedSynchronizers

`func (o *HealthThreadsBusyThreadsInner) GetLockedSynchronizers() []map[string]interface{}`

GetLockedSynchronizers returns the LockedSynchronizers field if non-nil, zero value otherwise.

### GetLockedSynchronizersOk

`func (o *HealthThreadsBusyThreadsInner) GetLockedSynchronizersOk() (*[]map[string]interface{}, bool)`

GetLockedSynchronizersOk returns a tuple with the LockedSynchronizers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedSynchronizers

`func (o *HealthThreadsBusyThreadsInner) SetLockedSynchronizers(v []map[string]interface{})`

SetLockedSynchronizers sets LockedSynchronizers field to given value.

### HasLockedSynchronizers

`func (o *HealthThreadsBusyThreadsInner) HasLockedSynchronizers() bool`

HasLockedSynchronizers returns a boolean if a field has been set.

### SetLockedSynchronizersNil

`func (o *HealthThreadsBusyThreadsInner) SetLockedSynchronizersNil(b bool)`

 SetLockedSynchronizersNil sets the value for LockedSynchronizers to be an explicit nil

### UnsetLockedSynchronizers
`func (o *HealthThreadsBusyThreadsInner) UnsetLockedSynchronizers()`

UnsetLockedSynchronizers ensures that no value is present for LockedSynchronizers, not even an explicit nil
### GetLockInfo

`func (o *HealthThreadsBusyThreadsInner) GetLockInfo() string`

GetLockInfo returns the LockInfo field if non-nil, zero value otherwise.

### GetLockInfoOk

`func (o *HealthThreadsBusyThreadsInner) GetLockInfoOk() (*string, bool)`

GetLockInfoOk returns a tuple with the LockInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockInfo

`func (o *HealthThreadsBusyThreadsInner) SetLockInfo(v string)`

SetLockInfo sets LockInfo field to given value.

### HasLockInfo

`func (o *HealthThreadsBusyThreadsInner) HasLockInfo() bool`

HasLockInfo returns a boolean if a field has been set.

### SetLockInfoNil

`func (o *HealthThreadsBusyThreadsInner) SetLockInfoNil(b bool)`

 SetLockInfoNil sets the value for LockInfo to be an explicit nil

### UnsetLockInfo
`func (o *HealthThreadsBusyThreadsInner) UnsetLockInfo()`

UnsetLockInfo ensures that no value is present for LockInfo, not even an explicit nil
### GetCurrentLines

`func (o *HealthThreadsBusyThreadsInner) GetCurrentLines() string`

GetCurrentLines returns the CurrentLines field if non-nil, zero value otherwise.

### GetCurrentLinesOk

`func (o *HealthThreadsBusyThreadsInner) GetCurrentLinesOk() (*string, bool)`

GetCurrentLinesOk returns a tuple with the CurrentLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentLines

`func (o *HealthThreadsBusyThreadsInner) SetCurrentLines(v string)`

SetCurrentLines sets CurrentLines field to given value.

### HasCurrentLines

`func (o *HealthThreadsBusyThreadsInner) HasCurrentLines() bool`

HasCurrentLines returns a boolean if a field has been set.

### GetCpuPercent

`func (o *HealthThreadsBusyThreadsInner) GetCpuPercent() float32`

GetCpuPercent returns the CpuPercent field if non-nil, zero value otherwise.

### GetCpuPercentOk

`func (o *HealthThreadsBusyThreadsInner) GetCpuPercentOk() (*float32, bool)`

GetCpuPercentOk returns a tuple with the CpuPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuPercent

`func (o *HealthThreadsBusyThreadsInner) SetCpuPercent(v float32)`

SetCpuPercent sets CpuPercent field to given value.

### HasCpuPercent

`func (o *HealthThreadsBusyThreadsInner) HasCpuPercent() bool`

HasCpuPercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


