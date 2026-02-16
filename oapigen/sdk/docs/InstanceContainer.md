# InstanceContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**InternalIp** | Pointer to **string** |  | [optional] 
**InternalHostname** | Pointer to **string** |  | [optional] 
**ExternalHostname** | Pointer to **string** |  | [optional] 
**ExternalDomain** | Pointer to **string** |  | [optional] 
**ExternalFqdn** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Instance** | Pointer to [**InstanceContainerInstance**](InstanceContainerInstance.md) |  | [optional] 
**ContainerType** | Pointer to [**InstanceContainerContainerType**](InstanceContainerContainerType.md) |  | [optional] 
**ContainerTypeSet** | Pointer to [**InstanceContainerContainerTypeSet**](InstanceContainerContainerTypeSet.md) |  | [optional] 
**Server** | Pointer to [**InstanceContainerServer**](InstanceContainerServer.md) |  | [optional] 
**Cloud** | Pointer to [**InstanceContainerCloud**](InstanceContainerCloud.md) |  | [optional] 
**Ports** | Pointer to [**[]InstanceContainerPortsInner**](InstanceContainerPortsInner.md) |  | [optional] 
**Plan** | Pointer to [**InstanceContainerPlan**](InstanceContainerPlan.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**VolumeCreated** | Pointer to **bool** |  | [optional] 
**ContainerCreated** | Pointer to **bool** |  | [optional] 

## Methods

### NewInstanceContainer

`func NewInstanceContainer() *InstanceContainer`

NewInstanceContainer instantiates a new InstanceContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceContainerWithDefaults

`func NewInstanceContainerWithDefaults() *InstanceContainer`

NewInstanceContainerWithDefaults instantiates a new InstanceContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceContainer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceContainer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceContainer) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceContainer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *InstanceContainer) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *InstanceContainer) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *InstanceContainer) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *InstanceContainer) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *InstanceContainer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceContainer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceContainer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceContainer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIp

`func (o *InstanceContainer) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *InstanceContainer) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *InstanceContainer) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *InstanceContainer) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetInternalIp

`func (o *InstanceContainer) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *InstanceContainer) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *InstanceContainer) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *InstanceContainer) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### GetInternalHostname

`func (o *InstanceContainer) GetInternalHostname() string`

GetInternalHostname returns the InternalHostname field if non-nil, zero value otherwise.

### GetInternalHostnameOk

`func (o *InstanceContainer) GetInternalHostnameOk() (*string, bool)`

GetInternalHostnameOk returns a tuple with the InternalHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalHostname

`func (o *InstanceContainer) SetInternalHostname(v string)`

SetInternalHostname sets InternalHostname field to given value.

### HasInternalHostname

`func (o *InstanceContainer) HasInternalHostname() bool`

HasInternalHostname returns a boolean if a field has been set.

### GetExternalHostname

`func (o *InstanceContainer) GetExternalHostname() string`

GetExternalHostname returns the ExternalHostname field if non-nil, zero value otherwise.

### GetExternalHostnameOk

`func (o *InstanceContainer) GetExternalHostnameOk() (*string, bool)`

GetExternalHostnameOk returns a tuple with the ExternalHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalHostname

`func (o *InstanceContainer) SetExternalHostname(v string)`

SetExternalHostname sets ExternalHostname field to given value.

### HasExternalHostname

`func (o *InstanceContainer) HasExternalHostname() bool`

HasExternalHostname returns a boolean if a field has been set.

### GetExternalDomain

`func (o *InstanceContainer) GetExternalDomain() string`

GetExternalDomain returns the ExternalDomain field if non-nil, zero value otherwise.

### GetExternalDomainOk

`func (o *InstanceContainer) GetExternalDomainOk() (*string, bool)`

GetExternalDomainOk returns a tuple with the ExternalDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDomain

`func (o *InstanceContainer) SetExternalDomain(v string)`

SetExternalDomain sets ExternalDomain field to given value.

### HasExternalDomain

`func (o *InstanceContainer) HasExternalDomain() bool`

HasExternalDomain returns a boolean if a field has been set.

### GetExternalFqdn

`func (o *InstanceContainer) GetExternalFqdn() string`

GetExternalFqdn returns the ExternalFqdn field if non-nil, zero value otherwise.

### GetExternalFqdnOk

`func (o *InstanceContainer) GetExternalFqdnOk() (*string, bool)`

GetExternalFqdnOk returns a tuple with the ExternalFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFqdn

`func (o *InstanceContainer) SetExternalFqdn(v string)`

SetExternalFqdn sets ExternalFqdn field to given value.

### HasExternalFqdn

`func (o *InstanceContainer) HasExternalFqdn() bool`

HasExternalFqdn returns a boolean if a field has been set.

### GetAccountId

`func (o *InstanceContainer) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *InstanceContainer) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *InstanceContainer) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *InstanceContainer) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetInstance

`func (o *InstanceContainer) GetInstance() InstanceContainerInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *InstanceContainer) GetInstanceOk() (*InstanceContainerInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *InstanceContainer) SetInstance(v InstanceContainerInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *InstanceContainer) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetContainerType

`func (o *InstanceContainer) GetContainerType() InstanceContainerContainerType`

GetContainerType returns the ContainerType field if non-nil, zero value otherwise.

### GetContainerTypeOk

`func (o *InstanceContainer) GetContainerTypeOk() (*InstanceContainerContainerType, bool)`

GetContainerTypeOk returns a tuple with the ContainerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerType

`func (o *InstanceContainer) SetContainerType(v InstanceContainerContainerType)`

SetContainerType sets ContainerType field to given value.

### HasContainerType

`func (o *InstanceContainer) HasContainerType() bool`

HasContainerType returns a boolean if a field has been set.

### GetContainerTypeSet

`func (o *InstanceContainer) GetContainerTypeSet() InstanceContainerContainerTypeSet`

GetContainerTypeSet returns the ContainerTypeSet field if non-nil, zero value otherwise.

### GetContainerTypeSetOk

`func (o *InstanceContainer) GetContainerTypeSetOk() (*InstanceContainerContainerTypeSet, bool)`

GetContainerTypeSetOk returns a tuple with the ContainerTypeSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTypeSet

`func (o *InstanceContainer) SetContainerTypeSet(v InstanceContainerContainerTypeSet)`

SetContainerTypeSet sets ContainerTypeSet field to given value.

### HasContainerTypeSet

`func (o *InstanceContainer) HasContainerTypeSet() bool`

HasContainerTypeSet returns a boolean if a field has been set.

### GetServer

`func (o *InstanceContainer) GetServer() InstanceContainerServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *InstanceContainer) GetServerOk() (*InstanceContainerServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *InstanceContainer) SetServer(v InstanceContainerServer)`

SetServer sets Server field to given value.

### HasServer

`func (o *InstanceContainer) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetCloud

`func (o *InstanceContainer) GetCloud() InstanceContainerCloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *InstanceContainer) GetCloudOk() (*InstanceContainerCloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *InstanceContainer) SetCloud(v InstanceContainerCloud)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *InstanceContainer) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetPorts

`func (o *InstanceContainer) GetPorts() []InstanceContainerPortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InstanceContainer) GetPortsOk() (*[]InstanceContainerPortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InstanceContainer) SetPorts(v []InstanceContainerPortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InstanceContainer) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetPlan

`func (o *InstanceContainer) GetPlan() InstanceContainerPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *InstanceContainer) GetPlanOk() (*InstanceContainerPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *InstanceContainer) SetPlan(v InstanceContainerPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *InstanceContainer) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetDateCreated

`func (o *InstanceContainer) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *InstanceContainer) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *InstanceContainer) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *InstanceContainer) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *InstanceContainer) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *InstanceContainer) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *InstanceContainer) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *InstanceContainer) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetHostname

`func (o *InstanceContainer) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *InstanceContainer) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *InstanceContainer) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *InstanceContainer) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetVolumeCreated

`func (o *InstanceContainer) GetVolumeCreated() bool`

GetVolumeCreated returns the VolumeCreated field if non-nil, zero value otherwise.

### GetVolumeCreatedOk

`func (o *InstanceContainer) GetVolumeCreatedOk() (*bool, bool)`

GetVolumeCreatedOk returns a tuple with the VolumeCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeCreated

`func (o *InstanceContainer) SetVolumeCreated(v bool)`

SetVolumeCreated sets VolumeCreated field to given value.

### HasVolumeCreated

`func (o *InstanceContainer) HasVolumeCreated() bool`

HasVolumeCreated returns a boolean if a field has been set.

### GetContainerCreated

`func (o *InstanceContainer) GetContainerCreated() bool`

GetContainerCreated returns the ContainerCreated field if non-nil, zero value otherwise.

### GetContainerCreatedOk

`func (o *InstanceContainer) GetContainerCreatedOk() (*bool, bool)`

GetContainerCreatedOk returns a tuple with the ContainerCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerCreated

`func (o *InstanceContainer) SetContainerCreated(v bool)`

SetContainerCreated sets ContainerCreated field to given value.

### HasContainerCreated

`func (o *InstanceContainer) HasContainerCreated() bool`

HasContainerCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


