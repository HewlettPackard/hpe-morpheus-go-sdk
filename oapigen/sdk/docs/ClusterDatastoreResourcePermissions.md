# ClusterDatastoreResourcePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllGroups** | Pointer to **bool** |  | [optional] 
**DefaultStore** | Pointer to **bool** |  | [optional] 
**AllPlans** | Pointer to **bool** |  | [optional] 
**DefaultTarget** | Pointer to **bool** |  | [optional] 
**MorpheusResourceType** | Pointer to **string** |  | [optional] 
**MorpheusResourceId** | Pointer to **int64** |  | [optional] 
**CanManage** | Pointer to **bool** |  | [optional] 
**All** | Pointer to **bool** |  | [optional] 
**Account** | Pointer to [**ClusterDatastoreResourcePermissionsAccount**](ClusterDatastoreResourcePermissionsAccount.md) |  | [optional] 
**Sites** | Pointer to [**[]ClusterDatastoreResourcePermissionsSitesInner**](ClusterDatastoreResourcePermissionsSitesInner.md) |  | [optional] 
**Plans** | Pointer to [**[]ClusterDatastoreResourcePermissionsPlansInner**](ClusterDatastoreResourcePermissionsPlansInner.md) |  | [optional] 

## Methods

### NewClusterDatastoreResourcePermissions

`func NewClusterDatastoreResourcePermissions() *ClusterDatastoreResourcePermissions`

NewClusterDatastoreResourcePermissions instantiates a new ClusterDatastoreResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDatastoreResourcePermissionsWithDefaults

`func NewClusterDatastoreResourcePermissionsWithDefaults() *ClusterDatastoreResourcePermissions`

NewClusterDatastoreResourcePermissionsWithDefaults instantiates a new ClusterDatastoreResourcePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllGroups

`func (o *ClusterDatastoreResourcePermissions) GetAllGroups() bool`

GetAllGroups returns the AllGroups field if non-nil, zero value otherwise.

### GetAllGroupsOk

`func (o *ClusterDatastoreResourcePermissions) GetAllGroupsOk() (*bool, bool)`

GetAllGroupsOk returns a tuple with the AllGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllGroups

`func (o *ClusterDatastoreResourcePermissions) SetAllGroups(v bool)`

SetAllGroups sets AllGroups field to given value.

### HasAllGroups

`func (o *ClusterDatastoreResourcePermissions) HasAllGroups() bool`

HasAllGroups returns a boolean if a field has been set.

### GetDefaultStore

`func (o *ClusterDatastoreResourcePermissions) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *ClusterDatastoreResourcePermissions) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *ClusterDatastoreResourcePermissions) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *ClusterDatastoreResourcePermissions) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetAllPlans

`func (o *ClusterDatastoreResourcePermissions) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *ClusterDatastoreResourcePermissions) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *ClusterDatastoreResourcePermissions) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *ClusterDatastoreResourcePermissions) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetDefaultTarget

`func (o *ClusterDatastoreResourcePermissions) GetDefaultTarget() bool`

GetDefaultTarget returns the DefaultTarget field if non-nil, zero value otherwise.

### GetDefaultTargetOk

`func (o *ClusterDatastoreResourcePermissions) GetDefaultTargetOk() (*bool, bool)`

GetDefaultTargetOk returns a tuple with the DefaultTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTarget

`func (o *ClusterDatastoreResourcePermissions) SetDefaultTarget(v bool)`

SetDefaultTarget sets DefaultTarget field to given value.

### HasDefaultTarget

`func (o *ClusterDatastoreResourcePermissions) HasDefaultTarget() bool`

HasDefaultTarget returns a boolean if a field has been set.

### GetMorpheusResourceType

`func (o *ClusterDatastoreResourcePermissions) GetMorpheusResourceType() string`

GetMorpheusResourceType returns the MorpheusResourceType field if non-nil, zero value otherwise.

### GetMorpheusResourceTypeOk

`func (o *ClusterDatastoreResourcePermissions) GetMorpheusResourceTypeOk() (*string, bool)`

GetMorpheusResourceTypeOk returns a tuple with the MorpheusResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMorpheusResourceType

`func (o *ClusterDatastoreResourcePermissions) SetMorpheusResourceType(v string)`

SetMorpheusResourceType sets MorpheusResourceType field to given value.

### HasMorpheusResourceType

`func (o *ClusterDatastoreResourcePermissions) HasMorpheusResourceType() bool`

HasMorpheusResourceType returns a boolean if a field has been set.

### GetMorpheusResourceId

`func (o *ClusterDatastoreResourcePermissions) GetMorpheusResourceId() int64`

GetMorpheusResourceId returns the MorpheusResourceId field if non-nil, zero value otherwise.

### GetMorpheusResourceIdOk

`func (o *ClusterDatastoreResourcePermissions) GetMorpheusResourceIdOk() (*int64, bool)`

GetMorpheusResourceIdOk returns a tuple with the MorpheusResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMorpheusResourceId

`func (o *ClusterDatastoreResourcePermissions) SetMorpheusResourceId(v int64)`

SetMorpheusResourceId sets MorpheusResourceId field to given value.

### HasMorpheusResourceId

`func (o *ClusterDatastoreResourcePermissions) HasMorpheusResourceId() bool`

HasMorpheusResourceId returns a boolean if a field has been set.

### GetCanManage

`func (o *ClusterDatastoreResourcePermissions) GetCanManage() bool`

GetCanManage returns the CanManage field if non-nil, zero value otherwise.

### GetCanManageOk

`func (o *ClusterDatastoreResourcePermissions) GetCanManageOk() (*bool, bool)`

GetCanManageOk returns a tuple with the CanManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManage

`func (o *ClusterDatastoreResourcePermissions) SetCanManage(v bool)`

SetCanManage sets CanManage field to given value.

### HasCanManage

`func (o *ClusterDatastoreResourcePermissions) HasCanManage() bool`

HasCanManage returns a boolean if a field has been set.

### GetAll

`func (o *ClusterDatastoreResourcePermissions) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *ClusterDatastoreResourcePermissions) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *ClusterDatastoreResourcePermissions) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *ClusterDatastoreResourcePermissions) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetAccount

`func (o *ClusterDatastoreResourcePermissions) GetAccount() ClusterDatastoreResourcePermissionsAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ClusterDatastoreResourcePermissions) GetAccountOk() (*ClusterDatastoreResourcePermissionsAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ClusterDatastoreResourcePermissions) SetAccount(v ClusterDatastoreResourcePermissionsAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ClusterDatastoreResourcePermissions) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSites

`func (o *ClusterDatastoreResourcePermissions) GetSites() []ClusterDatastoreResourcePermissionsSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ClusterDatastoreResourcePermissions) GetSitesOk() (*[]ClusterDatastoreResourcePermissionsSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ClusterDatastoreResourcePermissions) SetSites(v []ClusterDatastoreResourcePermissionsSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *ClusterDatastoreResourcePermissions) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetPlans

`func (o *ClusterDatastoreResourcePermissions) GetPlans() []ClusterDatastoreResourcePermissionsPlansInner`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *ClusterDatastoreResourcePermissions) GetPlansOk() (*[]ClusterDatastoreResourcePermissionsPlansInner, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *ClusterDatastoreResourcePermissions) SetPlans(v []ClusterDatastoreResourcePermissionsPlansInner)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *ClusterDatastoreResourcePermissions) HasPlans() bool`

HasPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


