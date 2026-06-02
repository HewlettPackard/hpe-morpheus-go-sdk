# GetClusterAffinityGroup200ResponseAffinityGroup

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
**Pool** | Pointer to [**GetClusterAffinityGroup200ResponseAffinityGroupPool**](GetClusterAffinityGroup200ResponseAffinityGroupPool.md) |  | [optional] 
**Servers** | Pointer to [**[]GetClusterAffinityGroup200ResponseAffinityGroupServersInner**](GetClusterAffinityGroup200ResponseAffinityGroupServersInner.md) | List of Servers in the Affinity Group | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Tenants** | Pointer to [**[]GetClusterAffinityGroup200ResponseAffinityGroupTenantsInner**](GetClusterAffinityGroup200ResponseAffinityGroupTenantsInner.md) | Array of tenant account ids that are allowed access | [optional] 
**ResourcePermissions** | Pointer to [**GetClusterAffinityGroup200ResponseAffinityGroupResourcePermissions**](GetClusterAffinityGroup200ResponseAffinityGroupResourcePermissions.md) |  | [optional] 

## Methods

### NewGetClusterAffinityGroup200ResponseAffinityGroup

`func NewGetClusterAffinityGroup200ResponseAffinityGroup() *GetClusterAffinityGroup200ResponseAffinityGroup`

NewGetClusterAffinityGroup200ResponseAffinityGroup instantiates a new GetClusterAffinityGroup200ResponseAffinityGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAffinityType

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) GetAffinityType() string`

GetAffinityType returns the AffinityType field if non-nil, zero value otherwise.

### GetAffinityTypeOk

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) GetAffinityTypeOk() (*string, bool)`

GetAffinityTypeOk returns a tuple with the AffinityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinityType

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) SetAffinityType(v string)`

SetAffinityType sets AffinityType field to given value.

### HasAffinityType

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) HasAffinityType() bool`

HasAffinityType returns a boolean if a field has been set.

### GetSource

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetRefType

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetActive

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPool

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) GetPool() GetClusterAffinityGroup200ResponseAffinityGroupPool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) GetPoolOk() (*GetClusterAffinityGroup200ResponseAffinityGroupPool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) SetPool(v GetClusterAffinityGroup200ResponseAffinityGroupPool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetServers

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) GetServers() []GetClusterAffinityGroup200ResponseAffinityGroupServersInner`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) GetServersOk() (*[]GetClusterAffinityGroup200ResponseAffinityGroupServersInner, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) SetServers(v []GetClusterAffinityGroup200ResponseAffinityGroupServersInner)`

SetServers sets Servers field to given value.

### HasServers

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetVisibility

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTenants

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) GetTenants() []GetClusterAffinityGroup200ResponseAffinityGroupTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) GetTenantsOk() (*[]GetClusterAffinityGroup200ResponseAffinityGroupTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) SetTenants(v []GetClusterAffinityGroup200ResponseAffinityGroupTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) GetResourcePermissions() GetClusterAffinityGroup200ResponseAffinityGroupResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) GetResourcePermissionsOk() (*GetClusterAffinityGroup200ResponseAffinityGroupResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) SetResourcePermissions(v GetClusterAffinityGroup200ResponseAffinityGroupResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *GetClusterAffinityGroup200ResponseAffinityGroup) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


