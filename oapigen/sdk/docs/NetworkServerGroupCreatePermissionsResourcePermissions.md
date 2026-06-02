# NetworkServerGroupCreatePermissionsResourcePermissions

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

### NewNetworkServerGroupCreatePermissionsResourcePermissions

`func NewNetworkServerGroupCreatePermissionsResourcePermissions() *NetworkServerGroupCreatePermissionsResourcePermissions`

NewNetworkServerGroupCreatePermissionsResourcePermissions instantiates a new NetworkServerGroupCreatePermissionsResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetDefaultStore

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetDefaultTarget

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) GetDefaultTarget() bool`

GetDefaultTarget returns the DefaultTarget field if non-nil, zero value otherwise.

### GetDefaultTargetOk

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) GetDefaultTargetOk() (*bool, bool)`

GetDefaultTargetOk returns a tuple with the DefaultTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTarget

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) SetDefaultTarget(v bool)`

SetDefaultTarget sets DefaultTarget field to given value.

### HasDefaultTarget

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) HasDefaultTarget() bool`

HasDefaultTarget returns a boolean if a field has been set.

### GetCanManage

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) GetCanManage() bool`

GetCanManage returns the CanManage field if non-nil, zero value otherwise.

### GetCanManageOk

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) GetCanManageOk() (*bool, bool)`

GetCanManageOk returns a tuple with the CanManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManage

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) SetCanManage(v bool)`

SetCanManage sets CanManage field to given value.

### HasCanManage

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) HasCanManage() bool`

HasCanManage returns a boolean if a field has been set.

### GetAll

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetAccount

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) GetAccount() ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) GetAccountOk() (*ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) SetAccount(v ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSites

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) GetSites() []ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) GetSitesOk() (*[]ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) SetSites(v []ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) HasSites() bool`

HasSites returns a boolean if a field has been set.

### SetSitesNil

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) SetSitesNil(b bool)`

 SetSitesNil sets the value for Sites to be an explicit nil

### UnsetSites
`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) UnsetSites()`

UnsetSites ensures that no value is present for Sites, not even an explicit nil
### GetAllPlans

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetPlans

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) GetPlans() []ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsPlansInner`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) GetPlansOk() (*[]ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsPlansInner, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) SetPlans(v []ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsPlansInner)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### SetPlansNil

`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) SetPlansNil(b bool)`

 SetPlansNil sets the value for Plans to be an explicit nil

### UnsetPlans
`func (o *NetworkServerGroupCreatePermissionsResourcePermissions) UnsetPlans()`

UnsetPlans ensures that no value is present for Plans, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


