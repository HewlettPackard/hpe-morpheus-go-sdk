# ListHealth200ResponseAllOfHealthThreads

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**ThreadList** | Pointer to **[]map[string]interface{}** |  | [optional] 
**BusyThreads** | Pointer to [**[]ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner**](ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner.md) |  | [optional] 
**BlockedThreads** | Pointer to **[]map[string]interface{}** |  | [optional] 
**RunningThreads** | Pointer to **[]map[string]interface{}** |  | [optional] 
**TotalCpuTime** | Pointer to **int64** |  | [optional] 
**TotalThreads** | Pointer to **int64** |  | [optional] 
**RunningWebThreads** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewListHealth200ResponseAllOfHealthThreads

`func NewListHealth200ResponseAllOfHealthThreads() *ListHealth200ResponseAllOfHealthThreads`

NewListHealth200ResponseAllOfHealthThreads instantiates a new ListHealth200ResponseAllOfHealthThreads object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHealth200ResponseAllOfHealthThreadsWithDefaults

`func NewListHealth200ResponseAllOfHealthThreadsWithDefaults() *ListHealth200ResponseAllOfHealthThreads`

NewListHealth200ResponseAllOfHealthThreadsWithDefaults instantiates a new ListHealth200ResponseAllOfHealthThreads object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ListHealth200ResponseAllOfHealthThreads) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListHealth200ResponseAllOfHealthThreads) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListHealth200ResponseAllOfHealthThreads) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ListHealth200ResponseAllOfHealthThreads) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetThreadList

`func (o *ListHealth200ResponseAllOfHealthThreads) GetThreadList() []map[string]interface{}`

GetThreadList returns the ThreadList field if non-nil, zero value otherwise.

### GetThreadListOk

`func (o *ListHealth200ResponseAllOfHealthThreads) GetThreadListOk() (*[]map[string]interface{}, bool)`

GetThreadListOk returns a tuple with the ThreadList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadList

`func (o *ListHealth200ResponseAllOfHealthThreads) SetThreadList(v []map[string]interface{})`

SetThreadList sets ThreadList field to given value.

### HasThreadList

`func (o *ListHealth200ResponseAllOfHealthThreads) HasThreadList() bool`

HasThreadList returns a boolean if a field has been set.

### GetBusyThreads

`func (o *ListHealth200ResponseAllOfHealthThreads) GetBusyThreads() []ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner`

GetBusyThreads returns the BusyThreads field if non-nil, zero value otherwise.

### GetBusyThreadsOk

`func (o *ListHealth200ResponseAllOfHealthThreads) GetBusyThreadsOk() (*[]ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner, bool)`

GetBusyThreadsOk returns a tuple with the BusyThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusyThreads

`func (o *ListHealth200ResponseAllOfHealthThreads) SetBusyThreads(v []ListHealth200ResponseAllOfHealthThreadsBusyThreadsInner)`

SetBusyThreads sets BusyThreads field to given value.

### HasBusyThreads

`func (o *ListHealth200ResponseAllOfHealthThreads) HasBusyThreads() bool`

HasBusyThreads returns a boolean if a field has been set.

### GetBlockedThreads

`func (o *ListHealth200ResponseAllOfHealthThreads) GetBlockedThreads() []map[string]interface{}`

GetBlockedThreads returns the BlockedThreads field if non-nil, zero value otherwise.

### GetBlockedThreadsOk

`func (o *ListHealth200ResponseAllOfHealthThreads) GetBlockedThreadsOk() (*[]map[string]interface{}, bool)`

GetBlockedThreadsOk returns a tuple with the BlockedThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedThreads

`func (o *ListHealth200ResponseAllOfHealthThreads) SetBlockedThreads(v []map[string]interface{})`

SetBlockedThreads sets BlockedThreads field to given value.

### HasBlockedThreads

`func (o *ListHealth200ResponseAllOfHealthThreads) HasBlockedThreads() bool`

HasBlockedThreads returns a boolean if a field has been set.

### GetRunningThreads

`func (o *ListHealth200ResponseAllOfHealthThreads) GetRunningThreads() []map[string]interface{}`

GetRunningThreads returns the RunningThreads field if non-nil, zero value otherwise.

### GetRunningThreadsOk

`func (o *ListHealth200ResponseAllOfHealthThreads) GetRunningThreadsOk() (*[]map[string]interface{}, bool)`

GetRunningThreadsOk returns a tuple with the RunningThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningThreads

`func (o *ListHealth200ResponseAllOfHealthThreads) SetRunningThreads(v []map[string]interface{})`

SetRunningThreads sets RunningThreads field to given value.

### HasRunningThreads

`func (o *ListHealth200ResponseAllOfHealthThreads) HasRunningThreads() bool`

HasRunningThreads returns a boolean if a field has been set.

### GetTotalCpuTime

`func (o *ListHealth200ResponseAllOfHealthThreads) GetTotalCpuTime() int64`

GetTotalCpuTime returns the TotalCpuTime field if non-nil, zero value otherwise.

### GetTotalCpuTimeOk

`func (o *ListHealth200ResponseAllOfHealthThreads) GetTotalCpuTimeOk() (*int64, bool)`

GetTotalCpuTimeOk returns a tuple with the TotalCpuTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCpuTime

`func (o *ListHealth200ResponseAllOfHealthThreads) SetTotalCpuTime(v int64)`

SetTotalCpuTime sets TotalCpuTime field to given value.

### HasTotalCpuTime

`func (o *ListHealth200ResponseAllOfHealthThreads) HasTotalCpuTime() bool`

HasTotalCpuTime returns a boolean if a field has been set.

### GetTotalThreads

`func (o *ListHealth200ResponseAllOfHealthThreads) GetTotalThreads() int64`

GetTotalThreads returns the TotalThreads field if non-nil, zero value otherwise.

### GetTotalThreadsOk

`func (o *ListHealth200ResponseAllOfHealthThreads) GetTotalThreadsOk() (*int64, bool)`

GetTotalThreadsOk returns a tuple with the TotalThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalThreads

`func (o *ListHealth200ResponseAllOfHealthThreads) SetTotalThreads(v int64)`

SetTotalThreads sets TotalThreads field to given value.

### HasTotalThreads

`func (o *ListHealth200ResponseAllOfHealthThreads) HasTotalThreads() bool`

HasTotalThreads returns a boolean if a field has been set.

### GetRunningWebThreads

`func (o *ListHealth200ResponseAllOfHealthThreads) GetRunningWebThreads() int64`

GetRunningWebThreads returns the RunningWebThreads field if non-nil, zero value otherwise.

### GetRunningWebThreadsOk

`func (o *ListHealth200ResponseAllOfHealthThreads) GetRunningWebThreadsOk() (*int64, bool)`

GetRunningWebThreadsOk returns a tuple with the RunningWebThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningWebThreads

`func (o *ListHealth200ResponseAllOfHealthThreads) SetRunningWebThreads(v int64)`

SetRunningWebThreads sets RunningWebThreads field to given value.

### HasRunningWebThreads

`func (o *ListHealth200ResponseAllOfHealthThreads) HasRunningWebThreads() bool`

HasRunningWebThreads returns a boolean if a field has been set.

### GetStatus

`func (o *ListHealth200ResponseAllOfHealthThreads) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListHealth200ResponseAllOfHealthThreads) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListHealth200ResponseAllOfHealthThreads) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListHealth200ResponseAllOfHealthThreads) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


