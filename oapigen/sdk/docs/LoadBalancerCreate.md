# LoadBalancerCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Load Balancer Type Code | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**NetworkServerId** | Pointer to **int64** | Network Server ID | [optional] 
**Site** | Pointer to [**LoadBalancerCreateSite**](LoadBalancerCreateSite.md) |  | [optional] 
**Zone** | Pointer to [**LoadBalancerCreateZone**](LoadBalancerCreateZone.md) |  | [optional] 
**Config** | Pointer to [**LoadBalancerCreateConfig**](LoadBalancerCreateConfig.md) |  | [optional] 
**Visibility** | Pointer to **string** | private or public | [optional] [default to "public"]
**Tenants** | Pointer to [**[]LoadBalancerCreateTenantsInner**](LoadBalancerCreateTenantsInner.md) | Array of tenant account ids that are allowed access | [optional] 
**ResourcePermissions** | Pointer to [**LoadBalancerCreateResourcePermissions**](LoadBalancerCreateResourcePermissions.md) |  | [optional] 

## Methods

### NewLoadBalancerCreate

`func NewLoadBalancerCreate() *LoadBalancerCreate`

NewLoadBalancerCreate instantiates a new LoadBalancerCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerCreateWithDefaults

`func NewLoadBalancerCreateWithDefaults() *LoadBalancerCreate`

NewLoadBalancerCreateWithDefaults instantiates a new LoadBalancerCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LoadBalancerCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LoadBalancerCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LoadBalancerCreate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LoadBalancerCreate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *LoadBalancerCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoadBalancerCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoadBalancerCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoadBalancerCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *LoadBalancerCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LoadBalancerCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LoadBalancerCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LoadBalancerCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNetworkServerId

`func (o *LoadBalancerCreate) GetNetworkServerId() int64`

GetNetworkServerId returns the NetworkServerId field if non-nil, zero value otherwise.

### GetNetworkServerIdOk

`func (o *LoadBalancerCreate) GetNetworkServerIdOk() (*int64, bool)`

GetNetworkServerIdOk returns a tuple with the NetworkServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServerId

`func (o *LoadBalancerCreate) SetNetworkServerId(v int64)`

SetNetworkServerId sets NetworkServerId field to given value.

### HasNetworkServerId

`func (o *LoadBalancerCreate) HasNetworkServerId() bool`

HasNetworkServerId returns a boolean if a field has been set.

### GetSite

`func (o *LoadBalancerCreate) GetSite() LoadBalancerCreateSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *LoadBalancerCreate) GetSiteOk() (*LoadBalancerCreateSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *LoadBalancerCreate) SetSite(v LoadBalancerCreateSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *LoadBalancerCreate) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetZone

`func (o *LoadBalancerCreate) GetZone() LoadBalancerCreateZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *LoadBalancerCreate) GetZoneOk() (*LoadBalancerCreateZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *LoadBalancerCreate) SetZone(v LoadBalancerCreateZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *LoadBalancerCreate) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetConfig

`func (o *LoadBalancerCreate) GetConfig() LoadBalancerCreateConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *LoadBalancerCreate) GetConfigOk() (*LoadBalancerCreateConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *LoadBalancerCreate) SetConfig(v LoadBalancerCreateConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *LoadBalancerCreate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetVisibility

`func (o *LoadBalancerCreate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *LoadBalancerCreate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *LoadBalancerCreate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *LoadBalancerCreate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTenants

`func (o *LoadBalancerCreate) GetTenants() []LoadBalancerCreateTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *LoadBalancerCreate) GetTenantsOk() (*[]LoadBalancerCreateTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *LoadBalancerCreate) SetTenants(v []LoadBalancerCreateTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *LoadBalancerCreate) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *LoadBalancerCreate) GetResourcePermissions() LoadBalancerCreateResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *LoadBalancerCreate) GetResourcePermissionsOk() (*LoadBalancerCreateResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *LoadBalancerCreate) SetResourcePermissions(v LoadBalancerCreateResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *LoadBalancerCreate) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


