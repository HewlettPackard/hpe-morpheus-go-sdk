# UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup

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
**Pool** | Pointer to [**UpdateCloudAffinityGroup200ResponseAllOfAffinityGroupPool**](UpdateCloudAffinityGroup200ResponseAllOfAffinityGroupPool.md) |  | [optional] 
**Servers** | Pointer to [**[]UpdateCloudAffinityGroup200ResponseAllOfAffinityGroupServersInner**](UpdateCloudAffinityGroup200ResponseAllOfAffinityGroupServersInner.md) | List of Servers in the Affinity Group | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Tenants** | Pointer to [**[]UpdateCloudAffinityGroup200ResponseAllOfAffinityGroupTenantsInner**](UpdateCloudAffinityGroup200ResponseAllOfAffinityGroupTenantsInner.md) | Array of tenant account ids that are allowed access | [optional] 
**ResourcePermissions** | Pointer to [**UpdateCloudAffinityGroup200ResponseAllOfAffinityGroupResourcePermissions**](UpdateCloudAffinityGroup200ResponseAllOfAffinityGroupResourcePermissions.md) |  | [optional] 

## Methods

### NewUpdateCloudAffinityGroup200ResponseAllOfAffinityGroup

`func NewUpdateCloudAffinityGroup200ResponseAllOfAffinityGroup() *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup`

NewUpdateCloudAffinityGroup200ResponseAllOfAffinityGroup instantiates a new UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCloudAffinityGroup200ResponseAllOfAffinityGroupWithDefaults

`func NewUpdateCloudAffinityGroup200ResponseAllOfAffinityGroupWithDefaults() *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup`

NewUpdateCloudAffinityGroup200ResponseAllOfAffinityGroupWithDefaults instantiates a new UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAffinityType

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) GetAffinityType() string`

GetAffinityType returns the AffinityType field if non-nil, zero value otherwise.

### GetAffinityTypeOk

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) GetAffinityTypeOk() (*string, bool)`

GetAffinityTypeOk returns a tuple with the AffinityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinityType

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) SetAffinityType(v string)`

SetAffinityType sets AffinityType field to given value.

### HasAffinityType

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) HasAffinityType() bool`

HasAffinityType returns a boolean if a field has been set.

### GetSource

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetRefType

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetActive

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPool

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) GetPool() UpdateCloudAffinityGroup200ResponseAllOfAffinityGroupPool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) GetPoolOk() (*UpdateCloudAffinityGroup200ResponseAllOfAffinityGroupPool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) SetPool(v UpdateCloudAffinityGroup200ResponseAllOfAffinityGroupPool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetServers

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) GetServers() []UpdateCloudAffinityGroup200ResponseAllOfAffinityGroupServersInner`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) GetServersOk() (*[]UpdateCloudAffinityGroup200ResponseAllOfAffinityGroupServersInner, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) SetServers(v []UpdateCloudAffinityGroup200ResponseAllOfAffinityGroupServersInner)`

SetServers sets Servers field to given value.

### HasServers

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTenants

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) GetTenants() []UpdateCloudAffinityGroup200ResponseAllOfAffinityGroupTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) GetTenantsOk() (*[]UpdateCloudAffinityGroup200ResponseAllOfAffinityGroupTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) SetTenants(v []UpdateCloudAffinityGroup200ResponseAllOfAffinityGroupTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) GetResourcePermissions() UpdateCloudAffinityGroup200ResponseAllOfAffinityGroupResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) GetResourcePermissionsOk() (*UpdateCloudAffinityGroup200ResponseAllOfAffinityGroupResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) SetResourcePermissions(v UpdateCloudAffinityGroup200ResponseAllOfAffinityGroupResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *UpdateCloudAffinityGroup200ResponseAllOfAffinityGroup) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


