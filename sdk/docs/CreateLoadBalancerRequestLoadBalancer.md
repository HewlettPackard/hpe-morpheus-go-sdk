# CreateLoadBalancerRequestLoadBalancer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**NetworkServerId** | Pointer to **int64** | Network Server ID | [optional] 
**Config** | Pointer to **map[string]interface{}** | Configuration object with parameters that vary by load balancer type. | [optional] 
**Visibility** | Pointer to **string** | private or public | [optional] [default to "public"]
**Tenants** | Pointer to [**[]CreateLoadBalancerRequestLoadBalancerTenantsInner**](CreateLoadBalancerRequestLoadBalancerTenantsInner.md) | Array of tenant account ids that are allowed access | [optional] 
**ResourcePermission** | Pointer to [**CreateLoadBalancerRequestLoadBalancerResourcePermission**](CreateLoadBalancerRequestLoadBalancerResourcePermission.md) |  | [optional] 

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

### GetConfig

`func (o *CreateLoadBalancerRequestLoadBalancer) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateLoadBalancerRequestLoadBalancer) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateLoadBalancerRequestLoadBalancer) SetConfig(v map[string]interface{})`

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

### GetResourcePermission

`func (o *CreateLoadBalancerRequestLoadBalancer) GetResourcePermission() CreateLoadBalancerRequestLoadBalancerResourcePermission`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *CreateLoadBalancerRequestLoadBalancer) GetResourcePermissionOk() (*CreateLoadBalancerRequestLoadBalancerResourcePermission, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *CreateLoadBalancerRequestLoadBalancer) SetResourcePermission(v CreateLoadBalancerRequestLoadBalancerResourcePermission)`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *CreateLoadBalancerRequestLoadBalancer) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


