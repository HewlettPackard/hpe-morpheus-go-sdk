# SaveClusterAffinityGroupRequestAffinityGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**AffinityType** | Pointer to **string** | Affinity Type | [optional] 
**Active** | Pointer to **bool** | Active | [optional] 
**Pool** | Pointer to [**SaveClusterAffinityGroupRequestAffinityGroupPool**](SaveClusterAffinityGroupRequestAffinityGroupPool.md) |  | [optional] 
**Servers** | Pointer to **[]int32** | List of Server IDs to include in the Affinity Group | [optional] 
**Visibility** | Pointer to **string** | Visibility - Set to public to allow all tenants | [optional] [default to "private"]
**Tenants** | Pointer to [**[]SaveClusterAffinityGroupRequestAffinityGroupTenantsInner**](SaveClusterAffinityGroupRequestAffinityGroupTenantsInner.md) | Array of tenant account ids that are allowed access | [optional] 
**ResourcePermissions** | Pointer to [**SaveClusterAffinityGroupRequestAffinityGroupResourcePermissions**](SaveClusterAffinityGroupRequestAffinityGroupResourcePermissions.md) |  | [optional] 

## Methods

### NewSaveClusterAffinityGroupRequestAffinityGroup

`func NewSaveClusterAffinityGroupRequestAffinityGroup() *SaveClusterAffinityGroupRequestAffinityGroup`

NewSaveClusterAffinityGroupRequestAffinityGroup instantiates a new SaveClusterAffinityGroupRequestAffinityGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveClusterAffinityGroupRequestAffinityGroupWithDefaults

`func NewSaveClusterAffinityGroupRequestAffinityGroupWithDefaults() *SaveClusterAffinityGroupRequestAffinityGroup`

NewSaveClusterAffinityGroupRequestAffinityGroupWithDefaults instantiates a new SaveClusterAffinityGroupRequestAffinityGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAffinityType

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) GetAffinityType() string`

GetAffinityType returns the AffinityType field if non-nil, zero value otherwise.

### GetAffinityTypeOk

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) GetAffinityTypeOk() (*string, bool)`

GetAffinityTypeOk returns a tuple with the AffinityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinityType

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) SetAffinityType(v string)`

SetAffinityType sets AffinityType field to given value.

### HasAffinityType

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) HasAffinityType() bool`

HasAffinityType returns a boolean if a field has been set.

### GetActive

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPool

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) GetPool() SaveClusterAffinityGroupRequestAffinityGroupPool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) GetPoolOk() (*SaveClusterAffinityGroupRequestAffinityGroupPool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) SetPool(v SaveClusterAffinityGroupRequestAffinityGroupPool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetServers

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) GetServers() []int32`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) GetServersOk() (*[]int32, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) SetServers(v []int32)`

SetServers sets Servers field to given value.

### HasServers

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetVisibility

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTenants

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) GetTenants() []SaveClusterAffinityGroupRequestAffinityGroupTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) GetTenantsOk() (*[]SaveClusterAffinityGroupRequestAffinityGroupTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) SetTenants(v []SaveClusterAffinityGroupRequestAffinityGroupTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) GetResourcePermissions() SaveClusterAffinityGroupRequestAffinityGroupResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) GetResourcePermissionsOk() (*SaveClusterAffinityGroupRequestAffinityGroupResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) SetResourcePermissions(v SaveClusterAffinityGroupRequestAffinityGroupResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *SaveClusterAffinityGroupRequestAffinityGroup) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


