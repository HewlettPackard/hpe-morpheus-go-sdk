# InstanceContainer4

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
**Instance** | Pointer to [**InstanceContainer4Instance**](InstanceContainer4Instance.md) |  | [optional] 
**ContainerType** | Pointer to [**InstanceContainer4ContainerType**](InstanceContainer4ContainerType.md) |  | [optional] 
**ContainerTypeSet** | Pointer to [**InstanceContainer4ContainerTypeSet**](InstanceContainer4ContainerTypeSet.md) |  | [optional] 
**Server** | Pointer to [**InstanceContainerServer4**](InstanceContainerServer4.md) |  | [optional] 
**Cloud** | Pointer to [**InstanceContainer4Cloud**](InstanceContainer4Cloud.md) |  | [optional] 
**Ports** | Pointer to [**[]InstanceContainer1PortsInner**](InstanceContainer1PortsInner.md) |  | [optional] 
**Plan** | Pointer to [**InstanceContainer4Plan**](InstanceContainer4Plan.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**VolumeCreated** | Pointer to **bool** |  | [optional] 
**ContainerCreated** | Pointer to **bool** |  | [optional] 

## Methods

### NewInstanceContainer4

`func NewInstanceContainer4() *InstanceContainer4`

NewInstanceContainer4 instantiates a new InstanceContainer4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceContainer4WithDefaults

`func NewInstanceContainer4WithDefaults() *InstanceContainer4`

NewInstanceContainer4WithDefaults instantiates a new InstanceContainer4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceContainer4) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceContainer4) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceContainer4) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceContainer4) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *InstanceContainer4) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *InstanceContainer4) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *InstanceContainer4) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *InstanceContainer4) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *InstanceContainer4) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceContainer4) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceContainer4) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceContainer4) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIp

`func (o *InstanceContainer4) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *InstanceContainer4) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *InstanceContainer4) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *InstanceContainer4) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetInternalIp

`func (o *InstanceContainer4) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *InstanceContainer4) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *InstanceContainer4) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *InstanceContainer4) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### GetInternalHostname

`func (o *InstanceContainer4) GetInternalHostname() string`

GetInternalHostname returns the InternalHostname field if non-nil, zero value otherwise.

### GetInternalHostnameOk

`func (o *InstanceContainer4) GetInternalHostnameOk() (*string, bool)`

GetInternalHostnameOk returns a tuple with the InternalHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalHostname

`func (o *InstanceContainer4) SetInternalHostname(v string)`

SetInternalHostname sets InternalHostname field to given value.

### HasInternalHostname

`func (o *InstanceContainer4) HasInternalHostname() bool`

HasInternalHostname returns a boolean if a field has been set.

### GetExternalHostname

`func (o *InstanceContainer4) GetExternalHostname() string`

GetExternalHostname returns the ExternalHostname field if non-nil, zero value otherwise.

### GetExternalHostnameOk

`func (o *InstanceContainer4) GetExternalHostnameOk() (*string, bool)`

GetExternalHostnameOk returns a tuple with the ExternalHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalHostname

`func (o *InstanceContainer4) SetExternalHostname(v string)`

SetExternalHostname sets ExternalHostname field to given value.

### HasExternalHostname

`func (o *InstanceContainer4) HasExternalHostname() bool`

HasExternalHostname returns a boolean if a field has been set.

### GetExternalDomain

`func (o *InstanceContainer4) GetExternalDomain() string`

GetExternalDomain returns the ExternalDomain field if non-nil, zero value otherwise.

### GetExternalDomainOk

`func (o *InstanceContainer4) GetExternalDomainOk() (*string, bool)`

GetExternalDomainOk returns a tuple with the ExternalDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDomain

`func (o *InstanceContainer4) SetExternalDomain(v string)`

SetExternalDomain sets ExternalDomain field to given value.

### HasExternalDomain

`func (o *InstanceContainer4) HasExternalDomain() bool`

HasExternalDomain returns a boolean if a field has been set.

### GetExternalFqdn

`func (o *InstanceContainer4) GetExternalFqdn() string`

GetExternalFqdn returns the ExternalFqdn field if non-nil, zero value otherwise.

### GetExternalFqdnOk

`func (o *InstanceContainer4) GetExternalFqdnOk() (*string, bool)`

GetExternalFqdnOk returns a tuple with the ExternalFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFqdn

`func (o *InstanceContainer4) SetExternalFqdn(v string)`

SetExternalFqdn sets ExternalFqdn field to given value.

### HasExternalFqdn

`func (o *InstanceContainer4) HasExternalFqdn() bool`

HasExternalFqdn returns a boolean if a field has been set.

### GetAccountId

`func (o *InstanceContainer4) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *InstanceContainer4) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *InstanceContainer4) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *InstanceContainer4) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetInstance

`func (o *InstanceContainer4) GetInstance() InstanceContainer4Instance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *InstanceContainer4) GetInstanceOk() (*InstanceContainer4Instance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *InstanceContainer4) SetInstance(v InstanceContainer4Instance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *InstanceContainer4) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetContainerType

`func (o *InstanceContainer4) GetContainerType() InstanceContainer4ContainerType`

GetContainerType returns the ContainerType field if non-nil, zero value otherwise.

### GetContainerTypeOk

`func (o *InstanceContainer4) GetContainerTypeOk() (*InstanceContainer4ContainerType, bool)`

GetContainerTypeOk returns a tuple with the ContainerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerType

`func (o *InstanceContainer4) SetContainerType(v InstanceContainer4ContainerType)`

SetContainerType sets ContainerType field to given value.

### HasContainerType

`func (o *InstanceContainer4) HasContainerType() bool`

HasContainerType returns a boolean if a field has been set.

### GetContainerTypeSet

`func (o *InstanceContainer4) GetContainerTypeSet() InstanceContainer4ContainerTypeSet`

GetContainerTypeSet returns the ContainerTypeSet field if non-nil, zero value otherwise.

### GetContainerTypeSetOk

`func (o *InstanceContainer4) GetContainerTypeSetOk() (*InstanceContainer4ContainerTypeSet, bool)`

GetContainerTypeSetOk returns a tuple with the ContainerTypeSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTypeSet

`func (o *InstanceContainer4) SetContainerTypeSet(v InstanceContainer4ContainerTypeSet)`

SetContainerTypeSet sets ContainerTypeSet field to given value.

### HasContainerTypeSet

`func (o *InstanceContainer4) HasContainerTypeSet() bool`

HasContainerTypeSet returns a boolean if a field has been set.

### GetServer

`func (o *InstanceContainer4) GetServer() InstanceContainerServer4`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *InstanceContainer4) GetServerOk() (*InstanceContainerServer4, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *InstanceContainer4) SetServer(v InstanceContainerServer4)`

SetServer sets Server field to given value.

### HasServer

`func (o *InstanceContainer4) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetCloud

`func (o *InstanceContainer4) GetCloud() InstanceContainer4Cloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *InstanceContainer4) GetCloudOk() (*InstanceContainer4Cloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *InstanceContainer4) SetCloud(v InstanceContainer4Cloud)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *InstanceContainer4) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetPorts

`func (o *InstanceContainer4) GetPorts() []InstanceContainer1PortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InstanceContainer4) GetPortsOk() (*[]InstanceContainer1PortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InstanceContainer4) SetPorts(v []InstanceContainer1PortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InstanceContainer4) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetPlan

`func (o *InstanceContainer4) GetPlan() InstanceContainer4Plan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *InstanceContainer4) GetPlanOk() (*InstanceContainer4Plan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *InstanceContainer4) SetPlan(v InstanceContainer4Plan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *InstanceContainer4) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetDateCreated

`func (o *InstanceContainer4) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *InstanceContainer4) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *InstanceContainer4) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *InstanceContainer4) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *InstanceContainer4) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *InstanceContainer4) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *InstanceContainer4) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *InstanceContainer4) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetHostname

`func (o *InstanceContainer4) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *InstanceContainer4) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *InstanceContainer4) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *InstanceContainer4) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetVolumeCreated

`func (o *InstanceContainer4) GetVolumeCreated() bool`

GetVolumeCreated returns the VolumeCreated field if non-nil, zero value otherwise.

### GetVolumeCreatedOk

`func (o *InstanceContainer4) GetVolumeCreatedOk() (*bool, bool)`

GetVolumeCreatedOk returns a tuple with the VolumeCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeCreated

`func (o *InstanceContainer4) SetVolumeCreated(v bool)`

SetVolumeCreated sets VolumeCreated field to given value.

### HasVolumeCreated

`func (o *InstanceContainer4) HasVolumeCreated() bool`

HasVolumeCreated returns a boolean if a field has been set.

### GetContainerCreated

`func (o *InstanceContainer4) GetContainerCreated() bool`

GetContainerCreated returns the ContainerCreated field if non-nil, zero value otherwise.

### GetContainerCreatedOk

`func (o *InstanceContainer4) GetContainerCreatedOk() (*bool, bool)`

GetContainerCreatedOk returns a tuple with the ContainerCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerCreated

`func (o *InstanceContainer4) SetContainerCreated(v bool)`

SetContainerCreated sets ContainerCreated field to given value.

### HasContainerCreated

`func (o *InstanceContainer4) HasContainerCreated() bool`

HasContainerCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


