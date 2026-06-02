# ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission

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

### NewListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission

`func NewListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission() *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission`

NewListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission instantiates a new ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetDefaultStore

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetDefaultTarget

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) GetDefaultTarget() bool`

GetDefaultTarget returns the DefaultTarget field if non-nil, zero value otherwise.

### GetDefaultTargetOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) GetDefaultTargetOk() (*bool, bool)`

GetDefaultTargetOk returns a tuple with the DefaultTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTarget

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) SetDefaultTarget(v bool)`

SetDefaultTarget sets DefaultTarget field to given value.

### HasDefaultTarget

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) HasDefaultTarget() bool`

HasDefaultTarget returns a boolean if a field has been set.

### GetCanManage

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) GetCanManage() bool`

GetCanManage returns the CanManage field if non-nil, zero value otherwise.

### GetCanManageOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) GetCanManageOk() (*bool, bool)`

GetCanManageOk returns a tuple with the CanManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManage

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) SetCanManage(v bool)`

SetCanManage sets CanManage field to given value.

### HasCanManage

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) HasCanManage() bool`

HasCanManage returns a boolean if a field has been set.

### GetAll

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetAccount

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) GetAccount() ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) GetAccountOk() (*ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) SetAccount(v ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSites

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) GetSites() []ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) GetSitesOk() (*[]ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) SetSites(v []ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) HasSites() bool`

HasSites returns a boolean if a field has been set.

### SetSitesNil

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) SetSitesNil(b bool)`

 SetSitesNil sets the value for Sites to be an explicit nil

### UnsetSites
`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) UnsetSites()`

UnsetSites ensures that no value is present for Sites, not even an explicit nil
### GetAllPlans

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetPlans

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) GetPlans() []ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsPlansInner`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) GetPlansOk() (*[]ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsPlansInner, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) SetPlans(v []ListCloudFolders200ResponseAllOfFoldersInnerResourcePermissionsPlansInner)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### SetPlansNil

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) SetPlansNil(b bool)`

 SetPlansNil sets the value for Plans to be an explicit nil

### UnsetPlans
`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission) UnsetPlans()`

UnsetPlans ensures that no value is present for Plans, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


