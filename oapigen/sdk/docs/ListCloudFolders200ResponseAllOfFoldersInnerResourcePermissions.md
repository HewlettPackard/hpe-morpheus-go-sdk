# ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultStore** | Pointer to **bool** |  | [optional] 
**DefaultTarget** | Pointer to **bool** |  | [optional] 
**CanManage** | Pointer to **bool** |  | [optional] 
**All** | Pointer to **bool** |  | [optional] 
**Account** | Pointer to [**ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsAccount**](ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsAccount.md) |  | [optional] 
**Sites** | Pointer to [**[]ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsSitesInner**](ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsSitesInner.md) |  | [optional] 
**AllPlans** | Pointer to **bool** |  | [optional] 
**Plans** | Pointer to [**[]ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsPlansInner**](ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsPlansInner.md) |  | [optional] 

## Methods

### NewListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions

`func NewListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions() *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions`

NewListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions instantiates a new ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetDefaultStore

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetDefaultTarget

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) GetDefaultTarget() bool`

GetDefaultTarget returns the DefaultTarget field if non-nil, zero value otherwise.

### GetDefaultTargetOk

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) GetDefaultTargetOk() (*bool, bool)`

GetDefaultTargetOk returns a tuple with the DefaultTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTarget

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) SetDefaultTarget(v bool)`

SetDefaultTarget sets DefaultTarget field to given value.

### HasDefaultTarget

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) HasDefaultTarget() bool`

HasDefaultTarget returns a boolean if a field has been set.

### GetCanManage

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) GetCanManage() bool`

GetCanManage returns the CanManage field if non-nil, zero value otherwise.

### GetCanManageOk

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) GetCanManageOk() (*bool, bool)`

GetCanManageOk returns a tuple with the CanManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManage

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) SetCanManage(v bool)`

SetCanManage sets CanManage field to given value.

### HasCanManage

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) HasCanManage() bool`

HasCanManage returns a boolean if a field has been set.

### GetAll

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetAccount

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) GetAccount() ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) GetAccountOk() (*ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) SetAccount(v ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSites

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) GetSites() []ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) GetSitesOk() (*[]ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) SetSites(v []ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) HasSites() bool`

HasSites returns a boolean if a field has been set.

### SetSitesNil

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) SetSitesNil(b bool)`

 SetSitesNil sets the value for Sites to be an explicit nil

### UnsetSites
`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) UnsetSites()`

UnsetSites ensures that no value is present for Sites, not even an explicit nil
### GetAllPlans

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetPlans

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) GetPlans() []ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsPlansInner`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) GetPlansOk() (*[]ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsPlansInner, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) SetPlans(v []ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsPlansInner)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### SetPlansNil

`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) SetPlansNil(b bool)`

 SetPlansNil sets the value for Plans to be an explicit nil

### UnsetPlans
`func (o *ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissions) UnsetPlans()`

UnsetPlans ensures that no value is present for Plans, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


