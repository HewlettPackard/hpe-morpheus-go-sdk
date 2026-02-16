# GetClusterDatastore200ResponseDatastoreResourcePermissions

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
**Account** | Pointer to [**GetClusterDatastore200ResponseDatastoreResourcePermissionsAccount**](GetClusterDatastore200ResponseDatastoreResourcePermissionsAccount.md) |  | [optional] 
**Sites** | Pointer to [**[]GetClusterDatastore200ResponseDatastoreResourcePermissionsSitesInner**](GetClusterDatastore200ResponseDatastoreResourcePermissionsSitesInner.md) |  | [optional] 
**Plans** | Pointer to [**[]GetClusterDatastore200ResponseDatastoreResourcePermissionsPlansInner**](GetClusterDatastore200ResponseDatastoreResourcePermissionsPlansInner.md) |  | [optional] 

## Methods

### NewGetClusterDatastore200ResponseDatastoreResourcePermissions

`func NewGetClusterDatastore200ResponseDatastoreResourcePermissions() *GetClusterDatastore200ResponseDatastoreResourcePermissions`

NewGetClusterDatastore200ResponseDatastoreResourcePermissions instantiates a new GetClusterDatastore200ResponseDatastoreResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClusterDatastore200ResponseDatastoreResourcePermissionsWithDefaults

`func NewGetClusterDatastore200ResponseDatastoreResourcePermissionsWithDefaults() *GetClusterDatastore200ResponseDatastoreResourcePermissions`

NewGetClusterDatastore200ResponseDatastoreResourcePermissionsWithDefaults instantiates a new GetClusterDatastore200ResponseDatastoreResourcePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllGroups

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) GetAllGroups() bool`

GetAllGroups returns the AllGroups field if non-nil, zero value otherwise.

### GetAllGroupsOk

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) GetAllGroupsOk() (*bool, bool)`

GetAllGroupsOk returns a tuple with the AllGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllGroups

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) SetAllGroups(v bool)`

SetAllGroups sets AllGroups field to given value.

### HasAllGroups

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) HasAllGroups() bool`

HasAllGroups returns a boolean if a field has been set.

### GetDefaultStore

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetAllPlans

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetDefaultTarget

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) GetDefaultTarget() bool`

GetDefaultTarget returns the DefaultTarget field if non-nil, zero value otherwise.

### GetDefaultTargetOk

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) GetDefaultTargetOk() (*bool, bool)`

GetDefaultTargetOk returns a tuple with the DefaultTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTarget

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) SetDefaultTarget(v bool)`

SetDefaultTarget sets DefaultTarget field to given value.

### HasDefaultTarget

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) HasDefaultTarget() bool`

HasDefaultTarget returns a boolean if a field has been set.

### GetMorpheusResourceType

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) GetMorpheusResourceType() string`

GetMorpheusResourceType returns the MorpheusResourceType field if non-nil, zero value otherwise.

### GetMorpheusResourceTypeOk

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) GetMorpheusResourceTypeOk() (*string, bool)`

GetMorpheusResourceTypeOk returns a tuple with the MorpheusResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMorpheusResourceType

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) SetMorpheusResourceType(v string)`

SetMorpheusResourceType sets MorpheusResourceType field to given value.

### HasMorpheusResourceType

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) HasMorpheusResourceType() bool`

HasMorpheusResourceType returns a boolean if a field has been set.

### GetMorpheusResourceId

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) GetMorpheusResourceId() int64`

GetMorpheusResourceId returns the MorpheusResourceId field if non-nil, zero value otherwise.

### GetMorpheusResourceIdOk

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) GetMorpheusResourceIdOk() (*int64, bool)`

GetMorpheusResourceIdOk returns a tuple with the MorpheusResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMorpheusResourceId

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) SetMorpheusResourceId(v int64)`

SetMorpheusResourceId sets MorpheusResourceId field to given value.

### HasMorpheusResourceId

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) HasMorpheusResourceId() bool`

HasMorpheusResourceId returns a boolean if a field has been set.

### GetCanManage

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) GetCanManage() bool`

GetCanManage returns the CanManage field if non-nil, zero value otherwise.

### GetCanManageOk

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) GetCanManageOk() (*bool, bool)`

GetCanManageOk returns a tuple with the CanManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManage

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) SetCanManage(v bool)`

SetCanManage sets CanManage field to given value.

### HasCanManage

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) HasCanManage() bool`

HasCanManage returns a boolean if a field has been set.

### GetAll

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetAccount

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) GetAccount() GetClusterDatastore200ResponseDatastoreResourcePermissionsAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) GetAccountOk() (*GetClusterDatastore200ResponseDatastoreResourcePermissionsAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) SetAccount(v GetClusterDatastore200ResponseDatastoreResourcePermissionsAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSites

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) GetSites() []GetClusterDatastore200ResponseDatastoreResourcePermissionsSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) GetSitesOk() (*[]GetClusterDatastore200ResponseDatastoreResourcePermissionsSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) SetSites(v []GetClusterDatastore200ResponseDatastoreResourcePermissionsSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetPlans

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) GetPlans() []GetClusterDatastore200ResponseDatastoreResourcePermissionsPlansInner`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) GetPlansOk() (*[]GetClusterDatastore200ResponseDatastoreResourcePermissionsPlansInner, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) SetPlans(v []GetClusterDatastore200ResponseDatastoreResourcePermissionsPlansInner)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *GetClusterDatastore200ResponseDatastoreResourcePermissions) HasPlans() bool`

HasPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


