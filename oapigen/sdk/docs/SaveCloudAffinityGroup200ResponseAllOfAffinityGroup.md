# SaveCloudAffinityGroup200ResponseAllOfAffinityGroup

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
**Pool** | Pointer to [**SaveCloudAffinityGroup200ResponseAllOfAffinityGroupPool**](SaveCloudAffinityGroup200ResponseAllOfAffinityGroupPool.md) |  | [optional] 
**Servers** | Pointer to [**[]SaveCloudAffinityGroup200ResponseAllOfAffinityGroupServersInner**](SaveCloudAffinityGroup200ResponseAllOfAffinityGroupServersInner.md) | List of Servers in the Affinity Group | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Tenants** | Pointer to [**[]SaveCloudAffinityGroup200ResponseAllOfAffinityGroupTenantsInner**](SaveCloudAffinityGroup200ResponseAllOfAffinityGroupTenantsInner.md) | Array of tenant account ids that are allowed access | [optional] 
**ResourcePermissions** | Pointer to [**SaveCloudAffinityGroup200ResponseAllOfAffinityGroupResourcePermissions**](SaveCloudAffinityGroup200ResponseAllOfAffinityGroupResourcePermissions.md) |  | [optional] 

## Methods

### NewSaveCloudAffinityGroup200ResponseAllOfAffinityGroup

`func NewSaveCloudAffinityGroup200ResponseAllOfAffinityGroup() *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup`

NewSaveCloudAffinityGroup200ResponseAllOfAffinityGroup instantiates a new SaveCloudAffinityGroup200ResponseAllOfAffinityGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAffinityType

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) GetAffinityType() string`

GetAffinityType returns the AffinityType field if non-nil, zero value otherwise.

### GetAffinityTypeOk

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) GetAffinityTypeOk() (*string, bool)`

GetAffinityTypeOk returns a tuple with the AffinityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinityType

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) SetAffinityType(v string)`

SetAffinityType sets AffinityType field to given value.

### HasAffinityType

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) HasAffinityType() bool`

HasAffinityType returns a boolean if a field has been set.

### GetSource

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetRefType

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetActive

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPool

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) GetPool() SaveCloudAffinityGroup200ResponseAllOfAffinityGroupPool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) GetPoolOk() (*SaveCloudAffinityGroup200ResponseAllOfAffinityGroupPool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) SetPool(v SaveCloudAffinityGroup200ResponseAllOfAffinityGroupPool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetServers

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) GetServers() []SaveCloudAffinityGroup200ResponseAllOfAffinityGroupServersInner`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) GetServersOk() (*[]SaveCloudAffinityGroup200ResponseAllOfAffinityGroupServersInner, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) SetServers(v []SaveCloudAffinityGroup200ResponseAllOfAffinityGroupServersInner)`

SetServers sets Servers field to given value.

### HasServers

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetVisibility

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTenants

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) GetTenants() []SaveCloudAffinityGroup200ResponseAllOfAffinityGroupTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) GetTenantsOk() (*[]SaveCloudAffinityGroup200ResponseAllOfAffinityGroupTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) SetTenants(v []SaveCloudAffinityGroup200ResponseAllOfAffinityGroupTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) GetResourcePermissions() SaveCloudAffinityGroup200ResponseAllOfAffinityGroupResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) GetResourcePermissionsOk() (*SaveCloudAffinityGroup200ResponseAllOfAffinityGroupResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) SetResourcePermissions(v SaveCloudAffinityGroup200ResponseAllOfAffinityGroupResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *SaveCloudAffinityGroup200ResponseAllOfAffinityGroup) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


