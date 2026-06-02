# GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission

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

### NewGetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission

`func NewGetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission() *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission`

NewGetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission instantiates a new GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetDefaultStore

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetDefaultTarget

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) GetDefaultTarget() bool`

GetDefaultTarget returns the DefaultTarget field if non-nil, zero value otherwise.

### GetDefaultTargetOk

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) GetDefaultTargetOk() (*bool, bool)`

GetDefaultTargetOk returns a tuple with the DefaultTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTarget

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) SetDefaultTarget(v bool)`

SetDefaultTarget sets DefaultTarget field to given value.

### HasDefaultTarget

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) HasDefaultTarget() bool`

HasDefaultTarget returns a boolean if a field has been set.

### GetCanManage

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) GetCanManage() bool`

GetCanManage returns the CanManage field if non-nil, zero value otherwise.

### GetCanManageOk

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) GetCanManageOk() (*bool, bool)`

GetCanManageOk returns a tuple with the CanManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManage

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) SetCanManage(v bool)`

SetCanManage sets CanManage field to given value.

### HasCanManage

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) HasCanManage() bool`

HasCanManage returns a boolean if a field has been set.

### GetAll

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetAccount

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) GetAccount() ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) GetAccountOk() (*ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) SetAccount(v ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSites

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) GetSites() []ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) GetSitesOk() (*[]ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) SetSites(v []ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) HasSites() bool`

HasSites returns a boolean if a field has been set.

### SetSitesNil

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) SetSitesNil(b bool)`

 SetSitesNil sets the value for Sites to be an explicit nil

### UnsetSites
`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) UnsetSites()`

UnsetSites ensures that no value is present for Sites, not even an explicit nil
### GetAllPlans

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetPlans

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) GetPlans() []ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsPlansInner`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) GetPlansOk() (*[]ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsPlansInner, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) SetPlans(v []ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsPlansInner)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### SetPlansNil

`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) SetPlansNil(b bool)`

 SetPlansNil sets the value for Plans to be an explicit nil

### UnsetPlans
`func (o *GetCloudResourcePools200ResponseAllOfResourcePoolResourcePermission) UnsetPlans()`

UnsetPlans ensures that no value is present for Plans, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


