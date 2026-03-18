# HealthDatabaseInnodbStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LargeMemory** | Pointer to **int64** |  | [optional] 
**DictionaryMemory** | Pointer to **int64** |  | [optional] 
**BufferPoolSize** | Pointer to **int64** |  | [optional] 
**FreeBuffers** | Pointer to **int64** |  | [optional] 
**DatabasePages** | Pointer to **int64** |  | [optional] 
**OldPages** | Pointer to **int64** |  | [optional] 
**PendingReads** | Pointer to **int64** |  | [optional] 
**InsertsPerSecond** | Pointer to **float32** |  | [optional] 
**UpdatesPerSecond** | Pointer to **float32** |  | [optional] 
**DeletesPerSecond** | Pointer to **float32** |  | [optional] 
**ReadsPerSecond** | Pointer to **float32** |  | [optional] 
**BufferHitRate** | Pointer to **int64** |  | [optional] 

## Methods

### NewHealthDatabaseInnodbStats

`func NewHealthDatabaseInnodbStats() *HealthDatabaseInnodbStats`

NewHealthDatabaseInnodbStats instantiates a new HealthDatabaseInnodbStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthDatabaseInnodbStatsWithDefaults

`func NewHealthDatabaseInnodbStatsWithDefaults() *HealthDatabaseInnodbStats`

NewHealthDatabaseInnodbStatsWithDefaults instantiates a new HealthDatabaseInnodbStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLargeMemory

`func (o *HealthDatabaseInnodbStats) GetLargeMemory() int64`

GetLargeMemory returns the LargeMemory field if non-nil, zero value otherwise.

### GetLargeMemoryOk

`func (o *HealthDatabaseInnodbStats) GetLargeMemoryOk() (*int64, bool)`

GetLargeMemoryOk returns a tuple with the LargeMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargeMemory

`func (o *HealthDatabaseInnodbStats) SetLargeMemory(v int64)`

SetLargeMemory sets LargeMemory field to given value.

### HasLargeMemory

`func (o *HealthDatabaseInnodbStats) HasLargeMemory() bool`

HasLargeMemory returns a boolean if a field has been set.

### GetDictionaryMemory

`func (o *HealthDatabaseInnodbStats) GetDictionaryMemory() int64`

GetDictionaryMemory returns the DictionaryMemory field if non-nil, zero value otherwise.

### GetDictionaryMemoryOk

`func (o *HealthDatabaseInnodbStats) GetDictionaryMemoryOk() (*int64, bool)`

GetDictionaryMemoryOk returns a tuple with the DictionaryMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDictionaryMemory

`func (o *HealthDatabaseInnodbStats) SetDictionaryMemory(v int64)`

SetDictionaryMemory sets DictionaryMemory field to given value.

### HasDictionaryMemory

`func (o *HealthDatabaseInnodbStats) HasDictionaryMemory() bool`

HasDictionaryMemory returns a boolean if a field has been set.

### GetBufferPoolSize

`func (o *HealthDatabaseInnodbStats) GetBufferPoolSize() int64`

GetBufferPoolSize returns the BufferPoolSize field if non-nil, zero value otherwise.

### GetBufferPoolSizeOk

`func (o *HealthDatabaseInnodbStats) GetBufferPoolSizeOk() (*int64, bool)`

GetBufferPoolSizeOk returns a tuple with the BufferPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferPoolSize

`func (o *HealthDatabaseInnodbStats) SetBufferPoolSize(v int64)`

SetBufferPoolSize sets BufferPoolSize field to given value.

### HasBufferPoolSize

`func (o *HealthDatabaseInnodbStats) HasBufferPoolSize() bool`

HasBufferPoolSize returns a boolean if a field has been set.

### GetFreeBuffers

`func (o *HealthDatabaseInnodbStats) GetFreeBuffers() int64`

GetFreeBuffers returns the FreeBuffers field if non-nil, zero value otherwise.

### GetFreeBuffersOk

`func (o *HealthDatabaseInnodbStats) GetFreeBuffersOk() (*int64, bool)`

GetFreeBuffersOk returns a tuple with the FreeBuffers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeBuffers

`func (o *HealthDatabaseInnodbStats) SetFreeBuffers(v int64)`

SetFreeBuffers sets FreeBuffers field to given value.

### HasFreeBuffers

`func (o *HealthDatabaseInnodbStats) HasFreeBuffers() bool`

HasFreeBuffers returns a boolean if a field has been set.

### GetDatabasePages

`func (o *HealthDatabaseInnodbStats) GetDatabasePages() int64`

GetDatabasePages returns the DatabasePages field if non-nil, zero value otherwise.

### GetDatabasePagesOk

`func (o *HealthDatabaseInnodbStats) GetDatabasePagesOk() (*int64, bool)`

GetDatabasePagesOk returns a tuple with the DatabasePages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabasePages

`func (o *HealthDatabaseInnodbStats) SetDatabasePages(v int64)`

SetDatabasePages sets DatabasePages field to given value.

### HasDatabasePages

`func (o *HealthDatabaseInnodbStats) HasDatabasePages() bool`

HasDatabasePages returns a boolean if a field has been set.

### GetOldPages

`func (o *HealthDatabaseInnodbStats) GetOldPages() int64`

GetOldPages returns the OldPages field if non-nil, zero value otherwise.

### GetOldPagesOk

`func (o *HealthDatabaseInnodbStats) GetOldPagesOk() (*int64, bool)`

GetOldPagesOk returns a tuple with the OldPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPages

`func (o *HealthDatabaseInnodbStats) SetOldPages(v int64)`

SetOldPages sets OldPages field to given value.

### HasOldPages

`func (o *HealthDatabaseInnodbStats) HasOldPages() bool`

HasOldPages returns a boolean if a field has been set.

### GetPendingReads

`func (o *HealthDatabaseInnodbStats) GetPendingReads() int64`

GetPendingReads returns the PendingReads field if non-nil, zero value otherwise.

### GetPendingReadsOk

`func (o *HealthDatabaseInnodbStats) GetPendingReadsOk() (*int64, bool)`

GetPendingReadsOk returns a tuple with the PendingReads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingReads

`func (o *HealthDatabaseInnodbStats) SetPendingReads(v int64)`

SetPendingReads sets PendingReads field to given value.

### HasPendingReads

`func (o *HealthDatabaseInnodbStats) HasPendingReads() bool`

HasPendingReads returns a boolean if a field has been set.

### GetInsertsPerSecond

`func (o *HealthDatabaseInnodbStats) GetInsertsPerSecond() float32`

GetInsertsPerSecond returns the InsertsPerSecond field if non-nil, zero value otherwise.

### GetInsertsPerSecondOk

`func (o *HealthDatabaseInnodbStats) GetInsertsPerSecondOk() (*float32, bool)`

GetInsertsPerSecondOk returns a tuple with the InsertsPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertsPerSecond

`func (o *HealthDatabaseInnodbStats) SetInsertsPerSecond(v float32)`

SetInsertsPerSecond sets InsertsPerSecond field to given value.

### HasInsertsPerSecond

`func (o *HealthDatabaseInnodbStats) HasInsertsPerSecond() bool`

HasInsertsPerSecond returns a boolean if a field has been set.

### GetUpdatesPerSecond

`func (o *HealthDatabaseInnodbStats) GetUpdatesPerSecond() float32`

GetUpdatesPerSecond returns the UpdatesPerSecond field if non-nil, zero value otherwise.

### GetUpdatesPerSecondOk

`func (o *HealthDatabaseInnodbStats) GetUpdatesPerSecondOk() (*float32, bool)`

GetUpdatesPerSecondOk returns a tuple with the UpdatesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatesPerSecond

`func (o *HealthDatabaseInnodbStats) SetUpdatesPerSecond(v float32)`

SetUpdatesPerSecond sets UpdatesPerSecond field to given value.

### HasUpdatesPerSecond

`func (o *HealthDatabaseInnodbStats) HasUpdatesPerSecond() bool`

HasUpdatesPerSecond returns a boolean if a field has been set.

### GetDeletesPerSecond

`func (o *HealthDatabaseInnodbStats) GetDeletesPerSecond() float32`

GetDeletesPerSecond returns the DeletesPerSecond field if non-nil, zero value otherwise.

### GetDeletesPerSecondOk

`func (o *HealthDatabaseInnodbStats) GetDeletesPerSecondOk() (*float32, bool)`

GetDeletesPerSecondOk returns a tuple with the DeletesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletesPerSecond

`func (o *HealthDatabaseInnodbStats) SetDeletesPerSecond(v float32)`

SetDeletesPerSecond sets DeletesPerSecond field to given value.

### HasDeletesPerSecond

`func (o *HealthDatabaseInnodbStats) HasDeletesPerSecond() bool`

HasDeletesPerSecond returns a boolean if a field has been set.

### GetReadsPerSecond

`func (o *HealthDatabaseInnodbStats) GetReadsPerSecond() float32`

GetReadsPerSecond returns the ReadsPerSecond field if non-nil, zero value otherwise.

### GetReadsPerSecondOk

`func (o *HealthDatabaseInnodbStats) GetReadsPerSecondOk() (*float32, bool)`

GetReadsPerSecondOk returns a tuple with the ReadsPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadsPerSecond

`func (o *HealthDatabaseInnodbStats) SetReadsPerSecond(v float32)`

SetReadsPerSecond sets ReadsPerSecond field to given value.

### HasReadsPerSecond

`func (o *HealthDatabaseInnodbStats) HasReadsPerSecond() bool`

HasReadsPerSecond returns a boolean if a field has been set.

### GetBufferHitRate

`func (o *HealthDatabaseInnodbStats) GetBufferHitRate() int64`

GetBufferHitRate returns the BufferHitRate field if non-nil, zero value otherwise.

### GetBufferHitRateOk

`func (o *HealthDatabaseInnodbStats) GetBufferHitRateOk() (*int64, bool)`

GetBufferHitRateOk returns a tuple with the BufferHitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferHitRate

`func (o *HealthDatabaseInnodbStats) SetBufferHitRate(v int64)`

SetBufferHitRate sets BufferHitRate field to given value.

### HasBufferHitRate

`func (o *HealthDatabaseInnodbStats) HasBufferHitRate() bool`

HasBufferHitRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


