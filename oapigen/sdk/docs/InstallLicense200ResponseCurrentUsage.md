# InstallLicense200ResponseCurrentUsage

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

### NewInstallLicense200ResponseCurrentUsage

`func NewInstallLicense200ResponseCurrentUsage() *InstallLicense200ResponseCurrentUsage`

NewInstallLicense200ResponseCurrentUsage instantiates a new InstallLicense200ResponseCurrentUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetMemory

`func (o *InstallLicense200ResponseCurrentUsage) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *InstallLicense200ResponseCurrentUsage) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *InstallLicense200ResponseCurrentUsage) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *InstallLicense200ResponseCurrentUsage) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetStorage

`func (o *InstallLicense200ResponseCurrentUsage) GetStorage() int64`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *InstallLicense200ResponseCurrentUsage) GetStorageOk() (*int64, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *InstallLicense200ResponseCurrentUsage) SetStorage(v int64)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *InstallLicense200ResponseCurrentUsage) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetWorkloads

`func (o *InstallLicense200ResponseCurrentUsage) GetWorkloads() int64`

GetWorkloads returns the Workloads field if non-nil, zero value otherwise.

### GetWorkloadsOk

`func (o *InstallLicense200ResponseCurrentUsage) GetWorkloadsOk() (*int64, bool)`

GetWorkloadsOk returns a tuple with the Workloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloads

`func (o *InstallLicense200ResponseCurrentUsage) SetWorkloads(v int64)`

SetWorkloads sets Workloads field to given value.

### HasWorkloads

`func (o *InstallLicense200ResponseCurrentUsage) HasWorkloads() bool`

HasWorkloads returns a boolean if a field has been set.

### GetDiscoveredServers

`func (o *InstallLicense200ResponseCurrentUsage) GetDiscoveredServers() int64`

GetDiscoveredServers returns the DiscoveredServers field if non-nil, zero value otherwise.

### GetDiscoveredServersOk

`func (o *InstallLicense200ResponseCurrentUsage) GetDiscoveredServersOk() (*int64, bool)`

GetDiscoveredServersOk returns a tuple with the DiscoveredServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredServers

`func (o *InstallLicense200ResponseCurrentUsage) SetDiscoveredServers(v int64)`

SetDiscoveredServers sets DiscoveredServers field to given value.

### HasDiscoveredServers

`func (o *InstallLicense200ResponseCurrentUsage) HasDiscoveredServers() bool`

HasDiscoveredServers returns a boolean if a field has been set.

### GetHosts

`func (o *InstallLicense200ResponseCurrentUsage) GetHosts() int64`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *InstallLicense200ResponseCurrentUsage) GetHostsOk() (*int64, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *InstallLicense200ResponseCurrentUsage) SetHosts(v int64)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *InstallLicense200ResponseCurrentUsage) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetMvm

`func (o *InstallLicense200ResponseCurrentUsage) GetMvm() int64`

GetMvm returns the Mvm field if non-nil, zero value otherwise.

### GetMvmOk

`func (o *InstallLicense200ResponseCurrentUsage) GetMvmOk() (*int64, bool)`

GetMvmOk returns a tuple with the Mvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMvm

`func (o *InstallLicense200ResponseCurrentUsage) SetMvm(v int64)`

SetMvm sets Mvm field to given value.

### HasMvm

`func (o *InstallLicense200ResponseCurrentUsage) HasMvm() bool`

HasMvm returns a boolean if a field has been set.

### GetMvmSockets

`func (o *InstallLicense200ResponseCurrentUsage) GetMvmSockets() int64`

GetMvmSockets returns the MvmSockets field if non-nil, zero value otherwise.

### GetMvmSocketsOk

`func (o *InstallLicense200ResponseCurrentUsage) GetMvmSocketsOk() (*int64, bool)`

GetMvmSocketsOk returns a tuple with the MvmSockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMvmSockets

`func (o *InstallLicense200ResponseCurrentUsage) SetMvmSockets(v int64)`

SetMvmSockets sets MvmSockets field to given value.

### HasMvmSockets

`func (o *InstallLicense200ResponseCurrentUsage) HasMvmSockets() bool`

HasMvmSockets returns a boolean if a field has been set.

### GetIac

`func (o *InstallLicense200ResponseCurrentUsage) GetIac() int64`

GetIac returns the Iac field if non-nil, zero value otherwise.

### GetIacOk

`func (o *InstallLicense200ResponseCurrentUsage) GetIacOk() (*int64, bool)`

GetIacOk returns a tuple with the Iac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIac

`func (o *InstallLicense200ResponseCurrentUsage) SetIac(v int64)`

SetIac sets Iac field to given value.

### HasIac

`func (o *InstallLicense200ResponseCurrentUsage) HasIac() bool`

HasIac returns a boolean if a field has been set.

### GetXaas

`func (o *InstallLicense200ResponseCurrentUsage) GetXaas() int64`

GetXaas returns the Xaas field if non-nil, zero value otherwise.

### GetXaasOk

`func (o *InstallLicense200ResponseCurrentUsage) GetXaasOk() (*int64, bool)`

GetXaasOk returns a tuple with the Xaas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXaas

`func (o *InstallLicense200ResponseCurrentUsage) SetXaas(v int64)`

SetXaas sets Xaas field to given value.

### HasXaas

`func (o *InstallLicense200ResponseCurrentUsage) HasXaas() bool`

HasXaas returns a boolean if a field has been set.

### GetExecutions

`func (o *InstallLicense200ResponseCurrentUsage) GetExecutions() int64`

GetExecutions returns the Executions field if non-nil, zero value otherwise.

### GetExecutionsOk

`func (o *InstallLicense200ResponseCurrentUsage) GetExecutionsOk() (*int64, bool)`

GetExecutionsOk returns a tuple with the Executions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutions

`func (o *InstallLicense200ResponseCurrentUsage) SetExecutions(v int64)`

SetExecutions sets Executions field to given value.

### HasExecutions

`func (o *InstallLicense200ResponseCurrentUsage) HasExecutions() bool`

HasExecutions returns a boolean if a field has been set.

### GetDistributedWorkers

`func (o *InstallLicense200ResponseCurrentUsage) GetDistributedWorkers() int64`

GetDistributedWorkers returns the DistributedWorkers field if non-nil, zero value otherwise.

### GetDistributedWorkersOk

`func (o *InstallLicense200ResponseCurrentUsage) GetDistributedWorkersOk() (*int64, bool)`

GetDistributedWorkersOk returns a tuple with the DistributedWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributedWorkers

`func (o *InstallLicense200ResponseCurrentUsage) SetDistributedWorkers(v int64)`

SetDistributedWorkers sets DistributedWorkers field to given value.

### HasDistributedWorkers

`func (o *InstallLicense200ResponseCurrentUsage) HasDistributedWorkers() bool`

HasDistributedWorkers returns a boolean if a field has been set.

### GetDiscoveredObjects

`func (o *InstallLicense200ResponseCurrentUsage) GetDiscoveredObjects() int64`

GetDiscoveredObjects returns the DiscoveredObjects field if non-nil, zero value otherwise.

### GetDiscoveredObjectsOk

`func (o *InstallLicense200ResponseCurrentUsage) GetDiscoveredObjectsOk() (*int64, bool)`

GetDiscoveredObjectsOk returns a tuple with the DiscoveredObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredObjects

`func (o *InstallLicense200ResponseCurrentUsage) SetDiscoveredObjects(v int64)`

SetDiscoveredObjects sets DiscoveredObjects field to given value.

### HasDiscoveredObjects

`func (o *InstallLicense200ResponseCurrentUsage) HasDiscoveredObjects() bool`

HasDiscoveredObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


