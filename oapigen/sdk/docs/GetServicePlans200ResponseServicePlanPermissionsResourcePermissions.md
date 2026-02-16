# GetServicePlans200ResponseServicePlanPermissionsResourcePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultStore** | Pointer to **bool** |  | [optional] 
**AllPlans** | Pointer to **bool** |  | [optional] 
**DefaultTarget** | Pointer to **bool** |  | [optional] 
**CanManage** | Pointer to **bool** |  | [optional] 
**All** | Pointer to **bool** |  | [optional] 
**Account** | Pointer to [**GetServicePlans200ResponseServicePlanPermissionsResourcePermissionsAccount**](GetServicePlans200ResponseServicePlanPermissionsResourcePermissionsAccount.md) |  | [optional] 
**Sites** | Pointer to [**[]GetServicePlans200ResponseServicePlanPermissionsResourcePermissionsSitesInner**](GetServicePlans200ResponseServicePlanPermissionsResourcePermissionsSitesInner.md) |  | [optional] 
**Plans** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewGetServicePlans200ResponseServicePlanPermissionsResourcePermissions

`func NewGetServicePlans200ResponseServicePlanPermissionsResourcePermissions() *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions`

NewGetServicePlans200ResponseServicePlanPermissionsResourcePermissions instantiates a new GetServicePlans200ResponseServicePlanPermissionsResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServicePlans200ResponseServicePlanPermissionsResourcePermissionsWithDefaults

`func NewGetServicePlans200ResponseServicePlanPermissionsResourcePermissionsWithDefaults() *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions`

NewGetServicePlans200ResponseServicePlanPermissionsResourcePermissionsWithDefaults instantiates a new GetServicePlans200ResponseServicePlanPermissionsResourcePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultStore

`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetAllPlans

`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetDefaultTarget

`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) GetDefaultTarget() bool`

GetDefaultTarget returns the DefaultTarget field if non-nil, zero value otherwise.

### GetDefaultTargetOk

`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) GetDefaultTargetOk() (*bool, bool)`

GetDefaultTargetOk returns a tuple with the DefaultTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTarget

`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) SetDefaultTarget(v bool)`

SetDefaultTarget sets DefaultTarget field to given value.

### HasDefaultTarget

`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) HasDefaultTarget() bool`

HasDefaultTarget returns a boolean if a field has been set.

### GetCanManage

`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) GetCanManage() bool`

GetCanManage returns the CanManage field if non-nil, zero value otherwise.

### GetCanManageOk

`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) GetCanManageOk() (*bool, bool)`

GetCanManageOk returns a tuple with the CanManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManage

`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) SetCanManage(v bool)`

SetCanManage sets CanManage field to given value.

### HasCanManage

`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) HasCanManage() bool`

HasCanManage returns a boolean if a field has been set.

### GetAll

`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetAccount

`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) GetAccount() GetServicePlans200ResponseServicePlanPermissionsResourcePermissionsAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) GetAccountOk() (*GetServicePlans200ResponseServicePlanPermissionsResourcePermissionsAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) SetAccount(v GetServicePlans200ResponseServicePlanPermissionsResourcePermissionsAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSites

`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) GetSites() []GetServicePlans200ResponseServicePlanPermissionsResourcePermissionsSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) GetSitesOk() (*[]GetServicePlans200ResponseServicePlanPermissionsResourcePermissionsSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) SetSites(v []GetServicePlans200ResponseServicePlanPermissionsResourcePermissionsSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetPlans

`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) GetPlans() []map[string]interface{}`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) GetPlansOk() (*[]map[string]interface{}, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) SetPlans(v []map[string]interface{})`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### SetPlansNil

`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) SetPlansNil(b bool)`

 SetPlansNil sets the value for Plans to be an explicit nil

### UnsetPlans
`func (o *GetServicePlans200ResponseServicePlanPermissionsResourcePermissions) UnsetPlans()`

UnsetPlans ensures that no value is present for Plans, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


