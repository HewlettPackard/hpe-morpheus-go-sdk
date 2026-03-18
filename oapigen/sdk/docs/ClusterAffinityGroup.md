# ClusterAffinityGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**AffinityType** | Pointer to **string** | Affinity Type | [optional] 
**Source** | Pointer to **string** | Source | [optional] 
**RefType** | Pointer to **string** | Reference Type for the Affinity Group. Can be ComputeZone or ComputeServerGroup for cloud or cluster respectively | [optional] 
**RefId** | Pointer to **int64** | Reference ID for the Affinity Group. The ID of the Cloud or Clusterfor cloud or cluster respectively | [optional] 
**Active** | Pointer to **bool** | Active | [optional] 
**Pool** | Pointer to [**SaveClusterAffinityGroup200ResponseAllOfAffinityGroupPool**](SaveClusterAffinityGroup200ResponseAllOfAffinityGroupPool.md) |  | [optional] 
**Servers** | Pointer to [**[]SaveClusterAffinityGroup200ResponseAllOfAffinityGroupServersInner**](SaveClusterAffinityGroup200ResponseAllOfAffinityGroupServersInner.md) | List of Servers in the Affinity Group | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Tenants** | Pointer to [**[]SaveClusterAffinityGroup200ResponseAllOfAffinityGroupTenantsInner**](SaveClusterAffinityGroup200ResponseAllOfAffinityGroupTenantsInner.md) | Array of tenant account ids that are allowed access | [optional] 
**ResourcePermissions** | Pointer to [**SaveClusterAffinityGroup200ResponseAllOfAffinityGroupResourcePermissions**](SaveClusterAffinityGroup200ResponseAllOfAffinityGroupResourcePermissions.md) |  | [optional] 

## Methods

### NewClusterAffinityGroup

`func NewClusterAffinityGroup() *ClusterAffinityGroup`

NewClusterAffinityGroup instantiates a new ClusterAffinityGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAffinityGroupWithDefaults

`func NewClusterAffinityGroupWithDefaults() *ClusterAffinityGroup`

NewClusterAffinityGroupWithDefaults instantiates a new ClusterAffinityGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterAffinityGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterAffinityGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterAffinityGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterAffinityGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ClusterAffinityGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterAffinityGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterAffinityGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterAffinityGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAffinityType

`func (o *ClusterAffinityGroup) GetAffinityType() string`

GetAffinityType returns the AffinityType field if non-nil, zero value otherwise.

### GetAffinityTypeOk

`func (o *ClusterAffinityGroup) GetAffinityTypeOk() (*string, bool)`

GetAffinityTypeOk returns a tuple with the AffinityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinityType

`func (o *ClusterAffinityGroup) SetAffinityType(v string)`

SetAffinityType sets AffinityType field to given value.

### HasAffinityType

`func (o *ClusterAffinityGroup) HasAffinityType() bool`

HasAffinityType returns a boolean if a field has been set.

### GetSource

`func (o *ClusterAffinityGroup) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ClusterAffinityGroup) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ClusterAffinityGroup) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ClusterAffinityGroup) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetRefType

`func (o *ClusterAffinityGroup) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ClusterAffinityGroup) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ClusterAffinityGroup) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *ClusterAffinityGroup) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *ClusterAffinityGroup) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ClusterAffinityGroup) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ClusterAffinityGroup) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *ClusterAffinityGroup) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetActive

`func (o *ClusterAffinityGroup) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ClusterAffinityGroup) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ClusterAffinityGroup) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ClusterAffinityGroup) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPool

`func (o *ClusterAffinityGroup) GetPool() SaveClusterAffinityGroup200ResponseAllOfAffinityGroupPool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *ClusterAffinityGroup) GetPoolOk() (*SaveClusterAffinityGroup200ResponseAllOfAffinityGroupPool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *ClusterAffinityGroup) SetPool(v SaveClusterAffinityGroup200ResponseAllOfAffinityGroupPool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *ClusterAffinityGroup) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetServers

`func (o *ClusterAffinityGroup) GetServers() []SaveClusterAffinityGroup200ResponseAllOfAffinityGroupServersInner`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *ClusterAffinityGroup) GetServersOk() (*[]SaveClusterAffinityGroup200ResponseAllOfAffinityGroupServersInner, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *ClusterAffinityGroup) SetServers(v []SaveClusterAffinityGroup200ResponseAllOfAffinityGroupServersInner)`

SetServers sets Servers field to given value.

### HasServers

`func (o *ClusterAffinityGroup) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetVisibility

`func (o *ClusterAffinityGroup) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ClusterAffinityGroup) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ClusterAffinityGroup) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ClusterAffinityGroup) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTenants

`func (o *ClusterAffinityGroup) GetTenants() []SaveClusterAffinityGroup200ResponseAllOfAffinityGroupTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ClusterAffinityGroup) GetTenantsOk() (*[]SaveClusterAffinityGroup200ResponseAllOfAffinityGroupTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ClusterAffinityGroup) SetTenants(v []SaveClusterAffinityGroup200ResponseAllOfAffinityGroupTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ClusterAffinityGroup) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *ClusterAffinityGroup) GetResourcePermissions() SaveClusterAffinityGroup200ResponseAllOfAffinityGroupResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *ClusterAffinityGroup) GetResourcePermissionsOk() (*SaveClusterAffinityGroup200ResponseAllOfAffinityGroupResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *ClusterAffinityGroup) SetResourcePermissions(v SaveClusterAffinityGroup200ResponseAllOfAffinityGroupResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *ClusterAffinityGroup) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


