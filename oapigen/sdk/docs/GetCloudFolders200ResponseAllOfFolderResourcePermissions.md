# GetCloudFolders200ResponseAllOfFolderResourcePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultStore** | Pointer to **bool** |  | [optional] 
**DefaultTarget** | Pointer to **bool** |  | [optional] 
**CanManage** | Pointer to **bool** |  | [optional] 
**All** | Pointer to **bool** |  | [optional] 
**Account** | Pointer to [**GetCloudFolders200ResponseAllOfFolderResourcePermissionsAccount**](GetCloudFolders200ResponseAllOfFolderResourcePermissionsAccount.md) |  | [optional] 
**Sites** | Pointer to [**[]GetCloudFolders200ResponseAllOfFolderResourcePermissionsSitesInner**](GetCloudFolders200ResponseAllOfFolderResourcePermissionsSitesInner.md) |  | [optional] 
**AllPlans** | Pointer to **bool** |  | [optional] 
**Plans** | Pointer to [**[]GetCloudFolders200ResponseAllOfFolderResourcePermissionsPlansInner**](GetCloudFolders200ResponseAllOfFolderResourcePermissionsPlansInner.md) |  | [optional] 

## Methods

### NewGetCloudFolders200ResponseAllOfFolderResourcePermissions

`func NewGetCloudFolders200ResponseAllOfFolderResourcePermissions() *GetCloudFolders200ResponseAllOfFolderResourcePermissions`

NewGetCloudFolders200ResponseAllOfFolderResourcePermissions instantiates a new GetCloudFolders200ResponseAllOfFolderResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCloudFolders200ResponseAllOfFolderResourcePermissionsWithDefaults

`func NewGetCloudFolders200ResponseAllOfFolderResourcePermissionsWithDefaults() *GetCloudFolders200ResponseAllOfFolderResourcePermissions`

NewGetCloudFolders200ResponseAllOfFolderResourcePermissionsWithDefaults instantiates a new GetCloudFolders200ResponseAllOfFolderResourcePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultStore

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetDefaultTarget

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) GetDefaultTarget() bool`

GetDefaultTarget returns the DefaultTarget field if non-nil, zero value otherwise.

### GetDefaultTargetOk

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) GetDefaultTargetOk() (*bool, bool)`

GetDefaultTargetOk returns a tuple with the DefaultTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTarget

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) SetDefaultTarget(v bool)`

SetDefaultTarget sets DefaultTarget field to given value.

### HasDefaultTarget

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) HasDefaultTarget() bool`

HasDefaultTarget returns a boolean if a field has been set.

### GetCanManage

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) GetCanManage() bool`

GetCanManage returns the CanManage field if non-nil, zero value otherwise.

### GetCanManageOk

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) GetCanManageOk() (*bool, bool)`

GetCanManageOk returns a tuple with the CanManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManage

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) SetCanManage(v bool)`

SetCanManage sets CanManage field to given value.

### HasCanManage

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) HasCanManage() bool`

HasCanManage returns a boolean if a field has been set.

### GetAll

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetAccount

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) GetAccount() GetCloudFolders200ResponseAllOfFolderResourcePermissionsAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) GetAccountOk() (*GetCloudFolders200ResponseAllOfFolderResourcePermissionsAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) SetAccount(v GetCloudFolders200ResponseAllOfFolderResourcePermissionsAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSites

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) GetSites() []GetCloudFolders200ResponseAllOfFolderResourcePermissionsSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) GetSitesOk() (*[]GetCloudFolders200ResponseAllOfFolderResourcePermissionsSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) SetSites(v []GetCloudFolders200ResponseAllOfFolderResourcePermissionsSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) HasSites() bool`

HasSites returns a boolean if a field has been set.

### SetSitesNil

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) SetSitesNil(b bool)`

 SetSitesNil sets the value for Sites to be an explicit nil

### UnsetSites
`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) UnsetSites()`

UnsetSites ensures that no value is present for Sites, not even an explicit nil
### GetAllPlans

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetPlans

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) GetPlans() []GetCloudFolders200ResponseAllOfFolderResourcePermissionsPlansInner`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) GetPlansOk() (*[]GetCloudFolders200ResponseAllOfFolderResourcePermissionsPlansInner, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) SetPlans(v []GetCloudFolders200ResponseAllOfFolderResourcePermissionsPlansInner)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### SetPlansNil

`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) SetPlansNil(b bool)`

 SetPlansNil sets the value for Plans to be an explicit nil

### UnsetPlans
`func (o *GetCloudFolders200ResponseAllOfFolderResourcePermissions) UnsetPlans()`

UnsetPlans ensures that no value is present for Plans, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


