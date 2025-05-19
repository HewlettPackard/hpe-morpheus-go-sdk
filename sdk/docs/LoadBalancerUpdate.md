# LoadBalancerUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**Enabled** | Pointer to **bool** | Activate (true) or disable (false) | [optional] 
**Config** | Pointer to **map[string]interface{}** | Configuration object with parameters that vary by load balancer type. | [optional] 
**Visibility** | Pointer to **string** | private or public | [optional] [default to "public"]
**Tenants** | Pointer to [**[]CreateLoadBalancerRequestLoadBalancerTenantsInner**](CreateLoadBalancerRequestLoadBalancerTenantsInner.md) | Array of tenant account ids that are allowed access | [optional] 
**ResourcePermission** | Pointer to [**CreateLoadBalancerRequestLoadBalancerResourcePermission**](CreateLoadBalancerRequestLoadBalancerResourcePermission.md) |  | [optional] 

## Methods

### NewLoadBalancerUpdate

`func NewLoadBalancerUpdate() *LoadBalancerUpdate`

NewLoadBalancerUpdate instantiates a new LoadBalancerUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerUpdateWithDefaults

`func NewLoadBalancerUpdateWithDefaults() *LoadBalancerUpdate`

NewLoadBalancerUpdateWithDefaults instantiates a new LoadBalancerUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LoadBalancerUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoadBalancerUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoadBalancerUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoadBalancerUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *LoadBalancerUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LoadBalancerUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LoadBalancerUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LoadBalancerUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *LoadBalancerUpdate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LoadBalancerUpdate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LoadBalancerUpdate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *LoadBalancerUpdate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetConfig

`func (o *LoadBalancerUpdate) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *LoadBalancerUpdate) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *LoadBalancerUpdate) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *LoadBalancerUpdate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetVisibility

`func (o *LoadBalancerUpdate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *LoadBalancerUpdate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *LoadBalancerUpdate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *LoadBalancerUpdate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTenants

`func (o *LoadBalancerUpdate) GetTenants() []CreateLoadBalancerRequestLoadBalancerTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *LoadBalancerUpdate) GetTenantsOk() (*[]CreateLoadBalancerRequestLoadBalancerTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *LoadBalancerUpdate) SetTenants(v []CreateLoadBalancerRequestLoadBalancerTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *LoadBalancerUpdate) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermission

`func (o *LoadBalancerUpdate) GetResourcePermission() CreateLoadBalancerRequestLoadBalancerResourcePermission`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *LoadBalancerUpdate) GetResourcePermissionOk() (*CreateLoadBalancerRequestLoadBalancerResourcePermission, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *LoadBalancerUpdate) SetResourcePermission(v CreateLoadBalancerRequestLoadBalancerResourcePermission)`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *LoadBalancerUpdate) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


