# UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission

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

### NewUpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission

`func NewUpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission() *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission`

NewUpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission instantiates a new UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermissionWithDefaults

`func NewUpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermissionWithDefaults() *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission`

NewUpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermissionWithDefaults instantiates a new UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultStore

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetDefaultTarget

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) GetDefaultTarget() bool`

GetDefaultTarget returns the DefaultTarget field if non-nil, zero value otherwise.

### GetDefaultTargetOk

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) GetDefaultTargetOk() (*bool, bool)`

GetDefaultTargetOk returns a tuple with the DefaultTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTarget

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) SetDefaultTarget(v bool)`

SetDefaultTarget sets DefaultTarget field to given value.

### HasDefaultTarget

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) HasDefaultTarget() bool`

HasDefaultTarget returns a boolean if a field has been set.

### GetCanManage

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) GetCanManage() bool`

GetCanManage returns the CanManage field if non-nil, zero value otherwise.

### GetCanManageOk

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) GetCanManageOk() (*bool, bool)`

GetCanManageOk returns a tuple with the CanManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManage

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) SetCanManage(v bool)`

SetCanManage sets CanManage field to given value.

### HasCanManage

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) HasCanManage() bool`

HasCanManage returns a boolean if a field has been set.

### GetAll

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetAccount

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) GetAccount() ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) GetAccountOk() (*ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) SetAccount(v ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSites

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) GetSites() []ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) GetSitesOk() (*[]ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) SetSites(v []ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) HasSites() bool`

HasSites returns a boolean if a field has been set.

### SetSitesNil

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) SetSitesNil(b bool)`

 SetSitesNil sets the value for Sites to be an explicit nil

### UnsetSites
`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) UnsetSites()`

UnsetSites ensures that no value is present for Sites, not even an explicit nil
### GetAllPlans

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetPlans

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) GetPlans() []ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsPlansInner`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) GetPlansOk() (*[]ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsPlansInner, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) SetPlans(v []ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsPlansInner)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### SetPlansNil

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) SetPlansNil(b bool)`

 SetPlansNil sets the value for Plans to be an explicit nil

### UnsetPlans
`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission) UnsetPlans()`

UnsetPlans ensures that no value is present for Plans, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


