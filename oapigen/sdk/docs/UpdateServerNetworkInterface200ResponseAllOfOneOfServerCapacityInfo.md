# UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**MaxMemory** | Pointer to **int64** |  | [optional] 
**Server** | Pointer to [**UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfoServer**](UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfoServer.md) |  | [optional] 
**UsedStorage** | Pointer to **int64** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**MaxCpu** | Pointer to **NullableString** |  | [optional] 
**UsedCores** | Pointer to **int64** |  | [optional] 
**UsedMemory** | Pointer to **int64** |  | [optional] 
**MaxCores** | Pointer to **int64** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 

## Methods

### NewUpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo

`func NewUpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo() *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo`

NewUpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo instantiates a new UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaxMemory

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetServer

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) GetServer() UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfoServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) GetServerOk() (*UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfoServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) SetServer(v UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfoServer)`

SetServer sets Server field to given value.

### HasServer

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetUsedStorage

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) GetUsedStorage() int64`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) GetUsedStorageOk() (*int64, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) SetUsedStorage(v int64)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### GetVersion

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetMaxCpu

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) GetMaxCpu() string`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) GetMaxCpuOk() (*string, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) SetMaxCpu(v string)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### SetMaxCpuNil

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) SetMaxCpuNil(b bool)`

 SetMaxCpuNil sets the value for MaxCpu to be an explicit nil

### UnsetMaxCpu
`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) UnsetMaxCpu()`

UnsetMaxCpu ensures that no value is present for MaxCpu, not even an explicit nil
### GetUsedCores

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) GetUsedCores() int64`

GetUsedCores returns the UsedCores field if non-nil, zero value otherwise.

### GetUsedCoresOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) GetUsedCoresOk() (*int64, bool)`

GetUsedCoresOk returns a tuple with the UsedCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCores

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) SetUsedCores(v int64)`

SetUsedCores sets UsedCores field to given value.

### HasUsedCores

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) HasUsedCores() bool`

HasUsedCores returns a boolean if a field has been set.

### GetUsedMemory

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) GetUsedMemory() int64`

GetUsedMemory returns the UsedMemory field if non-nil, zero value otherwise.

### GetUsedMemoryOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) GetUsedMemoryOk() (*int64, bool)`

GetUsedMemoryOk returns a tuple with the UsedMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedMemory

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) SetUsedMemory(v int64)`

SetUsedMemory sets UsedMemory field to given value.

### HasUsedMemory

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) HasUsedMemory() bool`

HasUsedMemory returns a boolean if a field has been set.

### GetMaxCores

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) GetMaxCores() int64`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) GetMaxCoresOk() (*int64, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) SetMaxCores(v int64)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetMaxStorage

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *UpdateServerNetworkInterface200ResponseAllOfOneOfServerCapacityInfo) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


