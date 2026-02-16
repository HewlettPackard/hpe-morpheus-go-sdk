# ZoneDatastoreCreateResourcePermissions

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
**Account** | Pointer to [**ZoneDatastoreCreateResourcePermissionsAccount**](ZoneDatastoreCreateResourcePermissionsAccount.md) |  | [optional] 
**Sites** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Plans** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewZoneDatastoreCreateResourcePermissions

`func NewZoneDatastoreCreateResourcePermissions() *ZoneDatastoreCreateResourcePermissions`

NewZoneDatastoreCreateResourcePermissions instantiates a new ZoneDatastoreCreateResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneDatastoreCreateResourcePermissionsWithDefaults

`func NewZoneDatastoreCreateResourcePermissionsWithDefaults() *ZoneDatastoreCreateResourcePermissions`

NewZoneDatastoreCreateResourcePermissionsWithDefaults instantiates a new ZoneDatastoreCreateResourcePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllGroups

`func (o *ZoneDatastoreCreateResourcePermissions) GetAllGroups() bool`

GetAllGroups returns the AllGroups field if non-nil, zero value otherwise.

### GetAllGroupsOk

`func (o *ZoneDatastoreCreateResourcePermissions) GetAllGroupsOk() (*bool, bool)`

GetAllGroupsOk returns a tuple with the AllGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllGroups

`func (o *ZoneDatastoreCreateResourcePermissions) SetAllGroups(v bool)`

SetAllGroups sets AllGroups field to given value.

### HasAllGroups

`func (o *ZoneDatastoreCreateResourcePermissions) HasAllGroups() bool`

HasAllGroups returns a boolean if a field has been set.

### GetDefaultStore

`func (o *ZoneDatastoreCreateResourcePermissions) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *ZoneDatastoreCreateResourcePermissions) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *ZoneDatastoreCreateResourcePermissions) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *ZoneDatastoreCreateResourcePermissions) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetAllPlans

`func (o *ZoneDatastoreCreateResourcePermissions) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *ZoneDatastoreCreateResourcePermissions) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *ZoneDatastoreCreateResourcePermissions) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *ZoneDatastoreCreateResourcePermissions) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetDefaultTarget

`func (o *ZoneDatastoreCreateResourcePermissions) GetDefaultTarget() bool`

GetDefaultTarget returns the DefaultTarget field if non-nil, zero value otherwise.

### GetDefaultTargetOk

`func (o *ZoneDatastoreCreateResourcePermissions) GetDefaultTargetOk() (*bool, bool)`

GetDefaultTargetOk returns a tuple with the DefaultTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTarget

`func (o *ZoneDatastoreCreateResourcePermissions) SetDefaultTarget(v bool)`

SetDefaultTarget sets DefaultTarget field to given value.

### HasDefaultTarget

`func (o *ZoneDatastoreCreateResourcePermissions) HasDefaultTarget() bool`

HasDefaultTarget returns a boolean if a field has been set.

### GetMorpheusResourceType

`func (o *ZoneDatastoreCreateResourcePermissions) GetMorpheusResourceType() string`

GetMorpheusResourceType returns the MorpheusResourceType field if non-nil, zero value otherwise.

### GetMorpheusResourceTypeOk

`func (o *ZoneDatastoreCreateResourcePermissions) GetMorpheusResourceTypeOk() (*string, bool)`

GetMorpheusResourceTypeOk returns a tuple with the MorpheusResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMorpheusResourceType

`func (o *ZoneDatastoreCreateResourcePermissions) SetMorpheusResourceType(v string)`

SetMorpheusResourceType sets MorpheusResourceType field to given value.

### HasMorpheusResourceType

`func (o *ZoneDatastoreCreateResourcePermissions) HasMorpheusResourceType() bool`

HasMorpheusResourceType returns a boolean if a field has been set.

### GetMorpheusResourceId

`func (o *ZoneDatastoreCreateResourcePermissions) GetMorpheusResourceId() int64`

GetMorpheusResourceId returns the MorpheusResourceId field if non-nil, zero value otherwise.

### GetMorpheusResourceIdOk

`func (o *ZoneDatastoreCreateResourcePermissions) GetMorpheusResourceIdOk() (*int64, bool)`

GetMorpheusResourceIdOk returns a tuple with the MorpheusResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMorpheusResourceId

`func (o *ZoneDatastoreCreateResourcePermissions) SetMorpheusResourceId(v int64)`

SetMorpheusResourceId sets MorpheusResourceId field to given value.

### HasMorpheusResourceId

`func (o *ZoneDatastoreCreateResourcePermissions) HasMorpheusResourceId() bool`

HasMorpheusResourceId returns a boolean if a field has been set.

### GetCanManage

`func (o *ZoneDatastoreCreateResourcePermissions) GetCanManage() bool`

GetCanManage returns the CanManage field if non-nil, zero value otherwise.

### GetCanManageOk

`func (o *ZoneDatastoreCreateResourcePermissions) GetCanManageOk() (*bool, bool)`

GetCanManageOk returns a tuple with the CanManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManage

`func (o *ZoneDatastoreCreateResourcePermissions) SetCanManage(v bool)`

SetCanManage sets CanManage field to given value.

### HasCanManage

`func (o *ZoneDatastoreCreateResourcePermissions) HasCanManage() bool`

HasCanManage returns a boolean if a field has been set.

### GetAll

`func (o *ZoneDatastoreCreateResourcePermissions) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *ZoneDatastoreCreateResourcePermissions) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *ZoneDatastoreCreateResourcePermissions) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *ZoneDatastoreCreateResourcePermissions) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetAccount

`func (o *ZoneDatastoreCreateResourcePermissions) GetAccount() ZoneDatastoreCreateResourcePermissionsAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ZoneDatastoreCreateResourcePermissions) GetAccountOk() (*ZoneDatastoreCreateResourcePermissionsAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ZoneDatastoreCreateResourcePermissions) SetAccount(v ZoneDatastoreCreateResourcePermissionsAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ZoneDatastoreCreateResourcePermissions) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSites

`func (o *ZoneDatastoreCreateResourcePermissions) GetSites() []map[string]interface{}`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ZoneDatastoreCreateResourcePermissions) GetSitesOk() (*[]map[string]interface{}, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ZoneDatastoreCreateResourcePermissions) SetSites(v []map[string]interface{})`

SetSites sets Sites field to given value.

### HasSites

`func (o *ZoneDatastoreCreateResourcePermissions) HasSites() bool`

HasSites returns a boolean if a field has been set.

### SetSitesNil

`func (o *ZoneDatastoreCreateResourcePermissions) SetSitesNil(b bool)`

 SetSitesNil sets the value for Sites to be an explicit nil

### UnsetSites
`func (o *ZoneDatastoreCreateResourcePermissions) UnsetSites()`

UnsetSites ensures that no value is present for Sites, not even an explicit nil
### GetPlans

`func (o *ZoneDatastoreCreateResourcePermissions) GetPlans() []map[string]interface{}`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *ZoneDatastoreCreateResourcePermissions) GetPlansOk() (*[]map[string]interface{}, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *ZoneDatastoreCreateResourcePermissions) SetPlans(v []map[string]interface{})`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *ZoneDatastoreCreateResourcePermissions) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### SetPlansNil

`func (o *ZoneDatastoreCreateResourcePermissions) SetPlansNil(b bool)`

 SetPlansNil sets the value for Plans to be an explicit nil

### UnsetPlans
`func (o *ZoneDatastoreCreateResourcePermissions) UnsetPlans()`

UnsetPlans ensures that no value is present for Plans, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


