# AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner

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
**Instance** | Pointer to [**ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance**](ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance.md) |  | [optional] 
**ContainerType** | Pointer to [**ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType**](ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType.md) |  | [optional] 
**ContainerTypeSet** | Pointer to [**ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerTypeSet**](ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerTypeSet.md) |  | [optional] 
**Server** | Pointer to [**AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer**](AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer.md) |  | [optional] 
**Cloud** | Pointer to [**ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance**](ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance.md) |  | [optional] 
**Ports** | Pointer to [**[]AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerPortsInner**](AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerPortsInner.md) |  | [optional] 
**Plan** | Pointer to [**ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType**](ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**VolumeCreated** | Pointer to **bool** |  | [optional] 
**ContainerCreated** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner

`func NewAddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner() *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner`

NewAddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner instantiates a new AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerWithDefaults

`func NewAddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerWithDefaults() *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner`

NewAddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerWithDefaults instantiates a new AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIp

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetInternalIp

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### GetInternalHostname

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetInternalHostname() string`

GetInternalHostname returns the InternalHostname field if non-nil, zero value otherwise.

### GetInternalHostnameOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetInternalHostnameOk() (*string, bool)`

GetInternalHostnameOk returns a tuple with the InternalHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalHostname

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) SetInternalHostname(v string)`

SetInternalHostname sets InternalHostname field to given value.

### HasInternalHostname

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) HasInternalHostname() bool`

HasInternalHostname returns a boolean if a field has been set.

### GetExternalHostname

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetExternalHostname() string`

GetExternalHostname returns the ExternalHostname field if non-nil, zero value otherwise.

### GetExternalHostnameOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetExternalHostnameOk() (*string, bool)`

GetExternalHostnameOk returns a tuple with the ExternalHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalHostname

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) SetExternalHostname(v string)`

SetExternalHostname sets ExternalHostname field to given value.

### HasExternalHostname

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) HasExternalHostname() bool`

HasExternalHostname returns a boolean if a field has been set.

### GetExternalDomain

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetExternalDomain() string`

GetExternalDomain returns the ExternalDomain field if non-nil, zero value otherwise.

### GetExternalDomainOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetExternalDomainOk() (*string, bool)`

GetExternalDomainOk returns a tuple with the ExternalDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDomain

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) SetExternalDomain(v string)`

SetExternalDomain sets ExternalDomain field to given value.

### HasExternalDomain

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) HasExternalDomain() bool`

HasExternalDomain returns a boolean if a field has been set.

### GetExternalFqdn

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetExternalFqdn() string`

GetExternalFqdn returns the ExternalFqdn field if non-nil, zero value otherwise.

### GetExternalFqdnOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetExternalFqdnOk() (*string, bool)`

GetExternalFqdnOk returns a tuple with the ExternalFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFqdn

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) SetExternalFqdn(v string)`

SetExternalFqdn sets ExternalFqdn field to given value.

### HasExternalFqdn

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) HasExternalFqdn() bool`

HasExternalFqdn returns a boolean if a field has been set.

### GetAccountId

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetInstance

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetInstance() ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetInstanceOk() (*ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) SetInstance(v ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetContainerType

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetContainerType() ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType`

GetContainerType returns the ContainerType field if non-nil, zero value otherwise.

### GetContainerTypeOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetContainerTypeOk() (*ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType, bool)`

GetContainerTypeOk returns a tuple with the ContainerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerType

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) SetContainerType(v ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType)`

SetContainerType sets ContainerType field to given value.

### HasContainerType

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) HasContainerType() bool`

HasContainerType returns a boolean if a field has been set.

### GetContainerTypeSet

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetContainerTypeSet() ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerTypeSet`

GetContainerTypeSet returns the ContainerTypeSet field if non-nil, zero value otherwise.

### GetContainerTypeSetOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetContainerTypeSetOk() (*ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerTypeSet, bool)`

GetContainerTypeSetOk returns a tuple with the ContainerTypeSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTypeSet

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) SetContainerTypeSet(v ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerTypeSet)`

SetContainerTypeSet sets ContainerTypeSet field to given value.

### HasContainerTypeSet

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) HasContainerTypeSet() bool`

HasContainerTypeSet returns a boolean if a field has been set.

### GetServer

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetServer() AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetServerOk() (*AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) SetServer(v AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer)`

SetServer sets Server field to given value.

### HasServer

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetCloud

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetCloud() ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetCloudOk() (*ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) SetCloud(v ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetPorts

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetPorts() []AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerPortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetPortsOk() (*[]AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerPortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) SetPorts(v []AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerPortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetPlan

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetPlan() ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetPlanOk() (*ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) SetPlan(v ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetDateCreated

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetHostname

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetVolumeCreated

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetVolumeCreated() bool`

GetVolumeCreated returns the VolumeCreated field if non-nil, zero value otherwise.

### GetVolumeCreatedOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetVolumeCreatedOk() (*bool, bool)`

GetVolumeCreatedOk returns a tuple with the VolumeCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeCreated

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) SetVolumeCreated(v bool)`

SetVolumeCreated sets VolumeCreated field to given value.

### HasVolumeCreated

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) HasVolumeCreated() bool`

HasVolumeCreated returns a boolean if a field has been set.

### GetContainerCreated

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetContainerCreated() bool`

GetContainerCreated returns the ContainerCreated field if non-nil, zero value otherwise.

### GetContainerCreatedOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) GetContainerCreatedOk() (*bool, bool)`

GetContainerCreatedOk returns a tuple with the ContainerCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerCreated

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) SetContainerCreated(v bool)`

SetContainerCreated sets ContainerCreated field to given value.

### HasContainerCreated

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInner) HasContainerCreated() bool`

HasContainerCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


