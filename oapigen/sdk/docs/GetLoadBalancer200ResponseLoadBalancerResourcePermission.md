# GetLoadBalancer200ResponseLoadBalancerResourcePermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultStore** | Pointer to **bool** |  | [optional] 
**DefaultTarget** | Pointer to **bool** |  | [optional] 
**CanManage** | Pointer to **bool** |  | [optional] 
**All** | Pointer to **bool** |  | [optional] 
**Account** | Pointer to [**GetLoadBalancer200ResponseLoadBalancerResourcePermissionAccount**](GetLoadBalancer200ResponseLoadBalancerResourcePermissionAccount.md) |  | [optional] 
**Sites** | Pointer to [**[]GetLoadBalancer200ResponseLoadBalancerResourcePermissionSitesInner**](GetLoadBalancer200ResponseLoadBalancerResourcePermissionSitesInner.md) |  | [optional] 
**AllPlans** | Pointer to **bool** |  | [optional] 
**Plans** | Pointer to [**[]GetLoadBalancer200ResponseLoadBalancerResourcePermissionPlansInner**](GetLoadBalancer200ResponseLoadBalancerResourcePermissionPlansInner.md) |  | [optional] 

## Methods

### NewGetLoadBalancer200ResponseLoadBalancerResourcePermission

`func NewGetLoadBalancer200ResponseLoadBalancerResourcePermission() *GetLoadBalancer200ResponseLoadBalancerResourcePermission`

NewGetLoadBalancer200ResponseLoadBalancerResourcePermission instantiates a new GetLoadBalancer200ResponseLoadBalancerResourcePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetDefaultStore

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetDefaultTarget

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) GetDefaultTarget() bool`

GetDefaultTarget returns the DefaultTarget field if non-nil, zero value otherwise.

### GetDefaultTargetOk

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) GetDefaultTargetOk() (*bool, bool)`

GetDefaultTargetOk returns a tuple with the DefaultTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTarget

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) SetDefaultTarget(v bool)`

SetDefaultTarget sets DefaultTarget field to given value.

### HasDefaultTarget

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) HasDefaultTarget() bool`

HasDefaultTarget returns a boolean if a field has been set.

### GetCanManage

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) GetCanManage() bool`

GetCanManage returns the CanManage field if non-nil, zero value otherwise.

### GetCanManageOk

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) GetCanManageOk() (*bool, bool)`

GetCanManageOk returns a tuple with the CanManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManage

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) SetCanManage(v bool)`

SetCanManage sets CanManage field to given value.

### HasCanManage

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) HasCanManage() bool`

HasCanManage returns a boolean if a field has been set.

### GetAll

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetAccount

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) GetAccount() GetLoadBalancer200ResponseLoadBalancerResourcePermissionAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) GetAccountOk() (*GetLoadBalancer200ResponseLoadBalancerResourcePermissionAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) SetAccount(v GetLoadBalancer200ResponseLoadBalancerResourcePermissionAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSites

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) GetSites() []GetLoadBalancer200ResponseLoadBalancerResourcePermissionSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) GetSitesOk() (*[]GetLoadBalancer200ResponseLoadBalancerResourcePermissionSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) SetSites(v []GetLoadBalancer200ResponseLoadBalancerResourcePermissionSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) HasSites() bool`

HasSites returns a boolean if a field has been set.

### SetSitesNil

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) SetSitesNil(b bool)`

 SetSitesNil sets the value for Sites to be an explicit nil

### UnsetSites
`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) UnsetSites()`

UnsetSites ensures that no value is present for Sites, not even an explicit nil
### GetAllPlans

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetPlans

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) GetPlans() []GetLoadBalancer200ResponseLoadBalancerResourcePermissionPlansInner`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) GetPlansOk() (*[]GetLoadBalancer200ResponseLoadBalancerResourcePermissionPlansInner, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) SetPlans(v []GetLoadBalancer200ResponseLoadBalancerResourcePermissionPlansInner)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### SetPlansNil

`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) SetPlansNil(b bool)`

 SetPlansNil sets the value for Plans to be an explicit nil

### UnsetPlans
`func (o *GetLoadBalancer200ResponseLoadBalancerResourcePermission) UnsetPlans()`

UnsetPlans ensures that no value is present for Plans, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


