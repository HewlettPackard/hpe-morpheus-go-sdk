# AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions

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
**Plans** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewAddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions

`func NewAddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions() *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions`

NewAddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions instantiates a new AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissionsWithDefaults

`func NewAddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissionsWithDefaults() *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions`

NewAddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissionsWithDefaults instantiates a new AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultStore

`func (o *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetAllPlans

`func (o *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetDefaultTarget

`func (o *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions) GetDefaultTarget() bool`

GetDefaultTarget returns the DefaultTarget field if non-nil, zero value otherwise.

### GetDefaultTargetOk

`func (o *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions) GetDefaultTargetOk() (*bool, bool)`

GetDefaultTargetOk returns a tuple with the DefaultTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTarget

`func (o *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions) SetDefaultTarget(v bool)`

SetDefaultTarget sets DefaultTarget field to given value.

### HasDefaultTarget

`func (o *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions) HasDefaultTarget() bool`

HasDefaultTarget returns a boolean if a field has been set.

### GetCanManage

`func (o *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions) GetCanManage() bool`

GetCanManage returns the CanManage field if non-nil, zero value otherwise.

### GetCanManageOk

`func (o *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions) GetCanManageOk() (*bool, bool)`

GetCanManageOk returns a tuple with the CanManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManage

`func (o *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions) SetCanManage(v bool)`

SetCanManage sets CanManage field to given value.

### HasCanManage

`func (o *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions) HasCanManage() bool`

HasCanManage returns a boolean if a field has been set.

### GetAll

`func (o *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetAccount

`func (o *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions) GetAccount() GetAlerts200ResponseAllOfChecksInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions) GetAccountOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions) SetAccount(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSites

`func (o *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions) GetSites() []ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermissionSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions) GetSitesOk() (*[]ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermissionSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions) SetSites(v []ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermissionSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetPlans

`func (o *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions) GetPlans() []map[string]interface{}`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions) GetPlansOk() (*[]map[string]interface{}, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions) SetPlans(v []map[string]interface{})`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *AddServicePlans200ResponseAllOfServicePlanPermissionsResourcePermissions) HasPlans() bool`

HasPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


