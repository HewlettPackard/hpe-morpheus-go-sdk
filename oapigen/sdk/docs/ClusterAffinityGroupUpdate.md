# ClusterAffinityGroupUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**AffinityType** | Pointer to **string** | Affinity Type | [optional] 
**Active** | Pointer to **bool** | Active | [optional] 
**Pool** | Pointer to [**ZoneAffinityGroupCreatePool**](ZoneAffinityGroupCreatePool.md) |  | [optional] 
**Servers** | Pointer to **[]int32** | List of Server IDs to include in the Affinity Group | [optional] 
**Visibility** | Pointer to **string** | Visibility - Set to public to allow all tenants | [optional] [default to "private"]
**Tenants** | Pointer to [**[]ZoneAffinityGroupCreateTenantsInner**](ZoneAffinityGroupCreateTenantsInner.md) | Array of tenant account ids that are allowed access | [optional] 
**ResourcePermissions** | Pointer to [**ZoneAffinityGroupCreateResourcePermissions**](ZoneAffinityGroupCreateResourcePermissions.md) |  | [optional] 

## Methods

### NewClusterAffinityGroupUpdate

`func NewClusterAffinityGroupUpdate() *ClusterAffinityGroupUpdate`

NewClusterAffinityGroupUpdate instantiates a new ClusterAffinityGroupUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAffinityGroupUpdateWithDefaults

`func NewClusterAffinityGroupUpdateWithDefaults() *ClusterAffinityGroupUpdate`

NewClusterAffinityGroupUpdateWithDefaults instantiates a new ClusterAffinityGroupUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClusterAffinityGroupUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterAffinityGroupUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterAffinityGroupUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterAffinityGroupUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAffinityType

`func (o *ClusterAffinityGroupUpdate) GetAffinityType() string`

GetAffinityType returns the AffinityType field if non-nil, zero value otherwise.

### GetAffinityTypeOk

`func (o *ClusterAffinityGroupUpdate) GetAffinityTypeOk() (*string, bool)`

GetAffinityTypeOk returns a tuple with the AffinityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinityType

`func (o *ClusterAffinityGroupUpdate) SetAffinityType(v string)`

SetAffinityType sets AffinityType field to given value.

### HasAffinityType

`func (o *ClusterAffinityGroupUpdate) HasAffinityType() bool`

HasAffinityType returns a boolean if a field has been set.

### GetActive

`func (o *ClusterAffinityGroupUpdate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ClusterAffinityGroupUpdate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ClusterAffinityGroupUpdate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ClusterAffinityGroupUpdate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPool

`func (o *ClusterAffinityGroupUpdate) GetPool() ZoneAffinityGroupCreatePool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *ClusterAffinityGroupUpdate) GetPoolOk() (*ZoneAffinityGroupCreatePool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *ClusterAffinityGroupUpdate) SetPool(v ZoneAffinityGroupCreatePool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *ClusterAffinityGroupUpdate) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetServers

`func (o *ClusterAffinityGroupUpdate) GetServers() []int32`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *ClusterAffinityGroupUpdate) GetServersOk() (*[]int32, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *ClusterAffinityGroupUpdate) SetServers(v []int32)`

SetServers sets Servers field to given value.

### HasServers

`func (o *ClusterAffinityGroupUpdate) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetVisibility

`func (o *ClusterAffinityGroupUpdate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ClusterAffinityGroupUpdate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ClusterAffinityGroupUpdate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ClusterAffinityGroupUpdate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTenants

`func (o *ClusterAffinityGroupUpdate) GetTenants() []ZoneAffinityGroupCreateTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ClusterAffinityGroupUpdate) GetTenantsOk() (*[]ZoneAffinityGroupCreateTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ClusterAffinityGroupUpdate) SetTenants(v []ZoneAffinityGroupCreateTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ClusterAffinityGroupUpdate) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *ClusterAffinityGroupUpdate) GetResourcePermissions() ZoneAffinityGroupCreateResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *ClusterAffinityGroupUpdate) GetResourcePermissionsOk() (*ZoneAffinityGroupCreateResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *ClusterAffinityGroupUpdate) SetResourcePermissions(v ZoneAffinityGroupCreateResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *ClusterAffinityGroupUpdate) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


