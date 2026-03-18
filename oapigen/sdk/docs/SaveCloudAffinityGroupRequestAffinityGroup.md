# SaveCloudAffinityGroupRequestAffinityGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**AffinityType** | Pointer to **string** | Affinity Type | [optional] 
**Active** | Pointer to **bool** | Active | [optional] 
**Pool** | Pointer to [**SaveCloudAffinityGroupRequestAffinityGroupPool**](SaveCloudAffinityGroupRequestAffinityGroupPool.md) |  | [optional] 
**Servers** | Pointer to **[]int32** | List of Server IDs to include in the Affinity Group | [optional] 
**Visibility** | Pointer to **string** | Visibility - Set to public to allow all tenants | [optional] [default to "private"]
**Tenants** | Pointer to [**[]SaveCloudAffinityGroupRequestAffinityGroupTenantsInner**](SaveCloudAffinityGroupRequestAffinityGroupTenantsInner.md) | Array of tenant account ids that are allowed access | [optional] 
**ResourcePermissions** | Pointer to [**SaveCloudAffinityGroupRequestAffinityGroupResourcePermissions**](SaveCloudAffinityGroupRequestAffinityGroupResourcePermissions.md) |  | [optional] 

## Methods

### NewSaveCloudAffinityGroupRequestAffinityGroup

`func NewSaveCloudAffinityGroupRequestAffinityGroup() *SaveCloudAffinityGroupRequestAffinityGroup`

NewSaveCloudAffinityGroupRequestAffinityGroup instantiates a new SaveCloudAffinityGroupRequestAffinityGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveCloudAffinityGroupRequestAffinityGroupWithDefaults

`func NewSaveCloudAffinityGroupRequestAffinityGroupWithDefaults() *SaveCloudAffinityGroupRequestAffinityGroup`

NewSaveCloudAffinityGroupRequestAffinityGroupWithDefaults instantiates a new SaveCloudAffinityGroupRequestAffinityGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SaveCloudAffinityGroupRequestAffinityGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SaveCloudAffinityGroupRequestAffinityGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SaveCloudAffinityGroupRequestAffinityGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SaveCloudAffinityGroupRequestAffinityGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAffinityType

`func (o *SaveCloudAffinityGroupRequestAffinityGroup) GetAffinityType() string`

GetAffinityType returns the AffinityType field if non-nil, zero value otherwise.

### GetAffinityTypeOk

`func (o *SaveCloudAffinityGroupRequestAffinityGroup) GetAffinityTypeOk() (*string, bool)`

GetAffinityTypeOk returns a tuple with the AffinityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinityType

`func (o *SaveCloudAffinityGroupRequestAffinityGroup) SetAffinityType(v string)`

SetAffinityType sets AffinityType field to given value.

### HasAffinityType

`func (o *SaveCloudAffinityGroupRequestAffinityGroup) HasAffinityType() bool`

HasAffinityType returns a boolean if a field has been set.

### GetActive

`func (o *SaveCloudAffinityGroupRequestAffinityGroup) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SaveCloudAffinityGroupRequestAffinityGroup) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SaveCloudAffinityGroupRequestAffinityGroup) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *SaveCloudAffinityGroupRequestAffinityGroup) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPool

`func (o *SaveCloudAffinityGroupRequestAffinityGroup) GetPool() SaveCloudAffinityGroupRequestAffinityGroupPool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *SaveCloudAffinityGroupRequestAffinityGroup) GetPoolOk() (*SaveCloudAffinityGroupRequestAffinityGroupPool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *SaveCloudAffinityGroupRequestAffinityGroup) SetPool(v SaveCloudAffinityGroupRequestAffinityGroupPool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *SaveCloudAffinityGroupRequestAffinityGroup) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetServers

`func (o *SaveCloudAffinityGroupRequestAffinityGroup) GetServers() []int32`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *SaveCloudAffinityGroupRequestAffinityGroup) GetServersOk() (*[]int32, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *SaveCloudAffinityGroupRequestAffinityGroup) SetServers(v []int32)`

SetServers sets Servers field to given value.

### HasServers

`func (o *SaveCloudAffinityGroupRequestAffinityGroup) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetVisibility

`func (o *SaveCloudAffinityGroupRequestAffinityGroup) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *SaveCloudAffinityGroupRequestAffinityGroup) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *SaveCloudAffinityGroupRequestAffinityGroup) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *SaveCloudAffinityGroupRequestAffinityGroup) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTenants

`func (o *SaveCloudAffinityGroupRequestAffinityGroup) GetTenants() []SaveCloudAffinityGroupRequestAffinityGroupTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *SaveCloudAffinityGroupRequestAffinityGroup) GetTenantsOk() (*[]SaveCloudAffinityGroupRequestAffinityGroupTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *SaveCloudAffinityGroupRequestAffinityGroup) SetTenants(v []SaveCloudAffinityGroupRequestAffinityGroupTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *SaveCloudAffinityGroupRequestAffinityGroup) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *SaveCloudAffinityGroupRequestAffinityGroup) GetResourcePermissions() SaveCloudAffinityGroupRequestAffinityGroupResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *SaveCloudAffinityGroupRequestAffinityGroup) GetResourcePermissionsOk() (*SaveCloudAffinityGroupRequestAffinityGroupResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *SaveCloudAffinityGroupRequestAffinityGroup) SetResourcePermissions(v SaveCloudAffinityGroupRequestAffinityGroupResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *SaveCloudAffinityGroupRequestAffinityGroup) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


