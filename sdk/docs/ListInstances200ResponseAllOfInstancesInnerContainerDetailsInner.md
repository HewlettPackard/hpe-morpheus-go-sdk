# ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner

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
**Server** | Pointer to [**ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer**](ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer.md) |  | [optional] 
**Cloud** | Pointer to [**ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance**](ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance.md) |  | [optional] 
**Ports** | Pointer to [**[]ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerPortsInner**](ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerPortsInner.md) |  | [optional] 
**Plan** | Pointer to [**ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType**](ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**VolumeCreated** | Pointer to **bool** |  | [optional] 
**ContainerCreated** | Pointer to **bool** |  | [optional] 

## Methods

### NewListInstances200ResponseAllOfInstancesInnerContainerDetailsInner

`func NewListInstances200ResponseAllOfInstancesInnerContainerDetailsInner() *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner`

NewListInstances200ResponseAllOfInstancesInnerContainerDetailsInner instantiates a new ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerWithDefaults

`func NewListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerWithDefaults() *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner`

NewListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerWithDefaults instantiates a new ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIp

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetInternalIp

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### GetInternalHostname

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetInternalHostname() string`

GetInternalHostname returns the InternalHostname field if non-nil, zero value otherwise.

### GetInternalHostnameOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetInternalHostnameOk() (*string, bool)`

GetInternalHostnameOk returns a tuple with the InternalHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalHostname

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) SetInternalHostname(v string)`

SetInternalHostname sets InternalHostname field to given value.

### HasInternalHostname

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) HasInternalHostname() bool`

HasInternalHostname returns a boolean if a field has been set.

### GetExternalHostname

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetExternalHostname() string`

GetExternalHostname returns the ExternalHostname field if non-nil, zero value otherwise.

### GetExternalHostnameOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetExternalHostnameOk() (*string, bool)`

GetExternalHostnameOk returns a tuple with the ExternalHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalHostname

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) SetExternalHostname(v string)`

SetExternalHostname sets ExternalHostname field to given value.

### HasExternalHostname

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) HasExternalHostname() bool`

HasExternalHostname returns a boolean if a field has been set.

### GetExternalDomain

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetExternalDomain() string`

GetExternalDomain returns the ExternalDomain field if non-nil, zero value otherwise.

### GetExternalDomainOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetExternalDomainOk() (*string, bool)`

GetExternalDomainOk returns a tuple with the ExternalDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDomain

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) SetExternalDomain(v string)`

SetExternalDomain sets ExternalDomain field to given value.

### HasExternalDomain

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) HasExternalDomain() bool`

HasExternalDomain returns a boolean if a field has been set.

### GetExternalFqdn

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetExternalFqdn() string`

GetExternalFqdn returns the ExternalFqdn field if non-nil, zero value otherwise.

### GetExternalFqdnOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetExternalFqdnOk() (*string, bool)`

GetExternalFqdnOk returns a tuple with the ExternalFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFqdn

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) SetExternalFqdn(v string)`

SetExternalFqdn sets ExternalFqdn field to given value.

### HasExternalFqdn

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) HasExternalFqdn() bool`

HasExternalFqdn returns a boolean if a field has been set.

### GetAccountId

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetInstance

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetInstance() ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetInstanceOk() (*ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) SetInstance(v ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetContainerType

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetContainerType() ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType`

GetContainerType returns the ContainerType field if non-nil, zero value otherwise.

### GetContainerTypeOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetContainerTypeOk() (*ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType, bool)`

GetContainerTypeOk returns a tuple with the ContainerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerType

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) SetContainerType(v ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType)`

SetContainerType sets ContainerType field to given value.

### HasContainerType

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) HasContainerType() bool`

HasContainerType returns a boolean if a field has been set.

### GetContainerTypeSet

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetContainerTypeSet() ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerTypeSet`

GetContainerTypeSet returns the ContainerTypeSet field if non-nil, zero value otherwise.

### GetContainerTypeSetOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetContainerTypeSetOk() (*ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerTypeSet, bool)`

GetContainerTypeSetOk returns a tuple with the ContainerTypeSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTypeSet

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) SetContainerTypeSet(v ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerTypeSet)`

SetContainerTypeSet sets ContainerTypeSet field to given value.

### HasContainerTypeSet

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) HasContainerTypeSet() bool`

HasContainerTypeSet returns a boolean if a field has been set.

### GetServer

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetServer() ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetServerOk() (*ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) SetServer(v ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer)`

SetServer sets Server field to given value.

### HasServer

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetCloud

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetCloud() ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetCloudOk() (*ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) SetCloud(v ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetPorts

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetPorts() []ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerPortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetPortsOk() (*[]ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerPortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) SetPorts(v []ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerPortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetPlan

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetPlan() ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetPlanOk() (*ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) SetPlan(v ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetHostname

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetVolumeCreated

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetVolumeCreated() bool`

GetVolumeCreated returns the VolumeCreated field if non-nil, zero value otherwise.

### GetVolumeCreatedOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetVolumeCreatedOk() (*bool, bool)`

GetVolumeCreatedOk returns a tuple with the VolumeCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeCreated

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) SetVolumeCreated(v bool)`

SetVolumeCreated sets VolumeCreated field to given value.

### HasVolumeCreated

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) HasVolumeCreated() bool`

HasVolumeCreated returns a boolean if a field has been set.

### GetContainerCreated

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetContainerCreated() bool`

GetContainerCreated returns the ContainerCreated field if non-nil, zero value otherwise.

### GetContainerCreatedOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) GetContainerCreatedOk() (*bool, bool)`

GetContainerCreatedOk returns a tuple with the ContainerCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerCreated

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) SetContainerCreated(v bool)`

SetContainerCreated sets ContainerCreated field to given value.

### HasContainerCreated

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInner) HasContainerCreated() bool`

HasContainerCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


