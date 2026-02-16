# InstanceContainer3

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
**Instance** | Pointer to [**InstanceContainer3Instance**](InstanceContainer3Instance.md) |  | [optional] 
**ContainerType** | Pointer to [**InstanceContainer3ContainerType**](InstanceContainer3ContainerType.md) |  | [optional] 
**ContainerTypeSet** | Pointer to [**InstanceContainer3ContainerTypeSet**](InstanceContainer3ContainerTypeSet.md) |  | [optional] 
**Server** | Pointer to [**InstanceContainerServer3**](InstanceContainerServer3.md) |  | [optional] 
**Cloud** | Pointer to [**InstanceContainer3Cloud**](InstanceContainer3Cloud.md) |  | [optional] 
**Ports** | Pointer to [**[]InstanceContainer1PortsInner**](InstanceContainer1PortsInner.md) |  | [optional] 
**Plan** | Pointer to [**InstanceContainer3Plan**](InstanceContainer3Plan.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**VolumeCreated** | Pointer to **bool** |  | [optional] 
**ContainerCreated** | Pointer to **bool** |  | [optional] 

## Methods

### NewInstanceContainer3

`func NewInstanceContainer3() *InstanceContainer3`

NewInstanceContainer3 instantiates a new InstanceContainer3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceContainer3WithDefaults

`func NewInstanceContainer3WithDefaults() *InstanceContainer3`

NewInstanceContainer3WithDefaults instantiates a new InstanceContainer3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceContainer3) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceContainer3) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceContainer3) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceContainer3) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *InstanceContainer3) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *InstanceContainer3) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *InstanceContainer3) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *InstanceContainer3) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *InstanceContainer3) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceContainer3) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceContainer3) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceContainer3) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIp

`func (o *InstanceContainer3) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *InstanceContainer3) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *InstanceContainer3) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *InstanceContainer3) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetInternalIp

`func (o *InstanceContainer3) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *InstanceContainer3) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *InstanceContainer3) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *InstanceContainer3) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### GetInternalHostname

`func (o *InstanceContainer3) GetInternalHostname() string`

GetInternalHostname returns the InternalHostname field if non-nil, zero value otherwise.

### GetInternalHostnameOk

`func (o *InstanceContainer3) GetInternalHostnameOk() (*string, bool)`

GetInternalHostnameOk returns a tuple with the InternalHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalHostname

`func (o *InstanceContainer3) SetInternalHostname(v string)`

SetInternalHostname sets InternalHostname field to given value.

### HasInternalHostname

`func (o *InstanceContainer3) HasInternalHostname() bool`

HasInternalHostname returns a boolean if a field has been set.

### GetExternalHostname

`func (o *InstanceContainer3) GetExternalHostname() string`

GetExternalHostname returns the ExternalHostname field if non-nil, zero value otherwise.

### GetExternalHostnameOk

`func (o *InstanceContainer3) GetExternalHostnameOk() (*string, bool)`

GetExternalHostnameOk returns a tuple with the ExternalHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalHostname

`func (o *InstanceContainer3) SetExternalHostname(v string)`

SetExternalHostname sets ExternalHostname field to given value.

### HasExternalHostname

`func (o *InstanceContainer3) HasExternalHostname() bool`

HasExternalHostname returns a boolean if a field has been set.

### GetExternalDomain

`func (o *InstanceContainer3) GetExternalDomain() string`

GetExternalDomain returns the ExternalDomain field if non-nil, zero value otherwise.

### GetExternalDomainOk

`func (o *InstanceContainer3) GetExternalDomainOk() (*string, bool)`

GetExternalDomainOk returns a tuple with the ExternalDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDomain

`func (o *InstanceContainer3) SetExternalDomain(v string)`

SetExternalDomain sets ExternalDomain field to given value.

### HasExternalDomain

`func (o *InstanceContainer3) HasExternalDomain() bool`

HasExternalDomain returns a boolean if a field has been set.

### GetExternalFqdn

`func (o *InstanceContainer3) GetExternalFqdn() string`

GetExternalFqdn returns the ExternalFqdn field if non-nil, zero value otherwise.

### GetExternalFqdnOk

`func (o *InstanceContainer3) GetExternalFqdnOk() (*string, bool)`

GetExternalFqdnOk returns a tuple with the ExternalFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFqdn

`func (o *InstanceContainer3) SetExternalFqdn(v string)`

SetExternalFqdn sets ExternalFqdn field to given value.

### HasExternalFqdn

`func (o *InstanceContainer3) HasExternalFqdn() bool`

HasExternalFqdn returns a boolean if a field has been set.

### GetAccountId

`func (o *InstanceContainer3) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *InstanceContainer3) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *InstanceContainer3) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *InstanceContainer3) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetInstance

`func (o *InstanceContainer3) GetInstance() InstanceContainer3Instance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *InstanceContainer3) GetInstanceOk() (*InstanceContainer3Instance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *InstanceContainer3) SetInstance(v InstanceContainer3Instance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *InstanceContainer3) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetContainerType

`func (o *InstanceContainer3) GetContainerType() InstanceContainer3ContainerType`

GetContainerType returns the ContainerType field if non-nil, zero value otherwise.

### GetContainerTypeOk

`func (o *InstanceContainer3) GetContainerTypeOk() (*InstanceContainer3ContainerType, bool)`

GetContainerTypeOk returns a tuple with the ContainerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerType

`func (o *InstanceContainer3) SetContainerType(v InstanceContainer3ContainerType)`

SetContainerType sets ContainerType field to given value.

### HasContainerType

`func (o *InstanceContainer3) HasContainerType() bool`

HasContainerType returns a boolean if a field has been set.

### GetContainerTypeSet

`func (o *InstanceContainer3) GetContainerTypeSet() InstanceContainer3ContainerTypeSet`

GetContainerTypeSet returns the ContainerTypeSet field if non-nil, zero value otherwise.

### GetContainerTypeSetOk

`func (o *InstanceContainer3) GetContainerTypeSetOk() (*InstanceContainer3ContainerTypeSet, bool)`

GetContainerTypeSetOk returns a tuple with the ContainerTypeSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTypeSet

`func (o *InstanceContainer3) SetContainerTypeSet(v InstanceContainer3ContainerTypeSet)`

SetContainerTypeSet sets ContainerTypeSet field to given value.

### HasContainerTypeSet

`func (o *InstanceContainer3) HasContainerTypeSet() bool`

HasContainerTypeSet returns a boolean if a field has been set.

### GetServer

`func (o *InstanceContainer3) GetServer() InstanceContainerServer3`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *InstanceContainer3) GetServerOk() (*InstanceContainerServer3, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *InstanceContainer3) SetServer(v InstanceContainerServer3)`

SetServer sets Server field to given value.

### HasServer

`func (o *InstanceContainer3) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetCloud

`func (o *InstanceContainer3) GetCloud() InstanceContainer3Cloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *InstanceContainer3) GetCloudOk() (*InstanceContainer3Cloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *InstanceContainer3) SetCloud(v InstanceContainer3Cloud)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *InstanceContainer3) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetPorts

`func (o *InstanceContainer3) GetPorts() []InstanceContainer1PortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InstanceContainer3) GetPortsOk() (*[]InstanceContainer1PortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InstanceContainer3) SetPorts(v []InstanceContainer1PortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InstanceContainer3) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetPlan

`func (o *InstanceContainer3) GetPlan() InstanceContainer3Plan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *InstanceContainer3) GetPlanOk() (*InstanceContainer3Plan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *InstanceContainer3) SetPlan(v InstanceContainer3Plan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *InstanceContainer3) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetDateCreated

`func (o *InstanceContainer3) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *InstanceContainer3) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *InstanceContainer3) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *InstanceContainer3) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *InstanceContainer3) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *InstanceContainer3) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *InstanceContainer3) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *InstanceContainer3) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetHostname

`func (o *InstanceContainer3) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *InstanceContainer3) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *InstanceContainer3) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *InstanceContainer3) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetVolumeCreated

`func (o *InstanceContainer3) GetVolumeCreated() bool`

GetVolumeCreated returns the VolumeCreated field if non-nil, zero value otherwise.

### GetVolumeCreatedOk

`func (o *InstanceContainer3) GetVolumeCreatedOk() (*bool, bool)`

GetVolumeCreatedOk returns a tuple with the VolumeCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeCreated

`func (o *InstanceContainer3) SetVolumeCreated(v bool)`

SetVolumeCreated sets VolumeCreated field to given value.

### HasVolumeCreated

`func (o *InstanceContainer3) HasVolumeCreated() bool`

HasVolumeCreated returns a boolean if a field has been set.

### GetContainerCreated

`func (o *InstanceContainer3) GetContainerCreated() bool`

GetContainerCreated returns the ContainerCreated field if non-nil, zero value otherwise.

### GetContainerCreatedOk

`func (o *InstanceContainer3) GetContainerCreatedOk() (*bool, bool)`

GetContainerCreatedOk returns a tuple with the ContainerCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerCreated

`func (o *InstanceContainer3) SetContainerCreated(v bool)`

SetContainerCreated sets ContainerCreated field to given value.

### HasContainerCreated

`func (o *InstanceContainer3) HasContainerCreated() bool`

HasContainerCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


