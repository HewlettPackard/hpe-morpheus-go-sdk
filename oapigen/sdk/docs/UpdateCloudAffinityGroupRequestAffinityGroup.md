# UpdateCloudAffinityGroupRequestAffinityGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**AffinityType** | Pointer to **string** | Affinity Type | [optional] 
**Active** | Pointer to **bool** | Active | [optional] 
**Pool** | Pointer to [**UpdateCloudAffinityGroupRequestAffinityGroupPool**](UpdateCloudAffinityGroupRequestAffinityGroupPool.md) |  | [optional] 
**Servers** | Pointer to **[]int32** | List of Server IDs to include in the Affinity Group | [optional] 
**Visibility** | Pointer to **string** | Visibility - Set to public to allow all tenants | [optional] [default to "private"]
**Tenants** | Pointer to [**[]UpdateCloudAffinityGroupRequestAffinityGroupTenantsInner**](UpdateCloudAffinityGroupRequestAffinityGroupTenantsInner.md) | Array of tenant account ids that are allowed access | [optional] 
**ResourcePermissions** | Pointer to [**UpdateCloudAffinityGroupRequestAffinityGroupResourcePermissions**](UpdateCloudAffinityGroupRequestAffinityGroupResourcePermissions.md) |  | [optional] 

## Methods

### NewUpdateCloudAffinityGroupRequestAffinityGroup

`func NewUpdateCloudAffinityGroupRequestAffinityGroup() *UpdateCloudAffinityGroupRequestAffinityGroup`

NewUpdateCloudAffinityGroupRequestAffinityGroup instantiates a new UpdateCloudAffinityGroupRequestAffinityGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetName

`func (o *UpdateCloudAffinityGroupRequestAffinityGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCloudAffinityGroupRequestAffinityGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCloudAffinityGroupRequestAffinityGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCloudAffinityGroupRequestAffinityGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAffinityType

`func (o *UpdateCloudAffinityGroupRequestAffinityGroup) GetAffinityType() string`

GetAffinityType returns the AffinityType field if non-nil, zero value otherwise.

### GetAffinityTypeOk

`func (o *UpdateCloudAffinityGroupRequestAffinityGroup) GetAffinityTypeOk() (*string, bool)`

GetAffinityTypeOk returns a tuple with the AffinityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinityType

`func (o *UpdateCloudAffinityGroupRequestAffinityGroup) SetAffinityType(v string)`

SetAffinityType sets AffinityType field to given value.

### HasAffinityType

`func (o *UpdateCloudAffinityGroupRequestAffinityGroup) HasAffinityType() bool`

HasAffinityType returns a boolean if a field has been set.

### GetActive

`func (o *UpdateCloudAffinityGroupRequestAffinityGroup) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateCloudAffinityGroupRequestAffinityGroup) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateCloudAffinityGroupRequestAffinityGroup) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateCloudAffinityGroupRequestAffinityGroup) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPool

`func (o *UpdateCloudAffinityGroupRequestAffinityGroup) GetPool() UpdateCloudAffinityGroupRequestAffinityGroupPool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *UpdateCloudAffinityGroupRequestAffinityGroup) GetPoolOk() (*UpdateCloudAffinityGroupRequestAffinityGroupPool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *UpdateCloudAffinityGroupRequestAffinityGroup) SetPool(v UpdateCloudAffinityGroupRequestAffinityGroupPool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *UpdateCloudAffinityGroupRequestAffinityGroup) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetServers

`func (o *UpdateCloudAffinityGroupRequestAffinityGroup) GetServers() []int32`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *UpdateCloudAffinityGroupRequestAffinityGroup) GetServersOk() (*[]int32, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *UpdateCloudAffinityGroupRequestAffinityGroup) SetServers(v []int32)`

SetServers sets Servers field to given value.

### HasServers

`func (o *UpdateCloudAffinityGroupRequestAffinityGroup) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateCloudAffinityGroupRequestAffinityGroup) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateCloudAffinityGroupRequestAffinityGroup) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateCloudAffinityGroupRequestAffinityGroup) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateCloudAffinityGroupRequestAffinityGroup) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTenants

`func (o *UpdateCloudAffinityGroupRequestAffinityGroup) GetTenants() []UpdateCloudAffinityGroupRequestAffinityGroupTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *UpdateCloudAffinityGroupRequestAffinityGroup) GetTenantsOk() (*[]UpdateCloudAffinityGroupRequestAffinityGroupTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *UpdateCloudAffinityGroupRequestAffinityGroup) SetTenants(v []UpdateCloudAffinityGroupRequestAffinityGroupTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *UpdateCloudAffinityGroupRequestAffinityGroup) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *UpdateCloudAffinityGroupRequestAffinityGroup) GetResourcePermissions() UpdateCloudAffinityGroupRequestAffinityGroupResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *UpdateCloudAffinityGroupRequestAffinityGroup) GetResourcePermissionsOk() (*UpdateCloudAffinityGroupRequestAffinityGroupResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *UpdateCloudAffinityGroupRequestAffinityGroup) SetResourcePermissions(v UpdateCloudAffinityGroupRequestAffinityGroupResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *UpdateCloudAffinityGroupRequestAffinityGroup) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


