# ClusterAffinityGroupCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**AffinityType** | Pointer to **string** | Affinity Type | [optional] 
**Active** | Pointer to **bool** | Active | [optional] 
**Pool** | Pointer to [**ClusterAffinityGroupCreatePool**](ClusterAffinityGroupCreatePool.md) |  | [optional] 
**Servers** | Pointer to **[]int32** | List of Server IDs to include in the Affinity Group | [optional] 
**Visibility** | Pointer to **string** | Visibility - Set to public to allow all tenants | [optional] [default to "private"]
**Tenants** | Pointer to [**[]ClusterAffinityGroupCreateTenantsInner**](ClusterAffinityGroupCreateTenantsInner.md) | Array of tenant account ids that are allowed access | [optional] 
**ResourcePermissions** | Pointer to [**ClusterAffinityGroupCreateResourcePermissions**](ClusterAffinityGroupCreateResourcePermissions.md) |  | [optional] 

## Methods

### NewClusterAffinityGroupCreate

`func NewClusterAffinityGroupCreate() *ClusterAffinityGroupCreate`

NewClusterAffinityGroupCreate instantiates a new ClusterAffinityGroupCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAffinityGroupCreateWithDefaults

`func NewClusterAffinityGroupCreateWithDefaults() *ClusterAffinityGroupCreate`

NewClusterAffinityGroupCreateWithDefaults instantiates a new ClusterAffinityGroupCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterAffinityGroupCreate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterAffinityGroupCreate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterAffinityGroupCreate) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterAffinityGroupCreate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ClusterAffinityGroupCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterAffinityGroupCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterAffinityGroupCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterAffinityGroupCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAffinityType

`func (o *ClusterAffinityGroupCreate) GetAffinityType() string`

GetAffinityType returns the AffinityType field if non-nil, zero value otherwise.

### GetAffinityTypeOk

`func (o *ClusterAffinityGroupCreate) GetAffinityTypeOk() (*string, bool)`

GetAffinityTypeOk returns a tuple with the AffinityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinityType

`func (o *ClusterAffinityGroupCreate) SetAffinityType(v string)`

SetAffinityType sets AffinityType field to given value.

### HasAffinityType

`func (o *ClusterAffinityGroupCreate) HasAffinityType() bool`

HasAffinityType returns a boolean if a field has been set.

### GetActive

`func (o *ClusterAffinityGroupCreate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ClusterAffinityGroupCreate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ClusterAffinityGroupCreate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ClusterAffinityGroupCreate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPool

`func (o *ClusterAffinityGroupCreate) GetPool() ClusterAffinityGroupCreatePool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *ClusterAffinityGroupCreate) GetPoolOk() (*ClusterAffinityGroupCreatePool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *ClusterAffinityGroupCreate) SetPool(v ClusterAffinityGroupCreatePool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *ClusterAffinityGroupCreate) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetServers

`func (o *ClusterAffinityGroupCreate) GetServers() []int32`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *ClusterAffinityGroupCreate) GetServersOk() (*[]int32, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *ClusterAffinityGroupCreate) SetServers(v []int32)`

SetServers sets Servers field to given value.

### HasServers

`func (o *ClusterAffinityGroupCreate) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetVisibility

`func (o *ClusterAffinityGroupCreate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ClusterAffinityGroupCreate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ClusterAffinityGroupCreate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ClusterAffinityGroupCreate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTenants

`func (o *ClusterAffinityGroupCreate) GetTenants() []ClusterAffinityGroupCreateTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ClusterAffinityGroupCreate) GetTenantsOk() (*[]ClusterAffinityGroupCreateTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ClusterAffinityGroupCreate) SetTenants(v []ClusterAffinityGroupCreateTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ClusterAffinityGroupCreate) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *ClusterAffinityGroupCreate) GetResourcePermissions() ClusterAffinityGroupCreateResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *ClusterAffinityGroupCreate) GetResourcePermissionsOk() (*ClusterAffinityGroupCreateResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *ClusterAffinityGroupCreate) SetResourcePermissions(v ClusterAffinityGroupCreateResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *ClusterAffinityGroupCreate) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


