# NetworkInterfaceUpdateSuccessServerCapacityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**MaxMemory** | Pointer to **int64** |  | [optional] 
**Server** | Pointer to [**NetworkInterfaceUpdateSuccessServerCapacityInfoServer**](NetworkInterfaceUpdateSuccessServerCapacityInfoServer.md) |  | [optional] 
**UsedStorage** | Pointer to **int64** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**MaxCpu** | Pointer to **NullableString** |  | [optional] 
**UsedCores** | Pointer to **int64** |  | [optional] 
**UsedMemory** | Pointer to **int64** |  | [optional] 
**MaxCores** | Pointer to **int64** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 

## Methods

### NewNetworkInterfaceUpdateSuccessServerCapacityInfo

`func NewNetworkInterfaceUpdateSuccessServerCapacityInfo() *NetworkInterfaceUpdateSuccessServerCapacityInfo`

NewNetworkInterfaceUpdateSuccessServerCapacityInfo instantiates a new NetworkInterfaceUpdateSuccessServerCapacityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkInterfaceUpdateSuccessServerCapacityInfoWithDefaults

`func NewNetworkInterfaceUpdateSuccessServerCapacityInfoWithDefaults() *NetworkInterfaceUpdateSuccessServerCapacityInfo`

NewNetworkInterfaceUpdateSuccessServerCapacityInfoWithDefaults instantiates a new NetworkInterfaceUpdateSuccessServerCapacityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaxMemory

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetServer

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) GetServer() NetworkInterfaceUpdateSuccessServerCapacityInfoServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) GetServerOk() (*NetworkInterfaceUpdateSuccessServerCapacityInfoServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) SetServer(v NetworkInterfaceUpdateSuccessServerCapacityInfoServer)`

SetServer sets Server field to given value.

### HasServer

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetUsedStorage

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) GetUsedStorage() int64`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) GetUsedStorageOk() (*int64, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) SetUsedStorage(v int64)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### GetVersion

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetMaxCpu

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) GetMaxCpu() string`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) GetMaxCpuOk() (*string, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) SetMaxCpu(v string)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### SetMaxCpuNil

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) SetMaxCpuNil(b bool)`

 SetMaxCpuNil sets the value for MaxCpu to be an explicit nil

### UnsetMaxCpu
`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) UnsetMaxCpu()`

UnsetMaxCpu ensures that no value is present for MaxCpu, not even an explicit nil
### GetUsedCores

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) GetUsedCores() int64`

GetUsedCores returns the UsedCores field if non-nil, zero value otherwise.

### GetUsedCoresOk

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) GetUsedCoresOk() (*int64, bool)`

GetUsedCoresOk returns a tuple with the UsedCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCores

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) SetUsedCores(v int64)`

SetUsedCores sets UsedCores field to given value.

### HasUsedCores

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) HasUsedCores() bool`

HasUsedCores returns a boolean if a field has been set.

### GetUsedMemory

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) GetUsedMemory() int64`

GetUsedMemory returns the UsedMemory field if non-nil, zero value otherwise.

### GetUsedMemoryOk

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) GetUsedMemoryOk() (*int64, bool)`

GetUsedMemoryOk returns a tuple with the UsedMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedMemory

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) SetUsedMemory(v int64)`

SetUsedMemory sets UsedMemory field to given value.

### HasUsedMemory

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) HasUsedMemory() bool`

HasUsedMemory returns a boolean if a field has been set.

### GetMaxCores

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) GetMaxCores() int64`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) GetMaxCoresOk() (*int64, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) SetMaxCores(v int64)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetMaxStorage

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *NetworkInterfaceUpdateSuccessServerCapacityInfo) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


