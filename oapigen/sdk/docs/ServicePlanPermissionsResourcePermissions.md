# ServicePlanPermissionsResourcePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultStore** | Pointer to **bool** |  | [optional] 
**AllPlans** | Pointer to **bool** |  | [optional] 
**DefaultTarget** | Pointer to **bool** |  | [optional] 
**CanManage** | Pointer to **bool** |  | [optional] 
**All** | Pointer to **bool** |  | [optional] 
**Account** | Pointer to [**ServicePlanPermissionsResourcePermissionsAccount**](ServicePlanPermissionsResourcePermissionsAccount.md) |  | [optional] 
**Sites** | Pointer to [**[]ServicePlanPermissionsResourcePermissionsSitesInner**](ServicePlanPermissionsResourcePermissionsSitesInner.md) |  | [optional] 
**Plans** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewServicePlanPermissionsResourcePermissions

`func NewServicePlanPermissionsResourcePermissions() *ServicePlanPermissionsResourcePermissions`

NewServicePlanPermissionsResourcePermissions instantiates a new ServicePlanPermissionsResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePlanPermissionsResourcePermissionsWithDefaults

`func NewServicePlanPermissionsResourcePermissionsWithDefaults() *ServicePlanPermissionsResourcePermissions`

NewServicePlanPermissionsResourcePermissionsWithDefaults instantiates a new ServicePlanPermissionsResourcePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultStore

`func (o *ServicePlanPermissionsResourcePermissions) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *ServicePlanPermissionsResourcePermissions) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *ServicePlanPermissionsResourcePermissions) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *ServicePlanPermissionsResourcePermissions) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetAllPlans

`func (o *ServicePlanPermissionsResourcePermissions) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *ServicePlanPermissionsResourcePermissions) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *ServicePlanPermissionsResourcePermissions) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *ServicePlanPermissionsResourcePermissions) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetDefaultTarget

`func (o *ServicePlanPermissionsResourcePermissions) GetDefaultTarget() bool`

GetDefaultTarget returns the DefaultTarget field if non-nil, zero value otherwise.

### GetDefaultTargetOk

`func (o *ServicePlanPermissionsResourcePermissions) GetDefaultTargetOk() (*bool, bool)`

GetDefaultTargetOk returns a tuple with the DefaultTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTarget

`func (o *ServicePlanPermissionsResourcePermissions) SetDefaultTarget(v bool)`

SetDefaultTarget sets DefaultTarget field to given value.

### HasDefaultTarget

`func (o *ServicePlanPermissionsResourcePermissions) HasDefaultTarget() bool`

HasDefaultTarget returns a boolean if a field has been set.

### GetCanManage

`func (o *ServicePlanPermissionsResourcePermissions) GetCanManage() bool`

GetCanManage returns the CanManage field if non-nil, zero value otherwise.

### GetCanManageOk

`func (o *ServicePlanPermissionsResourcePermissions) GetCanManageOk() (*bool, bool)`

GetCanManageOk returns a tuple with the CanManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManage

`func (o *ServicePlanPermissionsResourcePermissions) SetCanManage(v bool)`

SetCanManage sets CanManage field to given value.

### HasCanManage

`func (o *ServicePlanPermissionsResourcePermissions) HasCanManage() bool`

HasCanManage returns a boolean if a field has been set.

### GetAll

`func (o *ServicePlanPermissionsResourcePermissions) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *ServicePlanPermissionsResourcePermissions) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *ServicePlanPermissionsResourcePermissions) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *ServicePlanPermissionsResourcePermissions) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetAccount

`func (o *ServicePlanPermissionsResourcePermissions) GetAccount() ServicePlanPermissionsResourcePermissionsAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ServicePlanPermissionsResourcePermissions) GetAccountOk() (*ServicePlanPermissionsResourcePermissionsAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ServicePlanPermissionsResourcePermissions) SetAccount(v ServicePlanPermissionsResourcePermissionsAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ServicePlanPermissionsResourcePermissions) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSites

`func (o *ServicePlanPermissionsResourcePermissions) GetSites() []ServicePlanPermissionsResourcePermissionsSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ServicePlanPermissionsResourcePermissions) GetSitesOk() (*[]ServicePlanPermissionsResourcePermissionsSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ServicePlanPermissionsResourcePermissions) SetSites(v []ServicePlanPermissionsResourcePermissionsSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *ServicePlanPermissionsResourcePermissions) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetPlans

`func (o *ServicePlanPermissionsResourcePermissions) GetPlans() []map[string]interface{}`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *ServicePlanPermissionsResourcePermissions) GetPlansOk() (*[]map[string]interface{}, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *ServicePlanPermissionsResourcePermissions) SetPlans(v []map[string]interface{})`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *ServicePlanPermissionsResourcePermissions) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### SetPlansNil

`func (o *ServicePlanPermissionsResourcePermissions) SetPlansNil(b bool)`

 SetPlansNil sets the value for Plans to be an explicit nil

### UnsetPlans
`func (o *ServicePlanPermissionsResourcePermissions) UnsetPlans()`

UnsetPlans ensures that no value is present for Plans, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


