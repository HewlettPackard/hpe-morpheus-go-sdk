# GetLicense200ResponseCurrentUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Memory** | Pointer to **int64** | Total Used Memory (bytes) | [optional] 
**Storage** | Pointer to **int64** | Total Used Storage (bytes) | [optional] 
**Workloads** | Pointer to **int64** | Total Workloads | [optional] 
**DiscoveredServers** | Pointer to **int64** | Total Discovered Servers | [optional] 
**Hosts** | Pointer to **int64** | Total Hosts | [optional] 
**Mvm** | Pointer to **int64** | Total HPE VM Hosts | [optional] 
**MvmSockets** | Pointer to **int64** | Total HPE VM Sockets | [optional] 
**Iac** | Pointer to **int64** | Total IAC Deployments | [optional] 
**Xaas** | Pointer to **int64** | Total Xaas Instances | [optional] 
**Executions** | Pointer to **int64** | Total Executions | [optional] 
**DistributedWorkers** | Pointer to **int64** | Total Distributed Workers | [optional] 
**DiscoveredObjects** | Pointer to **int64** | Total Discovered Objects | [optional] 

## Methods

### NewGetLicense200ResponseCurrentUsage

`func NewGetLicense200ResponseCurrentUsage() *GetLicense200ResponseCurrentUsage`

NewGetLicense200ResponseCurrentUsage instantiates a new GetLicense200ResponseCurrentUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLicense200ResponseCurrentUsageWithDefaults

`func NewGetLicense200ResponseCurrentUsageWithDefaults() *GetLicense200ResponseCurrentUsage`

NewGetLicense200ResponseCurrentUsageWithDefaults instantiates a new GetLicense200ResponseCurrentUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemory

`func (o *GetLicense200ResponseCurrentUsage) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *GetLicense200ResponseCurrentUsage) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *GetLicense200ResponseCurrentUsage) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *GetLicense200ResponseCurrentUsage) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetStorage

`func (o *GetLicense200ResponseCurrentUsage) GetStorage() int64`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *GetLicense200ResponseCurrentUsage) GetStorageOk() (*int64, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *GetLicense200ResponseCurrentUsage) SetStorage(v int64)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *GetLicense200ResponseCurrentUsage) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetWorkloads

`func (o *GetLicense200ResponseCurrentUsage) GetWorkloads() int64`

GetWorkloads returns the Workloads field if non-nil, zero value otherwise.

### GetWorkloadsOk

`func (o *GetLicense200ResponseCurrentUsage) GetWorkloadsOk() (*int64, bool)`

GetWorkloadsOk returns a tuple with the Workloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloads

`func (o *GetLicense200ResponseCurrentUsage) SetWorkloads(v int64)`

SetWorkloads sets Workloads field to given value.

### HasWorkloads

`func (o *GetLicense200ResponseCurrentUsage) HasWorkloads() bool`

HasWorkloads returns a boolean if a field has been set.

### GetDiscoveredServers

`func (o *GetLicense200ResponseCurrentUsage) GetDiscoveredServers() int64`

GetDiscoveredServers returns the DiscoveredServers field if non-nil, zero value otherwise.

### GetDiscoveredServersOk

`func (o *GetLicense200ResponseCurrentUsage) GetDiscoveredServersOk() (*int64, bool)`

GetDiscoveredServersOk returns a tuple with the DiscoveredServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredServers

`func (o *GetLicense200ResponseCurrentUsage) SetDiscoveredServers(v int64)`

SetDiscoveredServers sets DiscoveredServers field to given value.

### HasDiscoveredServers

`func (o *GetLicense200ResponseCurrentUsage) HasDiscoveredServers() bool`

HasDiscoveredServers returns a boolean if a field has been set.

### GetHosts

`func (o *GetLicense200ResponseCurrentUsage) GetHosts() int64`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *GetLicense200ResponseCurrentUsage) GetHostsOk() (*int64, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *GetLicense200ResponseCurrentUsage) SetHosts(v int64)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *GetLicense200ResponseCurrentUsage) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetMvm

`func (o *GetLicense200ResponseCurrentUsage) GetMvm() int64`

GetMvm returns the Mvm field if non-nil, zero value otherwise.

### GetMvmOk

`func (o *GetLicense200ResponseCurrentUsage) GetMvmOk() (*int64, bool)`

GetMvmOk returns a tuple with the Mvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMvm

`func (o *GetLicense200ResponseCurrentUsage) SetMvm(v int64)`

SetMvm sets Mvm field to given value.

### HasMvm

`func (o *GetLicense200ResponseCurrentUsage) HasMvm() bool`

HasMvm returns a boolean if a field has been set.

### GetMvmSockets

`func (o *GetLicense200ResponseCurrentUsage) GetMvmSockets() int64`

GetMvmSockets returns the MvmSockets field if non-nil, zero value otherwise.

### GetMvmSocketsOk

`func (o *GetLicense200ResponseCurrentUsage) GetMvmSocketsOk() (*int64, bool)`

GetMvmSocketsOk returns a tuple with the MvmSockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMvmSockets

`func (o *GetLicense200ResponseCurrentUsage) SetMvmSockets(v int64)`

SetMvmSockets sets MvmSockets field to given value.

### HasMvmSockets

`func (o *GetLicense200ResponseCurrentUsage) HasMvmSockets() bool`

HasMvmSockets returns a boolean if a field has been set.

### GetIac

`func (o *GetLicense200ResponseCurrentUsage) GetIac() int64`

GetIac returns the Iac field if non-nil, zero value otherwise.

### GetIacOk

`func (o *GetLicense200ResponseCurrentUsage) GetIacOk() (*int64, bool)`

GetIacOk returns a tuple with the Iac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIac

`func (o *GetLicense200ResponseCurrentUsage) SetIac(v int64)`

SetIac sets Iac field to given value.

### HasIac

`func (o *GetLicense200ResponseCurrentUsage) HasIac() bool`

HasIac returns a boolean if a field has been set.

### GetXaas

`func (o *GetLicense200ResponseCurrentUsage) GetXaas() int64`

GetXaas returns the Xaas field if non-nil, zero value otherwise.

### GetXaasOk

`func (o *GetLicense200ResponseCurrentUsage) GetXaasOk() (*int64, bool)`

GetXaasOk returns a tuple with the Xaas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXaas

`func (o *GetLicense200ResponseCurrentUsage) SetXaas(v int64)`

SetXaas sets Xaas field to given value.

### HasXaas

`func (o *GetLicense200ResponseCurrentUsage) HasXaas() bool`

HasXaas returns a boolean if a field has been set.

### GetExecutions

`func (o *GetLicense200ResponseCurrentUsage) GetExecutions() int64`

GetExecutions returns the Executions field if non-nil, zero value otherwise.

### GetExecutionsOk

`func (o *GetLicense200ResponseCurrentUsage) GetExecutionsOk() (*int64, bool)`

GetExecutionsOk returns a tuple with the Executions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutions

`func (o *GetLicense200ResponseCurrentUsage) SetExecutions(v int64)`

SetExecutions sets Executions field to given value.

### HasExecutions

`func (o *GetLicense200ResponseCurrentUsage) HasExecutions() bool`

HasExecutions returns a boolean if a field has been set.

### GetDistributedWorkers

`func (o *GetLicense200ResponseCurrentUsage) GetDistributedWorkers() int64`

GetDistributedWorkers returns the DistributedWorkers field if non-nil, zero value otherwise.

### GetDistributedWorkersOk

`func (o *GetLicense200ResponseCurrentUsage) GetDistributedWorkersOk() (*int64, bool)`

GetDistributedWorkersOk returns a tuple with the DistributedWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributedWorkers

`func (o *GetLicense200ResponseCurrentUsage) SetDistributedWorkers(v int64)`

SetDistributedWorkers sets DistributedWorkers field to given value.

### HasDistributedWorkers

`func (o *GetLicense200ResponseCurrentUsage) HasDistributedWorkers() bool`

HasDistributedWorkers returns a boolean if a field has been set.

### GetDiscoveredObjects

`func (o *GetLicense200ResponseCurrentUsage) GetDiscoveredObjects() int64`

GetDiscoveredObjects returns the DiscoveredObjects field if non-nil, zero value otherwise.

### GetDiscoveredObjectsOk

`func (o *GetLicense200ResponseCurrentUsage) GetDiscoveredObjectsOk() (*int64, bool)`

GetDiscoveredObjectsOk returns a tuple with the DiscoveredObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredObjects

`func (o *GetLicense200ResponseCurrentUsage) SetDiscoveredObjects(v int64)`

SetDiscoveredObjects sets DiscoveredObjects field to given value.

### HasDiscoveredObjects

`func (o *GetLicense200ResponseCurrentUsage) HasDiscoveredObjects() bool`

HasDiscoveredObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


