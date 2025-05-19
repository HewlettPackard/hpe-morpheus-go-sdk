# ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultStore** | Pointer to **bool** |  | [optional] 
**AllPlans** | Pointer to **bool** |  | [optional] 
**DefaultTarget** | Pointer to **bool** |  | [optional] 
**CanManage** | Pointer to **bool** |  | [optional] 
**All** | Pointer to **bool** |  | [optional] 
**Account** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**Sites** | Pointer to [**[]ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermissionSitesInner**](ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermissionSitesInner.md) |  | [optional] 
**Plans** | Pointer to [**[]ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermissionSitesInner**](ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermissionSitesInner.md) |  | [optional] 

## Methods

### NewListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions

`func NewListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions() *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions`

NewListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions instantiates a new ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissionsWithDefaults

`func NewListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissionsWithDefaults() *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions`

NewListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissionsWithDefaults instantiates a new ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultStore

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetAllPlans

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetDefaultTarget

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions) GetDefaultTarget() bool`

GetDefaultTarget returns the DefaultTarget field if non-nil, zero value otherwise.

### GetDefaultTargetOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions) GetDefaultTargetOk() (*bool, bool)`

GetDefaultTargetOk returns a tuple with the DefaultTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTarget

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions) SetDefaultTarget(v bool)`

SetDefaultTarget sets DefaultTarget field to given value.

### HasDefaultTarget

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions) HasDefaultTarget() bool`

HasDefaultTarget returns a boolean if a field has been set.

### GetCanManage

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions) GetCanManage() bool`

GetCanManage returns the CanManage field if non-nil, zero value otherwise.

### GetCanManageOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions) GetCanManageOk() (*bool, bool)`

GetCanManageOk returns a tuple with the CanManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManage

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions) SetCanManage(v bool)`

SetCanManage sets CanManage field to given value.

### HasCanManage

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions) HasCanManage() bool`

HasCanManage returns a boolean if a field has been set.

### GetAll

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetAccount

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions) GetAccount() GetAlerts200ResponseAllOfChecksInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions) GetAccountOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions) SetAccount(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSites

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions) GetSites() []ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermissionSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions) GetSitesOk() (*[]ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermissionSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions) SetSites(v []ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermissionSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetPlans

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions) GetPlans() []ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermissionSitesInner`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions) GetPlansOk() (*[]ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermissionSitesInner, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions) SetPlans(v []ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermissionSitesInner)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *ListClusterDatastores200ResponseAllOfDatastoresInnerResourcePermissions) HasPlans() bool`

HasPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


