# GetNetworkEdgeCluster200ResponseNetworkEdgeCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**ProviderId** | Pointer to **string** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**GetNetworkEdgeCluster200ResponseNetworkEdgeClusterConfig**](GetNetworkEdgeCluster200ResponseNetworkEdgeClusterConfig.md) |  | [optional] 
**Owner** | Pointer to [**CreateLoadBalancerRequestLoadBalancerTenantsInner**](CreateLoadBalancerRequestLoadBalancerTenantsInner.md) |  | [optional] 
**NetworkServer** | Pointer to [**CreateLoadBalancerRequestLoadBalancerTenantsInner**](CreateLoadBalancerRequestLoadBalancerTenantsInner.md) |  | [optional] 
**Zone** | Pointer to [**CreateLoadBalancerRequestLoadBalancerTenantsInner**](CreateLoadBalancerRequestLoadBalancerTenantsInner.md) |  | [optional] 
**Tenants** | Pointer to [**[]SnapshotsInstance200ResponseSnapshotsInnerZone**](SnapshotsInstance200ResponseSnapshotsInnerZone.md) |  | [optional] 

## Methods

### NewGetNetworkEdgeCluster200ResponseNetworkEdgeCluster

`func NewGetNetworkEdgeCluster200ResponseNetworkEdgeCluster() *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster`

NewGetNetworkEdgeCluster200ResponseNetworkEdgeCluster instantiates a new GetNetworkEdgeCluster200ResponseNetworkEdgeCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkEdgeCluster200ResponseNetworkEdgeClusterWithDefaults

`func NewGetNetworkEdgeCluster200ResponseNetworkEdgeClusterWithDefaults() *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster`

NewGetNetworkEdgeCluster200ResponseNetworkEdgeClusterWithDefaults instantiates a new GetNetworkEdgeCluster200ResponseNetworkEdgeCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternalId

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetVisibility

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetProviderId

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetActive

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExternalId

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetConfig

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) GetConfig() GetNetworkEdgeCluster200ResponseNetworkEdgeClusterConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) GetConfigOk() (*GetNetworkEdgeCluster200ResponseNetworkEdgeClusterConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) SetConfig(v GetNetworkEdgeCluster200ResponseNetworkEdgeClusterConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetOwner

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) GetOwner() CreateLoadBalancerRequestLoadBalancerTenantsInner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) GetOwnerOk() (*CreateLoadBalancerRequestLoadBalancerTenantsInner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) SetOwner(v CreateLoadBalancerRequestLoadBalancerTenantsInner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetNetworkServer

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) GetNetworkServer() CreateLoadBalancerRequestLoadBalancerTenantsInner`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) GetNetworkServerOk() (*CreateLoadBalancerRequestLoadBalancerTenantsInner, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) SetNetworkServer(v CreateLoadBalancerRequestLoadBalancerTenantsInner)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.

### GetZone

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) GetZone() CreateLoadBalancerRequestLoadBalancerTenantsInner`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) GetZoneOk() (*CreateLoadBalancerRequestLoadBalancerTenantsInner, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) SetZone(v CreateLoadBalancerRequestLoadBalancerTenantsInner)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetTenants

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) GetTenants() []SnapshotsInstance200ResponseSnapshotsInnerZone`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) GetTenantsOk() (*[]SnapshotsInstance200ResponseSnapshotsInnerZone, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) SetTenants(v []SnapshotsInstance200ResponseSnapshotsInnerZone)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *GetNetworkEdgeCluster200ResponseNetworkEdgeCluster) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


