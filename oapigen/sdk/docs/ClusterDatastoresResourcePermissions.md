# ClusterDatastoresResourcePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultStore** | Pointer to **bool** |  | [optional] 
**AllPlans** | Pointer to **bool** |  | [optional] 
**DefaultTarget** | Pointer to **bool** |  | [optional] 
**CanManage** | Pointer to **bool** |  | [optional] 
**All** | Pointer to **bool** |  | [optional] 
**Account** | Pointer to [**ClusterDatastoresResourcePermissionsAccount**](ClusterDatastoresResourcePermissionsAccount.md) |  | [optional] 
**Sites** | Pointer to [**[]ClusterDatastoresResourcePermissionsSitesInner**](ClusterDatastoresResourcePermissionsSitesInner.md) |  | [optional] 
**Plans** | Pointer to [**[]ClusterDatastoresResourcePermissionsPlansInner**](ClusterDatastoresResourcePermissionsPlansInner.md) |  | [optional] 

## Methods

### NewClusterDatastoresResourcePermissions

`func NewClusterDatastoresResourcePermissions() *ClusterDatastoresResourcePermissions`

NewClusterDatastoresResourcePermissions instantiates a new ClusterDatastoresResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDatastoresResourcePermissionsWithDefaults

`func NewClusterDatastoresResourcePermissionsWithDefaults() *ClusterDatastoresResourcePermissions`

NewClusterDatastoresResourcePermissionsWithDefaults instantiates a new ClusterDatastoresResourcePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultStore

`func (o *ClusterDatastoresResourcePermissions) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *ClusterDatastoresResourcePermissions) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *ClusterDatastoresResourcePermissions) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *ClusterDatastoresResourcePermissions) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetAllPlans

`func (o *ClusterDatastoresResourcePermissions) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *ClusterDatastoresResourcePermissions) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *ClusterDatastoresResourcePermissions) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *ClusterDatastoresResourcePermissions) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetDefaultTarget

`func (o *ClusterDatastoresResourcePermissions) GetDefaultTarget() bool`

GetDefaultTarget returns the DefaultTarget field if non-nil, zero value otherwise.

### GetDefaultTargetOk

`func (o *ClusterDatastoresResourcePermissions) GetDefaultTargetOk() (*bool, bool)`

GetDefaultTargetOk returns a tuple with the DefaultTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTarget

`func (o *ClusterDatastoresResourcePermissions) SetDefaultTarget(v bool)`

SetDefaultTarget sets DefaultTarget field to given value.

### HasDefaultTarget

`func (o *ClusterDatastoresResourcePermissions) HasDefaultTarget() bool`

HasDefaultTarget returns a boolean if a field has been set.

### GetCanManage

`func (o *ClusterDatastoresResourcePermissions) GetCanManage() bool`

GetCanManage returns the CanManage field if non-nil, zero value otherwise.

### GetCanManageOk

`func (o *ClusterDatastoresResourcePermissions) GetCanManageOk() (*bool, bool)`

GetCanManageOk returns a tuple with the CanManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManage

`func (o *ClusterDatastoresResourcePermissions) SetCanManage(v bool)`

SetCanManage sets CanManage field to given value.

### HasCanManage

`func (o *ClusterDatastoresResourcePermissions) HasCanManage() bool`

HasCanManage returns a boolean if a field has been set.

### GetAll

`func (o *ClusterDatastoresResourcePermissions) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *ClusterDatastoresResourcePermissions) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *ClusterDatastoresResourcePermissions) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *ClusterDatastoresResourcePermissions) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetAccount

`func (o *ClusterDatastoresResourcePermissions) GetAccount() ClusterDatastoresResourcePermissionsAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ClusterDatastoresResourcePermissions) GetAccountOk() (*ClusterDatastoresResourcePermissionsAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ClusterDatastoresResourcePermissions) SetAccount(v ClusterDatastoresResourcePermissionsAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ClusterDatastoresResourcePermissions) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSites

`func (o *ClusterDatastoresResourcePermissions) GetSites() []ClusterDatastoresResourcePermissionsSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ClusterDatastoresResourcePermissions) GetSitesOk() (*[]ClusterDatastoresResourcePermissionsSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ClusterDatastoresResourcePermissions) SetSites(v []ClusterDatastoresResourcePermissionsSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *ClusterDatastoresResourcePermissions) HasSites() bool`

HasSites returns a boolean if a field has been set.

### SetSitesNil

`func (o *ClusterDatastoresResourcePermissions) SetSitesNil(b bool)`

 SetSitesNil sets the value for Sites to be an explicit nil

### UnsetSites
`func (o *ClusterDatastoresResourcePermissions) UnsetSites()`

UnsetSites ensures that no value is present for Sites, not even an explicit nil
### GetPlans

`func (o *ClusterDatastoresResourcePermissions) GetPlans() []ClusterDatastoresResourcePermissionsPlansInner`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *ClusterDatastoresResourcePermissions) GetPlansOk() (*[]ClusterDatastoresResourcePermissionsPlansInner, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *ClusterDatastoresResourcePermissions) SetPlans(v []ClusterDatastoresResourcePermissionsPlansInner)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *ClusterDatastoresResourcePermissions) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### SetPlansNil

`func (o *ClusterDatastoresResourcePermissions) SetPlansNil(b bool)`

 SetPlansNil sets the value for Plans to be an explicit nil

### UnsetPlans
`func (o *ClusterDatastoresResourcePermissions) UnsetPlans()`

UnsetPlans ensures that no value is present for Plans, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


