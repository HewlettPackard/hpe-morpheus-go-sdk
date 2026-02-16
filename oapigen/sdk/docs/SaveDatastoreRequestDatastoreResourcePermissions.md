# SaveDatastoreRequestDatastoreResourcePermissions

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
**Account** | Pointer to [**SaveDatastoreRequestDatastoreResourcePermissionsAccount**](SaveDatastoreRequestDatastoreResourcePermissionsAccount.md) |  | [optional] 
**Sites** | Pointer to [**[]SaveDatastoreRequestDatastoreResourcePermissionsSitesInner**](SaveDatastoreRequestDatastoreResourcePermissionsSitesInner.md) |  | [optional] 
**Plans** | Pointer to [**[]SaveDatastoreRequestDatastoreResourcePermissionsPlansInner**](SaveDatastoreRequestDatastoreResourcePermissionsPlansInner.md) |  | [optional] 

## Methods

### NewSaveDatastoreRequestDatastoreResourcePermissions

`func NewSaveDatastoreRequestDatastoreResourcePermissions() *SaveDatastoreRequestDatastoreResourcePermissions`

NewSaveDatastoreRequestDatastoreResourcePermissions instantiates a new SaveDatastoreRequestDatastoreResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveDatastoreRequestDatastoreResourcePermissionsWithDefaults

`func NewSaveDatastoreRequestDatastoreResourcePermissionsWithDefaults() *SaveDatastoreRequestDatastoreResourcePermissions`

NewSaveDatastoreRequestDatastoreResourcePermissionsWithDefaults instantiates a new SaveDatastoreRequestDatastoreResourcePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllGroups

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) GetAllGroups() bool`

GetAllGroups returns the AllGroups field if non-nil, zero value otherwise.

### GetAllGroupsOk

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) GetAllGroupsOk() (*bool, bool)`

GetAllGroupsOk returns a tuple with the AllGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllGroups

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) SetAllGroups(v bool)`

SetAllGroups sets AllGroups field to given value.

### HasAllGroups

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) HasAllGroups() bool`

HasAllGroups returns a boolean if a field has been set.

### GetDefaultStore

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetAllPlans

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetDefaultTarget

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) GetDefaultTarget() bool`

GetDefaultTarget returns the DefaultTarget field if non-nil, zero value otherwise.

### GetDefaultTargetOk

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) GetDefaultTargetOk() (*bool, bool)`

GetDefaultTargetOk returns a tuple with the DefaultTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTarget

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) SetDefaultTarget(v bool)`

SetDefaultTarget sets DefaultTarget field to given value.

### HasDefaultTarget

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) HasDefaultTarget() bool`

HasDefaultTarget returns a boolean if a field has been set.

### GetMorpheusResourceType

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) GetMorpheusResourceType() string`

GetMorpheusResourceType returns the MorpheusResourceType field if non-nil, zero value otherwise.

### GetMorpheusResourceTypeOk

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) GetMorpheusResourceTypeOk() (*string, bool)`

GetMorpheusResourceTypeOk returns a tuple with the MorpheusResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMorpheusResourceType

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) SetMorpheusResourceType(v string)`

SetMorpheusResourceType sets MorpheusResourceType field to given value.

### HasMorpheusResourceType

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) HasMorpheusResourceType() bool`

HasMorpheusResourceType returns a boolean if a field has been set.

### GetMorpheusResourceId

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) GetMorpheusResourceId() int64`

GetMorpheusResourceId returns the MorpheusResourceId field if non-nil, zero value otherwise.

### GetMorpheusResourceIdOk

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) GetMorpheusResourceIdOk() (*int64, bool)`

GetMorpheusResourceIdOk returns a tuple with the MorpheusResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMorpheusResourceId

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) SetMorpheusResourceId(v int64)`

SetMorpheusResourceId sets MorpheusResourceId field to given value.

### HasMorpheusResourceId

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) HasMorpheusResourceId() bool`

HasMorpheusResourceId returns a boolean if a field has been set.

### GetCanManage

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) GetCanManage() bool`

GetCanManage returns the CanManage field if non-nil, zero value otherwise.

### GetCanManageOk

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) GetCanManageOk() (*bool, bool)`

GetCanManageOk returns a tuple with the CanManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManage

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) SetCanManage(v bool)`

SetCanManage sets CanManage field to given value.

### HasCanManage

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) HasCanManage() bool`

HasCanManage returns a boolean if a field has been set.

### GetAll

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetAccount

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) GetAccount() SaveDatastoreRequestDatastoreResourcePermissionsAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) GetAccountOk() (*SaveDatastoreRequestDatastoreResourcePermissionsAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) SetAccount(v SaveDatastoreRequestDatastoreResourcePermissionsAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSites

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) GetSites() []SaveDatastoreRequestDatastoreResourcePermissionsSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) GetSitesOk() (*[]SaveDatastoreRequestDatastoreResourcePermissionsSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) SetSites(v []SaveDatastoreRequestDatastoreResourcePermissionsSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetPlans

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) GetPlans() []SaveDatastoreRequestDatastoreResourcePermissionsPlansInner`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) GetPlansOk() (*[]SaveDatastoreRequestDatastoreResourcePermissionsPlansInner, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) SetPlans(v []SaveDatastoreRequestDatastoreResourcePermissionsPlansInner)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *SaveDatastoreRequestDatastoreResourcePermissions) HasPlans() bool`

HasPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


