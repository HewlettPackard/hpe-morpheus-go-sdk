# CreateLoadBalancer200ResponseLoadBalancerResourcePermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultStore** | Pointer to **bool** |  | [optional] 
**DefaultTarget** | Pointer to **bool** |  | [optional] 
**CanManage** | Pointer to **bool** |  | [optional] 
**All** | Pointer to **bool** |  | [optional] 
**Account** | Pointer to [**CreateLoadBalancer200ResponseLoadBalancerResourcePermissionAccount**](CreateLoadBalancer200ResponseLoadBalancerResourcePermissionAccount.md) |  | [optional] 
**Sites** | Pointer to [**[]CreateLoadBalancer200ResponseLoadBalancerResourcePermissionSitesInner**](CreateLoadBalancer200ResponseLoadBalancerResourcePermissionSitesInner.md) |  | [optional] 
**AllPlans** | Pointer to **bool** |  | [optional] 
**Plans** | Pointer to [**[]CreateLoadBalancer200ResponseLoadBalancerResourcePermissionPlansInner**](CreateLoadBalancer200ResponseLoadBalancerResourcePermissionPlansInner.md) |  | [optional] 

## Methods

### NewCreateLoadBalancer200ResponseLoadBalancerResourcePermission

`func NewCreateLoadBalancer200ResponseLoadBalancerResourcePermission() *CreateLoadBalancer200ResponseLoadBalancerResourcePermission`

NewCreateLoadBalancer200ResponseLoadBalancerResourcePermission instantiates a new CreateLoadBalancer200ResponseLoadBalancerResourcePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoadBalancer200ResponseLoadBalancerResourcePermissionWithDefaults

`func NewCreateLoadBalancer200ResponseLoadBalancerResourcePermissionWithDefaults() *CreateLoadBalancer200ResponseLoadBalancerResourcePermission`

NewCreateLoadBalancer200ResponseLoadBalancerResourcePermissionWithDefaults instantiates a new CreateLoadBalancer200ResponseLoadBalancerResourcePermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultStore

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetDefaultTarget

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) GetDefaultTarget() bool`

GetDefaultTarget returns the DefaultTarget field if non-nil, zero value otherwise.

### GetDefaultTargetOk

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) GetDefaultTargetOk() (*bool, bool)`

GetDefaultTargetOk returns a tuple with the DefaultTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTarget

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) SetDefaultTarget(v bool)`

SetDefaultTarget sets DefaultTarget field to given value.

### HasDefaultTarget

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) HasDefaultTarget() bool`

HasDefaultTarget returns a boolean if a field has been set.

### GetCanManage

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) GetCanManage() bool`

GetCanManage returns the CanManage field if non-nil, zero value otherwise.

### GetCanManageOk

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) GetCanManageOk() (*bool, bool)`

GetCanManageOk returns a tuple with the CanManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManage

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) SetCanManage(v bool)`

SetCanManage sets CanManage field to given value.

### HasCanManage

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) HasCanManage() bool`

HasCanManage returns a boolean if a field has been set.

### GetAll

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetAccount

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) GetAccount() CreateLoadBalancer200ResponseLoadBalancerResourcePermissionAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) GetAccountOk() (*CreateLoadBalancer200ResponseLoadBalancerResourcePermissionAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) SetAccount(v CreateLoadBalancer200ResponseLoadBalancerResourcePermissionAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSites

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) GetSites() []CreateLoadBalancer200ResponseLoadBalancerResourcePermissionSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) GetSitesOk() (*[]CreateLoadBalancer200ResponseLoadBalancerResourcePermissionSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) SetSites(v []CreateLoadBalancer200ResponseLoadBalancerResourcePermissionSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) HasSites() bool`

HasSites returns a boolean if a field has been set.

### SetSitesNil

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) SetSitesNil(b bool)`

 SetSitesNil sets the value for Sites to be an explicit nil

### UnsetSites
`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) UnsetSites()`

UnsetSites ensures that no value is present for Sites, not even an explicit nil
### GetAllPlans

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetPlans

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) GetPlans() []CreateLoadBalancer200ResponseLoadBalancerResourcePermissionPlansInner`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) GetPlansOk() (*[]CreateLoadBalancer200ResponseLoadBalancerResourcePermissionPlansInner, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) SetPlans(v []CreateLoadBalancer200ResponseLoadBalancerResourcePermissionPlansInner)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### SetPlansNil

`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) SetPlansNil(b bool)`

 SetPlansNil sets the value for Plans to be an explicit nil

### UnsetPlans
`func (o *CreateLoadBalancer200ResponseLoadBalancerResourcePermission) UnsetPlans()`

UnsetPlans ensures that no value is present for Plans, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


