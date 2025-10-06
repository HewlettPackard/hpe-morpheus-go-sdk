# SaveClusterDatastoreRequestDatastoreResourcePermissions

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
**Account** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**Sites** | Pointer to [**[]GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**Plans** | Pointer to [**[]ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 

## Methods

### NewSaveClusterDatastoreRequestDatastoreResourcePermissions

`func NewSaveClusterDatastoreRequestDatastoreResourcePermissions() *SaveClusterDatastoreRequestDatastoreResourcePermissions`

NewSaveClusterDatastoreRequestDatastoreResourcePermissions instantiates a new SaveClusterDatastoreRequestDatastoreResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveClusterDatastoreRequestDatastoreResourcePermissionsWithDefaults

`func NewSaveClusterDatastoreRequestDatastoreResourcePermissionsWithDefaults() *SaveClusterDatastoreRequestDatastoreResourcePermissions`

NewSaveClusterDatastoreRequestDatastoreResourcePermissionsWithDefaults instantiates a new SaveClusterDatastoreRequestDatastoreResourcePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllGroups

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) GetAllGroups() bool`

GetAllGroups returns the AllGroups field if non-nil, zero value otherwise.

### GetAllGroupsOk

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) GetAllGroupsOk() (*bool, bool)`

GetAllGroupsOk returns a tuple with the AllGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllGroups

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) SetAllGroups(v bool)`

SetAllGroups sets AllGroups field to given value.

### HasAllGroups

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) HasAllGroups() bool`

HasAllGroups returns a boolean if a field has been set.

### GetDefaultStore

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetAllPlans

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetDefaultTarget

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) GetDefaultTarget() bool`

GetDefaultTarget returns the DefaultTarget field if non-nil, zero value otherwise.

### GetDefaultTargetOk

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) GetDefaultTargetOk() (*bool, bool)`

GetDefaultTargetOk returns a tuple with the DefaultTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTarget

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) SetDefaultTarget(v bool)`

SetDefaultTarget sets DefaultTarget field to given value.

### HasDefaultTarget

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) HasDefaultTarget() bool`

HasDefaultTarget returns a boolean if a field has been set.

### GetMorpheusResourceType

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) GetMorpheusResourceType() string`

GetMorpheusResourceType returns the MorpheusResourceType field if non-nil, zero value otherwise.

### GetMorpheusResourceTypeOk

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) GetMorpheusResourceTypeOk() (*string, bool)`

GetMorpheusResourceTypeOk returns a tuple with the MorpheusResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMorpheusResourceType

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) SetMorpheusResourceType(v string)`

SetMorpheusResourceType sets MorpheusResourceType field to given value.

### HasMorpheusResourceType

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) HasMorpheusResourceType() bool`

HasMorpheusResourceType returns a boolean if a field has been set.

### GetMorpheusResourceId

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) GetMorpheusResourceId() int64`

GetMorpheusResourceId returns the MorpheusResourceId field if non-nil, zero value otherwise.

### GetMorpheusResourceIdOk

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) GetMorpheusResourceIdOk() (*int64, bool)`

GetMorpheusResourceIdOk returns a tuple with the MorpheusResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMorpheusResourceId

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) SetMorpheusResourceId(v int64)`

SetMorpheusResourceId sets MorpheusResourceId field to given value.

### HasMorpheusResourceId

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) HasMorpheusResourceId() bool`

HasMorpheusResourceId returns a boolean if a field has been set.

### GetCanManage

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) GetCanManage() bool`

GetCanManage returns the CanManage field if non-nil, zero value otherwise.

### GetCanManageOk

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) GetCanManageOk() (*bool, bool)`

GetCanManageOk returns a tuple with the CanManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManage

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) SetCanManage(v bool)`

SetCanManage sets CanManage field to given value.

### HasCanManage

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) HasCanManage() bool`

HasCanManage returns a boolean if a field has been set.

### GetAll

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetAccount

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) GetAccount() GetAlerts200ResponseAllOfChecksInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) GetAccountOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) SetAccount(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSites

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) GetSites() []GetAlerts200ResponseAllOfChecksInnerAccount`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) GetSitesOk() (*[]GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) SetSites(v []GetAlerts200ResponseAllOfChecksInnerAccount)`

SetSites sets Sites field to given value.

### HasSites

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetPlans

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) GetPlans() []ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) GetPlansOk() (*[]ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) SetPlans(v []ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *SaveClusterDatastoreRequestDatastoreResourcePermissions) HasPlans() bool`

HasPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


