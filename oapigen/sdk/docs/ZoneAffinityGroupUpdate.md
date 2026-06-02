# ZoneAffinityGroupUpdate

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

### NewZoneAffinityGroupUpdate

`func NewZoneAffinityGroupUpdate() *ZoneAffinityGroupUpdate`

NewZoneAffinityGroupUpdate instantiates a new ZoneAffinityGroupUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetName

`func (o *ZoneAffinityGroupUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZoneAffinityGroupUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZoneAffinityGroupUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ZoneAffinityGroupUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAffinityType

`func (o *ZoneAffinityGroupUpdate) GetAffinityType() string`

GetAffinityType returns the AffinityType field if non-nil, zero value otherwise.

### GetAffinityTypeOk

`func (o *ZoneAffinityGroupUpdate) GetAffinityTypeOk() (*string, bool)`

GetAffinityTypeOk returns a tuple with the AffinityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinityType

`func (o *ZoneAffinityGroupUpdate) SetAffinityType(v string)`

SetAffinityType sets AffinityType field to given value.

### HasAffinityType

`func (o *ZoneAffinityGroupUpdate) HasAffinityType() bool`

HasAffinityType returns a boolean if a field has been set.

### GetActive

`func (o *ZoneAffinityGroupUpdate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ZoneAffinityGroupUpdate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ZoneAffinityGroupUpdate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ZoneAffinityGroupUpdate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPool

`func (o *ZoneAffinityGroupUpdate) GetPool() ZoneAffinityGroupCreatePool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *ZoneAffinityGroupUpdate) GetPoolOk() (*ZoneAffinityGroupCreatePool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *ZoneAffinityGroupUpdate) SetPool(v ZoneAffinityGroupCreatePool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *ZoneAffinityGroupUpdate) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetServers

`func (o *ZoneAffinityGroupUpdate) GetServers() []int32`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *ZoneAffinityGroupUpdate) GetServersOk() (*[]int32, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *ZoneAffinityGroupUpdate) SetServers(v []int32)`

SetServers sets Servers field to given value.

### HasServers

`func (o *ZoneAffinityGroupUpdate) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetVisibility

`func (o *ZoneAffinityGroupUpdate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ZoneAffinityGroupUpdate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ZoneAffinityGroupUpdate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ZoneAffinityGroupUpdate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTenants

`func (o *ZoneAffinityGroupUpdate) GetTenants() []ZoneAffinityGroupCreateTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ZoneAffinityGroupUpdate) GetTenantsOk() (*[]ZoneAffinityGroupCreateTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ZoneAffinityGroupUpdate) SetTenants(v []ZoneAffinityGroupCreateTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ZoneAffinityGroupUpdate) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *ZoneAffinityGroupUpdate) GetResourcePermissions() ZoneAffinityGroupCreateResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *ZoneAffinityGroupUpdate) GetResourcePermissionsOk() (*ZoneAffinityGroupCreateResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *ZoneAffinityGroupUpdate) SetResourcePermissions(v ZoneAffinityGroupCreateResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *ZoneAffinityGroupUpdate) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


