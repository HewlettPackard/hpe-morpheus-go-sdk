# InstanceContainer1

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
**Server** | Pointer to [**InstanceContainerServer1**](InstanceContainerServer1.md) |  | [optional] 
**Cloud** | Pointer to [**InstanceContainerInstance**](InstanceContainerInstance.md) |  | [optional] 
**Ports** | Pointer to [**[]InstanceContainer1PortsInner**](InstanceContainer1PortsInner.md) |  | [optional] 
**Plan** | Pointer to [**InstanceContainerContainerType**](InstanceContainerContainerType.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**VolumeCreated** | Pointer to **bool** |  | [optional] 
**ContainerCreated** | Pointer to **bool** |  | [optional] 

## Methods

### NewInstanceContainer1

`func NewInstanceContainer1() *InstanceContainer1`

NewInstanceContainer1 instantiates a new InstanceContainer1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceContainer1WithDefaults

`func NewInstanceContainer1WithDefaults() *InstanceContainer1`

NewInstanceContainer1WithDefaults instantiates a new InstanceContainer1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceContainer1) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceContainer1) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceContainer1) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceContainer1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *InstanceContainer1) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *InstanceContainer1) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *InstanceContainer1) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *InstanceContainer1) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *InstanceContainer1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceContainer1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceContainer1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceContainer1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIp

`func (o *InstanceContainer1) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *InstanceContainer1) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *InstanceContainer1) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *InstanceContainer1) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetInternalIp

`func (o *InstanceContainer1) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *InstanceContainer1) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *InstanceContainer1) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *InstanceContainer1) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### GetInternalHostname

`func (o *InstanceContainer1) GetInternalHostname() string`

GetInternalHostname returns the InternalHostname field if non-nil, zero value otherwise.

### GetInternalHostnameOk

`func (o *InstanceContainer1) GetInternalHostnameOk() (*string, bool)`

GetInternalHostnameOk returns a tuple with the InternalHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalHostname

`func (o *InstanceContainer1) SetInternalHostname(v string)`

SetInternalHostname sets InternalHostname field to given value.

### HasInternalHostname

`func (o *InstanceContainer1) HasInternalHostname() bool`

HasInternalHostname returns a boolean if a field has been set.

### GetExternalHostname

`func (o *InstanceContainer1) GetExternalHostname() string`

GetExternalHostname returns the ExternalHostname field if non-nil, zero value otherwise.

### GetExternalHostnameOk

`func (o *InstanceContainer1) GetExternalHostnameOk() (*string, bool)`

GetExternalHostnameOk returns a tuple with the ExternalHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalHostname

`func (o *InstanceContainer1) SetExternalHostname(v string)`

SetExternalHostname sets ExternalHostname field to given value.

### HasExternalHostname

`func (o *InstanceContainer1) HasExternalHostname() bool`

HasExternalHostname returns a boolean if a field has been set.

### GetExternalDomain

`func (o *InstanceContainer1) GetExternalDomain() string`

GetExternalDomain returns the ExternalDomain field if non-nil, zero value otherwise.

### GetExternalDomainOk

`func (o *InstanceContainer1) GetExternalDomainOk() (*string, bool)`

GetExternalDomainOk returns a tuple with the ExternalDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDomain

`func (o *InstanceContainer1) SetExternalDomain(v string)`

SetExternalDomain sets ExternalDomain field to given value.

### HasExternalDomain

`func (o *InstanceContainer1) HasExternalDomain() bool`

HasExternalDomain returns a boolean if a field has been set.

### GetExternalFqdn

`func (o *InstanceContainer1) GetExternalFqdn() string`

GetExternalFqdn returns the ExternalFqdn field if non-nil, zero value otherwise.

### GetExternalFqdnOk

`func (o *InstanceContainer1) GetExternalFqdnOk() (*string, bool)`

GetExternalFqdnOk returns a tuple with the ExternalFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFqdn

`func (o *InstanceContainer1) SetExternalFqdn(v string)`

SetExternalFqdn sets ExternalFqdn field to given value.

### HasExternalFqdn

`func (o *InstanceContainer1) HasExternalFqdn() bool`

HasExternalFqdn returns a boolean if a field has been set.

### GetAccountId

`func (o *InstanceContainer1) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *InstanceContainer1) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *InstanceContainer1) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *InstanceContainer1) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetInstance

`func (o *InstanceContainer1) GetInstance() InstanceContainerInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *InstanceContainer1) GetInstanceOk() (*InstanceContainerInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *InstanceContainer1) SetInstance(v InstanceContainerInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *InstanceContainer1) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetContainerType

`func (o *InstanceContainer1) GetContainerType() InstanceContainerContainerType`

GetContainerType returns the ContainerType field if non-nil, zero value otherwise.

### GetContainerTypeOk

`func (o *InstanceContainer1) GetContainerTypeOk() (*InstanceContainerContainerType, bool)`

GetContainerTypeOk returns a tuple with the ContainerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerType

`func (o *InstanceContainer1) SetContainerType(v InstanceContainerContainerType)`

SetContainerType sets ContainerType field to given value.

### HasContainerType

`func (o *InstanceContainer1) HasContainerType() bool`

HasContainerType returns a boolean if a field has been set.

### GetContainerTypeSet

`func (o *InstanceContainer1) GetContainerTypeSet() InstanceContainerContainerTypeSet`

GetContainerTypeSet returns the ContainerTypeSet field if non-nil, zero value otherwise.

### GetContainerTypeSetOk

`func (o *InstanceContainer1) GetContainerTypeSetOk() (*InstanceContainerContainerTypeSet, bool)`

GetContainerTypeSetOk returns a tuple with the ContainerTypeSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTypeSet

`func (o *InstanceContainer1) SetContainerTypeSet(v InstanceContainerContainerTypeSet)`

SetContainerTypeSet sets ContainerTypeSet field to given value.

### HasContainerTypeSet

`func (o *InstanceContainer1) HasContainerTypeSet() bool`

HasContainerTypeSet returns a boolean if a field has been set.

### GetServer

`func (o *InstanceContainer1) GetServer() InstanceContainerServer1`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *InstanceContainer1) GetServerOk() (*InstanceContainerServer1, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *InstanceContainer1) SetServer(v InstanceContainerServer1)`

SetServer sets Server field to given value.

### HasServer

`func (o *InstanceContainer1) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetCloud

`func (o *InstanceContainer1) GetCloud() InstanceContainerInstance`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *InstanceContainer1) GetCloudOk() (*InstanceContainerInstance, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *InstanceContainer1) SetCloud(v InstanceContainerInstance)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *InstanceContainer1) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetPorts

`func (o *InstanceContainer1) GetPorts() []InstanceContainer1PortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InstanceContainer1) GetPortsOk() (*[]InstanceContainer1PortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InstanceContainer1) SetPorts(v []InstanceContainer1PortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InstanceContainer1) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetPlan

`func (o *InstanceContainer1) GetPlan() InstanceContainerContainerType`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *InstanceContainer1) GetPlanOk() (*InstanceContainerContainerType, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *InstanceContainer1) SetPlan(v InstanceContainerContainerType)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *InstanceContainer1) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetDateCreated

`func (o *InstanceContainer1) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *InstanceContainer1) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *InstanceContainer1) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *InstanceContainer1) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *InstanceContainer1) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *InstanceContainer1) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *InstanceContainer1) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *InstanceContainer1) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetHostname

`func (o *InstanceContainer1) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *InstanceContainer1) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *InstanceContainer1) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *InstanceContainer1) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetVolumeCreated

`func (o *InstanceContainer1) GetVolumeCreated() bool`

GetVolumeCreated returns the VolumeCreated field if non-nil, zero value otherwise.

### GetVolumeCreatedOk

`func (o *InstanceContainer1) GetVolumeCreatedOk() (*bool, bool)`

GetVolumeCreatedOk returns a tuple with the VolumeCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeCreated

`func (o *InstanceContainer1) SetVolumeCreated(v bool)`

SetVolumeCreated sets VolumeCreated field to given value.

### HasVolumeCreated

`func (o *InstanceContainer1) HasVolumeCreated() bool`

HasVolumeCreated returns a boolean if a field has been set.

### GetContainerCreated

`func (o *InstanceContainer1) GetContainerCreated() bool`

GetContainerCreated returns the ContainerCreated field if non-nil, zero value otherwise.

### GetContainerCreatedOk

`func (o *InstanceContainer1) GetContainerCreatedOk() (*bool, bool)`

GetContainerCreatedOk returns a tuple with the ContainerCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerCreated

`func (o *InstanceContainer1) SetContainerCreated(v bool)`

SetContainerCreated sets ContainerCreated field to given value.

### HasContainerCreated

`func (o *InstanceContainer1) HasContainerCreated() bool`

HasContainerCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


