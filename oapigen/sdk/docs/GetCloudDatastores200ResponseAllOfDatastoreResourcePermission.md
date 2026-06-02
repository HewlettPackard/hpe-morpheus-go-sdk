# GetCloudDatastores200ResponseAllOfDatastoreResourcePermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultStore** | Pointer to **bool** |  | [optional] 
**DefaultTarget** | Pointer to **bool** |  | [optional] 
**CanManage** | Pointer to **bool** |  | [optional] 
**All** | Pointer to **bool** |  | [optional] 
**Account** | Pointer to [**GetCloudDatastores200ResponseAllOfDatastoreResourcePermissionAccount**](GetCloudDatastores200ResponseAllOfDatastoreResourcePermissionAccount.md) |  | [optional] 
**Sites** | Pointer to [**[]GetCloudDatastores200ResponseAllOfDatastoreResourcePermissionSitesInner**](GetCloudDatastores200ResponseAllOfDatastoreResourcePermissionSitesInner.md) |  | [optional] 
**AllPlans** | Pointer to **bool** |  | [optional] 
**Plans** | Pointer to [**[]GetCloudDatastores200ResponseAllOfDatastoreResourcePermissionPlansInner**](GetCloudDatastores200ResponseAllOfDatastoreResourcePermissionPlansInner.md) |  | [optional] 

## Methods

### NewGetCloudDatastores200ResponseAllOfDatastoreResourcePermission

`func NewGetCloudDatastores200ResponseAllOfDatastoreResourcePermission() *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission`

NewGetCloudDatastores200ResponseAllOfDatastoreResourcePermission instantiates a new GetCloudDatastores200ResponseAllOfDatastoreResourcePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetDefaultStore

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetDefaultTarget

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) GetDefaultTarget() bool`

GetDefaultTarget returns the DefaultTarget field if non-nil, zero value otherwise.

### GetDefaultTargetOk

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) GetDefaultTargetOk() (*bool, bool)`

GetDefaultTargetOk returns a tuple with the DefaultTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTarget

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) SetDefaultTarget(v bool)`

SetDefaultTarget sets DefaultTarget field to given value.

### HasDefaultTarget

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) HasDefaultTarget() bool`

HasDefaultTarget returns a boolean if a field has been set.

### GetCanManage

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) GetCanManage() bool`

GetCanManage returns the CanManage field if non-nil, zero value otherwise.

### GetCanManageOk

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) GetCanManageOk() (*bool, bool)`

GetCanManageOk returns a tuple with the CanManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManage

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) SetCanManage(v bool)`

SetCanManage sets CanManage field to given value.

### HasCanManage

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) HasCanManage() bool`

HasCanManage returns a boolean if a field has been set.

### GetAll

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetAccount

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) GetAccount() GetCloudDatastores200ResponseAllOfDatastoreResourcePermissionAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) GetAccountOk() (*GetCloudDatastores200ResponseAllOfDatastoreResourcePermissionAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) SetAccount(v GetCloudDatastores200ResponseAllOfDatastoreResourcePermissionAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSites

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) GetSites() []GetCloudDatastores200ResponseAllOfDatastoreResourcePermissionSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) GetSitesOk() (*[]GetCloudDatastores200ResponseAllOfDatastoreResourcePermissionSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) SetSites(v []GetCloudDatastores200ResponseAllOfDatastoreResourcePermissionSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) HasSites() bool`

HasSites returns a boolean if a field has been set.

### SetSitesNil

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) SetSitesNil(b bool)`

 SetSitesNil sets the value for Sites to be an explicit nil

### UnsetSites
`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) UnsetSites()`

UnsetSites ensures that no value is present for Sites, not even an explicit nil
### GetAllPlans

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetPlans

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) GetPlans() []GetCloudDatastores200ResponseAllOfDatastoreResourcePermissionPlansInner`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) GetPlansOk() (*[]GetCloudDatastores200ResponseAllOfDatastoreResourcePermissionPlansInner, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) SetPlans(v []GetCloudDatastores200ResponseAllOfDatastoreResourcePermissionPlansInner)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### SetPlansNil

`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) SetPlansNil(b bool)`

 SetPlansNil sets the value for Plans to be an explicit nil

### UnsetPlans
`func (o *GetCloudDatastores200ResponseAllOfDatastoreResourcePermission) UnsetPlans()`

UnsetPlans ensures that no value is present for Plans, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


