# ClusterDatastoreCreateResourcePermissions

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
**Account** | Pointer to [**ClusterDatastoreCreateResourcePermissionsAccount**](ClusterDatastoreCreateResourcePermissionsAccount.md) |  | [optional] 
**Sites** | Pointer to [**[]ClusterDatastoreCreateResourcePermissionsSitesInner**](ClusterDatastoreCreateResourcePermissionsSitesInner.md) |  | [optional] 
**Plans** | Pointer to [**[]ClusterDatastoreCreateResourcePermissionsPlansInner**](ClusterDatastoreCreateResourcePermissionsPlansInner.md) |  | [optional] 

## Methods

### NewClusterDatastoreCreateResourcePermissions

`func NewClusterDatastoreCreateResourcePermissions() *ClusterDatastoreCreateResourcePermissions`

NewClusterDatastoreCreateResourcePermissions instantiates a new ClusterDatastoreCreateResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDatastoreCreateResourcePermissionsWithDefaults

`func NewClusterDatastoreCreateResourcePermissionsWithDefaults() *ClusterDatastoreCreateResourcePermissions`

NewClusterDatastoreCreateResourcePermissionsWithDefaults instantiates a new ClusterDatastoreCreateResourcePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllGroups

`func (o *ClusterDatastoreCreateResourcePermissions) GetAllGroups() bool`

GetAllGroups returns the AllGroups field if non-nil, zero value otherwise.

### GetAllGroupsOk

`func (o *ClusterDatastoreCreateResourcePermissions) GetAllGroupsOk() (*bool, bool)`

GetAllGroupsOk returns a tuple with the AllGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllGroups

`func (o *ClusterDatastoreCreateResourcePermissions) SetAllGroups(v bool)`

SetAllGroups sets AllGroups field to given value.

### HasAllGroups

`func (o *ClusterDatastoreCreateResourcePermissions) HasAllGroups() bool`

HasAllGroups returns a boolean if a field has been set.

### GetDefaultStore

`func (o *ClusterDatastoreCreateResourcePermissions) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *ClusterDatastoreCreateResourcePermissions) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *ClusterDatastoreCreateResourcePermissions) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *ClusterDatastoreCreateResourcePermissions) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetAllPlans

`func (o *ClusterDatastoreCreateResourcePermissions) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *ClusterDatastoreCreateResourcePermissions) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *ClusterDatastoreCreateResourcePermissions) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *ClusterDatastoreCreateResourcePermissions) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetDefaultTarget

`func (o *ClusterDatastoreCreateResourcePermissions) GetDefaultTarget() bool`

GetDefaultTarget returns the DefaultTarget field if non-nil, zero value otherwise.

### GetDefaultTargetOk

`func (o *ClusterDatastoreCreateResourcePermissions) GetDefaultTargetOk() (*bool, bool)`

GetDefaultTargetOk returns a tuple with the DefaultTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTarget

`func (o *ClusterDatastoreCreateResourcePermissions) SetDefaultTarget(v bool)`

SetDefaultTarget sets DefaultTarget field to given value.

### HasDefaultTarget

`func (o *ClusterDatastoreCreateResourcePermissions) HasDefaultTarget() bool`

HasDefaultTarget returns a boolean if a field has been set.

### GetMorpheusResourceType

`func (o *ClusterDatastoreCreateResourcePermissions) GetMorpheusResourceType() string`

GetMorpheusResourceType returns the MorpheusResourceType field if non-nil, zero value otherwise.

### GetMorpheusResourceTypeOk

`func (o *ClusterDatastoreCreateResourcePermissions) GetMorpheusResourceTypeOk() (*string, bool)`

GetMorpheusResourceTypeOk returns a tuple with the MorpheusResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMorpheusResourceType

`func (o *ClusterDatastoreCreateResourcePermissions) SetMorpheusResourceType(v string)`

SetMorpheusResourceType sets MorpheusResourceType field to given value.

### HasMorpheusResourceType

`func (o *ClusterDatastoreCreateResourcePermissions) HasMorpheusResourceType() bool`

HasMorpheusResourceType returns a boolean if a field has been set.

### GetMorpheusResourceId

`func (o *ClusterDatastoreCreateResourcePermissions) GetMorpheusResourceId() int64`

GetMorpheusResourceId returns the MorpheusResourceId field if non-nil, zero value otherwise.

### GetMorpheusResourceIdOk

`func (o *ClusterDatastoreCreateResourcePermissions) GetMorpheusResourceIdOk() (*int64, bool)`

GetMorpheusResourceIdOk returns a tuple with the MorpheusResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMorpheusResourceId

`func (o *ClusterDatastoreCreateResourcePermissions) SetMorpheusResourceId(v int64)`

SetMorpheusResourceId sets MorpheusResourceId field to given value.

### HasMorpheusResourceId

`func (o *ClusterDatastoreCreateResourcePermissions) HasMorpheusResourceId() bool`

HasMorpheusResourceId returns a boolean if a field has been set.

### GetCanManage

`func (o *ClusterDatastoreCreateResourcePermissions) GetCanManage() bool`

GetCanManage returns the CanManage field if non-nil, zero value otherwise.

### GetCanManageOk

`func (o *ClusterDatastoreCreateResourcePermissions) GetCanManageOk() (*bool, bool)`

GetCanManageOk returns a tuple with the CanManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManage

`func (o *ClusterDatastoreCreateResourcePermissions) SetCanManage(v bool)`

SetCanManage sets CanManage field to given value.

### HasCanManage

`func (o *ClusterDatastoreCreateResourcePermissions) HasCanManage() bool`

HasCanManage returns a boolean if a field has been set.

### GetAll

`func (o *ClusterDatastoreCreateResourcePermissions) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *ClusterDatastoreCreateResourcePermissions) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *ClusterDatastoreCreateResourcePermissions) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *ClusterDatastoreCreateResourcePermissions) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetAccount

`func (o *ClusterDatastoreCreateResourcePermissions) GetAccount() ClusterDatastoreCreateResourcePermissionsAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ClusterDatastoreCreateResourcePermissions) GetAccountOk() (*ClusterDatastoreCreateResourcePermissionsAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ClusterDatastoreCreateResourcePermissions) SetAccount(v ClusterDatastoreCreateResourcePermissionsAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ClusterDatastoreCreateResourcePermissions) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSites

`func (o *ClusterDatastoreCreateResourcePermissions) GetSites() []ClusterDatastoreCreateResourcePermissionsSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ClusterDatastoreCreateResourcePermissions) GetSitesOk() (*[]ClusterDatastoreCreateResourcePermissionsSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ClusterDatastoreCreateResourcePermissions) SetSites(v []ClusterDatastoreCreateResourcePermissionsSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *ClusterDatastoreCreateResourcePermissions) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetPlans

`func (o *ClusterDatastoreCreateResourcePermissions) GetPlans() []ClusterDatastoreCreateResourcePermissionsPlansInner`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *ClusterDatastoreCreateResourcePermissions) GetPlansOk() (*[]ClusterDatastoreCreateResourcePermissionsPlansInner, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *ClusterDatastoreCreateResourcePermissions) SetPlans(v []ClusterDatastoreCreateResourcePermissionsPlansInner)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *ClusterDatastoreCreateResourcePermissions) HasPlans() bool`

HasPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


