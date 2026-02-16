# ServicePlanRowConfigRanges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinStorage** | Pointer to **NullableInt64** |  | [optional] 
**MaxStorage** | Pointer to **NullableInt64** |  | [optional] 
**MinPerDiskSize** | Pointer to **NullableInt64** |  | [optional] 
**MaxPerDiskSize** | Pointer to **NullableInt64** |  | [optional] 
**MinMemory** | Pointer to **NullableInt64** |  | [optional] 
**MaxMemory** | Pointer to **NullableInt64** |  | [optional] 
**MinCores** | Pointer to **NullableInt64** |  | [optional] 
**MaxCores** | Pointer to **NullableInt64** |  | [optional] 
**MinSockets** | Pointer to **NullableInt64** |  | [optional] 
**MaxSockets** | Pointer to **NullableInt64** |  | [optional] 
**MinCoresPerSocket** | Pointer to **NullableInt64** |  | [optional] 
**MaxCoresPerSocket** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewServicePlanRowConfigRanges

`func NewServicePlanRowConfigRanges() *ServicePlanRowConfigRanges`

NewServicePlanRowConfigRanges instantiates a new ServicePlanRowConfigRanges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePlanRowConfigRangesWithDefaults

`func NewServicePlanRowConfigRangesWithDefaults() *ServicePlanRowConfigRanges`

NewServicePlanRowConfigRangesWithDefaults instantiates a new ServicePlanRowConfigRanges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinStorage

`func (o *ServicePlanRowConfigRanges) GetMinStorage() int64`

GetMinStorage returns the MinStorage field if non-nil, zero value otherwise.

### GetMinStorageOk

`func (o *ServicePlanRowConfigRanges) GetMinStorageOk() (*int64, bool)`

GetMinStorageOk returns a tuple with the MinStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinStorage

`func (o *ServicePlanRowConfigRanges) SetMinStorage(v int64)`

SetMinStorage sets MinStorage field to given value.

### HasMinStorage

`func (o *ServicePlanRowConfigRanges) HasMinStorage() bool`

HasMinStorage returns a boolean if a field has been set.

### SetMinStorageNil

`func (o *ServicePlanRowConfigRanges) SetMinStorageNil(b bool)`

 SetMinStorageNil sets the value for MinStorage to be an explicit nil

### UnsetMinStorage
`func (o *ServicePlanRowConfigRanges) UnsetMinStorage()`

UnsetMinStorage ensures that no value is present for MinStorage, not even an explicit nil
### GetMaxStorage

`func (o *ServicePlanRowConfigRanges) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ServicePlanRowConfigRanges) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ServicePlanRowConfigRanges) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *ServicePlanRowConfigRanges) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### SetMaxStorageNil

`func (o *ServicePlanRowConfigRanges) SetMaxStorageNil(b bool)`

 SetMaxStorageNil sets the value for MaxStorage to be an explicit nil

### UnsetMaxStorage
`func (o *ServicePlanRowConfigRanges) UnsetMaxStorage()`

UnsetMaxStorage ensures that no value is present for MaxStorage, not even an explicit nil
### GetMinPerDiskSize

`func (o *ServicePlanRowConfigRanges) GetMinPerDiskSize() int64`

GetMinPerDiskSize returns the MinPerDiskSize field if non-nil, zero value otherwise.

### GetMinPerDiskSizeOk

`func (o *ServicePlanRowConfigRanges) GetMinPerDiskSizeOk() (*int64, bool)`

GetMinPerDiskSizeOk returns a tuple with the MinPerDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPerDiskSize

`func (o *ServicePlanRowConfigRanges) SetMinPerDiskSize(v int64)`

SetMinPerDiskSize sets MinPerDiskSize field to given value.

### HasMinPerDiskSize

`func (o *ServicePlanRowConfigRanges) HasMinPerDiskSize() bool`

HasMinPerDiskSize returns a boolean if a field has been set.

### SetMinPerDiskSizeNil

`func (o *ServicePlanRowConfigRanges) SetMinPerDiskSizeNil(b bool)`

 SetMinPerDiskSizeNil sets the value for MinPerDiskSize to be an explicit nil

### UnsetMinPerDiskSize
`func (o *ServicePlanRowConfigRanges) UnsetMinPerDiskSize()`

UnsetMinPerDiskSize ensures that no value is present for MinPerDiskSize, not even an explicit nil
### GetMaxPerDiskSize

`func (o *ServicePlanRowConfigRanges) GetMaxPerDiskSize() int64`

GetMaxPerDiskSize returns the MaxPerDiskSize field if non-nil, zero value otherwise.

### GetMaxPerDiskSizeOk

`func (o *ServicePlanRowConfigRanges) GetMaxPerDiskSizeOk() (*int64, bool)`

GetMaxPerDiskSizeOk returns a tuple with the MaxPerDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPerDiskSize

`func (o *ServicePlanRowConfigRanges) SetMaxPerDiskSize(v int64)`

SetMaxPerDiskSize sets MaxPerDiskSize field to given value.

### HasMaxPerDiskSize

`func (o *ServicePlanRowConfigRanges) HasMaxPerDiskSize() bool`

HasMaxPerDiskSize returns a boolean if a field has been set.

### SetMaxPerDiskSizeNil

`func (o *ServicePlanRowConfigRanges) SetMaxPerDiskSizeNil(b bool)`

 SetMaxPerDiskSizeNil sets the value for MaxPerDiskSize to be an explicit nil

### UnsetMaxPerDiskSize
`func (o *ServicePlanRowConfigRanges) UnsetMaxPerDiskSize()`

UnsetMaxPerDiskSize ensures that no value is present for MaxPerDiskSize, not even an explicit nil
### GetMinMemory

`func (o *ServicePlanRowConfigRanges) GetMinMemory() int64`

GetMinMemory returns the MinMemory field if non-nil, zero value otherwise.

### GetMinMemoryOk

`func (o *ServicePlanRowConfigRanges) GetMinMemoryOk() (*int64, bool)`

GetMinMemoryOk returns a tuple with the MinMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMemory

`func (o *ServicePlanRowConfigRanges) SetMinMemory(v int64)`

SetMinMemory sets MinMemory field to given value.

### HasMinMemory

`func (o *ServicePlanRowConfigRanges) HasMinMemory() bool`

HasMinMemory returns a boolean if a field has been set.

### SetMinMemoryNil

`func (o *ServicePlanRowConfigRanges) SetMinMemoryNil(b bool)`

 SetMinMemoryNil sets the value for MinMemory to be an explicit nil

### UnsetMinMemory
`func (o *ServicePlanRowConfigRanges) UnsetMinMemory()`

UnsetMinMemory ensures that no value is present for MinMemory, not even an explicit nil
### GetMaxMemory

`func (o *ServicePlanRowConfigRanges) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *ServicePlanRowConfigRanges) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *ServicePlanRowConfigRanges) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *ServicePlanRowConfigRanges) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### SetMaxMemoryNil

`func (o *ServicePlanRowConfigRanges) SetMaxMemoryNil(b bool)`

 SetMaxMemoryNil sets the value for MaxMemory to be an explicit nil

### UnsetMaxMemory
`func (o *ServicePlanRowConfigRanges) UnsetMaxMemory()`

UnsetMaxMemory ensures that no value is present for MaxMemory, not even an explicit nil
### GetMinCores

`func (o *ServicePlanRowConfigRanges) GetMinCores() int64`

GetMinCores returns the MinCores field if non-nil, zero value otherwise.

### GetMinCoresOk

`func (o *ServicePlanRowConfigRanges) GetMinCoresOk() (*int64, bool)`

GetMinCoresOk returns a tuple with the MinCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCores

`func (o *ServicePlanRowConfigRanges) SetMinCores(v int64)`

SetMinCores sets MinCores field to given value.

### HasMinCores

`func (o *ServicePlanRowConfigRanges) HasMinCores() bool`

HasMinCores returns a boolean if a field has been set.

### SetMinCoresNil

`func (o *ServicePlanRowConfigRanges) SetMinCoresNil(b bool)`

 SetMinCoresNil sets the value for MinCores to be an explicit nil

### UnsetMinCores
`func (o *ServicePlanRowConfigRanges) UnsetMinCores()`

UnsetMinCores ensures that no value is present for MinCores, not even an explicit nil
### GetMaxCores

`func (o *ServicePlanRowConfigRanges) GetMaxCores() int64`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *ServicePlanRowConfigRanges) GetMaxCoresOk() (*int64, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *ServicePlanRowConfigRanges) SetMaxCores(v int64)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *ServicePlanRowConfigRanges) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### SetMaxCoresNil

`func (o *ServicePlanRowConfigRanges) SetMaxCoresNil(b bool)`

 SetMaxCoresNil sets the value for MaxCores to be an explicit nil

### UnsetMaxCores
`func (o *ServicePlanRowConfigRanges) UnsetMaxCores()`

UnsetMaxCores ensures that no value is present for MaxCores, not even an explicit nil
### GetMinSockets

`func (o *ServicePlanRowConfigRanges) GetMinSockets() int64`

GetMinSockets returns the MinSockets field if non-nil, zero value otherwise.

### GetMinSocketsOk

`func (o *ServicePlanRowConfigRanges) GetMinSocketsOk() (*int64, bool)`

GetMinSocketsOk returns a tuple with the MinSockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSockets

`func (o *ServicePlanRowConfigRanges) SetMinSockets(v int64)`

SetMinSockets sets MinSockets field to given value.

### HasMinSockets

`func (o *ServicePlanRowConfigRanges) HasMinSockets() bool`

HasMinSockets returns a boolean if a field has been set.

### SetMinSocketsNil

`func (o *ServicePlanRowConfigRanges) SetMinSocketsNil(b bool)`

 SetMinSocketsNil sets the value for MinSockets to be an explicit nil

### UnsetMinSockets
`func (o *ServicePlanRowConfigRanges) UnsetMinSockets()`

UnsetMinSockets ensures that no value is present for MinSockets, not even an explicit nil
### GetMaxSockets

`func (o *ServicePlanRowConfigRanges) GetMaxSockets() int64`

GetMaxSockets returns the MaxSockets field if non-nil, zero value otherwise.

### GetMaxSocketsOk

`func (o *ServicePlanRowConfigRanges) GetMaxSocketsOk() (*int64, bool)`

GetMaxSocketsOk returns a tuple with the MaxSockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSockets

`func (o *ServicePlanRowConfigRanges) SetMaxSockets(v int64)`

SetMaxSockets sets MaxSockets field to given value.

### HasMaxSockets

`func (o *ServicePlanRowConfigRanges) HasMaxSockets() bool`

HasMaxSockets returns a boolean if a field has been set.

### SetMaxSocketsNil

`func (o *ServicePlanRowConfigRanges) SetMaxSocketsNil(b bool)`

 SetMaxSocketsNil sets the value for MaxSockets to be an explicit nil

### UnsetMaxSockets
`func (o *ServicePlanRowConfigRanges) UnsetMaxSockets()`

UnsetMaxSockets ensures that no value is present for MaxSockets, not even an explicit nil
### GetMinCoresPerSocket

`func (o *ServicePlanRowConfigRanges) GetMinCoresPerSocket() int64`

GetMinCoresPerSocket returns the MinCoresPerSocket field if non-nil, zero value otherwise.

### GetMinCoresPerSocketOk

`func (o *ServicePlanRowConfigRanges) GetMinCoresPerSocketOk() (*int64, bool)`

GetMinCoresPerSocketOk returns a tuple with the MinCoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCoresPerSocket

`func (o *ServicePlanRowConfigRanges) SetMinCoresPerSocket(v int64)`

SetMinCoresPerSocket sets MinCoresPerSocket field to given value.

### HasMinCoresPerSocket

`func (o *ServicePlanRowConfigRanges) HasMinCoresPerSocket() bool`

HasMinCoresPerSocket returns a boolean if a field has been set.

### SetMinCoresPerSocketNil

`func (o *ServicePlanRowConfigRanges) SetMinCoresPerSocketNil(b bool)`

 SetMinCoresPerSocketNil sets the value for MinCoresPerSocket to be an explicit nil

### UnsetMinCoresPerSocket
`func (o *ServicePlanRowConfigRanges) UnsetMinCoresPerSocket()`

UnsetMinCoresPerSocket ensures that no value is present for MinCoresPerSocket, not even an explicit nil
### GetMaxCoresPerSocket

`func (o *ServicePlanRowConfigRanges) GetMaxCoresPerSocket() int64`

GetMaxCoresPerSocket returns the MaxCoresPerSocket field if non-nil, zero value otherwise.

### GetMaxCoresPerSocketOk

`func (o *ServicePlanRowConfigRanges) GetMaxCoresPerSocketOk() (*int64, bool)`

GetMaxCoresPerSocketOk returns a tuple with the MaxCoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCoresPerSocket

`func (o *ServicePlanRowConfigRanges) SetMaxCoresPerSocket(v int64)`

SetMaxCoresPerSocket sets MaxCoresPerSocket field to given value.

### HasMaxCoresPerSocket

`func (o *ServicePlanRowConfigRanges) HasMaxCoresPerSocket() bool`

HasMaxCoresPerSocket returns a boolean if a field has been set.

### SetMaxCoresPerSocketNil

`func (o *ServicePlanRowConfigRanges) SetMaxCoresPerSocketNil(b bool)`

 SetMaxCoresPerSocketNil sets the value for MaxCoresPerSocket to be an explicit nil

### UnsetMaxCoresPerSocket
`func (o *ServicePlanRowConfigRanges) UnsetMaxCoresPerSocket()`

UnsetMaxCoresPerSocket ensures that no value is present for MaxCoresPerSocket, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


