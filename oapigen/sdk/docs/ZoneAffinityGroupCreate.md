# ZoneAffinityGroupCreate

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

### NewZoneAffinityGroupCreate

`func NewZoneAffinityGroupCreate() *ZoneAffinityGroupCreate`

NewZoneAffinityGroupCreate instantiates a new ZoneAffinityGroupCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetName

`func (o *ZoneAffinityGroupCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZoneAffinityGroupCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZoneAffinityGroupCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ZoneAffinityGroupCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAffinityType

`func (o *ZoneAffinityGroupCreate) GetAffinityType() string`

GetAffinityType returns the AffinityType field if non-nil, zero value otherwise.

### GetAffinityTypeOk

`func (o *ZoneAffinityGroupCreate) GetAffinityTypeOk() (*string, bool)`

GetAffinityTypeOk returns a tuple with the AffinityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinityType

`func (o *ZoneAffinityGroupCreate) SetAffinityType(v string)`

SetAffinityType sets AffinityType field to given value.

### HasAffinityType

`func (o *ZoneAffinityGroupCreate) HasAffinityType() bool`

HasAffinityType returns a boolean if a field has been set.

### GetActive

`func (o *ZoneAffinityGroupCreate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ZoneAffinityGroupCreate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ZoneAffinityGroupCreate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ZoneAffinityGroupCreate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPool

`func (o *ZoneAffinityGroupCreate) GetPool() ZoneAffinityGroupCreatePool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *ZoneAffinityGroupCreate) GetPoolOk() (*ZoneAffinityGroupCreatePool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *ZoneAffinityGroupCreate) SetPool(v ZoneAffinityGroupCreatePool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *ZoneAffinityGroupCreate) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetServers

`func (o *ZoneAffinityGroupCreate) GetServers() []int32`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *ZoneAffinityGroupCreate) GetServersOk() (*[]int32, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *ZoneAffinityGroupCreate) SetServers(v []int32)`

SetServers sets Servers field to given value.

### HasServers

`func (o *ZoneAffinityGroupCreate) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetVisibility

`func (o *ZoneAffinityGroupCreate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ZoneAffinityGroupCreate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ZoneAffinityGroupCreate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ZoneAffinityGroupCreate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTenants

`func (o *ZoneAffinityGroupCreate) GetTenants() []ZoneAffinityGroupCreateTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ZoneAffinityGroupCreate) GetTenantsOk() (*[]ZoneAffinityGroupCreateTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ZoneAffinityGroupCreate) SetTenants(v []ZoneAffinityGroupCreateTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ZoneAffinityGroupCreate) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *ZoneAffinityGroupCreate) GetResourcePermissions() ZoneAffinityGroupCreateResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *ZoneAffinityGroupCreate) GetResourcePermissionsOk() (*ZoneAffinityGroupCreateResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *ZoneAffinityGroupCreate) SetResourcePermissions(v ZoneAffinityGroupCreateResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *ZoneAffinityGroupCreate) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


