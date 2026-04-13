# CreateLoadBalancerRequestLoadBalancer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Load Balancer Type Code | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**NetworkServerId** | Pointer to **int64** | Network Server ID | [optional] 
**Site** | Pointer to [**CreateLoadBalancerRequestLoadBalancerSite**](CreateLoadBalancerRequestLoadBalancerSite.md) |  | [optional] 
**Zone** | Pointer to [**CreateLoadBalancerRequestLoadBalancerZone**](CreateLoadBalancerRequestLoadBalancerZone.md) |  | [optional] 
**Config** | Pointer to [**CreateLoadBalancerRequestLoadBalancerConfig**](CreateLoadBalancerRequestLoadBalancerConfig.md) |  | [optional] 
**Visibility** | Pointer to **string** | private or public | [optional] [default to "public"]
**Tenants** | Pointer to [**[]CreateLoadBalancerRequestLoadBalancerTenantsInner**](CreateLoadBalancerRequestLoadBalancerTenantsInner.md) | Array of tenant account ids that are allowed access | [optional] 
**ResourcePermissions** | Pointer to [**CreateLoadBalancerRequestLoadBalancerResourcePermissions**](CreateLoadBalancerRequestLoadBalancerResourcePermissions.md) |  | [optional] 

## Methods

### NewCreateLoadBalancerRequestLoadBalancer

`func NewCreateLoadBalancerRequestLoadBalancer() *CreateLoadBalancerRequestLoadBalancer`

NewCreateLoadBalancerRequestLoadBalancer instantiates a new CreateLoadBalancerRequestLoadBalancer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoadBalancerRequestLoadBalancerWithDefaults

`func NewCreateLoadBalancerRequestLoadBalancerWithDefaults() *CreateLoadBalancerRequestLoadBalancer`

NewCreateLoadBalancerRequestLoadBalancerWithDefaults instantiates a new CreateLoadBalancerRequestLoadBalancer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateLoadBalancerRequestLoadBalancer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateLoadBalancerRequestLoadBalancer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateLoadBalancerRequestLoadBalancer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateLoadBalancerRequestLoadBalancer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *CreateLoadBalancerRequestLoadBalancer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLoadBalancerRequestLoadBalancer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLoadBalancerRequestLoadBalancer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateLoadBalancerRequestLoadBalancer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateLoadBalancerRequestLoadBalancer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateLoadBalancerRequestLoadBalancer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateLoadBalancerRequestLoadBalancer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateLoadBalancerRequestLoadBalancer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNetworkServerId

`func (o *CreateLoadBalancerRequestLoadBalancer) GetNetworkServerId() int64`

GetNetworkServerId returns the NetworkServerId field if non-nil, zero value otherwise.

### GetNetworkServerIdOk

`func (o *CreateLoadBalancerRequestLoadBalancer) GetNetworkServerIdOk() (*int64, bool)`

GetNetworkServerIdOk returns a tuple with the NetworkServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServerId

`func (o *CreateLoadBalancerRequestLoadBalancer) SetNetworkServerId(v int64)`

SetNetworkServerId sets NetworkServerId field to given value.

### HasNetworkServerId

`func (o *CreateLoadBalancerRequestLoadBalancer) HasNetworkServerId() bool`

HasNetworkServerId returns a boolean if a field has been set.

### GetSite

`func (o *CreateLoadBalancerRequestLoadBalancer) GetSite() CreateLoadBalancerRequestLoadBalancerSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *CreateLoadBalancerRequestLoadBalancer) GetSiteOk() (*CreateLoadBalancerRequestLoadBalancerSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *CreateLoadBalancerRequestLoadBalancer) SetSite(v CreateLoadBalancerRequestLoadBalancerSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *CreateLoadBalancerRequestLoadBalancer) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetZone

`func (o *CreateLoadBalancerRequestLoadBalancer) GetZone() CreateLoadBalancerRequestLoadBalancerZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *CreateLoadBalancerRequestLoadBalancer) GetZoneOk() (*CreateLoadBalancerRequestLoadBalancerZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *CreateLoadBalancerRequestLoadBalancer) SetZone(v CreateLoadBalancerRequestLoadBalancerZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *CreateLoadBalancerRequestLoadBalancer) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetConfig

`func (o *CreateLoadBalancerRequestLoadBalancer) GetConfig() CreateLoadBalancerRequestLoadBalancerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateLoadBalancerRequestLoadBalancer) GetConfigOk() (*CreateLoadBalancerRequestLoadBalancerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateLoadBalancerRequestLoadBalancer) SetConfig(v CreateLoadBalancerRequestLoadBalancerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateLoadBalancerRequestLoadBalancer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetVisibility

`func (o *CreateLoadBalancerRequestLoadBalancer) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *CreateLoadBalancerRequestLoadBalancer) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *CreateLoadBalancerRequestLoadBalancer) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *CreateLoadBalancerRequestLoadBalancer) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTenants

`func (o *CreateLoadBalancerRequestLoadBalancer) GetTenants() []CreateLoadBalancerRequestLoadBalancerTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *CreateLoadBalancerRequestLoadBalancer) GetTenantsOk() (*[]CreateLoadBalancerRequestLoadBalancerTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *CreateLoadBalancerRequestLoadBalancer) SetTenants(v []CreateLoadBalancerRequestLoadBalancerTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *CreateLoadBalancerRequestLoadBalancer) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *CreateLoadBalancerRequestLoadBalancer) GetResourcePermissions() CreateLoadBalancerRequestLoadBalancerResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *CreateLoadBalancerRequestLoadBalancer) GetResourcePermissionsOk() (*CreateLoadBalancerRequestLoadBalancerResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *CreateLoadBalancerRequestLoadBalancer) SetResourcePermissions(v CreateLoadBalancerRequestLoadBalancerResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *CreateLoadBalancerRequestLoadBalancer) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


