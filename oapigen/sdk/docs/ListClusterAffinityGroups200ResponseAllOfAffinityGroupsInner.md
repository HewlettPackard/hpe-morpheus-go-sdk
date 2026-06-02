# ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner

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
**Pool** | Pointer to [**ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInnerPool**](ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInnerPool.md) |  | [optional] 
**Servers** | Pointer to [**[]ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInnerServersInner**](ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInnerServersInner.md) | List of Servers in the Affinity Group | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Tenants** | Pointer to [**[]ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInnerTenantsInner**](ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInnerTenantsInner.md) | Array of tenant account ids that are allowed access | [optional] 
**ResourcePermissions** | Pointer to [**ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInnerResourcePermissions**](ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInnerResourcePermissions.md) |  | [optional] 

## Methods

### NewListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner

`func NewListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner() *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner`

NewListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner instantiates a new ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAffinityType

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) GetAffinityType() string`

GetAffinityType returns the AffinityType field if non-nil, zero value otherwise.

### GetAffinityTypeOk

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) GetAffinityTypeOk() (*string, bool)`

GetAffinityTypeOk returns a tuple with the AffinityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinityType

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) SetAffinityType(v string)`

SetAffinityType sets AffinityType field to given value.

### HasAffinityType

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) HasAffinityType() bool`

HasAffinityType returns a boolean if a field has been set.

### GetSource

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetRefType

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetActive

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPool

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) GetPool() ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInnerPool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) GetPoolOk() (*ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInnerPool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) SetPool(v ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInnerPool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetServers

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) GetServers() []ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInnerServersInner`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) GetServersOk() (*[]ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInnerServersInner, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) SetServers(v []ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInnerServersInner)`

SetServers sets Servers field to given value.

### HasServers

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetVisibility

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTenants

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) GetTenants() []ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInnerTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) GetTenantsOk() (*[]ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInnerTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) SetTenants(v []ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInnerTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) GetResourcePermissions() ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInnerResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) GetResourcePermissionsOk() (*ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInnerResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) SetResourcePermissions(v ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInnerResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


